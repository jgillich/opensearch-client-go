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

// checks if the GetCertificatesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCertificatesResponseContent{}

// GetCertificatesResponseContent struct for GetCertificatesResponseContent
type GetCertificatesResponseContent struct {
	HttpCertificatesList []CertificatesDetail `json:"http_certificates_list,omitempty"`
	TransportCertificatesList []CertificatesDetail `json:"transport_certificates_list,omitempty"`
}

// NewGetCertificatesResponseContent instantiates a new GetCertificatesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCertificatesResponseContent() *GetCertificatesResponseContent {
	this := GetCertificatesResponseContent{}
	return &this
}

// NewGetCertificatesResponseContentWithDefaults instantiates a new GetCertificatesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCertificatesResponseContentWithDefaults() *GetCertificatesResponseContent {
	this := GetCertificatesResponseContent{}
	return &this
}

// GetHttpCertificatesList returns the HttpCertificatesList field value if set, zero value otherwise.
func (o *GetCertificatesResponseContent) GetHttpCertificatesList() []CertificatesDetail {
	if o == nil || IsNil(o.HttpCertificatesList) {
		var ret []CertificatesDetail
		return ret
	}
	return o.HttpCertificatesList
}

// GetHttpCertificatesListOk returns a tuple with the HttpCertificatesList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCertificatesResponseContent) GetHttpCertificatesListOk() ([]CertificatesDetail, bool) {
	if o == nil || IsNil(o.HttpCertificatesList) {
		return nil, false
	}
	return o.HttpCertificatesList, true
}

// HasHttpCertificatesList returns a boolean if a field has been set.
func (o *GetCertificatesResponseContent) HasHttpCertificatesList() bool {
	if o != nil && !IsNil(o.HttpCertificatesList) {
		return true
	}

	return false
}

// SetHttpCertificatesList gets a reference to the given []CertificatesDetail and assigns it to the HttpCertificatesList field.
func (o *GetCertificatesResponseContent) SetHttpCertificatesList(v []CertificatesDetail) {
	o.HttpCertificatesList = v
}

// GetTransportCertificatesList returns the TransportCertificatesList field value if set, zero value otherwise.
func (o *GetCertificatesResponseContent) GetTransportCertificatesList() []CertificatesDetail {
	if o == nil || IsNil(o.TransportCertificatesList) {
		var ret []CertificatesDetail
		return ret
	}
	return o.TransportCertificatesList
}

// GetTransportCertificatesListOk returns a tuple with the TransportCertificatesList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCertificatesResponseContent) GetTransportCertificatesListOk() ([]CertificatesDetail, bool) {
	if o == nil || IsNil(o.TransportCertificatesList) {
		return nil, false
	}
	return o.TransportCertificatesList, true
}

// HasTransportCertificatesList returns a boolean if a field has been set.
func (o *GetCertificatesResponseContent) HasTransportCertificatesList() bool {
	if o != nil && !IsNil(o.TransportCertificatesList) {
		return true
	}

	return false
}

// SetTransportCertificatesList gets a reference to the given []CertificatesDetail and assigns it to the TransportCertificatesList field.
func (o *GetCertificatesResponseContent) SetTransportCertificatesList(v []CertificatesDetail) {
	o.TransportCertificatesList = v
}

func (o GetCertificatesResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCertificatesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HttpCertificatesList) {
		toSerialize["http_certificates_list"] = o.HttpCertificatesList
	}
	if !IsNil(o.TransportCertificatesList) {
		toSerialize["transport_certificates_list"] = o.TransportCertificatesList
	}
	return toSerialize, nil
}

type NullableGetCertificatesResponseContent struct {
	value *GetCertificatesResponseContent
	isSet bool
}

func (v NullableGetCertificatesResponseContent) Get() *GetCertificatesResponseContent {
	return v.value
}

func (v *NullableGetCertificatesResponseContent) Set(val *GetCertificatesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCertificatesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCertificatesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCertificatesResponseContent(val *GetCertificatesResponseContent) *NullableGetCertificatesResponseContent {
	return &NullableGetCertificatesResponseContent{value: val, isSet: true}
}

func (v NullableGetCertificatesResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCertificatesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


