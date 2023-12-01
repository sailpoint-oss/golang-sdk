package cc

// checks if the EventList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventList{}

// EventList
type EventList struct {
	Total *int             `json:"total"`
	Items *[]Event `json:"items"`
}

type _EventList EventList

type NullableEventList struct {
	value *EventList
	isSet bool
}

// NewEventList instantiates a new EventList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventList() *EventList {
	this := EventList{}
	return &this
}

func (o EventList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}

	return toSerialize, nil
}
