# GenerateRandomString

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | This must always be set to \&quot;Cloud Services Deployment Utility\&quot; | 
**Operation** | **string** | The operation to perform &#x60;generateRandomString&#x60; | 
**IncludeNumbers** | **bool** | This must be either \&quot;true\&quot; or \&quot;false\&quot; to indicate whether the generator logic should include numbers | 
**IncludeSpecialChars** | **bool** | This must be either \&quot;true\&quot; or \&quot;false\&quot; to indicate whether the generator logic should include special characters | 
**Length** | **string** | This specifies how long the randomly generated string needs to be   &gt;NOTE Due to identity attribute data constraints, the maximum allowable value is 450 characters  | 
**RequiresPeriodicRefresh** | Pointer to **bool** | A value that indicates whether the transform logic should be re-evaluated every evening as part of the identity refresh process | [optional] 

## Methods

### NewGenerateRandomString

`func NewGenerateRandomString(name string, operation string, includeNumbers bool, includeSpecialChars bool, length string, ) *GenerateRandomString`

NewGenerateRandomString instantiates a new GenerateRandomString object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateRandomStringWithDefaults

`func NewGenerateRandomStringWithDefaults() *GenerateRandomString`

NewGenerateRandomStringWithDefaults instantiates a new GenerateRandomString object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GenerateRandomString) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GenerateRandomString) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GenerateRandomString) SetName(v string)`

SetName sets Name field to given value.


### GetOperation

`func (o *GenerateRandomString) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *GenerateRandomString) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *GenerateRandomString) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetIncludeNumbers

`func (o *GenerateRandomString) GetIncludeNumbers() bool`

GetIncludeNumbers returns the IncludeNumbers field if non-nil, zero value otherwise.

### GetIncludeNumbersOk

`func (o *GenerateRandomString) GetIncludeNumbersOk() (*bool, bool)`

GetIncludeNumbersOk returns a tuple with the IncludeNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeNumbers

`func (o *GenerateRandomString) SetIncludeNumbers(v bool)`

SetIncludeNumbers sets IncludeNumbers field to given value.


### GetIncludeSpecialChars

`func (o *GenerateRandomString) GetIncludeSpecialChars() bool`

GetIncludeSpecialChars returns the IncludeSpecialChars field if non-nil, zero value otherwise.

### GetIncludeSpecialCharsOk

`func (o *GenerateRandomString) GetIncludeSpecialCharsOk() (*bool, bool)`

GetIncludeSpecialCharsOk returns a tuple with the IncludeSpecialChars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSpecialChars

`func (o *GenerateRandomString) SetIncludeSpecialChars(v bool)`

SetIncludeSpecialChars sets IncludeSpecialChars field to given value.


### GetLength

`func (o *GenerateRandomString) GetLength() string`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *GenerateRandomString) GetLengthOk() (*string, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *GenerateRandomString) SetLength(v string)`

SetLength sets Length field to given value.


### GetRequiresPeriodicRefresh

`func (o *GenerateRandomString) GetRequiresPeriodicRefresh() bool`

GetRequiresPeriodicRefresh returns the RequiresPeriodicRefresh field if non-nil, zero value otherwise.

### GetRequiresPeriodicRefreshOk

`func (o *GenerateRandomString) GetRequiresPeriodicRefreshOk() (*bool, bool)`

GetRequiresPeriodicRefreshOk returns a tuple with the RequiresPeriodicRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPeriodicRefresh

`func (o *GenerateRandomString) SetRequiresPeriodicRefresh(v bool)`

SetRequiresPeriodicRefresh sets RequiresPeriodicRefresh field to given value.

### HasRequiresPeriodicRefresh

`func (o *GenerateRandomString) HasRequiresPeriodicRefresh() bool`

HasRequiresPeriodicRefresh returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


