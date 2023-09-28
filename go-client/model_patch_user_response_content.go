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

// checks if the PatchUserResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchUserResponseContent{}

// PatchUserResponseContent struct for PatchUserResponseContent
type PatchUserResponseContent struct {
	// Security Operation Status
	Status *string `json:"status,omitempty"`
	// Security Operation Message
	Message *string `json:"message,omitempty"`
}

// NewPatchUserResponseContent instantiates a new PatchUserResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchUserResponseContent() *PatchUserResponseContent {
	this := PatchUserResponseContent{}
	return &this
}

// NewPatchUserResponseContentWithDefaults instantiates a new PatchUserResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchUserResponseContentWithDefaults() *PatchUserResponseContent {
	this := PatchUserResponseContent{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PatchUserResponseContent) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchUserResponseContent) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchUserResponseContent) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PatchUserResponseContent) SetStatus(v string) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *PatchUserResponseContent) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchUserResponseContent) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *PatchUserResponseContent) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *PatchUserResponseContent) SetMessage(v string) {
	o.Message = &v
}

func (o PatchUserResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchUserResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullablePatchUserResponseContent struct {
	value *PatchUserResponseContent
	isSet bool
}

func (v NullablePatchUserResponseContent) Get() *PatchUserResponseContent {
	return v.value
}

func (v *NullablePatchUserResponseContent) Set(val *PatchUserResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchUserResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchUserResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchUserResponseContent(val *PatchUserResponseContent) *NullablePatchUserResponseContent {
	return &NullablePatchUserResponseContent{value: val, isSet: true}
}

func (v NullablePatchUserResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchUserResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


