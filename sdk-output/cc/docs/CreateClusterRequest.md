# CreateClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**GmtOffset** | Pointer to **float32** |  | [optional] 

## Methods

### NewCreateClusterRequest

`func NewCreateClusterRequest() *CreateClusterRequest`

NewCreateClusterRequest instantiates a new CreateClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClusterRequestWithDefaults

`func NewCreateClusterRequestWithDefaults() *CreateClusterRequest`

NewCreateClusterRequestWithDefaults instantiates a new CreateClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateClusterRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateClusterRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateClusterRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateClusterRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetGmtOffset

`func (o *CreateClusterRequest) GetGmtOffset() float32`

GetGmtOffset returns the GmtOffset field if non-nil, zero value otherwise.

### GetGmtOffsetOk

`func (o *CreateClusterRequest) GetGmtOffsetOk() (*float32, bool)`

GetGmtOffsetOk returns a tuple with the GmtOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtOffset

`func (o *CreateClusterRequest) SetGmtOffset(v float32)`

SetGmtOffset sets GmtOffset field to given value.

### HasGmtOffset

`func (o *CreateClusterRequest) HasGmtOffset() bool`

HasGmtOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


