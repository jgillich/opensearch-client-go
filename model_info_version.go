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

// checks if the InfoVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InfoVersion{}

// InfoVersion struct for InfoVersion
type InfoVersion struct {
	Distribution *string `json:"distribution,omitempty"`
	Number *string `json:"number,omitempty"`
	BuildType *string `json:"build_type,omitempty"`
	BuildHash *string `json:"build_hash,omitempty"`
	BuildDate *string `json:"build_date,omitempty"`
	BuildSnapshot *bool `json:"build_snapshot,omitempty"`
	LuceneVersion *string `json:"lucene_version,omitempty"`
	MinimumWireCompatibilityVersion *string `json:"minimum_wire_compatibility_version,omitempty"`
	MinimumIndexCompatibilityVersion *string `json:"minimum_index_compatibility_version,omitempty"`
}

// NewInfoVersion instantiates a new InfoVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInfoVersion() *InfoVersion {
	this := InfoVersion{}
	return &this
}

// NewInfoVersionWithDefaults instantiates a new InfoVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInfoVersionWithDefaults() *InfoVersion {
	this := InfoVersion{}
	return &this
}

// GetDistribution returns the Distribution field value if set, zero value otherwise.
func (o *InfoVersion) GetDistribution() string {
	if o == nil || IsNil(o.Distribution) {
		var ret string
		return ret
	}
	return *o.Distribution
}

// GetDistributionOk returns a tuple with the Distribution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfoVersion) GetDistributionOk() (*string, bool) {
	if o == nil || IsNil(o.Distribution) {
		return nil, false
	}
	return o.Distribution, true
}

// HasDistribution returns a boolean if a field has been set.
func (o *InfoVersion) HasDistribution() bool {
	if o != nil && !IsNil(o.Distribution) {
		return true
	}

	return false
}

// SetDistribution gets a reference to the given string and assigns it to the Distribution field.
func (o *InfoVersion) SetDistribution(v string) {
	o.Distribution = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *InfoVersion) GetNumber() string {
	if o == nil || IsNil(o.Number) {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfoVersion) GetNumberOk() (*string, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *InfoVersion) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *InfoVersion) SetNumber(v string) {
	o.Number = &v
}

// GetBuildType returns the BuildType field value if set, zero value otherwise.
func (o *InfoVersion) GetBuildType() string {
	if o == nil || IsNil(o.BuildType) {
		var ret string
		return ret
	}
	return *o.BuildType
}

// GetBuildTypeOk returns a tuple with the BuildType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfoVersion) GetBuildTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BuildType) {
		return nil, false
	}
	return o.BuildType, true
}

// HasBuildType returns a boolean if a field has been set.
func (o *InfoVersion) HasBuildType() bool {
	if o != nil && !IsNil(o.BuildType) {
		return true
	}

	return false
}

// SetBuildType gets a reference to the given string and assigns it to the BuildType field.
func (o *InfoVersion) SetBuildType(v string) {
	o.BuildType = &v
}

// GetBuildHash returns the BuildHash field value if set, zero value otherwise.
func (o *InfoVersion) GetBuildHash() string {
	if o == nil || IsNil(o.BuildHash) {
		var ret string
		return ret
	}
	return *o.BuildHash
}

// GetBuildHashOk returns a tuple with the BuildHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfoVersion) GetBuildHashOk() (*string, bool) {
	if o == nil || IsNil(o.BuildHash) {
		return nil, false
	}
	return o.BuildHash, true
}

// HasBuildHash returns a boolean if a field has been set.
func (o *InfoVersion) HasBuildHash() bool {
	if o != nil && !IsNil(o.BuildHash) {
		return true
	}

	return false
}

// SetBuildHash gets a reference to the given string and assigns it to the BuildHash field.
func (o *InfoVersion) SetBuildHash(v string) {
	o.BuildHash = &v
}

// GetBuildDate returns the BuildDate field value if set, zero value otherwise.
func (o *InfoVersion) GetBuildDate() string {
	if o == nil || IsNil(o.BuildDate) {
		var ret string
		return ret
	}
	return *o.BuildDate
}

// GetBuildDateOk returns a tuple with the BuildDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfoVersion) GetBuildDateOk() (*string, bool) {
	if o == nil || IsNil(o.BuildDate) {
		return nil, false
	}
	return o.BuildDate, true
}

// HasBuildDate returns a boolean if a field has been set.
func (o *InfoVersion) HasBuildDate() bool {
	if o != nil && !IsNil(o.BuildDate) {
		return true
	}

	return false
}

// SetBuildDate gets a reference to the given string and assigns it to the BuildDate field.
func (o *InfoVersion) SetBuildDate(v string) {
	o.BuildDate = &v
}

// GetBuildSnapshot returns the BuildSnapshot field value if set, zero value otherwise.
func (o *InfoVersion) GetBuildSnapshot() bool {
	if o == nil || IsNil(o.BuildSnapshot) {
		var ret bool
		return ret
	}
	return *o.BuildSnapshot
}

// GetBuildSnapshotOk returns a tuple with the BuildSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfoVersion) GetBuildSnapshotOk() (*bool, bool) {
	if o == nil || IsNil(o.BuildSnapshot) {
		return nil, false
	}
	return o.BuildSnapshot, true
}

