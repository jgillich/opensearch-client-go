/*
OpenSearch

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2021-11-23
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package opensearch

import (
	"encoding/json"
)

// checks if the DynamicOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DynamicOptions{}

// DynamicOptions struct for DynamicOptions
type DynamicOptions struct {
	FilteredAliasMode *string `json:"filteredAliasMode,omitempty"`
	DisableRestAuth *bool `json:"disableRestAuth,omitempty"`
	DisableIntertransportAuth *bool `json:"disableIntertransportAuth,omitempty"`
	RespectRequestIndicesOptions *bool `json:"respectRequestIndicesOptions,omitempty"`
	Kibana interface{} `json:"kibana,omitempty"`
	Http interface{} `json:"http,omitempty"`
	Authc interface{} `json:"authc,omitempty"`
	Authz interface{} `json:"authz,omitempty"`
	AuthFailureListeners interface{} `json:"authFailureListeners,omitempty"`
	DoNotFailOnForbidden *bool `json:"doNotFailOnForbidden,omitempty"`
	MultiRolespanEnabled *bool `json:"multiRolespanEnabled,omitempty"`
	HostsResolverMode *string `json:"hostsResolverMode,omitempty"`
	DoNotFailOnForbiddenEmpty *bool `json:"doNotFailOnForbiddenEmpty,omitempty"`
}

// NewDynamicOptions instantiates a new DynamicOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicOptions() *DynamicOptions {
	this := DynamicOptions{}
	return &this
}

// NewDynamicOptionsWithDefaults instantiates a new DynamicOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicOptionsWithDefaults() *DynamicOptions {
	this := DynamicOptions{}
	return &this
}

// GetFilteredAliasMode returns the FilteredAliasMode field value if set, zero value otherwise.
func (o *DynamicOptions) GetFilteredAliasMode() string {
	if o == nil || IsNil(o.FilteredAliasMode) {
		var ret string
		return ret
	}
	return *o.FilteredAliasMode
}

// GetFilteredAliasModeOk returns a tuple with the FilteredAliasMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicOptions) GetFilteredAliasModeOk() (*string, bool) {
	if o == nil || IsNil(o.FilteredAliasMode) {
		return nil, false
	}
	return o.FilteredAliasMode, true
}

// HasFilteredAliasMode returns a boolean if a field has been set.
func (o *DynamicOptions) HasFilteredAliasMode() bool {
	if o != nil && !IsNil(o.FilteredAliasMode) {
		return true
	}

	return false
}

// SetFilteredAliasMode gets a reference to the given string and assigns it to the FilteredAliasMode field.
func (o *DynamicOptions) SetFilteredAliasMode(v string) {
	o.FilteredAliasMode = &v
}

// GetDisableRestAuth returns the DisableRestAuth field value if set, zero value otherwise.
func (o *DynamicOptions) GetDisableRestAuth() bool {
	if o == nil || IsNil(o.DisableRestAuth) {
		var ret bool
		return ret
	}
	return *o.DisableRestAuth
}

// GetDisableRestAuthOk returns a tuple with the DisableRestAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicOptions) GetDisableRestAuthOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableRestAuth) {
		return nil, false
	}
	return o.DisableRestAuth, true
}

// HasDisableRestAuth returns a boolean if a field has been set.
func (o *DynamicOptions) HasDisableRestAuth() bool {
	if o != nil && !IsNil(o.DisableRestAuth) {
		return true
	}

	return false
}

// SetDisableRestAuth gets a reference to the given bool and assigns it to the DisableRestAuth field.
func (o *DynamicOptions) SetDisableRestAuth(v bool) {
	o.DisableRestAuth = &v
}

// GetDisableIntertransportAuth returns the DisableIntertransportAuth field value if set, zero value otherwise.
func (o *DynamicOptions) GetDisableIntertransportAuth() bool {
	if o == nil || IsNil(o.DisableIntertransportAuth) {
		var ret bool
		return ret
	}
	return *o.DisableIntertransportAuth
}

// GetDisableIntertransportAuthOk returns a tuple with the DisableIntertransportAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicOptions) GetDisableIntertransportAuthOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableIntertransportAuth) {
		return nil, false
	}
	return o.DisableIntertransportAuth, true
}

// HasDisableIntertransportAuth returns a boolean if a field has been set.
func (o *DynamicOptions) HasDisableIntertransportAuth() bool {
	if o != nil && !IsNil(o.DisableIntertransportAuth) {
		return true
	}

	return false
}

// SetDisableIntertransportAuth gets a reference to the given bool and assigns it to the DisableIntertransportAuth field.
func (o *DynamicOptions) SetDisableIntertransportAuth(v bool) {
	o.DisableIntertransportAuth = &v
}

// GetRespectRequestIndicesOptions returns the RespectRequestIndicesOptions field value if set, zero value otherwise.
func (o *DynamicOptions) GetRespectRequestIndicesOptions() bool {
	if o == nil || IsNil(o.RespectRequestIndicesOptions) {
		var ret bool
		return ret
	}
	return *o.RespectRequestIndicesOptions
}

// GetRespectRequestIndicesOptionsOk returns a tuple with the RespectRequestIndicesOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicOptions) GetRespectRequestIndicesOptionsOk() (*bool, bool) {
	if o == nil || IsNil(o.RespectRequestIndicesOptions) {
		return nil, false
	}
	return o.RespectRequestIndicesOptions, true
}

// HasRespectRequestIndicesOptions returns a boolean if a field has been set.
func (o *DynamicOptions) HasRespectRequestIndicesOptions() bool {
	if o != nil && !IsNil(o.RespectRequestIndicesOptions) {
		return true
	}

	return false
}

// SetRespectRequestIndicesOptions gets a reference to the given bool and assigns it to the RespectRequestIndicesOptions field.
func (o *DynamicOptions) SetRespectRequestIndicesOptions(v bool) {
	o.RespectRequestIndicesOptions = &v
}

// GetKibana returns the Kibana field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DynamicOptions) GetKibana() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Kibana
}

// GetKibanaOk returns a tuple with the Kibana field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DynamicOptions) GetKibanaOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Kibana) {
		return nil, false
	}
	return &o.Kibana, true
}

// HasKibana returns a boolean if a field has been set.
func (o *DynamicOptions) HasKibana() bool {
	if o != nil && IsNil(o.Kibana) {
		return true
	}

	return false
}

// SetKibana gets a reference to the given interface{} and assigns it to the Kibana field.
func (o *DynamicOptions) SetKibana(v interface{}) {
	o.Kibana = v
}

// GetHttp returns the Http field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DynamicOptions) GetHttp() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Http
}

// GetHttpOk returns a tuple with the Http field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DynamicOptions) GetHttpOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Http) {
		return nil, false
	}
	return &o.Http, true
}

// HasHttp returns a boolean if a field has been set.
func (o *DynamicOptions) HasHttp() bool {
	if o != nil && IsNil(o.Http) {
		return true
	}

	return false
}

// SetHttp gets a reference to the given interface{} and assigns it to the Http field.
func (o *DynamicOptions) SetHttp(v interface{}) {
	o.Http = v
}

// GetAuthc returns the Authc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DynamicOptions) GetAuthc() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Authc
}

// GetAuthcOk returns a tuple with the Authc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DynamicOptions) GetAuthcOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Authc) {
		return nil, false
	}
	return &o.Authc, true
}

// HasAuthc returns a boolean if a field has been set.
func (o *DynamicOptions) HasAuthc() bool {
	if o != nil && IsNil(o.Authc) {
		return true
	}

	return false
}

// SetAuthc gets a reference to the given interface{} and assigns it to the Authc field.
func (o *DynamicOptions) SetAuthc(v interface{}) {
	o.Authc = v
}

// GetAuthz returns the Authz field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DynamicOptions) GetAuthz() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Authz
}

// GetAuthzOk returns a tuple with the Authz field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DynamicOptions) GetAuthzOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Authz) {
		return nil, false
	}
	return &o.Authz, true
}

// HasAuthz returns a boolean if a field has been set.
func (o *DynamicOptions) HasAuthz() bool {
	if o != nil && IsNil(o.Authz) {
		return true
	}

	return false
}

// SetAuthz gets a reference to the given interface{} and assigns it to the Authz field.
func (o *DynamicOptions) SetAuthz(v interface{}) {
	o.Authz = v
}

// GetAuthFailureListeners returns the AuthFailureListeners field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DynamicOptions) GetAuthFailureListeners() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AuthFailureListeners
}

// GetAuthFailureListenersOk returns a tuple with the AuthFailureListeners field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DynamicOptions) GetAuthFailureListenersOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AuthFailureListeners) {
		return nil, false
	}
	return &o.AuthFailureListeners, true
}

// HasAuthFailureListeners returns a boolean if a field has been set.
func (o *DynamicOptions) HasAuthFailureListeners() bool {
	if o != nil && IsNil(o.AuthFailureListeners) {
		return true
	}

	return false
}

// SetAuthFailureListeners gets a reference to the given interface{} and assigns it to the AuthFailureListeners field.
func (o *DynamicOptions) SetAuthFailureListeners(v interface{}) {
	o.AuthFailureListeners = v
}

// GetDoNotFailOnForbidden returns the DoNotFailOnForbidden field value if set, zero value otherwise.
func (o *DynamicOptions) GetDoNotFailOnForbidden() bool {
	if o == nil || IsNil(o.DoNotFailOnForbidden) {
		var ret bool
		return ret
	}
	return *o.DoNotFailOnForbidden
}

// GetDoNotFailOnForbiddenOk returns a tuple with the DoNotFailOnForbidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicOptions) GetDoNotFailOnForbiddenOk() (*bool, bool) {
	if o == nil || IsNil(o.DoNotFailOnForbidden) {
		return nil, false
	}
	return o.DoNotFailOnForbidden, true
}

// HasDoNotFailOnForbidden returns a boolean if a field has been set.
func (o *DynamicOptions) HasDoNotFailOnForbidden() bool {
	if o != nil && !IsNil(o.DoNotFailOnForbidden) {
		return true
	}

	return false
}

// SetDoNotFailOnForbidden gets a reference to the given bool and assigns it to the DoNotFailOnForbidden field.
func (o *DynamicOptions) SetDoNotFailOnForbidden(v bool) {
	o.DoNotFailOnForbidden = &v
}

// GetMultiRolespanEnabled returns the MultiRolespanEnabled field value if set, zero value otherwise.
func (o *DynamicOptions) GetMultiRolespanEnabled() bool {
	if o == nil || IsNil(o.MultiRolespanEnabled) {
		var ret bool
		return ret
	}
	return *o.MultiRolespanEnabled
}

// GetMultiRolespanEnabledOk returns a tuple with the MultiRolespanEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicOptions) GetMultiRolespanEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.MultiRolespanEnabled) {
		return nil, false
	}
	return o.MultiRolespanEnabled, true
}

// HasMultiRolespanEnabled returns a boolean if a field has been set.
func (o *DynamicOptions) HasMultiRolespanEnabled() bool {
	if o != nil && !IsNil(o.MultiRolespanEnabled) {
		return true
	}

	return false
}

// SetMultiRolespanEnabled gets a reference to the given bool and assigns it to the MultiRolespanEnabled field.
func (o *DynamicOptions) SetMultiRolespanEnabled(v bool) {
	o.MultiRolespanEnabled = &v
}

// GetHostsResolverMode returns the HostsResolverMode field value if set, zero value otherwise.
func (o *DynamicOptions) GetHostsResolverMode() string {
	if o == nil || IsNil(o.HostsResolverMode) {
		var ret string
		return ret
	}
	return *o.HostsResolverMode
}

// GetHostsResolverModeOk returns a tuple with the HostsResolverMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicOptions) GetHostsResolverModeOk() (*string, bool) {
	if o == nil || IsNil(o.HostsResolverMode) {
		return nil, false
	}
	return o.HostsResolverMode, true
}

// HasHostsResolverMode returns a boolean if a field has been set.
func (o *DynamicOptions) HasHostsResolverMode() bool {
	if o != nil && !IsNil(o.HostsResolverMode) {
		return true
	}

	return false
}

// SetHostsResolverMode gets a reference to the given string and assigns it to the HostsResolverMode field.
func (o *DynamicOptions) SetHostsResolverMode(v string) {
	o.HostsResolverMode = &v
}

// GetDoNotFailOnForbiddenEmpty returns the DoNotFailOnForbiddenEmpty field value if set, zero value otherwise.
func (o *DynamicOptions) GetDoNotFailOnForbiddenEmpty() bool {
	if o == nil || IsNil(o.DoNotFailOnForbiddenEmpty) {
		var ret bool
		return ret
	}
	return *o.DoNotFailOnForbiddenEmpty
}

// GetDoNotFailOnForbiddenEmptyOk returns a tuple with the DoNotFailOnForbiddenEmpty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicOptions) GetDoNotFailOnForbiddenEmptyOk() (*bool, bool) {
	if o == nil || IsNil(o.DoNotFailOnForbiddenEmpty) {
		return nil, false
	}
	return o.DoNotFailOnForbiddenEmpty, true
}

// HasDoNotFailOnForbiddenEmpty returns a boolean if a field has been set.
func (o *DynamicOptions) HasDoNotFailOnForbiddenEmpty() bool {
	if o != nil && !IsNil(o.DoNotFailOnForbiddenEmpty) {
		return true
	}

	return false
}

// SetDoNotFailOnForbiddenEmpty gets a reference to the given bool and assigns it to the DoNotFailOnForbiddenEmpty field.
func (o *DynamicOptions) SetDoNotFailOnForbiddenEmpty(v bool) {
	o.DoNotFailOnForbiddenEmpty = &v
}

func (o DynamicOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DynamicOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FilteredAliasMode) {
		toSerialize["filteredAliasMode"] = o.FilteredAliasMode
	}
	if !IsNil(o.DisableRestAuth) {
		toSerialize["disableRestAuth"] = o.DisableRestAuth
	}
	if !IsNil(o.DisableIntertransportAuth) {
		toSerialize["disableIntertransportAuth"] = o.DisableIntertransportAuth
	}
	if !IsNil(o.RespectRequestIndicesOptions) {
		toSerialize["respectRequestIndicesOptions"] = o.RespectRequestIndicesOptions
	}
	if o.Kibana != nil {
		toSerialize["kibana"] = o.Kibana
	}
	if o.Http != nil {
		toSerialize["http"] = o.Http
	}
	if o.Authc != nil {
		toSerialize["authc"] = o.Authc
	}
	if o.Authz != nil {
		toSerialize["authz"] = o.Authz
	}
	if o.AuthFailureListeners != nil {
		toSerialize["authFailureListeners"] = o.AuthFailureListeners
	}
	if !IsNil(o.DoNotFailOnForbidden) {
		toSerialize["doNotFailOnForbidden"] = o.DoNotFailOnForbidden
	}
	if !IsNil(o.MultiRolespanEnabled) {
		toSerialize["multiRolespanEnabled"] = o.MultiRolespanEnabled
	}
	if !IsNil(o.HostsResolverMode) {
		toSerialize["hostsResolverMode"] = o.HostsResolverMode
	}
	if !IsNil(o.DoNotFailOnForbiddenEmpty) {
		toSerialize["doNotFailOnForbiddenEmpty"] = o.DoNotFailOnForbiddenEmpty
	}
	return toSerialize, nil
}

type NullableDynamicOptions struct {
	value *DynamicOptions
	isSet bool
}

func (v NullableDynamicOptions) Get() *DynamicOptions {
	return v.value
}

func (v *NullableDynamicOptions) Set(val *DynamicOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicOptions(val *DynamicOptions) *NullableDynamicOptions {
	return &NullableDynamicOptions{value: val, isSet: true}
}

func (v NullableDynamicOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

