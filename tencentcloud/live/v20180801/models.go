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

package v20180801

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AddDelayLiveStreamRequest struct {
	*tchttp.BaseRequest

	// Push path, which is the same as the `AppName` in push and playback addresses and is `live` by default.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Delay time in seconds, up to 600s.
	DelayTime *uint64 `json:"DelayTime,omitempty" name:"DelayTime"`

	// Expiration time of the configured delayed playback in UTC format, such as 2018-11-29T19:00:00Z.
	// Notes:
	// 1. The configuration will expire after 7 days by default and can last up to 7 days.
	// 2. The Beijing time is in UTC+8. This value should be in the format as required by ISO 8601. For more information, please see [ISO Date and Time Format](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

func (r *AddDelayLiveStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddDelayLiveStreamRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddDelayLiveStreamResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddDelayLiveStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddDelayLiveStreamResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddLiveDomainRequest struct {
	*tchttp.BaseRequest

	// Domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Domain name type.
	// 0: push domain name.
	// 1: playback domain name.
	DomainType *uint64 `json:"DomainType,omitempty" name:"DomainType"`

	// Pull domain name type:
	// 1: Mainland China.
	// 2: global.
	// 3: outside Mainland China.
	// Default value: 1.
	PlayType *uint64 `json:"PlayType,omitempty" name:"PlayType"`

	// Whether it is LCB:
	// 0: LVB,
	// 1: LCB.
	// Default value: 0.
	IsDelayLive *int64 `json:"IsDelayLive,omitempty" name:"IsDelayLive"`

	// Whether it is LVB on Mini Program.
	// 0: LVB.
	// 1: LVB on Mini Program.
	// Default value: 0.
	IsMiniProgramLive *int64 `json:"IsMiniProgramLive,omitempty" name:"IsMiniProgramLive"`
}

func (r *AddLiveDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddLiveDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddLiveDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddLiveDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddLiveDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddLiveWatermarkRequest struct {
	*tchttp.BaseRequest

	// Watermark image URL.
	PictureUrl *string `json:"PictureUrl,omitempty" name:"PictureUrl"`

	// Watermark name.
	// Up to 16 bytes.
	WatermarkName *string `json:"WatermarkName,omitempty" name:"WatermarkName"`

	// Display position: X-axis offset in %. Default value: 0.
	XPosition *int64 `json:"XPosition,omitempty" name:"XPosition"`

	// Display position: Y-axis offset in %. Default value: 0.
	YPosition *int64 `json:"YPosition,omitempty" name:"YPosition"`

	// Watermark width or its percentage of the live streaming video width. It is recommended to just specify either height or width as the other will be scaled proportionally to avoid distortions. The original width is used by default.
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Watermark height or its percentage of the live streaming video width. It is recommended to just specify either height or width as the other will be scaled proportionally to avoid distortions. The original height is used by default.
	Height *int64 `json:"Height,omitempty" name:"Height"`
}

func (r *AddLiveWatermarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddLiveWatermarkRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddLiveWatermarkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Watermark ID.
		WatermarkId *uint64 `json:"WatermarkId,omitempty" name:"WatermarkId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddLiveWatermarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddLiveWatermarkResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BillDataInfo struct {

	// Time point in the format of `yyyy-mm-dd HH:MM:SS`.
	Time *string `json:"Time,omitempty" name:"Time"`

	// Bandwidth in Mbps.
	Bandwidth *float64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// Traffic in MB.
	Flux *float64 `json:"Flux,omitempty" name:"Flux"`

	// Time point of peak value in the format of `yyyy-mm-dd HH:MM:SS`. As raw data is at a 5-minute granularity, if data at a 1-hour or 1-day granularity is queried, the time point of peak bandwidth value at the corresponding granularity will be returned.
	PeakTime *string `json:"PeakTime,omitempty" name:"PeakTime"`
}

type BindLiveDomainCertRequest struct {
	*tchttp.BaseRequest

	// Certificate ID, which can be obtained through the `CreateLiveCert` API.
	CertId *int64 `json:"CertId,omitempty" name:"CertId"`

	// Playback domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// HTTPS status. 0: disabled, 1: enabled.
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *BindLiveDomainCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindLiveDomainCertRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindLiveDomainCertResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindLiveDomainCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindLiveDomainCertResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CallBackRuleInfo struct {

	// Rule creation time.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Rule update time.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Push path.
	AppName *string `json:"AppName,omitempty" name:"AppName"`
}

type CallBackTemplateInfo struct {

	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// Template name.
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// Description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Stream starting callback URL.
	StreamBeginNotifyUrl *string `json:"StreamBeginNotifyUrl,omitempty" name:"StreamBeginNotifyUrl"`

	// Interruption callback URL.
	StreamEndNotifyUrl *string `json:"StreamEndNotifyUrl,omitempty" name:"StreamEndNotifyUrl"`

	// Stream mixing callback URL.
	StreamMixNotifyUrl *string `json:"StreamMixNotifyUrl,omitempty" name:"StreamMixNotifyUrl"`

	// Recording callback URL.
	RecordNotifyUrl *string `json:"RecordNotifyUrl,omitempty" name:"RecordNotifyUrl"`

	// Screencapturing callback URL.
	SnapshotNotifyUrl *string `json:"SnapshotNotifyUrl,omitempty" name:"SnapshotNotifyUrl"`

	// Porn detection callback URL.
	PornCensorshipNotifyUrl *string `json:"PornCensorshipNotifyUrl,omitempty" name:"PornCensorshipNotifyUrl"`

	// Callback authentication key.
	CallbackKey *string `json:"CallbackKey,omitempty" name:"CallbackKey"`
}

type CancelCommonMixStreamRequest struct {
	*tchttp.BaseRequest

	// ID of stream mix session (from applying for stream mix to canceling stream mix).
	// This value is the same as the `MixStreamSessionId` in `CreateCommonMixStream`.
	MixStreamSessionId *string `json:"MixStreamSessionId,omitempty" name:"MixStreamSessionId"`
}

func (r *CancelCommonMixStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CancelCommonMixStreamRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CancelCommonMixStreamResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelCommonMixStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CancelCommonMixStreamResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CdnPlayStatData struct {

	// Time point in the format of `yyyy-mm-dd HH:MM:SS`.
	Time *string `json:"Time,omitempty" name:"Time"`

	// Bandwidth in Mbps.
	Bandwidth *float64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// Traffic in MB.
	Flux *float64 `json:"Flux,omitempty" name:"Flux"`

	// Number of new requests.
	Request *uint64 `json:"Request,omitempty" name:"Request"`

	// Number of concurrent connections.
	Online *uint64 `json:"Online,omitempty" name:"Online"`
}

type CertInfo struct {

	// Certificate ID.
	CertId *int64 `json:"CertId,omitempty" name:"CertId"`

	// Certificate name.
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// Description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Creation time in UTC format.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Certificate content.
	HttpsCrt *string `json:"HttpsCrt,omitempty" name:"HttpsCrt"`

	// Certificate type:
	// 0: Tencent Cloud-hosted certificate.
	// 1: user-added certificate.
	CertType *int64 `json:"CertType,omitempty" name:"CertType"`

	// Certificate expiration time in UTC format.
	CertExpireTime *string `json:"CertExpireTime,omitempty" name:"CertExpireTime"`

	// List of domain names that use this certificate.
	DomainList []*string `json:"DomainList,omitempty" name:"DomainList" list`
}

type ClientIpPlaySumInfo struct {

	// Client IP in dotted-decimal notation.
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// District where the client is located.
	Province *string `json:"Province,omitempty" name:"Province"`

	// Total traffic.
	TotalFlux *float64 `json:"TotalFlux,omitempty" name:"TotalFlux"`

	// Total number of requests.
	TotalRequest *uint64 `json:"TotalRequest,omitempty" name:"TotalRequest"`

	// Total number of failed requests.
	TotalFailedRequest *uint64 `json:"TotalFailedRequest,omitempty" name:"TotalFailedRequest"`

	// Country/region where the client is located.
	CountryArea *string `json:"CountryArea,omitempty" name:"CountryArea"`
}

type CommonMixControlParams struct {

	// Valid values: [0,1].
	// If 1 is entered, when the layer resolution in the parameter is different from the actual video resolution, the video will be automatically cropped according to the resolution set by the layer.
	UseMixCropCenter *int64 `json:"UseMixCropCenter,omitempty" name:"UseMixCropCenter"`
}

type CommonMixCropParams struct {

	// Crop width. Value range: [0,2000].
	CropWidth *float64 `json:"CropWidth,omitempty" name:"CropWidth"`

	// Crop height. Value range: [0,2000].
	CropHeight *float64 `json:"CropHeight,omitempty" name:"CropHeight"`

	// Starting crop X coordinate. Value range: [0,2000].
	CropStartLocationX *float64 `json:"CropStartLocationX,omitempty" name:"CropStartLocationX"`

	// Starting crop Y coordinate. Value range: [0,2000].
	CropStartLocationY *float64 `json:"CropStartLocationY,omitempty" name:"CropStartLocationY"`
}

type CommonMixInputParam struct {

	// Input stream name of up to 80 bytes, which is a string containing letters, digits, and underscores.
	InputStreamName *string `json:"InputStreamName,omitempty" name:"InputStreamName"`

	// Input stream layout parameter.
	LayoutParams *CommonMixLayoutParams `json:"LayoutParams,omitempty" name:"LayoutParams"`

	// Input stream crop parameter.
	CropParams *CommonMixCropParams `json:"CropParams,omitempty" name:"CropParams"`
}

type CommonMixLayoutParams struct {

	// Input layer. Value range: [1,16].
	// 1) For `image_layer` of background stream (i.e., main host video image or canvas), enter 1.
	// 2) For audio stream mix, this parameter is also required.
	ImageLayer *int64 `json:"ImageLayer,omitempty" name:"ImageLayer"`

	// Input type. Value range: [0,5].
	// If this parameter is left empty, 0 will be used by default.
	// 0: the input stream is audio/video.
	// 2: the input stream is image.
	// 3: the input stream is canvas. 
	// 4: the input stream is audio.
	// 5: the input stream is pure video.
	InputType *int64 `json:"InputType,omitempty" name:"InputType"`

	// Output width of input video image. Value range:
	// Pixel: [0,2000]
	// Percentage: [0.01,0.99]
	// If this parameter is left empty, the input stream width will be used by default.
	// If percentage is used, the expected output is (percentage * background width).
	ImageWidth *float64 `json:"ImageWidth,omitempty" name:"ImageWidth"`

	// Output height of input video image. Value range:
	// Pixel: [0,2000]
	// Percentage: [0.01,0.99]
	// If this parameter is left empty, the input stream height will be used by default.
	// If percentage is used, the expected output is (percentage * background height).
	ImageHeight *float64 `json:"ImageHeight,omitempty" name:"ImageHeight"`

	// X-axis offset of input in output video image. Value range:
	// Pixel: [0,2000]
	// Percentage: [0.01,0.99]
	// If this parameter is left empty, 0 will be used by default.
	// Horizontal offset from the top-left corner of main host background video image. 
	// If percentage is used, the expected output is (percentage * background width).
	LocationX *float64 `json:"LocationX,omitempty" name:"LocationX"`

	// Y-axis offset of input in output video image. Value range:
	// Pixel: [0,2000]
	// Percentage: [0.01,0.99]
	// If this parameter is left empty, 0 will be used by default.
	// Vertical offset from the top-left corner of main host background video image. 
	// If percentage is used, the expected output is (percentage * background width)
	LocationY *float64 `json:"LocationY,omitempty" name:"LocationY"`

	// When `InputType` is 3 (canvas), this value indicates the canvas color.
	// Commonly used colors include:
	// Red: 0xcc0033.
	// Yellow: 0xcc9900.
	// Green: 0xcccc33.
	// Blue: 0x99CCFF.
	// Black: 0x000000.
	// White: 0xFFFFFF.
	// Gray: 0x999999
	Color *string `json:"Color,omitempty" name:"Color"`

	// When `InputType` is 2 (image), this value is the watermark ID.
	WatermarkId *int64 `json:"WatermarkId,omitempty" name:"WatermarkId"`
}

type CommonMixOutputParams struct {

	// Output stream name.
	OutputStreamName *string `json:"OutputStreamName,omitempty" name:"OutputStreamName"`

	// Output stream type. Valid values: [0,1].
	// If this parameter is left empty, 0 will be used by default.
	// If the output stream is a stream in the input stream list, enter 0.
	// If you want the stream mix result to be a new stream, enter 1.
	// If this value is 1, `output_stream_id` cannot appear in `input_stram_list`, and there cannot be a stream with the same ID on the LVB backend.
	OutputStreamType *int64 `json:"OutputStreamType,omitempty" name:"OutputStreamType"`

	// Output stream bitrate. Value range: [1,50000].
	// If this parameter is left empty, the system will automatically determine.
	OutputStreamBitRate *int64 `json:"OutputStreamBitRate,omitempty" name:"OutputStreamBitRate"`

	// Output stream GOP size. Value range: [1,10].
	// If this parameter is left empty, the system will automatically determine.
	OutputStreamGop *int64 `json:"OutputStreamGop,omitempty" name:"OutputStreamGop"`

	// Output stream frame rate. Value range: [1,60].
	// If this parameter is left empty, the system will automatically determine.
	OutputStreamFrameRate *int64 `json:"OutputStreamFrameRate,omitempty" name:"OutputStreamFrameRate"`

	// Output stream audio bitrate. Value range: [1,500]
	// If this parameter is left empty, the system will automatically determine.
	OutputAudioBitRate *int64 `json:"OutputAudioBitRate,omitempty" name:"OutputAudioBitRate"`

	// Output stream audio sample rate. Valid values: [96000, 88200, 64000, 48000, 44100, 32000,24000, 22050, 16000, 12000, 11025, 8000].
	// If this parameter is left empty, the system will automatically determine.
	OutputAudioSampleRate *int64 `json:"OutputAudioSampleRate,omitempty" name:"OutputAudioSampleRate"`

	// Output stream audio sound channel. Valid values: [1,2].
	// If this parameter is left empty, the system will automatically determine.
	OutputAudioChannels *int64 `json:"OutputAudioChannels,omitempty" name:"OutputAudioChannels"`

	// SEI information in output stream. If there are no special needs, leave it empty.
	MixSei *string `json:"MixSei,omitempty" name:"MixSei"`
}

type ConcurrentRecordStreamNum struct {

	// Time point.
	Time *string `json:"Time,omitempty" name:"Time"`

	// Number of channels.
	Num *uint64 `json:"Num,omitempty" name:"Num"`
}

type CreateCommonMixStreamRequest struct {
	*tchttp.BaseRequest

	// ID of stream mix session (from applying for stream mix to canceling stream mix).
	MixStreamSessionId *string `json:"MixStreamSessionId,omitempty" name:"MixStreamSessionId"`

	// Input stream list for stream mix.
	InputStreamList []*CommonMixInputParam `json:"InputStreamList,omitempty" name:"InputStreamList" list`

	// Output stream parameter for stream mix.
	OutputParams *CommonMixOutputParams `json:"OutputParams,omitempty" name:"OutputParams"`

	// Input template ID. If this parameter is set, the output will be generated according to the default template layout, and there is no need to enter the custom position parameters.
	// If this parameter is left empty, 0 will be used by default.
	// For two input sources, 10, 20, 30, 40, and 50 are supported.
	// For three input sources, 310, 390, and 391 are supported.
	// For four input sources, 410 is supported.
	// For five input sources, 510 and 590 are supported.
	// For six input sources, 610 is supported.
	MixStreamTemplateId *int64 `json:"MixStreamTemplateId,omitempty" name:"MixStreamTemplateId"`

	// Special control parameter for stream mix. If there are no special needs, leave it empty.
	ControlParams *CommonMixControlParams `json:"ControlParams,omitempty" name:"ControlParams"`
}

func (r *CreateCommonMixStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateCommonMixStreamRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateCommonMixStreamResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCommonMixStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateCommonMixStreamResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveCallbackRuleRequest struct {
	*tchttp.BaseRequest

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Push path, which is the same as the `AppName` in push and playback addresses and is `live` by default.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *CreateLiveCallbackRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveCallbackRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveCallbackRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveCallbackRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveCallbackRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveCallbackTemplateRequest struct {
	*tchttp.BaseRequest

	// Template name.
	// Maximum length: 255 bytes.
	// Only letters, digits, underscores, and hyphens can be contained.
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// Description.
	// Maximum length: 1,024 bytes.
	// Only letters, digits, underscores, and hyphens can be contained.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Stream starting callback URL,
	// Protocol document: [Event Message Notification](/document/product/267/32744).
	StreamBeginNotifyUrl *string `json:"StreamBeginNotifyUrl,omitempty" name:"StreamBeginNotifyUrl"`

	// Interruption callback URL,
	// Protocol document: [Event Message Notification](/document/product/267/32744).
	StreamEndNotifyUrl *string `json:"StreamEndNotifyUrl,omitempty" name:"StreamEndNotifyUrl"`

	// Recording callback URL,
	// Protocol document: [Event Message Notification](/document/product/267/32744).
	RecordNotifyUrl *string `json:"RecordNotifyUrl,omitempty" name:"RecordNotifyUrl"`

	// Screencapturing callback URL,
	// Protocol document: [Event Message Notification](/document/product/267/32744).
	SnapshotNotifyUrl *string `json:"SnapshotNotifyUrl,omitempty" name:"SnapshotNotifyUrl"`

	// Porn detection callback URL,
	// Protocol document: [Event Message Notification](/document/product/267/32741).
	PornCensorshipNotifyUrl *string `json:"PornCensorshipNotifyUrl,omitempty" name:"PornCensorshipNotifyUrl"`

	// Callback key. The callback URL is public. For the callback signature, please see the event message notification document.
	// [Event Message Notification](/document/product/267/32744).
	CallbackKey *string `json:"CallbackKey,omitempty" name:"CallbackKey"`
}

func (r *CreateLiveCallbackTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveCallbackTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveCallbackTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Template ID.
		TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveCallbackTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveCallbackTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveCertRequest struct {
	*tchttp.BaseRequest

	// Certificate type. 0: user-added certificate, 1: Tencent Cloud-hosted certificate.
	// Note: if the certificate type is 0, `HttpsCrt` and `HttpsKey` are required;
	// If the certificate type is 1, the certificate corresponding to `CloudCertId` will be used first. If `CloudCertId` is empty, `HttpsCrt` and `HttpsKey` will be used.
	CertType *uint64 `json:"CertType,omitempty" name:"CertType"`

	// Certificate name.
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// Certificate content, i.e., public key.
	HttpsCrt *string `json:"HttpsCrt,omitempty" name:"HttpsCrt"`

	// Private key.
	HttpsKey *string `json:"HttpsKey,omitempty" name:"HttpsKey"`

	// Description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Tencent Cloud-hosted certificate ID.
	CloudCertId *string `json:"CloudCertId,omitempty" name:"CloudCertId"`
}

func (r *CreateLiveCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveCertRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveCertResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Certificate ID
		CertId *int64 `json:"CertId,omitempty" name:"CertId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveCertResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveRecordRequest struct {
	*tchttp.BaseRequest

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Push path, which is the same as the `AppName` in push and playback addresses and is `live` by default.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Push domain name. This parameter must be set for multi-domain name push.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Recording start time, which is China standard time and should be URL-encoded (RFC3986). For example, the encoding of 2017-01-01 10:10:01 is 2017-01-01+10%3a10%3a01.
	// In scheduled recording mode, this field must be set; in real-time video recording mode, this field is ignored.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Recording end time, which is China standard time and should be URL-encoded (RFC3986). For example, the encoding of 2017-01-01 10:30:01 is 2017-01-01+10%3a30%3a01.
	// In scheduled recording mode, this field must be set; in real-time video recording mode, this field is optional. If the recording is set to real-time video recording mode through the `Highlight` parameter, the set end time should not be more than 30 minutes after the current time. If the set end time is more than 30 minutes after the current time, earlier than the current time, or left empty, the actual end time will be 30 minutes after the current time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Recording type.
	// "video": Audio-video recording **(default)**.
	// "audio": audio recording.
	// In both scheduled and real-time video recording modes, this parameter is valid and is not case sensitive.
	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`

	// Recording file format. Valid values:
	// "flv" **(default)**, "hls", "mp4", "aac", "mp3".
	// In both scheduled and real-time video recording modes, this parameter is valid and is not case sensitive.
	FileFormat *string `json:"FileFormat,omitempty" name:"FileFormat"`

	// Mark for enabling real-time video recording mode.
	// 0: Real-time video recording mode is not enabled, i.e., the scheduled recording mode is used **(default)**. See [Sample 1](#.E7.A4.BA.E4.BE.8B1-.E5.88.9B.E5.BB.BA.E5.AE.9A.E6.97.B6.E5.BD.95.E5.88.B6.E4.BB.BB.E5.8A.A1).
	// 1: Real-time video recording mode is enabled. See [Sample 2](#.E7.A4.BA.E4.BE.8B2-.E5.88.9B.E5.BB.BA.E5.AE.9E.E6.97.B6.E5.BD.95.E5.88.B6.E4.BB.BB.E5.8A.A1).
	Highlight *int64 `json:"Highlight,omitempty" name:"Highlight"`

	// Flag for enabling A+B=C mixed stream recording.
	// 0: A+B=C mixed stream recording is not enabled **(default)**.
	// 1: A+B=C mixed stream recording is enabled.
	// In both scheduled and real-time video recording modes, this parameter is valid.
	MixStream *int64 `json:"MixStream,omitempty" name:"MixStream"`

	// Recording stream parameter. The following parameters are supported currently:
	// record_interval: recording interval in seconds. Value range: 1800-7200.
	// storage_time: recording file storage duration in seconds.
	// Example: record_interval=3600&storage_time=2592000.
	// Note: the parameter needs to be URL-encoded.
	// In both scheduled and real-time video recording modes, this parameter is valid.
	StreamParam *string `json:"StreamParam,omitempty" name:"StreamParam"`
}

func (r *CreateLiveRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveRecordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task ID, which uniquely identifies a recording task globally.
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveRecordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveRecordRuleRequest struct {
	*tchttp.BaseRequest

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// Push path, which is the same as the AppName in push and playback addresses and is "live" by default.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Stream name.
	// Note: If the parameter is a non-empty string, the rule will be only applicable to the particular stream.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *CreateLiveRecordRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveRecordRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveRecordRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveRecordRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveRecordRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveRecordTemplateRequest struct {
	*tchttp.BaseRequest

	// Template name. Only letters, digits, underscores, and hyphens can be contained.
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// Message description
	Description *string `json:"Description,omitempty" name:"Description"`

	// FLV recording parameter, which is set when FLV recording is enabled.
	FlvParam *RecordParam `json:"FlvParam,omitempty" name:"FlvParam"`

	// HLS recording parameter, which is set when HLS recording is enabled.
	HlsParam *RecordParam `json:"HlsParam,omitempty" name:"HlsParam"`

	// Mp4 recording parameter, which is set when Mp4 recording is enabled.
	Mp4Param *RecordParam `json:"Mp4Param,omitempty" name:"Mp4Param"`

	// AAC recording parameter, which is set when AAC recording is enabled.
	AacParam *RecordParam `json:"AacParam,omitempty" name:"AacParam"`

	// LVB type. Default value: 0.
	// 0: LVB.
	// 1: LCB.
	IsDelayLive *int64 `json:"IsDelayLive,omitempty" name:"IsDelayLive"`

	// HLS-specific recording parameter.
	HlsSpecialParam *HlsSpecialParam `json:"HlsSpecialParam,omitempty" name:"HlsSpecialParam"`

	// Mp3 recording parameter, which is set when Mp3 recording is enabled.
	Mp3Param *RecordParam `json:"Mp3Param,omitempty" name:"Mp3Param"`
}

func (r *CreateLiveRecordTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveRecordTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveRecordTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Template ID.
		TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveRecordTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveRecordTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveSnapshotRuleRequest struct {
	*tchttp.BaseRequest

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// Push path, which is the same as the `AppName` in push and playback addresses and is `live` by default.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Stream name.
	// Note: if this parameter is a non-empty string, the rule will take effect only for the particular stream.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *CreateLiveSnapshotRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveSnapshotRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveSnapshotRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveSnapshotRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveSnapshotRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveSnapshotTemplateRequest struct {
	*tchttp.BaseRequest

	// Template name.
	// Maximum length: 255 bytes.
	// Only letters, digits, underscores, and hyphens can be contained.
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// COS application ID.
	CosAppId *int64 `json:"CosAppId,omitempty" name:"CosAppId"`

	// COS bucket name.
	CosBucket *string `json:"CosBucket,omitempty" name:"CosBucket"`

	// COS region.
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// Description.
	// Maximum length: 1,024 bytes.
	// Only letters, digits, underscores, and hyphens can be contained.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Screencapturing interval in seconds. Default value: 10s.
	// Value range: 5–600s.
	SnapshotInterval *int64 `json:"SnapshotInterval,omitempty" name:"SnapshotInterval"`

	// Screenshot width. Default value: 0 (original width).
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Screenshot height. Default value: 0 (original height).
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// Whether to enable porn detection. 0: no, 1: yes. Default value: 0
	PornFlag *int64 `json:"PornFlag,omitempty" name:"PornFlag"`

	// COS bucket folder prefix.
	CosPrefix *string `json:"CosPrefix,omitempty" name:"CosPrefix"`

	// COS filename.
	CosFileName *string `json:"CosFileName,omitempty" name:"CosFileName"`
}

func (r *CreateLiveSnapshotTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveSnapshotTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveSnapshotTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Template ID.
		TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveSnapshotTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveSnapshotTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveTranscodeRuleRequest struct {
	*tchttp.BaseRequest

	// Playback domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Push path, which is the same as the `AppName` in push and playback addresses and is `live` by default. If you only bind a domain name, leave this parameter empty.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Stream name. If only the domain name or path is bound, leave this parameter blank.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Designates an existing template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *CreateLiveTranscodeRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveTranscodeRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveTranscodeRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveTranscodeRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveTranscodeRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveTranscodeTemplateRequest struct {
	*tchttp.BaseRequest

	// Template name, such as 900 900p. This can be only a combination of letters and digits.
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// Video bitrate. Value range: 100–8,000.
	// Note: The bitrate must be a multiple of 100.
	VideoBitrate *int64 `json:"VideoBitrate,omitempty" name:"VideoBitrate"`

	// Video encoding format. Valid values: h264, h265. Default value: h264.
	Vcodec *string `json:"Vcodec,omitempty" name:"Vcodec"`

	// Audio encoding in ACC format. Default value: original audio format.
	// Note: This parameter will take effect later.
	Acodec *string `json:"Acodec,omitempty" name:"Acodec"`

	// Audio bitrate. Value range: 0–500. Default value: 0.
	AudioBitrate *int64 `json:"AudioBitrate,omitempty" name:"AudioBitrate"`

	// Template description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Width. Default value: 0.
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Whether to keep the video. 0: no; 1: yes. Default value: 1.
	NeedVideo *int64 `json:"NeedVideo,omitempty" name:"NeedVideo"`

	// Whether to keep the audio. 0: no; 1: yes. Default value: 1.
	NeedAudio *int64 `json:"NeedAudio,omitempty" name:"NeedAudio"`

	// Height. Default value: 0.
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// Frame rate. Default value: 0.
	Fps *int64 `json:"Fps,omitempty" name:"Fps"`

	// Keyframe interval in seconds. Original interval by default
	Gop *int64 `json:"Gop,omitempty" name:"Gop"`

	// Whether to rotate. 0: no; 1: yes. Default value: 0.
	Rotate *int64 `json:"Rotate,omitempty" name:"Rotate"`

	// Encoding quality:
	// baseline/main/high. Default value: baseline.
	Profile *string `json:"Profile,omitempty" name:"Profile"`

	// Whether to not exceed the original bitrate. 0: no; 1: yes. Default value: 0.
	BitrateToOrig *int64 `json:"BitrateToOrig,omitempty" name:"BitrateToOrig"`

	// Whether to not exceed the original height. 0: no; 1: yes. Default value: 0.
	HeightToOrig *int64 `json:"HeightToOrig,omitempty" name:"HeightToOrig"`

	// Whether to not exceed the original frame rate. 0: no; 1: yes. Default value: 0.
	FpsToOrig *int64 `json:"FpsToOrig,omitempty" name:"FpsToOrig"`

	// Whether it is a top speed codec template. 0: no, 1: yes. Default value: 0.
	AiTransCode *int64 `json:"AiTransCode,omitempty" name:"AiTransCode"`

	// `VideoBitrate` minus top speed codec bitrate. Value range: 0.1–0.5.
	AdaptBitratePercent *float64 `json:"AdaptBitratePercent,omitempty" name:"AdaptBitratePercent"`
}

func (r *CreateLiveTranscodeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveTranscodeTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveTranscodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Template ID.
		TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveTranscodeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveTranscodeTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveWatermarkRuleRequest struct {
	*tchttp.BaseRequest

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Push path, which is the same as the `AppName` in push and playback addresses and is `live` by default.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Watermark ID, which is the `WatermarkId` returned by the [AddLiveWatermark](/document/product/267/30154) API.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *CreateLiveWatermarkRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveWatermarkRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveWatermarkRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveWatermarkRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveWatermarkRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRecordTaskRequest struct {
	*tchttp.BaseRequest

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Push path.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Recording task end time in UNIX timestamp, which must be after `StartTime` and within 24 hours from the current time.
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// Recording task start time in UNIX timestamp. If this parameter is left empty, it indicates to start recording immediately. It must be within 24 hours from the current time.
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// Push type. Default value: 0. Valid values:
	// 0: LVB push.
	// 1: mixed stream, i.e., A + B = C mixed stream.
	StreamType *uint64 `json:"StreamType,omitempty" name:"StreamType"`

	// Recording template ID, which is the returned value of `CreateLiveRecordTemplate`. If this parameter is left empty or incorrect, the stream will be recorded in HLS format and retained permanently by default.
	TemplateId *uint64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// Extended field, which is empty by default.
	Extension *string `json:"Extension,omitempty" name:"Extension"`
}

func (r *CreateRecordTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRecordTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRecordTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task ID, which uniquely identifies the recording task globally.
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRecordTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRecordTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DayStreamPlayInfo struct {

	// Data point in time in the format of `yyyy-mm-dd HH:MM:SS`.
	Time *string `json:"Time,omitempty" name:"Time"`

	// Bandwidth in Mbps.
	Bandwidth *float64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// Traffic in MB.
	Flux *float64 `json:"Flux,omitempty" name:"Flux"`

	// Number of requests.
	Request *uint64 `json:"Request,omitempty" name:"Request"`

	// Number of online viewers.
	Online *uint64 `json:"Online,omitempty" name:"Online"`
}

type DelayInfo struct {

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Push path, which is the same as the 
	//  `AppName` in push and playback addresses and is `live` by default.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Delay time in seconds.
	DelayInterval *uint64 `json:"DelayInterval,omitempty" name:"DelayInterval"`

	// Creation time in UTC time.
	// Note: the difference between UTC time and Beijing time is 8 hours.
	// Example: 2019-06-18T12:00:00Z (i.e., June 18, 2019 20:00:00 Beijing time).
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Expiration time in UTC time.
	// Note: the difference between UTC time and Beijing time is 8 hours.
	// Example: 2019-06-18T12:00:00Z (i.e., June 18, 2019 20:00:00 Beijing time).
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// Current status:
	// -1: expired.
	// 1: in effect.
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type DeleteLiveCallbackRuleRequest struct {
	*tchttp.BaseRequest

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Push path, which is the same as the `AppName` in push and playback addresses and is `live` by default.
	AppName *string `json:"AppName,omitempty" name:"AppName"`
}

func (r *DeleteLiveCallbackRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveCallbackRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveCallbackRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveCallbackRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveCallbackRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveCallbackTemplateRequest struct {
	*tchttp.BaseRequest

	// Template ID.
	// 1. Get the template ID in the returned value of the [CreateLiveCallbackTemplate](/document/product/267/32637) API call.
	// 2. You can query the list of created templates through the [DescribeLiveCallbackTemplates](/document/product/267/32632) API.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DeleteLiveCallbackTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveCallbackTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveCallbackTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveCallbackTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveCallbackTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveCertRequest struct {
	*tchttp.BaseRequest

	// Certificate ID obtained through the `DescribeLiveCerts` API.
	CertId *int64 `json:"CertId,omitempty" name:"CertId"`
}

func (r *DeleteLiveCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveCertRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveCertResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveCertResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveDomainRequest struct {
	*tchttp.BaseRequest

	// Domain name to be deleted.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Type. 0: push, 1: playback.
	DomainType *uint64 `json:"DomainType,omitempty" name:"DomainType"`
}

func (r *DeleteLiveDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveRecordRequest struct {
	*tchttp.BaseRequest

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Task ID returned by the `CreateLiveRecord` API.
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DeleteLiveRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveRecordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveRecordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveRecordRuleRequest struct {
	*tchttp.BaseRequest

	// Push domain name.
	// Domain name+AppName+StreamName uniquely identifies a single transcoding rule. If you need to delete it, strong match is required. For example, even if AppName is blank, you need to pass in a blank string to make a strong match.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Push path, which is the same as the AppName in push and playback addresses and is "live" by default.
	// Domain name+AppName+StreamName uniquely identifies a single transcoding rule. If you need to delete it, strong match is required. For example, even if AppName is blank, you need to pass in a blank string to make a strong match.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Stream name.
	// Domain name+AppName+StreamName uniquely identifies a single transcoding rule. If you need to delete it, strong match is required. For example, even if AppName is blank, you need to pass in a blank string to make a strong match.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *DeleteLiveRecordRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveRecordRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveRecordRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveRecordRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveRecordRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveRecordTemplateRequest struct {
	*tchttp.BaseRequest

	// Template ID obtained through the `DescribeRecordTemplates` API.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DeleteLiveRecordTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveRecordTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveRecordTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveRecordTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveRecordTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveSnapshotRuleRequest struct {
	*tchttp.BaseRequest

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Push path, which is the same as the `AppName` in push and playback addresses and is `live` by default.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *DeleteLiveSnapshotRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveSnapshotRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveSnapshotRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveSnapshotRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveSnapshotRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveSnapshotTemplateRequest struct {
	*tchttp.BaseRequest

	// Template ID.
	// 1. Get from the returned value of the [CreateLiveSnapshotTemplate](/document/product/267/32624) API call.
	// 2. You can query the list of created screencapturing templates through the [DescribeLiveSnapshotTemplates](/document/product/267/32619) API.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DeleteLiveSnapshotTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveSnapshotTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveSnapshotTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveSnapshotTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveSnapshotTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveTranscodeRuleRequest struct {
	*tchttp.BaseRequest

	// Playback domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Push path, which is the same as the `AppName` in push and playback addresses and is `live` by default.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DeleteLiveTranscodeRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveTranscodeRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveTranscodeRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveTranscodeRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveTranscodeRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveTranscodeTemplateRequest struct {
	*tchttp.BaseRequest

	// Template ID.
	// 1. Get the template ID in the returned value of the [CreateLiveTranscodeTemplate](/document/product/267/32646) API call.
	// 2. You can query the list of created templates through the [DescribeLiveTranscodeTemplates](/document/product/267/32641) API.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DeleteLiveTranscodeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveTranscodeTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveTranscodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveTranscodeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveTranscodeTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveWatermarkRequest struct {
	*tchttp.BaseRequest

	// Watermark ID.
	// Watermark ID obtained in the returned value of the [AddLiveWatermark](/document/product/267/30154) API call.
	// Watermark ID returned by the `DescribeLiveWatermarks` API.
	WatermarkId *int64 `json:"WatermarkId,omitempty" name:"WatermarkId"`
}

func (r *DeleteLiveWatermarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveWatermarkRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveWatermarkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveWatermarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveWatermarkResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveWatermarkRuleRequest struct {
	*tchttp.BaseRequest

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Push path, which is the same as the `AppName` in push and playback addresses and is `live` by default.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *DeleteLiveWatermarkRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveWatermarkRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveWatermarkRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveWatermarkRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveWatermarkRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRecordTaskRequest struct {
	*tchttp.BaseRequest

	// Task ID returned by `CreateRecordTask`. The recording task specified by `TaskId` will be deleted.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DeleteRecordTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRecordTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRecordTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRecordTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRecordTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAllStreamPlayInfoListRequest struct {
	*tchttp.BaseRequest

	// Query time accurate down to the minute in the format of `yyyy-mm-dd HH:MM:SS`. Data for the last month can be queried. The data has a delay of about 5 minutes; therefore, if you want to query real-time data, we recommend you pass in a point in time 5 minutes ago.
	QueryTime *string `json:"QueryTime,omitempty" name:"QueryTime"`
}

func (r *DescribeAllStreamPlayInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAllStreamPlayInfoListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAllStreamPlayInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Query point in time in the returned input parameters.
		QueryTime *string `json:"QueryTime,omitempty" name:"QueryTime"`

		// Data information list.
		DataInfoList []*MonitorStreamPlayInfo `json:"DataInfoList,omitempty" name:"DataInfoList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllStreamPlayInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAllStreamPlayInfoListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillBandwidthAndFluxListRequest struct {
	*tchttp.BaseRequest

	// Start time point in the format of `yyyy-mm-dd HH:MM:SS`.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time point in the format of `yyyy-mm-dd HH:MM:SS`. The difference between the start time and end time cannot be greater than 31 days.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// LVB playback domain name. If this parameter is left empty, full data will be queried.
	PlayDomains []*string `json:"PlayDomains,omitempty" name:"PlayDomains" list`

	// Valid values:
	// Mainland: query data for Mainland China,
	// Oversea: query data for regions outside Mainland China,
	// Default: query data for all regions.
	// Note: LEB only supports querying data for all regions.
	MainlandOrOversea *string `json:"MainlandOrOversea,omitempty" name:"MainlandOrOversea"`

	// Data granularity. Valid values:
	// 5: 5-minute granularity (the query time span should be within 1 day),
	// 60: 1-hour granularity (the query time span should be within one month),
	// 1440: 1-day granularity (the query time span should be within one month).
	// Default value: 5.
	Granularity *uint64 `json:"Granularity,omitempty" name:"Granularity"`

	// Service name. Valid values: LVB, LEB. Default value: LVB.
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

func (r *DescribeBillBandwidthAndFluxListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillBandwidthAndFluxListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillBandwidthAndFluxListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Time point of peak bandwidth value in the format of `yyyy-mm-dd HH:MM:SS`.
		PeakBandwidthTime *string `json:"PeakBandwidthTime,omitempty" name:"PeakBandwidthTime"`

		// Peak bandwidth in Mbps.
		PeakBandwidth *float64 `json:"PeakBandwidth,omitempty" name:"PeakBandwidth"`

		// Time point of 95th percentile bandwidth value in the format of `yyyy-mm-dd HH:MM:SS`.
		P95PeakBandwidthTime *string `json:"P95PeakBandwidthTime,omitempty" name:"P95PeakBandwidthTime"`

		// 95th percentile bandwidth in Mbps.
		P95PeakBandwidth *float64 `json:"P95PeakBandwidth,omitempty" name:"P95PeakBandwidth"`

		// Total traffic in MB.
		SumFlux *float64 `json:"SumFlux,omitempty" name:"SumFlux"`

		// Detailed data information.
		DataInfoList []*BillDataInfo `json:"DataInfoList,omitempty" name:"DataInfoList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillBandwidthAndFluxListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillBandwidthAndFluxListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConcurrentRecordStreamNumRequest struct {
	*tchttp.BaseRequest

	// Live streaming type. SlowLive: LCB.
	// NormalLive: LVB.
	LiveType *string `json:"LiveType,omitempty" name:"LiveType"`

	// Start time in the format of `yyyy-mm-dd HH:MM:SS`.
	// Data for the last 180 days can be queried.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time in the format of `yyyy-mm-dd HH:MM:SS`.
	// The maximum time span supported is 31 days.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Valid values: Mainland (data for Mainland China), Oversea (data for regions outside Mainland China). If this parameter is left empty, data for all regions will be queried.
	MainlandOrOversea *string `json:"MainlandOrOversea,omitempty" name:"MainlandOrOversea"`

	// Playback domain name list. If this parameter is left empty, full data will be queried.
	PushDomains []*string `json:"PushDomains,omitempty" name:"PushDomains" list`
}

func (r *DescribeConcurrentRecordStreamNumRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConcurrentRecordStreamNumRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConcurrentRecordStreamNumResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Statistics list.
		DataInfoList []*ConcurrentRecordStreamNum `json:"DataInfoList,omitempty" name:"DataInfoList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConcurrentRecordStreamNumResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConcurrentRecordStreamNumResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupProIspPlayInfoListRequest struct {
	*tchttp.BaseRequest

	// Start time point in the format of `yyyy-mm-dd HH:MM:SS`.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time point in the format of `yyyy-mm-dd HH:MM:SS`
	// The time span is (0,3 hours]. Data for the last month can be queried.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Playback domain name. If this parameter is left empty, full data will be queried.
	PlayDomains []*string `json:"PlayDomains,omitempty" name:"PlayDomains" list`

	// District list. If this parameter is left empty, data for all districts will be returned.
	ProvinceNames []*string `json:"ProvinceNames,omitempty" name:"ProvinceNames" list`

	// ISP list. If this parameter is left empty, data of all ISPs will be returned.
	IspNames []*string `json:"IspNames,omitempty" name:"IspNames" list`

	// Within or outside Mainland China. Valid values: Mainland (data for Mainland China), Oversea (data for regions outside Mainland China). If this parameter is left empty, data for all regions will be queried.
	MainlandOrOversea *string `json:"MainlandOrOversea,omitempty" name:"MainlandOrOversea"`
}

func (r *DescribeGroupProIspPlayInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupProIspPlayInfoListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupProIspPlayInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Data content.
		DataInfoList []*GroupProIspDataInfo `json:"DataInfoList,omitempty" name:"DataInfoList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupProIspPlayInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupProIspPlayInfoListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeHttpStatusInfoListRequest struct {
	*tchttp.BaseRequest

	// Start time (Beijing time),
	// In the format of `yyyy-mm-dd HH:MM:SS`.
	// `StartTime` cannot be more than 3 months ago.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time (Beijing time),
	// In the format of `yyyy-mm-dd HH:MM:SS`.
	// Note: `EndTime` and `StartTime` only support querying data for the last day.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Playback domain name list.
	PlayDomains []*string `json:"PlayDomains,omitempty" name:"PlayDomains" list`
}

func (r *DescribeHttpStatusInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeHttpStatusInfoListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeHttpStatusInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Playback status code list.
		DataInfoList []*HttpStatusData `json:"DataInfoList,omitempty" name:"DataInfoList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHttpStatusInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeHttpStatusInfoListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveCallbackRulesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLiveCallbackRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveCallbackRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveCallbackRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Rule information list.
		Rules []*CallBackRuleInfo `json:"Rules,omitempty" name:"Rules" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveCallbackRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveCallbackRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveCallbackTemplateRequest struct {
	*tchttp.BaseRequest

	// Template ID.
	// 1. Get the template ID in the returned value of the [CreateLiveCallbackTemplate](/document/product/267/32637) API call.
	// 2. You can query the list of created templates through the [DescribeLiveCallbackTemplates](/document/product/267/32632) API.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DescribeLiveCallbackTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveCallbackTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveCallbackTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Callback template information.
		Template *CallBackTemplateInfo `json:"Template,omitempty" name:"Template"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveCallbackTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveCallbackTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveCallbackTemplatesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLiveCallbackTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveCallbackTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveCallbackTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Template information list.
		Templates []*CallBackTemplateInfo `json:"Templates,omitempty" name:"Templates" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveCallbackTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveCallbackTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveCertRequest struct {
	*tchttp.BaseRequest

	// Certificate ID obtained through the `DescribeLiveCerts` API.
	CertId *int64 `json:"CertId,omitempty" name:"CertId"`
}

func (r *DescribeLiveCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveCertRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveCertResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Certificate information.
		CertInfo *CertInfo `json:"CertInfo,omitempty" name:"CertInfo"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveCertResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveCertsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLiveCertsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveCertsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveCertsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Certificate information list.
		CertInfoSet []*CertInfo `json:"CertInfoSet,omitempty" name:"CertInfoSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveCertsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveCertsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveDelayInfoListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLiveDelayInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveDelayInfoListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveDelayInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Delayed playback information list.
		DelayInfoList []*DelayInfo `json:"DelayInfoList,omitempty" name:"DelayInfoList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveDelayInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveDelayInfoListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveDomainCertRequest struct {
	*tchttp.BaseRequest

	// Playback domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
}

func (r *DescribeLiveDomainCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveDomainCertRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveDomainCertResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Certificate information.
		DomainCertInfo *DomainCertInfo `json:"DomainCertInfo,omitempty" name:"DomainCertInfo"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveDomainCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveDomainCertResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveDomainPlayInfoListRequest struct {
	*tchttp.BaseRequest

	// Playback domain name list.
	PlayDomains []*string `json:"PlayDomains,omitempty" name:"PlayDomains" list`
}

func (r *DescribeLiveDomainPlayInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveDomainPlayInfoListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveDomainPlayInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Data time in the format of `yyyy-mm-dd HH:MM:SS`.
		Time *string `json:"Time,omitempty" name:"Time"`

		// Real-time total bandwidth.
		TotalBandwidth *float64 `json:"TotalBandwidth,omitempty" name:"TotalBandwidth"`

		// Real-time total traffic.
		TotalFlux *float64 `json:"TotalFlux,omitempty" name:"TotalFlux"`

		// Total number of requests.
		TotalRequest *uint64 `json:"TotalRequest,omitempty" name:"TotalRequest"`

		// Real-time total number of connections.
		TotalOnline *uint64 `json:"TotalOnline,omitempty" name:"TotalOnline"`

		// Data by domain name.
		DomainInfoList []*DomainInfoList `json:"DomainInfoList,omitempty" name:"DomainInfoList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveDomainPlayInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveDomainPlayInfoListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveDomainRequest struct {
	*tchttp.BaseRequest

	// Domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
}

func (r *DescribeLiveDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Domain name information.
		DomainInfo *DomainInfo `json:"DomainInfo,omitempty" name:"DomainInfo"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveDomainsRequest struct {
	*tchttp.BaseRequest

	// Filter by domain name status. 0: disabled, 1: enabled.
	DomainStatus *uint64 `json:"DomainStatus,omitempty" name:"DomainStatus"`

	// Filter by domain name type. 0: push. 1: playback
	DomainType *uint64 `json:"DomainType,omitempty" name:"DomainType"`

	// Number of entries per page. Value range: 10–100. Default value: 10.
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// Page number to get. Value range: 1–100000. Default value: 1.
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// 0: LVB, 1: LCB. Default value: 0.
	IsDelayLive *uint64 `json:"IsDelayLive,omitempty" name:"IsDelayLive"`

	// Domain name prefix.
	DomainPrefix *string `json:"DomainPrefix,omitempty" name:"DomainPrefix"`
}

func (r *DescribeLiveDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveDomainsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveDomainsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of results.
		AllCount *uint64 `json:"AllCount,omitempty" name:"AllCount"`

		// List of domain name details.
		DomainList []*DomainInfo `json:"DomainList,omitempty" name:"DomainList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveDomainsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveForbidStreamListRequest struct {
	*tchttp.BaseRequest

	// Page number to get. Default value: 1.
	PageNum *int64 `json:"PageNum,omitempty" name:"PageNum"`

	// Number of entries per page. Maximum value: 100. 
	// Value: any integer between 1 and 100.
	// Default value: 10.
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeLiveForbidStreamListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveForbidStreamListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveForbidStreamListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of eligible ones.
		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// Total number of pages.
		TotalPage *int64 `json:"TotalPage,omitempty" name:"TotalPage"`

		// Page number.
		PageNum *int64 `json:"PageNum,omitempty" name:"PageNum"`

		// Number of entries displayed per page.
		PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

		// List of forbidden streams.
		ForbidStreamList []*ForbidStreamInfo `json:"ForbidStreamList,omitempty" name:"ForbidStreamList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveForbidStreamListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveForbidStreamListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLivePlayAuthKeyRequest struct {
	*tchttp.BaseRequest

	// Domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
}

func (r *DescribeLivePlayAuthKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLivePlayAuthKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLivePlayAuthKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Playback authentication key information.
		PlayAuthKeyInfo *PlayAuthKeyInfo `json:"PlayAuthKeyInfo,omitempty" name:"PlayAuthKeyInfo"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLivePlayAuthKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLivePlayAuthKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLivePushAuthKeyRequest struct {
	*tchttp.BaseRequest

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
}

func (r *DescribeLivePushAuthKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLivePushAuthKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLivePushAuthKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Push authentication key information.
		PushAuthKeyInfo *PushAuthKeyInfo `json:"PushAuthKeyInfo,omitempty" name:"PushAuthKeyInfo"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLivePushAuthKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLivePushAuthKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveRecordRulesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLiveRecordRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveRecordRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveRecordRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of rules.
		Rules []*RuleInfo `json:"Rules,omitempty" name:"Rules" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveRecordRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveRecordRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveRecordTemplateRequest struct {
	*tchttp.BaseRequest

	// Template ID obtained through the `DescribeRecordTemplates` API.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DescribeLiveRecordTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveRecordTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveRecordTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Recording template information.
		Template *RecordTemplateInfo `json:"Template,omitempty" name:"Template"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveRecordTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveRecordTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveRecordTemplatesRequest struct {
	*tchttp.BaseRequest

	// Whether it is an LCB template. Default value: 0.
	// 0: LVB.
	// 1: LCB.
	IsDelayLive *int64 `json:"IsDelayLive,omitempty" name:"IsDelayLive"`
}

func (r *DescribeLiveRecordTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveRecordTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveRecordTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Recording template information list.
		Templates []*RecordTemplateInfo `json:"Templates,omitempty" name:"Templates" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveRecordTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveRecordTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveSnapshotRulesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLiveSnapshotRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveSnapshotRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveSnapshotRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Rule list.
		Rules []*RuleInfo `json:"Rules,omitempty" name:"Rules" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveSnapshotRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveSnapshotRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveSnapshotTemplateRequest struct {
	*tchttp.BaseRequest

	// Template ID.
	// Template ID returned by the [CreateLiveSnapshotTemplate](/document/product/267/32624) API call.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DescribeLiveSnapshotTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveSnapshotTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveSnapshotTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Screencapturing template information.
		Template *SnapshotTemplateInfo `json:"Template,omitempty" name:"Template"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveSnapshotTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveSnapshotTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveSnapshotTemplatesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLiveSnapshotTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveSnapshotTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveSnapshotTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Screencapturing template list.
		Templates []*SnapshotTemplateInfo `json:"Templates,omitempty" name:"Templates" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveSnapshotTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveSnapshotTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamEventListRequest struct {
	*tchttp.BaseRequest

	// Start time. 
	// In UTC format, such as 2018-12-29T19:00:00Z.
	// This supports querying the history of 60 days.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time.
	// In UTC format, such as 2018-12-29T20:00:00Z.
	// This cannot be after the current time and cannot be more than 30 days after the start time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Push path, which is the same as the AppName in push and playback addresses and is "live" by default.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Stream name; query with wildcard (*) is not supported; fuzzy match by default.
	// The IsStrict field can be used to change to exact query.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Page number to get.
	// Default value: 1.
	// Note: Currently, query for up to 10,000 entries is supported.
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// Number of entries per page.
	// Maximum value: 100.
	// Value range: any integer between 1 and 100.
	// Default value: 10.
	// Note: currently, query for up to 10,000 entries is supported.
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// Whether to filter. No filtering by default.
	// 0: No filtering at all.
	// 1: Filter out the failing streams and return only the successful ones.
	IsFilter *int64 `json:"IsFilter,omitempty" name:"IsFilter"`

	// Whether to query exactly. Fuzzy match by default.
	// 0: Fuzzy match.
	// 1: Exact query.
	// Note: This parameter takes effect when StreamName is used.
	IsStrict *int64 `json:"IsStrict,omitempty" name:"IsStrict"`

	// Whether to display in ascending order by end time. Descending order by default.
	// 0: Descending.
	// 1: Ascending.
	IsAsc *int64 `json:"IsAsc,omitempty" name:"IsAsc"`
}

func (r *DescribeLiveStreamEventListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveStreamEventListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamEventListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of streaming events.
		EventList []*StreamEventInfo `json:"EventList,omitempty" name:"EventList" list`

		// Page number.
		PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

		// Number of entries per page.
		PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

		// Total number of eligible ones.
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// Total number of pages.
		TotalPage *uint64 `json:"TotalPage,omitempty" name:"TotalPage"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveStreamEventListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveStreamEventListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamOnlineListRequest struct {
	*tchttp.BaseRequest

	// Push domain name. If you use multiple paths, enter the `DomainName`.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Push path, which is the same as the `AppName` in push and playback addresses and is `live` by default. If you use multiple paths, enter the `AppName`.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Page number to get. Default value: 1.
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// Number of entries per page. Maximum value: 100. 
	// Value: any integer between 10 and 100.
	// Default value: 10.
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// Stream name, which is used for exact query.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *DescribeLiveStreamOnlineListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveStreamOnlineListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamOnlineListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of eligible ones.
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// Total number of pages.
		TotalPage *uint64 `json:"TotalPage,omitempty" name:"TotalPage"`

		// Page number.
		PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

		// Number of entries displayed per page.
		PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

		// Active push information list.
		OnlineInfo []*StreamOnlineInfo `json:"OnlineInfo,omitempty" name:"OnlineInfo" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveStreamOnlineListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveStreamOnlineListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamPublishedListRequest struct {
	*tchttp.BaseRequest

	// Your push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// End time.
	// In UTC format, such as 2016-06-30T19:00:00Z.
	// This cannot be after the current time.
	// Note: The difference between EndTime and StartTime cannot be greater than 30 days.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Start time. 
	// In UTC format, such as 2016-06-29T19:00:00Z.
	// This supports querying data in the past 60 days.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Push path, which is the same as the `AppName` in push and playback addresses and is `live` by default. Fuzzy match is not supported.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Page number to get.
	// Default value: 1.
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// Number of entries per page.
	// Maximum value: 100.
	// Value range: any integer between 1 and 100.
	// Default value: 10.
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// Stream name, which supports fuzzy match.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *DescribeLiveStreamPublishedListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveStreamPublishedListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamPublishedListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Push record information.
		PublishInfo []*StreamName `json:"PublishInfo,omitempty" name:"PublishInfo" list`

		// Page number.
		PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

		// Number of entries per page
		PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

		// Total number of eligible ones.
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// Total number of pages.
		TotalPage *uint64 `json:"TotalPage,omitempty" name:"TotalPage"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveStreamPublishedListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveStreamPublishedListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamPushInfoListRequest struct {
	*tchttp.BaseRequest

	// Push domain name.
	PushDomain *string `json:"PushDomain,omitempty" name:"PushDomain"`

	// Push path, which is the same as the `AppName` in push and playback addresses and is `live` by default.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Number of pages,
	// Value range: [1,10000],
	// Default value: 1.
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// Number of entries per page,
	// Value range: [1,1000],
	// Default value: 200.
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeLiveStreamPushInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveStreamPushInfoListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamPushInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Live stream statistics list.
		DataInfoList []*PushDataInfo `json:"DataInfoList,omitempty" name:"DataInfoList" list`

		// Total number of live streams.
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// Total number of pages.
		TotalPage *uint64 `json:"TotalPage,omitempty" name:"TotalPage"`

		// Page number where the current data resides.
		PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

		// Number of live streams per page.
		PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveStreamPushInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveStreamPushInfoListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamStateRequest struct {
	*tchttp.BaseRequest

	// Push path, which is the same as the AppName in push and playback addresses and is "live" by default.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Your push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *DescribeLiveStreamStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveStreamStateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamStateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Stream status,
	// active: active
	// inactive: Inactive
	// forbid: forbidden.
		StreamState *string `json:"StreamState,omitempty" name:"StreamState"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveStreamStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveStreamStateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveTranscodeDetailInfoRequest struct {
	*tchttp.BaseRequest

	// Push domain name.
	PushDomain *string `json:"PushDomain,omitempty" name:"PushDomain"`

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Start time (Beijing time).
	// In the format of `yyyymmdd`.
	// Note: details for a specified day in the last month can be queried.
	DayTime *string `json:"DayTime,omitempty" name:"DayTime"`

	// Number of pages. Default value: 1.
	// Up to 100 pages.
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// Number of entries per page. Default value: 20,
	// Value range: [10,1000].
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// Start day time (Beijing time),
	// In the format of `yyyymmdd`.
	// Note: details for the last month can be queried.
	StartDayTime *string `json:"StartDayTime,omitempty" name:"StartDayTime"`

	// End day time (Beijing time),
	// In the format of `yyyymmdd`.
	// Note: detailed data for the last month can be queried. Either `DayTime` or `(StartDayTime,EndDayTime)` must be passed in. If both are passed in, `DayTime` shall prevail.
	EndDayTime *string `json:"EndDayTime,omitempty" name:"EndDayTime"`
}

func (r *DescribeLiveTranscodeDetailInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveTranscodeDetailInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveTranscodeDetailInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Statistics list.
		DataInfoList []*TranscodeDetailInfo `json:"DataInfoList,omitempty" name:"DataInfoList" list`

		// Page number.
		PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

		// Number of entries per page.
		PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

		// Total number.
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// Total number of pages.
		TotalPage *uint64 `json:"TotalPage,omitempty" name:"TotalPage"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveTranscodeDetailInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveTranscodeDetailInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveTranscodeRulesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLiveTranscodeRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveTranscodeRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveTranscodeRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of transcoding rules.
		Rules []*RuleInfo `json:"Rules,omitempty" name:"Rules" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveTranscodeRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveTranscodeRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveTranscodeTemplateRequest struct {
	*tchttp.BaseRequest

	// Template ID.
	// Note: get the template ID in the returned value of the [CreateLiveTranscodeTemplate](/document/product/267/32646) API call.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DescribeLiveTranscodeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveTranscodeTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveTranscodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Template information.
		Template *TemplateInfo `json:"Template,omitempty" name:"Template"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveTranscodeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveTranscodeTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveTranscodeTemplatesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLiveTranscodeTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveTranscodeTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveTranscodeTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of transcoding templates.
		Templates []*TemplateInfo `json:"Templates,omitempty" name:"Templates" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveTranscodeTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveTranscodeTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveWatermarkRequest struct {
	*tchttp.BaseRequest

	// Watermark ID returned by the `DescribeLiveWatermarks` API.
	WatermarkId *uint64 `json:"WatermarkId,omitempty" name:"WatermarkId"`
}

func (r *DescribeLiveWatermarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveWatermarkRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveWatermarkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Watermark information.
		Watermark *WatermarkInfo `json:"Watermark,omitempty" name:"Watermark"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveWatermarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveWatermarkResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveWatermarkRulesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLiveWatermarkRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveWatermarkRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveWatermarkRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Watermarking rule list.
		Rules []*RuleInfo `json:"Rules,omitempty" name:"Rules" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveWatermarkRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveWatermarkRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveWatermarksRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLiveWatermarksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveWatermarksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveWatermarksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of watermarks.
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// Watermark information list.
		WatermarkList []*WatermarkInfo `json:"WatermarkList,omitempty" name:"WatermarkList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveWatermarksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveWatermarksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePlayErrorCodeDetailInfoListRequest struct {
	*tchttp.BaseRequest

	// Start time (Beijing time),
	// In the format of `yyyy-mm-dd HH:MM:SS`.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time (Beijing time),
	// In the format of `yyyy-mm-dd HH:MM:SS`.
	// Note: `EndTime` and `StartTime` only support querying data for the last day.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Query granularity:
	// 1: 1-minute granularity.
	Granularity *uint64 `json:"Granularity,omitempty" name:"Granularity"`

	// Yes. Valid values: "4xx", "5xx". Mixed codes in the format of `4xx,5xx` are also supported.
	StatType *string `json:"StatType,omitempty" name:"StatType"`

	// Playback domain name list.
	PlayDomains []*string `json:"PlayDomains,omitempty" name:"PlayDomains" list`

	// Region. Valid values: Mainland (data for Mainland China), Oversea (data for regions outside Mainland China), China (data for China, including Hong Kong, Macao, and Taiwan), Foreign (data for regions outside China, excluding Hong Kong, Macao, and Taiwan), Global (default). If this parameter is left empty, data for all regions will be queried.
	MainlandOrOversea *string `json:"MainlandOrOversea,omitempty" name:"MainlandOrOversea"`
}

func (r *DescribePlayErrorCodeDetailInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePlayErrorCodeDetailInfoListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePlayErrorCodeDetailInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Statistics list.
		HttpCodeList []*HttpCodeInfo `json:"HttpCodeList,omitempty" name:"HttpCodeList" list`

		// Statistics type.
		StatType *string `json:"StatType,omitempty" name:"StatType"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePlayErrorCodeDetailInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePlayErrorCodeDetailInfoListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePlayErrorCodeSumInfoListRequest struct {
	*tchttp.BaseRequest

	// Start point in time (Beijing time).
	// In the format of `yyyy-mm-dd HH:MM:SS`.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End point in time (Beijing time).
	// In the format of `yyyy-mm-dd HH:MM:SS`.
	// Note: `EndTime` and `StartTime` only support querying data for the last day.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Playback domain name list. If this parameter is left empty, full data will be queried.
	PlayDomains []*string `json:"PlayDomains,omitempty" name:"PlayDomains" list`

	// Number of pages. Value range: [1,1000]. Default value: 1.
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// Number of entries per page. Value range: [1,1000]. Default value: 20.
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// Region. Valid values: Mainland (data for Mainland China), Oversea (data for regions outside Mainland China), China (data for China, including Hong Kong, Macao, and Taiwan), Foreign (data for regions outside China, excluding Hong Kong, Macao, and Taiwan), Global (default). If this parameter is left empty, data for all regions will be queried.
	MainlandOrOversea *string `json:"MainlandOrOversea,omitempty" name:"MainlandOrOversea"`

	// Grouping parameter. Valid values: CountryProIsp (default value), Country (country/region). Grouping is made by country/region + district + ISP by default. Currently, districts and ISPs outside Mainland China cannot be recognized.
	GroupType *string `json:"GroupType,omitempty" name:"GroupType"`

	// Language used in the output field. Valid values: Chinese (default), English. Currently, country/region, district, and ISP parameters support multiple languages.
	OutLanguage *string `json:"OutLanguage,omitempty" name:"OutLanguage"`
}

func (r *DescribePlayErrorCodeSumInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePlayErrorCodeSumInfoListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePlayErrorCodeSumInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Information of 4xx or 5xx error codes by district and ISP.
		ProIspInfoList []*ProIspPlayCodeDataInfo `json:"ProIspInfoList,omitempty" name:"ProIspInfoList" list`

		// Total occurrences of all status codes.
		TotalCodeAll *uint64 `json:"TotalCodeAll,omitempty" name:"TotalCodeAll"`

		// Occurrences of 4xx status codes.
		TotalCode4xx *uint64 `json:"TotalCode4xx,omitempty" name:"TotalCode4xx"`

		// Occurrences of 5xx status codes.
		TotalCode5xx *uint64 `json:"TotalCode5xx,omitempty" name:"TotalCode5xx"`

		// Total occurrences of each status code.
		TotalCodeList []*PlayCodeTotalInfo `json:"TotalCodeList,omitempty" name:"TotalCodeList" list`

		// Page number.
		PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

		// Number of entries per page.
		PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

		// Total number of pages.
		TotalPage *uint64 `json:"TotalPage,omitempty" name:"TotalPage"`

		// Total number of results.
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// Occurrences of 2xx status codes.
		TotalCode2xx *uint64 `json:"TotalCode2xx,omitempty" name:"TotalCode2xx"`

		// Occurrences of 3xx status codes.
		TotalCode3xx *uint64 `json:"TotalCode3xx,omitempty" name:"TotalCode3xx"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePlayErrorCodeSumInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePlayErrorCodeSumInfoListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProIspPlaySumInfoListRequest struct {
	*tchttp.BaseRequest

	// Start time (Beijing time).
	// In the format of `yyyy-mm-dd HH:MM:SS`.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time (Beijing time).
	// In the format of `yyyy-mm-dd HH:MM:SS`.
	// Note: `EndTime` and `StartTime` only support querying data for the last day.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Statistics type. Valid values: Province, Isp, CountryOrArea.
	StatType *string `json:"StatType,omitempty" name:"StatType"`

	// If this parameter is left empty, full data will be queried.
	PlayDomains []*string `json:"PlayDomains,omitempty" name:"PlayDomains" list`

	// Page number. Value range: [1,1000]. Default value: 1.
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// Number of entries per page. Value range: [1,1000]. Default value: 20.
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// Region. Valid values: Mainland (data for Mainland China), Oversea (data for regions outside Mainland China), China (data for China, including Hong Kong, Macao, and Taiwan), Foreign (data for regions outside China, excluding Hong Kong, Macao, and Taiwan), Global (default). If this parameter is left empty, data for all regions will be queried.
	MainlandOrOversea *string `json:"MainlandOrOversea,omitempty" name:"MainlandOrOversea"`

	// Language used in the output field. Valid values: Chinese (default), English. Currently, country/region, district, and ISP parameters support multiple languages.
	OutLanguage *string `json:"OutLanguage,omitempty" name:"OutLanguage"`
}

func (r *DescribeProIspPlaySumInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProIspPlaySumInfoListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProIspPlaySumInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total traffic.
		TotalFlux *float64 `json:"TotalFlux,omitempty" name:"TotalFlux"`

		// Total number of requests.
		TotalRequest *uint64 `json:"TotalRequest,omitempty" name:"TotalRequest"`

		// Statistics type.
		StatType *string `json:"StatType,omitempty" name:"StatType"`

		// Number of results per page.
		PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

		// Page number.
		PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

		// Total number of results.
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// Total number of pages.
		TotalPage *uint64 `json:"TotalPage,omitempty" name:"TotalPage"`

		// Aggregated data list by district, ISP, or country/region.
		DataInfoList []*ProIspPlaySumInfo `json:"DataInfoList,omitempty" name:"DataInfoList" list`

		// Download speed in MB/s. Calculation method: total traffic/total duration.
		AvgFluxPerSecond *float64 `json:"AvgFluxPerSecond,omitempty" name:"AvgFluxPerSecond"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProIspPlaySumInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProIspPlaySumInfoListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProvinceIspPlayInfoListRequest struct {
	*tchttp.BaseRequest

	// Start point in time (Beijing time).
	// Example: 2019-02-21 10:00:00.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End point in time (Beijing time).
	// Example: 2019-02-21 12:00:00.
	// Note: `EndTime` and `StartTime` only support querying data for the last day.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Supported granularities:
	// 1: 1-minute granularity (the query interval should be within 1 day)
	Granularity *uint64 `json:"Granularity,omitempty" name:"Granularity"`

	// Statistical metric type:
	// "Bandwidth": bandwidth
	// "FluxPerSecond": average traffic
	// "Flux": traffic
	// "Request": number of requests
	// "Online": number of concurrent connections
	StatType *string `json:"StatType,omitempty" name:"StatType"`

	// Playback domain name list.
	PlayDomains []*string `json:"PlayDomains,omitempty" name:"PlayDomains" list`

	// List of the districts to be queried, such as Beijing.
	ProvinceNames []*string `json:"ProvinceNames,omitempty" name:"ProvinceNames" list`

	// List of the ISPs to be queried, such as China Mobile. If this parameter is left empty, the data of all ISPs will be queried.
	IspNames []*string `json:"IspNames,omitempty" name:"IspNames" list`

	// Region. Valid values: Mainland (data for Mainland China), Oversea (data for regions outside Mainland China), China (data for China, including Hong Kong, Macao, and Taiwan), Foreign (data for regions outside China, excluding Hong Kong, Macao, and Taiwan), Global (default). If this parameter is left empty, data for all regions will be queried.
	MainlandOrOversea *string `json:"MainlandOrOversea,omitempty" name:"MainlandOrOversea"`
}

func (r *DescribeProvinceIspPlayInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProvinceIspPlayInfoListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProvinceIspPlayInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Playback information list.
		DataInfoList []*PlayStatInfo `json:"DataInfoList,omitempty" name:"DataInfoList" list`

		// Statistics type, which is the same as the input parameter.
		StatType *string `json:"StatType,omitempty" name:"StatType"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProvinceIspPlayInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProvinceIspPlayInfoListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeScreenShotSheetNumListRequest struct {
	*tchttp.BaseRequest

	// Start time in UTC time in the format of `yyyy-mm-ddTHH:MM:SSZ`.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time in UTC time in the format of `yyyy-mm-ddTHH:MM:SSZ`. Data for the last year can be queried.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Region information. Valid values: Mainland, Oversea. The former is to query data within Mainland China, while the latter outside Mainland China. If this parameter is left empty, data of all regions will be queried.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Push domain name (data at the domain name level after November 1, 2019 can be queried).
	PushDomains []*string `json:"PushDomains,omitempty" name:"PushDomains" list`

	// Data dimension. The data has a delay of one and a half hours. Valid values: 1. Minute (5-minute granularity, which supports a maximum query time range of 31 days); 2. Day (1-day granularity, which is the default value and supports a maximum query time range of 186 days).
	Granularity *string `json:"Granularity,omitempty" name:"Granularity"`
}

func (r *DescribeScreenShotSheetNumListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeScreenShotSheetNumListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeScreenShotSheetNumListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Data information list.
		DataInfoList []*TimeValue `json:"DataInfoList,omitempty" name:"DataInfoList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScreenShotSheetNumListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeScreenShotSheetNumListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeStreamDayPlayInfoListRequest struct {
	*tchttp.BaseRequest

	// Date in the format of `YYYY-mm-dd`.
	// Data is available at 3 AM the next day. You are recommended to query the latest data after this time point.
	DayTime *string `json:"DayTime,omitempty" name:"DayTime"`

	// Playback domain name.
	PlayDomain *string `json:"PlayDomain,omitempty" name:"PlayDomain"`

	// Page number. Value range: [1,1000]. Default value: 1.
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// Number of entries per page. Value range: [100,1000]. Default value: 1,000.
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeStreamDayPlayInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeStreamDayPlayInfoListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeStreamDayPlayInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Playback data information list.
		DataInfoList []*PlayDataInfoByStream `json:"DataInfoList,omitempty" name:"DataInfoList" list`

		// Total number.
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// Total number of pages.
		TotalPage *uint64 `json:"TotalPage,omitempty" name:"TotalPage"`

		// Page number where the current data resides.
		PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

		// Number of entries per page.
		PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStreamDayPlayInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeStreamDayPlayInfoListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeStreamPlayInfoListRequest struct {
	*tchttp.BaseRequest

	// Start time (Beijing time) in the format of `yyyy-mm-dd HH:MM:SS`,
	// The start time cannot be more than 30 days after the current time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time (Beijing time) in the format of `yyyy-mm-dd HH:MM:SS`.
	// The end time and start time must be on the same day.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Playback domain name,
	// If this parameter is left empty, data of live streams of all playback domain names will be queried.
	PlayDomain *string `json:"PlayDomain,omitempty" name:"PlayDomain"`

	// Stream name (exact match).
	// If this parameter is left empty, full playback data will be queried.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Push path, which is the same as the `AppName` in the playback address, subject to exact match, and valid if `StreamName` is passed in.
	// If this parameter is left empty, full playback data will be queried.
	// Note: to query by `AppName`, you need to submit a ticket for application.
	AppName *string `json:"AppName,omitempty" name:"AppName"`
}

func (r *DescribeStreamPlayInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeStreamPlayInfoListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeStreamPlayInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Statistics list at a 1-minute granularity.
		DataInfoList []*DayStreamPlayInfo `json:"DataInfoList,omitempty" name:"DataInfoList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStreamPlayInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeStreamPlayInfoListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeStreamPushInfoListRequest struct {
	*tchttp.BaseRequest

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Start time point in the format of `yyyy-mm-dd HH:MM:SS`.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time point in the format of `yyyy-mm-dd HH:MM:SS`. The maximum time span is 6 hours. Data for the last 6 days can be queried.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Push domain name.
	PushDomain *string `json:"PushDomain,omitempty" name:"PushDomain"`

	// Push path, which is the same as the `AppName` in push and playback addresses and is `live` by default.
	AppName *string `json:"AppName,omitempty" name:"AppName"`
}

func (r *DescribeStreamPushInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeStreamPushInfoListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeStreamPushInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned data list.
		DataInfoList []*PushQualityData `json:"DataInfoList,omitempty" name:"DataInfoList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStreamPushInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeStreamPushInfoListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTopClientIpSumInfoListRequest struct {
	*tchttp.BaseRequest

	// Start point in time in the format of `yyyy-mm-dd HH:MM:SS`.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End point in time in the format of `yyyy-mm-dd HH:MM:SS`
	// The time span is [0,4 hours]. Data for the last day can be queried.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Playback domain name. If this parameter is left empty, full data will be queried by default.
	PlayDomains []*string `json:"PlayDomains,omitempty" name:"PlayDomains" list`

	// Page number. Value range: [1,1000]. Default value: 1.
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// Number of entries per page. Value range: [1,1000]. Default value: 20.
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// Sorting metric. Valid values: TotalRequest (default value), FailedRequest, TotalFlux.
	OrderParam *string `json:"OrderParam,omitempty" name:"OrderParam"`

	// Region. Valid values: Mainland (data for Mainland China), Oversea (data for regions outside Mainland China), China (data for China, including Hong Kong, Macao, and Taiwan), Foreign (data for regions outside China, excluding Hong Kong, Macao, and Taiwan), Global (default). If this parameter is left empty, data for all regions will be queried.
	MainlandOrOversea *string `json:"MainlandOrOversea,omitempty" name:"MainlandOrOversea"`

	// Language used in the output field. Valid values: Chinese (default), English. Currently, country/region, district, and ISP parameters support multiple languages.
	OutLanguage *string `json:"OutLanguage,omitempty" name:"OutLanguage"`
}

func (r *DescribeTopClientIpSumInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTopClientIpSumInfoListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTopClientIpSumInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Page number. Value range: [1,1000]. Default value: 1.
		PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

		// Number of entries per page. Value range: [1,1000]. Default value: 20.
		PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

		// Sorting metric. Valid values: "TotalRequest", "FailedRequest", "TotalFlux".
		OrderParam *string `json:"OrderParam,omitempty" name:"OrderParam"`

		// Total number of results.
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// Total number of result pages.
		TotalPage *uint64 `json:"TotalPage,omitempty" name:"TotalPage"`

		// Data content.
		DataInfoList []*ClientIpPlaySumInfo `json:"DataInfoList,omitempty" name:"DataInfoList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopClientIpSumInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTopClientIpSumInfoListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVisitTopSumInfoListRequest struct {
	*tchttp.BaseRequest

	// Start point in time in the format of `yyyy-mm-dd HH:MM:SS`.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End point in time in the format of `yyyy-mm-dd HH:MM:SS`
	// The time span is (0,4 hours]. Data for the last day can be queried.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Bandwidth metric. Valid values: "Domain", "StreamId".
	TopIndex *string `json:"TopIndex,omitempty" name:"TopIndex"`

	// Playback domain name. If this parameter is left empty, full data will be queried by default.
	PlayDomains []*string `json:"PlayDomains,omitempty" name:"PlayDomains" list`

	// Page number,
	// Value range: [1,1000],
	// Default value: 1.
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// Number of entries per page. Value range: [1,1000].
	// Default value: 20.
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// Sorting metric. Valid values: "AvgFluxPerSecond", "TotalRequest" (default), "TotalFlux".
	OrderParam *string `json:"OrderParam,omitempty" name:"OrderParam"`
}

func (r *DescribeVisitTopSumInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVisitTopSumInfoListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVisitTopSumInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Page number,
	// Value range: [1,1000],
	// Default value: 1.
		PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

		// Number of entries per page. Value range: [1,1000].
	// Default value: 20.
		PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

		// Bandwidth metric. Valid values: "Domain", "StreamId".
		TopIndex *string `json:"TopIndex,omitempty" name:"TopIndex"`

		// Sorting metric. Valid values: "AvgFluxPerSecond", "TotalRequest" (default), "TotalFlux".
		OrderParam *string `json:"OrderParam,omitempty" name:"OrderParam"`

		// Total number of results.
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// Total number of result pages.
		TotalPage *uint64 `json:"TotalPage,omitempty" name:"TotalPage"`

		// Data content.
		DataInfoList []*PlaySumStatInfo `json:"DataInfoList,omitempty" name:"DataInfoList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVisitTopSumInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVisitTopSumInfoListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DomainCertInfo struct {

	// Certificate ID.
	CertId *int64 `json:"CertId,omitempty" name:"CertId"`

	// Certificate name.
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// Description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Creation time in UTC format.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Certificate content.
	HttpsCrt *string `json:"HttpsCrt,omitempty" name:"HttpsCrt"`

	// Certificate type.
	// 0: user-added certificate
	// 1: Tencent Cloud-hosted certificate.
	CertType *int64 `json:"CertType,omitempty" name:"CertType"`

	// Certificate expiration time in UTC format.
	CertExpireTime *string `json:"CertExpireTime,omitempty" name:"CertExpireTime"`

	// Domain name that uses this certificate.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Certificate status.
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type DomainDetailInfo struct {

	// In or outside Mainland China:
	// Mainland: data in Mainland China.
	// Oversea: data outside Mainland China.
	MainlandOrOversea *string `json:"MainlandOrOversea,omitempty" name:"MainlandOrOversea"`

	// Bandwidth in Mbps.
	Bandwidth *float64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// Traffic in MB.
	Flux *float64 `json:"Flux,omitempty" name:"Flux"`

	// Number of viewers.
	Online *uint64 `json:"Online,omitempty" name:"Online"`

	// Number of requests.
	Request *uint64 `json:"Request,omitempty" name:"Request"`
}

type DomainInfo struct {

	// LVB domain name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Domain name type:
	// 0: push.
	// 1: playback.
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// Domain name status:
	// 0: deactivated.
	// 1: activated.
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Creation time.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Whether there is a CNAME record pointing to a fixed rule domain name:
	// 0: no.
	// 1: yes.
	BCName *uint64 `json:"BCName,omitempty" name:"BCName"`

	// Domain name corresponding to CNAME record.
	TargetDomain *string `json:"TargetDomain,omitempty" name:"TargetDomain"`

	// Playback region. This parameter is valid only if `Type` is 1.
	// 1: in Mainland China.
	// 2: global.
	// 3: outside Mainland China.
	PlayType *int64 `json:"PlayType,omitempty" name:"PlayType"`

	// Whether it is LCB:
	// 0: LVB.
	// 1: LCB.
	IsDelayLive *int64 `json:"IsDelayLive,omitempty" name:"IsDelayLive"`

	// Information of currently used CNAME record.
	CurrentCName *string `json:"CurrentCName,omitempty" name:"CurrentCName"`

	// Disused parameter, which can be ignored.
	RentTag *int64 `json:"RentTag,omitempty" name:"RentTag"`

	// Disused parameter, which can be ignored.
	RentExpireTime *string `json:"RentExpireTime,omitempty" name:"RentExpireTime"`

	// 0: LVB.
	// 1: LVB on Mini Program.
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsMiniProgramLive *int64 `json:"IsMiniProgramLive,omitempty" name:"IsMiniProgramLive"`
}

type DomainInfoList struct {

	// Domain name.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Details.
	DetailInfoList []*DomainDetailInfo `json:"DetailInfoList,omitempty" name:"DetailInfoList" list`
}

type DropLiveStreamRequest struct {
	*tchttp.BaseRequest

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Your acceleration domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Push path, which is the same as the AppName in push and playback addresses and is "live" by default.
	AppName *string `json:"AppName,omitempty" name:"AppName"`
}

func (r *DropLiveStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DropLiveStreamRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DropLiveStreamResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DropLiveStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DropLiveStreamResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableLiveDomainRequest struct {
	*tchttp.BaseRequest

	// LVB domain name to be enabled.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
}

func (r *EnableLiveDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableLiveDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableLiveDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableLiveDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableLiveDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ForbidLiveDomainRequest struct {
	*tchttp.BaseRequest

	// LVB domain name to be disabled.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
}

func (r *ForbidLiveDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ForbidLiveDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ForbidLiveDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ForbidLiveDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ForbidLiveDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ForbidLiveStreamRequest struct {
	*tchttp.BaseRequest

	// Push path, which is the same as the AppName in push and playback addresses and is "live" by default.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Your push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Time to resume the stream in UTC format, such as 2018-11-29T19:00:00Z.
	// Notes:
	// 1. The duration of forbidding is 7 days by default and can be up to 90 days.
	// 2. The Beijing time is in UTC+8. This value should be in the format as required by ISO 8601. For more information, please see [ISO Date and Time Format](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F).
	ResumeTime *string `json:"ResumeTime,omitempty" name:"ResumeTime"`

	// Reason for forbidding.
	// Note: Be sure to enter the reason for forbidding to avoid any faulty operations.
	// Length limit: 2,048 bytes.
	Reason *string `json:"Reason,omitempty" name:"Reason"`
}

func (r *ForbidLiveStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ForbidLiveStreamRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ForbidLiveStreamResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ForbidLiveStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ForbidLiveStreamResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ForbidStreamInfo struct {

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Creation time.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Forbidding expiration time.
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

type GroupProIspDataInfo struct {

	// District.
	ProvinceName *string `json:"ProvinceName,omitempty" name:"ProvinceName"`

	// ISP.
	IspName *string `json:"IspName,omitempty" name:"IspName"`

	// Detailed data at the minute level.
	DetailInfoList []*CdnPlayStatData `json:"DetailInfoList,omitempty" name:"DetailInfoList" list`
}

type HlsSpecialParam struct {

	// HLS timeout period.
	FlowContinueDuration *uint64 `json:"FlowContinueDuration,omitempty" name:"FlowContinueDuration"`
}

type HttpCodeInfo struct {

	// HTTP return code.
	// Example: "2xx", "3xx", "4xx", "5xx".
	HttpCode *string `json:"HttpCode,omitempty" name:"HttpCode"`

	// Statistics. 0 will be added for points in time when there is no data.
	ValueList []*HttpCodeValue `json:"ValueList,omitempty" name:"ValueList" list`
}

type HttpCodeValue struct {

	// Time in the format of `yyyy-mm-dd HH:MM:SS`.
	Time *string `json:"Time,omitempty" name:"Time"`

	// Occurrences.
	Numbers *uint64 `json:"Numbers,omitempty" name:"Numbers"`

	// Proportion.
	Percentage *float64 `json:"Percentage,omitempty" name:"Percentage"`
}

type HttpStatusData struct {

	// Data point in time,
	// In the format of `yyyy-mm-dd HH:MM:SS`.
	Time *string `json:"Time,omitempty" name:"Time"`

	// Playback status code details.
	HttpStatusInfoList []*HttpStatusInfo `json:"HttpStatusInfoList,omitempty" name:"HttpStatusInfoList" list`
}

type HttpStatusInfo struct {

	// Playback HTTP status code.
	HttpStatus *string `json:"HttpStatus,omitempty" name:"HttpStatus"`

	// Quantity.
	Num *uint64 `json:"Num,omitempty" name:"Num"`
}

type ModifyLiveCallbackTemplateRequest struct {
	*tchttp.BaseRequest

	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// Template name.
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// Description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Stream starting callback URL.
	StreamBeginNotifyUrl *string `json:"StreamBeginNotifyUrl,omitempty" name:"StreamBeginNotifyUrl"`

	// Interruption callback URL.
	StreamEndNotifyUrl *string `json:"StreamEndNotifyUrl,omitempty" name:"StreamEndNotifyUrl"`

	// Recording callback URL.
	RecordNotifyUrl *string `json:"RecordNotifyUrl,omitempty" name:"RecordNotifyUrl"`

	// Screencapturing callback URL.
	SnapshotNotifyUrl *string `json:"SnapshotNotifyUrl,omitempty" name:"SnapshotNotifyUrl"`

	// Porn detection callback URL.
	PornCensorshipNotifyUrl *string `json:"PornCensorshipNotifyUrl,omitempty" name:"PornCensorshipNotifyUrl"`

	// Callback key. The callback URL is public. For the callback signature, please see the event message notification document.
	// [Event Message Notification](/document/product/267/32744).
	CallbackKey *string `json:"CallbackKey,omitempty" name:"CallbackKey"`
}

func (r *ModifyLiveCallbackTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLiveCallbackTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveCallbackTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLiveCallbackTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLiveCallbackTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveCertRequest struct {
	*tchttp.BaseRequest

	// Certificate ID.
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// Certificate type. 0: user-added certificate, 1: Tencent Cloud-hosted certificate.
	CertType *uint64 `json:"CertType,omitempty" name:"CertType"`

	// Certificate name.
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// Certificate content, i.e., public key.
	HttpsCrt *string `json:"HttpsCrt,omitempty" name:"HttpsCrt"`

	// Private key.
	HttpsKey *string `json:"HttpsKey,omitempty" name:"HttpsKey"`

	// Description.
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyLiveCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLiveCertRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveCertResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLiveCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLiveCertResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveDomainCertRequest struct {
	*tchttp.BaseRequest

	// Playback domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Certificate ID.
	CertId *int64 `json:"CertId,omitempty" name:"CertId"`

	// Status. 0: off, 1: on.
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyLiveDomainCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLiveDomainCertRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveDomainCertResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLiveDomainCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLiveDomainCertResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLivePlayAuthKeyRequest struct {
	*tchttp.BaseRequest

	// Playback domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Whether to enable. 0: disabled; 1: enabled.
	// If this parameter is left empty, the current value will not be modified.
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// Authentication key.
	// If this parameter is left empty, the current value will not be modified.
	AuthKey *string `json:"AuthKey,omitempty" name:"AuthKey"`

	// Validity period in seconds.
	// If this parameter is left empty, the current value will not be modified.
	AuthDelta *uint64 `json:"AuthDelta,omitempty" name:"AuthDelta"`

	// Backup authentication key.
	// If this parameter is left empty, the current value will not be modified.
	AuthBackKey *string `json:"AuthBackKey,omitempty" name:"AuthBackKey"`
}

func (r *ModifyLivePlayAuthKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLivePlayAuthKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLivePlayAuthKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLivePlayAuthKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLivePlayAuthKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLivePlayDomainRequest struct {
	*tchttp.BaseRequest

	// Playback domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Pull domain name type. 1: Mainland China. 2: global, 3: outside Mainland China
	PlayType *int64 `json:"PlayType,omitempty" name:"PlayType"`
}

func (r *ModifyLivePlayDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLivePlayDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLivePlayDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLivePlayDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLivePlayDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLivePushAuthKeyRequest struct {
	*tchttp.BaseRequest

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Whether to enable. 0: disabled; 1: enabled.
	// If this parameter is left empty, the current value will not be modified.
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// Master authentication key.
	// If this parameter is left empty, the current value will not be modified.
	MasterAuthKey *string `json:"MasterAuthKey,omitempty" name:"MasterAuthKey"`

	// Backup authentication key.
	// If this parameter is left empty, the current value will not be modified.
	BackupAuthKey *string `json:"BackupAuthKey,omitempty" name:"BackupAuthKey"`

	// Validity period in seconds.
	AuthDelta *uint64 `json:"AuthDelta,omitempty" name:"AuthDelta"`
}

func (r *ModifyLivePushAuthKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLivePushAuthKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLivePushAuthKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLivePushAuthKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLivePushAuthKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveRecordTemplateRequest struct {
	*tchttp.BaseRequest

	// Template ID obtained through the `DescribeRecordTemplates` API.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// Template name.
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// Message description
	Description *string `json:"Description,omitempty" name:"Description"`

	// FLV recording parameter, which is set when FLV recording is enabled.
	FlvParam *RecordParam `json:"FlvParam,omitempty" name:"FlvParam"`

	// HLS recording parameter, which is set when HLS recording is enabled.
	HlsParam *RecordParam `json:"HlsParam,omitempty" name:"HlsParam"`

	// MP4 recording parameter, which is set when MP4 recording is enabled.
	Mp4Param *RecordParam `json:"Mp4Param,omitempty" name:"Mp4Param"`

	// AAC recording parameter, which is set when AAC recording is enabled.
	AacParam *RecordParam `json:"AacParam,omitempty" name:"AacParam"`

	// Custom HLS recording parameter.
	HlsSpecialParam *HlsSpecialParam `json:"HlsSpecialParam,omitempty" name:"HlsSpecialParam"`

	// MP3 recording parameter, which is set when MP3 recording is enabled.
	Mp3Param *RecordParam `json:"Mp3Param,omitempty" name:"Mp3Param"`
}

func (r *ModifyLiveRecordTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLiveRecordTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveRecordTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLiveRecordTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLiveRecordTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveSnapshotTemplateRequest struct {
	*tchttp.BaseRequest

	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// Template name.
	// Maximum length: 255 bytes.
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// Description.
	// Maximum length: 1,024 bytes.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Screencapturing interval in seconds. Default value: 10s.
	// Value range: 5–600s.
	SnapshotInterval *int64 `json:"SnapshotInterval,omitempty" name:"SnapshotInterval"`

	// Screenshot width. Default value: 0 (original width).
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Screenshot height. Default value: 0 (original height).
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// Whether to enable porn detection. Default value: 0.
	// 0: do not enable.
	// 1: enable.
	PornFlag *int64 `json:"PornFlag,omitempty" name:"PornFlag"`

	// COS application ID.
	CosAppId *int64 `json:"CosAppId,omitempty" name:"CosAppId"`

	// COS bucket name.
	CosBucket *string `json:"CosBucket,omitempty" name:"CosBucket"`

	// COS region.
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// COS bucket folder prefix.
	CosPrefix *string `json:"CosPrefix,omitempty" name:"CosPrefix"`

	// COS filename.
	CosFileName *string `json:"CosFileName,omitempty" name:"CosFileName"`
}

func (r *ModifyLiveSnapshotTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLiveSnapshotTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveSnapshotTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLiveSnapshotTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLiveSnapshotTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveTranscodeTemplateRequest struct {
	*tchttp.BaseRequest

	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// Video encoding format:
	// h264/h265.
	Vcodec *string `json:"Vcodec,omitempty" name:"Vcodec"`

	// Audio encoding format:
	// aac/mp3.
	Acodec *string `json:"Acodec,omitempty" name:"Acodec"`

	// Audio bitrate. Default value: 0.
	// Value range: 0–500.
	AudioBitrate *int64 `json:"AudioBitrate,omitempty" name:"AudioBitrate"`

	// Template description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Video bitrate. Value range: 100–8000 Kbps.
	// Note: the bitrate value must be a multiple of 100.
	VideoBitrate *int64 `json:"VideoBitrate,omitempty" name:"VideoBitrate"`

	// Width. Value range: 0-3000.
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Whether to keep the video. 0: no; 1: yes. Default value: 1.
	NeedVideo *int64 `json:"NeedVideo,omitempty" name:"NeedVideo"`

	// Whether to keep the audio. 0: no; 1: yes. Default value: 1.
	NeedAudio *int64 `json:"NeedAudio,omitempty" name:"NeedAudio"`

	// Height. Value range: 0-3000.
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// Frame rate. Value range: 0–200.
	Fps *int64 `json:"Fps,omitempty" name:"Fps"`

	// Keyframe interval in seconds. Value range: 0–50.
	Gop *int64 `json:"Gop,omitempty" name:"Gop"`

	// Rotation angle.
	// 0, 90, 180, 270.
	Rotate *int64 `json:"Rotate,omitempty" name:"Rotate"`

	// Encoding quality:
	// baseline/main/high.
	Profile *string `json:"Profile,omitempty" name:"Profile"`

	// Whether to not exceed the original bitrate. 0: no; 1: yes. Default value: 0.
	BitrateToOrig *int64 `json:"BitrateToOrig,omitempty" name:"BitrateToOrig"`

	// Whether to not exceed the original height. 0: no; 1: yes. Default value: 0.
	HeightToOrig *int64 `json:"HeightToOrig,omitempty" name:"HeightToOrig"`

	// Whether to not exceed the original frame rate. 0: no; 1: yes. Default value: 0.
	FpsToOrig *int64 `json:"FpsToOrig,omitempty" name:"FpsToOrig"`

	// `VideoBitrate` minus top speed codec bitrate. Value range: 0.1–0.5.
	AdaptBitratePercent *float64 `json:"AdaptBitratePercent,omitempty" name:"AdaptBitratePercent"`
}

func (r *ModifyLiveTranscodeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLiveTranscodeTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveTranscodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLiveTranscodeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLiveTranscodeTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MonitorStreamPlayInfo struct {

	// Playback domain name.
	PlayDomain *string `json:"PlayDomain,omitempty" name:"PlayDomain"`

	// Stream ID.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Playback bitrate. 0 indicates the original bitrate.
	Rate *uint64 `json:"Rate,omitempty" name:"Rate"`

	// Playback protocol. Valid values: Unknown, Flv, Hls, Rtmp, Huyap2p.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Bandwidth in Mbps.
	Bandwidth *float64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// Number of online viewers. A data point is sampled per minute, and the number of TCP connections across the sample points is calculated.
	Online *uint64 `json:"Online,omitempty" name:"Online"`

	// Number of requests.
	Request *uint64 `json:"Request,omitempty" name:"Request"`
}

type PlayAuthKeyInfo struct {

	// Domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Whether to enable:
	// 0: disable.
	// 1: enable.
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// Authentication key.
	AuthKey *string `json:"AuthKey,omitempty" name:"AuthKey"`

	// Validity period in seconds.
	AuthDelta *uint64 `json:"AuthDelta,omitempty" name:"AuthDelta"`

	// Authentication `BackKey`.
	AuthBackKey *string `json:"AuthBackKey,omitempty" name:"AuthBackKey"`
}

type PlayCodeTotalInfo struct {

	// HTTP code. Valid values:
	// 400, 403, 404, 500, 502, 503, 504.
	Code *string `json:"Code,omitempty" name:"Code"`

	// Total occurrences.
	Num *uint64 `json:"Num,omitempty" name:"Num"`
}

type PlayDataInfoByStream struct {

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Total traffic in MB.
	TotalFlux *float64 `json:"TotalFlux,omitempty" name:"TotalFlux"`
}

type PlayStatInfo struct {

	// Data point in time.
	Time *string `json:"Time,omitempty" name:"Time"`

	// Value of bandwidth/traffic/number of requests/number of concurrent connections/download speed. If there is no data returned, the value is 0.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type PlaySumStatInfo struct {

	// Domain name or stream ID.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Average download speed,
	// In MB/s.
	// Calculation formula: average download speed per minute.
	AvgFluxPerSecond *float64 `json:"AvgFluxPerSecond,omitempty" name:"AvgFluxPerSecond"`

	// Total traffic in MB.
	TotalFlux *float64 `json:"TotalFlux,omitempty" name:"TotalFlux"`

	// Total number of requests.
	TotalRequest *uint64 `json:"TotalRequest,omitempty" name:"TotalRequest"`
}

type ProIspPlayCodeDataInfo struct {

	// Country or region.
	CountryAreaName *string `json:"CountryAreaName,omitempty" name:"CountryAreaName"`

	// District.
	ProvinceName *string `json:"ProvinceName,omitempty" name:"ProvinceName"`

	// ISP.
	IspName *string `json:"IspName,omitempty" name:"IspName"`

	// Occurrences of 2xx error codes.
	Code2xx *uint64 `json:"Code2xx,omitempty" name:"Code2xx"`

	// Occurrences of 3xx error codes.
	Code3xx *uint64 `json:"Code3xx,omitempty" name:"Code3xx"`

	// Occurrences of 4xx error codes.
	Code4xx *uint64 `json:"Code4xx,omitempty" name:"Code4xx"`

	// Occurrences of 5xx error codes.
	Code5xx *uint64 `json:"Code5xx,omitempty" name:"Code5xx"`
}

type ProIspPlaySumInfo struct {

	// District/ISP/country/region.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Total traffic in MB.
	TotalFlux *float64 `json:"TotalFlux,omitempty" name:"TotalFlux"`

	// Total number of requests.
	TotalRequest *uint64 `json:"TotalRequest,omitempty" name:"TotalRequest"`

	// Average download traffic in MB/s.
	AvgFluxPerSecond *float64 `json:"AvgFluxPerSecond,omitempty" name:"AvgFluxPerSecond"`
}

type PublishTime struct {

	// Push time.
	// In UTC format, such as 2018-06-29T19:00:00Z.
	PublishTime *string `json:"PublishTime,omitempty" name:"PublishTime"`
}

type PushAuthKeyInfo struct {

	// Domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Whether to enable. 0: disabled; 1: enabled.
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// Master authentication key.
	MasterAuthKey *string `json:"MasterAuthKey,omitempty" name:"MasterAuthKey"`

	// Standby authentication key.
	BackupAuthKey *string `json:"BackupAuthKey,omitempty" name:"BackupAuthKey"`

	// Validity period in seconds.
	AuthDelta *uint64 `json:"AuthDelta,omitempty" name:"AuthDelta"`
}

type PushDataInfo struct {

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Push path.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Push client IP.
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// IP of the server that receives the stream.
	ServerIp *string `json:"ServerIp,omitempty" name:"ServerIp"`

	// Pushed video frame rate in Hz.
	VideoFps *uint64 `json:"VideoFps,omitempty" name:"VideoFps"`

	// Pushed video bitrate in bps.
	VideoSpeed *uint64 `json:"VideoSpeed,omitempty" name:"VideoSpeed"`

	// Pushed audio frame rate in Hz.
	AudioFps *uint64 `json:"AudioFps,omitempty" name:"AudioFps"`

	// Pushed audio bitrate in bps.
	AudioSpeed *uint64 `json:"AudioSpeed,omitempty" name:"AudioSpeed"`

	// Push domain name.
	PushDomain *string `json:"PushDomain,omitempty" name:"PushDomain"`

	// Push start time.
	BeginPushTime *string `json:"BeginPushTime,omitempty" name:"BeginPushTime"`

	// Audio codec,
	// Example: AAC.
	Acodec *string `json:"Acodec,omitempty" name:"Acodec"`

	// Video codec,
	// Example: H.264.
	Vcodec *string `json:"Vcodec,omitempty" name:"Vcodec"`

	// Resolution.
	Resolution *string `json:"Resolution,omitempty" name:"Resolution"`

	// Sample rate.
	AsampleRate *uint64 `json:"AsampleRate,omitempty" name:"AsampleRate"`

	// Audio bitrate in `metadata` in Kbps.
	MetaAudioSpeed *uint64 `json:"MetaAudioSpeed,omitempty" name:"MetaAudioSpeed"`

	// Video bitrate in `metadata` in Kbps.
	MetaVideoSpeed *uint64 `json:"MetaVideoSpeed,omitempty" name:"MetaVideoSpeed"`

	// Frame rate in `metadata`.
	MetaFps *uint64 `json:"MetaFps,omitempty" name:"MetaFps"`
}

type PushQualityData struct {

	// Data time in the format of `%Y-%m-%d %H:%M:%S.%ms` and accurate down to the millisecond level.
	Time *string `json:"Time,omitempty" name:"Time"`

	// Push domain name.
	PushDomain *string `json:"PushDomain,omitempty" name:"PushDomain"`

	// Push path.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Push client IP.
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// Push start time in the format of `%Y-%m-%d %H:%M:%S.%ms` and accurate down to the millisecond level.
	BeginPushTime *string `json:"BeginPushTime,omitempty" name:"BeginPushTime"`

	// Resolution information.
	Resolution *string `json:"Resolution,omitempty" name:"Resolution"`

	// Video codec.
	VCodec *string `json:"VCodec,omitempty" name:"VCodec"`

	// Audio codec.
	ACodec *string `json:"ACodec,omitempty" name:"ACodec"`

	// Push serial number, which uniquely identifies a push.
	Sequence *string `json:"Sequence,omitempty" name:"Sequence"`

	// Video frame rate.
	VideoFps *uint64 `json:"VideoFps,omitempty" name:"VideoFps"`

	// Video bitrate in bps.
	VideoRate *uint64 `json:"VideoRate,omitempty" name:"VideoRate"`

	// Audio frame rate.
	AudioFps *uint64 `json:"AudioFps,omitempty" name:"AudioFps"`

	// Audio bitrate in bps.
	AudioRate *uint64 `json:"AudioRate,omitempty" name:"AudioRate"`

	// Local elapsed time in milliseconds. The greater the difference between audio/video elapsed time and local elapsed time, the poorer the push quality and the more serious the upstream lag.
	LocalTs *uint64 `json:"LocalTs,omitempty" name:"LocalTs"`

	// Video elapsed time in milliseconds.
	VideoTs *uint64 `json:"VideoTs,omitempty" name:"VideoTs"`

	// Audio elapsed time in milliseconds.
	AudioTs *uint64 `json:"AudioTs,omitempty" name:"AudioTs"`

	// Video bitrate in `metadata` in Kbps.
	MetaVideoRate *uint64 `json:"MetaVideoRate,omitempty" name:"MetaVideoRate"`

	// Audio bitrate in `metadata` in Kbps.
	MetaAudioRate *uint64 `json:"MetaAudioRate,omitempty" name:"MetaAudioRate"`

	// Frame rate in `metadata`.
	MateFps *uint64 `json:"MateFps,omitempty" name:"MateFps"`
}

type RecordParam struct {

	// Recording interval.
	// In seconds. Default value: 1800.
	// Value range: 300–7200.
	// This parameter is not valid for HLS, and a file will be generated from push start to interruption during HLS recording.
	RecordInterval *int64 `json:"RecordInterval,omitempty" name:"RecordInterval"`

	// Recording storage period.
	// In seconds. Value range: 0–93312000.
	// 0: permanent storage.
	StorageTime *int64 `json:"StorageTime,omitempty" name:"StorageTime"`

	// Whether to enable recording in the current format. Default value: 0. 0: no, 1: yes.
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// VOD subapplication ID.
	VodSubAppId *int64 `json:"VodSubAppId,omitempty" name:"VodSubAppId"`

	// Recording filename.
	// Supported special placeholders include:
	// {StreamID}: stream ID
	// {StartYear}: start time - year
	// {StartMonth}: start time - month
	// {StartDay}: start time - day
	// {StartHour}: start time - hour
	// {StartMinute}: start time - minute
	// {StartSecond}: start time - second
	// {StartMillisecond}: start time - millisecond
	// {EndYear}: end time - year
	// {EndMonth}: end time - month
	// {EndDay}: end time - day
	// {EndHour}: end time - hour
	// {EndMinute}: end time - minute
	// {EndSecond}: end time - second
	// {EndMillisecond}: end time - millisecond
	// 
	// If this parameter is not set, the recording filename will be `{StreamID}_{StartYear}-{StartMonth}-{StartDay}-{StartHour}-{StartMinute}-{StartSecond}_{EndYear}-{EndMonth}-{EndDay}-{EndHour}-{EndMinute}-{EndSecond}` by default
	VodFileName *string `json:"VodFileName,omitempty" name:"VodFileName"`
}

type RecordTemplateInfo struct {

	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// Template name.
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// Message description
	Description *string `json:"Description,omitempty" name:"Description"`

	// FLV recording parameter.
	FlvParam *RecordParam `json:"FlvParam,omitempty" name:"FlvParam"`

	// HLS recording parameter.
	HlsParam *RecordParam `json:"HlsParam,omitempty" name:"HlsParam"`

	// MP4 recording parameter.
	Mp4Param *RecordParam `json:"Mp4Param,omitempty" name:"Mp4Param"`

	// AAC recording parameter.
	AacParam *RecordParam `json:"AacParam,omitempty" name:"AacParam"`

	// 0: LVB,
	// 1: LCB.
	IsDelayLive *int64 `json:"IsDelayLive,omitempty" name:"IsDelayLive"`

	// Custom HLS recording parameter
	HlsSpecialParam *HlsSpecialParam `json:"HlsSpecialParam,omitempty" name:"HlsSpecialParam"`

	// MP3 recording parameter.
	Mp3Param *RecordParam `json:"Mp3Param,omitempty" name:"Mp3Param"`
}

type ResumeDelayLiveStreamRequest struct {
	*tchttp.BaseRequest

	// Push path, which is the same as the AppName in push and playback addresses and is "live" by default.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *ResumeDelayLiveStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResumeDelayLiveStreamRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResumeDelayLiveStreamResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResumeDelayLiveStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResumeDelayLiveStreamResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResumeLiveStreamRequest struct {
	*tchttp.BaseRequest

	// Push path, which is the same as the AppName in push and playback addresses and is "live" by default.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Your push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *ResumeLiveStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResumeLiveStreamRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResumeLiveStreamResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResumeLiveStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResumeLiveStreamResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RuleInfo struct {

	// Rule creation time.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Rule update time.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Push path.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

type SnapshotTemplateInfo struct {

	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// Template name.
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// Screencapturing interval. Value range: 5–300s.
	SnapshotInterval *int64 `json:"SnapshotInterval,omitempty" name:"SnapshotInterval"`

	// Screenshot width. Value range: 0–3000. 
	// 0: original width and fit to the original ratio.
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Screenshot height. Value range: 0–2000.
	// 0: original height and fit to the original ratio.
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// Whether to enable porn detection. 0: no, 1: yes.
	PornFlag *int64 `json:"PornFlag,omitempty" name:"PornFlag"`

	// COS application ID.
	CosAppId *int64 `json:"CosAppId,omitempty" name:"CosAppId"`

	// COS bucket name.
	CosBucket *string `json:"CosBucket,omitempty" name:"CosBucket"`

	// COS region.
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// Template description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// COS bucket folder prefix.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CosPrefix *string `json:"CosPrefix,omitempty" name:"CosPrefix"`

	// COS filename.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CosFileName *string `json:"CosFileName,omitempty" name:"CosFileName"`
}

type StopLiveRecordRequest struct {
	*tchttp.BaseRequest

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Task ID returned by the `CreateLiveRecord` API.
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *StopLiveRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopLiveRecordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopLiveRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopLiveRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopLiveRecordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopRecordTaskRequest struct {
	*tchttp.BaseRequest

	// Recording task ID.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *StopRecordTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopRecordTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopRecordTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopRecordTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopRecordTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StreamEventInfo struct {

	// Application name.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Push start time.
	// In UTC format, such as 2019-01-07T12:00:00Z.
	StreamStartTime *string `json:"StreamStartTime,omitempty" name:"StreamStartTime"`

	// Push end time.
	// In UTC format, such as 2019-01-07T15:00:00Z.
	StreamEndTime *string `json:"StreamEndTime,omitempty" name:"StreamEndTime"`

	// Stop reason.
	StopReason *string `json:"StopReason,omitempty" name:"StopReason"`

	// Push duration in seconds.
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`

	// Host IP.
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// Resolution.
	Resolution *string `json:"Resolution,omitempty" name:"Resolution"`
}

type StreamName struct {

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Application name.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Push start time.
	// In UTC format, such as 2019-01-07T12:00:00Z.
	StreamStartTime *string `json:"StreamStartTime,omitempty" name:"StreamStartTime"`

	// Push end time.
	// In UTC format, such as 2019-01-07T15:00:00Z.
	StreamEndTime *string `json:"StreamEndTime,omitempty" name:"StreamEndTime"`

	// Stop reason.
	StopReason *string `json:"StopReason,omitempty" name:"StopReason"`

	// Push duration in seconds.
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`

	// Host IP.
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// Resolution.
	Resolution *string `json:"Resolution,omitempty" name:"Resolution"`
}

type StreamOnlineInfo struct {

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Push time list
	PublishTimeList []*PublishTime `json:"PublishTimeList,omitempty" name:"PublishTimeList" list`

	// Application name.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Push domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
}

type TemplateInfo struct {

	// Video encoding format:
	// h264/h265.
	Vcodec *string `json:"Vcodec,omitempty" name:"Vcodec"`

	// Video bitrate. Value range: 100–8000 Kbps.
	VideoBitrate *int64 `json:"VideoBitrate,omitempty" name:"VideoBitrate"`

	// Audio codec. Valid values: aac, mp3.
	Acodec *string `json:"Acodec,omitempty" name:"Acodec"`

	// Audio bitrate. Value range: 0–500 Kbps.
	AudioBitrate *int64 `json:"AudioBitrate,omitempty" name:"AudioBitrate"`

	// Width. Value range: 0–3000.
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Height. Value range: 0–3000.
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// Frame rate. Value range: 0–200 FPS.
	Fps *int64 `json:"Fps,omitempty" name:"Fps"`

	// Keyframe interval. Value range: 1–50s.
	Gop *int64 `json:"Gop,omitempty" name:"Gop"`

	// Rotation angle. Valid values: 0, 90, 180, 270.
	Rotate *int64 `json:"Rotate,omitempty" name:"Rotate"`

	// Encoding quality. Valid values:
	// baseline, main, high.
	Profile *string `json:"Profile,omitempty" name:"Profile"`

	// Whether to not exceed the original bitrate. 0: no; 1: yes.
	BitrateToOrig *int64 `json:"BitrateToOrig,omitempty" name:"BitrateToOrig"`

	// Whether to not exceed the original height. 0: no; 1: yes.
	HeightToOrig *int64 `json:"HeightToOrig,omitempty" name:"HeightToOrig"`

	// Whether to not exceed the original frame rate. 0: no; 1: yes.
	FpsToOrig *int64 `json:"FpsToOrig,omitempty" name:"FpsToOrig"`

	// Whether to keep the video. 0: no; 1: yes.
	NeedVideo *int64 `json:"NeedVideo,omitempty" name:"NeedVideo"`

	// Whether to keep the audio. 0: no; 1: yes.
	NeedAudio *int64 `json:"NeedAudio,omitempty" name:"NeedAudio"`

	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// Template name.
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// Template description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Whether it is a top speed codec template. 0: no, 1: yes. Default value: 0.
	AiTransCode *int64 `json:"AiTransCode,omitempty" name:"AiTransCode"`

	// `VideoBitrate` minus top speed codec bitrate. Value range: 0.1–0.5.
	AdaptBitratePercent *float64 `json:"AdaptBitratePercent,omitempty" name:"AdaptBitratePercent"`
}

type TimeValue struct {

	// UTC time in the format of `yyyy-mm-ddTHH:MM:SSZ`.
	Time *string `json:"Time,omitempty" name:"Time"`

	// Value.
	Num *uint64 `json:"Num,omitempty" name:"Num"`
}

type TranscodeDetailInfo struct {

	// Stream name.
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// Start time (Beijing time) in the format of `yyyy-mm-dd HH:MM`.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time (Beijing time) in the format of `yyyy-mm-dd HH:MM`.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Transcoding duration in minutes.
	// Note: given the possible interruptions during push, duration here is the sum of actual duration of transcoding instead of the interval between the start time and end time.
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`

	// Codec with modules,
	// Example:
	// liveprocessor_H264: LVB transcoding - H264,
	// liveprocessor_H265: LVB transcoding - H265,
	// topspeed_H264: top speed codec - H264,
	// topspeed_H265: top speed codec - H265.
	ModuleCodec *string `json:"ModuleCodec,omitempty" name:"ModuleCodec"`

	// Bitrate.
	Bitrate *uint64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// Type. Valid values: Transcode, MixStream, WaterMark.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Push domain name.
	PushDomain *string `json:"PushDomain,omitempty" name:"PushDomain"`

	// Resolution.
	Resolution *string `json:"Resolution,omitempty" name:"Resolution"`
}

type UnBindLiveDomainCertRequest struct {
	*tchttp.BaseRequest

	// Playback domain name.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
}

func (r *UnBindLiveDomainCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindLiveDomainCertRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnBindLiveDomainCertResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnBindLiveDomainCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindLiveDomainCertResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateLiveWatermarkRequest struct {
	*tchttp.BaseRequest

	// Watermark ID.
	// Get the watermark ID in the returned value of the [AddLiveWatermark](/document/product/267/30154) API call.
	WatermarkId *int64 `json:"WatermarkId,omitempty" name:"WatermarkId"`

	// Watermark image URL.
	PictureUrl *string `json:"PictureUrl,omitempty" name:"PictureUrl"`

	// Display position: X-axis offset in %. Default value: 0.
	XPosition *int64 `json:"XPosition,omitempty" name:"XPosition"`

	// Display position: Y-axis offset in %. Default value: 0.
	YPosition *int64 `json:"YPosition,omitempty" name:"YPosition"`

	// Watermark name.
	// Up to 16 bytes.
	WatermarkName *string `json:"WatermarkName,omitempty" name:"WatermarkName"`

	// Watermark width or its percentage of the live streaming video width. It is recommended to just specify either height or width as the other will be scaled proportionally to avoid distortions. The original width is used by default.
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Watermark height or its percentage of the live streaming video width. It is recommended to just specify either height or width as the other will be scaled proportionally to avoid distortions. The original height is used by default.
	Height *int64 `json:"Height,omitempty" name:"Height"`
}

func (r *UpdateLiveWatermarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateLiveWatermarkRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateLiveWatermarkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateLiveWatermarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateLiveWatermarkResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type WatermarkInfo struct {

	// Watermark ID.
	WatermarkId *int64 `json:"WatermarkId,omitempty" name:"WatermarkId"`

	// Watermark image URL.
	PictureUrl *string `json:"PictureUrl,omitempty" name:"PictureUrl"`

	// Display position: X-axis offset.
	XPosition *int64 `json:"XPosition,omitempty" name:"XPosition"`

	// Display position: Y-axis offset.
	YPosition *int64 `json:"YPosition,omitempty" name:"YPosition"`

	// Watermark name.
	WatermarkName *string `json:"WatermarkName,omitempty" name:"WatermarkName"`

	// Current status. 0: not used. 1: in use.
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Creation time.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Watermark width.
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// Watermark height.
	Height *int64 `json:"Height,omitempty" name:"Height"`
}
