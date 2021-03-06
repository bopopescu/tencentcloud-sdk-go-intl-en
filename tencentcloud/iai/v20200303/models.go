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

package v20200303

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AnalyzeFaceRequest struct {
	*tchttp.BaseRequest

	// Detection mode. 0: detect all faces that appear; 1: detect the largest face. Default value: 0. The facial feature localization information (facial keypoints) of up to 10 faces can be returned.
	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`

	// Base64-encoded image data, which cannot exceed 5 MB.
	// .png, .jpg, .jpeg, and .bmp images are supported, while .gif images are not.
	Image *string `json:"Image,omitempty" name:"Image"`

	// Image URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used.  
	// You are recommended to store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. 
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// .png, .jpg, .jpeg, and .bmp images are supported, while .gif images are not.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Algorithm model version used by the Face Recognition service. Valid values: 2.0, 3.0.  
	// This parameter is `3.0` by default starting from April 2, 2020. If it is left empty for accounts that used this API previously, `2.0` will be used by default.  
	// Different algorithm model versions correspond to different face recognition algorithms. The new version has a better overall effect than the legacy version and is thus recommended.
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

func (r *AnalyzeFaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AnalyzeFaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AnalyzeFaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Width of requested image.
		ImageWidth *uint64 `json:"ImageWidth,omitempty" name:"ImageWidth"`

		// Height of requested image.
		ImageHeight *uint64 `json:"ImageHeight,omitempty" name:"ImageHeight"`

		// Specific information of facial feature localization (facial keypoints).
		FaceShapeSet []*FaceShape `json:"FaceShapeSet,omitempty" name:"FaceShapeSet" list`

		// Algorithm model version used for face recognition.
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AnalyzeFaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AnalyzeFaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Candidate struct {

	// Person ID
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// Face ID
	FaceId *string `json:"FaceId,omitempty" name:"FaceId"`

	// Match score of candidate. 
	// 
	// In a face base library containing 10,000 faces, the 1%, 0.1%, and 0.01% FARs correspond to scores of 70, 80, and 90 points, respectively;
	// In a face base library containing 100,000 faces, the 1%, 0.1%, and 0.01% FARs correspond to scores of 80, 90, and 100 points, respectively;
	// In a face base library containing 300,000 faces, the 1% and 0.1% FARs correspond to scores of 85 and 95, respectively.
	// 
	// Generally, the score of 80 is suitable for most scenarios. You are recommended to choose an appropriate score based on the actual situation, preferably no more than 90.
	Score *float64 `json:"Score,omitempty" name:"Score"`

	// Person name
	// Note: this field may return null, indicating that no valid values can be obtained.
	PersonName *string `json:"PersonName,omitempty" name:"PersonName"`

	// Person gender
	// Note: this field may return null, indicating that no valid values can be obtained.
	Gender *int64 `json:"Gender,omitempty" name:"Gender"`

	// List of groups containing this person and their description fields
	// Note: this field may return null, indicating that no valid values can be obtained.
	PersonGroupInfos []*PersonGroupInfo `json:"PersonGroupInfos,omitempty" name:"PersonGroupInfos" list`
}

type CheckSimilarPersonRequest struct {
	*tchttp.BaseRequest

	// List of groups to be checked. 
	// There can be up to 2 million persons in one group and up to 10 groups.
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds" list`

	// Control over the strictness of duplicate person check.
	// 1: archive sorting with high strictness, which can eliminate more duplicate identities but leads to higher false elimination rate for non-duplicate identities.
	// 2: archive sorting with low strictness, which leads to lower false elimination rate for non-duplicate identities and lower elimination rate for duplicate identities.
	UniquePersonControl *int64 `json:"UniquePersonControl,omitempty" name:"UniquePersonControl"`
}

func (r *CheckSimilarPersonRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckSimilarPersonRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckSimilarPersonResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Duplicate check task ID, which is used to query and get the progress and result of the task.
		JobId *string `json:"JobId,omitempty" name:"JobId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckSimilarPersonResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckSimilarPersonResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CompareFaceRequest struct {
	*tchttp.BaseRequest

	// Base64-encoded image A data, which cannot exceed 5 MB.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// .png, .jpg, .jpeg, and .bmp images are supported, while .gif images are not.
	ImageA *string `json:"ImageA,omitempty" name:"ImageA"`

	// Base64-encoded image B data, which cannot exceed 5 MB.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// .png, .jpg, .jpeg, and .bmp images are supported, while .gif images are not.
	ImageB *string `json:"ImageB,omitempty" name:"ImageB"`

	// Image A URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// Either `Url` or `Image` of image A must be provided; if both are provided, only `Url` will be used. 
	// You are recommended to store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. 
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// .png, .jpg, .jpeg, and .bmp images are supported, while .gif images are not.
	UrlA *string `json:"UrlA,omitempty" name:"UrlA"`

	// Image B URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// Either `Url` or `Image` of image B must be provided; if both are provided, only `Url` will be used. 
	// You are recommended to store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. 
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// .png, .jpg, .jpeg, and .bmp images are supported, while .gif images are not.
	UrlB *string `json:"UrlB,omitempty" name:"UrlB"`

	// Algorithm model version used by the Face Recognition service. Valid values: 2.0, 3.0. 
	// This parameter is `3.0` by default starting from April 2, 2020. If it is left empty for accounts that used this API previously, `2.0` will be used by default. 
	// Different algorithm model versions correspond to different face recognition algorithms. The 3.0 version has a better overall effect than the legacy version and is thus recommended.
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	// Image quality control. 
	// 0: no control. 
	// 1: low quality requirement. The image has one or more of the following problems: extreme blurriness, covered eyes, covered nose, and covered mouth. 
	// 2: average quality requirement. The image has at least three of following problems: extreme brightness, extreme dimness, blurriness or average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 3: high quality requirement. The image has one to two of following problems: extreme brightness, extreme dimness, average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 4: very high quality requirement. The image is optimal in all dimensions or only has a slight problem in one dimension. 
	// Default value: 0. 
	// If the image quality does not meet the requirement, the returned result will prompt that the detected image quality is unsatisfactory.
	QualityControl *uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

func (r *CompareFaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CompareFaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CompareFaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Face similarity score between two images.
	// The returned similarity score varies by algorithm version. 
	// If you need to verify whether the faces in the two images are the same person, then the 0.1%, 0.01%, and 0.001% FARs on v3.0 correspond to scores of 40, 50, and 60, respectively. Generally, if the score is above 50, it can be judged that they are the same person. 
	// The 0.1%, 0.01%, and 0.001% FARs on v2.0 correspond to scores of 70, 80, and 90, respectively. Generally, if the score is above 80, it can be judged that they are the same person. 
	// If you need to verify whether the faces in the two images are the same person, you are recommended to use the `VerifyFace` API.
		Score *float64 `json:"Score,omitempty" name:"Score"`

		// Algorithm model version used for face recognition.
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CompareFaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CompareFaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CopyPersonRequest struct {
	*tchttp.BaseRequest

	// Person ID
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// List of the groups to add to.
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds" list`
}

