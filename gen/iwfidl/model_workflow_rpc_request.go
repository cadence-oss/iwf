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

// checks if the WorkflowRpcRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowRpcRequest{}

// WorkflowRpcRequest struct for WorkflowRpcRequest
type WorkflowRpcRequest struct {
	WorkflowId                        string                    `json:"workflowId"`
	WorkflowRunId                     *string                   `json:"workflowRunId,omitempty"`
	RpcName                           string                    `json:"rpcName"`
	Input                             *EncodedObject            `json:"input,omitempty"`
	SearchAttributesLoadingPolicy     *PersistenceLoadingPolicy `json:"searchAttributesLoadingPolicy,omitempty"`
	DataAttributesLoadingPolicy       *PersistenceLoadingPolicy `json:"dataAttributesLoadingPolicy,omitempty"`
	CachedDataAttributesLoadingPolicy *PersistenceLoadingPolicy `json:"cachedDataAttributesLoadingPolicy,omitempty"`
	TimeoutSeconds                    *int32                    `json:"timeoutSeconds,omitempty"`
}

// NewWorkflowRpcRequest instantiates a new WorkflowRpcRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowRpcRequest(workflowId string, rpcName string) *WorkflowRpcRequest {
	this := WorkflowRpcRequest{}
	this.WorkflowId = workflowId
	this.RpcName = rpcName
	return &this
}

// NewWorkflowRpcRequestWithDefaults instantiates a new WorkflowRpcRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowRpcRequestWithDefaults() *WorkflowRpcRequest {
	this := WorkflowRpcRequest{}
	return &this
}

// GetWorkflowId returns the WorkflowId field value
func (o *WorkflowRpcRequest) GetWorkflowId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkflowId
}

// GetWorkflowIdOk returns a tuple with the WorkflowId field value
// and a boolean to check if the value has been set.
func (o *WorkflowRpcRequest) GetWorkflowIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkflowId, true
}

// SetWorkflowId sets field value
func (o *WorkflowRpcRequest) SetWorkflowId(v string) {
	o.WorkflowId = v
}

// GetWorkflowRunId returns the WorkflowRunId field value if set, zero value otherwise.
func (o *WorkflowRpcRequest) GetWorkflowRunId() string {
	if o == nil || IsNil(o.WorkflowRunId) {
		var ret string
		return ret
	}
	return *o.WorkflowRunId
}

// GetWorkflowRunIdOk returns a tuple with the WorkflowRunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRpcRequest) GetWorkflowRunIdOk() (*string, bool) {
	if o == nil || IsNil(o.WorkflowRunId) {
		return nil, false
	}
	return o.WorkflowRunId, true
}

// HasWorkflowRunId returns a boolean if a field has been set.
func (o *WorkflowRpcRequest) HasWorkflowRunId() bool {
	if o != nil && !IsNil(o.WorkflowRunId) {
		return true
	}

	return false
}

// SetWorkflowRunId gets a reference to the given string and assigns it to the WorkflowRunId field.
func (o *WorkflowRpcRequest) SetWorkflowRunId(v string) {
	o.WorkflowRunId = &v
}

// GetRpcName returns the RpcName field value
func (o *WorkflowRpcRequest) GetRpcName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RpcName
}

// GetRpcNameOk returns a tuple with the RpcName field value
// and a boolean to check if the value has been set.
func (o *WorkflowRpcRequest) GetRpcNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RpcName, true
}

// SetRpcName sets field value
func (o *WorkflowRpcRequest) SetRpcName(v string) {
	o.RpcName = v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *WorkflowRpcRequest) GetInput() EncodedObject {
	if o == nil || IsNil(o.Input) {
		var ret EncodedObject
		return ret
	}
	return *o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRpcRequest) GetInputOk() (*EncodedObject, bool) {
	if o == nil || IsNil(o.Input) {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *WorkflowRpcRequest) HasInput() bool {
	if o != nil && !IsNil(o.Input) {
		return true
	}

	return false
}

// SetInput gets a reference to the given EncodedObject and assigns it to the Input field.
func (o *WorkflowRpcRequest) SetInput(v EncodedObject) {
	o.Input = &v
}

// GetSearchAttributesLoadingPolicy returns the SearchAttributesLoadingPolicy field value if set, zero value otherwise.
func (o *WorkflowRpcRequest) GetSearchAttributesLoadingPolicy() PersistenceLoadingPolicy {
	if o == nil || IsNil(o.SearchAttributesLoadingPolicy) {
		var ret PersistenceLoadingPolicy
		return ret
	}
	return *o.SearchAttributesLoadingPolicy
}

// GetSearchAttributesLoadingPolicyOk returns a tuple with the SearchAttributesLoadingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRpcRequest) GetSearchAttributesLoadingPolicyOk() (*PersistenceLoadingPolicy, bool) {
	if o == nil || IsNil(o.SearchAttributesLoadingPolicy) {
		return nil, false
	}
	return o.SearchAttributesLoadingPolicy, true
}

// HasSearchAttributesLoadingPolicy returns a boolean if a field has been set.
func (o *WorkflowRpcRequest) HasSearchAttributesLoadingPolicy() bool {
	if o != nil && !IsNil(o.SearchAttributesLoadingPolicy) {
		return true
	}

	return false
}

// SetSearchAttributesLoadingPolicy gets a reference to the given PersistenceLoadingPolicy and assigns it to the SearchAttributesLoadingPolicy field.
func (o *WorkflowRpcRequest) SetSearchAttributesLoadingPolicy(v PersistenceLoadingPolicy) {
	o.SearchAttributesLoadingPolicy = &v
}

