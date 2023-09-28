/*
OpenSearch

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2021-11-23
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// WaitForEvents Wait until all currently queued events with the given priority are processed.
type WaitForEvents string

// List of WaitForEvents
const (
	IMMEDIATE WaitForEvents = "immediate"
	URGENT WaitForEvents = "urgent"
	HIGH WaitForEvents = "high"
	NORMAL WaitForEvents = "normal"
	LOW WaitForEvents = "low"
	LANGUID WaitForEvents = "languid"
)

// All allowed values of WaitForEvents enum
var AllowedWaitForEventsEnumValues = []WaitForEvents{
	"immediate",
	"urgent",
	"high",
	"normal",
	"low",
	"languid",
}

func (v *WaitForEvents) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WaitForEvents(value)
	for _, existing := range AllowedWaitForEventsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WaitForEvents", value)
}

// NewWaitForEventsFromValue returns a pointer to a valid WaitForEvents
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWaitForEventsFromValue(v string) (*WaitForEvents, error) {
	ev := WaitForEvents(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WaitForEvents: valid values are %v", v, AllowedWaitForEventsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WaitForEvents) IsValid() bool {
	for _, existing := range AllowedWaitForEventsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WaitForEvents value
func (v WaitForEvents) Ptr() *WaitForEvents {
	return &v
}

type NullableWaitForEvents struct {
	value *WaitForEvents
	isSet bool
}

func (v NullableWaitForEvents) Get() *WaitForEvents {
	return v.value
}

func (v *NullableWaitForEvents) Set(val *WaitForEvents) {
	v.value = val
	v.isSet = true
}

func (v NullableWaitForEvents) IsSet() bool {
	return v.isSet
}

func (v *NullableWaitForEvents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWaitForEvents(val *WaitForEvents) *NullableWaitForEvents {
	return &NullableWaitForEvents{value: val, isSet: true}
}

func (v NullableWaitForEvents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWaitForEvents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

