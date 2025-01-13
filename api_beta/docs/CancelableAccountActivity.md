# CancelableAccountActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the account activity itself | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**Completed** | Pointer to **time.Time** |  | [optional] 
**CompletionStatus** | Pointer to [**NullableCompletionStatus**](CompletionStatus.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**RequesterIdentitySummary** | Pointer to [**NullableIdentitySummary**](IdentitySummary.md) |  | [optional] 
**TargetIdentitySummary** | Pointer to [**NullableIdentitySummary**](IdentitySummary.md) |  | [optional] 
**Errors** | Pointer to **[]string** |  | [optional] 
**Warnings** | Pointer to **[]string** |  | [optional] 
**Items** | Pointer to [**[]AccountActivityItem**](AccountActivityItem.md) |  | [optional] 
**ExecutionStatus** | Pointer to [**ExecutionStatus**](ExecutionStatus.md) |  | [optional] 
**ClientMetadata** | Pointer to **map[string]string** | Arbitrary key-value pairs, if any were included in the corresponding access request | [optional] 
**Cancelable** | Pointer to **bool** | Whether the account activity can be canceled before completion | [optional] 
**CancelComment** | Pointer to [**NullableComment**](Comment.md) |  | [optional] 

## Methods

### NewCancelableAccountActivity

`func NewCancelableAccountActivity() *CancelableAccountActivity`

NewCancelableAccountActivity instantiates a new CancelableAccountActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelableAccountActivityWithDefaults

`func NewCancelableAccountActivityWithDefaults() *CancelableAccountActivity`

NewCancelableAccountActivityWithDefaults instantiates a new CancelableAccountActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CancelableAccountActivity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CancelableAccountActivity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CancelableAccountActivity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CancelableAccountActivity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CancelableAccountActivity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CancelableAccountActivity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CancelableAccountActivity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CancelableAccountActivity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreated

`func (o *CancelableAccountActivity) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CancelableAccountActivity) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CancelableAccountActivity) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CancelableAccountActivity) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *CancelableAccountActivity) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *CancelableAccountActivity) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *CancelableAccountActivity) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *CancelableAccountActivity) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetCompleted

`func (o *CancelableAccountActivity) GetCompleted() time.Time`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *CancelableAccountActivity) GetCompletedOk() (*time.Time, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *CancelableAccountActivity) SetCompleted(v time.Time)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *CancelableAccountActivity) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetCompletionStatus

`func (o *CancelableAccountActivity) GetCompletionStatus() CompletionStatus`

GetCompletionStatus returns the CompletionStatus field if non-nil, zero value otherwise.

### GetCompletionStatusOk

`func (o *CancelableAccountActivity) GetCompletionStatusOk() (*CompletionStatus, bool)`

GetCompletionStatusOk returns a tuple with the CompletionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionStatus

`func (o *CancelableAccountActivity) SetCompletionStatus(v CompletionStatus)`

SetCompletionStatus sets CompletionStatus field to given value.

### HasCompletionStatus

`func (o *CancelableAccountActivity) HasCompletionStatus() bool`

HasCompletionStatus returns a boolean if a field has been set.

### SetCompletionStatusNil

`func (o *CancelableAccountActivity) SetCompletionStatusNil(b bool)`

 SetCompletionStatusNil sets the value for CompletionStatus to be an explicit nil

### UnsetCompletionStatus
`func (o *CancelableAccountActivity) UnsetCompletionStatus()`

UnsetCompletionStatus ensures that no value is present for CompletionStatus, not even an explicit nil
### GetType

`func (o *CancelableAccountActivity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CancelableAccountActivity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CancelableAccountActivity) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CancelableAccountActivity) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRequesterIdentitySummary

`func (o *CancelableAccountActivity) GetRequesterIdentitySummary() IdentitySummary`

GetRequesterIdentitySummary returns the RequesterIdentitySummary field if non-nil, zero value otherwise.

### GetRequesterIdentitySummaryOk

`func (o *CancelableAccountActivity) GetRequesterIdentitySummaryOk() (*IdentitySummary, bool)`

GetRequesterIdentitySummaryOk returns a tuple with the RequesterIdentitySummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterIdentitySummary

`func (o *CancelableAccountActivity) SetRequesterIdentitySummary(v IdentitySummary)`

SetRequesterIdentitySummary sets RequesterIdentitySummary field to given value.

### HasRequesterIdentitySummary

`func (o *CancelableAccountActivity) HasRequesterIdentitySummary() bool`

HasRequesterIdentitySummary returns a boolean if a field has been set.

### SetRequesterIdentitySummaryNil

`func (o *CancelableAccountActivity) SetRequesterIdentitySummaryNil(b bool)`

 SetRequesterIdentitySummaryNil sets the value for RequesterIdentitySummary to be an explicit nil

### UnsetRequesterIdentitySummary
`func (o *CancelableAccountActivity) UnsetRequesterIdentitySummary()`

