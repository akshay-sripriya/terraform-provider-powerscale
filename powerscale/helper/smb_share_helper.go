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

package helper

import (
	"context"
	powerscale "dell/powerscale-go-client"
	"terraform-provider-powerscale/client"
)

// DeleteSmbShare delete smb share
func DeleteSmbShare(ctx context.Context, client *client.Client, shareID string) error {
	_, err := client.PscaleOpenAPIClient.ProtocolsApi.DeleteProtocolsv7SmbShare(ctx, shareID).Execute()
	return err
}

// CreateSmbShare create smb share
func CreateSmbShare(ctx context.Context, client *client.Client, share powerscale.V7SmbShare) (*powerscale.Createv12SmbShareResponse, error) {
	shareID, _, err := client.PscaleOpenAPIClient.ProtocolsApi.CreateProtocolsv7SmbShare(ctx).V7SmbShare(share).Execute()
	return shareID, err
}

// GetSmbShare get smb share
func GetSmbShare(ctx context.Context, client *client.Client, shareID string) (*powerscale.V7SmbSharesExtended, error) {
	response, _, err := client.PscaleOpenAPIClient.ProtocolsApi.GetProtocolsv7SmbShare(ctx, shareID).Execute()
	return response, err
}

// UpdateSmbShare update smb share
func UpdateSmbShare(ctx context.Context, client *client.Client, shareID string, shareToUpdate powerscale.V7SmbShareExtendedExtended) error {
	updateParam := client.PscaleOpenAPIClient.ProtocolsApi.UpdateProtocolsv7SmbShare(ctx, shareID)
	_, err := updateParam.V7SmbShare(shareToUpdate).Execute()
	return err
}
