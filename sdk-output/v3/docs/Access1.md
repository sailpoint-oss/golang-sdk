# Access1

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
**Privileged** | Pointer to **bool** |  | [optional] 
**Attribute** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Standalone** | Pointer to **bool** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewAccess1

`func NewAccess1() *Access1`

NewAccess1 instantiates a new Access1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccess1WithDefaults

`func NewAccess1WithDefaults() *Access1`

NewAccess1WithDefaults instantiates a new Access1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Access1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Access1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Access1) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Access1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Access1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Access1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Access1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Access1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *Access1) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Access1) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Access1) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Access1) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetType

`func (o *Access1) GetType() DtoType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Access1) GetTypeOk() (*DtoType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Access1) SetType(v DtoType)`

SetType sets Type field to given value.

### HasType

`func (o *Access1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *Access1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Access1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Access1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Access1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Access1) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Access1) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSource

`func (o *Access1) GetSource() Reference1`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Access1) GetSourceOk() (*Reference1, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Access1) SetSource(v Reference1)`

SetSource sets Source field to given value.

### HasSource

`func (o *Access1) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetOwner

`func (o *Access1) GetOwner() DisplayReference`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Access1) GetOwnerOk() (*DisplayReference, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Access1) SetOwner(v DisplayReference)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Access1) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetRevocable

`func (o *Access1) GetRevocable() bool`

GetRevocable returns the Revocable field if non-nil, zero value otherwise.

### GetRevocableOk

`func (o *Access1) GetRevocableOk() (*bool, bool)`

GetRevocableOk returns a tuple with the Revocable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocable

`func (o *Access1) SetRevocable(v bool)`

SetRevocable sets Revocable field to given value.

### HasRevocable

`func (o *Access1) HasRevocable() bool`

HasRevocable returns a boolean if a field has been set.

### GetPrivileged

`func (o *Access1) GetPrivileged() bool`

GetPrivileged returns the Privileged field if non-nil, zero value otherwise.

### GetPrivilegedOk

`func (o *Access1) GetPrivilegedOk() (*bool, bool)`

GetPrivilegedOk returns a tuple with the Privileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileged

`func (o *Access1) SetPrivileged(v bool)`

SetPrivileged sets Privileged field to given value.

### HasPrivileged

`func (o *Access1) HasPrivileged() bool`

HasPrivileged returns a boolean if a field has been set.

### GetAttribute

`func (o *Access1) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *Access1) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *Access1) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *Access1) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetValue

`func (o *Access1) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Access1) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Access1) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Access1) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetStandalone

`func (o *Access1) GetStandalone() bool`

GetStandalone returns the Standalone field if non-nil, zero value otherwise.

### GetStandaloneOk

`func (o *Access1) GetStandaloneOk() (*bool, bool)`

GetStandaloneOk returns a tuple with the Standalone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandalone

`func (o *Access1) SetStandalone(v bool)`

SetStandalone sets Standalone field to given value.

### HasStandalone

`func (o *Access1) HasStandalone() bool`

HasStandalone returns a boolean if a field has been set.

### GetDisabled

`func (o *Access1) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *Access1) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *Access1) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *Access1) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


