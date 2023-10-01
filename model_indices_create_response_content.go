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

// checks if the IndicesCreateResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndicesCreateResponseContent{}

// IndicesCreateResponseContent struct for IndicesCreateResponseContent
type IndicesCreateResponseContent struct {
	Index string `json:"index"`
	ShardsAcknowledged bool `json:"shards_acknowledged"`
	Acknowledged bool `json:"acknowledged"`
}

// NewIndicesCreateResponseContent instantiates a new IndicesCreateResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndicesCreateResponseContent(index string, shardsAcknowledged bool, acknowledged bool) *IndicesCreateResponseContent {
	this := IndicesCreateResponseContent{}
	this.Index = index
	this.ShardsAcknowledged = shardsAcknowledged
	this.Acknowledged = acknowledged
	return &this
}

// NewIndicesCreateResponseContentWithDefaults instantiates a new IndicesCreateResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndicesCreateResponseContentWithDefaults() *IndicesCreateResponseContent {
	this := IndicesCreateResponseContent{}
	return &this
}

// GetIndex returns the Index field value
func (o *IndicesCreateResponseContent) GetIndex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *IndicesCreateResponseContent) GetIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *IndicesCreateResponseContent) SetIndex(v string) {
	o.Index = v
}

// GetShardsAcknowledged returns the ShardsAcknowledged field value
func (o *IndicesCreateResponseContent) GetShardsAcknowledged() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ShardsAcknowledged
}

// GetShardsAcknowledgedOk returns a tuple with the ShardsAcknowledged field value
// and a boolean to check if the value has been set.
func (o *IndicesCreateResponseContent) GetShardsAcknowledgedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShardsAcknowledged, true
}

// SetShardsAcknowledged sets field value
func (o *IndicesCreateResponseContent) SetShardsAcknowledged(v bool) {
	o.ShardsAcknowledged = v
}

// GetAcknowledged returns the Acknowledged field value
func (o *IndicesCreateResponseContent) GetAcknowledged() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Acknowledged
}

// GetAcknowledgedOk returns a tuple with the Acknowledged field value
// and a boolean to check if the value has been set.
func (o *IndicesCreateResponseContent) GetAcknowledgedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Acknowledged, true
}

// SetAcknowledged sets field value
func (o *IndicesCreateResponseContent) SetAcknowledged(v bool) {
	o.Acknowledged = v
}

func (o IndicesCreateResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndicesCreateResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	toSerialize["shards_acknowledged"] = o.ShardsAcknowledged
	toSerialize["acknowledged"] = o.Acknowledged
	return toSerialize, nil
}

type NullableIndicesCreateResponseContent struct {
	value *IndicesCreateResponseContent
	isSet bool
}

func (v NullableIndicesCreateResponseContent) Get() *IndicesCreateResponseContent {
	return v.value
}

func (v *NullableIndicesCreateResponseContent) Set(val *IndicesCreateResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableIndicesCreateResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableIndicesCreateResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndicesCreateResponseContent(val *IndicesCreateResponseContent) *NullableIndicesCreateResponseContent {
	return &NullableIndicesCreateResponseContent{value: val, isSet: true}
}

func (v NullableIndicesCreateResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndicesCreateResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

