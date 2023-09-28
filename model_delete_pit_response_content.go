/*
OpenSearch

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2021-11-23
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DeletePitResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeletePitResponseContent{}

// DeletePitResponseContent struct for DeletePitResponseContent
type DeletePitResponseContent struct {
	Pits []DeletedPit `json:"pits,omitempty"`
}

// NewDeletePitResponseContent instantiates a new DeletePitResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeletePitResponseContent() *DeletePitResponseContent {
	this := DeletePitResponseContent{}
	return &this
}

// NewDeletePitResponseContentWithDefaults instantiates a new DeletePitResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeletePitResponseContentWithDefaults() *DeletePitResponseContent {
	this := DeletePitResponseContent{}
	return &this
}

// GetPits returns the Pits field value if set, zero value otherwise.
func (o *DeletePitResponseContent) GetPits() []DeletedPit {
	if o == nil || IsNil(o.Pits) {
		var ret []DeletedPit
		return ret
	}
	return o.Pits
}

// GetPitsOk returns a tuple with the Pits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeletePitResponseContent) GetPitsOk() ([]DeletedPit, bool) {
	if o == nil || IsNil(o.Pits) {
		return nil, false
	}
	return o.Pits, true
}

// HasPits returns a boolean if a field has been set.
func (o *DeletePitResponseContent) HasPits() bool {
	if o != nil && !IsNil(o.Pits) {
		return true
	}

	return false
}

// SetPits gets a reference to the given []DeletedPit and assigns it to the Pits field.
func (o *DeletePitResponseContent) SetPits(v []DeletedPit) {
	o.Pits = v
}

func (o DeletePitResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeletePitResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pits) {
		toSerialize["pits"] = o.Pits
	}
	return toSerialize, nil
}

type NullableDeletePitResponseContent struct {
	value *DeletePitResponseContent
	isSet bool
}

func (v NullableDeletePitResponseContent) Get() *DeletePitResponseContent {
	return v.value
}

func (v *NullableDeletePitResponseContent) Set(val *DeletePitResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableDeletePitResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableDeletePitResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeletePitResponseContent(val *DeletePitResponseContent) *NullableDeletePitResponseContent {
	return &NullableDeletePitResponseContent{value: val, isSet: true}
}

func (v NullableDeletePitResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeletePitResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


