# UserDefinedObjectStructure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bool** | Pointer to **interface{}** |  | [optional] 
**Boosting** | Pointer to **interface{}** |  | [optional] 
**CombinedFields** | Pointer to **interface{}** |  | [optional] 
**ConstantScore** | Pointer to **interface{}** |  | [optional] 
**DisMax** | Pointer to **interface{}** |  | [optional] 
**DistanceFeature** | Pointer to **interface{}** |  | [optional] 
**Exists** | Pointer to **interface{}** |  | [optional] 
**FunctionScore** | Pointer to **interface{}** |  | [optional] 
**Fuzzy** | Pointer to **map[string]interface{}** |  | [optional] 
**GeoBoundingBox** | Pointer to **interface{}** |  | [optional] 
**GeoDistance** | Pointer to **interface{}** |  | [optional] 
**GeoPolygon** | Pointer to **interface{}** |  | [optional] 
**GeoShape** | Pointer to **interface{}** |  | [optional] 
**HasChild** | Pointer to **interface{}** |  | [optional] 
**HasParent** | Pointer to **interface{}** |  | [optional] 
**Ids** | Pointer to **interface{}** |  | [optional] 
**Intervals** | Pointer to **map[string]interface{}** |  | [optional] 
**Knn** | Pointer to **interface{}** |  | [optional] 
**Match** | Pointer to **map[string]interface{}** |  | [optional] 
**MatchAll** | Pointer to **interface{}** |  | [optional] 
**MatchBoolPrefix** | Pointer to **map[string]interface{}** |  | [optional] 
**MatchNone** | Pointer to **interface{}** |  | [optional] 
**MatchPhrase** | Pointer to **map[string]interface{}** |  | [optional] 
**MatchPhrasePrefix** | Pointer to **map[string]interface{}** |  | [optional] 
**MoreLikeThis** | Pointer to **interface{}** |  | [optional] 
**MultiMatch** | Pointer to **interface{}** |  | [optional] 
**Nested** | Pointer to **interface{}** |  | [optional] 
**ParentId** | Pointer to **interface{}** |  | [optional] 
**Percolate** | Pointer to **interface{}** |  | [optional] 
**Pinned** | Pointer to **interface{}** |  | [optional] 
**Prefix** | Pointer to **map[string]interface{}** |  | [optional] 
**QueryString** | Pointer to **interface{}** |  | [optional] 
**Range** | Pointer to **map[string]interface{}** |  | [optional] 
**RankFeature** | Pointer to **interface{}** |  | [optional] 
**Regexp** | Pointer to **map[string]interface{}** |  | [optional] 
**Script** | Pointer to **interface{}** |  | [optional] 
**ScriptScore** | Pointer to **interface{}** |  | [optional] 
**Shape** | Pointer to **interface{}** |  | [optional] 
**SimpleQueryString** | Pointer to **interface{}** |  | [optional] 
**SpanContaining** | Pointer to **interface{}** |  | [optional] 
**FieldMaskingSpan** | Pointer to **interface{}** |  | [optional] 
**SpanFirst** | Pointer to **interface{}** |  | [optional] 
**SpanMulti** | Pointer to **interface{}** |  | [optional] 
**SpanNear** | Pointer to **interface{}** |  | [optional] 
**SpanNot** | Pointer to **interface{}** |  | [optional] 
**SpanOr** | Pointer to **interface{}** |  | [optional] 
**SpanTerm** | Pointer to **map[string]interface{}** |  | [optional] 
**SpanWithin** | Pointer to **interface{}** |  | [optional] 
**Term** | Pointer to **map[string]interface{}** |  | [optional] 
**Terms** | Pointer to **interface{}** |  | [optional] 
**TermsSet** | Pointer to **map[string]interface{}** |  | [optional] 
**Wildcard** | Pointer to **map[string]interface{}** |  | [optional] 
**Wrapper** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewUserDefinedObjectStructure

`func NewUserDefinedObjectStructure() *UserDefinedObjectStructure`

NewUserDefinedObjectStructure instantiates a new UserDefinedObjectStructure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDefinedObjectStructureWithDefaults

`func NewUserDefinedObjectStructureWithDefaults() *UserDefinedObjectStructure`

NewUserDefinedObjectStructureWithDefaults instantiates a new UserDefinedObjectStructure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBool

`func (o *UserDefinedObjectStructure) GetBool() interface{}`

GetBool returns the Bool field if non-nil, zero value otherwise.

### GetBoolOk

`func (o *UserDefinedObjectStructure) GetBoolOk() (*interface{}, bool)`

GetBoolOk returns a tuple with the Bool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBool

`func (o *UserDefinedObjectStructure) SetBool(v interface{})`

SetBool sets Bool field to given value.

### HasBool

`func (o *UserDefinedObjectStructure) HasBool() bool`

HasBool returns a boolean if a field has been set.

### SetBoolNil

`func (o *UserDefinedObjectStructure) SetBoolNil(b bool)`

 SetBoolNil sets the value for Bool to be an explicit nil

### UnsetBool
`func (o *UserDefinedObjectStructure) UnsetBool()`

UnsetBool ensures that no value is present for Bool, not even an explicit nil
### GetBoosting

`func (o *UserDefinedObjectStructure) GetBoosting() interface{}`

GetBoosting returns the Boosting field if non-nil, zero value otherwise.

### GetBoostingOk

`func (o *UserDefinedObjectStructure) GetBoostingOk() (*interface{}, bool)`

GetBoostingOk returns a tuple with the Boosting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoosting

`func (o *UserDefinedObjectStructure) SetBoosting(v interface{})`

SetBoosting sets Boosting field to given value.

### HasBoosting

`func (o *UserDefinedObjectStructure) HasBoosting() bool`

HasBoosting returns a boolean if a field has been set.

### SetBoostingNil

`func (o *UserDefinedObjectStructure) SetBoostingNil(b bool)`

 SetBoostingNil sets the value for Boosting to be an explicit nil

### UnsetBoosting
`func (o *UserDefinedObjectStructure) UnsetBoosting()`

UnsetBoosting ensures that no value is present for Boosting, not even an explicit nil
### GetCombinedFields

`func (o *UserDefinedObjectStructure) GetCombinedFields() interface{}`

GetCombinedFields returns the CombinedFields field if non-nil, zero value otherwise.

### GetCombinedFieldsOk

`func (o *UserDefinedObjectStructure) GetCombinedFieldsOk() (*interface{}, bool)`

