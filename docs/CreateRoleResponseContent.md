# CreateRoleResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Security Operation Status | [optional] 
**Message** | Pointer to **string** | Security Operation Message | [optional] 

## Methods

### NewCreateRoleResponseContent

`func NewCreateRoleResponseContent() *CreateRoleResponseContent`

NewCreateRoleResponseContent instantiates a new CreateRoleResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRoleResponseContentWithDefaults

`func NewCreateRoleResponseContentWithDefaults() *CreateRoleResponseContent`

NewCreateRoleResponseContentWithDefaults instantiates a new CreateRoleResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CreateRoleResponseContent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateRoleResponseContent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateRoleResponseContent) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateRoleResponseContent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *CreateRoleResponseContent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateRoleResponseContent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateRoleResponseContent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreateRoleResponseContent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


