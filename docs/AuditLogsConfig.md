# AuditLogsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IgnoreUsers** | Pointer to **[]string** |  | [optional] 
**IgnoreRequests** | Pointer to **[]string** |  | [optional] 
**DisabledRestCategories** | Pointer to **[]string** |  | [optional] 
**DisabledTransportCategories** | Pointer to **[]string** |  | [optional] 
**LogRequestBody** | Pointer to **bool** |  | [optional] 
**ResolveIndices** | Pointer to **bool** |  | [optional] 
**ResolveBulkRequests** | Pointer to **bool** |  | [optional] 
**ExcludeSensitiveHeaders** | Pointer to **bool** |  | [optional] 
**EnableTransport** | Pointer to **bool** |  | [optional] 
**EnableRest** | Pointer to **bool** |  | [optional] 

## Methods

### NewAuditLogsConfig

`func NewAuditLogsConfig() *AuditLogsConfig`

NewAuditLogsConfig instantiates a new AuditLogsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogsConfigWithDefaults

`func NewAuditLogsConfigWithDefaults() *AuditLogsConfig`

NewAuditLogsConfigWithDefaults instantiates a new AuditLogsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIgnoreUsers

`func (o *AuditLogsConfig) GetIgnoreUsers() []string`

GetIgnoreUsers returns the IgnoreUsers field if non-nil, zero value otherwise.

### GetIgnoreUsersOk

`func (o *AuditLogsConfig) GetIgnoreUsersOk() (*[]string, bool)`

GetIgnoreUsersOk returns a tuple with the IgnoreUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreUsers

`func (o *AuditLogsConfig) SetIgnoreUsers(v []string)`

SetIgnoreUsers sets IgnoreUsers field to given value.

### HasIgnoreUsers

`func (o *AuditLogsConfig) HasIgnoreUsers() bool`

HasIgnoreUsers returns a boolean if a field has been set.

### GetIgnoreRequests

`func (o *AuditLogsConfig) GetIgnoreRequests() []string`

GetIgnoreRequests returns the IgnoreRequests field if non-nil, zero value otherwise.

### GetIgnoreRequestsOk

`func (o *AuditLogsConfig) GetIgnoreRequestsOk() (*[]string, bool)`

GetIgnoreRequestsOk returns a tuple with the IgnoreRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreRequests

`func (o *AuditLogsConfig) SetIgnoreRequests(v []string)`

SetIgnoreRequests sets IgnoreRequests field to given value.

### HasIgnoreRequests

`func (o *AuditLogsConfig) HasIgnoreRequests() bool`

HasIgnoreRequests returns a boolean if a field has been set.

### GetDisabledRestCategories

`func (o *AuditLogsConfig) GetDisabledRestCategories() []string`

GetDisabledRestCategories returns the DisabledRestCategories field if non-nil, zero value otherwise.

### GetDisabledRestCategoriesOk

`func (o *AuditLogsConfig) GetDisabledRestCategoriesOk() (*[]string, bool)`

GetDisabledRestCategoriesOk returns a tuple with the DisabledRestCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledRestCategories

`func (o *AuditLogsConfig) SetDisabledRestCategories(v []string)`

SetDisabledRestCategories sets DisabledRestCategories field to given value.

### HasDisabledRestCategories

`func (o *AuditLogsConfig) HasDisabledRestCategories() bool`

HasDisabledRestCategories returns a boolean if a field has been set.

### GetDisabledTransportCategories

`func (o *AuditLogsConfig) GetDisabledTransportCategories() []string`

GetDisabledTransportCategories returns the DisabledTransportCategories field if non-nil, zero value otherwise.

### GetDisabledTransportCategoriesOk

`func (o *AuditLogsConfig) GetDisabledTransportCategoriesOk() (*[]string, bool)`

GetDisabledTransportCategoriesOk returns a tuple with the DisabledTransportCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledTransportCategories

`func (o *AuditLogsConfig) SetDisabledTransportCategories(v []string)`

SetDisabledTransportCategories sets DisabledTransportCategories field to given value.

### HasDisabledTransportCategories

`func (o *AuditLogsConfig) HasDisabledTransportCategories() bool`

HasDisabledTransportCategories returns a boolean if a field has been set.

### GetLogRequestBody

