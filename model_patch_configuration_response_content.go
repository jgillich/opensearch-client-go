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

// checks if the PatchConfigurationResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchConfigurationResponseContent{}

// PatchConfigurationResponseContent struct for PatchConfigurationResponseContent
type PatchConfigurationResponseContent struct {
	// Security Operation Status
	Status *string `json:"status,omitempty"`
	// Security Operation Message
	Message *string `json:"message,omitempty"`
}

// NewPatchConfigurationResponseContent instantiates a new PatchConfigurationResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchConfigurationResponseContent() *PatchConfigurationResponseContent {
	this := PatchConfigurationResponseContent{}
	return &this
}

// NewPatchConfigurationResponseContentWithDefaults instantiates a new PatchConfigurationResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchConfigurationResponseContentWithDefaults() *PatchConfigurationResponseContent {
	this := PatchConfigurationResponseContent{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PatchConfigurationResponseContent) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchConfigurationResponseContent) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchConfigurationResponseContent) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PatchConfigurationResponseContent) SetStatus(v string) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *PatchConfigurationResponseContent) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchConfigurationResponseContent) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *PatchConfigurationResponseContent) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *PatchConfigurationResponseContent) SetMessage(v string) {
	o.Message = &v
}

func (o PatchConfigurationResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchConfigurationResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullablePatchConfigurationResponseContent struct {
	value *PatchConfigurationResponseContent
	isSet bool
}

func (v NullablePatchConfigurationResponseContent) Get() *PatchConfigurationResponseContent {
	return v.value
}

func (v *NullablePatchConfigurationResponseContent) Set(val *PatchConfigurationResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchConfigurationResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchConfigurationResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchConfigurationResponseContent(val *PatchConfigurationResponseContent) *NullablePatchConfigurationResponseContent {
	return &NullablePatchConfigurationResponseContent{value: val, isSet: true}
}

func (v NullablePatchConfigurationResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchConfigurationResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


