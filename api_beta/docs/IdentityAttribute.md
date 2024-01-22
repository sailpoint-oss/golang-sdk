# IdentityAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The technical name of the identity attribute | [optional] 
**DisplayName** | Pointer to **string** | The business-friendly name of the identity attribute | [optional] 
**Standard** | Pointer to **bool** | Shows if the attribute is &#39;standard&#39; or default | [optional] [default to false]
**Type** | Pointer to **string** | The type of the identity attribute | [optional] 
**Multi** | Pointer to **bool** | Shows if the identity attribute is multi-valued | [optional] [default to false]
**Searchable** | Pointer to **bool** | Shows if the identity attribute is searchable | [optional] [default to false]
**System** | Pointer to **bool** | Shows this is &#39;system&#39; identity attribute that does not have a source and is not configurable. | [optional] [default to false]
**Sources** | Pointer to [**[]Source1**](Source1.md) | List of sources for an attribute, this specifies how the value of the rule is derived | [optional] 

## Methods

### NewIdentityAttribute

`func NewIdentityAttribute() *IdentityAttribute`

NewIdentityAttribute instantiates a new IdentityAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityAttributeWithDefaults

`func NewIdentityAttributeWithDefaults() *IdentityAttribute`

NewIdentityAttributeWithDefaults instantiates a new IdentityAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IdentityAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityAttribute) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityAttribute) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *IdentityAttribute) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IdentityAttribute) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IdentityAttribute) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IdentityAttribute) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetStandard

`func (o *IdentityAttribute) GetStandard() bool`

GetStandard returns the Standard field if non-nil, zero value otherwise.

### GetStandardOk

`func (o *IdentityAttribute) GetStandardOk() (*bool, bool)`

GetStandardOk returns a tuple with the Standard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandard

`func (o *IdentityAttribute) SetStandard(v bool)`

SetStandard sets Standard field to given value.

### HasStandard

`func (o *IdentityAttribute) HasStandard() bool`

HasStandard returns a boolean if a field has been set.

### GetType

`func (o *IdentityAttribute) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdentityAttribute) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdentityAttribute) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IdentityAttribute) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMulti

`func (o *IdentityAttribute) GetMulti() bool`

GetMulti returns the Multi field if non-nil, zero value otherwise.

### GetMultiOk

`func (o *IdentityAttribute) GetMultiOk() (*bool, bool)`

GetMultiOk returns a tuple with the Multi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulti

`func (o *IdentityAttribute) SetMulti(v bool)`

SetMulti sets Multi field to given value.

### HasMulti

`func (o *IdentityAttribute) HasMulti() bool`

HasMulti returns a boolean if a field has been set.

### GetSearchable

`func (o *IdentityAttribute) GetSearchable() bool`

GetSearchable returns the Searchable field if non-nil, zero value otherwise.

### GetSearchableOk

`func (o *IdentityAttribute) GetSearchableOk() (*bool, bool)`

GetSearchableOk returns a tuple with the Searchable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchable

`func (o *IdentityAttribute) SetSearchable(v bool)`

SetSearchable sets Searchable field to given value.

### HasSearchable

`func (o *IdentityAttribute) HasSearchable() bool`

HasSearchable returns a boolean if a field has been set.

### GetSystem

`func (o *IdentityAttribute) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *IdentityAttribute) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *IdentityAttribute) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *IdentityAttribute) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetSources

`func (o *IdentityAttribute) GetSources() []Source1`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *IdentityAttribute) GetSourcesOk() (*[]Source1, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *IdentityAttribute) SetSources(v []Source1)`

SetSources sets Sources field to given value.

### HasSources

`func (o *IdentityAttribute) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


