# FirstValid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Values** | **[]map[string]interface{}** | An array of attributes to evaluate for existence. | 
**IgnoreErrors** | Pointer to **bool** | a true or false value representing to move on to the next option if an error (like an Null Pointer Exception) were to occur. | [optional] 
**RequiresPeriodicRefresh** | Pointer to **bool** | A value that indicates whether the transform logic should be re-evaluated every evening as part of the identity refresh process | [optional] [default to false]

## Methods

### NewFirstValid

`func NewFirstValid(values []map[string]interface{}, ) *FirstValid`

NewFirstValid instantiates a new FirstValid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirstValidWithDefaults

`func NewFirstValidWithDefaults() *FirstValid`

NewFirstValidWithDefaults instantiates a new FirstValid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValues

`func (o *FirstValid) GetValues() []map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *FirstValid) GetValuesOk() (*[]map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *FirstValid) SetValues(v []map[string]interface{})`

SetValues sets Values field to given value.


### GetIgnoreErrors

`func (o *FirstValid) GetIgnoreErrors() bool`

GetIgnoreErrors returns the IgnoreErrors field if non-nil, zero value otherwise.

### GetIgnoreErrorsOk

`func (o *FirstValid) GetIgnoreErrorsOk() (*bool, bool)`

GetIgnoreErrorsOk returns a tuple with the IgnoreErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreErrors

`func (o *FirstValid) SetIgnoreErrors(v bool)`

SetIgnoreErrors sets IgnoreErrors field to given value.

### HasIgnoreErrors

`func (o *FirstValid) HasIgnoreErrors() bool`

HasIgnoreErrors returns a boolean if a field has been set.

### GetRequiresPeriodicRefresh

`func (o *FirstValid) GetRequiresPeriodicRefresh() bool`

GetRequiresPeriodicRefresh returns the RequiresPeriodicRefresh field if non-nil, zero value otherwise.

### GetRequiresPeriodicRefreshOk

`func (o *FirstValid) GetRequiresPeriodicRefreshOk() (*bool, bool)`

GetRequiresPeriodicRefreshOk returns a tuple with the RequiresPeriodicRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPeriodicRefresh

`func (o *FirstValid) SetRequiresPeriodicRefresh(v bool)`

SetRequiresPeriodicRefresh sets RequiresPeriodicRefresh field to given value.

### HasRequiresPeriodicRefresh

`func (o *FirstValid) HasRequiresPeriodicRefresh() bool`

HasRequiresPeriodicRefresh returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


