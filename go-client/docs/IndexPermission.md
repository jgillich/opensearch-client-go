# IndexPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IndexPatterns** | Pointer to **[]string** |  | [optional] 
**Fls** | Pointer to **[]string** |  | [optional] 
**MaskedFields** | Pointer to **[]string** |  | [optional] 
**AllowedActions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewIndexPermission

`func NewIndexPermission() *IndexPermission`

NewIndexPermission instantiates a new IndexPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexPermissionWithDefaults

`func NewIndexPermissionWithDefaults() *IndexPermission`

NewIndexPermissionWithDefaults instantiates a new IndexPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexPatterns

`func (o *IndexPermission) GetIndexPatterns() []string`

GetIndexPatterns returns the IndexPatterns field if non-nil, zero value otherwise.

### GetIndexPatternsOk

`func (o *IndexPermission) GetIndexPatternsOk() (*[]string, bool)`

GetIndexPatternsOk returns a tuple with the IndexPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexPatterns

`func (o *IndexPermission) SetIndexPatterns(v []string)`

SetIndexPatterns sets IndexPatterns field to given value.

### HasIndexPatterns

`func (o *IndexPermission) HasIndexPatterns() bool`

HasIndexPatterns returns a boolean if a field has been set.

### GetFls

`func (o *IndexPermission) GetFls() []string`

GetFls returns the Fls field if non-nil, zero value otherwise.

### GetFlsOk

`func (o *IndexPermission) GetFlsOk() (*[]string, bool)`

GetFlsOk returns a tuple with the Fls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFls

`func (o *IndexPermission) SetFls(v []string)`

SetFls sets Fls field to given value.

### HasFls

`func (o *IndexPermission) HasFls() bool`

HasFls returns a boolean if a field has been set.

### GetMaskedFields

`func (o *IndexPermission) GetMaskedFields() []string`

GetMaskedFields returns the MaskedFields field if non-nil, zero value otherwise.

### GetMaskedFieldsOk

`func (o *IndexPermission) GetMaskedFieldsOk() (*[]string, bool)`

GetMaskedFieldsOk returns a tuple with the MaskedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedFields

`func (o *IndexPermission) SetMaskedFields(v []string)`

SetMaskedFields sets MaskedFields field to given value.

### HasMaskedFields

`func (o *IndexPermission) HasMaskedFields() bool`

HasMaskedFields returns a boolean if a field has been set.

### GetAllowedActions

`func (o *IndexPermission) GetAllowedActions() []string`

GetAllowedActions returns the AllowedActions field if non-nil, zero value otherwise.

### GetAllowedActionsOk

`func (o *IndexPermission) GetAllowedActionsOk() (*[]string, bool)`

GetAllowedActionsOk returns a tuple with the AllowedActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedActions

`func (o *IndexPermission) SetAllowedActions(v []string)`

SetAllowedActions sets AllowedActions field to given value.

### HasAllowedActions

`func (o *IndexPermission) HasAllowedActions() bool`

HasAllowedActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


