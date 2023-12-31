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

// checks if the IndicesDeleteResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndicesDeleteResponseContent{}

// IndicesDeleteResponseContent struct for IndicesDeleteResponseContent
type IndicesDeleteResponseContent struct {
	Acknowledged *bool `json:"acknowledged,omitempty"`
}

// NewIndicesDeleteResponseContent instantiates a new IndicesDeleteResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndicesDeleteResponseContent() *IndicesDeleteResponseContent {
	this := IndicesDeleteResponseContent{}
	return &this
}

// NewIndicesDeleteResponseContentWithDefaults instantiates a new IndicesDeleteResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndicesDeleteResponseContentWithDefaults() *IndicesDeleteResponseContent {
	this := IndicesDeleteResponseContent{}
	return &this
}

// GetAcknowledged returns the Acknowledged field value if set, zero value otherwise.
func (o *IndicesDeleteResponseContent) GetAcknowledged() bool {
	if o == nil || IsNil(o.Acknowledged) {
		var ret bool
		return ret
	}
	return *o.Acknowledged
}

// GetAcknowledgedOk returns a tuple with the Acknowledged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicesDeleteResponseContent) GetAcknowledgedOk() (*bool, bool) {
	if o == nil || IsNil(o.Acknowledged) {
		return nil, false
	}
	return o.Acknowledged, true
}

// HasAcknowledged returns a boolean if a field has been set.
func (o *IndicesDeleteResponseContent) HasAcknowledged() bool {
	if o != nil && !IsNil(o.Acknowledged) {
		return true
	}

	return false
}

// SetAcknowledged gets a reference to the given bool and assigns it to the Acknowledged field.
func (o *IndicesDeleteResponseContent) SetAcknowledged(v bool) {
	o.Acknowledged = &v
}

func (o IndicesDeleteResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndicesDeleteResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Acknowledged) {
		toSerialize["acknowledged"] = o.Acknowledged
	}
	return toSerialize, nil
}

type NullableIndicesDeleteResponseContent struct {
	value *IndicesDeleteResponseContent
	isSet bool
}

func (v NullableIndicesDeleteResponseContent) Get() *IndicesDeleteResponseContent {
	return v.value
}

func (v *NullableIndicesDeleteResponseContent) Set(val *IndicesDeleteResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableIndicesDeleteResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableIndicesDeleteResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndicesDeleteResponseContent(val *IndicesDeleteResponseContent) *NullableIndicesDeleteResponseContent {
	return &NullableIndicesDeleteResponseContent{value: val, isSet: true}
}

func (v NullableIndicesDeleteResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndicesDeleteResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


