# ApprovalComment1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | The comment text | [optional] 
**Commenter** | Pointer to **string** | The name of the commenter | [optional] 
**Date** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 

## Methods

### NewApprovalComment1

`func NewApprovalComment1() *ApprovalComment1`

NewApprovalComment1 instantiates a new ApprovalComment1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalComment1WithDefaults

`func NewApprovalComment1WithDefaults() *ApprovalComment1`

NewApprovalComment1WithDefaults instantiates a new ApprovalComment1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *ApprovalComment1) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ApprovalComment1) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ApprovalComment1) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ApprovalComment1) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCommenter

`func (o *ApprovalComment1) GetCommenter() string`

GetCommenter returns the Commenter field if non-nil, zero value otherwise.

### GetCommenterOk

`func (o *ApprovalComment1) GetCommenterOk() (*string, bool)`

GetCommenterOk returns a tuple with the Commenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenter

`func (o *ApprovalComment1) SetCommenter(v string)`

SetCommenter sets Commenter field to given value.

### HasCommenter

`func (o *ApprovalComment1) HasCommenter() bool`

HasCommenter returns a boolean if a field has been set.

### GetDate

`func (o *ApprovalComment1) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ApprovalComment1) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ApprovalComment1) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *ApprovalComment1) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *ApprovalComment1) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *ApprovalComment1) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


