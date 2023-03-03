# TriggerInputAccessRequestPreApproval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessRequestId** | **string** | The unique ID of the access request. | 
**RequestedFor** | [**TriggerInputAccessRequestPostApprovalRequestedFor**](TriggerInputAccessRequestPostApprovalRequestedFor.md) |  | 
**RequestedItems** | [**[]TriggerInputAccessRequestPreApprovalRequestedItemsInner**](TriggerInputAccessRequestPreApprovalRequestedItemsInner.md) | Details of the access items being requested. | 
**RequestedBy** | [**TriggerInputAccessRequestPostApprovalRequestedBy**](TriggerInputAccessRequestPostApprovalRequestedBy.md) |  | 

## Methods

### NewTriggerInputAccessRequestPreApproval

`func NewTriggerInputAccessRequestPreApproval(accessRequestId string, requestedFor TriggerInputAccessRequestPostApprovalRequestedFor, requestedItems []TriggerInputAccessRequestPreApprovalRequestedItemsInner, requestedBy TriggerInputAccessRequestPostApprovalRequestedBy, ) *TriggerInputAccessRequestPreApproval`

NewTriggerInputAccessRequestPreApproval instantiates a new TriggerInputAccessRequestPreApproval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerInputAccessRequestPreApprovalWithDefaults

`func NewTriggerInputAccessRequestPreApprovalWithDefaults() *TriggerInputAccessRequestPreApproval`

NewTriggerInputAccessRequestPreApprovalWithDefaults instantiates a new TriggerInputAccessRequestPreApproval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessRequestId

`func (o *TriggerInputAccessRequestPreApproval) GetAccessRequestId() string`

GetAccessRequestId returns the AccessRequestId field if non-nil, zero value otherwise.

### GetAccessRequestIdOk

`func (o *TriggerInputAccessRequestPreApproval) GetAccessRequestIdOk() (*string, bool)`

GetAccessRequestIdOk returns a tuple with the AccessRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRequestId

`func (o *TriggerInputAccessRequestPreApproval) SetAccessRequestId(v string)`

SetAccessRequestId sets AccessRequestId field to given value.


### GetRequestedFor

`func (o *TriggerInputAccessRequestPreApproval) GetRequestedFor() TriggerInputAccessRequestPostApprovalRequestedFor`

GetRequestedFor returns the RequestedFor field if non-nil, zero value otherwise.

### GetRequestedForOk

`func (o *TriggerInputAccessRequestPreApproval) GetRequestedForOk() (*TriggerInputAccessRequestPostApprovalRequestedFor, bool)`

GetRequestedForOk returns a tuple with the RequestedFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedFor

`func (o *TriggerInputAccessRequestPreApproval) SetRequestedFor(v TriggerInputAccessRequestPostApprovalRequestedFor)`

SetRequestedFor sets RequestedFor field to given value.


### GetRequestedItems

`func (o *TriggerInputAccessRequestPreApproval) GetRequestedItems() []TriggerInputAccessRequestPreApprovalRequestedItemsInner`

GetRequestedItems returns the RequestedItems field if non-nil, zero value otherwise.

### GetRequestedItemsOk

`func (o *TriggerInputAccessRequestPreApproval) GetRequestedItemsOk() (*[]TriggerInputAccessRequestPreApprovalRequestedItemsInner, bool)`

GetRequestedItemsOk returns a tuple with the RequestedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedItems

`func (o *TriggerInputAccessRequestPreApproval) SetRequestedItems(v []TriggerInputAccessRequestPreApprovalRequestedItemsInner)`

SetRequestedItems sets RequestedItems field to given value.


### GetRequestedBy

`func (o *TriggerInputAccessRequestPreApproval) GetRequestedBy() TriggerInputAccessRequestPostApprovalRequestedBy`

GetRequestedBy returns the RequestedBy field if non-nil, zero value otherwise.

### GetRequestedByOk

`func (o *TriggerInputAccessRequestPreApproval) GetRequestedByOk() (*TriggerInputAccessRequestPostApprovalRequestedBy, bool)`

GetRequestedByOk returns a tuple with the RequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBy

`func (o *TriggerInputAccessRequestPreApproval) SetRequestedBy(v TriggerInputAccessRequestPostApprovalRequestedBy)`

SetRequestedBy sets RequestedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


