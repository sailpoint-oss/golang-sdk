# Search1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Indices** | Pointer to [**[]Index**](Index.md) | The names of the Elasticsearch indices in which to search. If none are provided, then all indices will be searched. | [optional] 
**QueryType** | Pointer to [**QueryType**](QueryType.md) |  | [optional] [default to QUERYTYPE_SAILPOINT]
**QueryVersion** | Pointer to **string** |  | [optional] 
**Query** | Pointer to [**Query**](Query.md) |  | [optional] 
**QueryDsl** | Pointer to **map[string]interface{}** | The search query using the Elasticsearch [Query DSL](https://www.elastic.co/guide/en/elasticsearch/reference/7.10/query-dsl.html) syntax. | [optional] 
**TypeAheadQuery** | Pointer to [**TypeAheadQuery**](TypeAheadQuery.md) |  | [optional] 
**IncludeNested** | Pointer to **bool** | Indicates whether nested objects from returned search results should be included. | [optional] [default to true]
**QueryResultFilter** | Pointer to [**QueryResultFilter**](QueryResultFilter.md) |  | [optional] 
**AggregationType** | Pointer to [**AggregationType**](AggregationType.md) |  | [optional] [default to AGGREGATIONTYPE_DSL]
**AggregationsVersion** | Pointer to **string** |  | [optional] 
**AggregationsDsl** | Pointer to **map[string]interface{}** | The aggregation search query using Elasticsearch [Aggregations](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/search-aggregations.html) syntax. | [optional] 
**Aggregations** | Pointer to [**Aggregation1**](Aggregation1.md) |  | [optional] 
**Sort** | Pointer to **[]string** | The fields to be used to sort the search results. Use + or - to specify the sort direction. | [optional] 
**SearchAfter** | Pointer to **[]string** | Used to begin the search window at the values specified. This parameter consists of the last values of the sorted fields in the current record set. This is used to expand the Elasticsearch limit of 10K records by shifting the 10K window to begin at this value. For example, when searching for identities, if the last idenitity ID in the search result is 2c91808375d8e80a0175e1f88a575221, then using that ID in this property will start a new search after this identity. | [optional] 
**Filters** | Pointer to [**map[string]Filter**](Filter.md) | The filters to be applied for each filtered field name. | [optional] 

## Methods

### NewSearch1

`func NewSearch1() *Search1`

NewSearch1 instantiates a new Search1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearch1WithDefaults

`func NewSearch1WithDefaults() *Search1`

NewSearch1WithDefaults instantiates a new Search1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndices

`func (o *Search1) GetIndices() []Index`

GetIndices returns the Indices field if non-nil, zero value otherwise.

### GetIndicesOk

`func (o *Search1) GetIndicesOk() (*[]Index, bool)`

GetIndicesOk returns a tuple with the Indices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndices

`func (o *Search1) SetIndices(v []Index)`

SetIndices sets Indices field to given value.

### HasIndices

`func (o *Search1) HasIndices() bool`

HasIndices returns a boolean if a field has been set.

### GetQueryType

`func (o *Search1) GetQueryType() QueryType`

GetQueryType returns the QueryType field if non-nil, zero value otherwise.

### GetQueryTypeOk

`func (o *Search1) GetQueryTypeOk() (*QueryType, bool)`

GetQueryTypeOk returns a tuple with the QueryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryType

`func (o *Search1) SetQueryType(v QueryType)`

SetQueryType sets QueryType field to given value.

### HasQueryType

`func (o *Search1) HasQueryType() bool`

HasQueryType returns a boolean if a field has been set.

### GetQueryVersion

`func (o *Search1) GetQueryVersion() string`

GetQueryVersion returns the QueryVersion field if non-nil, zero value otherwise.

### GetQueryVersionOk

`func (o *Search1) GetQueryVersionOk() (*string, bool)`

GetQueryVersionOk returns a tuple with the QueryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryVersion

`func (o *Search1) SetQueryVersion(v string)`

SetQueryVersion sets QueryVersion field to given value.

### HasQueryVersion

`func (o *Search1) HasQueryVersion() bool`

HasQueryVersion returns a boolean if a field has been set.

### GetQuery

`func (o *Search1) GetQuery() Query`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Search1) GetQueryOk() (*Query, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Search1) SetQuery(v Query)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *Search1) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetQueryDsl

`func (o *Search1) GetQueryDsl() map[string]interface{}`

GetQueryDsl returns the QueryDsl field if non-nil, zero value otherwise.

### GetQueryDslOk

`func (o *Search1) GetQueryDslOk() (*map[string]interface{}, bool)`

GetQueryDslOk returns a tuple with the QueryDsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryDsl

`func (o *Search1) SetQueryDsl(v map[string]interface{})`

SetQueryDsl sets QueryDsl field to given value.

### HasQueryDsl

`func (o *Search1) HasQueryDsl() bool`

HasQueryDsl returns a boolean if a field has been set.

### GetTypeAheadQuery

`func (o *Search1) GetTypeAheadQuery() TypeAheadQuery`

GetTypeAheadQuery returns the TypeAheadQuery field if non-nil, zero value otherwise.

