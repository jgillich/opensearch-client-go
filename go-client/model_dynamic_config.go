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

// checks if the DynamicConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DynamicConfig{}

// DynamicConfig struct for DynamicConfig
type DynamicConfig struct {
	Dynamic *DynamicOptions `json:"dynamic,omitempty"`
}

// NewDynamicConfig instantiates a new DynamicConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicConfig() *DynamicConfig {
	this := DynamicConfig{}
	return &this
}

// NewDynamicConfigWithDefaults instantiates a new DynamicConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicConfigWithDefaults() *DynamicConfig {
	this := DynamicConfig{}
	return &this
}

// GetDynamic returns the Dynamic field value if set, zero value otherwise.
func (o *DynamicConfig) GetDynamic() DynamicOptions {
	if o == nil || IsNil(o.Dynamic) {
		var ret DynamicOptions
		return ret
	}
	return *o.Dynamic
}

// GetDynamicOk returns a tuple with the Dynamic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicConfig) GetDynamicOk() (*DynamicOptions, bool) {
	if o == nil || IsNil(o.Dynamic) {
		return nil, false
	}
	return o.Dynamic, true
}

// HasDynamic returns a boolean if a field has been set.
func (o *DynamicConfig) HasDynamic() bool {
	if o != nil && !IsNil(o.Dynamic) {
		return true
	}

	return false
}

// SetDynamic gets a reference to the given DynamicOptions and assigns it to the Dynamic field.
func (o *DynamicConfig) SetDynamic(v DynamicOptions) {
	o.Dynamic = &v
}

func (o DynamicConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DynamicConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Dynamic) {
		toSerialize["dynamic"] = o.Dynamic
	}
	return toSerialize, nil
}

type NullableDynamicConfig struct {
	value *DynamicConfig
	isSet bool
}

func (v NullableDynamicConfig) Get() *DynamicConfig {
	return v.value
}

func (v *NullableDynamicConfig) Set(val *DynamicConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicConfig(val *DynamicConfig) *NullableDynamicConfig {
	return &NullableDynamicConfig{value: val, isSet: true}
}

func (v NullableDynamicConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


