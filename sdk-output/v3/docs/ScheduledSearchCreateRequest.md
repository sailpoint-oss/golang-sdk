# ScheduledSearchCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the scheduled search.  | [optional] 
**Description** | Pointer to **string** | The description of the scheduled search.  | [optional] 
**SavedSearchId** | **string** | The ID of the saved search that will be executed.  | 
**Created** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**Modified** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**Schedule** | [**Schedule1**](Schedule1.md) |  | 
**Recipients** | [**[]TypedReference**](TypedReference.md) | The email recipients.  | 
**Enabled** | Pointer to **bool** | Indicates if the scheduled search is enabled.  | [optional] [default to false]
**EmailEmptyResults** | Pointer to **bool** | Indicates if email generation should not be suppressed if search returns no results.  | [optional] [default to false]
**DisplayQueryDetails** | Pointer to **bool** | Indicates if the generated email should include the query and search results preview (which could include PII).  | [optional] [default to false]

## Methods

### NewScheduledSearchCreateRequest

`func NewScheduledSearchCreateRequest(savedSearchId string, schedule Schedule1, recipients []TypedReference, ) *ScheduledSearchCreateRequest`

NewScheduledSearchCreateRequest instantiates a new ScheduledSearchCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledSearchCreateRequestWithDefaults

`func NewScheduledSearchCreateRequestWithDefaults() *ScheduledSearchCreateRequest`

NewScheduledSearchCreateRequestWithDefaults instantiates a new ScheduledSearchCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ScheduledSearchCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScheduledSearchCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScheduledSearchCreateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScheduledSearchCreateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ScheduledSearchCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScheduledSearchCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScheduledSearchCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScheduledSearchCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSavedSearchId

`func (o *ScheduledSearchCreateRequest) GetSavedSearchId() string`

GetSavedSearchId returns the SavedSearchId field if non-nil, zero value otherwise.

### GetSavedSearchIdOk

`func (o *ScheduledSearchCreateRequest) GetSavedSearchIdOk() (*string, bool)`

GetSavedSearchIdOk returns a tuple with the SavedSearchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedSearchId

`func (o *ScheduledSearchCreateRequest) SetSavedSearchId(v string)`

SetSavedSearchId sets SavedSearchId field to given value.


### GetCreated

`func (o *ScheduledSearchCreateRequest) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ScheduledSearchCreateRequest) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ScheduledSearchCreateRequest) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ScheduledSearchCreateRequest) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *ScheduledSearchCreateRequest) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *ScheduledSearchCreateRequest) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetModified

`func (o *ScheduledSearchCreateRequest) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ScheduledSearchCreateRequest) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ScheduledSearchCreateRequest) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *ScheduledSearchCreateRequest) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModifiedNil

`func (o *ScheduledSearchCreateRequest) SetModifiedNil(b bool)`

 SetModifiedNil sets the value for Modified to be an explicit nil

### UnsetModified
`func (o *ScheduledSearchCreateRequest) UnsetModified()`

UnsetModified ensures that no value is present for Modified, not even an explicit nil
### GetSchedule

`func (o *ScheduledSearchCreateRequest) GetSchedule() Schedule1`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ScheduledSearchCreateRequest) GetScheduleOk() (*Schedule1, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ScheduledSearchCreateRequest) SetSchedule(v Schedule1)`

SetSchedule sets Schedule field to given value.


### GetRecipients

`func (o *ScheduledSearchCreateRequest) GetRecipients() []TypedReference`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *ScheduledSearchCreateRequest) GetRecipientsOk() (*[]TypedReference, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *ScheduledSearchCreateRequest) SetRecipients(v []TypedReference)`

SetRecipients sets Recipients field to given value.


### GetEnabled

`func (o *ScheduledSearchCreateRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ScheduledSearchCreateRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ScheduledSearchCreateRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ScheduledSearchCreateRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEmailEmptyResults

`func (o *ScheduledSearchCreateRequest) GetEmailEmptyResults() bool`

GetEmailEmptyResults returns the EmailEmptyResults field if non-nil, zero value otherwise.

### GetEmailEmptyResultsOk

`func (o *ScheduledSearchCreateRequest) GetEmailEmptyResultsOk() (*bool, bool)`

GetEmailEmptyResultsOk returns a tuple with the EmailEmptyResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailEmptyResults

`func (o *ScheduledSearchCreateRequest) SetEmailEmptyResults(v bool)`

SetEmailEmptyResults sets EmailEmptyResults field to given value.

### HasEmailEmptyResults

`func (o *ScheduledSearchCreateRequest) HasEmailEmptyResults() bool`

HasEmailEmptyResults returns a boolean if a field has been set.

### GetDisplayQueryDetails

`func (o *ScheduledSearchCreateRequest) GetDisplayQueryDetails() bool`

GetDisplayQueryDetails returns the DisplayQueryDetails field if non-nil, zero value otherwise.

### GetDisplayQueryDetailsOk

`func (o *ScheduledSearchCreateRequest) GetDisplayQueryDetailsOk() (*bool, bool)`

GetDisplayQueryDetailsOk returns a tuple with the DisplayQueryDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayQueryDetails

`func (o *ScheduledSearchCreateRequest) SetDisplayQueryDetails(v bool)`

SetDisplayQueryDetails sets DisplayQueryDetails field to given value.

### HasDisplayQueryDetails

`func (o *ScheduledSearchCreateRequest) HasDisplayQueryDetails() bool`

HasDisplayQueryDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