GetCombinedFieldsOk returns a tuple with the CombinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedFields

`func (o *UserDefinedObjectStructure) SetCombinedFields(v interface{})`

SetCombinedFields sets CombinedFields field to given value.

### HasCombinedFields

`func (o *UserDefinedObjectStructure) HasCombinedFields() bool`

HasCombinedFields returns a boolean if a field has been set.

### SetCombinedFieldsNil

`func (o *UserDefinedObjectStructure) SetCombinedFieldsNil(b bool)`

 SetCombinedFieldsNil sets the value for CombinedFields to be an explicit nil

### UnsetCombinedFields
`func (o *UserDefinedObjectStructure) UnsetCombinedFields()`

UnsetCombinedFields ensures that no value is present for CombinedFields, not even an explicit nil
### GetConstantScore

`func (o *UserDefinedObjectStructure) GetConstantScore() interface{}`

GetConstantScore returns the ConstantScore field if non-nil, zero value otherwise.

### GetConstantScoreOk

`func (o *UserDefinedObjectStructure) GetConstantScoreOk() (*interface{}, bool)`

GetConstantScoreOk returns a tuple with the ConstantScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstantScore

`func (o *UserDefinedObjectStructure) SetConstantScore(v interface{})`

SetConstantScore sets ConstantScore field to given value.

### HasConstantScore

`func (o *UserDefinedObjectStructure) HasConstantScore() bool`

HasConstantScore returns a boolean if a field has been set.

### SetConstantScoreNil

`func (o *UserDefinedObjectStructure) SetConstantScoreNil(b bool)`

 SetConstantScoreNil sets the value for ConstantScore to be an explicit nil

### UnsetConstantScore
`func (o *UserDefinedObjectStructure) UnsetConstantScore()`

UnsetConstantScore ensures that no value is present for ConstantScore, not even an explicit nil
### GetDisMax

`func (o *UserDefinedObjectStructure) GetDisMax() interface{}`

GetDisMax returns the DisMax field if non-nil, zero value otherwise.

### GetDisMaxOk

`func (o *UserDefinedObjectStructure) GetDisMaxOk() (*interface{}, bool)`

GetDisMaxOk returns a tuple with the DisMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisMax

`func (o *UserDefinedObjectStructure) SetDisMax(v interface{})`

SetDisMax sets DisMax field to given value.

### HasDisMax

`func (o *UserDefinedObjectStructure) HasDisMax() bool`

HasDisMax returns a boolean if a field has been set.

### SetDisMaxNil

`func (o *UserDefinedObjectStructure) SetDisMaxNil(b bool)`

 SetDisMaxNil sets the value for DisMax to be an explicit nil

### UnsetDisMax
`func (o *UserDefinedObjectStructure) UnsetDisMax()`

UnsetDisMax ensures that no value is present for DisMax, not even an explicit nil
### GetDistanceFeature

`func (o *UserDefinedObjectStructure) GetDistanceFeature() interface{}`

GetDistanceFeature returns the DistanceFeature field if non-nil, zero value otherwise.

### GetDistanceFeatureOk

`func (o *UserDefinedObjectStructure) GetDistanceFeatureOk() (*interface{}, bool)`

GetDistanceFeatureOk returns a tuple with the DistanceFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceFeature

`func (o *UserDefinedObjectStructure) SetDistanceFeature(v interface{})`

SetDistanceFeature sets DistanceFeature field to given value.

### HasDistanceFeature

`func (o *UserDefinedObjectStructure) HasDistanceFeature() bool`

HasDistanceFeature returns a boolean if a field has been set.

### SetDistanceFeatureNil

`func (o *UserDefinedObjectStructure) SetDistanceFeatureNil(b bool)`

 SetDistanceFeatureNil sets the value for DistanceFeature to be an explicit nil

### UnsetDistanceFeature
`func (o *UserDefinedObjectStructure) UnsetDistanceFeature()`

UnsetDistanceFeature ensures that no value is present for DistanceFeature, not even an explicit nil
### GetExists

`func (o *UserDefinedObjectStructure) GetExists() interface{}`

GetExists returns the Exists field if non-nil, zero value otherwise.

### GetExistsOk

`func (o *UserDefinedObjectStructure) GetExistsOk() (*interface{}, bool)`

GetExistsOk returns a tuple with the Exists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExists

`func (o *UserDefinedObjectStructure) SetExists(v interface{})`

SetExists sets Exists field to given value.

### HasExists

`func (o *UserDefinedObjectStructure) HasExists() bool`

HasExists returns a boolean if a field has been set.

### SetExistsNil

`func (o *UserDefinedObjectStructure) SetExistsNil(b bool)`

 SetExistsNil sets the value for Exists to be an explicit nil

### UnsetExists
`func (o *UserDefinedObjectStructure) UnsetExists()`

UnsetExists ensures that no value is present for Exists, not even an explicit nil
### GetFunctionScore

`func (o *UserDefinedObjectStructure) GetFunctionScore() interface{}`

GetFunctionScore returns the FunctionScore field if non-nil, zero value otherwise.

### GetFunctionScoreOk

`func (o *UserDefinedObjectStructure) GetFunctionScoreOk() (*interface{}, bool)`

GetFunctionScoreOk returns a tuple with the FunctionScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionScore

`func (o *UserDefinedObjectStructure) SetFunctionScore(v interface{})`

SetFunctionScore sets FunctionScore field to given value.

### HasFunctionScore

`func (o *UserDefinedObjectStructure) HasFunctionScore() bool`

HasFunctionScore returns a boolean if a field has been set.

### SetFunctionScoreNil

`func (o *UserDefinedObjectStructure) SetFunctionScoreNil(b bool)`

 SetFunctionScoreNil sets the value for FunctionScore to be an explicit nil

### UnsetFunctionScore
`func (o *UserDefinedObjectStructure) UnsetFunctionScore()`

UnsetFunctionScore ensures that no value is present for FunctionScore, not even an explicit nil
### GetFuzzy

`func (o *UserDefinedObjectStructure) GetFuzzy() map[string]interface{}`

GetFuzzy returns the Fuzzy field if non-nil, zero value otherwise.

### GetFuzzyOk

`func (o *UserDefinedObjectStructure) GetFuzzyOk() (*map[string]interface{}, bool)`

GetFuzzyOk returns a tuple with the Fuzzy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuzzy

`func (o *UserDefinedObjectStructure) SetFuzzy(v map[string]interface{})`

