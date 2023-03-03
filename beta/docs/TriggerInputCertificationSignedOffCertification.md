# TriggerInputCertificationSignedOffCertification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID of the certification. | 
**Name** | **string** | The name of the certification. | 
**Created** | **time.Time** | The date and time the certification was created. | 
**Modified** | Pointer to **NullableTime** | The date and time the certification was last modified. | [optional] 
**CampaignRef** | [**CampaignReference**](CampaignReference.md) |  | 
**Phase** | [**CertificationPhase**](CertificationPhase.md) |  | 
**Due** | **time.Time** | The due date of the certification. | 
**Signed** | **time.Time** | The date the reviewer signed off on the certification. | 
**Reviewer** | [**Reviewer**](Reviewer.md) |  | 
**Reassignment** | Pointer to [**Reassignment**](Reassignment.md) |  | [optional] 
**HasErrors** | **bool** | Indicates it the certification has any errors. | 
**ErrorMessage** | Pointer to **NullableString** | A message indicating what the error is. | [optional] 
**Completed** | **bool** | Indicates if all certification decisions have been made. | 
**DecisionsMade** | **int32** | The number of approve/revoke/acknowledge decisions that have been made by the reviewer. | 
**DecisionsTotal** | **int32** | The total number of approve/revoke/acknowledge decisions for the certification. | 
**EntitiesCompleted** | **int32** | The number of entities (identities, access profiles, roles, etc.) for which all decisions have been made and are complete. | 
**EntitiesTotal** | **int32** | The total number of entities (identities, access profiles, roles, etc.) in the certification, both complete and incomplete. | 

## Methods

### NewTriggerInputCertificationSignedOffCertification

`func NewTriggerInputCertificationSignedOffCertification(id string, name string, created time.Time, campaignRef CampaignReference, phase CertificationPhase, due time.Time, signed time.Time, reviewer Reviewer, hasErrors bool, completed bool, decisionsMade int32, decisionsTotal int32, entitiesCompleted int32, entitiesTotal int32, ) *TriggerInputCertificationSignedOffCertification`

NewTriggerInputCertificationSignedOffCertification instantiates a new TriggerInputCertificationSignedOffCertification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerInputCertificationSignedOffCertificationWithDefaults

`func NewTriggerInputCertificationSignedOffCertificationWithDefaults() *TriggerInputCertificationSignedOffCertification`

NewTriggerInputCertificationSignedOffCertificationWithDefaults instantiates a new TriggerInputCertificationSignedOffCertification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TriggerInputCertificationSignedOffCertification) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TriggerInputCertificationSignedOffCertification) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TriggerInputCertificationSignedOffCertification) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TriggerInputCertificationSignedOffCertification) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TriggerInputCertificationSignedOffCertification) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TriggerInputCertificationSignedOffCertification) SetName(v string)`

SetName sets Name field to given value.


### GetCreated

`func (o *TriggerInputCertificationSignedOffCertification) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TriggerInputCertificationSignedOffCertification) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TriggerInputCertificationSignedOffCertification) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *TriggerInputCertificationSignedOffCertification) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *TriggerInputCertificationSignedOffCertification) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *TriggerInputCertificationSignedOffCertification) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *TriggerInputCertificationSignedOffCertification) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModifiedNil

`func (o *TriggerInputCertificationSignedOffCertification) SetModifiedNil(b bool)`

 SetModifiedNil sets the value for Modified to be an explicit nil

### UnsetModified
`func (o *TriggerInputCertificationSignedOffCertification) UnsetModified()`

UnsetModified ensures that no value is present for Modified, not even an explicit nil
### GetCampaignRef

`func (o *TriggerInputCertificationSignedOffCertification) GetCampaignRef() CampaignReference`

GetCampaignRef returns the CampaignRef field if non-nil, zero value otherwise.

### GetCampaignRefOk

`func (o *TriggerInputCertificationSignedOffCertification) GetCampaignRefOk() (*CampaignReference, bool)`

GetCampaignRefOk returns a tuple with the CampaignRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignRef

`func (o *TriggerInputCertificationSignedOffCertification) SetCampaignRef(v CampaignReference)`

SetCampaignRef sets CampaignRef field to given value.


### GetPhase

