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

// ExpandWildcards Whether to expand wildcard expression to concrete indices that are open, closed or both.
type ExpandWildcards string

// List of ExpandWildcards
const (
	ALL ExpandWildcards = "all"
	OPEN ExpandWildcards = "open"
	CLOSED ExpandWildcards = "closed"
	HIDDEN ExpandWildcards = "hidden"
	NONE ExpandWildcards = "none"
)

// All allowed values of ExpandWildcards enum
var AllowedExpandWildcardsEnumValues = []ExpandWildcards{
	"all",
	"open",
	"closed",
	"hidden",
	"none",
}

func (v *ExpandWildcards) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExpandWildcards(value)
	for _, existing := range AllowedExpandWildcardsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExpandWildcards", value)
}

// NewExpandWildcardsFromValue returns a pointer to a valid ExpandWildcards
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExpandWildcardsFromValue(v string) (*ExpandWildcards, error) {
	ev := ExpandWildcards(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExpandWildcards: valid values are %v", v, AllowedExpandWildcardsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExpandWildcards) IsValid() bool {
	for _, existing := range AllowedExpandWildcardsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExpandWildcards value
func (v ExpandWildcards) Ptr() *ExpandWildcards {
	return &v
}

type NullableExpandWildcards struct {
	value *ExpandWildcards
	isSet bool
}

func (v NullableExpandWildcards) Get() *ExpandWildcards {
	return v.value
}

func (v *NullableExpandWildcards) Set(val *ExpandWildcards) {
	v.value = val
	v.isSet = true
}

func (v NullableExpandWildcards) IsSet() bool {
	return v.isSet
}

func (v *NullableExpandWildcards) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpandWildcards(val *ExpandWildcards) *NullableExpandWildcards {
	return &NullableExpandWildcards{value: val, isSet: true}
}

func (v NullableExpandWildcards) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpandWildcards) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

