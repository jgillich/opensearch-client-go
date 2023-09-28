# ComplianceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**WriteLogDiffs** | Pointer to **bool** |  | [optional] 
**ReadWatchedFields** | Pointer to **interface{}** |  | [optional] 
**ReadIgnoreUsers** | Pointer to **[]string** |  | [optional] 
**WriteWatchedIndices** | Pointer to **[]string** |  | [optional] 
**WriteIgnoreUsers** | Pointer to **[]string** |  | [optional] 
**ReadMetadataOnly** | Pointer to **bool** |  | [optional] 
**WriteMetadataOnly** | Pointer to **bool** |  | [optional] 
**ExternalConfig** | Pointer to **bool** |  | [optional] 
**InternalConfig** | Pointer to **bool** |  | [optional] 

## Methods

### NewComplianceConfig

`func NewComplianceConfig() *ComplianceConfig`

NewComplianceConfig instantiates a new ComplianceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComplianceConfigWithDefaults

`func NewComplianceConfigWithDefaults() *ComplianceConfig`

NewComplianceConfigWithDefaults instantiates a new ComplianceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ComplianceConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ComplianceConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ComplianceConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ComplianceConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetWriteLogDiffs

`func (o *ComplianceConfig) GetWriteLogDiffs() bool`

GetWriteLogDiffs returns the WriteLogDiffs field if non-nil, zero value otherwise.

### GetWriteLogDiffsOk

`func (o *ComplianceConfig) GetWriteLogDiffsOk() (*bool, bool)`

GetWriteLogDiffsOk returns a tuple with the WriteLogDiffs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteLogDiffs

`func (o *ComplianceConfig) SetWriteLogDiffs(v bool)`

SetWriteLogDiffs sets WriteLogDiffs field to given value.

### HasWriteLogDiffs

`func (o *ComplianceConfig) HasWriteLogDiffs() bool`

HasWriteLogDiffs returns a boolean if a field has been set.

### GetReadWatchedFields

`func (o *ComplianceConfig) GetReadWatchedFields() interface{}`

GetReadWatchedFields returns the ReadWatchedFields field if non-nil, zero value otherwise.

### GetReadWatchedFieldsOk

`func (o *ComplianceConfig) GetReadWatchedFieldsOk() (*interface{}, bool)`

GetReadWatchedFieldsOk returns a tuple with the ReadWatchedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadWatchedFields

`func (o *ComplianceConfig) SetReadWatchedFields(v interface{})`

SetReadWatchedFields sets ReadWatchedFields field to given value.

### HasReadWatchedFields

`func (o *ComplianceConfig) HasReadWatchedFields() bool`

HasReadWatchedFields returns a boolean if a field has been set.

### SetReadWatchedFieldsNil

`func (o *ComplianceConfig) SetReadWatchedFieldsNil(b bool)`

 SetReadWatchedFieldsNil sets the value for ReadWatchedFields to be an explicit nil

### UnsetReadWatchedFields
`func (o *ComplianceConfig) UnsetReadWatchedFields()`

UnsetReadWatchedFields ensures that no value is present for ReadWatchedFields, not even an explicit nil
### GetReadIgnoreUsers

`func (o *ComplianceConfig) GetReadIgnoreUsers() []string`

GetReadIgnoreUsers returns the ReadIgnoreUsers field if non-nil, zero value otherwise.

### GetReadIgnoreUsersOk

`func (o *ComplianceConfig) GetReadIgnoreUsersOk() (*[]string, bool)`

GetReadIgnoreUsersOk returns a tuple with the ReadIgnoreUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadIgnoreUsers

`func (o *ComplianceConfig) SetReadIgnoreUsers(v []string)`

SetReadIgnoreUsers sets ReadIgnoreUsers field to given value.

### HasReadIgnoreUsers

`func (o *ComplianceConfig) HasReadIgnoreUsers() bool`

HasReadIgnoreUsers returns a boolean if a field has been set.

### GetWriteWatchedIndices

`func (o *ComplianceConfig) GetWriteWatchedIndices() []string`

GetWriteWatchedIndices returns the WriteWatchedIndices field if non-nil, zero value otherwise.

### GetWriteWatchedIndicesOk

