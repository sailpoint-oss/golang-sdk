# Role1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**DtoType**](DtoType.md) |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Owner** | Pointer to [**DisplayReference**](DisplayReference.md) |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Revocable** | Pointer to **bool** |  | [optional] 

## Methods

### NewRole1

`func NewRole1() *Role1`

NewRole1 instantiates a new Role1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRole1WithDefaults

`func NewRole1WithDefaults() *Role1`

NewRole1WithDefaults instantiates a new Role1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Role1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Role1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Role1) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Role1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Role1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Role1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Role1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Role1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *Role1) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Role1) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Role1) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Role1) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetType

`func (o *Role1) GetType() DtoType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Role1) GetTypeOk() (*DtoType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Role1) SetType(v DtoType)`

SetType sets Type field to given value.

### HasType

`func (o *Role1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *Role1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Role1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Role1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Role1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Role1) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Role1) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetOwner

`func (o *Role1) GetOwner() DisplayReference`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Role1) GetOwnerOk() (*DisplayReference, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Role1) SetOwner(v DisplayReference)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Role1) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetDisabled

`func (o *Role1) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *Role1) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *Role1) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *Role1) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetRevocable

`func (o *Role1) GetRevocable() bool`

GetRevocable returns the Revocable field if non-nil, zero value otherwise.

### GetRevocableOk

`func (o *Role1) GetRevocableOk() (*bool, bool)`

GetRevocableOk returns a tuple with the Revocable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocable

`func (o *Role1) SetRevocable(v bool)`

SetRevocable sets Revocable field to given value.

### HasRevocable

`func (o *Role1) HasRevocable() bool`

HasRevocable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