SetFuzzy sets Fuzzy field to given value.

### HasFuzzy

`func (o *UserDefinedObjectStructure) HasFuzzy() bool`

HasFuzzy returns a boolean if a field has been set.

### GetGeoBoundingBox

`func (o *UserDefinedObjectStructure) GetGeoBoundingBox() interface{}`

GetGeoBoundingBox returns the GeoBoundingBox field if non-nil, zero value otherwise.

### GetGeoBoundingBoxOk

`func (o *UserDefinedObjectStructure) GetGeoBoundingBoxOk() (*interface{}, bool)`

GetGeoBoundingBoxOk returns a tuple with the GeoBoundingBox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoBoundingBox

`func (o *UserDefinedObjectStructure) SetGeoBoundingBox(v interface{})`

SetGeoBoundingBox sets GeoBoundingBox field to given value.

### HasGeoBoundingBox

`func (o *UserDefinedObjectStructure) HasGeoBoundingBox() bool`

HasGeoBoundingBox returns a boolean if a field has been set.

### SetGeoBoundingBoxNil

`func (o *UserDefinedObjectStructure) SetGeoBoundingBoxNil(b bool)`

 SetGeoBoundingBoxNil sets the value for GeoBoundingBox to be an explicit nil

### UnsetGeoBoundingBox
`func (o *UserDefinedObjectStructure) UnsetGeoBoundingBox()`

UnsetGeoBoundingBox ensures that no value is present for GeoBoundingBox, not even an explicit nil
### GetGeoDistance

`func (o *UserDefinedObjectStructure) GetGeoDistance() interface{}`

GetGeoDistance returns the GeoDistance field if non-nil, zero value otherwise.

### GetGeoDistanceOk

`func (o *UserDefinedObjectStructure) GetGeoDistanceOk() (*interface{}, bool)`

GetGeoDistanceOk returns a tuple with the GeoDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoDistance

`func (o *UserDefinedObjectStructure) SetGeoDistance(v interface{})`

SetGeoDistance sets GeoDistance field to given value.

### HasGeoDistance

`func (o *UserDefinedObjectStructure) HasGeoDistance() bool`

HasGeoDistance returns a boolean if a field has been set.

### SetGeoDistanceNil

`func (o *UserDefinedObjectStructure) SetGeoDistanceNil(b bool)`

 SetGeoDistanceNil sets the value for GeoDistance to be an explicit nil

### UnsetGeoDistance
`func (o *UserDefinedObjectStructure) UnsetGeoDistance()`

UnsetGeoDistance ensures that no value is present for GeoDistance, not even an explicit nil
### GetGeoPolygon

`func (o *UserDefinedObjectStructure) GetGeoPolygon() interface{}`

GetGeoPolygon returns the GeoPolygon field if non-nil, zero value otherwise.

### GetGeoPolygonOk

`func (o *UserDefinedObjectStructure) GetGeoPolygonOk() (*interface{}, bool)`

GetGeoPolygonOk returns a tuple with the GeoPolygon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoPolygon

`func (o *UserDefinedObjectStructure) SetGeoPolygon(v interface{})`

SetGeoPolygon sets GeoPolygon field to given value.

### HasGeoPolygon

`func (o *UserDefinedObjectStructure) HasGeoPolygon() bool`

HasGeoPolygon returns a boolean if a field has been set.

### SetGeoPolygonNil

`func (o *UserDefinedObjectStructure) SetGeoPolygonNil(b bool)`

 SetGeoPolygonNil sets the value for GeoPolygon to be an explicit nil

### UnsetGeoPolygon
`func (o *UserDefinedObjectStructure) UnsetGeoPolygon()`

UnsetGeoPolygon ensures that no value is present for GeoPolygon, not even an explicit nil
### GetGeoShape

`func (o *UserDefinedObjectStructure) GetGeoShape() interface{}`

GetGeoShape returns the GeoShape field if non-nil, zero value otherwise.

### GetGeoShapeOk

`func (o *UserDefinedObjectStructure) GetGeoShapeOk() (*interface{}, bool)`

GetGeoShapeOk returns a tuple with the GeoShape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoShape

`func (o *UserDefinedObjectStructure) SetGeoShape(v interface{})`

SetGeoShape sets GeoShape field to given value.

### HasGeoShape

`func (o *UserDefinedObjectStructure) HasGeoShape() bool`

HasGeoShape returns a boolean if a field has been set.

### SetGeoShapeNil

`func (o *UserDefinedObjectStructure) SetGeoShapeNil(b bool)`

 SetGeoShapeNil sets the value for GeoShape to be an explicit nil

### UnsetGeoShape
`func (o *UserDefinedObjectStructure) UnsetGeoShape()`

UnsetGeoShape ensures that no value is present for GeoShape, not even an explicit nil
### GetHasChild

`func (o *UserDefinedObjectStructure) GetHasChild() interface{}`

GetHasChild returns the HasChild field if non-nil, zero value otherwise.

### GetHasChildOk

`func (o *UserDefinedObjectStructure) GetHasChildOk() (*interface{}, bool)`

GetHasChildOk returns a tuple with the HasChild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasChild

`func (o *UserDefinedObjectStructure) SetHasChild(v interface{})`

SetHasChild sets HasChild field to given value.

### HasHasChild

`func (o *UserDefinedObjectStructure) HasHasChild() bool`

HasHasChild returns a boolean if a field has been set.

### SetHasChildNil

`func (o *UserDefinedObjectStructure) SetHasChildNil(b bool)`

 SetHasChildNil sets the value for HasChild to be an explicit nil

### UnsetHasChild
`func (o *UserDefinedObjectStructure) UnsetHasChild()`

UnsetHasChild ensures that no value is present for HasChild, not even an explicit nil
### GetHasParent

`func (o *UserDefinedObjectStructure) GetHasParent() interface{}`

GetHasParent returns the HasParent field if non-nil, zero value otherwise.

### GetHasParentOk

`func (o *UserDefinedObjectStructure) GetHasParentOk() (*interface{}, bool)`

GetHasParentOk returns a tuple with the HasParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasParent

`func (o *UserDefinedObjectStructure) SetHasParent(v interface{})`

SetHasParent sets HasParent field to given value.

### HasHasParent

`func (o *UserDefinedObjectStructure) HasHasParent() bool`

HasHasParent returns a boolean if a field has been set.

### SetHasParentNil

