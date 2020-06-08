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

package v20180808

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type ApiEnvironmentStrategy struct {

	// Unique API ID.
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// Custom API name.
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// API path, such as `/path`.
	Path *string `json:"Path,omitempty" name:"Path"`

	// API method, such as `GET`.
	Method *string `json:"Method,omitempty" name:"Method"`

	// Environment throttling information.
	EnvironmentStrategySet []*EnvironmentStrategy `json:"EnvironmentStrategySet,omitempty" name:"EnvironmentStrategySet" list`
}

type ApiEnvironmentStrategyStataus struct {

	// Number of throttling policies bound to API.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of throttling policies bound to API.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiEnvironmentStrategySet []*ApiEnvironmentStrategy `json:"ApiEnvironmentStrategySet,omitempty" name:"ApiEnvironmentStrategySet" list`
}

type ApiIdStatus struct {

	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique API ID.
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// API description
	ApiDesc *string `json:"ApiDesc,omitempty" name:"ApiDesc"`

	// API path.
	Path *string `json:"Path,omitempty" name:"Path"`

	// API method.
	Method *string `json:"Method,omitempty" name:"Method"`

	// Service creation time.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Service modification time.
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// API name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// Unique VPC ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// API type.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`

	// API protocol.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Whether to enable debugging on purchase.
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsDebugAfterCharge *bool `json:"IsDebugAfterCharge,omitempty" name:"IsDebugAfterCharge"`

	// Authorization type.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// API business type.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiBusinessType *string `json:"ApiBusinessType,omitempty" name:"ApiBusinessType"`

	// Unique ID of associated authorization API.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AuthRelationApiId *string `json:"AuthRelationApiId,omitempty" name:"AuthRelationApiId"`

	// List of business APIs associated with authorization API.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RelationBuniessApiIds []*string `json:"RelationBuniessApiIds,omitempty" name:"RelationBuniessApiIds" list`

	// OAuth configuration information.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OauthConfig *OauthConfig `json:"OauthConfig,omitempty" name:"OauthConfig"`

	// Token storage position, which is an OAuth 2.0 API request.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TokenLocation *string `json:"TokenLocation,omitempty" name:"TokenLocation"`
}

type ApiInfo struct {

	// Unique ID of API's service.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Name of API's service.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// Description of API's service.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`

	// Unique API ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// API description.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiDesc *string `json:"ApiDesc,omitempty" name:"ApiDesc"`

	// Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// API name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// API type. Valid values: NORMAL (general API), TSF (microservice API).
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`

	// API frontend request type, such as HTTP, HTTPS, or HTTP and HTTPS.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// API authentication type. Valid values: SECRET (key pair authentication), NONE (no authentication), OAUTH.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// OAuth API type. Valid values: NORMAL (business API), OAUTH (authorization API).
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiBusinessType *string `json:"ApiBusinessType,omitempty" name:"ApiBusinessType"`

	// Unique ID of the authorization API associated with OAuth business API.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AuthRelationApiId *string `json:"AuthRelationApiId,omitempty" name:"AuthRelationApiId"`

	// OAuth configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OauthConfig *OauthConfig `json:"OauthConfig,omitempty" name:"OauthConfig"`

	// Whether to enable debugging on purchase (reserved for the marketplace).
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsDebugAfterCharge *bool `json:"IsDebugAfterCharge,omitempty" name:"IsDebugAfterCharge"`

	// Request frontend configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RequestConfig *RequestConfig `json:"RequestConfig,omitempty" name:"RequestConfig"`

	// Return type.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ResponseType *string `json:"ResponseType,omitempty" name:"ResponseType"`

	// Successful response sample of custom response configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ResponseSuccessExample *string `json:"ResponseSuccessExample,omitempty" name:"ResponseSuccessExample"`

	// Response failure sample of custom response configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ResponseFailExample *string `json:"ResponseFailExample,omitempty" name:"ResponseFailExample"`

	// Custom error code configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ResponseErrorCodes []*ErrorCodes `json:"ResponseErrorCodes,omitempty" name:"ResponseErrorCodes" list`

	// Frontend request parameters.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RequestParameters []*ReqParameter `json:"RequestParameters,omitempty" name:"RequestParameters" list`

	// API backend service timeout period in seconds.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceTimeout *int64 `json:"ServiceTimeout,omitempty" name:"ServiceTimeout"`

	// API backend service type. Valid values: HTTP, MOCK, TSF, CLB, SCF, WEBSOCKET, TARGET (in beta test).
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// API backend service configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceConfig *ServiceConfig `json:"ServiceConfig,omitempty" name:"ServiceConfig"`

	// API backend service parameters.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceParameters []*ServiceParameter `json:"ServiceParameters,omitempty" name:"ServiceParameters" list`

	// Constant parameters.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ConstantParameters []*ConstantParameter `json:"ConstantParameters,omitempty" name:"ConstantParameters" list`

	// Returned information of API backend mocking. This parameter is required when `ServiceType` is `Mock`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceMockReturnMessage *string `json:"ServiceMockReturnMessage,omitempty" name:"ServiceMockReturnMessage"`

	// SCF function name. This parameter takes effect when the backend type is `SCF`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceScfFunctionName *string `json:"ServiceScfFunctionName,omitempty" name:"ServiceScfFunctionName"`

	// SCF function namespace. This parameter takes effect when the backend type is `SCF`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceScfFunctionNamespace *string `json:"ServiceScfFunctionNamespace,omitempty" name:"ServiceScfFunctionNamespace"`

	// SCF function version. This parameter takes effect when the backend type is `SCF`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceScfFunctionQualifier *string `json:"ServiceScfFunctionQualifier,omitempty" name:"ServiceScfFunctionQualifier"`

	// Whether to enable integrated response.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceScfIsIntegratedResponse *bool `json:"ServiceScfIsIntegratedResponse,omitempty" name:"ServiceScfIsIntegratedResponse"`

	// SCF WebSocket registration function namespace. This parameter takes effect when the frontend type is `WEBSOCKET` and the backend type is `SCF`
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceWebsocketRegisterFunctionName *string `json:"ServiceWebsocketRegisterFunctionName,omitempty" name:"ServiceWebsocketRegisterFunctionName"`

	// SCF WebSocket registration function namespace. This parameter takes effect when the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceWebsocketRegisterFunctionNamespace *string `json:"ServiceWebsocketRegisterFunctionNamespace,omitempty" name:"ServiceWebsocketRegisterFunctionNamespace"`

	// SCF WebSocket transfer function version. This parameter takes effect when the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceWebsocketRegisterFunctionQualifier *string `json:"ServiceWebsocketRegisterFunctionQualifier,omitempty" name:"ServiceWebsocketRegisterFunctionQualifier"`

	// SCF WebSocket cleanup function. This parameter takes effect when the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceWebsocketCleanupFunctionName *string `json:"ServiceWebsocketCleanupFunctionName,omitempty" name:"ServiceWebsocketCleanupFunctionName"`

	// SCF WebSocket cleanup function namespace. This parameter takes effect when the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceWebsocketCleanupFunctionNamespace *string `json:"ServiceWebsocketCleanupFunctionNamespace,omitempty" name:"ServiceWebsocketCleanupFunctionNamespace"`

	// SCF WebSocket cleanup function version. This parameter takes effect when the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceWebsocketCleanupFunctionQualifier *string `json:"ServiceWebsocketCleanupFunctionQualifier,omitempty" name:"ServiceWebsocketCleanupFunctionQualifier"`

	// WebSocket pushback address.
	// Note: this field may return null, indicating that no valid values can be obtained.
	InternalDomain *string `json:"InternalDomain,omitempty" name:"InternalDomain"`

	// SCF WebSocket transfer function. This parameter takes effect when the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceWebsocketTransportFunctionName *string `json:"ServiceWebsocketTransportFunctionName,omitempty" name:"ServiceWebsocketTransportFunctionName"`

	// SCF WebSocket transfer function namespace. This parameter takes effect when the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceWebsocketTransportFunctionNamespace *string `json:"ServiceWebsocketTransportFunctionNamespace,omitempty" name:"ServiceWebsocketTransportFunctionNamespace"`

	// SCF WebSocket transfer function version. This parameter takes effect when the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceWebsocketTransportFunctionQualifier *string `json:"ServiceWebsocketTransportFunctionQualifier,omitempty" name:"ServiceWebsocketTransportFunctionQualifier"`

	// List of microservices bound to API.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MicroServices []*MicroService `json:"MicroServices,omitempty" name:"MicroServices" list`

	// Microservice detailed information.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MicroServicesInfo []*int64 `json:"MicroServicesInfo,omitempty" name:"MicroServicesInfo" list`

	// Microservice load balancing configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceTsfLoadBalanceConf *TsfLoadBalanceConfResp `json:"ServiceTsfLoadBalanceConf,omitempty" name:"ServiceTsfLoadBalanceConf"`

	// Microservice health check configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceTsfHealthCheckConf *HealthCheckConf `json:"ServiceTsfHealthCheckConf,omitempty" name:"ServiceTsfHealthCheckConf"`

	// Whether to enable CORS.
	// Note: this field may return null, indicating that no valid values can be obtained.
	EnableCORS *bool `json:"EnableCORS,omitempty" name:"EnableCORS"`

	// Information of tags bound to API.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`
}

