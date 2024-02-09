# Rule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | This must always be set to \&quot;Cloud Services Deployment Utility\&quot; | 
**RequiresPeriodicRefresh** | Pointer to **bool** | A value that indicates whether the transform logic should be re-evaluated every evening as part of the identity refresh process | [optional] 
**Operation** | **string** | The operation to perform &#x60;getReferenceIdentityAttribute&#x60; | 
**IncludeNumbers** | **bool** | This must be either \&quot;true\&quot; or \&quot;false\&quot; to indicate whether the generator logic should include numbers | 
**IncludeSpecialChars** | **bool** | This must be either \&quot;true\&quot; or \&quot;false\&quot; to indicate whether the generator logic should include special characters | 
**Length** | **string** | This specifies how long the randomly generated string needs to be   &gt;NOTE Due to identity attribute data constraints, the maximum allowable value is 450 characters  | 
**Uid** | **string** | This is the SailPoint User Name (uid) value of the identity whose attribute is desired  As a convenience feature, you can use the &#x60;manager&#x60; keyword to dynamically look up the user&#39;s manager and then get that manager&#39;s identity attribute.  | 

## Methods

### NewRule

`func NewRule(name string, operation string, includeNumbers bool, includeSpecialChars bool, length string, uid string, ) *Rule`

NewRule instantiates a new Rule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleWithDefaults

`func NewRuleWithDefaults() *Rule`

NewRuleWithDefaults instantiates a new Rule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Rule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Rule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Rule) SetName(v string)`

SetName sets Name field to given value.


### GetRequiresPeriodicRefresh

`func (o *Rule) GetRequiresPeriodicRefresh() bool`

GetRequiresPeriodicRefresh returns the RequiresPeriodicRefresh field if non-nil, zero value otherwise.

### GetRequiresPeriodicRefreshOk

`func (o *Rule) GetRequiresPeriodicRefreshOk() (*bool, bool)`

GetRequiresPeriodicRefreshOk returns a tuple with the RequiresPeriodicRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPeriodicRefresh

`func (o *Rule) SetRequiresPeriodicRefresh(v bool)`

SetRequiresPeriodicRefresh sets RequiresPeriodicRefresh field to given value.

### HasRequiresPeriodicRefresh

`func (o *Rule) HasRequiresPeriodicRefresh() bool`

HasRequiresPeriodicRefresh returns a boolean if a field has been set.

### GetOperation

`func (o *Rule) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *Rule) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *Rule) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetIncludeNumbers

`func (o *Rule) GetIncludeNumbers() bool`

GetIncludeNumbers returns the IncludeNumbers field if non-nil, zero value otherwise.

### GetIncludeNumbersOk

`func (o *Rule) GetIncludeNumbersOk() (*bool, bool)`

GetIncludeNumbersOk returns a tuple with the IncludeNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeNumbers

`func (o *Rule) SetIncludeNumbers(v bool)`

SetIncludeNumbers sets IncludeNumbers field to given value.


### GetIncludeSpecialChars

`func (o *Rule) GetIncludeSpecialChars() bool`

GetIncludeSpecialChars returns the IncludeSpecialChars field if non-nil, zero value otherwise.

### GetIncludeSpecialCharsOk

`func (o *Rule) GetIncludeSpecialCharsOk() (*bool, bool)`

GetIncludeSpecialCharsOk returns a tuple with the IncludeSpecialChars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSpecialChars

`func (o *Rule) SetIncludeSpecialChars(v bool)`

SetIncludeSpecialChars sets IncludeSpecialChars field to given value.


### GetLength

`func (o *Rule) GetLength() string`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *Rule) GetLengthOk() (*string, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *Rule) SetLength(v string)`

SetLength sets Length field to given value.


### GetUid

`func (o *Rule) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Rule) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Rule) SetUid(v string)`

SetUid sets Uid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


