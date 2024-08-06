# Reviewer1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The reviewer&#39;s DTO type. | 
**Id** | **string** | The reviewer&#39;s ID. | 
**Name** | **string** | The reviewer&#39;s display name. | 
**Email** | Pointer to **NullableString** | The reviewing identity&#39;s email. Only applicable to &#x60;IDENTITY&#x60;. | [optional] 

## Methods

### NewReviewer1

`func NewReviewer1(type_ string, id string, name string, ) *Reviewer1`

NewReviewer1 instantiates a new Reviewer1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewer1WithDefaults

`func NewReviewer1WithDefaults() *Reviewer1`

NewReviewer1WithDefaults instantiates a new Reviewer1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Reviewer1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Reviewer1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Reviewer1) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *Reviewer1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Reviewer1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Reviewer1) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Reviewer1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Reviewer1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Reviewer1) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *Reviewer1) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Reviewer1) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Reviewer1) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Reviewer1) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *Reviewer1) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *Reviewer1) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


