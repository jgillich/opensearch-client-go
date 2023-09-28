# Tenant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reserved** | Pointer to **bool** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Static** | Pointer to **bool** |  | [optional] 

## Methods

### NewTenant

`func NewTenant() *Tenant`

NewTenant instantiates a new Tenant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantWithDefaults

`func NewTenantWithDefaults() *Tenant`

NewTenantWithDefaults instantiates a new Tenant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReserved

`func (o *Tenant) GetReserved() bool`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *Tenant) GetReservedOk() (*bool, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *Tenant) SetReserved(v bool)`

SetReserved sets Reserved field to given value.

### HasReserved

`func (o *Tenant) HasReserved() bool`

HasReserved returns a boolean if a field has been set.

### GetHidden

`func (o *Tenant) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *Tenant) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *Tenant) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *Tenant) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetDescription

`func (o *Tenant) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Tenant) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Tenant) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Tenant) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatic

`func (o *Tenant) GetStatic() bool`

GetStatic returns the Static field if non-nil, zero value otherwise.

### GetStaticOk

`func (o *Tenant) GetStaticOk() (*bool, bool)`

GetStaticOk returns a tuple with the Static field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic

`func (o *Tenant) SetStatic(v bool)`

SetStatic sets Static field to given value.

### HasStatic

`func (o *Tenant) HasStatic() bool`

HasStatic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


