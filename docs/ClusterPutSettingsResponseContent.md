# ClusterPutSettingsResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acknowledged** | Pointer to **bool** |  | [optional] 
**Persistent** | Pointer to **map[string]interface{}** |  | [optional] 
**Transient** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewClusterPutSettingsResponseContent

`func NewClusterPutSettingsResponseContent() *ClusterPutSettingsResponseContent`

NewClusterPutSettingsResponseContent instantiates a new ClusterPutSettingsResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterPutSettingsResponseContentWithDefaults

`func NewClusterPutSettingsResponseContentWithDefaults() *ClusterPutSettingsResponseContent`

NewClusterPutSettingsResponseContentWithDefaults instantiates a new ClusterPutSettingsResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcknowledged

`func (o *ClusterPutSettingsResponseContent) GetAcknowledged() bool`

GetAcknowledged returns the Acknowledged field if non-nil, zero value otherwise.

### GetAcknowledgedOk

`func (o *ClusterPutSettingsResponseContent) GetAcknowledgedOk() (*bool, bool)`

GetAcknowledgedOk returns a tuple with the Acknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledged

`func (o *ClusterPutSettingsResponseContent) SetAcknowledged(v bool)`

SetAcknowledged sets Acknowledged field to given value.

### HasAcknowledged

`func (o *ClusterPutSettingsResponseContent) HasAcknowledged() bool`

HasAcknowledged returns a boolean if a field has been set.

### GetPersistent

`func (o *ClusterPutSettingsResponseContent) GetPersistent() map[string]interface{}`

GetPersistent returns the Persistent field if non-nil, zero value otherwise.

### GetPersistentOk

`func (o *ClusterPutSettingsResponseContent) GetPersistentOk() (*map[string]interface{}, bool)`

GetPersistentOk returns a tuple with the Persistent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistent

`func (o *ClusterPutSettingsResponseContent) SetPersistent(v map[string]interface{})`

SetPersistent sets Persistent field to given value.

### HasPersistent

`func (o *ClusterPutSettingsResponseContent) HasPersistent() bool`

HasPersistent returns a boolean if a field has been set.

### GetTransient

`func (o *ClusterPutSettingsResponseContent) GetTransient() map[string]interface{}`

GetTransient returns the Transient field if non-nil, zero value otherwise.

### GetTransientOk

`func (o *ClusterPutSettingsResponseContent) GetTransientOk() (*map[string]interface{}, bool)`

GetTransientOk returns a tuple with the Transient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransient

`func (o *ClusterPutSettingsResponseContent) SetTransient(v map[string]interface{})`

SetTransient sets Transient field to given value.

### HasTransient

`func (o *ClusterPutSettingsResponseContent) HasTransient() bool`

HasTransient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


