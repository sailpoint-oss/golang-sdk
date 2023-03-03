# ApprovalItemDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the approval item | [optional] 
**Account** | Pointer to **string** | The account referenced by the approval item | [optional] 
**Application** | Pointer to **string** | The name the application/source | [optional] 
**AttributeName** | Pointer to **string** | The name of the attribute | [optional] 
**AttributeOperation** | Pointer to **string** | The operation of the attribute | [optional] 
**AttributeValue** | Pointer to **string** | The value of the attribute | [optional] 
**State** | Pointer to [**WorkItemState**](WorkItemState.md) |  | [optional] 

## Methods

### NewApprovalItemDetails

`func NewApprovalItemDetails() *ApprovalItemDetails`

NewApprovalItemDetails instantiates a new ApprovalItemDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalItemDetailsWithDefaults

`func NewApprovalItemDetailsWithDefaults() *ApprovalItemDetails`

NewApprovalItemDetailsWithDefaults instantiates a new ApprovalItemDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApprovalItemDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApprovalItemDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApprovalItemDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApprovalItemDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *ApprovalItemDetails) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ApprovalItemDetails) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ApprovalItemDetails) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ApprovalItemDetails) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetApplication

`func (o *ApprovalItemDetails) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *ApprovalItemDetails) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *ApprovalItemDetails) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *ApprovalItemDetails) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetAttributeName

`func (o *ApprovalItemDetails) GetAttributeName() string`

GetAttributeName returns the AttributeName field if non-nil, zero value otherwise.

### GetAttributeNameOk

`func (o *ApprovalItemDetails) GetAttributeNameOk() (*string, bool)`

GetAttributeNameOk returns a tuple with the AttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeName

`func (o *ApprovalItemDetails) SetAttributeName(v string)`

SetAttributeName sets AttributeName field to given value.

### HasAttributeName

`func (o *ApprovalItemDetails) HasAttributeName() bool`

HasAttributeName returns a boolean if a field has been set.

### GetAttributeOperation

`func (o *ApprovalItemDetails) GetAttributeOperation() string`

GetAttributeOperation returns the AttributeOperation field if non-nil, zero value otherwise.

### GetAttributeOperationOk

`func (o *ApprovalItemDetails) GetAttributeOperationOk() (*string, bool)`

GetAttributeOperationOk returns a tuple with the AttributeOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeOperation

`func (o *ApprovalItemDetails) SetAttributeOperation(v string)`

SetAttributeOperation sets AttributeOperation field to given value.

### HasAttributeOperation

`func (o *ApprovalItemDetails) HasAttributeOperation() bool`

HasAttributeOperation returns a boolean if a field has been set.

### GetAttributeValue

`func (o *ApprovalItemDetails) GetAttributeValue() string`

GetAttributeValue returns the AttributeValue field if non-nil, zero value otherwise.

### GetAttributeValueOk

`func (o *ApprovalItemDetails) GetAttributeValueOk() (*string, bool)`

GetAttributeValueOk returns a tuple with the AttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeValue

`func (o *ApprovalItemDetails) SetAttributeValue(v string)`

SetAttributeValue sets AttributeValue field to given value.

### HasAttributeValue

`func (o *ApprovalItemDetails) HasAttributeValue() bool`

HasAttributeValue returns a boolean if a field has been set.

### GetState

`func (o *ApprovalItemDetails) GetState() WorkItemState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApprovalItemDetails) GetStateOk() (*WorkItemState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ApprovalItemDetails) SetState(v WorkItemState)`

SetState sets State field to given value.

### HasState

`func (o *ApprovalItemDetails) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


