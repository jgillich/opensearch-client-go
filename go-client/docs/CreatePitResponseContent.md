# CreatePitResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PitId** | Pointer to **string** |  | [optional] 
**Shard** | Pointer to [**ShardStatistics**](ShardStatistics.md) |  | [optional] 
**CreationTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewCreatePitResponseContent

`func NewCreatePitResponseContent() *CreatePitResponseContent`

NewCreatePitResponseContent instantiates a new CreatePitResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePitResponseContentWithDefaults

`func NewCreatePitResponseContentWithDefaults() *CreatePitResponseContent`

NewCreatePitResponseContentWithDefaults instantiates a new CreatePitResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPitId

`func (o *CreatePitResponseContent) GetPitId() string`

GetPitId returns the PitId field if non-nil, zero value otherwise.

### GetPitIdOk

`func (o *CreatePitResponseContent) GetPitIdOk() (*string, bool)`

GetPitIdOk returns a tuple with the PitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPitId

`func (o *CreatePitResponseContent) SetPitId(v string)`

SetPitId sets PitId field to given value.

### HasPitId

`func (o *CreatePitResponseContent) HasPitId() bool`

HasPitId returns a boolean if a field has been set.

### GetShard

`func (o *CreatePitResponseContent) GetShard() ShardStatistics`

GetShard returns the Shard field if non-nil, zero value otherwise.

### GetShardOk

`func (o *CreatePitResponseContent) GetShardOk() (*ShardStatistics, bool)`

GetShardOk returns a tuple with the Shard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShard

`func (o *CreatePitResponseContent) SetShard(v ShardStatistics)`

SetShard sets Shard field to given value.

### HasShard

`func (o *CreatePitResponseContent) HasShard() bool`

HasShard returns a boolean if a field has been set.

### GetCreationTime

`func (o *CreatePitResponseContent) GetCreationTime() int64`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CreatePitResponseContent) GetCreationTimeOk() (*int64, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CreatePitResponseContent) SetCreationTime(v int64)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *CreatePitResponseContent) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


