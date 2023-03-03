# TriggerInputAccountAggregationCompleted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**TriggerInputAccountAggregationCompletedSource**](TriggerInputAccountAggregationCompletedSource.md) |  | 
**Status** | **map[string]interface{}** | The overall status of the aggregation. | 
**Started** | **time.Time** | The date and time when the account aggregation started. | 
**Completed** | **time.Time** | The date and time when the account aggregation finished. | 
**Errors** | **[]string** | A list of errors that occurred during the aggregation. | 
**Warnings** | **[]string** | A list of warnings that occurred during the aggregation. | 
**Stats** | [**TriggerInputAccountAggregationCompletedStats**](TriggerInputAccountAggregationCompletedStats.md) |  | 

## Methods

### NewTriggerInputAccountAggregationCompleted

`func NewTriggerInputAccountAggregationCompleted(source TriggerInputAccountAggregationCompletedSource, status map[string]interface{}, started time.Time, completed time.Time, errors []string, warnings []string, stats TriggerInputAccountAggregationCompletedStats, ) *TriggerInputAccountAggregationCompleted`

NewTriggerInputAccountAggregationCompleted instantiates a new TriggerInputAccountAggregationCompleted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerInputAccountAggregationCompletedWithDefaults

`func NewTriggerInputAccountAggregationCompletedWithDefaults() *TriggerInputAccountAggregationCompleted`

NewTriggerInputAccountAggregationCompletedWithDefaults instantiates a new TriggerInputAccountAggregationCompleted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *TriggerInputAccountAggregationCompleted) GetSource() TriggerInputAccountAggregationCompletedSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TriggerInputAccountAggregationCompleted) GetSourceOk() (*TriggerInputAccountAggregationCompletedSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TriggerInputAccountAggregationCompleted) SetSource(v TriggerInputAccountAggregationCompletedSource)`

SetSource sets Source field to given value.


### GetStatus

`func (o *TriggerInputAccountAggregationCompleted) GetStatus() map[string]interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TriggerInputAccountAggregationCompleted) GetStatusOk() (*map[string]interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TriggerInputAccountAggregationCompleted) SetStatus(v map[string]interface{})`

SetStatus sets Status field to given value.


### GetStarted

`func (o *TriggerInputAccountAggregationCompleted) GetStarted() time.Time`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *TriggerInputAccountAggregationCompleted) GetStartedOk() (*time.Time, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *TriggerInputAccountAggregationCompleted) SetStarted(v time.Time)`

SetStarted sets Started field to given value.


### GetCompleted

`func (o *TriggerInputAccountAggregationCompleted) GetCompleted() time.Time`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *TriggerInputAccountAggregationCompleted) GetCompletedOk() (*time.Time, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *TriggerInputAccountAggregationCompleted) SetCompleted(v time.Time)`

SetCompleted sets Completed field to given value.


### GetErrors

`func (o *TriggerInputAccountAggregationCompleted) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *TriggerInputAccountAggregationCompleted) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *TriggerInputAccountAggregationCompleted) SetErrors(v []string)`

SetErrors sets Errors field to given value.


### SetErrorsNil

`func (o *TriggerInputAccountAggregationCompleted) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *TriggerInputAccountAggregationCompleted) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetWarnings

`func (o *TriggerInputAccountAggregationCompleted) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *TriggerInputAccountAggregationCompleted) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *TriggerInputAccountAggregationCompleted) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.


### SetWarningsNil

`func (o *TriggerInputAccountAggregationCompleted) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *TriggerInputAccountAggregationCompleted) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil
### GetStats

`func (o *TriggerInputAccountAggregationCompleted) GetStats() TriggerInputAccountAggregationCompletedStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *TriggerInputAccountAggregationCompleted) GetStatsOk() (*TriggerInputAccountAggregationCompletedStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *TriggerInputAccountAggregationCompleted) SetStats(v TriggerInputAccountAggregationCompletedStats)`

SetStats sets Stats field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


