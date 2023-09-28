# SearchBodyParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocvalueFields** | Pointer to **string** |  | [optional] 
**Explain** | Pointer to **bool** |  | [optional] 
**From** | Pointer to **int32** |  | [optional] 
**SeqNoPrimaryTerm** | Pointer to **bool** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Stats** | Pointer to **string** |  | [optional] 
**TerminateAfter** | Pointer to **int32** |  | [optional] 
**Timeout** | Pointer to [**Time**](Time.md) |  | [optional] 
**Version** | Pointer to **bool** |  | [optional] 
**Fields** | Pointer to **[]string** |  | [optional] 
**MinScore** | Pointer to **int32** |  | [optional] 
**IndicesBoost** | Pointer to **[]interface{}** |  | [optional] 
**Query** | Pointer to [**UserDefinedObjectStructure**](UserDefinedObjectStructure.md) |  | [optional] 

## Methods

### NewSearchBodyParams

`func NewSearchBodyParams() *SearchBodyParams`

NewSearchBodyParams instantiates a new SearchBodyParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchBodyParamsWithDefaults

`func NewSearchBodyParamsWithDefaults() *SearchBodyParams`

NewSearchBodyParamsWithDefaults instantiates a new SearchBodyParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocvalueFields

`func (o *SearchBodyParams) GetDocvalueFields() string`

GetDocvalueFields returns the DocvalueFields field if non-nil, zero value otherwise.

### GetDocvalueFieldsOk

`func (o *SearchBodyParams) GetDocvalueFieldsOk() (*string, bool)`

GetDocvalueFieldsOk returns a tuple with the DocvalueFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocvalueFields

`func (o *SearchBodyParams) SetDocvalueFields(v string)`

SetDocvalueFields sets DocvalueFields field to given value.

### HasDocvalueFields

`func (o *SearchBodyParams) HasDocvalueFields() bool`

HasDocvalueFields returns a boolean if a field has been set.

### GetExplain

`func (o *SearchBodyParams) GetExplain() bool`

GetExplain returns the Explain field if non-nil, zero value otherwise.

### GetExplainOk

`func (o *SearchBodyParams) GetExplainOk() (*bool, bool)`

GetExplainOk returns a tuple with the Explain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplain

`func (o *SearchBodyParams) SetExplain(v bool)`

SetExplain sets Explain field to given value.

### HasExplain

`func (o *SearchBodyParams) HasExplain() bool`

HasExplain returns a boolean if a field has been set.

### GetFrom

`func (o *SearchBodyParams) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SearchBodyParams) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SearchBodyParams) SetFrom(v int32)`

SetFrom sets From field to given value.

### HasFrom

`func (o *SearchBodyParams) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetSeqNoPrimaryTerm

`func (o *SearchBodyParams) GetSeqNoPrimaryTerm() bool`

GetSeqNoPrimaryTerm returns the SeqNoPrimaryTerm field if non-nil, zero value otherwise.

### GetSeqNoPrimaryTermOk

`func (o *SearchBodyParams) GetSeqNoPrimaryTermOk() (*bool, bool)`

GetSeqNoPrimaryTermOk returns a tuple with the SeqNoPrimaryTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeqNoPrimaryTerm

`func (o *SearchBodyParams) SetSeqNoPrimaryTerm(v bool)`

SetSeqNoPrimaryTerm sets SeqNoPrimaryTerm field to given value.

### HasSeqNoPrimaryTerm

`func (o *SearchBodyParams) HasSeqNoPrimaryTerm() bool`

HasSeqNoPrimaryTerm returns a boolean if a field has been set.

### GetSize

`func (o *SearchBodyParams) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SearchBodyParams) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SearchBodyParams) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *SearchBodyParams) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSource

`func (o *SearchBodyParams) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SearchBodyParams) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SearchBodyParams) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *SearchBodyParams) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStats

`func (o *SearchBodyParams) GetStats() string`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *SearchBodyParams) GetStatsOk() (*string, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *SearchBodyParams) SetStats(v string)`

SetStats sets Stats field to given value.

### HasStats

`func (o *SearchBodyParams) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetTerminateAfter

`func (o *SearchBodyParams) GetTerminateAfter() int32`

GetTerminateAfter returns the TerminateAfter field if non-nil, zero value otherwise.

### GetTerminateAfterOk

`func (o *SearchBodyParams) GetTerminateAfterOk() (*int32, bool)`

GetTerminateAfterOk returns a tuple with the TerminateAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminateAfter

`func (o *SearchBodyParams) SetTerminateAfter(v int32)`

SetTerminateAfter sets TerminateAfter field to given value.

### HasTerminateAfter

`func (o *SearchBodyParams) HasTerminateAfter() bool`

HasTerminateAfter returns a boolean if a field has been set.

### GetTimeout

`func (o *SearchBodyParams) GetTimeout() Time`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *SearchBodyParams) GetTimeoutOk() (*Time, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *SearchBodyParams) SetTimeout(v Time)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *SearchBodyParams) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetVersion

`func (o *SearchBodyParams) GetVersion() bool`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SearchBodyParams) GetVersionOk() (*bool, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SearchBodyParams) SetVersion(v bool)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SearchBodyParams) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetFields

`func (o *SearchBodyParams) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *SearchBodyParams) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *SearchBodyParams) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *SearchBodyParams) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetMinScore

`func (o *SearchBodyParams) GetMinScore() int32`

GetMinScore returns the MinScore field if non-nil, zero value otherwise.

### GetMinScoreOk

`func (o *SearchBodyParams) GetMinScoreOk() (*int32, bool)`

GetMinScoreOk returns a tuple with the MinScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinScore

`func (o *SearchBodyParams) SetMinScore(v int32)`

SetMinScore sets MinScore field to given value.

### HasMinScore

`func (o *SearchBodyParams) HasMinScore() bool`

HasMinScore returns a boolean if a field has been set.

### GetIndicesBoost

`func (o *SearchBodyParams) GetIndicesBoost() []interface{}`

GetIndicesBoost returns the IndicesBoost field if non-nil, zero value otherwise.

### GetIndicesBoostOk

`func (o *SearchBodyParams) GetIndicesBoostOk() (*[]interface{}, bool)`

GetIndicesBoostOk returns a tuple with the IndicesBoost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicesBoost

`func (o *SearchBodyParams) SetIndicesBoost(v []interface{})`

SetIndicesBoost sets IndicesBoost field to given value.

### HasIndicesBoost

`func (o *SearchBodyParams) HasIndicesBoost() bool`

HasIndicesBoost returns a boolean if a field has been set.

### GetQuery

`func (o *SearchBodyParams) GetQuery() UserDefinedObjectStructure`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SearchBodyParams) GetQueryOk() (*UserDefinedObjectStructure, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SearchBodyParams) SetQuery(v UserDefinedObjectStructure)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *SearchBodyParams) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


