# ConnectorRuleResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | the ID of the rule | 
**Created** | **string** | an ISO 8601 UTC timestamp when this rule was created | 
**Modified** | Pointer to **NullableString** | an ISO 8601 UTC timestamp when this rule was last modified | [optional] 

## Methods

### NewConnectorRuleResponseAllOf

`func NewConnectorRuleResponseAllOf(id string, created string, ) *ConnectorRuleResponseAllOf`

NewConnectorRuleResponseAllOf instantiates a new ConnectorRuleResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorRuleResponseAllOfWithDefaults

`func NewConnectorRuleResponseAllOfWithDefaults() *ConnectorRuleResponseAllOf`

NewConnectorRuleResponseAllOfWithDefaults instantiates a new ConnectorRuleResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConnectorRuleResponseAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectorRuleResponseAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectorRuleResponseAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetCreated

`func (o *ConnectorRuleResponseAllOf) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ConnectorRuleResponseAllOf) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ConnectorRuleResponseAllOf) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetModified

`func (o *ConnectorRuleResponseAllOf) GetModified() string`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ConnectorRuleResponseAllOf) GetModifiedOk() (*string, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ConnectorRuleResponseAllOf) SetModified(v string)`

SetModified sets Modified field to given value.

### HasModified

`func (o *ConnectorRuleResponseAllOf) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModifiedNil

`func (o *ConnectorRuleResponseAllOf) SetModifiedNil(b bool)`

 SetModifiedNil sets the value for Modified to be an explicit nil

### UnsetModified
`func (o *ConnectorRuleResponseAllOf) UnsetModified()`

UnsetModified ensures that no value is present for Modified, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


