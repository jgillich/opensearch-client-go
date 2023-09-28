# CatPitSegment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **string** |  | [optional] 
**Shard** | Pointer to **int32** |  | [optional] 
**Prirep** | Pointer to **bool** | Set to true to return stats only for primary shards. | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**Segment** | Pointer to **string** |  | [optional] 
**Generation** | Pointer to **int32** |  | [optional] 
**DocsCount** | Pointer to **int32** |  | [optional] 
**DocsDeleted** | Pointer to **int32** |  | [optional] 
**Size** | Pointer to **string** |  | [optional] 
**SizeMemory** | Pointer to **int32** |  | [optional] 
**Committed** | Pointer to **bool** |  | [optional] 
**Searchable** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Compound** | Pointer to **bool** |  | [optional] 

## Methods

### NewCatPitSegment

`func NewCatPitSegment() *CatPitSegment`

NewCatPitSegment instantiates a new CatPitSegment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatPitSegmentWithDefaults

`func NewCatPitSegmentWithDefaults() *CatPitSegment`

NewCatPitSegmentWithDefaults instantiates a new CatPitSegment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *CatPitSegment) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *CatPitSegment) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *CatPitSegment) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *CatPitSegment) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetShard

`func (o *CatPitSegment) GetShard() int32`

GetShard returns the Shard field if non-nil, zero value otherwise.

### GetShardOk

`func (o *CatPitSegment) GetShardOk() (*int32, bool)`

GetShardOk returns a tuple with the Shard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShard

`func (o *CatPitSegment) SetShard(v int32)`

SetShard sets Shard field to given value.

### HasShard

`func (o *CatPitSegment) HasShard() bool`

HasShard returns a boolean if a field has been set.

### GetPrirep

`func (o *CatPitSegment) GetPrirep() bool`

GetPrirep returns the Prirep field if non-nil, zero value otherwise.

### GetPrirepOk

`func (o *CatPitSegment) GetPrirepOk() (*bool, bool)`

GetPrirepOk returns a tuple with the Prirep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrirep

`func (o *CatPitSegment) SetPrirep(v bool)`

SetPrirep sets Prirep field to given value.

### HasPrirep

`func (o *CatPitSegment) HasPrirep() bool`

HasPrirep returns a boolean if a field has been set.

### GetIp

`func (o *CatPitSegment) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *CatPitSegment) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *CatPitSegment) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *CatPitSegment) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetSegment

`func (o *CatPitSegment) GetSegment() string`

GetSegment returns the Segment field if non-nil, zero value otherwise.

### GetSegmentOk

`func (o *CatPitSegment) GetSegmentOk() (*string, bool)`

GetSegmentOk returns a tuple with the Segment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment

`func (o *CatPitSegment) SetSegment(v string)`

SetSegment sets Segment field to given value.

### HasSegment

`func (o *CatPitSegment) HasSegment() bool`

HasSegment returns a boolean if a field has been set.

### GetGeneration

`func (o *CatPitSegment) GetGeneration() int32`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *CatPitSegment) GetGenerationOk() (*int32, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *CatPitSegment) SetGeneration(v int32)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *CatPitSegment) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.

### GetDocsCount

`func (o *CatPitSegment) GetDocsCount() int32`

GetDocsCount returns the DocsCount field if non-nil, zero value otherwise.

### GetDocsCountOk

`func (o *CatPitSegment) GetDocsCountOk() (*int32, bool)`

GetDocsCountOk returns a tuple with the DocsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocsCount

`func (o *CatPitSegment) SetDocsCount(v int32)`

SetDocsCount sets DocsCount field to given value.

### HasDocsCount

`func (o *CatPitSegment) HasDocsCount() bool`

HasDocsCount returns a boolean if a field has been set.

### GetDocsDeleted

`func (o *CatPitSegment) GetDocsDeleted() int32`

GetDocsDeleted returns the DocsDeleted field if non-nil, zero value otherwise.

### GetDocsDeletedOk

`func (o *CatPitSegment) GetDocsDeletedOk() (*int32, bool)`

GetDocsDeletedOk returns a tuple with the DocsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocsDeleted

`func (o *CatPitSegment) SetDocsDeleted(v int32)`

SetDocsDeleted sets DocsDeleted field to given value.

### HasDocsDeleted

`func (o *CatPitSegment) HasDocsDeleted() bool`

HasDocsDeleted returns a boolean if a field has been set.

### GetSize

`func (o *CatPitSegment) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CatPitSegment) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CatPitSegment) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *CatPitSegment) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSizeMemory

`func (o *CatPitSegment) GetSizeMemory() int32`

GetSizeMemory returns the SizeMemory field if non-nil, zero value otherwise.

### GetSizeMemoryOk

`func (o *CatPitSegment) GetSizeMemoryOk() (*int32, bool)`

GetSizeMemoryOk returns a tuple with the SizeMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeMemory

`func (o *CatPitSegment) SetSizeMemory(v int32)`

SetSizeMemory sets SizeMemory field to given value.

### HasSizeMemory

`func (o *CatPitSegment) HasSizeMemory() bool`

HasSizeMemory returns a boolean if a field has been set.

### GetCommitted

`func (o *CatPitSegment) GetCommitted() bool`

GetCommitted returns the Committed field if non-nil, zero value otherwise.

### GetCommittedOk

`func (o *CatPitSegment) GetCommittedOk() (*bool, bool)`

GetCommittedOk returns a tuple with the Committed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitted

`func (o *CatPitSegment) SetCommitted(v bool)`

SetCommitted sets Committed field to given value.

### HasCommitted

`func (o *CatPitSegment) HasCommitted() bool`

HasCommitted returns a boolean if a field has been set.

### GetSearchable

`func (o *CatPitSegment) GetSearchable() bool`

GetSearchable returns the Searchable field if non-nil, zero value otherwise.

### GetSearchableOk

`func (o *CatPitSegment) GetSearchableOk() (*bool, bool)`

GetSearchableOk returns a tuple with the Searchable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchable

`func (o *CatPitSegment) SetSearchable(v bool)`

SetSearchable sets Searchable field to given value.

### HasSearchable

`func (o *CatPitSegment) HasSearchable() bool`

HasSearchable returns a boolean if a field has been set.

### GetVersion

`func (o *CatPitSegment) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CatPitSegment) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CatPitSegment) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CatPitSegment) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCompound

`func (o *CatPitSegment) GetCompound() bool`

GetCompound returns the Compound field if non-nil, zero value otherwise.

### GetCompoundOk

`func (o *CatPitSegment) GetCompoundOk() (*bool, bool)`

GetCompoundOk returns a tuple with the Compound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompound

`func (o *CatPitSegment) SetCompound(v bool)`

SetCompound sets Compound field to given value.

### HasCompound

`func (o *CatPitSegment) HasCompound() bool`

HasCompound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


