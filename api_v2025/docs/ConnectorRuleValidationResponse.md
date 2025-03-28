# ConnectorRuleValidationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** |  | 
**Details** | [**[]ConnectorRuleValidationResponseDetailsInner**](ConnectorRuleValidationResponseDetailsInner.md) |  | 

## Methods

### NewConnectorRuleValidationResponse

`func NewConnectorRuleValidationResponse(state string, details []ConnectorRuleValidationResponseDetailsInner, ) *ConnectorRuleValidationResponse`

NewConnectorRuleValidationResponse instantiates a new ConnectorRuleValidationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorRuleValidationResponseWithDefaults

`func NewConnectorRuleValidationResponseWithDefaults() *ConnectorRuleValidationResponse`

NewConnectorRuleValidationResponseWithDefaults instantiates a new ConnectorRuleValidationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ConnectorRuleValidationResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ConnectorRuleValidationResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ConnectorRuleValidationResponse) SetState(v string)`

SetState sets State field to given value.


### GetDetails

`func (o *ConnectorRuleValidationResponse) GetDetails() []ConnectorRuleValidationResponseDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ConnectorRuleValidationResponse) GetDetailsOk() (*[]ConnectorRuleValidationResponseDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ConnectorRuleValidationResponse) SetDetails(v []ConnectorRuleValidationResponseDetailsInner)`

SetDetails sets Details field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


