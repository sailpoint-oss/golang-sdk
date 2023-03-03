# EventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**Synced** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**Action** | Pointer to **string** | The action that was performed | [optional] 
**Type** | Pointer to **string** | The type of event | [optional] 
**Actor** | Pointer to [**NameType**](NameType.md) |  | [optional] 
**Target** | Pointer to [**NameType**](NameType.md) |  | [optional] 
**Stack** | Pointer to **string** |  | [optional] 
**TrackingNumber** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**Objects** | Pointer to **[]string** |  | [optional] 
**Operation** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TechnicalName** | Pointer to **string** |  | [optional] 

## Methods

### NewEventAllOf

`func NewEventAllOf() *EventAllOf`

NewEventAllOf instantiates a new EventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventAllOfWithDefaults

`func NewEventAllOfWithDefaults() *EventAllOf`

NewEventAllOfWithDefaults instantiates a new EventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *EventAllOf) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *EventAllOf) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *EventAllOf) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *EventAllOf) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *EventAllOf) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *EventAllOf) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetSynced

`func (o *EventAllOf) GetSynced() time.Time`

GetSynced returns the Synced field if non-nil, zero value otherwise.

### GetSyncedOk

`func (o *EventAllOf) GetSyncedOk() (*time.Time, bool)`

GetSyncedOk returns a tuple with the Synced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynced

`func (o *EventAllOf) SetSynced(v time.Time)`

SetSynced sets Synced field to given value.

### HasSynced

`func (o *EventAllOf) HasSynced() bool`

HasSynced returns a boolean if a field has been set.

### SetSyncedNil

`func (o *EventAllOf) SetSyncedNil(b bool)`

 SetSyncedNil sets the value for Synced to be an explicit nil

### UnsetSynced
`func (o *EventAllOf) UnsetSynced()`

UnsetSynced ensures that no value is present for Synced, not even an explicit nil
### GetAction

`func (o *EventAllOf) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *EventAllOf) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *EventAllOf) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *EventAllOf) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetType

`func (o *EventAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EventAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetActor

`func (o *EventAllOf) GetActor() NameType`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *EventAllOf) GetActorOk() (*NameType, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *EventAllOf) SetActor(v NameType)`

SetActor sets Actor field to given value.

### HasActor

`func (o *EventAllOf) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetTarget

`func (o *EventAllOf) GetTarget() NameType`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *EventAllOf) GetTargetOk() (*NameType, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *EventAllOf) SetTarget(v NameType)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *EventAllOf) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetStack

`func (o *EventAllOf) GetStack() string`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *EventAllOf) GetStackOk() (*string, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *EventAllOf) SetStack(v string)`

SetStack sets Stack field to given value.

### HasStack

`func (o *EventAllOf) HasStack() bool`

HasStack returns a boolean if a field has been set.

### GetTrackingNumber

`func (o *EventAllOf) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *EventAllOf) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *EventAllOf) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *EventAllOf) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### GetIpAddress

`func (o *EventAllOf) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *EventAllOf) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *EventAllOf) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *EventAllOf) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetDetails

`func (o *EventAllOf) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *EventAllOf) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *EventAllOf) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *EventAllOf) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetAttributes

`func (o *EventAllOf) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EventAllOf) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EventAllOf) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *EventAllOf) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetObjects

`func (o *EventAllOf) GetObjects() []string`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *EventAllOf) GetObjectsOk() (*[]string, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *EventAllOf) SetObjects(v []string)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *EventAllOf) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetOperation

`func (o *EventAllOf) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *EventAllOf) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *EventAllOf) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *EventAllOf) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetStatus

`func (o *EventAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EventAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EventAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTechnicalName

`func (o *EventAllOf) GetTechnicalName() string`

GetTechnicalName returns the TechnicalName field if non-nil, zero value otherwise.

### GetTechnicalNameOk

`func (o *EventAllOf) GetTechnicalNameOk() (*string, bool)`

GetTechnicalNameOk returns a tuple with the TechnicalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalName

`func (o *EventAllOf) SetTechnicalName(v string)`

SetTechnicalName sets TechnicalName field to given value.

### HasTechnicalName

`func (o *EventAllOf) HasTechnicalName() bool`

HasTechnicalName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


