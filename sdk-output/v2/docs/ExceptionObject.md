# ExceptionObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** |  | 
**Message** | **string** |  | 
**Errors** | Pointer to [**[]ExceptionObjectErrorsInner**](ExceptionObjectErrorsInner.md) |  | [optional] 

## Methods

### NewExceptionObject

`func NewExceptionObject(code int32, message string, ) *ExceptionObject`

NewExceptionObject instantiates a new ExceptionObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExceptionObjectWithDefaults

`func NewExceptionObjectWithDefaults() *ExceptionObject`

NewExceptionObjectWithDefaults instantiates a new ExceptionObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ExceptionObject) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ExceptionObject) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ExceptionObject) SetCode(v int32)`

SetCode sets Code field to given value.


### GetMessage

`func (o *ExceptionObject) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ExceptionObject) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ExceptionObject) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetErrors

`func (o *ExceptionObject) GetErrors() []ExceptionObjectErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ExceptionObject) GetErrorsOk() (*[]ExceptionObjectErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ExceptionObject) SetErrors(v []ExceptionObjectErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ExceptionObject) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


