# RequestedItemDtoRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the item being requested. | 
**Id** | **string** | ID of Role, Access Profile or Entitlement being requested. | 
**Comment** | Pointer to **string** | Comment provided by requester. * Comment is required when the request is of type Revoke Access.  | [optional] 
**ClientMetadata** | Pointer to **map[string]string** | Arbitrary key-value pairs. They will never be processed by the IdentityNow system but will be returned on associated APIs such as /account-activities and /access-request-status. | [optional] 
**RemoveDate** | Pointer to **time.Time** | The date the role or access profile or entitlement is no longer assigned to the specified identity. Also known as the expiration date. * Specify a date in the future. * The current SLA for the deprovisioning is 24 hours. * This date can be modified to either extend or decrease the duration of access item assignments for the specified identity. You can change the expiration date for requests for yourself or direct reports, but you cannot remove an expiration date on an already approved item. If the access request has not been approved, you can cancel it and submit a new one without the expiration. If it has already been approved, then you have to revoke the access and then re-request without the expiration.  | [optional] 
**AssignmentId** | Pointer to **NullableString** | The assignmentId for a specific role assignment on the identity. This id is used to revoke that specific roleAssignment on that identity. * For use with REVOKE_ACCESS requests for roles for identities with multiple accounts on a single source.  | [optional] 
**NativeIdentity** | Pointer to **NullableString** | The &#39;distinguishedName&#39; field for an account on the identity, also called nativeIdentity. This nativeIdentity is used to revoke a specific attributeAssignment on the identity. * For use with REVOKE_ACCESS requests for entitlements for identities with multiple accounts on a single source.  | [optional] 
**AccountSelection** | Pointer to [**[]SourceItemRef**](SourceItemRef.md) | The accounts where the access item will be provisioned to * Includes selections performed by the user in the event of multiple accounts existing on the same source * Also includes details for sources where user only has one account  | [optional] 

## Methods

### NewRequestedItemDtoRef

`func NewRequestedItemDtoRef(type_ string, id string, ) *RequestedItemDtoRef`

NewRequestedItemDtoRef instantiates a new RequestedItemDtoRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestedItemDtoRefWithDefaults

`func NewRequestedItemDtoRefWithDefaults() *RequestedItemDtoRef`

NewRequestedItemDtoRefWithDefaults instantiates a new RequestedItemDtoRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RequestedItemDtoRef) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequestedItemDtoRef) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequestedItemDtoRef) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *RequestedItemDtoRef) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestedItemDtoRef) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestedItemDtoRef) SetId(v string)`

SetId sets Id field to given value.


### GetComment

`func (o *RequestedItemDtoRef) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RequestedItemDtoRef) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RequestedItemDtoRef) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RequestedItemDtoRef) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetClientMetadata

`func (o *RequestedItemDtoRef) GetClientMetadata() map[string]string`

GetClientMetadata returns the ClientMetadata field if non-nil, zero value otherwise.

### GetClientMetadataOk

`func (o *RequestedItemDtoRef) GetClientMetadataOk() (*map[string]string, bool)`

GetClientMetadataOk returns a tuple with the ClientMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMetadata

`func (o *RequestedItemDtoRef) SetClientMetadata(v map[string]string)`

SetClientMetadata sets ClientMetadata field to given value.

### HasClientMetadata

`func (o *RequestedItemDtoRef) HasClientMetadata() bool`

HasClientMetadata returns a boolean if a field has been set.

### GetRemoveDate

`func (o *RequestedItemDtoRef) GetRemoveDate() time.Time`

GetRemoveDate returns the RemoveDate field if non-nil, zero value otherwise.

### GetRemoveDateOk

`func (o *RequestedItemDtoRef) GetRemoveDateOk() (*time.Time, bool)`

GetRemoveDateOk returns a tuple with the RemoveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveDate

`func (o *RequestedItemDtoRef) SetRemoveDate(v time.Time)`

SetRemoveDate sets RemoveDate field to given value.

### HasRemoveDate

`func (o *RequestedItemDtoRef) HasRemoveDate() bool`

HasRemoveDate returns a boolean if a field has been set.

### GetAssignmentId

`func (o *RequestedItemDtoRef) GetAssignmentId() string`

GetAssignmentId returns the AssignmentId field if non-nil, zero value otherwise.

### GetAssignmentIdOk

`func (o *RequestedItemDtoRef) GetAssignmentIdOk() (*string, bool)`

GetAssignmentIdOk returns a tuple with the AssignmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentId

`func (o *RequestedItemDtoRef) SetAssignmentId(v string)`

SetAssignmentId sets AssignmentId field to given value.

### HasAssignmentId

`func (o *RequestedItemDtoRef) HasAssignmentId() bool`

HasAssignmentId returns a boolean if a field has been set.

### SetAssignmentIdNil

`func (o *RequestedItemDtoRef) SetAssignmentIdNil(b bool)`

 SetAssignmentIdNil sets the value for AssignmentId to be an explicit nil

### UnsetAssignmentId
`func (o *RequestedItemDtoRef) UnsetAssignmentId()`

UnsetAssignmentId ensures that no value is present for AssignmentId, not even an explicit nil
### GetNativeIdentity

`func (o *RequestedItemDtoRef) GetNativeIdentity() string`

GetNativeIdentity returns the NativeIdentity field if non-nil, zero value otherwise.

### GetNativeIdentityOk

`func (o *RequestedItemDtoRef) GetNativeIdentityOk() (*string, bool)`

GetNativeIdentityOk returns a tuple with the NativeIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeIdentity

`func (o *RequestedItemDtoRef) SetNativeIdentity(v string)`

SetNativeIdentity sets NativeIdentity field to given value.

### HasNativeIdentity

`func (o *RequestedItemDtoRef) HasNativeIdentity() bool`

HasNativeIdentity returns a boolean if a field has been set.

### SetNativeIdentityNil

`func (o *RequestedItemDtoRef) SetNativeIdentityNil(b bool)`

 SetNativeIdentityNil sets the value for NativeIdentity to be an explicit nil

### UnsetNativeIdentity
`func (o *RequestedItemDtoRef) UnsetNativeIdentity()`

UnsetNativeIdentity ensures that no value is present for NativeIdentity, not even an explicit nil
### GetAccountSelection

`func (o *RequestedItemDtoRef) GetAccountSelection() []SourceItemRef`

GetAccountSelection returns the AccountSelection field if non-nil, zero value otherwise.

### GetAccountSelectionOk

`func (o *RequestedItemDtoRef) GetAccountSelectionOk() (*[]SourceItemRef, bool)`

GetAccountSelectionOk returns a tuple with the AccountSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSelection

`func (o *RequestedItemDtoRef) SetAccountSelection(v []SourceItemRef)`

SetAccountSelection sets AccountSelection field to given value.

### HasAccountSelection

`func (o *RequestedItemDtoRef) HasAccountSelection() bool`

HasAccountSelection returns a boolean if a field has been set.

### SetAccountSelectionNil

`func (o *RequestedItemDtoRef) SetAccountSelectionNil(b bool)`

 SetAccountSelectionNil sets the value for AccountSelection to be an explicit nil

### UnsetAccountSelection
`func (o *RequestedItemDtoRef) UnsetAccountSelection()`

UnsetAccountSelection ensures that no value is present for AccountSelection, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


