# Entitlement1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The entitlement id | [optional] 
**Name** | Pointer to **string** | The entitlement name | [optional] 
**Created** | Pointer to **time.Time** | Time when the entitlement was created | [optional] 
**Modified** | Pointer to **time.Time** | Time when the entitlement was last modified | [optional] 
**Attribute** | Pointer to **NullableString** | The entitlement attribute name | [optional] 
**Value** | Pointer to **string** | The value of the entitlement | [optional] 
**SourceSchemaObjectType** | Pointer to **string** | The object type of the entitlement from the source schema | [optional] 
**Privileged** | Pointer to **bool** | True if the entitlement is privileged | [optional] [default to false]
**CloudGoverned** | Pointer to **bool** | True if the entitlement is cloud governed | [optional] [default to false]
**Description** | Pointer to **NullableString** | The description of the entitlement | [optional] 
**Requestable** | Pointer to **bool** | True if the entitlement is requestable | [optional] [default to false]
**Attributes** | Pointer to **map[string]interface{}** | A map of free-form key-value pairs from the source system | [optional] 
**Source** | Pointer to [**Entitlement1Source**](Entitlement1Source.md) |  | [optional] 
**Owner** | Pointer to [**Entitlement1Owner**](Entitlement1Owner.md) |  | [optional] 
**DirectPermissions** | Pointer to [**[]PermissionDto**](PermissionDto.md) |  | [optional] 
**Segments** | Pointer to **[]string** | List of IDs of segments, if any, to which this Entitlement is assigned. | [optional] 
**ManuallyUpdatedFields** | Pointer to [**Entitlement1ManuallyUpdatedFields**](Entitlement1ManuallyUpdatedFields.md) |  | [optional] 
**AccessModelMetadata** | Pointer to [**Entitlement1AccessModelMetadata**](Entitlement1AccessModelMetadata.md) |  | [optional] 

## Methods

### NewEntitlement1

`func NewEntitlement1() *Entitlement1`

NewEntitlement1 instantiates a new Entitlement1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlement1WithDefaults

`func NewEntitlement1WithDefaults() *Entitlement1`

NewEntitlement1WithDefaults instantiates a new Entitlement1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Entitlement1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Entitlement1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Entitlement1) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Entitlement1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Entitlement1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Entitlement1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Entitlement1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Entitlement1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreated

`func (o *Entitlement1) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Entitlement1) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Entitlement1) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Entitlement1) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *Entitlement1) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Entitlement1) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Entitlement1) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *Entitlement1) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetAttribute

`func (o *Entitlement1) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *Entitlement1) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *Entitlement1) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *Entitlement1) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### SetAttributeNil

`func (o *Entitlement1) SetAttributeNil(b bool)`

 SetAttributeNil sets the value for Attribute to be an explicit nil

### UnsetAttribute
`func (o *Entitlement1) UnsetAttribute()`

UnsetAttribute ensures that no value is present for Attribute, not even an explicit nil
### GetValue

`func (o *Entitlement1) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Entitlement1) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Entitlement1) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Entitlement1) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetSourceSchemaObjectType

`func (o *Entitlement1) GetSourceSchemaObjectType() string`

GetSourceSchemaObjectType returns the SourceSchemaObjectType field if non-nil, zero value otherwise.

### GetSourceSchemaObjectTypeOk

`func (o *Entitlement1) GetSourceSchemaObjectTypeOk() (*string, bool)`

GetSourceSchemaObjectTypeOk returns a tuple with the SourceSchemaObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSchemaObjectType

`func (o *Entitlement1) SetSourceSchemaObjectType(v string)`

SetSourceSchemaObjectType sets SourceSchemaObjectType field to given value.

### HasSourceSchemaObjectType

`func (o *Entitlement1) HasSourceSchemaObjectType() bool`

HasSourceSchemaObjectType returns a boolean if a field has been set.

### GetPrivileged

`func (o *Entitlement1) GetPrivileged() bool`

GetPrivileged returns the Privileged field if non-nil, zero value otherwise.

### GetPrivilegedOk

`func (o *Entitlement1) GetPrivilegedOk() (*bool, bool)`

GetPrivilegedOk returns a tuple with the Privileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileged

`func (o *Entitlement1) SetPrivileged(v bool)`

SetPrivileged sets Privileged field to given value.

### HasPrivileged

`func (o *Entitlement1) HasPrivileged() bool`

HasPrivileged returns a boolean if a field has been set.

### GetCloudGoverned

`func (o *Entitlement1) GetCloudGoverned() bool`

GetCloudGoverned returns the CloudGoverned field if non-nil, zero value otherwise.

### GetCloudGovernedOk

`func (o *Entitlement1) GetCloudGovernedOk() (*bool, bool)`

GetCloudGovernedOk returns a tuple with the CloudGoverned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudGoverned

`func (o *Entitlement1) SetCloudGoverned(v bool)`

SetCloudGoverned sets CloudGoverned field to given value.

### HasCloudGoverned

`func (o *Entitlement1) HasCloudGoverned() bool`

HasCloudGoverned returns a boolean if a field has been set.

### GetDescription

`func (o *Entitlement1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Entitlement1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Entitlement1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Entitlement1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Entitlement1) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Entitlement1) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRequestable

