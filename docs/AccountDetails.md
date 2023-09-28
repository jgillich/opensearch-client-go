# AccountDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserName** | Pointer to **string** |  | [optional] 
**IsReserved** | Pointer to **bool** |  | [optional] 
**IsHidden** | Pointer to **bool** |  | [optional] 
**IsInternalUser** | Pointer to **bool** |  | [optional] 
**UserRequestedTenant** | Pointer to **string** |  | [optional] 
**BackendRoles** | Pointer to **[]string** |  | [optional] 
**CustomAttributeNames** | Pointer to **[]string** |  | [optional] 
**Tenants** | Pointer to [**UserTenants**](UserTenants.md) |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAccountDetails

`func NewAccountDetails() *AccountDetails`

NewAccountDetails instantiates a new AccountDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountDetailsWithDefaults

`func NewAccountDetailsWithDefaults() *AccountDetails`

NewAccountDetailsWithDefaults instantiates a new AccountDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserName

`func (o *AccountDetails) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *AccountDetails) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *AccountDetails) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *AccountDetails) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetIsReserved

`func (o *AccountDetails) GetIsReserved() bool`

GetIsReserved returns the IsReserved field if non-nil, zero value otherwise.

### GetIsReservedOk

`func (o *AccountDetails) GetIsReservedOk() (*bool, bool)`

GetIsReservedOk returns a tuple with the IsReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReserved

`func (o *AccountDetails) SetIsReserved(v bool)`

SetIsReserved sets IsReserved field to given value.

### HasIsReserved

`func (o *AccountDetails) HasIsReserved() bool`

HasIsReserved returns a boolean if a field has been set.

### GetIsHidden

`func (o *AccountDetails) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *AccountDetails) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *AccountDetails) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *AccountDetails) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetIsInternalUser

`func (o *AccountDetails) GetIsInternalUser() bool`

GetIsInternalUser returns the IsInternalUser field if non-nil, zero value otherwise.

### GetIsInternalUserOk

`func (o *AccountDetails) GetIsInternalUserOk() (*bool, bool)`

GetIsInternalUserOk returns a tuple with the IsInternalUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInternalUser

`func (o *AccountDetails) SetIsInternalUser(v bool)`

SetIsInternalUser sets IsInternalUser field to given value.

### HasIsInternalUser

`func (o *AccountDetails) HasIsInternalUser() bool`

HasIsInternalUser returns a boolean if a field has been set.

### GetUserRequestedTenant

`func (o *AccountDetails) GetUserRequestedTenant() string`

GetUserRequestedTenant returns the UserRequestedTenant field if non-nil, zero value otherwise.

### GetUserRequestedTenantOk

`func (o *AccountDetails) GetUserRequestedTenantOk() (*string, bool)`

GetUserRequestedTenantOk returns a tuple with the UserRequestedTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRequestedTenant

`func (o *AccountDetails) SetUserRequestedTenant(v string)`

SetUserRequestedTenant sets UserRequestedTenant field to given value.

### HasUserRequestedTenant

`func (o *AccountDetails) HasUserRequestedTenant() bool`

HasUserRequestedTenant returns a boolean if a field has been set.

### GetBackendRoles

`func (o *AccountDetails) GetBackendRoles() []string`

GetBackendRoles returns the BackendRoles field if non-nil, zero value otherwise.

### GetBackendRolesOk

`func (o *AccountDetails) GetBackendRolesOk() (*[]string, bool)`

GetBackendRolesOk returns a tuple with the BackendRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendRoles

`func (o *AccountDetails) SetBackendRoles(v []string)`

SetBackendRoles sets BackendRoles field to given value.

### HasBackendRoles

`func (o *AccountDetails) HasBackendRoles() bool`

HasBackendRoles returns a boolean if a field has been set.

### GetCustomAttributeNames

`func (o *AccountDetails) GetCustomAttributeNames() []string`

GetCustomAttributeNames returns the CustomAttributeNames field if non-nil, zero value otherwise.

### GetCustomAttributeNamesOk

`func (o *AccountDetails) GetCustomAttributeNamesOk() (*[]string, bool)`

GetCustomAttributeNamesOk returns a tuple with the CustomAttributeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributeNames

`func (o *AccountDetails) SetCustomAttributeNames(v []string)`

SetCustomAttributeNames sets CustomAttributeNames field to given value.

### HasCustomAttributeNames

`func (o *AccountDetails) HasCustomAttributeNames() bool`

HasCustomAttributeNames returns a boolean if a field has been set.

### GetTenants

`func (o *AccountDetails) GetTenants() UserTenants`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *AccountDetails) GetTenantsOk() (*UserTenants, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *AccountDetails) SetTenants(v UserTenants)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *AccountDetails) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetRoles

`func (o *AccountDetails) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *AccountDetails) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *AccountDetails) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *AccountDetails) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


