# Owner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**DtoType**](DtoType.md) |  | [optional] 
**Id** | Pointer to **string** | ID of the object to which this reference applies | [optional] 
**Name** | Pointer to **string** | Human-readable name of the owner | [optional] 
**DisplayName** | Pointer to **string** | Human-readable display name of the owner | [optional] 
**EmailAddress** | Pointer to **string** | Email ID of the owner | [optional] 

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

### GetDisplayName

`func (o *Owner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Owner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Owner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Owner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmailAddress

`func (o *Owner) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *Owner) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *Owner) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *Owner) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


