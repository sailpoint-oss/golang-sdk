# TransformReadAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID of this transform | 
**Internal** | **bool** | Indicates whether this is an internal SailPoint-created transform or a customer-created transform | [default to false]

## Methods

### NewTransformReadAllOf

`func NewTransformReadAllOf(id string, internal bool, ) *TransformReadAllOf`

NewTransformReadAllOf instantiates a new TransformReadAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransformReadAllOfWithDefaults

`func NewTransformReadAllOfWithDefaults() *TransformReadAllOf`

NewTransformReadAllOfWithDefaults instantiates a new TransformReadAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TransformReadAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransformReadAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransformReadAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetInternal

`func (o *TransformReadAllOf) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *TransformReadAllOf) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *TransformReadAllOf) SetInternal(v bool)`

SetInternal sets Internal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


