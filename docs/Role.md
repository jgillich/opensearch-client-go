# Role

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reserved** | Pointer to **bool** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ClusterPermissions** | Pointer to **[]string** |  | [optional] 
**IndexPermissions** | Pointer to [**[]IndexPermission**](IndexPermission.md) |  | [optional] 
**TenantPermissions** | Pointer to [**[]TenantPermission**](TenantPermission.md) |  | [optional] 
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

### GetClusterPermissions

`func (o *Role) GetClusterPermissions() []string`

GetClusterPermissions returns the ClusterPermissions field if non-nil, zero value otherwise.

### GetClusterPermissionsOk

`func (o *Role) GetClusterPermissionsOk() (*[]string, bool)`

GetClusterPermissionsOk returns a tuple with the ClusterPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterPermissions

`func (o *Role) SetClusterPermissions(v []string)`

SetClusterPermissions sets ClusterPermissions field to given value.

### HasClusterPermissions

`func (o *Role) HasClusterPermissions() bool`

HasClusterPermissions returns a boolean if a field has been set.

### GetIndexPermissions

`func (o *Role) GetIndexPermissions() []IndexPermission`

GetIndexPermissions returns the IndexPermissions field if non-nil, zero value otherwise.

### GetIndexPermissionsOk

`func (o *Role) GetIndexPermissionsOk() (*[]IndexPermission, bool)`

GetIndexPermissionsOk returns a tuple with the IndexPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexPermissions

`func (o *Role) SetIndexPermissions(v []IndexPermission)`

SetIndexPermissions sets IndexPermissions field to given value.

### HasIndexPermissions

`func (o *Role) HasIndexPermissions() bool`

HasIndexPermissions returns a boolean if a field has been set.

### GetTenantPermissions

`func (o *Role) GetTenantPermissions() []TenantPermission`

GetTenantPermissions returns the TenantPermissions field if non-nil, zero value otherwise.

### GetTenantPermissionsOk

`func (o *Role) GetTenantPermissionsOk() (*[]TenantPermission, bool)`

GetTenantPermissionsOk returns a tuple with the TenantPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantPermissions

`func (o *Role) SetTenantPermissions(v []TenantPermission)`

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


