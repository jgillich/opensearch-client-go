/*
OpenSearch

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2021-11-23
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package opensearch

import (
	"encoding/json"
	"fmt"
)

// GroupBy Group tasks by nodes or parent/child relationships.
type GroupBy string

// List of GroupBy
const (
	NODES GroupBy = "nodes"
	PARENTS GroupBy = "parents"
	NONE GroupBy = "none"
)

// All allowed values of GroupBy enum
var AllowedGroupByEnumValues = []GroupBy{
	"nodes",
	"parents",
	"none",
}

func (v *GroupBy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GroupBy(value)
	for _, existing := range AllowedGroupByEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GroupBy", value)
}

// NewGroupByFromValue returns a pointer to a valid GroupBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGroupByFromValue(v string) (*GroupBy, error) {
	ev := GroupBy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GroupBy: valid values are %v", v, AllowedGroupByEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GroupBy) IsValid() bool {
	for _, existing := range AllowedGroupByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GroupBy value
func (v GroupBy) Ptr() *GroupBy {
	return &v
}

type NullableGroupBy struct {
	value *GroupBy
	isSet bool
}

func (v NullableGroupBy) Get() *GroupBy {
	return v.value
}

func (v *NullableGroupBy) Set(val *GroupBy) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupBy) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupBy(val *GroupBy) *NullableGroupBy {
	return &NullableGroupBy{value: val, isSet: true}
}

func (v NullableGroupBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

