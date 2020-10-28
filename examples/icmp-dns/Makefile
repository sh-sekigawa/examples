NAME = icmp-dns
DESCRIPTION = "Basic kernel interface ICMP reposnder"
CONTAINERS = nse-dns-init sidecar-nse
PODS = simple-client simple-endpoint
NETWORK_SERVICES = icmp-responder
CHECK = scripts/check_icmp.sh

# should not fail go linter
FAIL_GOLINT =

include $(TOP)/mk/targets.mk
