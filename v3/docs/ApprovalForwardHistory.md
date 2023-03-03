# ApprovalForwardHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OldApproverName** | Pointer to **string** | Display name of approver that forwarded the approval. | [optional] 
**NewApproverName** | Pointer to **string** | Display name of approver to whom the approval was forwarded. | [optional] 
**Comment** | Pointer to **string** | Comment made by old approver when forwarding. | [optional] 
**Modified** | Pointer to **time.Time** | Time at which approval was forwarded. | [optional] 

## Methods

### NewApprovalForwardHistory

`func NewApprovalForwardHistory() *ApprovalForwardHistory`

NewApprovalForwardHistory instantiates a new ApprovalForwardHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalForwardHistoryWithDefaults

`func NewApprovalForwardHistoryWithDefaults() *ApprovalForwardHistory`

NewApprovalForwardHistoryWithDefaults instantiates a new ApprovalForwardHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOldApproverName

`func (o *ApprovalForwardHistory) GetOldApproverName() string`

GetOldApproverName returns the OldApproverName field if non-nil, zero value otherwise.

### GetOldApproverNameOk

`func (o *ApprovalForwardHistory) GetOldApproverNameOk() (*string, bool)`

GetOldApproverNameOk returns a tuple with the OldApproverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldApproverName

`func (o *ApprovalForwardHistory) SetOldApproverName(v string)`

SetOldApproverName sets OldApproverName field to given value.

### HasOldApproverName

`func (o *ApprovalForwardHistory) HasOldApproverName() bool`

HasOldApproverName returns a boolean if a field has been set.

### GetNewApproverName

`func (o *ApprovalForwardHistory) GetNewApproverName() string`

GetNewApproverName returns the NewApproverName field if non-nil, zero value otherwise.

### GetNewApproverNameOk

`func (o *ApprovalForwardHistory) GetNewApproverNameOk() (*string, bool)`

GetNewApproverNameOk returns a tuple with the NewApproverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewApproverName

`func (o *ApprovalForwardHistory) SetNewApproverName(v string)`

SetNewApproverName sets NewApproverName field to given value.

### HasNewApproverName

`func (o *ApprovalForwardHistory) HasNewApproverName() bool`

HasNewApproverName returns a boolean if a field has been set.

### GetComment

`func (o *ApprovalForwardHistory) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ApprovalForwardHistory) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ApprovalForwardHistory) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ApprovalForwardHistory) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetModified

`func (o *ApprovalForwardHistory) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ApprovalForwardHistory) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ApprovalForwardHistory) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *ApprovalForwardHistory) HasModified() bool`

HasModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


