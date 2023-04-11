# UpdateClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the cluster to update | [optional] 
**GmtOffset** | Pointer to **string** | The GMT offset for the timezone the cluster operates with | [optional] 

## Methods

### NewUpdateClusterRequest

`func NewUpdateClusterRequest() *UpdateClusterRequest`

NewUpdateClusterRequest instantiates a new UpdateClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClusterRequestWithDefaults

`func NewUpdateClusterRequestWithDefaults() *UpdateClusterRequest`

NewUpdateClusterRequestWithDefaults instantiates a new UpdateClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateClusterRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateClusterRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateClusterRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateClusterRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGmtOffset

`func (o *UpdateClusterRequest) GetGmtOffset() string`

GetGmtOffset returns the GmtOffset field if non-nil, zero value otherwise.

### GetGmtOffsetOk

`func (o *UpdateClusterRequest) GetGmtOffsetOk() (*string, bool)`

GetGmtOffsetOk returns a tuple with the GmtOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtOffset

`func (o *UpdateClusterRequest) SetGmtOffset(v string)`

SetGmtOffset sets GmtOffset field to given value.

### HasGmtOffset

`func (o *UpdateClusterRequest) HasGmtOffset() bool`

HasGmtOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


