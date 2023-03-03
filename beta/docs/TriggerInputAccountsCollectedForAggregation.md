# TriggerInputAccountsCollectedForAggregation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**TriggerInputAccountsCollectedForAggregationSource**](TriggerInputAccountsCollectedForAggregationSource.md) |  | 
**Status** | **map[string]interface{}** | The overall status of the collection. | 
**Started** | **time.Time** | The date and time when the account collection started. | 
**Completed** | **time.Time** | The date and time when the account collection finished. | 
**Errors** | **[]string** | A list of errors that occurred during the collection. | 
**Warnings** | **[]string** | A list of warnings that occurred during the collection. | 
**Stats** | [**TriggerInputAccountsCollectedForAggregationStats**](TriggerInputAccountsCollectedForAggregationStats.md) |  | 

## Methods

### NewTriggerInputAccountsCollectedForAggregation

`func NewTriggerInputAccountsCollectedForAggregation(source TriggerInputAccountsCollectedForAggregationSource, status map[string]interface{}, started time.Time, completed time.Time, errors []string, warnings []string, stats TriggerInputAccountsCollectedForAggregationStats, ) *TriggerInputAccountsCollectedForAggregation`

NewTriggerInputAccountsCollectedForAggregation instantiates a new TriggerInputAccountsCollectedForAggregation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerInputAccountsCollectedForAggregationWithDefaults

`func NewTriggerInputAccountsCollectedForAggregationWithDefaults() *TriggerInputAccountsCollectedForAggregation`

NewTriggerInputAccountsCollectedForAggregationWithDefaults instantiates a new TriggerInputAccountsCollectedForAggregation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *TriggerInputAccountsCollectedForAggregation) GetSource() TriggerInputAccountsCollectedForAggregationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TriggerInputAccountsCollectedForAggregation) GetSourceOk() (*TriggerInputAccountsCollectedForAggregationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TriggerInputAccountsCollectedForAggregation) SetSource(v TriggerInputAccountsCollectedForAggregationSource)`

SetSource sets Source field to given value.


### GetStatus

`func (o *TriggerInputAccountsCollectedForAggregation) GetStatus() map[string]interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TriggerInputAccountsCollectedForAggregation) GetStatusOk() (*map[string]interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TriggerInputAccountsCollectedForAggregation) SetStatus(v map[string]interface{})`

SetStatus sets Status field to given value.


### GetStarted

`func (o *TriggerInputAccountsCollectedForAggregation) GetStarted() time.Time`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *TriggerInputAccountsCollectedForAggregation) GetStartedOk() (*time.Time, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *TriggerInputAccountsCollectedForAggregation) SetStarted(v time.Time)`

SetStarted sets Started field to given value.


### GetCompleted

`func (o *TriggerInputAccountsCollectedForAggregation) GetCompleted() time.Time`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *TriggerInputAccountsCollectedForAggregation) GetCompletedOk() (*time.Time, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *TriggerInputAccountsCollectedForAggregation) SetCompleted(v time.Time)`

SetCompleted sets Completed field to given value.


### GetErrors

`func (o *TriggerInputAccountsCollectedForAggregation) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *TriggerInputAccountsCollectedForAggregation) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *TriggerInputAccountsCollectedForAggregation) SetErrors(v []string)`

SetErrors sets Errors field to given value.


### SetErrorsNil

`func (o *TriggerInputAccountsCollectedForAggregation) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *TriggerInputAccountsCollectedForAggregation) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetWarnings

`func (o *TriggerInputAccountsCollectedForAggregation) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *TriggerInputAccountsCollectedForAggregation) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *TriggerInputAccountsCollectedForAggregation) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.


### SetWarningsNil

`func (o *TriggerInputAccountsCollectedForAggregation) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *TriggerInputAccountsCollectedForAggregation) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil
### GetStats

`func (o *TriggerInputAccountsCollectedForAggregation) GetStats() TriggerInputAccountsCollectedForAggregationStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *TriggerInputAccountsCollectedForAggregation) GetStatsOk() (*TriggerInputAccountsCollectedForAggregationStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *TriggerInputAccountsCollectedForAggregation) SetStats(v TriggerInputAccountsCollectedForAggregationStats)`

SetStats sets Stats field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


