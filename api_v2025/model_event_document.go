/*
Identity Security Cloud V2025 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2025
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2025

import (
	"encoding/json"
	
)

// checks if the EventDocument type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventDocument{}

// EventDocument Event
type EventDocument struct {
	// ID of the entitlement.
	Id *string `json:"id,omitempty"`
	// Name of the entitlement.
	Name *string `json:"name,omitempty"`
	// ISO-8601 date-time referring to the time when the object was created.
	Created NullableTime `json:"created,omitempty"`
	// ISO-8601 date-time referring to the date-time when object was queued to be synced into search database for use in the search API.   This date-time changes anytime there is an update to the object, which triggers a synchronization event being sent to the search database.  There may be some delay between the `synced` time and the time when the updated data is actually available in the search API. 
	Synced *string `json:"synced,omitempty"`
	// Name of the event as it's displayed in audit reports.
	Action *string `json:"action,omitempty"`
	// Event type. Refer to [Event Types](https://documentation.sailpoint.com/saas/help/search/index.html#event-types) for a list of event types and their meanings.
	Type *string `json:"type,omitempty"`
	Actor *EventActor `json:"actor,omitempty"`
	Target *EventTarget `json:"target,omitempty"`
	// The event's stack.
	Stack *string `json:"stack,omitempty"`
	// ID of the group of events.
	TrackingNumber *string `json:"trackingNumber,omitempty"`
	// Target system's IP address.
	IpAddress *string `json:"ipAddress,omitempty"`
	// ID of event's details.
	Details *string `json:"details,omitempty"`
	// Attributes involved in the event.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// Objects the event is happening to.
	Objects []string `json:"objects,omitempty"`
	// Operation, or action, performed during the event.
	Operation *string `json:"operation,omitempty"`
	// Event status. Refer to [Event Statuses](https://documentation.sailpoint.com/saas/help/search/index.html#event-statuses) for a list of event statuses and their meanings.
	Status *string `json:"status,omitempty"`
	// Event's normalized name. This normalized name always follows the pattern of 'objects_operation_status'.
	TechnicalName *string `json:"technicalName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EventDocument EventDocument

// NewEventDocument instantiates a new EventDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventDocument() *EventDocument {
	this := EventDocument{}
	return &this
}

// NewEventDocumentWithDefaults instantiates a new EventDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventDocumentWithDefaults() *EventDocument {
	this := EventDocument{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EventDocument) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventDocument) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EventDocument) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EventDocument) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EventDocument) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventDocument) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EventDocument) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EventDocument) SetName(v string) {
	o.Name = &v
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventDocument) GetCreated() SailPointTime {
	if o == nil || IsNil(o.Created.Get()) {
		var ret SailPointTime
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventDocument) GetCreatedOk() (*SailPointTime, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *EventDocument) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableTime and assigns it to the Created field.
func (o *EventDocument) SetCreated(v SailPointTime) {
	o.Created.Set(&v)
}
// SetCreatedNil sets the value for Created to be an explicit nil
func (o *EventDocument) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *EventDocument) UnsetCreated() {
	o.Created.Unset()
}

// GetSynced returns the Synced field value if set, zero value otherwise.
func (o *EventDocument) GetSynced() string {
	if o == nil || IsNil(o.Synced) {
		var ret string
		return ret
	}
	return *o.Synced
}

// GetSyncedOk returns a tuple with the Synced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventDocument) GetSyncedOk() (*string, bool) {
	if o == nil || IsNil(o.Synced) {
		return nil, false
	}
	return o.Synced, true
}

// HasSynced returns a boolean if a field has been set.
func (o *EventDocument) HasSynced() bool {
	if o != nil && !IsNil(o.Synced) {
		return true
	}

	return false
}

// SetSynced gets a reference to the given string and assigns it to the Synced field.
func (o *EventDocument) SetSynced(v string) {
	o.Synced = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *EventDocument) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventDocument) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *EventDocument) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *EventDocument) SetAction(v string) {
	o.Action = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EventDocument) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventDocument) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EventDocument) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EventDocument) SetType(v string) {
	o.Type = &v
}

// GetActor returns the Actor field value if set, zero value otherwise.
func (o *EventDocument) GetActor() EventActor {
	if o == nil || IsNil(o.Actor) {
		var ret EventActor
		return ret
	}
	return *o.Actor
}

// GetActorOk returns a tuple with the Actor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventDocument) GetActorOk() (*EventActor, bool) {
	if o == nil || IsNil(o.Actor) {
		return nil, false
	}
	return o.Actor, true
}

// HasActor returns a boolean if a field has been set.
func (o *EventDocument) HasActor() bool {
	if o != nil && !IsNil(o.Actor) {
		return true
	}

	return false
}

// SetActor gets a reference to the given EventActor and assigns it to the Actor field.
func (o *EventDocument) SetActor(v EventActor) {
	o.Actor = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *EventDocument) GetTarget() EventTarget {
	if o == nil || IsNil(o.Target) {
		var ret EventTarget
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventDocument) GetTargetOk() (*EventTarget, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *EventDocument) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given EventTarget and assigns it to the Target field.
func (o *EventDocument) SetTarget(v EventTarget) {
	o.Target = &v
}

// GetStack returns the Stack field value if set, zero value otherwise.
func (o *EventDocument) GetStack() string {
	if o == nil || IsNil(o.Stack) {
		var ret string
		return ret
	}
	return *o.Stack
}

// GetStackOk returns a tuple with the Stack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventDocument) GetStackOk() (*string, bool) {
	if o == nil || IsNil(o.Stack) {
		return nil, false
	}
	return o.Stack, true
}

// HasStack returns a boolean if a field has been set.
func (o *EventDocument) HasStack() bool {
	if o != nil && !IsNil(o.Stack) {
		return true
	}

	return false
}

// SetStack gets a reference to the given string and assigns it to the Stack field.
func (o *EventDocument) SetStack(v string) {
	o.Stack = &v
}

// GetTrackingNumber returns the TrackingNumber field value if set, zero value otherwise.
func (o *EventDocument) GetTrackingNumber() string {
	if o == nil || IsNil(o.TrackingNumber) {
		var ret string
		return ret
	}
	return *o.TrackingNumber
}

// GetTrackingNumberOk returns a tuple with the TrackingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventDocument) GetTrackingNumberOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingNumber) {
		return nil, false
	}
	return o.TrackingNumber, true
}

// HasTrackingNumber returns a boolean if a field has been set.
func (o *EventDocument) HasTrackingNumber() bool {
	if o != nil && !IsNil(o.TrackingNumber) {
		return true
	}

	return false
}

// SetTrackingNumber gets a reference to the given string and assigns it to the TrackingNumber field.
func (o *EventDocument) SetTrackingNumber(v string) {
	o.TrackingNumber = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *EventDocument) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventDocument) GetIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *EventDocument) HasIpAddress() bool {
	if o != nil && !IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *EventDocument) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *EventDocument) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventDocument) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *EventDocument) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *EventDocument) SetDetails(v string) {
	o.Details = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *EventDocument) GetAttributes() map[string]interface{} {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventDocument) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *EventDocument) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *EventDocument) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetObjects returns the Objects field value if set, zero value otherwise.
func (o *EventDocument) GetObjects() []string {
	if o == nil || IsNil(o.Objects) {
		var ret []string
		return ret
	}
	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventDocument) GetObjectsOk() ([]string, bool) {
	if o == nil || IsNil(o.Objects) {
		return nil, false
	}
	return o.Objects, true
}

// HasObjects returns a boolean if a field has been set.
func (o *EventDocument) HasObjects() bool {
	if o != nil && !IsNil(o.Objects) {
		return true
	}

	return false
}

// SetObjects gets a reference to the given []string and assigns it to the Objects field.
func (o *EventDocument) SetObjects(v []string) {
	o.Objects = v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *EventDocument) GetOperation() string {
	if o == nil || IsNil(o.Operation) {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventDocument) GetOperationOk() (*string, bool) {
	if o == nil || IsNil(o.Operation) {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *EventDocument) HasOperation() bool {
	if o != nil && !IsNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *EventDocument) SetOperation(v string) {
	o.Operation = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EventDocument) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventDocument) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EventDocument) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *EventDocument) SetStatus(v string) {
	o.Status = &v
}

// GetTechnicalName returns the TechnicalName field value if set, zero value otherwise.
func (o *EventDocument) GetTechnicalName() string {
	if o == nil || IsNil(o.TechnicalName) {
		var ret string
		return ret
	}
	return *o.TechnicalName
}

// GetTechnicalNameOk returns a tuple with the TechnicalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventDocument) GetTechnicalNameOk() (*string, bool) {
	if o == nil || IsNil(o.TechnicalName) {
		return nil, false
	}
	return o.TechnicalName, true
}

// HasTechnicalName returns a boolean if a field has been set.
func (o *EventDocument) HasTechnicalName() bool {
	if o != nil && !IsNil(o.TechnicalName) {
		return true
	}

	return false
}

// SetTechnicalName gets a reference to the given string and assigns it to the TechnicalName field.
func (o *EventDocument) SetTechnicalName(v string) {
	o.TechnicalName = &v
}

func (o EventDocument) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventDocument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Created.IsSet() {
		toSerialize["created"] = o.Created.Get()
	}
	if !IsNil(o.Synced) {
		toSerialize["synced"] = o.Synced
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Actor) {
		toSerialize["actor"] = o.Actor
	}
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	if !IsNil(o.Stack) {
		toSerialize["stack"] = o.Stack
	}
	if !IsNil(o.TrackingNumber) {
		toSerialize["trackingNumber"] = o.TrackingNumber
	}
	if !IsNil(o.IpAddress) {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Objects) {
		toSerialize["objects"] = o.Objects
	}
	if !IsNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TechnicalName) {
		toSerialize["technicalName"] = o.TechnicalName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EventDocument) UnmarshalJSON(data []byte) (err error) {
	varEventDocument := _EventDocument{}

	err = json.Unmarshal(data, &varEventDocument)

	if err != nil {
		return err
	}

	*o = EventDocument(varEventDocument)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "created")
		delete(additionalProperties, "synced")
		delete(additionalProperties, "action")
		delete(additionalProperties, "type")
		delete(additionalProperties, "actor")
		delete(additionalProperties, "target")
		delete(additionalProperties, "stack")
		delete(additionalProperties, "trackingNumber")
		delete(additionalProperties, "ipAddress")
		delete(additionalProperties, "details")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "objects")
		delete(additionalProperties, "operation")
		delete(additionalProperties, "status")
		delete(additionalProperties, "technicalName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEventDocument struct {
	value *EventDocument
	isSet bool
}

func (v NullableEventDocument) Get() *EventDocument {
	return v.value
}

func (v *NullableEventDocument) Set(val *EventDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableEventDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableEventDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventDocument(val *EventDocument) *NullableEventDocument {
	return &NullableEventDocument{value: val, isSet: true}
}

func (v NullableEventDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


