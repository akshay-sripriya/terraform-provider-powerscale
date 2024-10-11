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
	"context"
	"fmt"
	"strings"
	"terraform-provider-powerscale/client"
	"terraform-provider-powerscale/powerscale/helper"
	"terraform-provider-powerscale/powerscale/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = &GroupnetDataSource{}
	_ datasource.DataSourceWithConfigure = &GroupnetDataSource{}
)

// NewGroupnetDataSource creates a new Groupnet data source.
func NewGroupnetDataSource() datasource.DataSource {
	return &GroupnetDataSource{}
}

// GroupnetDataSource defines the data source implementation.
type GroupnetDataSource struct {
	client *client.Client
}

// Metadata describes the data source arguments.
func (d *GroupnetDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_groupnet"
}

// Schema describes the data source arguments.
func (d *GroupnetDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "This datasource is used to query the existing Groupnets from PowerScale array. The information fetched from this datasource can be used for getting the details or for further processing in resource block. PowerScale Groupnet sits above subnets and pools and allows separate Access Zones to contain distinct DNS settings.",
		Description:         "This datasource is used to query the existing Groupnets from PowerScale array. The information fetched from this datasource can be used for getting the details or for further processing in resource block. PowerScale Groupnet sits above subnets and pools and allows separate Access Zones to contain distinct DNS settings.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Unique identifier of the groupnet instance.",
				Description:         "Unique identifier of the groupnet instance.",
			},
			"groupnets": schema.ListNestedAttribute{
				MarkdownDescription: "List of groupnets.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description:         "The name of the groupnet.",
							MarkdownDescription: "The name of the groupnet.",
							Computed:            true,
						},
						"id": schema.StringAttribute{
							Description:         "Unique Interface ID.",
							MarkdownDescription: "Unique Interface ID.",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							Description:         "A description of the groupnet.",
							MarkdownDescription: "A description of the groupnet.",
							Computed:            true,
						},
						"allow_wildcard_subdomains": schema.BoolAttribute{
							Description:         "If enabled, SmartConnect treats subdomains of known dns zones as the known dns zone. This is required for S3 Virtual Host domains.",
							MarkdownDescription: "If enabled, SmartConnect treats subdomains of known dns zones as the known dns zone. This is required for S3 Virtual Host domains.",
							Computed:            true,
						},
						"dns_cache_enabled": schema.BoolAttribute{
							Description:         "DNS caching is enabled or disabled.",
							MarkdownDescription: "DNS caching is enabled or disabled.",
							Computed:            true,
						},
						"server_side_dns_search": schema.BoolAttribute{
							Description:         "Enable or disable appending nodes DNS search list to client DNS inquiries directed at SmartConnect service IP.",
							MarkdownDescription: "Enable or disable appending nodes DNS search list to client DNS inquiries directed at SmartConnect service IP.",
							Computed:            true,
						},
						"dns_resolver_rotate": schema.BoolAttribute{
							Description:         "Enable or disable DNS resolver rotate.",
							MarkdownDescription: "Enable or disable DNS resolver rotate.",
							Computed:            true,
						},
						"dns_search": schema.ListAttribute{
							Description:         "List of DNS search suffixes.",
							MarkdownDescription: "List of DNS search suffixes.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"dns_servers": schema.ListAttribute{
							Description:         "List of Domain Name Server IP addresses.",
							MarkdownDescription: "List of Domain Name Server IP addresses.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"subnets": schema.ListAttribute{
							Description:         "Name of the subnets in the groupnet.",
							MarkdownDescription: "Name of the subnets in the groupnet.",
							ElementType:         types.StringType,
							Computed:            true,
						},
					},
				},
			},
		},
		Blocks: map[string]schema.Block{
			"filter": schema.SingleNestedBlock{
				Attributes: map[string]schema.Attribute{
					"names": schema.SetAttribute{
						Description:         "Filter groupnets by name.",
						MarkdownDescription: "Filter groupnets by name.",
						Optional:            true,
						ElementType:         types.StringType,
					},

					"dns_cache_enabled": schema.BoolAttribute{
						Description:         "Filter groupnets by DNS cache enabled (true) or disabled (false).",
						MarkdownDescription: "Filter groupnets by DNS cache enabled (true) or disabled (false).",
						Optional:            true,
					},

					"allow_wildcard_subdomains": schema.BoolAttribute{
						Description:         "Filter groupnets by allow wildcard subdomains (true) or disabled (false).",
						MarkdownDescription: "Filter groupnets by allow wildcard subdomains (true) or disabled (false).",
						Optional:            true,
					},

					"dns_resolver_rotate": schema.BoolAttribute{
						Description:         "Filter groupnets by DNS resolver rotate (true) or disabled (false).",
						MarkdownDescription: "Filter groupnets by DNS resolver rotate (true) or disabled (false).",
						Optional:            true,
					},

					"server_side_dns_search": schema.BoolAttribute{
						Description:         "Filter groupnets by server side DNS search (true) or disabled (false).",
						MarkdownDescription: "Filter groupnets by server side DNS search (true) or disabled (false).",
						Optional:            true,
					},

					"dns_search": schema.ListAttribute{
						Description:         "Filter groupnets by DNS search suffixes.",
						MarkdownDescription: "Filter groupnets by DNS search suffixes.",
						Optional:            true,
						ElementType:         types.StringType,
					},

					"dns_servers": schema.ListAttribute{
						Description:         "Filter groupnets by Domain Name Server IP addresses.",
						MarkdownDescription: "Filter groupnets by Domain Name Server IP addresses.",
						Optional:            true,
						ElementType:         types.StringType,
					},
				},
			},
		},
	}
}