`func (o *UserDefinedObjectStructure) SetHasParentNil(b bool)`

 SetHasParentNil sets the value for HasParent to be an explicit nil

### UnsetHasParent
`func (o *UserDefinedObjectStructure) UnsetHasParent()`

UnsetHasParent ensures that no value is present for HasParent, not even an explicit nil
### GetIds

`func (o *UserDefinedObjectStructure) GetIds() interface{}`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *UserDefinedObjectStructure) GetIdsOk() (*interface{}, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *UserDefinedObjectStructure) SetIds(v interface{})`

SetIds sets Ids field to given value.

### HasIds

`func (o *UserDefinedObjectStructure) HasIds() bool`

HasIds returns a boolean if a field has been set.

### SetIdsNil

`func (o *UserDefinedObjectStructure) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *UserDefinedObjectStructure) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil
### GetIntervals

`func (o *UserDefinedObjectStructure) GetIntervals() map[string]interface{}`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *UserDefinedObjectStructure) GetIntervalsOk() (*map[string]interface{}, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *UserDefinedObjectStructure) SetIntervals(v map[string]interface{})`

SetIntervals sets Intervals field to given value.

### HasIntervals

`func (o *UserDefinedObjectStructure) HasIntervals() bool`

HasIntervals returns a boolean if a field has been set.

### GetKnn

`func (o *UserDefinedObjectStructure) GetKnn() interface{}`

GetKnn returns the Knn field if non-nil, zero value otherwise.

### GetKnnOk

`func (o *UserDefinedObjectStructure) GetKnnOk() (*interface{}, bool)`

GetKnnOk returns a tuple with the Knn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnn

`func (o *UserDefinedObjectStructure) SetKnn(v interface{})`

SetKnn sets Knn field to given value.

### HasKnn

`func (o *UserDefinedObjectStructure) HasKnn() bool`

HasKnn returns a boolean if a field has been set.

### SetKnnNil

`func (o *UserDefinedObjectStructure) SetKnnNil(b bool)`

 SetKnnNil sets the value for Knn to be an explicit nil

### UnsetKnn
`func (o *UserDefinedObjectStructure) UnsetKnn()`

UnsetKnn ensures that no value is present for Knn, not even an explicit nil
### GetMatch

`func (o *UserDefinedObjectStructure) GetMatch() map[string]interface{}`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *UserDefinedObjectStructure) GetMatchOk() (*map[string]interface{}, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *UserDefinedObjectStructure) SetMatch(v map[string]interface{})`

SetMatch sets Match field to given value.

### HasMatch

