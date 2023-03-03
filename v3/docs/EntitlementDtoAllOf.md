# EntitlementDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attribute** | Pointer to **string** | Name of the entitlement attribute | [optional] 
**Value** | Pointer to **string** | Raw value of the entitlement | [optional] 
**Description** | Pointer to **string** | Entitlment description | [optional] 
**Attributes** | Pointer to **map[string]interface{}** | Entitlement attributes | [optional] 
**SourceSchemaObjectType** | Pointer to **string** | Schema objectType on the given application that maps to an Account Group | [optional] 
**Privileged** | Pointer to **bool** | Determines if this Entitlement is privileged. | [optional] 
**CloudGoverned** | Pointer to **bool** | Determines if this Entitlement is goverened in the cloud. | [optional] 
**Source** | Pointer to [**BaseReferenceDto**](BaseReferenceDto.md) |  | [optional] 

## Methods

### NewEntitlementDtoAllOf

`func NewEntitlementDtoAllOf() *EntitlementDtoAllOf`

NewEntitlementDtoAllOf instantiates a new EntitlementDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementDtoAllOfWithDefaults

`func NewEntitlementDtoAllOfWithDefaults() *EntitlementDtoAllOf`

NewEntitlementDtoAllOfWithDefaults instantiates a new EntitlementDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttribute

`func (o *EntitlementDtoAllOf) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *EntitlementDtoAllOf) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *EntitlementDtoAllOf) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *EntitlementDtoAllOf) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetValue

`func (o *EntitlementDtoAllOf) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EntitlementDtoAllOf) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EntitlementDtoAllOf) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *EntitlementDtoAllOf) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDescription

`func (o *EntitlementDtoAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntitlementDtoAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntitlementDtoAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntitlementDtoAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAttributes

`func (o *EntitlementDtoAllOf) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EntitlementDtoAllOf) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EntitlementDtoAllOf) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *EntitlementDtoAllOf) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetSourceSchemaObjectType

`func (o *EntitlementDtoAllOf) GetSourceSchemaObjectType() string`

GetSourceSchemaObjectType returns the SourceSchemaObjectType field if non-nil, zero value otherwise.

### GetSourceSchemaObjectTypeOk

`func (o *EntitlementDtoAllOf) GetSourceSchemaObjectTypeOk() (*string, bool)`

GetSourceSchemaObjectTypeOk returns a tuple with the SourceSchemaObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSchemaObjectType

`func (o *EntitlementDtoAllOf) SetSourceSchemaObjectType(v string)`

SetSourceSchemaObjectType sets SourceSchemaObjectType field to given value.

### HasSourceSchemaObjectType

`func (o *EntitlementDtoAllOf) HasSourceSchemaObjectType() bool`

HasSourceSchemaObjectType returns a boolean if a field has been set.

### GetPrivileged

`func (o *EntitlementDtoAllOf) GetPrivileged() bool`

GetPrivileged returns the Privileged field if non-nil, zero value otherwise.

### GetPrivilegedOk

`func (o *EntitlementDtoAllOf) GetPrivilegedOk() (*bool, bool)`

GetPrivilegedOk returns a tuple with the Privileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileged

`func (o *EntitlementDtoAllOf) SetPrivileged(v bool)`

SetPrivileged sets Privileged field to given value.

### HasPrivileged

`func (o *EntitlementDtoAllOf) HasPrivileged() bool`

HasPrivileged returns a boolean if a field has been set.

### GetCloudGoverned

`func (o *EntitlementDtoAllOf) GetCloudGoverned() bool`

GetCloudGoverned returns the CloudGoverned field if non-nil, zero value otherwise.

### GetCloudGovernedOk

`func (o *EntitlementDtoAllOf) GetCloudGovernedOk() (*bool, bool)`

GetCloudGovernedOk returns a tuple with the CloudGoverned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudGoverned

`func (o *EntitlementDtoAllOf) SetCloudGoverned(v bool)`

SetCloudGoverned sets CloudGoverned field to given value.

### HasCloudGoverned

`func (o *EntitlementDtoAllOf) HasCloudGoverned() bool`

HasCloudGoverned returns a boolean if a field has been set.

### GetSource

`func (o *EntitlementDtoAllOf) GetSource() BaseReferenceDto`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EntitlementDtoAllOf) GetSourceOk() (*BaseReferenceDto, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EntitlementDtoAllOf) SetSource(v BaseReferenceDto)`

SetSource sets Source field to given value.

### HasSource

`func (o *EntitlementDtoAllOf) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


