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

// Relation the model 'Relation'
type Relation string

// List of Relation
const (
	EQ Relation = "eq"
	GTE Relation = "gte"
)

// All allowed values of Relation enum
var AllowedRelationEnumValues = []Relation{
	"eq",
	"gte",
}

func (v *Relation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Relation(value)
	for _, existing := range AllowedRelationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Relation", value)
}

// NewRelationFromValue returns a pointer to a valid Relation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRelationFromValue(v string) (*Relation, error) {
	ev := Relation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Relation: valid values are %v", v, AllowedRelationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Relation) IsValid() bool {
	for _, existing := range AllowedRelationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Relation value
func (v Relation) Ptr() *Relation {
	return &v
}

type NullableRelation struct {
	value *Relation
	isSet bool
}

func (v NullableRelation) Get() *Relation {
	return v.value
}

func (v *NullableRelation) Set(val *Relation) {
	v.value = val
	v.isSet = true
}

func (v NullableRelation) IsSet() bool {
	return v.isSet
}

func (v *NullableRelation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelation(val *Relation) *NullableRelation {
	return &NullableRelation{value: val, isSet: true}
}

func (v NullableRelation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

