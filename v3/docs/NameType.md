# NameType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | the actor or target name | [optional] 
**Type** | Pointer to [**DtoType**](DtoType.md) |  | [optional] 

## Methods

### NewNameType

`func NewNameType() *NameType`

NewNameType instantiates a new NameType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNameTypeWithDefaults

`func NewNameTypeWithDefaults() *NameType`

NewNameTypeWithDefaults instantiates a new NameType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NameType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NameType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NameType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NameType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *NameType) GetType() DtoType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NameType) GetTypeOk() (*DtoType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NameType) SetType(v DtoType)`

SetType sets Type field to given value.

### HasType

`func (o *NameType) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


