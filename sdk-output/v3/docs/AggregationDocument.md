# AggregationDocument

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

### NewAggregationDocument

`func NewAggregationDocument(id string, name string, type_ DocumentType, ) *AggregationDocument`

NewAggregationDocument instantiates a new AggregationDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregationDocumentWithDefaults

`func NewAggregationDocumentWithDefaults() *AggregationDocument`

NewAggregationDocumentWithDefaults instantiates a new AggregationDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AggregationDocument) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AggregationDocument) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AggregationDocument) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AggregationDocument) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AggregationDocument) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AggregationDocument) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *AggregationDocument) GetType() DocumentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AggregationDocument) GetTypeOk() (*DocumentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AggregationDocument) SetType(v DocumentType)`

SetType sets Type field to given value.


### GetStatus

`func (o *AggregationDocument) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AggregationDocument) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AggregationDocument) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AggregationDocument) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDuration

`func (o *AggregationDocument) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *AggregationDocument) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *AggregationDocument) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *AggregationDocument) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetAvgDuration

`func (o *AggregationDocument) GetAvgDuration() int32`

GetAvgDuration returns the AvgDuration field if non-nil, zero value otherwise.

### GetAvgDurationOk

`func (o *AggregationDocument) GetAvgDurationOk() (*int32, bool)`

GetAvgDurationOk returns a tuple with the AvgDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgDuration

`func (o *AggregationDocument) SetAvgDuration(v int32)`

SetAvgDuration sets AvgDuration field to given value.

### HasAvgDuration

`func (o *AggregationDocument) HasAvgDuration() bool`

HasAvgDuration returns a boolean if a field has been set.

### GetChangedAccounts

`func (o *AggregationDocument) GetChangedAccounts() int32`

GetChangedAccounts returns the ChangedAccounts field if non-nil, zero value otherwise.

### GetChangedAccountsOk

`func (o *AggregationDocument) GetChangedAccountsOk() (*int32, bool)`

GetChangedAccountsOk returns a tuple with the ChangedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedAccounts

`func (o *AggregationDocument) SetChangedAccounts(v int32)`

SetChangedAccounts sets ChangedAccounts field to given value.

### HasChangedAccounts

`func (o *AggregationDocument) HasChangedAccounts() bool`

HasChangedAccounts returns a boolean if a field has been set.

### GetNextScheduled

`func (o *AggregationDocument) GetNextScheduled() time.Time`

GetNextScheduled returns the NextScheduled field if non-nil, zero value otherwise.

### GetNextScheduledOk

`func (o *AggregationDocument) GetNextScheduledOk() (*time.Time, bool)`

GetNextScheduledOk returns a tuple with the NextScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextScheduled

`func (o *AggregationDocument) SetNextScheduled(v time.Time)`

SetNextScheduled sets NextScheduled field to given value.

### HasNextScheduled

`func (o *AggregationDocument) HasNextScheduled() bool`

HasNextScheduled returns a boolean if a field has been set.

### SetNextScheduledNil

`func (o *AggregationDocument) SetNextScheduledNil(b bool)`

 SetNextScheduledNil sets the value for NextScheduled to be an explicit nil

### UnsetNextScheduled
`func (o *AggregationDocument) UnsetNextScheduled()`

UnsetNextScheduled ensures that no value is present for NextScheduled, not even an explicit nil
### GetStartTime

`func (o *AggregationDocument) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *AggregationDocument) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *AggregationDocument) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *AggregationDocument) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *AggregationDocument) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *AggregationDocument) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetSourceOwner

`func (o *AggregationDocument) GetSourceOwner() string`

GetSourceOwner returns the SourceOwner field if non-nil, zero value otherwise.

### GetSourceOwnerOk

`func (o *AggregationDocument) GetSourceOwnerOk() (*string, bool)`

GetSourceOwnerOk returns a tuple with the SourceOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOwner

`func (o *AggregationDocument) SetSourceOwner(v string)`

SetSourceOwner sets SourceOwner field to given value.

### HasSourceOwner

`func (o *AggregationDocument) HasSourceOwner() bool`

HasSourceOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


