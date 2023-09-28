# SearchPostWithIndexResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScrollId** | Pointer to **string** |  | [optional] 
**Took** | Pointer to **int64** |  | [optional] 
**TimedOut** | Pointer to **bool** |  | [optional] 
**Shards** | Pointer to [**ShardStatistics**](ShardStatistics.md) |  | [optional] 
**Hits** | Pointer to [**HitsMetadata**](HitsMetadata.md) |  | [optional] 

## Methods

### NewSearchPostWithIndexResponseContent

`func NewSearchPostWithIndexResponseContent() *SearchPostWithIndexResponseContent`

NewSearchPostWithIndexResponseContent instantiates a new SearchPostWithIndexResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchPostWithIndexResponseContentWithDefaults

`func NewSearchPostWithIndexResponseContentWithDefaults() *SearchPostWithIndexResponseContent`

NewSearchPostWithIndexResponseContentWithDefaults instantiates a new SearchPostWithIndexResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScrollId

`func (o *SearchPostWithIndexResponseContent) GetScrollId() string`

GetScrollId returns the ScrollId field if non-nil, zero value otherwise.

### GetScrollIdOk

`func (o *SearchPostWithIndexResponseContent) GetScrollIdOk() (*string, bool)`

GetScrollIdOk returns a tuple with the ScrollId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrollId

`func (o *SearchPostWithIndexResponseContent) SetScrollId(v string)`

SetScrollId sets ScrollId field to given value.

### HasScrollId

`func (o *SearchPostWithIndexResponseContent) HasScrollId() bool`

HasScrollId returns a boolean if a field has been set.

### GetTook

`func (o *SearchPostWithIndexResponseContent) GetTook() int64`

GetTook returns the Took field if non-nil, zero value otherwise.

### GetTookOk

`func (o *SearchPostWithIndexResponseContent) GetTookOk() (*int64, bool)`

GetTookOk returns a tuple with the Took field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTook

`func (o *SearchPostWithIndexResponseContent) SetTook(v int64)`

SetTook sets Took field to given value.

### HasTook

`func (o *SearchPostWithIndexResponseContent) HasTook() bool`

HasTook returns a boolean if a field has been set.

### GetTimedOut

`func (o *SearchPostWithIndexResponseContent) GetTimedOut() bool`

GetTimedOut returns the TimedOut field if non-nil, zero value otherwise.

### GetTimedOutOk

`func (o *SearchPostWithIndexResponseContent) GetTimedOutOk() (*bool, bool)`

GetTimedOutOk returns a tuple with the TimedOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimedOut

`func (o *SearchPostWithIndexResponseContent) SetTimedOut(v bool)`

SetTimedOut sets TimedOut field to given value.

### HasTimedOut

`func (o *SearchPostWithIndexResponseContent) HasTimedOut() bool`

HasTimedOut returns a boolean if a field has been set.

### GetShards

`func (o *SearchPostWithIndexResponseContent) GetShards() ShardStatistics`

GetShards returns the Shards field if non-nil, zero value otherwise.

### GetShardsOk

`func (o *SearchPostWithIndexResponseContent) GetShardsOk() (*ShardStatistics, bool)`

GetShardsOk returns a tuple with the Shards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShards

`func (o *SearchPostWithIndexResponseContent) SetShards(v ShardStatistics)`

SetShards sets Shards field to given value.

### HasShards

`func (o *SearchPostWithIndexResponseContent) HasShards() bool`

HasShards returns a boolean if a field has been set.

### GetHits

`func (o *SearchPostWithIndexResponseContent) GetHits() HitsMetadata`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *SearchPostWithIndexResponseContent) GetHitsOk() (*HitsMetadata, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *SearchPostWithIndexResponseContent) SetHits(v HitsMetadata)`

SetHits sets Hits field to given value.

### HasHits

`func (o *SearchPostWithIndexResponseContent) HasHits() bool`

HasHits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


