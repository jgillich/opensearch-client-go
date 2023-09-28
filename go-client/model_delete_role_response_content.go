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

// checks if the DeleteRoleResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteRoleResponseContent{}

// DeleteRoleResponseContent struct for DeleteRoleResponseContent
type DeleteRoleResponseContent struct {
	// Security Operation Status
	Status *string `json:"status,omitempty"`
	// Security Operation Message
	Message *string `json:"message,omitempty"`
}

// NewDeleteRoleResponseContent instantiates a new DeleteRoleResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteRoleResponseContent() *DeleteRoleResponseContent {
	this := DeleteRoleResponseContent{}
	return &this
}

// NewDeleteRoleResponseContentWithDefaults instantiates a new DeleteRoleResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteRoleResponseContentWithDefaults() *DeleteRoleResponseContent {
	this := DeleteRoleResponseContent{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeleteRoleResponseContent) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteRoleResponseContent) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeleteRoleResponseContent) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DeleteRoleResponseContent) SetStatus(v string) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DeleteRoleResponseContent) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteRoleResponseContent) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DeleteRoleResponseContent) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *DeleteRoleResponseContent) SetMessage(v string) {
	o.Message = &v
}

func (o DeleteRoleResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteRoleResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableDeleteRoleResponseContent struct {
	value *DeleteRoleResponseContent
	isSet bool
}

func (v NullableDeleteRoleResponseContent) Get() *DeleteRoleResponseContent {
	return v.value
}

func (v *NullableDeleteRoleResponseContent) Set(val *DeleteRoleResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteRoleResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteRoleResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteRoleResponseContent(val *DeleteRoleResponseContent) *NullableDeleteRoleResponseContent {
	return &NullableDeleteRoleResponseContent{value: val, isSet: true}
}

func (v NullableDeleteRoleResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteRoleResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


