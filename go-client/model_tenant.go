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

// checks if the Tenant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Tenant{}

// Tenant struct for Tenant
type Tenant struct {
	Reserved *bool `json:"reserved,omitempty"`
	Hidden *bool `json:"hidden,omitempty"`
	Description *string `json:"description,omitempty"`
	Static *bool `json:"static,omitempty"`
}

// NewTenant instantiates a new Tenant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenant() *Tenant {
	this := Tenant{}
	return &this
}

// NewTenantWithDefaults instantiates a new Tenant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantWithDefaults() *Tenant {
	this := Tenant{}
	return &this
}

// GetReserved returns the Reserved field value if set, zero value otherwise.
func (o *Tenant) GetReserved() bool {
	if o == nil || IsNil(o.Reserved) {
		var ret bool
		return ret
	}
	return *o.Reserved
}

// GetReservedOk returns a tuple with the Reserved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetReservedOk() (*bool, bool) {
	if o == nil || IsNil(o.Reserved) {
		return nil, false
	}
	return o.Reserved, true
}

// HasReserved returns a boolean if a field has been set.
func (o *Tenant) HasReserved() bool {
	if o != nil && !IsNil(o.Reserved) {
		return true
	}

	return false
}

// SetReserved gets a reference to the given bool and assigns it to the Reserved field.
func (o *Tenant) SetReserved(v bool) {
	o.Reserved = &v
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *Tenant) GetHidden() bool {
	if o == nil || IsNil(o.Hidden) {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetHiddenOk() (*bool, bool) {
	if o == nil || IsNil(o.Hidden) {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *Tenant) HasHidden() bool {
	if o != nil && !IsNil(o.Hidden) {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *Tenant) SetHidden(v bool) {
	o.Hidden = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Tenant) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Tenant) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Tenant) SetDescription(v string) {
	o.Description = &v
}

// GetStatic returns the Static field value if set, zero value otherwise.
func (o *Tenant) GetStatic() bool {
	if o == nil || IsNil(o.Static) {
		var ret bool
		return ret
	}
	return *o.Static
}

// GetStaticOk returns a tuple with the Static field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetStaticOk() (*bool, bool) {
	if o == nil || IsNil(o.Static) {
		return nil, false
	}
	return o.Static, true
}

// HasStatic returns a boolean if a field has been set.
func (o *Tenant) HasStatic() bool {
	if o != nil && !IsNil(o.Static) {
		return true
	}

	return false
}

// SetStatic gets a reference to the given bool and assigns it to the Static field.
func (o *Tenant) SetStatic(v bool) {
	o.Static = &v
}

func (o Tenant) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Tenant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Reserved) {
		toSerialize["reserved"] = o.Reserved
	}
	if !IsNil(o.Hidden) {
		toSerialize["hidden"] = o.Hidden
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Static) {
		toSerialize["static"] = o.Static
	}
	return toSerialize, nil
}

type NullableTenant struct {
	value *Tenant
	isSet bool
}

func (v NullableTenant) Get() *Tenant {
	return v.value
}

func (v *NullableTenant) Set(val *Tenant) {
	v.value = val
	v.isSet = true
}

func (v NullableTenant) IsSet() bool {
	return v.isSet
}

func (v *NullableTenant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenant(val *Tenant) *NullableTenant {
	return &NullableTenant{value: val, isSet: true}
}

func (v NullableTenant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


