# ConditionEffect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to **map[string]map[string]interface{}** | Config is a arbitrary map that holds a configuration based on EffectType | [optional] 
**EffectType** | Pointer to **string** | EffectType is the type of effect to perform when the conditions are evaluated for this logic block HIDE ConditionEffectTypeHide  ConditionEffectTypeHide disables validations SHOW ConditionEffectTypeShow  ConditionEffectTypeShow enables validations DISABLE ConditionEffectTypeDisable  ConditionEffectTypeDisable disables validations ENABLE ConditionEffectTypeEnable  ConditionEffectTypeEnable enables validations REQUIRE ConditionEffectTypeRequire OPTIONAL ConditionEffectTypeOptional SUBMIT_MESSAGE ConditionEffectTypeSubmitMessage SUBMIT_NOTIFICATION ConditionEffectTypeSubmitNotification SET_DEFAULT_VALUE ConditionEffectTypeSetDefaultValue  ConditionEffectTypeSetDefaultValue is ignored on purpose | [optional] 

## Methods

### NewConditionEffect

`func NewConditionEffect() *ConditionEffect`

NewConditionEffect instantiates a new ConditionEffect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionEffectWithDefaults

`func NewConditionEffectWithDefaults() *ConditionEffect`

NewConditionEffectWithDefaults instantiates a new ConditionEffect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *ConditionEffect) GetConfig() map[string]map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ConditionEffect) GetConfigOk() (*map[string]map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ConditionEffect) SetConfig(v map[string]map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ConditionEffect) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEffectType

`func (o *ConditionEffect) GetEffectType() string`

GetEffectType returns the EffectType field if non-nil, zero value otherwise.

### GetEffectTypeOk

`func (o *ConditionEffect) GetEffectTypeOk() (*string, bool)`

GetEffectTypeOk returns a tuple with the EffectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectType

`func (o *ConditionEffect) SetEffectType(v string)`

SetEffectType sets EffectType field to given value.

### HasEffectType

`func (o *ConditionEffect) HasEffectType() bool`

HasEffectType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


