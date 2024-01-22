# SendTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserAlias** | **string** | User alias from table spt_identity field named &#39;name&#39; | 
**DeliveryType** | **string** | Token delivery type | 

## Methods

### NewSendTokenRequest

`func NewSendTokenRequest(userAlias string, deliveryType string, ) *SendTokenRequest`

NewSendTokenRequest instantiates a new SendTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendTokenRequestWithDefaults

`func NewSendTokenRequestWithDefaults() *SendTokenRequest`

NewSendTokenRequestWithDefaults instantiates a new SendTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserAlias

`func (o *SendTokenRequest) GetUserAlias() string`

GetUserAlias returns the UserAlias field if non-nil, zero value otherwise.

### GetUserAliasOk

`func (o *SendTokenRequest) GetUserAliasOk() (*string, bool)`

GetUserAliasOk returns a tuple with the UserAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAlias

`func (o *SendTokenRequest) SetUserAlias(v string)`

SetUserAlias sets UserAlias field to given value.


### GetDeliveryType

`func (o *SendTokenRequest) GetDeliveryType() string`

GetDeliveryType returns the DeliveryType field if non-nil, zero value otherwise.

### GetDeliveryTypeOk

`func (o *SendTokenRequest) GetDeliveryTypeOk() (*string, bool)`

GetDeliveryTypeOk returns a tuple with the DeliveryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryType

`func (o *SendTokenRequest) SetDeliveryType(v string)`

SetDeliveryType sets DeliveryType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


