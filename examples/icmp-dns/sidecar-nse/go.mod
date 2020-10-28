module github.com/networkservicemesh/examples

go 1.14

require (
	dnspopulate v0.0.0-00010101000000-000000000000
	github.com/Nordix/simple-ipam v1.0.0
	github.com/davecgh/go-spew v1.1.1
	github.com/fsnotify/fsnotify v1.4.9
	github.com/golang/protobuf v1.4.2
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645
	github.com/networkservicemesh/networkservicemesh/controlplane/api v0.3.0
	github.com/networkservicemesh/networkservicemesh/pkg v0.3.0
	github.com/networkservicemesh/networkservicemesh/sdk v0.3.0
	github.com/networkservicemesh/networkservicemesh/test/applications/cmd/icmp-responder-nse/flags v0.0.0-00010101000000-000000000000
	github.com/networkservicemesh/networkservicemesh/utils v0.3.0
	github.com/opentracing/opentracing-go v1.1.0
	github.com/pkg/errors v0.9.1
	github.com/sh-sekigawa/dnsendpoint/api/types/v1alpha1 v0.0.0-00010101000000-000000000000 // indirect
	github.com/sirupsen/logrus v1.5.0
	github.com/spf13/viper v1.6.2
	go.ligato.io/vpp-agent/v3 v3.1.0
	google.golang.org/grpc v1.28.0
	gopkg.in/yaml.v2 v2.2.8
	k8s.io/apimachinery v0.17.0
	k8s.io/client-go v0.17.0 // indirect
)

replace (
	dnspopulate => ./examples/icmp-dns/sidecar-nse/cmd/dnspopulate
	github.com/census-instrumentation/opencensus-proto v0.1.0-0.20181214143942-ba49f56771b8 => github.com/census-instrumentation/opencensus-proto v0.0.3-0.20181214143942-ba49f56771b8
	github.com/networkservicemesh/networkservicemesh => github.com/networkservicemesh/networkservicemesh v0.0.0-20200328192804-8d64ff42c90d
	github.com/networkservicemesh/networkservicemesh/controlplane/api => github.com/networkservicemesh/networkservicemesh/controlplane/api v0.0.0-20200328192804-8d64ff42c90d
	github.com/networkservicemesh/networkservicemesh/pkg => github.com/networkservicemesh/networkservicemesh/pkg v0.0.0-20200328192804-8d64ff42c90d
	github.com/networkservicemesh/networkservicemesh/sdk => github.com/networkservicemesh/networkservicemesh/sdk v0.0.0-20200328192804-8d64ff42c90d
	github.com/networkservicemesh/networkservicemesh/test/applications/cmd/icmp-responder-nse/flags v0.0.0-00010101000000-000000000000 => ./examples/icmp-dns/sidecar-nse/cmd/flags
	github.com/networkservicemesh/networkservicemesh/utils => github.com/networkservicemesh/networkservicemesh/utils v0.0.0-20200328192804-8d64ff42c90d
	github.com/sh-sekigawa/dnsendpoint/api/types/v1alpha1 => ./examples/icmp-dns/sidecar-nse/cmd/dnsendpoint/api/types/v1alpha1
)
