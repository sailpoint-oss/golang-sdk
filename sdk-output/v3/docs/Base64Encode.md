# Base64Encode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequiresPeriodicRefresh** | Pointer to **bool** | A value that indicates whether the transform logic should be re-evaluated every evening as part of the identity refresh process | [optional] [default to false]
**Input** | Pointer to **map[string]interface{}** | This is an optional attribute that can explicitly define the input data which will be fed into the transform logic. If input is not provided, the transform will take its input from the source and attribute combination configured via the UI. | [optional] 

## Methods

### NewBase64Encode

`func NewBase64Encode() *Base64Encode`

NewBase64Encode instantiates a new Base64Encode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBase64EncodeWithDefaults

`func NewBase64EncodeWithDefaults() *Base64Encode`

NewBase64EncodeWithDefaults instantiates a new Base64Encode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequiresPeriodicRefresh

`func (o *Base64Encode) GetRequiresPeriodicRefresh() bool`

GetRequiresPeriodicRefresh returns the RequiresPeriodicRefresh field if non-nil, zero value otherwise.

### GetRequiresPeriodicRefreshOk

`func (o *Base64Encode) GetRequiresPeriodicRefreshOk() (*bool, bool)`

GetRequiresPeriodicRefreshOk returns a tuple with the RequiresPeriodicRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPeriodicRefresh

`func (o *Base64Encode) SetRequiresPeriodicRefresh(v bool)`

SetRequiresPeriodicRefresh sets RequiresPeriodicRefresh field to given value.

### HasRequiresPeriodicRefresh

`func (o *Base64Encode) HasRequiresPeriodicRefresh() bool`

HasRequiresPeriodicRefresh returns a boolean if a field has been set.

### GetInput

`func (o *Base64Encode) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *Base64Encode) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *Base64Encode) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *Base64Encode) HasInput() bool`

HasInput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


