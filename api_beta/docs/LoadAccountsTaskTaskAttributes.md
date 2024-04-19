# LoadAccountsTaskTaskAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** | The id of the source | [optional] 
**OptimizedAggregation** | Pointer to **map[string]interface{}** | The indicator if the aggregation process was enabled/disabled for the aggregation job | [optional] 

## Methods

### NewLoadAccountsTaskTaskAttributes

`func NewLoadAccountsTaskTaskAttributes() *LoadAccountsTaskTaskAttributes`

NewLoadAccountsTaskTaskAttributes instantiates a new LoadAccountsTaskTaskAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadAccountsTaskTaskAttributesWithDefaults

`func NewLoadAccountsTaskTaskAttributesWithDefaults() *LoadAccountsTaskTaskAttributes`

NewLoadAccountsTaskTaskAttributesWithDefaults instantiates a new LoadAccountsTaskTaskAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *LoadAccountsTaskTaskAttributes) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *LoadAccountsTaskTaskAttributes) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *LoadAccountsTaskTaskAttributes) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *LoadAccountsTaskTaskAttributes) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetOptimizedAggregation

`func (o *LoadAccountsTaskTaskAttributes) GetOptimizedAggregation() map[string]interface{}`

GetOptimizedAggregation returns the OptimizedAggregation field if non-nil, zero value otherwise.

### GetOptimizedAggregationOk

`func (o *LoadAccountsTaskTaskAttributes) GetOptimizedAggregationOk() (*map[string]interface{}, bool)`

GetOptimizedAggregationOk returns a tuple with the OptimizedAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimizedAggregation

`func (o *LoadAccountsTaskTaskAttributes) SetOptimizedAggregation(v map[string]interface{})`

SetOptimizedAggregation sets OptimizedAggregation field to given value.

### HasOptimizedAggregation

`func (o *LoadAccountsTaskTaskAttributes) HasOptimizedAggregation() bool`

HasOptimizedAggregation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


