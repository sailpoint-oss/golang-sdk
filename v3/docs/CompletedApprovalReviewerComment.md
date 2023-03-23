# CompletedApprovalReviewerComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Content of the comment | [optional] 
**Author** | Pointer to [**CommentDtoAuthor**](CommentDtoAuthor.md) |  | [optional] 
**Created** | Pointer to **time.Time** | Date and time comment was created | [optional] 

## Methods

### NewCompletedApprovalReviewerComment

`func NewCompletedApprovalReviewerComment() *CompletedApprovalReviewerComment`

NewCompletedApprovalReviewerComment instantiates a new CompletedApprovalReviewerComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompletedApprovalReviewerCommentWithDefaults

`func NewCompletedApprovalReviewerCommentWithDefaults() *CompletedApprovalReviewerComment`

NewCompletedApprovalReviewerCommentWithDefaults instantiates a new CompletedApprovalReviewerComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *CompletedApprovalReviewerComment) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CompletedApprovalReviewerComment) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CompletedApprovalReviewerComment) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CompletedApprovalReviewerComment) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetAuthor

`func (o *CompletedApprovalReviewerComment) GetAuthor() CommentDtoAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *CompletedApprovalReviewerComment) GetAuthorOk() (*CommentDtoAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *CompletedApprovalReviewerComment) SetAuthor(v CommentDtoAuthor)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *CompletedApprovalReviewerComment) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetCreated

`func (o *CompletedApprovalReviewerComment) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CompletedApprovalReviewerComment) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CompletedApprovalReviewerComment) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CompletedApprovalReviewerComment) HasCreated() bool`

HasCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


