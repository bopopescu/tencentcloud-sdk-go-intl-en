// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20190719

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AvailableProtoStatus struct {

	// Sale status. Valid values: sale_out (sold out), saling (purchasable), no_saling (non-purchasable)
	SaleStatus *string `json:"SaleStatus,omitempty" name:"SaleStatus"`

	// Protocol type. Valid values: NFS, CIFS
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type AvailableRegion struct {

	// Region name, such as "ap-beijing"
	Region *string `json:"Region,omitempty" name:"Region"`

	// Region name, such as "bj"
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// Region availability. If a region has at least one AZ where resources are purchasable, this value will be `AVAILABLE`; otherwise, it will be `UNAVAILABLE`
	RegionStatus *string `json:"RegionStatus,omitempty" name:"RegionStatus"`

	// Array of AZs
	Zones []*AvailableZone `json:"Zones,omitempty" name:"Zones" list`

	// Region name, such as "Guangzhou"
	RegionCnName *string `json:"RegionCnName,omitempty" name:"RegionCnName"`
}

type AvailableType struct {

	// Protocol and sale details
	Protocols []*AvailableProtoStatus `json:"Protocols,omitempty" name:"Protocols" list`

	// Storage class. Valid values: SD (standard), HP (high-performance)
	Type *string `json:"Type,omitempty" name:"Type"`
}

type AvailableZone struct {

	// AZ name
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// AZ ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// Chinese name of an AZ
	ZoneCnName *string `json:"ZoneCnName,omitempty" name:"ZoneCnName"`

	// Array of classes
	Types []*AvailableType `json:"Types,omitempty" name:"Types" list`

	// Chinese and English names of an AZ
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

type CreateCfsFileSystemRequest struct {
	*tchttp.BaseRequest

	// AZ name, such as "ap-beijing-1". For the list of regions and AZs, please see [Overview](https://cloud.tencent.com/document/product/582/13225)
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Network type. Valid values: VPC (VPC), BASIC (basic network)
	NetInterface *string `json:"NetInterface,omitempty" name:"NetInterface"`

	// Permission group ID
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// File system protocol type. Valid values: NFS, CIFS. If this parameter is left empty, NFS will be used by default
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// File system storage class. Valid values: SD (standard), HP (high-performance)
	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`

	// VPC ID. This field is required if network type is VPC.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID. This field is required if network type is VPC.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Specifies an IP address, which is supported only for VPC. If this parameter is left empty, a random IP will be assigned in the subnet
	MountIP *string `json:"MountIP,omitempty" name:"MountIP"`

	// Custom file system name
	FsName *string `json:"FsName,omitempty" name:"FsName"`

	// File system tag
	ResourceTags []*TagInfo `json:"ResourceTags,omitempty" name:"ResourceTags" list`
}

func (r *CreateCfsFileSystemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateCfsFileSystemRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateCfsFileSystemResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// File system creation time
		CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

		// Custom file system name
		CreationToken *string `json:"CreationToken,omitempty" name:"CreationToken"`

		// File system ID
		FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

		// File system status
		LifeCycleState *string `json:"LifeCycleState,omitempty" name:"LifeCycleState"`

		// Used file system capacity
		SizeByte *uint64 `json:"SizeByte,omitempty" name:"SizeByte"`

		// AZ ID
		ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

		// Custom file system name
		FsName *string `json:"FsName,omitempty" name:"FsName"`

		// Whether a file system is encrypted
		Encrypted *bool `json:"Encrypted,omitempty" name:"Encrypted"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCfsFileSystemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateCfsFileSystemResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateCfsPGroupRequest struct {
	*tchttp.BaseRequest

	// Permission group name, which can contain 1–64 Chinese characters, letters, numbers, underscores, or dashes
	Name *string `json:"Name,omitempty" name:"Name"`

	// Permission group description, which can contain 1–255 characters
	DescInfo *string `json:"DescInfo,omitempty" name:"DescInfo"`
}

func (r *CreateCfsPGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateCfsPGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateCfsPGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Permission group ID
		PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

		// Permission group name
		Name *string `json:"Name,omitempty" name:"Name"`

		// Permission group description
		DescInfo *string `json:"DescInfo,omitempty" name:"DescInfo"`

		// The number of file systems bound to this permission group
		BindCfsNum *int64 `json:"BindCfsNum,omitempty" name:"BindCfsNum"`

		// Permission group creation time
		CDate *string `json:"CDate,omitempty" name:"CDate"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCfsPGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateCfsPGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateCfsRuleRequest struct {
	*tchttp.BaseRequest

	// Permission group ID
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// You can enter a single IP or IP range, such as 10.1.10.11 or 10.10.1.0/24. The default visiting address is `*`, indicating that all IPs are allowed. Please note that you need to enter the CVM instance's private IP here.
	AuthClientIp *string `json:"AuthClientIp,omitempty" name:"AuthClientIp"`

	// Rule priority. Value range: 1–100. 1 represents the highest priority, while 100 the lowest
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`

	// Read/write permission. Valid values: RO (read-only), RW (read & write). Default value: RO
	RWPermission *string `json:"RWPermission,omitempty" name:"RWPermission"`

	// User permission. Valid values: all_squash, no_all_squash, root_squash, no_root_squash. Specifically, all_squash: any visiting user will be mapped to an anonymous user or user group; no_all_squash: a visiting user will be first matched with a local user, and if the match fails, it will be mapped to an anonymous user or user group; root_squash: a visiting root user will be mapped to an anonymous user or user group; no_root_squash: a visiting root user will be allowed to maintain root account permissions. Default value: root_squash.
	UserPermission *string `json:"UserPermission,omitempty" name:"UserPermission"`
}

func (r *CreateCfsRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateCfsRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateCfsRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Rule ID
		RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

		// Permission group ID
		PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

		// Client IP
		AuthClientIp *string `json:"AuthClientIp,omitempty" name:"AuthClientIp"`

		// Read & write permissions
		RWPermission *string `json:"RWPermission,omitempty" name:"RWPermission"`

		// User permission
		UserPermission *string `json:"UserPermission,omitempty" name:"UserPermission"`

		// Priority
		Priority *int64 `json:"Priority,omitempty" name:"Priority"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCfsRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateCfsRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteCfsFileSystemRequest struct {
	*tchttp.BaseRequest

	// File system ID. Note: please call the `DeleteMountTarget` API to delete the mount target first before deleting a file system; otherwise, the delete operation will fail.
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

func (r *DeleteCfsFileSystemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteCfsFileSystemRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteCfsFileSystemResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCfsFileSystemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteCfsFileSystemResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteCfsPGroupRequest struct {
	*tchttp.BaseRequest

	// Permission group ID
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`
}

func (r *DeleteCfsPGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteCfsPGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteCfsPGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Permission group ID
		PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

		// User ID
		AppId *int64 `json:"AppId,omitempty" name:"AppId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCfsPGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteCfsPGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteCfsRuleRequest struct {
	*tchttp.BaseRequest

	// Permission group ID
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DeleteCfsRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteCfsRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteCfsRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Rule ID
		RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

		// Permission group ID
		PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCfsRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteCfsRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMountTargetRequest struct {
	*tchttp.BaseRequest

	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// Mount target ID
	MountTargetId *string `json:"MountTargetId,omitempty" name:"MountTargetId"`
}

func (r *DeleteMountTargetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMountTargetRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMountTargetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMountTargetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMountTargetResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableZoneInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAvailableZoneInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAvailableZoneInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableZoneInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Information such as resource availability in each AZ and the supported storage classes and protocols
		RegionZones []*AvailableRegion `json:"RegionZones,omitempty" name:"RegionZones" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAvailableZoneInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAvailableZoneInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCfsFileSystemsRequest struct {
	*tchttp.BaseRequest

	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

func (r *DescribeCfsFileSystemsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCfsFileSystemsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCfsFileSystemsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// File system information
		FileSystems []*FileSystemInfo `json:"FileSystems,omitempty" name:"FileSystems" list`

		// Total number of file systems
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCfsFileSystemsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCfsFileSystemsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCfsPGroupsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCfsPGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCfsPGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCfsPGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Permission group information list
		PGroupList []*PGroupInfo `json:"PGroupList,omitempty" name:"PGroupList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCfsPGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCfsPGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCfsRulesRequest struct {
	*tchttp.BaseRequest

	// Permission group ID
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`
}

func (r *DescribeCfsRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCfsRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCfsRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of permission group rules
		RuleList []*PGroupRuleInfo `json:"RuleList,omitempty" name:"RuleList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCfsRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCfsRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCfsServiceStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCfsServiceStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCfsServiceStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCfsServiceStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Current status of the CFS service for this user. Valid values: none (not activated), creating (activating), created (activated)
		CfsServiceStatus *string `json:"CfsServiceStatus,omitempty" name:"CfsServiceStatus"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCfsServiceStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCfsServiceStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMountTargetsRequest struct {
	*tchttp.BaseRequest

	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

func (r *DescribeMountTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMountTargetsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMountTargetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Mount target details
		MountTargets []*MountInfo `json:"MountTargets,omitempty" name:"MountTargets" list`

		// The number of mount targets
		NumberOfMountTargets *int64 `json:"NumberOfMountTargets,omitempty" name:"NumberOfMountTargets"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMountTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMountTargetsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type FileSystemInfo struct {

	// Creation time
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// Custom name
	CreationToken *string `json:"CreationToken,omitempty" name:"CreationToken"`

	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// File system status
	LifeCycleState *string `json:"LifeCycleState,omitempty" name:"LifeCycleState"`

	// Used file system capacity
	SizeByte *uint64 `json:"SizeByte,omitempty" name:"SizeByte"`

	// Maximum storage limit of a file system
	SizeLimit *uint64 `json:"SizeLimit,omitempty" name:"SizeLimit"`

	// Region ID
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// Region name
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// File system protocol type
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// File system storage class
	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`

	// Prepaid storage pack bound to a file system (not supported currently)
	StorageResourcePkg *string `json:"StorageResourcePkg,omitempty" name:"StorageResourcePkg"`

	// Prepaid bandwidth pack bound to a file system (not supported currently)
	BandwidthResourcePkg *string `json:"BandwidthResourcePkg,omitempty" name:"BandwidthResourcePkg"`

	// Information of permission groups bound to a file system
	PGroup *PGroup `json:"PGroup,omitempty" name:"PGroup"`

	// Custom name
	FsName *string `json:"FsName,omitempty" name:"FsName"`

	// Whether a file system is encrypted
	Encrypted *bool `json:"Encrypted,omitempty" name:"Encrypted"`

	// Key used for encryption, which can be the key ID or ARN
	KmsKeyId *string `json:"KmsKeyId,omitempty" name:"KmsKeyId"`
}

type MountInfo struct {

	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// Mount target ID
	MountTargetId *string `json:"MountTargetId,omitempty" name:"MountTargetId"`

	// Mount target IP
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`

	// Mount root-directory
	FSID *string `json:"FSID,omitempty" name:"FSID"`

	// Mount target status
	LifeCycleState *string `json:"LifeCycleState,omitempty" name:"LifeCycleState"`

	// Network type
	NetworkInterface *string `json:"NetworkInterface,omitempty" name:"NetworkInterface"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VPC name
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Subnet name
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
}

type PGroup struct {

	// Permission group ID
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// Permission group name
	Name *string `json:"Name,omitempty" name:"Name"`
}

type PGroupInfo struct {

	// Permission group ID
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// Permission group name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Description
	DescInfo *string `json:"DescInfo,omitempty" name:"DescInfo"`

	// Creation time
	CDate *string `json:"CDate,omitempty" name:"CDate"`

	// The number of bound file system
	BindCfsNum *int64 `json:"BindCfsNum,omitempty" name:"BindCfsNum"`
}

type PGroupRuleInfo struct {

	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Client IP allowed for access
	AuthClientIp *string `json:"AuthClientIp,omitempty" name:"AuthClientIp"`

	// Read/write permission. ro: read-only; rw: read & write
	RWPermission *string `json:"RWPermission,omitempty" name:"RWPermission"`

	// User permission. all_squash: any visiting user will be mapped to an anonymous user or user group; no_all_squash: a visiting user will be first matched with a local user, and if the match fails, it will be mapped to an anonymous user or user group; root_squash: a visiting root user will be mapped to an anonymous user or user group; no_root_squash: a visiting root user will be allowed to maintain root account permissions.
	UserPermission *string `json:"UserPermission,omitempty" name:"UserPermission"`

	// Rule priority. Value range: 1–100. 1 represents the highest priority, while 100 the lowest
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
}

type SignUpCfsServiceRequest struct {
	*tchttp.BaseRequest
}

func (r *SignUpCfsServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SignUpCfsServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SignUpCfsServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Current status of the CFS service for this user. Valid values: none (not activated), creating (activating), created (activated)
		CfsServiceStatus *string `json:"CfsServiceStatus,omitempty" name:"CfsServiceStatus"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SignUpCfsServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SignUpCfsServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TagInfo struct {

	// Tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type UpdateCfsFileSystemNameRequest struct {
	*tchttp.BaseRequest

	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// Custom file system name
	FsName *string `json:"FsName,omitempty" name:"FsName"`
}

func (r *UpdateCfsFileSystemNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateCfsFileSystemNameRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateCfsFileSystemNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Custom file system name
		CreationToken *string `json:"CreationToken,omitempty" name:"CreationToken"`

		// File system ID
		FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

		// Custom file system name
		FsName *string `json:"FsName,omitempty" name:"FsName"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateCfsFileSystemNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateCfsFileSystemNameResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateCfsFileSystemPGroupRequest struct {
	*tchttp.BaseRequest

	// Permission group ID
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

func (r *UpdateCfsFileSystemPGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateCfsFileSystemPGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateCfsFileSystemPGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Permission group ID
		PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

		// File system ID
		FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateCfsFileSystemPGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateCfsFileSystemPGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateCfsFileSystemSizeLimitRequest struct {
	*tchttp.BaseRequest

	// File system capacity limit in GB. Value range: 0–1,073,741,824. If 0 is entered, no limit will be imposed on the file system capacity.
	FsLimit *uint64 `json:"FsLimit,omitempty" name:"FsLimit"`

	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

func (r *UpdateCfsFileSystemSizeLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateCfsFileSystemSizeLimitRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateCfsFileSystemSizeLimitResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateCfsFileSystemSizeLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateCfsFileSystemSizeLimitResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateCfsPGroupRequest struct {
	*tchttp.BaseRequest

	// Permission group ID
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// Permission group name, which can contain 1–64 Chinese characters, letters, numbers, underscores, or dashes
	Name *string `json:"Name,omitempty" name:"Name"`

	// Permission group description, which can contain 1–255 characters
	DescInfo *string `json:"DescInfo,omitempty" name:"DescInfo"`
}

func (r *UpdateCfsPGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateCfsPGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateCfsPGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Permission group ID
		PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

		// Permission group name
		Name *string `json:"Name,omitempty" name:"Name"`

		// Description
		DescInfo *string `json:"DescInfo,omitempty" name:"DescInfo"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateCfsPGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateCfsPGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateCfsRuleRequest struct {
	*tchttp.BaseRequest

	// Permission group ID
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// You can enter a single IP or IP range, such as 10.1.10.11 or 10.10.1.0/24. The default visiting address is `*`, indicating that all IPs are allowed. Please note that you need to enter the CVM instance's private IP here.
	AuthClientIp *string `json:"AuthClientIp,omitempty" name:"AuthClientIp"`

	// Read/write permission. Valid values: RO (read-only), RW (read & write). Default value: RO
	RWPermission *string `json:"RWPermission,omitempty" name:"RWPermission"`

	// User permission. Valid values: all_squash, no_all_squash, root_squash, no_root_squash. Specifically, all_squash: any visiting user will be mapped to an anonymous user or user group; no_all_squash: a visiting user will be first matched with a local user, and if the match fails, it will be mapped to an anonymous user or user group; root_squash: a visiting root user will be mapped to an anonymous user or user group; no_root_squash: a visiting root user will be allowed to maintain root account permissions. Default value: root_squash.
	UserPermission *string `json:"UserPermission,omitempty" name:"UserPermission"`

	// Rule priority. Value range: 1–100. 1 represents the highest priority, while 100 the lowest
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
}

func (r *UpdateCfsRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateCfsRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateCfsRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Permission group ID
		PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

		// Rule ID
		RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

		// Client IP or IP range allowed for access
		AuthClientIp *string `json:"AuthClientIp,omitempty" name:"AuthClientIp"`

		// Read & write permission
		RWPermission *string `json:"RWPermission,omitempty" name:"RWPermission"`

		// User permission
		UserPermission *string `json:"UserPermission,omitempty" name:"UserPermission"`

		// Priority
		Priority *int64 `json:"Priority,omitempty" name:"Priority"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateCfsRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateCfsRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
