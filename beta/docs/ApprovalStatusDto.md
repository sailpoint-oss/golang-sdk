# ApprovalStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Forwarded** | Pointer to **bool** | True if the request for this item was forwarded from one owner to another. | [optional] 
**OriginalOwner** | Pointer to [**BaseReferenceDto**](BaseReferenceDto.md) |  | [optional] 
**CurrentOwner** | Pointer to [**BaseReferenceDto**](BaseReferenceDto.md) |  | [optional] 
**ReviewedBy** | Pointer to [**BaseReferenceDto**](BaseReferenceDto.md) |  | [optional] 
**Modified** | Pointer to **time.Time** | Time at which item was modified. | [optional] 
**Status** | Pointer to [**ManualWorkItemState**](ManualWorkItemState.md) |  | [optional] 
**Scheme** | Pointer to [**ApprovalScheme**](ApprovalScheme.md) |  | [optional] 
**ErrorMessages** | Pointer to [**[]ErrorMessageDto**](ErrorMessageDto.md) | If the request failed, includes any error messages that were generated. | [optional] 
**Comment** | Pointer to **string** | Comment, if any, provided by the approver. | [optional] 
**RemoveDate** | Pointer to **time.Time** | The date the role or access profile is no longer assigned to the specified identity. | [optional] 

## Methods

### NewApprovalStatusDto

`func NewApprovalStatusDto() *ApprovalStatusDto`

NewApprovalStatusDto instantiates a new ApprovalStatusDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalStatusDtoWithDefaults

`func NewApprovalStatusDtoWithDefaults() *ApprovalStatusDto`

NewApprovalStatusDtoWithDefaults instantiates a new ApprovalStatusDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForwarded

`func (o *ApprovalStatusDto) GetForwarded() bool`

GetForwarded returns the Forwarded field if non-nil, zero value otherwise.

### GetForwardedOk

`func (o *ApprovalStatusDto) GetForwardedOk() (*bool, bool)`

GetForwardedOk returns a tuple with the Forwarded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwarded

`func (o *ApprovalStatusDto) SetForwarded(v bool)`

SetForwarded sets Forwarded field to given value.

### HasForwarded

`func (o *ApprovalStatusDto) HasForwarded() bool`

HasForwarded returns a boolean if a field has been set.

### GetOriginalOwner

`func (o *ApprovalStatusDto) GetOriginalOwner() BaseReferenceDto`

GetOriginalOwner returns the OriginalOwner field if non-nil, zero value otherwise.

### GetOriginalOwnerOk

`func (o *ApprovalStatusDto) GetOriginalOwnerOk() (*BaseReferenceDto, bool)`

GetOriginalOwnerOk returns a tuple with the OriginalOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalOwner

`func (o *ApprovalStatusDto) SetOriginalOwner(v BaseReferenceDto)`

SetOriginalOwner sets OriginalOwner field to given value.

### HasOriginalOwner

`func (o *ApprovalStatusDto) HasOriginalOwner() bool`

HasOriginalOwner returns a boolean if a field has been set.

### GetCurrentOwner

`func (o *ApprovalStatusDto) GetCurrentOwner() BaseReferenceDto`

GetCurrentOwner returns the CurrentOwner field if non-nil, zero value otherwise.

### GetCurrentOwnerOk

`func (o *ApprovalStatusDto) GetCurrentOwnerOk() (*BaseReferenceDto, bool)`

GetCurrentOwnerOk returns a tuple with the CurrentOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentOwner

`func (o *ApprovalStatusDto) SetCurrentOwner(v BaseReferenceDto)`

SetCurrentOwner sets CurrentOwner field to given value.

### HasCurrentOwner

`func (o *ApprovalStatusDto) HasCurrentOwner() bool`

HasCurrentOwner returns a boolean if a field has been set.

### GetReviewedBy

`func (o *ApprovalStatusDto) GetReviewedBy() BaseReferenceDto`

GetReviewedBy returns the ReviewedBy field if non-nil, zero value otherwise.

### GetReviewedByOk

`func (o *ApprovalStatusDto) GetReviewedByOk() (*BaseReferenceDto, bool)`

GetReviewedByOk returns a tuple with the ReviewedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedBy

`func (o *ApprovalStatusDto) SetReviewedBy(v BaseReferenceDto)`

SetReviewedBy sets ReviewedBy field to given value.

### HasReviewedBy

`func (o *ApprovalStatusDto) HasReviewedBy() bool`

HasReviewedBy returns a boolean if a field has been set.

### GetModified

`func (o *ApprovalStatusDto) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ApprovalStatusDto) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ApprovalStatusDto) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *ApprovalStatusDto) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetStatus

`func (o *ApprovalStatusDto) GetStatus() ManualWorkItemState`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApprovalStatusDto) GetStatusOk() (*ManualWorkItemState, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApprovalStatusDto) SetStatus(v ManualWorkItemState)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApprovalStatusDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetScheme

`func (o *ApprovalStatusDto) GetScheme() ApprovalScheme`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *ApprovalStatusDto) GetSchemeOk() (*ApprovalScheme, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *ApprovalStatusDto) SetScheme(v ApprovalScheme)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *ApprovalStatusDto) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetErrorMessages

`func (o *ApprovalStatusDto) GetErrorMessages() []ErrorMessageDto`

GetErrorMessages returns the ErrorMessages field if non-nil, zero value otherwise.

### GetErrorMessagesOk

`func (o *ApprovalStatusDto) GetErrorMessagesOk() (*[]ErrorMessageDto, bool)`

GetErrorMessagesOk returns a tuple with the ErrorMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessages

`func (o *ApprovalStatusDto) SetErrorMessages(v []ErrorMessageDto)`

SetErrorMessages sets ErrorMessages field to given value.

### HasErrorMessages

`func (o *ApprovalStatusDto) HasErrorMessages() bool`

HasErrorMessages returns a boolean if a field has been set.

### GetComment

`func (o *ApprovalStatusDto) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ApprovalStatusDto) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ApprovalStatusDto) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ApprovalStatusDto) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetRemoveDate

`func (o *ApprovalStatusDto) GetRemoveDate() time.Time`

GetRemoveDate returns the RemoveDate field if non-nil, zero value otherwise.

### GetRemoveDateOk

`func (o *ApprovalStatusDto) GetRemoveDateOk() (*time.Time, bool)`

GetRemoveDateOk returns a tuple with the RemoveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveDate

`func (o *ApprovalStatusDto) SetRemoveDate(v time.Time)`

SetRemoveDate sets RemoveDate field to given value.

### HasRemoveDate

`func (o *ApprovalStatusDto) HasRemoveDate() bool`

HasRemoveDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


