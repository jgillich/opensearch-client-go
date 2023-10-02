# TenantPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantPatterns** | Pointer to **[]string** |  | [optional] 
**AllowedActions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewTenantPermission

`func NewTenantPermission() *TenantPermission`

NewTenantPermission instantiates a new TenantPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantPermissionWithDefaults

`func NewTenantPermissionWithDefaults() *TenantPermission`

NewTenantPermissionWithDefaults instantiates a new TenantPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantPatterns

`func (o *TenantPermission) GetTenantPatterns() []string`

GetTenantPatterns returns the TenantPatterns field if non-nil, zero value otherwise.

### GetTenantPatternsOk

`func (o *TenantPermission) GetTenantPatternsOk() (*[]string, bool)`

GetTenantPatternsOk returns a tuple with the TenantPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantPatterns

`func (o *TenantPermission) SetTenantPatterns(v []string)`

SetTenantPatterns sets TenantPatterns field to given value.

### HasTenantPatterns

`func (o *TenantPermission) HasTenantPatterns() bool`

HasTenantPatterns returns a boolean if a field has been set.

### GetAllowedActions

`func (o *TenantPermission) GetAllowedActions() []string`

GetAllowedActions returns the AllowedActions field if non-nil, zero value otherwise.

### GetAllowedActionsOk

`func (o *TenantPermission) GetAllowedActionsOk() (*[]string, bool)`

GetAllowedActionsOk returns a tuple with the AllowedActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedActions

`func (o *TenantPermission) SetAllowedActions(v []string)`

SetAllowedActions sets AllowedActions field to given value.

### HasAllowedActions

`func (o *TenantPermission) HasAllowedActions() bool`

HasAllowedActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


