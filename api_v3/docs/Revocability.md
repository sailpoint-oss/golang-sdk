# Revocability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommentsRequired** | Pointer to **NullableBool** | Whether the requester of the containing object must provide comments justifying the request | [optional] [default to false]
**DenialCommentsRequired** | Pointer to **NullableBool** | Whether an approver must provide comments when denying the request | [optional] [default to false]
**ApprovalSchemes** | Pointer to [**[]AccessProfileApprovalScheme**](AccessProfileApprovalScheme.md) | List describing the steps in approving the revocation request | [optional] 

## Methods

### NewRevocability

`func NewRevocability() *Revocability`

NewRevocability instantiates a new Revocability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevocabilityWithDefaults

`func NewRevocabilityWithDefaults() *Revocability`

NewRevocabilityWithDefaults instantiates a new Revocability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommentsRequired

`func (o *Revocability) GetCommentsRequired() bool`

GetCommentsRequired returns the CommentsRequired field if non-nil, zero value otherwise.

### GetCommentsRequiredOk

`func (o *Revocability) GetCommentsRequiredOk() (*bool, bool)`

GetCommentsRequiredOk returns a tuple with the CommentsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsRequired

`func (o *Revocability) SetCommentsRequired(v bool)`

SetCommentsRequired sets CommentsRequired field to given value.

### HasCommentsRequired

`func (o *Revocability) HasCommentsRequired() bool`

HasCommentsRequired returns a boolean if a field has been set.

### SetCommentsRequiredNil

`func (o *Revocability) SetCommentsRequiredNil(b bool)`

 SetCommentsRequiredNil sets the value for CommentsRequired to be an explicit nil

### UnsetCommentsRequired
`func (o *Revocability) UnsetCommentsRequired()`

UnsetCommentsRequired ensures that no value is present for CommentsRequired, not even an explicit nil
### GetDenialCommentsRequired

`func (o *Revocability) GetDenialCommentsRequired() bool`

GetDenialCommentsRequired returns the DenialCommentsRequired field if non-nil, zero value otherwise.

### GetDenialCommentsRequiredOk

`func (o *Revocability) GetDenialCommentsRequiredOk() (*bool, bool)`

GetDenialCommentsRequiredOk returns a tuple with the DenialCommentsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenialCommentsRequired

`func (o *Revocability) SetDenialCommentsRequired(v bool)`

SetDenialCommentsRequired sets DenialCommentsRequired field to given value.

### HasDenialCommentsRequired

`func (o *Revocability) HasDenialCommentsRequired() bool`

HasDenialCommentsRequired returns a boolean if a field has been set.

### SetDenialCommentsRequiredNil

`func (o *Revocability) SetDenialCommentsRequiredNil(b bool)`

 SetDenialCommentsRequiredNil sets the value for DenialCommentsRequired to be an explicit nil

### UnsetDenialCommentsRequired
`func (o *Revocability) UnsetDenialCommentsRequired()`

UnsetDenialCommentsRequired ensures that no value is present for DenialCommentsRequired, not even an explicit nil
### GetApprovalSchemes

`func (o *Revocability) GetApprovalSchemes() []AccessProfileApprovalScheme`

GetApprovalSchemes returns the ApprovalSchemes field if non-nil, zero value otherwise.

### GetApprovalSchemesOk

`func (o *Revocability) GetApprovalSchemesOk() (*[]AccessProfileApprovalScheme, bool)`

GetApprovalSchemesOk returns a tuple with the ApprovalSchemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalSchemes

`func (o *Revocability) SetApprovalSchemes(v []AccessProfileApprovalScheme)`

SetApprovalSchemes sets ApprovalSchemes field to given value.

### HasApprovalSchemes

`func (o *Revocability) HasApprovalSchemes() bool`

HasApprovalSchemes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