`func (o *Entitlement1) GetRequestable() bool`

GetRequestable returns the Requestable field if non-nil, zero value otherwise.

### GetRequestableOk

`func (o *Entitlement1) GetRequestableOk() (*bool, bool)`

GetRequestableOk returns a tuple with the Requestable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestable

`func (o *Entitlement1) SetRequestable(v bool)`

SetRequestable sets Requestable field to given value.

### HasRequestable

`func (o *Entitlement1) HasRequestable() bool`

HasRequestable returns a boolean if a field has been set.

### GetAttributes

`func (o *Entitlement1) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Entitlement1) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Entitlement1) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Entitlement1) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetSource

`func (o *Entitlement1) GetSource() Entitlement1Source`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Entitlement1) GetSourceOk() (*Entitlement1Source, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Entitlement1) SetSource(v Entitlement1Source)`

SetSource sets Source field to given value.

### HasSource

`func (o *Entitlement1) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetOwner

`func (o *Entitlement1) GetOwner() Entitlement1Owner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Entitlement1) GetOwnerOk() (*Entitlement1Owner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Entitlement1) SetOwner(v Entitlement1Owner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Entitlement1) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetDirectPermissions

`func (o *Entitlement1) GetDirectPermissions() []PermissionDto`

GetDirectPermissions returns the DirectPermissions field if non-nil, zero value otherwise.

### GetDirectPermissionsOk

`func (o *Entitlement1) GetDirectPermissionsOk() (*[]PermissionDto, bool)`

GetDirectPermissionsOk returns a tuple with the DirectPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectPermissions

`func (o *Entitlement1) SetDirectPermissions(v []PermissionDto)`

SetDirectPermissions sets DirectPermissions field to given value.

### HasDirectPermissions

`func (o *Entitlement1) HasDirectPermissions() bool`

HasDirectPermissions returns a boolean if a field has been set.

### GetSegments

`func (o *Entitlement1) GetSegments() []string`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *Entitlement1) GetSegmentsOk() (*[]string, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *Entitlement1) SetSegments(v []string)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *Entitlement1) HasSegments() bool`

HasSegments returns a boolean if a field has been set.

### SetSegmentsNil

`func (o *Entitlement1) SetSegmentsNil(b bool)`

 SetSegmentsNil sets the value for Segments to be an explicit nil

### UnsetSegments
`func (o *Entitlement1) UnsetSegments()`

UnsetSegments ensures that no value is present for Segments, not even an explicit nil
### GetManuallyUpdatedFields

`func (o *Entitlement1) GetManuallyUpdatedFields() Entitlement1ManuallyUpdatedFields`

GetManuallyUpdatedFields returns the ManuallyUpdatedFields field if non-nil, zero value otherwise.

### GetManuallyUpdatedFieldsOk

`func (o *Entitlement1) GetManuallyUpdatedFieldsOk() (*Entitlement1ManuallyUpdatedFields, bool)`

GetManuallyUpdatedFieldsOk returns a tuple with the ManuallyUpdatedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManuallyUpdatedFields

`func (o *Entitlement1) SetManuallyUpdatedFields(v Entitlement1ManuallyUpdatedFields)`

SetManuallyUpdatedFields sets ManuallyUpdatedFields field to given value.

### HasManuallyUpdatedFields

`func (o *Entitlement1) HasManuallyUpdatedFields() bool`

HasManuallyUpdatedFields returns a boolean if a field has been set.

### GetAccessModelMetadata

`func (o *Entitlement1) GetAccessModelMetadata() Entitlement1AccessModelMetadata`

GetAccessModelMetadata returns the AccessModelMetadata field if non-nil, zero value otherwise.

### GetAccessModelMetadataOk

`func (o *Entitlement1) GetAccessModelMetadataOk() (*Entitlement1AccessModelMetadata, bool)`

GetAccessModelMetadataOk returns a tuple with the AccessModelMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessModelMetadata

`func (o *Entitlement1) SetAccessModelMetadata(v Entitlement1AccessModelMetadata)`

SetAccessModelMetadata sets AccessModelMetadata field to given value.

### HasAccessModelMetadata

`func (o *Entitlement1) HasAccessModelMetadata() bool`

HasAccessModelMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


