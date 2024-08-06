# Reassignment1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to [**CertificationReference1**](CertificationReference1.md) |  | [optional] 
**Comment** | Pointer to **string** | Comments from the previous reviewer. | [optional] 

## Methods

### NewReassignment1

`func NewReassignment1() *Reassignment1`

NewReassignment1 instantiates a new Reassignment1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReassignment1WithDefaults

`func NewReassignment1WithDefaults() *Reassignment1`

NewReassignment1WithDefaults instantiates a new Reassignment1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *Reassignment1) GetFrom() CertificationReference1`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Reassignment1) GetFromOk() (*CertificationReference1, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *Reassignment1) SetFrom(v CertificationReference1)`

SetFrom sets From field to given value.

### HasFrom

`func (o *Reassignment1) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetComment

`func (o *Reassignment1) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Reassignment1) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Reassignment1) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Reassignment1) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


