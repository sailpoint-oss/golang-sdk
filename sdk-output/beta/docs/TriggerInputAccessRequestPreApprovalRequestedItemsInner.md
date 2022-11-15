# TriggerInputAccessRequestPreApprovalRequestedItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID of the access item being requested. | 
**Name** | **string** | The human friendly name of the access item. | 
**Description** | Pointer to **NullableString** | Detailed description of the access item. | [optional] 
**Type** | **map[string]interface{}** | The type of access item. | 
**Operation** | **map[string]interface{}** | The action to perform on the access item. | 
**Comment** | Pointer to **NullableString** | A comment from the identity requesting the access. | [optional] 

## Methods

### NewTriggerInputAccessRequestPreApprovalRequestedItemsInner

`func NewTriggerInputAccessRequestPreApprovalRequestedItemsInner(id string, name string, type_ map[string]interface{}, operation map[string]interface{}, ) *TriggerInputAccessRequestPreApprovalRequestedItemsInner`

NewTriggerInputAccessRequestPreApprovalRequestedItemsInner instantiates a new TriggerInputAccessRequestPreApprovalRequestedItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerInputAccessRequestPreApprovalRequestedItemsInnerWithDefaults

`func NewTriggerInputAccessRequestPreApprovalRequestedItemsInnerWithDefaults() *TriggerInputAccessRequestPreApprovalRequestedItemsInner`

NewTriggerInputAccessRequestPreApprovalRequestedItemsInnerWithDefaults instantiates a new TriggerInputAccessRequestPreApprovalRequestedItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TriggerInputAccessRequestPreApprovalRequestedItemsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TriggerInputAccessRequestPreApprovalRequestedItemsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TriggerInputAccessRequestPreApprovalRequestedItemsInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TriggerInputAccessRequestPreApprovalRequestedItemsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TriggerInputAccessRequestPreApprovalRequestedItemsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TriggerInputAccessRequestPreApprovalRequestedItemsInner) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TriggerInputAccessRequestPreApprovalRequestedItemsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TriggerInputAccessRequestPreApprovalRequestedItemsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TriggerInputAccessRequestPreApprovalRequestedItemsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TriggerInputAccessRequestPreApprovalRequestedItemsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TriggerInputAccessRequestPreApprovalRequestedItemsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TriggerInputAccessRequestPreApprovalRequestedItemsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *TriggerInputAccessRequestPreApprovalRequestedItemsInner) GetType() map[string]interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TriggerInputAccessRequestPreApprovalRequestedItemsInner) GetTypeOk() (*map[string]interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TriggerInputAccessRequestPreApprovalRequestedItemsInner) SetType(v map[string]interface{})`

SetType sets Type field to given value.


### GetOperation

`func (o *TriggerInputAccessRequestPreApprovalRequestedItemsInner) GetOperation() map[string]interface{}`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *TriggerInputAccessRequestPreApprovalRequestedItemsInner) GetOperationOk() (*map[string]interface{}, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *TriggerInputAccessRequestPreApprovalRequestedItemsInner) SetOperation(v map[string]interface{})`

SetOperation sets Operation field to given value.


### GetComment

`func (o *TriggerInputAccessRequestPreApprovalRequestedItemsInner) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TriggerInputAccessRequestPreApprovalRequestedItemsInner) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TriggerInputAccessRequestPreApprovalRequestedItemsInner) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TriggerInputAccessRequestPreApprovalRequestedItemsInner) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *TriggerInputAccessRequestPreApprovalRequestedItemsInner) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *TriggerInputAccessRequestPreApprovalRequestedItemsInner) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


