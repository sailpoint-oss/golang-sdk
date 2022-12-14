# AccessRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Approvals** | Pointer to [**[]Approval**](Approval.md) |  | [optional] 
**CompletionDate** | Pointer to **time.Time** |  | [optional] 
**CompletionStatus** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**Errors** | Pointer to **[]string** |  | [optional] 
**ExecutionStatus** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **string** |  | [optional] 
**RequestSummaries** | Pointer to [**[]RequestSummary**](RequestSummary.md) |  | [optional] 
**RequesterId** | Pointer to **string** |  | [optional] 
**RequesterDisplayName** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**SourceName** | Pointer to **string** |  | [optional] 
**TargetName** | Pointer to **string** |  | [optional] 
**TargetDisplayName** | Pointer to **string** |  | [optional] 
**TargetId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Warnings** | Pointer to **[]string** |  | [optional] 
**WorkflowStep** | Pointer to **string** |  | [optional] 

## Methods

### NewAccessRequest

`func NewAccessRequest() *AccessRequest`

NewAccessRequest instantiates a new AccessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessRequestWithDefaults

`func NewAccessRequestWithDefaults() *AccessRequest`

NewAccessRequestWithDefaults instantiates a new AccessRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovals

`func (o *AccessRequest) GetApprovals() []Approval`

GetApprovals returns the Approvals field if non-nil, zero value otherwise.

### GetApprovalsOk

`func (o *AccessRequest) GetApprovalsOk() (*[]Approval, bool)`

GetApprovalsOk returns a tuple with the Approvals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovals

`func (o *AccessRequest) SetApprovals(v []Approval)`

SetApprovals sets Approvals field to given value.

### HasApprovals

`func (o *AccessRequest) HasApprovals() bool`

HasApprovals returns a boolean if a field has been set.

### GetCompletionDate

`func (o *AccessRequest) GetCompletionDate() time.Time`

GetCompletionDate returns the CompletionDate field if non-nil, zero value otherwise.

### GetCompletionDateOk

`func (o *AccessRequest) GetCompletionDateOk() (*time.Time, bool)`

GetCompletionDateOk returns a tuple with the CompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDate

`func (o *AccessRequest) SetCompletionDate(v time.Time)`

SetCompletionDate sets CompletionDate field to given value.

### HasCompletionDate

`func (o *AccessRequest) HasCompletionDate() bool`

HasCompletionDate returns a boolean if a field has been set.

### GetCompletionStatus

`func (o *AccessRequest) GetCompletionStatus() string`

GetCompletionStatus returns the CompletionStatus field if non-nil, zero value otherwise.

### GetCompletionStatusOk

`func (o *AccessRequest) GetCompletionStatusOk() (*string, bool)`

GetCompletionStatusOk returns a tuple with the CompletionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionStatus

`func (o *AccessRequest) SetCompletionStatus(v string)`

SetCompletionStatus sets CompletionStatus field to given value.

### HasCompletionStatus

`func (o *AccessRequest) HasCompletionStatus() bool`

HasCompletionStatus returns a boolean if a field has been set.

### GetDateCreated

`func (o *AccessRequest) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AccessRequest) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AccessRequest) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AccessRequest) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetErrors

`func (o *AccessRequest) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *AccessRequest) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *AccessRequest) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *AccessRequest) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetExecutionStatus

`func (o *AccessRequest) GetExecutionStatus() string`

GetExecutionStatus returns the ExecutionStatus field if non-nil, zero value otherwise.

### GetExecutionStatusOk

`func (o *AccessRequest) GetExecutionStatusOk() (*string, bool)`

GetExecutionStatusOk returns a tuple with the ExecutionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionStatus

`func (o *AccessRequest) SetExecutionStatus(v string)`

SetExecutionStatus sets ExecutionStatus field to given value.

### HasExecutionStatus

`func (o *AccessRequest) HasExecutionStatus() bool`

HasExecutionStatus returns a boolean if a field has been set.

### GetId