// HasBuildSnapshot returns a boolean if a field has been set.
func (o *InfoVersion) HasBuildSnapshot() bool {
	if o != nil && !IsNil(o.BuildSnapshot) {
		return true
	}

	return false
}

// SetBuildSnapshot gets a reference to the given bool and assigns it to the BuildSnapshot field.
func (o *InfoVersion) SetBuildSnapshot(v bool) {
	o.BuildSnapshot = &v
}

// GetLuceneVersion returns the LuceneVersion field value if set, zero value otherwise.
func (o *InfoVersion) GetLuceneVersion() string {
	if o == nil || IsNil(o.LuceneVersion) {
		var ret string
		return ret
	}
	return *o.LuceneVersion
}

// GetLuceneVersionOk returns a tuple with the LuceneVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfoVersion) GetLuceneVersionOk() (*string, bool) {
	if o == nil || IsNil(o.LuceneVersion) {
		return nil, false
	}
	return o.LuceneVersion, true
}

// HasLuceneVersion returns a boolean if a field has been set.
func (o *InfoVersion) HasLuceneVersion() bool {
	if o != nil && !IsNil(o.LuceneVersion) {
		return true
	}

	return false
}

// SetLuceneVersion gets a reference to the given string and assigns it to the LuceneVersion field.
func (o *InfoVersion) SetLuceneVersion(v string) {
	o.LuceneVersion = &v
}

// GetMinimumWireCompatibilityVersion returns the MinimumWireCompatibilityVersion field value if set, zero value otherwise.
func (o *InfoVersion) GetMinimumWireCompatibilityVersion() string {
	if o == nil || IsNil(o.MinimumWireCompatibilityVersion) {
		var ret string
		return ret
	}
	return *o.MinimumWireCompatibilityVersion
}

// GetMinimumWireCompatibilityVersionOk returns a tuple with the MinimumWireCompatibilityVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfoVersion) GetMinimumWireCompatibilityVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MinimumWireCompatibilityVersion) {
		return nil, false
	}
	return o.MinimumWireCompatibilityVersion, true
}

// HasMinimumWireCompatibilityVersion returns a boolean if a field has been set.
func (o *InfoVersion) HasMinimumWireCompatibilityVersion() bool {
	if o != nil && !IsNil(o.MinimumWireCompatibilityVersion) {
		return true
	}

	return false
}

// SetMinimumWireCompatibilityVersion gets a reference to the given string and assigns it to the MinimumWireCompatibilityVersion field.
func (o *InfoVersion) SetMinimumWireCompatibilityVersion(v string) {
	o.MinimumWireCompatibilityVersion = &v
}

// GetMinimumIndexCompatibilityVersion returns the MinimumIndexCompatibilityVersion field value if set, zero value otherwise.
func (o *InfoVersion) GetMinimumIndexCompatibilityVersion() string {
	if o == nil || IsNil(o.MinimumIndexCompatibilityVersion) {
		var ret string
		return ret
	}
	return *o.MinimumIndexCompatibilityVersion
}

// GetMinimumIndexCompatibilityVersionOk returns a tuple with the MinimumIndexCompatibilityVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfoVersion) GetMinimumIndexCompatibilityVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MinimumIndexCompatibilityVersion) {
		return nil, false
	}
	return o.MinimumIndexCompatibilityVersion, true
}

// HasMinimumIndexCompatibilityVersion returns a boolean if a field has been set.
func (o *InfoVersion) HasMinimumIndexCompatibilityVersion() bool {
	if o != nil && !IsNil(o.MinimumIndexCompatibilityVersion) {
		return true
	}

	return false
}

// SetMinimumIndexCompatibilityVersion gets a reference to the given string and assigns it to the MinimumIndexCompatibilityVersion field.
func (o *InfoVersion) SetMinimumIndexCompatibilityVersion(v string) {
	o.MinimumIndexCompatibilityVersion = &v
}

func (o InfoVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InfoVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Distribution) {
		toSerialize["distribution"] = o.Distribution
	}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !IsNil(o.BuildType) {
		toSerialize["build_type"] = o.BuildType
	}
	if !IsNil(o.BuildHash) {
		toSerialize["build_hash"] = o.BuildHash
	}
	if !IsNil(o.BuildDate) {
		toSerialize["build_date"] = o.BuildDate
	}
	if !IsNil(o.BuildSnapshot) {
		toSerialize["build_snapshot"] = o.BuildSnapshot
	}
	if !IsNil(o.LuceneVersion) {
		toSerialize["lucene_version"] = o.LuceneVersion
	}
	if !IsNil(o.MinimumWireCompatibilityVersion) {
		toSerialize["minimum_wire_compatibility_version"] = o.MinimumWireCompatibilityVersion
	}
	if !IsNil(o.MinimumIndexCompatibilityVersion) {
		toSerialize["minimum_index_compatibility_version"] = o.MinimumIndexCompatibilityVersion
	}
	return toSerialize, nil
}

type NullableInfoVersion struct {
	value *InfoVersion
	isSet bool
}

func (v NullableInfoVersion) Get() *InfoVersion {
	return v.value
}

func (v *NullableInfoVersion) Set(val *InfoVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableInfoVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableInfoVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInfoVersion(val *InfoVersion) *NullableInfoVersion {
	return &NullableInfoVersion{value: val, isSet: true}
}

func (v NullableInfoVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInfoVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

