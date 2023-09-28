# ShardStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**Successful** | Pointer to **int32** |  | [optional] 
**Skipped** | Pointer to **int32** |  | [optional] 
**Failed** | Pointer to **int32** |  | [optional] 

## Methods

### NewShardStatistics

`func NewShardStatistics() *ShardStatistics`

NewShardStatistics instantiates a new ShardStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShardStatisticsWithDefaults

`func NewShardStatisticsWithDefaults() *ShardStatistics`

NewShardStatisticsWithDefaults instantiates a new ShardStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ShardStatistics) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ShardStatistics) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ShardStatistics) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ShardStatistics) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetSuccessful

`func (o *ShardStatistics) GetSuccessful() int32`

GetSuccessful returns the Successful field if non-nil, zero value otherwise.

### GetSuccessfulOk

`func (o *ShardStatistics) GetSuccessfulOk() (*int32, bool)`

GetSuccessfulOk returns a tuple with the Successful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessful

`func (o *ShardStatistics) SetSuccessful(v int32)`

SetSuccessful sets Successful field to given value.

### HasSuccessful

`func (o *ShardStatistics) HasSuccessful() bool`

HasSuccessful returns a boolean if a field has been set.

### GetSkipped

`func (o *ShardStatistics) GetSkipped() int32`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *ShardStatistics) GetSkippedOk() (*int32, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *ShardStatistics) SetSkipped(v int32)`

SetSkipped sets Skipped field to given value.

### HasSkipped

`func (o *ShardStatistics) HasSkipped() bool`

HasSkipped returns a boolean if a field has been set.

### GetFailed

`func (o *ShardStatistics) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *ShardStatistics) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *ShardStatistics) SetFailed(v int32)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *ShardStatistics) HasFailed() bool`

HasFailed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