func (r *CopyPersonRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CopyPersonRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CopyPersonResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of groups successfully added to.
		SucGroupNum *uint64 `json:"SucGroupNum,omitempty" name:"SucGroupNum"`

		// List of groups successfully added to.
		SucGroupIds []*string `json:"SucGroupIds,omitempty" name:"SucGroupIds" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CopyPersonResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CopyPersonResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateFaceRequest struct {
	*tchttp.BaseRequest

	// Person ID.
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// Base64-encoded image data, which cannot exceed 5 MB.
	// There can be up to 5 faces in one image.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// .png, .jpg, .jpeg, and .bmp images are supported, while .gif images are not.
	Images []*string `json:"Images,omitempty" name:"Images" list`

	// Image URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used.  
	// You are recommended to store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. 
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// .png, .jpg, .jpeg, and .bmp images are supported, while .gif images are not.
	// There can be up to 5 faces in one image.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	Urls []*string `json:"Urls,omitempty" name:"Urls" list`

	// Only faces whose similarity to an existing face of the person is above the value of `FaceMatchThreshold` can be added successfully. 
	// Default value: 60. Value range: [0,100].
	FaceMatchThreshold *float64 `json:"FaceMatchThreshold,omitempty" name:"FaceMatchThreshold"`

	// Image quality control. 
	// 0: no control. 
	// 1: low quality requirement. The image has one or more of the following problems: extreme blurriness, covered eyes, covered nose, and covered mouth. 
	// 2: average quality requirement. The image has at least three of following problems: extreme brightness, extreme dimness, blurriness or average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 3: high quality requirement. The image has one to two of following problems: extreme brightness, extreme dimness, average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 4: very high quality requirement. The image is optimal in all dimensions or only has a slight problem in one dimension. 
	// Default value: 0. 
	// If the image quality does not meet the requirement, the returned result will prompt that the detected image quality is unsatisfactory.
	QualityControl *uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

func (r *CreateFaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateFaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateFaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of successfully added faces
		SucFaceNum *uint64 `json:"SucFaceNum,omitempty" name:"SucFaceNum"`

		// List of IDs of successfully added faces
		SucFaceIds []*string `json:"SucFaceIds,omitempty" name:"SucFaceIds" list`

		// Adding result for each face image. -1101: no face detected; -1102: image decoding failed; 
	// -1601: the image quality control requirement is not met; -1604: the face similarity is not above `FaceMatchThreshold`. 
	// Other non-zero values: algorithm service exception. 
	// The order of `RetCode` values is the same as the order of `Images` or `Urls` in the input parameter.
		RetCode []*int64 `json:"RetCode,omitempty" name:"RetCode" list`

		// Indexes of successfully added faces. The order of indexes is the same as the order of `Images` or `Urls` in the input parameter. 
	// For example, if there are 3 URLs in `Urls`, and the second URL fails, then the value of `SucIndexes` will be [0,2].
		SucIndexes []*uint64 `json:"SucIndexes,omitempty" name:"SucIndexes" list`

		// Frame positions of successfully added faces. The order is the same as the order of `Images` or `Urls` in the input parameter.
		SucFaceRects []*FaceRect `json:"SucFaceRects,omitempty" name:"SucFaceRects" list`

		// Algorithm model version used for face recognition.
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateFaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateFaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateGroupRequest struct {
	*tchttp.BaseRequest

	// Group name, which is modifiable, must be unique, and can contain 1–60 characters.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Group ID, which is unmodifiable, must be unique, and can contain letters, digits, and special symbols (-%@#&_) of up to 64B.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Custom group description field that describes the person attributes in the group, which will be applied to all persons in the group. 
	// Up to 5 ones can be created. 
	// Each custom description field can contain 1–30 characters. 
	// The custom description field must be unique in the group. 
	// Example: if you set the "custom description field" of a group to ["student ID","employee ID","mobile number"], 
	// then all the persons in the group will have description fields named "student ID", "employee ID", and "mobile number". 
	// You can enter content in the corresponding field to register a person's student ID, employee ID, and mobile number.
	GroupExDescriptions []*string `json:"GroupExDescriptions,omitempty" name:"GroupExDescriptions" list`

	// Group remarks, which can contain 0–40 characters.
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// Algorithm model version used by the Face Recognition service. Valid values: 2.0, 3.0.
	// This parameter is `3.0` by default starting from April 2, 2020. If it is left empty for accounts that used this API previously, `2.0` will be used by default. 
	// Different algorithm model versions correspond to different face recognition algorithms. The 3.0 version has a better overall effect than the legacy version and is thus recommended.
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`
}

