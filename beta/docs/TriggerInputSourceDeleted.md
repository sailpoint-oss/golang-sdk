# TriggerInputSourceDeleted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID of the source. | 
**Name** | **string** | Human friendly name of the source. | 
**Type** | **string** | The connection type. | 
**Deleted** | **time.Time** | The date and time the source was deleted. | 
**Connector** | **string** | The connector type used to connect to the source. | 
**Actor** | [**TriggerInputSourceDeletedActor**](TriggerInputSourceDeletedActor.md) |  | 

## Methods

### NewTriggerInputSourceDeleted

`func NewTriggerInputSourceDeleted(id string, name string, type_ string, deleted time.Time, connector string, actor TriggerInputSourceDeletedActor, ) *TriggerInputSourceDeleted`

NewTriggerInputSourceDeleted instantiates a new TriggerInputSourceDeleted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerInputSourceDeletedWithDefaults

`func NewTriggerInputSourceDeletedWithDefaults() *TriggerInputSourceDeleted`

NewTriggerInputSourceDeletedWithDefaults instantiates a new TriggerInputSourceDeleted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TriggerInputSourceDeleted) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TriggerInputSourceDeleted) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TriggerInputSourceDeleted) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TriggerInputSourceDeleted) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TriggerInputSourceDeleted) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TriggerInputSourceDeleted) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *TriggerInputSourceDeleted) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TriggerInputSourceDeleted) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TriggerInputSourceDeleted) SetType(v string)`

SetType sets Type field to given value.


### GetDeleted

`func (o *TriggerInputSourceDeleted) GetDeleted() time.Time`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *TriggerInputSourceDeleted) GetDeletedOk() (*time.Time, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *TriggerInputSourceDeleted) SetDeleted(v time.Time)`

SetDeleted sets Deleted field to given value.


### GetConnector

`func (o *TriggerInputSourceDeleted) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *TriggerInputSourceDeleted) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *TriggerInputSourceDeleted) SetConnector(v string)`

SetConnector sets Connector field to given value.


### GetActor

`func (o *TriggerInputSourceDeleted) GetActor() TriggerInputSourceDeletedActor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *TriggerInputSourceDeleted) GetActorOk() (*TriggerInputSourceDeletedActor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *TriggerInputSourceDeleted) SetActor(v TriggerInputSourceDeletedActor)`

SetActor sets Actor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


