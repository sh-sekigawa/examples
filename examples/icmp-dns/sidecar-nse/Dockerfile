FROM golang:alpine as build
RUN apk --no-cache add git
ENV PACKAGEPATH=github.com/sh-sekigawa/networkservicemesh/
ENV GO111MODULE=on

RUN mkdir /root/networkservicemesh
ADD ["examples/icmp-dns/sidecar-nse/go.mod","/root/networkservicemesh"]
ADD ["examples/icmp-dns/sidecar-nse/cmd/dnspopulate/","/root/networkservicemesh/examples/icmp-dns/sidecar-nse/cmd/dnspopulate"]
ADD ["examples/icmp-dns/sidecar-nse/cmd/flags/","/root/networkservicemesh/examples/icmp-dns/sidecar-nse/cmd/flags"]
ADD ["examples/icmp-dns/sidecar-nse/cmd/dnsendpoint/","/root/networkservicemesh/examples/icmp-dns/sidecar-nse/cmd/dnsendpoint"]
WORKDIR /root/networkservicemesh/
RUN go mod download

ADD ["examples/icmp-dns/sidecar-nse","/root/networkservicemesh/examples/icmp-dns/sidecar-nse"]

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-extldflags "-static"' -o /go/bin/sidecar-nse ./examples/icmp-dns/sidecar-nse/cmd/main.go ./examples/icmp-dns/sidecar-nse/cmd/env.go
FROM alpine as runtime
COPY --from=build /go/bin/sidecar-nse /bin/sidecar-nse
ENTRYPOINT ["/bin/sidecar-nse", "-dns"]
