# SodViolationContextCheckCompleted1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **NullableString** | The status of SOD violation check | [optional] 
**Uuid** | Pointer to **NullableString** | The id of the Violation check event | [optional] 
**ViolationCheckResult** | Pointer to [**SodViolationCheckResult1**](SodViolationCheckResult1.md) |  | [optional] 

## Methods

### NewSodViolationContextCheckCompleted1

`func NewSodViolationContextCheckCompleted1() *SodViolationContextCheckCompleted1`

NewSodViolationContextCheckCompleted1 instantiates a new SodViolationContextCheckCompleted1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSodViolationContextCheckCompleted1WithDefaults

`func NewSodViolationContextCheckCompleted1WithDefaults() *SodViolationContextCheckCompleted1`

NewSodViolationContextCheckCompleted1WithDefaults instantiates a new SodViolationContextCheckCompleted1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *SodViolationContextCheckCompleted1) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SodViolationContextCheckCompleted1) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SodViolationContextCheckCompleted1) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *SodViolationContextCheckCompleted1) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *SodViolationContextCheckCompleted1) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *SodViolationContextCheckCompleted1) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetUuid

`func (o *SodViolationContextCheckCompleted1) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SodViolationContextCheckCompleted1) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SodViolationContextCheckCompleted1) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *SodViolationContextCheckCompleted1) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *SodViolationContextCheckCompleted1) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *SodViolationContextCheckCompleted1) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetViolationCheckResult

`func (o *SodViolationContextCheckCompleted1) GetViolationCheckResult() SodViolationCheckResult1`

GetViolationCheckResult returns the ViolationCheckResult field if non-nil, zero value otherwise.

### GetViolationCheckResultOk

`func (o *SodViolationContextCheckCompleted1) GetViolationCheckResultOk() (*SodViolationCheckResult1, bool)`

GetViolationCheckResultOk returns a tuple with the ViolationCheckResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolationCheckResult

`func (o *SodViolationContextCheckCompleted1) SetViolationCheckResult(v SodViolationCheckResult1)`

SetViolationCheckResult sets ViolationCheckResult field to given value.

### HasViolationCheckResult

`func (o *SodViolationContextCheckCompleted1) HasViolationCheckResult() bool`

HasViolationCheckResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


