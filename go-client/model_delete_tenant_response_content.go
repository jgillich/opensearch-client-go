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

// checks if the DeleteTenantResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteTenantResponseContent{}

// DeleteTenantResponseContent struct for DeleteTenantResponseContent
type DeleteTenantResponseContent struct {
	// Security Operation Status
	Status *string `json:"status,omitempty"`
	// Security Operation Message
	Message *string `json:"message,omitempty"`
}

// NewDeleteTenantResponseContent instantiates a new DeleteTenantResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteTenantResponseContent() *DeleteTenantResponseContent {
	this := DeleteTenantResponseContent{}
	return &this
}

// NewDeleteTenantResponseContentWithDefaults instantiates a new DeleteTenantResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteTenantResponseContentWithDefaults() *DeleteTenantResponseContent {
	this := DeleteTenantResponseContent{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeleteTenantResponseContent) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteTenantResponseContent) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeleteTenantResponseContent) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DeleteTenantResponseContent) SetStatus(v string) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DeleteTenantResponseContent) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteTenantResponseContent) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DeleteTenantResponseContent) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *DeleteTenantResponseContent) SetMessage(v string) {
	o.Message = &v
}

func (o DeleteTenantResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteTenantResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableDeleteTenantResponseContent struct {
	value *DeleteTenantResponseContent
	isSet bool
}

func (v NullableDeleteTenantResponseContent) Get() *DeleteTenantResponseContent {
	return v.value
}

func (v *NullableDeleteTenantResponseContent) Set(val *DeleteTenantResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteTenantResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteTenantResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteTenantResponseContent(val *DeleteTenantResponseContent) *NullableDeleteTenantResponseContent {
	return &NullableDeleteTenantResponseContent{value: val, isSet: true}
}

func (v NullableDeleteTenantResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteTenantResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


