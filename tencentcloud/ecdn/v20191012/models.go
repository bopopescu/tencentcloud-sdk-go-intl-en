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

package v20191012

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AddEcdnDomainRequest struct {
	*tchttp.BaseRequest

	// Domain name.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Origin server configuration.
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// Domain name acceleration region. Valid values: mainland (acceleration in Mainland China), overseas (acceleration outside Mainland China), global (global acceleration).
	Area *string `json:"Area,omitempty" name:"Area"`

	// Project ID. Default value: 0.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// IP blacklist/whitelist configuration.
	IpFilter *IpFilter `json:"IpFilter,omitempty" name:"IpFilter"`

	// IP access limit configuration.
	IpFreqLimit *IpFreqLimit `json:"IpFreqLimit,omitempty" name:"IpFreqLimit"`

	// Origin server response header configuration.
	ResponseHeader *ResponseHeader `json:"ResponseHeader,omitempty" name:"ResponseHeader"`

	// Node caching configuration.
	CacheKey *CacheKey `json:"CacheKey,omitempty" name:"CacheKey"`

	// Caching rule configuration.
	Cache *Cache `json:"Cache,omitempty" name:"Cache"`

	// HTTPS configuration.
	Https *Https `json:"Https,omitempty" name:"Https"`

	// Forced access protocol redirection configuration.
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitempty" name:"ForceRedirect"`
}

