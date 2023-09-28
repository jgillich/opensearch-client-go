# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | Pointer to **string** |  | [optional] 
**Reserved** | Pointer to **bool** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**BackendRoles** | Pointer to **[]string** |  | [optional] 
**Attributes** | Pointer to **map[string]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**OpendistroSecurityRoles** | Pointer to **[]string** |  | [optional] 
**Static** | Pointer to **bool** |  | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *User) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *User) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *User) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *User) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetReserved

`func (o *User) GetReserved() bool`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *User) GetReservedOk() (*bool, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *User) SetReserved(v bool)`

SetReserved sets Reserved field to given value.

### HasReserved

`func (o *User) HasReserved() bool`

HasReserved returns a boolean if a field has been set.

### GetHidden

`func (o *User) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *User) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *User) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *User) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetBackendRoles

`func (o *User) GetBackendRoles() []string`

GetBackendRoles returns the BackendRoles field if non-nil, zero value otherwise.

### GetBackendRolesOk

`func (o *User) GetBackendRolesOk() (*[]string, bool)`

GetBackendRolesOk returns a tuple with the BackendRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendRoles

`func (o *User) SetBackendRoles(v []string)`

SetBackendRoles sets BackendRoles field to given value.

### HasBackendRoles

`func (o *User) HasBackendRoles() bool`

HasBackendRoles returns a boolean if a field has been set.

### GetAttributes

`func (o *User) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *User) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *User) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *User) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetDescription

`func (o *User) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *User) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *User) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *User) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOpendistroSecurityRoles

`func (o *User) GetOpendistroSecurityRoles() []string`

GetOpendistroSecurityRoles returns the OpendistroSecurityRoles field if non-nil, zero value otherwise.

### GetOpendistroSecurityRolesOk

`func (o *User) GetOpendistroSecurityRolesOk() (*[]string, bool)`

GetOpendistroSecurityRolesOk returns a tuple with the OpendistroSecurityRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpendistroSecurityRoles

`func (o *User) SetOpendistroSecurityRoles(v []string)`

SetOpendistroSecurityRoles sets OpendistroSecurityRoles field to given value.

### HasOpendistroSecurityRoles

`func (o *User) HasOpendistroSecurityRoles() bool`

HasOpendistroSecurityRoles returns a boolean if a field has been set.

### GetStatic

`func (o *User) GetStatic() bool`

GetStatic returns the Static field if non-nil, zero value otherwise.

### GetStaticOk

`func (o *User) GetStaticOk() (*bool, bool)`

GetStaticOk returns a tuple with the Static field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic

`func (o *User) SetStatic(v bool)`

SetStatic sets Static field to given value.

### HasStatic

`func (o *User) HasStatic() bool`

HasStatic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


