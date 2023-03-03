# TriggerInputProvisioningCompleted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackingNumber** | **string** | The reference number of the provisioning request. Useful for tracking status in the Account Activity search interface. | 
**Sources** | **string** | One or more sources that the provisioning transaction(s) were done against.  Sources are comma separated. | 
**Action** | Pointer to **NullableString** | Origin of where the provisioning request came from. | [optional] 
**Errors** | Pointer to **[]string** | A list of any accumulated error messages that occurred during provisioning. | [optional] 
**Warnings** | Pointer to **[]string** | A list of any accumulated warning messages that occurred during provisioning. | [optional] 
**Recipient** | [**TriggerInputProvisioningCompletedRecipient**](TriggerInputProvisioningCompletedRecipient.md) |  | 
**Requester** | Pointer to [**NullableTriggerInputProvisioningCompletedRequester**](TriggerInputProvisioningCompletedRequester.md) |  | [optional] 
**AccountRequests** | [**[]TriggerInputProvisioningCompletedAccountRequestsInner**](TriggerInputProvisioningCompletedAccountRequestsInner.md) | A list of provisioning instructions to perform on an account-by-account basis. | 

## Methods

### NewTriggerInputProvisioningCompleted

`func NewTriggerInputProvisioningCompleted(trackingNumber string, sources string, recipient TriggerInputProvisioningCompletedRecipient, accountRequests []TriggerInputProvisioningCompletedAccountRequestsInner, ) *TriggerInputProvisioningCompleted`

NewTriggerInputProvisioningCompleted instantiates a new TriggerInputProvisioningCompleted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerInputProvisioningCompletedWithDefaults

`func NewTriggerInputProvisioningCompletedWithDefaults() *TriggerInputProvisioningCompleted`

NewTriggerInputProvisioningCompletedWithDefaults instantiates a new TriggerInputProvisioningCompleted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrackingNumber

`func (o *TriggerInputProvisioningCompleted) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *TriggerInputProvisioningCompleted) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *TriggerInputProvisioningCompleted) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.


### GetSources

`func (o *TriggerInputProvisioningCompleted) GetSources() string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *TriggerInputProvisioningCompleted) GetSourcesOk() (*string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *TriggerInputProvisioningCompleted) SetSources(v string)`

SetSources sets Sources field to given value.


### GetAction

`func (o *TriggerInputProvisioningCompleted) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TriggerInputProvisioningCompleted) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TriggerInputProvisioningCompleted) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *TriggerInputProvisioningCompleted) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *TriggerInputProvisioningCompleted) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *TriggerInputProvisioningCompleted) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetErrors

`func (o *TriggerInputProvisioningCompleted) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *TriggerInputProvisioningCompleted) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *TriggerInputProvisioningCompleted) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *TriggerInputProvisioningCompleted) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *TriggerInputProvisioningCompleted) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *TriggerInputProvisioningCompleted) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetWarnings

`func (o *TriggerInputProvisioningCompleted) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *TriggerInputProvisioningCompleted) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *TriggerInputProvisioningCompleted) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *TriggerInputProvisioningCompleted) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarningsNil

`func (o *TriggerInputProvisioningCompleted) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *TriggerInputProvisioningCompleted) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil
### GetRecipient

`func (o *TriggerInputProvisioningCompleted) GetRecipient() TriggerInputProvisioningCompletedRecipient`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *TriggerInputProvisioningCompleted) GetRecipientOk() (*TriggerInputProvisioningCompletedRecipient, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *TriggerInputProvisioningCompleted) SetRecipient(v TriggerInputProvisioningCompletedRecipient)`

SetRecipient sets Recipient field to given value.


### GetRequester

`func (o *TriggerInputProvisioningCompleted) GetRequester() TriggerInputProvisioningCompletedRequester`

GetRequester returns the Requester field if non-nil, zero value otherwise.

### GetRequesterOk

`func (o *TriggerInputProvisioningCompleted) GetRequesterOk() (*TriggerInputProvisioningCompletedRequester, bool)`

GetRequesterOk returns a tuple with the Requester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequester

`func (o *TriggerInputProvisioningCompleted) SetRequester(v TriggerInputProvisioningCompletedRequester)`

SetRequester sets Requester field to given value.

### HasRequester

`func (o *TriggerInputProvisioningCompleted) HasRequester() bool`

HasRequester returns a boolean if a field has been set.

### SetRequesterNil

`func (o *TriggerInputProvisioningCompleted) SetRequesterNil(b bool)`

 SetRequesterNil sets the value for Requester to be an explicit nil

### UnsetRequester
`func (o *TriggerInputProvisioningCompleted) UnsetRequester()`

UnsetRequester ensures that no value is present for Requester, not even an explicit nil
### GetAccountRequests

`func (o *TriggerInputProvisioningCompleted) GetAccountRequests() []TriggerInputProvisioningCompletedAccountRequestsInner`

GetAccountRequests returns the AccountRequests field if non-nil, zero value otherwise.

### GetAccountRequestsOk

`func (o *TriggerInputProvisioningCompleted) GetAccountRequestsOk() (*[]TriggerInputProvisioningCompletedAccountRequestsInner, bool)`

GetAccountRequestsOk returns a tuple with the AccountRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountRequests

`func (o *TriggerInputProvisioningCompleted) SetAccountRequests(v []TriggerInputProvisioningCompletedAccountRequestsInner)`

SetAccountRequests sets AccountRequests field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


