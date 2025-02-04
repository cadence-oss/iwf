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

// checks if the WorkflowConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowConfig{}

// WorkflowConfig struct for WorkflowConfig
type WorkflowConfig struct {
	DisableSystemSearchAttribute *bool                 `json:"disableSystemSearchAttribute,omitempty"`
	ExecutingStateIdMode         *ExecutingStateIdMode `json:"executingStateIdMode,omitempty"`
	ContinueAsNewThreshold       *int32                `json:"continueAsNewThreshold,omitempty"`
	ContinueAsNewPageSizeInBytes *int32                `json:"continueAsNewPageSizeInBytes,omitempty"`
	OptimizeActivity             *bool                 `json:"optimizeActivity,omitempty"`
	OptimizeTimer                *bool                 `json:"optimizeTimer,omitempty"`
}

// NewWorkflowConfig instantiates a new WorkflowConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowConfig() *WorkflowConfig {
	this := WorkflowConfig{}
	return &this
}

// NewWorkflowConfigWithDefaults instantiates a new WorkflowConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowConfigWithDefaults() *WorkflowConfig {
	this := WorkflowConfig{}
	return &this
}

// GetDisableSystemSearchAttribute returns the DisableSystemSearchAttribute field value if set, zero value otherwise.
func (o *WorkflowConfig) GetDisableSystemSearchAttribute() bool {
	if o == nil || IsNil(o.DisableSystemSearchAttribute) {
		var ret bool
		return ret
	}
	return *o.DisableSystemSearchAttribute
}

// GetDisableSystemSearchAttributeOk returns a tuple with the DisableSystemSearchAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowConfig) GetDisableSystemSearchAttributeOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableSystemSearchAttribute) {
		return nil, false
	}
	return o.DisableSystemSearchAttribute, true
}

// HasDisableSystemSearchAttribute returns a boolean if a field has been set.
func (o *WorkflowConfig) HasDisableSystemSearchAttribute() bool {
	if o != nil && !IsNil(o.DisableSystemSearchAttribute) {
		return true
	}

	return false
}

// SetDisableSystemSearchAttribute gets a reference to the given bool and assigns it to the DisableSystemSearchAttribute field.
func (o *WorkflowConfig) SetDisableSystemSearchAttribute(v bool) {
	o.DisableSystemSearchAttribute = &v
}

// GetExecutingStateIdMode returns the ExecutingStateIdMode field value if set, zero value otherwise.
func (o *WorkflowConfig) GetExecutingStateIdMode() ExecutingStateIdMode {
	if o == nil || IsNil(o.ExecutingStateIdMode) {
		var ret ExecutingStateIdMode
		return ret
	}
	return *o.ExecutingStateIdMode
}

// GetExecutingStateIdModeOk returns a tuple with the ExecutingStateIdMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowConfig) GetExecutingStateIdModeOk() (*ExecutingStateIdMode, bool) {
	if o == nil || IsNil(o.ExecutingStateIdMode) {
		return nil, false
	}
	return o.ExecutingStateIdMode, true
}

// HasExecutingStateIdMode returns a boolean if a field has been set.
func (o *WorkflowConfig) HasExecutingStateIdMode() bool {
	if o != nil && !IsNil(o.ExecutingStateIdMode) {
		return true
	}

	return false
}

// SetExecutingStateIdMode gets a reference to the given ExecutingStateIdMode and assigns it to the ExecutingStateIdMode field.
func (o *WorkflowConfig) SetExecutingStateIdMode(v ExecutingStateIdMode) {
	o.ExecutingStateIdMode = &v
}

// GetContinueAsNewThreshold returns the ContinueAsNewThreshold field value if set, zero value otherwise.
func (o *WorkflowConfig) GetContinueAsNewThreshold() int32 {
	if o == nil || IsNil(o.ContinueAsNewThreshold) {
		var ret int32
		return ret
	}
	return *o.ContinueAsNewThreshold
}

// GetContinueAsNewThresholdOk returns a tuple with the ContinueAsNewThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowConfig) GetContinueAsNewThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.ContinueAsNewThreshold) {
		return nil, false
	}
	return o.ContinueAsNewThreshold, true
}

// HasContinueAsNewThreshold returns a boolean if a field has been set.
func (o *WorkflowConfig) HasContinueAsNewThreshold() bool {
	if o != nil && !IsNil(o.ContinueAsNewThreshold) {
		return true
	}

	return false
}

// SetContinueAsNewThreshold gets a reference to the given int32 and assigns it to the ContinueAsNewThreshold field.
func (o *WorkflowConfig) SetContinueAsNewThreshold(v int32) {
	o.ContinueAsNewThreshold = &v
}

