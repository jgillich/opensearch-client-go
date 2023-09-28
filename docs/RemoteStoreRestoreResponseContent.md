# RemoteStoreRestoreResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accepted** | Pointer to **bool** |  | [optional] 
**RemoteStore** | Pointer to [**RemoteStoreRestoreInfo**](RemoteStoreRestoreInfo.md) |  | [optional] 

## Methods

### NewRemoteStoreRestoreResponseContent

`func NewRemoteStoreRestoreResponseContent() *RemoteStoreRestoreResponseContent`

NewRemoteStoreRestoreResponseContent instantiates a new RemoteStoreRestoreResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteStoreRestoreResponseContentWithDefaults

`func NewRemoteStoreRestoreResponseContentWithDefaults() *RemoteStoreRestoreResponseContent`

NewRemoteStoreRestoreResponseContentWithDefaults instantiates a new RemoteStoreRestoreResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccepted

`func (o *RemoteStoreRestoreResponseContent) GetAccepted() bool`

GetAccepted returns the Accepted field if non-nil, zero value otherwise.

### GetAcceptedOk

`func (o *RemoteStoreRestoreResponseContent) GetAcceptedOk() (*bool, bool)`

GetAcceptedOk returns a tuple with the Accepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepted

`func (o *RemoteStoreRestoreResponseContent) SetAccepted(v bool)`

SetAccepted sets Accepted field to given value.

### HasAccepted

`func (o *RemoteStoreRestoreResponseContent) HasAccepted() bool`

HasAccepted returns a boolean if a field has been set.

### GetRemoteStore

`func (o *RemoteStoreRestoreResponseContent) GetRemoteStore() RemoteStoreRestoreInfo`

GetRemoteStore returns the RemoteStore field if non-nil, zero value otherwise.

### GetRemoteStoreOk

`func (o *RemoteStoreRestoreResponseContent) GetRemoteStoreOk() (*RemoteStoreRestoreInfo, bool)`

GetRemoteStoreOk returns a tuple with the RemoteStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteStore

`func (o *RemoteStoreRestoreResponseContent) SetRemoteStore(v RemoteStoreRestoreInfo)`

SetRemoteStore sets RemoteStore field to given value.

### HasRemoteStore

`func (o *RemoteStoreRestoreResponseContent) HasRemoteStore() bool`

HasRemoteStore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


