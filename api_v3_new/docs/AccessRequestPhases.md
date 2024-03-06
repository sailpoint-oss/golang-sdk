# AccessRequestPhases

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Started** | Pointer to **time.Time** | The time that this phase started. | [optional] 
**Finished** | Pointer to **time.Time** | The time that this phase finished. | [optional] 
**Name** | Pointer to **string** | The name of this phase. | [optional] 
**State** | Pointer to **string** | The state of this phase. | [optional] 
**Result** | Pointer to **string** | The state of this phase. | [optional] 
**PhaseReference** | Pointer to **string** | A reference to another object on the RequestedItemStatus that contains more details about the phase. Note that for the Provisioning phase, this will be empty if there are no manual work items. | [optional] 

## Methods

### NewAccessRequestPhases

`func NewAccessRequestPhases() *AccessRequestPhases`

NewAccessRequestPhases instantiates a new AccessRequestPhases object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessRequestPhasesWithDefaults

`func NewAccessRequestPhasesWithDefaults() *AccessRequestPhases`

NewAccessRequestPhasesWithDefaults instantiates a new AccessRequestPhases object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStarted

`func (o *AccessRequestPhases) GetStarted() time.Time`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *AccessRequestPhases) GetStartedOk() (*time.Time, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *AccessRequestPhases) SetStarted(v time.Time)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *AccessRequestPhases) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### GetFinished

`func (o *AccessRequestPhases) GetFinished() time.Time`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *AccessRequestPhases) GetFinishedOk() (*time.Time, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *AccessRequestPhases) SetFinished(v time.Time)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *AccessRequestPhases) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### GetName

`func (o *AccessRequestPhases) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessRequestPhases) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessRequestPhases) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccessRequestPhases) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *AccessRequestPhases) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AccessRequestPhases) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AccessRequestPhases) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AccessRequestPhases) HasState() bool`

HasState returns a boolean if a field has been set.

### GetResult

`func (o *AccessRequestPhases) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AccessRequestPhases) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AccessRequestPhases) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *AccessRequestPhases) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetPhaseReference

`func (o *AccessRequestPhases) GetPhaseReference() string`

GetPhaseReference returns the PhaseReference field if non-nil, zero value otherwise.

### GetPhaseReferenceOk

`func (o *AccessRequestPhases) GetPhaseReferenceOk() (*string, bool)`

GetPhaseReferenceOk returns a tuple with the PhaseReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhaseReference

`func (o *AccessRequestPhases) SetPhaseReference(v string)`

SetPhaseReference sets PhaseReference field to given value.

### HasPhaseReference

`func (o *AccessRequestPhases) HasPhaseReference() bool`

HasPhaseReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


