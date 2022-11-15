# Schedule1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ScheduleType**](ScheduleType.md) |  | 
**Months** | Pointer to [**NullableSelector**](Selector.md) |  | [optional] 
**Days** | Pointer to [**NullableSelector**](Selector.md) |  | [optional] 
**Hours** | [**NullableSelector**](Selector.md) |  | 
**Expiration** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**TimeZoneId** | Pointer to **string** | The ID of the time zone for the schedule.  | [optional] 

## Methods

### NewSchedule1

`func NewSchedule1(type_ ScheduleType, hours NullableSelector, ) *Schedule1`

NewSchedule1 instantiates a new Schedule1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedule1WithDefaults

`func NewSchedule1WithDefaults() *Schedule1`

NewSchedule1WithDefaults instantiates a new Schedule1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Schedule1) GetType() ScheduleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Schedule1) GetTypeOk() (*ScheduleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Schedule1) SetType(v ScheduleType)`

SetType sets Type field to given value.


### GetMonths

`func (o *Schedule1) GetMonths() Selector`

GetMonths returns the Months field if non-nil, zero value otherwise.

### GetMonthsOk

`func (o *Schedule1) GetMonthsOk() (*Selector, bool)`

GetMonthsOk returns a tuple with the Months field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonths

`func (o *Schedule1) SetMonths(v Selector)`

SetMonths sets Months field to given value.

### HasMonths

`func (o *Schedule1) HasMonths() bool`

HasMonths returns a boolean if a field has been set.

### SetMonthsNil

`func (o *Schedule1) SetMonthsNil(b bool)`

 SetMonthsNil sets the value for Months to be an explicit nil

### UnsetMonths
`func (o *Schedule1) UnsetMonths()`

UnsetMonths ensures that no value is present for Months, not even an explicit nil
### GetDays

`func (o *Schedule1) GetDays() Selector`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *Schedule1) GetDaysOk() (*Selector, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *Schedule1) SetDays(v Selector)`

SetDays sets Days field to given value.

### HasDays

`func (o *Schedule1) HasDays() bool`

HasDays returns a boolean if a field has been set.

### SetDaysNil

`func (o *Schedule1) SetDaysNil(b bool)`

 SetDaysNil sets the value for Days to be an explicit nil

### UnsetDays
`func (o *Schedule1) UnsetDays()`

UnsetDays ensures that no value is present for Days, not even an explicit nil
### GetHours

`func (o *Schedule1) GetHours() Selector`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *Schedule1) GetHoursOk() (*Selector, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *Schedule1) SetHours(v Selector)`

SetHours sets Hours field to given value.


### SetHoursNil

`func (o *Schedule1) SetHoursNil(b bool)`

 SetHoursNil sets the value for Hours to be an explicit nil

### UnsetHours
`func (o *Schedule1) UnsetHours()`

UnsetHours ensures that no value is present for Hours, not even an explicit nil
### GetExpiration

`func (o *Schedule1) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *Schedule1) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *Schedule1) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *Schedule1) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### SetExpirationNil

`func (o *Schedule1) SetExpirationNil(b bool)`

 SetExpirationNil sets the value for Expiration to be an explicit nil

### UnsetExpiration
`func (o *Schedule1) UnsetExpiration()`

UnsetExpiration ensures that no value is present for Expiration, not even an explicit nil
### GetTimeZoneId

`func (o *Schedule1) GetTimeZoneId() string`

GetTimeZoneId returns the TimeZoneId field if non-nil, zero value otherwise.

### GetTimeZoneIdOk

`func (o *Schedule1) GetTimeZoneIdOk() (*string, bool)`

GetTimeZoneIdOk returns a tuple with the TimeZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneId

`func (o *Schedule1) SetTimeZoneId(v string)`

SetTimeZoneId sets TimeZoneId field to given value.

### HasTimeZoneId

`func (o *Schedule1) HasTimeZoneId() bool`

HasTimeZoneId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


