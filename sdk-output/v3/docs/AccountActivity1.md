# AccountActivity1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Type** | [**DocumentType**](DocumentType.md) |  | 
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

### NewAccountActivity1

`func NewAccountActivity1(id string, name string, type_ DocumentType, ) *AccountActivity1`

NewAccountActivity1 instantiates a new AccountActivity1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountActivity1WithDefaults

`func NewAccountActivity1WithDefaults() *AccountActivity1`

NewAccountActivity1WithDefaults instantiates a new AccountActivity1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountActivity1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountActivity1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountActivity1) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AccountActivity1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountActivity1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountActivity1) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *AccountActivity1) GetType() DocumentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountActivity1) GetTypeOk() (*DocumentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountActivity1) SetType(v DocumentType)`

SetType sets Type field to given value.


### GetAction

`func (o *AccountActivity1) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AccountActivity1) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AccountActivity1) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *AccountActivity1) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCreated

`func (o *AccountActivity1) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AccountActivity1) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AccountActivity1) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AccountActivity1) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *AccountActivity1) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *AccountActivity1) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetModified

`func (o *AccountActivity1) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *AccountActivity1) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *AccountActivity1) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *AccountActivity1) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModifiedNil

`func (o *AccountActivity1) SetModifiedNil(b bool)`

 SetModifiedNil sets the value for Modified to be an explicit nil

### UnsetModified
`func (o *AccountActivity1) UnsetModified()`

UnsetModified ensures that no value is present for Modified, not even an explicit nil
### GetStage

`func (o *AccountActivity1) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *AccountActivity1) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *AccountActivity1) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *AccountActivity1) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetOrigin

`func (o *AccountActivity1) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *AccountActivity1) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *AccountActivity1) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *AccountActivity1) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### SetOriginNil

`func (o *AccountActivity1) SetOriginNil(b bool)`

 SetOriginNil sets the value for Origin to be an explicit nil

### UnsetOrigin
`func (o *AccountActivity1) UnsetOrigin()`

UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
### GetStatus

`func (o *AccountActivity1) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountActivity1) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountActivity1) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccountActivity1) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRequester

`func (o *AccountActivity1) GetRequester() Source1`

GetRequester returns the Requester field if non-nil, zero value otherwise.

### GetRequesterOk

`func (o *AccountActivity1) GetRequesterOk() (*Source1, bool)`

GetRequesterOk returns a tuple with the Requester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequester

`func (o *AccountActivity1) SetRequester(v Source1)`

SetRequester sets Requester field to given value.

### HasRequester

`func (o *AccountActivity1) HasRequester() bool`

HasRequester returns a boolean if a field has been set.

### GetRecipient

`func (o *AccountActivity1) GetRecipient() Source1`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *AccountActivity1) GetRecipientOk() (*Source1, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *AccountActivity1) SetRecipient(v Source1)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *AccountActivity1) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetTrackingNumber

`func (o *AccountActivity1) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *AccountActivity1) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *AccountActivity1) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *AccountActivity1) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### GetErrors

`func (o *AccountActivity1) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *AccountActivity1) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *AccountActivity1) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *AccountActivity1) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *AccountActivity1) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *AccountActivity1) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetWarnings

`func (o *AccountActivity1) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *AccountActivity1) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *AccountActivity1) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *AccountActivity1) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarningsNil

`func (o *AccountActivity1) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *AccountActivity1) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil
### GetApprovals

`func (o *AccountActivity1) GetApprovals() []Approval`

GetApprovals returns the Approvals field if non-nil, zero value otherwise.

### GetApprovalsOk

`func (o *AccountActivity1) GetApprovalsOk() (*[]Approval, bool)`

GetApprovalsOk returns a tuple with the Approvals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovals

`func (o *AccountActivity1) SetApprovals(v []Approval)`

SetApprovals sets Approvals field to given value.

### HasApprovals

`func (o *AccountActivity1) HasApprovals() bool`

HasApprovals returns a boolean if a field has been set.

### GetOriginalRequests

`func (o *AccountActivity1) GetOriginalRequests() []OriginalRequest`

GetOriginalRequests returns the OriginalRequests field if non-nil, zero value otherwise.

### GetOriginalRequestsOk

`func (o *AccountActivity1) GetOriginalRequestsOk() (*[]OriginalRequest, bool)`

GetOriginalRequestsOk returns a tuple with the OriginalRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalRequests

`func (o *AccountActivity1) SetOriginalRequests(v []OriginalRequest)`

SetOriginalRequests sets OriginalRequests field to given value.

### HasOriginalRequests

`func (o *AccountActivity1) HasOriginalRequests() bool`

HasOriginalRequests returns a boolean if a field has been set.

### GetExpansionItems

`func (o *AccountActivity1) GetExpansionItems() []ExpansionItem`

GetExpansionItems returns the ExpansionItems field if non-nil, zero value otherwise.

### GetExpansionItemsOk

`func (o *AccountActivity1) GetExpansionItemsOk() (*[]ExpansionItem, bool)`

GetExpansionItemsOk returns a tuple with the ExpansionItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpansionItems

`func (o *AccountActivity1) SetExpansionItems(v []ExpansionItem)`

SetExpansionItems sets ExpansionItems field to given value.

### HasExpansionItems

`func (o *AccountActivity1) HasExpansionItems() bool`

HasExpansionItems returns a boolean if a field has been set.

### GetAccountRequests

`func (o *AccountActivity1) GetAccountRequests() []AccountRequest`

GetAccountRequests returns the AccountRequests field if non-nil, zero value otherwise.

### GetAccountRequestsOk

`func (o *AccountActivity1) GetAccountRequestsOk() (*[]AccountRequest, bool)`

GetAccountRequestsOk returns a tuple with the AccountRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountRequests

`func (o *AccountActivity1) SetAccountRequests(v []AccountRequest)`

SetAccountRequests sets AccountRequests field to given value.

### HasAccountRequests

`func (o *AccountActivity1) HasAccountRequests() bool`

HasAccountRequests returns a boolean if a field has been set.

### GetSources

`func (o *AccountActivity1) GetSources() string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *AccountActivity1) GetSourcesOk() (*string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *AccountActivity1) SetSources(v string)`

SetSources sets Sources field to given value.

### HasSources

`func (o *AccountActivity1) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


