package cc

// checks if the Event type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Event{}

// Event
type Event struct {
	Id          int    `json:"id"`
	Type        string `json:"type"`
	Status      string `json:"status"`
	Timestamp   string `json:"timestamp"`
	DateCreated string `json:"dateCreated"`
	Details     struct {
		Id string `json:"id"`
	} `json:"details"`
}

type _Event Event

type NullableEvent struct {
	value *Event
	isSet bool
}

// NewEvent instantiates a new Event object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvent() *Event {
	this := Event{}
	return &this
}

func (o Event) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !isNil(o.DateCreated) {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if !isNil(o.Details) {
		toSerialize["details"] = o.Details
	}

	return toSerialize, nil
}
