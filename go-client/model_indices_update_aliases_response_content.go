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

// checks if the IndicesUpdateAliasesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndicesUpdateAliasesResponseContent{}

// IndicesUpdateAliasesResponseContent struct for IndicesUpdateAliasesResponseContent
type IndicesUpdateAliasesResponseContent struct {
	Acknowledged bool `json:"acknowledged"`
}

// NewIndicesUpdateAliasesResponseContent instantiates a new IndicesUpdateAliasesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndicesUpdateAliasesResponseContent(acknowledged bool) *IndicesUpdateAliasesResponseContent {
	this := IndicesUpdateAliasesResponseContent{}
	this.Acknowledged = acknowledged
	return &this
}

// NewIndicesUpdateAliasesResponseContentWithDefaults instantiates a new IndicesUpdateAliasesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndicesUpdateAliasesResponseContentWithDefaults() *IndicesUpdateAliasesResponseContent {
	this := IndicesUpdateAliasesResponseContent{}
	return &this
}

// GetAcknowledged returns the Acknowledged field value
func (o *IndicesUpdateAliasesResponseContent) GetAcknowledged() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Acknowledged
}

// GetAcknowledgedOk returns a tuple with the Acknowledged field value
// and a boolean to check if the value has been set.
func (o *IndicesUpdateAliasesResponseContent) GetAcknowledgedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Acknowledged, true
}

// SetAcknowledged sets field value
func (o *IndicesUpdateAliasesResponseContent) SetAcknowledged(v bool) {
	o.Acknowledged = v
}

func (o IndicesUpdateAliasesResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndicesUpdateAliasesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["acknowledged"] = o.Acknowledged
	return toSerialize, nil
}

type NullableIndicesUpdateAliasesResponseContent struct {
	value *IndicesUpdateAliasesResponseContent
	isSet bool
}

func (v NullableIndicesUpdateAliasesResponseContent) Get() *IndicesUpdateAliasesResponseContent {
	return v.value
}

func (v *NullableIndicesUpdateAliasesResponseContent) Set(val *IndicesUpdateAliasesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableIndicesUpdateAliasesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableIndicesUpdateAliasesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndicesUpdateAliasesResponseContent(val *IndicesUpdateAliasesResponseContent) *NullableIndicesUpdateAliasesResponseContent {
	return &NullableIndicesUpdateAliasesResponseContent{value: val, isSet: true}
}

func (v NullableIndicesUpdateAliasesResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndicesUpdateAliasesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


