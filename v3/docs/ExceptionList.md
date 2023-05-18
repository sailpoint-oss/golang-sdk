# ExceptionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exceptions** | Pointer to [**[]Exception**](Exception.md) | List of exceptions. | [optional] 

## Methods

### NewExceptionList

`func NewExceptionList() *ExceptionList`

NewExceptionList instantiates a new ExceptionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExceptionListWithDefaults

`func NewExceptionListWithDefaults() *ExceptionList`

NewExceptionListWithDefaults instantiates a new ExceptionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExceptions

`func (o *ExceptionList) GetExceptions() []Exception`

GetExceptions returns the Exceptions field if non-nil, zero value otherwise.

### GetExceptionsOk

`func (o *ExceptionList) GetExceptionsOk() (*[]Exception, bool)`

GetExceptionsOk returns a tuple with the Exceptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptions

`func (o *ExceptionList) SetExceptions(v []Exception)`

SetExceptions sets Exceptions field to given value.

### HasExceptions

`func (o *ExceptionList) HasExceptions() bool`

HasExceptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


