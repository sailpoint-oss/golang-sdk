# ApprovalItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The approval item&#39;s ID | [optional] 
**Account** | Pointer to **string** | The account referenced by the approval item | [optional] 
**Application** | Pointer to **string** | The name of the application/source | [optional] 
**Name** | Pointer to **string** | The attribute&#39;s name | [optional] 
**Operation** | Pointer to **string** | The attribute&#39;s operation | [optional] 
**Value** | Pointer to **string** | The attribute&#39;s value | [optional] 
**State** | Pointer to [**WorkItemState**](WorkItemState.md) |  | [optional] 

## Methods

### NewApprovalItems

`func NewApprovalItems() *ApprovalItems`

NewApprovalItems instantiates a new ApprovalItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalItemsWithDefaults

`func NewApprovalItemsWithDefaults() *ApprovalItems`

NewApprovalItemsWithDefaults instantiates a new ApprovalItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApprovalItems) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApprovalItems) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApprovalItems) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApprovalItems) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *ApprovalItems) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ApprovalItems) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ApprovalItems) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ApprovalItems) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetApplication

`func (o *ApprovalItems) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *ApprovalItems) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *ApprovalItems) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *ApprovalItems) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetName

`func (o *ApprovalItems) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApprovalItems) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApprovalItems) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApprovalItems) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperation

`func (o *ApprovalItems) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ApprovalItems) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ApprovalItems) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *ApprovalItems) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetValue

`func (o *ApprovalItems) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ApprovalItems) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ApprovalItems) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ApprovalItems) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetState

`func (o *ApprovalItems) GetState() WorkItemState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApprovalItems) GetStateOk() (*WorkItemState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ApprovalItems) SetState(v WorkItemState)`

SetState sets State field to given value.

### HasState

`func (o *ApprovalItems) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


