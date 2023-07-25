---
# Copyright (c) 2023 Dell Inc., or its subsidiaries. All Rights Reserved.
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

title: "powerscale_cluster data source"
linkTitle: "powerscale_cluster"
page_title: "powerscale_cluster Data Source - terraform-provider-powerscale"
subcategory: ""
description: |-
  The cluster attributes and cluster node information.
---

# powerscale_cluster (Data Source)

The cluster attributes and cluster node information.



<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `config` (Attributes) The configuration information of cluster. (see [below for nested schema](#nestedatt--config))
- `id` (String) Unique identifier of the cluster.
- `identity` (Attributes) Unprivileged cluster information for display when logging in. (see [below for nested schema](#nestedatt--identity))
- `internal_networks` (Attributes) V7ClusterInternalNetworks Configuration fields for internal networks. (see [below for nested schema](#nestedatt--internal_networks))
- `nodes` (Attributes) IsiClusterNodes struct for IsiClusterNodes (see [below for nested schema](#nestedatt--nodes))

<a id="nestedatt--config"></a>
### Nested Schema for `config`

Read-Only:

- `description` (String) Customer configurable description.
- `devices` (Attributes List) device (see [below for nested schema](#nestedatt--config--devices))
- `guid` (String) Cluster GUID.
- `join_mode` (String) Node join mode: 'manual' or 'secure'.
- `local_devid` (Number) Device ID of the queried node.
- `local_lnn` (Number) Device logical node number of the queried node.
- `local_serial` (String) Device serial number of the queried node.
- `name` (String) Cluster name.
- `onefs_version` (Attributes) version (see [below for nested schema](#nestedatt--config--onefs_version))
- `timezone` (Attributes) version (see [below for nested schema](#nestedatt--config--timezone))

<a id="nestedatt--config--devices"></a>
### Nested Schema for `config.devices`

Read-Only:

- `devid` (Number) Device ID.
- `guid` (String) Device GUID.
- `is_up` (Boolean) If true, this node is online and communicating with the local node and every other node with the is_up property normally
- `lnn` (Number) Device logical node number.


<a id="nestedatt--config--onefs_version"></a>
### Nested Schema for `config.onefs_version`

Read-Only:

- `build` (String) OneFS build string.
- `release` (String) Kernel release number.
- `revision` (String) OneFS build number.
- `type` (String) Kernel release type.
- `version` (String) Kernel full version information.


<a id="nestedatt--config--timezone"></a>
### Nested Schema for `config.timezone`

Read-Only:

- `abbreviation` (String) Timezone abbreviation.
- `custom` (String) Customer timezone information.
- `name` (String) Timezone full name.
- `path` (String) Timezone hierarchical name.



<a id="nestedatt--identity"></a>
### Nested Schema for `identity`

Read-Only:

- `description` (String) A description of the cluster.
- `logon` (Attributes) // (see [below for nested schema](#nestedatt--identity--logon))
- `name` (String) The name of the cluster.

<a id="nestedatt--identity--logon"></a>
### Nested Schema for `identity.logon`

Read-Only:

- `motd` (String) The message of the day.
- `motd_header` (String) The header to the message of the day.



<a id="nestedatt--internal_networks"></a>
### Nested Schema for `internal_networks`

Read-Only:

- `failover_ip_addresses` (Attributes List) Array of IP address ranges to be used to configure the internal failover network of the OneFS cluster. (see [below for nested schema](#nestedatt--internal_networks--failover_ip_addresses))
- `failover_status` (String) Status of failover network.
- `int_a_fabric` (String) Network fabric used for the primary network int-a.
- `int_a_ip_addresses` (Attributes List) Array of IP address ranges to be used to configure the internal int-a network of the OneFS cluster. (see [below for nested schema](#nestedatt--internal_networks--int_a_ip_addresses))
- `int_a_mtu` (Number) Maximum Transfer Unit (MTU) of the primary network int-a.
- `int_a_prefix_length` (Number) Prefixlen specifies the length of network bits used in an IP address. This field is the right-hand part of the CIDR notation representing the subnet mask.
- `int_a_status` (String) Status of the primary network int-a.
- `int_b_fabric` (String) Network fabric used for the failover network.
- `int_b_ip_addresses` (Attributes List) Array of IP address ranges to be used to configure the internal int-b network of the OneFS cluster. (see [below for nested schema](#nestedatt--internal_networks--int_b_ip_addresses))
- `int_b_mtu` (Number) Maximum Transfer Unit (MTU) of the failover network int-b.
- `int_b_prefix_length` (Number) Prefixlen specifies the length of network bits used in an IP address. This field is the right-hand part of the CIDR notation representing the subnet mask.

<a id="nestedatt--internal_networks--failover_ip_addresses"></a>
### Nested Schema for `internal_networks.failover_ip_addresses`

Read-Only:

- `high` (String) IPv4 address in the format: xxx.xxx.xxx.xxx
- `low` (String) IPv4 address in the format: xxx.xxx.xxx.xxx


<a id="nestedatt--internal_networks--int_a_ip_addresses"></a>
### Nested Schema for `internal_networks.int_a_ip_addresses`

Read-Only:

- `high` (String) IPv4 address in the format: xxx.xxx.xxx.xxx
- `low` (String) IPv4 address in the format: xxx.xxx.xxx.xxx


<a id="nestedatt--internal_networks--int_b_ip_addresses"></a>
### Nested Schema for `internal_networks.int_b_ip_addresses`

Read-Only:

- `high` (String) IPv4 address in the format: xxx.xxx.xxx.xxx
- `low` (String) IPv4 address in the format: xxx.xxx.xxx.xxx



<a id="nestedatt--nodes"></a>
### Nested Schema for `nodes`

Read-Only:

- `errors` (Attributes List) A list of errors encountered by the individual nodes involved in this request, or an empty list if there were no errors. (see [below for nested schema](#nestedatt--nodes--errors))
- `nodes` (Attributes List) The responses from the individual nodes involved in this request. (see [below for nested schema](#nestedatt--nodes--nodes))
- `total` (Number) The total number of nodes responding.

<a id="nestedatt--nodes--errors"></a>
### Nested Schema for `nodes.errors`

Read-Only:

- `code` (String) The error code.
- `field` (String) The field with the error if applicable.
- `id` (Number) Node ID (Device Number) of a node.
- `lnn` (Number) Logical Node Number (LNN) of a node.
- `message` (String) The error message.
- `status` (Number) HTTP Status code returned by this node.


<a id="nestedatt--nodes--nodes"></a>
### Nested Schema for `nodes.nodes`

Read-Only:

- `drives` (Attributes List) List of the drives in this node. (see [below for nested schema](#nestedatt--nodes--nodes--drives))
- `error` (String) Error message, if the HTTP status returned from this node was not 200.
- `hardware` (Attributes) // (see [below for nested schema](#nestedatt--nodes--nodes--hardware))
- `id` (Number) Node ID (Device Number) of a node.
- `lnn` (Number) Logical Node Number (LNN) of a node.
- `partitions` (Attributes) // (see [below for nested schema](#nestedatt--nodes--nodes--partitions))
- `sensors` (Attributes) // (see [below for nested schema](#nestedatt--nodes--nodes--sensors))
- `state` (Attributes) // (see [below for nested schema](#nestedatt--nodes--nodes--state))
- `status` (Attributes) // (see [below for nested schema](#nestedatt--nodes--nodes--status))

<a id="nestedatt--nodes--nodes--drives"></a>
### Nested Schema for `nodes.nodes.drives`

Read-Only:

- `baynum` (Number) Numerical representation of this drive's bay.
- `blocks` (Number) Number of blocks on this drive.
- `chassis` (Number) The chassis number which contains this drive.
- `devname` (String) This drive's device name.
- `firmware` (Attributes) Drive firmware information (see [below for nested schema](#nestedatt--nodes--nodes--drives--firmware))
- `handle` (Number) Drive_d's handle representation for this driveIf we fail to retrieve the handle for this drive from drive_d: -1
- `interface_type` (String) String representation of this drive's interface type.
- `lnum` (Number) This drive's logical drive number in IFS.
- `locnstr` (String) String representation of this drive's physical location.
- `logical_block_length` (Number) Size of a logical block on this drive.
- `media_type` (String) String representation of this drive's media type.
- `model` (String) This drive's manufacturer and model.
- `physical_block_length` (Number) Size of a physical block on this drive.
- `present` (Boolean) Indicates whether this drive is physically present in the node.
- `purpose` (String) This drive's purpose in the DRV state machine.
- `purpose_description` (String) Description of this drive's purpose.
- `serial` (String) Serial number for this drive.
- `ui_state` (String) This drive's state as presented to the UI.
- `wwn` (String) The drive's 'worldwide name' from its NAA identifiers.
- `x_loc` (Number) This drive's x-axis grid location.
- `y_loc` (Number) This drive's y-axis grid location.

<a id="nestedatt--nodes--nodes--drives--firmware"></a>
### Nested Schema for `nodes.nodes.drives.y_loc`

Read-Only:

- `current_firmware` (String) This drive's current firmware revision
- `desired_firmware` (String) This drive's desired firmware revision.



<a id="nestedatt--nodes--nodes--hardware"></a>
### Nested Schema for `nodes.nodes.hardware`

Read-Only:

- `chassis` (String) Name of this node's chassis.
- `chassis_code` (String) Chassis code of this node (1U, 2U, etc.).
- `chassis_count` (String) Number of chassis making up this node.
- `class` (String) Class of this node (storage, accelerator, etc.).
- `configuration_id` (String) Node configuration ID.
- `cpu` (String) Manufacturer and model of this node's CPU.
- `disk_controller` (String) Manufacturer and model of this node's disk controller.
- `disk_expander` (String) Manufacturer and model of this node's disk expander.
- `family_code` (String) Family code of this node (X, S, NL, etc.).
- `flash_drive` (String) Manufacturer, model, and device id of this node's flash drive.
- `generation_code` (String) Generation code of this node.
- `hwgen` (String) PowerScale hardware generation name.
- `imb_version` (String) Version of this node's PowerScale Management Board.
- `infiniband` (String) Infiniband card type.
- `lcd_version` (String) Version of the LCD panel.
- `motherboard` (String) Manufacturer and model of this node's motherboard.
- `net_interfaces` (String) Description of all this node's network interfaces.
- `nvram` (String) Manufacturer and model of this node's NVRAM board.
- `powersupplies` (List of String) Description strings for each power supply on this node.
- `processor` (String) Number of processors and cores on this node.
- `product` (String) PowerScale product name.
- `ram` (Number) Size of RAM in bytes.
- `serial_number` (String) Serial number of this node.
- `series` (String) Series of this node (X, I, NL, etc.).
- `storage_class` (String) Storage class of this node (storage or diskless).


<a id="nestedatt--nodes--nodes--partitions"></a>
### Nested Schema for `nodes.nodes.partitions`

Read-Only:

- `count` (Number) Count of how many partitions are included.
- `partitions` (Attributes List) Partition information. (see [below for nested schema](#nestedatt--nodes--nodes--partitions--partitions))

<a id="nestedatt--nodes--nodes--partitions--partitions"></a>
### Nested Schema for `nodes.nodes.partitions.partitions`

Read-Only:

- `block_size` (Number) The block size used for the reported partition information.
- `capacity` (Number) Total blocks on this file system partition.
- `component_devices` (String) Comma separated list of devices used for this file system partition.
- `mount_point` (String) Directory on which this partition is mounted.
- `percent_used` (String) Used blocks on this file system partition, expressed as a percentage.
- `statfs` (Attributes) // (see [below for nested schema](#nestedatt--nodes--nodes--partitions--partitions--statfs))
- `used` (Number) Used blocks on this file system partition.

<a id="nestedatt--nodes--nodes--partitions--partitions--statfs"></a>
### Nested Schema for `nodes.nodes.partitions.partitions.statfs`

Read-Only:

- `f_bavail` (Number) Free blocks available to non-superuser on this partition.
- `f_bfree` (Number) Free blocks on this partition.
- `f_blocks` (Number) Total data blocks on this partition.
- `f_bsize` (Number) Filesystem fragment size; block size in OneFS.
- `f_ffree` (Number) Free file nodes avail to non-superuser.
- `f_files` (Number) Total file nodes in filesystem.
- `f_flags` (Number) Mount exported flags.
- `f_fstypename` (String) File system type name.
- `f_iosize` (Number) Optimal transfer block size.
- `f_mntfromname` (String) Names of devices this partition is mounted from.
- `f_mntonname` (String) Directory this partition is mounted to.
- `f_namemax` (Number) Maximum filename length.
- `f_owner` (Number) UID of user that mounted the filesystem.
- `f_type` (Number) Type of filesystem.
- `f_version` (Number) statfs() structure version number.




<a id="nestedatt--nodes--nodes--sensors"></a>
### Nested Schema for `nodes.nodes.sensors`

Read-Only:

- `sensors` (Attributes List) This node's sensor information. (see [below for nested schema](#nestedatt--nodes--nodes--sensors--sensors))

<a id="nestedatt--nodes--nodes--sensors--sensors"></a>
### Nested Schema for `nodes.nodes.sensors.sensors`

Read-Only:

- `count` (Number) The count of values in this sensor group.
- `name` (String) The name of this sensor group.
- `values` (Attributes List) The list of specific sensor value info in this sensor group. (see [below for nested schema](#nestedatt--nodes--nodes--sensors--sensors--values))

<a id="nestedatt--nodes--nodes--sensors--sensors--values"></a>
### Nested Schema for `nodes.nodes.sensors.sensors.values`

Read-Only:

- `desc` (String) The descriptive name of this sensor.
- `name` (String) The identifier name of this sensor.
- `units` (String) The units of this sensor.
- `value` (String) The value of this sensor.




<a id="nestedatt--nodes--nodes--state"></a>
### Nested Schema for `nodes.nodes.state`

Read-Only:

- `readonly` (Attributes) // (see [below for nested schema](#nestedatt--nodes--nodes--state--readonly))
- `servicelight` (Attributes) // (see [below for nested schema](#nestedatt--nodes--nodes--state--servicelight))
- `smartfail` (Attributes) // (see [below for nested schema](#nestedatt--nodes--nodes--state--smartfail))

<a id="nestedatt--nodes--nodes--state--readonly"></a>
### Nested Schema for `nodes.nodes.state.smartfail`

Read-Only:

- `allowed` (Boolean) The current read-only mode allowed status for the node.
- `enabled` (Boolean) The current read-only user mode status for the node. NOTE: If read-only mode is currently disallowed for this node, it will remain read/write until read-only mode is allowed again. This value only sets or clears any user-specified requests for read-only mode. If the node has been placed into read-only mode by the system, it will remain in read-only mode until the system conditions which triggered read-only mode have cleared.
- `mode` (Boolean) The current read-only mode status for the node.
- `status` (String) The current read-only mode status description for the node.
- `valid` (Boolean) The read-only state values are valid (False = Error).
- `value` (Number) The current read-only value (enumerated bitfield) for the node.


<a id="nestedatt--nodes--nodes--state--servicelight"></a>
### Nested Schema for `nodes.nodes.state.smartfail`

Read-Only:

- `enabled` (Boolean) The node service light state (True = on).


<a id="nestedatt--nodes--nodes--state--smartfail"></a>
### Nested Schema for `nodes.nodes.state.smartfail`

Read-Only:

- `smartfailed` (Boolean) This node is smartfailed (soft_devs).



<a id="nestedatt--nodes--nodes--status"></a>
### Nested Schema for `nodes.nodes.status`

Read-Only:

- `batterystatus` (Attributes) // (see [below for nested schema](#nestedatt--nodes--nodes--status--batterystatus))
- `capacity` (Attributes List) Storage capacity of this node. (see [below for nested schema](#nestedatt--nodes--nodes--status--capacity))
- `cpu` (Attributes) // (see [below for nested schema](#nestedatt--nodes--nodes--status--cpu))
- `nvram` (Attributes) // (see [below for nested schema](#nestedatt--nodes--nodes--status--nvram))
- `powersupplies` (Attributes) // (see [below for nested schema](#nestedatt--nodes--nodes--status--powersupplies))
- `release` (String) OneFS release.
- `uptime` (Number) Seconds this node has been online.
- `version` (String) OneFS version.

<a id="nestedatt--nodes--nodes--status--batterystatus"></a>
### Nested Schema for `nodes.nodes.status.version`

Read-Only:

- `last_test_time1` (String) The last battery test time for battery 1.
- `last_test_time2` (String) The last battery test time for battery 2.
- `next_test_time1` (String) The next checkup for battery 1.
- `next_test_time2` (String) The next checkup for battery 2.
- `present` (Boolean) Node has battery status.
- `result1` (String) The result of the last battery test for battery 1.
- `result2` (String) The result of the last battery test for battery 2.
- `status1` (String) The status of battery 1.
- `status2` (String) The status of battery 2.
- `supported` (Boolean) Node supports battery status.


<a id="nestedatt--nodes--nodes--status--capacity"></a>
### Nested Schema for `nodes.nodes.status.version`

Read-Only:

- `bytes` (Number) Total device storage bytes.
- `count` (Number) Total device count.
- `type` (String) Device type.


<a id="nestedatt--nodes--nodes--status--cpu"></a>
### Nested Schema for `nodes.nodes.status.version`

Read-Only:

- `model` (String) Manufacturer model description of this CPU.
- `overtemp` (String) CPU overtemp state.
- `proc` (String) Type of processor and core of this CPU.
- `speed_limit` (String) CPU throttling (expressed as a percentage).


<a id="nestedatt--nodes--nodes--status--nvram"></a>
### Nested Schema for `nodes.nodes.status.version`

Read-Only:

- `batteries` (Attributes List) This node's NVRAM battery status information. (see [below for nested schema](#nestedatt--nodes--nodes--status--version--batteries))
- `battery_count` (Number) This node's NVRAM battery count. On failure: -1, otherwise 1 or 2.
- `charge_status` (String) This node's NVRAM battery charge status, as a color.
- `charge_status_number` (Number) This node's NVRAM battery charge status, as a number. Error or not supported: -1. BR_BLACK: 0. BR_GREEN: 1. BR_YELLOW: 2. BR_RED: 3.
- `device` (String) This node's NVRAM device name with path.
- `present` (Boolean) This node has NVRAM.
- `present_flash` (Boolean) This node has NVRAM with flash storage.
- `present_size` (Number) The size of the NVRAM, in bytes.
- `present_type` (String) This node's NVRAM type.
- `ship_mode` (Number) This node's current ship mode state for NVRAM batteries. If not supported or on failure: -1. Disabled: 0. Enabled: 1.
- `supported` (Boolean) This node supports NVRAM.
- `supported_flash` (Boolean) This node supports NVRAM with flash storage.
- `supported_size` (Number) The maximum size of the NVRAM, in bytes.
- `supported_type` (String) This node's supported NVRAM type.

<a id="nestedatt--nodes--nodes--status--version--batteries"></a>
### Nested Schema for `nodes.nodes.status.version.batteries`

Read-Only:

- `color` (String) The current status color of the NVRAM battery.
- `id` (Number) Identifying index for the NVRAM battery.
- `status` (String) The current status message of the NVRAM battery.
- `voltage` (String) The current voltage of the NVRAM battery.



<a id="nestedatt--nodes--nodes--status--powersupplies"></a>
### Nested Schema for `nodes.nodes.status.version`

Read-Only:

- `count` (Number) Count of how many power supplies are supported.
- `failures` (Number) Count of how many power supplies have failed.
- `has_cff` (Boolean) Does this node have a CFF power supply.
- `status` (String) A descriptive status string for this node's power supplies.
- `supplies` (Attributes List) List of this node's power supplies. (see [below for nested schema](#nestedatt--nodes--nodes--status--version--supplies))
- `supports_cff` (Boolean) Does this node support CFF power supplies.

<a id="nestedatt--nodes--nodes--status--version--supplies"></a>
### Nested Schema for `nodes.nodes.status.version.supplies`

Read-Only:

- `chassis` (Number) Which node chassis is this power supply in.
- `firmware` (String) The current firmware revision of this power supply.
- `good` (String) Is this power supply in a failure state.
- `id` (Number) Identifying index for this power supply.
- `name` (String) Complete identifying string for this power supply.
- `status` (String) A descriptive status string for this power supply.
- `type` (String) The type of this power supply.