func (r *CreateGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Algorithm model version used for face recognition.
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePersonRequest struct {
	*tchttp.BaseRequest

	// ID of the group to add to.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Person name, which can contain 1–60 characters and is modifiable and repeatable.
	PersonName *string `json:"PersonName,omitempty" name:"PersonName"`

	// Person ID, which is unmodifiable, must be unique under a Tencent Cloud account, and can contain letters, digits, and special symbols (-%@#&_) of up to 64B.
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 0: empty; 1: male; 2: female.
	Gender *int64 `json:"Gender,omitempty" name:"Gender"`

	// Content of person description field, which is a `key-value` pair, can contain 0–60 characters, and is modifiable and repeatable.
	PersonExDescriptionInfos []*PersonExDescriptionInfo `json:"PersonExDescriptionInfos,omitempty" name:"PersonExDescriptionInfos" list`

	// Base64-encoded image data, which cannot exceed 5 MB.
	// .png, .jpg, .jpeg, and .bmp images are supported, while .gif images are not.
	Image *string `json:"Image,omitempty" name:"Image"`

	// Image URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used.  
	// You are recommended to store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. 
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// .png, .jpg, .jpeg, and .bmp images are supported, while .gif images are not.
	Url *string `json:"Url,omitempty" name:"Url"`

	// This parameter is used to control the judgment whether the face contained in the image in `Image` or `Url` corresponds to an existing person in the group. 
	// If it is judged that a duplicate person exists in the group, no new person will be created, and information of the suspected duplicate person will be returned. 
	// Otherwise, the new person will be created. 
	// 0: do not judge, i.e., the person will be created no matter whether a duplicate person exists in the group. 
	// 1: low duplicate person judgment requirement (1% FAR); 
	// 2: average duplicate person judgment requirement (0.1% FAR); 
	// 3: high duplicate person judgment requirement (0.01% FAR); 
	// 4: very high duplicate person judgment requirement (0.001% FAR). 
	// Default value: 0.  
	// Note: the higher the requirement, the lower the probability of duplicate person. The FARs corresponding to different requirements are for reference only and can be adjusted as needed.
	UniquePersonControl *uint64 `json:"UniquePersonControl,omitempty" name:"UniquePersonControl"`

	// Image quality control. 
	// 0: no control. 
	// 1: low quality requirement. The image has one or more of the following problems: extreme blurriness, covered eyes, covered nose, and covered mouth. 
	// 2: average quality requirement. The image has at least three of following problems: extreme brightness, extreme dimness, blurriness or average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 3: high quality requirement. The image has one to two of following problems: extreme brightness, extreme dimness, average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 4: very high quality requirement. The image is optimal in all dimensions or only has a slight problem in one dimension. 
	// Default value: 0. 
	// If the image quality does not meet the requirement, the returned result will prompt that the detected image quality is unsatisfactory.
	QualityControl *uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

func (r *CreatePersonRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePersonRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePersonResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Unique ID of face image.
		FaceId *string `json:"FaceId,omitempty" name:"FaceId"`

		// Position of detected face frame.
	// Note: this field may return null, indicating that no valid values can be obtained.
		FaceRect *FaceRect `json:"FaceRect,omitempty" name:"FaceRect"`

		// `PersonId` of suspected duplicate person. 
	// This parameter is meaningful only if the `UniquePersonControl` parameter is not 0 and there is a suspected duplicate person in the group.
		SimilarPersonId *string `json:"SimilarPersonId,omitempty" name:"SimilarPersonId"`

		// Algorithm model version used for face recognition.
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePersonResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePersonResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteFaceRequest struct {
	*tchttp.BaseRequest

	// Person ID
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// List of IDs of the faces to be deleted
	FaceIds []*string `json:"FaceIds,omitempty" name:"FaceIds" list`
}

func (r *DeleteFaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteFaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteFaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of successfully deleted faces
		SucDeletedNum *uint64 `json:"SucDeletedNum,omitempty" name:"SucDeletedNum"`

		// List of IDs of successfully deleted faces
		SucFaceIds []*string `json:"SucFaceIds,omitempty" name:"SucFaceIds" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteFaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteFaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteGroupRequest struct {
	*tchttp.BaseRequest

	// Group ID.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeleteGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePersonFromGroupRequest struct {
	*tchttp.BaseRequest

	// Person ID
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// Group ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeletePersonFromGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePersonFromGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePersonFromGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePersonFromGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePersonFromGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePersonRequest struct {
	*tchttp.BaseRequest

	// Person ID
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`
}

func (r *DeletePersonRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePersonRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePersonResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePersonResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePersonResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetectFaceRequest struct {
	*tchttp.BaseRequest

	// Maximum number of processable faces. Default value: 1 (i.e., detecting only the face with the largest size in the image). Maximum value: 120. 
	// This parameter is used to control the number of faces in the image to be detected. The smaller the value, the faster the processing.
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitempty" name:"MaxFaceNum"`

	// Minimum height and width of face in px.
	// Default value: 34. You are recommended to keep it at or above 34.
	// Faces below the `MinFaceSize` value will not be detected.
	MinFaceSize *uint64 `json:"MinFaceSize,omitempty" name:"MinFaceSize"`

	// Base64-encoded image data, which cannot exceed 5 MB.
	// .png, .jpg, .jpeg, and .bmp images are supported, while .gif images are not.
	Image *string `json:"Image,omitempty" name:"Image"`

	// Image URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used.  
	// You are recommended to store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. 
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// .png, .jpg, .jpeg, and .bmp images are supported, while .gif images are not.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Whether the face attribute information (FaceAttributesInfo) needs to be returned. 0: no; 1: yes. Default value: 0. 
	// If the value is not 1, it will be deemed as no need to return, and `FaceAttributesInfo` is meaningless in this case.  
	// The face attribute information of up to 5 largest faces in the image will be returned, and `FaceAttributesInfo` of the 6th and rest faces is meaningless.  
	// Extracting face attribute information is quite time-consuming. If face attribute information is not required, you are recommended to disable this feature to speed up face detection.
	NeedFaceAttributes *uint64 `json:"NeedFaceAttributes,omitempty" name:"NeedFaceAttributes"`

	// Whether to enable quality detection. 0: no; 1: yes. Default value: 0. 
	// If the value is not 1, it will be deemed not to perform quality detection.
	// The face quality score information of up to 30 largest faces in the image will be returned, and `FaceQualityInfo` of the 31st and rest faces is meaningless.  
	// You are recommended to enable this feature for the face adding operation.
	NeedQualityDetection *uint64 `json:"NeedQualityDetection,omitempty" name:"NeedQualityDetection"`

	// Algorithm model version used by the Face Recognition service. Valid values: 2.0, 3.0.  
	// This parameter is `3.0` by default starting from April 2, 2020. If it is left empty for accounts that used this API previously, `2.0` will be used by default. 
	// Different algorithm model versions correspond to different face recognition algorithms. The 3.0 version has a better overall effect than the legacy version and is thus recommended.
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

func (r *DetectFaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetectFaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetectFaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Width of requested image.
		ImageWidth *int64 `json:"ImageWidth,omitempty" name:"ImageWidth"`

		// Height of requested image.
		ImageHeight *int64 `json:"ImageHeight,omitempty" name:"ImageHeight"`

		// Face information list, including face coordinate information, attribute information (if needed), and quality score information (if needed).
		FaceInfos []*FaceInfo `json:"FaceInfos,omitempty" name:"FaceInfos" list`

		// Algorithm model version used for face recognition.
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetectFaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetectFaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetectLiveFaceRequest struct {
	*tchttp.BaseRequest

	// Base64-encoded image data, which cannot exceed 5 MB. (The aspect ratio of the image should be close to 3:4 (width:height); otherwise, the score returned for the image will be meaningless.)
	// .png, .jpg, .jpeg, and .bmp images are supported, while .gif images are not.
	Image *string `json:"Image,omitempty" name:"Image"`

	// Image URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used. 
	// (The aspect ratio of the image should be close to 3:4 (width:height); otherwise, the score returned for the image will be meaningless.) 
	// You are recommended to store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. 
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// .png, .jpg, .jpeg, and .bmp images are supported, while .gif images are not.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Algorithm model version used by the Face Recognition service. Valid values: 2.0, 3.0.  
	// This parameter is `3.0` by default starting from April 2, 2020. If it is left empty for accounts that used this API previously, `2.0` will be used by default. 
	// Different algorithm model versions correspond to different face recognition algorithms. The 3.0 version has a better overall effect than the legacy version and is thus recommended.
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`
}

func (r *DetectLiveFaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetectLiveFaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetectLiveFaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Liveness score. Value range: [0,100]. The score is generally between 80 and 100, but 0 is also a common value. As a recommendation, when the score is greater than 87, it can be judged that the person in the image is alive. You can adjust the threshold according to your specific scenario.
	// This field is meaningful only if `FaceModelVersion` is 2.0.
		Score *float64 `json:"Score,omitempty" name:"Score"`

		// Algorithm model version used for face recognition.
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// Whether liveness detection is passed.
	// This field is meaningful only if `FaceModelVersion` is 3.0.
		IsLiveness *bool `json:"IsLiveness,omitempty" name:"IsLiveness"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetectLiveFaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetectLiveFaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EstimateCheckSimilarPersonCostTimeRequest struct {
	*tchttp.BaseRequest

	// List of groups to be checked. 
	// There can be up to 2 million persons in one group and up to 10 groups.
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds" list`
}

func (r *EstimateCheckSimilarPersonCostTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EstimateCheckSimilarPersonCostTimeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EstimateCheckSimilarPersonCostTimeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Estimated duration of duplicate person check task in minutes.
		EstimatedTimeCost *uint64 `json:"EstimatedTimeCost,omitempty" name:"EstimatedTimeCost"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EstimateCheckSimilarPersonCostTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EstimateCheckSimilarPersonCostTimeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type FaceAttributesInfo struct {

	// Gender. The gender is female for the value range [0,49] and male for the value range [50,100]. The closer the value to 0 or 100, the higher the confidence. If `NeedFaceAttributes` is not 1 or more than 5 faces are detected, this parameter will still be returned but meaningless.
	Gender *int64 `json:"Gender,omitempty" name:"Gender"`

	// Age. Value range: [0,100]. If `NeedFaceAttributes` is not 1 or more than 5 faces are detected, this parameter will still be returned but meaningless.
	Age *int64 `json:"Age,omitempty" name:"Age"`

	// Expression. Value range: [0 (normal)–50 (smile)–100 (laugh)]. If `NeedFaceAttributes` is not 1 or more than 5 faces are detected, this parameter will still be returned but meaningless.
	Expression *int64 `json:"Expression,omitempty" name:"Expression"`

	// Whether glasses are present. Valid values: [true,false]. If `NeedFaceAttributes` is not 1 or more than 5 faces are detected, this parameter will still be returned but meaningless.
	Glass *bool `json:"Glass,omitempty" name:"Glass"`

	// Vertical offset in degrees. Value range: [-30,30]. If `NeedFaceAttributes` is not 1 or more than 5 faces are detected, this parameter will still be returned but meaningless. 
	// You are recommended to select images in the [-10,10] range for adding faces.
	Pitch *int64 `json:"Pitch,omitempty" name:"Pitch"`

	// Horizontal offset in degrees. Value range: [-30,30]. If `NeedFaceAttributes` is not 1 or more than 5 faces are detected, this parameter will still be returned but meaningless. 
	// You are recommended to select images in the [-10,10] range for adding faces.
	Yaw *int64 `json:"Yaw,omitempty" name:"Yaw"`

	// Horizontal rotation in degrees. Value range: [-180,180]. If `NeedFaceAttributes` is not 1 or more than 5 faces are detected, this parameter will still be returned but meaningless.  
	// You are recommended to select images in the [-20,20] range for adding faces.
	Roll *int64 `json:"Roll,omitempty" name:"Roll"`

	// Beauty. Value range: [0–100]. If `NeedFaceAttributes` is not 1 or more than 5 faces are detected, this parameter will still be returned but meaningless.
	Beauty *int64 `json:"Beauty,omitempty" name:"Beauty"`

	// Whether hat is present. Valid values: [true,false]. If `NeedFaceAttributes` is not 1 or more than 5 faces are detected, this parameter will still be returned but meaningless.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Hat *bool `json:"Hat,omitempty" name:"Hat"`

	// Whether mask is present. Valid values: [true,false]. If `NeedFaceAttributes` is not 1 or more than 5 faces are detected, this parameter will still be returned but meaningless.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Mask *bool `json:"Mask,omitempty" name:"Mask"`

	// Hair information, including length, bang, and color. If `NeedFaceAttributes` is not 1 or more than 5 faces are detected, this parameter will still be returned but meaningless.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Hair *FaceHairAttributesInfo `json:"Hair,omitempty" name:"Hair"`

	// Whether the eyes are open. Valid values: [true,false]. As long as there is more than one eye closed, `false` will be returned. If `NeedFaceAttributes` is not 1 or more than 5 faces are detected, this parameter will still be returned but meaningless.
	// Note: this field may return null, indicating that no valid values can be obtained.
	EyeOpen *bool `json:"EyeOpen,omitempty" name:"EyeOpen"`
}

type FaceHairAttributesInfo struct {

	// 0: shaved head, 1: short hair, 2: medium hair, 3: long hair, 4: braid
	// Note: this field may return null, indicating that no valid values can be obtained.
	Length *int64 `json:"Length,omitempty" name:"Length"`

	// 0: with bangs, 1: no bangs
	// Note: this field may return null, indicating that no valid values can be obtained.
	Bang *int64 `json:"Bang,omitempty" name:"Bang"`

	// 0: black, 1: golden, 2: brown, 3: gray
	// Note: this field may return null, indicating that no valid values can be obtained.
	Color *int64 `json:"Color,omitempty" name:"Color"`
}

type FaceInfo struct {

	// 人脸框左上角横坐标。
	// 人脸框包含人脸五官位置并在此基础上进行一定的扩展，若人脸框超出图片范围，会导致坐标负值。 
	// 若需截取完整人脸，可以在完整分completeness满足需求的情况下，将负值坐标取0。
	X *int64 `json:"X,omitempty" name:"X"`

	// 人脸框左上角纵坐标。 
	// 人脸框包含人脸五官位置并在此基础上进行一定的扩展，若人脸框超出图片范围，会导致坐标负值。 
	// 若需截取完整人脸，可以在完整分completeness满足需求的情况下，将负值坐标取0。
	Y *int64 `json:"Y,omitempty" name:"Y"`

	// 人脸框宽度。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 人脸框高度。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 人脸属性信息，包含性别( gender )、年龄( age )、表情( expression )、 
	// 魅力( beauty )、眼镜( glass )、口罩（mask）、头发（hair）和姿态 (pitch，roll，yaw )。只有当 NeedFaceAttributes 设为 1 时才返回有效信息。
	FaceAttributesInfo *FaceAttributesInfo `json:"FaceAttributesInfo,omitempty" name:"FaceAttributesInfo"`

	// 人脸质量信息，包含质量分（score）、模糊分（sharpness）、光照分（brightness）、遮挡分（completeness）。只有当NeedFaceDetection设为1时才返回有效信息。
	FaceQualityInfo *FaceQualityInfo `json:"FaceQualityInfo,omitempty" name:"FaceQualityInfo"`
}

type FaceQualityCompleteness struct {

	// Eyebrow completeness. Value range: [0,100]. The higher the score, the higher the completeness. 
	// Reference range: [0,80], which means incomplete.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Eyebrow *int64 `json:"Eyebrow,omitempty" name:"Eyebrow"`

	// Eye completeness. Value range: [0,100]. The higher the score, the higher the completeness. 
	// Reference range: [0,80], which means incomplete.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Eye *int64 `json:"Eye,omitempty" name:"Eye"`

	// Nose completeness. Value range: [0,100]. The higher the score, the higher the completeness. 
	// Reference range: [0,60], which means incomplete.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Nose *int64 `json:"Nose,omitempty" name:"Nose"`

	// Cheek completeness. Value range: [0,100]. The higher the score, the higher the completeness. 
	// Reference range: [0,70], which means incomplete.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Cheek *int64 `json:"Cheek,omitempty" name:"Cheek"`

	// Mouth completeness. Value range: [0,100]. The higher the score, the higher the completeness. 
	// Reference range: [0,50], which means incomplete.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Mouth *int64 `json:"Mouth,omitempty" name:"Mouth"`

	// Chin completeness. Value range: [0,100]. The higher the score, the higher the completeness. 
	// Reference range: [0,70], which means incomplete.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Chin *int64 `json:"Chin,omitempty" name:"Chin"`
}

type FaceQualityInfo struct {

	// Quality score. Value range: [0,100]. It comprehensively evaluates whether the image quality is suitable for face recognition; the higher the score, the higher the quality. 
	// In normal cases, you only need to use `Score` as the overall quality standard score. Specific item scores such as `Sharpness`, `Brightness`, `Completeness` are for reference only.
	// Reference range: [0,40]: poor; [40,60]: fine; [60,80]: good; [80,100]: excellent. 
	// You are recommended to select images with a score above 70 for adding faces.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Score *int64 `json:"Score,omitempty" name:"Score"`

	// Sharpness. Value range: [0,100]. It evaluates the sharpness of the image. The higher the score, the sharper the image. 
	// Reference range: [0,40]: very blurry; [40,60]: blurry; [60,80]: fine; [80,100]: sharp. 
	// You are recommended to select images with a score above 80 for adding faces.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Sharpness *int64 `json:"Sharpness,omitempty" name:"Sharpness"`

	// Brightness. Value range: [0,100]. The brighter the image, the higher the score. 
	// Reference range: [0,30]: dark; [30,70]: normal; [70,100]: bright. 
	// You are recommended to select images in the [30,70] range for adding faces.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Brightness *int64 `json:"Brightness,omitempty" name:"Brightness"`

	// Completeness of facial features, which assesses the completeness of the eyebrows, eyes, nose, cheeks, mouth, and chin.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Completeness *FaceQualityCompleteness `json:"Completeness,omitempty" name:"Completeness"`
}

type FaceRect struct {

	// Horizontal coordinate of the top-left corner of face frame. 
	// The face frame encompasses the facial features and is extended accordingly. If it is larger than the image, the coordinates will be negative. 
	// If you want to capture a complete face, you can set the negative coordinates to 0 if the `completeness` score meets the requirement.
	X *int64 `json:"X,omitempty" name:"X"`

	// Vertical coordinate of the top-left corner of face frame. 
	// The face frame encompasses the facial features and is extended accordingly. If it is larger than the image, the coordinates will be negative. 
	// If you want to capture a complete face, you can set the negative coordinates to 0 if the `completeness` score meets the requirement.
	Y *int64 `json:"Y,omitempty" name:"Y"`

	// Face width
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// Face height
	Height *uint64 `json:"Height,omitempty" name:"Height"`
}

type FaceShape struct {

	// 21 points that describe the face contour.
	FaceProfile []*Point `json:"FaceProfile,omitempty" name:"FaceProfile" list`

	// 8 points that describe the left eye.
	LeftEye []*Point `json:"LeftEye,omitempty" name:"LeftEye" list`

	// 8 points that describe the right eye.
	RightEye []*Point `json:"RightEye,omitempty" name:"RightEye" list`

	// 8 points that describe the left eyebrow.
	LeftEyeBrow []*Point `json:"LeftEyeBrow,omitempty" name:"LeftEyeBrow" list`

	// 8 points that describe the right eyebrow.
	RightEyeBrow []*Point `json:"RightEyeBrow,omitempty" name:"RightEyeBrow" list`

	// 22 points that describe the mouth.
	Mouth []*Point `json:"Mouth,omitempty" name:"Mouth" list`

	// 13 points that describe the nose.
	Nose []*Point `json:"Nose,omitempty" name:"Nose" list`

	// 1 point that describes the left pupil.
	LeftPupil []*Point `json:"LeftPupil,omitempty" name:"LeftPupil" list`

	// 1 point that describes the right pupil.
	RightPupil []*Point `json:"RightPupil,omitempty" name:"RightPupil" list`
}

type GetCheckSimilarPersonJobIdListRequest struct {
	*tchttp.BaseRequest

	// Starting number. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 10. Maximum value: 1000.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetCheckSimilarPersonJobIdListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetCheckSimilarPersonJobIdListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetCheckSimilarPersonJobIdListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of duplicate person check task information.
		JobIdInfos []*JobIdInfo `json:"JobIdInfos,omitempty" name:"JobIdInfos" list`

		// Total number of duplicate check tasks.
		JobIdNum *uint64 `json:"JobIdNum,omitempty" name:"JobIdNum"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCheckSimilarPersonJobIdListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetCheckSimilarPersonJobIdListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetGroupInfoRequest struct {
	*tchttp.BaseRequest

	// Group ID.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *GetGroupInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetGroupInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Group name
		GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

		// Group ID
		GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

		// Custom group description field
		GroupExDescriptions []*string `json:"GroupExDescriptions,omitempty" name:"GroupExDescriptions" list`

		// Group remarks
		Tag *string `json:"Tag,omitempty" name:"Tag"`

		// Algorithm model version used for face recognition.
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// Group creation time and date (`CreationTimestamp`), whose value is the number of milliseconds between the UNIX epoch time and the group creation time.
		CreationTimestamp *uint64 `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetGroupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetGroupInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetGroupListRequest struct {
	*tchttp.BaseRequest

	// Starting number. Default value: 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 10. Maximum value: 1000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetGroupListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetGroupListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned group information
		GroupInfos []*GroupInfo `json:"GroupInfos,omitempty" name:"GroupInfos" list`

		// Total number of groups
	// Note: this field may return null, indicating that no valid values can be obtained.
		GroupNum *uint64 `json:"GroupNum,omitempty" name:"GroupNum"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetGroupListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPersonBaseInfoRequest struct {
	*tchttp.BaseRequest

	// Person ID
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`
}

func (r *GetPersonBaseInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPersonBaseInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPersonBaseInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Person name
		PersonName *string `json:"PersonName,omitempty" name:"PersonName"`

		// Person gender
		Gender *int64 `json:"Gender,omitempty" name:"Gender"`

		// List of the IDs of included faces
		FaceIds []*string `json:"FaceIds,omitempty" name:"FaceIds" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPersonBaseInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPersonBaseInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPersonGroupInfoRequest struct {
	*tchttp.BaseRequest

	// Person ID
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// Starting number. Default value: 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 10. Maximum value: 100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetPersonGroupInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPersonGroupInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPersonGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of groups containing this person and their description fields
		PersonGroupInfos []*PersonGroupInfo `json:"PersonGroupInfos,omitempty" name:"PersonGroupInfos" list`

		// Total number of groups
	// Note: this field may return null, indicating that no valid values can be obtained.
		GroupNum *uint64 `json:"GroupNum,omitempty" name:"GroupNum"`

		// Algorithm model version used by the Face Recognition service.
	// Note: this field may return null, indicating that no valid values can be obtained.
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPersonGroupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPersonGroupInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPersonListNumRequest struct {
	*tchttp.BaseRequest

	// Group ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *GetPersonListNumRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPersonListNumRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPersonListNumResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of persons
		PersonNum *uint64 `json:"PersonNum,omitempty" name:"PersonNum"`

		// Number of faces
		FaceNum *uint64 `json:"FaceNum,omitempty" name:"FaceNum"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPersonListNumResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPersonListNumResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPersonListRequest struct {
	*tchttp.BaseRequest

	// Group ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Starting number. Default value: 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 10. Maximum value: 1000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetPersonListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPersonListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPersonListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned person information
		PersonInfos []*PersonInfo `json:"PersonInfos,omitempty" name:"PersonInfos" list`

		// Number of persons in the group
	// Note: this field may return null, indicating that no valid values can be obtained.
		PersonNum *uint64 `json:"PersonNum,omitempty" name:"PersonNum"`

		// Number of faces in the group
	// Note: this field may return null, indicating that no valid values can be obtained.
		FaceNum *uint64 `json:"FaceNum,omitempty" name:"FaceNum"`

		// Algorithm model version used for face recognition.
	// Note: this field may return null, indicating that no valid values can be obtained.
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPersonListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPersonListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetSimilarPersonResultRequest struct {
	*tchttp.BaseRequest

	// Duplicate check task ID, which is used to query and get the progress and result of the task.
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *GetSimilarPersonResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetSimilarPersonResultRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetSimilarPersonResultResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Duplicate check task completion progress. Value range: [0.0,100.0]. `SimilarPersons` takes effect only if this parameter value is 100.
		Progress *float64 `json:"Progress,omitempty" name:"Progress"`

		// Temporary download link for the information file of the persons suspected to be duplicate. The validity period is 5 minutes, and the result file retention duration is 90 days.
	// The file content is an array of `SimilarPerson` values.
		SimilarPersonsUrl *string `json:"SimilarPersonsUrl,omitempty" name:"SimilarPersonsUrl"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSimilarPersonResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetSimilarPersonResultResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GroupCandidate struct {

	// Group ID.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Most matching candidates recognized.
	Candidates []*Candidate `json:"Candidates,omitempty" name:"Candidates" list`
}

type GroupExDescriptionInfo struct {

	// Custom group description field index, whose value starts from 0
	// Note: this field may return null, indicating that no valid values can be obtained.
	GroupExDescriptionIndex *uint64 `json:"GroupExDescriptionIndex,omitempty" name:"GroupExDescriptionIndex"`

	// Content of the custom group description field to be updated
	GroupExDescription *string `json:"GroupExDescription,omitempty" name:"GroupExDescription"`
}

type GroupInfo struct {

	// Group name
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Group ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Custom group description field
	// Note: this field may return null, indicating that no valid values can be obtained.
	GroupExDescriptions []*string `json:"GroupExDescriptions,omitempty" name:"GroupExDescriptions" list`

	// Group remarks
	// Note: this field may return null, indicating that no valid values can be obtained.
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// Algorithm model version used for face recognition.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	// Group creation time and date (`CreationTimestamp`), whose value is the number of milliseconds between the UNIX epoch time and the group creation time. 
	// The UNIX epoch time is 00:00:00, Thursday, January 1, 1970, Coordinated Universal Time (UTC). For more information, please see the UNIX time document.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreationTimestamp *uint64 `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
}

type JobIdInfo struct {

	// Duplicate check task ID, which is used to query and get the progress and result of the task.
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Start time. 
	// The `StartTime` value is the number of milliseconds between the UNIX epoch time and the group creation time. 
	// The UNIX epoch time is 00:00:00, Thursday, January 1, 1970, Coordinated Universal Time (UTC). 
	// For more information, please see the UNIX time document.
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// Whether the duplicate check task is completed. 0: completed; 1: uncompleted; 2: failed.
	JobStatus *int64 `json:"JobStatus,omitempty" name:"JobStatus"`
}

type ModifyGroupRequest struct {
	*tchttp.BaseRequest

	// Group ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Group name
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Custom description field of the group to be modified, which is a `key-value`
	GroupExDescriptionInfos []*GroupExDescriptionInfo `json:"GroupExDescriptionInfos,omitempty" name:"GroupExDescriptionInfos" list`

	// Group remarks
	Tag *string `json:"Tag,omitempty" name:"Tag"`
}

func (r *ModifyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPersonBaseInfoRequest struct {
	*tchttp.BaseRequest

	// Person ID
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// Name of the person to be modified
	PersonName *string `json:"PersonName,omitempty" name:"PersonName"`

	// Gender of the person to be modified
	Gender *int64 `json:"Gender,omitempty" name:"Gender"`
}

func (r *ModifyPersonBaseInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPersonBaseInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPersonBaseInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPersonBaseInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPersonBaseInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPersonGroupInfoRequest struct {
	*tchttp.BaseRequest

	// Group ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Person ID
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// Custom description field of the person to be modified, which is a `key-value`
	PersonExDescriptionInfos []*PersonExDescriptionInfo `json:"PersonExDescriptionInfos,omitempty" name:"PersonExDescriptionInfos" list`
}

func (r *ModifyPersonGroupInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPersonGroupInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPersonGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPersonGroupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPersonGroupInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PersonExDescriptionInfo struct {

	// Person description field index, whose value starts from 0
	// Note: this field may return null, indicating that no valid values can be obtained.
	PersonExDescriptionIndex *uint64 `json:"PersonExDescriptionIndex,omitempty" name:"PersonExDescriptionIndex"`

	// Content of the person description field to be updated
	PersonExDescription *string `json:"PersonExDescription,omitempty" name:"PersonExDescription"`
}

type PersonGroupInfo struct {

	// ID of the group that contains this person
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Content of person description field
	PersonExDescriptions []*string `json:"PersonExDescriptions,omitempty" name:"PersonExDescriptions" list`
}

type PersonInfo struct {

	// Person name
	PersonName *string `json:"PersonName,omitempty" name:"PersonName"`

	// Person ID
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// Person gender
	Gender *int64 `json:"Gender,omitempty" name:"Gender"`

	// Content of person description field
	PersonExDescriptions []*string `json:"PersonExDescriptions,omitempty" name:"PersonExDescriptions" list`

	// List of contained face images
	FaceIds []*string `json:"FaceIds,omitempty" name:"FaceIds" list`

	// Person creation time and date (`CreationTimestamp`), whose value is the number of milliseconds between the UNIX epoch time and the group creation time. 
	// The UNIX epoch time is 00:00:00, Thursday, January 1, 1970, Coordinated Universal Time (UTC). For more information, please see the UNIX time document.
	CreationTimestamp *uint64 `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
}

type Point struct {

	// X coordinate
	X *int64 `json:"X,omitempty" name:"X"`

	// Y coordinate
	Y *int64 `json:"Y,omitempty" name:"Y"`
}

type Result struct {

	// Most matching candidate recognized
	Candidates []*Candidate `json:"Candidates,omitempty" name:"Candidates" list`

	// Position of detected face frame
	FaceRect *FaceRect `json:"FaceRect,omitempty" name:"FaceRect"`

	// Status return code of detected face image. 0: normal. 
	// -1601: the image quality control requirement is not met; in this case, `Candidate` is empty.
	RetCode *int64 `json:"RetCode,omitempty" name:"RetCode"`
}

type ResultsReturnsByGroup struct {

	// Position of detected face frame.
	FaceRect *FaceRect `json:"FaceRect,omitempty" name:"FaceRect"`

	// Recognition result.
	GroupCandidates []*GroupCandidate `json:"GroupCandidates,omitempty" name:"GroupCandidates" list`

	// Status return code of detected face image. 0: normal. 
	// -1601: the image quality control requirement is not met; in this case, `Candidate` is empty.
	RetCode *int64 `json:"RetCode,omitempty" name:"RetCode"`
}

type SearchFacesRequest struct {
	*tchttp.BaseRequest

	// List of groups to be searched in. Up to 100 groups are supported.
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds" list`

	// Base64-encoded image data, which cannot exceed 5 MB.
	// .png, .jpg, .jpeg, and .bmp images are supported, while .gif images are not.
	Image *string `json:"Image,omitempty" name:"Image"`

	// Image URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used.  
	// You are recommended to store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. 
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// .png, .jpg, .jpeg, and .bmp images are supported, while .gif images are not.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Maximum number of recognizable faces. Default value: 1 (i.e., detecting only the face with the largest size in the image). Maximum value: 10. 
	// `MaxFaceNum` is used to control the number of faces to be searched for if there are multiple faces in the input image to be recognized. 
	// For example, if the input image in `Image` or `Url` contains multiple faces and `MaxFaceNum` is 5, top 5 faces with the largest size in the image will be recognized.
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitempty" name:"MaxFaceNum"`

	// Minimum height and width of face in px. Default value: 34. Face images whose value is below 34 cannot be recognized. You are recommended to set this parameter to 80.
	MinFaceSize *uint64 `json:"MinFaceSize,omitempty" name:"MinFaceSize"`

	// Number of the most similar persons returned for faces recognized in one single image. Default value: 5. Maximum value: 100. 
	// For example, if `MaxFaceNum` is 1 and `MaxPersonNum` is 8, information of the top 8 most similar persons will be returned.
	// The greater the value, the longer the processing time. You are recommended to set a value below 10.
	MaxPersonNum *uint64 `json:"MaxPersonNum,omitempty" name:"MaxPersonNum"`

	// Whether to return person details. 0: no; 1: yes. Default value: 0. Other values will be considered as 0 by default
	NeedPersonInfo *int64 `json:"NeedPersonInfo,omitempty" name:"NeedPersonInfo"`

	// Image quality control. 
	// 0: no control. 
	// 1: low quality requirement. The image has one or more of the following problems: extreme blurriness, covered eyes, covered nose, and covered mouth. 
	// 2: average quality requirement. The image has at least three of following problems: extreme brightness, extreme dimness, blurriness or average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 3: high quality requirement. The image has one to two of following problems: extreme brightness, extreme dimness, average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 4: very high quality requirement. The image is optimal in all dimensions or only has a slight problem in one dimension. 
	// Default value: 0. 
	// If the image quality does not meet the requirement, the returned result will prompt that the detected image quality is unsatisfactory.
	QualityControl *uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// In the output parameter `Score`, the result will be returned only if the result value is above the `FaceMatchThreshold` value. Default value: 0.
	FaceMatchThreshold *float64 `json:"FaceMatchThreshold,omitempty" name:"FaceMatchThreshold"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

func (r *SearchFacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchFacesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchFacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Recognition result.
		Results []*Result `json:"Results,omitempty" name:"Results" list`

		// Number of faces included in searched groups.
		FaceNum *uint64 `json:"FaceNum,omitempty" name:"FaceNum"`

		// Algorithm model version used for face recognition.
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchFacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchFacesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchFacesReturnsByGroupRequest struct {
	*tchttp.BaseRequest

	// List of groups to be searched in. Up to 60 groups are supported.
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds" list`

	// Base64-encoded image data, which cannot exceed 5 MB.
	// .png, .jpg, .jpeg, and .bmp images are supported, while .gif images are not.
	Image *string `json:"Image,omitempty" name:"Image"`

	// Image URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used.
	// You are recommended to store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// .png, .jpg, .jpeg, and .bmp images are supported, while .gif images are not.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Maximum number of recognizable faces. Default value: 1 (i.e., detecting only the face with the largest size in the image). Maximum value: 10.
	// `MaxFaceNum` is used to control the number of faces to be searched for if there are multiple faces in the input image to be recognized.
	// For example, if the input image in `Image` or `Url` contains multiple faces and `MaxFaceNum` is 5, top 5 faces with the largest size in the image will be recognized.
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitempty" name:"MaxFaceNum"`

	// Minimum height and width of face in px. Default value: 34. A value below 34 will affect the search accuracy. You are recommended to set this parameter to 80.
	MinFaceSize *uint64 `json:"MinFaceSize,omitempty" name:"MinFaceSize"`

	// Detected faces, which is corresponding to the maximum number of returned most matching persons. Default value: 5. Maximum value: 10.  
	// For example, if `MaxFaceNum` is 3 and `MaxPersonNum` is 5, up to 15 (3 * 5) persons will be returned.
	MaxPersonNumPerGroup *uint64 `json:"MaxPersonNumPerGroup,omitempty" name:"MaxPersonNumPerGroup"`

	// Whether to return person details. 0: no; 1: yes. Default value: 0. Other values will be considered as 0 by default
	NeedPersonInfo *int64 `json:"NeedPersonInfo,omitempty" name:"NeedPersonInfo"`

	// Image quality control. 
	// 0: no control. 
	// 1: low quality requirement. The image has one or more of the following problems: extreme blurriness, covered eyes, covered nose, and covered mouth. 
	// 2: average quality requirement. The image has at least three of following problems: extreme brightness, extreme dimness, blurriness or average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 3: high quality requirement. The image has one to two of following problems: extreme brightness, extreme dimness, average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 4: very high quality requirement. The image is optimal in all dimensions or only has a slight problem in one dimension. 
	// Default value: 0. 
	// If the image quality does not meet the requirement, the returned result will prompt that the detected image quality is unsatisfactory.
	QualityControl *uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// In the output parameter `Score`, the result will be returned only if the result value is greater than or equal to the `FaceMatchThreshold` value.
	// Default value: 0.
	// Value range: [0.0,100.0).
	FaceMatchThreshold *float64 `json:"FaceMatchThreshold,omitempty" name:"FaceMatchThreshold"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

func (r *SearchFacesReturnsByGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchFacesReturnsByGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchFacesReturnsByGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of faces included in searched groups.
		FaceNum *uint64 `json:"FaceNum,omitempty" name:"FaceNum"`

		// Recognition result.
		ResultsReturnsByGroup []*ResultsReturnsByGroup `json:"ResultsReturnsByGroup,omitempty" name:"ResultsReturnsByGroup" list`

		// Algorithm model version used for face recognition.
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchFacesReturnsByGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchFacesReturnsByGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchPersonsRequest struct {
	*tchttp.BaseRequest

	// List of groups to be searched in. Up to 100 groups are supported.
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds" list`

	// Base64-encoded image data, which cannot exceed 5 MB.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// .png, .jpg, .jpeg, and .bmp images are supported, while .gif images are not.
	Image *string `json:"Image,omitempty" name:"Image"`

	// Image URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used.
	// You are recommended to store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// .png, .jpg, .jpeg, and .bmp images are supported, while .gif images are not.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Maximum number of recognizable faces. Default value: 1 (i.e., detecting only the face with the largest size in the image). Maximum value: 10.
	// `MaxFaceNum` is used to control the number of faces to be searched for if there are multiple faces in the input image to be recognized.
	// For example, if the input image in `Image` or `Url` contains multiple faces and `MaxFaceNum` is 5, top 5 faces with the largest size in the image will be recognized.
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitempty" name:"MaxFaceNum"`

	// Minimum height and width of face in px. Default value: 34. A value below 34 will affect the search accuracy. You are recommended to set this parameter to 80.
	MinFaceSize *uint64 `json:"MinFaceSize,omitempty" name:"MinFaceSize"`

	// Number of the most similar persons returned for faces recognized in one single image. Default value: 5. Maximum value: 100.
	// For example, if `MaxFaceNum` is 1 and `MaxPersonNum` is 8, information of the top 8 most similar persons will be returned.
	// The greater the value, the longer the processing time. You are recommended to set a value below 10.
	MaxPersonNum *uint64 `json:"MaxPersonNum,omitempty" name:"MaxPersonNum"`

	// Image quality control. 
	// 0: no control. 
	// 1: low quality requirement. The image has one or more of the following problems: extreme blurriness, covered eyes, covered nose, and covered mouth. 
	// 2: average quality requirement. The image has at least three of following problems: extreme brightness, extreme dimness, blurriness or average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 3: high quality requirement. The image has one to two of following problems: extreme brightness, extreme dimness, average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 4: very high quality requirement. The image is optimal in all dimensions or only has a slight problem in one dimension. 
	// Default value: 0. 
	// If the image quality does not meet the requirement, the returned result will prompt that the detected image quality is unsatisfactory.
	QualityControl *uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// In the output parameter `Score`, the result will be returned only if the result value is greater than or equal to the `FaceMatchThreshold` value. Default value: 0. Value range: [0.0,100.0).
	FaceMatchThreshold *float64 `json:"FaceMatchThreshold,omitempty" name:"FaceMatchThreshold"`

	// Whether to return person details. 0: no; 1: yes. Default value: 0. Other values will be considered as 0 by default
	NeedPersonInfo *int64 `json:"NeedPersonInfo,omitempty" name:"NeedPersonInfo"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

func (r *SearchPersonsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchPersonsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchPersonsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Recognition result.
		Results []*Result `json:"Results,omitempty" name:"Results" list`

		// Number of the persons included in searched groups. If the quality of all faces in the input image does not meet the requirement, 0 will be returned.
		PersonNum *uint64 `json:"PersonNum,omitempty" name:"PersonNum"`

		// Algorithm model version used for face recognition.
	// Note: this field may return null, indicating that no valid values can be obtained.
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchPersonsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchPersonsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchPersonsReturnsByGroupRequest struct {
	*tchttp.BaseRequest

	// List of groups to be searched in. Up to 60 groups are supported.
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds" list`

	// Base64-encoded image data, which cannot exceed 5 MB.
	// .png, .jpg, .jpeg, and .bmp images are supported, while .gif images are not.
	Image *string `json:"Image,omitempty" name:"Image"`

	// Image URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used.
	// You are recommended to store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability.
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// .png, .jpg, .jpeg, and .bmp images are supported, while .gif images are not.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Maximum number of recognizable faces. Default value: 1 (i.e., detecting only the face with the largest size in the image). Maximum value: 10.
	// `MaxFaceNum` is used to control the number of faces to be searched for if there are multiple faces in the input image to be recognized.
	// For example, if the input image in `Image` or `Url` contains multiple faces and `MaxFaceNum` is 5, top 5 faces with the largest size in the image will be recognized.
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitempty" name:"MaxFaceNum"`

	// Minimum height and width of face in px. Default value: 34. A value below 34 will affect the search accuracy. You are recommended to set this parameter to 80.
	MinFaceSize *uint64 `json:"MinFaceSize,omitempty" name:"MinFaceSize"`

	// Detected faces, which is corresponding to the maximum number of returned most matching persons. Default value: 5. Maximum value: 10.  
	// For example, if `MaxFaceNum` is 3, `MaxPersonNumPerGroup` is 3, and the `GroupIds` length is 3, up to 45 (3 * 5 * 3) persons will be returned.
	MaxPersonNumPerGroup *uint64 `json:"MaxPersonNumPerGroup,omitempty" name:"MaxPersonNumPerGroup"`

	// Image quality control. 
	// 0: no control. 
	// 1: low quality requirement. The image has one or more of the following problems: extreme blurriness, covered eyes, covered nose, and covered mouth. 
	// 2: average quality requirement. The image has at least three of following problems: extreme brightness, extreme dimness, blurriness or average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 3: high quality requirement. The image has one to two of following problems: extreme brightness, extreme dimness, average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 4: very high quality requirement. The image is optimal in all dimensions or only has a slight problem in one dimension. 
	// Default value: 0. 
	// If the image quality does not meet the requirement, the returned result will prompt that the detected image quality is unsatisfactory.
	QualityControl *uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// In the output parameter `Score`, the result will be returned only if the result value is above the `FaceMatchThreshold` value. Default value: 0.
	FaceMatchThreshold *float64 `json:"FaceMatchThreshold,omitempty" name:"FaceMatchThreshold"`

	// Whether to return person details. 0: no; 1: yes. Default value: 0. Other values will be considered as 0 by default
	NeedPersonInfo *int64 `json:"NeedPersonInfo,omitempty" name:"NeedPersonInfo"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

func (r *SearchPersonsReturnsByGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchPersonsReturnsByGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchPersonsReturnsByGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of the persons included in searched groups. If the quality of all faces in the input image does not meet the requirement, 0 will be returned.
		PersonNum *uint64 `json:"PersonNum,omitempty" name:"PersonNum"`

		// Recognition result.
		ResultsReturnsByGroup []*ResultsReturnsByGroup `json:"ResultsReturnsByGroup,omitempty" name:"ResultsReturnsByGroup" list`

		// Algorithm model version used for face recognition.
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchPersonsReturnsByGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchPersonsReturnsByGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VerifyFaceRequest struct {
	*tchttp.BaseRequest

	// ID of the person to be verified. For more information on `PersonId`, please see the group management APIs.
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// Base64-encoded image data, which cannot exceed 5 MB.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// .png, .jpg, .jpeg, and .bmp images are supported, while .gif images are not.
	Image *string `json:"Image,omitempty" name:"Image"`

	// Image URL. The image cannot exceed 5 MB in size after being Base64-encoded.
	// Either `Url` or `Image` must be provided; if both are provided, only `Url` will be used.  
	// You are recommended to store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. 
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// .png, .jpg, .jpeg, and .bmp images are supported, while .gif images are not.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Image quality control. 
	// 0: no control. 
	// 1: low quality requirement. The image has one or more of the following problems: extreme blurriness, covered eyes, covered nose, and covered mouth. 
	// 2: average quality requirement. The image has at least three of following problems: extreme brightness, extreme dimness, blurriness or average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 3: high quality requirement. The image has one to two of following problems: extreme brightness, extreme dimness, average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 4: very high quality requirement. The image is optimal in all dimensions or only has a slight problem in one dimension. 
	// Default value: 0. 
	// If the image quality does not meet the requirement, the returned result will prompt that the detected image quality is unsatisfactory.
	QualityControl *uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

func (r *VerifyFaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *VerifyFaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VerifyFaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Similarity between given face image and `PersonId`. If there are multiple faces under the `PersonId`, the score of the highest similarity will be returned.
	// 
	// The returned similarity score varies by algorithm version.
	// If you need to verify whether the faces in the two images are the same person, then the 0.1%, 0.01%, and 0.001% FARs on v3.0 correspond to scores of 40, 50, and 60, respectively. Generally, if the score is above 50, it can be judged that they are the same person.
	// The 0.1%, 0.01%, and 0.001% FARs on v2.0 correspond to scores of 70, 80, and 90, respectively. Generally, if the score is above 80, it can be judged that they are the same person.
		Score *float64 `json:"Score,omitempty" name:"Score"`

		// Whether the person in the image matches the `PersonId`.
		IsMatch *bool `json:"IsMatch,omitempty" name:"IsMatch"`

		// Algorithm model version used for face recognition.
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VerifyFaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *VerifyFaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VerifyPersonRequest struct {
	*tchttp.BaseRequest

	// Base64-encoded data of image.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// .png, .jpg, .jpeg, and .bmp images are supported, while .gif images are not.
	Image *string `json:"Image,omitempty" name:"Image"`

	// Image URL. Either the `Url` or `Image` of the image must be provided; if both are provided, only `Url` will be used. 
	// You are recommended to store the image in Tencent Cloud, as a Tencent Cloud URL can guarantee higher download speed and stability. 
	// The download speed and stability of non-Tencent Cloud URLs may be low.
	// If there are multiple faces in the image, only the face with the largest size will be selected.
	// .png, .jpg, .jpeg, and .bmp images are supported, while .gif images are not.
	Url *string `json:"Url,omitempty" name:"Url"`

	// ID of the person to be verified. For more information on `PersonId`, please see the group management APIs.
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// Image quality control. 
	// 0: no control. 
	// 1: low quality requirement. The image has one or more of the following problems: extreme blurriness, covered eyes, covered nose, and covered mouth. 
	// 2: average quality requirement. The image has at least three of following problems: extreme brightness, extreme dimness, blurriness or average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 3: high quality requirement. The image has one to two of following problems: extreme brightness, extreme dimness, average blurriness, covered eyebrows, covered cheeks, and covered chin. 
	// 4: very high quality requirement. The image is optimal in all dimensions or only has a slight problem in one dimension. 
	// Default value: 0. 
	// If the image quality does not meet the requirement, the returned result will prompt that the detected image quality is unsatisfactory.
	QualityControl *uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// Whether to enable the support for rotated image recognition. 0: no; 1: yes. Default value: 0. When the face in the image is rotated and the image has no EXIF information, if this parameter is not enabled, the face in the image cannot be correctly detected and recognized. If you are sure that the input image contains EXIF information or the face in the image will not be rotated, do not enable this parameter, as the overall time consumption may increase by hundreds of milliseconds after it is enabled.
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

func (r *VerifyPersonRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *VerifyPersonRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VerifyPersonResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Similarity between given face image and `PersonId`. If there are multiple faces under the `PersonId`, their information will be fused for verification.
		Score *float64 `json:"Score,omitempty" name:"Score"`

		// Whether the person in the image matches the `PersonId`.
		IsMatch *bool `json:"IsMatch,omitempty" name:"IsMatch"`

		// Algorithm model version used for face recognition.
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VerifyPersonResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *VerifyPersonResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
