module github.com/sh-seigawa/nsm-dns-init

go 1.14

require (
    github.com/networkservicemesh/networkservicemesh/controlplane/api v0.3.0
	github.com/networkservicemesh/networkservicemesh/pkg v0.3.0
	github.com/networkservicemesh/networkservicemesh/sdk v0.3.0
	github.com/networkservicemesh/networkservicemesh/utils v0.3.0
	github.com/onsi/gomega v1.7.0
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.4.2
	github.com/spiffe/spire/proto/spire v0.0.0-20200103215556-34b7e3785007
	google.golang.org/grpc v1.27.1
)

replace github.com/census-instrumentation/opencensus-proto v0.1.0-0.20181214143942-ba49f56771b8 => github.com/census-instrumentation/opencensus-proto v0.0.3-0.20181214143942-ba49f56771b8

replace (
	github.com/networkservicemesh/networkservicemesh => github.com/networkservicemesh/networkservicemesh v0.0.0-20200328192804-8d64ff42c90d
	github.com/networkservicemesh/networkservicemesh/controlplane/api => github.com/networkservicemesh/networkservicemesh/controlplane/api v0.0.0-20200328192804-8d64ff42c90d
	github.com/networkservicemesh/networkservicemesh/pkg => github.com/networkservicemesh/networkservicemesh/pkg v0.0.0-20200328192804-8d64ff42c90d
	github.com/networkservicemesh/networkservicemesh/sdk => github.com/networkservicemesh/networkservicemesh/sdk v0.0.0-20200328192804-8d64ff42c90d
	github.com/networkservicemesh/networkservicemesh/test/applications/cmd/icmp-responder-nse/flags v0.0.0-00010101000000-000000000000 => ./examples/icmp-dns/sidecar-nse/cmd/flags
	github.com/networkservicemesh/networkservicemesh/utils => github.com/networkservicemesh/networkservicemesh/utils v0.0.0-20200328192804-8d64ff42c90d
	github.com/sh-sekigawa/dnsendpoint/api/types/v1alpha1 => ./examples/icmp-dns/sidecar-nse/cmd/dnsendpoint/api/types/v1alpha1
)
