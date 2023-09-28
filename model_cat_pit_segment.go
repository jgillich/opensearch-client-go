/*
OpenSearch

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2021-11-23
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CatPitSegment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatPitSegment{}

// CatPitSegment struct for CatPitSegment
type CatPitSegment struct {
	Index *string `json:"index,omitempty"`
	Shard *int32 `json:"shard,omitempty"`
	// Set to true to return stats only for primary shards.
	Prirep *bool `json:"prirep,omitempty"`
	Ip *string `json:"ip,omitempty"`
	Segment *string `json:"segment,omitempty"`
	Generation *int32 `json:"generation,omitempty"`
	DocsCount *int32 `json:"docs_count,omitempty"`
	DocsDeleted *int32 `json:"docs_deleted,omitempty"`
	Size *string `json:"size,omitempty"`
	SizeMemory *int32 `json:"size_memory,omitempty"`
	Committed *bool `json:"committed,omitempty"`
	Searchable *bool `json:"searchable,omitempty"`
	Version *string `json:"version,omitempty"`
	Compound *bool `json:"compound,omitempty"`
}

// NewCatPitSegment instantiates a new CatPitSegment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatPitSegment() *CatPitSegment {
	this := CatPitSegment{}
	return &this
}

// NewCatPitSegmentWithDefaults instantiates a new CatPitSegment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatPitSegmentWithDefaults() *CatPitSegment {
	this := CatPitSegment{}
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *CatPitSegment) GetIndex() string {
	if o == nil || IsNil(o.Index) {
		var ret string
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatPitSegment) GetIndexOk() (*string, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *CatPitSegment) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given string and assigns it to the Index field.
func (o *CatPitSegment) SetIndex(v string) {
	o.Index = &v
}

// GetShard returns the Shard field value if set, zero value otherwise.
func (o *CatPitSegment) GetShard() int32 {
	if o == nil || IsNil(o.Shard) {
		var ret int32
		return ret
	}
	return *o.Shard
}

// GetShardOk returns a tuple with the Shard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatPitSegment) GetShardOk() (*int32, bool) {
	if o == nil || IsNil(o.Shard) {
		return nil, false
	}
	return o.Shard, true
}

// HasShard returns a boolean if a field has been set.
func (o *CatPitSegment) HasShard() bool {
	if o != nil && !IsNil(o.Shard) {
		return true
	}

	return false
}

// SetShard gets a reference to the given int32 and assigns it to the Shard field.
func (o *CatPitSegment) SetShard(v int32) {
	o.Shard = &v
}

// GetPrirep returns the Prirep field value if set, zero value otherwise.
func (o *CatPitSegment) GetPrirep() bool {
	if o == nil || IsNil(o.Prirep) {
		var ret bool
		return ret
	}
	return *o.Prirep
}

// GetPrirepOk returns a tuple with the Prirep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatPitSegment) GetPrirepOk() (*bool, bool) {
	if o == nil || IsNil(o.Prirep) {
		return nil, false
	}
	return o.Prirep, true
}

// HasPrirep returns a boolean if a field has been set.
func (o *CatPitSegment) HasPrirep() bool {
	if o != nil && !IsNil(o.Prirep) {
		return true
	}

	return false
}

// SetPrirep gets a reference to the given bool and assigns it to the Prirep field.
func (o *CatPitSegment) SetPrirep(v bool) {
	o.Prirep = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *CatPitSegment) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatPitSegment) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *CatPitSegment) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *CatPitSegment) SetIp(v string) {
	o.Ip = &v
}

// GetSegment returns the Segment field value if set, zero value otherwise.
func (o *CatPitSegment) GetSegment() string {
	if o == nil || IsNil(o.Segment) {
		var ret string
		return ret
	}
	return *o.Segment
}

// GetSegmentOk returns a tuple with the Segment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatPitSegment) GetSegmentOk() (*string, bool) {
	if o == nil || IsNil(o.Segment) {
		return nil, false
	}
	return o.Segment, true
}

// HasSegment returns a boolean if a field has been set.
func (o *CatPitSegment) HasSegment() bool {
	if o != nil && !IsNil(o.Segment) {
		return true
	}

	return false
}

// SetSegment gets a reference to the given string and assigns it to the Segment field.
func (o *CatPitSegment) SetSegment(v string) {
	o.Segment = &v
}

// GetGeneration returns the Generation field value if set, zero value otherwise.
func (o *CatPitSegment) GetGeneration() int32 {
	if o == nil || IsNil(o.Generation) {
		var ret int32
		return ret
	}
	return *o.Generation
}

// GetGenerationOk returns a tuple with the Generation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatPitSegment) GetGenerationOk() (*int32, bool) {
	if o == nil || IsNil(o.Generation) {
		return nil, false
	}
	return o.Generation, true
}

// HasGeneration returns a boolean if a field has been set.
func (o *CatPitSegment) HasGeneration() bool {
	if o != nil && !IsNil(o.Generation) {
		return true
	}

	return false
}

// SetGeneration gets a reference to the given int32 and assigns it to the Generation field.
func (o *CatPitSegment) SetGeneration(v int32) {
	o.Generation = &v
}

// GetDocsCount returns the DocsCount field value if set, zero value otherwise.
func (o *CatPitSegment) GetDocsCount() int32 {
	if o == nil || IsNil(o.DocsCount) {
		var ret int32
		return ret
	}
	return *o.DocsCount
}

// GetDocsCountOk returns a tuple with the DocsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatPitSegment) GetDocsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.DocsCount) {
		return nil, false
	}
	return o.DocsCount, true
}

// HasDocsCount returns a boolean if a field has been set.
func (o *CatPitSegment) HasDocsCount() bool {
	if o != nil && !IsNil(o.DocsCount) {
		return true
	}

	return false
}

// SetDocsCount gets a reference to the given int32 and assigns it to the DocsCount field.
func (o *CatPitSegment) SetDocsCount(v int32) {
	o.DocsCount = &v
}

// GetDocsDeleted returns the DocsDeleted field value if set, zero value otherwise.
func (o *CatPitSegment) GetDocsDeleted() int32 {
	if o == nil || IsNil(o.DocsDeleted) {
		var ret int32
		return ret
	}
	return *o.DocsDeleted
}

// GetDocsDeletedOk returns a tuple with the DocsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatPitSegment) GetDocsDeletedOk() (*int32, bool) {
	if o == nil || IsNil(o.DocsDeleted) {
		return nil, false
	}
	return o.DocsDeleted, true
}

// HasDocsDeleted returns a boolean if a field has been set.
func (o *CatPitSegment) HasDocsDeleted() bool {
	if o != nil && !IsNil(o.DocsDeleted) {
		return true
	}

	return false
}

// SetDocsDeleted gets a reference to the given int32 and assigns it to the DocsDeleted field.
func (o *CatPitSegment) SetDocsDeleted(v int32) {
	o.DocsDeleted = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *CatPitSegment) GetSize() string {
	if o == nil || IsNil(o.Size) {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatPitSegment) GetSizeOk() (*string, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *CatPitSegment) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *CatPitSegment) SetSize(v string) {
	o.Size = &v
}

// GetSizeMemory returns the SizeMemory field value if set, zero value otherwise.
func (o *CatPitSegment) GetSizeMemory() int32 {
	if o == nil || IsNil(o.SizeMemory) {
		var ret int32
		return ret
	}
	return *o.SizeMemory
}

// GetSizeMemoryOk returns a tuple with the SizeMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatPitSegment) GetSizeMemoryOk() (*int32, bool) {
	if o == nil || IsNil(o.SizeMemory) {
		return nil, false
	}
	return o.SizeMemory, true
}

// HasSizeMemory returns a boolean if a field has been set.
func (o *CatPitSegment) HasSizeMemory() bool {
	if o != nil && !IsNil(o.SizeMemory) {
		return true
	}

	return false
}

// SetSizeMemory gets a reference to the given int32 and assigns it to the SizeMemory field.
func (o *CatPitSegment) SetSizeMemory(v int32) {
	o.SizeMemory = &v
}

// GetCommitted returns the Committed field value if set, zero value otherwise.
func (o *CatPitSegment) GetCommitted() bool {
	if o == nil || IsNil(o.Committed) {
		var ret bool
		return ret
	}
	return *o.Committed
}

// GetCommittedOk returns a tuple with the Committed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatPitSegment) GetCommittedOk() (*bool, bool) {
	if o == nil || IsNil(o.Committed) {
		return nil, false
	}
	return o.Committed, true
}

// HasCommitted returns a boolean if a field has been set.
func (o *CatPitSegment) HasCommitted() bool {
	if o != nil && !IsNil(o.Committed) {
		return true
	}

	return false
}

// SetCommitted gets a reference to the given bool and assigns it to the Committed field.
func (o *CatPitSegment) SetCommitted(v bool) {
	o.Committed = &v
}

// GetSearchable returns the Searchable field value if set, zero value otherwise.
func (o *CatPitSegment) GetSearchable() bool {
	if o == nil || IsNil(o.Searchable) {
		var ret bool
		return ret
	}
	return *o.Searchable
}

// GetSearchableOk returns a tuple with the Searchable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatPitSegment) GetSearchableOk() (*bool, bool) {
	if o == nil || IsNil(o.Searchable) {
		return nil, false
	}
	return o.Searchable, true
}

// HasSearchable returns a boolean if a field has been set.
func (o *CatPitSegment) HasSearchable() bool {
	if o != nil && !IsNil(o.Searchable) {
		return true
	}

	return false
}

// SetSearchable gets a reference to the given bool and assigns it to the Searchable field.
func (o *CatPitSegment) SetSearchable(v bool) {
	o.Searchable = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CatPitSegment) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatPitSegment) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CatPitSegment) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *CatPitSegment) SetVersion(v string) {
	o.Version = &v
}

// GetCompound returns the Compound field value if set, zero value otherwise.
func (o *CatPitSegment) GetCompound() bool {
	if o == nil || IsNil(o.Compound) {
		var ret bool
		return ret
	}
	return *o.Compound
}

// GetCompoundOk returns a tuple with the Compound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatPitSegment) GetCompoundOk() (*bool, bool) {
	if o == nil || IsNil(o.Compound) {
		return nil, false
	}
	return o.Compound, true
}

// HasCompound returns a boolean if a field has been set.
func (o *CatPitSegment) HasCompound() bool {
	if o != nil && !IsNil(o.Compound) {
		return true
	}

	return false
}

// SetCompound gets a reference to the given bool and assigns it to the Compound field.
func (o *CatPitSegment) SetCompound(v bool) {
	o.Compound = &v
}

func (o CatPitSegment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatPitSegment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	if !IsNil(o.Shard) {
		toSerialize["shard"] = o.Shard
	}
	if !IsNil(o.Prirep) {
		toSerialize["prirep"] = o.Prirep
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.Segment) {
		toSerialize["segment"] = o.Segment
	}
	if !IsNil(o.Generation) {
		toSerialize["generation"] = o.Generation
	}
	if !IsNil(o.DocsCount) {
		toSerialize["docs_count"] = o.DocsCount
	}
	if !IsNil(o.DocsDeleted) {
		toSerialize["docs_deleted"] = o.DocsDeleted
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.SizeMemory) {
		toSerialize["size_memory"] = o.SizeMemory
	}
	if !IsNil(o.Committed) {
		toSerialize["committed"] = o.Committed
	}
	if !IsNil(o.Searchable) {
		toSerialize["searchable"] = o.Searchable
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Compound) {
		toSerialize["compound"] = o.Compound
	}
	return toSerialize, nil
}

type NullableCatPitSegment struct {
	value *CatPitSegment
	isSet bool
}

func (v NullableCatPitSegment) Get() *CatPitSegment {
	return v.value
}

func (v *NullableCatPitSegment) Set(val *CatPitSegment) {
	v.value = val
	v.isSet = true
}

func (v NullableCatPitSegment) IsSet() bool {
	return v.isSet
}

func (v *NullableCatPitSegment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatPitSegment(val *CatPitSegment) *NullableCatPitSegment {
	return &NullableCatPitSegment{value: val, isSet: true}
}

func (v NullableCatPitSegment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatPitSegment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


