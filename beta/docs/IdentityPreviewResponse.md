# IdentityPreviewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identity** | Pointer to [**BaseReferenceDto1**](BaseReferenceDto1.md) |  | [optional] 
**PreviewAttributes** | Pointer to [**[]IdentityAttributePreview**](IdentityAttributePreview.md) |  | [optional] 

## Methods

### NewIdentityPreviewResponse

`func NewIdentityPreviewResponse() *IdentityPreviewResponse`

NewIdentityPreviewResponse instantiates a new IdentityPreviewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityPreviewResponseWithDefaults

`func NewIdentityPreviewResponseWithDefaults() *IdentityPreviewResponse`

NewIdentityPreviewResponseWithDefaults instantiates a new IdentityPreviewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentity

`func (o *IdentityPreviewResponse) GetIdentity() BaseReferenceDto1`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *IdentityPreviewResponse) GetIdentityOk() (*BaseReferenceDto1, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *IdentityPreviewResponse) SetIdentity(v BaseReferenceDto1)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *IdentityPreviewResponse) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetPreviewAttributes

`func (o *IdentityPreviewResponse) GetPreviewAttributes() []IdentityAttributePreview`

GetPreviewAttributes returns the PreviewAttributes field if non-nil, zero value otherwise.

### GetPreviewAttributesOk

`func (o *IdentityPreviewResponse) GetPreviewAttributesOk() (*[]IdentityAttributePreview, bool)`

GetPreviewAttributesOk returns a tuple with the PreviewAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewAttributes

`func (o *IdentityPreviewResponse) SetPreviewAttributes(v []IdentityAttributePreview)`

SetPreviewAttributes sets PreviewAttributes field to given value.

### HasPreviewAttributes

`func (o *IdentityPreviewResponse) HasPreviewAttributes() bool`

HasPreviewAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


