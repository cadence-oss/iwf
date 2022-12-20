/*
Workflow APIs

This APIs for iwf SDKs to operate workflows

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iwfidl

import (
	"encoding/json"
	"fmt"
)

// SignalRequestStatus the model 'SignalRequestStatus'
type SignalRequestStatus string

// List of SignalRequestStatus
const (
	WAITING SignalRequestStatus = "WAITING"
	RECEIVED SignalRequestStatus = "RECEIVED"
)

// All allowed values of SignalRequestStatus enum
var AllowedSignalRequestStatusEnumValues = []SignalRequestStatus{
	"WAITING",
	"RECEIVED",
}

func (v *SignalRequestStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SignalRequestStatus(value)
	for _, existing := range AllowedSignalRequestStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SignalRequestStatus", value)
}

// NewSignalRequestStatusFromValue returns a pointer to a valid SignalRequestStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSignalRequestStatusFromValue(v string) (*SignalRequestStatus, error) {
	ev := SignalRequestStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SignalRequestStatus: valid values are %v", v, AllowedSignalRequestStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SignalRequestStatus) IsValid() bool {
	for _, existing := range AllowedSignalRequestStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SignalRequestStatus value
func (v SignalRequestStatus) Ptr() *SignalRequestStatus {
	return &v
}

type NullableSignalRequestStatus struct {
	value *SignalRequestStatus
	isSet bool
}

func (v NullableSignalRequestStatus) Get() *SignalRequestStatus {
	return v.value
}

func (v *NullableSignalRequestStatus) Set(val *SignalRequestStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSignalRequestStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSignalRequestStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignalRequestStatus(val *SignalRequestStatus) *NullableSignalRequestStatus {
	return &NullableSignalRequestStatus{value: val, isSet: true}
}

func (v NullableSignalRequestStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignalRequestStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

