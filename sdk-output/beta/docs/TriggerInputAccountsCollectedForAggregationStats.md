# TriggerInputAccountsCollectedForAggregationStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scanned** | **int32** | The number of accounts which were scanned / iterated over. | 
**Unchanged** | **int32** | The number of accounts which existed before, but had no changes. | 
**Changed** | **int32** | The number of accounts which existed before, but had changes. | 
**Added** | **int32** | The number of accounts which are new - have not existed before. | 
**Removed** | **int32** | The number accounts which existed before, but no longer exist (thus getting removed). | 

## Methods

### NewTriggerInputAccountsCollectedForAggregationStats

`func NewTriggerInputAccountsCollectedForAggregationStats(scanned int32, unchanged int32, changed int32, added int32, removed int32, ) *TriggerInputAccountsCollectedForAggregationStats`

NewTriggerInputAccountsCollectedForAggregationStats instantiates a new TriggerInputAccountsCollectedForAggregationStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerInputAccountsCollectedForAggregationStatsWithDefaults

`func NewTriggerInputAccountsCollectedForAggregationStatsWithDefaults() *TriggerInputAccountsCollectedForAggregationStats`

NewTriggerInputAccountsCollectedForAggregationStatsWithDefaults instantiates a new TriggerInputAccountsCollectedForAggregationStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScanned

`func (o *TriggerInputAccountsCollectedForAggregationStats) GetScanned() int32`

GetScanned returns the Scanned field if non-nil, zero value otherwise.

### GetScannedOk

`func (o *TriggerInputAccountsCollectedForAggregationStats) GetScannedOk() (*int32, bool)`

GetScannedOk returns a tuple with the Scanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanned

`func (o *TriggerInputAccountsCollectedForAggregationStats) SetScanned(v int32)`

SetScanned sets Scanned field to given value.


### GetUnchanged

`func (o *TriggerInputAccountsCollectedForAggregationStats) GetUnchanged() int32`

GetUnchanged returns the Unchanged field if non-nil, zero value otherwise.

### GetUnchangedOk

`func (o *TriggerInputAccountsCollectedForAggregationStats) GetUnchangedOk() (*int32, bool)`

GetUnchangedOk returns a tuple with the Unchanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnchanged

`func (o *TriggerInputAccountsCollectedForAggregationStats) SetUnchanged(v int32)`

SetUnchanged sets Unchanged field to given value.


### GetChanged

`func (o *TriggerInputAccountsCollectedForAggregationStats) GetChanged() int32`

GetChanged returns the Changed field if non-nil, zero value otherwise.

### GetChangedOk

`func (o *TriggerInputAccountsCollectedForAggregationStats) GetChangedOk() (*int32, bool)`

GetChangedOk returns a tuple with the Changed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanged

`func (o *TriggerInputAccountsCollectedForAggregationStats) SetChanged(v int32)`

SetChanged sets Changed field to given value.


### GetAdded

`func (o *TriggerInputAccountsCollectedForAggregationStats) GetAdded() int32`

GetAdded returns the Added field if non-nil, zero value otherwise.

### GetAddedOk

`func (o *TriggerInputAccountsCollectedForAggregationStats) GetAddedOk() (*int32, bool)`

GetAddedOk returns a tuple with the Added field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdded

`func (o *TriggerInputAccountsCollectedForAggregationStats) SetAdded(v int32)`

SetAdded sets Added field to given value.


### GetRemoved

`func (o *TriggerInputAccountsCollectedForAggregationStats) GetRemoved() int32`

GetRemoved returns the Removed field if non-nil, zero value otherwise.

### GetRemovedOk

`func (o *TriggerInputAccountsCollectedForAggregationStats) GetRemovedOk() (*int32, bool)`

GetRemovedOk returns a tuple with the Removed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoved

`func (o *TriggerInputAccountsCollectedForAggregationStats) SetRemoved(v int32)`

SetRemoved sets Removed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


