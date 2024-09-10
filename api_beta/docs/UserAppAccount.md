# UserAppAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | the account ID | [optional] 
**Type** | Pointer to **string** | It will always be \&quot;ACCOUNT\&quot; | [optional] 
**Name** | Pointer to **string** | the account name | [optional] 

## Methods

### NewUserAppAccount

`func NewUserAppAccount() *UserAppAccount`

NewUserAppAccount instantiates a new UserAppAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAppAccountWithDefaults

`func NewUserAppAccountWithDefaults() *UserAppAccount`

NewUserAppAccountWithDefaults instantiates a new UserAppAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserAppAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserAppAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserAppAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserAppAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UserAppAccount) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserAppAccount) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserAppAccount) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserAppAccount) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *UserAppAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserAppAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserAppAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserAppAccount) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

