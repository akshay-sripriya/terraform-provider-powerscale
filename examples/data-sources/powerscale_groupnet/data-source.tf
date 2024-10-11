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

# PowerScale Groupnet sits above subnets and pools and allows separate Access Zones to contain distinct DNS settings.

# Returns a list of PowerScale Groupnets based on names filter block. 
data "powerscale_groupnet" "example_groupnet" {
  filter {
    # Optional list of names to filter upon
    names = ["groupnet_name"]
  }
}

data "powerscale_groupnet" "example_groupnet" {
  # Filter by dns_cache_enabled
  filter {
    dns_cache_enabled = false
  }
}

data "powerscale_groupnet" "example_groupnet" {
  # Filter by allow_wildcard_subdomains
  filter {
    allow_wildcard_subdomains = false
    names = ["test"]
  }
}

data "powerscale_groupnet" "example_groupnet" {
  # Filter by dns_resolver_rotate
  filter {
    dns_resolver_rotate = true
  }
}

data "powerscale_groupnet" "example_groupnet" {
  # Filter by all attributes
  filter {
    names = ["groupnet0"]
    dns_cache_enabled = true
    dns_resolver_rotate = false
    dns_search = ["pie.lab.emc.com"]
    dns_servers = ["10.230.44.169"]
    allow_wildcard_subdomains = true
    server_side_dns_search = true
  }
}

# Any combination of filters will return the intersection of the filters. multiple values within the same filter return the union of the values
# For example, 
# name = ["name1", "name2"]
# The above will return the union of name1 and name2

# name = ["name1", "name2"]
# dns_cache_enabled = true
# The above will return the intersection of name and dns_cache_enabled



# Output value of above block by executing 'terraform output' command.
# The user can use the fetched information by the variable data.powerscale_groupnet.example_groupnet
output "powerscale_groupnet_filter" {
  value = data.powerscale_groupnet.example_groupnet
}


# Returns all of the PowerScale Groupnets
data "powerscale_groupnet" "all" {
}

# Output value of above block by executing 'terraform output' command
# The user can use the fetched information by the variable data.powerscale_groupnet.all
output "powerscale_groupnet_all" {
  value = data.powerscale_groupnet.all
}
