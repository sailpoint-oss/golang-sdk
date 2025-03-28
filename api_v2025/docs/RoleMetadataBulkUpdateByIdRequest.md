# RoleMetadataBulkUpdateByIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | **[]string** | Roles&#39; Id to be updated | 
**Operation** | **string** | The operation to be performed | 
**ReplaceScope** | Pointer to **string** | The choice of update scope. | [optional] 
**Values** | [**[]RoleMetadataBulkUpdateByIdRequestValuesInner**](RoleMetadataBulkUpdateByIdRequestValuesInner.md) | The metadata to be updated, including attribute key and value. | 

## Methods

### NewRoleMetadataBulkUpdateByIdRequest

`func NewRoleMetadataBulkUpdateByIdRequest(roles []string, operation string, values []RoleMetadataBulkUpdateByIdRequestValuesInner, ) *RoleMetadataBulkUpdateByIdRequest`

NewRoleMetadataBulkUpdateByIdRequest instantiates a new RoleMetadataBulkUpdateByIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleMetadataBulkUpdateByIdRequestWithDefaults

`func NewRoleMetadataBulkUpdateByIdRequestWithDefaults() *RoleMetadataBulkUpdateByIdRequest`

NewRoleMetadataBulkUpdateByIdRequestWithDefaults instantiates a new RoleMetadataBulkUpdateByIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoles

`func (o *RoleMetadataBulkUpdateByIdRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *RoleMetadataBulkUpdateByIdRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *RoleMetadataBulkUpdateByIdRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.


### GetOperation

`func (o *RoleMetadataBulkUpdateByIdRequest) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *RoleMetadataBulkUpdateByIdRequest) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *RoleMetadataBulkUpdateByIdRequest) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetReplaceScope

`func (o *RoleMetadataBulkUpdateByIdRequest) GetReplaceScope() string`

GetReplaceScope returns the ReplaceScope field if non-nil, zero value otherwise.

### GetReplaceScopeOk

`func (o *RoleMetadataBulkUpdateByIdRequest) GetReplaceScopeOk() (*string, bool)`

GetReplaceScopeOk returns a tuple with the ReplaceScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceScope

`func (o *RoleMetadataBulkUpdateByIdRequest) SetReplaceScope(v string)`

SetReplaceScope sets ReplaceScope field to given value.

### HasReplaceScope

`func (o *RoleMetadataBulkUpdateByIdRequest) HasReplaceScope() bool`

HasReplaceScope returns a boolean if a field has been set.

### GetValues

`func (o *RoleMetadataBulkUpdateByIdRequest) GetValues() []RoleMetadataBulkUpdateByIdRequestValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *RoleMetadataBulkUpdateByIdRequest) GetValuesOk() (*[]RoleMetadataBulkUpdateByIdRequestValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *RoleMetadataBulkUpdateByIdRequest) SetValues(v []RoleMetadataBulkUpdateByIdRequestValuesInner)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


