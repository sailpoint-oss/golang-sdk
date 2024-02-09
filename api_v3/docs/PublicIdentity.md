# PublicIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identity id | [optional] 
**Name** | Pointer to **string** | Human-readable display name of identity. | [optional] 
**Alias** | Pointer to **string** | Alternate unique identifier for the identity. | [optional] 
**Email** | Pointer to **NullableString** | Email address of identity. | [optional] 
**Status** | Pointer to **NullableString** | The lifecycle status for the identity | [optional] 
**Manager** | Pointer to [**NullableIdentityReference**](IdentityReference.md) |  | [optional] 
**Attributes** | Pointer to [**[]IdentityAttribute1**](IdentityAttribute1.md) | The public identity attributes of the identity | [optional] 

## Methods

### NewPublicIdentity

`func NewPublicIdentity() *PublicIdentity`

NewPublicIdentity instantiates a new PublicIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicIdentityWithDefaults

`func NewPublicIdentityWithDefaults() *PublicIdentity`

NewPublicIdentityWithDefaults instantiates a new PublicIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PublicIdentity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicIdentity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicIdentity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PublicIdentity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PublicIdentity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicIdentity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicIdentity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PublicIdentity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAlias

`func (o *PublicIdentity) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *PublicIdentity) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *PublicIdentity) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *PublicIdentity) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetEmail

`func (o *PublicIdentity) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PublicIdentity) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PublicIdentity) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PublicIdentity) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *PublicIdentity) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *PublicIdentity) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetStatus

`func (o *PublicIdentity) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PublicIdentity) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PublicIdentity) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PublicIdentity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *PublicIdentity) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *PublicIdentity) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetManager

`func (o *PublicIdentity) GetManager() IdentityReference`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *PublicIdentity) GetManagerOk() (*IdentityReference, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *PublicIdentity) SetManager(v IdentityReference)`

SetManager sets Manager field to given value.

### HasManager

`func (o *PublicIdentity) HasManager() bool`

HasManager returns a boolean if a field has been set.

### SetManagerNil

`func (o *PublicIdentity) SetManagerNil(b bool)`

 SetManagerNil sets the value for Manager to be an explicit nil

### UnsetManager
`func (o *PublicIdentity) UnsetManager()`

UnsetManager ensures that no value is present for Manager, not even an explicit nil
### GetAttributes

`func (o *PublicIdentity) GetAttributes() []IdentityAttribute1`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PublicIdentity) GetAttributesOk() (*[]IdentityAttribute1, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PublicIdentity) SetAttributes(v []IdentityAttribute1)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PublicIdentity) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


