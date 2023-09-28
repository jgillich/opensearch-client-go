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

// VersionType Specific version type.
type VersionType string

// List of VersionType
const (
	INTERNAL VersionType = "internal"
	EXTERNAL VersionType = "external"
	EXTERNAL_GTE VersionType = "external_gte"
	FORCE VersionType = "force"
)

// All allowed values of VersionType enum
var AllowedVersionTypeEnumValues = []VersionType{
	"internal",
	"external",
	"external_gte",
	"force",
}

func (v *VersionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VersionType(value)
	for _, existing := range AllowedVersionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VersionType", value)
}

// NewVersionTypeFromValue returns a pointer to a valid VersionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVersionTypeFromValue(v string) (*VersionType, error) {
	ev := VersionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VersionType: valid values are %v", v, AllowedVersionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VersionType) IsValid() bool {
	for _, existing := range AllowedVersionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VersionType value
func (v VersionType) Ptr() *VersionType {
	return &v
}

type NullableVersionType struct {
	value *VersionType
	isSet bool
}

func (v NullableVersionType) Get() *VersionType {
	return v.value
}

func (v *NullableVersionType) Set(val *VersionType) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionType) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionType(val *VersionType) *NullableVersionType {
	return &NullableVersionType{value: val, isSet: true}
}

func (v NullableVersionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