`func (o *UserDefinedObjectStructure) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### GetMatchAll

`func (o *UserDefinedObjectStructure) GetMatchAll() interface{}`

GetMatchAll returns the MatchAll field if non-nil, zero value otherwise.

### GetMatchAllOk

`func (o *UserDefinedObjectStructure) GetMatchAllOk() (*interface{}, bool)`

GetMatchAllOk returns a tuple with the MatchAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchAll

`func (o *UserDefinedObjectStructure) SetMatchAll(v interface{})`

SetMatchAll sets MatchAll field to given value.

### HasMatchAll

`func (o *UserDefinedObjectStructure) HasMatchAll() bool`

HasMatchAll returns a boolean if a field has been set.

### SetMatchAllNil

`func (o *UserDefinedObjectStructure) SetMatchAllNil(b bool)`

 SetMatchAllNil sets the value for MatchAll to be an explicit nil

### UnsetMatchAll
`func (o *UserDefinedObjectStructure) UnsetMatchAll()`

UnsetMatchAll ensures that no value is present for MatchAll, not even an explicit nil
### GetMatchBoolPrefix

`func (o *UserDefinedObjectStructure) GetMatchBoolPrefix() map[string]interface{}`

GetMatchBoolPrefix returns the MatchBoolPrefix field if non-nil, zero value otherwise.

### GetMatchBoolPrefixOk

`func (o *UserDefinedObjectStructure) GetMatchBoolPrefixOk() (*map[string]interface{}, bool)`

GetMatchBoolPrefixOk returns a tuple with the MatchBoolPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchBoolPrefix

`func (o *UserDefinedObjectStructure) SetMatchBoolPrefix(v map[string]interface{})`

SetMatchBoolPrefix sets MatchBoolPrefix field to given value.

### HasMatchBoolPrefix

`func (o *UserDefinedObjectStructure) HasMatchBoolPrefix() bool`

HasMatchBoolPrefix returns a boolean if a field has been set.

### GetMatchNone

`func (o *UserDefinedObjectStructure) GetMatchNone() interface{}`

GetMatchNone returns the MatchNone field if non-nil, zero value otherwise.

### GetMatchNoneOk

`func (o *UserDefinedObjectStructure) GetMatchNoneOk() (*interface{}, bool)`

GetMatchNoneOk returns a tuple with the MatchNone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchNone

`func (o *UserDefinedObjectStructure) SetMatchNone(v interface{})`

SetMatchNone sets MatchNone field to given value.

### HasMatchNone

`func (o *UserDefinedObjectStructure) HasMatchNone() bool`

HasMatchNone returns a boolean if a field has been set.

### SetMatchNoneNil

`func (o *UserDefinedObjectStructure) SetMatchNoneNil(b bool)`

 SetMatchNoneNil sets the value for MatchNone to be an explicit nil

### UnsetMatchNone
`func (o *UserDefinedObjectStructure) UnsetMatchNone()`

UnsetMatchNone ensures that no value is present for MatchNone, not even an explicit nil
### GetMatchPhrase

`func (o *UserDefinedObjectStructure) GetMatchPhrase() map[string]interface{}`

GetMatchPhrase returns the MatchPhrase field if non-nil, zero value otherwise.

### GetMatchPhraseOk

`func (o *UserDefinedObjectStructure) GetMatchPhraseOk() (*map[string]interface{}, bool)`

GetMatchPhraseOk returns a tuple with the MatchPhrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchPhrase

`func (o *UserDefinedObjectStructure) SetMatchPhrase(v map[string]interface{})`

SetMatchPhrase sets MatchPhrase field to given value.

### HasMatchPhrase

`func (o *UserDefinedObjectStructure) HasMatchPhrase() bool`

HasMatchPhrase returns a boolean if a field has been set.

### GetMatchPhrasePrefix

`func (o *UserDefinedObjectStructure) GetMatchPhrasePrefix() map[string]interface{}`

GetMatchPhrasePrefix returns the MatchPhrasePrefix field if non-nil, zero value otherwise.

### GetMatchPhrasePrefixOk

`func (o *UserDefinedObjectStructure) GetMatchPhrasePrefixOk() (*map[string]interface{}, bool)`

GetMatchPhrasePrefixOk returns a tuple with the MatchPhrasePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchPhrasePrefix

`func (o *UserDefinedObjectStructure) SetMatchPhrasePrefix(v map[string]interface{})`

SetMatchPhrasePrefix sets MatchPhrasePrefix field to given value.

### HasMatchPhrasePrefix

`func (o *UserDefinedObjectStructure) HasMatchPhrasePrefix() bool`

HasMatchPhrasePrefix returns a boolean if a field has been set.

### GetMoreLikeThis

`func (o *UserDefinedObjectStructure) GetMoreLikeThis() interface{}`

GetMoreLikeThis returns the MoreLikeThis field if non-nil, zero value otherwise.

### GetMoreLikeThisOk

`func (o *UserDefinedObjectStructure) GetMoreLikeThisOk() (*interface{}, bool)`

GetMoreLikeThisOk returns a tuple with the MoreLikeThis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoreLikeThis

`func (o *UserDefinedObjectStructure) SetMoreLikeThis(v interface{})`

SetMoreLikeThis sets MoreLikeThis field to given value.

### HasMoreLikeThis

`func (o *UserDefinedObjectStructure) HasMoreLikeThis() bool`

HasMoreLikeThis returns a boolean if a field has been set.

### SetMoreLikeThisNil

`func (o *UserDefinedObjectStructure) SetMoreLikeThisNil(b bool)`

 SetMoreLikeThisNil sets the value for MoreLikeThis to be an explicit nil

### UnsetMoreLikeThis
`func (o *UserDefinedObjectStructure) UnsetMoreLikeThis()`

UnsetMoreLikeThis ensures that no value is present for MoreLikeThis, not even an explicit nil
### GetMultiMatch

`func (o *UserDefinedObjectStructure) GetMultiMatch() interface{}`

GetMultiMatch returns the MultiMatch field if non-nil, zero value otherwise.

### GetMultiMatchOk

`func (o *UserDefinedObjectStructure) GetMultiMatchOk() (*interface{}, bool)`

GetMultiMatchOk returns a tuple with the MultiMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiMatch

`func (o *UserDefinedObjectStructure) SetMultiMatch(v interface{})`

SetMultiMatch sets MultiMatch field to given value.

### HasMultiMatch

`func (o *UserDefinedObjectStructure) HasMultiMatch() bool`

HasMultiMatch returns a boolean if a field has been set.

### SetMultiMatchNil

`func (o *UserDefinedObjectStructure) SetMultiMatchNil(b bool)`

 SetMultiMatchNil sets the value for MultiMatch to be an explicit nil

### UnsetMultiMatch
`func (o *UserDefinedObjectStructure) UnsetMultiMatch()`

UnsetMultiMatch ensures that no value is present for MultiMatch, not even an explicit nil
### GetNested

`func (o *UserDefinedObjectStructure) GetNested() interface{}`

GetNested returns the Nested field if non-nil, zero value otherwise.

### GetNestedOk

`func (o *UserDefinedObjectStructure) GetNestedOk() (*interface{}, bool)`

GetNestedOk returns a tuple with the Nested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNested

`func (o *UserDefinedObjectStructure) SetNested(v interface{})`

SetNested sets Nested field to given value.

### HasNested

`func (o *UserDefinedObjectStructure) HasNested() bool`

HasNested returns a boolean if a field has been set.

### SetNestedNil

`func (o *UserDefinedObjectStructure) SetNestedNil(b bool)`

 SetNestedNil sets the value for Nested to be an explicit nil

### UnsetNested
`func (o *UserDefinedObjectStructure) UnsetNested()`

UnsetNested ensures that no value is present for Nested, not even an explicit nil
### GetParentId

`func (o *UserDefinedObjectStructure) GetParentId() interface{}`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *UserDefinedObjectStructure) GetParentIdOk() (*interface{}, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *UserDefinedObjectStructure) SetParentId(v interface{})`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *UserDefinedObjectStructure) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *UserDefinedObjectStructure) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *UserDefinedObjectStructure) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetPercolate

`func (o *UserDefinedObjectStructure) GetPercolate() interface{}`

GetPercolate returns the Percolate field if non-nil, zero value otherwise.

### GetPercolateOk

`func (o *UserDefinedObjectStructure) GetPercolateOk() (*interface{}, bool)`

GetPercolateOk returns a tuple with the Percolate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercolate

`func (o *UserDefinedObjectStructure) SetPercolate(v interface{})`

SetPercolate sets Percolate field to given value.

### HasPercolate

`func (o *UserDefinedObjectStructure) HasPercolate() bool`

HasPercolate returns a boolean if a field has been set.

### SetPercolateNil

`func (o *UserDefinedObjectStructure) SetPercolateNil(b bool)`

 SetPercolateNil sets the value for Percolate to be an explicit nil

### UnsetPercolate
`func (o *UserDefinedObjectStructure) UnsetPercolate()`

UnsetPercolate ensures that no value is present for Percolate, not even an explicit nil
### GetPinned

`func (o *UserDefinedObjectStructure) GetPinned() interface{}`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *UserDefinedObjectStructure) GetPinnedOk() (*interface{}, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *UserDefinedObjectStructure) SetPinned(v interface{})`

SetPinned sets Pinned field to given value.

### HasPinned

`func (o *UserDefinedObjectStructure) HasPinned() bool`

HasPinned returns a boolean if a field has been set.

### SetPinnedNil

`func (o *UserDefinedObjectStructure) SetPinnedNil(b bool)`

 SetPinnedNil sets the value for Pinned to be an explicit nil

### UnsetPinned
`func (o *UserDefinedObjectStructure) UnsetPinned()`

UnsetPinned ensures that no value is present for Pinned, not even an explicit nil
### GetPrefix

`func (o *UserDefinedObjectStructure) GetPrefix() map[string]interface{}`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *UserDefinedObjectStructure) GetPrefixOk() (*map[string]interface{}, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *UserDefinedObjectStructure) SetPrefix(v map[string]interface{})`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *UserDefinedObjectStructure) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetQueryString

