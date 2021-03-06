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

package v20180419

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type Activity struct {

	// Auto scaling group ID.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// Scaling activity ID.
	ActivityId *string `json:"ActivityId,omitempty" name:"ActivityId"`

	// Scaling activity type. Value range:<br>
	// <li>SCALE_OUT: scale-out <li>SCALE_IN: scale-in <li>ATTACH_INSTANCES: adding an instance <li>REMOVE_INSTANCES: terminating an instance <li>DETACH_INSTANCES: removing an instance <li>TERMINATE_INSTANCES_UNEXPECTEDLY: terminating an instance in the CVM console <li>REPLACE_UNHEALTHY_INSTANCE: replacing an unhealthy instance
	ActivityType *string `json:"ActivityType,omitempty" name:"ActivityType"`

	// Scaling activity status. Value range:<br>
	// <li>INIT: initializing
	// <li>RUNNING: running
	// <li>SUCCESSFUL: succeeded
	// <li>PARTIALLY_SUCCESSFUL: partially succeeded
	// <li>FAILED: failed
	// <li>CANCELLED: canceled
	StatusCode *string `json:"StatusCode,omitempty" name:"StatusCode"`

	// Description of the scaling activity status.
	StatusMessage *string `json:"StatusMessage,omitempty" name:"StatusMessage"`

	// Cause of the scaling activity.
	Cause *string `json:"Cause,omitempty" name:"Cause"`

	// Description of the scaling activity.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Start time of the scaling activity.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the scaling activity.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Creation time of the scaling activity.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Information set of the instances related to the scaling activity.
	ActivityRelatedInstanceSet []*ActivtyRelatedInstance `json:"ActivityRelatedInstanceSet,omitempty" name:"ActivityRelatedInstanceSet" list`

	// Brief description of the scaling activity status.
	StatusMessageSimplified *string `json:"StatusMessageSimplified,omitempty" name:"StatusMessageSimplified"`

	// Result of the lifecycle hook action in the scaling activity
	LifecycleActionResultSet []*LifecycleActionResultInfo `json:"LifecycleActionResultSet,omitempty" name:"LifecycleActionResultSet" list`
}

type ActivtyRelatedInstance struct {

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Status of the instance in the scaling activity. Value range:
	// <li>INIT: initializing
	// <li>RUNNING: running
	// <li>SUCCESSFUL: succeeded
	// <li>FAILED: failed
	InstanceStatus *string `json:"InstanceStatus,omitempty" name:"InstanceStatus"`
}

type AttachInstancesRequest struct {
	*tchttp.BaseRequest

	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// List of CVM instance IDs
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *AttachInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AttachInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Scaling activity ID
		ActivityId *string `json:"ActivityId,omitempty" name:"ActivityId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AttachInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AutoScalingGroup struct {

	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// Auto scaling group name
	AutoScalingGroupName *string `json:"AutoScalingGroupName,omitempty" name:"AutoScalingGroupName"`

	// Current status of the auto scaling group. Value range: <br><li>NORMAL: normal <br><li>CVM_ABNORMAL: Exception with the launch configuration <br><li>LB_ABNORMAL: exception with the load balancer <br><li>VPC_ABNORMAL: exception with the VPC <br><li>INSUFFICIENT_BALANCE: insufficient balance <br><li>LB_BACKEND_REGION_NOT_MATCH: the backend region of the CLB instance is not the same as the one of AS service.<br>
	AutoScalingGroupStatus *string `json:"AutoScalingGroupStatus,omitempty" name:"AutoScalingGroupStatus"`

	// Creation time in UTC format
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Default cooldown period in seconds
	DefaultCooldown *int64 `json:"DefaultCooldown,omitempty" name:"DefaultCooldown"`

	// Desired number of instances
	DesiredCapacity *int64 `json:"DesiredCapacity,omitempty" name:"DesiredCapacity"`

	// Enabled status. Value range: `ENABLED`, `DISABLED`
	EnabledStatus *string `json:"EnabledStatus,omitempty" name:"EnabledStatus"`

	// List of application load balancers
	ForwardLoadBalancerSet []*ForwardLoadBalancer `json:"ForwardLoadBalancerSet,omitempty" name:"ForwardLoadBalancerSet" list`

	// Number of instances
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// Number of instances in `IN_SERVICE` status
	InServiceInstanceCount *int64 `json:"InServiceInstanceCount,omitempty" name:"InServiceInstanceCount"`

	// Launch configuration ID
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitempty" name:"LaunchConfigurationId"`

	// Launch configuration name
	LaunchConfigurationName *string `json:"LaunchConfigurationName,omitempty" name:"LaunchConfigurationName"`

	// List of Classic load balancer IDs
	LoadBalancerIdSet []*string `json:"LoadBalancerIdSet,omitempty" name:"LoadBalancerIdSet" list`

	// Maximum number of instances
	MaxSize *int64 `json:"MaxSize,omitempty" name:"MaxSize"`

	// Minimum number of instances
	MinSize *int64 `json:"MinSize,omitempty" name:"MinSize"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// List of subnet IDs
	SubnetIdSet []*string `json:"SubnetIdSet,omitempty" name:"SubnetIdSet" list`

	// Termination policy
	TerminationPolicySet []*string `json:"TerminationPolicySet,omitempty" name:"TerminationPolicySet" list`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// List of availability zones
	ZoneSet []*string `json:"ZoneSet,omitempty" name:"ZoneSet" list`

	// Retry policy
	RetryPolicy *string `json:"RetryPolicy,omitempty" name:"RetryPolicy"`

	// Whether the auto scaling group is performing a scaling activity. `IN_ACTIVITY` indicates yes, and `NOT_IN_ACTIVITY` indicates no.
	InActivityStatus *string `json:"InActivityStatus,omitempty" name:"InActivityStatus"`

	// List of auto scaling group tags
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`

	// Service settings
	ServiceSettings *ServiceSettings `json:"ServiceSettings,omitempty" name:"ServiceSettings"`

	// 
	Ipv6AddressCount *int64 `json:"Ipv6AddressCount,omitempty" name:"Ipv6AddressCount"`

	// 
	MultiZoneSubnetPolicy *string `json:"MultiZoneSubnetPolicy,omitempty" name:"MultiZoneSubnetPolicy"`
}

type AutoScalingGroupAbstract struct {

	// Auto scaling group ID.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// Auto scaling group name.
	AutoScalingGroupName *string `json:"AutoScalingGroupName,omitempty" name:"AutoScalingGroupName"`
}

type AutoScalingNotification struct {

	// Auto scaling group ID.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// List of user group IDs.
	NotificationUserGroupIds []*string `json:"NotificationUserGroupIds,omitempty" name:"NotificationUserGroupIds" list`

	// List of notification events.
	NotificationTypes []*string `json:"NotificationTypes,omitempty" name:"NotificationTypes" list`

	// Event notification ID.
	AutoScalingNotificationId *string `json:"AutoScalingNotificationId,omitempty" name:"AutoScalingNotificationId"`
}

type CompleteLifecycleActionRequest struct {
	*tchttp.BaseRequest

	// Lifecycle hook ID
	LifecycleHookId *string `json:"LifecycleHookId,omitempty" name:"LifecycleHookId"`

	// Result of the lifecycle action. Value range: "CONTINUE", "ABANDON"
	LifecycleActionResult *string `json:"LifecycleActionResult,omitempty" name:"LifecycleActionResult"`

	// Instance ID. Either "InstanceId" or "LifecycleActionToken" must be specified
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Either "InstanceId" or "LifecycleActionToken" must be specified
	LifecycleActionToken *string `json:"LifecycleActionToken,omitempty" name:"LifecycleActionToken"`
}

