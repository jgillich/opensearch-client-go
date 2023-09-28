# DynamicOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilteredAliasMode** | Pointer to **string** |  | [optional] 
**DisableRestAuth** | Pointer to **bool** |  | [optional] 
**DisableIntertransportAuth** | Pointer to **bool** |  | [optional] 
**RespectRequestIndicesOptions** | Pointer to **bool** |  | [optional] 
**Kibana** | Pointer to **interface{}** |  | [optional] 
**Http** | Pointer to **interface{}** |  | [optional] 
**Authc** | Pointer to **interface{}** |  | [optional] 
**Authz** | Pointer to **interface{}** |  | [optional] 
**AuthFailureListeners** | Pointer to **interface{}** |  | [optional] 
**DoNotFailOnForbidden** | Pointer to **bool** |  | [optional] 
**MultiRolespanEnabled** | Pointer to **bool** |  | [optional] 
**HostsResolverMode** | Pointer to **string** |  | [optional] 
**DoNotFailOnForbiddenEmpty** | Pointer to **bool** |  | [optional] 

## Methods

### NewDynamicOptions

`func NewDynamicOptions() *DynamicOptions`

NewDynamicOptions instantiates a new DynamicOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicOptionsWithDefaults

`func NewDynamicOptionsWithDefaults() *DynamicOptions`

NewDynamicOptionsWithDefaults instantiates a new DynamicOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilteredAliasMode

`func (o *DynamicOptions) GetFilteredAliasMode() string`

GetFilteredAliasMode returns the FilteredAliasMode field if non-nil, zero value otherwise.

### GetFilteredAliasModeOk

`func (o *DynamicOptions) GetFilteredAliasModeOk() (*string, bool)`

GetFilteredAliasModeOk returns a tuple with the FilteredAliasMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilteredAliasMode

`func (o *DynamicOptions) SetFilteredAliasMode(v string)`

SetFilteredAliasMode sets FilteredAliasMode field to given value.

### HasFilteredAliasMode

`func (o *DynamicOptions) HasFilteredAliasMode() bool`

HasFilteredAliasMode returns a boolean if a field has been set.

### GetDisableRestAuth

`func (o *DynamicOptions) GetDisableRestAuth() bool`

GetDisableRestAuth returns the DisableRestAuth field if non-nil, zero value otherwise.

### GetDisableRestAuthOk

`func (o *DynamicOptions) GetDisableRestAuthOk() (*bool, bool)`

GetDisableRestAuthOk returns a tuple with the DisableRestAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableRestAuth

`func (o *DynamicOptions) SetDisableRestAuth(v bool)`

SetDisableRestAuth sets DisableRestAuth field to given value.

### HasDisableRestAuth

`func (o *DynamicOptions) HasDisableRestAuth() bool`

HasDisableRestAuth returns a boolean if a field has been set.

### GetDisableIntertransportAuth

`func (o *DynamicOptions) GetDisableIntertransportAuth() bool`

GetDisableIntertransportAuth returns the DisableIntertransportAuth field if non-nil, zero value otherwise.

### GetDisableIntertransportAuthOk

`func (o *DynamicOptions) GetDisableIntertransportAuthOk() (*bool, bool)`

GetDisableIntertransportAuthOk returns a tuple with the DisableIntertransportAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableIntertransportAuth

`func (o *DynamicOptions) SetDisableIntertransportAuth(v bool)`

SetDisableIntertransportAuth sets DisableIntertransportAuth field to given value.

### HasDisableIntertransportAuth

`func (o *DynamicOptions) HasDisableIntertransportAuth() bool`

HasDisableIntertransportAuth returns a boolean if a field has been set.

### GetRespectRequestIndicesOptions

`func (o *DynamicOptions) GetRespectRequestIndicesOptions() bool`

GetRespectRequestIndicesOptions returns the RespectRequestIndicesOptions field if non-nil, zero value otherwise.

### GetRespectRequestIndicesOptionsOk

`func (o *DynamicOptions) GetRespectRequestIndicesOptionsOk() (*bool, bool)`

GetRespectRequestIndicesOptionsOk returns a tuple with the RespectRequestIndicesOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespectRequestIndicesOptions

