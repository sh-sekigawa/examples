FROM golang:alpine as build
RUN apk --no-cache add git
ENV PACKAGEPATH=github.com/sh-sekigawa/networkservicemesh/
ENV GO111MODULE=on

RUN mkdir /root/networkservicemesh
ADD ["examples/icmp-dns/nse-dns-init/go.mod","/root/networkservicemesh"]
WORKDIR /root/networkservicemesh/
RUN go mod download

ADD ["examples/icmp-dns/nse-dns-init","/root/networkservicemesh/examples/icmp-dns/nse-dns-init"]

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-extldflags "-static"' -o /go/bin/nse-dns-init ./examples/icmp-dns/nse-dns-init/cmd/main.go ./examples/icmp-dns/nse-dns-init/cmd/const.go ./examples/icmp-dns/nse-dns-init/cmd/resolv_conf.go
FROM alpine as runtime
COPY --from=build /go/bin/nse-dns-init /bin/nse-dns-init
ENTRYPOINT ["/bin/nse-dns-init"]