`func (o *AccessRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccessRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AccessRequest) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AccessRequest) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AccessRequest) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AccessRequest) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *AccessRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccessRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriority

`func (o *AccessRequest) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *AccessRequest) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *AccessRequest) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *AccessRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRequestSummaries

`func (o *AccessRequest) GetRequestSummaries() []RequestSummary`

GetRequestSummaries returns the RequestSummaries field if non-nil, zero value otherwise.

### GetRequestSummariesOk

`func (o *AccessRequest) GetRequestSummariesOk() (*[]RequestSummary, bool)`

GetRequestSummariesOk returns a tuple with the RequestSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSummaries

`func (o *AccessRequest) SetRequestSummaries(v []RequestSummary)`

SetRequestSummaries sets RequestSummaries field to given value.

### HasRequestSummaries

`func (o *AccessRequest) HasRequestSummaries() bool`

HasRequestSummaries returns a boolean if a field has been set.

### GetRequesterId

`func (o *AccessRequest) GetRequesterId() string`

GetRequesterId returns the RequesterId field if non-nil, zero value otherwise.

### GetRequesterIdOk

`func (o *AccessRequest) GetRequesterIdOk() (*string, bool)`

GetRequesterIdOk returns a tuple with the RequesterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterId

`func (o *AccessRequest) SetRequesterId(v string)`

SetRequesterId sets RequesterId field to given value.

### HasRequesterId

`func (o *AccessRequest) HasRequesterId() bool`

HasRequesterId returns a boolean if a field has been set.

### GetRequesterDisplayName

`func (o *AccessRequest) GetRequesterDisplayName() string`

GetRequesterDisplayName returns the RequesterDisplayName field if non-nil, zero value otherwise.

### GetRequesterDisplayNameOk

`func (o *AccessRequest) GetRequesterDisplayNameOk() (*string, bool)`

GetRequesterDisplayNameOk returns a tuple with the RequesterDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterDisplayName

`func (o *AccessRequest) SetRequesterDisplayName(v string)`

SetRequesterDisplayName sets RequesterDisplayName field to given value.

### HasRequesterDisplayName

`func (o *AccessRequest) HasRequesterDisplayName() bool`

HasRequesterDisplayName returns a boolean if a field has been set.

### GetSource

`func (o *AccessRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AccessRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AccessRequest) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *AccessRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceName

`func (o *AccessRequest) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *AccessRequest) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *AccessRequest) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *AccessRequest) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### GetTargetName

`func (o *AccessRequest) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *AccessRequest) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *AccessRequest) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *AccessRequest) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetTargetDisplayName

`func (o *AccessRequest) GetTargetDisplayName() string`

GetTargetDisplayName returns the TargetDisplayName field if non-nil, zero value otherwise.

### GetTargetDisplayNameOk

`func (o *AccessRequest) GetTargetDisplayNameOk() (*string, bool)`

GetTargetDisplayNameOk returns a tuple with the TargetDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDisplayName

`func (o *AccessRequest) SetTargetDisplayName(v string)`

SetTargetDisplayName sets TargetDisplayName field to given value.

### HasTargetDisplayName

`func (o *AccessRequest) HasTargetDisplayName() bool`

HasTargetDisplayName returns a boolean if a field has been set.

### GetTargetId

`func (o *AccessRequest) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *AccessRequest) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *AccessRequest) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *AccessRequest) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetType

`func (o *AccessRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccessRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccessRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AccessRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWarnings

`func (o *AccessRequest) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *AccessRequest) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *AccessRequest) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *AccessRequest) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetWorkflowStep

`func (o *AccessRequest) GetWorkflowStep() string`

GetWorkflowStep returns the WorkflowStep field if non-nil, zero value otherwise.

### GetWorkflowStepOk

`func (o *AccessRequest) GetWorkflowStepOk() (*string, bool)`

GetWorkflowStepOk returns a tuple with the WorkflowStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStep

`func (o *AccessRequest) SetWorkflowStep(v string)`

SetWorkflowStep sets WorkflowStep field to given value.

### HasWorkflowStep

`func (o *AccessRequest) HasWorkflowStep() bool`

HasWorkflowStep returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


