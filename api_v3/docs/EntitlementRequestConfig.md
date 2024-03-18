# EntitlementRequestConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowEntitlementRequest** | Pointer to **bool** | Flag for allowing entitlement request. | [optional] 
**RequestCommentsRequired** | Pointer to **bool** | Flag for requiring comments while submitting an entitlement request. | [optional] [default to false]
**DeniedCommentsRequired** | Pointer to **bool** | Flag for requiring comments while rejecting an entitlement request. | [optional] [default to false]
**GrantRequestApprovalSchemes** | Pointer to **NullableString** | Approval schemes for granting entitlement request. This can be empty if no approval is needed. Multiple schemes must be comma-separated. The valid schemes are \&quot;entitlementOwner\&quot;, \&quot;sourceOwner\&quot;, \&quot;manager\&quot; and \&quot;workgroup:{id}\&quot;. Multiple workgroups (governance groups) can be used.  | [optional] [default to "sourceOwner"]

## Methods

### NewEntitlementRequestConfig

`func NewEntitlementRequestConfig() *EntitlementRequestConfig`

NewEntitlementRequestConfig instantiates a new EntitlementRequestConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementRequestConfigWithDefaults

`func NewEntitlementRequestConfigWithDefaults() *EntitlementRequestConfig`

NewEntitlementRequestConfigWithDefaults instantiates a new EntitlementRequestConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowEntitlementRequest

`func (o *EntitlementRequestConfig) GetAllowEntitlementRequest() bool`

GetAllowEntitlementRequest returns the AllowEntitlementRequest field if non-nil, zero value otherwise.

### GetAllowEntitlementRequestOk

`func (o *EntitlementRequestConfig) GetAllowEntitlementRequestOk() (*bool, bool)`

GetAllowEntitlementRequestOk returns a tuple with the AllowEntitlementRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEntitlementRequest

`func (o *EntitlementRequestConfig) SetAllowEntitlementRequest(v bool)`

SetAllowEntitlementRequest sets AllowEntitlementRequest field to given value.

### HasAllowEntitlementRequest

`func (o *EntitlementRequestConfig) HasAllowEntitlementRequest() bool`

HasAllowEntitlementRequest returns a boolean if a field has been set.

### GetRequestCommentsRequired

`func (o *EntitlementRequestConfig) GetRequestCommentsRequired() bool`

GetRequestCommentsRequired returns the RequestCommentsRequired field if non-nil, zero value otherwise.

### GetRequestCommentsRequiredOk

`func (o *EntitlementRequestConfig) GetRequestCommentsRequiredOk() (*bool, bool)`

GetRequestCommentsRequiredOk returns a tuple with the RequestCommentsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCommentsRequired

`func (o *EntitlementRequestConfig) SetRequestCommentsRequired(v bool)`

SetRequestCommentsRequired sets RequestCommentsRequired field to given value.

### HasRequestCommentsRequired

`func (o *EntitlementRequestConfig) HasRequestCommentsRequired() bool`

HasRequestCommentsRequired returns a boolean if a field has been set.

### GetDeniedCommentsRequired

`func (o *EntitlementRequestConfig) GetDeniedCommentsRequired() bool`

GetDeniedCommentsRequired returns the DeniedCommentsRequired field if non-nil, zero value otherwise.

### GetDeniedCommentsRequiredOk

`func (o *EntitlementRequestConfig) GetDeniedCommentsRequiredOk() (*bool, bool)`

GetDeniedCommentsRequiredOk returns a tuple with the DeniedCommentsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedCommentsRequired

`func (o *EntitlementRequestConfig) SetDeniedCommentsRequired(v bool)`

SetDeniedCommentsRequired sets DeniedCommentsRequired field to given value.

### HasDeniedCommentsRequired

`func (o *EntitlementRequestConfig) HasDeniedCommentsRequired() bool`

HasDeniedCommentsRequired returns a boolean if a field has been set.

### GetGrantRequestApprovalSchemes

`func (o *EntitlementRequestConfig) GetGrantRequestApprovalSchemes() string`

GetGrantRequestApprovalSchemes returns the GrantRequestApprovalSchemes field if non-nil, zero value otherwise.

### GetGrantRequestApprovalSchemesOk

`func (o *EntitlementRequestConfig) GetGrantRequestApprovalSchemesOk() (*string, bool)`

GetGrantRequestApprovalSchemesOk returns a tuple with the GrantRequestApprovalSchemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantRequestApprovalSchemes

`func (o *EntitlementRequestConfig) SetGrantRequestApprovalSchemes(v string)`

SetGrantRequestApprovalSchemes sets GrantRequestApprovalSchemes field to given value.

### HasGrantRequestApprovalSchemes

`func (o *EntitlementRequestConfig) HasGrantRequestApprovalSchemes() bool`

HasGrantRequestApprovalSchemes returns a boolean if a field has been set.

### SetGrantRequestApprovalSchemesNil

`func (o *EntitlementRequestConfig) SetGrantRequestApprovalSchemesNil(b bool)`

 SetGrantRequestApprovalSchemesNil sets the value for GrantRequestApprovalSchemes to be an explicit nil

### UnsetGrantRequestApprovalSchemes
`func (o *EntitlementRequestConfig) UnsetGrantRequestApprovalSchemes()`

UnsetGrantRequestApprovalSchemes ensures that no value is present for GrantRequestApprovalSchemes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


