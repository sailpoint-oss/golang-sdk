# SearchIdentityAccessInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Source** | Pointer to [**SearchIdentityAppsInnerSource**](SearchIdentityAppsInnerSource.md) |  | [optional] 
**Owner** | Pointer to [**SearchIdentityAccessInnerOwner**](SearchIdentityAccessInnerOwner.md) |  | [optional] 
**Privileged** | Pointer to **bool** |  | [optional] 
**Attribute** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewSearchIdentityAccessInner

`func NewSearchIdentityAccessInner() *SearchIdentityAccessInner`

NewSearchIdentityAccessInner instantiates a new SearchIdentityAccessInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchIdentityAccessInnerWithDefaults

`func NewSearchIdentityAccessInnerWithDefaults() *SearchIdentityAccessInner`

NewSearchIdentityAccessInnerWithDefaults instantiates a new SearchIdentityAccessInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SearchIdentityAccessInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchIdentityAccessInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchIdentityAccessInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SearchIdentityAccessInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *SearchIdentityAccessInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SearchIdentityAccessInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SearchIdentityAccessInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SearchIdentityAccessInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDisplayName

`func (o *SearchIdentityAccessInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SearchIdentityAccessInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SearchIdentityAccessInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SearchIdentityAccessInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetName

`func (o *SearchIdentityAccessInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchIdentityAccessInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchIdentityAccessInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SearchIdentityAccessInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *SearchIdentityAccessInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SearchIdentityAccessInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SearchIdentityAccessInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SearchIdentityAccessInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSource

`func (o *SearchIdentityAccessInner) GetSource() SearchIdentityAppsInnerSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SearchIdentityAccessInner) GetSourceOk() (*SearchIdentityAppsInnerSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SearchIdentityAccessInner) SetSource(v SearchIdentityAppsInnerSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *SearchIdentityAccessInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetOwner

`func (o *SearchIdentityAccessInner) GetOwner() SearchIdentityAccessInnerOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *SearchIdentityAccessInner) GetOwnerOk() (*SearchIdentityAccessInnerOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *SearchIdentityAccessInner) SetOwner(v SearchIdentityAccessInnerOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *SearchIdentityAccessInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPrivileged

`func (o *SearchIdentityAccessInner) GetPrivileged() bool`

GetPrivileged returns the Privileged field if non-nil, zero value otherwise.

### GetPrivilegedOk

`func (o *SearchIdentityAccessInner) GetPrivilegedOk() (*bool, bool)`

GetPrivilegedOk returns a tuple with the Privileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileged

`func (o *SearchIdentityAccessInner) SetPrivileged(v bool)`

SetPrivileged sets Privileged field to given value.

### HasPrivileged

`func (o *SearchIdentityAccessInner) HasPrivileged() bool`

HasPrivileged returns a boolean if a field has been set.

### GetAttribute

`func (o *SearchIdentityAccessInner) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *SearchIdentityAccessInner) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *SearchIdentityAccessInner) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *SearchIdentityAccessInner) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetValue

`func (o *SearchIdentityAccessInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SearchIdentityAccessInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SearchIdentityAccessInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SearchIdentityAccessInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDisabled

`func (o *SearchIdentityAccessInner) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *SearchIdentityAccessInner) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *SearchIdentityAccessInner) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *SearchIdentityAccessInner) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


