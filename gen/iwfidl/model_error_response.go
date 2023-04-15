/*
Workflow APIs

This APIs for iwf SDKs to operate workflows

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iwfidl

import (
	"encoding/json"
)

// checks if the ErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorResponse{}

// ErrorResponse struct for ErrorResponse
type ErrorResponse struct {
	Detail                    *string         `json:"detail,omitempty"`
	SubStatus                 *ErrorSubStatus `json:"subStatus,omitempty"`
	OriginalWorkerErrorDetail *string         `json:"originalWorkerErrorDetail,omitempty"`
	OriginalWorkerErrorType   *string         `json:"originalWorkerErrorType,omitempty"`
	OriginalWorkerErrorStatus *int32          `json:"originalWorkerErrorStatus,omitempty"`
}

// NewErrorResponse instantiates a new ErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorResponse() *ErrorResponse {
	this := ErrorResponse{}
	return &this
}

// NewErrorResponseWithDefaults instantiates a new ErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorResponseWithDefaults() *ErrorResponse {
	this := ErrorResponse{}
	return &this
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ErrorResponse) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ErrorResponse) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ErrorResponse) SetDetail(v string) {
	o.Detail = &v
}

// GetSubStatus returns the SubStatus field value if set, zero value otherwise.
func (o *ErrorResponse) GetSubStatus() ErrorSubStatus {
	if o == nil || IsNil(o.SubStatus) {
		var ret ErrorSubStatus
		return ret
	}
	return *o.SubStatus
}

// GetSubStatusOk returns a tuple with the SubStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetSubStatusOk() (*ErrorSubStatus, bool) {
	if o == nil || IsNil(o.SubStatus) {
		return nil, false
	}
	return o.SubStatus, true
}

// HasSubStatus returns a boolean if a field has been set.
func (o *ErrorResponse) HasSubStatus() bool {
	if o != nil && !IsNil(o.SubStatus) {
		return true
	}

	return false
}

// SetSubStatus gets a reference to the given ErrorSubStatus and assigns it to the SubStatus field.
func (o *ErrorResponse) SetSubStatus(v ErrorSubStatus) {
	o.SubStatus = &v
}

// GetOriginalWorkerErrorDetail returns the OriginalWorkerErrorDetail field value if set, zero value otherwise.
func (o *ErrorResponse) GetOriginalWorkerErrorDetail() string {
	if o == nil || IsNil(o.OriginalWorkerErrorDetail) {
		var ret string
		return ret
	}
	return *o.OriginalWorkerErrorDetail
}

// GetOriginalWorkerErrorDetailOk returns a tuple with the OriginalWorkerErrorDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetOriginalWorkerErrorDetailOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalWorkerErrorDetail) {
		return nil, false
	}
	return o.OriginalWorkerErrorDetail, true
}

// HasOriginalWorkerErrorDetail returns a boolean if a field has been set.
func (o *ErrorResponse) HasOriginalWorkerErrorDetail() bool {
	if o != nil && !IsNil(o.OriginalWorkerErrorDetail) {
		return true
	}

	return false
}

// SetOriginalWorkerErrorDetail gets a reference to the given string and assigns it to the OriginalWorkerErrorDetail field.
func (o *ErrorResponse) SetOriginalWorkerErrorDetail(v string) {
	o.OriginalWorkerErrorDetail = &v
}

// GetOriginalWorkerErrorType returns the OriginalWorkerErrorType field value if set, zero value otherwise.
func (o *ErrorResponse) GetOriginalWorkerErrorType() string {
	if o == nil || IsNil(o.OriginalWorkerErrorType) {
		var ret string
		return ret
	}
	return *o.OriginalWorkerErrorType
}

// GetOriginalWorkerErrorTypeOk returns a tuple with the OriginalWorkerErrorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetOriginalWorkerErrorTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalWorkerErrorType) {
		return nil, false
	}
	return o.OriginalWorkerErrorType, true
}

// HasOriginalWorkerErrorType returns a boolean if a field has been set.
func (o *ErrorResponse) HasOriginalWorkerErrorType() bool {
	if o != nil && !IsNil(o.OriginalWorkerErrorType) {
		return true
	}

	return false
}

// SetOriginalWorkerErrorType gets a reference to the given string and assigns it to the OriginalWorkerErrorType field.
func (o *ErrorResponse) SetOriginalWorkerErrorType(v string) {
	o.OriginalWorkerErrorType = &v
}

// GetOriginalWorkerErrorStatus returns the OriginalWorkerErrorStatus field value if set, zero value otherwise.
func (o *ErrorResponse) GetOriginalWorkerErrorStatus() int32 {
	if o == nil || IsNil(o.OriginalWorkerErrorStatus) {
		var ret int32
		return ret
	}
	return *o.OriginalWorkerErrorStatus
}

// GetOriginalWorkerErrorStatusOk returns a tuple with the OriginalWorkerErrorStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetOriginalWorkerErrorStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.OriginalWorkerErrorStatus) {
		return nil, false
	}
	return o.OriginalWorkerErrorStatus, true
}

// HasOriginalWorkerErrorStatus returns a boolean if a field has been set.
func (o *ErrorResponse) HasOriginalWorkerErrorStatus() bool {
	if o != nil && !IsNil(o.OriginalWorkerErrorStatus) {
		return true
	}

	return false
}

// SetOriginalWorkerErrorStatus gets a reference to the given int32 and assigns it to the OriginalWorkerErrorStatus field.
func (o *ErrorResponse) SetOriginalWorkerErrorStatus(v int32) {
	o.OriginalWorkerErrorStatus = &v
}

func (o ErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	if !IsNil(o.SubStatus) {
		toSerialize["subStatus"] = o.SubStatus
	}
	if !IsNil(o.OriginalWorkerErrorDetail) {
		toSerialize["originalWorkerErrorDetail"] = o.OriginalWorkerErrorDetail
	}
	if !IsNil(o.OriginalWorkerErrorType) {
		toSerialize["originalWorkerErrorType"] = o.OriginalWorkerErrorType
	}
	if !IsNil(o.OriginalWorkerErrorStatus) {
		toSerialize["originalWorkerErrorStatus"] = o.OriginalWorkerErrorStatus
	}
	return toSerialize, nil
}

type NullableErrorResponse struct {
	value *ErrorResponse
	isSet bool
}

func (v NullableErrorResponse) Get() *ErrorResponse {
	return v.value
}

func (v *NullableErrorResponse) Set(val *ErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorResponse(val *ErrorResponse) *NullableErrorResponse {
	return &NullableErrorResponse{value: val, isSet: true}
}

func (v NullableErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
