# NonEmployeeRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountName** | Pointer to **string** | Requested identity account name. | [optional] 
**FirstName** | Pointer to **string** | Non-Employee&#39;s first name. | [optional] 
**LastName** | Pointer to **string** | Non-Employee&#39;s last name. | [optional] 
**Email** | Pointer to **string** | Non-Employee&#39;s email. | [optional] 
**Phone** | Pointer to **string** | Non-Employee&#39;s phone. | [optional] 
**Manager** | Pointer to **string** | The account ID of a valid identity to serve as this non-employee&#39;s manager. | [optional] 
**NonEmployeeSource** | Pointer to [**NonEmployeeSourceLite**](NonEmployeeSourceLite.md) |  | [optional] 
**Data** | Pointer to **map[string]string** | Attribute blob/bag for a non-employee. | [optional] 
**ApprovalItems** | Pointer to [**[]NonEmployeeApprovalItemBase**](NonEmployeeApprovalItemBase.md) | List of approval item for the request | [optional] 
**ApprovalStatus** | Pointer to [**ApprovalStatus**](ApprovalStatus.md) |  | [optional] 
**Comment** | Pointer to **string** | comment of requester | [optional] 
**CompletionDate** | Pointer to **time.Time** | When the request was completely approved. | [optional] 
**StartDate** | Pointer to **time.Time** | Non-Employee employment start date. | [optional] 
**EndDate** | Pointer to **time.Time** | Non-Employee employment end date. | [optional] 
**Modified** | Pointer to **time.Time** | When the request was last modified. | [optional] 
**Created** | Pointer to **time.Time** | When the request was created. | [optional] 

## Methods

### NewNonEmployeeRequestAllOf

`func NewNonEmployeeRequestAllOf() *NonEmployeeRequestAllOf`

NewNonEmployeeRequestAllOf instantiates a new NonEmployeeRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonEmployeeRequestAllOfWithDefaults

`func NewNonEmployeeRequestAllOfWithDefaults() *NonEmployeeRequestAllOf`

NewNonEmployeeRequestAllOfWithDefaults instantiates a new NonEmployeeRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountName

`func (o *NonEmployeeRequestAllOf) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *NonEmployeeRequestAllOf) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *NonEmployeeRequestAllOf) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *NonEmployeeRequestAllOf) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetFirstName

`func (o *NonEmployeeRequestAllOf) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *NonEmployeeRequestAllOf) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *NonEmployeeRequestAllOf) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *NonEmployeeRequestAllOf) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *NonEmployeeRequestAllOf) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *NonEmployeeRequestAllOf) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *NonEmployeeRequestAllOf) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *NonEmployeeRequestAllOf) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *NonEmployeeRequestAllOf) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *NonEmployeeRequestAllOf) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *NonEmployeeRequestAllOf) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *NonEmployeeRequestAllOf) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *NonEmployeeRequestAllOf) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *NonEmployeeRequestAllOf) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *NonEmployeeRequestAllOf) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *NonEmployeeRequestAllOf) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetManager

`func (o *NonEmployeeRequestAllOf) GetManager() string`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *NonEmployeeRequestAllOf) GetManagerOk() (*string, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *NonEmployeeRequestAllOf) SetManager(v string)`

SetManager sets Manager field to given value.

### HasManager

`func (o *NonEmployeeRequestAllOf) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetNonEmployeeSource

`func (o *NonEmployeeRequestAllOf) GetNonEmployeeSource() NonEmployeeSourceLite`

GetNonEmployeeSource returns the NonEmployeeSource field if non-nil, zero value otherwise.

### GetNonEmployeeSourceOk

`func (o *NonEmployeeRequestAllOf) GetNonEmployeeSourceOk() (*NonEmployeeSourceLite, bool)`

GetNonEmployeeSourceOk returns a tuple with the NonEmployeeSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonEmployeeSource

`func (o *NonEmployeeRequestAllOf) SetNonEmployeeSource(v NonEmployeeSourceLite)`

