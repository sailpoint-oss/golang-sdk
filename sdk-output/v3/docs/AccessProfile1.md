# AccessProfile1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**DtoType**](DtoType.md) |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Source** | Pointer to [**Reference1**](Reference1.md) |  | [optional] 
**Owner** | Pointer to [**DisplayReference**](DisplayReference.md) |  | [optional] 
**Revocable** | Pointer to **bool** |  | [optional] 

## Methods

### NewAccessProfile1

`func NewAccessProfile1() *AccessProfile1`

NewAccessProfile1 instantiates a new AccessProfile1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessProfile1WithDefaults

`func NewAccessProfile1WithDefaults() *AccessProfile1`

NewAccessProfile1WithDefaults instantiates a new AccessProfile1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccessProfile1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessProfile1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessProfile1) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccessProfile1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AccessProfile1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessProfile1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessProfile1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccessProfile1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *AccessProfile1) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AccessProfile1) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AccessProfile1) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AccessProfile1) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetType

`func (o *AccessProfile1) GetType() DtoType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccessProfile1) GetTypeOk() (*DtoType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccessProfile1) SetType(v DtoType)`

SetType sets Type field to given value.

### HasType

`func (o *AccessProfile1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *AccessProfile1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccessProfile1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccessProfile1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccessProfile1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AccessProfile1) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AccessProfile1) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSource

`func (o *AccessProfile1) GetSource() Reference1`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AccessProfile1) GetSourceOk() (*Reference1, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AccessProfile1) SetSource(v Reference1)`

SetSource sets Source field to given value.

### HasSource

`func (o *AccessProfile1) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetOwner

`func (o *AccessProfile1) GetOwner() DisplayReference`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AccessProfile1) GetOwnerOk() (*DisplayReference, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AccessProfile1) SetOwner(v DisplayReference)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AccessProfile1) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetRevocable

`func (o *AccessProfile1) GetRevocable() bool`

GetRevocable returns the Revocable field if non-nil, zero value otherwise.

### GetRevocableOk

`func (o *AccessProfile1) GetRevocableOk() (*bool, bool)`

GetRevocableOk returns a tuple with the Revocable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocable

`func (o *AccessProfile1) SetRevocable(v bool)`

SetRevocable sets Revocable field to given value.

### HasRevocable

`func (o *AccessProfile1) HasRevocable() bool`

HasRevocable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


