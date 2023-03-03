# TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApprovalComment** | Pointer to **NullableString** | A comment left by the approver. | [optional] 
**ApprovalDecision** | **map[string]interface{}** | The final decision of the approver. | 
**ApproverName** | **string** | The name of the approver | 
**Approver** | [**TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover**](TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover.md) |  | 

## Methods

### NewTriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner

`func NewTriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner(approvalDecision map[string]interface{}, approverName string, approver TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover, ) *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner`

NewTriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner instantiates a new TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerWithDefaults

`func NewTriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerWithDefaults() *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner`

NewTriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerWithDefaults instantiates a new TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovalComment

`func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) GetApprovalComment() string`

GetApprovalComment returns the ApprovalComment field if non-nil, zero value otherwise.

### GetApprovalCommentOk

`func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) GetApprovalCommentOk() (*string, bool)`

GetApprovalCommentOk returns a tuple with the ApprovalComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalComment

`func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) SetApprovalComment(v string)`

SetApprovalComment sets ApprovalComment field to given value.

### HasApprovalComment

`func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) HasApprovalComment() bool`

HasApprovalComment returns a boolean if a field has been set.

### SetApprovalCommentNil

`func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) SetApprovalCommentNil(b bool)`

 SetApprovalCommentNil sets the value for ApprovalComment to be an explicit nil

### UnsetApprovalComment
`func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) UnsetApprovalComment()`

UnsetApprovalComment ensures that no value is present for ApprovalComment, not even an explicit nil
### GetApprovalDecision

`func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) GetApprovalDecision() map[string]interface{}`

GetApprovalDecision returns the ApprovalDecision field if non-nil, zero value otherwise.

### GetApprovalDecisionOk

`func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) GetApprovalDecisionOk() (*map[string]interface{}, bool)`

GetApprovalDecisionOk returns a tuple with the ApprovalDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalDecision

`func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) SetApprovalDecision(v map[string]interface{})`

SetApprovalDecision sets ApprovalDecision field to given value.


### GetApproverName

`func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) GetApproverName() string`

GetApproverName returns the ApproverName field if non-nil, zero value otherwise.

### GetApproverNameOk

`func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) GetApproverNameOk() (*string, bool)`

GetApproverNameOk returns a tuple with the ApproverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverName

`func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) SetApproverName(v string)`

SetApproverName sets ApproverName field to given value.


### GetApprover

`func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) GetApprover() TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover`

GetApprover returns the Approver field if non-nil, zero value otherwise.

### GetApproverOk

`func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) GetApproverOk() (*TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover, bool)`

GetApproverOk returns a tuple with the Approver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprover

`func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) SetApprover(v TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInnerApprover)`

SetApprover sets Approver field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


