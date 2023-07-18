# FormOwner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is a unique identifier | [optional] 
**Type** | Pointer to **string** | Type is a FormOwnerType value IDENTITY FormOwnerTypeIdentity | [optional] 

## Methods

### NewFormOwner

`func NewFormOwner() *FormOwner`

NewFormOwner instantiates a new FormOwner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormOwnerWithDefaults

`func NewFormOwnerWithDefaults() *FormOwner`

NewFormOwnerWithDefaults instantiates a new FormOwner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FormOwner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormOwner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormOwner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FormOwner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *FormOwner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormOwner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormOwner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FormOwner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


