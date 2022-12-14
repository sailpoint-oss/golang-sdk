# SearchEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Synced** | Pointer to **time.Time** |  | [optional] 
**Actor** | Pointer to [**SearchEventActor**](SearchEventActor.md) |  | [optional] 
**Target** | Pointer to [**SearchEventActor**](SearchEventActor.md) |  | [optional] 
**Stack** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to [**SearchEventAttributes**](SearchEventAttributes.md) |  | [optional] 

## Methods

### NewSearchEvent

`func NewSearchEvent() *SearchEvent`

NewSearchEvent instantiates a new SearchEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchEventWithDefaults

`func NewSearchEventWithDefaults() *SearchEvent`

NewSearchEventWithDefaults instantiates a new SearchEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SearchEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SearchEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAction

`func (o *SearchEvent) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SearchEvent) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SearchEvent) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *SearchEvent) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetType

`func (o *SearchEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SearchEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SearchEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SearchEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreated

`func (o *SearchEvent) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SearchEvent) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SearchEvent) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SearchEvent) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetSynced

`func (o *SearchEvent) GetSynced() time.Time`

GetSynced returns the Synced field if non-nil, zero value otherwise.

### GetSyncedOk

`func (o *SearchEvent) GetSyncedOk() (*time.Time, bool)`

GetSyncedOk returns a tuple with the Synced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynced

`func (o *SearchEvent) SetSynced(v time.Time)`

SetSynced sets Synced field to given value.

### HasSynced

`func (o *SearchEvent) HasSynced() bool`

HasSynced returns a boolean if a field has been set.

### GetActor

`func (o *SearchEvent) GetActor() SearchEventActor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *SearchEvent) GetActorOk() (*SearchEventActor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *SearchEvent) SetActor(v SearchEventActor)`

SetActor sets Actor field to given value.

### HasActor

`func (o *SearchEvent) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetTarget

`func (o *SearchEvent) GetTarget() SearchEventActor`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *SearchEvent) GetTargetOk() (*SearchEventActor, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *SearchEvent) SetTarget(v SearchEventActor)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *SearchEvent) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetStack

`func (o *SearchEvent) GetStack() string`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *SearchEvent) GetStackOk() (*string, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *SearchEvent) SetStack(v string)`

SetStack sets Stack field to given value.

### HasStack

`func (o *SearchEvent) HasStack() bool`

HasStack returns a boolean if a field has been set.

### GetRequestId

`func (o *SearchEvent) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *SearchEvent) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *SearchEvent) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *SearchEvent) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetHostname

`func (o *SearchEvent) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *SearchEvent) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *SearchEvent) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *SearchEvent) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIpAddress

`func (o *SearchEvent) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *SearchEvent) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *SearchEvent) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *SearchEvent) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetAttributes

`func (o *SearchEvent) GetAttributes() SearchEventAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SearchEvent) GetAttributesOk() (*SearchEventAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SearchEvent) SetAttributes(v SearchEventAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SearchEvent) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


