# DataStream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**TimestampField** | Pointer to [**DataStreamTimestampField**](DataStreamTimestampField.md) |  | [optional] 
**Indices** | Pointer to [**[]DataStreamIndex**](DataStreamIndex.md) |  | [optional] 
**Generation** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to [**DataStreamStatus**](DataStreamStatus.md) |  | [optional] 
**Template** | Pointer to **string** |  | [optional] 

## Methods

### NewDataStream

`func NewDataStream() *DataStream`

NewDataStream instantiates a new DataStream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataStreamWithDefaults

`func NewDataStreamWithDefaults() *DataStream`

NewDataStreamWithDefaults instantiates a new DataStream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DataStream) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataStream) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataStream) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataStream) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTimestampField

`func (o *DataStream) GetTimestampField() DataStreamTimestampField`

GetTimestampField returns the TimestampField field if non-nil, zero value otherwise.

### GetTimestampFieldOk

`func (o *DataStream) GetTimestampFieldOk() (*DataStreamTimestampField, bool)`

GetTimestampFieldOk returns a tuple with the TimestampField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampField

`func (o *DataStream) SetTimestampField(v DataStreamTimestampField)`

SetTimestampField sets TimestampField field to given value.

### HasTimestampField

`func (o *DataStream) HasTimestampField() bool`

HasTimestampField returns a boolean if a field has been set.

### GetIndices

`func (o *DataStream) GetIndices() []DataStreamIndex`

GetIndices returns the Indices field if non-nil, zero value otherwise.

### GetIndicesOk

`func (o *DataStream) GetIndicesOk() (*[]DataStreamIndex, bool)`

GetIndicesOk returns a tuple with the Indices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndices

`func (o *DataStream) SetIndices(v []DataStreamIndex)`

SetIndices sets Indices field to given value.

### HasIndices

`func (o *DataStream) HasIndices() bool`

HasIndices returns a boolean if a field has been set.

### GetGeneration

`func (o *DataStream) GetGeneration() int64`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *DataStream) GetGenerationOk() (*int64, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *DataStream) SetGeneration(v int64)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *DataStream) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.

### GetStatus

`func (o *DataStream) GetStatus() DataStreamStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DataStream) GetStatusOk() (*DataStreamStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DataStream) SetStatus(v DataStreamStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DataStream) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTemplate

`func (o *DataStream) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *DataStream) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *DataStream) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *DataStream) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


