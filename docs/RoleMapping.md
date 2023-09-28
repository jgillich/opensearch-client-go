# RoleMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hosts** | Pointer to **[]string** |  | [optional] 
**Users** | Pointer to **[]string** |  | [optional] 
**Reserved** | Pointer to **bool** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**BackendRoles** | Pointer to **[]string** |  | [optional] 
**AndBackendRoles** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewRoleMapping

`func NewRoleMapping() *RoleMapping`

NewRoleMapping instantiates a new RoleMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleMappingWithDefaults

`func NewRoleMappingWithDefaults() *RoleMapping`

NewRoleMappingWithDefaults instantiates a new RoleMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHosts

`func (o *RoleMapping) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *RoleMapping) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *RoleMapping) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *RoleMapping) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetUsers

`func (o *RoleMapping) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *RoleMapping) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *RoleMapping) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *RoleMapping) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetReserved

`func (o *RoleMapping) GetReserved() bool`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *RoleMapping) GetReservedOk() (*bool, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *RoleMapping) SetReserved(v bool)`

SetReserved sets Reserved field to given value.

### HasReserved

`func (o *RoleMapping) HasReserved() bool`

HasReserved returns a boolean if a field has been set.

### GetHidden

`func (o *RoleMapping) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *RoleMapping) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *RoleMapping) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *RoleMapping) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetBackendRoles

`func (o *RoleMapping) GetBackendRoles() []string`

GetBackendRoles returns the BackendRoles field if non-nil, zero value otherwise.

### GetBackendRolesOk

`func (o *RoleMapping) GetBackendRolesOk() (*[]string, bool)`

GetBackendRolesOk returns a tuple with the BackendRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendRoles

`func (o *RoleMapping) SetBackendRoles(v []string)`

SetBackendRoles sets BackendRoles field to given value.

### HasBackendRoles

`func (o *RoleMapping) HasBackendRoles() bool`

HasBackendRoles returns a boolean if a field has been set.

### GetAndBackendRoles

`func (o *RoleMapping) GetAndBackendRoles() []string`

GetAndBackendRoles returns the AndBackendRoles field if non-nil, zero value otherwise.

### GetAndBackendRolesOk

`func (o *RoleMapping) GetAndBackendRolesOk() (*[]string, bool)`

GetAndBackendRolesOk returns a tuple with the AndBackendRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndBackendRoles

`func (o *RoleMapping) SetAndBackendRoles(v []string)`

SetAndBackendRoles sets AndBackendRoles field to given value.

### HasAndBackendRoles

`func (o *RoleMapping) HasAndBackendRoles() bool`

HasAndBackendRoles returns a boolean if a field has been set.

### GetDescription

`func (o *RoleMapping) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoleMapping) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoleMapping) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoleMapping) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


