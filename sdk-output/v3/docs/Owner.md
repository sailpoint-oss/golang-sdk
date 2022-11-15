# Owner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**DtoType**](DtoType.md) |  | [optional] 
**Email** | Pointer to **string** | The email of the identity | [optional] 

## Methods

### NewOwner

`func NewOwner() *Owner`

NewOwner instantiates a new Owner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOwnerWithDefaults

`func NewOwnerWithDefaults() *Owner`

NewOwnerWithDefaults instantiates a new Owner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Owner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Owner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Owner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Owner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Owner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Owner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Owner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Owner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *Owner) GetType() DtoType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Owner) GetTypeOk() (*DtoType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Owner) SetType(v DtoType)`

SetType sets Type field to given value.

### HasType

`func (o *Owner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEmail

`func (o *Owner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Owner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Owner) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Owner) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


