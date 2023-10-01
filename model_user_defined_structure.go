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

// checks if the UserDefinedStructure type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserDefinedStructure{}

// UserDefinedStructure struct for UserDefinedStructure
type UserDefinedStructure struct {
	Alias *string `json:"alias,omitempty"`
	Aliases []string `json:"aliases,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
	Index *string `json:"index,omitempty"`
	Indices []string `json:"indices,omitempty"`
	IndexRouting *string `json:"index_routing,omitempty"`
	IsHidden *bool `json:"is_hidden,omitempty"`
	IsWriteIndex *bool `json:"is_write_index,omitempty"`
	MustExist *string `json:"must_exist,omitempty"`
	Routing *string `json:"routing,omitempty"`
	SearchRouting *string `json:"search_routing,omitempty"`
}

// NewUserDefinedStructure instantiates a new UserDefinedStructure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDefinedStructure() *UserDefinedStructure {
	this := UserDefinedStructure{}
	return &this
}

// NewUserDefinedStructureWithDefaults instantiates a new UserDefinedStructure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDefinedStructureWithDefaults() *UserDefinedStructure {
	this := UserDefinedStructure{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *UserDefinedStructure) GetAlias() string {
	if o == nil || IsNil(o.Alias) {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedStructure) GetAliasOk() (*string, bool) {
	if o == nil || IsNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *UserDefinedStructure) HasAlias() bool {
	if o != nil && !IsNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *UserDefinedStructure) SetAlias(v string) {
	o.Alias = &v
}

// GetAliases returns the Aliases field value if set, zero value otherwise.
func (o *UserDefinedStructure) GetAliases() []string {
	if o == nil || IsNil(o.Aliases) {
		var ret []string
		return ret
	}
	return o.Aliases
}

// GetAliasesOk returns a tuple with the Aliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedStructure) GetAliasesOk() ([]string, bool) {
	if o == nil || IsNil(o.Aliases) {
		return nil, false
	}
	return o.Aliases, true
}

// HasAliases returns a boolean if a field has been set.
func (o *UserDefinedStructure) HasAliases() bool {
	if o != nil && !IsNil(o.Aliases) {
		return true
	}

	return false
}

// SetAliases gets a reference to the given []string and assigns it to the Aliases field.
func (o *UserDefinedStructure) SetAliases(v []string) {
	o.Aliases = v
}

// GetFilter returns the Filter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDefinedStructure) GetFilter() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDefinedStructure) GetFilterOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return &o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *UserDefinedStructure) HasFilter() bool {
	if o != nil && IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given interface{} and assigns it to the Filter field.
func (o *UserDefinedStructure) SetFilter(v interface{}) {
	o.Filter = v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *UserDefinedStructure) GetIndex() string {
	if o == nil || IsNil(o.Index) {
		var ret string
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedStructure) GetIndexOk() (*string, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *UserDefinedStructure) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given string and assigns it to the Index field.
func (o *UserDefinedStructure) SetIndex(v string) {
	o.Index = &v
}

// GetIndices returns the Indices field value if set, zero value otherwise.
func (o *UserDefinedStructure) GetIndices() []string {
	if o == nil || IsNil(o.Indices) {
		var ret []string
		return ret
	}
	return o.Indices
}

// GetIndicesOk returns a tuple with the Indices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedStructure) GetIndicesOk() ([]string, bool) {
	if o == nil || IsNil(o.Indices) {
		return nil, false
	}
	return o.Indices, true
}

// HasIndices returns a boolean if a field has been set.
func (o *UserDefinedStructure) HasIndices() bool {
	if o != nil && !IsNil(o.Indices) {
		return true
	}

	return false
}

// SetIndices gets a reference to the given []string and assigns it to the Indices field.
func (o *UserDefinedStructure) SetIndices(v []string) {
	o.Indices = v
}

// GetIndexRouting returns the IndexRouting field value if set, zero value otherwise.
func (o *UserDefinedStructure) GetIndexRouting() string {
	if o == nil || IsNil(o.IndexRouting) {
		var ret string
		return ret
	}
	return *o.IndexRouting
}

// GetIndexRoutingOk returns a tuple with the IndexRouting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedStructure) GetIndexRoutingOk() (*string, bool) {
	if o == nil || IsNil(o.IndexRouting) {
		return nil, false
	}
	return o.IndexRouting, true
}

// HasIndexRouting returns a boolean if a field has been set.
func (o *UserDefinedStructure) HasIndexRouting() bool {
	if o != nil && !IsNil(o.IndexRouting) {
		return true
	}

	return false
}

// SetIndexRouting gets a reference to the given string and assigns it to the IndexRouting field.
func (o *UserDefinedStructure) SetIndexRouting(v string) {
	o.IndexRouting = &v
}

// GetIsHidden returns the IsHidden field value if set, zero value otherwise.
func (o *UserDefinedStructure) GetIsHidden() bool {
	if o == nil || IsNil(o.IsHidden) {
		var ret bool
		return ret
	}
	return *o.IsHidden
}

// GetIsHiddenOk returns a tuple with the IsHidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedStructure) GetIsHiddenOk() (*bool, bool) {
	if o == nil || IsNil(o.IsHidden) {
		return nil, false
	}
	return o.IsHidden, true
}

// HasIsHidden returns a boolean if a field has been set.
func (o *UserDefinedStructure) HasIsHidden() bool {
	if o != nil && !IsNil(o.IsHidden) {
		return true
	}

	return false
}

// SetIsHidden gets a reference to the given bool and assigns it to the IsHidden field.
func (o *UserDefinedStructure) SetIsHidden(v bool) {
	o.IsHidden = &v
}

// GetIsWriteIndex returns the IsWriteIndex field value if set, zero value otherwise.
func (o *UserDefinedStructure) GetIsWriteIndex() bool {
	if o == nil || IsNil(o.IsWriteIndex) {
		var ret bool
		return ret
	}
	return *o.IsWriteIndex
}

// GetIsWriteIndexOk returns a tuple with the IsWriteIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedStructure) GetIsWriteIndexOk() (*bool, bool) {
	if o == nil || IsNil(o.IsWriteIndex) {
		return nil, false
	}
	return o.IsWriteIndex, true
}

// HasIsWriteIndex returns a boolean if a field has been set.
func (o *UserDefinedStructure) HasIsWriteIndex() bool {
	if o != nil && !IsNil(o.IsWriteIndex) {
		return true
	}

	return false
}

// SetIsWriteIndex gets a reference to the given bool and assigns it to the IsWriteIndex field.
func (o *UserDefinedStructure) SetIsWriteIndex(v bool) {
	o.IsWriteIndex = &v
}

// GetMustExist returns the MustExist field value if set, zero value otherwise.
func (o *UserDefinedStructure) GetMustExist() string {
	if o == nil || IsNil(o.MustExist) {
		var ret string
		return ret
	}
	return *o.MustExist
}

// GetMustExistOk returns a tuple with the MustExist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedStructure) GetMustExistOk() (*string, bool) {
	if o == nil || IsNil(o.MustExist) {
		return nil, false
	}
	return o.MustExist, true
}

// HasMustExist returns a boolean if a field has been set.
func (o *UserDefinedStructure) HasMustExist() bool {
	if o != nil && !IsNil(o.MustExist) {
		return true
	}

	return false
}

// SetMustExist gets a reference to the given string and assigns it to the MustExist field.
func (o *UserDefinedStructure) SetMustExist(v string) {
	o.MustExist = &v
}

// GetRouting returns the Routing field value if set, zero value otherwise.
func (o *UserDefinedStructure) GetRouting() string {
	if o == nil || IsNil(o.Routing) {
		var ret string
		return ret
	}
	return *o.Routing
}

// GetRoutingOk returns a tuple with the Routing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedStructure) GetRoutingOk() (*string, bool) {
	if o == nil || IsNil(o.Routing) {
		return nil, false
	}
	return o.Routing, true
}

// HasRouting returns a boolean if a field has been set.
func (o *UserDefinedStructure) HasRouting() bool {
	if o != nil && !IsNil(o.Routing) {
		return true
	}

	return false
}

// SetRouting gets a reference to the given string and assigns it to the Routing field.
func (o *UserDefinedStructure) SetRouting(v string) {
	o.Routing = &v
}

// GetSearchRouting returns the SearchRouting field value if set, zero value otherwise.
func (o *UserDefinedStructure) GetSearchRouting() string {
	if o == nil || IsNil(o.SearchRouting) {
		var ret string
		return ret
	}
	return *o.SearchRouting
}

// GetSearchRoutingOk returns a tuple with the SearchRouting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedStructure) GetSearchRoutingOk() (*string, bool) {
	if o == nil || IsNil(o.SearchRouting) {
		return nil, false
	}
	return o.SearchRouting, true
}

// HasSearchRouting returns a boolean if a field has been set.
func (o *UserDefinedStructure) HasSearchRouting() bool {
	if o != nil && !IsNil(o.SearchRouting) {
		return true
	}

	return false
}

// SetSearchRouting gets a reference to the given string and assigns it to the SearchRouting field.
func (o *UserDefinedStructure) SetSearchRouting(v string) {
	o.SearchRouting = &v
}

func (o UserDefinedStructure) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserDefinedStructure) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Alias) {
		toSerialize["alias"] = o.Alias
	}
	if !IsNil(o.Aliases) {
		toSerialize["aliases"] = o.Aliases
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if !IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	if !IsNil(o.Indices) {
		toSerialize["indices"] = o.Indices
	}
	if !IsNil(o.IndexRouting) {
		toSerialize["index_routing"] = o.IndexRouting
	}
	if !IsNil(o.IsHidden) {
		toSerialize["is_hidden"] = o.IsHidden
	}
	if !IsNil(o.IsWriteIndex) {
		toSerialize["is_write_index"] = o.IsWriteIndex
	}
	if !IsNil(o.MustExist) {
		toSerialize["must_exist"] = o.MustExist
	}
	if !IsNil(o.Routing) {
		toSerialize["routing"] = o.Routing
	}
	if !IsNil(o.SearchRouting) {
		toSerialize["search_routing"] = o.SearchRouting
	}
	return toSerialize, nil
}

type NullableUserDefinedStructure struct {
	value *UserDefinedStructure
	isSet bool
}

func (v NullableUserDefinedStructure) Get() *UserDefinedStructure {
	return v.value
}

func (v *NullableUserDefinedStructure) Set(val *UserDefinedStructure) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDefinedStructure) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDefinedStructure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDefinedStructure(val *UserDefinedStructure) *NullableUserDefinedStructure {
	return &NullableUserDefinedStructure{value: val, isSet: true}
}

func (v NullableUserDefinedStructure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDefinedStructure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

