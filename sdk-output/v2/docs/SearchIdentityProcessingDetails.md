# SearchIdentityProcessingDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **time.Time** |  | [optional] 
**Stage** | Pointer to **string** |  | [optional] 
**RetryCount** | Pointer to **int32** |  | [optional] 
**StackTrace** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewSearchIdentityProcessingDetails

`func NewSearchIdentityProcessingDetails() *SearchIdentityProcessingDetails`

NewSearchIdentityProcessingDetails instantiates a new SearchIdentityProcessingDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchIdentityProcessingDetailsWithDefaults

`func NewSearchIdentityProcessingDetailsWithDefaults() *SearchIdentityProcessingDetails`

NewSearchIdentityProcessingDetailsWithDefaults instantiates a new SearchIdentityProcessingDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *SearchIdentityProcessingDetails) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *SearchIdentityProcessingDetails) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *SearchIdentityProcessingDetails) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *SearchIdentityProcessingDetails) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetStage

`func (o *SearchIdentityProcessingDetails) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *SearchIdentityProcessingDetails) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *SearchIdentityProcessingDetails) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *SearchIdentityProcessingDetails) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetRetryCount

`func (o *SearchIdentityProcessingDetails) GetRetryCount() int32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *SearchIdentityProcessingDetails) GetRetryCountOk() (*int32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *SearchIdentityProcessingDetails) SetRetryCount(v int32)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *SearchIdentityProcessingDetails) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetStackTrace

`func (o *SearchIdentityProcessingDetails) GetStackTrace() string`

GetStackTrace returns the StackTrace field if non-nil, zero value otherwise.

### GetStackTraceOk

`func (o *SearchIdentityProcessingDetails) GetStackTraceOk() (*string, bool)`

GetStackTraceOk returns a tuple with the StackTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackTrace

`func (o *SearchIdentityProcessingDetails) SetStackTrace(v string)`

SetStackTrace sets StackTrace field to given value.

### HasStackTrace

`func (o *SearchIdentityProcessingDetails) HasStackTrace() bool`

HasStackTrace returns a boolean if a field has been set.

### GetMessage

`func (o *SearchIdentityProcessingDetails) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SearchIdentityProcessingDetails) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SearchIdentityProcessingDetails) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SearchIdentityProcessingDetails) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


