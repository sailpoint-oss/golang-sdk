# Schedule2Days

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **string** | The application id | [optional] 
**AccountMatchConfig** | Pointer to [**SelectorAccountMatchConfig**](SelectorAccountMatchConfig.md) |  | [optional] 

## Methods

### NewSchedule2Days

`func NewSchedule2Days() *Schedule2Days`

NewSchedule2Days instantiates a new Schedule2Days object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedule2DaysWithDefaults

`func NewSchedule2DaysWithDefaults() *Schedule2Days`

NewSchedule2DaysWithDefaults instantiates a new Schedule2Days object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *Schedule2Days) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *Schedule2Days) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *Schedule2Days) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *Schedule2Days) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetAccountMatchConfig

`func (o *Schedule2Days) GetAccountMatchConfig() SelectorAccountMatchConfig`

GetAccountMatchConfig returns the AccountMatchConfig field if non-nil, zero value otherwise.

### GetAccountMatchConfigOk

`func (o *Schedule2Days) GetAccountMatchConfigOk() (*SelectorAccountMatchConfig, bool)`

GetAccountMatchConfigOk returns a tuple with the AccountMatchConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMatchConfig

`func (o *Schedule2Days) SetAccountMatchConfig(v SelectorAccountMatchConfig)`

SetAccountMatchConfig sets AccountMatchConfig field to given value.

### HasAccountMatchConfig

`func (o *Schedule2Days) HasAccountMatchConfig() bool`

HasAccountMatchConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


