# RequestedItemStatusCancelledRequestDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Comment made by the owner when cancelling the associated request. | [optional] 
**Owner** | Pointer to [**OwnerDto**](OwnerDto.md) |  | [optional] 
**Modified** | Pointer to **time.Time** | Date comment was added by the owner when cancelling the associated request. | [optional] 

## Methods

### NewRequestedItemStatusCancelledRequestDetails

`func NewRequestedItemStatusCancelledRequestDetails() *RequestedItemStatusCancelledRequestDetails`

NewRequestedItemStatusCancelledRequestDetails instantiates a new RequestedItemStatusCancelledRequestDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestedItemStatusCancelledRequestDetailsWithDefaults

`func NewRequestedItemStatusCancelledRequestDetailsWithDefaults() *RequestedItemStatusCancelledRequestDetails`

NewRequestedItemStatusCancelledRequestDetailsWithDefaults instantiates a new RequestedItemStatusCancelledRequestDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *RequestedItemStatusCancelledRequestDetails) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RequestedItemStatusCancelledRequestDetails) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RequestedItemStatusCancelledRequestDetails) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RequestedItemStatusCancelledRequestDetails) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetOwner

`func (o *RequestedItemStatusCancelledRequestDetails) GetOwner() OwnerDto`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *RequestedItemStatusCancelledRequestDetails) GetOwnerOk() (*OwnerDto, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *RequestedItemStatusCancelledRequestDetails) SetOwner(v OwnerDto)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *RequestedItemStatusCancelledRequestDetails) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetModified

`func (o *RequestedItemStatusCancelledRequestDetails) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *RequestedItemStatusCancelledRequestDetails) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *RequestedItemStatusCancelledRequestDetails) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *RequestedItemStatusCancelledRequestDetails) HasModified() bool`

HasModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