`func (o *TriggerInputCertificationSignedOffCertification) GetPhase() CertificationPhase`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *TriggerInputCertificationSignedOffCertification) GetPhaseOk() (*CertificationPhase, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *TriggerInputCertificationSignedOffCertification) SetPhase(v CertificationPhase)`

SetPhase sets Phase field to given value.


### GetDue

`func (o *TriggerInputCertificationSignedOffCertification) GetDue() time.Time`

GetDue returns the Due field if non-nil, zero value otherwise.

### GetDueOk

`func (o *TriggerInputCertificationSignedOffCertification) GetDueOk() (*time.Time, bool)`

GetDueOk returns a tuple with the Due field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDue

`func (o *TriggerInputCertificationSignedOffCertification) SetDue(v time.Time)`

SetDue sets Due field to given value.


### GetSigned

`func (o *TriggerInputCertificationSignedOffCertification) GetSigned() time.Time`

GetSigned returns the Signed field if non-nil, zero value otherwise.

### GetSignedOk

`func (o *TriggerInputCertificationSignedOffCertification) GetSignedOk() (*time.Time, bool)`

GetSignedOk returns a tuple with the Signed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigned

`func (o *TriggerInputCertificationSignedOffCertification) SetSigned(v time.Time)`

SetSigned sets Signed field to given value.


### GetReviewer

`func (o *TriggerInputCertificationSignedOffCertification) GetReviewer() Reviewer`

GetReviewer returns the Reviewer field if non-nil, zero value otherwise.

### GetReviewerOk

`func (o *TriggerInputCertificationSignedOffCertification) GetReviewerOk() (*Reviewer, bool)`

GetReviewerOk returns a tuple with the Reviewer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewer

`func (o *TriggerInputCertificationSignedOffCertification) SetReviewer(v Reviewer)`

SetReviewer sets Reviewer field to given value.


### GetReassignment

`func (o *TriggerInputCertificationSignedOffCertification) GetReassignment() Reassignment`

GetReassignment returns the Reassignment field if non-nil, zero value otherwise.

### GetReassignmentOk

`func (o *TriggerInputCertificationSignedOffCertification) GetReassignmentOk() (*Reassignment, bool)`

GetReassignmentOk returns a tuple with the Reassignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReassignment

`func (o *TriggerInputCertificationSignedOffCertification) SetReassignment(v Reassignment)`

SetReassignment sets Reassignment field to given value.

### HasReassignment

`func (o *TriggerInputCertificationSignedOffCertification) HasReassignment() bool`

HasReassignment returns a boolean if a field has been set.

### GetHasErrors

`func (o *TriggerInputCertificationSignedOffCertification) GetHasErrors() bool`

GetHasErrors returns the HasErrors field if non-nil, zero value otherwise.

### GetHasErrorsOk

`func (o *TriggerInputCertificationSignedOffCertification) GetHasErrorsOk() (*bool, bool)`

GetHasErrorsOk returns a tuple with the HasErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasErrors

`func (o *TriggerInputCertificationSignedOffCertification) SetHasErrors(v bool)`

SetHasErrors sets HasErrors field to given value.


### GetErrorMessage

`func (o *TriggerInputCertificationSignedOffCertification) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *TriggerInputCertificationSignedOffCertification) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *TriggerInputCertificationSignedOffCertification) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *TriggerInputCertificationSignedOffCertification) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *TriggerInputCertificationSignedOffCertification) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *TriggerInputCertificationSignedOffCertification) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetCompleted

`func (o *TriggerInputCertificationSignedOffCertification) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *TriggerInputCertificationSignedOffCertification) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *TriggerInputCertificationSignedOffCertification) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.


### GetDecisionsMade

`func (o *TriggerInputCertificationSignedOffCertification) GetDecisionsMade() int32`

GetDecisionsMade returns the DecisionsMade field if non-nil, zero value otherwise.

### GetDecisionsMadeOk

`func (o *TriggerInputCertificationSignedOffCertification) GetDecisionsMadeOk() (*int32, bool)`

GetDecisionsMadeOk returns a tuple with the DecisionsMade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionsMade

`func (o *TriggerInputCertificationSignedOffCertification) SetDecisionsMade(v int32)`

SetDecisionsMade sets DecisionsMade field to given value.


### GetDecisionsTotal

`func (o *TriggerInputCertificationSignedOffCertification) GetDecisionsTotal() int32`

GetDecisionsTotal returns the DecisionsTotal field if non-nil, zero value otherwise.

### GetDecisionsTotalOk

`func (o *TriggerInputCertificationSignedOffCertification) GetDecisionsTotalOk() (*int32, bool)`

GetDecisionsTotalOk returns a tuple with the DecisionsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionsTotal

`func (o *TriggerInputCertificationSignedOffCertification) SetDecisionsTotal(v int32)`

SetDecisionsTotal sets DecisionsTotal field to given value.


### GetEntitiesCompleted

`func (o *TriggerInputCertificationSignedOffCertification) GetEntitiesCompleted() int32`

GetEntitiesCompleted returns the EntitiesCompleted field if non-nil, zero value otherwise.

### GetEntitiesCompletedOk

`func (o *TriggerInputCertificationSignedOffCertification) GetEntitiesCompletedOk() (*int32, bool)`

GetEntitiesCompletedOk returns a tuple with the EntitiesCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitiesCompleted

`func (o *TriggerInputCertificationSignedOffCertification) SetEntitiesCompleted(v int32)`

SetEntitiesCompleted sets EntitiesCompleted field to given value.


### GetEntitiesTotal

`func (o *TriggerInputCertificationSignedOffCertification) GetEntitiesTotal() int32`

GetEntitiesTotal returns the EntitiesTotal field if non-nil, zero value otherwise.

### GetEntitiesTotalOk

`func (o *TriggerInputCertificationSignedOffCertification) GetEntitiesTotalOk() (*int32, bool)`

GetEntitiesTotalOk returns a tuple with the EntitiesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitiesTotal

`func (o *TriggerInputCertificationSignedOffCertification) SetEntitiesTotal(v int32)`

SetEntitiesTotal sets EntitiesTotal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


