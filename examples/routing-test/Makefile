NAME = routing-test
DESCRIPTION = "Test for static routing of kernel forwarding"
CONTAINERS =
AUX_CONTAINERS = universal-cnf-vppagent
PODS = simple-client ucnf-client gateway server gateway2
NETWORK_SERVICES = routing-test
CHECK = scripts/check_routing-test.sh

# should not fail go linter
FAIL_GOLINT =

include $(TOP)/mk/targets.mk