`func (o *DynamicOptions) SetRespectRequestIndicesOptions(v bool)`

SetRespectRequestIndicesOptions sets RespectRequestIndicesOptions field to given value.

### HasRespectRequestIndicesOptions

`func (o *DynamicOptions) HasRespectRequestIndicesOptions() bool`

HasRespectRequestIndicesOptions returns a boolean if a field has been set.

### GetKibana

`func (o *DynamicOptions) GetKibana() interface{}`

GetKibana returns the Kibana field if non-nil, zero value otherwise.

### GetKibanaOk

`func (o *DynamicOptions) GetKibanaOk() (*interface{}, bool)`

GetKibanaOk returns a tuple with the Kibana field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKibana

`func (o *DynamicOptions) SetKibana(v interface{})`

SetKibana sets Kibana field to given value.

### HasKibana

`func (o *DynamicOptions) HasKibana() bool`

HasKibana returns a boolean if a field has been set.

### SetKibanaNil

`func (o *DynamicOptions) SetKibanaNil(b bool)`

 SetKibanaNil sets the value for Kibana to be an explicit nil

### UnsetKibana
`func (o *DynamicOptions) UnsetKibana()`

UnsetKibana ensures that no value is present for Kibana, not even an explicit nil
### GetHttp

`func (o *DynamicOptions) GetHttp() interface{}`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *DynamicOptions) GetHttpOk() (*interface{}, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *DynamicOptions) SetHttp(v interface{})`

SetHttp sets Http field to given value.

### HasHttp

`func (o *DynamicOptions) HasHttp() bool`

HasHttp returns a boolean if a field has been set.

### SetHttpNil

`func (o *DynamicOptions) SetHttpNil(b bool)`

 SetHttpNil sets the value for Http to be an explicit nil

### UnsetHttp
`func (o *DynamicOptions) UnsetHttp()`

UnsetHttp ensures that no value is present for Http, not even an explicit nil
### GetAuthc

`func (o *DynamicOptions) GetAuthc() interface{}`

GetAuthc returns the Authc field if non-nil, zero value otherwise.

### GetAuthcOk

`func (o *DynamicOptions) GetAuthcOk() (*interface{}, bool)`

GetAuthcOk returns a tuple with the Authc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthc

`func (o *DynamicOptions) SetAuthc(v interface{})`

SetAuthc sets Authc field to given value.

### HasAuthc

`func (o *DynamicOptions) HasAuthc() bool`

HasAuthc returns a boolean if a field has been set.

### SetAuthcNil

`func (o *DynamicOptions) SetAuthcNil(b bool)`

 SetAuthcNil sets the value for Authc to be an explicit nil

### UnsetAuthc
`func (o *DynamicOptions) UnsetAuthc()`

UnsetAuthc ensures that no value is present for Authc, not even an explicit nil
### GetAuthz

`func (o *DynamicOptions) GetAuthz() interface{}`

GetAuthz returns the Authz field if non-nil, zero value otherwise.

### GetAuthzOk

`func (o *DynamicOptions) GetAuthzOk() (*interface{}, bool)`

GetAuthzOk returns a tuple with the Authz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthz

`func (o *DynamicOptions) SetAuthz(v interface{})`

SetAuthz sets Authz field to given value.

### HasAuthz

`func (o *DynamicOptions) HasAuthz() bool`

HasAuthz returns a boolean if a field has been set.

### SetAuthzNil

`func (o *DynamicOptions) SetAuthzNil(b bool)`

 SetAuthzNil sets the value for Authz to be an explicit nil

### UnsetAuthz
`func (o *DynamicOptions) UnsetAuthz()`

UnsetAuthz ensures that no value is present for Authz, not even an explicit nil
### GetAuthFailureListeners

`func (o *DynamicOptions) GetAuthFailureListeners() interface{}`

GetAuthFailureListeners returns the AuthFailureListeners field if non-nil, zero value otherwise.

### GetAuthFailureListenersOk

`func (o *DynamicOptions) GetAuthFailureListenersOk() (*interface{}, bool)`

GetAuthFailureListenersOk returns a tuple with the AuthFailureListeners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthFailureListeners

`func (o *DynamicOptions) SetAuthFailureListeners(v interface{})`

SetAuthFailureListeners sets AuthFailureListeners field to given value.

### HasAuthFailureListeners

`func (o *DynamicOptions) HasAuthFailureListeners() bool`

HasAuthFailureListeners returns a boolean if a field has been set.

### SetAuthFailureListenersNil

`func (o *DynamicOptions) SetAuthFailureListenersNil(b bool)`

 SetAuthFailureListenersNil sets the value for AuthFailureListeners to be an explicit nil

### UnsetAuthFailureListeners
`func (o *DynamicOptions) UnsetAuthFailureListeners()`

UnsetAuthFailureListeners ensures that no value is present for AuthFailureListeners, not even an explicit nil
### GetDoNotFailOnForbidden

`func (o *DynamicOptions) GetDoNotFailOnForbidden() bool`

GetDoNotFailOnForbidden returns the DoNotFailOnForbidden field if non-nil, zero value otherwise.

### GetDoNotFailOnForbiddenOk

`func (o *DynamicOptions) GetDoNotFailOnForbiddenOk() (*bool, bool)`

GetDoNotFailOnForbiddenOk returns a tuple with the DoNotFailOnForbidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotFailOnForbidden

`func (o *DynamicOptions) SetDoNotFailOnForbidden(v bool)`

SetDoNotFailOnForbidden sets DoNotFailOnForbidden field to given value.

### HasDoNotFailOnForbidden

`func (o *DynamicOptions) HasDoNotFailOnForbidden() bool`

HasDoNotFailOnForbidden returns a boolean if a field has been set.

### GetMultiRolespanEnabled

`func (o *DynamicOptions) GetMultiRolespanEnabled() bool`

GetMultiRolespanEnabled returns the MultiRolespanEnabled field if non-nil, zero value otherwise.

### GetMultiRolespanEnabledOk

`func (o *DynamicOptions) GetMultiRolespanEnabledOk() (*bool, bool)`

GetMultiRolespanEnabledOk returns a tuple with the MultiRolespanEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiRolespanEnabled

`func (o *DynamicOptions) SetMultiRolespanEnabled(v bool)`

SetMultiRolespanEnabled sets MultiRolespanEnabled field to given value.

### HasMultiRolespanEnabled

`func (o *DynamicOptions) HasMultiRolespanEnabled() bool`

HasMultiRolespanEnabled returns a boolean if a field has been set.

### GetHostsResolverMode

`func (o *DynamicOptions) GetHostsResolverMode() string`

GetHostsResolverMode returns the HostsResolverMode field if non-nil, zero value otherwise.

### GetHostsResolverModeOk

`func (o *DynamicOptions) GetHostsResolverModeOk() (*string, bool)`

GetHostsResolverModeOk returns a tuple with the HostsResolverMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsResolverMode

`func (o *DynamicOptions) SetHostsResolverMode(v string)`

SetHostsResolverMode sets HostsResolverMode field to given value.

### HasHostsResolverMode

`func (o *DynamicOptions) HasHostsResolverMode() bool`

HasHostsResolverMode returns a boolean if a field has been set.

### GetDoNotFailOnForbiddenEmpty

`func (o *DynamicOptions) GetDoNotFailOnForbiddenEmpty() bool`

GetDoNotFailOnForbiddenEmpty returns the DoNotFailOnForbiddenEmpty field if non-nil, zero value otherwise.

### GetDoNotFailOnForbiddenEmptyOk

`func (o *DynamicOptions) GetDoNotFailOnForbiddenEmptyOk() (*bool, bool)`

GetDoNotFailOnForbiddenEmptyOk returns a tuple with the DoNotFailOnForbiddenEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotFailOnForbiddenEmpty

`func (o *DynamicOptions) SetDoNotFailOnForbiddenEmpty(v bool)`

SetDoNotFailOnForbiddenEmpty sets DoNotFailOnForbiddenEmpty field to given value.

### HasDoNotFailOnForbiddenEmpty

`func (o *DynamicOptions) HasDoNotFailOnForbiddenEmpty() bool`

HasDoNotFailOnForbiddenEmpty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