// GetDataAttributesLoadingPolicy returns the DataAttributesLoadingPolicy field value if set, zero value otherwise.
func (o *WorkflowRpcRequest) GetDataAttributesLoadingPolicy() PersistenceLoadingPolicy {
	if o == nil || IsNil(o.DataAttributesLoadingPolicy) {
		var ret PersistenceLoadingPolicy
		return ret
	}
	return *o.DataAttributesLoadingPolicy
}

// GetDataAttributesLoadingPolicyOk returns a tuple with the DataAttributesLoadingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRpcRequest) GetDataAttributesLoadingPolicyOk() (*PersistenceLoadingPolicy, bool) {
	if o == nil || IsNil(o.DataAttributesLoadingPolicy) {
		return nil, false
	}
	return o.DataAttributesLoadingPolicy, true
}

// HasDataAttributesLoadingPolicy returns a boolean if a field has been set.
func (o *WorkflowRpcRequest) HasDataAttributesLoadingPolicy() bool {
	if o != nil && !IsNil(o.DataAttributesLoadingPolicy) {
		return true
	}

	return false
}

// SetDataAttributesLoadingPolicy gets a reference to the given PersistenceLoadingPolicy and assigns it to the DataAttributesLoadingPolicy field.
func (o *WorkflowRpcRequest) SetDataAttributesLoadingPolicy(v PersistenceLoadingPolicy) {
	o.DataAttributesLoadingPolicy = &v
}

// GetCachedDataAttributesLoadingPolicy returns the CachedDataAttributesLoadingPolicy field value if set, zero value otherwise.
func (o *WorkflowRpcRequest) GetCachedDataAttributesLoadingPolicy() PersistenceLoadingPolicy {
	if o == nil || IsNil(o.CachedDataAttributesLoadingPolicy) {
		var ret PersistenceLoadingPolicy
		return ret
	}
	return *o.CachedDataAttributesLoadingPolicy
}

// GetCachedDataAttributesLoadingPolicyOk returns a tuple with the CachedDataAttributesLoadingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRpcRequest) GetCachedDataAttributesLoadingPolicyOk() (*PersistenceLoadingPolicy, bool) {
	if o == nil || IsNil(o.CachedDataAttributesLoadingPolicy) {
		return nil, false
	}
	return o.CachedDataAttributesLoadingPolicy, true
}

// HasCachedDataAttributesLoadingPolicy returns a boolean if a field has been set.
func (o *WorkflowRpcRequest) HasCachedDataAttributesLoadingPolicy() bool {
	if o != nil && !IsNil(o.CachedDataAttributesLoadingPolicy) {
		return true
	}

	return false
}

// SetCachedDataAttributesLoadingPolicy gets a reference to the given PersistenceLoadingPolicy and assigns it to the CachedDataAttributesLoadingPolicy field.
func (o *WorkflowRpcRequest) SetCachedDataAttributesLoadingPolicy(v PersistenceLoadingPolicy) {
	o.CachedDataAttributesLoadingPolicy = &v
}

// GetTimeoutSeconds returns the TimeoutSeconds field value if set, zero value otherwise.
func (o *WorkflowRpcRequest) GetTimeoutSeconds() int32 {
	if o == nil || IsNil(o.TimeoutSeconds) {
		var ret int32
		return ret
	}
	return *o.TimeoutSeconds
}

// GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRpcRequest) GetTimeoutSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.TimeoutSeconds) {
		return nil, false
	}
	return o.TimeoutSeconds, true
}

// HasTimeoutSeconds returns a boolean if a field has been set.
func (o *WorkflowRpcRequest) HasTimeoutSeconds() bool {
	if o != nil && !IsNil(o.TimeoutSeconds) {
		return true
	}

	return false
}

// SetTimeoutSeconds gets a reference to the given int32 and assigns it to the TimeoutSeconds field.
func (o *WorkflowRpcRequest) SetTimeoutSeconds(v int32) {
	o.TimeoutSeconds = &v
}

func (o WorkflowRpcRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowRpcRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["workflowId"] = o.WorkflowId
	if !IsNil(o.WorkflowRunId) {
		toSerialize["workflowRunId"] = o.WorkflowRunId
	}
	toSerialize["rpcName"] = o.RpcName
	if !IsNil(o.Input) {
		toSerialize["input"] = o.Input
	}
	if !IsNil(o.SearchAttributesLoadingPolicy) {
		toSerialize["searchAttributesLoadingPolicy"] = o.SearchAttributesLoadingPolicy
	}
	if !IsNil(o.DataAttributesLoadingPolicy) {
		toSerialize["dataAttributesLoadingPolicy"] = o.DataAttributesLoadingPolicy
	}
	if !IsNil(o.CachedDataAttributesLoadingPolicy) {
		toSerialize["cachedDataAttributesLoadingPolicy"] = o.CachedDataAttributesLoadingPolicy
	}
	if !IsNil(o.TimeoutSeconds) {
		toSerialize["timeoutSeconds"] = o.TimeoutSeconds
	}
	return toSerialize, nil
}

type NullableWorkflowRpcRequest struct {
	value *WorkflowRpcRequest
	isSet bool
}

func (v NullableWorkflowRpcRequest) Get() *WorkflowRpcRequest {
	return v.value
}

func (v *NullableWorkflowRpcRequest) Set(val *WorkflowRpcRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowRpcRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowRpcRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowRpcRequest(val *WorkflowRpcRequest) *NullableWorkflowRpcRequest {
	return &NullableWorkflowRpcRequest{value: val, isSet: true}
}

func (v NullableWorkflowRpcRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowRpcRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
