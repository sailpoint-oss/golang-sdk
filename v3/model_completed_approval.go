/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
	"time"
)

// CompletedApproval struct for CompletedApproval
type CompletedApproval struct {
	// The approval id.
	Id *string `json:"id,omitempty"`
	// The name of the approval.
	Name *string `json:"name,omitempty"`
	// When the approval was created.
	Created *time.Time `json:"created,omitempty"`
	// When the approval was modified last time.
	Modified *time.Time `json:"modified,omitempty"`
	// When the access-request was created.
	RequestCreated *time.Time `json:"requestCreated,omitempty"`
	RequestType *AccessRequestType `json:"requestType,omitempty"`
	Requester *BaseReferenceDto `json:"requester,omitempty"`
	RequestedFor *BaseReferenceDto `json:"requestedFor,omitempty"`
	ReviewedBy *BaseReferenceDto `json:"reviewedBy,omitempty"`
	Owner *BaseReferenceDto `json:"owner,omitempty"`
	RequestedObject *RequestableObjectReference `json:"requestedObject,omitempty"`
	RequesterComment *CommentDto `json:"requesterComment,omitempty"`
	ReviewerComment *CommentDto `json:"reviewerComment,omitempty"`
	// The history of the previous reviewers comments.
	PreviousReviewersComments []CommentDto `json:"previousReviewersComments,omitempty"`
	// The history of approval forward action.
	ForwardHistory []ApprovalForwardHistory `json:"forwardHistory,omitempty"`
	// When true the rejector has to provide comments when rejecting
	CommentRequiredWhenRejected *bool `json:"commentRequiredWhenRejected,omitempty"`
	State *CompletedApprovalState `json:"state,omitempty"`
	// The date the role or access profile is no longer assigned to the specified identity.
	RemoveDate *time.Time `json:"removeDate,omitempty"`
	// If true, then the request was to change the remove date or sunset date.
	RemoveDateUpdateRequested *bool `json:"removeDateUpdateRequested,omitempty"`
	// The remove date or sunset date that was assigned at the time of the request.
	CurrentRemoveDate *time.Time `json:"currentRemoveDate,omitempty"`
	SodViolationContext *SodViolationContextCheckCompleted `json:"sodViolationContext,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CompletedApproval CompletedApproval

// NewCompletedApproval instantiates a new CompletedApproval object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompletedApproval() *CompletedApproval {
	this := CompletedApproval{}
	return &this
}

// NewCompletedApprovalWithDefaults instantiates a new CompletedApproval object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompletedApprovalWithDefaults() *CompletedApproval {
	this := CompletedApproval{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CompletedApproval) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletedApproval) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CompletedApproval) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CompletedApproval) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CompletedApproval) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletedApproval) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CompletedApproval) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CompletedApproval) SetName(v string) {
	o.Name = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *CompletedApproval) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletedApproval) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *CompletedApproval) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *CompletedApproval) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *CompletedApproval) GetModified() time.Time {
	if o == nil || isNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletedApproval) GetModifiedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *CompletedApproval) HasModified() bool {
	if o != nil && !isNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *CompletedApproval) SetModified(v time.Time) {
	o.Modified = &v
}

// GetRequestCreated returns the RequestCreated field value if set, zero value otherwise.
func (o *CompletedApproval) GetRequestCreated() time.Time {
	if o == nil || isNil(o.RequestCreated) {
		var ret time.Time
		return ret
	}
	return *o.RequestCreated
}

// GetRequestCreatedOk returns a tuple with the RequestCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletedApproval) GetRequestCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.RequestCreated) {
		return nil, false
	}
	return o.RequestCreated, true
}

// HasRequestCreated returns a boolean if a field has been set.
func (o *CompletedApproval) HasRequestCreated() bool {
	if o != nil && !isNil(o.RequestCreated) {
		return true
	}

	return false
}

// SetRequestCreated gets a reference to the given time.Time and assigns it to the RequestCreated field.
func (o *CompletedApproval) SetRequestCreated(v time.Time) {
	o.RequestCreated = &v
}

// GetRequestType returns the RequestType field value if set, zero value otherwise.
func (o *CompletedApproval) GetRequestType() AccessRequestType {
	if o == nil || isNil(o.RequestType) {
		var ret AccessRequestType
		return ret
	}
	return *o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletedApproval) GetRequestTypeOk() (*AccessRequestType, bool) {
	if o == nil || isNil(o.RequestType) {
		return nil, false
	}
	return o.RequestType, true
}

// HasRequestType returns a boolean if a field has been set.
func (o *CompletedApproval) HasRequestType() bool {
	if o != nil && !isNil(o.RequestType) {
		return true
	}

	return false
}

// SetRequestType gets a reference to the given AccessRequestType and assigns it to the RequestType field.
func (o *CompletedApproval) SetRequestType(v AccessRequestType) {
	o.RequestType = &v
}

// GetRequester returns the Requester field value if set, zero value otherwise.
func (o *CompletedApproval) GetRequester() BaseReferenceDto {
	if o == nil || isNil(o.Requester) {
		var ret BaseReferenceDto
		return ret
	}
	return *o.Requester
}

// GetRequesterOk returns a tuple with the Requester field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletedApproval) GetRequesterOk() (*BaseReferenceDto, bool) {
	if o == nil || isNil(o.Requester) {
		return nil, false
	}
	return o.Requester, true
}

// HasRequester returns a boolean if a field has been set.
func (o *CompletedApproval) HasRequester() bool {
	if o != nil && !isNil(o.Requester) {
		return true
	}

	return false
}

// SetRequester gets a reference to the given BaseReferenceDto and assigns it to the Requester field.
func (o *CompletedApproval) SetRequester(v BaseReferenceDto) {
	o.Requester = &v
}

// GetRequestedFor returns the RequestedFor field value if set, zero value otherwise.
func (o *CompletedApproval) GetRequestedFor() BaseReferenceDto {
	if o == nil || isNil(o.RequestedFor) {
		var ret BaseReferenceDto
		return ret
	}
	return *o.RequestedFor
}

// GetRequestedForOk returns a tuple with the RequestedFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletedApproval) GetRequestedForOk() (*BaseReferenceDto, bool) {
	if o == nil || isNil(o.RequestedFor) {
		return nil, false
	}
	return o.RequestedFor, true
}

// HasRequestedFor returns a boolean if a field has been set.
func (o *CompletedApproval) HasRequestedFor() bool {
	if o != nil && !isNil(o.RequestedFor) {
		return true
	}

	return false
}

// SetRequestedFor gets a reference to the given BaseReferenceDto and assigns it to the RequestedFor field.
func (o *CompletedApproval) SetRequestedFor(v BaseReferenceDto) {
	o.RequestedFor = &v
}

// GetReviewedBy returns the ReviewedBy field value if set, zero value otherwise.
func (o *CompletedApproval) GetReviewedBy() BaseReferenceDto {
	if o == nil || isNil(o.ReviewedBy) {
		var ret BaseReferenceDto
		return ret
	}
	return *o.ReviewedBy
}

// GetReviewedByOk returns a tuple with the ReviewedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletedApproval) GetReviewedByOk() (*BaseReferenceDto, bool) {
	if o == nil || isNil(o.ReviewedBy) {
		return nil, false
	}
	return o.ReviewedBy, true
}

// HasReviewedBy returns a boolean if a field has been set.
func (o *CompletedApproval) HasReviewedBy() bool {
	if o != nil && !isNil(o.ReviewedBy) {
		return true
	}

	return false
}

// SetReviewedBy gets a reference to the given BaseReferenceDto and assigns it to the ReviewedBy field.
func (o *CompletedApproval) SetReviewedBy(v BaseReferenceDto) {
	o.ReviewedBy = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *CompletedApproval) GetOwner() BaseReferenceDto {
	if o == nil || isNil(o.Owner) {
		var ret BaseReferenceDto
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletedApproval) GetOwnerOk() (*BaseReferenceDto, bool) {
	if o == nil || isNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *CompletedApproval) HasOwner() bool {
	if o != nil && !isNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given BaseReferenceDto and assigns it to the Owner field.
func (o *CompletedApproval) SetOwner(v BaseReferenceDto) {
	o.Owner = &v
}

// GetRequestedObject returns the RequestedObject field value if set, zero value otherwise.
func (o *CompletedApproval) GetRequestedObject() RequestableObjectReference {
	if o == nil || isNil(o.RequestedObject) {
		var ret RequestableObjectReference
		return ret
	}
	return *o.RequestedObject
}

// GetRequestedObjectOk returns a tuple with the RequestedObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletedApproval) GetRequestedObjectOk() (*RequestableObjectReference, bool) {
	if o == nil || isNil(o.RequestedObject) {
		return nil, false
	}
	return o.RequestedObject, true
}

// HasRequestedObject returns a boolean if a field has been set.
func (o *CompletedApproval) HasRequestedObject() bool {
	if o != nil && !isNil(o.RequestedObject) {
		return true
	}

	return false
}

// SetRequestedObject gets a reference to the given RequestableObjectReference and assigns it to the RequestedObject field.
func (o *CompletedApproval) SetRequestedObject(v RequestableObjectReference) {
	o.RequestedObject = &v
}

// GetRequesterComment returns the RequesterComment field value if set, zero value otherwise.
func (o *CompletedApproval) GetRequesterComment() CommentDto {
	if o == nil || isNil(o.RequesterComment) {
		var ret CommentDto
		return ret
	}
	return *o.RequesterComment
}

// GetRequesterCommentOk returns a tuple with the RequesterComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletedApproval) GetRequesterCommentOk() (*CommentDto, bool) {
	if o == nil || isNil(o.RequesterComment) {
		return nil, false
	}
	return o.RequesterComment, true
}

// HasRequesterComment returns a boolean if a field has been set.
func (o *CompletedApproval) HasRequesterComment() bool {
	if o != nil && !isNil(o.RequesterComment) {
		return true
	}

	return false
}

// SetRequesterComment gets a reference to the given CommentDto and assigns it to the RequesterComment field.
func (o *CompletedApproval) SetRequesterComment(v CommentDto) {
	o.RequesterComment = &v
}

// GetReviewerComment returns the ReviewerComment field value if set, zero value otherwise.
func (o *CompletedApproval) GetReviewerComment() CommentDto {
	if o == nil || isNil(o.ReviewerComment) {
		var ret CommentDto
		return ret
	}
	return *o.ReviewerComment
}

// GetReviewerCommentOk returns a tuple with the ReviewerComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletedApproval) GetReviewerCommentOk() (*CommentDto, bool) {
	if o == nil || isNil(o.ReviewerComment) {
		return nil, false
	}
	return o.ReviewerComment, true
}

// HasReviewerComment returns a boolean if a field has been set.
func (o *CompletedApproval) HasReviewerComment() bool {
	if o != nil && !isNil(o.ReviewerComment) {
		return true
	}

	return false
}

// SetReviewerComment gets a reference to the given CommentDto and assigns it to the ReviewerComment field.
func (o *CompletedApproval) SetReviewerComment(v CommentDto) {
	o.ReviewerComment = &v
}

// GetPreviousReviewersComments returns the PreviousReviewersComments field value if set, zero value otherwise.
func (o *CompletedApproval) GetPreviousReviewersComments() []CommentDto {
	if o == nil || isNil(o.PreviousReviewersComments) {
		var ret []CommentDto
		return ret
	}
	return o.PreviousReviewersComments
}

// GetPreviousReviewersCommentsOk returns a tuple with the PreviousReviewersComments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletedApproval) GetPreviousReviewersCommentsOk() ([]CommentDto, bool) {
	if o == nil || isNil(o.PreviousReviewersComments) {
		return nil, false
	}
	return o.PreviousReviewersComments, true
}

// HasPreviousReviewersComments returns a boolean if a field has been set.
func (o *CompletedApproval) HasPreviousReviewersComments() bool {
	if o != nil && !isNil(o.PreviousReviewersComments) {
		return true
	}

	return false
}

// SetPreviousReviewersComments gets a reference to the given []CommentDto and assigns it to the PreviousReviewersComments field.
func (o *CompletedApproval) SetPreviousReviewersComments(v []CommentDto) {
	o.PreviousReviewersComments = v
}

// GetForwardHistory returns the ForwardHistory field value if set, zero value otherwise.
func (o *CompletedApproval) GetForwardHistory() []ApprovalForwardHistory {
	if o == nil || isNil(o.ForwardHistory) {
		var ret []ApprovalForwardHistory
		return ret
	}
	return o.ForwardHistory
}

// GetForwardHistoryOk returns a tuple with the ForwardHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletedApproval) GetForwardHistoryOk() ([]ApprovalForwardHistory, bool) {
	if o == nil || isNil(o.ForwardHistory) {
		return nil, false
	}
	return o.ForwardHistory, true
}

// HasForwardHistory returns a boolean if a field has been set.
func (o *CompletedApproval) HasForwardHistory() bool {
	if o != nil && !isNil(o.ForwardHistory) {
		return true
	}

	return false
}

// SetForwardHistory gets a reference to the given []ApprovalForwardHistory and assigns it to the ForwardHistory field.
func (o *CompletedApproval) SetForwardHistory(v []ApprovalForwardHistory) {
	o.ForwardHistory = v
}

// GetCommentRequiredWhenRejected returns the CommentRequiredWhenRejected field value if set, zero value otherwise.
func (o *CompletedApproval) GetCommentRequiredWhenRejected() bool {
	if o == nil || isNil(o.CommentRequiredWhenRejected) {
		var ret bool
		return ret
	}
	return *o.CommentRequiredWhenRejected
}

// GetCommentRequiredWhenRejectedOk returns a tuple with the CommentRequiredWhenRejected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletedApproval) GetCommentRequiredWhenRejectedOk() (*bool, bool) {
	if o == nil || isNil(o.CommentRequiredWhenRejected) {
		return nil, false
	}
	return o.CommentRequiredWhenRejected, true
}

// HasCommentRequiredWhenRejected returns a boolean if a field has been set.
func (o *CompletedApproval) HasCommentRequiredWhenRejected() bool {
	if o != nil && !isNil(o.CommentRequiredWhenRejected) {
		return true
	}

	return false
}

// SetCommentRequiredWhenRejected gets a reference to the given bool and assigns it to the CommentRequiredWhenRejected field.
func (o *CompletedApproval) SetCommentRequiredWhenRejected(v bool) {
	o.CommentRequiredWhenRejected = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CompletedApproval) GetState() CompletedApprovalState {
	if o == nil || isNil(o.State) {
		var ret CompletedApprovalState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletedApproval) GetStateOk() (*CompletedApprovalState, bool) {
	if o == nil || isNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CompletedApproval) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given CompletedApprovalState and assigns it to the State field.
func (o *CompletedApproval) SetState(v CompletedApprovalState) {
	o.State = &v
}

// GetRemoveDate returns the RemoveDate field value if set, zero value otherwise.
func (o *CompletedApproval) GetRemoveDate() time.Time {
	if o == nil || isNil(o.RemoveDate) {
		var ret time.Time
		return ret
	}
	return *o.RemoveDate
}

// GetRemoveDateOk returns a tuple with the RemoveDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletedApproval) GetRemoveDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.RemoveDate) {
		return nil, false
	}
	return o.RemoveDate, true
}

// HasRemoveDate returns a boolean if a field has been set.
func (o *CompletedApproval) HasRemoveDate() bool {
	if o != nil && !isNil(o.RemoveDate) {
		return true
	}

	return false
}

// SetRemoveDate gets a reference to the given time.Time and assigns it to the RemoveDate field.
func (o *CompletedApproval) SetRemoveDate(v time.Time) {
	o.RemoveDate = &v
}

// GetRemoveDateUpdateRequested returns the RemoveDateUpdateRequested field value if set, zero value otherwise.
func (o *CompletedApproval) GetRemoveDateUpdateRequested() bool {
	if o == nil || isNil(o.RemoveDateUpdateRequested) {
		var ret bool
		return ret
	}
	return *o.RemoveDateUpdateRequested
}

// GetRemoveDateUpdateRequestedOk returns a tuple with the RemoveDateUpdateRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletedApproval) GetRemoveDateUpdateRequestedOk() (*bool, bool) {
	if o == nil || isNil(o.RemoveDateUpdateRequested) {
		return nil, false
	}
	return o.RemoveDateUpdateRequested, true
}

// HasRemoveDateUpdateRequested returns a boolean if a field has been set.
func (o *CompletedApproval) HasRemoveDateUpdateRequested() bool {
	if o != nil && !isNil(o.RemoveDateUpdateRequested) {
		return true
	}

	return false
}

// SetRemoveDateUpdateRequested gets a reference to the given bool and assigns it to the RemoveDateUpdateRequested field.
func (o *CompletedApproval) SetRemoveDateUpdateRequested(v bool) {
	o.RemoveDateUpdateRequested = &v
}

// GetCurrentRemoveDate returns the CurrentRemoveDate field value if set, zero value otherwise.
func (o *CompletedApproval) GetCurrentRemoveDate() time.Time {
	if o == nil || isNil(o.CurrentRemoveDate) {
		var ret time.Time
		return ret
	}
	return *o.CurrentRemoveDate
}

// GetCurrentRemoveDateOk returns a tuple with the CurrentRemoveDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletedApproval) GetCurrentRemoveDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.CurrentRemoveDate) {
		return nil, false
	}
	return o.CurrentRemoveDate, true
}

// HasCurrentRemoveDate returns a boolean if a field has been set.
func (o *CompletedApproval) HasCurrentRemoveDate() bool {
	if o != nil && !isNil(o.CurrentRemoveDate) {
		return true
	}

	return false
}

// SetCurrentRemoveDate gets a reference to the given time.Time and assigns it to the CurrentRemoveDate field.
func (o *CompletedApproval) SetCurrentRemoveDate(v time.Time) {
	o.CurrentRemoveDate = &v
}

// GetSodViolationContext returns the SodViolationContext field value if set, zero value otherwise.
func (o *CompletedApproval) GetSodViolationContext() SodViolationContextCheckCompleted {
	if o == nil || isNil(o.SodViolationContext) {
		var ret SodViolationContextCheckCompleted
		return ret
	}
	return *o.SodViolationContext
}

// GetSodViolationContextOk returns a tuple with the SodViolationContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletedApproval) GetSodViolationContextOk() (*SodViolationContextCheckCompleted, bool) {
	if o == nil || isNil(o.SodViolationContext) {
		return nil, false
	}
	return o.SodViolationContext, true
}

// HasSodViolationContext returns a boolean if a field has been set.
func (o *CompletedApproval) HasSodViolationContext() bool {
	if o != nil && !isNil(o.SodViolationContext) {
		return true
	}

	return false
}

// SetSodViolationContext gets a reference to the given SodViolationContextCheckCompleted and assigns it to the SodViolationContext field.
func (o *CompletedApproval) SetSodViolationContext(v SodViolationContextCheckCompleted) {
	o.SodViolationContext = &v
}

func (o CompletedApproval) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !isNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !isNil(o.RequestCreated) {
		toSerialize["requestCreated"] = o.RequestCreated
	}
	if !isNil(o.RequestType) {
		toSerialize["requestType"] = o.RequestType
	}
	if !isNil(o.Requester) {
		toSerialize["requester"] = o.Requester
	}
	if !isNil(o.RequestedFor) {
		toSerialize["requestedFor"] = o.RequestedFor
	}
	if !isNil(o.ReviewedBy) {
		toSerialize["reviewedBy"] = o.ReviewedBy
	}
	if !isNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !isNil(o.RequestedObject) {
		toSerialize["requestedObject"] = o.RequestedObject
	}
	if !isNil(o.RequesterComment) {
		toSerialize["requesterComment"] = o.RequesterComment
	}
	if !isNil(o.ReviewerComment) {
		toSerialize["reviewerComment"] = o.ReviewerComment
	}
	if !isNil(o.PreviousReviewersComments) {
		toSerialize["previousReviewersComments"] = o.PreviousReviewersComments
	}
	if !isNil(o.ForwardHistory) {
		toSerialize["forwardHistory"] = o.ForwardHistory
	}
	if !isNil(o.CommentRequiredWhenRejected) {
		toSerialize["commentRequiredWhenRejected"] = o.CommentRequiredWhenRejected
	}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !isNil(o.RemoveDate) {
		toSerialize["removeDate"] = o.RemoveDate
	}
	if !isNil(o.RemoveDateUpdateRequested) {
		toSerialize["removeDateUpdateRequested"] = o.RemoveDateUpdateRequested
	}
	if !isNil(o.CurrentRemoveDate) {
		toSerialize["currentRemoveDate"] = o.CurrentRemoveDate
	}
	if !isNil(o.SodViolationContext) {
		toSerialize["sodViolationContext"] = o.SodViolationContext
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CompletedApproval) UnmarshalJSON(bytes []byte) (err error) {
	varCompletedApproval := _CompletedApproval{}

	if err = json.Unmarshal(bytes, &varCompletedApproval); err == nil {
		*o = CompletedApproval(varCompletedApproval)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "requestCreated")
		delete(additionalProperties, "requestType")
		delete(additionalProperties, "requester")
		delete(additionalProperties, "requestedFor")
		delete(additionalProperties, "reviewedBy")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "requestedObject")
		delete(additionalProperties, "requesterComment")
		delete(additionalProperties, "reviewerComment")
		delete(additionalProperties, "previousReviewersComments")
		delete(additionalProperties, "forwardHistory")
		delete(additionalProperties, "commentRequiredWhenRejected")
		delete(additionalProperties, "state")
		delete(additionalProperties, "removeDate")
		delete(additionalProperties, "removeDateUpdateRequested")
		delete(additionalProperties, "currentRemoveDate")
		delete(additionalProperties, "sodViolationContext")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCompletedApproval struct {
	value *CompletedApproval
	isSet bool
}

func (v NullableCompletedApproval) Get() *CompletedApproval {
	return v.value
}

func (v *NullableCompletedApproval) Set(val *CompletedApproval) {
	v.value = val
	v.isSet = true
}

func (v NullableCompletedApproval) IsSet() bool {
	return v.isSet
}

func (v *NullableCompletedApproval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompletedApproval(val *CompletedApproval) *NullableCompletedApproval {
	return &NullableCompletedApproval{value: val, isSet: true}
}

func (v NullableCompletedApproval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompletedApproval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

