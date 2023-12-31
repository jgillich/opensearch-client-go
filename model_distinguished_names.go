/*
OpenSearch

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2021-11-23
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package opensearch

import (
	"encoding/json"
)

// checks if the DistinguishedNames type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DistinguishedNames{}

// DistinguishedNames struct for DistinguishedNames
type DistinguishedNames struct {
	NodesDn []string `json:"nodes_dn,omitempty"`
}

// NewDistinguishedNames instantiates a new DistinguishedNames object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDistinguishedNames() *DistinguishedNames {
	this := DistinguishedNames{}
	return &this
}

// NewDistinguishedNamesWithDefaults instantiates a new DistinguishedNames object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDistinguishedNamesWithDefaults() *DistinguishedNames {
	this := DistinguishedNames{}
	return &this
}

// GetNodesDn returns the NodesDn field value if set, zero value otherwise.
func (o *DistinguishedNames) GetNodesDn() []string {
	if o == nil || IsNil(o.NodesDn) {
		var ret []string
		return ret
	}
	return o.NodesDn
}

// GetNodesDnOk returns a tuple with the NodesDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistinguishedNames) GetNodesDnOk() ([]string, bool) {
	if o == nil || IsNil(o.NodesDn) {
		return nil, false
	}
	return o.NodesDn, true
}

// HasNodesDn returns a boolean if a field has been set.
func (o *DistinguishedNames) HasNodesDn() bool {
	if o != nil && !IsNil(o.NodesDn) {
		return true
	}

	return false
}

// SetNodesDn gets a reference to the given []string and assigns it to the NodesDn field.
func (o *DistinguishedNames) SetNodesDn(v []string) {
	o.NodesDn = v
}

func (o DistinguishedNames) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DistinguishedNames) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NodesDn) {
		toSerialize["nodes_dn"] = o.NodesDn
	}
	return toSerialize, nil
}

type NullableDistinguishedNames struct {
	value *DistinguishedNames
	isSet bool
}

func (v NullableDistinguishedNames) Get() *DistinguishedNames {
	return v.value
}

func (v *NullableDistinguishedNames) Set(val *DistinguishedNames) {
	v.value = val
	v.isSet = true
}

func (v NullableDistinguishedNames) IsSet() bool {
	return v.isSet
}

func (v *NullableDistinguishedNames) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistinguishedNames(val *DistinguishedNames) *NullableDistinguishedNames {
	return &NullableDistinguishedNames{value: val, isSet: true}
}

func (v NullableDistinguishedNames) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistinguishedNames) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