// GetContinueAsNewPageSizeInBytes returns the ContinueAsNewPageSizeInBytes field value if set, zero value otherwise.
func (o *WorkflowConfig) GetContinueAsNewPageSizeInBytes() int32 {
	if o == nil || IsNil(o.ContinueAsNewPageSizeInBytes) {
		var ret int32
		return ret
	}
	return *o.ContinueAsNewPageSizeInBytes
}

// GetContinueAsNewPageSizeInBytesOk returns a tuple with the ContinueAsNewPageSizeInBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowConfig) GetContinueAsNewPageSizeInBytesOk() (*int32, bool) {
	if o == nil || IsNil(o.ContinueAsNewPageSizeInBytes) {
		return nil, false
	}
	return o.ContinueAsNewPageSizeInBytes, true
}

// HasContinueAsNewPageSizeInBytes returns a boolean if a field has been set.
func (o *WorkflowConfig) HasContinueAsNewPageSizeInBytes() bool {
	if o != nil && !IsNil(o.ContinueAsNewPageSizeInBytes) {
		return true
	}

	return false
}

// SetContinueAsNewPageSizeInBytes gets a reference to the given int32 and assigns it to the ContinueAsNewPageSizeInBytes field.
func (o *WorkflowConfig) SetContinueAsNewPageSizeInBytes(v int32) {
	o.ContinueAsNewPageSizeInBytes = &v
}

// GetOptimizeActivity returns the OptimizeActivity field value if set, zero value otherwise.
func (o *WorkflowConfig) GetOptimizeActivity() bool {
	if o == nil || IsNil(o.OptimizeActivity) {
		var ret bool
		return ret
	}
	return *o.OptimizeActivity
}

// GetOptimizeActivityOk returns a tuple with the OptimizeActivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowConfig) GetOptimizeActivityOk() (*bool, bool) {
	if o == nil || IsNil(o.OptimizeActivity) {
		return nil, false
	}
	return o.OptimizeActivity, true
}

// HasOptimizeActivity returns a boolean if a field has been set.
func (o *WorkflowConfig) HasOptimizeActivity() bool {
	if o != nil && !IsNil(o.OptimizeActivity) {
		return true
	}

	return false
}

// SetOptimizeActivity gets a reference to the given bool and assigns it to the OptimizeActivity field.
func (o *WorkflowConfig) SetOptimizeActivity(v bool) {
	o.OptimizeActivity = &v
}

// GetOptimizeTimer returns the OptimizeTimer field value if set, zero value otherwise.
func (o *WorkflowConfig) GetOptimizeTimer() bool {
	if o == nil || IsNil(o.OptimizeTimer) {
		var ret bool
		return ret
	}
	return *o.OptimizeTimer
}

// GetOptimizeTimerOk returns a tuple with the OptimizeTimer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowConfig) GetOptimizeTimerOk() (*bool, bool) {
	if o == nil || IsNil(o.OptimizeTimer) {
		return nil, false
	}
	return o.OptimizeTimer, true
}

// HasOptimizeTimer returns a boolean if a field has been set.
func (o *WorkflowConfig) HasOptimizeTimer() bool {
	if o != nil && !IsNil(o.OptimizeTimer) {
		return true
	}

	return false
}

// SetOptimizeTimer gets a reference to the given bool and assigns it to the OptimizeTimer field.
func (o *WorkflowConfig) SetOptimizeTimer(v bool) {
	o.OptimizeTimer = &v
}

func (o WorkflowConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisableSystemSearchAttribute) {
		toSerialize["disableSystemSearchAttribute"] = o.DisableSystemSearchAttribute
	}
	if !IsNil(o.ExecutingStateIdMode) {
		toSerialize["executingStateIdMode"] = o.ExecutingStateIdMode
	}
	if !IsNil(o.ContinueAsNewThreshold) {
		toSerialize["continueAsNewThreshold"] = o.ContinueAsNewThreshold
	}
	if !IsNil(o.ContinueAsNewPageSizeInBytes) {
		toSerialize["continueAsNewPageSizeInBytes"] = o.ContinueAsNewPageSizeInBytes
	}
	if !IsNil(o.OptimizeActivity) {
		toSerialize["optimizeActivity"] = o.OptimizeActivity
	}
	if !IsNil(o.OptimizeTimer) {
		toSerialize["optimizeTimer"] = o.OptimizeTimer
	}
	return toSerialize, nil
}

type NullableWorkflowConfig struct {
	value *WorkflowConfig
	isSet bool
}

func (v NullableWorkflowConfig) Get() *WorkflowConfig {
	return v.value
}

func (v *NullableWorkflowConfig) Set(val *WorkflowConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowConfig(val *WorkflowConfig) *NullableWorkflowConfig {
	return &NullableWorkflowConfig{value: val, isSet: true}
}

func (v NullableWorkflowConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
