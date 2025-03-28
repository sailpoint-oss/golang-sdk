# Ref

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**DtoType**](DtoType.md) |  | [optional] 
**Id** | Pointer to **string** | ID of the object to which this reference applies | [optional] 

## Methods

### NewRef

`func NewRef() *Ref`

NewRef instantiates a new Ref object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefWithDefaults

`func NewRefWithDefaults() *Ref`

NewRefWithDefaults instantiates a new Ref object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Ref) GetType() DtoType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Ref) GetTypeOk() (*DtoType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Ref) SetType(v DtoType)`

SetType sets Type field to given value.

### HasType

`func (o *Ref) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *Ref) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ref) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Ref) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Ref) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


