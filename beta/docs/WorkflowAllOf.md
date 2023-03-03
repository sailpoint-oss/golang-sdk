# WorkflowAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Workflow ID. This is a UUID generated upon creation. | [optional] 
**ExecutionCount** | Pointer to **int32** | The number of times this workflow has been executed | [optional] 
**FailureCount** | Pointer to **int32** | The number of times this workflow has failed during execution | [optional] 
**Created** | Pointer to **time.Time** | The date and time the workflow was created | [optional] 
**Creator** | Pointer to [**BaseReferenceDto**](BaseReferenceDto.md) | The identity that created the workflow. | [optional] 

## Methods

### NewWorkflowAllOf

`func NewWorkflowAllOf() *WorkflowAllOf`

NewWorkflowAllOf instantiates a new WorkflowAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowAllOfWithDefaults

`func NewWorkflowAllOfWithDefaults() *WorkflowAllOf`

NewWorkflowAllOfWithDefaults instantiates a new WorkflowAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExecutionCount

`func (o *WorkflowAllOf) GetExecutionCount() int32`

GetExecutionCount returns the ExecutionCount field if non-nil, zero value otherwise.

### GetExecutionCountOk

`func (o *WorkflowAllOf) GetExecutionCountOk() (*int32, bool)`

GetExecutionCountOk returns a tuple with the ExecutionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCount

`func (o *WorkflowAllOf) SetExecutionCount(v int32)`

SetExecutionCount sets ExecutionCount field to given value.

### HasExecutionCount

`func (o *WorkflowAllOf) HasExecutionCount() bool`

HasExecutionCount returns a boolean if a field has been set.

### GetFailureCount

`func (o *WorkflowAllOf) GetFailureCount() int32`

GetFailureCount returns the FailureCount field if non-nil, zero value otherwise.

### GetFailureCountOk

`func (o *WorkflowAllOf) GetFailureCountOk() (*int32, bool)`

GetFailureCountOk returns a tuple with the FailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCount

`func (o *WorkflowAllOf) SetFailureCount(v int32)`

SetFailureCount sets FailureCount field to given value.

### HasFailureCount

`func (o *WorkflowAllOf) HasFailureCount() bool`

HasFailureCount returns a boolean if a field has been set.

### GetCreated

`func (o *WorkflowAllOf) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WorkflowAllOf) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WorkflowAllOf) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *WorkflowAllOf) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreator

`func (o *WorkflowAllOf) GetCreator() BaseReferenceDto`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *WorkflowAllOf) GetCreatorOk() (*BaseReferenceDto, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *WorkflowAllOf) SetCreator(v BaseReferenceDto)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *WorkflowAllOf) HasCreator() bool`

HasCreator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