`func (o *UserDefinedObjectStructure) GetQueryString() interface{}`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *UserDefinedObjectStructure) GetQueryStringOk() (*interface{}, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *UserDefinedObjectStructure) SetQueryString(v interface{})`

SetQueryString sets QueryString field to given value.

### HasQueryString

`func (o *UserDefinedObjectStructure) HasQueryString() bool`

HasQueryString returns a boolean if a field has been set.

### SetQueryStringNil

`func (o *UserDefinedObjectStructure) SetQueryStringNil(b bool)`

 SetQueryStringNil sets the value for QueryString to be an explicit nil

### UnsetQueryString
`func (o *UserDefinedObjectStructure) UnsetQueryString()`

UnsetQueryString ensures that no value is present for QueryString, not even an explicit nil
### GetRange

`func (o *UserDefinedObjectStructure) GetRange() map[string]interface{}`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *UserDefinedObjectStructure) GetRangeOk() (*map[string]interface{}, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *UserDefinedObjectStructure) SetRange(v map[string]interface{})`

SetRange sets Range field to given value.

### HasRange

`func (o *UserDefinedObjectStructure) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetRankFeature

`func (o *UserDefinedObjectStructure) GetRankFeature() interface{}`

GetRankFeature returns the RankFeature field if non-nil, zero value otherwise.

### GetRankFeatureOk

`func (o *UserDefinedObjectStructure) GetRankFeatureOk() (*interface{}, bool)`

GetRankFeatureOk returns a tuple with the RankFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRankFeature

`func (o *UserDefinedObjectStructure) SetRankFeature(v interface{})`

SetRankFeature sets RankFeature field to given value.

### HasRankFeature

`func (o *UserDefinedObjectStructure) HasRankFeature() bool`

HasRankFeature returns a boolean if a field has been set.

### SetRankFeatureNil

`func (o *UserDefinedObjectStructure) SetRankFeatureNil(b bool)`

 SetRankFeatureNil sets the value for RankFeature to be an explicit nil

### UnsetRankFeature
`func (o *UserDefinedObjectStructure) UnsetRankFeature()`

UnsetRankFeature ensures that no value is present for RankFeature, not even an explicit nil
### GetRegexp

`func (o *UserDefinedObjectStructure) GetRegexp() map[string]interface{}`

GetRegexp returns the Regexp field if non-nil, zero value otherwise.

### GetRegexpOk

`func (o *UserDefinedObjectStructure) GetRegexpOk() (*map[string]interface{}, bool)`

GetRegexpOk returns a tuple with the Regexp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexp

`func (o *UserDefinedObjectStructure) SetRegexp(v map[string]interface{})`

SetRegexp sets Regexp field to given value.

### HasRegexp

`func (o *UserDefinedObjectStructure) HasRegexp() bool`

HasRegexp returns a boolean if a field has been set.

### GetScript

`func (o *UserDefinedObjectStructure) GetScript() interface{}`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *UserDefinedObjectStructure) GetScriptOk() (*interface{}, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *UserDefinedObjectStructure) SetScript(v interface{})`

SetScript sets Script field to given value.

### HasScript

`func (o *UserDefinedObjectStructure) HasScript() bool`

HasScript returns a boolean if a field has been set.

### SetScriptNil

`func (o *UserDefinedObjectStructure) SetScriptNil(b bool)`

 SetScriptNil sets the value for Script to be an explicit nil

### UnsetScript
`func (o *UserDefinedObjectStructure) UnsetScript()`

UnsetScript ensures that no value is present for Script, not even an explicit nil
### GetScriptScore

`func (o *UserDefinedObjectStructure) GetScriptScore() interface{}`

GetScriptScore returns the ScriptScore field if non-nil, zero value otherwise.

### GetScriptScoreOk

`func (o *UserDefinedObjectStructure) GetScriptScoreOk() (*interface{}, bool)`

GetScriptScoreOk returns a tuple with the ScriptScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptScore

`func (o *UserDefinedObjectStructure) SetScriptScore(v interface{})`

SetScriptScore sets ScriptScore field to given value.

### HasScriptScore

`func (o *UserDefinedObjectStructure) HasScriptScore() bool`

HasScriptScore returns a boolean if a field has been set.

### SetScriptScoreNil

`func (o *UserDefinedObjectStructure) SetScriptScoreNil(b bool)`

 SetScriptScoreNil sets the value for ScriptScore to be an explicit nil

### UnsetScriptScore
`func (o *UserDefinedObjectStructure) UnsetScriptScore()`

UnsetScriptScore ensures that no value is present for ScriptScore, not even an explicit nil
### GetShape

`func (o *UserDefinedObjectStructure) GetShape() interface{}`

GetShape returns the Shape field if non-nil, zero value otherwise.

### GetShapeOk

`func (o *UserDefinedObjectStructure) GetShapeOk() (*interface{}, bool)`

GetShapeOk returns a tuple with the Shape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShape

`func (o *UserDefinedObjectStructure) SetShape(v interface{})`

SetShape sets Shape field to given value.

### HasShape

`func (o *UserDefinedObjectStructure) HasShape() bool`

HasShape returns a boolean if a field has been set.

### SetShapeNil

`func (o *UserDefinedObjectStructure) SetShapeNil(b bool)`

 SetShapeNil sets the value for Shape to be an explicit nil

### UnsetShape
`func (o *UserDefinedObjectStructure) UnsetShape()`

UnsetShape ensures that no value is present for Shape, not even an explicit nil
### GetSimpleQueryString

`func (o *UserDefinedObjectStructure) GetSimpleQueryString() interface{}`

GetSimpleQueryString returns the SimpleQueryString field if non-nil, zero value otherwise.

### GetSimpleQueryStringOk

`func (o *UserDefinedObjectStructure) GetSimpleQueryStringOk() (*interface{}, bool)`

GetSimpleQueryStringOk returns a tuple with the SimpleQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimpleQueryString

`func (o *UserDefinedObjectStructure) SetSimpleQueryString(v interface{})`

SetSimpleQueryString sets SimpleQueryString field to given value.

### HasSimpleQueryString

`func (o *UserDefinedObjectStructure) HasSimpleQueryString() bool`

HasSimpleQueryString returns a boolean if a field has been set.

### SetSimpleQueryStringNil

`func (o *UserDefinedObjectStructure) SetSimpleQueryStringNil(b bool)`

 SetSimpleQueryStringNil sets the value for SimpleQueryString to be an explicit nil

### UnsetSimpleQueryString
`func (o *UserDefinedObjectStructure) UnsetSimpleQueryString()`

UnsetSimpleQueryString ensures that no value is present for SimpleQueryString, not even an explicit nil
### GetSpanContaining

