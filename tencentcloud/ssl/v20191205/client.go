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

package v20191205

import (
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2019-12-05"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewApplyCertificateRequest() (request *ApplyCertificateRequest) {
    request = &ApplyCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssl", APIVersion, "ApplyCertificate")
    return
}

func NewApplyCertificateResponse() (response *ApplyCertificateResponse) {
    response = &ApplyCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to apply for a free certificate.
func (c *Client) ApplyCertificate(request *ApplyCertificateRequest) (response *ApplyCertificateResponse, err error) {
    if request == nil {
        request = NewApplyCertificateRequest()
    }
    response = NewApplyCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewCancelCertificateOrderRequest() (request *CancelCertificateOrderRequest) {
    request = &CancelCertificateOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssl", APIVersion, "CancelCertificateOrder")
    return
}

func NewCancelCertificateOrderResponse() (response *CancelCertificateOrderResponse) {
    response = &CancelCertificateOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to cancel a certificate order.
func (c *Client) CancelCertificateOrder(request *CancelCertificateOrderRequest) (response *CancelCertificateOrderResponse, err error) {
    if request == nil {
        request = NewCancelCertificateOrderRequest()
    }
    response = NewCancelCertificateOrderResponse()
    err = c.Send(request, response)
    return
}

func NewCommitCertificateInformationRequest() (request *CommitCertificateInformationRequest) {
    request = &CommitCertificateInformationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssl", APIVersion, "CommitCertificateInformation")
    return
}

func NewCommitCertificateInformationResponse() (response *CommitCertificateInformationResponse) {
    response = &CommitCertificateInformationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to submit a certificate order.
func (c *Client) CommitCertificateInformation(request *CommitCertificateInformationRequest) (response *CommitCertificateInformationResponse, err error) {
    if request == nil {
        request = NewCommitCertificateInformationRequest()
    }
    response = NewCommitCertificateInformationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCertificateRequest() (request *DeleteCertificateRequest) {
    request = &DeleteCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssl", APIVersion, "DeleteCertificate")
    return
}

func NewDeleteCertificateResponse() (response *DeleteCertificateResponse) {
    response = &DeleteCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete a certificate.
func (c *Client) DeleteCertificate(request *DeleteCertificateRequest) (response *DeleteCertificateResponse, err error) {
    if request == nil {
        request = NewDeleteCertificateRequest()
    }
    response = NewDeleteCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCertificateRequest() (request *DescribeCertificateRequest) {
    request = &DescribeCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeCertificate")
    return
}

func NewDescribeCertificateResponse() (response *DescribeCertificateResponse) {
    response = &DescribeCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to obtain certificate information.
func (c *Client) DescribeCertificate(request *DescribeCertificateRequest) (response *DescribeCertificateResponse, err error) {
    if request == nil {
        request = NewDescribeCertificateRequest()
    }
    response = NewDescribeCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCertificateDetailRequest() (request *DescribeCertificateDetailRequest) {
    request = &DescribeCertificateDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeCertificateDetail")
    return
}

func NewDescribeCertificateDetailResponse() (response *DescribeCertificateDetailResponse) {
    response = &DescribeCertificateDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to obtain certificate details.
func (c *Client) DescribeCertificateDetail(request *DescribeCertificateDetailRequest) (response *DescribeCertificateDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCertificateDetailRequest()
    }
    response = NewDescribeCertificateDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCertificateOperateLogsRequest() (request *DescribeCertificateOperateLogsRequest) {
    request = &DescribeCertificateOperateLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeCertificateOperateLogs")
    return
}

func NewDescribeCertificateOperateLogsResponse() (response *DescribeCertificateOperateLogsResponse) {
    response = &DescribeCertificateOperateLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to obtain certificate operation logs in the current account.
func (c *Client) DescribeCertificateOperateLogs(request *DescribeCertificateOperateLogsRequest) (response *DescribeCertificateOperateLogsResponse, err error) {
    if request == nil {
        request = NewDescribeCertificateOperateLogsRequest()
    }
    response = NewDescribeCertificateOperateLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCertificatesRequest() (request *DescribeCertificatesRequest) {
    request = &DescribeCertificatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeCertificates")
    return
}

func NewDescribeCertificatesResponse() (response *DescribeCertificatesResponse) {
    response = &DescribeCertificatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to obtain the certificate list.
func (c *Client) DescribeCertificates(request *DescribeCertificatesRequest) (response *DescribeCertificatesResponse, err error) {
    if request == nil {
        request = NewDescribeCertificatesRequest()
    }
    response = NewDescribeCertificatesResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadCertificateRequest() (request *DownloadCertificateRequest) {
    request = &DownloadCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssl", APIVersion, "DownloadCertificate")
    return
}

func NewDownloadCertificateResponse() (response *DownloadCertificateResponse) {
    response = &DownloadCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to download a certificate.
func (c *Client) DownloadCertificate(request *DownloadCertificateRequest) (response *DownloadCertificateResponse, err error) {
    if request == nil {
        request = NewDownloadCertificateRequest()
    }
    response = NewDownloadCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCertificateAliasRequest() (request *ModifyCertificateAliasRequest) {
    request = &ModifyCertificateAliasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssl", APIVersion, "ModifyCertificateAlias")
    return
}

func NewModifyCertificateAliasResponse() (response *ModifyCertificateAliasResponse) {
    response = &ModifyCertificateAliasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify certificate alias by inputting the certificate ID and new alias.
func (c *Client) ModifyCertificateAlias(request *ModifyCertificateAliasRequest) (response *ModifyCertificateAliasResponse, err error) {
    if request == nil {
        request = NewModifyCertificateAliasRequest()
    }
    response = NewModifyCertificateAliasResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCertificateProjectRequest() (request *ModifyCertificateProjectRequest) {
    request = &ModifyCertificateProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssl", APIVersion, "ModifyCertificateProject")
    return
}

func NewModifyCertificateProjectResponse() (response *ModifyCertificateProjectResponse) {
    response = &ModifyCertificateProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the projects of multiple certificates.
func (c *Client) ModifyCertificateProject(request *ModifyCertificateProjectRequest) (response *ModifyCertificateProjectResponse, err error) {
    if request == nil {
        request = NewModifyCertificateProjectRequest()
    }
    response = NewModifyCertificateProjectResponse()
    err = c.Send(request, response)
    return
}

func NewReplaceCertificateRequest() (request *ReplaceCertificateRequest) {
    request = &ReplaceCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssl", APIVersion, "ReplaceCertificate")
    return
}

func NewReplaceCertificateResponse() (response *ReplaceCertificateResponse) {
    response = &ReplaceCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to reissue a certificate. Note that if you have applied for a free certificate, only an RSA-2048 certificate will be reissued, and the certificate can be reissued only once.
func (c *Client) ReplaceCertificate(request *ReplaceCertificateRequest) (response *ReplaceCertificateResponse, err error) {
    if request == nil {
        request = NewReplaceCertificateRequest()
    }
    response = NewReplaceCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitCertificateInformationRequest() (request *SubmitCertificateInformationRequest) {
    request = &SubmitCertificateInformationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssl", APIVersion, "SubmitCertificateInformation")
    return
}

func NewSubmitCertificateInformationResponse() (response *SubmitCertificateInformationResponse) {
    response = &SubmitCertificateInformationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to submit certificate information.
func (c *Client) SubmitCertificateInformation(request *SubmitCertificateInformationRequest) (response *SubmitCertificateInformationResponse, err error) {
    if request == nil {
        request = NewSubmitCertificateInformationRequest()
    }
    response = NewSubmitCertificateInformationResponse()
    err = c.Send(request, response)
    return
}

func NewUploadCertificateRequest() (request *UploadCertificateRequest) {
    request = &UploadCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssl", APIVersion, "UploadCertificate")
    return
}

func NewUploadCertificateResponse() (response *UploadCertificateResponse) {
    response = &UploadCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to upload a certificate.
func (c *Client) UploadCertificate(request *UploadCertificateRequest) (response *UploadCertificateResponse, err error) {
    if request == nil {
        request = NewUploadCertificateRequest()
    }
    response = NewUploadCertificateResponse()
    err = c.Send(request, response)
    return
}
