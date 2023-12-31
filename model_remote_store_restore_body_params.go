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

// checks if the RemoteStoreRestoreBodyParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoteStoreRestoreBodyParams{}

// RemoteStoreRestoreBodyParams Comma-separated list of index IDs
type RemoteStoreRestoreBodyParams struct {
	Indices []string `json:"indices"`
}

// NewRemoteStoreRestoreBodyParams instantiates a new RemoteStoreRestoreBodyParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteStoreRestoreBodyParams(indices []string) *RemoteStoreRestoreBodyParams {
	this := RemoteStoreRestoreBodyParams{}
	this.Indices = indices
	return &this
}

// NewRemoteStoreRestoreBodyParamsWithDefaults instantiates a new RemoteStoreRestoreBodyParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteStoreRestoreBodyParamsWithDefaults() *RemoteStoreRestoreBodyParams {
	this := RemoteStoreRestoreBodyParams{}
	return &this
}

// GetIndices returns the Indices field value
func (o *RemoteStoreRestoreBodyParams) GetIndices() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Indices
}

// GetIndicesOk returns a tuple with the Indices field value
// and a boolean to check if the value has been set.
func (o *RemoteStoreRestoreBodyParams) GetIndicesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Indices, true
}

// SetIndices sets field value
func (o *RemoteStoreRestoreBodyParams) SetIndices(v []string) {
	o.Indices = v
}

func (o RemoteStoreRestoreBodyParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoteStoreRestoreBodyParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["indices"] = o.Indices
	return toSerialize, nil
}

type NullableRemoteStoreRestoreBodyParams struct {
	value *RemoteStoreRestoreBodyParams
	isSet bool
}

func (v NullableRemoteStoreRestoreBodyParams) Get() *RemoteStoreRestoreBodyParams {
	return v.value
}

func (v *NullableRemoteStoreRestoreBodyParams) Set(val *RemoteStoreRestoreBodyParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteStoreRestoreBodyParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteStoreRestoreBodyParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteStoreRestoreBodyParams(val *RemoteStoreRestoreBodyParams) *NullableRemoteStoreRestoreBodyParams {
	return &NullableRemoteStoreRestoreBodyParams{value: val, isSet: true}
}

func (v NullableRemoteStoreRestoreBodyParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteStoreRestoreBodyParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


