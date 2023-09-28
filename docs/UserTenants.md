# UserTenants

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalTenant** | Pointer to **bool** |  | [optional] 
**AdminTenant** | Pointer to **bool** |  | [optional] 
**Admin** | Pointer to **bool** |  | [optional] 

## Methods

### NewUserTenants

`func NewUserTenants() *UserTenants`

NewUserTenants instantiates a new UserTenants object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserTenantsWithDefaults

`func NewUserTenantsWithDefaults() *UserTenants`

NewUserTenantsWithDefaults instantiates a new UserTenants object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobalTenant

`func (o *UserTenants) GetGlobalTenant() bool`

GetGlobalTenant returns the GlobalTenant field if non-nil, zero value otherwise.

### GetGlobalTenantOk

`func (o *UserTenants) GetGlobalTenantOk() (*bool, bool)`

GetGlobalTenantOk returns a tuple with the GlobalTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalTenant

`func (o *UserTenants) SetGlobalTenant(v bool)`

SetGlobalTenant sets GlobalTenant field to given value.

### HasGlobalTenant

`func (o *UserTenants) HasGlobalTenant() bool`

HasGlobalTenant returns a boolean if a field has been set.

### GetAdminTenant

`func (o *UserTenants) GetAdminTenant() bool`

GetAdminTenant returns the AdminTenant field if non-nil, zero value otherwise.

### GetAdminTenantOk

`func (o *UserTenants) GetAdminTenantOk() (*bool, bool)`

GetAdminTenantOk returns a tuple with the AdminTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminTenant

`func (o *UserTenants) SetAdminTenant(v bool)`

SetAdminTenant sets AdminTenant field to given value.

### HasAdminTenant

`func (o *UserTenants) HasAdminTenant() bool`

HasAdminTenant returns a boolean if a field has been set.

### GetAdmin

`func (o *UserTenants) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *UserTenants) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *UserTenants) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *UserTenants) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


