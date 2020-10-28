NAME = 5g-network
DESCRIPTION = "free5GC"
CONTAINERS = 
AUX_CONTAINERS = icmp-dns-sidecar-nse icmp-dns-nse-dns-init
PODS = free5gc-config amf ausf db nrf nssf pcf ran smf udm udr upf webui
NETWORK_SERVICES = 5g-network
CHECK = 

# should not fail go linter
FAIL_GOLINT =

include $(TOP)/mk/targets.mk
