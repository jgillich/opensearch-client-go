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

// checks if the AccountDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountDetails{}

// AccountDetails struct for AccountDetails
type AccountDetails struct {
	UserName *string `json:"user_name,omitempty"`
	IsReserved *bool `json:"is_reserved,omitempty"`
	IsHidden *bool `json:"is_hidden,omitempty"`
	IsInternalUser *bool `json:"is_internal_user,omitempty"`
	UserRequestedTenant *string `json:"user_requested_tenant,omitempty"`
	BackendRoles []string `json:"backend_roles,omitempty"`
	CustomAttributeNames []string `json:"custom_attribute_names,omitempty"`
	Tenants *UserTenants `json:"tenants,omitempty"`
	Roles []string `json:"roles,omitempty"`
}

// NewAccountDetails instantiates a new AccountDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountDetails() *AccountDetails {
	this := AccountDetails{}
	return &this
}

// NewAccountDetailsWithDefaults instantiates a new AccountDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountDetailsWithDefaults() *AccountDetails {
	this := AccountDetails{}
	return &this
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *AccountDetails) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetails) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *AccountDetails) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *AccountDetails) SetUserName(v string) {
	o.UserName = &v
}

// GetIsReserved returns the IsReserved field value if set, zero value otherwise.
func (o *AccountDetails) GetIsReserved() bool {
	if o == nil || IsNil(o.IsReserved) {
		var ret bool
		return ret
	}
	return *o.IsReserved
}

// GetIsReservedOk returns a tuple with the IsReserved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetails) GetIsReservedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsReserved) {
		return nil, false
	}
	return o.IsReserved, true
}

// HasIsReserved returns a boolean if a field has been set.
func (o *AccountDetails) HasIsReserved() bool {
	if o != nil && !IsNil(o.IsReserved) {
		return true
	}

	return false
}

// SetIsReserved gets a reference to the given bool and assigns it to the IsReserved field.
func (o *AccountDetails) SetIsReserved(v bool) {
	o.IsReserved = &v
}

// GetIsHidden returns the IsHidden field value if set, zero value otherwise.
func (o *AccountDetails) GetIsHidden() bool {
	if o == nil || IsNil(o.IsHidden) {
		var ret bool
		return ret
	}
	return *o.IsHidden
}

// GetIsHiddenOk returns a tuple with the IsHidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetails) GetIsHiddenOk() (*bool, bool) {
	if o == nil || IsNil(o.IsHidden) {
		return nil, false
	}
	return o.IsHidden, true
}

// HasIsHidden returns a boolean if a field has been set.
func (o *AccountDetails) HasIsHidden() bool {
	if o != nil && !IsNil(o.IsHidden) {
		return true
	}

	return false
}

// SetIsHidden gets a reference to the given bool and assigns it to the IsHidden field.
func (o *AccountDetails) SetIsHidden(v bool) {
	o.IsHidden = &v
}

// GetIsInternalUser returns the IsInternalUser field value if set, zero value otherwise.
func (o *AccountDetails) GetIsInternalUser() bool {
	if o == nil || IsNil(o.IsInternalUser) {
		var ret bool
		return ret
	}
	return *o.IsInternalUser
}

// GetIsInternalUserOk returns a tuple with the IsInternalUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetails) GetIsInternalUserOk() (*bool, bool) {
	if o == nil || IsNil(o.IsInternalUser) {
		return nil, false
	}
	return o.IsInternalUser, true
}

// HasIsInternalUser returns a boolean if a field has been set.
func (o *AccountDetails) HasIsInternalUser() bool {
	if o != nil && !IsNil(o.IsInternalUser) {
		return true
	}

	return false
}

// SetIsInternalUser gets a reference to the given bool and assigns it to the IsInternalUser field.
func (o *AccountDetails) SetIsInternalUser(v bool) {
	o.IsInternalUser = &v
}

// GetUserRequestedTenant returns the UserRequestedTenant field value if set, zero value otherwise.
func (o *AccountDetails) GetUserRequestedTenant() string {
	if o == nil || IsNil(o.UserRequestedTenant) {
		var ret string
		return ret
	}
	return *o.UserRequestedTenant
}

// GetUserRequestedTenantOk returns a tuple with the UserRequestedTenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetails) GetUserRequestedTenantOk() (*string, bool) {
	if o == nil || IsNil(o.UserRequestedTenant) {
		return nil, false
	}
	return o.UserRequestedTenant, true
}

