# AccessProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**SourceId** | Pointer to **string** |  | [optional] 
**Entitlements** | Pointer to **[]string** |  | [optional] 
**ApprovalSchemes** | Pointer to **string** | Comma-separated list of approval schemes. Each approval scheme is one of - manager - appOwner - sourceOwner - accessProfileOwner - workgroup:\\&lt;workgroupId&gt; | [optional] 
**RevokeRequestApprovalSchemes** | Pointer to **string** | Comma-separated list of revoke request approval schemes. Each approval scheme is one of - manager - sourceOwner - accessProfileOwner - workgroup:\\&lt;workgroupId&gt; | [optional] 
**RequestCommentsRequired** | Pointer to **bool** |  | [optional] 
**DeniedCommentsRequired** | Pointer to **bool** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewAccessProfile

`func NewAccessProfile() *AccessProfile`

NewAccessProfile instantiates a new AccessProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessProfileWithDefaults

`func NewAccessProfileWithDefaults() *AccessProfile`

NewAccessProfileWithDefaults instantiates a new AccessProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AccessProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccessProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AccessProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccessProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccessProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccessProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwnerId

`func (o *AccessProfile) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *AccessProfile) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *AccessProfile) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *AccessProfile) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetSourceId

`func (o *AccessProfile) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *AccessProfile) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *AccessProfile) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *AccessProfile) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetEntitlements

`func (o *AccessProfile) GetEntitlements() []string`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *AccessProfile) GetEntitlementsOk() (*[]string, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *AccessProfile) SetEntitlements(v []string)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *AccessProfile) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### GetApprovalSchemes

`func (o *AccessProfile) GetApprovalSchemes() string`

GetApprovalSchemes returns the ApprovalSchemes field if non-nil, zero value otherwise.

### GetApprovalSchemesOk

`func (o *AccessProfile) GetApprovalSchemesOk() (*string, bool)`

GetApprovalSchemesOk returns a tuple with the ApprovalSchemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalSchemes

`func (o *AccessProfile) SetApprovalSchemes(v string)`

SetApprovalSchemes sets ApprovalSchemes field to given value.

### HasApprovalSchemes

`func (o *AccessProfile) HasApprovalSchemes() bool`

HasApprovalSchemes returns a boolean if a field has been set.

### GetRevokeRequestApprovalSchemes

`func (o *AccessProfile) GetRevokeRequestApprovalSchemes() string`

GetRevokeRequestApprovalSchemes returns the RevokeRequestApprovalSchemes field if non-nil, zero value otherwise.

### GetRevokeRequestApprovalSchemesOk

`func (o *AccessProfile) GetRevokeRequestApprovalSchemesOk() (*string, bool)`

GetRevokeRequestApprovalSchemesOk returns a tuple with the RevokeRequestApprovalSchemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeRequestApprovalSchemes

`func (o *AccessProfile) SetRevokeRequestApprovalSchemes(v string)`

SetRevokeRequestApprovalSchemes sets RevokeRequestApprovalSchemes field to given value.

### HasRevokeRequestApprovalSchemes

`func (o *AccessProfile) HasRevokeRequestApprovalSchemes() bool`

HasRevokeRequestApprovalSchemes returns a boolean if a field has been set.

### GetRequestCommentsRequired

`func (o *AccessProfile) GetRequestCommentsRequired() bool`

GetRequestCommentsRequired returns the RequestCommentsRequired field if non-nil, zero value otherwise.

### GetRequestCommentsRequiredOk

`func (o *AccessProfile) GetRequestCommentsRequiredOk() (*bool, bool)`

GetRequestCommentsRequiredOk returns a tuple with the RequestCommentsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCommentsRequired

`func (o *AccessProfile) SetRequestCommentsRequired(v bool)`

SetRequestCommentsRequired sets RequestCommentsRequired field to given value.

### HasRequestCommentsRequired

`func (o *AccessProfile) HasRequestCommentsRequired() bool`

HasRequestCommentsRequired returns a boolean if a field has been set.

### GetDeniedCommentsRequired

`func (o *AccessProfile) GetDeniedCommentsRequired() bool`

GetDeniedCommentsRequired returns the DeniedCommentsRequired field if non-nil, zero value otherwise.

### GetDeniedCommentsRequiredOk

`func (o *AccessProfile) GetDeniedCommentsRequiredOk() (*bool, bool)`

GetDeniedCommentsRequiredOk returns a tuple with the DeniedCommentsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedCommentsRequired

`func (o *AccessProfile) SetDeniedCommentsRequired(v bool)`

SetDeniedCommentsRequired sets DeniedCommentsRequired field to given value.

### HasDeniedCommentsRequired

`func (o *AccessProfile) HasDeniedCommentsRequired() bool`

HasDeniedCommentsRequired returns a boolean if a field has been set.

### GetDisabled

`func (o *AccessProfile) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *AccessProfile) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *AccessProfile) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *AccessProfile) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


