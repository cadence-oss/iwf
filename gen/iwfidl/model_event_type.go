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

// EventType the model 'EventType'
type EventType string

// List of EventType
const (
	STATE_EXECUTE_ATTEMPT_FAIL_EVENT    EventType = "STATE_EXECUTE_ATTEMPT_FAIL_EVENT"
	STATE_EXECUTE_ATTEMPT_SUCC_EVENT    EventType = "STATE_EXECUTE_ATTEMPT_SUCC_EVENT"
	STATE_EXECUTE_EE_COMPLETE_EVENT     EventType = "STATE_EXECUTE_EE_COMPLETE_EVENT"
	STATE_EXECUTE_EE_FAIL_EVENT         EventType = "STATE_EXECUTE_EE_FAIL_EVENT"
	STATE_EXECUTE_EE_START_EVENT        EventType = "STATE_EXECUTE_EE_START_EVENT"
	STATE_WAIT_UNTIL_ATTEMPT_FAIL_EVENT EventType = "STATE_WAIT_UNTIL_ATTEMPT_FAIL_EVENT"
	STATE_WAIT_UNTIL_ATTEMPT_SUCC_EVENT EventType = "STATE_WAIT_UNTIL_ATTEMPT_SUCC_EVENT"
	STATE_WAIT_UNTIL_EE_COMPLETE_EVENT  EventType = "STATE_WAIT_UNTIL_EE_COMPLETE_EVENT"
	STATE_WAIT_UNTIL_EE_FAIL_EVENT      EventType = "STATE_WAIT_UNTIL_EE_FAIL_EVENT"
	STATE_WAIT_UNTIL_EE_START_EVENT     EventType = "STATE_WAIT_UNTIL_EE_START_EVENT"
	WORKFLOW_COMPLETE_EVENT             EventType = "WORKFLOW_COMPLETE_EVENT"
	WORKFLOW_FAIL_EVENT                 EventType = "WORKFLOW_FAIL_EVENT"
	WORKFLOW_START_EVENT                EventType = "WORKFLOW_START_EVENT"
)

// All allowed values of EventType enum
var AllowedEventTypeEnumValues = []EventType{
	"STATE_EXECUTE_ATTEMPT_FAIL_EVENT",
	"STATE_EXECUTE_ATTEMPT_SUCC_EVENT",
	"STATE_EXECUTE_EE_COMPLETE_EVENT",
	"STATE_EXECUTE_EE_FAIL_EVENT",
	"STATE_EXECUTE_EE_START_EVENT",
	"STATE_WAIT_UNTIL_ATTEMPT_FAIL_EVENT",
	"STATE_WAIT_UNTIL_ATTEMPT_SUCC_EVENT",
	"STATE_WAIT_UNTIL_EE_COMPLETE_EVENT",
	"STATE_WAIT_UNTIL_EE_FAIL_EVENT",
	"STATE_WAIT_UNTIL_EE_START_EVENT",
	"WORKFLOW_COMPLETE_EVENT",
	"WORKFLOW_FAIL_EVENT",
	"WORKFLOW_START_EVENT",
}

func (v *EventType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventType(value)
	for _, existing := range AllowedEventTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventType", value)
}

// NewEventTypeFromValue returns a pointer to a valid EventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventTypeFromValue(v string) (*EventType, error) {
	ev := EventType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventType: valid values are %v", v, AllowedEventTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventType) IsValid() bool {
	for _, existing := range AllowedEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventType value
func (v EventType) Ptr() *EventType {
	return &v
}

type NullableEventType struct {
	value *EventType
	isSet bool
}

func (v NullableEventType) Get() *EventType {
	return v.value
}

func (v *NullableEventType) Set(val *EventType) {
	v.value = val
	v.isSet = true
}

func (v NullableEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventType(val *EventType) *NullableEventType {
	return &NullableEventType{value: val, isSet: true}
}

func (v NullableEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}