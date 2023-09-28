# AuditConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compliance** | Pointer to [**ComplianceConfig**](ComplianceConfig.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Audit** | Pointer to [**AuditLogsConfig**](AuditLogsConfig.md) |  | [optional] 

## Methods

### NewAuditConfig

`func NewAuditConfig() *AuditConfig`

NewAuditConfig instantiates a new AuditConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditConfigWithDefaults

`func NewAuditConfigWithDefaults() *AuditConfig`

NewAuditConfigWithDefaults instantiates a new AuditConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompliance

`func (o *AuditConfig) GetCompliance() ComplianceConfig`

GetCompliance returns the Compliance field if non-nil, zero value otherwise.

### GetComplianceOk

`func (o *AuditConfig) GetComplianceOk() (*ComplianceConfig, bool)`

GetComplianceOk returns a tuple with the Compliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliance

`func (o *AuditConfig) SetCompliance(v ComplianceConfig)`

SetCompliance sets Compliance field to given value.

### HasCompliance

`func (o *AuditConfig) HasCompliance() bool`

HasCompliance returns a boolean if a field has been set.

### GetEnabled

`func (o *AuditConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AuditConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AuditConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AuditConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAudit

`func (o *AuditConfig) GetAudit() AuditLogsConfig`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *AuditConfig) GetAuditOk() (*AuditLogsConfig, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *AuditConfig) SetAudit(v AuditLogsConfig)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *AuditConfig) HasAudit() bool`

HasAudit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


