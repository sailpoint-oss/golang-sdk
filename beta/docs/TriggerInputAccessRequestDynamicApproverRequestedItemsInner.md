# TriggerInputAccessRequestDynamicApproverRequestedItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID of the access item. | 
**Name** | **string** | Human friendly name of the access item. | 
**Description** | Pointer to **NullableString** | Extended description of the access item. | [optional] 
**Type** | **map[string]interface{}** | The type of access item being requested. | 
**Operation** | **map[string]interface{}** | Grant or revoke the access item | 
**Comment** | Pointer to **NullableString** | A comment from the requestor on why the access is needed. | [optional] 

## Methods

### NewTriggerInputAccessRequestDynamicApproverRequestedItemsInner

`func NewTriggerInputAccessRequestDynamicApproverRequestedItemsInner(id string, name string, type_ map[string]interface{}, operation map[string]interface{}, ) *TriggerInputAccessRequestDynamicApproverRequestedItemsInner`

NewTriggerInputAccessRequestDynamicApproverRequestedItemsInner instantiates a new TriggerInputAccessRequestDynamicApproverRequestedItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerInputAccessRequestDynamicApproverRequestedItemsInnerWithDefaults

`func NewTriggerInputAccessRequestDynamicApproverRequestedItemsInnerWithDefaults() *TriggerInputAccessRequestDynamicApproverRequestedItemsInner`

NewTriggerInputAccessRequestDynamicApproverRequestedItemsInnerWithDefaults instantiates a new TriggerInputAccessRequestDynamicApproverRequestedItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TriggerInputAccessRequestDynamicApproverRequestedItemsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TriggerInputAccessRequestDynamicApproverRequestedItemsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TriggerInputAccessRequestDynamicApproverRequestedItemsInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TriggerInputAccessRequestDynamicApproverRequestedItemsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TriggerInputAccessRequestDynamicApproverRequestedItemsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TriggerInputAccessRequestDynamicApproverRequestedItemsInner) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TriggerInputAccessRequestDynamicApproverRequestedItemsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TriggerInputAccessRequestDynamicApproverRequestedItemsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TriggerInputAccessRequestDynamicApproverRequestedItemsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TriggerInputAccessRequestDynamicApproverRequestedItemsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TriggerInputAccessRequestDynamicApproverRequestedItemsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TriggerInputAccessRequestDynamicApproverRequestedItemsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *TriggerInputAccessRequestDynamicApproverRequestedItemsInner) GetType() map[string]interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TriggerInputAccessRequestDynamicApproverRequestedItemsInner) GetTypeOk() (*map[string]interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TriggerInputAccessRequestDynamicApproverRequestedItemsInner) SetType(v map[string]interface{})`

SetType sets Type field to given value.


### GetOperation

`func (o *TriggerInputAccessRequestDynamicApproverRequestedItemsInner) GetOperation() map[string]interface{}`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *TriggerInputAccessRequestDynamicApproverRequestedItemsInner) GetOperationOk() (*map[string]interface{}, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *TriggerInputAccessRequestDynamicApproverRequestedItemsInner) SetOperation(v map[string]interface{})`

SetOperation sets Operation field to given value.


### GetComment

`func (o *TriggerInputAccessRequestDynamicApproverRequestedItemsInner) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TriggerInputAccessRequestDynamicApproverRequestedItemsInner) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TriggerInputAccessRequestDynamicApproverRequestedItemsInner) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TriggerInputAccessRequestDynamicApproverRequestedItemsInner) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *TriggerInputAccessRequestDynamicApproverRequestedItemsInner) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *TriggerInputAccessRequestDynamicApproverRequestedItemsInner) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


