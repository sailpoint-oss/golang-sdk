# SearchExportReportArguments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Indices** | Pointer to [**[]Index**](Index.md) | The names of the Elasticsearch indices in which to search. If none are provided, then all indices will be searched. | [optional] 
**Filters** | Pointer to [**map[string]Filter**](Filter.md) | The filters to be applied for each filtered field name. | [optional] 
**Query** | [**Query**](Query.md) |  | 
**IncludeNested** | Pointer to **bool** | Indicates whether nested objects from returned search results should be included. | [optional] [default to true]
**Sort** | Pointer to **[]string** | The fields to be used to sort the search results. Use + or - to specify the sort direction. | [optional] 

## Methods

### NewSearchExportReportArguments

`func NewSearchExportReportArguments(query Query, ) *SearchExportReportArguments`

NewSearchExportReportArguments instantiates a new SearchExportReportArguments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchExportReportArgumentsWithDefaults

`func NewSearchExportReportArgumentsWithDefaults() *SearchExportReportArguments`

NewSearchExportReportArgumentsWithDefaults instantiates a new SearchExportReportArguments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndices

`func (o *SearchExportReportArguments) GetIndices() []Index`

GetIndices returns the Indices field if non-nil, zero value otherwise.

### GetIndicesOk

`func (o *SearchExportReportArguments) GetIndicesOk() (*[]Index, bool)`

GetIndicesOk returns a tuple with the Indices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndices

`func (o *SearchExportReportArguments) SetIndices(v []Index)`

SetIndices sets Indices field to given value.

### HasIndices

`func (o *SearchExportReportArguments) HasIndices() bool`

HasIndices returns a boolean if a field has been set.

### GetFilters

`func (o *SearchExportReportArguments) GetFilters() map[string]Filter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *SearchExportReportArguments) GetFiltersOk() (*map[string]Filter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *SearchExportReportArguments) SetFilters(v map[string]Filter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *SearchExportReportArguments) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetQuery

`func (o *SearchExportReportArguments) GetQuery() Query`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SearchExportReportArguments) GetQueryOk() (*Query, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SearchExportReportArguments) SetQuery(v Query)`

SetQuery sets Query field to given value.


### GetIncludeNested

`func (o *SearchExportReportArguments) GetIncludeNested() bool`

GetIncludeNested returns the IncludeNested field if non-nil, zero value otherwise.

### GetIncludeNestedOk

`func (o *SearchExportReportArguments) GetIncludeNestedOk() (*bool, bool)`

GetIncludeNestedOk returns a tuple with the IncludeNested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeNested

`func (o *SearchExportReportArguments) SetIncludeNested(v bool)`

SetIncludeNested sets IncludeNested field to given value.

### HasIncludeNested

`func (o *SearchExportReportArguments) HasIncludeNested() bool`

HasIncludeNested returns a boolean if a field has been set.

### GetSort

`func (o *SearchExportReportArguments) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *SearchExportReportArguments) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *SearchExportReportArguments) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *SearchExportReportArguments) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