func (r *CompleteLifecycleActionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CompleteLifecycleActionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CompleteLifecycleActionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CompleteLifecycleActionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CompleteLifecycleActionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAutoScalingGroupFromInstanceRequest struct {
	*tchttp.BaseRequest

	// The scaling group name. It must be unique under your account. The name can only contain Chinese characters, English letters, numbers, underscore, hyphen “-” and periods. It cannot exceed 55 bytes.
	AutoScalingGroupName *string `json:"AutoScalingGroupName,omitempty" name:"AutoScalingGroupName"`

	// The instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The maximum number of instances. Value range: 0-2000.
	MinSize *int64 `json:"MinSize,omitempty" name:"MinSize"`

	// The minimum number of instances. Value range: 0-2000.
	MaxSize *int64 `json:"MaxSize,omitempty" name:"MaxSize"`

	// The desired capacity. Its value must be greater than the minimum and smaller than the maximum.
	DesiredCapacity *int64 `json:"DesiredCapacity,omitempty" name:"DesiredCapacity"`

	// Whether to inherit the instance tag. Default value: False
	InheritInstanceTag *bool `json:"InheritInstanceTag,omitempty" name:"InheritInstanceTag"`
}

func (r *CreateAutoScalingGroupFromInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAutoScalingGroupFromInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAutoScalingGroupFromInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The scaling group ID.
		AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAutoScalingGroupFromInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAutoScalingGroupFromInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAutoScalingGroupRequest struct {
	*tchttp.BaseRequest

	// Auto scaling group name, which can only contain letters, numbers, underscores, hyphens ("-"), and decimal points with a maximum length of 55 bytes and must be unique under your account.
	AutoScalingGroupName *string `json:"AutoScalingGroupName,omitempty" name:"AutoScalingGroupName"`

	// Launch configuration ID
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitempty" name:"LaunchConfigurationId"`

	// Maximum number of instances. Value range: 0-2,000.
	MaxSize *uint64 `json:"MaxSize,omitempty" name:"MaxSize"`

	// Minimum number of instances. Value range: 0-2,000.
	MinSize *uint64 `json:"MinSize,omitempty" name:"MinSize"`

	// VPC ID; if on a basic network, enter an empty string
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Default cooldown period in seconds. Default value: 300
	DefaultCooldown *uint64 `json:"DefaultCooldown,omitempty" name:"DefaultCooldown"`

	// Desired number of instances. The number should be no larger than the maximum and no smaller than minimum number of instances
	DesiredCapacity *uint64 `json:"DesiredCapacity,omitempty" name:"DesiredCapacity"`

	// List of classic CLB IDs. Currently, the maximum length is 20. You cannot specify LoadBalancerIds and ForwardLoadBalancers at the same time.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds" list`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// List of CLBs. Currently, the maximum length is 20. You cannot specify LoadBalancerIds and ForwardLoadBalancers at the same time.
	ForwardLoadBalancers []*ForwardLoadBalancer `json:"ForwardLoadBalancers,omitempty" name:"ForwardLoadBalancers" list`

	// List of subnet IDs. A subnet must be specified in the VPC scenario. If multiple subnets are entered, their priority will be determined by the order in which they are entered, and they will be tried one by one until instances can be successfully created.
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds" list`

	// Termination policy. Currently, the maximum length is 1. Value range: OLDEST_INSTANCE, NEWEST_INSTANCE. Default value: OLDEST_INSTANCE.
	// <br><li> OLDEST_INSTANCE: The oldest instance in the auto scaling group will be terminated first.
	// <br><li> NEWEST_INSTANCE: The newest instance in the auto scaling group will be terminated first.
	TerminationPolicies []*string `json:"TerminationPolicies,omitempty" name:"TerminationPolicies" list`

	// List of availability zones. An availability zone must be specified in the basic network scenario. If multiple availability zones are entered, their priority will be determined by the order in which they are entered, and they will be tried one by one until instances can be successfully created.
	Zones []*string `json:"Zones,omitempty" name:"Zones" list`

	// Retry policy. Value range: IMMEDIATE_RETRY, INCREMENTAL_INTERVALS, and NO_RETRY. Default value: IMMEDIATE_RETRY.
	// <br><li> IMMEDIATE_RETRY: Retrying immediately in a short period of time and stopping after a number of consecutive failures (5).
	// <br><li> INCREMENTAL_INTERVALS: Retrying at incremental intervals, i.e., as the number of consecutive failures increases, the retry interval gradually increases, ranging from one second to one day.
	// <br><li> NO_RETRY: No retry until a user call or alarm message is received again.
	RetryPolicy *string `json:"RetryPolicy,omitempty" name:"RetryPolicy"`

	// Availability zone verification policy. Value range: ALL, ANY. Default value: ANY.
	// <br><li> ALL: The verification will succeed only if all availability zones (Zone) or subnets (SubnetId) are available; otherwise, an error will be reported.
	// <br><li> ANY: The verification will success if any availability zone (Zone) or subnet (SubnetId) is available; otherwise, an error will be reported.
	// 
	// Common reasons why an availability zone or subnet is unavailable include stock-out of CVM instances or CBS cloud disks in the availability zone, insufficient quota in the availability zone, or insufficient IPs in the subnet.
	// If an availability zone or subnet in Zones/SubnetIds does not exist, a verification error will be reported regardless of the value of ZonesCheckPolicy.
	ZonesCheckPolicy *string `json:"ZonesCheckPolicy,omitempty" name:"ZonesCheckPolicy"`

	// List of tag descriptions. This parameter is used to bind a tag to a scaling group as well as corresponding resource instances. Each scaling group can have up to 30 tags.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`

	// Service settings such as unhealthy instance replacement.
	ServiceSettings *ServiceSettings `json:"ServiceSettings,omitempty" name:"ServiceSettings"`

	// 
	Ipv6AddressCount *int64 `json:"Ipv6AddressCount,omitempty" name:"Ipv6AddressCount"`

	// 
	MultiZoneSubnetPolicy *string `json:"MultiZoneSubnetPolicy,omitempty" name:"MultiZoneSubnetPolicy"`
}

func (r *CreateAutoScalingGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAutoScalingGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAutoScalingGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Auto scaling group ID
		AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAutoScalingGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAutoScalingGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLaunchConfigurationRequest struct {
	*tchttp.BaseRequest

	// Display name of the launch configuration, which can contain Chinese characters, letters, numbers, underscores, separators ("-"), and decimal points with a maximum length of 60 bytes.
	LaunchConfigurationName *string `json:"LaunchConfigurationName,omitempty" name:"LaunchConfigurationName"`

	// Valid [image](https://cloud.tencent.com/document/product/213/4940) ID in the format of `img-8toqc6s3`. There are four types of images: <br/><li>Public images </li><li>Custom images </li><li>Shared images </li><li>Marketplace images </li><br/>You can obtain the available image IDs in the following ways: <br/><li>For `public images`, `custom images`, and `shared images`, log in to the [console](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE) to query the image IDs; for `marketplace images`, query the image IDs through [Cloud Marketplace](https://market.cloud.tencent.com/list). </li><li>This value can be obtained from the `ImageId` field in the return value of the [DescribeImages API](https://cloud.tencent.com/document/api/213/15715).</li>
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// ID of the project to which the instance belongs. This parameter can be obtained from the `projectId` field in the returned values of [DescribeProject](https://cloud.tencent.com/document/api/378/4400). If this is left empty, default project is used.
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Instance model. Different instance models specify different resource specifications. The specific value can be obtained by calling the [DescribeInstanceTypeConfigs](https://cloud.tencent.com/document/api/213/15749) API to get the latest specification table or referring to the descriptions in [Instance Types](https://cloud.tencent.com/document/product/213/11518).
	// `InstanceType` and `InstanceTypes` are mutually exclusive, and one and only one of them must be entered.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// System disk configuration of the instance. If this parameter is not specified, the default value will be assigned to it.
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// Information of the instance's data disk configuration. If this parameter is not specified, no data disk is purchased by default. Up to 11 data disks can be supported.
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks" list`

	// Configuration information of public network bandwidth. If this parameter is not specified, the default public network bandwidth is 0 Mbps.
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// Login settings of the instance. This parameter is used to set the login password and key for the instance, or to keep the original login settings for the image. By default, a random password is generated and sent to the user via the internal message.
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// The security group to which the instance belongs. This parameter can be obtained by calling the `SecurityGroupId` field in the returned value of [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808). If this parameter is not specified, no security group will be bound by default.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`

	// Enhanced service. This parameter is used to specify whether to enable Cloud Security, Cloud Monitoring and other services. If this parameter is not specified, Cloud Monitoring and Cloud Security will be enabled by default.
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// Base64-encoded custom data of up to 16 KB.
	UserData *string `json:"UserData,omitempty" name:"UserData"`

	// Instance billing type. CVM instances are POSTPAID_BY_HOUR by default.
	// <br><li>POSTPAID_BY_HOUR: Pay-as-you-go on an hourly basis
	// <br><li>SPOTPAID: Bidding
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// Market-related options of the instance, such as the parameters related to stop instances. If the billing method of instance is specified as bidding, this parameter must be passed in.
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitempty" name:"InstanceMarketOptions"`

	// List of instance models. Different instance models specify different resource specifications. Up to 10 instance models can be supported.
	// `InstanceType` and `InstanceTypes` are mutually exclusive, and one and only one of them must be entered.
	InstanceTypes []*string `json:"InstanceTypes,omitempty" name:"InstanceTypes" list`

	// Instance type verification policy. Value range: ALL, ANY. Default value: ANY.
	// <br><li> ALL: The verification will success only if all instance types (InstanceType) are available; otherwise, an error will be reported.
	// <br><li> ANY: The verification will success if any instance type (InstanceType) is available; otherwise, an error will be reported.
	// 
	// Common reasons why an instance type is unavailable include stock-out of the instance type or the corresponding cloud disk.
	// If a model in InstanceTypes does not exist or has been discontinued, a verification error will be reported regardless of the value of InstanceTypesCheckPolicy.
	InstanceTypesCheckPolicy *string `json:"InstanceTypesCheckPolicy,omitempty" name:"InstanceTypesCheckPolicy"`

	// List of tags. This parameter is used to bind up to 10 tags to newly added instances.
	InstanceTags []*InstanceTag `json:"InstanceTags,omitempty" name:"InstanceTags" list`

	// CAM role name, which can be obtained from the roleName field in the return value of the DescribeRoleList API.
	CamRoleName *string `json:"CamRoleName,omitempty" name:"CamRoleName"`

	// CVM HostName settings.
	HostNameSettings *HostNameSettings `json:"HostNameSettings,omitempty" name:"HostNameSettings"`
}

func (r *CreateLaunchConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLaunchConfigurationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLaunchConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// This parameter is returned when a launch configuration is created through this API, indicating the launch configuration ID.
		LaunchConfigurationId *string `json:"LaunchConfigurationId,omitempty" name:"LaunchConfigurationId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLaunchConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLaunchConfigurationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLifecycleHookRequest struct {
	*tchttp.BaseRequest

	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// Lifecycle hook name, which can contain Chinese characters, letters, numbers, underscores (_), hyphens (-), and periods (.) with a maximum length of 128 bytes.
	LifecycleHookName *string `json:"LifecycleHookName,omitempty" name:"LifecycleHookName"`

	// Scenario for the lifecycle hook. Valid values: "INSTANCE_LAUNCHING" and "INSTANCE_TERMINATING"
	LifecycleTransition *string `json:"LifecycleTransition,omitempty" name:"LifecycleTransition"`

	// Defined actions when lifecycle hook times out. Valid values: "CONTINUE" and "ABANDON". Default value: "CONTINUE"
	DefaultResult *string `json:"DefaultResult,omitempty" name:"DefaultResult"`

	// The maximum length of time (in seconds) that can elapse before the lifecycle hook times out. Value range: 30-3,600. Default value: 300
	HeartbeatTimeout *int64 `json:"HeartbeatTimeout,omitempty" name:"HeartbeatTimeout"`

	// Additional information sent by Auto Scaling to the notification target. Default value is “”. Maximum length is 1024 characters.
	NotificationMetadata *string `json:"NotificationMetadata,omitempty" name:"NotificationMetadata"`

	// Notification target
	NotificationTarget *NotificationTarget `json:"NotificationTarget,omitempty" name:"NotificationTarget"`

	// The scenario where the lifecycle hook is applied. `EXTENSION`: the lifecycle hook will be triggered when AttachInstances, DetachInstances or RemoveInstaces is called. `NORMAL`: the lifecycle hook is not triggered by the above APIs. 
	LifecycleTransitionType *string `json:"LifecycleTransitionType,omitempty" name:"LifecycleTransitionType"`
}

func (r *CreateLifecycleHookRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLifecycleHookRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLifecycleHookResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Lifecycle hook ID
		LifecycleHookId *string `json:"LifecycleHookId,omitempty" name:"LifecycleHookId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLifecycleHookResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLifecycleHookResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateNotificationConfigurationRequest struct {
	*tchttp.BaseRequest

	// Auto scaling group ID.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// Notification type, i.e., the set of types of notifications to be subscribed to. Value range:
	// <li>SCALE_OUT_SUCCESSFUL: scale-out succeeded</li>
	// <li>SCALE_OUT_FAILED: scale-out failed</li>
	// <li>SCALE_IN_SUCCESSFUL: scale-in succeeded</li>
	// <li>SCALE_IN_FAILED: scale-in failed</li>
	// <li>REPLACE_UNHEALTHY_INSTANCE_SUCCESSFUL: unhealthy instance replacement succeeded</li>
	// <li>REPLACE_UNHEALTHY_INSTANCE_FAILED: unhealthy instance replacement failed</li>
	NotificationTypes []*string `json:"NotificationTypes,omitempty" name:"NotificationTypes" list`

	// Notification group ID, which is the set of user group IDs. You can query the user group IDs through the [ListGroups](https://cloud.tencent.com/document/product/598/34589) API.
	NotificationUserGroupIds []*string `json:"NotificationUserGroupIds,omitempty" name:"NotificationUserGroupIds" list`
}

func (r *CreateNotificationConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateNotificationConfigurationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateNotificationConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Notification ID.
		AutoScalingNotificationId *string `json:"AutoScalingNotificationId,omitempty" name:"AutoScalingNotificationId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNotificationConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateNotificationConfigurationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePaiInstanceRequest struct {
	*tchttp.BaseRequest

	// PAI instance domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Information of the public network bandwidth configuration.
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// Base64-encoded string of the launch script.
	InitScript *string `json:"InitScript,omitempty" name:"InitScript"`

	// List of availability zones.
	Zones []*string `json:"Zones,omitempty" name:"Zones" list`

	// VpcId.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// List of subnets.
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds" list`

	// Instance display name.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// List of instance models.
	InstanceTypes []*string `json:"InstanceTypes,omitempty" name:"InstanceTypes" list`

	// Instance login settings.
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// Instance billing type.
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// Relevant parameter settings for the prepaid mode (i.e., monthly subscription). This parameter can specify the purchased usage period, whether to set automatic renewal, and other attributes of the instance purchased on a prepaid basis. If the billing method of the specified instance is prepaid, this parameter is required.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
}

func (r *CreatePaiInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePaiInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePaiInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// This parameter is returned when an instance is created via this API, representing one or more instance `IDs`. The return of the instance `ID` list does not mean that the instance is created successfully. You can find out whether the instance is created by checking the status of the instance `ID` in the InstancesSet returned by the [DescribeInstances API](https://cloud.tencent.com/document/api/213/15728). If the status of the instance changes from "pending" to "running", the instance is created successfully.
		InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePaiInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePaiInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateScalingPolicyRequest struct {
	*tchttp.BaseRequest

	// Auto scaling group ID.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// Alarm trigger policy name.
	ScalingPolicyName *string `json:"ScalingPolicyName,omitempty" name:"ScalingPolicyName"`

	// The method to adjust the desired number of instances after the alarm is triggered. Value range: <br><li>CHANGE_IN_CAPACITY: Increase or decrease the desired number of instances </li><li>EXACT_CAPACITY: Adjust to the specified desired number of instances </li> <li>PERCENT_CHANGE_IN_CAPACITY: Adjust the desired number of instances by percentage </li>
	AdjustmentType *string `json:"AdjustmentType,omitempty" name:"AdjustmentType"`

	// The adjusted value of desired number of instances after the alarm is triggered. Value range: <br><li>When AdjustmentType is CHANGE_IN_CAPACITY, if AdjustmentValue is a positive value, some new instances will be added after the alarm is triggered, and if it is a negative value, some existing instances will be removed after the alarm is triggered </li> <li> When AdjustmentType is EXACT_CAPACITY, the value of AdjustmentValue is the desired number of instances after the alarm is triggered, which should be equal to or greater than 0 </li> <li> When AdjustmentType is PERCENT_CHANGE_IN_CAPACITY, if AdjusmentValue (in %) is a positive value, new instances will be added by percentage after the alarm is triggered; if it is a negative value, existing instances will be removed by percentage after the alarm is triggered.
	AdjustmentValue *int64 `json:"AdjustmentValue,omitempty" name:"AdjustmentValue"`

	// Alarm monitoring metric.
	MetricAlarm *MetricAlarm `json:"MetricAlarm,omitempty" name:"MetricAlarm"`

	// Cooldown period in seconds. Default value: 300 seconds.
	Cooldown *uint64 `json:"Cooldown,omitempty" name:"Cooldown"`

	// Notification group ID, which is the set of user group IDs. You can query the user group IDs through the [ListGroups](https://cloud.tencent.com/document/product/598/34589) API.
	NotificationUserGroupIds []*string `json:"NotificationUserGroupIds,omitempty" name:"NotificationUserGroupIds" list`
}

func (r *CreateScalingPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateScalingPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateScalingPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Alarm trigger policy ID.
		AutoScalingPolicyId *string `json:"AutoScalingPolicyId,omitempty" name:"AutoScalingPolicyId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateScalingPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateScalingPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateScheduledActionRequest struct {
	*tchttp.BaseRequest

	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// Scheduled task name, which can only contain letters, numbers, underscores, hyphens ("-"), and decimal points with a maximum length of 60 bytes and must be unique in an auto scaling group.
	ScheduledActionName *string `json:"ScheduledActionName,omitempty" name:"ScheduledActionName"`

	// The maximum number of instances set for the auto scaling group when the scheduled task is triggered.
	MaxSize *uint64 `json:"MaxSize,omitempty" name:"MaxSize"`

	// The minimum number of instances set for the auto scaling group when the scheduled task is triggered.
	MinSize *uint64 `json:"MinSize,omitempty" name:"MinSize"`

	// The desired number of instances set for the auto scaling group when the scheduled task is triggered.
	DesiredCapacity *uint64 `json:"DesiredCapacity,omitempty" name:"DesiredCapacity"`

	// Initial triggered time of the scheduled task. The value is in `Beijing time` (UTC+8) in the format of `YYYY-MM-DDThh:mm:ss+08:00` according to the `ISO8601` standard.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the scheduled task. The value is in `Beijing time` (UTC+8) in the format of `YYYY-MM-DDThh:mm:ss+08:00` according to the `ISO8601` standard. <br><br>This parameter and `Recurrence` need to be specified at the same time. After the end time, the scheduled task will no longer take effect.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Repeating mode of the scheduled task, which is in standard cron format. <br><br>This parameter and `EndTime` need to be specified at the same time.
	Recurrence *string `json:"Recurrence,omitempty" name:"Recurrence"`
}

func (r *CreateScheduledActionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateScheduledActionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateScheduledActionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Scheduled task ID
		ScheduledActionId *string `json:"ScheduledActionId,omitempty" name:"ScheduledActionId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateScheduledActionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateScheduledActionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DataDisk struct {

	// Data disk type. For more information on limits of data disk types, see [CVM Instance Configuration](https://cloud.tencent.com/document/product/213/2177). Value range: <br><li>LOCAL_BASIC: Local disk <br><li>LOCAL_SSD: Local SSD disk <br><li>CLOUD_BASIC: HDD cloud disk <br><li>CLOUD_PREMIUM: Premium cloud disk <br><li>CLOUD_SSD: SSD cloud disk <br><br>Default value: LOCAL_BASIC.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// Data disk size (in GB). The minimum adjustment increment is 10 GB. The value range varies by data disk type. For more information on limits, see [CVM Instance Configuration](https://cloud.tencent.com/document/product/213/2177). The default value is 0, indicating that no data disk is purchased. For more information, see the product documentation.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// Data disk snapshot ID, such as `snap-l8psqwnt`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
}

type DeleteAutoScalingGroupRequest struct {
	*tchttp.BaseRequest

	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`
}

func (r *DeleteAutoScalingGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAutoScalingGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAutoScalingGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAutoScalingGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAutoScalingGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLaunchConfigurationRequest struct {
	*tchttp.BaseRequest

	// ID of the launch configuration to be deleted.
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitempty" name:"LaunchConfigurationId"`
}

func (r *DeleteLaunchConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLaunchConfigurationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLaunchConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLaunchConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLaunchConfigurationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLifecycleHookRequest struct {
	*tchttp.BaseRequest

	// Lifecycle hook ID
	LifecycleHookId *string `json:"LifecycleHookId,omitempty" name:"LifecycleHookId"`
}

func (r *DeleteLifecycleHookRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLifecycleHookRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLifecycleHookResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLifecycleHookResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLifecycleHookResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNotificationConfigurationRequest struct {
	*tchttp.BaseRequest

	// ID of the notification to be deleted.
	AutoScalingNotificationId *string `json:"AutoScalingNotificationId,omitempty" name:"AutoScalingNotificationId"`
}

func (r *DeleteNotificationConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteNotificationConfigurationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNotificationConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNotificationConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteNotificationConfigurationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteScalingPolicyRequest struct {
	*tchttp.BaseRequest

	// ID of the alarm policy to be deleted.
	AutoScalingPolicyId *string `json:"AutoScalingPolicyId,omitempty" name:"AutoScalingPolicyId"`
}

func (r *DeleteScalingPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteScalingPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteScalingPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteScalingPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteScalingPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteScheduledActionRequest struct {
	*tchttp.BaseRequest

	// ID of the scheduled task to be deleted.
	ScheduledActionId *string `json:"ScheduledActionId,omitempty" name:"ScheduledActionId"`
}

func (r *DeleteScheduledActionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteScheduledActionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteScheduledActionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteScheduledActionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteScheduledActionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountLimitsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAccountLimitsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccountLimitsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountLimitsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Maximum number of launch configurations allowed for creation by the user account
		MaxNumberOfLaunchConfigurations *int64 `json:"MaxNumberOfLaunchConfigurations,omitempty" name:"MaxNumberOfLaunchConfigurations"`

		// Current number of launch configurations under the user account
		NumberOfLaunchConfigurations *int64 `json:"NumberOfLaunchConfigurations,omitempty" name:"NumberOfLaunchConfigurations"`

		// Maximum number of auto scaling groups allowed for creation by the user account
		MaxNumberOfAutoScalingGroups *int64 `json:"MaxNumberOfAutoScalingGroups,omitempty" name:"MaxNumberOfAutoScalingGroups"`

		// Current number of auto scaling groups under the user account
		NumberOfAutoScalingGroups *int64 `json:"NumberOfAutoScalingGroups,omitempty" name:"NumberOfAutoScalingGroups"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountLimitsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccountLimitsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoScalingActivitiesRequest struct {
	*tchttp.BaseRequest

	// Queries by one or more scaling activity IDs in the format of `asa-5l2ejpfo`. The maximum quantity per request is 100. This parameter does not support specifying both `ActivityIds` and `Filters` at the same time.
	ActivityIds []*string `json:"ActivityIds,omitempty" name:"ActivityIds" list`

	// Filter.
	// <li> auto-scaling-group-id - String - Required: No - (Filter) Filter by auto scaling group ID.</li>
	// <li> activity-status-code - String - Required: No - (Filter) Filter by scaling activity status . (INIT: initializing | RUNNING: running | SUCCESSFUL: succeeded | PARTIALLY_SUCCESSFUL: partially succeeded | FAILED: failed | CANCELLED: canceled)</li>
	// <li> activity-type - String - Required: No - (Filter) Filter by scaling activity type. (SCALE_OUT: scale-out | SCALE_IN: scale-in | ATTACH_INSTANCES: adding an instance | REMOVE_INSTANCES: terminating an instance | DETACH_INSTANCES: removing an instance | TERMINATE_INSTANCES_UNEXPECTEDLY: terminating an instance in the CVM console | REPLACE_UNHEALTHY_INSTANCE: replacing an unhealthy instance | UPDATE_LOAD_BALANCERS: updating a load balancer)</li>
	// <li> activity-id - String - Required: No - (Filter) Filter by scaling activity ID.</li>
	// The maximum number of `Filters` per request is 10. The upper limit for `Filter.Values` is 5. This parameter does not support specifying both `ActivityIds` and `Filters` at the same time.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant section in the API [overview](https://cloud.tencent.com/document/api/213/15688).
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0. For more information on `Offset`, see the relevant section in the API [overview](https://cloud.tencent.com/document/api/213/15688).
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// The earliest start time of the scaling activity, which will be ignored if ActivityIds is specified. The value is in `UTC time` in the format of `YYYY-MM-DDThh:mm:ssZ` according to the `ISO8601` standard.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The latest end time of the scaling activity, which will be ignored if ActivityIds is specified. The value is in `UTC time` in the format of `YYYY-MM-DDThh:mm:ssZ` according to the `ISO8601` standard.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeAutoScalingActivitiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAutoScalingActivitiesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoScalingActivitiesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible scaling activities.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Information set of eligible scaling activities.
		ActivitySet []*Activity `json:"ActivitySet,omitempty" name:"ActivitySet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAutoScalingActivitiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAutoScalingActivitiesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoScalingGroupLastActivitiesRequest struct {
	*tchttp.BaseRequest

	// ID list of an auto scaling group.
	AutoScalingGroupIds []*string `json:"AutoScalingGroupIds,omitempty" name:"AutoScalingGroupIds" list`
}

func (r *DescribeAutoScalingGroupLastActivitiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAutoScalingGroupLastActivitiesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoScalingGroupLastActivitiesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Information set of eligible scaling activities. Scaling groups without scaling activities are not returned. For example, if there are 50 auto scaling group IDs but only 45 records are returned, it indicates that 5 of the auto scaling groups do not have scaling activities.
		ActivitySet []*Activity `json:"ActivitySet,omitempty" name:"ActivitySet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAutoScalingGroupLastActivitiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAutoScalingGroupLastActivitiesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoScalingGroupsRequest struct {
	*tchttp.BaseRequest

	// Queries by one or more auto scaling group IDs in the format of `asg-nkdwoui0`. The maximum quantity per request is 100. This parameter does not support specifying both `AutoScalingGroupIds` and `Filters` at the same time.
	AutoScalingGroupIds []*string `json:"AutoScalingGroupIds,omitempty" name:"AutoScalingGroupIds" list`

	// Filter.
	// <li> auto-scaling-group-id - String - Required: No - (Filter) Filter by auto scaling group ID.</li>
	// <li> auto-scaling-group-name - String - Required: No - (Filter) Filter by auto scaling group name.</li>
	// <li> launch-configuration-id - String - Required: No - (Filter) Filter by launch configuration ID.</li>
	// <li> tag-key - String - Required: No - (Filter) Filter by tag key.</li>
	// <li> tag-value - String - Required: No - (Filter) Filter by tag value.</li>
	// <li> tag:tag-key - String - Required: No - (Filter) Filter by tag key-value pair. The tag-key should be replaced with a specified tag key. For detailed usage, see sample 2</li>
	// The maximum number of `Filters` per request is 10. The upper limit for `Filter.Values` is 5. This parameter does not support specifying both `AutoScalingGroupIds` and `Filters` at the same time.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant section in the API [overview](https://cloud.tencent.com/document/api/213/15688).
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0. For more information on `Offset`, see the relevant section in the API [overview](https://cloud.tencent.com/document/api/213/15688).
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeAutoScalingGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAutoScalingGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoScalingGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of auto scaling group details.
		AutoScalingGroupSet []*AutoScalingGroup `json:"AutoScalingGroupSet,omitempty" name:"AutoScalingGroupSet" list`

		// Number of eligible auto scaling groups.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAutoScalingGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAutoScalingGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoScalingInstancesRequest struct {
	*tchttp.BaseRequest

	// ID of the CVM instance to be queried. This parameter does not support specifying both InstanceIds and Filters at the same time.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// Filter.
	// <li> instance-id - String - Required: No - (Filter) Filter by instance ID.</li>
	// <li> auto-scaling-group-id - String - Required: No - (Filter) Filter by auto scaling group ID.</li>
	// The maximum number of `Filters` per request is 10. The upper limit for `Filter.Values` is 5. This parameter does not support specifying both `InstanceIds` and `Filters` at the same time.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Offset. Default value: 0. For more information on `Offset`, see the relevant section in the API [overview](https://cloud.tencent.com/document/api/213/15688).
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant section in the API [overview](https://cloud.tencent.com/document/api/213/15688).
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAutoScalingInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAutoScalingInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoScalingInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of instance details.
		AutoScalingInstanceSet []*Instance `json:"AutoScalingInstanceSet,omitempty" name:"AutoScalingInstanceSet" list`

		// Number of eligible instances.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAutoScalingInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAutoScalingInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLaunchConfigurationsRequest struct {
	*tchttp.BaseRequest

	// Queries by one or more launch configuration IDs in the format of `asc-ouy1ax38`. The maximum quantity per request is 100. This parameter does not support specifying both `LaunchConfigurationIds` and `Filters` at the same time.
	LaunchConfigurationIds []*string `json:"LaunchConfigurationIds,omitempty" name:"LaunchConfigurationIds" list`

	// Filter.
	// <li> launch-configuration-id - String - Required: No - (Filter) Filter by launch configuration ID.</li>
	// <li> launch-configuration-name - String - Required: No - (Filter) Filter by launch configuration name.</li>
	// The maximum number of `Filters` per request is 10. The upper limit for `Filter.Values` is 5. This parameter does not support specifying both `LaunchConfigurationIds` and `Filters` at the same time.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant section in the API [overview](https://cloud.tencent.com/document/api/213/15688).
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0. For more information on `Offset`, see the relevant section in the API [overview](https://cloud.tencent.com/document/api/213/15688).
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeLaunchConfigurationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLaunchConfigurationsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLaunchConfigurationsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible launch configurations.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of launch configuration details.
		LaunchConfigurationSet []*LaunchConfiguration `json:"LaunchConfigurationSet,omitempty" name:"LaunchConfigurationSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLaunchConfigurationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLaunchConfigurationsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLifecycleHooksRequest struct {
	*tchttp.BaseRequest

	// Queries by one or more lifecycle hook IDs in the format of `ash-8azjzxcl`. The maximum quantity per request is 100. This parameter does not support specifying both `LifecycleHookIds` and `Filters` at the same time.
	LifecycleHookIds []*string `json:"LifecycleHookIds,omitempty" name:"LifecycleHookIds" list`

	// Filter.
	// <li> lifecycle-hook-id - String - Required: No - (Filter) Filter by lifecycle hook ID.</li>
	// <li> lifecycle-hook-name - String - Required: No - (Filter) Filter by lifecycle hook name.</li>
	// <li> auto-scaling-group-id - String - Required: No - (Filter) Filter by auto scaling group ID.</li>
	// Filter.
	// <li> lifecycle-hook-id - String - Required: No - (Filter) Filter by lifecycle hook ID.</li>
	// <li> lifecycle-hook-name - String - Required: No - (Filter) Filter by lifecycle hook name.</li>
	// <li> auto-scaling-group-id - String - Required: No - (Filter) Filter by auto scaling group ID.</li>
	// The maximum number of `Filters` per request is 10. The upper limit for `Filter.Values` is 5. This parameter does not support specifying both `LifecycleHookIds` and `Filters` at the same time.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant section in the API [overview](https://cloud.tencent.com/document/api/213/15688).
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0. For more information on `Offset`, see the relevant section in the API [overview](https://cloud.tencent.com/document/api/213/15688).
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeLifecycleHooksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLifecycleHooksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLifecycleHooksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Array of lifecycle hooks
		LifecycleHookSet []*LifecycleHook `json:"LifecycleHookSet,omitempty" name:"LifecycleHookSet" list`

		// Total quantity
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLifecycleHooksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLifecycleHooksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNotificationConfigurationsRequest struct {
	*tchttp.BaseRequest

	// Queries by one or more notification IDs in the format of asn-2sestqbr. The maximum number of instances per request is 100. This parameter does not support specifying both `AutoScalingNotificationIds` and `Filters` at the same time.
	AutoScalingNotificationIds []*string `json:"AutoScalingNotificationIds,omitempty" name:"AutoScalingNotificationIds" list`

	// Filter.
	// <li> auto-scaling-notification-id - String - Required: No - (Filter) Filter by notification ID.</li>
	// <li> auto-scaling-group-id - String - Required: No - (Filter) Filter by auto scaling group ID.</li>
	// The maximum number of `Filters` per request is 10. The upper limit for `Filter.Values` is 5. This parameter does not support specifying both `AutoScalingNotificationIds` and `Filters` at the same time.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant section in the API [overview](https://cloud.tencent.com/document/api/213/15688).
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0. For more information on `Offset`, see the relevant section in the API [overview](https://cloud.tencent.com/document/api/213/15688).
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeNotificationConfigurationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNotificationConfigurationsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNotificationConfigurationsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible notifications.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of AS event notification details.
		AutoScalingNotificationSet []*AutoScalingNotification `json:"AutoScalingNotificationSet,omitempty" name:"AutoScalingNotificationSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNotificationConfigurationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNotificationConfigurationsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePaiInstancesRequest struct {
	*tchttp.BaseRequest

	// Queries by PAI instance ID.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// Filter.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribePaiInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePaiInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePaiInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible PAI instances
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// PAI instance details
		PaiInstanceSet []*PaiInstance `json:"PaiInstanceSet,omitempty" name:"PaiInstanceSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePaiInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePaiInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeScalingPoliciesRequest struct {
	*tchttp.BaseRequest

	// Queries by one or more alarm policy IDs in the format of asp-i9vkg894. The maximum number of instances per request is 100. This parameter does not support specifying both `AutoScalingPolicyIds` and `Filters` at the same time.
	AutoScalingPolicyIds []*string `json:"AutoScalingPolicyIds,omitempty" name:"AutoScalingPolicyIds" list`

	// Filter.
	// <li> auto-scaling-policy-id - String - Required: No - (Filter) Filter by alarm policy ID.</li>
	// <li> auto-scaling-group-id - String - Required: No - (Filter) Filter by auto scaling group ID.</li>
	// <li> scaling-policy-name - String - Required: No - (Filter) Filter by alarm policy name.</li>
	// The maximum number of `Filters` per request is 10. The upper limit for `Filter.Values` is 5. This parameter does not support specifying both `AutoScalingPolicyIds` and `Filters` at the same time.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant section in the API [overview](https://cloud.tencent.com/document/api/213/15688).
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0. For more information on `Offset`, see the relevant section in the API [overview](https://cloud.tencent.com/document/api/213/15688).
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeScalingPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeScalingPoliciesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeScalingPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of AS alarm trigger policy details.
		ScalingPolicySet []*ScalingPolicy `json:"ScalingPolicySet,omitempty" name:"ScalingPolicySet" list`

		// Number of eligible notifications.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScalingPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeScalingPoliciesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeScheduledActionsRequest struct {
	*tchttp.BaseRequest

	// Queries by one or more scheduled task IDs in the format of asst-am691zxo. The maximum number of instances per request is 100. This parameter does not support specifying both ScheduledActionIds` and `Filters` at the same time.
	ScheduledActionIds []*string `json:"ScheduledActionIds,omitempty" name:"ScheduledActionIds" list`

	// Filter.
	// <li> scheduled-action-id - String - Required: No - (Filter) Filter by scheduled task ID.</li>
	// <li> scheduled-action-name - String - Required: No - (Filter) Filter by scheduled task name.</li>
	// <li> auto-scaling-group-id - String - Required: No - (Filter) Filter by auto scaling group ID.</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Offset. Default value: 0. For more information on `Offset`, see the relevant section in the API [overview](https://cloud.tencent.com/document/api/213/15688).
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant section in the API [overview](https://cloud.tencent.com/document/api/213/15688).
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeScheduledActionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeScheduledActionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeScheduledActionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible scheduled tasks.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of scheduled task details.
		ScheduledActionSet []*ScheduledAction `json:"ScheduledActionSet,omitempty" name:"ScheduledActionSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScheduledActionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeScheduledActionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetachInstancesRequest struct {
	*tchttp.BaseRequest

	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// List of CVM instance IDs
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *DetachInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetachInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Scaling activity ID
		ActivityId *string `json:"ActivityId,omitempty" name:"ActivityId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetachInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableAutoScalingGroupRequest struct {
	*tchttp.BaseRequest

	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`
}

func (r *DisableAutoScalingGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableAutoScalingGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableAutoScalingGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableAutoScalingGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableAutoScalingGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableAutoScalingGroupRequest struct {
	*tchttp.BaseRequest

	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`
}

func (r *EnableAutoScalingGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableAutoScalingGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableAutoScalingGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableAutoScalingGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableAutoScalingGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnhancedService struct {

	// Enables the Cloud Security service. If this parameter is not specified, the Cloud Security service will be enabled by default.
	SecurityService *RunSecurityServiceEnabled `json:"SecurityService,omitempty" name:"SecurityService"`

	// Enables the Cloud Monitor service. If this parameter is not specified, the Cloud Monitor service will be enabled by default.
	MonitorService *RunMonitorServiceEnabled `json:"MonitorService,omitempty" name:"MonitorService"`
}

type ExecuteScalingPolicyRequest struct {
	*tchttp.BaseRequest

	// Alarm-based scaling policy ID
	AutoScalingPolicyId *string `json:"AutoScalingPolicyId,omitempty" name:"AutoScalingPolicyId"`

	// Whether to check if the auto scaling group is in the cooldown period. Default value: false
	HonorCooldown *bool `json:"HonorCooldown,omitempty" name:"HonorCooldown"`

	// Trigger source that executes a scaling policy. Valid values: API and CLOUD_MONITOR. Default value: API. The value `CLOUD_MONITOR` is specific to the Cloud Monitor service.
	TriggerSource *string `json:"TriggerSource,omitempty" name:"TriggerSource"`
}

func (r *ExecuteScalingPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExecuteScalingPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExecuteScalingPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Scaling activity ID
		ActivityId *string `json:"ActivityId,omitempty" name:"ActivityId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExecuteScalingPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExecuteScalingPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// Field to be filtered.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Filter value of the field.
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type ForwardLoadBalancer struct {

	// Load balancer ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// Application load balancer listener ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// List of target rule attributes
	TargetAttributes []*TargetAttribute `json:"TargetAttributes,omitempty" name:"TargetAttributes" list`

	// ID of a forwarding rule. This parameter is required for layer-7 listeners.
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// The region of CLB instance. It defaults to the region of AS service and is in the format of the common parameter `Region`, such as `ap-guangzhou`.
	Region *string `json:"Region,omitempty" name:"Region"`
}

type HostNameSettings struct {

	// Host name of a CVM.
	// <br><li> A period (.) and hyphen (-) cannot be used as the first and the last characters of HostName, and multiple consecutive hyphens (-) or periods (.) are not allowed.
	// <br><li> No support for Windows instances.
	// <br><li> Other types of instances (such as Linux): the name should be a combination of 2 to 40 characters, supports multiple periods (.). The string between two periods can be composed of letters (case insensitive), numbers, and hyphens (-).
	// Note: This field may return null, indicating that no valid values can be obtained.
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// Type of CVM host name. Valid values: "ORIGINAL" and "UNIQUE". Default value: "ORIGINAL"
	// <br><li> ORIGINAL. Auto Scaling transfers the HostName set in input parameters to the CVM directly. CVM may adds serial numbers for the HostName. The HostName of instances within the auto scaling group may conflict.
	// <br><li> UNIQUE. The HostName set in input parameters is the prefix of a host name. Auto Scaling and CVM expand it. The HostName of an instance in the auto scaling group is unique.
	// Note: This field may return null, indicating that no valid values can be obtained.
	HostNameStyle *string `json:"HostNameStyle,omitempty" name:"HostNameStyle"`
}

type Instance struct {

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// Launch configuration ID
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitempty" name:"LaunchConfigurationId"`

	// Launch configuration name
	LaunchConfigurationName *string `json:"LaunchConfigurationName,omitempty" name:"LaunchConfigurationName"`

	// Lifecycle status. Value range: IN_SERVICE, CREATING, TERMINATING, ATTACHING, DETACHING, ATTACHING_LB, DETACHING_LB
	LifeCycleState *string `json:"LifeCycleState,omitempty" name:"LifeCycleState"`

	// Health status. Value range: HEALTHY, UNHEALTHY
	HealthStatus *string `json:"HealthStatus,omitempty" name:"HealthStatus"`

	// Whether to add scale-in protection
	ProtectedFromScaleIn *bool `json:"ProtectedFromScaleIn,omitempty" name:"ProtectedFromScaleIn"`

	// Availability zone
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Creation type. Value range: AUTO_CREATION, MANUAL_ATTACHING.
	CreationType *string `json:"CreationType,omitempty" name:"CreationType"`

	// Instance addition time
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// Instance type
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Version number
	VersionNumber *int64 `json:"VersionNumber,omitempty" name:"VersionNumber"`

	// Auto scaling group name
	AutoScalingGroupName *string `json:"AutoScalingGroupName,omitempty" name:"AutoScalingGroupName"`
}

type InstanceChargePrepaid struct {

	// Purchased usage period of an instance in months. Value range: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36.
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// Auto-renewal flag. Value range: <br><li>NOTIFY_AND_AUTO_RENEW: Notify expiry and renew automatically <br><li>NOTIFY_AND_MANUAL_RENEW: Notify expiry but not renew automatically <br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW: Neither notify expiry nor renew automatically <br><br>Default value: NOTIFY_AND_MANUAL_RENEW. If this parameter is specified as NOTIFY_AND_AUTO_RENEW, the instance will be automatically renewed on a monthly basis if the account balance is sufficient.
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type InstanceMarketOptionsRequest struct {

	// Bidding-related options
	SpotOptions *SpotMarketOptions `json:"SpotOptions,omitempty" name:"SpotOptions"`

	// Market option type. Currently, this only supports the value "spot"
	// Note: This field may return null, indicating that no valid values can be obtained.
	MarketType *string `json:"MarketType,omitempty" name:"MarketType"`
}

type InstanceTag struct {

	// Tag key
	Key *string `json:"Key,omitempty" name:"Key"`

	// Tag value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type InternetAccessible struct {

	// Network billing method. Value range: <br><li>BANDWIDTH_PREPAID: Prepaid by bandwidth <br><li>TRAFFIC_POSTPAID_BY_HOUR: Postpaid by traffic on a per hour basis <br><li>BANDWIDTH_POSTPAID_BY_HOUR: Postpaid by bandwidth on a per hour basis <br><li>BANDWIDTH_PACKAGE: BWP user <br>Default value: TRAFFIC_POSTPAID_BY_HOUR.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`

	// The maximum outbound bandwidth in Mbps of the public network. The default value is 0 Mbps. The upper limit of bandwidth varies by model. For more information, see [Purchase Network Bandwidth](https://cloud.tencent.com/document/product/213/509).
	// Note: This field may return null, indicating that no valid values can be obtained.
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// Whether to assign a public IP. Value range: <br><li>TRUE: Assign a public IP <br><li>FALSE: Do not assign a public IP <br><br>If the public network bandwidth is greater than 0 Mbps, you are free to choose whether to enable the public IP (which is enabled by default). If the public network bandwidth is 0 Mbps, no public IP will be allowed to be assigned.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PublicIpAssigned *bool `json:"PublicIpAssigned,omitempty" name:"PublicIpAssigned"`

	// Bandwidth package ID. You can obtain the parameter value from the `BandwidthPackageId` field in the response of the [DescribeBandwidthPackages](https://cloud.tencent.com/document/api/215/19209) API.
	// Note: this field may return null, indicating that no valid value was found.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" name:"BandwidthPackageId"`
}

type LaunchConfiguration struct {

	// Project ID of the instance.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Launch configuration ID.
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitempty" name:"LaunchConfigurationId"`

	// Launch configuration name.
	LaunchConfigurationName *string `json:"LaunchConfigurationName,omitempty" name:"LaunchConfigurationName"`

	// Instance model.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Information of the instance's system disk configuration.
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// Information of the instance's data disk configuration.
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks" list`

	// Instance login settings.
	LoginSettings *LimitedLoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// Information of the public network bandwidth configuration.
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// Security group of the instance.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`

	// Auto scaling group associated with the launch configuration.
	AutoScalingGroupAbstractSet []*AutoScalingGroupAbstract `json:"AutoScalingGroupAbstractSet,omitempty" name:"AutoScalingGroupAbstractSet" list`

	// Custom data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserData *string `json:"UserData,omitempty" name:"UserData"`

	// Creation time of the launch configuration.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Conditions of enhancement services for the instance and their settings.
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// Image ID.
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// Current status of the launch configuration. Value range: <br><li>NORMAL: normal <br><li>IMAGE_ABNORMAL: Exception with the image of the launch configuration <br><li>CBS_SNAP_ABNORMAL: Exception with the data disk snapshot of the launch configuration <br><li>SECURITY_GROUP_ABNORMAL: Exception with the security group of the launch configuration<br>
	LaunchConfigurationStatus *string `json:"LaunchConfigurationStatus,omitempty" name:"LaunchConfigurationStatus"`

	// Instance billing type. CVM instances are POSTPAID_BY_HOUR by default.
	// <br><li>POSTPAID_BY_HOUR: Pay-as-you-go on an hourly basis
	// <br><li>SPOTPAID: Bidding
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// Market-related options of the instance, such as the parameters related to stop instances. If the billing method of instance is specified as bidding, this parameter must be passed in.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitempty" name:"InstanceMarketOptions"`

	// List of instance models.
	InstanceTypes []*string `json:"InstanceTypes,omitempty" name:"InstanceTypes" list`

	// List of tags.
	InstanceTags []*InstanceTag `json:"InstanceTags,omitempty" name:"InstanceTags" list`

	// Version number.
	VersionNumber *int64 `json:"VersionNumber,omitempty" name:"VersionNumber"`

	// Update time.
	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`

	// CAM role name, which can be obtained from the roleName field in the return value of the DescribeRoleList API.
	CamRoleName *string `json:"CamRoleName,omitempty" name:"CamRoleName"`

	// Value of InstanceTypesCheckPolicy upon the last operation.
	LastOperationInstanceTypesCheckPolicy *string `json:"LastOperationInstanceTypesCheckPolicy,omitempty" name:"LastOperationInstanceTypesCheckPolicy"`

	// CVM HostName settings.
	HostNameSettings *HostNameSettings `json:"HostNameSettings,omitempty" name:"HostNameSettings"`
}

type LifecycleActionResultInfo struct {

	// ID of the lifecycle hook
	LifecycleHookId *string `json:"LifecycleHookId,omitempty" name:"LifecycleHookId"`

	// ID of the instance
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Whether the notification is sent to CMQ successfully
	NotificationResult *string `json:"NotificationResult,omitempty" name:"NotificationResult"`

	// Result of the lifecyle hook action. Values: CONTINUE, ABANDON
	LifecycleActionResult *string `json:"LifecycleActionResult,omitempty" name:"LifecycleActionResult"`

	// Cause of the result
	ResultReason *string `json:"ResultReason,omitempty" name:"ResultReason"`
}

type LifecycleHook struct {

	// Lifecycle hook ID
	LifecycleHookId *string `json:"LifecycleHookId,omitempty" name:"LifecycleHookId"`

	// Lifecycle hook name
	LifecycleHookName *string `json:"LifecycleHookName,omitempty" name:"LifecycleHookName"`

	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// Default result of the lifecycle hook
	DefaultResult *string `json:"DefaultResult,omitempty" name:"DefaultResult"`

	// Wait timeout period of the lifecycle hook
	HeartbeatTimeout *int64 `json:"HeartbeatTimeout,omitempty" name:"HeartbeatTimeout"`

	// Applicable scenario of the lifecycle hook
	LifecycleTransition *string `json:"LifecycleTransition,omitempty" name:"LifecycleTransition"`

	// Additional information for the notification target
	NotificationMetadata *string `json:"NotificationMetadata,omitempty" name:"NotificationMetadata"`

	// Creation time
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Notification target
	NotificationTarget *NotificationTarget `json:"NotificationTarget,omitempty" name:"NotificationTarget"`

	// Applicable scenario of the lifecycle hook
	LifecycleTransitionType *string `json:"LifecycleTransitionType,omitempty" name:"LifecycleTransitionType"`
}

type LimitedLoginSettings struct {

	// List of key IDs.
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds" list`
}

type LoginSettings struct {

	// Login password of the instance. The rule of password complexity varies by operating system: <br><li>For Linux instances, the password must be a combination of 8-16 characters comprised of at least two of the following types: [a-z, A-Z], [0-9] and [( ) ` ~ ! @ # $ % ^ & * - + = | { } [ ] : ; ' , . ? / ]. <br><li>For Windows instances, the password must be a combination of 12-16 characters comprised of at least three of the following types: [a-z], [A-Z], [0-9] and [( ) ` ~ ! @ # $ % ^ & * - + = { } [ ] : ; ' , . ? /]. <br><br>If this parameter is not specified, a password will be randomly generated and sent to you via internal message.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Password *string `json:"Password,omitempty" name:"Password"`

	// List of key IDs. An instance associated with the key can be accessed using the corresponding private key. KeyId can be obtained via the API DescribeKeyPairs. A key and a password cannot be specified at the same time, and specifying the key is not supported in Windows. You can specify only one key when purchasing an instance.
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds" list`

	// Keep the original settings for an image. You cannot specify this parameter if Password or KeyIds.N is specified. You can specify this parameter to TRUE only when you create an instance using a custom image, shared image, or image imported from external resources. Value range: <br><li>TRUE: Keep the login settings for the image <br><li>FALSE: Do not keep the login settings for the image <br><br>Default value: FALSE.
	// Note: This field may return null, indicating that no valid values can be obtained.
	KeepImageLogin *bool `json:"KeepImageLogin,omitempty" name:"KeepImageLogin"`
}

type MetricAlarm struct {

	// Comparison operator. Value range: <br><li>GREATER_THAN: greater than </li><li>GREATER_THAN_OR_EQUAL_TO: greater than or equal to </li><li>LESS_THAN: less than </li><li> LESS_THAN_OR_EQUAL_TO: less than or equal to </li><li> EQUAL_TO: equal to </li> <li>NOT_EQUAL_TO: not equal to </li>
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" name:"ComparisonOperator"`

	// Metric name. Value range: <br><li>CPU_UTILIZATION: CPU utilization </li><li>MEM_UTILIZATION: memory utilization </li><li>LAN_TRAFFIC_OUT: private network outbound bandwidth </li><li>LAN_TRAFFIC_IN: private network inbound bandwidth </li><li>WAN_TRAFFIC_OUT: public network outbound bandwidth </li><li>WAN_TRAFFIC_IN: public network inbound bandwidth </li>
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// Alarming threshold: <br><li>CPU_UTILIZATION: [1, 100] in % </li><li>MEM_UTILIZATION: [1, 100] in % </li><li>LAN_TRAFFIC_OUT: >0 in Mbps </li><li>LAN_TRAFFIC_IN: >0 in Mbps </li><li>WAN_TRAFFIC_OUT: >0 in Mbps </li><li>WAN_TRAFFIC_IN: >0 in Mbps </li>
	Threshold *uint64 `json:"Threshold,omitempty" name:"Threshold"`

	// Time period in seconds. Enumerated values: 60, 300.
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// Number of repetitions. Value range: [1, 10]
	ContinuousTime *uint64 `json:"ContinuousTime,omitempty" name:"ContinuousTime"`

	// Statistics type. Value range: <br><li>AVERAGE: average </li><li>MAXIMUM: maximum <li>MINIMUM: minimum </li><br> Default value: AVERAGE
	Statistic *string `json:"Statistic,omitempty" name:"Statistic"`
}

type ModifyAutoScalingGroupRequest struct {
	*tchttp.BaseRequest

	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// Auto scaling group name, which can only contain letters, numbers, underscores, hyphens ("-"), and decimal points with a maximum length of 55 bytes and must be unique under your account.
	AutoScalingGroupName *string `json:"AutoScalingGroupName,omitempty" name:"AutoScalingGroupName"`

	// Default cooldown period in seconds. Default value: 300
	DefaultCooldown *uint64 `json:"DefaultCooldown,omitempty" name:"DefaultCooldown"`

	// Desired number of instances. The number should be no larger than the maximum and no smaller than minimum number of instances
	DesiredCapacity *uint64 `json:"DesiredCapacity,omitempty" name:"DesiredCapacity"`

	// Launch configuration ID
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitempty" name:"LaunchConfigurationId"`

	// Maximum number of instances. Value range: 0-2,000.
	MaxSize *uint64 `json:"MaxSize,omitempty" name:"MaxSize"`

	// Minimum number of instances. Value range: 0-2,000.
	MinSize *uint64 `json:"MinSize,omitempty" name:"MinSize"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// List of subnet IDs
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds" list`

	// Termination policy. Currently, the maximum length is 1. Value range: OLDEST_INSTANCE, NEWEST_INSTANCE.
	// <br><li> OLDEST_INSTANCE: The oldest instance in the auto scaling group will be terminated first.
	// <br><li> NEWEST_INSTANCE: The newest instance in the auto scaling group will be terminated first.
	TerminationPolicies []*string `json:"TerminationPolicies,omitempty" name:"TerminationPolicies" list`

	// VPC ID. This field is left empty for basic networks. You need to specify SubnetIds when modifying the network of the auto scaling group to a VPC with a specified VPC ID. Specify Zones when modifying the network to a basic network.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// List of availability zones
	Zones []*string `json:"Zones,omitempty" name:"Zones" list`

	// Retry policy. Value range: IMMEDIATE_RETRY, INCREMENTAL_INTERVALS, and NO_RETRY. Default value: IMMEDIATE_RETRY.
	// <br><li> IMMEDIATE_RETRY: Retrying immediately in a short period of time and stopping after a number of consecutive failures (5).
	// <br><li> INCREMENTAL_INTERVALS: Retrying at incremental intervals, i.e., as the number of consecutive failures increases, the retry interval gradually increases, ranging from one second to one day.
	// <br><li> NO_RETRY: No retry until a user call or alarm message is received again.
	RetryPolicy *string `json:"RetryPolicy,omitempty" name:"RetryPolicy"`

	// Availability zone verification policy. Value range: ALL, ANY. Default value: ANY. This will work when the resource-related fields (launch configuration, availability zone, or subnet) of the auto scaling group are actually modified.
	// <br><li> ALL: The verification will succeed only if all availability zones (Zone) or subnets (SubnetId) are available; otherwise, an error will be reported.
	// <br><li> ANY: The verification will success if any availability zone (Zone) or subnet (SubnetId) is available; otherwise, an error will be reported.
	// 
	// Common reasons why an availability zone or subnet is unavailable include stock-out of CVM instances or CBS cloud disks in the availability zone, insufficient quota in the availability zone, or insufficient IPs in the subnet.
	// If an availability zone or subnet in Zones/SubnetIds does not exist, a verification error will be reported regardless of the value of ZonesCheckPolicy.
	ZonesCheckPolicy *string `json:"ZonesCheckPolicy,omitempty" name:"ZonesCheckPolicy"`

	// Service settings such as unhealthy instance replacement.
	ServiceSettings *ServiceSettings `json:"ServiceSettings,omitempty" name:"ServiceSettings"`

	// 
	Ipv6AddressCount *int64 `json:"Ipv6AddressCount,omitempty" name:"Ipv6AddressCount"`
}

func (r *ModifyAutoScalingGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAutoScalingGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoScalingGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAutoScalingGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAutoScalingGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDesiredCapacityRequest struct {
	*tchttp.BaseRequest

	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// Desired capacity
	DesiredCapacity *uint64 `json:"DesiredCapacity,omitempty" name:"DesiredCapacity"`
}

func (r *ModifyDesiredCapacityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDesiredCapacityRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDesiredCapacityResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDesiredCapacityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDesiredCapacityResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLaunchConfigurationAttributesRequest struct {
	*tchttp.BaseRequest

	// Launch configuration ID
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitempty" name:"LaunchConfigurationId"`

	// Valid [image](https://cloud.tencent.com/document/product/213/4940) ID in the format of `img-8toqc6s3`. There are four types of images: <br/><li>Public images </li><li>Custom images </li><li>Shared images </li><li>Marketplace images </li><br/>You can obtain the available image IDs in the following ways: <br/><li>For `public images`, `custom images`, and `shared images`, log in to the [console](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE) to query the image IDs; for `marketplace images`, query the image IDs through [Cloud Marketplace](https://market.cloud.tencent.com/list). </li><li>This value can be obtained from the `ImageId` field in the return value of the [DescribeImages API](https://cloud.tencent.com/document/api/213/15715).</li>
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// List of instance types. Different instance models specify different resource specifications. Up to 5 instance models can be supported.
	// The launch configuration uses InstanceType to indicate one single instance type and InstanceTypes to indicate multiple instance types. After InstanceTypes is successfully specified for the launch configuration, the original InstanceType will be automatically invalidated.
	InstanceTypes []*string `json:"InstanceTypes,omitempty" name:"InstanceTypes" list`

	// Instance type verification policy which works when InstanceTypes is actually modified. Value range: ALL, ANY. Default value: ANY.
	// <br><li> ALL: The verification will success only if all instance types (InstanceType) are available; otherwise, an error will be reported.
	// <br><li> ANY: The verification will success if any instance type (InstanceType) is available; otherwise, an error will be reported.
	// 
	// Common reasons why an instance type is unavailable include stock-out of the instance type or the corresponding cloud disk.
	// If a model in InstanceTypes does not exist or has been discontinued, a verification error will be reported regardless of the value of InstanceTypesCheckPolicy.
	InstanceTypesCheckPolicy *string `json:"InstanceTypesCheckPolicy,omitempty" name:"InstanceTypesCheckPolicy"`

	// Display name of the launch configuration, which can contain Chinese characters, letters, numbers, underscores, separators ("-"), and decimal points with a maximum length of 60 bytes.
	LaunchConfigurationName *string `json:"LaunchConfigurationName,omitempty" name:"LaunchConfigurationName"`

	// Base64-encoded custom data of up to 16 KB. If you want to clear UserData, specify it as an empty string
	UserData *string `json:"UserData,omitempty" name:"UserData"`
}

func (r *ModifyLaunchConfigurationAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLaunchConfigurationAttributesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLaunchConfigurationAttributesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLaunchConfigurationAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLaunchConfigurationAttributesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLoadBalancersRequest struct {
	*tchttp.BaseRequest

	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// List of classic CLB IDs. Currently, the maximum length is 20. You cannot specify LoadBalancerIds and ForwardLoadBalancers at the same time.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds" list`

	// List of CLBs. Currently, the maximum length is 20. You cannot specify LoadBalancerIds and ForwardLoadBalancers at the same time.
	ForwardLoadBalancers []*ForwardLoadBalancer `json:"ForwardLoadBalancers,omitempty" name:"ForwardLoadBalancers" list`

	// CLB verification policy. Valid values: "ALL" and "DIFF". Default value: "ALL"
	// <br><li> ALL. Verification is successful only when all CLBs are valid. Otherwise, verification fails.
	// <br><li> DIFF. Only the changes in the CLB parameters are verified. If valid, the verification is successful. Otherwise, verification fails.
	LoadBalancersCheckPolicy *string `json:"LoadBalancersCheckPolicy,omitempty" name:"LoadBalancersCheckPolicy"`
}

func (r *ModifyLoadBalancersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLoadBalancersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLoadBalancersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Scaling activity ID
		ActivityId *string `json:"ActivityId,omitempty" name:"ActivityId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLoadBalancersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLoadBalancersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNotificationConfigurationRequest struct {
	*tchttp.BaseRequest

	// ID of the notification to be modified.
	AutoScalingNotificationId *string `json:"AutoScalingNotificationId,omitempty" name:"AutoScalingNotificationId"`

	// Notification type, i.e., the set of types of notifications to be subscribed to. Value range:
	// <li>SCALE_OUT_SUCCESSFUL: scale-out succeeded</li>
	// <li>SCALE_OUT_FAILED: scale-out failed</li>
	// <li>SCALE_IN_SUCCESSFUL: scale-in succeeded</li>
	// <li>SCALE_IN_FAILED: scale-in failed</li>
	// <li>REPLACE_UNHEALTHY_INSTANCE_SUCCESSFUL: unhealthy instance replacement succeeded</li>
	// <li>REPLACE_UNHEALTHY_INSTANCE_FAILED: unhealthy instance replacement failed</li>
	NotificationTypes []*string `json:"NotificationTypes,omitempty" name:"NotificationTypes" list`

	// Notification group ID, which is the set of user group IDs. You can query the user group IDs through the [ListGroups](https://cloud.tencent.com/document/product/598/34589) API.
	NotificationUserGroupIds []*string `json:"NotificationUserGroupIds,omitempty" name:"NotificationUserGroupIds" list`
}

func (r *ModifyNotificationConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNotificationConfigurationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNotificationConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNotificationConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNotificationConfigurationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyScalingPolicyRequest struct {
	*tchttp.BaseRequest

	// Alarm policy ID.
	AutoScalingPolicyId *string `json:"AutoScalingPolicyId,omitempty" name:"AutoScalingPolicyId"`

	// Alarm policy name.
	ScalingPolicyName *string `json:"ScalingPolicyName,omitempty" name:"ScalingPolicyName"`

	// The method to adjust the desired number of instances after the alarm is triggered. Value range: <br><li>CHANGE_IN_CAPACITY: Increase or decrease the desired number of instances </li><li>EXACT_CAPACITY: Adjust to the specified desired number of instances </li> <li>PERCENT_CHANGE_IN_CAPACITY: Adjust the desired number of instances by percentage </li>
	AdjustmentType *string `json:"AdjustmentType,omitempty" name:"AdjustmentType"`

	// The adjusted value of desired number of instances after the alarm is triggered. Value range: <br><li>When AdjustmentType is CHANGE_IN_CAPACITY, if AdjustmentValue is a positive value, some new instances will be added after the alarm is triggered, and if it is a negative value, some existing instances will be removed after the alarm is triggered </li> <li> When AdjustmentType is EXACT_CAPACITY, the value of AdjustmentValue is the desired number of instances after the alarm is triggered, which should be equal to or greater than 0 </li> <li> When AdjustmentType is PERCENT_CHANGE_IN_CAPACITY, if AdjusmentValue (in %) is a positive value, new instances will be added by percentage after the alarm is triggered; if it is a negative value, existing instances will be removed by percentage after the alarm is triggered.
	AdjustmentValue *int64 `json:"AdjustmentValue,omitempty" name:"AdjustmentValue"`

	// Cooldown period in seconds.
	Cooldown *uint64 `json:"Cooldown,omitempty" name:"Cooldown"`

	// Alarm monitoring metric.
	MetricAlarm *MetricAlarm `json:"MetricAlarm,omitempty" name:"MetricAlarm"`

	// Notification group ID, which is the set of user group IDs. You can query the user group IDs through the [ListGroups](https://cloud.tencent.com/document/product/598/34589) API.
	// If you want to clear the user group, you need to pass in the specific string "NULL" to the list.
	NotificationUserGroupIds []*string `json:"NotificationUserGroupIds,omitempty" name:"NotificationUserGroupIds" list`
}

func (r *ModifyScalingPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyScalingPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyScalingPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyScalingPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyScalingPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyScheduledActionRequest struct {
	*tchttp.BaseRequest

	// ID of the scheduled task to be edited
	ScheduledActionId *string `json:"ScheduledActionId,omitempty" name:"ScheduledActionId"`

	// Scheduled task name, which can only contain letters, numbers, underscores, hyphens ("-"), and decimal points with a maximum length of 60 bytes and must be unique in an auto scaling group.
	ScheduledActionName *string `json:"ScheduledActionName,omitempty" name:"ScheduledActionName"`

	// The maximum number of instances set for the auto scaling group when the scheduled task is triggered.
	MaxSize *uint64 `json:"MaxSize,omitempty" name:"MaxSize"`

	// The minimum number of instances set for the auto scaling group when the scheduled task is triggered.
	MinSize *uint64 `json:"MinSize,omitempty" name:"MinSize"`

	// The desired number of instances set for the auto scaling group when the scheduled task is triggered.
	DesiredCapacity *uint64 `json:"DesiredCapacity,omitempty" name:"DesiredCapacity"`

	// Initial triggered time of the scheduled task. The value is in `Beijing time` (UTC+8) in the format of `YYYY-MM-DDThh:mm:ss+08:00` according to the `ISO8601` standard.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the scheduled task. The value is in `Beijing time` (UTC+8) in the format of `YYYY-MM-DDThh:mm:ss+08:00` according to the `ISO8601` standard. <br>This parameter and `Recurrence` need to be specified at the same time. After the end time, the scheduled task will no longer take effect.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Repeating mode of the scheduled task, which is in standard cron format. <br>This parameter and `EndTime` need to be specified at the same time.
	Recurrence *string `json:"Recurrence,omitempty" name:"Recurrence"`
}

func (r *ModifyScheduledActionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyScheduledActionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyScheduledActionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyScheduledActionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyScheduledActionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type NotificationTarget struct {

	// Target type. Value range: `CMQ_QUEUE`, `CMQ_TOPIC`.
	// <li> CMQ_QUEUE: CMQ_QUEUE: CMQ queue model.</li>
	// <li> CMQ_TOPIC: CMQ topic model.</li>
	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`

	// Queue name. If `TargetType` is `CMQ_QUEUE`, this parameter is required.
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`

	// Topic name. If `TargetType` is `CMQ_TOPIC`, this parameter is required.
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type PaiInstance struct {

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance domain name
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 
	PaiMateUrl *string `json:"PaiMateUrl,omitempty" name:"PaiMateUrl"`
}

type PreviewPaiDomainNameRequest struct {
	*tchttp.BaseRequest

	// Domain name type
	DomainNameType *string `json:"DomainNameType,omitempty" name:"DomainNameType"`
}

func (r *PreviewPaiDomainNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PreviewPaiDomainNameRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PreviewPaiDomainNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Available PAI domain name
		DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PreviewPaiDomainNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PreviewPaiDomainNameResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RemoveInstancesRequest struct {
	*tchttp.BaseRequest

	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// List of CVM instance IDs
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *RemoveInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RemoveInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RemoveInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Scaling activity ID
		ActivityId *string `json:"ActivityId,omitempty" name:"ActivityId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RemoveInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunMonitorServiceEnabled struct {

	// Whether to enable the [Cloud Monitor](https://cloud.tencent.com/document/product/248) service. Value range: <br><li>TRUE: Cloud Monitor is enabled <br><li>FALSE: Cloud Monitor is disabled <br><br>Default value: TRUE. |
	// Note: This field may return null, indicating that no valid values can be obtained.
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type RunSecurityServiceEnabled struct {

	// Whether to enable the [Cloud Security](https://cloud.tencent.com/document/product/296) service. Value range: <br><li>TRUE: Cloud Security is enabled <br><li>FALSE: Cloud Security is disabled <br><br>Default value: TRUE.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type ScalingPolicy struct {

	// Auto scaling group ID.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// Alarm trigger policy ID.
	AutoScalingPolicyId *string `json:"AutoScalingPolicyId,omitempty" name:"AutoScalingPolicyId"`

	// Alarm trigger policy name.
	ScalingPolicyName *string `json:"ScalingPolicyName,omitempty" name:"ScalingPolicyName"`

	// The method to adjust the desired number of instances after the alarm is triggered. Value range: <br><li>CHANGE_IN_CAPACITY: Increase or decrease the desired number of instances </li><li>EXACT_CAPACITY: Adjust to the specified desired number of instances </li> <li>PERCENT_CHANGE_IN_CAPACITY: Adjust the desired number of instances by percentage </li>
	AdjustmentType *string `json:"AdjustmentType,omitempty" name:"AdjustmentType"`

	// The adjusted value of desired number of instances after the alarm is triggered.
	AdjustmentValue *int64 `json:"AdjustmentValue,omitempty" name:"AdjustmentValue"`

	// Cooldown period.
	Cooldown *uint64 `json:"Cooldown,omitempty" name:"Cooldown"`

	// Alarm monitoring metric.
	MetricAlarm *MetricAlarm `json:"MetricAlarm,omitempty" name:"MetricAlarm"`

	// Notification group ID, which is the set of user group IDs.
	NotificationUserGroupIds []*string `json:"NotificationUserGroupIds,omitempty" name:"NotificationUserGroupIds" list`
}

type ScheduledAction struct {

	// Scheduled task ID.
	ScheduledActionId *string `json:"ScheduledActionId,omitempty" name:"ScheduledActionId"`

	// Scheduled task name.
	ScheduledActionName *string `json:"ScheduledActionName,omitempty" name:"ScheduledActionName"`

	// ID of the auto scaling group where the scheduled task is located.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// Start time of the scheduled task. The value is in `Beijing time` (UTC+8) in the format of `YYYY-MM-DDThh:mm:ss+08:00` according to the `ISO8601` standard.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Repeating mode of the scheduled task.
	Recurrence *string `json:"Recurrence,omitempty" name:"Recurrence"`

	// End time of the scheduled task. The value is in `Beijing time` (UTC+8) in the format of `YYYY-MM-DDThh:mm:ss+08:00` according to the `ISO8601` standard.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Maximum number of instances set by the scheduled task.
	MaxSize *uint64 `json:"MaxSize,omitempty" name:"MaxSize"`

	// Desired number of instances set by the scheduled task.
	DesiredCapacity *uint64 `json:"DesiredCapacity,omitempty" name:"DesiredCapacity"`

	// Minimum number of instances set by the scheduled task.
	MinSize *uint64 `json:"MinSize,omitempty" name:"MinSize"`

	// Creation time of the scheduled task. The value is in `UTC time` in the format of `YYYY-MM-DDThh:mm:ssZ` according to the `ISO8601` standard.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type ServiceSettings struct {

	// Enables unhealthy instance replacement. If this feature is enabled, AS will replace instances that are flagged as unhealthy by Cloud Monitor. If this parameter is not specified, the value will be False by default.
	ReplaceMonitorUnhealthy *bool `json:"ReplaceMonitorUnhealthy,omitempty" name:"ReplaceMonitorUnhealthy"`
}

type SetInstancesProtectionRequest struct {
	*tchttp.BaseRequest

	// Auto scaling group ID.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// Instance ID.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// Whether the instance needs to be protected from scale-in.
	ProtectedFromScaleIn *bool `json:"ProtectedFromScaleIn,omitempty" name:"ProtectedFromScaleIn"`
}

func (r *SetInstancesProtectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetInstancesProtectionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetInstancesProtectionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetInstancesProtectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetInstancesProtectionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SpotMarketOptions struct {

	// Bidding price such as "1.05"
	MaxPrice *string `json:"MaxPrice,omitempty" name:"MaxPrice"`

	// Bid request type. Currently, only "one-time" type is supported. Default value: one-time
	// Note: This field may return null, indicating that no valid values can be obtained.
	SpotInstanceType *string `json:"SpotInstanceType,omitempty" name:"SpotInstanceType"`
}

type StartAutoScalingInstancesRequest struct {
	*tchttp.BaseRequest

	// The scaling group ID.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// The list of the CVM instances you want to launch.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *StartAutoScalingInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartAutoScalingInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartAutoScalingInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The scaling activity ID.
		ActivityId *string `json:"ActivityId,omitempty" name:"ActivityId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartAutoScalingInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartAutoScalingInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopAutoScalingInstancesRequest struct {
	*tchttp.BaseRequest

	// The scaling group ID.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// The list of the CVM instances you want to shut down.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// Whether the shutdown instances will be charged. Valid values:  
	// KEEP_CHARGING: keep charging after shutdown.  
	// STOP_CHARGING: stop charging after shutdown.
	// Default value: KEEP_CHARGING.
	StoppedMode *string `json:"StoppedMode,omitempty" name:"StoppedMode"`
}

func (r *StopAutoScalingInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopAutoScalingInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopAutoScalingInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The scaling activity ID.
		ActivityId *string `json:"ActivityId,omitempty" name:"ActivityId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopAutoScalingInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopAutoScalingInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SystemDisk struct {

	// System disk type. For more information on limits of system disk types, see [CVM Instance Configuration](https://cloud.tencent.com/document/product/213/2177). Value range: <br><li>LOCAL_BASIC: Local disk <br><li>LOCAL_SSD: Local SSD disk <br><li>CLOUD_BASIC: HDD cloud disk <br><li>CLOUD_PREMIUM: Premium cloud disk <br><li>CLOUD_SSD: SSD cloud disk <br><br>Default value: LOCAL_BASIC.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// System disk size in GB. Default value: 50
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
}

type Tag struct {

	// Tag key
	Key *string `json:"Key,omitempty" name:"Key"`

	// Tag value
	Value *string `json:"Value,omitempty" name:"Value"`

	// Type of the resource binded to the tag. Currently supported types include "auto-scaling-group"
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
}

type TargetAttribute struct {

	// Port
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// Weight
	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`
}

type UpgradeLaunchConfigurationRequest struct {
	*tchttp.BaseRequest

	// Launch configuration ID.
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitempty" name:"LaunchConfigurationId"`

	// Valid [image](https://cloud.tencent.com/document/product/213/4940) ID in the format of `img-8toqc6s3`. There are four types of images: <br/><li>Public images </li><li>Custom images </li><li>Shared images </li><li>Marketplace images </li><br/>You can obtain the available image IDs in the following ways: <br/><li>For `public images`, `custom images`, and `shared images`, log in to the [console](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE) to query the image IDs; for `marketplace images`, query the image IDs through [Cloud Marketplace](https://market.cloud.tencent.com/list). </li><li>This value can be obtained from the `ImageId` field in the return value of the [DescribeImages API](https://cloud.tencent.com/document/api/213/15715).</li>
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// List of instance models. Different instance models specify different resource specifications. Up to 5 instance models can be supported.
	InstanceTypes []*string `json:"InstanceTypes,omitempty" name:"InstanceTypes" list`

	// Display name of the launch configuration, which can contain Chinese characters, letters, numbers, underscores, separators ("-"), and decimal points with a maximum length of 60 bytes.
	LaunchConfigurationName *string `json:"LaunchConfigurationName,omitempty" name:"LaunchConfigurationName"`

	// Information of the instance's data disk configuration. If this parameter is not specified, no data disk is purchased by default. Up to 11 data disks can be supported.
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks" list`

	// Enhanced service. This parameter is used to specify whether to enable Cloud Security, Cloud Monitoring and other services. If this parameter is not specified, Cloud Monitoring and Cloud Security will be enabled by default.
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// Instance billing type. CVM instances are POSTPAID_BY_HOUR by default.
	// <br><li>POSTPAID_BY_HOUR: Pay-as-you-go on an hourly basis
	// <br><li>SPOTPAID: Bidding
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// Market-related options of the instance, such as the parameters related to stop instances. If the billing method of instance is specified as bidding, this parameter must be passed in.
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitempty" name:"InstanceMarketOptions"`

	// Instance type verification policy. Value range: ALL, ANY. Default value: ANY.
	// <br><li> ALL: The verification will success only if all instance types (InstanceType) are available; otherwise, an error will be reported.
	// <br><li> ANY: The verification will success if any instance type (InstanceType) is available; otherwise, an error will be reported.
	// 
	// Common reasons why an instance type is unavailable include stock-out of the instance type and the corresponding cloud disk.
	// If a model in InstanceTypes does not exist or has been deactivated, a verification error will be reported regardless of the value of InstanceTypesCheckPolicy.
	InstanceTypesCheckPolicy *string `json:"InstanceTypesCheckPolicy,omitempty" name:"InstanceTypesCheckPolicy"`

	// Configuration information of public network bandwidth. If this parameter is not specified, the default public network bandwidth is 0 Mbps.
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// Login settings of the instance. This parameter is used to set the login password and key for the instance, or to keep the original login settings for the image. By default, a random password is generated and sent to the user via the internal message.
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// Project ID of the instance. This parameter can be obtained from the `projectId` field in the returned values of [DescribeProject](https://cloud.tencent.com/document/api/378/4400). If this is left empty, default project is used.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// The security group of instance. This parameter can be obtained by calling the `SecurityGroupId` field in the returned value of [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808). If this parameter is not specified, no security group will be bound by default.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`

	// System disk configuration of the instance. If this parameter is not specified, the default value will be assigned to it.
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// Base64-encoded custom data of up to 16 KB.
	UserData *string `json:"UserData,omitempty" name:"UserData"`

	// List of tags. This parameter is used to bind up to 10 tags to newly added instances.
	InstanceTags []*InstanceTag `json:"InstanceTags,omitempty" name:"InstanceTags" list`

	// CAM role name, which can be obtained from the roleName field in the return value of the DescribeRoleList API.
	CamRoleName *string `json:"CamRoleName,omitempty" name:"CamRoleName"`

	// CVM HostName settings.
	HostNameSettings *HostNameSettings `json:"HostNameSettings,omitempty" name:"HostNameSettings"`
}

func (r *UpgradeLaunchConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeLaunchConfigurationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeLaunchConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeLaunchConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeLaunchConfigurationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeLifecycleHookRequest struct {
	*tchttp.BaseRequest

	// Lifecycle hook ID
	LifecycleHookId *string `json:"LifecycleHookId,omitempty" name:"LifecycleHookId"`

	// Lifecycle hook name
	LifecycleHookName *string `json:"LifecycleHookName,omitempty" name:"LifecycleHookName"`

	// Scenario for the lifecycle hook. Value range: "INSTANCE_LAUNCHING", "INSTANCE_TERMINATING"
	LifecycleTransition *string `json:"LifecycleTransition,omitempty" name:"LifecycleTransition"`

	// Defines the action to be taken by the auto scaling group upon lifecycle hook timeout. Value range: "CONTINUE", "ABANDON". Default value: "CONTINUE"
	DefaultResult *string `json:"DefaultResult,omitempty" name:"DefaultResult"`

	// The maximum length of time (in seconds) that can elapse before the lifecycle hook times out. Value range: 30-3,600. Default value: 300
	HeartbeatTimeout *int64 `json:"HeartbeatTimeout,omitempty" name:"HeartbeatTimeout"`

	// Additional information sent by AS to the notification target. The default value is ''
	NotificationMetadata *string `json:"NotificationMetadata,omitempty" name:"NotificationMetadata"`

	// Notification target
	NotificationTarget *NotificationTarget `json:"NotificationTarget,omitempty" name:"NotificationTarget"`

	// The scenario where the lifecycle hook is applied. `EXTENSION`: the lifecycle hook will be triggered when AttachInstances, DetachInstances or RemoveInstaces is called. `NORMAL`: the lifecycle hook is not triggered by the above APIs. 
	LifecycleTransitionType *string `json:"LifecycleTransitionType,omitempty" name:"LifecycleTransitionType"`
}

func (r *UpgradeLifecycleHookRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeLifecycleHookRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeLifecycleHookResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeLifecycleHookResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeLifecycleHookResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
