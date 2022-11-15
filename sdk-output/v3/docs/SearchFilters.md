# SearchFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**FilterType**](FilterType.md) |  | [optional] 
**Range** | Pointer to [**Range**](Range.md) |  | [optional] 
**Terms** | Pointer to **[]string** | The terms to be filtered. | [optional] 
**Exclude** | Pointer to **bool** | Indicates if the filter excludes results. | [optional] [default to false]

## Methods

### NewSearchFilters

`func NewSearchFilters() *SearchFilters`

NewSearchFilters instantiates a new SearchFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchFiltersWithDefaults

`func NewSearchFiltersWithDefaults() *SearchFilters`

NewSearchFiltersWithDefaults instantiates a new SearchFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SearchFilters) GetType() FilterType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SearchFilters) GetTypeOk() (*FilterType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SearchFilters) SetType(v FilterType)`

SetType sets Type field to given value.

### HasType

`func (o *SearchFilters) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRange

`func (o *SearchFilters) GetRange() Range`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *SearchFilters) GetRangeOk() (*Range, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *SearchFilters) SetRange(v Range)`

SetRange sets Range field to given value.

### HasRange

`func (o *SearchFilters) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetTerms

`func (o *SearchFilters) GetTerms() []string`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *SearchFilters) GetTermsOk() (*[]string, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *SearchFilters) SetTerms(v []string)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *SearchFilters) HasTerms() bool`

HasTerms returns a boolean if a field has been set.

### GetExclude

`func (o *SearchFilters) GetExclude() bool`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *SearchFilters) GetExcludeOk() (*bool, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *SearchFilters) SetExclude(v bool)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *SearchFilters) HasExclude() bool`

HasExclude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


