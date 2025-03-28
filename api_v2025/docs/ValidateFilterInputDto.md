# ValidateFilterInputDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | **map[string]interface{}** | Mock input to evaluate filter expression against. | 
**Filter** | **string** | JSONPath filter to conditionally invoke trigger when expression evaluates to true. | 

## Methods

### NewValidateFilterInputDto

`func NewValidateFilterInputDto(input map[string]interface{}, filter string, ) *ValidateFilterInputDto`

NewValidateFilterInputDto instantiates a new ValidateFilterInputDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateFilterInputDtoWithDefaults

`func NewValidateFilterInputDtoWithDefaults() *ValidateFilterInputDto`

NewValidateFilterInputDtoWithDefaults instantiates a new ValidateFilterInputDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *ValidateFilterInputDto) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ValidateFilterInputDto) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ValidateFilterInputDto) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.


### GetFilter

`func (o *ValidateFilterInputDto) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ValidateFilterInputDto) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ValidateFilterInputDto) SetFilter(v string)`

SetFilter sets Filter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


