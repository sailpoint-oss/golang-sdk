# NonEmployeeSourceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Approvers** | Pointer to [**[]NonEmployeeIdentityReferenceWithId**](NonEmployeeIdentityReferenceWithId.md) | List of approvers | [optional] 
**AccountManagers** | Pointer to [**[]NonEmployeeIdentityReferenceWithId**](NonEmployeeIdentityReferenceWithId.md) | List of account managers | [optional] 
**Modified** | Pointer to **time.Time** | When the request was last modified. | [optional] 
**Created** | Pointer to **time.Time** | When the request was created. | [optional] 

## Methods

### NewNonEmployeeSourceAllOf

`func NewNonEmployeeSourceAllOf() *NonEmployeeSourceAllOf`

NewNonEmployeeSourceAllOf instantiates a new NonEmployeeSourceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonEmployeeSourceAllOfWithDefaults

`func NewNonEmployeeSourceAllOfWithDefaults() *NonEmployeeSourceAllOf`

NewNonEmployeeSourceAllOfWithDefaults instantiates a new NonEmployeeSourceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovers

`func (o *NonEmployeeSourceAllOf) GetApprovers() []NonEmployeeIdentityReferenceWithId`

GetApprovers returns the Approvers field if non-nil, zero value otherwise.

### GetApproversOk

`func (o *NonEmployeeSourceAllOf) GetApproversOk() (*[]NonEmployeeIdentityReferenceWithId, bool)`

GetApproversOk returns a tuple with the Approvers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovers

`func (o *NonEmployeeSourceAllOf) SetApprovers(v []NonEmployeeIdentityReferenceWithId)`

SetApprovers sets Approvers field to given value.

### HasApprovers

`func (o *NonEmployeeSourceAllOf) HasApprovers() bool`

HasApprovers returns a boolean if a field has been set.

### GetAccountManagers

`func (o *NonEmployeeSourceAllOf) GetAccountManagers() []NonEmployeeIdentityReferenceWithId`

GetAccountManagers returns the AccountManagers field if non-nil, zero value otherwise.

### GetAccountManagersOk

`func (o *NonEmployeeSourceAllOf) GetAccountManagersOk() (*[]NonEmployeeIdentityReferenceWithId, bool)`

GetAccountManagersOk returns a tuple with the AccountManagers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountManagers

`func (o *NonEmployeeSourceAllOf) SetAccountManagers(v []NonEmployeeIdentityReferenceWithId)`

SetAccountManagers sets AccountManagers field to given value.

### HasAccountManagers

`func (o *NonEmployeeSourceAllOf) HasAccountManagers() bool`

HasAccountManagers returns a boolean if a field has been set.

### GetModified

`func (o *NonEmployeeSourceAllOf) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *NonEmployeeSourceAllOf) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *NonEmployeeSourceAllOf) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *NonEmployeeSourceAllOf) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetCreated

`func (o *NonEmployeeSourceAllOf) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *NonEmployeeSourceAllOf) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *NonEmployeeSourceAllOf) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *NonEmployeeSourceAllOf) HasCreated() bool`

HasCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


