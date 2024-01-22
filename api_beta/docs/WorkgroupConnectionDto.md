# WorkgroupConnectionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to [**ConnectedObject**](ConnectedObject.md) |  | [optional] 
**ConnectionType** | Pointer to **string** | Connection Type. | [optional] 

## Methods

### NewWorkgroupConnectionDto

`func NewWorkgroupConnectionDto() *WorkgroupConnectionDto`

NewWorkgroupConnectionDto instantiates a new WorkgroupConnectionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkgroupConnectionDtoWithDefaults

`func NewWorkgroupConnectionDtoWithDefaults() *WorkgroupConnectionDto`

NewWorkgroupConnectionDtoWithDefaults instantiates a new WorkgroupConnectionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *WorkgroupConnectionDto) GetObject() ConnectedObject`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *WorkgroupConnectionDto) GetObjectOk() (*ConnectedObject, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *WorkgroupConnectionDto) SetObject(v ConnectedObject)`

SetObject sets Object field to given value.

### HasObject

`func (o *WorkgroupConnectionDto) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetConnectionType

`func (o *WorkgroupConnectionDto) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *WorkgroupConnectionDto) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *WorkgroupConnectionDto) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *WorkgroupConnectionDto) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


