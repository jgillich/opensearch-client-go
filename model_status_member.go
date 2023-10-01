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

// StatusMember the model 'StatusMember'
type StatusMember string

// List of Status_Member
const (
	STATUSMEMBER_GREEN StatusMember = "green"
	STATUSMEMBER_YELLOW StatusMember = "yellow"
	STATUSMEMBER_RED StatusMember = "red"
	STATUSMEMBER_ALL StatusMember = "all"
)

// All allowed values of StatusMember enum
var AllowedStatusMemberEnumValues = []StatusMember{
	"green",
	"yellow",
	"red",
	"all",
}

func (v *StatusMember) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StatusMember(value)
	for _, existing := range AllowedStatusMemberEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StatusMember", value)
}

// NewStatusMemberFromValue returns a pointer to a valid StatusMember
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStatusMemberFromValue(v string) (*StatusMember, error) {
	ev := StatusMember(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StatusMember: valid values are %v", v, AllowedStatusMemberEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StatusMember) IsValid() bool {
	for _, existing := range AllowedStatusMemberEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Status_Member value
func (v StatusMember) Ptr() *StatusMember {
	return &v
}

type NullableStatusMember struct {
	value *StatusMember
	isSet bool
}

func (v NullableStatusMember) Get() *StatusMember {
	return v.value
}

func (v *NullableStatusMember) Set(val *StatusMember) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusMember) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusMember(val *StatusMember) *NullableStatusMember {
	return &NullableStatusMember{value: val, isSet: true}
}

func (v NullableStatusMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