`func (o *UserDefinedObjectStructure) GetSpanContaining() interface{}`

GetSpanContaining returns the SpanContaining field if non-nil, zero value otherwise.

### GetSpanContainingOk

`func (o *UserDefinedObjectStructure) GetSpanContainingOk() (*interface{}, bool)`

GetSpanContainingOk returns a tuple with the SpanContaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanContaining

`func (o *UserDefinedObjectStructure) SetSpanContaining(v interface{})`

SetSpanContaining sets SpanContaining field to given value.

### HasSpanContaining

`func (o *UserDefinedObjectStructure) HasSpanContaining() bool`

HasSpanContaining returns a boolean if a field has been set.

### SetSpanContainingNil

`func (o *UserDefinedObjectStructure) SetSpanContainingNil(b bool)`

 SetSpanContainingNil sets the value for SpanContaining to be an explicit nil

### UnsetSpanContaining
`func (o *UserDefinedObjectStructure) UnsetSpanContaining()`

UnsetSpanContaining ensures that no value is present for SpanContaining, not even an explicit nil
### GetFieldMaskingSpan

`func (o *UserDefinedObjectStructure) GetFieldMaskingSpan() interface{}`

GetFieldMaskingSpan returns the FieldMaskingSpan field if non-nil, zero value otherwise.

### GetFieldMaskingSpanOk

`func (o *UserDefinedObjectStructure) GetFieldMaskingSpanOk() (*interface{}, bool)`

GetFieldMaskingSpanOk returns a tuple with the FieldMaskingSpan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldMaskingSpan

`func (o *UserDefinedObjectStructure) SetFieldMaskingSpan(v interface{})`

SetFieldMaskingSpan sets FieldMaskingSpan field to given value.

### HasFieldMaskingSpan

`func (o *UserDefinedObjectStructure) HasFieldMaskingSpan() bool`

HasFieldMaskingSpan returns a boolean if a field has been set.

### SetFieldMaskingSpanNil

`func (o *UserDefinedObjectStructure) SetFieldMaskingSpanNil(b bool)`

 SetFieldMaskingSpanNil sets the value for FieldMaskingSpan to be an explicit nil

### UnsetFieldMaskingSpan
`func (o *UserDefinedObjectStructure) UnsetFieldMaskingSpan()`

UnsetFieldMaskingSpan ensures that no value is present for FieldMaskingSpan, not even an explicit nil
### GetSpanFirst

`func (o *UserDefinedObjectStructure) GetSpanFirst() interface{}`

GetSpanFirst returns the SpanFirst field if non-nil, zero value otherwise.

### GetSpanFirstOk

`func (o *UserDefinedObjectStructure) GetSpanFirstOk() (*interface{}, bool)`

GetSpanFirstOk returns a tuple with the SpanFirst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanFirst

`func (o *UserDefinedObjectStructure) SetSpanFirst(v interface{})`

SetSpanFirst sets SpanFirst field to given value.

### HasSpanFirst

`func (o *UserDefinedObjectStructure) HasSpanFirst() bool`

HasSpanFirst returns a boolean if a field has been set.

### SetSpanFirstNil

`func (o *UserDefinedObjectStructure) SetSpanFirstNil(b bool)`

 SetSpanFirstNil sets the value for SpanFirst to be an explicit nil

### UnsetSpanFirst
`func (o *UserDefinedObjectStructure) UnsetSpanFirst()`

UnsetSpanFirst ensures that no value is present for SpanFirst, not even an explicit nil
### GetSpanMulti

`func (o *UserDefinedObjectStructure) GetSpanMulti() interface{}`

GetSpanMulti returns the SpanMulti field if non-nil, zero value otherwise.

### GetSpanMultiOk

`func (o *UserDefinedObjectStructure) GetSpanMultiOk() (*interface{}, bool)`

GetSpanMultiOk returns a tuple with the SpanMulti field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanMulti

`func (o *UserDefinedObjectStructure) SetSpanMulti(v interface{})`

SetSpanMulti sets SpanMulti field to given value.

### HasSpanMulti

`func (o *UserDefinedObjectStructure) HasSpanMulti() bool`

HasSpanMulti returns a boolean if a field has been set.

### SetSpanMultiNil

`func (o *UserDefinedObjectStructure) SetSpanMultiNil(b bool)`

 SetSpanMultiNil sets the value for SpanMulti to be an explicit nil

### UnsetSpanMulti
`func (o *UserDefinedObjectStructure) UnsetSpanMulti()`

UnsetSpanMulti ensures that no value is present for SpanMulti, not even an explicit nil
### GetSpanNear

`func (o *UserDefinedObjectStructure) GetSpanNear() interface{}`

GetSpanNear returns the SpanNear field if non-nil, zero value otherwise.

### GetSpanNearOk

`func (o *UserDefinedObjectStructure) GetSpanNearOk() (*interface{}, bool)`

GetSpanNearOk returns a tuple with the SpanNear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanNear

`func (o *UserDefinedObjectStructure) SetSpanNear(v interface{})`

SetSpanNear sets SpanNear field to given value.

### HasSpanNear

`func (o *UserDefinedObjectStructure) HasSpanNear() bool`

HasSpanNear returns a boolean if a field has been set.

### SetSpanNearNil

`func (o *UserDefinedObjectStructure) SetSpanNearNil(b bool)`

 SetSpanNearNil sets the value for SpanNear to be an explicit nil

### UnsetSpanNear
`func (o *UserDefinedObjectStructure) UnsetSpanNear()`

UnsetSpanNear ensures that no value is present for SpanNear, not even an explicit nil
### GetSpanNot

`func (o *UserDefinedObjectStructure) GetSpanNot() interface{}`

GetSpanNot returns the SpanNot field if non-nil, zero value otherwise.

### GetSpanNotOk

`func (o *UserDefinedObjectStructure) GetSpanNotOk() (*interface{}, bool)`

GetSpanNotOk returns a tuple with the SpanNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanNot

`func (o *UserDefinedObjectStructure) SetSpanNot(v interface{})`

SetSpanNot sets SpanNot field to given value.

### HasSpanNot

`func (o *UserDefinedObjectStructure) HasSpanNot() bool`

HasSpanNot returns a boolean if a field has been set.

### SetSpanNotNil

`func (o *UserDefinedObjectStructure) SetSpanNotNil(b bool)`

 SetSpanNotNil sets the value for SpanNot to be an explicit nil

