# GetAllPitsResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pits** | Pointer to [**[]PitDetail**](PitDetail.md) |  | [optional] 

## Methods

### NewGetAllPitsResponseContent

`func NewGetAllPitsResponseContent() *GetAllPitsResponseContent`

NewGetAllPitsResponseContent instantiates a new GetAllPitsResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllPitsResponseContentWithDefaults

`func NewGetAllPitsResponseContentWithDefaults() *GetAllPitsResponseContent`

NewGetAllPitsResponseContentWithDefaults instantiates a new GetAllPitsResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPits

`func (o *GetAllPitsResponseContent) GetPits() []PitDetail`

GetPits returns the Pits field if non-nil, zero value otherwise.

### GetPitsOk

`func (o *GetAllPitsResponseContent) GetPitsOk() (*[]PitDetail, bool)`

GetPitsOk returns a tuple with the Pits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPits

`func (o *GetAllPitsResponseContent) SetPits(v []PitDetail)`

SetPits sets Pits field to given value.

### HasPits

`func (o *GetAllPitsResponseContent) HasPits() bool`

HasPits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


