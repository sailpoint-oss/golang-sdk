# InvocationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Invocation ID | 
**TriggerId** | **string** | Trigger ID | 
**SubscriptionId** | **string** | Subscription ID | 
**Type** | [**InvocationStatusType**](InvocationStatusType.md) |  | 
**Created** | **time.Time** | Invocation created timestamp. ISO-8601 in UTC. | 
**Completed** | Pointer to **time.Time** | Invocation completed timestamp; empty fields imply invocation is in-flight or not completed. ISO-8601 in UTC. | [optional] 
**StartInvocationInput** | [**StartInvocationInput**](StartInvocationInput.md) |  | 
**CompleteInvocationInput** | Pointer to [**CompleteInvocationInput**](CompleteInvocationInput.md) |  | [optional] 

## Methods

### NewInvocationStatus

`func NewInvocationStatus(id string, triggerId string, subscriptionId string, type_ InvocationStatusType, created time.Time, startInvocationInput StartInvocationInput, ) *InvocationStatus`

NewInvocationStatus instantiates a new InvocationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvocationStatusWithDefaults

`func NewInvocationStatusWithDefaults() *InvocationStatus`

NewInvocationStatusWithDefaults instantiates a new InvocationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvocationStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvocationStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvocationStatus) SetId(v string)`

SetId sets Id field to given value.


### GetTriggerId

`func (o *InvocationStatus) GetTriggerId() string`

GetTriggerId returns the TriggerId field if non-nil, zero value otherwise.

### GetTriggerIdOk

`func (o *InvocationStatus) GetTriggerIdOk() (*string, bool)`

GetTriggerIdOk returns a tuple with the TriggerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerId

`func (o *InvocationStatus) SetTriggerId(v string)`

SetTriggerId sets TriggerId field to given value.


### GetSubscriptionId

`func (o *InvocationStatus) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *InvocationStatus) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *InvocationStatus) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetType

`func (o *InvocationStatus) GetType() InvocationStatusType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvocationStatus) GetTypeOk() (*InvocationStatusType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvocationStatus) SetType(v InvocationStatusType)`

SetType sets Type field to given value.


### GetCreated

`func (o *InvocationStatus) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *InvocationStatus) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *InvocationStatus) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCompleted

`func (o *InvocationStatus) GetCompleted() time.Time`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *InvocationStatus) GetCompletedOk() (*time.Time, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *InvocationStatus) SetCompleted(v time.Time)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *InvocationStatus) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetStartInvocationInput

`func (o *InvocationStatus) GetStartInvocationInput() StartInvocationInput`

GetStartInvocationInput returns the StartInvocationInput field if non-nil, zero value otherwise.

### GetStartInvocationInputOk

`func (o *InvocationStatus) GetStartInvocationInputOk() (*StartInvocationInput, bool)`

GetStartInvocationInputOk returns a tuple with the StartInvocationInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartInvocationInput

`func (o *InvocationStatus) SetStartInvocationInput(v StartInvocationInput)`

SetStartInvocationInput sets StartInvocationInput field to given value.


### GetCompleteInvocationInput

`func (o *InvocationStatus) GetCompleteInvocationInput() CompleteInvocationInput`

GetCompleteInvocationInput returns the CompleteInvocationInput field if non-nil, zero value otherwise.

### GetCompleteInvocationInputOk

`func (o *InvocationStatus) GetCompleteInvocationInputOk() (*CompleteInvocationInput, bool)`

GetCompleteInvocationInputOk returns a tuple with the CompleteInvocationInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteInvocationInput

`func (o *InvocationStatus) SetCompleteInvocationInput(v CompleteInvocationInput)`

SetCompleteInvocationInput sets CompleteInvocationInput field to given value.

### HasCompleteInvocationInput

`func (o *InvocationStatus) HasCompleteInvocationInput() bool`

HasCompleteInvocationInput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


