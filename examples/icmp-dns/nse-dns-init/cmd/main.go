// Copyright (c) 2020 Doc.ai, Inc and/or its affiliates.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"os"
	"github.com/sirupsen/logrus"

	"github.com/networkservicemesh/networkservicemesh/utils/dnsconfig"

	"github.com/networkservicemesh/networkservicemesh/controlplane/api/connectioncontext"
	"github.com/networkservicemesh/networkservicemesh/utils/caddyfile"
)

func main() {
	externalDnsServerIp := os.Getenv("EXTERNAL_DNS_IP")
	externalDnsSearchDomain := os.Getenv("DNS_SEARCH_DOMAIN")
	r := resolvConfFile{path: resolvConfFilePath}
	defaultDNSConfig := []*connectioncontext.DNSConfig{
		{
			// specific zones config
			DnsServerIps:  r.Nameservers(),
			SearchDomains: r.Searches(),
		},
		{
			//any zone config
			DnsServerIps: r.Nameservers(),
		},
	}
	if externalDnsServerIp != "" && externalDnsSearchDomain != "" {
		defaultDNSConfig = append(defaultDNSConfig,
			&connectioncontext.DNSConfig{
				DnsServerIps:  []string{externalDnsServerIp},
				SearchDomains: []string{externalDnsSearchDomain},
			},
		)
	}

	properties := []resolvConfProperty{
		{nameserverProperty, []string{"127.0.0.1"}},
		{searchProperty, r.Searches()},
		{optionsProperty, r.Options()},
	}
	r.ReplaceProperties(properties)
	m := dnsconfig.NewManager(defaultDNSConfig...)
	f := m.Caddyfile(caddyfile.Path())
	err := f.Save()
	if err != nil {
		logrus.Fatalf("An error during save caddy file %v", err)
	}
}
