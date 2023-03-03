# EventDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Type** | [**DocumentType**](DocumentType.md) |  | 
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

### NewEventDocument

`func NewEventDocument(id string, name string, type_ DocumentType, ) *EventDocument`

NewEventDocument instantiates a new EventDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventDocumentWithDefaults

`func NewEventDocumentWithDefaults() *EventDocument`

NewEventDocumentWithDefaults instantiates a new EventDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EventDocument) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventDocument) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventDocument) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *EventDocument) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventDocument) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventDocument) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *EventDocument) GetType() DocumentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventDocument) GetTypeOk() (*DocumentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventDocument) SetType(v DocumentType)`

SetType sets Type field to given value.


### GetCreated

`func (o *EventDocument) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *EventDocument) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *EventDocument) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *EventDocument) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *EventDocument) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *EventDocument) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetSynced

`func (o *EventDocument) GetSynced() time.Time`

GetSynced returns the Synced field if non-nil, zero value otherwise.

### GetSyncedOk

`func (o *EventDocument) GetSyncedOk() (*time.Time, bool)`

GetSyncedOk returns a tuple with the Synced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynced

`func (o *EventDocument) SetSynced(v time.Time)`

SetSynced sets Synced field to given value.

### HasSynced

`func (o *EventDocument) HasSynced() bool`

HasSynced returns a boolean if a field has been set.

### SetSyncedNil

`func (o *EventDocument) SetSyncedNil(b bool)`

 SetSyncedNil sets the value for Synced to be an explicit nil

### UnsetSynced
`func (o *EventDocument) UnsetSynced()`

UnsetSynced ensures that no value is present for Synced, not even an explicit nil
### GetAction

`func (o *EventDocument) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *EventDocument) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *EventDocument) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *EventDocument) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetType

`func (o *EventDocument) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventDocument) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventDocument) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EventDocument) HasType() bool`

HasType returns a boolean if a field has been set.

### GetActor

`func (o *EventDocument) GetActor() NameType`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *EventDocument) GetActorOk() (*NameType, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *EventDocument) SetActor(v NameType)`

SetActor sets Actor field to given value.

### HasActor

`func (o *EventDocument) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetTarget

`func (o *EventDocument) GetTarget() NameType`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *EventDocument) GetTargetOk() (*NameType, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *EventDocument) SetTarget(v NameType)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *EventDocument) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetStack

`func (o *EventDocument) GetStack() string`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *EventDocument) GetStackOk() (*string, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *EventDocument) SetStack(v string)`

SetStack sets Stack field to given value.

### HasStack

`func (o *EventDocument) HasStack() bool`

HasStack returns a boolean if a field has been set.

### GetTrackingNumber

`func (o *EventDocument) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *EventDocument) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *EventDocument) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *EventDocument) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### GetIpAddress

`func (o *EventDocument) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *EventDocument) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *EventDocument) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *EventDocument) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetDetails

`func (o *EventDocument) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *EventDocument) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *EventDocument) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *EventDocument) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetAttributes

`func (o *EventDocument) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EventDocument) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EventDocument) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *EventDocument) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetObjects

`func (o *EventDocument) GetObjects() []string`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *EventDocument) GetObjectsOk() (*[]string, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *EventDocument) SetObjects(v []string)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *EventDocument) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetOperation

`func (o *EventDocument) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *EventDocument) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *EventDocument) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *EventDocument) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetStatus

`func (o *EventDocument) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventDocument) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EventDocument) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EventDocument) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTechnicalName

`func (o *EventDocument) GetTechnicalName() string`

GetTechnicalName returns the TechnicalName field if non-nil, zero value otherwise.

### GetTechnicalNameOk

`func (o *EventDocument) GetTechnicalNameOk() (*string, bool)`

GetTechnicalNameOk returns a tuple with the TechnicalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalName

`func (o *EventDocument) SetTechnicalName(v string)`

SetTechnicalName sets TechnicalName field to given value.

### HasTechnicalName

`func (o *EventDocument) HasTechnicalName() bool`

HasTechnicalName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


