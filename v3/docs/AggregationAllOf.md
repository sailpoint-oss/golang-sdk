# AggregationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **int32** |  | [optional] 
**AvgDuration** | Pointer to **int32** |  | [optional] 
**ChangedAccounts** | Pointer to **int32** |  | [optional] 
**NextScheduled** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**StartTime** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**SourceOwner** | Pointer to **string** | John Doe | [optional] 

## Methods

### NewAggregationAllOf

`func NewAggregationAllOf() *AggregationAllOf`

NewAggregationAllOf instantiates a new AggregationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregationAllOfWithDefaults

`func NewAggregationAllOfWithDefaults() *AggregationAllOf`

NewAggregationAllOfWithDefaults instantiates a new AggregationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AggregationAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AggregationAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AggregationAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AggregationAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDuration

`func (o *AggregationAllOf) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *AggregationAllOf) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *AggregationAllOf) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *AggregationAllOf) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetAvgDuration

`func (o *AggregationAllOf) GetAvgDuration() int32`

GetAvgDuration returns the AvgDuration field if non-nil, zero value otherwise.

### GetAvgDurationOk

`func (o *AggregationAllOf) GetAvgDurationOk() (*int32, bool)`

GetAvgDurationOk returns a tuple with the AvgDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgDuration

`func (o *AggregationAllOf) SetAvgDuration(v int32)`

SetAvgDuration sets AvgDuration field to given value.

### HasAvgDuration

`func (o *AggregationAllOf) HasAvgDuration() bool`

HasAvgDuration returns a boolean if a field has been set.

### GetChangedAccounts

`func (o *AggregationAllOf) GetChangedAccounts() int32`

GetChangedAccounts returns the ChangedAccounts field if non-nil, zero value otherwise.

### GetChangedAccountsOk

`func (o *AggregationAllOf) GetChangedAccountsOk() (*int32, bool)`

GetChangedAccountsOk returns a tuple with the ChangedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedAccounts

`func (o *AggregationAllOf) SetChangedAccounts(v int32)`

SetChangedAccounts sets ChangedAccounts field to given value.

### HasChangedAccounts

`func (o *AggregationAllOf) HasChangedAccounts() bool`

HasChangedAccounts returns a boolean if a field has been set.

### GetNextScheduled

`func (o *AggregationAllOf) GetNextScheduled() time.Time`

GetNextScheduled returns the NextScheduled field if non-nil, zero value otherwise.

### GetNextScheduledOk

`func (o *AggregationAllOf) GetNextScheduledOk() (*time.Time, bool)`

GetNextScheduledOk returns a tuple with the NextScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextScheduled

`func (o *AggregationAllOf) SetNextScheduled(v time.Time)`

SetNextScheduled sets NextScheduled field to given value.

### HasNextScheduled

`func (o *AggregationAllOf) HasNextScheduled() bool`

HasNextScheduled returns a boolean if a field has been set.

### SetNextScheduledNil

`func (o *AggregationAllOf) SetNextScheduledNil(b bool)`

 SetNextScheduledNil sets the value for NextScheduled to be an explicit nil

### UnsetNextScheduled
`func (o *AggregationAllOf) UnsetNextScheduled()`

UnsetNextScheduled ensures that no value is present for NextScheduled, not even an explicit nil
### GetStartTime

`func (o *AggregationAllOf) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *AggregationAllOf) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *AggregationAllOf) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *AggregationAllOf) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *AggregationAllOf) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *AggregationAllOf) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetSourceOwner

`func (o *AggregationAllOf) GetSourceOwner() string`

GetSourceOwner returns the SourceOwner field if non-nil, zero value otherwise.

### GetSourceOwnerOk

`func (o *AggregationAllOf) GetSourceOwnerOk() (*string, bool)`

GetSourceOwnerOk returns a tuple with the SourceOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOwner

`func (o *AggregationAllOf) SetSourceOwner(v string)`

SetSourceOwner sets SourceOwner field to given value.

### HasSourceOwner

`func (o *AggregationAllOf) HasSourceOwner() bool`

HasSourceOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


