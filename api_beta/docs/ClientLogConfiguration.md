# ClientLogConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Log configuration&#39;s client ID | [optional] 
**DurationMinutes** | **int32** | Duration in minutes for log configuration to remain in effect before resetting to defaults | 
**Expiration** | Pointer to **time.Time** | Expiration date-time of the log configuration request | [optional] 
**RootLevel** | [**StandardLevel**](StandardLevel.md) |  | 
**LogLevels** | Pointer to [**map[string]StandardLevel**](StandardLevel.md) | Mapping of identifiers to Standard Log Level values | [optional] 

## Methods

### NewClientLogConfiguration

`func NewClientLogConfiguration(durationMinutes int32, rootLevel StandardLevel, ) *ClientLogConfiguration`

NewClientLogConfiguration instantiates a new ClientLogConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientLogConfigurationWithDefaults

`func NewClientLogConfigurationWithDefaults() *ClientLogConfiguration`

NewClientLogConfigurationWithDefaults instantiates a new ClientLogConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *ClientLogConfiguration) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ClientLogConfiguration) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ClientLogConfiguration) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ClientLogConfiguration) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetDurationMinutes

`func (o *ClientLogConfiguration) GetDurationMinutes() int32`

GetDurationMinutes returns the DurationMinutes field if non-nil, zero value otherwise.

### GetDurationMinutesOk

`func (o *ClientLogConfiguration) GetDurationMinutesOk() (*int32, bool)`

GetDurationMinutesOk returns a tuple with the DurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMinutes

`func (o *ClientLogConfiguration) SetDurationMinutes(v int32)`

SetDurationMinutes sets DurationMinutes field to given value.


### GetExpiration

`func (o *ClientLogConfiguration) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *ClientLogConfiguration) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *ClientLogConfiguration) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *ClientLogConfiguration) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetRootLevel

`func (o *ClientLogConfiguration) GetRootLevel() StandardLevel`

GetRootLevel returns the RootLevel field if non-nil, zero value otherwise.

### GetRootLevelOk

`func (o *ClientLogConfiguration) GetRootLevelOk() (*StandardLevel, bool)`

GetRootLevelOk returns a tuple with the RootLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootLevel

`func (o *ClientLogConfiguration) SetRootLevel(v StandardLevel)`

SetRootLevel sets RootLevel field to given value.


### GetLogLevels

`func (o *ClientLogConfiguration) GetLogLevels() map[string]StandardLevel`

GetLogLevels returns the LogLevels field if non-nil, zero value otherwise.

### GetLogLevelsOk

`func (o *ClientLogConfiguration) GetLogLevelsOk() (*map[string]StandardLevel, bool)`

GetLogLevelsOk returns a tuple with the LogLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevels

`func (o *ClientLogConfiguration) SetLogLevels(v map[string]StandardLevel)`

SetLogLevels sets LogLevels field to given value.

### HasLogLevels

`func (o *ClientLogConfiguration) HasLogLevels() bool`

HasLogLevels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


