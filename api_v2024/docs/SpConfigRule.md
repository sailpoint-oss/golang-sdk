# SpConfigRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | JSONPath expression denoting the path within the object where a value substitution should be applied | [optional] 
**Value** | Pointer to **map[string]interface{}** | Value to be assigned at the jsonPath location within the object | [optional] 
**Mode** | Pointer to **[]string** | Draft modes to which this rule will apply | [optional] 

## Methods

### NewSpConfigRule

`func NewSpConfigRule() *SpConfigRule`

NewSpConfigRule instantiates a new SpConfigRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpConfigRuleWithDefaults

`func NewSpConfigRuleWithDefaults() *SpConfigRule`

NewSpConfigRuleWithDefaults instantiates a new SpConfigRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *SpConfigRule) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SpConfigRule) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SpConfigRule) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SpConfigRule) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetValue

`func (o *SpConfigRule) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SpConfigRule) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SpConfigRule) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *SpConfigRule) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *SpConfigRule) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *SpConfigRule) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetMode

`func (o *SpConfigRule) GetMode() []string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SpConfigRule) GetModeOk() (*[]string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SpConfigRule) SetMode(v []string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *SpConfigRule) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


