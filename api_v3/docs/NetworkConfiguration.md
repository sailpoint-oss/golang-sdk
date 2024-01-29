# NetworkConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Range** | Pointer to **[]string** | The collection of ip ranges. | [optional] 
**Geolocation** | Pointer to **[]string** | The collection of country codes. | [optional] 
**Whitelisted** | Pointer to **bool** | Denotes whether the provided lists are whitelisted or blacklisted for geo location. | [optional] [default to false]

## Methods

### NewNetworkConfiguration

`func NewNetworkConfiguration() *NetworkConfiguration`

NewNetworkConfiguration instantiates a new NetworkConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkConfigurationWithDefaults

`func NewNetworkConfigurationWithDefaults() *NetworkConfiguration`

NewNetworkConfigurationWithDefaults instantiates a new NetworkConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRange

`func (o *NetworkConfiguration) GetRange() []string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *NetworkConfiguration) GetRangeOk() (*[]string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *NetworkConfiguration) SetRange(v []string)`

SetRange sets Range field to given value.

### HasRange

`func (o *NetworkConfiguration) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetGeolocation

`func (o *NetworkConfiguration) GetGeolocation() []string`

GetGeolocation returns the Geolocation field if non-nil, zero value otherwise.

### GetGeolocationOk

`func (o *NetworkConfiguration) GetGeolocationOk() (*[]string, bool)`

GetGeolocationOk returns a tuple with the Geolocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeolocation

`func (o *NetworkConfiguration) SetGeolocation(v []string)`

SetGeolocation sets Geolocation field to given value.

### HasGeolocation

`func (o *NetworkConfiguration) HasGeolocation() bool`

HasGeolocation returns a boolean if a field has been set.

### GetWhitelisted

`func (o *NetworkConfiguration) GetWhitelisted() bool`

GetWhitelisted returns the Whitelisted field if non-nil, zero value otherwise.

### GetWhitelistedOk

`func (o *NetworkConfiguration) GetWhitelistedOk() (*bool, bool)`

GetWhitelistedOk returns a tuple with the Whitelisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelisted

`func (o *NetworkConfiguration) SetWhitelisted(v bool)`

SetWhitelisted sets Whitelisted field to given value.

### HasWhitelisted

`func (o *NetworkConfiguration) HasWhitelisted() bool`

HasWhitelisted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


