# UpdateUserPermissionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **string** |  | [optional] 
**IsAdmin** | Pointer to **string** | Indicates if user should be an IDN Admin.  \&quot;0\&quot; for false, \&quot;1\&quot; for true. | [optional] 
**AdminType** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateUserPermissionsRequest

`func NewUpdateUserPermissionsRequest() *UpdateUserPermissionsRequest`

NewUpdateUserPermissionsRequest instantiates a new UpdateUserPermissionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserPermissionsRequestWithDefaults

`func NewUpdateUserPermissionsRequestWithDefaults() *UpdateUserPermissionsRequest`

NewUpdateUserPermissionsRequestWithDefaults instantiates a new UpdateUserPermissionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *UpdateUserPermissionsRequest) GetIds() string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *UpdateUserPermissionsRequest) GetIdsOk() (*string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *UpdateUserPermissionsRequest) SetIds(v string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *UpdateUserPermissionsRequest) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetIsAdmin

`func (o *UpdateUserPermissionsRequest) GetIsAdmin() string`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *UpdateUserPermissionsRequest) GetIsAdminOk() (*string, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *UpdateUserPermissionsRequest) SetIsAdmin(v string)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *UpdateUserPermissionsRequest) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetAdminType

`func (o *UpdateUserPermissionsRequest) GetAdminType() string`

GetAdminType returns the AdminType field if non-nil, zero value otherwise.

### GetAdminTypeOk

`func (o *UpdateUserPermissionsRequest) GetAdminTypeOk() (*string, bool)`

GetAdminTypeOk returns a tuple with the AdminType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminType

`func (o *UpdateUserPermissionsRequest) SetAdminType(v string)`

SetAdminType sets AdminType field to given value.

### HasAdminType

`func (o *UpdateUserPermissionsRequest) HasAdminType() bool`

HasAdminType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


