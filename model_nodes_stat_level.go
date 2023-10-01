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

// NodesStatLevel Return indices stats aggregated at index, node or shard level.
type NodesStatLevel string

// List of NodesStatLevel
const (
	NODESSTATLEVEL_INDICES NodesStatLevel = "indices"
	NODESSTATLEVEL_NODE NodesStatLevel = "node"
	NODESSTATLEVEL_SHARDS NodesStatLevel = "shards"
)

// All allowed values of NodesStatLevel enum
var AllowedNodesStatLevelEnumValues = []NodesStatLevel{
	"indices",
	"node",
	"shards",
}

func (v *NodesStatLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NodesStatLevel(value)
	for _, existing := range AllowedNodesStatLevelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NodesStatLevel", value)
}

// NewNodesStatLevelFromValue returns a pointer to a valid NodesStatLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNodesStatLevelFromValue(v string) (*NodesStatLevel, error) {
	ev := NodesStatLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NodesStatLevel: valid values are %v", v, AllowedNodesStatLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NodesStatLevel) IsValid() bool {
	for _, existing := range AllowedNodesStatLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NodesStatLevel value
func (v NodesStatLevel) Ptr() *NodesStatLevel {
	return &v
}

type NullableNodesStatLevel struct {
	value *NodesStatLevel
	isSet bool
}

func (v NullableNodesStatLevel) Get() *NodesStatLevel {
	return v.value
}

func (v *NullableNodesStatLevel) Set(val *NodesStatLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableNodesStatLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableNodesStatLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodesStatLevel(val *NodesStatLevel) *NullableNodesStatLevel {
	return &NullableNodesStatLevel{value: val, isSet: true}
}

func (v NullableNodesStatLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodesStatLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