`func (o *AuditLogsConfig) GetLogRequestBody() bool`

GetLogRequestBody returns the LogRequestBody field if non-nil, zero value otherwise.

### GetLogRequestBodyOk

`func (o *AuditLogsConfig) GetLogRequestBodyOk() (*bool, bool)`

GetLogRequestBodyOk returns a tuple with the LogRequestBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRequestBody

`func (o *AuditLogsConfig) SetLogRequestBody(v bool)`

SetLogRequestBody sets LogRequestBody field to given value.

### HasLogRequestBody

`func (o *AuditLogsConfig) HasLogRequestBody() bool`

HasLogRequestBody returns a boolean if a field has been set.

### GetResolveIndices

`func (o *AuditLogsConfig) GetResolveIndices() bool`

GetResolveIndices returns the ResolveIndices field if non-nil, zero value otherwise.

### GetResolveIndicesOk

`func (o *AuditLogsConfig) GetResolveIndicesOk() (*bool, bool)`

GetResolveIndicesOk returns a tuple with the ResolveIndices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolveIndices

`func (o *AuditLogsConfig) SetResolveIndices(v bool)`

SetResolveIndices sets ResolveIndices field to given value.

### HasResolveIndices

`func (o *AuditLogsConfig) HasResolveIndices() bool`

HasResolveIndices returns a boolean if a field has been set.

### GetResolveBulkRequests

`func (o *AuditLogsConfig) GetResolveBulkRequests() bool`

GetResolveBulkRequests returns the ResolveBulkRequests field if non-nil, zero value otherwise.

### GetResolveBulkRequestsOk

`func (o *AuditLogsConfig) GetResolveBulkRequestsOk() (*bool, bool)`

GetResolveBulkRequestsOk returns a tuple with the ResolveBulkRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolveBulkRequests

`func (o *AuditLogsConfig) SetResolveBulkRequests(v bool)`

SetResolveBulkRequests sets ResolveBulkRequests field to given value.

### HasResolveBulkRequests

`func (o *AuditLogsConfig) HasResolveBulkRequests() bool`

HasResolveBulkRequests returns a boolean if a field has been set.

### GetExcludeSensitiveHeaders

`func (o *AuditLogsConfig) GetExcludeSensitiveHeaders() bool`

GetExcludeSensitiveHeaders returns the ExcludeSensitiveHeaders field if non-nil, zero value otherwise.

### GetExcludeSensitiveHeadersOk

`func (o *AuditLogsConfig) GetExcludeSensitiveHeadersOk() (*bool, bool)`

GetExcludeSensitiveHeadersOk returns a tuple with the ExcludeSensitiveHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeSensitiveHeaders

`func (o *AuditLogsConfig) SetExcludeSensitiveHeaders(v bool)`

SetExcludeSensitiveHeaders sets ExcludeSensitiveHeaders field to given value.

### HasExcludeSensitiveHeaders

`func (o *AuditLogsConfig) HasExcludeSensitiveHeaders() bool`

HasExcludeSensitiveHeaders returns a boolean if a field has been set.

### GetEnableTransport

`func (o *AuditLogsConfig) GetEnableTransport() bool`

GetEnableTransport returns the EnableTransport field if non-nil, zero value otherwise.

### GetEnableTransportOk

`func (o *AuditLogsConfig) GetEnableTransportOk() (*bool, bool)`

GetEnableTransportOk returns a tuple with the EnableTransport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTransport

`func (o *AuditLogsConfig) SetEnableTransport(v bool)`

SetEnableTransport sets EnableTransport field to given value.

### HasEnableTransport

`func (o *AuditLogsConfig) HasEnableTransport() bool`

HasEnableTransport returns a boolean if a field has been set.

### GetEnableRest

`func (o *AuditLogsConfig) GetEnableRest() bool`

GetEnableRest returns the EnableRest field if non-nil, zero value otherwise.

### GetEnableRestOk

`func (o *AuditLogsConfig) GetEnableRestOk() (*bool, bool)`

GetEnableRestOk returns a tuple with the EnableRest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRest

`func (o *AuditLogsConfig) SetEnableRest(v bool)`

SetEnableRest sets EnableRest field to given value.

### HasEnableRest

`func (o *AuditLogsConfig) HasEnableRest() bool`

HasEnableRest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


