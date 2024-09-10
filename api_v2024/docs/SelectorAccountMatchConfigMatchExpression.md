# SelectorAccountMatchConfigMatchExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchTerms** | Pointer to [**[]MatchTerm**](MatchTerm.md) |  | [optional] 
**And** | Pointer to **bool** | If it is AND operators for match terms | [optional] [default to true]

## Methods

### NewSelectorAccountMatchConfigMatchExpression

`func NewSelectorAccountMatchConfigMatchExpression() *SelectorAccountMatchConfigMatchExpression`

NewSelectorAccountMatchConfigMatchExpression instantiates a new SelectorAccountMatchConfigMatchExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectorAccountMatchConfigMatchExpressionWithDefaults

`func NewSelectorAccountMatchConfigMatchExpressionWithDefaults() *SelectorAccountMatchConfigMatchExpression`

NewSelectorAccountMatchConfigMatchExpressionWithDefaults instantiates a new SelectorAccountMatchConfigMatchExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchTerms

`func (o *SelectorAccountMatchConfigMatchExpression) GetMatchTerms() []MatchTerm`

GetMatchTerms returns the MatchTerms field if non-nil, zero value otherwise.

### GetMatchTermsOk

`func (o *SelectorAccountMatchConfigMatchExpression) GetMatchTermsOk() (*[]MatchTerm, bool)`

GetMatchTermsOk returns a tuple with the MatchTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchTerms

`func (o *SelectorAccountMatchConfigMatchExpression) SetMatchTerms(v []MatchTerm)`

SetMatchTerms sets MatchTerms field to given value.

### HasMatchTerms

`func (o *SelectorAccountMatchConfigMatchExpression) HasMatchTerms() bool`

HasMatchTerms returns a boolean if a field has been set.

### GetAnd

`func (o *SelectorAccountMatchConfigMatchExpression) GetAnd() bool`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *SelectorAccountMatchConfigMatchExpression) GetAndOk() (*bool, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *SelectorAccountMatchConfigMatchExpression) SetAnd(v bool)`

SetAnd sets And field to given value.

### HasAnd

`func (o *SelectorAccountMatchConfigMatchExpression) HasAnd() bool`

HasAnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

