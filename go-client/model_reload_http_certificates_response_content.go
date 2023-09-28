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

// checks if the ReloadHttpCertificatesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReloadHttpCertificatesResponseContent{}

// ReloadHttpCertificatesResponseContent struct for ReloadHttpCertificatesResponseContent
type ReloadHttpCertificatesResponseContent struct {
	// Security Operation Status
	Status *string `json:"status,omitempty"`
	// Security Operation Message
	Message *string `json:"message,omitempty"`
}

// NewReloadHttpCertificatesResponseContent instantiates a new ReloadHttpCertificatesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReloadHttpCertificatesResponseContent() *ReloadHttpCertificatesResponseContent {
	this := ReloadHttpCertificatesResponseContent{}
	return &this
}

// NewReloadHttpCertificatesResponseContentWithDefaults instantiates a new ReloadHttpCertificatesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReloadHttpCertificatesResponseContentWithDefaults() *ReloadHttpCertificatesResponseContent {
	this := ReloadHttpCertificatesResponseContent{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ReloadHttpCertificatesResponseContent) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReloadHttpCertificatesResponseContent) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ReloadHttpCertificatesResponseContent) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ReloadHttpCertificatesResponseContent) SetStatus(v string) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ReloadHttpCertificatesResponseContent) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReloadHttpCertificatesResponseContent) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ReloadHttpCertificatesResponseContent) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ReloadHttpCertificatesResponseContent) SetMessage(v string) {
	o.Message = &v
}

func (o ReloadHttpCertificatesResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReloadHttpCertificatesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableReloadHttpCertificatesResponseContent struct {
	value *ReloadHttpCertificatesResponseContent
	isSet bool
}

func (v NullableReloadHttpCertificatesResponseContent) Get() *ReloadHttpCertificatesResponseContent {
	return v.value
}

func (v *NullableReloadHttpCertificatesResponseContent) Set(val *ReloadHttpCertificatesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableReloadHttpCertificatesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableReloadHttpCertificatesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReloadHttpCertificatesResponseContent(val *ReloadHttpCertificatesResponseContent) *NullableReloadHttpCertificatesResponseContent {
	return &NullableReloadHttpCertificatesResponseContent{value: val, isSet: true}
}

func (v NullableReloadHttpCertificatesResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReloadHttpCertificatesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


