# LeftPad

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Length** | **string** | An integer value for the desired length of the final output string | 
**Padding** | Pointer to **string** | A string value representing the character that the incoming data should be padded with to get to the desired length   If not provided, the transform will default to a single space (\&quot; \&quot;) character for padding  | [optional] 
**RequiresPeriodicRefresh** | Pointer to **bool** | A value that indicates whether the transform logic should be re-evaluated every evening as part of the identity refresh process | [optional] [default to false]
**Input** | Pointer to **map[string]interface{}** | This is an optional attribute that can explicitly define the input data which will be fed into the transform logic. If input is not provided, the transform will take its input from the source and attribute combination configured via the UI. | [optional] 

## Methods

### NewLeftPad

`func NewLeftPad(length string, ) *LeftPad`

NewLeftPad instantiates a new LeftPad object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeftPadWithDefaults

`func NewLeftPadWithDefaults() *LeftPad`

NewLeftPadWithDefaults instantiates a new LeftPad object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLength

`func (o *LeftPad) GetLength() string`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *LeftPad) GetLengthOk() (*string, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *LeftPad) SetLength(v string)`

SetLength sets Length field to given value.


### GetPadding

`func (o *LeftPad) GetPadding() string`

GetPadding returns the Padding field if non-nil, zero value otherwise.

### GetPaddingOk

`func (o *LeftPad) GetPaddingOk() (*string, bool)`

GetPaddingOk returns a tuple with the Padding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPadding

`func (o *LeftPad) SetPadding(v string)`

SetPadding sets Padding field to given value.

### HasPadding

`func (o *LeftPad) HasPadding() bool`

HasPadding returns a boolean if a field has been set.

### GetRequiresPeriodicRefresh

`func (o *LeftPad) GetRequiresPeriodicRefresh() bool`

GetRequiresPeriodicRefresh returns the RequiresPeriodicRefresh field if non-nil, zero value otherwise.

### GetRequiresPeriodicRefreshOk

`func (o *LeftPad) GetRequiresPeriodicRefreshOk() (*bool, bool)`

GetRequiresPeriodicRefreshOk returns a tuple with the RequiresPeriodicRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPeriodicRefresh

`func (o *LeftPad) SetRequiresPeriodicRefresh(v bool)`

SetRequiresPeriodicRefresh sets RequiresPeriodicRefresh field to given value.

### HasRequiresPeriodicRefresh

`func (o *LeftPad) HasRequiresPeriodicRefresh() bool`

HasRequiresPeriodicRefresh returns a boolean if a field has been set.

### GetInput

`func (o *LeftPad) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *LeftPad) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *LeftPad) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *LeftPad) HasInput() bool`

HasInput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


