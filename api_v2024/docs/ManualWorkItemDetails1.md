# ManualWorkItemDetails1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Forwarded** | Pointer to **bool** | True if the request for this item was forwarded from one owner to another. | [optional] [default to false]
**OriginalOwner** | Pointer to [**NullableManualWorkItemDetailsOriginalOwner**](ManualWorkItemDetailsOriginalOwner.md) |  | [optional] 
**CurrentOwner** | Pointer to [**NullableManualWorkItemDetailsCurrentOwner**](ManualWorkItemDetailsCurrentOwner.md) |  | [optional] 
**Modified** | Pointer to **time.Time** | Time at which item was modified. | [optional] 
**Status** | Pointer to [**ManualWorkItemState**](ManualWorkItemState.md) |  | [optional] 
**ForwardHistory** | Pointer to [**[]ApprovalForwardHistory1**](ApprovalForwardHistory1.md) | The history of approval forward action. | [optional] 

## Methods

### NewManualWorkItemDetails1

`func NewManualWorkItemDetails1() *ManualWorkItemDetails1`

NewManualWorkItemDetails1 instantiates a new ManualWorkItemDetails1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualWorkItemDetails1WithDefaults

`func NewManualWorkItemDetails1WithDefaults() *ManualWorkItemDetails1`

NewManualWorkItemDetails1WithDefaults instantiates a new ManualWorkItemDetails1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForwarded

`func (o *ManualWorkItemDetails1) GetForwarded() bool`

GetForwarded returns the Forwarded field if non-nil, zero value otherwise.

### GetForwardedOk

`func (o *ManualWorkItemDetails1) GetForwardedOk() (*bool, bool)`

GetForwardedOk returns a tuple with the Forwarded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwarded

`func (o *ManualWorkItemDetails1) SetForwarded(v bool)`

SetForwarded sets Forwarded field to given value.

### HasForwarded

`func (o *ManualWorkItemDetails1) HasForwarded() bool`

HasForwarded returns a boolean if a field has been set.

### GetOriginalOwner

`func (o *ManualWorkItemDetails1) GetOriginalOwner() ManualWorkItemDetailsOriginalOwner`

GetOriginalOwner returns the OriginalOwner field if non-nil, zero value otherwise.

### GetOriginalOwnerOk

`func (o *ManualWorkItemDetails1) GetOriginalOwnerOk() (*ManualWorkItemDetailsOriginalOwner, bool)`

GetOriginalOwnerOk returns a tuple with the OriginalOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalOwner

`func (o *ManualWorkItemDetails1) SetOriginalOwner(v ManualWorkItemDetailsOriginalOwner)`

SetOriginalOwner sets OriginalOwner field to given value.

### HasOriginalOwner

`func (o *ManualWorkItemDetails1) HasOriginalOwner() bool`

HasOriginalOwner returns a boolean if a field has been set.

### SetOriginalOwnerNil

`func (o *ManualWorkItemDetails1) SetOriginalOwnerNil(b bool)`

 SetOriginalOwnerNil sets the value for OriginalOwner to be an explicit nil

### UnsetOriginalOwner
`func (o *ManualWorkItemDetails1) UnsetOriginalOwner()`

UnsetOriginalOwner ensures that no value is present for OriginalOwner, not even an explicit nil
### GetCurrentOwner

`func (o *ManualWorkItemDetails1) GetCurrentOwner() ManualWorkItemDetailsCurrentOwner`

GetCurrentOwner returns the CurrentOwner field if non-nil, zero value otherwise.

### GetCurrentOwnerOk

`func (o *ManualWorkItemDetails1) GetCurrentOwnerOk() (*ManualWorkItemDetailsCurrentOwner, bool)`

GetCurrentOwnerOk returns a tuple with the CurrentOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentOwner

`func (o *ManualWorkItemDetails1) SetCurrentOwner(v ManualWorkItemDetailsCurrentOwner)`

SetCurrentOwner sets CurrentOwner field to given value.

### HasCurrentOwner

`func (o *ManualWorkItemDetails1) HasCurrentOwner() bool`

HasCurrentOwner returns a boolean if a field has been set.

### SetCurrentOwnerNil

`func (o *ManualWorkItemDetails1) SetCurrentOwnerNil(b bool)`

 SetCurrentOwnerNil sets the value for CurrentOwner to be an explicit nil

### UnsetCurrentOwner
`func (o *ManualWorkItemDetails1) UnsetCurrentOwner()`

UnsetCurrentOwner ensures that no value is present for CurrentOwner, not even an explicit nil
### GetModified

`func (o *ManualWorkItemDetails1) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ManualWorkItemDetails1) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ManualWorkItemDetails1) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *ManualWorkItemDetails1) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetStatus

`func (o *ManualWorkItemDetails1) GetStatus() ManualWorkItemState`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ManualWorkItemDetails1) GetStatusOk() (*ManualWorkItemState, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ManualWorkItemDetails1) SetStatus(v ManualWorkItemState)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ManualWorkItemDetails1) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetForwardHistory

`func (o *ManualWorkItemDetails1) GetForwardHistory() []ApprovalForwardHistory1`

GetForwardHistory returns the ForwardHistory field if non-nil, zero value otherwise.

### GetForwardHistoryOk

`func (o *ManualWorkItemDetails1) GetForwardHistoryOk() (*[]ApprovalForwardHistory1, bool)`

GetForwardHistoryOk returns a tuple with the ForwardHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardHistory

`func (o *ManualWorkItemDetails1) SetForwardHistory(v []ApprovalForwardHistory1)`

SetForwardHistory sets ForwardHistory field to given value.

### HasForwardHistory

`func (o *ManualWorkItemDetails1) HasForwardHistory() bool`

HasForwardHistory returns a boolean if a field has been set.

### SetForwardHistoryNil

`func (o *ManualWorkItemDetails1) SetForwardHistoryNil(b bool)`

 SetForwardHistoryNil sets the value for ForwardHistory to be an explicit nil

### UnsetForwardHistory
`func (o *ManualWorkItemDetails1) UnsetForwardHistory()`

UnsetForwardHistory ensures that no value is present for ForwardHistory, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


