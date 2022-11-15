# TriggerOutputAccessRequestDynamicApprover

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID of the identity to add to the approver list for the access request. | 
**Name** | **string** | The name of the identity to add to the approver list for the access request. | 
**Type** | **map[string]interface{}** | The type of object being referenced. | 

## Methods

### NewTriggerOutputAccessRequestDynamicApprover

`func NewTriggerOutputAccessRequestDynamicApprover(id string, name string, type_ map[string]interface{}, ) *TriggerOutputAccessRequestDynamicApprover`

NewTriggerOutputAccessRequestDynamicApprover instantiates a new TriggerOutputAccessRequestDynamicApprover object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerOutputAccessRequestDynamicApproverWithDefaults

`func NewTriggerOutputAccessRequestDynamicApproverWithDefaults() *TriggerOutputAccessRequestDynamicApprover`

NewTriggerOutputAccessRequestDynamicApproverWithDefaults instantiates a new TriggerOutputAccessRequestDynamicApprover object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TriggerOutputAccessRequestDynamicApprover) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TriggerOutputAccessRequestDynamicApprover) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TriggerOutputAccessRequestDynamicApprover) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TriggerOutputAccessRequestDynamicApprover) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TriggerOutputAccessRequestDynamicApprover) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TriggerOutputAccessRequestDynamicApprover) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *TriggerOutputAccessRequestDynamicApprover) GetType() map[string]interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TriggerOutputAccessRequestDynamicApprover) GetTypeOk() (*map[string]interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TriggerOutputAccessRequestDynamicApprover) SetType(v map[string]interface{})`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