### UnsetSpanNot
`func (o *UserDefinedObjectStructure) UnsetSpanNot()`

UnsetSpanNot ensures that no value is present for SpanNot, not even an explicit nil
### GetSpanOr

`func (o *UserDefinedObjectStructure) GetSpanOr() interface{}`

GetSpanOr returns the SpanOr field if non-nil, zero value otherwise.

### GetSpanOrOk

`func (o *UserDefinedObjectStructure) GetSpanOrOk() (*interface{}, bool)`

GetSpanOrOk returns a tuple with the SpanOr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanOr

`func (o *UserDefinedObjectStructure) SetSpanOr(v interface{})`

SetSpanOr sets SpanOr field to given value.

### HasSpanOr

`func (o *UserDefinedObjectStructure) HasSpanOr() bool`

HasSpanOr returns a boolean if a field has been set.

### SetSpanOrNil

`func (o *UserDefinedObjectStructure) SetSpanOrNil(b bool)`

 SetSpanOrNil sets the value for SpanOr to be an explicit nil

### UnsetSpanOr
`func (o *UserDefinedObjectStructure) UnsetSpanOr()`

UnsetSpanOr ensures that no value is present for SpanOr, not even an explicit nil
### GetSpanTerm

`func (o *UserDefinedObjectStructure) GetSpanTerm() map[string]interface{}`

GetSpanTerm returns the SpanTerm field if non-nil, zero value otherwise.

### GetSpanTermOk

`func (o *UserDefinedObjectStructure) GetSpanTermOk() (*map[string]interface{}, bool)`

GetSpanTermOk returns a tuple with the SpanTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanTerm

`func (o *UserDefinedObjectStructure) SetSpanTerm(v map[string]interface{})`

SetSpanTerm sets SpanTerm field to given value.

### HasSpanTerm

`func (o *UserDefinedObjectStructure) HasSpanTerm() bool`

HasSpanTerm returns a boolean if a field has been set.

### GetSpanWithin

`func (o *UserDefinedObjectStructure) GetSpanWithin() interface{}`

GetSpanWithin returns the SpanWithin field if non-nil, zero value otherwise.

### GetSpanWithinOk

`func (o *UserDefinedObjectStructure) GetSpanWithinOk() (*interface{}, bool)`

GetSpanWithinOk returns a tuple with the SpanWithin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanWithin

`func (o *UserDefinedObjectStructure) SetSpanWithin(v interface{})`

SetSpanWithin sets SpanWithin field to given value.

### HasSpanWithin

`func (o *UserDefinedObjectStructure) HasSpanWithin() bool`

HasSpanWithin returns a boolean if a field has been set.

### SetSpanWithinNil

`func (o *UserDefinedObjectStructure) SetSpanWithinNil(b bool)`

 SetSpanWithinNil sets the value for SpanWithin to be an explicit nil

### UnsetSpanWithin
`func (o *UserDefinedObjectStructure) UnsetSpanWithin()`

UnsetSpanWithin ensures that no value is present for SpanWithin, not even an explicit nil
### GetTerm

`func (o *UserDefinedObjectStructure) GetTerm() map[string]interface{}`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *UserDefinedObjectStructure) GetTermOk() (*map[string]interface{}, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *UserDefinedObjectStructure) SetTerm(v map[string]interface{})`

SetTerm sets Term field to given value.

### HasTerm

`func (o *UserDefinedObjectStructure) HasTerm() bool`

HasTerm returns a boolean if a field has been set.

### GetTerms

`func (o *UserDefinedObjectStructure) GetTerms() interface{}`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *UserDefinedObjectStructure) GetTermsOk() (*interface{}, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *UserDefinedObjectStructure) SetTerms(v interface{})`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *UserDefinedObjectStructure) HasTerms() bool`

HasTerms returns a boolean if a field has been set.

### SetTermsNil

`func (o *UserDefinedObjectStructure) SetTermsNil(b bool)`

 SetTermsNil sets the value for Terms to be an explicit nil

### UnsetTerms
`func (o *UserDefinedObjectStructure) UnsetTerms()`

UnsetTerms ensures that no value is present for Terms, not even an explicit nil
### GetTermsSet

`func (o *UserDefinedObjectStructure) GetTermsSet() map[string]interface{}`

GetTermsSet returns the TermsSet field if non-nil, zero value otherwise.

### GetTermsSetOk

`func (o *UserDefinedObjectStructure) GetTermsSetOk() (*map[string]interface{}, bool)`

GetTermsSetOk returns a tuple with the TermsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsSet

`func (o *UserDefinedObjectStructure) SetTermsSet(v map[string]interface{})`

SetTermsSet sets TermsSet field to given value.

### HasTermsSet

`func (o *UserDefinedObjectStructure) HasTermsSet() bool`

HasTermsSet returns a boolean if a field has been set.

### GetWildcard

`func (o *UserDefinedObjectStructure) GetWildcard() map[string]interface{}`

GetWildcard returns the Wildcard field if non-nil, zero value otherwise.

### GetWildcardOk

`func (o *UserDefinedObjectStructure) GetWildcardOk() (*map[string]interface{}, bool)`

GetWildcardOk returns a tuple with the Wildcard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWildcard

`func (o *UserDefinedObjectStructure) SetWildcard(v map[string]interface{})`

SetWildcard sets Wildcard field to given value.

### HasWildcard

`func (o *UserDefinedObjectStructure) HasWildcard() bool`

HasWildcard returns a boolean if a field has been set.

### GetWrapper

`func (o *UserDefinedObjectStructure) GetWrapper() interface{}`

GetWrapper returns the Wrapper field if non-nil, zero value otherwise.

### GetWrapperOk

`func (o *UserDefinedObjectStructure) GetWrapperOk() (*interface{}, bool)`

GetWrapperOk returns a tuple with the Wrapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrapper

`func (o *UserDefinedObjectStructure) SetWrapper(v interface{})`

SetWrapper sets Wrapper field to given value.

### HasWrapper

`func (o *UserDefinedObjectStructure) HasWrapper() bool`

HasWrapper returns a boolean if a field has been set.

### SetWrapperNil

`func (o *UserDefinedObjectStructure) SetWrapperNil(b bool)`

 SetWrapperNil sets the value for Wrapper to be an explicit nil

### UnsetWrapper
`func (o *UserDefinedObjectStructure) UnsetWrapper()`

UnsetWrapper ensures that no value is present for Wrapper, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


