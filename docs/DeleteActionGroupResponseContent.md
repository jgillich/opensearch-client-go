# DeleteActionGroupResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Security Operation Status | [optional] 
**Message** | Pointer to **string** | Security Operation Message | [optional] 

## Methods

### NewDeleteActionGroupResponseContent

`func NewDeleteActionGroupResponseContent() *DeleteActionGroupResponseContent`

NewDeleteActionGroupResponseContent instantiates a new DeleteActionGroupResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteActionGroupResponseContentWithDefaults

`func NewDeleteActionGroupResponseContentWithDefaults() *DeleteActionGroupResponseContent`

NewDeleteActionGroupResponseContentWithDefaults instantiates a new DeleteActionGroupResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DeleteActionGroupResponseContent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeleteActionGroupResponseContent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeleteActionGroupResponseContent) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeleteActionGroupResponseContent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *DeleteActionGroupResponseContent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DeleteActionGroupResponseContent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DeleteActionGroupResponseContent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DeleteActionGroupResponseContent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


