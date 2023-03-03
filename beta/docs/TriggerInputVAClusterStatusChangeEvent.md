# TriggerInputVAClusterStatusChangeEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** | The date and time the status change occurred. | 
**Type** | **map[string]interface{}** | The type of the object that initiated this event. | 
**Application** | [**TriggerInputVAClusterStatusChangeEventApplication**](TriggerInputVAClusterStatusChangeEventApplication.md) |  | 
**HealthCheckResult** | [**TriggerInputVAClusterStatusChangeEventHealthCheckResult**](TriggerInputVAClusterStatusChangeEventHealthCheckResult.md) |  | 
**PreviousHealthCheckResult** | [**TriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult**](TriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult.md) |  | 

## Methods

### NewTriggerInputVAClusterStatusChangeEvent

`func NewTriggerInputVAClusterStatusChangeEvent(created time.Time, type_ map[string]interface{}, application TriggerInputVAClusterStatusChangeEventApplication, healthCheckResult TriggerInputVAClusterStatusChangeEventHealthCheckResult, previousHealthCheckResult TriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult, ) *TriggerInputVAClusterStatusChangeEvent`

NewTriggerInputVAClusterStatusChangeEvent instantiates a new TriggerInputVAClusterStatusChangeEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerInputVAClusterStatusChangeEventWithDefaults

`func NewTriggerInputVAClusterStatusChangeEventWithDefaults() *TriggerInputVAClusterStatusChangeEvent`

NewTriggerInputVAClusterStatusChangeEventWithDefaults instantiates a new TriggerInputVAClusterStatusChangeEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *TriggerInputVAClusterStatusChangeEvent) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TriggerInputVAClusterStatusChangeEvent) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TriggerInputVAClusterStatusChangeEvent) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetType

`func (o *TriggerInputVAClusterStatusChangeEvent) GetType() map[string]interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TriggerInputVAClusterStatusChangeEvent) GetTypeOk() (*map[string]interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TriggerInputVAClusterStatusChangeEvent) SetType(v map[string]interface{})`

SetType sets Type field to given value.


### GetApplication

`func (o *TriggerInputVAClusterStatusChangeEvent) GetApplication() TriggerInputVAClusterStatusChangeEventApplication`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *TriggerInputVAClusterStatusChangeEvent) GetApplicationOk() (*TriggerInputVAClusterStatusChangeEventApplication, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *TriggerInputVAClusterStatusChangeEvent) SetApplication(v TriggerInputVAClusterStatusChangeEventApplication)`

SetApplication sets Application field to given value.


### GetHealthCheckResult

`func (o *TriggerInputVAClusterStatusChangeEvent) GetHealthCheckResult() TriggerInputVAClusterStatusChangeEventHealthCheckResult`

GetHealthCheckResult returns the HealthCheckResult field if non-nil, zero value otherwise.

### GetHealthCheckResultOk

`func (o *TriggerInputVAClusterStatusChangeEvent) GetHealthCheckResultOk() (*TriggerInputVAClusterStatusChangeEventHealthCheckResult, bool)`

GetHealthCheckResultOk returns a tuple with the HealthCheckResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckResult

`func (o *TriggerInputVAClusterStatusChangeEvent) SetHealthCheckResult(v TriggerInputVAClusterStatusChangeEventHealthCheckResult)`

SetHealthCheckResult sets HealthCheckResult field to given value.


### GetPreviousHealthCheckResult

`func (o *TriggerInputVAClusterStatusChangeEvent) GetPreviousHealthCheckResult() TriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult`

GetPreviousHealthCheckResult returns the PreviousHealthCheckResult field if non-nil, zero value otherwise.

### GetPreviousHealthCheckResultOk

`func (o *TriggerInputVAClusterStatusChangeEvent) GetPreviousHealthCheckResultOk() (*TriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult, bool)`

GetPreviousHealthCheckResultOk returns a tuple with the PreviousHealthCheckResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousHealthCheckResult

`func (o *TriggerInputVAClusterStatusChangeEvent) SetPreviousHealthCheckResult(v TriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult)`

SetPreviousHealthCheckResult sets PreviousHealthCheckResult field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


