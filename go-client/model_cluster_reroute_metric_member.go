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

// ClusterRerouteMetricMember the model 'ClusterRerouteMetricMember'
type ClusterRerouteMetricMember string

// List of ClusterRerouteMetric_Member
const (
	ALL ClusterRerouteMetricMember = "_all"
	BLOCKS ClusterRerouteMetricMember = "blocks"
	METADATA ClusterRerouteMetricMember = "metadata"
	NODES ClusterRerouteMetricMember = "nodes"
	ROUTING_TABLE ClusterRerouteMetricMember = "routing_table"
	MASTER_NODE ClusterRerouteMetricMember = "master_node"
	CLUSTER_MANAGER_NODE ClusterRerouteMetricMember = "cluster_manager_node"
	VERSION ClusterRerouteMetricMember = "version"
)

// All allowed values of ClusterRerouteMetricMember enum
var AllowedClusterRerouteMetricMemberEnumValues = []ClusterRerouteMetricMember{
	"_all",
	"blocks",
	"metadata",
	"nodes",
	"routing_table",
	"master_node",
	"cluster_manager_node",
	"version",
}

func (v *ClusterRerouteMetricMember) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ClusterRerouteMetricMember(value)
	for _, existing := range AllowedClusterRerouteMetricMemberEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ClusterRerouteMetricMember", value)
}

// NewClusterRerouteMetricMemberFromValue returns a pointer to a valid ClusterRerouteMetricMember
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClusterRerouteMetricMemberFromValue(v string) (*ClusterRerouteMetricMember, error) {
	ev := ClusterRerouteMetricMember(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClusterRerouteMetricMember: valid values are %v", v, AllowedClusterRerouteMetricMemberEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClusterRerouteMetricMember) IsValid() bool {
	for _, existing := range AllowedClusterRerouteMetricMemberEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ClusterRerouteMetric_Member value
func (v ClusterRerouteMetricMember) Ptr() *ClusterRerouteMetricMember {
	return &v
}

type NullableClusterRerouteMetricMember struct {
	value *ClusterRerouteMetricMember
	isSet bool
}

func (v NullableClusterRerouteMetricMember) Get() *ClusterRerouteMetricMember {
	return v.value
}

func (v *NullableClusterRerouteMetricMember) Set(val *ClusterRerouteMetricMember) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterRerouteMetricMember) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterRerouteMetricMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterRerouteMetricMember(val *ClusterRerouteMetricMember) *NullableClusterRerouteMetricMember {
	return &NullableClusterRerouteMetricMember{value: val, isSet: true}
}

func (v NullableClusterRerouteMetricMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterRerouteMetricMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

