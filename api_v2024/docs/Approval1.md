# Approval1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comments** | Pointer to [**[]ApprovalComment2**](ApprovalComment2.md) |  | [optional] 
**Modified** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**Owner** | Pointer to [**ActivityIdentity**](ActivityIdentity.md) |  | [optional] 
**Result** | Pointer to **string** | The result of the approval | [optional] 
**AttributeRequest** | Pointer to [**AttributeRequest**](AttributeRequest.md) |  | [optional] 
**Source** | Pointer to [**AccountSource**](AccountSource.md) |  | [optional] 

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

`func (o *Approval1) GetComments() []ApprovalComment2`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *Approval1) GetCommentsOk() (*[]ApprovalComment2, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *Approval1) SetComments(v []ApprovalComment2)`

SetComments sets Comments field to given value.

### HasComments

`func (o *Approval1) HasComments() bool`

HasComments returns a boolean if a field has been set.

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

`func (o *Approval1) GetOwner() ActivityIdentity`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Approval1) GetOwnerOk() (*ActivityIdentity, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Approval1) SetOwner(v ActivityIdentity)`

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

### GetAttributeRequest

`func (o *Approval1) GetAttributeRequest() AttributeRequest`

GetAttributeRequest returns the AttributeRequest field if non-nil, zero value otherwise.

### GetAttributeRequestOk

`func (o *Approval1) GetAttributeRequestOk() (*AttributeRequest, bool)`

GetAttributeRequestOk returns a tuple with the AttributeRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeRequest

`func (o *Approval1) SetAttributeRequest(v AttributeRequest)`

SetAttributeRequest sets AttributeRequest field to given value.

### HasAttributeRequest

`func (o *Approval1) HasAttributeRequest() bool`

HasAttributeRequest returns a boolean if a field has been set.

### GetSource

`func (o *Approval1) GetSource() AccountSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Approval1) GetSourceOk() (*AccountSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Approval1) SetSource(v AccountSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *Approval1) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


