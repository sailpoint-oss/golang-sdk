# IndexOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Substring** | **string** | A substring to search for, searches the entire calling string, and returns the index of the first occurrence of the specified substring. | 
**RequiresPeriodicRefresh** | Pointer to **bool** | A value that indicates whether the transform logic should be re-evaluated every evening as part of the identity refresh process | [optional] [default to false]
**Input** | Pointer to **map[string]interface{}** | This is an optional attribute that can explicitly define the input data which will be fed into the transform logic. If input is not provided, the transform will take its input from the source and attribute combination configured via the UI. | [optional] 

## Methods

### NewIndexOf

`func NewIndexOf(substring string, ) *IndexOf`

NewIndexOf instantiates a new IndexOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexOfWithDefaults

`func NewIndexOfWithDefaults() *IndexOf`

NewIndexOfWithDefaults instantiates a new IndexOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubstring

`func (o *IndexOf) GetSubstring() string`

GetSubstring returns the Substring field if non-nil, zero value otherwise.

### GetSubstringOk

`func (o *IndexOf) GetSubstringOk() (*string, bool)`

GetSubstringOk returns a tuple with the Substring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstring

`func (o *IndexOf) SetSubstring(v string)`

SetSubstring sets Substring field to given value.


### GetRequiresPeriodicRefresh

`func (o *IndexOf) GetRequiresPeriodicRefresh() bool`

GetRequiresPeriodicRefresh returns the RequiresPeriodicRefresh field if non-nil, zero value otherwise.

### GetRequiresPeriodicRefreshOk

`func (o *IndexOf) GetRequiresPeriodicRefreshOk() (*bool, bool)`

GetRequiresPeriodicRefreshOk returns a tuple with the RequiresPeriodicRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPeriodicRefresh

`func (o *IndexOf) SetRequiresPeriodicRefresh(v bool)`

SetRequiresPeriodicRefresh sets RequiresPeriodicRefresh field to given value.

### HasRequiresPeriodicRefresh

`func (o *IndexOf) HasRequiresPeriodicRefresh() bool`

HasRequiresPeriodicRefresh returns a boolean if a field has been set.

### GetInput

`func (o *IndexOf) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *IndexOf) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *IndexOf) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *IndexOf) HasInput() bool`

HasInput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


