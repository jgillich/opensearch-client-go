# UserDefinedStructure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** |  | [optional] 
**Aliases** | Pointer to **[]string** |  | [optional] 
**Filter** | Pointer to **interface{}** |  | [optional] 
**Index** | Pointer to **string** |  | [optional] 
**Indices** | Pointer to **[]string** |  | [optional] 
**IndexRouting** | Pointer to **string** |  | [optional] 
**IsHidden** | Pointer to **bool** |  | [optional] 
**IsWriteIndex** | Pointer to **bool** |  | [optional] 
**MustExist** | Pointer to **string** |  | [optional] 
**Routing** | Pointer to **string** |  | [optional] 
**SearchRouting** | Pointer to **string** |  | [optional] 

## Methods

### NewUserDefinedStructure

`func NewUserDefinedStructure() *UserDefinedStructure`

NewUserDefinedStructure instantiates a new UserDefinedStructure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDefinedStructureWithDefaults

`func NewUserDefinedStructureWithDefaults() *UserDefinedStructure`

NewUserDefinedStructureWithDefaults instantiates a new UserDefinedStructure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *UserDefinedStructure) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *UserDefinedStructure) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *UserDefinedStructure) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *UserDefinedStructure) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetAliases

`func (o *UserDefinedStructure) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *UserDefinedStructure) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *UserDefinedStructure) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *UserDefinedStructure) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetFilter

`func (o *UserDefinedStructure) GetFilter() interface{}`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *UserDefinedStructure) GetFilterOk() (*interface{}, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *UserDefinedStructure) SetFilter(v interface{})`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *UserDefinedStructure) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *UserDefinedStructure) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *UserDefinedStructure) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetIndex

`func (o *UserDefinedStructure) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *UserDefinedStructure) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *UserDefinedStructure) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *UserDefinedStructure) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetIndices

`func (o *UserDefinedStructure) GetIndices() []string`

GetIndices returns the Indices field if non-nil, zero value otherwise.

### GetIndicesOk

`func (o *UserDefinedStructure) GetIndicesOk() (*[]string, bool)`

GetIndicesOk returns a tuple with the Indices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndices

`func (o *UserDefinedStructure) SetIndices(v []string)`

SetIndices sets Indices field to given value.

### HasIndices

`func (o *UserDefinedStructure) HasIndices() bool`

HasIndices returns a boolean if a field has been set.

### GetIndexRouting

`func (o *UserDefinedStructure) GetIndexRouting() string`

GetIndexRouting returns the IndexRouting field if non-nil, zero value otherwise.

### GetIndexRoutingOk

`func (o *UserDefinedStructure) GetIndexRoutingOk() (*string, bool)`

GetIndexRoutingOk returns a tuple with the IndexRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexRouting

`func (o *UserDefinedStructure) SetIndexRouting(v string)`

SetIndexRouting sets IndexRouting field to given value.

### HasIndexRouting

`func (o *UserDefinedStructure) HasIndexRouting() bool`

HasIndexRouting returns a boolean if a field has been set.

### GetIsHidden

`func (o *UserDefinedStructure) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *UserDefinedStructure) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *UserDefinedStructure) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *UserDefinedStructure) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetIsWriteIndex

`func (o *UserDefinedStructure) GetIsWriteIndex() bool`

GetIsWriteIndex returns the IsWriteIndex field if non-nil, zero value otherwise.

### GetIsWriteIndexOk

`func (o *UserDefinedStructure) GetIsWriteIndexOk() (*bool, bool)`

GetIsWriteIndexOk returns a tuple with the IsWriteIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWriteIndex

`func (o *UserDefinedStructure) SetIsWriteIndex(v bool)`

SetIsWriteIndex sets IsWriteIndex field to given value.

### HasIsWriteIndex

`func (o *UserDefinedStructure) HasIsWriteIndex() bool`

HasIsWriteIndex returns a boolean if a field has been set.

### GetMustExist

`func (o *UserDefinedStructure) GetMustExist() string`

GetMustExist returns the MustExist field if non-nil, zero value otherwise.

### GetMustExistOk

`func (o *UserDefinedStructure) GetMustExistOk() (*string, bool)`

GetMustExistOk returns a tuple with the MustExist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMustExist

`func (o *UserDefinedStructure) SetMustExist(v string)`

SetMustExist sets MustExist field to given value.

### HasMustExist

`func (o *UserDefinedStructure) HasMustExist() bool`

HasMustExist returns a boolean if a field has been set.

### GetRouting

`func (o *UserDefinedStructure) GetRouting() string`

GetRouting returns the Routing field if non-nil, zero value otherwise.

### GetRoutingOk

`func (o *UserDefinedStructure) GetRoutingOk() (*string, bool)`

GetRoutingOk returns a tuple with the Routing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouting

`func (o *UserDefinedStructure) SetRouting(v string)`

SetRouting sets Routing field to given value.

### HasRouting

`func (o *UserDefinedStructure) HasRouting() bool`

HasRouting returns a boolean if a field has been set.

### GetSearchRouting

`func (o *UserDefinedStructure) GetSearchRouting() string`

GetSearchRouting returns the SearchRouting field if non-nil, zero value otherwise.

### GetSearchRoutingOk

`func (o *UserDefinedStructure) GetSearchRoutingOk() (*string, bool)`

GetSearchRoutingOk returns a tuple with the SearchRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchRouting

`func (o *UserDefinedStructure) SetSearchRouting(v string)`

SetSearchRouting sets SearchRouting field to given value.

### HasSearchRouting

`func (o *UserDefinedStructure) HasSearchRouting() bool`

HasSearchRouting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


