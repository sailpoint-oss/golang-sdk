# RunSearch200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identity** | Pointer to [**[]SearchIdentity**](SearchIdentity.md) |  | [optional] 
**Entitlement** | Pointer to [**[]SearchEntitlement**](SearchEntitlement.md) |  | [optional] 
**Event** | Pointer to [**[]SearchEvent**](SearchEvent.md) |  | [optional] 

## Methods

### NewRunSearch200Response

`func NewRunSearch200Response() *RunSearch200Response`

NewRunSearch200Response instantiates a new RunSearch200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunSearch200ResponseWithDefaults

`func NewRunSearch200ResponseWithDefaults() *RunSearch200Response`

NewRunSearch200ResponseWithDefaults instantiates a new RunSearch200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentity

`func (o *RunSearch200Response) GetIdentity() []SearchIdentity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *RunSearch200Response) GetIdentityOk() (*[]SearchIdentity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *RunSearch200Response) SetIdentity(v []SearchIdentity)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *RunSearch200Response) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetEntitlement

`func (o *RunSearch200Response) GetEntitlement() []SearchEntitlement`

GetEntitlement returns the Entitlement field if non-nil, zero value otherwise.

### GetEntitlementOk

`func (o *RunSearch200Response) GetEntitlementOk() (*[]SearchEntitlement, bool)`

GetEntitlementOk returns a tuple with the Entitlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlement

`func (o *RunSearch200Response) SetEntitlement(v []SearchEntitlement)`

SetEntitlement sets Entitlement field to given value.

### HasEntitlement

`func (o *RunSearch200Response) HasEntitlement() bool`

HasEntitlement returns a boolean if a field has been set.

### GetEvent

`func (o *RunSearch200Response) GetEvent() []SearchEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *RunSearch200Response) GetEventOk() (*[]SearchEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *RunSearch200Response) SetEvent(v []SearchEvent)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *RunSearch200Response) HasEvent() bool`

HasEvent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


