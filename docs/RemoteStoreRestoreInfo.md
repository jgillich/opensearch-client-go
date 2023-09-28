# RemoteStoreRestoreInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Snapshot** | Pointer to **string** |  | [optional] 
**Indices** | Pointer to **[]string** |  | [optional] 
**Shards** | Pointer to [**RemoteStoreRestoreShardsInfo**](RemoteStoreRestoreShardsInfo.md) |  | [optional] 

## Methods

### NewRemoteStoreRestoreInfo

`func NewRemoteStoreRestoreInfo() *RemoteStoreRestoreInfo`

NewRemoteStoreRestoreInfo instantiates a new RemoteStoreRestoreInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteStoreRestoreInfoWithDefaults

`func NewRemoteStoreRestoreInfoWithDefaults() *RemoteStoreRestoreInfo`

NewRemoteStoreRestoreInfoWithDefaults instantiates a new RemoteStoreRestoreInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshot

`func (o *RemoteStoreRestoreInfo) GetSnapshot() string`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *RemoteStoreRestoreInfo) GetSnapshotOk() (*string, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *RemoteStoreRestoreInfo) SetSnapshot(v string)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *RemoteStoreRestoreInfo) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### GetIndices

`func (o *RemoteStoreRestoreInfo) GetIndices() []string`

GetIndices returns the Indices field if non-nil, zero value otherwise.

### GetIndicesOk

`func (o *RemoteStoreRestoreInfo) GetIndicesOk() (*[]string, bool)`

GetIndicesOk returns a tuple with the Indices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndices

`func (o *RemoteStoreRestoreInfo) SetIndices(v []string)`

SetIndices sets Indices field to given value.

### HasIndices

`func (o *RemoteStoreRestoreInfo) HasIndices() bool`

HasIndices returns a boolean if a field has been set.

### GetShards

`func (o *RemoteStoreRestoreInfo) GetShards() RemoteStoreRestoreShardsInfo`

GetShards returns the Shards field if non-nil, zero value otherwise.

### GetShardsOk

`func (o *RemoteStoreRestoreInfo) GetShardsOk() (*RemoteStoreRestoreShardsInfo, bool)`

GetShardsOk returns a tuple with the Shards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShards

`func (o *RemoteStoreRestoreInfo) SetShards(v RemoteStoreRestoreShardsInfo)`

SetShards sets Shards field to given value.

### HasShards

`func (o *RemoteStoreRestoreInfo) HasShards() bool`

HasShards returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


