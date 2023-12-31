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

// checks if the FlushCacheResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlushCacheResponseContent{}

// FlushCacheResponseContent struct for FlushCacheResponseContent
type FlushCacheResponseContent struct {
	// Security Operation Status
	Status *string `json:"status,omitempty"`
	// Security Operation Message
	Message *string `json:"message,omitempty"`
}

// NewFlushCacheResponseContent instantiates a new FlushCacheResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlushCacheResponseContent() *FlushCacheResponseContent {
	this := FlushCacheResponseContent{}
	return &this
}

// NewFlushCacheResponseContentWithDefaults instantiates a new FlushCacheResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlushCacheResponseContentWithDefaults() *FlushCacheResponseContent {
	this := FlushCacheResponseContent{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FlushCacheResponseContent) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlushCacheResponseContent) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FlushCacheResponseContent) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *FlushCacheResponseContent) SetStatus(v string) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *FlushCacheResponseContent) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlushCacheResponseContent) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *FlushCacheResponseContent) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *FlushCacheResponseContent) SetMessage(v string) {
	o.Message = &v
}

func (o FlushCacheResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlushCacheResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableFlushCacheResponseContent struct {
	value *FlushCacheResponseContent
	isSet bool
}

func (v NullableFlushCacheResponseContent) Get() *FlushCacheResponseContent {
	return v.value
}

func (v *NullableFlushCacheResponseContent) Set(val *FlushCacheResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableFlushCacheResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableFlushCacheResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlushCacheResponseContent(val *FlushCacheResponseContent) *NullableFlushCacheResponseContent {
	return &NullableFlushCacheResponseContent{value: val, isSet: true}
}

func (v NullableFlushCacheResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlushCacheResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


