# RandomAlphaNumeric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Length** | Pointer to **string** | This is an integer value specifying the size/number of characters the random string must contain   * This value must be a positive number and cannot be blank   * If no length is provided, the transform will default to a value of &#x60;32&#x60;   * Due to identity attribute data constraints, the maximum allowable value is &#x60;450&#x60; characters  | [optional] 
**RequiresPeriodicRefresh** | Pointer to **bool** | A value that indicates whether the transform logic should be re-evaluated every evening as part of the identity refresh process | [optional] [default to false]
**Input** | Pointer to **map[string]interface{}** | This is an optional attribute that can explicitly define the input data which will be fed into the transform logic. If input is not provided, the transform will take its input from the source and attribute combination configured via the UI. | [optional] 

## Methods

### NewRandomAlphaNumeric

`func NewRandomAlphaNumeric() *RandomAlphaNumeric`

NewRandomAlphaNumeric instantiates a new RandomAlphaNumeric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRandomAlphaNumericWithDefaults

`func NewRandomAlphaNumericWithDefaults() *RandomAlphaNumeric`

NewRandomAlphaNumericWithDefaults instantiates a new RandomAlphaNumeric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLength

`func (o *RandomAlphaNumeric) GetLength() string`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *RandomAlphaNumeric) GetLengthOk() (*string, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *RandomAlphaNumeric) SetLength(v string)`

SetLength sets Length field to given value.

### HasLength

`func (o *RandomAlphaNumeric) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetRequiresPeriodicRefresh

`func (o *RandomAlphaNumeric) GetRequiresPeriodicRefresh() bool`

GetRequiresPeriodicRefresh returns the RequiresPeriodicRefresh field if non-nil, zero value otherwise.

### GetRequiresPeriodicRefreshOk

`func (o *RandomAlphaNumeric) GetRequiresPeriodicRefreshOk() (*bool, bool)`

GetRequiresPeriodicRefreshOk returns a tuple with the RequiresPeriodicRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPeriodicRefresh

`func (o *RandomAlphaNumeric) SetRequiresPeriodicRefresh(v bool)`

SetRequiresPeriodicRefresh sets RequiresPeriodicRefresh field to given value.

### HasRequiresPeriodicRefresh

`func (o *RandomAlphaNumeric) HasRequiresPeriodicRefresh() bool`

HasRequiresPeriodicRefresh returns a boolean if a field has been set.

### GetInput

`func (o *RandomAlphaNumeric) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *RandomAlphaNumeric) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *RandomAlphaNumeric) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *RandomAlphaNumeric) HasInput() bool`

HasInput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