type ApiKey struct {

	// Created API key ID.
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`

	// Created API key.
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" name:"AccessKeySecret"`

	// Key type. Valid values: auto, manual.
	AccessKeyType *string `json:"AccessKeyType,omitempty" name:"AccessKeyType"`

	// Custom key name.
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// Key status. 0: disabled. 1: enabled.
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type ApiRequestConfig struct {

	// path
	Path *string `json:"Path,omitempty" name:"Path"`

	// Method
	Method *string `json:"Method,omitempty" name:"Method"`
}

type ApiUsagePlan struct {

	// Unique service ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique API ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// API name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// API path.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Path *string `json:"Path,omitempty" name:"Path"`

	// API method.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Method *string `json:"Method,omitempty" name:"Method"`

	// Unique usage plan ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// Usage plan name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UsagePlanName *string `json:"UsagePlanName,omitempty" name:"UsagePlanName"`

	// Usage plan description.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UsagePlanDesc *string `json:"UsagePlanDesc,omitempty" name:"UsagePlanDesc"`

	// Service environment bound to usage plan.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// Used quota.
	// Note: this field may return null, indicating that no valid values can be obtained.
	InUseRequestNum *int64 `json:"InUseRequestNum,omitempty" name:"InUseRequestNum"`

	// Total number of requests allowed. `-1` indicates no limit.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxRequestNum *int64 `json:"MaxRequestNum,omitempty" name:"MaxRequestNum"`

	// Request QPS upper limit. `-1` indicates no limit.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxRequestNumPreSec *int64 `json:"MaxRequestNumPreSec,omitempty" name:"MaxRequestNumPreSec"`

	// Usage plan creation time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Last modified time of usage plan.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// Service name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

type ApiUsagePlanSet struct {

	// Total number of usage plans bound to API.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of usage plans bound to API.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiUsagePlanList []*ApiUsagePlan `json:"ApiUsagePlanList,omitempty" name:"ApiUsagePlanList" list`
}

type BindEnvironmentRequest struct {
	*tchttp.BaseRequest

	// List of unique IDs of the usage plans to be bound.
	UsagePlanIds []*string `json:"UsagePlanIds,omitempty" name:"UsagePlanIds" list`

	// Binding type. Valid values: API, SERVICE (default value).
	BindType *string `json:"BindType,omitempty" name:"BindType"`

	// Environment to be bound.
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// Unique ID of the service to be bound.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique API ID array. This parameter will be required when `bindType` is `API`.
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds" list`
}

func (r *BindEnvironmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindEnvironmentRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindEnvironmentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Whether binding succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *bool `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindEnvironmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindEnvironmentResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindIPStrategyRequest struct {
	*tchttp.BaseRequest

	// Unique ID of the service of the IP policy to be bound.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique ID of the IP policy to be bound.
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`

	// Environment to be bound to IP policy
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// List of APIs to be bound to IP policy.
	BindApiIds []*string `json:"BindApiIds,omitempty" name:"BindApiIds" list`
}

func (r *BindIPStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindIPStrategyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindIPStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Whether binding succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *bool `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindIPStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindIPStrategyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindSecretIdsRequest struct {
	*tchttp.BaseRequest

	// Unique ID of the usage plan to be bound.
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// Array of IDs of the keys to be bound.
	AccessKeyIds []*string `json:"AccessKeyIds,omitempty" name:"AccessKeyIds" list`
}

func (r *BindSecretIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindSecretIdsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindSecretIdsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Whether binding succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *bool `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindSecretIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindSecretIdsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindSubDomainRequest struct {
	*tchttp.BaseRequest

	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Custom domain name to be bound.
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// Protocols supported by service. Valid values: http, https, http&https.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Network type. Valid values: INNER, OUTER.
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Whether to use the default path mapping. The default value is `true`. If the value is `false`, the custom path mapping will be used and `PathMappingSet` is required.
	IsDefaultMapping *bool `json:"IsDefaultMapping,omitempty" name:"IsDefaultMapping"`

	// Default domain name.
	NetSubDomain *string `json:"NetSubDomain,omitempty" name:"NetSubDomain"`

	// Unique ID of the certificate of the custom domain name to be bound. The certificate can be uploaded only when `Protocol` is `https` or `http&https`.
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// Custom domain name path mapping. It can contain up to 3 `Environment` parameters which can be set to only `test`, `prepub`, or `release`.
	PathMappingSet []*PathMapping `json:"PathMappingSet,omitempty" name:"PathMappingSet" list`
}

func (r *BindSubDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindSubDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindSubDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindSubDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindSubDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ConstantParameter struct {

	// Constant parameter name. This parameter takes effect only when `ServiceType` is `HTTP`.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Constant parameter description. This parameter takes effect only when `ServiceType` is `HTTP`.
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// Constant parameter location. This parameter takes effect only when `ServiceType` is `HTTP`.
	Position *string `json:"Position,omitempty" name:"Position"`

	// Constant parameter default value. This parameter takes effect only when `ServiceType` is `HTTP`.
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`
}

type CreateApiKeyRequest struct {
	*tchttp.BaseRequest

	// Custom key name.
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// Key type. Valid values: auto, manual (custom key). Default value: auto.
	AccessKeyType *string `json:"AccessKeyType,omitempty" name:"AccessKeyType"`

	// Custom key ID, which is required when `AccessKeyType` is `manual`. It can contain 5 to 50 letters, digits, and underscores.
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`

	// Custom key, which is required when `AccessKeyType` is `manual`. It can contain 10 to 50 letters, digits, and underscores.
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" name:"AccessKeySecret"`
}

func (r *CreateApiKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateApiKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// New key details.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *ApiKey `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateApiKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateApiRequest struct {
	*tchttp.BaseRequest

	// Unique ID of API's service.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API backend service type. Valid values: HTTP, MOCK, TSF, CLB, SCF, WEBSOCKET, TARGET (in beta test).
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// API backend service timeout period in seconds.
	ServiceTimeout *int64 `json:"ServiceTimeout,omitempty" name:"ServiceTimeout"`

	// API frontend request type, such as HTTP, HTTPS, or HTTP and HTTPS.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Request frontend configuration.
	RequestConfig *ApiRequestConfig `json:"RequestConfig,omitempty" name:"RequestConfig"`

	// Custom API name.
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// Custom API description.
	ApiDesc *string `json:"ApiDesc,omitempty" name:"ApiDesc"`

	// API type. Valid values: NORMAL (general API), TSF (microservice API). Default value: NORMAL.
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`

	// API authentication type. Valid values: SECRET (key pair authentication), NONE (no authentication), OAUTH. Default value: NONE.
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// Whether to enable CORS.
	EnableCORS *bool `json:"EnableCORS,omitempty" name:"EnableCORS"`

	// Constant parameters.
	ConstantParameters []*ConstantParameter `json:"ConstantParameters,omitempty" name:"ConstantParameters" list`

	// Frontend request parameters.
	RequestParameters []*RequestParameter `json:"RequestParameters,omitempty" name:"RequestParameters" list`

	// This field takes effect when `AuthType` is `OAUTH`. NORMAL: business API. OAUTH: authorization API.
	ApiBusinessType *string `json:"ApiBusinessType,omitempty" name:"ApiBusinessType"`

	// Returned information of API backend mocking. This parameter is required when `ServiceType` is `Mock`.
	ServiceMockReturnMessage *string `json:"ServiceMockReturnMessage,omitempty" name:"ServiceMockReturnMessage"`

	// List of microservices bound to API.
	MicroServices []*MicroServiceReq `json:"MicroServices,omitempty" name:"MicroServices" list`

	// Microservice load balancing configuration.
	ServiceTsfLoadBalanceConf *TsfLoadBalanceConfResp `json:"ServiceTsfLoadBalanceConf,omitempty" name:"ServiceTsfLoadBalanceConf"`

	// Microservice health check configuration.
	ServiceTsfHealthCheckConf *HealthCheckConf `json:"ServiceTsfHealthCheckConf,omitempty" name:"ServiceTsfHealthCheckConf"`

	// `target` type backend resource information (in beta test).
	TargetServices []*TargetServicesReq `json:"TargetServices,omitempty" name:"TargetServices" list`

	// `target` type load balancing configuration (in beta test).
	TargetServicesLoadBalanceConf *int64 `json:"TargetServicesLoadBalanceConf,omitempty" name:"TargetServicesLoadBalanceConf"`

	// `target` health check configuration (in beta test).
	TargetServicesHealthCheckConf *HealthCheckConf `json:"TargetServicesHealthCheckConf,omitempty" name:"TargetServicesHealthCheckConf"`

	// SCF function name. This parameter takes effect when the backend type is `SCF`.
	ServiceScfFunctionName *string `json:"ServiceScfFunctionName,omitempty" name:"ServiceScfFunctionName"`

	// SCF WebSocket registration function. This parameter takes effect when the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketRegisterFunctionName *string `json:"ServiceWebsocketRegisterFunctionName,omitempty" name:"ServiceWebsocketRegisterFunctionName"`

	// SCF WebSocket cleanup function. This parameter takes effect when the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketCleanupFunctionName *string `json:"ServiceWebsocketCleanupFunctionName,omitempty" name:"ServiceWebsocketCleanupFunctionName"`

	// SCF WebSocket transfer function. This parameter takes effect when the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketTransportFunctionName *string `json:"ServiceWebsocketTransportFunctionName,omitempty" name:"ServiceWebsocketTransportFunctionName"`

	// SCF function namespace. This parameter takes effect when the backend type is `SCF`.
	ServiceScfFunctionNamespace *string `json:"ServiceScfFunctionNamespace,omitempty" name:"ServiceScfFunctionNamespace"`

	// SCF function version. This parameter takes effect when the backend type is `SCF`.
	ServiceScfFunctionQualifier *string `json:"ServiceScfFunctionQualifier,omitempty" name:"ServiceScfFunctionQualifier"`

	// SCF WebSocket registration function namespace. This parameter takes effect when the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketRegisterFunctionNamespace *string `json:"ServiceWebsocketRegisterFunctionNamespace,omitempty" name:"ServiceWebsocketRegisterFunctionNamespace"`

	// SCF WebSocket transfer function version. This parameter takes effect when the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketRegisterFunctionQualifier *string `json:"ServiceWebsocketRegisterFunctionQualifier,omitempty" name:"ServiceWebsocketRegisterFunctionQualifier"`

	// SCF WebSocket transfer function namespace. This parameter takes effect when the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketTransportFunctionNamespace *string `json:"ServiceWebsocketTransportFunctionNamespace,omitempty" name:"ServiceWebsocketTransportFunctionNamespace"`

	// SCF WebSocket transfer function version. This parameter takes effect when the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketTransportFunctionQualifier *string `json:"ServiceWebsocketTransportFunctionQualifier,omitempty" name:"ServiceWebsocketTransportFunctionQualifier"`

	// SCF WebSocket cleanup function namespace. This parameter takes effect when the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketCleanupFunctionNamespace *string `json:"ServiceWebsocketCleanupFunctionNamespace,omitempty" name:"ServiceWebsocketCleanupFunctionNamespace"`

	// SCF WebSocket cleanup function version. This parameter takes effect when the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketCleanupFunctionQualifier *string `json:"ServiceWebsocketCleanupFunctionQualifier,omitempty" name:"ServiceWebsocketCleanupFunctionQualifier"`

	// Whether to enable response integration. This parameter takes effect when the backend type is `SCF`.
	ServiceScfIsIntegratedResponse *bool `json:"ServiceScfIsIntegratedResponse,omitempty" name:"ServiceScfIsIntegratedResponse"`

	// Billing after debugging starts (reserved for marketplace).
	IsDebugAfterCharge *bool `json:"IsDebugAfterCharge,omitempty" name:"IsDebugAfterCharge"`

	// Whether to delete the custom response configuration error codes. If the value is left empty or `False`, the error codes will not be deleted. If the value is `True`, all custom response configuration error codes of the API will be deleted.
	IsDeleteResponseErrorCodes *bool `json:"IsDeleteResponseErrorCodes,omitempty" name:"IsDeleteResponseErrorCodes"`

	// Return type.
	ResponseType *string `json:"ResponseType,omitempty" name:"ResponseType"`

	// Successful response sample of custom response configuration.
	ResponseSuccessExample *string `json:"ResponseSuccessExample,omitempty" name:"ResponseSuccessExample"`

	// Response failure sample of custom response configuration.
	ResponseFailExample *string `json:"ResponseFailExample,omitempty" name:"ResponseFailExample"`

	// API backend service configuration.
	ServiceConfig *ServiceConfig `json:"ServiceConfig,omitempty" name:"ServiceConfig"`

	// Unique ID of associated authorization API. This parameter takes effect only when `AuthType` is `OAUTH` and `ApiBusinessType` is `NORMAL`. It is the unique ID of the OAuth 2.0 authorization API bound to the business API.
	AuthRelationApiId *string `json:"AuthRelationApiId,omitempty" name:"AuthRelationApiId"`

	// API backend service parameters.
	ServiceParameters []*ServiceParameter `json:"ServiceParameters,omitempty" name:"ServiceParameters" list`

	// OAuth configuration. This parameter takes effect when `AuthType` is `OAUTH`.
	OauthConfig *OauthConfig `json:"OauthConfig,omitempty" name:"OauthConfig"`

	// Custom error code configuration.
	ResponseErrorCodes []*ResponseErrorCodeReq `json:"ResponseErrorCodes,omitempty" name:"ResponseErrorCodes" list`

	// TSF Serverless namespace ID (in beta test).
	TargetNamespaceId *string `json:"TargetNamespaceId,omitempty" name:"TargetNamespaceId"`

	// User type.
	UserType *string `json:"UserType,omitempty" name:"UserType"`
}

func (r *CreateApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateApiRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateApiResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// API information
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *CreateApiRsp `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateApiResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateApiRsp struct {

	// API ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// path
	// Note: this field may return null, indicating that no valid values can be obtained.
	Path *string `json:"Path,omitempty" name:"Path"`

	// method
	// Note: this field may return null, indicating that no valid values can be obtained.
	Method *string `json:"Method,omitempty" name:"Method"`

	// Creation time
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type CreateIPStrategyRequest struct {
	*tchttp.BaseRequest

	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Custom policy name.
	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`

	// Policy type. Valid values: WHITE (whitelist), BLACK (blacklist).
	StrategyType *string `json:"StrategyType,omitempty" name:"StrategyType"`

	// Policy details.
	StrategyData *string `json:"StrategyData,omitempty" name:"StrategyData"`
}

func (r *CreateIPStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateIPStrategyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateIPStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// New IP policy details.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *IPStrategy `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateIPStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateIPStrategyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateServiceRequest struct {
	*tchttp.BaseRequest

	// Custom service name. If this parameter is left empty, the system will automatically generate a unique name.
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// Service frontend request type, such as `http`, `https`, and `http&https`.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Custom service description.
	ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`

	// Self-deployed cluster name, which is used to specify the self-deployed cluster where the service is to be created.
	ExclusiveSetName *string `json:"ExclusiveSetName,omitempty" name:"ExclusiveSetName"`

	// Network type list, which is used to specify the supported network types. `INNER` indicates access over private network, and `OUTER` indicates access over public network. The default value is `OUTER`.
	NetTypes []*string `json:"NetTypes,omitempty" name:"NetTypes" list`

	// IP version number. Valid values: IPv4 (default value), IPv6
	IpVersion *string `json:"IpVersion,omitempty" name:"IpVersion"`

	// Cluster name, which is reserved and used by the `tsf serverless` type.
	SetServerName *string `json:"SetServerName,omitempty" name:"SetServerName"`

	// User type, which is reserved and can be used by `serverless` users.
	AppIdType *string `json:"AppIdType,omitempty" name:"AppIdType"`
}

func (r *CreateServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Unique service ID.
		ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

		// Custom service name.
		ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

		// Custom service description.
		ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`

		// Default public network domain name.
		OuterSubDomain *string `json:"OuterSubDomain,omitempty" name:"OuterSubDomain"`

		// Default domain name of VPC private network
		InnerSubDomain *string `json:"InnerSubDomain,omitempty" name:"InnerSubDomain"`

		// Service creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
		CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

		// Network type list. `INNER` indicates access over private network, and `OUTER` indicates access over public network.
		NetTypes []*string `json:"NetTypes,omitempty" name:"NetTypes" list`

		// IP version number.
	// Note: this field may return null, indicating that no valid values can be obtained.
		IpVersion *string `json:"IpVersion,omitempty" name:"IpVersion"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateUsagePlanRequest struct {
	*tchttp.BaseRequest

	// Custom usage plan name.
	UsagePlanName *string `json:"UsagePlanName,omitempty" name:"UsagePlanName"`

	// Custom usage plan description.
	UsagePlanDesc *string `json:"UsagePlanDesc,omitempty" name:"UsagePlanDesc"`

	// Total number of requests allowed. Valid values: -1, [1,99999999]. The default value is `-1`, which indicates no limit.
	MaxRequestNum *int64 `json:"MaxRequestNum,omitempty" name:"MaxRequestNum"`

	// Limit of requests per second. Valid values: -1, [1,2000]. The default value is `-1`, which indicates no limit.
	MaxRequestNumPreSec *int64 `json:"MaxRequestNumPreSec,omitempty" name:"MaxRequestNumPreSec"`
}

func (r *CreateUsagePlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateUsagePlanRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateUsagePlanResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Usage plan details.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *UsagePlanInfo `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUsagePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateUsagePlanResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteApiKeyRequest struct {
	*tchttp.BaseRequest

	// ID of the key to be deleted.
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
}

func (r *DeleteApiKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteApiKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Whether deletion succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *bool `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteApiKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteApiRequest struct {
	*tchttp.BaseRequest

	// Unique ID of API's service.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique API ID.
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
}

func (r *DeleteApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteApiRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteApiResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Whether deletion succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *bool `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteApiResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteIPStrategyRequest struct {
	*tchttp.BaseRequest

	// Unique ID of the service of the IP policy to be deleted.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique ID of the IP policy to be deleted.
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`
}

func (r *DeleteIPStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteIPStrategyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteIPStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Whether deletion succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *bool `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteIPStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteIPStrategyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceRequest struct {
	*tchttp.BaseRequest

	// Unique ID of the service to be deleted.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
}

func (r *DeleteServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Whether deletion succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *bool `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceSubDomainMappingRequest struct {
	*tchttp.BaseRequest

	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Custom domain name bound to service.
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// Name of the environment whose mapping is to be deleted. Valid values: test (testing environment), prepub (pre-publish environment), release (release environment).
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *DeleteServiceSubDomainMappingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteServiceSubDomainMappingRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceSubDomainMappingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Whether the path mapping of the custom domain name is successfully deleted.
		Result *bool `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteServiceSubDomainMappingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteServiceSubDomainMappingResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteUsagePlanRequest struct {
	*tchttp.BaseRequest

	// Unique ID of the usage plan to be deleted.
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`
}

func (r *DeleteUsagePlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteUsagePlanRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteUsagePlanResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Whether deletion succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *bool `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUsagePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteUsagePlanResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DemoteServiceUsagePlanRequest struct {
	*tchttp.BaseRequest

	// Usage plan ID.
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// Unique ID of the service to be demoted.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Environment name.
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *DemoteServiceUsagePlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DemoteServiceUsagePlanRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DemoteServiceUsagePlanResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Whether demotion succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *bool `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DemoteServiceUsagePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DemoteServiceUsagePlanResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DesApisStatus struct {

	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique API ID.
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// Custom API description.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiDesc *string `json:"ApiDesc,omitempty" name:"ApiDesc"`

	// Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// API name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// VPC ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`

	// Unique VPC ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// API type. Valid values (general API), TSF (microservice API).
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`

	// API protocol.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Whether to enable debugging on purchase (reserved for the marketplace)
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsDebugAfterCharge *bool `json:"IsDebugAfterCharge,omitempty" name:"IsDebugAfterCharge"`

	// API authentication type. Valid values: SECRET (key pair authentication), NONE (no authentication), OAUTH.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// OAuth API type. This parameter takes effect when `AuthType` is `OAUTH`. Valid values: NORMAL (business API), OAUTH (authorization API).
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiBusinessType *string `json:"ApiBusinessType,omitempty" name:"ApiBusinessType"`

	// Unique ID of associated authorization API. This parameter takes effect only when `AuthType` is `OAUTH` and `ApiBusinessType` is `NORMAL`. It is the unique ID of the OAuth 2.0 authorization API bound to the business API.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AuthRelationApiId *string `json:"AuthRelationApiId,omitempty" name:"AuthRelationApiId"`

	// OAuth configuration information. This parameter takes effect when `AuthType` is `OAUTH`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OauthConfig *OauthConfig `json:"OauthConfig,omitempty" name:"OauthConfig"`

	// List of business APIs associated with authorization API.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RelationBuniessApiIds []*string `json:"RelationBuniessApiIds,omitempty" name:"RelationBuniessApiIds" list`

	// Information of tags associated with API.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Tags []*string `json:"Tags,omitempty" name:"Tags" list`

	// API path, such as `/path`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Path *string `json:"Path,omitempty" name:"Path"`

	// API request method, such as `GET`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Method *string `json:"Method,omitempty" name:"Method"`
}

type DescribeApiEnvironmentStrategyRequest struct {
	*tchttp.BaseRequest

	// Unique ID of the service of API.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Environment list.
	EnvironmentNames []*string `json:"EnvironmentNames,omitempty" name:"EnvironmentNames" list`

	// Unique API ID.
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeApiEnvironmentStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiEnvironmentStrategyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiEnvironmentStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Details of policies bound to API
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *ApiEnvironmentStrategyStataus `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiEnvironmentStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiEnvironmentStrategyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiKeyRequest struct {
	*tchttp.BaseRequest

	// API key ID.
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
}

func (r *DescribeApiKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Key details.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *ApiKey `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiRequest struct {
	*tchttp.BaseRequest

	// Unique ID of API's service.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique API ID.
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
}

func (r *DescribeApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// API details.
		Result *ApiInfo `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiUsagePlanRequest struct {
	*tchttp.BaseRequest

	// Unique ID of the service to be queried.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeApiUsagePlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiUsagePlanRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiUsagePlanResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of usage plans bound to API.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *ApiUsagePlanSet `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiUsagePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiUsagePlanResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLogSearchRequest struct {
	*tchttp.BaseRequest

	// Log start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Log end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Service ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Exact search by `apiid` or `reqid`
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Number of logs returned at a time. Up to 100 logs can be returned at a time
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Subsequent content can be obtained based on the `ConText` returned last time. Up to 10,000 data entries can be obtained
	ConText *string `json:"ConText,omitempty" name:"ConText"`

	// Log sorting by time in ascending order (asc) or descending order (desc). The default value is `desc`
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// Fuzzy search of logs by keyword
	Query *string `json:"Query,omitempty" name:"Query"`
}

func (r *DescribeLogSearchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLogSearchRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLogSearchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Cursor for getting more search results. If the value is `""`, there will be no subsequent results
		ConText *string `json:"ConText,omitempty" name:"ConText"`

		// The returned result contains zero or multiple logs, which are in the following format:
	// '[$app_id][$env_name][$service_id][$http_host][$api_id][$uri][$scheme][rsp_st:$status][ups_st:$upstream_status]'
	// '[cip:$remote_addr][uip:$upstream_addr][vip:$server_addr][rsp_len:$bytes_sent][req_len:$request_length]'
	// '[req_t:$request_time][ups_rsp_t:$upstream_response_time][ups_conn_t:$upstream_connect_time][ups_head_t:$upstream_header_time]’
	// '[err_msg:$err_msg][tcp_rtt:$tcpinfo_rtt][$pid][$time_local][req_id:$request_id]';
	// 
	// Note:
	// app_id: user ID.
	// env_name: environment name.
	// service_id: service ID.
	// http_host: domain name.
	// api_id: API ID.
	// uri: request path.
	// scheme: HTTP/HTTPS protocol.
	// rsp_st: request response status code.
	// ups_st: backend business server response status code (if the request is passed through to the backend, this variable will not be empty. If the request is blocked in API Gateway, this variable will be `-`).
	// cip: client IP.
	// uip: backend business service (upstream) IP.
	// vip: VIP requested to be accessed.
	// rsp_len: response length.
	// req_len: request length.
	// req_t: total request response time.
	// ups_rsp_t: total backend response time (time between connection establishment by API Gateway and backend response reception).
	// ups_conn_t: time when the backend business server is successfully connected.
	// ups_head_t: time when the backend response head arrives.
	// err_msg: error message.
	// tcp_rtt: client TCP connection information. RTT (Round Trip Time) consists of three parts: link propagation delay, end system processing delay, queuing and processing delay in router cache.
	// req_id: request ID.
		LogSet []*string `json:"LogSet,omitempty" name:"LogSet" list`

		// Number of logs returned for a search (`TotalCount <= Limit`)
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogSearchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLogSearchResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceEnvironmentListRequest struct {
	*tchttp.BaseRequest

	// Unique ID of the service to be queried.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeServiceEnvironmentListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceEnvironmentListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceEnvironmentListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Details of environments bound to service.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *ServiceEnvironmentSet `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceEnvironmentListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceEnvironmentListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceEnvironmentReleaseHistoryRequest struct {
	*tchttp.BaseRequest

	// Unique ID of the service to be queried.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Environment name.
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeServiceEnvironmentReleaseHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceEnvironmentReleaseHistoryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceEnvironmentReleaseHistoryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Service release history.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *ServiceReleaseHistory `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceEnvironmentReleaseHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceEnvironmentReleaseHistoryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceEnvironmentStrategyRequest struct {
	*tchttp.BaseRequest

	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeServiceEnvironmentStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceEnvironmentStrategyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceEnvironmentStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Throttling policy list.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *ServiceEnvironmentStrategyStatus `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceEnvironmentStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceEnvironmentStrategyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceReleaseVersionRequest struct {
	*tchttp.BaseRequest

	// Unique ID of the service to be queried.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeServiceReleaseVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceReleaseVersionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceReleaseVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Service release version list.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *ServiceReleaseVersion `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceReleaseVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceReleaseVersionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceRequest struct {
	*tchttp.BaseRequest

	// Unique ID of the service to be queried.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
}

func (r *DescribeServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Unique service ID.
		ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

		// Service environment list.
		AvailableEnvironments []*string `json:"AvailableEnvironments,omitempty" name:"AvailableEnvironments" list`

		// Service name.
		ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

		// Service description.
	// Note: this field may return null, indicating that no valid values can be obtained.
		ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`

		// Protocols supported by service. Valid values: http, https, http&https.
		Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

		// Service creation time.
		CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

		// Service modification time.
		ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

		// Self-deployed cluster name.
		ExclusiveSetName *string `json:"ExclusiveSetName,omitempty" name:"ExclusiveSetName"`

		// Network type list. `INNER` indicates access over private network, and `OUTER` indicates access over public network.
		NetTypes []*string `json:"NetTypes,omitempty" name:"NetTypes" list`

		// Private network access subdomain name.
		InternalSubDomain *string `json:"InternalSubDomain,omitempty" name:"InternalSubDomain"`

		// Public network access subdomain name.
		OuterSubDomain *string `json:"OuterSubDomain,omitempty" name:"OuterSubDomain"`

		// Port number for HTTP access over private network.
		InnerHttpPort *int64 `json:"InnerHttpPort,omitempty" name:"InnerHttpPort"`

		// Port number for HTTPS access over private network.
		InnerHttpsPort *int64 `json:"InnerHttpsPort,omitempty" name:"InnerHttpsPort"`

		// Total number of APIs.
		ApiTotalCount *int64 `json:"ApiTotalCount,omitempty" name:"ApiTotalCount"`

		// API list.
	// Note: this field may return null, indicating that no valid values can be obtained.
		ApiIdStatusSet []*ApiIdStatus `json:"ApiIdStatusSet,omitempty" name:"ApiIdStatusSet" list`

		// Total number of usage plans.
		UsagePlanTotalCount *int64 `json:"UsagePlanTotalCount,omitempty" name:"UsagePlanTotalCount"`

		// Usage plan array.
	// Note: this field may return null, indicating that no valid values can be obtained.
		UsagePlanList []*UsagePlan `json:"UsagePlanList,omitempty" name:"UsagePlanList" list`

		// IP version.
	// Note: this field may return null, indicating that no valid values can be obtained.
		IpVersion *string `json:"IpVersion,omitempty" name:"IpVersion"`

		// Service user type.
	// Note: this field may return null, indicating that no valid values can be obtained.
		UserType *string `json:"UserType,omitempty" name:"UserType"`

		// Reserved field.
	// Note: this field may return null, indicating that no valid values can be obtained.
		SetId *int64 `json:"SetId,omitempty" name:"SetId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceSubDomainMappingsRequest struct {
	*tchttp.BaseRequest

	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Custom domain name bound to service.
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`
}

func (r *DescribeServiceSubDomainMappingsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceSubDomainMappingsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceSubDomainMappingsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Custom path mapping list.
		Result *ServiceSubDomainMappings `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceSubDomainMappingsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceSubDomainMappingsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceSubDomainsRequest struct {
	*tchttp.BaseRequest

	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeServiceSubDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceSubDomainsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceSubDomainsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Custom service domain names.
		Result *DomainSets `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceSubDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceSubDomainsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceUsagePlanRequest struct {
	*tchttp.BaseRequest

	// Unique ID of the service to be queried.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeServiceUsagePlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceUsagePlanRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceUsagePlanResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of usage plans bound to service.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *ServiceUsagePlanSet `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceUsagePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceUsagePlanResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUsagePlanEnvironmentsRequest struct {
	*tchttp.BaseRequest

	// Unique ID of the usage plan to be queried.
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// Binding type. Valid values: API, SERVICE (default value).
	BindType *string `json:"BindType,omitempty" name:"BindType"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeUsagePlanEnvironmentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUsagePlanEnvironmentsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUsagePlanEnvironmentsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Usage plan binding details.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *UsagePlanEnvironmentStatus `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUsagePlanEnvironmentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUsagePlanEnvironmentsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUsagePlanRequest struct {
	*tchttp.BaseRequest

	// Unique ID of the usage plan to be queried.
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`
}

func (r *DescribeUsagePlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUsagePlanRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUsagePlanResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Usage plan details.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *UsagePlanInfo `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUsagePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUsagePlanResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUsagePlanSecretIdsRequest struct {
	*tchttp.BaseRequest

	// Unique ID of bound usage plan.
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeUsagePlanSecretIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUsagePlanSecretIdsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUsagePlanSecretIdsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of keys bound to usage plan.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *UsagePlanBindSecretStatus `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUsagePlanSecretIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUsagePlanSecretIdsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableApiKeyRequest struct {
	*tchttp.BaseRequest

	// ID of the key to be disabled.
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
}

func (r *DisableApiKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableApiKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Whether the key is successfully disabled.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *bool `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableApiKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DocumentSDK struct {

	// Download link of generated file. Generated documents are stored in COS.
	DocumentURL *string `json:"DocumentURL,omitempty" name:"DocumentURL"`

	// Download link of generated SDK file. Generated SDKs are stored in COS.
	SdkURL *string `json:"SdkURL,omitempty" name:"SdkURL"`
}

type DomainSets struct {

	// Number of custom domain names under service
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Custom service domain name list.
	DomainSet []*DomainSetList `json:"DomainSet,omitempty" name:"DomainSet" list`
}

type EnableApiKeyRequest struct {
	*tchttp.BaseRequest

	// ID of the key to be enabled.
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
}

func (r *EnableApiKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableApiKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Whether the key is successfully enabled.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *bool `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableApiKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Environment struct {

	// Environment name.
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// Access path.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Release status. 1: released. 0: not released.
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Running version.
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`
}

type EnvironmentStrategy struct {

	// Environment name
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// Throttling value
	Quota *int64 `json:"Quota,omitempty" name:"Quota"`
}

type ErrorCodes struct {

	// Custom response configuration error code.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// Custom response configuration error message.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// Remarks of the custom response configuration error code.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// Custom error code conversion.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ConvertedCode *int64 `json:"ConvertedCode,omitempty" name:"ConvertedCode"`

	// Whether to enable error code conversion.
	// Note: this field may return null, indicating that no valid values can be obtained.
	NeedConvert *bool `json:"NeedConvert,omitempty" name:"NeedConvert"`
}

type Filter struct {

	// Field to be filtered.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Filter value of field.
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type GenerateApiDocumentRequest struct {
	*tchttp.BaseRequest

	// Unique service ID of the document to be created.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Environment of the service for which the SDK is to be created.
	GenEnvironment *string `json:"GenEnvironment,omitempty" name:"GenEnvironment"`

	// Programming language of the SDK to be created. Currently, only Python and JavaScript are supported.
	GenLanguage *string `json:"GenLanguage,omitempty" name:"GenLanguage"`
}

func (r *GenerateApiDocumentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GenerateApiDocumentRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GenerateApiDocumentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// API documentation and SDK link.
		Result *DocumentSDK `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GenerateApiDocumentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GenerateApiDocumentResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type HealthCheckConf struct {

	// Whether to enable health check.
	IsHealthCheck *bool `json:"IsHealthCheck,omitempty" name:"IsHealthCheck"`

	// Health check threshold.
	RequestVolumeThreshold *int64 `json:"RequestVolumeThreshold,omitempty" name:"RequestVolumeThreshold"`

	// Window size.
	SleepWindowInMilliseconds *int64 `json:"SleepWindowInMilliseconds,omitempty" name:"SleepWindowInMilliseconds"`

	// Threshold percentage.
	ErrorThresholdPercentage *int64 `json:"ErrorThresholdPercentage,omitempty" name:"ErrorThresholdPercentage"`
}

type IPStrategy struct {

	// Unique policy ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`

	// Custom policy name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`

	// Policy type. Valid values: WHITE (whitelist), BLACK (blacklist).
	// Note: this field may return null, indicating that no valid values can be obtained.
	StrategyType *string `json:"StrategyType,omitempty" name:"StrategyType"`

	// IP list.
	// Note: this field may return null, indicating that no valid values can be obtained.
	StrategyData *string `json:"StrategyData,omitempty" name:"StrategyData"`

	// Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Modification time
	// Note: this field may return null, indicating that no valid values can be obtained.
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Number of APIs bound to policy.
	// Note: this field may return null, indicating that no valid values can be obtained.
	BindApiTotalCount *int64 `json:"BindApiTotalCount,omitempty" name:"BindApiTotalCount"`

	// Bound API details.
	// Note: this field may return null, indicating that no valid values can be obtained.
	BindApis []*DesApisStatus `json:"BindApis,omitempty" name:"BindApis" list`
}

type MicroService struct {

	// Microservice cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Microservice namespace ID.
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// Microservice name.
	MicroServiceName *string `json:"MicroServiceName,omitempty" name:"MicroServiceName"`
}

type MicroServiceReq struct {

	// Microservice cluster.
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Microservice namespace.
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// Microservice name.
	MicroServiceName *string `json:"MicroServiceName,omitempty" name:"MicroServiceName"`
}

type ModifyApiEnvironmentStrategyRequest struct {
	*tchttp.BaseRequest

	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Throttling value.
	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`

	// Environment name.
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// API list.
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds" list`
}

func (r *ModifyApiEnvironmentStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyApiEnvironmentStrategyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyApiEnvironmentStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Whether modification succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *bool `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApiEnvironmentStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyApiEnvironmentStrategyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyApiIncrementRequest struct {
	*tchttp.BaseRequest

	// Service ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API ID
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// Authorization type of the API to be modified (you can select `OAUTH`, i.e., authorization API)
	BusinessType *string `json:"BusinessType,omitempty" name:"BusinessType"`

	// Public key value to be modified of OAuth API
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`

	// OAuth API redirect address
	LoginRedirectUrl *string `json:"LoginRedirectUrl,omitempty" name:"LoginRedirectUrl"`
}

func (r *ModifyApiIncrementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyApiIncrementRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyApiIncrementResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApiIncrementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyApiIncrementResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyApiRequest struct {
	*tchttp.BaseRequest

	// Unique ID of API's service.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API backend service type. Valid values: HTTP, MOCK, TSF, CLB, SCF, WEBSOCKET, TARGET (in beta test).
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// Request frontend configuration.
	RequestConfig *RequestConfig `json:"RequestConfig,omitempty" name:"RequestConfig"`

	// Unique API ID.
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// Custom API name.
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// Custom API description.
	ApiDesc *string `json:"ApiDesc,omitempty" name:"ApiDesc"`

	// API type. Valid values: NORMAL (default value), TSF.
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`

	// API authentication type. Valid values: SECRET, NONE (default), OAUTH.
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// Whether signature authentication is required. `True` indicates yes while `False` indicates no. This parameter is to be disused.
	AuthRequired *bool `json:"AuthRequired,omitempty" name:"AuthRequired"`

	// API backend service timeout period in seconds.
	ServiceTimeout *int64 `json:"ServiceTimeout,omitempty" name:"ServiceTimeout"`

	// API frontend request type, such as HTTP, HTTPS, or HTTP and HTTPS.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Whether to enable CORS. `True` indicates yes while `False` indicates no.
	EnableCORS *bool `json:"EnableCORS,omitempty" name:"EnableCORS"`

	// Constant parameters.
	ConstantParameters []*ConstantParameter `json:"ConstantParameters,omitempty" name:"ConstantParameters" list`

	// Frontend request parameters.
	RequestParameters []*ReqParameter `json:"RequestParameters,omitempty" name:"RequestParameters" list`

	// This field takes effect when `AuthType` is `OAUTH`. NORMAL: business API. OAUTH: authorization API.
	ApiBusinessType *string `json:"ApiBusinessType,omitempty" name:"ApiBusinessType"`

	// Returned information of API backend mocking. This parameter is required when `ServiceType` is `Mock`.
	ServiceMockReturnMessage *string `json:"ServiceMockReturnMessage,omitempty" name:"ServiceMockReturnMessage"`

	// List of microservices bound to API.
	MicroServices []*MicroServiceReq `json:"MicroServices,omitempty" name:"MicroServices" list`

	// Microservice load balancing configuration.
	ServiceTsfLoadBalanceConf *TsfLoadBalanceConfResp `json:"ServiceTsfLoadBalanceConf,omitempty" name:"ServiceTsfLoadBalanceConf"`

	// Microservice health check configuration.
	ServiceTsfHealthCheckConf *HealthCheckConf `json:"ServiceTsfHealthCheckConf,omitempty" name:"ServiceTsfHealthCheckConf"`

	// `target` type load balancing configuration (in beta test).
	TargetServicesLoadBalanceConf *int64 `json:"TargetServicesLoadBalanceConf,omitempty" name:"TargetServicesLoadBalanceConf"`

	// `target` health check configuration (in beta test).
	TargetServicesHealthCheckConf *HealthCheckConf `json:"TargetServicesHealthCheckConf,omitempty" name:"TargetServicesHealthCheckConf"`

	// SCF function name. This parameter takes effect when the backend type is `SCF`.
	ServiceScfFunctionName *string `json:"ServiceScfFunctionName,omitempty" name:"ServiceScfFunctionName"`

	// SCF WebSocket registration function. This parameter takes effect when the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketRegisterFunctionName *string `json:"ServiceWebsocketRegisterFunctionName,omitempty" name:"ServiceWebsocketRegisterFunctionName"`

	// SCF WebSocket cleanup function. This parameter takes effect when the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketCleanupFunctionName *string `json:"ServiceWebsocketCleanupFunctionName,omitempty" name:"ServiceWebsocketCleanupFunctionName"`

	// SCF WebSocket transfer function. This parameter takes effect when the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketTransportFunctionName *string `json:"ServiceWebsocketTransportFunctionName,omitempty" name:"ServiceWebsocketTransportFunctionName"`

	// SCF function namespace. This parameter takes effect when the backend type is `SCF`.
	ServiceScfFunctionNamespace *string `json:"ServiceScfFunctionNamespace,omitempty" name:"ServiceScfFunctionNamespace"`

	// SCF function version. This parameter takes effect when the backend type is `SCF`.
	ServiceScfFunctionQualifier *string `json:"ServiceScfFunctionQualifier,omitempty" name:"ServiceScfFunctionQualifier"`

	// SCF WebSocket registration function namespace. This parameter takes effect when the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketRegisterFunctionNamespace *string `json:"ServiceWebsocketRegisterFunctionNamespace,omitempty" name:"ServiceWebsocketRegisterFunctionNamespace"`

	// SCF WebSocket transfer function version. This parameter takes effect when the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketRegisterFunctionQualifier *string `json:"ServiceWebsocketRegisterFunctionQualifier,omitempty" name:"ServiceWebsocketRegisterFunctionQualifier"`

	// SCF WebSocket transfer function namespace. This parameter takes effect when the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketTransportFunctionNamespace *string `json:"ServiceWebsocketTransportFunctionNamespace,omitempty" name:"ServiceWebsocketTransportFunctionNamespace"`

	// SCF WebSocket transfer function version. This parameter takes effect when the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketTransportFunctionQualifier *string `json:"ServiceWebsocketTransportFunctionQualifier,omitempty" name:"ServiceWebsocketTransportFunctionQualifier"`

	// SCF WebSocket cleanup function namespace. This parameter takes effect when the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketCleanupFunctionNamespace *string `json:"ServiceWebsocketCleanupFunctionNamespace,omitempty" name:"ServiceWebsocketCleanupFunctionNamespace"`

	// SCF WebSocket cleanup function version. This parameter takes effect when the frontend type is `WEBSOCKET` and the backend type is `SCF`.
	ServiceWebsocketCleanupFunctionQualifier *string `json:"ServiceWebsocketCleanupFunctionQualifier,omitempty" name:"ServiceWebsocketCleanupFunctionQualifier"`

	// Whether to enable response integration. This parameter takes effect when the backend type is `SCF`.
	ServiceScfIsIntegratedResponse *bool `json:"ServiceScfIsIntegratedResponse,omitempty" name:"ServiceScfIsIntegratedResponse"`

	// Billing after debugging starts (reserved for marketplace).
	IsDebugAfterCharge *bool `json:"IsDebugAfterCharge,omitempty" name:"IsDebugAfterCharge"`

	// Tag.
	TagSpecifications *Tag `json:"TagSpecifications,omitempty" name:"TagSpecifications"`

	// Whether to delete the custom response configuration error codes. If the value is left empty or `False`, the error codes will not be deleted. If the value is `True`, all custom response configuration error codes of the API will be deleted.
	IsDeleteResponseErrorCodes *bool `json:"IsDeleteResponseErrorCodes,omitempty" name:"IsDeleteResponseErrorCodes"`

	// Return type.
	ResponseType *string `json:"ResponseType,omitempty" name:"ResponseType"`

	// Successful response sample of custom response configuration.
	ResponseSuccessExample *string `json:"ResponseSuccessExample,omitempty" name:"ResponseSuccessExample"`

	// Response failure sample of custom response configuration.
	ResponseFailExample *string `json:"ResponseFailExample,omitempty" name:"ResponseFailExample"`

	// API backend service configuration.
	ServiceConfig *ServiceConfig `json:"ServiceConfig,omitempty" name:"ServiceConfig"`

	// Unique ID of associated authorization API. This parameter takes effect only when `AuthType` is `OAUTH` and `ApiBusinessType` is `NORMAL`. It is the unique ID of the OAuth 2.0 authorization API bound to the business API.
	AuthRelationApiId *string `json:"AuthRelationApiId,omitempty" name:"AuthRelationApiId"`

	// API backend service parameters.
	ServiceParameters []*ServiceParameter `json:"ServiceParameters,omitempty" name:"ServiceParameters" list`

	// OAuth configuration. This parameter takes effect when `AuthType` is `OAUTH`.
	OauthConfig *OauthConfig `json:"OauthConfig,omitempty" name:"OauthConfig"`

	// Custom error code configuration.
	ResponseErrorCodes []*ResponseErrorCodeReq `json:"ResponseErrorCodes,omitempty" name:"ResponseErrorCodes" list`
}

func (r *ModifyApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyApiRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyApiResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyApiResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyIPStrategyRequest struct {
	*tchttp.BaseRequest

	// Unique ID of the service of the policy to be modified.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique ID of the policy to be modified.
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`

	// Details of the policy to be modified.
	StrategyData *string `json:"StrategyData,omitempty" name:"StrategyData"`
}

func (r *ModifyIPStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyIPStrategyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyIPStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Whether modification succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *bool `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyIPStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyIPStrategyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyServiceEnvironmentStrategyRequest struct {
	*tchttp.BaseRequest

	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Throttling value.
	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`

	// Environment list.
	EnvironmentNames []*string `json:"EnvironmentNames,omitempty" name:"EnvironmentNames" list`
}

func (r *ModifyServiceEnvironmentStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyServiceEnvironmentStrategyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyServiceEnvironmentStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Whether modification succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *bool `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyServiceEnvironmentStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyServiceEnvironmentStrategyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyServiceRequest struct {
	*tchttp.BaseRequest

	// Unique ID of the service to be modified.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Modified service name.
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// Modified service description.
	ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`

	// Modified service frontend request type, such as `http`, `https`, and `http&https`.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Network type list, which is used to specify the supported network types. `INNER` indicates access over private network, and `OUTER` indicates access over public network. The default value is `OUTER`.
	NetTypes []*string `json:"NetTypes,omitempty" name:"NetTypes" list`
}

func (r *ModifyServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySubDomainRequest struct {
	*tchttp.BaseRequest

	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Custom domain name for which the path mapping is to be modified.
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// Whether to change to the default path mapping. true: use the default path mapping. false: use the custom path mapping.
	IsDefaultMapping *bool `json:"IsDefaultMapping,omitempty" name:"IsDefaultMapping"`

	// Certificate ID, which is required when the HTTPS protocol is included.
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// Modified custom domain name protocol type. Valid values: http, https, http&https.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Modified path mapping list.
	PathMappingSet []*PathMapping `json:"PathMappingSet,omitempty" name:"PathMappingSet" list`

	// Network type. Valid values: INNER, OUTER.
	NetType *string `json:"NetType,omitempty" name:"NetType"`
}

func (r *ModifySubDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySubDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySubDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Whether the custom domain name is successfully modified.
		Result *bool `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySubDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySubDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyUsagePlanRequest struct {
	*tchttp.BaseRequest

	// Unique usage plan ID.
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// Modified custom usage plan name.
	UsagePlanName *string `json:"UsagePlanName,omitempty" name:"UsagePlanName"`

	// Modified custom usage plan description.
	UsagePlanDesc *string `json:"UsagePlanDesc,omitempty" name:"UsagePlanDesc"`

	// Total number of requests allowed. Valid values: -1, [1,99999999]. The default value is `-1`, which indicates no limit.
	MaxRequestNum *int64 `json:"MaxRequestNum,omitempty" name:"MaxRequestNum"`

	// Limit of requests per second. Valid values: -1, [1,2000]. The default value is `-1`, which indicates no limit.
	MaxRequestNumPreSec *int64 `json:"MaxRequestNumPreSec,omitempty" name:"MaxRequestNumPreSec"`
}

func (r *ModifyUsagePlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyUsagePlanRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyUsagePlanResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Usage plan details.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *UsagePlanInfo `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUsagePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyUsagePlanResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OauthConfig struct {

	// Public key for user token verification.
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`

	// Token delivery location.
	TokenLocation *string `json:"TokenLocation,omitempty" name:"TokenLocation"`

	// Redirect address, which is used to guide user logins.
	LoginRedirectUrl *string `json:"LoginRedirectUrl,omitempty" name:"LoginRedirectUrl"`
}

type PathMapping struct {

	// Path.
	Path *string `json:"Path,omitempty" name:"Path"`

	// Release environment. Valid values: test, prepub, release.
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type ReleaseService struct {

	// Release remarks.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`

	// Release version ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ReleaseVersion *string `json:"ReleaseVersion,omitempty" name:"ReleaseVersion"`
}

type ReleaseServiceRequest struct {
	*tchttp.BaseRequest

	// Unique ID of the service to be published.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Name of the environment to be published. Valid values: test (testing environment), prepub (pre-publish environment), release (release environment).
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// Release description.
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`

	// `apiId` list, which is reserved. Full API release is used by default.
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds" list`
}

func (r *ReleaseServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleaseServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReleaseServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Release information.
		Result *ReleaseService `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReleaseServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleaseServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RequestConfig struct {

	// API path, such as `/path`.
	Path *string `json:"Path,omitempty" name:"Path"`

	// API request method, such as `GET`.
	Method *string `json:"Method,omitempty" name:"Method"`
}

type RequestParameter struct {

	// Request parameter name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Description
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// Parameter location
	Position *string `json:"Position,omitempty" name:"Position"`

	// Parameter type
	Type *string `json:"Type,omitempty" name:"Type"`

	// Default value
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`

	// Required
	Required *bool `json:"Required,omitempty" name:"Required"`
}

type ResponseErrorCodeReq struct {

	// Custom response configuration error code.
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// Custom response configuration error message.
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// Remarks of the custom response configuration error code.
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// Custom error code conversion.
	ConvertedCode *int64 `json:"ConvertedCode,omitempty" name:"ConvertedCode"`

	// Whether to enable error code conversion.
	NeedConvert *bool `json:"NeedConvert,omitempty" name:"NeedConvert"`
}

type ServiceConfig struct {

	// Backend type. This parameter takes effect when VPC is enabled. Currently, only `clb` is supported.
	Product *string `json:"Product,omitempty" name:"Product"`

	// Unique VPC ID.
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// API backend service URL. This parameter is required when `ServiceType` is `HTTP`.
	Url *string `json:"Url,omitempty" name:"Url"`

	// API backend service path, such as `/path`. If `ServiceType` is `HTTP`, this parameter will be required. The frontend and backend paths can be different.
	Path *string `json:"Path,omitempty" name:"Path"`

	// API backend service request method, such as `GET`. If `ServiceType` is `HTTP`, this parameter will be required. The frontend and backend methods can be different
	Method *string `json:"Method,omitempty" name:"Method"`
}

type ServiceEnvironmentSet struct {

	// Total number of environments bound to service.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of environments bound to service
	// Note: this field may return null, indicating that no valid values can be obtained.
	EnvironmentList []*Environment `json:"EnvironmentList,omitempty" name:"EnvironmentList" list`
}

type ServiceEnvironmentStrategy struct {

	// Environment name.
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// Access service environment URL.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Release status.
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Release version ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// Throttling value.
	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`
}

type ServiceEnvironmentStrategyStatus struct {

	// Throttling policy number.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Throttling policy list.
	// Note: this field may return null, indicating that no valid values can be obtained.
	EnvironmentList []*ServiceEnvironmentStrategy `json:"EnvironmentList,omitempty" name:"EnvironmentList" list`
}

type ServiceParameter struct {

	// API backend service parameter name. This parameter takes effect only when `ServiceType` is `HTTP`. Frontend and backend parameter names can be different.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitempty" name:"Name"`

	// API backend service parameter position, such as `head`. This parameter takes effect only when `ServiceType` is `HTTP`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Position *string `json:"Position,omitempty" name:"Position"`

	// Position of the API frontend parameter corresponding to backend service parameter, such as `head`. This parameter takes effect only when `ServiceType` is `HTTP`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RelevantRequestParameterPosition *string `json:"RelevantRequestParameterPosition,omitempty" name:"RelevantRequestParameterPosition"`

	// Name of the API frontend parameter corresponding to backend service parameter. This parameter takes effect only when `ServiceType` is `HTTP`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RelevantRequestParameterName *string `json:"RelevantRequestParameterName,omitempty" name:"RelevantRequestParameterName"`

	// API backend service parameter default value. This parameter takes effect only when `ServiceType` is `HTTP`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`

	// API backend service parameter remarks. This parameter takes effect only when `ServiceType` is `HTTP`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RelevantRequestParameterDesc *string `json:"RelevantRequestParameterDesc,omitempty" name:"RelevantRequestParameterDesc"`

	// API backend service parameter type. This parameter takes effect only when `ServiceType` is `HTTP`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RelevantRequestParameterType *string `json:"RelevantRequestParameterType,omitempty" name:"RelevantRequestParameterType"`
}

type ServiceReleaseHistory struct {

	// Total number of release versions.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Historical version list.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VersionList []*ServiceReleaseHistoryInfo `json:"VersionList,omitempty" name:"VersionList" list`
}

type ServiceReleaseHistoryInfo struct {

	// Version ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// Version description.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VersionDesc *string `json:"VersionDesc,omitempty" name:"VersionDesc"`

	// Version release time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ReleaseTime *string `json:"ReleaseTime,omitempty" name:"ReleaseTime"`
}

type ServiceReleaseVersion struct {

	// Total number of release versions.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Release version list.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VersionList []*ServiceReleaseHistoryInfo `json:"VersionList,omitempty" name:"VersionList" list`
}

type ServiceSubDomainMappings struct {

	// Whether to use the default path mapping. true: use the default path mapping. false: use the custom path mapping (`PathMappingSet` is required).
	IsDefaultMapping *bool `json:"IsDefaultMapping,omitempty" name:"IsDefaultMapping"`

	// Custom path mapping list.
	PathMappingSet []*PathMapping `json:"PathMappingSet,omitempty" name:"PathMappingSet" list`
}

type ServiceUsagePlanSet struct {

	// Total number of usage plans bound to service.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of usage plans bound to service.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceUsagePlanList []*ApiUsagePlan `json:"ServiceUsagePlanList,omitempty" name:"ServiceUsagePlanList" list`
}

type Tag struct {

	// Tag key.
	Key *string `json:"Key,omitempty" name:"Key"`

	// Tag value.
	Value *string `json:"Value,omitempty" name:"Value"`
}

type TargetServicesReq struct {

	// VM IP
	VmIp *string `json:"VmIp,omitempty" name:"VmIp"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VM Port
	VmPort *int64 `json:"VmPort,omitempty" name:"VmPort"`

	// IP of the host where the CVM instance resides
	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`

	// Docker IP
	DockerIp *string `json:"DockerIp,omitempty" name:"DockerIp"`
}

type TsfLoadBalanceConfResp struct {

	// Whether to enable load balancing.
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsLoadBalance *bool `json:"IsLoadBalance,omitempty" name:"IsLoadBalance"`

	// Load balancing method.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Method *string `json:"Method,omitempty" name:"Method"`

	// Whether to enable session persistence.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SessionStickRequired *bool `json:"SessionStickRequired,omitempty" name:"SessionStickRequired"`

	// Session persistence timeout period.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SessionStickTimeout *int64 `json:"SessionStickTimeout,omitempty" name:"SessionStickTimeout"`
}

type UnBindEnvironmentRequest struct {
	*tchttp.BaseRequest

	// Binding type. Valid values: API, SERVICE (default value).
	BindType *string `json:"BindType,omitempty" name:"BindType"`

	// List of unique IDs of the usage plans to be bound.
	UsagePlanIds []*string `json:"UsagePlanIds,omitempty" name:"UsagePlanIds" list`

	// Service environment to be unbound.
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// Unique ID of the service to be unbound.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique API ID array. This parameter will be required when `BindType` is `API`.
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds" list`
}

func (r *UnBindEnvironmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindEnvironmentRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnBindEnvironmentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Whether unbinding succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *bool `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnBindEnvironmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindEnvironmentResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnBindIPStrategyRequest struct {
	*tchttp.BaseRequest

	// Unique ID of the service to be unbound.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique ID of the IP policy to be unbound.
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`

	// Environment to be unbound.
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// List of APIs to be unbound.
	UnBindApiIds []*string `json:"UnBindApiIds,omitempty" name:"UnBindApiIds" list`
}

func (r *UnBindIPStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindIPStrategyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnBindIPStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Whether unbinding succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *bool `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnBindIPStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindIPStrategyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnBindSecretIdsRequest struct {
	*tchttp.BaseRequest

	// Unique ID of the usage plan to be unbound.
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// Array of IDs of the keys to be unbound.
	AccessKeyIds []*string `json:"AccessKeyIds,omitempty" name:"AccessKeyIds" list`
}

func (r *UnBindSecretIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindSecretIdsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnBindSecretIdsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Whether unbinding succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *bool `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnBindSecretIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindSecretIdsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnBindSubDomainRequest struct {
	*tchttp.BaseRequest

	// Unique service ID.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Custom domain name to be unbound.
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`
}

func (r *UnBindSubDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindSubDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnBindSubDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Whether the custom domain name is successfully unbound.
		Result *bool `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnBindSubDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindSubDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnReleaseServiceRequest struct {
	*tchttp.BaseRequest

	// Unique ID of the service to be deactivated.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Name of the environment to be deactivated. Valid values: test (testing environment), prepub (pre-publish environment), release (release environment).
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// List of APIs to be deactivated, which is a reserved field.
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds" list`
}

func (r *UnReleaseServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnReleaseServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnReleaseServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Whether deactivation succeeded.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *bool `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnReleaseServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnReleaseServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateApiKeyRequest struct {
	*tchttp.BaseRequest

	// ID of the key to be updated.
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`

	// Key to be updated, which is required when a custom key is updated. It can contain 10 to 50 letters, digits, and underscores.
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" name:"AccessKeySecret"`
}

func (r *UpdateApiKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateApiKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Changed key details.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *ApiKey `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateApiKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateServiceRequest struct {
	*tchttp.BaseRequest

	// Unique ID of the service to be switch to.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Name of the environment to be switched to. Valid values: test (testing environment), prepub (pre-publish environment), release (release environment).
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// Number of the version to be switched to.
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// Switch description.
	UpdateDesc *string `json:"UpdateDesc,omitempty" name:"UpdateDesc"`
}

func (r *UpdateServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Whether the version is successfully switched.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *bool `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UsagePlan struct {

	// Environment name.
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// Unique usage plan ID.
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// Usage plan name.
	UsagePlanName *string `json:"UsagePlanName,omitempty" name:"UsagePlanName"`

	// Usage plan description.
	UsagePlanDesc *string `json:"UsagePlanDesc,omitempty" name:"UsagePlanDesc"`

	// Usage plan QPS. `-1` indicates no limit.
	MaxRequestNumPreSec *int64 `json:"MaxRequestNumPreSec,omitempty" name:"MaxRequestNumPreSec"`

	// Usage plan time.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Usage plan modification time.
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`
}

type UsagePlanBindSecret struct {

	// Key ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`

	// Key name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// Key status. 0: disabled. 1: enabled.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type UsagePlanBindSecretStatus struct {

	// Number of keys bound to usage plan.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of key details.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AccessKeyList []*UsagePlanBindSecret `json:"AccessKeyList,omitempty" name:"AccessKeyList" list`
}

type UsagePlanEnvironment struct {

	// Unique ID of bound service.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Unique API ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// API name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// API path.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Path *string `json:"Path,omitempty" name:"Path"`

	// API method.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Method *string `json:"Method,omitempty" name:"Method"`

	// Name of the bound environment.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// Used quota.
	// Note: this field may return null, indicating that no valid values can be obtained.
	InUseRequestNum *int64 `json:"InUseRequestNum,omitempty" name:"InUseRequestNum"`

	// Maximum number of requests.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxRequestNum *int64 `json:"MaxRequestNum,omitempty" name:"MaxRequestNum"`

	// Maximum requests per second.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxRequestNumPreSec *int64 `json:"MaxRequestNumPreSec,omitempty" name:"MaxRequestNumPreSec"`

	// Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// Service name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

type UsagePlanEnvironmentStatus struct {

	// Number of environments of the service bound to usage plan.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Environment status of the services bound to usage plan.
	// Note: this field may return null, indicating that no valid values can be obtained.
	EnvironmentList []*UsagePlanEnvironment `json:"EnvironmentList,omitempty" name:"EnvironmentList" list`
}
