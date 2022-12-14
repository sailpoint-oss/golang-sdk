# AccessProfileCreateEto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | **string** |  | 
**SourceId** | **string** |  | 
**OwnerId** | **string** |  | 
**Entitlements** | **[]string** |  | 
**ApprovalSchemes** | Pointer to **string** | Comma-separated list of approval schemes. Each approval scheme is one of - manager - appOwner - sourceOwner - accessProfileOwner - workgroup:\\&lt;workgroupId&gt; | [optional] 
**RevokeRequestApprovalSchemes** | Pointer to **string** | Comma-separated list of revoke request approval schemes. Each approval scheme is one of - manager - sourceOwner - accessProfileOwner - workgroup:\\&lt;workgroupId&gt; | [optional] 
**RequestCommentsRequired** | Pointer to **bool** |  | [optional] 
**DeniedCommentsRequired** | Pointer to **bool** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewAccessProfileCreateEto

`func NewAccessProfileCreateEto(name string, description string, sourceId string, ownerId string, entitlements []string, ) *AccessProfileCreateEto`

NewAccessProfileCreateEto instantiates a new AccessProfileCreateEto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessProfileCreateEtoWithDefaults

`func NewAccessProfileCreateEtoWithDefaults() *AccessProfileCreateEto`

NewAccessProfileCreateEtoWithDefaults instantiates a new AccessProfileCreateEto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AccessProfileCreateEto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessProfileCreateEto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessProfileCreateEto) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AccessProfileCreateEto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccessProfileCreateEto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccessProfileCreateEto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSourceId

`func (o *AccessProfileCreateEto) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *AccessProfileCreateEto) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *AccessProfileCreateEto) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetOwnerId

`func (o *AccessProfileCreateEto) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *AccessProfileCreateEto) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *AccessProfileCreateEto) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetEntitlements

`func (o *AccessProfileCreateEto) GetEntitlements() []string`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *AccessProfileCreateEto) GetEntitlementsOk() (*[]string, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *AccessProfileCreateEto) SetEntitlements(v []string)`

SetEntitlements sets Entitlements field to given value.


### GetApprovalSchemes

`func (o *AccessProfileCreateEto) GetApprovalSchemes() string`

GetApprovalSchemes returns the ApprovalSchemes field if non-nil, zero value otherwise.

### GetApprovalSchemesOk

`func (o *AccessProfileCreateEto) GetApprovalSchemesOk() (*string, bool)`

GetApprovalSchemesOk returns a tuple with the ApprovalSchemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalSchemes

`func (o *AccessProfileCreateEto) SetApprovalSchemes(v string)`

SetApprovalSchemes sets ApprovalSchemes field to given value.

### HasApprovalSchemes

`func (o *AccessProfileCreateEto) HasApprovalSchemes() bool`

HasApprovalSchemes returns a boolean if a field has been set.

### GetRevokeRequestApprovalSchemes

`func (o *AccessProfileCreateEto) GetRevokeRequestApprovalSchemes() string`

GetRevokeRequestApprovalSchemes returns the RevokeRequestApprovalSchemes field if non-nil, zero value otherwise.

### GetRevokeRequestApprovalSchemesOk

`func (o *AccessProfileCreateEto) GetRevokeRequestApprovalSchemesOk() (*string, bool)`

GetRevokeRequestApprovalSchemesOk returns a tuple with the RevokeRequestApprovalSchemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeRequestApprovalSchemes

`func (o *AccessProfileCreateEto) SetRevokeRequestApprovalSchemes(v string)`

SetRevokeRequestApprovalSchemes sets RevokeRequestApprovalSchemes field to given value.

### HasRevokeRequestApprovalSchemes

`func (o *AccessProfileCreateEto) HasRevokeRequestApprovalSchemes() bool`

HasRevokeRequestApprovalSchemes returns a boolean if a field has been set.

### GetRequestCommentsRequired

`func (o *AccessProfileCreateEto) GetRequestCommentsRequired() bool`

GetRequestCommentsRequired returns the RequestCommentsRequired field if non-nil, zero value otherwise.

### GetRequestCommentsRequiredOk

`func (o *AccessProfileCreateEto) GetRequestCommentsRequiredOk() (*bool, bool)`

GetRequestCommentsRequiredOk returns a tuple with the RequestCommentsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCommentsRequired

`func (o *AccessProfileCreateEto) SetRequestCommentsRequired(v bool)`

SetRequestCommentsRequired sets RequestCommentsRequired field to given value.

### HasRequestCommentsRequired

`func (o *AccessProfileCreateEto) HasRequestCommentsRequired() bool`

HasRequestCommentsRequired returns a boolean if a field has been set.

### GetDeniedCommentsRequired

`func (o *AccessProfileCreateEto) GetDeniedCommentsRequired() bool`

GetDeniedCommentsRequired returns the DeniedCommentsRequired field if non-nil, zero value otherwise.

### GetDeniedCommentsRequiredOk

`func (o *AccessProfileCreateEto) GetDeniedCommentsRequiredOk() (*bool, bool)`

GetDeniedCommentsRequiredOk returns a tuple with the DeniedCommentsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedCommentsRequired

`func (o *AccessProfileCreateEto) SetDeniedCommentsRequired(v bool)`

SetDeniedCommentsRequired sets DeniedCommentsRequired field to given value.

### HasDeniedCommentsRequired

`func (o *AccessProfileCreateEto) HasDeniedCommentsRequired() bool`

HasDeniedCommentsRequired returns a boolean if a field has been set.

### GetDisabled

`func (o *AccessProfileCreateEto) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *AccessProfileCreateEto) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *AccessProfileCreateEto) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *AccessProfileCreateEto) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


