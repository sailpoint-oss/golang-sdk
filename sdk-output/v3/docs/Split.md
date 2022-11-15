# Split

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delimiter** | **string** | This can be either a single character or a regex expression, and is used by the transform to identify the break point between two substrings in the incoming data | 
**Index** | **string** | An integer value for the desired array element after the incoming data has been split into a list; the array is a 0-based object, so the first array element would be index 0, the second element would be index 1, etc. | 
**Throws** | Pointer to **bool** | A boolean (true/false) value which indicates whether an exception should be thrown and returned as an output when an index is out of bounds with the resultant array (i.e., the provided index value is larger than the size of the array)   &#x60;true&#x60; - The transform should return \&quot;IndexOutOfBoundsException\&quot;   &#x60;false&#x60; - The transform should return null   If not provided, the transform will default to false and return a null  | [optional] 
**RequiresPeriodicRefresh** | Pointer to **bool** | A value that indicates whether the transform logic should be re-evaluated every evening as part of the identity refresh process | [optional] [default to false]
**Input** | Pointer to **map[string]interface{}** | This is an optional attribute that can explicitly define the input data which will be fed into the transform logic. If input is not provided, the transform will take its input from the source and attribute combination configured via the UI. | [optional] 

## Methods

### NewSplit

`func NewSplit(delimiter string, index string, ) *Split`

NewSplit instantiates a new Split object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSplitWithDefaults

`func NewSplitWithDefaults() *Split`

NewSplitWithDefaults instantiates a new Split object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelimiter

`func (o *Split) GetDelimiter() string`

GetDelimiter returns the Delimiter field if non-nil, zero value otherwise.

### GetDelimiterOk

`func (o *Split) GetDelimiterOk() (*string, bool)`

GetDelimiterOk returns a tuple with the Delimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelimiter

`func (o *Split) SetDelimiter(v string)`

SetDelimiter sets Delimiter field to given value.


### GetIndex

`func (o *Split) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *Split) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *Split) SetIndex(v string)`

SetIndex sets Index field to given value.


### GetThrows

`func (o *Split) GetThrows() bool`

GetThrows returns the Throws field if non-nil, zero value otherwise.

### GetThrowsOk

`func (o *Split) GetThrowsOk() (*bool, bool)`

GetThrowsOk returns a tuple with the Throws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrows

`func (o *Split) SetThrows(v bool)`

SetThrows sets Throws field to given value.

### HasThrows

`func (o *Split) HasThrows() bool`

HasThrows returns a boolean if a field has been set.

### GetRequiresPeriodicRefresh

`func (o *Split) GetRequiresPeriodicRefresh() bool`

GetRequiresPeriodicRefresh returns the RequiresPeriodicRefresh field if non-nil, zero value otherwise.

### GetRequiresPeriodicRefreshOk

`func (o *Split) GetRequiresPeriodicRefreshOk() (*bool, bool)`

GetRequiresPeriodicRefreshOk returns a tuple with the RequiresPeriodicRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPeriodicRefresh

`func (o *Split) SetRequiresPeriodicRefresh(v bool)`

SetRequiresPeriodicRefresh sets RequiresPeriodicRefresh field to given value.

### HasRequiresPeriodicRefresh

`func (o *Split) HasRequiresPeriodicRefresh() bool`

HasRequiresPeriodicRefresh returns a boolean if a field has been set.

### GetInput

`func (o *Split) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *Split) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *Split) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *Split) HasInput() bool`

HasInput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