// Configure configures the data source.
func (d *GroupnetDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	pscaleClient, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = pscaleClient
}

// Read reads data from the data source.
func (d *GroupnetDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Info(ctx, "Reading Groupnet data source ")

	var state models.GroupnetDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	groupnets, err := helper.GetAllGroupnets(ctx, d.client)
	if err != nil {
		resp.Diagnostics.AddError("Error getting the list of PowerScale Groupnets.", err.Error())
		return
	}

	// parse groupnet response to state groupnet model
	if err := helper.UpdateGroupnetDataSourceState(ctx, &state, groupnets); err != nil {
		resp.Diagnostics.AddError("Error reading groupnets datasource plan",
			fmt.Sprintf("Could not list groupnets with error: %s", err.Error()))
		return
	}

	// filter groupnets by names
	if state.Filter != nil && len(state.Filter.Names) > 0 {
		var validGroupnets []string
		var filteredGroupnets []models.GroupnetModel

		for _, groupnet := range state.Groupnets {
			for _, name := range state.Filter.Names {
				if groupnet.Name.Equal(name) {
					filteredGroupnets = append(filteredGroupnets, groupnet)
					validGroupnets = append(validGroupnets, groupnet.Name.ValueString())
					break
				}
			}
		}

		state.Groupnets = filteredGroupnets

		if len(state.Groupnets) != len(state.Filter.Names) {
			resp.Diagnostics.AddError(
				"Error one or more of the filtered groupnet names is not a valid powerscale groupnet.",
				fmt.Sprintf("Valid groupnets: [%v], filtered list: [%v]", strings.Join(validGroupnets, " , "), state.Filter.Names),
			)
		}
	}

	// filter groupnets by dns_cache_enabled
	if state.Filter != nil && !state.Filter.DNSCache.IsNull() {
		var validGroupnets []string
		var filteredGroupnets []models.GroupnetModel

		for _, groupnet := range state.Groupnets {
			if groupnet.DNSCacheEnabled.Equal(state.Filter.DNSCache) {
				filteredGroupnets = append(filteredGroupnets, groupnet)
				validGroupnets = append(validGroupnets, groupnet.Name.ValueString())
			}
		}

		state.Groupnets = filteredGroupnets

		if len(state.Groupnets) == 0 {
			resp.Diagnostics.AddError(
				"Error no groupnets found with the specified dns_cache_enabled value.",
				fmt.Sprintf("Valid groupnets: [%v], filtered list: [%v]", strings.Join(validGroupnets, " , "), state.Filter.DNSCache),
			)
		}
	}

	// filter groupnets by allow_wildcard_subdomains
	if state.Filter != nil && !state.Filter.AllowWildcardSubdomains.IsNull() {
		var validGroupnets []string
		var filteredGroupnets []models.GroupnetModel

		for _, groupnet := range state.Groupnets {
			if groupnet.AllowWildcardSubdomains.Equal(state.Filter.AllowWildcardSubdomains) {
				filteredGroupnets = append(filteredGroupnets, groupnet)
				validGroupnets = append(validGroupnets, groupnet.Name.ValueString())
			}
		}

		state.Groupnets = filteredGroupnets

		if len(state.Groupnets) == 0 {
			resp.Diagnostics.AddError(
				"Error no groupnets found with the specified allow_wildcard_subdomains value.",
				fmt.Sprintf("Valid groupnets: [%v], filtered list: [%v]", strings.Join(validGroupnets, " , "), state.Filter.AllowWildcardSubdomains),
			)
		}
	}

	// filter groupnets by dns_resolver_rotate
	if state.Filter != nil && !state.Filter.DNSResolverRotate.IsNull() {
		var validGroupnets []string
		var filteredGroupnets []models.GroupnetModel

		for _, groupnet := range state.Groupnets {
			if groupnet.DNSResolverRotate.Equal(state.Filter.DNSResolverRotate) {
				filteredGroupnets = append(filteredGroupnets, groupnet)
				validGroupnets = append(validGroupnets, groupnet.Name.ValueString())
			}
		}

		state.Groupnets = filteredGroupnets

		if len(state.Groupnets) == 0 {
			resp.Diagnostics.AddError(
				"Error no groupnets found with the specified dns_resolver_rotate value.",
				fmt.Sprintf("Valid groupnets: [%v], filtered list: [%v]", strings.Join(validGroupnets, " , "), state.Filter.DNSResolverRotate),
			)
		}
	}

	// filter groupnets by server_side_dns_search
	if state.Filter != nil && !state.Filter.ServerSideDNSSearch.IsNull() {
		var validGroupnets []string
		var filteredGroupnets []models.GroupnetModel

		for _, groupnet := range state.Groupnets {
			if groupnet.ServerSideDNSSearch.Equal(state.Filter.ServerSideDNSSearch) {
				filteredGroupnets = append(filteredGroupnets, groupnet)
				validGroupnets = append(validGroupnets, groupnet.Name.ValueString())
			}
		}

		state.Groupnets = filteredGroupnets

		if len(state.Groupnets) == 0 {
			resp.Diagnostics.AddError(
				"Error no groupnets found with the specified server_side_dns_search value.",
				fmt.Sprintf("Valid groupnets: [%v], filtered list: [%v]", strings.Join(validGroupnets, " , "), state.Filter.ServerSideDNSSearch),
			)
		}
	}

	// filter groupnets by dns_search_suffixes
	if state.Filter != nil && !state.Filter.DNSSearch.IsNull() && len(state.Filter.DNSSearch.Elements()) > 0 {
		var validGroupnets []string
		var filteredGroupnets []models.GroupnetModel

		for _, groupnet := range state.Groupnets {
			for _, search := range state.Filter.DNSSearch.Elements() {
				if contains(groupnet.DNSSearch.Elements(), search) {
					filteredGroupnets = append(filteredGroupnets, groupnet)
					validGroupnets = append(validGroupnets, groupnet.Name.ValueString())
					break
				}
			}
		}

		state.Groupnets = filteredGroupnets

		if len(state.Groupnets) == 0 {
			resp.Diagnostics.AddError(
				"Error one or more of the filtered dns_search values is not a valid powerscale groupnet.",
				fmt.Sprintf("Valid groupnets: [%v], filtered list: [%v]", strings.Join(validGroupnets, " , "), state.Filter.DNSSearch),
			)
		}
	}

	// filter groupnets by dns_servers
	if state.Filter != nil && !state.Filter.DNSServers.IsNull() && len(state.Filter.DNSServers.Elements()) > 0 {
		var validGroupnets []string
		var filteredGroupnets []models.GroupnetModel

		for _, groupnet := range state.Groupnets {
			for _, search := range state.Filter.DNSServers.Elements() {
				if contains(groupnet.DNSServers.Elements(), search) {
					filteredGroupnets = append(filteredGroupnets, groupnet)
					validGroupnets = append(validGroupnets, groupnet.Name.ValueString())
					break
				}
			}
		}

		state.Groupnets = filteredGroupnets

		if len(state.Groupnets) == 0 {
			resp.Diagnostics.AddError(
				"Error one or more of the filtered dns_servers values is not a valid powerscale groupnet.",
				fmt.Sprintf("Valid groupnets: [%v], filtered list: [%v]", strings.Join(validGroupnets, " , "), state.Filter.DNSServers),
			)
		}
	}

	state.ID = types.StringValue("groupnet_datasource")

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	tflog.Info(ctx, "Done with Read Groupnet data source ")
}

// function to check if a string is in a list
func contains(list []attr.Value, value attr.Value) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}
