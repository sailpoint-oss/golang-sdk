# Aggregation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Type** | [**DocumentType**](DocumentType.md) |  | 
**Status** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **int32** |  | [optional] 
**AvgDuration** | Pointer to **int32** |  | [optional] 
**ChangedAccounts** | Pointer to **int32** |  | [optional] 
**NextScheduled** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**StartTime** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**SourceOwner** | Pointer to **string** | John Doe | [optional] 

## Methods

### NewAggregation

`func NewAggregation(id string, name string, type_ DocumentType, ) *Aggregation`

NewAggregation instantiates a new Aggregation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregationWithDefaults

`func NewAggregationWithDefaults() *Aggregation`

NewAggregationWithDefaults instantiates a new Aggregation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Aggregation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Aggregation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Aggregation) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Aggregation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Aggregation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Aggregation) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *Aggregation) GetType() DocumentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Aggregation) GetTypeOk() (*DocumentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Aggregation) SetType(v DocumentType)`

SetType sets Type field to given value.


### GetStatus

`func (o *Aggregation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Aggregation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Aggregation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Aggregation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDuration

`func (o *Aggregation) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Aggregation) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Aggregation) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *Aggregation) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetAvgDuration

`func (o *Aggregation) GetAvgDuration() int32`

GetAvgDuration returns the AvgDuration field if non-nil, zero value otherwise.

### GetAvgDurationOk

`func (o *Aggregation) GetAvgDurationOk() (*int32, bool)`

GetAvgDurationOk returns a tuple with the AvgDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgDuration

`func (o *Aggregation) SetAvgDuration(v int32)`

SetAvgDuration sets AvgDuration field to given value.

### HasAvgDuration

`func (o *Aggregation) HasAvgDuration() bool`

HasAvgDuration returns a boolean if a field has been set.

### GetChangedAccounts

`func (o *Aggregation) GetChangedAccounts() int32`

GetChangedAccounts returns the ChangedAccounts field if non-nil, zero value otherwise.

### GetChangedAccountsOk

`func (o *Aggregation) GetChangedAccountsOk() (*int32, bool)`

GetChangedAccountsOk returns a tuple with the ChangedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedAccounts

`func (o *Aggregation) SetChangedAccounts(v int32)`

SetChangedAccounts sets ChangedAccounts field to given value.

### HasChangedAccounts

`func (o *Aggregation) HasChangedAccounts() bool`

HasChangedAccounts returns a boolean if a field has been set.

### GetNextScheduled

`func (o *Aggregation) GetNextScheduled() time.Time`

GetNextScheduled returns the NextScheduled field if non-nil, zero value otherwise.

### GetNextScheduledOk

`func (o *Aggregation) GetNextScheduledOk() (*time.Time, bool)`

GetNextScheduledOk returns a tuple with the NextScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextScheduled

`func (o *Aggregation) SetNextScheduled(v time.Time)`

SetNextScheduled sets NextScheduled field to given value.

### HasNextScheduled

`func (o *Aggregation) HasNextScheduled() bool`

HasNextScheduled returns a boolean if a field has been set.

### SetNextScheduledNil

`func (o *Aggregation) SetNextScheduledNil(b bool)`

 SetNextScheduledNil sets the value for NextScheduled to be an explicit nil

### UnsetNextScheduled
`func (o *Aggregation) UnsetNextScheduled()`

UnsetNextScheduled ensures that no value is present for NextScheduled, not even an explicit nil
### GetStartTime

`func (o *Aggregation) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Aggregation) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Aggregation) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *Aggregation) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *Aggregation) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *Aggregation) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetSourceOwner

`func (o *Aggregation) GetSourceOwner() string`

GetSourceOwner returns the SourceOwner field if non-nil, zero value otherwise.

### GetSourceOwnerOk

`func (o *Aggregation) GetSourceOwnerOk() (*string, bool)`

GetSourceOwnerOk returns a tuple with the SourceOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOwner

`func (o *Aggregation) SetSourceOwner(v string)`

SetSourceOwner sets SourceOwner field to given value.

### HasSourceOwner

`func (o *Aggregation) HasSourceOwner() bool`

HasSourceOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


