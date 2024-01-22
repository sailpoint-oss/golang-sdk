# IdentityWithNewAccess1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityId** | **string** | Set of identity IDs to be checked. | 
**AccessRefs** | [**[]EntitlementRef**](EntitlementRef.md) | The bundle of access profiles to be added to the identities specified. All references must be ENTITLEMENT type. | 
**ClientMetadata** | Pointer to **map[string]string** | Arbitrary key-value pairs. They will never be processed by the IdentityNow system but will be returned on completion of the violation check. | [optional] 

## Methods

### NewIdentityWithNewAccess1

`func NewIdentityWithNewAccess1(identityId string, accessRefs []EntitlementRef, ) *IdentityWithNewAccess1`

NewIdentityWithNewAccess1 instantiates a new IdentityWithNewAccess1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityWithNewAccess1WithDefaults

`func NewIdentityWithNewAccess1WithDefaults() *IdentityWithNewAccess1`

NewIdentityWithNewAccess1WithDefaults instantiates a new IdentityWithNewAccess1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityId

`func (o *IdentityWithNewAccess1) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *IdentityWithNewAccess1) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *IdentityWithNewAccess1) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.


### GetAccessRefs

`func (o *IdentityWithNewAccess1) GetAccessRefs() []EntitlementRef`

GetAccessRefs returns the AccessRefs field if non-nil, zero value otherwise.

### GetAccessRefsOk

`func (o *IdentityWithNewAccess1) GetAccessRefsOk() (*[]EntitlementRef, bool)`

GetAccessRefsOk returns a tuple with the AccessRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRefs

`func (o *IdentityWithNewAccess1) SetAccessRefs(v []EntitlementRef)`

SetAccessRefs sets AccessRefs field to given value.


### GetClientMetadata

`func (o *IdentityWithNewAccess1) GetClientMetadata() map[string]string`

GetClientMetadata returns the ClientMetadata field if non-nil, zero value otherwise.

### GetClientMetadataOk

`func (o *IdentityWithNewAccess1) GetClientMetadataOk() (*map[string]string, bool)`

GetClientMetadataOk returns a tuple with the ClientMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMetadata

`func (o *IdentityWithNewAccess1) SetClientMetadata(v map[string]string)`

SetClientMetadata sets ClientMetadata field to given value.

### HasClientMetadata

`func (o *IdentityWithNewAccess1) HasClientMetadata() bool`

HasClientMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