`func (o *ComplianceConfig) GetWriteWatchedIndicesOk() (*[]string, bool)`

GetWriteWatchedIndicesOk returns a tuple with the WriteWatchedIndices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteWatchedIndices

`func (o *ComplianceConfig) SetWriteWatchedIndices(v []string)`

SetWriteWatchedIndices sets WriteWatchedIndices field to given value.

### HasWriteWatchedIndices

`func (o *ComplianceConfig) HasWriteWatchedIndices() bool`

HasWriteWatchedIndices returns a boolean if a field has been set.

### GetWriteIgnoreUsers

`func (o *ComplianceConfig) GetWriteIgnoreUsers() []string`

GetWriteIgnoreUsers returns the WriteIgnoreUsers field if non-nil, zero value otherwise.

### GetWriteIgnoreUsersOk

`func (o *ComplianceConfig) GetWriteIgnoreUsersOk() (*[]string, bool)`

GetWriteIgnoreUsersOk returns a tuple with the WriteIgnoreUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteIgnoreUsers

`func (o *ComplianceConfig) SetWriteIgnoreUsers(v []string)`

SetWriteIgnoreUsers sets WriteIgnoreUsers field to given value.

### HasWriteIgnoreUsers

`func (o *ComplianceConfig) HasWriteIgnoreUsers() bool`

HasWriteIgnoreUsers returns a boolean if a field has been set.

### GetReadMetadataOnly

`func (o *ComplianceConfig) GetReadMetadataOnly() bool`

GetReadMetadataOnly returns the ReadMetadataOnly field if non-nil, zero value otherwise.

### GetReadMetadataOnlyOk

`func (o *ComplianceConfig) GetReadMetadataOnlyOk() (*bool, bool)`

GetReadMetadataOnlyOk returns a tuple with the ReadMetadataOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadMetadataOnly

`func (o *ComplianceConfig) SetReadMetadataOnly(v bool)`

SetReadMetadataOnly sets ReadMetadataOnly field to given value.

### HasReadMetadataOnly

`func (o *ComplianceConfig) HasReadMetadataOnly() bool`

HasReadMetadataOnly returns a boolean if a field has been set.

### GetWriteMetadataOnly

`func (o *ComplianceConfig) GetWriteMetadataOnly() bool`

GetWriteMetadataOnly returns the WriteMetadataOnly field if non-nil, zero value otherwise.

### GetWriteMetadataOnlyOk

`func (o *ComplianceConfig) GetWriteMetadataOnlyOk() (*bool, bool)`

GetWriteMetadataOnlyOk returns a tuple with the WriteMetadataOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteMetadataOnly

`func (o *ComplianceConfig) SetWriteMetadataOnly(v bool)`

SetWriteMetadataOnly sets WriteMetadataOnly field to given value.

### HasWriteMetadataOnly

`func (o *ComplianceConfig) HasWriteMetadataOnly() bool`

HasWriteMetadataOnly returns a boolean if a field has been set.

### GetExternalConfig

`func (o *ComplianceConfig) GetExternalConfig() bool`

GetExternalConfig returns the ExternalConfig field if non-nil, zero value otherwise.

### GetExternalConfigOk

`func (o *ComplianceConfig) GetExternalConfigOk() (*bool, bool)`

GetExternalConfigOk returns a tuple with the ExternalConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalConfig

`func (o *ComplianceConfig) SetExternalConfig(v bool)`

SetExternalConfig sets ExternalConfig field to given value.

### HasExternalConfig

`func (o *ComplianceConfig) HasExternalConfig() bool`

HasExternalConfig returns a boolean if a field has been set.

### GetInternalConfig

`func (o *ComplianceConfig) GetInternalConfig() bool`

GetInternalConfig returns the InternalConfig field if non-nil, zero value otherwise.

### GetInternalConfigOk

`func (o *ComplianceConfig) GetInternalConfigOk() (*bool, bool)`

GetInternalConfigOk returns a tuple with the InternalConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalConfig

`func (o *ComplianceConfig) SetInternalConfig(v bool)`

SetInternalConfig sets InternalConfig field to given value.

### HasInternalConfig

`func (o *ComplianceConfig) HasInternalConfig() bool`

HasInternalConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


