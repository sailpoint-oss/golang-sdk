# SavedSearchCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the saved search.  | [optional] 
**Description** | Pointer to **string** | The description of the saved search.  | [optional] 
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

### NewSavedSearchCreateRequest

`func NewSavedSearchCreateRequest(indices []Index, query string, ) *SavedSearchCreateRequest`

NewSavedSearchCreateRequest instantiates a new SavedSearchCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavedSearchCreateRequestWithDefaults

`func NewSavedSearchCreateRequestWithDefaults() *SavedSearchCreateRequest`

NewSavedSearchCreateRequestWithDefaults instantiates a new SavedSearchCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SavedSearchCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SavedSearchCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SavedSearchCreateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SavedSearchCreateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *SavedSearchCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SavedSearchCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SavedSearchCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SavedSearchCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPublic

`func (o *SavedSearchCreateRequest) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *SavedSearchCreateRequest) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *SavedSearchCreateRequest) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *SavedSearchCreateRequest) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetCreated

`func (o *SavedSearchCreateRequest) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SavedSearchCreateRequest) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SavedSearchCreateRequest) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SavedSearchCreateRequest) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *SavedSearchCreateRequest) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *SavedSearchCreateRequest) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetModified

`func (o *SavedSearchCreateRequest) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *SavedSearchCreateRequest) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *SavedSearchCreateRequest) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *SavedSearchCreateRequest) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModifiedNil

`func (o *SavedSearchCreateRequest) SetModifiedNil(b bool)`

 SetModifiedNil sets the value for Modified to be an explicit nil

### UnsetModified
`func (o *SavedSearchCreateRequest) UnsetModified()`

UnsetModified ensures that no value is present for Modified, not even an explicit nil
### GetIndices

`func (o *SavedSearchCreateRequest) GetIndices() []Index`

GetIndices returns the Indices field if non-nil, zero value otherwise.

### GetIndicesOk

`func (o *SavedSearchCreateRequest) GetIndicesOk() (*[]Index, bool)`

GetIndicesOk returns a tuple with the Indices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndices

`func (o *SavedSearchCreateRequest) SetIndices(v []Index)`

SetIndices sets Indices field to given value.


### GetColumns

`func (o *SavedSearchCreateRequest) GetColumns() map[string][]Column`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *SavedSearchCreateRequest) GetColumnsOk() (*map[string][]Column, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *SavedSearchCreateRequest) SetColumns(v map[string][]Column)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *SavedSearchCreateRequest) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetQuery

`func (o *SavedSearchCreateRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SavedSearchCreateRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SavedSearchCreateRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetFields

`func (o *SavedSearchCreateRequest) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *SavedSearchCreateRequest) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *SavedSearchCreateRequest) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *SavedSearchCreateRequest) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetSort

`func (o *SavedSearchCreateRequest) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *SavedSearchCreateRequest) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *SavedSearchCreateRequest) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *SavedSearchCreateRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetFilters

`func (o *SavedSearchCreateRequest) GetFilters() SearchFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *SavedSearchCreateRequest) GetFiltersOk() (*SearchFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *SavedSearchCreateRequest) SetFilters(v SearchFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *SavedSearchCreateRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *SavedSearchCreateRequest) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *SavedSearchCreateRequest) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


