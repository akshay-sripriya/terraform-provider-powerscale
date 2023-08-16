/*
Copyright (c) 2023 Dell Inc., or its subsidiaries. All Rights Reserved.

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
	powerscale "dell/powerscale-go-client"
	"fmt"
	"terraform-provider-powerscale/client"
	"terraform-provider-powerscale/powerscale/constants"
	"terraform-provider-powerscale/powerscale/helper"
	"terraform-provider-powerscale/powerscale/models"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &FileSystemResource{}
var _ resource.ResourceWithImportState = &FileSystemResource{}

// NewFileSystemResource creates a new data source.
func NewFileSystemResource() resource.Resource {
	return &FileSystemResource{}
}

// FileSystemResource defines the data source implementation.
type FileSystemResource struct {
	client *client.Client
}

// Metadata describes the data source arguments.
func (r *FileSystemResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_filesystem"
}

// Schema describes the resource arguments.
func (r *FileSystemResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "FileSystem resource",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "FileSystem identifier",
				MarkdownDescription: "FileSystem identifier",
				Computed:            true,
				Optional:            true,
			},
			"name": schema.StringAttribute{
				Description:         "FileSystem directory name",
				MarkdownDescription: "FileSystem directory name",
				Required:            true,
			},
			"directory_path": schema.StringAttribute{
				Description:         "FileSystem directory path. If no directory path is specified, [/ifs] would be taken by default.",
				MarkdownDescription: "FileSystem directory path. If no directory path is specified, [/ifs] would be taken by default.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("/ifs"),
			},
			"type": schema.StringAttribute{
				Description:         "File System Resource type",
				MarkdownDescription: "File System Resource type",
				Computed:            true,
			},
			"creation_time": schema.StringAttribute{
				Description:         "File System Resource Creation time",
				MarkdownDescription: "File System Resource Creation time",
				Computed:            true,
			},
			"owner": schema.SingleNestedAttribute{
				Description:         "The owner of the Filesystem.",
				MarkdownDescription: "The owner of the Filesystem.",
				Required:            true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description:         "Owner identifier",
						MarkdownDescription: "Owner identifier",
						Required:            true,
					},
					"name": schema.StringAttribute{
						Description:         "Owner name",
						MarkdownDescription: "Owner name",
						Required:            true,
					},
					"type": schema.StringAttribute{
						Description:         "Owner type",
						MarkdownDescription: "Owner type",
						Required:            true,
					},
				},
			},
			"group": schema.SingleNestedAttribute{
				Description:         "The group of the Filesystem.",
				MarkdownDescription: "The group of the Filesystem.",
				Required:            true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description:         "group identifier",
						MarkdownDescription: "group identifier",
						Required:            true,
					},
					"name": schema.StringAttribute{
						Description:         "group name",
						MarkdownDescription: "group name",
						Required:            true,
					},
					"type": schema.StringAttribute{
						Description:         "group type",
						MarkdownDescription: "group type",
						Required:            true,
					},
				},
			},
			"access_control": schema.StringAttribute{
				Description: `The ACL value for the directory. Users can either provide access rights input such as 'private_read' , 'private' ,
				'public_read', 'public_read_write', 'public' or permissions in POSIX format as '0550', '0770', '0775','0777' or 0700. The Default value is (0700). 
				Modification of ACL is only supported from POSIX to POSIX mode.`,
				MarkdownDescription: `The ACL value for the directory. Users can either provide access rights input such as 'private_read' , 'private' ,
				'public_read', 'public_read_write', 'public' or permissions in POSIX format as '0550', '0770', '0775','0777' or 0700. The Default value is (0700). 
				Modification of ACL is only supported from POSIX to POSIX mode.`,
				Optional: true,
			},
			"authoritative": schema.StringAttribute{
				Description:         "If the directory has access rights set, then this field returns acl. Otherwise it returns mode.",
				MarkdownDescription: "If the directory has access rights set, then this field returns acl. Otherwise it returns mode.",
				Computed:            true,
			},
			"recursive": schema.BoolAttribute{
				Description:         "Creates intermediate folders recursively when set to true.",
				MarkdownDescription: "Creates intermediate folders recursively when set to true.",
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"overwrite": schema.BoolAttribute{
				Description:         "Deletes and replaces the existing user attributes and ACLs of the directory with user-specified attributes if set to true.",
				MarkdownDescription: "Deletes and replaces the existing user attributes and ACLs of the directory with user-specified attributes if set to true.",
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"mode": schema.StringAttribute{
				Description:         "Acl mode",
				MarkdownDescription: "Acl mode",
				Computed:            true,
			},
		},
	}
}

// Configure configures the data source.
func (r *FileSystemResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	pscaleClient, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}
	r.client = pscaleClient
}

// Create creates the File system resource.
func (r *FileSystemResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Creating File System..")
	var plan models.FileSystemResource

	if resp.Diagnostics.HasError() {
		return
	}
	// Read Terraform plan data into the model
	resp.Diagnostics = append(resp.Diagnostics, req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	dirPath := helper.GetDirectoryPath(plan.DirectoryPath.ValueString(), plan.Name.ValueString())

	createReq := r.client.PscaleOpenAPIClient.NamespaceApi.CreateDirectory(ctx, dirPath)

	createReq = createReq.XIsiIfsTargetType("container")
	createReq = createReq.Overwrite(plan.Overwrite.ValueBool())
	createReq = createReq.Recursive(plan.Recursive.ValueBool())
	if !plan.AccessControl.IsNull() && (plan.AccessControl.ValueString() != "") {
		createReq = createReq.XIsiIfsAccessControl(plan.AccessControl.ValueString())
	}

	_, _, errCR := createReq.Execute()
	if errCR != nil {
		errStr := constants.CreateFileSystemErrorMsg + "with error: "
		message := helper.GetErrorString(errCR, errStr)
		resp.Diagnostics.AddError(
			"Error creating File System",
			message,
		)
		return
	}

	// set owner and group for file system
	setACLReq := r.client.PscaleOpenAPIClient.NamespaceApi.SetAcl(ctx, dirPath)
	setACLReq = setACLReq.Acl(true)

	namespaceACLUserGroup := *powerscale.NewNamespaceAcl()
	namespaceACLUserGroup.SetAuthoritative("mode")

	owner := *powerscale.NewMemberObject()
	owner.Id = plan.Owner.ID.ValueStringPointer()
	owner.Name = plan.Owner.Name.ValueStringPointer()
	owner.Type = plan.Owner.Type.ValueStringPointer()
	namespaceACLUserGroup.SetOwner(owner)

	group := *powerscale.NewMemberObject()
	group.Id = plan.Group.ID.ValueStringPointer()
	group.Name = plan.Group.Name.ValueStringPointer()
	group.Type = plan.Group.Type.ValueStringPointer()
	namespaceACLUserGroup.SetGroup(group)

	setACLReq = setACLReq.NamespaceAcl(namespaceACLUserGroup)

	_, _, err := setACLReq.Execute()
	if err != nil {
		errStr := constants.CreateFileSystemErrorMsg + "with error: "
		message := helper.GetErrorString(err, errStr)
		resp.Diagnostics.AddError(
			"Error Setting User / Groups for the filesystem",
			message,
		)
		return
	}

	// Get File system metadata
	meta, err := helper.GetDirectoryMetadata(ctx, r.client, dirPath)

	if err != nil {
		errStr := constants.CreateFileSystemErrorMsg + "with error: "
		message := helper.GetErrorString(err, errStr)
		resp.Diagnostics.AddError(
			"Error getting the metadata for the filesystem",
			message,
		)
		return
	}

	// Get Acl
	acl, err := helper.GetDirectoryACL(ctx, r.client, dirPath)
	if err != nil {
		errStr := constants.CreateFileSystemErrorMsg + "with error: "
		message := helper.GetErrorString(err, errStr)
		resp.Diagnostics.AddError(
			"Error getting the acl for the filesystem",
			message,
		)
		return
	}

	// Update resource state
	var state models.FileSystemResource
	helper.UpdateFileSystemResourceState(ctx, &plan, &state, acl, meta)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	tflog.Info(ctx, "Done with Create File System resource")
}

// Read reads data from the resource.
func (r *FileSystemResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Info(ctx, "Read File System Resource..")
	var plan models.FileSystemResource

	// Read Terraform prior state plan into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &plan)...)

	if resp.Diagnostics.HasError() {
		return
	}
	dirPath := helper.GetDirectoryPath(plan.DirectoryPath.ValueString(), plan.Name.ValueString())

	// Get metadata
	meta, err := helper.GetDirectoryMetadata(ctx, r.client, dirPath)

	if err != nil {
		errStr := constants.ReadFileSystemErrorMsg + "with error: "
		message := helper.GetErrorString(err, errStr)
		resp.Diagnostics.AddError(
			"Error getting the metadata for the filesystem",
			message,
		)
		return
	}

	// GetAcl
	acl, err := helper.GetDirectoryACL(ctx, r.client, dirPath)
	if err != nil {
		errStr := constants.ReadFileSystemErrorMsg + "with error: "
		message := helper.GetErrorString(err, errStr)
		resp.Diagnostics.AddError(
			"Error getting the acl for the filesystem",
			message,
		)
		return
	}

	//copy to model
	var state models.FileSystemResource
	helper.UpdateFileSystemResourceState(ctx, &plan, &state, acl, meta)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	tflog.Info(ctx, "Read File System Resource Complete.")
}

// Delete deletes the resource.
func (r *FileSystemResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Info(ctx, "Deleting File System Resource..")
	var plan models.FileSystemResource

	// Read Terraform prior state plan into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &plan)...)

	if resp.Diagnostics.HasError() {
		return
	}
	dirPath := helper.GetDirectoryPath(plan.DirectoryPath.ValueString(), plan.Name.ValueString())

	_, _, err := r.client.PscaleOpenAPIClient.NamespaceApi.DeleteDirectory(ctx, dirPath).Execute()
	if err != nil {
		errStr := constants.DeleteFileSystemErrorMsg + "with error: "
		message := helper.GetErrorString(err, errStr)
		resp.Diagnostics.AddError(
			"Error Deleting filesystem",
			message,
		)
		return
	}
	tflog.Info(ctx, "Delete File system complete")
}

// ImportState imports the resource state.
func (r *FileSystemResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	//TODO
}

// Update updates the resource state.
func (r *FileSystemResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	//TODO
}
