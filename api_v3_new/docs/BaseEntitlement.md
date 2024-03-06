# BaseEntitlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasPermissions** | Pointer to **bool** | Indicates whether the entitlement has permissions. | [optional] [default to false]
**Description** | Pointer to **string** | Entitlement&#39;s description. | [optional] 
**Attribute** | Pointer to **string** | Entitlement attribute&#39;s name. | [optional] 
**Value** | Pointer to **string** | Entitlement&#39;s value. | [optional] 
**Schema** | Pointer to **string** | Entitlement&#39;s schema. | [optional] 
**Privileged** | Pointer to **bool** | Indicates whether the entitlement is privileged. | [optional] [default to false]
**Id** | Pointer to **string** | Entitlement&#39;s ID. | [optional] 
**Name** | Pointer to **string** | Entitlement&#39;s name. | [optional] 

## Methods

### NewBaseEntitlement

`func NewBaseEntitlement() *BaseEntitlement`

NewBaseEntitlement instantiates a new BaseEntitlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseEntitlementWithDefaults

`func NewBaseEntitlementWithDefaults() *BaseEntitlement`

NewBaseEntitlementWithDefaults instantiates a new BaseEntitlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasPermissions

`func (o *BaseEntitlement) GetHasPermissions() bool`

GetHasPermissions returns the HasPermissions field if non-nil, zero value otherwise.

### GetHasPermissionsOk

`func (o *BaseEntitlement) GetHasPermissionsOk() (*bool, bool)`

GetHasPermissionsOk returns a tuple with the HasPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPermissions

`func (o *BaseEntitlement) SetHasPermissions(v bool)`

SetHasPermissions sets HasPermissions field to given value.

### HasHasPermissions

`func (o *BaseEntitlement) HasHasPermissions() bool`

HasHasPermissions returns a boolean if a field has been set.

### GetDescription

`func (o *BaseEntitlement) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BaseEntitlement) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BaseEntitlement) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BaseEntitlement) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAttribute

`func (o *BaseEntitlement) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *BaseEntitlement) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *BaseEntitlement) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *BaseEntitlement) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetValue

`func (o *BaseEntitlement) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BaseEntitlement) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BaseEntitlement) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *BaseEntitlement) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetSchema

`func (o *BaseEntitlement) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *BaseEntitlement) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *BaseEntitlement) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *BaseEntitlement) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetPrivileged

`func (o *BaseEntitlement) GetPrivileged() bool`

GetPrivileged returns the Privileged field if non-nil, zero value otherwise.

### GetPrivilegedOk

`func (o *BaseEntitlement) GetPrivilegedOk() (*bool, bool)`

GetPrivilegedOk returns a tuple with the Privileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileged

`func (o *BaseEntitlement) SetPrivileged(v bool)`

SetPrivileged sets Privileged field to given value.

### HasPrivileged

`func (o *BaseEntitlement) HasPrivileged() bool`

HasPrivileged returns a boolean if a field has been set.

### GetId

`func (o *BaseEntitlement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseEntitlement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseEntitlement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BaseEntitlement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BaseEntitlement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseEntitlement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseEntitlement) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BaseEntitlement) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


