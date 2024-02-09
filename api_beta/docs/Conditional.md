# Conditional

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expression** | **string** | A comparison statement that follows the structure of &#x60;ValueA eq ValueB&#x60; where &#x60;ValueA&#x60; and &#x60;ValueB&#x60; are static strings or outputs of other transforms.   The &#x60;eq&#x60; operator is the only valid comparison | 
**PositiveCondition** | **string** | The output of the transform if the expression evalutes to true | 
**NegativeCondition** | **string** | The output of the transform if the expression evalutes to false | 
**RequiresPeriodicRefresh** | Pointer to **bool** | A value that indicates whether the transform logic should be re-evaluated every evening as part of the identity refresh process | [optional] [default to false]
**Input** | Pointer to **map[string]interface{}** | This is an optional attribute that can explicitly define the input data which will be fed into the transform logic. If input is not provided, the transform will take its input from the source and attribute combination configured via the UI. | [optional] 

## Methods

### NewConditional

`func NewConditional(expression string, positiveCondition string, negativeCondition string, ) *Conditional`

NewConditional instantiates a new Conditional object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionalWithDefaults

`func NewConditionalWithDefaults() *Conditional`

NewConditionalWithDefaults instantiates a new Conditional object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpression

`func (o *Conditional) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *Conditional) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *Conditional) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetPositiveCondition

`func (o *Conditional) GetPositiveCondition() string`

GetPositiveCondition returns the PositiveCondition field if non-nil, zero value otherwise.

### GetPositiveConditionOk

`func (o *Conditional) GetPositiveConditionOk() (*string, bool)`

GetPositiveConditionOk returns a tuple with the PositiveCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositiveCondition

`func (o *Conditional) SetPositiveCondition(v string)`

SetPositiveCondition sets PositiveCondition field to given value.


### GetNegativeCondition

`func (o *Conditional) GetNegativeCondition() string`

GetNegativeCondition returns the NegativeCondition field if non-nil, zero value otherwise.

### GetNegativeConditionOk

`func (o *Conditional) GetNegativeConditionOk() (*string, bool)`

GetNegativeConditionOk returns a tuple with the NegativeCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeCondition

`func (o *Conditional) SetNegativeCondition(v string)`

SetNegativeCondition sets NegativeCondition field to given value.


### GetRequiresPeriodicRefresh

`func (o *Conditional) GetRequiresPeriodicRefresh() bool`

GetRequiresPeriodicRefresh returns the RequiresPeriodicRefresh field if non-nil, zero value otherwise.

### GetRequiresPeriodicRefreshOk

`func (o *Conditional) GetRequiresPeriodicRefreshOk() (*bool, bool)`

GetRequiresPeriodicRefreshOk returns a tuple with the RequiresPeriodicRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPeriodicRefresh

`func (o *Conditional) SetRequiresPeriodicRefresh(v bool)`

SetRequiresPeriodicRefresh sets RequiresPeriodicRefresh field to given value.

### HasRequiresPeriodicRefresh

`func (o *Conditional) HasRequiresPeriodicRefresh() bool`

HasRequiresPeriodicRefresh returns a boolean if a field has been set.

### GetInput

`func (o *Conditional) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *Conditional) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *Conditional) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *Conditional) HasInput() bool`

HasInput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


