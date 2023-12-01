package cc

// checks if the UserList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserList{}

// UserList
type UserList struct {
	Total *int                      `json:"total"`
	Items *[]GetIdentity200Response `json:"items"`
}

type _UserList UserList

type NullableUserList struct {
	value *UserList
	isSet bool
}

// NewUserList instantiates a new UserList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserList() *UserList {
	this := UserList{}
	return &this
}

func (o UserList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}

	return toSerialize, nil
}
