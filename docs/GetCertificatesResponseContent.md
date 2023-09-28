# GetCertificatesResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpCertificatesList** | Pointer to [**[]CertificatesDetail**](CertificatesDetail.md) |  | [optional] 
**TransportCertificatesList** | Pointer to [**[]CertificatesDetail**](CertificatesDetail.md) |  | [optional] 

## Methods

### NewGetCertificatesResponseContent

`func NewGetCertificatesResponseContent() *GetCertificatesResponseContent`

NewGetCertificatesResponseContent instantiates a new GetCertificatesResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCertificatesResponseContentWithDefaults

`func NewGetCertificatesResponseContentWithDefaults() *GetCertificatesResponseContent`

NewGetCertificatesResponseContentWithDefaults instantiates a new GetCertificatesResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpCertificatesList

`func (o *GetCertificatesResponseContent) GetHttpCertificatesList() []CertificatesDetail`

GetHttpCertificatesList returns the HttpCertificatesList field if non-nil, zero value otherwise.

### GetHttpCertificatesListOk

`func (o *GetCertificatesResponseContent) GetHttpCertificatesListOk() (*[]CertificatesDetail, bool)`

GetHttpCertificatesListOk returns a tuple with the HttpCertificatesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpCertificatesList

`func (o *GetCertificatesResponseContent) SetHttpCertificatesList(v []CertificatesDetail)`

SetHttpCertificatesList sets HttpCertificatesList field to given value.

### HasHttpCertificatesList

`func (o *GetCertificatesResponseContent) HasHttpCertificatesList() bool`

HasHttpCertificatesList returns a boolean if a field has been set.

### GetTransportCertificatesList

`func (o *GetCertificatesResponseContent) GetTransportCertificatesList() []CertificatesDetail`

GetTransportCertificatesList returns the TransportCertificatesList field if non-nil, zero value otherwise.

### GetTransportCertificatesListOk

`func (o *GetCertificatesResponseContent) GetTransportCertificatesListOk() (*[]CertificatesDetail, bool)`

GetTransportCertificatesListOk returns a tuple with the TransportCertificatesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportCertificatesList

`func (o *GetCertificatesResponseContent) SetTransportCertificatesList(v []CertificatesDetail)`

SetTransportCertificatesList sets TransportCertificatesList field to given value.

### HasTransportCertificatesList

`func (o *GetCertificatesResponseContent) HasTransportCertificatesList() bool`

HasTransportCertificatesList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


