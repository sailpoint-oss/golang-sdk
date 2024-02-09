# Substring

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Begin** | **int32** | The index of the first character to include in the returned substring.   If &#x60;begin&#x60; is set to -1, the transform will begin at character 0 of the input data  | 
**BeginOffset** | Pointer to **int32** | This integer value is the number of characters to add to the begin attribute when returning a substring.   This attribute is only used if begin is not -1.  | [optional] 
**End** | Pointer to **int32** | The index of the first character to exclude from the returned substring.  If end is -1 or not provided at all, the substring transform will return everything up to the end of the input string.  | [optional] 
**EndOffset** | Pointer to **int32** | This integer value is the number of characters to add to the end attribute when returning a substring.   This attribute is only used if end is provided and is not -1.  | [optional] 
**RequiresPeriodicRefresh** | Pointer to **bool** | A value that indicates whether the transform logic should be re-evaluated every evening as part of the identity refresh process | [optional] [default to false]
**Input** | Pointer to **map[string]interface{}** | This is an optional attribute that can explicitly define the input data which will be fed into the transform logic. If input is not provided, the transform will take its input from the source and attribute combination configured via the UI. | [optional] 

## Methods

### NewSubstring

`func NewSubstring(begin int32, ) *Substring`

NewSubstring instantiates a new Substring object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubstringWithDefaults

`func NewSubstringWithDefaults() *Substring`

NewSubstringWithDefaults instantiates a new Substring object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBegin

`func (o *Substring) GetBegin() int32`

GetBegin returns the Begin field if non-nil, zero value otherwise.

### GetBeginOk

`func (o *Substring) GetBeginOk() (*int32, bool)`

GetBeginOk returns a tuple with the Begin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBegin

`func (o *Substring) SetBegin(v int32)`

SetBegin sets Begin field to given value.


### GetBeginOffset

`func (o *Substring) GetBeginOffset() int32`

GetBeginOffset returns the BeginOffset field if non-nil, zero value otherwise.

### GetBeginOffsetOk

`func (o *Substring) GetBeginOffsetOk() (*int32, bool)`

GetBeginOffsetOk returns a tuple with the BeginOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeginOffset

`func (o *Substring) SetBeginOffset(v int32)`

SetBeginOffset sets BeginOffset field to given value.

### HasBeginOffset

`func (o *Substring) HasBeginOffset() bool`

HasBeginOffset returns a boolean if a field has been set.

### GetEnd

`func (o *Substring) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Substring) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Substring) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *Substring) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetEndOffset

`func (o *Substring) GetEndOffset() int32`

GetEndOffset returns the EndOffset field if non-nil, zero value otherwise.

### GetEndOffsetOk

`func (o *Substring) GetEndOffsetOk() (*int32, bool)`

GetEndOffsetOk returns a tuple with the EndOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOffset

`func (o *Substring) SetEndOffset(v int32)`

SetEndOffset sets EndOffset field to given value.

### HasEndOffset

`func (o *Substring) HasEndOffset() bool`

HasEndOffset returns a boolean if a field has been set.

### GetRequiresPeriodicRefresh

`func (o *Substring) GetRequiresPeriodicRefresh() bool`

GetRequiresPeriodicRefresh returns the RequiresPeriodicRefresh field if non-nil, zero value otherwise.

### GetRequiresPeriodicRefreshOk

`func (o *Substring) GetRequiresPeriodicRefreshOk() (*bool, bool)`

GetRequiresPeriodicRefreshOk returns a tuple with the RequiresPeriodicRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPeriodicRefresh

`func (o *Substring) SetRequiresPeriodicRefresh(v bool)`

SetRequiresPeriodicRefresh sets RequiresPeriodicRefresh field to given value.

### HasRequiresPeriodicRefresh

`func (o *Substring) HasRequiresPeriodicRefresh() bool`

HasRequiresPeriodicRefresh returns a boolean if a field has been set.

### GetInput

`func (o *Substring) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *Substring) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *Substring) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *Substring) HasInput() bool`

HasInput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