func (r *AddEcdnDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddEcdnDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddEcdnDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddEcdnDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddEcdnDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Cache struct {

	// Caching configuration rule array.
	CacheRules []*CacheRule `json:"CacheRules,omitempty" name:"CacheRules" list`

	// Whether to follow origin server's `Cache-Control: max-age` configuration
	// on: enable.
	// off: disable.
	// After this feature is enabled, resources that do not match the `CacheRules` rule will be cached on nodes according to the `max-age` value returned by the origin server, while resources that match the `CacheRules` rule will be cached on nodes according to the cache expiration time set in `CacheRules`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FollowOrigin *string `json:"FollowOrigin,omitempty" name:"FollowOrigin"`
}

type CacheKey struct {

	// Whether to enable full path cache. Valid values: on, off.
	FullUrlCache *string `json:"FullUrlCache,omitempty" name:"FullUrlCache"`
}

type CacheRule struct {

	// Cache type. Valid values: all (all files), file (extension type), directory (directory), path (full path), index (homepage).
	CacheType *string `json:"CacheType,omitempty" name:"CacheType"`

	// Cached content list.
	CacheContents []*string `json:"CacheContents,omitempty" name:"CacheContents" list`

	// Cache time in seconds.
	CacheTime *int64 `json:"CacheTime,omitempty" name:"CacheTime"`
}

type ClientCert struct {

	// Client certificate in PEM format.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Certificate *string `json:"Certificate,omitempty" name:"Certificate"`

	// Client certificate name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// Certificate expiration time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// Certificate issuance time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DeployTime *string `json:"DeployTime,omitempty" name:"DeployTime"`
}

type DeleteEcdnDomainRequest struct {
	*tchttp.BaseRequest

	// Domain name to be deleted.
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DeleteEcdnDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteEcdnDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteEcdnDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteEcdnDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteEcdnDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainsConfigRequest struct {
	*tchttp.BaseRequest

	// Pagination offset address. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of domain names per page. Default value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Query filter.
	Filters []*DomainFilter `json:"Filters,omitempty" name:"Filters" list`

	// Query result sorting rule.
	Sort *Sort `json:"Sort,omitempty" name:"Sort"`
}

func (r *DescribeDomainsConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDomainsConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainsConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Domain name list.
		Domains []*DomainDetailInfo `json:"Domains,omitempty" name:"Domains" list`

		// Number of matched domain names. This is used for paginated query.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainsConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDomainsConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainsRequest struct {
	*tchttp.BaseRequest

	// Pagination offset address. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of domain names per page. Default value: 100. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Query filter.
	Filters []*DomainFilter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDomainsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Domain name information list.
		Domains []*DomainBriefInfo `json:"Domains,omitempty" name:"Domains" list`

		// Total number of domain names.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDomainsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEcdnDomainLogsRequest struct {
	*tchttp.BaseRequest

	// Domain name to be queried.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Log start time, such as 2019-10-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Log end time, such as 2019-10-02 00:00:00. Only logs for the last 30 days can be queried.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Pagination offset for log link list. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of log links per page. Default value: 100. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeEcdnDomainLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEcdnDomainLogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEcdnDomainLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Log link list.
	// Note: this field may return null, indicating that no valid values can be obtained.
		DomainLogs []*DomainLogs `json:"DomainLogs,omitempty" name:"DomainLogs" list`

		// Total number of log links.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEcdnDomainLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEcdnDomainLogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEcdnDomainStatisticsRequest struct {
	*tchttp.BaseRequest

	// Query start time, such as 2019-12-13 00:00:00.
	// The time span cannot exceed 90 days.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Query end time, such as 2019-12-13 23:59:59.
	// The time span cannot exceed 90 days.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Statistical metric name. flux: traffic in bytes
	// bandwidth: bandwidth in bps
	// request: number of requests
	// delay: response time in ms
	// static_request: number of static requests
	// static_flux: static traffic in bytes
	// static_bandwidth: static bandwidth in bps
	// dynamic_request: number of dynamic requests
	// dynamic_flux: dynamic traffic in bytes
	// dynamic_bandwidth: dynamic bandwidth in bps
	Metrics []*string `json:"Metrics,omitempty" name:"Metrics" list`

	// Specifies the list of domain names to be queried
	Domains []*string `json:"Domains,omitempty" name:"Domains" list`

	// Specifies the project ID to be queried, which can be viewed [here](https://console.cloud.tencent.com/project)
	// If no domain name is entered, the specified project will be queried; otherwise, the domain name will prevail
	Projects []*int64 `json:"Projects,omitempty" name:"Projects" list`

	// Pagination offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page. Default value: 1000. Maximum value: 3,000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeEcdnDomainStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEcdnDomainStatisticsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEcdnDomainStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Domain name data
		Data []*DomainData `json:"Data,omitempty" name:"Data" list`

		// Quantity
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEcdnDomainStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEcdnDomainStatisticsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEcdnStatisticsRequest struct {
	*tchttp.BaseRequest

	// Query start time, such as 2019-12-13 00:00:00
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Query end time, such as 2019-12-13 23:59:59
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Specifies the query metric, which can be:
	// flux: traffic in bytes
	// bandwidth: bandwidth in bps
	// request: number of requests
	// delay: response time in ms
	// 2xx: returns the number of 2xx status codes or details of status codes starting with 2
	// 3xx: returns the number of 3xx status codes or details of status codes starting with 3
	// 4xx: returns the number of 4xx status codes or details of status codes starting with 4
	// 5xx: returns the number of 5xx status codes or details of status codes starting with 5
	// static_request: number of static requests
	// static_flux: static traffic in bytes
	// static_bandwidth: static bandwidth in bps
	// dynamic_request: number of dynamic requests
	// dynamic_flux: dynamic traffic in bytes
	// dynamic_bandwidth: dynamic bandwidth in bps
	Metrics []*string `json:"Metrics,omitempty" name:"Metrics" list`

	// Time granularity, which can be:
	// 1 day	 1, 5, 15, 30, 60, 120, 240, 1440 
	// 2–3 days 15, 30, 60, 120, 240, 1440
	// 4–7 days 30, 60, 120, 240, 1440
	// 8–90 days	 60, 120, 240, 1440
	Interval *int64 `json:"Interval,omitempty" name:"Interval"`

	// Specifies the list of domain names to be queried
	// 
	// Up to 30 acceleration domain names can be queried at a time.
	Domains []*string `json:"Domains,omitempty" name:"Domains" list`

	// Specifies the project ID to be queried, which can be viewed [here](https://console.cloud.tencent.com/project)
	// If no domain name is entered, the specified project will be queried; otherwise, the domain name will prevail
	Projects []*int64 `json:"Projects,omitempty" name:"Projects" list`
}

func (r *DescribeEcdnStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEcdnStatisticsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEcdnStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned data details of the specified conditional query
		Data []*ResourceData `json:"Data,omitempty" name:"Data" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEcdnStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEcdnStatisticsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePurgeQuotaRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribePurgeQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePurgeQuotaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePurgeQuotaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// URL purge usage and quota.
		UrlPurge *Quota `json:"UrlPurge,omitempty" name:"UrlPurge"`

		// Directory purge usage and quota.
		PathPurge *Quota `json:"PathPurge,omitempty" name:"PathPurge"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePurgeQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePurgeQuotaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePurgeTasksRequest struct {
	*tchttp.BaseRequest

	// Purge type to be queried. url: query URL purge records; path: query directory purge records.
	PurgeType *string `json:"PurgeType,omitempty" name:"PurgeType"`

	// Start time, such as 2018-08-08 00:00:00
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time, such as 2018-08-08 23:59:59
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Task ID returned during submission. Either `TaskId` or start time must be specified for a query.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Pagination offset. Default value: 0 (starting from entry 0).
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Pagination limit. Default value: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Query keyword. Please enter a domain name or full URL beginning with `http(s)://`.
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// Specified task status to be queried. fail: failed, done: succeeded, process: purging.
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *DescribePurgeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePurgeTasksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePurgeTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Purge history.
		PurgeLogs []*PurgeTask `json:"PurgeLogs,omitempty" name:"PurgeLogs" list`

		// Total number of tasks, which is used for pagination.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePurgeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePurgeTasksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetailData struct {

	// Data type name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Data value
	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type DomainBriefInfo struct {

	// Domain name ID.
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Tencent Cloud account ID.
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// CDN acceleration domain name.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Domain name CNAME.
	Cname *string `json:"Cname,omitempty" name:"Cname"`

	// Domain name status. Valid values: pending (reviewing), rejected (rejected), processing (deploying after approval), online (enabled), offline (disabled), deleted (deleted).
	Status *string `json:"Status,omitempty" name:"Status"`

	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Domain name creation time.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Domain name update time.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Origin server configuration details.
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// Domain name blockage status. Valid values: normal (normal), overdue (service is suspended due to arrears), quota (trial traffic package is used up), malicious (malicious user), ddos (DDoS attack), idle (no traffic), unlicensed (no ICP filing), capping (bandwidth cap reached), readonly (read-only)
	Disable *string `json:"Disable,omitempty" name:"Disable"`

	// Acceleration region. Valid values: mainland, oversea, global.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Domain name lock status. normal: not locked; global: globally locked
	Readonly *string `json:"Readonly,omitempty" name:"Readonly"`
}

type DomainData struct {

	// Domain name
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// Result details.
	DetailData []*DetailData `json:"DetailData,omitempty" name:"DetailData" list`
}

type DomainDetailInfo struct {

	// Domain name ID.
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Tencent Cloud account ID.
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// Acceleration domain name.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Domain name CNAME.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Cname *string `json:"Cname,omitempty" name:"Cname"`

	// Domain name status. Valid values: pending (reviewing), rejected (rejected), processing (deploying after approval), online (enabled), offline (disabled), deleted (deleted).
	Status *string `json:"Status,omitempty" name:"Status"`

	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Domain name creation time.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Domain name update time.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Origin server configuration.
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// IP blacklist/whitelist configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	IpFilter *IpFilter `json:"IpFilter,omitempty" name:"IpFilter"`

	// IP access limit configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	IpFreqLimit *IpFreqLimit `json:"IpFreqLimit,omitempty" name:"IpFreqLimit"`

	// Origin server response header configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ResponseHeader *ResponseHeader `json:"ResponseHeader,omitempty" name:"ResponseHeader"`

	// Node caching configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CacheKey *CacheKey `json:"CacheKey,omitempty" name:"CacheKey"`

	// Caching rule configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Cache *Cache `json:"Cache,omitempty" name:"Cache"`

	// HTTPS configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Https *Https `json:"Https,omitempty" name:"Https"`

	// Domain name blockage status. Valid values: normal (normal), overdue (service is suspended due to arrears), quota (trial traffic package is used up), malicious (malicious user), ddos (DDoS attack), idle (no traffic), unlicensed (no ICP filing), capping (bandwidth cap reached), readonly (read-only).
	// Note: this field may return null, indicating that no valid values can be obtained.
	Disable *string `json:"Disable,omitempty" name:"Disable"`

	// Forced access protocol redirection configuration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitempty" name:"ForceRedirect"`

	// Acceleration region. Valid values: mainland, overseas, global.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Domain name lock status. normal: not locked; global: globally locked.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Readonly *string `json:"Readonly,omitempty" name:"Readonly"`
}

type DomainFilter struct {

	// Filter field name, which can be:
	// - origin: primary origin server.
	// - domain: domain name.
	// - resourceId: domain name ID.
	// - status: domain name status. Valid values: online, offline, processing.
	// - disable: domain name blockage status. Valid values: normal, unlicensed.
	// - projectId: project ID.
	// - fullUrlCache: full path cache. Valid values: on, off.
	// - https: whether to configure HTTPS. Valid values: on, off, processing.
	// - originPullProtocol: origin-pull protocol type. Valid values: http, follow, https.
	// - area: acceleration region. Valid values: mainland, overseas, global.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Filter field value.
	Value []*string `json:"Value,omitempty" name:"Value" list`

	// Whether to enable fuzzy query, which is supported only for filter fields `origin` and `domain`.
	Fuzzy *bool `json:"Fuzzy,omitempty" name:"Fuzzy"`
}

type DomainLogs struct {

	// Log start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Log end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Log download path.
	LogPath *string `json:"LogPath,omitempty" name:"LogPath"`
}

type EcdnData struct {

	// Queries the specified metric. Valid values: Bandwidth, Flux, Request, Delay, status code, LogBandwidth, LogFlux, LogRequest
	Metrics []*string `json:"Metrics,omitempty" name:"Metrics" list`

	// Detailed data collection
	DetailData []*TimestampData `json:"DetailData,omitempty" name:"DetailData" list`
}

type ForceRedirect struct {

	// Forced access protocol redirection configuration switch. Valid values: on, off.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Access protocol type for forced redirection. Valid values: http (forced redirection to HTTP protocol), https (forced redirection to HTTPS protocol).
	// Note: this field may return null, indicating that no valid values can be obtained.
	RedirectType *string `json:"RedirectType,omitempty" name:"RedirectType"`

	// HTTP status code returned when forced redirection is enabled. Valid values: 301, 302.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RedirectStatusCode *int64 `json:"RedirectStatusCode,omitempty" name:"RedirectStatusCode"`
}

type HttpHeaderPathRule struct {

	// HTTP header setting method. Valid values: add (add header), set (set header), del (delete header).
	// Request header currently does not support `set`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	HeaderMode *string `json:"HeaderMode,omitempty" name:"HeaderMode"`

	// HTTP header name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	HeaderName *string `json:"HeaderName,omitempty" name:"HeaderName"`

	// HTTP header value, which is optional when it is `del`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	HeaderValue *string `json:"HeaderValue,omitempty" name:"HeaderValue"`

	// Type of effective URL path rule. Valid values: all (all paths), file (file extension), directory (directory), path (absolute path).
	// Note: this field may return null, indicating that no valid values can be obtained.
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// URL path or file type list
	// Note: this field may return null, indicating that no valid values can be obtained.
	RulePaths []*string `json:"RulePaths,omitempty" name:"RulePaths" list`
}

type Https struct {

	// HTTPS configuration switch. Valid values: on, off. If the domain name with HTTPS configuration enabled is being deployed, this switch will be `off`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Whether to enable HTTP2. Valid values: on, off.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Http2 *string `json:"Http2,omitempty" name:"Http2"`

	// Whether to enable the OCSP feature. Valid values: on, off.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OcspStapling *string `json:"OcspStapling,omitempty" name:"OcspStapling"`

	// Whether to enable the client certificate verification feature. Valid values: on, off. The client certificate information must be uploaded if this feature is enabled.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VerifyClient *string `json:"VerifyClient,omitempty" name:"VerifyClient"`

	// Server certificate configuration information.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertInfo *ServerCert `json:"CertInfo,omitempty" name:"CertInfo"`

	// Client certificate configuration information.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClientCertInfo *ClientCert `json:"ClientCertInfo,omitempty" name:"ClientCertInfo"`

	// Whether to enable SPDY. Valid values: on, off.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Spdy *string `json:"Spdy,omitempty" name:"Spdy"`

	// HTTPS certificate deployment status. Valid values: closed (disabled), deploying (deploying), deployed (deployment succeeded), failed (deployment failed). This parameter cannot be used as an input parameter.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SslStatus *string `json:"SslStatus,omitempty" name:"SslStatus"`
}

type IpFilter struct {

	// IP blacklist/whitelist switch. Valid values: on, off.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// IP blacklist/whitelist type. Valid values: whitelist, blacklist.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FilterType *string `json:"FilterType,omitempty" name:"FilterType"`

	// IP blacklist/whitelist list.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Filters []*string `json:"Filters,omitempty" name:"Filters" list`
}

type IpFreqLimit struct {

	// IP access limit switch. Valid values: on, off.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Number of requests per second.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Qps *int64 `json:"Qps,omitempty" name:"Qps"`
}

type Origin struct {

	// Primary origin server list. The default format is ["ip1:port1", "ip2:port2"].
	// Weights can be configured in the origin server list. The weight format of IP origin servers is ["ip1:port1:weight1", "ip2:port2:weight2"].
	Origins []*string `json:"Origins,omitempty" name:"Origins" list`

	// Primary origin server type. Valid values: domain (domain name origin server), ip (IP origin server).
	// This is required when setting `Origins`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// Host header value during origin-pull.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// Origin-pull protocol type. Valid values: http (forced HTTP origin-pull), follow (protocol follow), https (HTTPS origin-pull).
	// Note: this field may return null, indicating that no valid values can be obtained.
	OriginPullProtocol *string `json:"OriginPullProtocol,omitempty" name:"OriginPullProtocol"`

	// Secondary origin server list.
	BackupOrigins []*string `json:"BackupOrigins,omitempty" name:"BackupOrigins" list`

	// Secondary origin server type, which is the same as `OriginType`.
	// This is required when setting `BackupOrigins`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	BackupOriginType *string `json:"BackupOriginType,omitempty" name:"BackupOriginType"`
}

type PurgePathCacheRequest struct {
	*tchttp.BaseRequest

	// List of directories to be purged. The protocol header must be included.
	Paths []*string `json:"Paths,omitempty" name:"Paths" list`

	// Purge type. flush: purges updated resources, delete: purges all resources.
	FlushType *string `json:"FlushType,omitempty" name:"FlushType"`
}

func (r *PurgePathCacheRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PurgePathCacheRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PurgePathCacheResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Purge task ID. The first ten digits are the UTC time when the task is submitted.
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PurgePathCacheResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PurgePathCacheResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PurgeTask struct {

	// Purge task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Purged URL.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Purge task status. fail: failed, done: succeeded, process: purging.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Purge type. url: URL purge; path: directory purge.
	PurgeType *string `json:"PurgeType,omitempty" name:"PurgeType"`

	// Resource purge method. flush: purges updated resources, delete: purges all resources.
	FlushType *string `json:"FlushType,omitempty" name:"FlushType"`

	// Purge task submission time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type PurgeUrlsCacheRequest struct {
	*tchttp.BaseRequest

	// List of URLs to be purged. The protocol header must be included.
	Urls []*string `json:"Urls,omitempty" name:"Urls" list`
}

func (r *PurgeUrlsCacheRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PurgeUrlsCacheRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PurgeUrlsCacheResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Purge task ID. The first ten digits are the UTC time when the task is submitted.
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PurgeUrlsCacheResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PurgeUrlsCacheResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Quota struct {

	// Quota limit for one batch submission request.
	Batch *int64 `json:"Batch,omitempty" name:"Batch"`

	// Daily submission quota limit.
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// Remaining daily submission quota.
	Available *int64 `json:"Available,omitempty" name:"Available"`
}

type ResourceData struct {

	// Resource name, which is categorized as follows based on different query conditions:
	// Specific domain name: indicates the details of the specific domain name
	// multiDomains: indicates aggregated details of multiple domain names
	// Project ID: displays the ID of the specified project to be queried
	// all: details at the account level
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// Data details of resource
	EcdnData *EcdnData `json:"EcdnData,omitempty" name:"EcdnData"`
}

type ResponseHeader struct {

	// Custom response header switch. Valid values: on, off.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// Custom response header rule array.
	// Note: this field may return null, indicating that no valid values can be obtained.
	HeaderRules []*HttpHeaderPathRule `json:"HeaderRules,omitempty" name:"HeaderRules" list`
}

type ServerCert struct {

	// Server certificate ID, which is required if the certificate is a Tencent Cloud-hosted certificate.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// Server certificate name, which is required if the certificate is a Tencent Cloud-hosted certificate.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// Server certificate information, which is required when uploading your own certificate and must contain complete certificate chain information.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Certificate *string `json:"Certificate,omitempty" name:"Certificate"`

	// Server key information, which is required when uploading your own certificate.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`

	// Certificate expiration time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// Certificate issuance time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DeployTime *string `json:"DeployTime,omitempty" name:"DeployTime"`

	// Certificate remarks.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`
}

type Sort struct {

	// Sort by field. Valid values:
	// createTime: domain name creation time
	// certExpireTime: certificate expiration time
	Key *string `json:"Key,omitempty" name:"Key"`

	// asc/desc. Default value: desc.
	Sequence *string `json:"Sequence,omitempty" name:"Sequence"`
}

type StartEcdnDomainRequest struct {
	*tchttp.BaseRequest

	// Domain name to be enabled.
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *StartEcdnDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartEcdnDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartEcdnDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartEcdnDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartEcdnDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopEcdnDomainRequest struct {
	*tchttp.BaseRequest

	// Domain name to be disabled.
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *StopEcdnDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopEcdnDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopEcdnDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopEcdnDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopEcdnDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TimestampData struct {

	// Statistical time point in forward rounding mode
	// Taking the 5-minute granularity as an example, 13:35:00 indicates that the statistical interval is between 13:35:00 and 13:39:59
	Time *string `json:"Time,omitempty" name:"Time"`

	// Data value
	Value []*float64 `json:"Value,omitempty" name:"Value" list`
}

type UpdateDomainConfigRequest struct {
	*tchttp.BaseRequest

	// Domain name.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Origin server configuration.
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// IP blacklist/whitelist configuration.
	IpFilter *IpFilter `json:"IpFilter,omitempty" name:"IpFilter"`

	// IP access limit configuration.
	IpFreqLimit *IpFreqLimit `json:"IpFreqLimit,omitempty" name:"IpFreqLimit"`

	// Origin server response header configuration.
	ResponseHeader *ResponseHeader `json:"ResponseHeader,omitempty" name:"ResponseHeader"`

	// Node caching configuration.
	CacheKey *CacheKey `json:"CacheKey,omitempty" name:"CacheKey"`

	// Caching rule configuration.
	Cache *Cache `json:"Cache,omitempty" name:"Cache"`

	// HTTPS configuration.
	Https *Https `json:"Https,omitempty" name:"Https"`

	// Forced access protocol redirection configuration.
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitempty" name:"ForceRedirect"`

	// Domain name acceleration region. Valid values: mainland (acceleration in Mainland China), overseas (acceleration outside Mainland China), global (global acceleration).
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *UpdateDomainConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateDomainConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateDomainConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateDomainConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateDomainConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
