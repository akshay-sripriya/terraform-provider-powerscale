---
# Copyright (c) 2023-2024 Dell Inc., or its subsidiaries. All Rights Reserved.
#
# Licensed under the Mozilla Public License Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://mozilla.org/MPL/2.0/
#
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

title: "powerscale_smartpool_settings resource"
linkTitle: "powerscale_smartpool_settings"
page_title: "powerscale_smartpool_settings Resource - terraform-provider-powerscale"
subcategory: ""
description: |-
  This resource is used to manage the SmartPools Settings of PowerScale Array. We can Create, Update and Delete the SmartPools Settings using this resource.Note that, SmartPools Settings is the native functionality of PowerScale. When creating the resource, we actually load SmartPools Settings from PowerScale to the resource.
---

# powerscale_smartpool_settings (Resource)

This resource is used to manage the SmartPools Settings of PowerScale Array. We can Create, Update and Delete the SmartPools Settings using this resource.  
Note that, SmartPools Settings is the native functionality of PowerScale. When creating the resource, we actually load SmartPools Settings from PowerScale to the resource.


## Example Usage

```terraform
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

# Available actions: Create, Update, Delete and Import
# If resource arguments are omitted, `terraform apply` will load SmartPools Settings from PowerScale, and save to
# terraform state file.
# If any resource arguments are specified, `terraform apply` will try to load SmartPools Settings (if not loaded) and update the settings.
# `terraform destroy` will delete the resource from terraform state file rather than deleting SmartPools Settings from PowerScale.
# For more information, Please check the terraform state file.

resource "powerscale_smartpool_settings" "settings" {
  #    global_namespace_acceleration_enabled = false
  #    manage_io_optimization                = true
  #    manage_io_optimization_apply_to_files = false
  #    manage_protection                     = true
  #    manage_protection_apply_to_files      = false
  #    protect_directories_one_level_higher  = true
  #    spillover_enabled                     = true
  #    spillover_target                      = {
  #      name    = "sample_storagepool"
  #      type    = "storagepool" // anywhere or storagepool. name should be empty string when type is anywhere
  #    }
  #    ssd_l3_cache_default_enabled          = true
  #    ssd_qab_mirrors                       = "all"
  #    ssd_system_btree_mirrors              = "all"
  #    ssd_system_delta_mirrors              = "all"
  #    virtual_hot_spare_deny_writes         = true
  #    virtual_hot_spare_hide_spare          = true
  #    virtual_hot_spare_limit_drives        = 4
  #    virtual_hot_spare_limit_percent       = 4
  #  # Note that, default_transfer_limit_state and default_transfer_limit_pct are mutually exclusive and only one can be specified.
  #    default_transfer_limit_state          = "disabled" // available for PowerScale 9.5 and above
  #    default_transfer_limit_pct            = 90 // available for PowerScale 9.5 and above
}

# After the execution of above resource block, SmartPools Settings would have been cached in terraform state file, or
# SmartPools Settings would have been updated on PowerScale.
# For more information, Please check the terraform state file.
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `default_transfer_limit_pct` (Number) Applies to all storagepools that fall back on the default transfer limit. Stop moving files to this pool when this limit is met. The value must be between 0 and 100. Only available for PowerScale 9.5 and above.
- `default_transfer_limit_state` (String) How the default transfer limit value is applied. Only available for PowerScale 9.5 and above.
- `global_namespace_acceleration_enabled` (Boolean) Enable global namespace acceleration.
- `global_namespace_acceleration_state` (String) Whether or not namespace operation optimizations are currently in effect.
- `manage_io_optimization` (Boolean) Manage I/O optimization settings.
- `manage_io_optimization_apply_to_files` (Boolean) Apply to files with manually-managed I/O optimization settings.
- `manage_protection` (Boolean) Manage protection settings.
- `manage_protection_apply_to_files` (Boolean) Apply to files with manually-managed protection.
- `protect_directories_one_level_higher` (Boolean) Increase directory protection to a higher requested protection than its contents.
- `spillover_enabled` (Boolean) Enable global spillover.
- `spillover_target` (Attributes) Spillover data target. (see [below for nested schema](#nestedatt--spillover_target))
- `ssd_l3_cache_default_enabled` (Boolean) Use SSDs as L3 cache by default for new node pools.
- `ssd_qab_mirrors` (String) Controls number of mirrors of QAB blocks to place on SSDs. Acceptable values: one, all
- `ssd_system_btree_mirrors` (String) Controls number of mirrors of system B-tree blocks to place on SSDs. Acceptable values: one, all
- `ssd_system_delta_mirrors` (String) Controls number of mirrors of system delta blocks to place on SSDs. Acceptable values: one, all
- `virtual_hot_spare_deny_writes` (Boolean) Deny data writes to reserved disk space
- `virtual_hot_spare_hide_spare` (Boolean) Subtract the space reserved for the virtual hot spare when calculating available free space
- `virtual_hot_spare_limit_drives` (Number) The number of drives to reserve for the virtual hot spare, from 0-4.
- `virtual_hot_spare_limit_percent` (Number) The percent space to reserve for the virtual hot spare, from 0-20.

### Read-Only

- `id` (String) Id of SmartPools settings. Readonly. Fixed value of "smartpools_settings"

<a id="nestedatt--spillover_target"></a>
### Nested Schema for `spillover_target`

Optional:

- `name` (String) Target pool name if target specified as storagepool, otherwise empty string.
- `type` (String) Type of target pool. Acceptable values: storagepool, anywhere

Unless specified otherwise, all fields of this resource can be updated.

## Import

Import is supported using the following syntax:

```shell
# Copyright (c) 2023-2024 Dell Inc., or its subsidiaries. All Rights Reserved.

# Licensed under the Mozilla Public License Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at

#     http://mozilla.org/MPL/2.0/


# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# The command is
# terraform import powerscale_smartpool_settings.settings <anystring>
# Example:
terraform import powerscale_smartpool_settings.settings smartpools_settings
# after running this command, populate the name field and other required parameters in the config file to start managing this resource.
# Note: running "terraform show" after importing shows the current config/state of the resource. You can copy/paste that config to make it easier to manage the resource.
```