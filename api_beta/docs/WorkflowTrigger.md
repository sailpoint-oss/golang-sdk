# WorkflowTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The trigger type | 
**Attributes** | **map[string]interface{}** | Workflow Trigger Attributes. | 

## Methods

### NewWorkflowTrigger

`func NewWorkflowTrigger(type_ string, attributes map[string]interface{}, ) *WorkflowTrigger`

NewWorkflowTrigger instantiates a new WorkflowTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTriggerWithDefaults

`func NewWorkflowTriggerWithDefaults() *WorkflowTrigger`

NewWorkflowTriggerWithDefaults instantiates a new WorkflowTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WorkflowTrigger) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowTrigger) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowTrigger) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *WorkflowTrigger) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WorkflowTrigger) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WorkflowTrigger) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


