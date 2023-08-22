# RoleMiningSessionScopingMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | The scoping method used in the current role mining session. Can be one of these states - MANUAL|AUTO_RM | [optional] 

## Methods

### NewRoleMiningSessionScopingMethod

`func NewRoleMiningSessionScopingMethod() *RoleMiningSessionScopingMethod`

NewRoleMiningSessionScopingMethod instantiates a new RoleMiningSessionScopingMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleMiningSessionScopingMethodWithDefaults

`func NewRoleMiningSessionScopingMethodWithDefaults() *RoleMiningSessionScopingMethod`

NewRoleMiningSessionScopingMethodWithDefaults instantiates a new RoleMiningSessionScopingMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *RoleMiningSessionScopingMethod) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RoleMiningSessionScopingMethod) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RoleMiningSessionScopingMethod) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *RoleMiningSessionScopingMethod) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


