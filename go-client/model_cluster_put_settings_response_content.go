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

// checks if the ClusterPutSettingsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterPutSettingsResponseContent{}

// ClusterPutSettingsResponseContent struct for ClusterPutSettingsResponseContent
type ClusterPutSettingsResponseContent struct {
	Acknowledged *bool `json:"acknowledged,omitempty"`
	Persistent map[string]interface{} `json:"persistent,omitempty"`
	Transient map[string]interface{} `json:"transient,omitempty"`
}

// NewClusterPutSettingsResponseContent instantiates a new ClusterPutSettingsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterPutSettingsResponseContent() *ClusterPutSettingsResponseContent {
	this := ClusterPutSettingsResponseContent{}
	return &this
}

// NewClusterPutSettingsResponseContentWithDefaults instantiates a new ClusterPutSettingsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterPutSettingsResponseContentWithDefaults() *ClusterPutSettingsResponseContent {
	this := ClusterPutSettingsResponseContent{}
	return &this
}

// GetAcknowledged returns the Acknowledged field value if set, zero value otherwise.
func (o *ClusterPutSettingsResponseContent) GetAcknowledged() bool {
	if o == nil || IsNil(o.Acknowledged) {
		var ret bool
		return ret
	}
	return *o.Acknowledged
}

// GetAcknowledgedOk returns a tuple with the Acknowledged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterPutSettingsResponseContent) GetAcknowledgedOk() (*bool, bool) {
	if o == nil || IsNil(o.Acknowledged) {
		return nil, false
	}
	return o.Acknowledged, true
}

// HasAcknowledged returns a boolean if a field has been set.
func (o *ClusterPutSettingsResponseContent) HasAcknowledged() bool {
	if o != nil && !IsNil(o.Acknowledged) {
		return true
	}

	return false
}

// SetAcknowledged gets a reference to the given bool and assigns it to the Acknowledged field.
func (o *ClusterPutSettingsResponseContent) SetAcknowledged(v bool) {
	o.Acknowledged = &v
}

// GetPersistent returns the Persistent field value if set, zero value otherwise.
func (o *ClusterPutSettingsResponseContent) GetPersistent() map[string]interface{} {
	if o == nil || IsNil(o.Persistent) {
		var ret map[string]interface{}
		return ret
	}
	return o.Persistent
}

// GetPersistentOk returns a tuple with the Persistent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterPutSettingsResponseContent) GetPersistentOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Persistent) {
		return map[string]interface{}{}, false
	}
	return o.Persistent, true
}

// HasPersistent returns a boolean if a field has been set.
func (o *ClusterPutSettingsResponseContent) HasPersistent() bool {
	if o != nil && !IsNil(o.Persistent) {
		return true
	}

	return false
}

// SetPersistent gets a reference to the given map[string]interface{} and assigns it to the Persistent field.
func (o *ClusterPutSettingsResponseContent) SetPersistent(v map[string]interface{}) {
	o.Persistent = v
}

// GetTransient returns the Transient field value if set, zero value otherwise.
func (o *ClusterPutSettingsResponseContent) GetTransient() map[string]interface{} {
	if o == nil || IsNil(o.Transient) {
		var ret map[string]interface{}
		return ret
	}
	return o.Transient
}

// GetTransientOk returns a tuple with the Transient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterPutSettingsResponseContent) GetTransientOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Transient) {
		return map[string]interface{}{}, false
	}
	return o.Transient, true
}

// HasTransient returns a boolean if a field has been set.
func (o *ClusterPutSettingsResponseContent) HasTransient() bool {
	if o != nil && !IsNil(o.Transient) {
		return true
	}

	return false
}

// SetTransient gets a reference to the given map[string]interface{} and assigns it to the Transient field.
func (o *ClusterPutSettingsResponseContent) SetTransient(v map[string]interface{}) {
	o.Transient = v
}

func (o ClusterPutSettingsResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterPutSettingsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Acknowledged) {
		toSerialize["acknowledged"] = o.Acknowledged
	}
	if !IsNil(o.Persistent) {
		toSerialize["persistent"] = o.Persistent
	}
	if !IsNil(o.Transient) {
		toSerialize["transient"] = o.Transient
	}
	return toSerialize, nil
}

type NullableClusterPutSettingsResponseContent struct {
	value *ClusterPutSettingsResponseContent
	isSet bool
}

func (v NullableClusterPutSettingsResponseContent) Get() *ClusterPutSettingsResponseContent {
	return v.value
}

func (v *NullableClusterPutSettingsResponseContent) Set(val *ClusterPutSettingsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterPutSettingsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterPutSettingsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterPutSettingsResponseContent(val *ClusterPutSettingsResponseContent) *NullableClusterPutSettingsResponseContent {
	return &NullableClusterPutSettingsResponseContent{value: val, isSet: true}
}

func (v NullableClusterPutSettingsResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterPutSettingsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


