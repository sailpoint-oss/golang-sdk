# AccountActivitySearchedItemAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The type of action that this activity performed | [optional] 
**Created** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**Modified** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**Stage** | Pointer to **string** | The current stage of the activity | [optional] 
**Origin** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** | the current status of the activity | [optional] 
**Requester** | Pointer to [**Source1**](Source1.md) |  | [optional] 
**Recipient** | Pointer to [**Source1**](Source1.md) |  | [optional] 
**TrackingNumber** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to **[]string** |  | [optional] 
**Warnings** | Pointer to **[]string** |  | [optional] 
**Approvals** | Pointer to [**[]Approval**](Approval.md) |  | [optional] 
**OriginalRequests** | Pointer to [**[]OriginalRequest**](OriginalRequest.md) |  | [optional] 
**ExpansionItems** | Pointer to [**[]ExpansionItem**](ExpansionItem.md) |  | [optional] 
**AccountRequests** | Pointer to [**[]AccountRequest**](AccountRequest.md) |  | [optional] 
**Sources** | Pointer to **string** |  | [optional] 

## Methods

### NewAccountActivitySearchedItemAllOf

`func NewAccountActivitySearchedItemAllOf() *AccountActivitySearchedItemAllOf`

NewAccountActivitySearchedItemAllOf instantiates a new AccountActivitySearchedItemAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountActivitySearchedItemAllOfWithDefaults

`func NewAccountActivitySearchedItemAllOfWithDefaults() *AccountActivitySearchedItemAllOf`

NewAccountActivitySearchedItemAllOfWithDefaults instantiates a new AccountActivitySearchedItemAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AccountActivitySearchedItemAllOf) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AccountActivitySearchedItemAllOf) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AccountActivitySearchedItemAllOf) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *AccountActivitySearchedItemAllOf) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCreated

`func (o *AccountActivitySearchedItemAllOf) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AccountActivitySearchedItemAllOf) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AccountActivitySearchedItemAllOf) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AccountActivitySearchedItemAllOf) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *AccountActivitySearchedItemAllOf) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *AccountActivitySearchedItemAllOf) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetModified

`func (o *AccountActivitySearchedItemAllOf) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *AccountActivitySearchedItemAllOf) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *AccountActivitySearchedItemAllOf) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *AccountActivitySearchedItemAllOf) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModifiedNil

`func (o *AccountActivitySearchedItemAllOf) SetModifiedNil(b bool)`

 SetModifiedNil sets the value for Modified to be an explicit nil

### UnsetModified
`func (o *AccountActivitySearchedItemAllOf) UnsetModified()`

UnsetModified ensures that no value is present for Modified, not even an explicit nil
### GetStage

`func (o *AccountActivitySearchedItemAllOf) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *AccountActivitySearchedItemAllOf) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *AccountActivitySearchedItemAllOf) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *AccountActivitySearchedItemAllOf) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetOrigin

`func (o *AccountActivitySearchedItemAllOf) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *AccountActivitySearchedItemAllOf) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *AccountActivitySearchedItemAllOf) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *AccountActivitySearchedItemAllOf) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### SetOriginNil

`func (o *AccountActivitySearchedItemAllOf) SetOriginNil(b bool)`

 SetOriginNil sets the value for Origin to be an explicit nil

### UnsetOrigin
`func (o *AccountActivitySearchedItemAllOf) UnsetOrigin()`

UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
### GetStatus

`func (o *AccountActivitySearchedItemAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountActivitySearchedItemAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountActivitySearchedItemAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccountActivitySearchedItemAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRequester

`func (o *AccountActivitySearchedItemAllOf) GetRequester() Source1`

GetRequester returns the Requester field if non-nil, zero value otherwise.

### GetRequesterOk

`func (o *AccountActivitySearchedItemAllOf) GetRequesterOk() (*Source1, bool)`

GetRequesterOk returns a tuple with the Requester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequester

`func (o *AccountActivitySearchedItemAllOf) SetRequester(v Source1)`

SetRequester sets Requester field to given value.

### HasRequester

`func (o *AccountActivitySearchedItemAllOf) HasRequester() bool`

HasRequester returns a boolean if a field has been set.

### GetRecipient

`func (o *AccountActivitySearchedItemAllOf) GetRecipient() Source1`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *AccountActivitySearchedItemAllOf) GetRecipientOk() (*Source1, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *AccountActivitySearchedItemAllOf) SetRecipient(v Source1)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *AccountActivitySearchedItemAllOf) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetTrackingNumber

`func (o *AccountActivitySearchedItemAllOf) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *AccountActivitySearchedItemAllOf) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *AccountActivitySearchedItemAllOf) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *AccountActivitySearchedItemAllOf) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### GetErrors