SetNonEmployeeSource sets NonEmployeeSource field to given value.

### HasNonEmployeeSource

`func (o *NonEmployeeRequestAllOf) HasNonEmployeeSource() bool`

HasNonEmployeeSource returns a boolean if a field has been set.

### GetData

`func (o *NonEmployeeRequestAllOf) GetData() map[string]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NonEmployeeRequestAllOf) GetDataOk() (*map[string]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NonEmployeeRequestAllOf) SetData(v map[string]string)`

SetData sets Data field to given value.

### HasData

`func (o *NonEmployeeRequestAllOf) HasData() bool`

HasData returns a boolean if a field has been set.

### GetApprovalItems

`func (o *NonEmployeeRequestAllOf) GetApprovalItems() []NonEmployeeApprovalItemBase`

GetApprovalItems returns the ApprovalItems field if non-nil, zero value otherwise.

### GetApprovalItemsOk

`func (o *NonEmployeeRequestAllOf) GetApprovalItemsOk() (*[]NonEmployeeApprovalItemBase, bool)`

GetApprovalItemsOk returns a tuple with the ApprovalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalItems

`func (o *NonEmployeeRequestAllOf) SetApprovalItems(v []NonEmployeeApprovalItemBase)`

SetApprovalItems sets ApprovalItems field to given value.

### HasApprovalItems

`func (o *NonEmployeeRequestAllOf) HasApprovalItems() bool`

HasApprovalItems returns a boolean if a field has been set.

### GetApprovalStatus

`func (o *NonEmployeeRequestAllOf) GetApprovalStatus() ApprovalStatus`

GetApprovalStatus returns the ApprovalStatus field if non-nil, zero value otherwise.

### GetApprovalStatusOk

`func (o *NonEmployeeRequestAllOf) GetApprovalStatusOk() (*ApprovalStatus, bool)`

GetApprovalStatusOk returns a tuple with the ApprovalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalStatus

`func (o *NonEmployeeRequestAllOf) SetApprovalStatus(v ApprovalStatus)`

SetApprovalStatus sets ApprovalStatus field to given value.

### HasApprovalStatus

`func (o *NonEmployeeRequestAllOf) HasApprovalStatus() bool`

HasApprovalStatus returns a boolean if a field has been set.

### GetComment

`func (o *NonEmployeeRequestAllOf) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *NonEmployeeRequestAllOf) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *NonEmployeeRequestAllOf) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *NonEmployeeRequestAllOf) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCompletionDate

`func (o *NonEmployeeRequestAllOf) GetCompletionDate() time.Time`

GetCompletionDate returns the CompletionDate field if non-nil, zero value otherwise.

### GetCompletionDateOk

`func (o *NonEmployeeRequestAllOf) GetCompletionDateOk() (*time.Time, bool)`

GetCompletionDateOk returns a tuple with the CompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDate

`func (o *NonEmployeeRequestAllOf) SetCompletionDate(v time.Time)`

SetCompletionDate sets CompletionDate field to given value.

### HasCompletionDate

`func (o *NonEmployeeRequestAllOf) HasCompletionDate() bool`

HasCompletionDate returns a boolean if a field has been set.

### GetStartDate

`func (o *NonEmployeeRequestAllOf) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *NonEmployeeRequestAllOf) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *NonEmployeeRequestAllOf) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *NonEmployeeRequestAllOf) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *NonEmployeeRequestAllOf) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *NonEmployeeRequestAllOf) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *NonEmployeeRequestAllOf) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *NonEmployeeRequestAllOf) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetModified

`func (o *NonEmployeeRequestAllOf) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *NonEmployeeRequestAllOf) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *NonEmployeeRequestAllOf) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *NonEmployeeRequestAllOf) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetCreated

`func (o *NonEmployeeRequestAllOf) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *NonEmployeeRequestAllOf) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *NonEmployeeRequestAllOf) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *NonEmployeeRequestAllOf) HasCreated() bool`

HasCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


