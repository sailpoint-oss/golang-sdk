# NonEmployeeRequestWithoutApprovalItemAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountName** | Pointer to **string** | Requested identity account name. | [optional] 
**FirstName** | Pointer to **string** | Non-Employee&#39;s first name. | [optional] 
**LastName** | Pointer to **string** | Non-Employee&#39;s last name. | [optional] 
**Email** | Pointer to **string** | Non-Employee&#39;s email. | [optional] 
**Phone** | Pointer to **string** | Non-Employee&#39;s phone. | [optional] 
**Manager** | Pointer to **string** | The account ID of a valid identity to serve as this non-employee&#39;s manager. | [optional] 
**NonEmployeeSource** | Pointer to [**NonEmployeeSourceLiteWithSchemaAttributes**](NonEmployeeSourceLiteWithSchemaAttributes.md) |  | [optional] 
**Data** | Pointer to **map[string]string** | Attribute blob/bag for a non-employee. | [optional] 
**ApprovalStatus** | Pointer to [**ApprovalStatus**](ApprovalStatus.md) |  | [optional] 
**Comment** | Pointer to **string** | comment of requester | [optional] 
**CompletionDate** | Pointer to **time.Time** | When the request was completely approved. | [optional] 
**StartDate** | Pointer to **string** | Non-Employee employment start date. | [optional] 
**EndDate** | Pointer to **string** | Non-Employee employment end date. | [optional] 
**Modified** | Pointer to **time.Time** | When the request was last modified. | [optional] 
**Created** | Pointer to **time.Time** | When the request was created. | [optional] 

## Methods

### NewNonEmployeeRequestWithoutApprovalItemAllOf

`func NewNonEmployeeRequestWithoutApprovalItemAllOf() *NonEmployeeRequestWithoutApprovalItemAllOf`

NewNonEmployeeRequestWithoutApprovalItemAllOf instantiates a new NonEmployeeRequestWithoutApprovalItemAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonEmployeeRequestWithoutApprovalItemAllOfWithDefaults

`func NewNonEmployeeRequestWithoutApprovalItemAllOfWithDefaults() *NonEmployeeRequestWithoutApprovalItemAllOf`

NewNonEmployeeRequestWithoutApprovalItemAllOfWithDefaults instantiates a new NonEmployeeRequestWithoutApprovalItemAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountName

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetFirstName

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetManager

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) GetManager() string`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) GetManagerOk() (*string, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) SetManager(v string)`

SetManager sets Manager field to given value.

### HasManager

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetNonEmployeeSource

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) GetNonEmployeeSource() NonEmployeeSourceLiteWithSchemaAttributes`

GetNonEmployeeSource returns the NonEmployeeSource field if non-nil, zero value otherwise.

### GetNonEmployeeSourceOk

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) GetNonEmployeeSourceOk() (*NonEmployeeSourceLiteWithSchemaAttributes, bool)`

GetNonEmployeeSourceOk returns a tuple with the NonEmployeeSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonEmployeeSource

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) SetNonEmployeeSource(v NonEmployeeSourceLiteWithSchemaAttributes)`

SetNonEmployeeSource sets NonEmployeeSource field to given value.

### HasNonEmployeeSource

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) HasNonEmployeeSource() bool`

HasNonEmployeeSource returns a boolean if a field has been set.

### GetData

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) GetData() map[string]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) GetDataOk() (*map[string]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) SetData(v map[string]string)`

SetData sets Data field to given value.

### HasData

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) HasData() bool`

HasData returns a boolean if a field has been set.

### GetApprovalStatus

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) GetApprovalStatus() ApprovalStatus`

GetApprovalStatus returns the ApprovalStatus field if non-nil, zero value otherwise.

### GetApprovalStatusOk

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) GetApprovalStatusOk() (*ApprovalStatus, bool)`

GetApprovalStatusOk returns a tuple with the ApprovalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalStatus

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) SetApprovalStatus(v ApprovalStatus)`

SetApprovalStatus sets ApprovalStatus field to given value.

### HasApprovalStatus

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) HasApprovalStatus() bool`

HasApprovalStatus returns a boolean if a field has been set.

### GetComment

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCompletionDate

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) GetCompletionDate() time.Time`

GetCompletionDate returns the CompletionDate field if non-nil, zero value otherwise.

### GetCompletionDateOk

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) GetCompletionDateOk() (*time.Time, bool)`

GetCompletionDateOk returns a tuple with the CompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDate

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) SetCompletionDate(v time.Time)`

SetCompletionDate sets CompletionDate field to given value.

### HasCompletionDate

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) HasCompletionDate() bool`

HasCompletionDate returns a boolean if a field has been set.

### GetStartDate

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetModified

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetCreated

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *NonEmployeeRequestWithoutApprovalItemAllOf) HasCreated() bool`

HasCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


