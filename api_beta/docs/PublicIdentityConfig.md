# PublicIdentityConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**[]PublicIdentityAttributeConfig**](PublicIdentityAttributeConfig.md) |  | [optional] 
**ModifiedBy** | Pointer to [**NullableIdentityReference**](IdentityReference.md) |  | [optional] 
**Modified** | Pointer to **time.Time** | the date/time of the modification | [optional] 

## Methods

### NewPublicIdentityConfig

`func NewPublicIdentityConfig() *PublicIdentityConfig`

NewPublicIdentityConfig instantiates a new PublicIdentityConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicIdentityConfigWithDefaults

`func NewPublicIdentityConfigWithDefaults() *PublicIdentityConfig`

NewPublicIdentityConfigWithDefaults instantiates a new PublicIdentityConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *PublicIdentityConfig) GetAttributes() []PublicIdentityAttributeConfig`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PublicIdentityConfig) GetAttributesOk() (*[]PublicIdentityAttributeConfig, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PublicIdentityConfig) SetAttributes(v []PublicIdentityAttributeConfig)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PublicIdentityConfig) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetModifiedBy

`func (o *PublicIdentityConfig) GetModifiedBy() IdentityReference`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *PublicIdentityConfig) GetModifiedByOk() (*IdentityReference, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *PublicIdentityConfig) SetModifiedBy(v IdentityReference)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *PublicIdentityConfig) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### SetModifiedByNil

`func (o *PublicIdentityConfig) SetModifiedByNil(b bool)`

 SetModifiedByNil sets the value for ModifiedBy to be an explicit nil

### UnsetModifiedBy
`func (o *PublicIdentityConfig) UnsetModifiedBy()`

UnsetModifiedBy ensures that no value is present for ModifiedBy, not even an explicit nil
### GetModified

`func (o *PublicIdentityConfig) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *PublicIdentityConfig) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *PublicIdentityConfig) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *PublicIdentityConfig) HasModified() bool`

HasModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


