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

// checks if the CreateRoleMappingResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRoleMappingResponseContent{}

// CreateRoleMappingResponseContent struct for CreateRoleMappingResponseContent
type CreateRoleMappingResponseContent struct {
	// Security Operation Status
	Status *string `json:"status,omitempty"`
	// Security Operation Message
	Message *string `json:"message,omitempty"`
}

// NewCreateRoleMappingResponseContent instantiates a new CreateRoleMappingResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRoleMappingResponseContent() *CreateRoleMappingResponseContent {
	this := CreateRoleMappingResponseContent{}
	return &this
}

// NewCreateRoleMappingResponseContentWithDefaults instantiates a new CreateRoleMappingResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRoleMappingResponseContentWithDefaults() *CreateRoleMappingResponseContent {
	this := CreateRoleMappingResponseContent{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CreateRoleMappingResponseContent) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRoleMappingResponseContent) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CreateRoleMappingResponseContent) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CreateRoleMappingResponseContent) SetStatus(v string) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CreateRoleMappingResponseContent) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRoleMappingResponseContent) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CreateRoleMappingResponseContent) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CreateRoleMappingResponseContent) SetMessage(v string) {
	o.Message = &v
}

func (o CreateRoleMappingResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRoleMappingResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableCreateRoleMappingResponseContent struct {
	value *CreateRoleMappingResponseContent
	isSet bool
}

func (v NullableCreateRoleMappingResponseContent) Get() *CreateRoleMappingResponseContent {
	return v.value
}

func (v *NullableCreateRoleMappingResponseContent) Set(val *CreateRoleMappingResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRoleMappingResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRoleMappingResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRoleMappingResponseContent(val *CreateRoleMappingResponseContent) *NullableCreateRoleMappingResponseContent {
	return &NullableCreateRoleMappingResponseContent{value: val, isSet: true}
}

func (v NullableCreateRoleMappingResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRoleMappingResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


