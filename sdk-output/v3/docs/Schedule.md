# Schedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SavedSearchId** | **string** | The ID of the saved search that will be executed.  | 
**Created** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**Modified** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**Schedule** | [**Schedule1**](Schedule1.md) |  | 
**Recipients** | [**[]TypedReference**](TypedReference.md) | The email recipients.  | 
**Enabled** | Pointer to **bool** | Indicates if the scheduled search is enabled.  | [optional] [default to false]
**EmailEmptyResults** | Pointer to **bool** | Indicates if email generation should not be suppressed if search returns no results.  | [optional] [default to false]
**DisplayQueryDetails** | Pointer to **bool** | Indicates if the generated email should include the query and search results preview (which could include PII).  | [optional] [default to false]

## Methods

### NewSchedule

`func NewSchedule(savedSearchId string, schedule Schedule1, recipients []TypedReference, ) *Schedule`

NewSchedule instantiates a new Schedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleWithDefaults

`func NewScheduleWithDefaults() *Schedule`

NewScheduleWithDefaults instantiates a new Schedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSavedSearchId

`func (o *Schedule) GetSavedSearchId() string`

GetSavedSearchId returns the SavedSearchId field if non-nil, zero value otherwise.

### GetSavedSearchIdOk

`func (o *Schedule) GetSavedSearchIdOk() (*string, bool)`

GetSavedSearchIdOk returns a tuple with the SavedSearchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedSearchId

`func (o *Schedule) SetSavedSearchId(v string)`

SetSavedSearchId sets SavedSearchId field to given value.


### GetCreated

`func (o *Schedule) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Schedule) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Schedule) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Schedule) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *Schedule) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *Schedule) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetModified

`func (o *Schedule) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Schedule) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Schedule) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *Schedule) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModifiedNil

`func (o *Schedule) SetModifiedNil(b bool)`

 SetModifiedNil sets the value for Modified to be an explicit nil

### UnsetModified
`func (o *Schedule) UnsetModified()`

UnsetModified ensures that no value is present for Modified, not even an explicit nil
### GetSchedule

`func (o *Schedule) GetSchedule() Schedule1`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *Schedule) GetScheduleOk() (*Schedule1, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *Schedule) SetSchedule(v Schedule1)`

SetSchedule sets Schedule field to given value.


### GetRecipients

`func (o *Schedule) GetRecipients() []TypedReference`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *Schedule) GetRecipientsOk() (*[]TypedReference, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *Schedule) SetRecipients(v []TypedReference)`

SetRecipients sets Recipients field to given value.


### GetEnabled

`func (o *Schedule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Schedule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Schedule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Schedule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEmailEmptyResults

`func (o *Schedule) GetEmailEmptyResults() bool`

GetEmailEmptyResults returns the EmailEmptyResults field if non-nil, zero value otherwise.

### GetEmailEmptyResultsOk

`func (o *Schedule) GetEmailEmptyResultsOk() (*bool, bool)`

GetEmailEmptyResultsOk returns a tuple with the EmailEmptyResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailEmptyResults

`func (o *Schedule) SetEmailEmptyResults(v bool)`

SetEmailEmptyResults sets EmailEmptyResults field to given value.

### HasEmailEmptyResults

`func (o *Schedule) HasEmailEmptyResults() bool`

HasEmailEmptyResults returns a boolean if a field has been set.

### GetDisplayQueryDetails

`func (o *Schedule) GetDisplayQueryDetails() bool`

GetDisplayQueryDetails returns the DisplayQueryDetails field if non-nil, zero value otherwise.

### GetDisplayQueryDetailsOk

`func (o *Schedule) GetDisplayQueryDetailsOk() (*bool, bool)`

GetDisplayQueryDetailsOk returns a tuple with the DisplayQueryDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayQueryDetails

`func (o *Schedule) SetDisplayQueryDetails(v bool)`

SetDisplayQueryDetails sets DisplayQueryDetails field to given value.

### HasDisplayQueryDetails

`func (o *Schedule) HasDisplayQueryDetails() bool`

HasDisplayQueryDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


