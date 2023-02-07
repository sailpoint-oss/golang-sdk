# ScheduleDays

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**SelectorType**](SelectorType.md) |  | 
**Values** | **[]string** | The selected values.  | 
**Interval** | Pointer to **NullableInt32** | The selected interval for RANGE selectors.  | [optional] 

## Methods

### NewScheduleDays

`func NewScheduleDays(type_ SelectorType, values []string, ) *ScheduleDays`

NewScheduleDays instantiates a new ScheduleDays object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleDaysWithDefaults

`func NewScheduleDaysWithDefaults() *ScheduleDays`

NewScheduleDaysWithDefaults instantiates a new ScheduleDays object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ScheduleDays) GetType() SelectorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScheduleDays) GetTypeOk() (*SelectorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScheduleDays) SetType(v SelectorType)`

SetType sets Type field to given value.


### GetValues

`func (o *ScheduleDays) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ScheduleDays) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ScheduleDays) SetValues(v []string)`

SetValues sets Values field to given value.


### GetInterval

`func (o *ScheduleDays) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *ScheduleDays) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *ScheduleDays) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *ScheduleDays) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### SetIntervalNil

`func (o *ScheduleDays) SetIntervalNil(b bool)`

 SetIntervalNil sets the value for Interval to be an explicit nil

### UnsetInterval
`func (o *ScheduleDays) UnsetInterval()`

UnsetInterval ensures that no value is present for Interval, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


