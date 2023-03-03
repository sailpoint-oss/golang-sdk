# TriggerInputSourceCreated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID of the source. | 
**Name** | **string** | Human friendly name of the source. | 
**Type** | **string** | The connection type. | 
**Created** | **time.Time** | The date and time the source was created. | 
**Connector** | **string** | The connector type used to connect to the source. | 
**Actor** | [**TriggerInputSourceCreatedActor**](TriggerInputSourceCreatedActor.md) |  | 

## Methods

### NewTriggerInputSourceCreated

`func NewTriggerInputSourceCreated(id string, name string, type_ string, created time.Time, connector string, actor TriggerInputSourceCreatedActor, ) *TriggerInputSourceCreated`

NewTriggerInputSourceCreated instantiates a new TriggerInputSourceCreated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerInputSourceCreatedWithDefaults

`func NewTriggerInputSourceCreatedWithDefaults() *TriggerInputSourceCreated`

NewTriggerInputSourceCreatedWithDefaults instantiates a new TriggerInputSourceCreated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TriggerInputSourceCreated) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TriggerInputSourceCreated) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TriggerInputSourceCreated) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TriggerInputSourceCreated) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TriggerInputSourceCreated) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TriggerInputSourceCreated) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *TriggerInputSourceCreated) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TriggerInputSourceCreated) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TriggerInputSourceCreated) SetType(v string)`

SetType sets Type field to given value.


### GetCreated

`func (o *TriggerInputSourceCreated) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TriggerInputSourceCreated) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TriggerInputSourceCreated) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetConnector

`func (o *TriggerInputSourceCreated) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *TriggerInputSourceCreated) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *TriggerInputSourceCreated) SetConnector(v string)`

SetConnector sets Connector field to given value.


### GetActor

`func (o *TriggerInputSourceCreated) GetActor() TriggerInputSourceCreatedActor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *TriggerInputSourceCreated) GetActorOk() (*TriggerInputSourceCreatedActor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *TriggerInputSourceCreated) SetActor(v TriggerInputSourceCreatedActor)`

SetActor sets Actor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


