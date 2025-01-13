# AccessRequestPreApproval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessRequestId** | **string** | The unique ID of the access request. | 
**RequestedFor** | [**AccessItemRequestedForDto**](AccessItemRequestedForDto.md) |  | 
**RequestedItems** | [**[]AccessRequestPreApprovalRequestedItemsInner**](AccessRequestPreApprovalRequestedItemsInner.md) | Details of the access items being requested. | 
**RequestedBy** | [**AccessItemRequesterDto**](AccessItemRequesterDto.md) |  | 

## Methods

### NewAccessRequestPreApproval

`func NewAccessRequestPreApproval(accessRequestId string, requestedFor AccessItemRequestedForDto, requestedItems []AccessRequestPreApprovalRequestedItemsInner, requestedBy AccessItemRequesterDto, ) *AccessRequestPreApproval`

NewAccessRequestPreApproval instantiates a new AccessRequestPreApproval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessRequestPreApprovalWithDefaults

`func NewAccessRequestPreApprovalWithDefaults() *AccessRequestPreApproval`

NewAccessRequestPreApprovalWithDefaults instantiates a new AccessRequestPreApproval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessRequestId

`func (o *AccessRequestPreApproval) GetAccessRequestId() string`

GetAccessRequestId returns the AccessRequestId field if non-nil, zero value otherwise.

### GetAccessRequestIdOk

`func (o *AccessRequestPreApproval) GetAccessRequestIdOk() (*string, bool)`

GetAccessRequestIdOk returns a tuple with the AccessRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRequestId

`func (o *AccessRequestPreApproval) SetAccessRequestId(v string)`

SetAccessRequestId sets AccessRequestId field to given value.


### GetRequestedFor

`func (o *AccessRequestPreApproval) GetRequestedFor() AccessItemRequestedForDto`

GetRequestedFor returns the RequestedFor field if non-nil, zero value otherwise.

### GetRequestedForOk

`func (o *AccessRequestPreApproval) GetRequestedForOk() (*AccessItemRequestedForDto, bool)`

GetRequestedForOk returns a tuple with the RequestedFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedFor

`func (o *AccessRequestPreApproval) SetRequestedFor(v AccessItemRequestedForDto)`

SetRequestedFor sets RequestedFor field to given value.


### GetRequestedItems

`func (o *AccessRequestPreApproval) GetRequestedItems() []AccessRequestPreApprovalRequestedItemsInner`

GetRequestedItems returns the RequestedItems field if non-nil, zero value otherwise.

### GetRequestedItemsOk

`func (o *AccessRequestPreApproval) GetRequestedItemsOk() (*[]AccessRequestPreApprovalRequestedItemsInner, bool)`

GetRequestedItemsOk returns a tuple with the RequestedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedItems

`func (o *AccessRequestPreApproval) SetRequestedItems(v []AccessRequestPreApprovalRequestedItemsInner)`

SetRequestedItems sets RequestedItems field to given value.


### GetRequestedBy

`func (o *AccessRequestPreApproval) GetRequestedBy() AccessItemRequesterDto`

GetRequestedBy returns the RequestedBy field if non-nil, zero value otherwise.

### GetRequestedByOk

`func (o *AccessRequestPreApproval) GetRequestedByOk() (*AccessItemRequesterDto, bool)`

GetRequestedByOk returns a tuple with the RequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBy

`func (o *AccessRequestPreApproval) SetRequestedBy(v AccessItemRequesterDto)`

SetRequestedBy sets RequestedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


