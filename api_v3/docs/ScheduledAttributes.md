# ScheduledAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CronString** | **string** | A valid CRON expression | 

## Methods

### NewScheduledAttributes

`func NewScheduledAttributes(cronString string, ) *ScheduledAttributes`

NewScheduledAttributes instantiates a new ScheduledAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledAttributesWithDefaults

`func NewScheduledAttributesWithDefaults() *ScheduledAttributes`

NewScheduledAttributesWithDefaults instantiates a new ScheduledAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCronString

`func (o *ScheduledAttributes) GetCronString() string`

GetCronString returns the CronString field if non-nil, zero value otherwise.

### GetCronStringOk

`func (o *ScheduledAttributes) GetCronStringOk() (*string, bool)`

GetCronStringOk returns a tuple with the CronString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronString

`func (o *ScheduledAttributes) SetCronString(v string)`

SetCronString sets CronString field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


