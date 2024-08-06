# Approval1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comments** | Pointer to [**[]ApprovalComment1**](ApprovalComment1.md) |  | [optional] 
**Created** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**Modified** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**Owner** | Pointer to [**AccountSource**](AccountSource.md) |  | [optional] 
**Result** | Pointer to **string** | The result of the approval | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewApproval1

`func NewApproval1() *Approval1`

NewApproval1 instantiates a new Approval1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApproval1WithDefaults

`func NewApproval1WithDefaults() *Approval1`

NewApproval1WithDefaults instantiates a new Approval1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComments

`func (o *Approval1) GetComments() []ApprovalComment1`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *Approval1) GetCommentsOk() (*[]ApprovalComment1, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *Approval1) SetComments(v []ApprovalComment1)`

SetComments sets Comments field to given value.

### HasComments

`func (o *Approval1) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCreated

`func (o *Approval1) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Approval1) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Approval1) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Approval1) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *Approval1) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *Approval1) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetModified

`func (o *Approval1) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Approval1) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Approval1) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *Approval1) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModifiedNil

`func (o *Approval1) SetModifiedNil(b bool)`

 SetModifiedNil sets the value for Modified to be an explicit nil

### UnsetModified
`func (o *Approval1) UnsetModified()`

UnsetModified ensures that no value is present for Modified, not even an explicit nil
### GetOwner

`func (o *Approval1) GetOwner() AccountSource`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Approval1) GetOwnerOk() (*AccountSource, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Approval1) SetOwner(v AccountSource)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Approval1) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetResult

`func (o *Approval1) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *Approval1) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *Approval1) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *Approval1) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetType

`func (o *Approval1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Approval1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Approval1) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Approval1) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *Approval1) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Approval1) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


