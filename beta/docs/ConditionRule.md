# ConditionRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to **string** | Operator is a ConditionRuleComparisonOperatorType value EQ ConditionRuleComparisonOperatorTypeEquals  ConditionRuleComparisonOperatorTypeEquals is a comparison operator, the source and target are compared for equality NE ConditionRuleComparisonOperatorTypeNotEquals  ConditionRuleComparisonOperatorTypeNotEquals is a comparison operator, the source and target are compared for the opposite of equality CO ConditionRuleComparisonOperatorTypeContains  ConditionRuleComparisonOperatorTypeContains is a comparison operator, the source is searched to see if it contains the value NOT_CO ConditionRuleComparisonOperatorTypeNotContains IN ConditionRuleComparisonOperatorTypeIncludes  ConditionRuleComparisonOperatorTypeIncludes is a comparison operator, the source will be searched if it equals any of the values NOT_IN ConditionRuleComparisonOperatorTypeNotIncludes EM ConditionRuleComparisonOperatorTypeEmpty NOT_EM ConditionRuleComparisonOperatorTypeNotEmpty SW ConditionRuleComparisonOperatorTypeStartsWith  ConditionRuleComparisonOperatorTypeStartsWith checks if a string starts with another substring of the same string, this operator is case-sensitive NOT_SW ConditionRuleComparisonOperatorTypeNotStartsWith EW ConditionRuleComparisonOperatorTypeEndsWith  ConditionRuleComparisonOperatorTypeEndsWith checks if a string ends with another substring of the same string, this operator is case-sensitive NOT_EW ConditionRuleComparisonOperatorTypeNotEndsWith | [optional] 
**Source** | Pointer to **string** | Source, if the sourceType is ConditionRuleSourceTypeInput then the source type is the name of the form input to accept. While if the sourceType is ConditionRuleSourceTypeElement then source is the name of a technical key of an element to retrieve its value | [optional] 
**SourceType** | Pointer to **string** | SourceType defines what type of object is being selected. Either a reference to a form input (by input name), or a form element (by technical key) INPUT ConditionRuleSourceTypeInput ELEMENT ConditionRuleSourceTypeElement | [optional] 
**Value** | Pointer to **map[string]interface{}** | Value is the value based on the ValueType | [optional] 
**ValueType** | Pointer to **string** | ValueType is a ConditionRuleValueType type STRING ConditionRuleValueTypeString  ConditionRuleValueTypeString the value field is a static string STRING_LIST ConditionRuleValueTypeStringList  ConditionRuleValueTypeStringList the value field is an array of string values INPUT ConditionRuleValueTypeInput  ConditionRuleValueTypeInput the value field is a reference to a form input by ELEMENT ConditionRuleValueTypeElement  ConditionRuleValueTypeElement the value field is a reference to form element (by technical key) LIST ConditionRuleValueTypeList BOOLEAN ConditionRuleValueTypeBoolean | [optional] 

## Methods

### NewConditionRule

`func NewConditionRule() *ConditionRule`

NewConditionRule instantiates a new ConditionRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionRuleWithDefaults

`func NewConditionRuleWithDefaults() *ConditionRule`

NewConditionRuleWithDefaults instantiates a new ConditionRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *ConditionRule) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ConditionRule) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ConditionRule) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *ConditionRule) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetSource

`func (o *ConditionRule) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ConditionRule) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ConditionRule) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ConditionRule) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceType

`func (o *ConditionRule) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *ConditionRule) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *ConditionRule) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *ConditionRule) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetValue

`func (o *ConditionRule) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConditionRule) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConditionRule) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *ConditionRule) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueType

`func (o *ConditionRule) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *ConditionRule) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *ConditionRule) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *ConditionRule) HasValueType() bool`

HasValueType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