UnsetRequesterIdentitySummary ensures that no value is present for RequesterIdentitySummary, not even an explicit nil
### GetTargetIdentitySummary

`func (o *CancelableAccountActivity) GetTargetIdentitySummary() IdentitySummary`

GetTargetIdentitySummary returns the TargetIdentitySummary field if non-nil, zero value otherwise.

### GetTargetIdentitySummaryOk

`func (o *CancelableAccountActivity) GetTargetIdentitySummaryOk() (*IdentitySummary, bool)`

GetTargetIdentitySummaryOk returns a tuple with the TargetIdentitySummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetIdentitySummary

`func (o *CancelableAccountActivity) SetTargetIdentitySummary(v IdentitySummary)`

SetTargetIdentitySummary sets TargetIdentitySummary field to given value.

### HasTargetIdentitySummary

`func (o *CancelableAccountActivity) HasTargetIdentitySummary() bool`

HasTargetIdentitySummary returns a boolean if a field has been set.

### SetTargetIdentitySummaryNil

`func (o *CancelableAccountActivity) SetTargetIdentitySummaryNil(b bool)`

 SetTargetIdentitySummaryNil sets the value for TargetIdentitySummary to be an explicit nil

### UnsetTargetIdentitySummary
`func (o *CancelableAccountActivity) UnsetTargetIdentitySummary()`

UnsetTargetIdentitySummary ensures that no value is present for TargetIdentitySummary, not even an explicit nil
### GetErrors

`func (o *CancelableAccountActivity) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CancelableAccountActivity) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CancelableAccountActivity) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CancelableAccountActivity) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *CancelableAccountActivity) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *CancelableAccountActivity) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *CancelableAccountActivity) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *CancelableAccountActivity) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetItems

`func (o *CancelableAccountActivity) GetItems() []AccountActivityItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CancelableAccountActivity) GetItemsOk() (*[]AccountActivityItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CancelableAccountActivity) SetItems(v []AccountActivityItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *CancelableAccountActivity) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetExecutionStatus

`func (o *CancelableAccountActivity) GetExecutionStatus() ExecutionStatus`

GetExecutionStatus returns the ExecutionStatus field if non-nil, zero value otherwise.

### GetExecutionStatusOk

`func (o *CancelableAccountActivity) GetExecutionStatusOk() (*ExecutionStatus, bool)`

GetExecutionStatusOk returns a tuple with the ExecutionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionStatus

`func (o *CancelableAccountActivity) SetExecutionStatus(v ExecutionStatus)`

SetExecutionStatus sets ExecutionStatus field to given value.

### HasExecutionStatus

`func (o *CancelableAccountActivity) HasExecutionStatus() bool`

HasExecutionStatus returns a boolean if a field has been set.

### GetClientMetadata

`func (o *CancelableAccountActivity) GetClientMetadata() map[string]string`

GetClientMetadata returns the ClientMetadata field if non-nil, zero value otherwise.

### GetClientMetadataOk

`func (o *CancelableAccountActivity) GetClientMetadataOk() (*map[string]string, bool)`

GetClientMetadataOk returns a tuple with the ClientMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMetadata

`func (o *CancelableAccountActivity) SetClientMetadata(v map[string]string)`

SetClientMetadata sets ClientMetadata field to given value.

### HasClientMetadata

`func (o *CancelableAccountActivity) HasClientMetadata() bool`

HasClientMetadata returns a boolean if a field has been set.

### GetCancelable

`func (o *CancelableAccountActivity) GetCancelable() bool`

GetCancelable returns the Cancelable field if non-nil, zero value otherwise.

### GetCancelableOk

`func (o *CancelableAccountActivity) GetCancelableOk() (*bool, bool)`

GetCancelableOk returns a tuple with the Cancelable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelable

`func (o *CancelableAccountActivity) SetCancelable(v bool)`

SetCancelable sets Cancelable field to given value.

### HasCancelable

`func (o *CancelableAccountActivity) HasCancelable() bool`

HasCancelable returns a boolean if a field has been set.

### GetCancelComment

`func (o *CancelableAccountActivity) GetCancelComment() Comment`

GetCancelComment returns the CancelComment field if non-nil, zero value otherwise.

### GetCancelCommentOk

`func (o *CancelableAccountActivity) GetCancelCommentOk() (*Comment, bool)`

GetCancelCommentOk returns a tuple with the CancelComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelComment

`func (o *CancelableAccountActivity) SetCancelComment(v Comment)`

SetCancelComment sets CancelComment field to given value.

### HasCancelComment

`func (o *CancelableAccountActivity) HasCancelComment() bool`

HasCancelComment returns a boolean if a field has been set.

### SetCancelCommentNil

`func (o *CancelableAccountActivity) SetCancelCommentNil(b bool)`

 SetCancelCommentNil sets the value for CancelComment to be an explicit nil

### UnsetCancelComment
`func (o *CancelableAccountActivity) UnsetCancelComment()`

UnsetCancelComment ensures that no value is present for CancelComment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


