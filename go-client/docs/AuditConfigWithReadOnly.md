# AuditConfigWithReadOnly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Readonly** | Pointer to **[]string** |  | [optional] 
**Config** | Pointer to [**AuditConfig**](AuditConfig.md) |  | [optional] 

## Methods

### NewAuditConfigWithReadOnly

`func NewAuditConfigWithReadOnly() *AuditConfigWithReadOnly`

NewAuditConfigWithReadOnly instantiates a new AuditConfigWithReadOnly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditConfigWithReadOnlyWithDefaults

`func NewAuditConfigWithReadOnlyWithDefaults() *AuditConfigWithReadOnly`

NewAuditConfigWithReadOnlyWithDefaults instantiates a new AuditConfigWithReadOnly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReadonly

`func (o *AuditConfigWithReadOnly) GetReadonly() []string`

GetReadonly returns the Readonly field if non-nil, zero value otherwise.

### GetReadonlyOk

`func (o *AuditConfigWithReadOnly) GetReadonlyOk() (*[]string, bool)`

GetReadonlyOk returns a tuple with the Readonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonly

`func (o *AuditConfigWithReadOnly) SetReadonly(v []string)`

SetReadonly sets Readonly field to given value.

### HasReadonly

`func (o *AuditConfigWithReadOnly) HasReadonly() bool`

HasReadonly returns a boolean if a field has been set.

### GetConfig

`func (o *AuditConfigWithReadOnly) GetConfig() AuditConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AuditConfigWithReadOnly) GetConfigOk() (*AuditConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AuditConfigWithReadOnly) SetConfig(v AuditConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AuditConfigWithReadOnly) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


