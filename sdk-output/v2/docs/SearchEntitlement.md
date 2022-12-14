# SearchEntitlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Source** | Pointer to [**LauncherAccount**](LauncherAccount.md) |  | [optional] 
**Privileged** | Pointer to **bool** |  | [optional] 
**Attribute** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**Synced** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSearchEntitlement

`func NewSearchEntitlement() *SearchEntitlement`

NewSearchEntitlement instantiates a new SearchEntitlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchEntitlementWithDefaults

`func NewSearchEntitlementWithDefaults() *SearchEntitlement`

NewSearchEntitlementWithDefaults instantiates a new SearchEntitlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SearchEntitlement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchEntitlement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchEntitlement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SearchEntitlement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SearchEntitlement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchEntitlement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchEntitlement) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SearchEntitlement) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *SearchEntitlement) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SearchEntitlement) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SearchEntitlement) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SearchEntitlement) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *SearchEntitlement) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SearchEntitlement) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SearchEntitlement) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SearchEntitlement) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSource

`func (o *SearchEntitlement) GetSource() LauncherAccount`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SearchEntitlement) GetSourceOk() (*LauncherAccount, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SearchEntitlement) SetSource(v LauncherAccount)`

SetSource sets Source field to given value.

### HasSource

`func (o *SearchEntitlement) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetPrivileged

`func (o *SearchEntitlement) GetPrivileged() bool`

GetPrivileged returns the Privileged field if non-nil, zero value otherwise.

### GetPrivilegedOk

`func (o *SearchEntitlement) GetPrivilegedOk() (*bool, bool)`

GetPrivilegedOk returns a tuple with the Privileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileged

`func (o *SearchEntitlement) SetPrivileged(v bool)`

SetPrivileged sets Privileged field to given value.

### HasPrivileged

`func (o *SearchEntitlement) HasPrivileged() bool`

HasPrivileged returns a boolean if a field has been set.

### GetAttribute

`func (o *SearchEntitlement) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *SearchEntitlement) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *SearchEntitlement) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *SearchEntitlement) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetValue

`func (o *SearchEntitlement) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SearchEntitlement) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SearchEntitlement) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SearchEntitlement) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetModified

`func (o *SearchEntitlement) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *SearchEntitlement) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *SearchEntitlement) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *SearchEntitlement) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetSynced

`func (o *SearchEntitlement) GetSynced() time.Time`

GetSynced returns the Synced field if non-nil, zero value otherwise.

### GetSyncedOk

`func (o *SearchEntitlement) GetSyncedOk() (*time.Time, bool)`

GetSyncedOk returns a tuple with the Synced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynced

`func (o *SearchEntitlement) SetSynced(v time.Time)`

SetSynced sets Synced field to given value.

### HasSynced

`func (o *SearchEntitlement) HasSynced() bool`

HasSynced returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


