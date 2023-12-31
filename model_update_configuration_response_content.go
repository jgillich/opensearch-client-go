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

// checks if the UpdateConfigurationResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateConfigurationResponseContent{}

// UpdateConfigurationResponseContent struct for UpdateConfigurationResponseContent
type UpdateConfigurationResponseContent struct {
	// Security Operation Status
	Status *string `json:"status,omitempty"`
	// Security Operation Message
	Message *string `json:"message,omitempty"`
}

// NewUpdateConfigurationResponseContent instantiates a new UpdateConfigurationResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateConfigurationResponseContent() *UpdateConfigurationResponseContent {
	this := UpdateConfigurationResponseContent{}
	return &this
}

// NewUpdateConfigurationResponseContentWithDefaults instantiates a new UpdateConfigurationResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateConfigurationResponseContentWithDefaults() *UpdateConfigurationResponseContent {
	this := UpdateConfigurationResponseContent{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UpdateConfigurationResponseContent) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateConfigurationResponseContent) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UpdateConfigurationResponseContent) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UpdateConfigurationResponseContent) SetStatus(v string) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *UpdateConfigurationResponseContent) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateConfigurationResponseContent) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *UpdateConfigurationResponseContent) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *UpdateConfigurationResponseContent) SetMessage(v string) {
	o.Message = &v
}

func (o UpdateConfigurationResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateConfigurationResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableUpdateConfigurationResponseContent struct {
	value *UpdateConfigurationResponseContent
	isSet bool
}

func (v NullableUpdateConfigurationResponseContent) Get() *UpdateConfigurationResponseContent {
	return v.value
}

func (v *NullableUpdateConfigurationResponseContent) Set(val *UpdateConfigurationResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateConfigurationResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateConfigurationResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateConfigurationResponseContent(val *UpdateConfigurationResponseContent) *NullableUpdateConfigurationResponseContent {
	return &NullableUpdateConfigurationResponseContent{value: val, isSet: true}
}

func (v NullableUpdateConfigurationResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateConfigurationResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


