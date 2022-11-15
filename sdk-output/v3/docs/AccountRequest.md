# AccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | Unique ID of the account | [optional] 
**AttributeRequests** | Pointer to [**[]AttributeRequest**](AttributeRequest.md) |  | [optional] 
**Op** | Pointer to **string** | The operation that was performed | [optional] 
**ProvisioningTarget** | Pointer to [**Source1**](Source1.md) |  | [optional] 
**Result** | Pointer to [**AccountRequestResult**](AccountRequestResult.md) |  | [optional] 
**Source** | Pointer to [**Source1**](Source1.md) |  | [optional] 

## Methods

### NewAccountRequest

`func NewAccountRequest() *AccountRequest`

NewAccountRequest instantiates a new AccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountRequestWithDefaults

`func NewAccountRequestWithDefaults() *AccountRequest`

NewAccountRequestWithDefaults instantiates a new AccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AccountRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AccountRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAttributeRequests

`func (o *AccountRequest) GetAttributeRequests() []AttributeRequest`

GetAttributeRequests returns the AttributeRequests field if non-nil, zero value otherwise.

### GetAttributeRequestsOk

`func (o *AccountRequest) GetAttributeRequestsOk() (*[]AttributeRequest, bool)`

GetAttributeRequestsOk returns a tuple with the AttributeRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeRequests

`func (o *AccountRequest) SetAttributeRequests(v []AttributeRequest)`

SetAttributeRequests sets AttributeRequests field to given value.

### HasAttributeRequests

`func (o *AccountRequest) HasAttributeRequests() bool`

HasAttributeRequests returns a boolean if a field has been set.

### GetOp

`func (o *AccountRequest) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *AccountRequest) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *AccountRequest) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *AccountRequest) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetProvisioningTarget

`func (o *AccountRequest) GetProvisioningTarget() Source1`

GetProvisioningTarget returns the ProvisioningTarget field if non-nil, zero value otherwise.

### GetProvisioningTargetOk

`func (o *AccountRequest) GetProvisioningTargetOk() (*Source1, bool)`

GetProvisioningTargetOk returns a tuple with the ProvisioningTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningTarget

`func (o *AccountRequest) SetProvisioningTarget(v Source1)`

SetProvisioningTarget sets ProvisioningTarget field to given value.

### HasProvisioningTarget

`func (o *AccountRequest) HasProvisioningTarget() bool`

HasProvisioningTarget returns a boolean if a field has been set.

### GetResult

`func (o *AccountRequest) GetResult() AccountRequestResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AccountRequest) GetResultOk() (*AccountRequestResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AccountRequest) SetResult(v AccountRequestResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *AccountRequest) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetSource

`func (o *AccountRequest) GetSource() Source1`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AccountRequest) GetSourceOk() (*Source1, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AccountRequest) SetSource(v Source1)`

SetSource sets Source field to given value.

### HasSource

`func (o *AccountRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


