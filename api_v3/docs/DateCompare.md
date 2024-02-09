# DateCompare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstDate** | [**DateCompareFirstDate**](DateCompareFirstDate.md) |  | 
**SecondDate** | [**DateCompareSecondDate**](DateCompareSecondDate.md) |  | 
**Operator** | **string** | This is the comparison to perform. | Operation | Description | | --------- | ------- | | LT        | Strictly less than: firstDate &lt; secondDate | | LTE       | Less than or equal to: firstDate &lt;&#x3D; secondDate | | GT        | Strictly greater than: firstDate &gt; secondDate | | GTE       | Greater than or equal to: firstDate &gt;&#x3D; secondDate |  | 
**PositiveCondition** | **string** | The output of the transform if the expression evalutes to true | 
**NegativeCondition** | **string** | The output of the transform if the expression evalutes to false | 
**RequiresPeriodicRefresh** | Pointer to **bool** | A value that indicates whether the transform logic should be re-evaluated every evening as part of the identity refresh process | [optional] [default to false]
**Input** | Pointer to **map[string]interface{}** | This is an optional attribute that can explicitly define the input data which will be fed into the transform logic. If input is not provided, the transform will take its input from the source and attribute combination configured via the UI. | [optional] 

## Methods

### NewDateCompare

`func NewDateCompare(firstDate DateCompareFirstDate, secondDate DateCompareSecondDate, operator string, positiveCondition string, negativeCondition string, ) *DateCompare`

NewDateCompare instantiates a new DateCompare object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDateCompareWithDefaults

`func NewDateCompareWithDefaults() *DateCompare`

NewDateCompareWithDefaults instantiates a new DateCompare object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstDate

`func (o *DateCompare) GetFirstDate() DateCompareFirstDate`

GetFirstDate returns the FirstDate field if non-nil, zero value otherwise.

### GetFirstDateOk

`func (o *DateCompare) GetFirstDateOk() (*DateCompareFirstDate, bool)`

GetFirstDateOk returns a tuple with the FirstDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstDate

`func (o *DateCompare) SetFirstDate(v DateCompareFirstDate)`

SetFirstDate sets FirstDate field to given value.


### GetSecondDate

`func (o *DateCompare) GetSecondDate() DateCompareSecondDate`

GetSecondDate returns the SecondDate field if non-nil, zero value otherwise.

### GetSecondDateOk

`func (o *DateCompare) GetSecondDateOk() (*DateCompareSecondDate, bool)`

GetSecondDateOk returns a tuple with the SecondDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondDate

`func (o *DateCompare) SetSecondDate(v DateCompareSecondDate)`

SetSecondDate sets SecondDate field to given value.


### GetOperator

`func (o *DateCompare) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *DateCompare) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *DateCompare) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetPositiveCondition

`func (o *DateCompare) GetPositiveCondition() string`

GetPositiveCondition returns the PositiveCondition field if non-nil, zero value otherwise.

### GetPositiveConditionOk

`func (o *DateCompare) GetPositiveConditionOk() (*string, bool)`

GetPositiveConditionOk returns a tuple with the PositiveCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositiveCondition

`func (o *DateCompare) SetPositiveCondition(v string)`

SetPositiveCondition sets PositiveCondition field to given value.


### GetNegativeCondition

`func (o *DateCompare) GetNegativeCondition() string`

GetNegativeCondition returns the NegativeCondition field if non-nil, zero value otherwise.

### GetNegativeConditionOk

`func (o *DateCompare) GetNegativeConditionOk() (*string, bool)`

GetNegativeConditionOk returns a tuple with the NegativeCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeCondition

`func (o *DateCompare) SetNegativeCondition(v string)`

SetNegativeCondition sets NegativeCondition field to given value.


### GetRequiresPeriodicRefresh

`func (o *DateCompare) GetRequiresPeriodicRefresh() bool`

GetRequiresPeriodicRefresh returns the RequiresPeriodicRefresh field if non-nil, zero value otherwise.

### GetRequiresPeriodicRefreshOk

`func (o *DateCompare) GetRequiresPeriodicRefreshOk() (*bool, bool)`

GetRequiresPeriodicRefreshOk returns a tuple with the RequiresPeriodicRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPeriodicRefresh

`func (o *DateCompare) SetRequiresPeriodicRefresh(v bool)`

SetRequiresPeriodicRefresh sets RequiresPeriodicRefresh field to given value.

### HasRequiresPeriodicRefresh

`func (o *DateCompare) HasRequiresPeriodicRefresh() bool`

HasRequiresPeriodicRefresh returns a boolean if a field has been set.

### GetInput

`func (o *DateCompare) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *DateCompare) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *DateCompare) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *DateCompare) HasInput() bool`

HasInput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


