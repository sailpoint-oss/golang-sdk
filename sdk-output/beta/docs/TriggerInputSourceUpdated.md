# TriggerInputSourceUpdated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID of the source. | 
**Name** | **string** | The user friendly name of the source. | 
**Type** | **string** | The connection type of the source. | 
**Modified** | **time.Time** | The date and time the source was modified. | 
**Connector** | **string** | The connector type used to connect to the source. | 
**Actor** | [**TriggerInputSourceUpdatedActor**](TriggerInputSourceUpdatedActor.md) |  | 

## Methods

### NewTriggerInputSourceUpdated

`func NewTriggerInputSourceUpdated(id string, name string, type_ string, modified time.Time, connector string, actor TriggerInputSourceUpdatedActor, ) *TriggerInputSourceUpdated`

NewTriggerInputSourceUpdated instantiates a new TriggerInputSourceUpdated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerInputSourceUpdatedWithDefaults

`func NewTriggerInputSourceUpdatedWithDefaults() *TriggerInputSourceUpdated`

NewTriggerInputSourceUpdatedWithDefaults instantiates a new TriggerInputSourceUpdated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TriggerInputSourceUpdated) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TriggerInputSourceUpdated) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TriggerInputSourceUpdated) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TriggerInputSourceUpdated) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TriggerInputSourceUpdated) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TriggerInputSourceUpdated) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *TriggerInputSourceUpdated) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TriggerInputSourceUpdated) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TriggerInputSourceUpdated) SetType(v string)`

SetType sets Type field to given value.


### GetModified

`func (o *TriggerInputSourceUpdated) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *TriggerInputSourceUpdated) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *TriggerInputSourceUpdated) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetConnector

`func (o *TriggerInputSourceUpdated) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *TriggerInputSourceUpdated) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *TriggerInputSourceUpdated) SetConnector(v string)`

SetConnector sets Connector field to given value.


### GetActor

`func (o *TriggerInputSourceUpdated) GetActor() TriggerInputSourceUpdatedActor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *TriggerInputSourceUpdated) GetActorOk() (*TriggerInputSourceUpdatedActor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *TriggerInputSourceUpdated) SetActor(v TriggerInputSourceUpdatedActor)`

SetActor sets Actor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


