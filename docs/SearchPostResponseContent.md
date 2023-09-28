# SearchPostResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScrollId** | Pointer to **string** |  | [optional] 
**Took** | Pointer to **int64** |  | [optional] 
**TimedOut** | Pointer to **bool** |  | [optional] 
**Shards** | Pointer to [**ShardStatistics**](ShardStatistics.md) |  | [optional] 
**Hits** | Pointer to [**HitsMetadata**](HitsMetadata.md) |  | [optional] 

## Methods

### NewSearchPostResponseContent

`func NewSearchPostResponseContent() *SearchPostResponseContent`

NewSearchPostResponseContent instantiates a new SearchPostResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchPostResponseContentWithDefaults

`func NewSearchPostResponseContentWithDefaults() *SearchPostResponseContent`

NewSearchPostResponseContentWithDefaults instantiates a new SearchPostResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScrollId

`func (o *SearchPostResponseContent) GetScrollId() string`

GetScrollId returns the ScrollId field if non-nil, zero value otherwise.

### GetScrollIdOk

`func (o *SearchPostResponseContent) GetScrollIdOk() (*string, bool)`

GetScrollIdOk returns a tuple with the ScrollId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrollId

`func (o *SearchPostResponseContent) SetScrollId(v string)`

SetScrollId sets ScrollId field to given value.

### HasScrollId

`func (o *SearchPostResponseContent) HasScrollId() bool`

HasScrollId returns a boolean if a field has been set.

### GetTook

`func (o *SearchPostResponseContent) GetTook() int64`

GetTook returns the Took field if non-nil, zero value otherwise.

### GetTookOk

`func (o *SearchPostResponseContent) GetTookOk() (*int64, bool)`

GetTookOk returns a tuple with the Took field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTook

`func (o *SearchPostResponseContent) SetTook(v int64)`

SetTook sets Took field to given value.

### HasTook

`func (o *SearchPostResponseContent) HasTook() bool`

HasTook returns a boolean if a field has been set.

### GetTimedOut

`func (o *SearchPostResponseContent) GetTimedOut() bool`

GetTimedOut returns the TimedOut field if non-nil, zero value otherwise.

### GetTimedOutOk

`func (o *SearchPostResponseContent) GetTimedOutOk() (*bool, bool)`

GetTimedOutOk returns a tuple with the TimedOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimedOut

`func (o *SearchPostResponseContent) SetTimedOut(v bool)`

SetTimedOut sets TimedOut field to given value.

### HasTimedOut

`func (o *SearchPostResponseContent) HasTimedOut() bool`

HasTimedOut returns a boolean if a field has been set.

### GetShards

`func (o *SearchPostResponseContent) GetShards() ShardStatistics`

GetShards returns the Shards field if non-nil, zero value otherwise.

### GetShardsOk

`func (o *SearchPostResponseContent) GetShardsOk() (*ShardStatistics, bool)`

GetShardsOk returns a tuple with the Shards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShards

`func (o *SearchPostResponseContent) SetShards(v ShardStatistics)`

SetShards sets Shards field to given value.

### HasShards

`func (o *SearchPostResponseContent) HasShards() bool`

HasShards returns a boolean if a field has been set.

### GetHits

`func (o *SearchPostResponseContent) GetHits() HitsMetadata`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *SearchPostResponseContent) GetHitsOk() (*HitsMetadata, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *SearchPostResponseContent) SetHits(v HitsMetadata)`

SetHits sets Hits field to given value.

### HasHits

`func (o *SearchPostResponseContent) HasHits() bool`

HasHits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


