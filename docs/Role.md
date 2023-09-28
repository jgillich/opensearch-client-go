# Role

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reserved** | Pointer to **bool** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ClusterPermission** | Pointer to **[]string** |  | [optional] 
**IndexPermission** | Pointer to [**IndexPermission**](IndexPermission.md) |  | [optional] 
**TenantPermissions** | Pointer to **[]string** |  | [optional] 
**Static** | Pointer to **bool** |  | [optional] 

## Methods

### NewRole

`func NewRole() *Role`

NewRole instantiates a new Role object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleWithDefaults

`func NewRoleWithDefaults() *Role`

NewRoleWithDefaults instantiates a new Role object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReserved

`func (o *Role) GetReserved() bool`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *Role) GetReservedOk() (*bool, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *Role) SetReserved(v bool)`

SetReserved sets Reserved field to given value.

### HasReserved

`func (o *Role) HasReserved() bool`

HasReserved returns a boolean if a field has been set.

### GetHidden

`func (o *Role) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *Role) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *Role) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *Role) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetDescription

`func (o *Role) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Role) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Role) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Role) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetClusterPermission

`func (o *Role) GetClusterPermission() []string`

GetClusterPermission returns the ClusterPermission field if non-nil, zero value otherwise.

### GetClusterPermissionOk

`func (o *Role) GetClusterPermissionOk() (*[]string, bool)`

GetClusterPermissionOk returns a tuple with the ClusterPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterPermission

`func (o *Role) SetClusterPermission(v []string)`

SetClusterPermission sets ClusterPermission field to given value.

### HasClusterPermission

`func (o *Role) HasClusterPermission() bool`

HasClusterPermission returns a boolean if a field has been set.

### GetIndexPermission

`func (o *Role) GetIndexPermission() IndexPermission`

GetIndexPermission returns the IndexPermission field if non-nil, zero value otherwise.

### GetIndexPermissionOk

`func (o *Role) GetIndexPermissionOk() (*IndexPermission, bool)`

GetIndexPermissionOk returns a tuple with the IndexPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexPermission

`func (o *Role) SetIndexPermission(v IndexPermission)`

SetIndexPermission sets IndexPermission field to given value.

### HasIndexPermission

`func (o *Role) HasIndexPermission() bool`

HasIndexPermission returns a boolean if a field has been set.

### GetTenantPermissions

`func (o *Role) GetTenantPermissions() []string`

GetTenantPermissions returns the TenantPermissions field if non-nil, zero value otherwise.

### GetTenantPermissionsOk

`func (o *Role) GetTenantPermissionsOk() (*[]string, bool)`

GetTenantPermissionsOk returns a tuple with the TenantPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantPermissions

`func (o *Role) SetTenantPermissions(v []string)`

SetTenantPermissions sets TenantPermissions field to given value.

### HasTenantPermissions

`func (o *Role) HasTenantPermissions() bool`

HasTenantPermissions returns a boolean if a field has been set.

### GetStatic

`func (o *Role) GetStatic() bool`

GetStatic returns the Static field if non-nil, zero value otherwise.

### GetStaticOk

`func (o *Role) GetStaticOk() (*bool, bool)`

GetStaticOk returns a tuple with the Static field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic

`func (o *Role) SetStatic(v bool)`

SetStatic sets Static field to given value.

### HasStatic

`func (o *Role) HasStatic() bool`

HasStatic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


