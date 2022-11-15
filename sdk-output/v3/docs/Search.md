# Search

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Public** | Pointer to **bool** | Indicates if the saved search is public.  | [optional] [default to false]
**Created** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**Modified** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**Indices** | [**[]Index**](Index.md) | The names of the Elasticsearch indices in which to search.  | 
**Columns** | Pointer to [**map[string][]Column**](array.md) | The columns to be returned (specifies the order in which they will be presented) for each document type.  The currently supported document types are: _accessprofile_, _accountactivity_, _account_, _aggregation_, _entitlement_, _event_, _identity_, and _role_.  | [optional] 
**Query** | **string** | The search query using Elasticsearch [Query String Query](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/query-dsl-query-string-query.html#query-string) syntax from the Query DSL.  | 
**Fields** | Pointer to **[]string** | The fields to be searched against in a multi-field query.  | [optional] 
**Sort** | Pointer to **[]string** | The fields to be used to sort the search results.  | [optional] 
**Filters** | Pointer to [**NullableSearchFilters**](SearchFilters.md) |  | [optional] 

## Methods

### NewSearch

`func NewSearch(indices []Index, query string, ) *Search`

NewSearch instantiates a new Search object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchWithDefaults

`func NewSearchWithDefaults() *Search`

NewSearchWithDefaults instantiates a new Search object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublic

`func (o *Search) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *Search) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *Search) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *Search) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetCreated

`func (o *Search) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Search) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Search) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Search) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *Search) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *Search) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetModified

`func (o *Search) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Search) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Search) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *Search) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModifiedNil

`func (o *Search) SetModifiedNil(b bool)`

 SetModifiedNil sets the value for Modified to be an explicit nil

### UnsetModified
`func (o *Search) UnsetModified()`

UnsetModified ensures that no value is present for Modified, not even an explicit nil
### GetIndices

`func (o *Search) GetIndices() []Index`

GetIndices returns the Indices field if non-nil, zero value otherwise.

### GetIndicesOk

`func (o *Search) GetIndicesOk() (*[]Index, bool)`

GetIndicesOk returns a tuple with the Indices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndices

`func (o *Search) SetIndices(v []Index)`

SetIndices sets Indices field to given value.


### GetColumns

`func (o *Search) GetColumns() map[string][]Column`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *Search) GetColumnsOk() (*map[string][]Column, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *Search) SetColumns(v map[string][]Column)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *Search) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetQuery

`func (o *Search) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Search) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Search) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetFields

`func (o *Search) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *Search) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *Search) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *Search) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetSort

`func (o *Search) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *Search) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *Search) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *Search) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetFilters

`func (o *Search) GetFilters() SearchFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *Search) GetFiltersOk() (*SearchFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *Search) SetFilters(v SearchFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *Search) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *Search) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *Search) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


