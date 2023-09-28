# HitsMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to [**Total**](Total.md) |  | [optional] 
**MaxScore** | Pointer to **float64** |  | [optional] 
**Hits** | Pointer to [**[]Hits**](Hits.md) |  | [optional] 

## Methods

### NewHitsMetadata

`func NewHitsMetadata() *HitsMetadata`

NewHitsMetadata instantiates a new HitsMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHitsMetadataWithDefaults

`func NewHitsMetadataWithDefaults() *HitsMetadata`

NewHitsMetadataWithDefaults instantiates a new HitsMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *HitsMetadata) GetTotal() Total`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *HitsMetadata) GetTotalOk() (*Total, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *HitsMetadata) SetTotal(v Total)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *HitsMetadata) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetMaxScore

`func (o *HitsMetadata) GetMaxScore() float64`

GetMaxScore returns the MaxScore field if non-nil, zero value otherwise.

### GetMaxScoreOk

`func (o *HitsMetadata) GetMaxScoreOk() (*float64, bool)`

GetMaxScoreOk returns a tuple with the MaxScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxScore

`func (o *HitsMetadata) SetMaxScore(v float64)`

SetMaxScore sets MaxScore field to given value.

### HasMaxScore

`func (o *HitsMetadata) HasMaxScore() bool`

HasMaxScore returns a boolean if a field has been set.

### GetHits

`func (o *HitsMetadata) GetHits() []Hits`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *HitsMetadata) GetHitsOk() (*[]Hits, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *HitsMetadata) SetHits(v []Hits)`

SetHits sets Hits field to given value.

### HasHits

`func (o *HitsMetadata) HasHits() bool`

HasHits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


