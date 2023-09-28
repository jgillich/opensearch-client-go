# GetResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**Version** | Pointer to **int32** |  | [optional] 
**SeqNo** | Pointer to **int64** |  | [optional] 
**PrimaryTerm** | Pointer to **int64** |  | [optional] 
**Found** | **bool** |  | 
**Routing** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **map[string]interface{}** |  | [optional] 
**Fields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGetResponseContent

`func NewGetResponseContent(index string, id string, found bool, ) *GetResponseContent`

NewGetResponseContent instantiates a new GetResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetResponseContentWithDefaults

`func NewGetResponseContentWithDefaults() *GetResponseContent`

NewGetResponseContentWithDefaults instantiates a new GetResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *GetResponseContent) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *GetResponseContent) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *GetResponseContent) SetIndex(v string)`

SetIndex sets Index field to given value.


### GetType

`func (o *GetResponseContent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetResponseContent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetResponseContent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetResponseContent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *GetResponseContent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetResponseContent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetResponseContent) SetId(v string)`

SetId sets Id field to given value.


### GetVersion

`func (o *GetResponseContent) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetResponseContent) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetResponseContent) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetResponseContent) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSeqNo

`func (o *GetResponseContent) GetSeqNo() int64`

GetSeqNo returns the SeqNo field if non-nil, zero value otherwise.

### GetSeqNoOk

`func (o *GetResponseContent) GetSeqNoOk() (*int64, bool)`

GetSeqNoOk returns a tuple with the SeqNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeqNo

`func (o *GetResponseContent) SetSeqNo(v int64)`

SetSeqNo sets SeqNo field to given value.

### HasSeqNo

`func (o *GetResponseContent) HasSeqNo() bool`

HasSeqNo returns a boolean if a field has been set.

### GetPrimaryTerm

`func (o *GetResponseContent) GetPrimaryTerm() int64`

GetPrimaryTerm returns the PrimaryTerm field if non-nil, zero value otherwise.

### GetPrimaryTermOk

`func (o *GetResponseContent) GetPrimaryTermOk() (*int64, bool)`

GetPrimaryTermOk returns a tuple with the PrimaryTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTerm

`func (o *GetResponseContent) SetPrimaryTerm(v int64)`

SetPrimaryTerm sets PrimaryTerm field to given value.

### HasPrimaryTerm

`func (o *GetResponseContent) HasPrimaryTerm() bool`

HasPrimaryTerm returns a boolean if a field has been set.

### GetFound

`func (o *GetResponseContent) GetFound() bool`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *GetResponseContent) GetFoundOk() (*bool, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *GetResponseContent) SetFound(v bool)`

SetFound sets Found field to given value.


### GetRouting

`func (o *GetResponseContent) GetRouting() string`

GetRouting returns the Routing field if non-nil, zero value otherwise.

### GetRoutingOk

`func (o *GetResponseContent) GetRoutingOk() (*string, bool)`

GetRoutingOk returns a tuple with the Routing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouting

`func (o *GetResponseContent) SetRouting(v string)`

SetRouting sets Routing field to given value.

### HasRouting

`func (o *GetResponseContent) HasRouting() bool`

HasRouting returns a boolean if a field has been set.

### GetSource

`func (o *GetResponseContent) GetSource() map[string]interface{}`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetResponseContent) GetSourceOk() (*map[string]interface{}, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetResponseContent) SetSource(v map[string]interface{})`

SetSource sets Source field to given value.

### HasSource

`func (o *GetResponseContent) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetFields

`func (o *GetResponseContent) GetFields() map[string]interface{}`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *GetResponseContent) GetFieldsOk() (*map[string]interface{}, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *GetResponseContent) SetFields(v map[string]interface{})`

SetFields sets Fields field to given value.

### HasFields

`func (o *GetResponseContent) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


