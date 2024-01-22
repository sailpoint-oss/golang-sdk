# Entitlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The entitlement id | [optional] 
**Name** | Pointer to **string** | The entitlement name | [optional] 
**Attribute** | Pointer to **string** | The entitlement attribute name | [optional] 
**Value** | Pointer to **string** | The value of the entitlement | [optional] 
**SourceSchemaObjectType** | Pointer to **string** | The object type of the entitlement from the source schema | [optional] 
**Description** | Pointer to **string** | The description of the entitlement | [optional] 
**Privileged** | Pointer to **bool** | True if the entitlement is privileged | [optional] 
**CloudGoverned** | Pointer to **bool** | True if the entitlement is cloud governed | [optional] 
**Created** | Pointer to **time.Time** | Time when the entitlement was created | [optional] 
**Modified** | Pointer to **time.Time** | Time when the entitlement was last modified | [optional] 
**Source** | Pointer to [**EntitlementSource**](EntitlementSource.md) |  | [optional] 
**Attributes** | Pointer to **map[string]interface{}** | A map of free-form key-value pairs from the source system | [optional] 
**Segments** | Pointer to **[]string** | List of IDs of segments, if any, to which this Entitlement is assigned. | [optional] 
**DirectPermissions** | Pointer to [**[]PermissionDto**](PermissionDto.md) |  | [optional] 

## Methods

### NewEntitlement

`func NewEntitlement() *Entitlement`

NewEntitlement instantiates a new Entitlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementWithDefaults

`func NewEntitlementWithDefaults() *Entitlement`

NewEntitlementWithDefaults instantiates a new Entitlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Entitlement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Entitlement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Entitlement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Entitlement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Entitlement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Entitlement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Entitlement) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Entitlement) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAttribute

`func (o *Entitlement) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *Entitlement) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *Entitlement) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *Entitlement) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetValue

`func (o *Entitlement) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Entitlement) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Entitlement) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Entitlement) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetSourceSchemaObjectType

`func (o *Entitlement) GetSourceSchemaObjectType() string`

GetSourceSchemaObjectType returns the SourceSchemaObjectType field if non-nil, zero value otherwise.

### GetSourceSchemaObjectTypeOk

`func (o *Entitlement) GetSourceSchemaObjectTypeOk() (*string, bool)`

GetSourceSchemaObjectTypeOk returns a tuple with the SourceSchemaObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSchemaObjectType

`func (o *Entitlement) SetSourceSchemaObjectType(v string)`

SetSourceSchemaObjectType sets SourceSchemaObjectType field to given value.

### HasSourceSchemaObjectType

`func (o *Entitlement) HasSourceSchemaObjectType() bool`

HasSourceSchemaObjectType returns a boolean if a field has been set.

### GetDescription

`func (o *Entitlement) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Entitlement) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Entitlement) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Entitlement) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPrivileged

`func (o *Entitlement) GetPrivileged() bool`

GetPrivileged returns the Privileged field if non-nil, zero value otherwise.

### GetPrivilegedOk

`func (o *Entitlement) GetPrivilegedOk() (*bool, bool)`

GetPrivilegedOk returns a tuple with the Privileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileged

`func (o *Entitlement) SetPrivileged(v bool)`

SetPrivileged sets Privileged field to given value.

### HasPrivileged

`func (o *Entitlement) HasPrivileged() bool`

HasPrivileged returns a boolean if a field has been set.

### GetCloudGoverned

`func (o *Entitlement) GetCloudGoverned() bool`

GetCloudGoverned returns the CloudGoverned field if non-nil, zero value otherwise.

### GetCloudGovernedOk

`func (o *Entitlement) GetCloudGovernedOk() (*bool, bool)`

GetCloudGovernedOk returns a tuple with the CloudGoverned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudGoverned

`func (o *Entitlement) SetCloudGoverned(v bool)`

SetCloudGoverned sets CloudGoverned field to given value.

### HasCloudGoverned

`func (o *Entitlement) HasCloudGoverned() bool`

HasCloudGoverned returns a boolean if a field has been set.

### GetCreated

`func (o *Entitlement) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Entitlement) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Entitlement) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Entitlement) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *Entitlement) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Entitlement) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Entitlement) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *Entitlement) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetSource

`func (o *Entitlement) GetSource() EntitlementSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Entitlement) GetSourceOk() (*EntitlementSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Entitlement) SetSource(v EntitlementSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *Entitlement) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetAttributes

`func (o *Entitlement) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Entitlement) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Entitlement) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Entitlement) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetSegments

`func (o *Entitlement) GetSegments() []string`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *Entitlement) GetSegmentsOk() (*[]string, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *Entitlement) SetSegments(v []string)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *Entitlement) HasSegments() bool`

HasSegments returns a boolean if a field has been set.

### SetSegmentsNil

`func (o *Entitlement) SetSegmentsNil(b bool)`

 SetSegmentsNil sets the value for Segments to be an explicit nil

### UnsetSegments
`func (o *Entitlement) UnsetSegments()`

UnsetSegments ensures that no value is present for Segments, not even an explicit nil
### GetDirectPermissions

`func (o *Entitlement) GetDirectPermissions() []PermissionDto`

GetDirectPermissions returns the DirectPermissions field if non-nil, zero value otherwise.

### GetDirectPermissionsOk

`func (o *Entitlement) GetDirectPermissionsOk() (*[]PermissionDto, bool)`

GetDirectPermissionsOk returns a tuple with the DirectPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectPermissions

`func (o *Entitlement) SetDirectPermissions(v []PermissionDto)`

SetDirectPermissions sets DirectPermissions field to given value.

### HasDirectPermissions

`func (o *Entitlement) HasDirectPermissions() bool`

HasDirectPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


