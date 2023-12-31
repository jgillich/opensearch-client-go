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

// Bytes The unit in which to display byte values.
type Bytes string

// List of Bytes
const (
	BYTES_B Bytes = "b"
	BYTES_K Bytes = "k"
	BYTES_KB Bytes = "kb"
	BYTES_M Bytes = "m"
	BYTES_MB Bytes = "mb"
	BYTES_G Bytes = "g"
	BYTES_GB Bytes = "gb"
	BYTES_T Bytes = "t"
	BYTES_TB Bytes = "tb"
	BYTES_P Bytes = "p"
	BYTES_PB Bytes = "pb"
)

// All allowed values of Bytes enum
var AllowedBytesEnumValues = []Bytes{
	"b",
	"k",
	"kb",
	"m",
	"mb",
	"g",
	"gb",
	"t",
	"tb",
	"p",
	"pb",
}

func (v *Bytes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Bytes(value)
	for _, existing := range AllowedBytesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Bytes", value)
}

// NewBytesFromValue returns a pointer to a valid Bytes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBytesFromValue(v string) (*Bytes, error) {
	ev := Bytes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Bytes: valid values are %v", v, AllowedBytesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Bytes) IsValid() bool {
	for _, existing := range AllowedBytesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Bytes value
func (v Bytes) Ptr() *Bytes {
	return &v
}

type NullableBytes struct {
	value *Bytes
	isSet bool
}

func (v NullableBytes) Get() *Bytes {
	return v.value
}

func (v *NullableBytes) Set(val *Bytes) {
	v.value = val
	v.isSet = true
}

func (v NullableBytes) IsSet() bool {
	return v.isSet
}

func (v *NullableBytes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBytes(val *Bytes) *NullableBytes {
	return &NullableBytes{value: val, isSet: true}
}

func (v NullableBytes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBytes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

