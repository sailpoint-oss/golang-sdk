# CertificationReference1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | DTO type of certification for review. | [optional] 
**Id** | Pointer to **string** | ID of certification for review. | [optional] 
**Name** | Pointer to **string** | Display name of certification for review. | [optional] 
**Reviewer** | Pointer to [**Reviewer1**](Reviewer1.md) |  | [optional] 

## Methods

### NewCertificationReference1

`func NewCertificationReference1() *CertificationReference1`

NewCertificationReference1 instantiates a new CertificationReference1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificationReference1WithDefaults

`func NewCertificationReference1WithDefaults() *CertificationReference1`

NewCertificationReference1WithDefaults instantiates a new CertificationReference1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CertificationReference1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CertificationReference1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CertificationReference1) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CertificationReference1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *CertificationReference1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificationReference1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificationReference1) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CertificationReference1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CertificationReference1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificationReference1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificationReference1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CertificationReference1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReviewer

`func (o *CertificationReference1) GetReviewer() Reviewer1`

GetReviewer returns the Reviewer field if non-nil, zero value otherwise.

### GetReviewerOk

`func (o *CertificationReference1) GetReviewerOk() (*Reviewer1, bool)`

GetReviewerOk returns a tuple with the Reviewer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewer

`func (o *CertificationReference1) SetReviewer(v Reviewer1)`

SetReviewer sets Reviewer field to given value.

### HasReviewer

`func (o *CertificationReference1) HasReviewer() bool`

HasReviewer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


