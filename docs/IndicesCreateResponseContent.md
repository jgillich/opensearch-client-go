# IndicesCreateResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **string** |  | 
**ShardsAcknowledged** | **bool** |  | 
**Acknowledged** | **bool** |  | 

## Methods

### NewIndicesCreateResponseContent

`func NewIndicesCreateResponseContent(index string, shardsAcknowledged bool, acknowledged bool, ) *IndicesCreateResponseContent`

NewIndicesCreateResponseContent instantiates a new IndicesCreateResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndicesCreateResponseContentWithDefaults

`func NewIndicesCreateResponseContentWithDefaults() *IndicesCreateResponseContent`

NewIndicesCreateResponseContentWithDefaults instantiates a new IndicesCreateResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *IndicesCreateResponseContent) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *IndicesCreateResponseContent) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *IndicesCreateResponseContent) SetIndex(v string)`

SetIndex sets Index field to given value.


### GetShardsAcknowledged

`func (o *IndicesCreateResponseContent) GetShardsAcknowledged() bool`

GetShardsAcknowledged returns the ShardsAcknowledged field if non-nil, zero value otherwise.

### GetShardsAcknowledgedOk

`func (o *IndicesCreateResponseContent) GetShardsAcknowledgedOk() (*bool, bool)`

GetShardsAcknowledgedOk returns a tuple with the ShardsAcknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardsAcknowledged

`func (o *IndicesCreateResponseContent) SetShardsAcknowledged(v bool)`

SetShardsAcknowledged sets ShardsAcknowledged field to given value.


### GetAcknowledged

`func (o *IndicesCreateResponseContent) GetAcknowledged() bool`

GetAcknowledged returns the Acknowledged field if non-nil, zero value otherwise.

### GetAcknowledgedOk

`func (o *IndicesCreateResponseContent) GetAcknowledgedOk() (*bool, bool)`

GetAcknowledgedOk returns a tuple with the Acknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledged

`func (o *IndicesCreateResponseContent) SetAcknowledged(v bool)`

SetAcknowledged sets Acknowledged field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


