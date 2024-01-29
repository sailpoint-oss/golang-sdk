# CompleteInvocationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalizedError** | Pointer to [**LocalizedMessage**](LocalizedMessage.md) |  | [optional] 
**Output** | Pointer to **map[string]interface{}** | Trigger output that completed the invocation. Its schema is defined in the trigger definition. | [optional] 

## Methods

### NewCompleteInvocationInput

`func NewCompleteInvocationInput() *CompleteInvocationInput`

NewCompleteInvocationInput instantiates a new CompleteInvocationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompleteInvocationInputWithDefaults

`func NewCompleteInvocationInputWithDefaults() *CompleteInvocationInput`

NewCompleteInvocationInputWithDefaults instantiates a new CompleteInvocationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalizedError

`func (o *CompleteInvocationInput) GetLocalizedError() LocalizedMessage`

GetLocalizedError returns the LocalizedError field if non-nil, zero value otherwise.

### GetLocalizedErrorOk

`func (o *CompleteInvocationInput) GetLocalizedErrorOk() (*LocalizedMessage, bool)`

GetLocalizedErrorOk returns a tuple with the LocalizedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedError

`func (o *CompleteInvocationInput) SetLocalizedError(v LocalizedMessage)`

SetLocalizedError sets LocalizedError field to given value.

### HasLocalizedError

`func (o *CompleteInvocationInput) HasLocalizedError() bool`

HasLocalizedError returns a boolean if a field has been set.

### GetOutput

`func (o *CompleteInvocationInput) GetOutput() map[string]interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *CompleteInvocationInput) GetOutputOk() (*map[string]interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *CompleteInvocationInput) SetOutput(v map[string]interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *CompleteInvocationInput) HasOutput() bool`

HasOutput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


