# RoleMiningSessionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | The role mining session status. Can be one of these states - CREATED|UPDATED|IDENTITIES_OBTAINED|PRUNE_THRESHOLD_OBTAINED|POTENTIAL_ROLES_PROCESSING|POTENTIAL_ROLES_CREATED | [optional] 

## Methods

### NewRoleMiningSessionStatus

`func NewRoleMiningSessionStatus() *RoleMiningSessionStatus`

NewRoleMiningSessionStatus instantiates a new RoleMiningSessionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleMiningSessionStatusWithDefaults

`func NewRoleMiningSessionStatusWithDefaults() *RoleMiningSessionStatus`

NewRoleMiningSessionStatusWithDefaults instantiates a new RoleMiningSessionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *RoleMiningSessionStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RoleMiningSessionStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RoleMiningSessionStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *RoleMiningSessionStatus) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


