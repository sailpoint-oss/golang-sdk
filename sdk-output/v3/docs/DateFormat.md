# DateFormat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputFormat** | Pointer to [**DateFormatInputFormat**](DateFormatInputFormat.md) |  | [optional] 
**OutputFormat** | Pointer to [**DateFormatOutputFormat**](DateFormatOutputFormat.md) |  | [optional] 
**RequiresPeriodicRefresh** | Pointer to **bool** | A value that indicates whether the transform logic should be re-evaluated every evening as part of the identity refresh process | [optional] [default to false]
**Input** | Pointer to **map[string]interface{}** | This is an optional attribute that can explicitly define the input data which will be fed into the transform logic. If input is not provided, the transform will take its input from the source and attribute combination configured via the UI. | [optional] 

## Methods

### NewDateFormat

`func NewDateFormat() *DateFormat`

NewDateFormat instantiates a new DateFormat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDateFormatWithDefaults

`func NewDateFormatWithDefaults() *DateFormat`

NewDateFormatWithDefaults instantiates a new DateFormat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputFormat

`func (o *DateFormat) GetInputFormat() DateFormatInputFormat`

GetInputFormat returns the InputFormat field if non-nil, zero value otherwise.

### GetInputFormatOk

`func (o *DateFormat) GetInputFormatOk() (*DateFormatInputFormat, bool)`

GetInputFormatOk returns a tuple with the InputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFormat

`func (o *DateFormat) SetInputFormat(v DateFormatInputFormat)`

SetInputFormat sets InputFormat field to given value.

### HasInputFormat

`func (o *DateFormat) HasInputFormat() bool`

HasInputFormat returns a boolean if a field has been set.

### GetOutputFormat

`func (o *DateFormat) GetOutputFormat() DateFormatOutputFormat`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *DateFormat) GetOutputFormatOk() (*DateFormatOutputFormat, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *DateFormat) SetOutputFormat(v DateFormatOutputFormat)`

SetOutputFormat sets OutputFormat field to given value.

### HasOutputFormat

`func (o *DateFormat) HasOutputFormat() bool`

HasOutputFormat returns a boolean if a field has been set.

### GetRequiresPeriodicRefresh

`func (o *DateFormat) GetRequiresPeriodicRefresh() bool`

GetRequiresPeriodicRefresh returns the RequiresPeriodicRefresh field if non-nil, zero value otherwise.

### GetRequiresPeriodicRefreshOk

`func (o *DateFormat) GetRequiresPeriodicRefreshOk() (*bool, bool)`

GetRequiresPeriodicRefreshOk returns a tuple with the RequiresPeriodicRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPeriodicRefresh

`func (o *DateFormat) SetRequiresPeriodicRefresh(v bool)`

SetRequiresPeriodicRefresh sets RequiresPeriodicRefresh field to given value.

### HasRequiresPeriodicRefresh

`func (o *DateFormat) HasRequiresPeriodicRefresh() bool`

HasRequiresPeriodicRefresh returns a boolean if a field has been set.

### GetInput

`func (o *DateFormat) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *DateFormat) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *DateFormat) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *DateFormat) HasInput() bool`

HasInput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


