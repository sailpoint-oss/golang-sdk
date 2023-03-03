# TriggerOutputAccessRequestPreApproval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Approved** | **bool** | Whether or not to approve the access request. | 
**Comment** | **string** | A comment about the decision to approve or deny the request. | 
**Approver** | **string** | The name of the entity that approved or denied the request. | 

## Methods

### NewTriggerOutputAccessRequestPreApproval

`func NewTriggerOutputAccessRequestPreApproval(approved bool, comment string, approver string, ) *TriggerOutputAccessRequestPreApproval`

NewTriggerOutputAccessRequestPreApproval instantiates a new TriggerOutputAccessRequestPreApproval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerOutputAccessRequestPreApprovalWithDefaults

`func NewTriggerOutputAccessRequestPreApprovalWithDefaults() *TriggerOutputAccessRequestPreApproval`

NewTriggerOutputAccessRequestPreApprovalWithDefaults instantiates a new TriggerOutputAccessRequestPreApproval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApproved

`func (o *TriggerOutputAccessRequestPreApproval) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *TriggerOutputAccessRequestPreApproval) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *TriggerOutputAccessRequestPreApproval) SetApproved(v bool)`

SetApproved sets Approved field to given value.


### GetComment

`func (o *TriggerOutputAccessRequestPreApproval) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TriggerOutputAccessRequestPreApproval) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TriggerOutputAccessRequestPreApproval) SetComment(v string)`

SetComment sets Comment field to given value.


### GetApprover

`func (o *TriggerOutputAccessRequestPreApproval) GetApprover() string`

GetApprover returns the Approver field if non-nil, zero value otherwise.

### GetApproverOk

`func (o *TriggerOutputAccessRequestPreApproval) GetApproverOk() (*string, bool)`

GetApproverOk returns a tuple with the Approver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprover

`func (o *TriggerOutputAccessRequestPreApproval) SetApprover(v string)`

SetApprover sets Approver field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


