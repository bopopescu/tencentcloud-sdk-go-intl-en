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

package v20180410

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AcceptDirectConnectTunnelRequest struct {
	*tchttp.BaseRequest

	// The connection owner accepts an application for sharing the dedicated tunnel
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`
}

func (r *AcceptDirectConnectTunnelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AcceptDirectConnectTunnelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AcceptDirectConnectTunnelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AcceptDirectConnectTunnelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AcceptDirectConnectTunnelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AccessPoint struct {

	// Access point name.
	AccessPointName *string `json:"AccessPointName,omitempty" name:"AccessPointName"`

	// Unique access point ID.
	AccessPointId *string `json:"AccessPointId,omitempty" name:"AccessPointId"`

	// Access point status. Valid values: available, unavailable.
	State *string `json:"State,omitempty" name:"State"`

	// Access point location.
	Location *string `json:"Location,omitempty" name:"Location"`

	// List of ISPs supported by access point.
	LineOperator []*string `json:"LineOperator,omitempty" name:"LineOperator" list`

	// ID of the region that manages the access point.
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

type BgpPeer struct {

	// User-side BGP Asn.
	Asn *int64 `json:"Asn,omitempty" name:"Asn"`

	// User-side BGP key.
	AuthKey *string `json:"AuthKey,omitempty" name:"AuthKey"`
}

type CreateDirectConnectRequest struct {
	*tchttp.BaseRequest

	// Connection name.
	DirectConnectName *string `json:"DirectConnectName,omitempty" name:"DirectConnectName"`

	// Access point of connection.
	// You can call `DescribeAccessPoints` to get the region ID. The selected access point must exist and be available.
	AccessPointId *string `json:"AccessPointId,omitempty" name:"AccessPointId"`

	// ISP that provides connections. Valid values: ChinaTelecom (China Telecom), ChinaMobile (China Mobile), ChinaUnicom (China Unicom), In-houseWiring (in-house wiring), ChinaOther (other Chinese ISPs), InternationalOperator (international ISPs).
	LineOperator *string `json:"LineOperator,omitempty" name:"LineOperator"`

	// Local IDC location.
	Location *string `json:"Location,omitempty" name:"Location"`

	// Port type of connection. Valid values: 100Base-T (100-Megabit electrical Ethernet interface), 1000Base-T (1-Gigabit electrical Ethernet interface), 1000Base-LX (1-Gigabit single-module optical Ethernet interface; 10 KM), 10GBase-T (10-Gigabit electrical Ethernet interface), 10GBase-LR (10-Gigabit single-module optical Ethernet interface; 10 KM). Default value: 1000Base-LX.
	PortType *string `json:"PortType,omitempty" name:"PortType"`

	// Circuit code of connection, which is provided by the ISP or connection provider.
	CircuitCode *string `json:"CircuitCode,omitempty" name:"CircuitCode"`

	// Connection port bandwidth in Mbps. Value range: [2,10240]. Default value: 1000.
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// ID of redundant connection.
	RedundantDirectConnectId *string `json:"RedundantDirectConnectId,omitempty" name:"RedundantDirectConnectId"`

	// VLAN for connection debugging, which is enabled and automatically assigned by default.
	Vlan *int64 `json:"Vlan,omitempty" name:"Vlan"`

	// Tencent-side IP address for connection debugging, which is automatically assigned by default.
	TencentAddress *string `json:"TencentAddress,omitempty" name:"TencentAddress"`

	// User-side IP address for connection debugging, which is automatically assigned by default.
	CustomerAddress *string `json:"CustomerAddress,omitempty" name:"CustomerAddress"`

	// Name of connection applicant, which is obtained from the account system by default.
	CustomerName *string `json:"CustomerName,omitempty" name:"CustomerName"`

	// Email address of connection applicant, which is obtained from the account system by default.
	CustomerContactMail *string `json:"CustomerContactMail,omitempty" name:"CustomerContactMail"`

	// Contact number of connection applicant, which is obtained from the account system by default.
	CustomerContactNumber *string `json:"CustomerContactNumber,omitempty" name:"CustomerContactNumber"`

	// Fault reporting contact person.
	FaultReportContactPerson *string `json:"FaultReportContactPerson,omitempty" name:"FaultReportContactPerson"`

	// Fault reporting contact number.
	FaultReportContactNumber *string `json:"FaultReportContactNumber,omitempty" name:"FaultReportContactNumber"`
}

func (r *CreateDirectConnectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDirectConnectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDirectConnectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Connection ID.
		DirectConnectIdSet []*string `json:"DirectConnectIdSet,omitempty" name:"DirectConnectIdSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDirectConnectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDirectConnectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDirectConnectTunnelRequest struct {
	*tchttp.BaseRequest

	// Direct Connect ID, such as `dc-kd7d06of`.
	DirectConnectId *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`

	// Dedicated tunnel name.
	DirectConnectTunnelName *string `json:"DirectConnectTunnelName,omitempty" name:"DirectConnectTunnelName"`

	// Connection owner, who is the current customer by default.
	// The developer account ID should be entered for shared connections.
	DirectConnectOwnerAccount *string `json:"DirectConnectOwnerAccount,omitempty" name:"DirectConnectOwnerAccount"`

	// Network type. Valid values: VPC, BMVPC, CCN. Default value: VPC.
	// VPC: Virtual Private Cloud.
	// BMVPC: BM VPC.
	// CCN: Cloud Connect Network.
	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`

	// Network region.
	NetworkRegion *string `json:"NetworkRegion,omitempty" name:"NetworkRegion"`

	// Unified VPC ID or BMVPC ID.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Direct connect gateway ID, such as `dcg-d545ddf`.
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`

	// Direct Connect bandwidth in Mbps.
	// Default value: connection bandwidth value.
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// BGP: BGP routing.
	// STATIC: Static routing.
	// Default value: BGP routing.
	RouteType *string `json:"RouteType,omitempty" name:"RouteType"`

	// BgpPeer, which is BGP information on the user side and includes Asn and AuthKey.
	BgpPeer *BgpPeer `json:"BgpPeer,omitempty" name:"BgpPeer"`

	// Static routing, i.e., IP range of the user's IDC.
	RouteFilterPrefixes []*RouteFilterPrefix `json:"RouteFilterPrefixes,omitempty" name:"RouteFilterPrefixes" list`

	// VLAN. Value range: 0-3,000.
	// 0: sub-interface not enabled.
	// Default value: Non-zero.
	Vlan *int64 `json:"Vlan,omitempty" name:"Vlan"`

	// TencentAddress: Tencent-side IP address.
	TencentAddress *string `json:"TencentAddress,omitempty" name:"TencentAddress"`

	// CustomerAddress: User-side IP address.
	CustomerAddress *string `json:"CustomerAddress,omitempty" name:"CustomerAddress"`

	// TencentBackupAddress, i.e., Tencent-side standby IP address
	TencentBackupAddress *string `json:"TencentBackupAddress,omitempty" name:"TencentBackupAddress"`
}

func (r *CreateDirectConnectTunnelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDirectConnectTunnelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDirectConnectTunnelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Dedicated tunnel ID.
		DirectConnectTunnelIdSet []*string `json:"DirectConnectTunnelIdSet,omitempty" name:"DirectConnectTunnelIdSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDirectConnectTunnelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDirectConnectTunnelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteDirectConnectRequest struct {
	*tchttp.BaseRequest

	// Connection ID.
	DirectConnectId *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`
}

func (r *DeleteDirectConnectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDirectConnectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteDirectConnectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDirectConnectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDirectConnectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteDirectConnectTunnelRequest struct {
	*tchttp.BaseRequest

	// Dedicated tunnel ID.
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`
}

func (r *DeleteDirectConnectTunnelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDirectConnectTunnelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteDirectConnectTunnelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDirectConnectTunnelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDirectConnectTunnelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessPointsRequest struct {
	*tchttp.BaseRequest

	// Access point region, which can be queried through `DescribeRegions`.
	// 
	// You can call `DescribeRegions` to get the region ID.
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAccessPointsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccessPointsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessPointsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Access point information.
		AccessPointSet []*AccessPoint `json:"AccessPointSet,omitempty" name:"AccessPointSet" list`

		// Number of eligible access points.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessPointsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccessPointsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectTunnelsRequest struct {
	*tchttp.BaseRequest

	// Filter conditions:
	// This parameter does not support specifying `DirectConnectTunnelIds` and `Filters` at the same time.
	// <li> direct-connect-tunnel-name: Dedicated tunnel name.</li>
	// <li> direct-connect-tunnel-id: Dedicated tunnel instance ID, such as `dcx-abcdefgh`.</li>
	// <li>direct-connect-id: Connection instance ID, such as `dc-abcdefgh`.</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Array of dedicated tunnel IDs.
	DirectConnectTunnelIds []*string `json:"DirectConnectTunnelIds,omitempty" name:"DirectConnectTunnelIds" list`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDirectConnectTunnelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDirectConnectTunnelsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectTunnelsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of dedicated tunnels.
		DirectConnectTunnelSet []*DirectConnectTunnel `json:"DirectConnectTunnelSet,omitempty" name:"DirectConnectTunnelSet" list`

		// Number of eligible dedicated tunnels.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDirectConnectTunnelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDirectConnectTunnelsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectsRequest struct {
	*tchttp.BaseRequest

	// Filter conditions:
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Array of connection IDs.
	DirectConnectIds []*string `json:"DirectConnectIds,omitempty" name:"DirectConnectIds" list`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDirectConnectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDirectConnectsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of connections.
		DirectConnectSet []*DirectConnect `json:"DirectConnectSet,omitempty" name:"DirectConnectSet" list`

		// Number of eligible connection lists.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDirectConnectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDirectConnectsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DirectConnect struct {

	// Connection ID.
	DirectConnectId *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`

	// Connection name.
	DirectConnectName *string `json:"DirectConnectName,omitempty" name:"DirectConnectName"`

	// Access point ID of a connection.
	AccessPointId *string `json:"AccessPointId,omitempty" name:"AccessPointId"`

	// Connection status.
	// PENDING: Applying. 
	// REJECTED: Application rejected.   
	// TOPAY: Payment pending. 
	// PAID: Paid. 
	// ALLOCATED: Constructing.   
	// AVAILABLE: Available.  
	// DELETING: Deleting.
	// DELETED: Deleted.
	State *string `json:"State,omitempty" name:"State"`

	// Connection creation time.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Connection activation time.
	EnabledTime *string `json:"EnabledTime,omitempty" name:"EnabledTime"`

	// ISP that provides connections. Valid values: ChinaTelecom (China Telecom), ChinaMobile (China Mobile), ChinaUnicom (China Unicom), In-houseWiring (in-house wiring), ChinaOther (other Chinese ISPs), InternationalOperator (international ISPs).
	LineOperator *string `json:"LineOperator,omitempty" name:"LineOperator"`

	// Location of a local IDC.
	Location *string `json:"Location,omitempty" name:"Location"`

	// Connection port bandwidth in Mbps.
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// User-side port type of a connection. Valid values: 100Base-T (100-Megabit electrical Ethernet interface), 1000Base-T (1-Gigabit electrical Ethernet interface; it is the default value), 1000Base-LX (1-Gigabit single-mode optical Ethernet interface; 10 KM), 10GBase-T (10-Gigabit electrical Ethernet interface), 10GBase-LR (10-Gigabit single-mode optical Ethernet interface; 10 KM).
	PortType *string `json:"PortType,omitempty" name:"PortType"`

	// Circuit code of a connection, which is provided by the ISP or service provider.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CircuitCode *string `json:"CircuitCode,omitempty" name:"CircuitCode"`

	// ID of a redundant connection.
	RedundantDirectConnectId *string `json:"RedundantDirectConnectId,omitempty" name:"RedundantDirectConnectId"`

	// VLAN for connection debugging, which is enabled and automatically assigned by default.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Vlan *int64 `json:"Vlan,omitempty" name:"Vlan"`

	// Tencent-side IP address for connection debugging.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TencentAddress *string `json:"TencentAddress,omitempty" name:"TencentAddress"`

	// User-side IP address for connection debugging.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CustomerAddress *string `json:"CustomerAddress,omitempty" name:"CustomerAddress"`

	// Name of the connection applicant, which is obtained from the account system by default.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CustomerName *string `json:"CustomerName,omitempty" name:"CustomerName"`

	// Email address of the connection applicant, which is obtained from the account system by default.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CustomerContactMail *string `json:"CustomerContactMail,omitempty" name:"CustomerContactMail"`

	// Contact number of the connection applicant, which is obtained from the account system by default.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CustomerContactNumber *string `json:"CustomerContactNumber,omitempty" name:"CustomerContactNumber"`

	// Connection expiration time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// Connection billing mode. NON_RECURRING_CHARGE: One-time charge for accessing service
	// Note: this field may return null, indicating that no valid values can be obtained.
	ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`

	// Fault reporting contact person.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FaultReportContactPerson *string `json:"FaultReportContactPerson,omitempty" name:"FaultReportContactPerson"`

	// Fault reporting contact number.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FaultReportContactNumber *string `json:"FaultReportContactNumber,omitempty" name:"FaultReportContactNumber"`

	// Tag key-value pair
	// Note: this field may return null, indicating that no valid values can be obtained.
	TagSet []*Tag `json:"TagSet,omitempty" name:"TagSet" list`

	// Access point type of a connection.
	AccessPointType *string `json:"AccessPointType,omitempty" name:"AccessPointType"`

	// IDC city.
	// Note: this field may return null, indicating that no valid values can be obtained.
	IdcCity *string `json:"IdcCity,omitempty" name:"IdcCity"`

	// Billing status
	// Note: this field may return null, indicating that no valid values can be obtained.
	ChargeState *string `json:"ChargeState,omitempty" name:"ChargeState"`

	// Connection activation time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

type DirectConnectTunnel struct {

	// Dedicated tunnel ID
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`

	// Connection ID.
	DirectConnectId *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`

	// Dedicated tunnel status
	// AVAILABLE: ready or connected
	// PENDING: applying
	// ALLOCATING: configuring
	// ALLOCATED: configured
	// ALTERING: modifying
	// DELETING: deleting
	// DELETED: deleted
	// CONFIRMING: to be accepted
	// REJECTED: rejected
	State *string `json:"State,omitempty" name:"State"`

	// Connection owner, i.e., developer account ID.
	DirectConnectOwnerAccount *string `json:"DirectConnectOwnerAccount,omitempty" name:"DirectConnectOwnerAccount"`

	// Dedicated tunnel owner, i.e., developer account ID
	OwnerAccount *string `json:"OwnerAccount,omitempty" name:"OwnerAccount"`

	// Network type. Valid values: VPC, BMVPC, CCN.
	//  VPC: Virtual Private Cloud; BMVPC: BM VPC; CCN: Cloud Connect Network.
	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`

	// Network of the VPC region, such as `ap-guangzhou`.
	NetworkRegion *string `json:"NetworkRegion,omitempty" name:"NetworkRegion"`

	// Unified VPC ID or BMVPC ID.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Direct connect gateway ID.
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`

	// BGP: BGP routing; STATIC: Static routing. Default value: BGP routing.
	RouteType *string `json:"RouteType,omitempty" name:"RouteType"`

	// User-side BGP, including Asn and AuthKey.
	BgpPeer *BgpPeer `json:"BgpPeer,omitempty" name:"BgpPeer"`

	// User-side IP range.
	RouteFilterPrefixes []*RouteFilterPrefix `json:"RouteFilterPrefixes,omitempty" name:"RouteFilterPrefixes" list`

	// Dedicated tunnel `Vlan`
	Vlan *int64 `json:"Vlan,omitempty" name:"Vlan"`

	// TencentAddress: Tencent-side IP address.
	TencentAddress *string `json:"TencentAddress,omitempty" name:"TencentAddress"`

	// CustomerAddress: User-side IP address.
	CustomerAddress *string `json:"CustomerAddress,omitempty" name:"CustomerAddress"`

	// Dedicated tunnel name
	DirectConnectTunnelName *string `json:"DirectConnectTunnelName,omitempty" name:"DirectConnectTunnelName"`

	// Dedicated tunnel creation time
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Dedicated tunnel bandwidth value
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// Dedicated tunnel tag value
	TagSet []*Tag `json:"TagSet,omitempty" name:"TagSet" list`

	// Associated custom network probe ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	NetDetectId *string `json:"NetDetectId,omitempty" name:"NetDetectId"`

	// BGP community switch
	// Note: this field may return null, indicating that no valid values can be obtained.
	EnableBGPCommunity *bool `json:"EnableBGPCommunity,omitempty" name:"EnableBGPCommunity"`

	// Whether it is a NAT tunnel
	// Note: this field may return null, indicating that no valid values can be obtained.
	NatType *int64 `json:"NatType,omitempty" name:"NatType"`

	// VPC region abbreviation, such as `gz`, `cd`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VpcRegion *string `json:"VpcRegion,omitempty" name:"VpcRegion"`

	// Whether to enable BFD
	// Note: this field may return null, indicating that no valid values can be obtained.
	BfdEnable *int64 `json:"BfdEnable,omitempty" name:"BfdEnable"`

	// Dedicated tunnel access point type
	// Note: this field may return null, indicating that no valid values can be obtained.
	AccessPointType *string `json:"AccessPointType,omitempty" name:"AccessPointType"`

	// Direct connect gateway name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DirectConnectGatewayName *string `json:"DirectConnectGatewayName,omitempty" name:"DirectConnectGatewayName"`

	// VPC name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// TencentBackupAddress, i.e., Tencent-side standby IP address
	// Note: this field may return null, indicating that no valid values can be obtained.
	TencentBackupAddress *string `json:"TencentBackupAddress,omitempty" name:"TencentBackupAddress"`
}

type Filter struct {

	// Fields to be filtered.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Filter values of the field.
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type ModifyDirectConnectAttributeRequest struct {
	*tchttp.BaseRequest

	// Connection ID.
	DirectConnectId *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`

	// Connection name.
	DirectConnectName *string `json:"DirectConnectName,omitempty" name:"DirectConnectName"`

	// Circuit code of connection, which is provided by the ISP or connection provider.
	CircuitCode *string `json:"CircuitCode,omitempty" name:"CircuitCode"`

	// VLAN for connection debugging.
	Vlan *int64 `json:"Vlan,omitempty" name:"Vlan"`

	// Tencent-side IP address for connection debugging.
	TencentAddress *string `json:"TencentAddress,omitempty" name:"TencentAddress"`

	// User-side IP address for connection debugging.
	CustomerAddress *string `json:"CustomerAddress,omitempty" name:"CustomerAddress"`

	// Name of connection applicant, which is obtained from the account system by default.
	CustomerName *string `json:"CustomerName,omitempty" name:"CustomerName"`

	// Email address of connection applicant, which is obtained from the account system by default.
	CustomerContactMail *string `json:"CustomerContactMail,omitempty" name:"CustomerContactMail"`

	// Contact number of connection applicant, which is obtained from the account system by default.
	CustomerContactNumber *string `json:"CustomerContactNumber,omitempty" name:"CustomerContactNumber"`

	// Fault reporting contact person.
	FaultReportContactPerson *string `json:"FaultReportContactPerson,omitempty" name:"FaultReportContactPerson"`

	// Fault reporting contact number.
	FaultReportContactNumber *string `json:"FaultReportContactNumber,omitempty" name:"FaultReportContactNumber"`
}

func (r *ModifyDirectConnectAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDirectConnectAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDirectConnectAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDirectConnectAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDirectConnectAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDirectConnectTunnelAttributeRequest struct {
	*tchttp.BaseRequest

	// Dedicated tunnel ID.
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`

	// Dedicated tunnel name.
	DirectConnectTunnelName *string `json:"DirectConnectTunnelName,omitempty" name:"DirectConnectTunnelName"`

	// User-side BGP, including Asn and AuthKey.
	BgpPeer *BgpPeer `json:"BgpPeer,omitempty" name:"BgpPeer"`

	// User-side IP range.
	RouteFilterPrefixes []*RouteFilterPrefix `json:"RouteFilterPrefixes,omitempty" name:"RouteFilterPrefixes" list`

	// Tencent-side IP address.
	TencentAddress *string `json:"TencentAddress,omitempty" name:"TencentAddress"`

	// User-side IP address.
	CustomerAddress *string `json:"CustomerAddress,omitempty" name:"CustomerAddress"`

	// Bandwidth value of a dedicated tunnel in Mbps.
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// Tencent-side standby IP address
	TencentBackupAddress *string `json:"TencentBackupAddress,omitempty" name:"TencentBackupAddress"`
}

func (r *ModifyDirectConnectTunnelAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDirectConnectTunnelAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDirectConnectTunnelAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDirectConnectTunnelAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDirectConnectTunnelAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RejectDirectConnectTunnelRequest struct {
	*tchttp.BaseRequest

	// None.
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`
}

func (r *RejectDirectConnectTunnelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RejectDirectConnectTunnelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RejectDirectConnectTunnelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RejectDirectConnectTunnelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RejectDirectConnectTunnelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RouteFilterPrefix struct {

	// User-side IP range.
	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`
}

type Tag struct {

	// Tag key
	// Note: this field may return null, indicating that no valid values can be obtained.
	Key *string `json:"Key,omitempty" name:"Key"`

	// Tag value
	// Note: this field may return null, indicating that no valid values can be obtained.
	Value *string `json:"Value,omitempty" name:"Value"`
}
