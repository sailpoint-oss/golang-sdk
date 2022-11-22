# Transform

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique ID of this transform | [optional] [readonly] 
**Name** | **string** | Unique name of this transform | 
**Type** | **string** | The type of transform operation | 
**Attributes** | Pointer to [**TransformAttributes**](TransformAttributes.md) |  | [optional] 
**Internal** | Pointer to **bool** | Indicates whether this is an internal SailPoint-created transform or a customer-created transform | [optional] [readonly] 

## Methods

### NewTransform

`func NewTransform(name string, type_ string, ) *Transform`

NewTransform instantiates a new Transform object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransformWithDefaults

`func NewTransformWithDefaults() *Transform`

NewTransformWithDefaults instantiates a new Transform object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Transform) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Transform) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Transform) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Transform) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Transform) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Transform) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Transform) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *Transform) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Transform) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Transform) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *Transform) GetAttributes() TransformAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Transform) GetAttributesOk() (*TransformAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Transform) SetAttributes(v TransformAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Transform) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetInternal

`func (o *Transform) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *Transform) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *Transform) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *Transform) HasInternal() bool`

HasInternal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