`func (o *AccountActivitySearchedItemAllOf) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *AccountActivitySearchedItemAllOf) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *AccountActivitySearchedItemAllOf) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *AccountActivitySearchedItemAllOf) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *AccountActivitySearchedItemAllOf) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *AccountActivitySearchedItemAllOf) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetWarnings

`func (o *AccountActivitySearchedItemAllOf) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *AccountActivitySearchedItemAllOf) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *AccountActivitySearchedItemAllOf) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *AccountActivitySearchedItemAllOf) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarningsNil

`func (o *AccountActivitySearchedItemAllOf) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *AccountActivitySearchedItemAllOf) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil
### GetApprovals

`func (o *AccountActivitySearchedItemAllOf) GetApprovals() []Approval`

GetApprovals returns the Approvals field if non-nil, zero value otherwise.

### GetApprovalsOk

`func (o *AccountActivitySearchedItemAllOf) GetApprovalsOk() (*[]Approval, bool)`

GetApprovalsOk returns a tuple with the Approvals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovals

`func (o *AccountActivitySearchedItemAllOf) SetApprovals(v []Approval)`

SetApprovals sets Approvals field to given value.

### HasApprovals

`func (o *AccountActivitySearchedItemAllOf) HasApprovals() bool`

HasApprovals returns a boolean if a field has been set.

### GetOriginalRequests

`func (o *AccountActivitySearchedItemAllOf) GetOriginalRequests() []OriginalRequest`

GetOriginalRequests returns the OriginalRequests field if non-nil, zero value otherwise.

### GetOriginalRequestsOk

`func (o *AccountActivitySearchedItemAllOf) GetOriginalRequestsOk() (*[]OriginalRequest, bool)`

GetOriginalRequestsOk returns a tuple with the OriginalRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalRequests

`func (o *AccountActivitySearchedItemAllOf) SetOriginalRequests(v []OriginalRequest)`

SetOriginalRequests sets OriginalRequests field to given value.

### HasOriginalRequests

`func (o *AccountActivitySearchedItemAllOf) HasOriginalRequests() bool`

HasOriginalRequests returns a boolean if a field has been set.

### GetExpansionItems

`func (o *AccountActivitySearchedItemAllOf) GetExpansionItems() []ExpansionItem`

GetExpansionItems returns the ExpansionItems field if non-nil, zero value otherwise.

### GetExpansionItemsOk

`func (o *AccountActivitySearchedItemAllOf) GetExpansionItemsOk() (*[]ExpansionItem, bool)`

GetExpansionItemsOk returns a tuple with the ExpansionItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpansionItems

`func (o *AccountActivitySearchedItemAllOf) SetExpansionItems(v []ExpansionItem)`

SetExpansionItems sets ExpansionItems field to given value.

### HasExpansionItems

`func (o *AccountActivitySearchedItemAllOf) HasExpansionItems() bool`

HasExpansionItems returns a boolean if a field has been set.

### GetAccountRequests

`func (o *AccountActivitySearchedItemAllOf) GetAccountRequests() []AccountRequest`

GetAccountRequests returns the AccountRequests field if non-nil, zero value otherwise.

### GetAccountRequestsOk

`func (o *AccountActivitySearchedItemAllOf) GetAccountRequestsOk() (*[]AccountRequest, bool)`

GetAccountRequestsOk returns a tuple with the AccountRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountRequests

`func (o *AccountActivitySearchedItemAllOf) SetAccountRequests(v []AccountRequest)`

SetAccountRequests sets AccountRequests field to given value.

### HasAccountRequests

`func (o *AccountActivitySearchedItemAllOf) HasAccountRequests() bool`

HasAccountRequests returns a boolean if a field has been set.

### GetSources

`func (o *AccountActivitySearchedItemAllOf) GetSources() string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *AccountActivitySearchedItemAllOf) GetSourcesOk() (*string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *AccountActivitySearchedItemAllOf) SetSources(v string)`

SetSources sets Sources field to given value.

### HasSources

`func (o *AccountActivitySearchedItemAllOf) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


