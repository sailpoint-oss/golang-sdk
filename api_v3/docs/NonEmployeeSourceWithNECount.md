# NonEmployeeSourceWithNECount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Non-Employee source id. | [optional] 
**SourceId** | Pointer to **string** | Source Id associated with this non-employee source. | [optional] 
**Name** | Pointer to **string** | Source name associated with this non-employee source. | [optional] 
**Description** | Pointer to **string** | Source description associated with this non-employee source. | [optional] 
**Approvers** | Pointer to [**[]NonEmployeeIdentityReferenceWithId**](NonEmployeeIdentityReferenceWithId.md) | List of approvers | [optional] 
**AccountManagers** | Pointer to [**[]NonEmployeeIdentityReferenceWithId**](NonEmployeeIdentityReferenceWithId.md) | List of account managers | [optional] 
**Modified** | Pointer to **time.Time** | When the request was last modified. | [optional] 
**Created** | Pointer to **time.Time** | When the request was created. | [optional] 
**NonEmployeeCount** | Pointer to **int32** | Number of non-employee records associated with this source. | [optional] 

## Methods

### NewNonEmployeeSourceWithNECount

`func NewNonEmployeeSourceWithNECount() *NonEmployeeSourceWithNECount`

NewNonEmployeeSourceWithNECount instantiates a new NonEmployeeSourceWithNECount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonEmployeeSourceWithNECountWithDefaults

`func NewNonEmployeeSourceWithNECountWithDefaults() *NonEmployeeSourceWithNECount`

NewNonEmployeeSourceWithNECountWithDefaults instantiates a new NonEmployeeSourceWithNECount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NonEmployeeSourceWithNECount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NonEmployeeSourceWithNECount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NonEmployeeSourceWithNECount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NonEmployeeSourceWithNECount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSourceId

`func (o *NonEmployeeSourceWithNECount) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *NonEmployeeSourceWithNECount) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *NonEmployeeSourceWithNECount) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *NonEmployeeSourceWithNECount) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetName

`func (o *NonEmployeeSourceWithNECount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NonEmployeeSourceWithNECount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NonEmployeeSourceWithNECount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NonEmployeeSourceWithNECount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *NonEmployeeSourceWithNECount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NonEmployeeSourceWithNECount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NonEmployeeSourceWithNECount) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NonEmployeeSourceWithNECount) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetApprovers

`func (o *NonEmployeeSourceWithNECount) GetApprovers() []NonEmployeeIdentityReferenceWithId`

GetApprovers returns the Approvers field if non-nil, zero value otherwise.

### GetApproversOk

`func (o *NonEmployeeSourceWithNECount) GetApproversOk() (*[]NonEmployeeIdentityReferenceWithId, bool)`

GetApproversOk returns a tuple with the Approvers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovers

`func (o *NonEmployeeSourceWithNECount) SetApprovers(v []NonEmployeeIdentityReferenceWithId)`

SetApprovers sets Approvers field to given value.

### HasApprovers

`func (o *NonEmployeeSourceWithNECount) HasApprovers() bool`

HasApprovers returns a boolean if a field has been set.

### GetAccountManagers

`func (o *NonEmployeeSourceWithNECount) GetAccountManagers() []NonEmployeeIdentityReferenceWithId`

GetAccountManagers returns the AccountManagers field if non-nil, zero value otherwise.

### GetAccountManagersOk

`func (o *NonEmployeeSourceWithNECount) GetAccountManagersOk() (*[]NonEmployeeIdentityReferenceWithId, bool)`

GetAccountManagersOk returns a tuple with the AccountManagers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountManagers

`func (o *NonEmployeeSourceWithNECount) SetAccountManagers(v []NonEmployeeIdentityReferenceWithId)`

SetAccountManagers sets AccountManagers field to given value.

### HasAccountManagers

`func (o *NonEmployeeSourceWithNECount) HasAccountManagers() bool`

HasAccountManagers returns a boolean if a field has been set.

### GetModified

`func (o *NonEmployeeSourceWithNECount) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *NonEmployeeSourceWithNECount) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *NonEmployeeSourceWithNECount) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *NonEmployeeSourceWithNECount) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetCreated

`func (o *NonEmployeeSourceWithNECount) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *NonEmployeeSourceWithNECount) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *NonEmployeeSourceWithNECount) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *NonEmployeeSourceWithNECount) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetNonEmployeeCount

`func (o *NonEmployeeSourceWithNECount) GetNonEmployeeCount() int32`

GetNonEmployeeCount returns the NonEmployeeCount field if non-nil, zero value otherwise.

### GetNonEmployeeCountOk

`func (o *NonEmployeeSourceWithNECount) GetNonEmployeeCountOk() (*int32, bool)`

GetNonEmployeeCountOk returns a tuple with the NonEmployeeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonEmployeeCount

`func (o *NonEmployeeSourceWithNECount) SetNonEmployeeCount(v int32)`

SetNonEmployeeCount sets NonEmployeeCount field to given value.

### HasNonEmployeeCount

`func (o *NonEmployeeSourceWithNECount) HasNonEmployeeCount() bool`

HasNonEmployeeCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


