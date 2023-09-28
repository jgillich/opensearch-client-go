# ActionGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reserved** | Pointer to **bool** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**AllowedActions** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Static** | Pointer to **bool** |  | [optional] 

## Methods

### NewActionGroup

`func NewActionGroup() *ActionGroup`

NewActionGroup instantiates a new ActionGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionGroupWithDefaults

`func NewActionGroupWithDefaults() *ActionGroup`

NewActionGroupWithDefaults instantiates a new ActionGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReserved

`func (o *ActionGroup) GetReserved() bool`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *ActionGroup) GetReservedOk() (*bool, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *ActionGroup) SetReserved(v bool)`

SetReserved sets Reserved field to given value.

### HasReserved

`func (o *ActionGroup) HasReserved() bool`

HasReserved returns a boolean if a field has been set.

### GetHidden

`func (o *ActionGroup) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *ActionGroup) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *ActionGroup) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *ActionGroup) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetAllowedActions

`func (o *ActionGroup) GetAllowedActions() []string`

GetAllowedActions returns the AllowedActions field if non-nil, zero value otherwise.

### GetAllowedActionsOk

`func (o *ActionGroup) GetAllowedActionsOk() (*[]string, bool)`

GetAllowedActionsOk returns a tuple with the AllowedActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedActions

`func (o *ActionGroup) SetAllowedActions(v []string)`

SetAllowedActions sets AllowedActions field to given value.

### HasAllowedActions

`func (o *ActionGroup) HasAllowedActions() bool`

HasAllowedActions returns a boolean if a field has been set.

### GetType

`func (o *ActionGroup) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActionGroup) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActionGroup) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ActionGroup) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *ActionGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ActionGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ActionGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ActionGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatic

`func (o *ActionGroup) GetStatic() bool`

GetStatic returns the Static field if non-nil, zero value otherwise.

### GetStaticOk

`func (o *ActionGroup) GetStaticOk() (*bool, bool)`

GetStaticOk returns a tuple with the Static field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic

`func (o *ActionGroup) SetStatic(v bool)`

SetStatic sets Static field to given value.

### HasStatic

`func (o *ActionGroup) HasStatic() bool`

HasStatic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


