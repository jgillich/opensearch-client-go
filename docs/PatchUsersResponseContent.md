# PatchUsersResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Security Operation Status | [optional] 
**Message** | Pointer to **string** | Security Operation Message | [optional] 

## Methods

### NewPatchUsersResponseContent

`func NewPatchUsersResponseContent() *PatchUsersResponseContent`

NewPatchUsersResponseContent instantiates a new PatchUsersResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchUsersResponseContentWithDefaults

`func NewPatchUsersResponseContentWithDefaults() *PatchUsersResponseContent`

NewPatchUsersResponseContentWithDefaults instantiates a new PatchUsersResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PatchUsersResponseContent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchUsersResponseContent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchUsersResponseContent) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchUsersResponseContent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *PatchUsersResponseContent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PatchUsersResponseContent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PatchUsersResponseContent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PatchUsersResponseContent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