### GetTypeAheadQueryOk

`func (o *Search1) GetTypeAheadQueryOk() (*TypeAheadQuery, bool)`

GetTypeAheadQueryOk returns a tuple with the TypeAheadQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeAheadQuery

`func (o *Search1) SetTypeAheadQuery(v TypeAheadQuery)`

SetTypeAheadQuery sets TypeAheadQuery field to given value.

### HasTypeAheadQuery

`func (o *Search1) HasTypeAheadQuery() bool`

HasTypeAheadQuery returns a boolean if a field has been set.

### GetIncludeNested

`func (o *Search1) GetIncludeNested() bool`

GetIncludeNested returns the IncludeNested field if non-nil, zero value otherwise.

### GetIncludeNestedOk

`func (o *Search1) GetIncludeNestedOk() (*bool, bool)`

GetIncludeNestedOk returns a tuple with the IncludeNested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeNested

`func (o *Search1) SetIncludeNested(v bool)`

SetIncludeNested sets IncludeNested field to given value.

### HasIncludeNested

`func (o *Search1) HasIncludeNested() bool`

HasIncludeNested returns a boolean if a field has been set.

### GetQueryResultFilter

`func (o *Search1) GetQueryResultFilter() QueryResultFilter`

GetQueryResultFilter returns the QueryResultFilter field if non-nil, zero value otherwise.

### GetQueryResultFilterOk

`func (o *Search1) GetQueryResultFilterOk() (*QueryResultFilter, bool)`

GetQueryResultFilterOk returns a tuple with the QueryResultFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryResultFilter

`func (o *Search1) SetQueryResultFilter(v QueryResultFilter)`

SetQueryResultFilter sets QueryResultFilter field to given value.

### HasQueryResultFilter

`func (o *Search1) HasQueryResultFilter() bool`

HasQueryResultFilter returns a boolean if a field has been set.

### GetAggregationType

`func (o *Search1) GetAggregationType() AggregationType`

GetAggregationType returns the AggregationType field if non-nil, zero value otherwise.

### GetAggregationTypeOk

`func (o *Search1) GetAggregationTypeOk() (*AggregationType, bool)`

GetAggregationTypeOk returns a tuple with the AggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationType

`func (o *Search1) SetAggregationType(v AggregationType)`

SetAggregationType sets AggregationType field to given value.

### HasAggregationType

`func (o *Search1) HasAggregationType() bool`

HasAggregationType returns a boolean if a field has been set.

### GetAggregationsVersion

`func (o *Search1) GetAggregationsVersion() string`

GetAggregationsVersion returns the AggregationsVersion field if non-nil, zero value otherwise.

### GetAggregationsVersionOk

`func (o *Search1) GetAggregationsVersionOk() (*string, bool)`

GetAggregationsVersionOk returns a tuple with the AggregationsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationsVersion

`func (o *Search1) SetAggregationsVersion(v string)`

SetAggregationsVersion sets AggregationsVersion field to given value.

### HasAggregationsVersion

`func (o *Search1) HasAggregationsVersion() bool`

HasAggregationsVersion returns a boolean if a field has been set.

### GetAggregationsDsl

`func (o *Search1) GetAggregationsDsl() map[string]interface{}`

GetAggregationsDsl returns the AggregationsDsl field if non-nil, zero value otherwise.

### GetAggregationsDslOk

`func (o *Search1) GetAggregationsDslOk() (*map[string]interface{}, bool)`

GetAggregationsDslOk returns a tuple with the AggregationsDsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationsDsl

`func (o *Search1) SetAggregationsDsl(v map[string]interface{})`

SetAggregationsDsl sets AggregationsDsl field to given value.

### HasAggregationsDsl

`func (o *Search1) HasAggregationsDsl() bool`

HasAggregationsDsl returns a boolean if a field has been set.

### GetAggregations

`func (o *Search1) GetAggregations() Aggregation1`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *Search1) GetAggregationsOk() (*Aggregation1, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *Search1) SetAggregations(v Aggregation1)`

SetAggregations sets Aggregations field to given value.

### HasAggregations

`func (o *Search1) HasAggregations() bool`

HasAggregations returns a boolean if a field has been set.

### GetSort

`func (o *Search1) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *Search1) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *Search1) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *Search1) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetSearchAfter

`func (o *Search1) GetSearchAfter() []string`

GetSearchAfter returns the SearchAfter field if non-nil, zero value otherwise.

### GetSearchAfterOk

`func (o *Search1) GetSearchAfterOk() (*[]string, bool)`

GetSearchAfterOk returns a tuple with the SearchAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchAfter

`func (o *Search1) SetSearchAfter(v []string)`

SetSearchAfter sets SearchAfter field to given value.

### HasSearchAfter

`func (o *Search1) HasSearchAfter() bool`

HasSearchAfter returns a boolean if a field has been set.

### GetFilters

`func (o *Search1) GetFilters() map[string]Filter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *Search1) GetFiltersOk() (*map[string]Filter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *Search1) SetFilters(v map[string]Filter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *Search1) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


