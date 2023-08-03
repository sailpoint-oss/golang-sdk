# TextQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Terms** | **[]string** | Words or characters that specify a particular thing to be searched for. | 
**Fields** | **[]string** | The fields to be searched. | 
**MatchAny** | Pointer to **bool** | Indicates if a match was found. | [optional] [default to false]
**Contains** | Pointer to **bool** | Indicates if the search contained a field. | [optional] [default to false]

## Methods

### NewTextQuery

`func NewTextQuery(terms []string, fields []string, ) *TextQuery`

NewTextQuery instantiates a new TextQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextQueryWithDefaults

`func NewTextQueryWithDefaults() *TextQuery`

NewTextQueryWithDefaults instantiates a new TextQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTerms

`func (o *TextQuery) GetTerms() []string`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *TextQuery) GetTermsOk() (*[]string, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *TextQuery) SetTerms(v []string)`

SetTerms sets Terms field to given value.


### GetFields

`func (o *TextQuery) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *TextQuery) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *TextQuery) SetFields(v []string)`

SetFields sets Fields field to given value.


### GetMatchAny

`func (o *TextQuery) GetMatchAny() bool`

GetMatchAny returns the MatchAny field if non-nil, zero value otherwise.

### GetMatchAnyOk

`func (o *TextQuery) GetMatchAnyOk() (*bool, bool)`

GetMatchAnyOk returns a tuple with the MatchAny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchAny

`func (o *TextQuery) SetMatchAny(v bool)`

SetMatchAny sets MatchAny field to given value.

### HasMatchAny

`func (o *TextQuery) HasMatchAny() bool`

HasMatchAny returns a boolean if a field has been set.

### GetContains

`func (o *TextQuery) GetContains() bool`

GetContains returns the Contains field if non-nil, zero value otherwise.

### GetContainsOk

`func (o *TextQuery) GetContainsOk() (*bool, bool)`

GetContainsOk returns a tuple with the Contains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContains

`func (o *TextQuery) SetContains(v bool)`

SetContains sets Contains field to given value.

### HasContains

`func (o *TextQuery) HasContains() bool`

HasContains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


