/*
Copyright (c) 2023-2024 Dell Inc., or its subsidiaries. All Rights Reserved.

Licensed under the Mozilla Public License Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://mozilla.org/MPL/2.0/


Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package provider

import (
	"fmt"
	"regexp"
	"terraform-provider-powerscale/powerscale/helper"
	"testing"

	"github.com/bytedance/mockey"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGroupnetDataSourceAll(t *testing.T) {
	var groupnetTerraformName = "data.powerscale_groupnet.all"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// read all testing
			{
				Config: ProviderConfig + groupnetAllDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(groupnetTerraformName, "groupnets.#"),
				),
			},
		},
	})
}

func TestAccGroupnetDataSourceFilterNames(t *testing.T) {
	var groupnetTerraformName = "data.powerscale_groupnet.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// filter with all filters read testing
			{
				Config: ProviderConfig + groupnetFilterDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(groupnetTerraformName, "groupnets.#", "1"),
					resource.TestCheckResourceAttr(groupnetTerraformName, "groupnets.0.dns_resolver_rotate", "true"),
					resource.TestCheckResourceAttr(groupnetTerraformName, "groupnets.0.allow_wildcard_subdomains", "false"),
					resource.TestCheckResourceAttr(groupnetTerraformName, "groupnets.0.server_side_dns_search", "true"),
					resource.TestCheckResourceAttr(groupnetTerraformName, "groupnets.0.dns_cache_enabled", "true"),
					resource.TestCheckResourceAttr(groupnetTerraformName, "groupnets.0.name", "tfaccGroupnetDatasourceDep"),
					resource.TestCheckResourceAttr(groupnetTerraformName, "groupnets.0.id", "tfaccGroupnetDatasourceDep"),
					resource.TestCheckResourceAttr(groupnetTerraformName, "groupnets.0.description", "terraform groupnet datasource"),
					resource.TestCheckResourceAttr(groupnetTerraformName, "groupnets.0.dns_search.0", "pie.lab.emc.com"),
					resource.TestCheckResourceAttr(groupnetTerraformName, "groupnets.0.dns_servers.0", "10.230.44.169"),
					resource.TestCheckResourceAttr(groupnetTerraformName, "groupnets.0.subnets.#", "0"),
				),
			},
		},
	})
}

func TestAccGroupnetDataSourceInvalidNames(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// filter with invalid names read testing
			{
				Config:      ProviderConfig + groupnetInvalidFilterDataSourceConfig,
				ExpectError: regexp.MustCompile(`.*Error one or more of the filtered groupnet names is not a valid powerscale groupnet.*.`),
			},
		},
	})
}

func TestAccGroupnetDataSourceFilterDNSCache(t *testing.T) {
	var groupnetTerraformName = "data.powerscale_groupnet.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// filter with DNS cache read testing
			{
				Config: ProviderConfig + groupnetFilterDNSCacheDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(groupnetTerraformName, "groupnets.#"),
					resource.TestCheckResourceAttr(groupnetTerraformName, "groupnets.0.dns_cache_enabled", "false"),
				),
			},
		},
	})
}

func TestAccGroupnetDataSourceFilterAllowWildcardSubdomains(t *testing.T) {
	var groupnetTerraformName = "data.powerscale_groupnet.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// filter with allow_wildcard_subdomains read testing
			{
				Config: ProviderConfig + groupnetFilterAllowWildcardSubdomainsDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(groupnetTerraformName, "groupnets.#"),
					resource.TestCheckResourceAttr(groupnetTerraformName, "groupnets.0.allow_wildcard_subdomains", "false"),
				),
			},
		},
	})
}

func TestAccGroupnetDataSourceFilterDNSResolverRotate(t *testing.T) {
	var groupnetTerraformName = "data.powerscale_groupnet.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// filter with DNS resolver rotate read testing
			{
				Config: ProviderConfig + groupnetFilterDNSResolverRotateDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(groupnetTerraformName, "groupnets.#"),
					resource.TestCheckResourceAttr(groupnetTerraformName, "groupnets.0.dns_resolver_rotate", "true"),
				),
			},
		},
	})
}

func TestAccGroupnetDataSourceFilterServerSideDNSSearch(t *testing.T) {
	var groupnetTerraformName = "data.powerscale_groupnet.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// filter with server_side_dns_search read testing
			{
				Config: ProviderConfig + groupnetFilterServerSideDNSSearchDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(groupnetTerraformName, "groupnets.#"),
					resource.TestCheckResourceAttr(groupnetTerraformName, "groupnets.0.server_side_dns_search", "true"),
				),
			},
		},
	})
}

func TestAccGroupnetDataSourceInvalidFilter(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// filter with invalid filter read testing
			{
				Config:      ProviderConfig + groupnetFilterInvalidDataSourceConfig,
				ExpectError: regexp.MustCompile(`.*is not expected here.*.`),
			},
		},
	})
}

func TestAccGroupnetDatasourceErrorCopyField(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					FunctionMocker = mockey.Mock(helper.CopyFields).Return(fmt.Errorf("mock error")).Build()
				},
				Config:      ProviderConfig + groupnetAllDataSourceConfig,
				ExpectError: regexp.MustCompile("mock error"),
			},
		},
	})
}

func TestAccGroupnetDatasourceErrorGetAll(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					FunctionMocker = mockey.Mock(helper.GetAllGroupnets).Return(nil, fmt.Errorf("mock error")).Build()
				},
				Config:      ProviderConfig + groupnetAllDataSourceConfig,
				ExpectError: regexp.MustCompile("mock error"),
			},
		},
	})
}

func TestAccGroupnetDatasourceReleaseMock(t *testing.T) {
	var groupnetTerraformName = "data.powerscale_groupnet.all"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
				},
				Config: ProviderConfig + groupnetAllDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(groupnetTerraformName, "groupnets.#"),
				),
			},
		},
	})
}

var groupnetFilterDataSourceConfig = `
resource "powerscale_groupnet" "test" {
	name = "tfaccGroupnetDatasourceDep"
	dns_cache_enabled = true
	description = "terraform groupnet datasource"
	allow_wildcard_subdomains = false
	server_side_dns_search = true
	dns_resolver_rotate = true
	dns_search = ["pie.lab.emc.com"]
	dns_servers = ["10.230.44.169"]
  }

data "powerscale_groupnet" "test" {
  filter {
    names = ["tfaccGroupnetDatasourceDep"]
	dns_cache_enabled = true
	allow_wildcard_subdomains = false
	server_side_dns_search = true
	dns_resolver_rotate = true
	dns_search = ["pie.lab.emc.com"]
	dns_servers = ["10.230.44.169"]
  }
  depends_on = [
	powerscale_groupnet.test
  ]
}
`

var groupnetFilterDNSCacheDataSourceConfig = `
resource "powerscale_groupnet" "test" {
	name = "tfaccGroupnetDatasourceDNS"
	dns_cache_enabled = false
	description = "terraform groupnet datasource"
	allow_wildcard_subdomains = false
	server_side_dns_search = true
	dns_resolver_rotate = true
	dns_search = ["pie.lab.emc.com"]
	dns_servers = ["10.230.44.169"]
  }

data "powerscale_groupnet" "test" {
  filter {
	dns_cache_enabled = false
  }
  depends_on = [
	powerscale_groupnet.test
  ]
}
`

var groupnetFilterAllowWildcardSubdomainsDataSourceConfig = `
resource "powerscale_groupnet" "test" {
	name = "tfaccGroupnetDatasourceWildCard"
	dns_cache_enabled = false
	description = "terraform groupnet datasource"
	allow_wildcard_subdomains = false
	server_side_dns_search = true
	dns_resolver_rotate = true
	dns_search = ["pie.lab.emc.com"]
	dns_servers = ["10.230.44.169"]
  }

data "powerscale_groupnet" "test" {
  filter {
	dns_cache_enabled = false
  }
  depends_on = [
	powerscale_groupnet.test
  ]
}
`

var groupnetFilterDNSResolverRotateDataSourceConfig = `
resource "powerscale_groupnet" "test" {
	name = "tfaccGroupnetDatasourceRotate"
	dns_cache_enabled = false
	description = "terraform groupnet datasource"
	allow_wildcard_subdomains = false
	server_side_dns_search = true
	dns_resolver_rotate = true
	dns_search = ["pie.lab.emc.com"]
	dns_servers = ["10.230.44.169"]
  }

data "powerscale_groupnet" "test" {
  filter {
	dns_resolver_rotate = true
  }
}
`

var groupnetFilterServerSideDNSSearchDataSourceConfig = `
resource "powerscale_groupnet" "test" {
	name = "tfaccGroupnetDatasourceDNSSearch"
	dns_cache_enabled = false
	description = "terraform groupnet datasource"
	allow_wildcard_subdomains = false
	server_side_dns_search = true
	dns_resolver_rotate = true
	dns_search = ["pie.lab.emc.com"]
	dns_servers = ["10.230.44.169"]
  }

data "powerscale_groupnet" "test" {
  filter {
	server_side_dns_search = true
  }
}
`

var groupnetAllDataSourceConfig = `
data "powerscale_groupnet" "all" {
}
`
var groupnetInvalidFilterDataSourceConfig = `
data "powerscale_groupnet" "test" {
  filter {
    names = ["invalidName"]
  }
}
`

var groupnetFilterInvalidDataSourceConfig = `
data "powerscale_groupnet" "test" {
  filter {
	badfilter = "invalidFilter"
  }
}
`