// HasUserRequestedTenant returns a boolean if a field has been set.
func (o *AccountDetails) HasUserRequestedTenant() bool {
	if o != nil && !IsNil(o.UserRequestedTenant) {
		return true
	}

	return false
}

// SetUserRequestedTenant gets a reference to the given string and assigns it to the UserRequestedTenant field.
func (o *AccountDetails) SetUserRequestedTenant(v string) {
	o.UserRequestedTenant = &v
}

// GetBackendRoles returns the BackendRoles field value if set, zero value otherwise.
func (o *AccountDetails) GetBackendRoles() []string {
	if o == nil || IsNil(o.BackendRoles) {
		var ret []string
		return ret
	}
	return o.BackendRoles
}

// GetBackendRolesOk returns a tuple with the BackendRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetails) GetBackendRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.BackendRoles) {
		return nil, false
	}
	return o.BackendRoles, true
}

// HasBackendRoles returns a boolean if a field has been set.
func (o *AccountDetails) HasBackendRoles() bool {
	if o != nil && !IsNil(o.BackendRoles) {
		return true
	}

	return false
}

// SetBackendRoles gets a reference to the given []string and assigns it to the BackendRoles field.
func (o *AccountDetails) SetBackendRoles(v []string) {
	o.BackendRoles = v
}

// GetCustomAttributeNames returns the CustomAttributeNames field value if set, zero value otherwise.
func (o *AccountDetails) GetCustomAttributeNames() []string {
	if o == nil || IsNil(o.CustomAttributeNames) {
		var ret []string
		return ret
	}
	return o.CustomAttributeNames
}

// GetCustomAttributeNamesOk returns a tuple with the CustomAttributeNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetails) GetCustomAttributeNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.CustomAttributeNames) {
		return nil, false
	}
	return o.CustomAttributeNames, true
}

// HasCustomAttributeNames returns a boolean if a field has been set.
func (o *AccountDetails) HasCustomAttributeNames() bool {
	if o != nil && !IsNil(o.CustomAttributeNames) {
		return true
	}

	return false
}

// SetCustomAttributeNames gets a reference to the given []string and assigns it to the CustomAttributeNames field.
func (o *AccountDetails) SetCustomAttributeNames(v []string) {
	o.CustomAttributeNames = v
}

// GetTenants returns the Tenants field value if set, zero value otherwise.
func (o *AccountDetails) GetTenants() UserTenants {
	if o == nil || IsNil(o.Tenants) {
		var ret UserTenants
		return ret
	}
	return *o.Tenants
}

// GetTenantsOk returns a tuple with the Tenants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetails) GetTenantsOk() (*UserTenants, bool) {
	if o == nil || IsNil(o.Tenants) {
		return nil, false
	}
	return o.Tenants, true
}

// HasTenants returns a boolean if a field has been set.
func (o *AccountDetails) HasTenants() bool {
	if o != nil && !IsNil(o.Tenants) {
		return true
	}

	return false
}

// SetTenants gets a reference to the given UserTenants and assigns it to the Tenants field.
func (o *AccountDetails) SetTenants(v UserTenants) {
	o.Tenants = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *AccountDetails) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetails) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *AccountDetails) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *AccountDetails) SetRoles(v []string) {
	o.Roles = v
}

func (o AccountDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserName) {
		toSerialize["user_name"] = o.UserName
	}
	if !IsNil(o.IsReserved) {
		toSerialize["is_reserved"] = o.IsReserved
	}
	if !IsNil(o.IsHidden) {
		toSerialize["is_hidden"] = o.IsHidden
	}
	if !IsNil(o.IsInternalUser) {
		toSerialize["is_internal_user"] = o.IsInternalUser
	}
	if !IsNil(o.UserRequestedTenant) {
		toSerialize["user_requested_tenant"] = o.UserRequestedTenant
	}
	if !IsNil(o.BackendRoles) {
		toSerialize["backend_roles"] = o.BackendRoles
	}
	if !IsNil(o.CustomAttributeNames) {
		toSerialize["custom_attribute_names"] = o.CustomAttributeNames
	}
	if !IsNil(o.Tenants) {
		toSerialize["tenants"] = o.Tenants
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}

type NullableAccountDetails struct {
	value *AccountDetails
	isSet bool
}

func (v NullableAccountDetails) Get() *AccountDetails {
	return v.value
}

func (v *NullableAccountDetails) Set(val *AccountDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountDetails(val *AccountDetails) *NullableAccountDetails {
	return &NullableAccountDetails{value: val, isSet: true}
}

func (v NullableAccountDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


