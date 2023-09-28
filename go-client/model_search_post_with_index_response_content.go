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

// checks if the SearchPostWithIndexResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchPostWithIndexResponseContent{}

// SearchPostWithIndexResponseContent struct for SearchPostWithIndexResponseContent
type SearchPostWithIndexResponseContent struct {
	ScrollId *string `json:"_scroll_id,omitempty"`
	Took *int64 `json:"took,omitempty"`
	TimedOut *bool `json:"timed_out,omitempty"`
	Shards *ShardStatistics `json:"_shards,omitempty"`
	Hits *HitsMetadata `json:"hits,omitempty"`
}

// NewSearchPostWithIndexResponseContent instantiates a new SearchPostWithIndexResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchPostWithIndexResponseContent() *SearchPostWithIndexResponseContent {
	this := SearchPostWithIndexResponseContent{}
	return &this
}

// NewSearchPostWithIndexResponseContentWithDefaults instantiates a new SearchPostWithIndexResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchPostWithIndexResponseContentWithDefaults() *SearchPostWithIndexResponseContent {
	this := SearchPostWithIndexResponseContent{}
	return &this
}

// GetScrollId returns the ScrollId field value if set, zero value otherwise.
func (o *SearchPostWithIndexResponseContent) GetScrollId() string {
	if o == nil || IsNil(o.ScrollId) {
		var ret string
		return ret
	}
	return *o.ScrollId
}

// GetScrollIdOk returns a tuple with the ScrollId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPostWithIndexResponseContent) GetScrollIdOk() (*string, bool) {
	if o == nil || IsNil(o.ScrollId) {
		return nil, false
	}
	return o.ScrollId, true
}

// HasScrollId returns a boolean if a field has been set.
func (o *SearchPostWithIndexResponseContent) HasScrollId() bool {
	if o != nil && !IsNil(o.ScrollId) {
		return true
	}

	return false
}

// SetScrollId gets a reference to the given string and assigns it to the ScrollId field.
func (o *SearchPostWithIndexResponseContent) SetScrollId(v string) {
	o.ScrollId = &v
}

// GetTook returns the Took field value if set, zero value otherwise.
func (o *SearchPostWithIndexResponseContent) GetTook() int64 {
	if o == nil || IsNil(o.Took) {
		var ret int64
		return ret
	}
	return *o.Took
}

// GetTookOk returns a tuple with the Took field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPostWithIndexResponseContent) GetTookOk() (*int64, bool) {
	if o == nil || IsNil(o.Took) {
		return nil, false
	}
	return o.Took, true
}

// HasTook returns a boolean if a field has been set.
func (o *SearchPostWithIndexResponseContent) HasTook() bool {
	if o != nil && !IsNil(o.Took) {
		return true
	}

	return false
}

// SetTook gets a reference to the given int64 and assigns it to the Took field.
func (o *SearchPostWithIndexResponseContent) SetTook(v int64) {
	o.Took = &v
}

// GetTimedOut returns the TimedOut field value if set, zero value otherwise.
func (o *SearchPostWithIndexResponseContent) GetTimedOut() bool {
	if o == nil || IsNil(o.TimedOut) {
		var ret bool
		return ret
	}
	return *o.TimedOut
}

// GetTimedOutOk returns a tuple with the TimedOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPostWithIndexResponseContent) GetTimedOutOk() (*bool, bool) {
	if o == nil || IsNil(o.TimedOut) {
		return nil, false
	}
	return o.TimedOut, true
}

// HasTimedOut returns a boolean if a field has been set.
func (o *SearchPostWithIndexResponseContent) HasTimedOut() bool {
	if o != nil && !IsNil(o.TimedOut) {
		return true
	}

	return false
}

// SetTimedOut gets a reference to the given bool and assigns it to the TimedOut field.
func (o *SearchPostWithIndexResponseContent) SetTimedOut(v bool) {
	o.TimedOut = &v
}

// GetShards returns the Shards field value if set, zero value otherwise.
func (o *SearchPostWithIndexResponseContent) GetShards() ShardStatistics {
	if o == nil || IsNil(o.Shards) {
		var ret ShardStatistics
		return ret
	}
	return *o.Shards
}

// GetShardsOk returns a tuple with the Shards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPostWithIndexResponseContent) GetShardsOk() (*ShardStatistics, bool) {
	if o == nil || IsNil(o.Shards) {
		return nil, false
	}
	return o.Shards, true
}

// HasShards returns a boolean if a field has been set.
func (o *SearchPostWithIndexResponseContent) HasShards() bool {
	if o != nil && !IsNil(o.Shards) {
		return true
	}

	return false
}

// SetShards gets a reference to the given ShardStatistics and assigns it to the Shards field.
func (o *SearchPostWithIndexResponseContent) SetShards(v ShardStatistics) {
	o.Shards = &v
}

// GetHits returns the Hits field value if set, zero value otherwise.
func (o *SearchPostWithIndexResponseContent) GetHits() HitsMetadata {
	if o == nil || IsNil(o.Hits) {
		var ret HitsMetadata
		return ret
	}
	return *o.Hits
}

// GetHitsOk returns a tuple with the Hits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPostWithIndexResponseContent) GetHitsOk() (*HitsMetadata, bool) {
	if o == nil || IsNil(o.Hits) {
		return nil, false
	}
	return o.Hits, true
}

// HasHits returns a boolean if a field has been set.
func (o *SearchPostWithIndexResponseContent) HasHits() bool {
	if o != nil && !IsNil(o.Hits) {
		return true
	}

	return false
}

// SetHits gets a reference to the given HitsMetadata and assigns it to the Hits field.
func (o *SearchPostWithIndexResponseContent) SetHits(v HitsMetadata) {
	o.Hits = &v
}

func (o SearchPostWithIndexResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchPostWithIndexResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ScrollId) {
		toSerialize["_scroll_id"] = o.ScrollId
	}
	if !IsNil(o.Took) {
		toSerialize["took"] = o.Took
	}
	if !IsNil(o.TimedOut) {
		toSerialize["timed_out"] = o.TimedOut
	}
	if !IsNil(o.Shards) {
		toSerialize["_shards"] = o.Shards
	}
	if !IsNil(o.Hits) {
		toSerialize["hits"] = o.Hits
	}
	return toSerialize, nil
}

type NullableSearchPostWithIndexResponseContent struct {
	value *SearchPostWithIndexResponseContent
	isSet bool
}

func (v NullableSearchPostWithIndexResponseContent) Get() *SearchPostWithIndexResponseContent {
	return v.value
}

func (v *NullableSearchPostWithIndexResponseContent) Set(val *SearchPostWithIndexResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchPostWithIndexResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchPostWithIndexResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchPostWithIndexResponseContent(val *SearchPostWithIndexResponseContent) *NullableSearchPostWithIndexResponseContent {
	return &NullableSearchPostWithIndexResponseContent{value: val, isSet: true}
}

func (v NullableSearchPostWithIndexResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchPostWithIndexResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

