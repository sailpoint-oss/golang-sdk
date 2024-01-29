/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
	"fmt"
)

// checks if the TypeAheadQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TypeAheadQuery{}

// TypeAheadQuery Query parameters used to construct an Elasticsearch type ahead query object.  The typeAheadQuery performs a search for top values beginning with the typed values. For example, typing \"Jo\" results in top hits matching \"Jo.\" Typing \"Job\" results in top hits matching \"Job.\" 
type TypeAheadQuery struct {
	// The type ahead query string used to construct a phrase prefix match query.
	Query string `json:"query"`
	// The field on which to perform the type ahead search.
	Field string `json:"field"`
	// The nested type.
	NestedType *string `json:"nestedType,omitempty"`
	// The number of suffixes the last term will be expanded into. Influences the performance of the query and the number results returned. Valid values: 1 to 1000.
	MaxExpansions *int32 `json:"maxExpansions,omitempty"`
	// The max amount of records the search will return.
	Size *int32 `json:"size,omitempty"`
	// The sort order of the returned records.
	Sort *string `json:"sort,omitempty"`
	// The flag that defines the sort type, by count or value.
	SortByValue *bool `json:"sortByValue,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TypeAheadQuery TypeAheadQuery

// NewTypeAheadQuery instantiates a new TypeAheadQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTypeAheadQuery(query string, field string) *TypeAheadQuery {
	this := TypeAheadQuery{}
	this.Query = query
	this.Field = field
	var maxExpansions int32 = 10
	this.MaxExpansions = &maxExpansions
	var size int32 = 100
	this.Size = &size
	var sort string = "desc"
	this.Sort = &sort
	var sortByValue bool = false
	this.SortByValue = &sortByValue
	return &this
}

// NewTypeAheadQueryWithDefaults instantiates a new TypeAheadQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypeAheadQueryWithDefaults() *TypeAheadQuery {
	this := TypeAheadQuery{}
	var maxExpansions int32 = 10
	this.MaxExpansions = &maxExpansions
	var size int32 = 100
	this.Size = &size
	var sort string = "desc"
	this.Sort = &sort
	var sortByValue bool = false
	this.SortByValue = &sortByValue
	return &this
}

// GetQuery returns the Query field value
func (o *TypeAheadQuery) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *TypeAheadQuery) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *TypeAheadQuery) SetQuery(v string) {
	o.Query = v
}

// GetField returns the Field field value
func (o *TypeAheadQuery) GetField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *TypeAheadQuery) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *TypeAheadQuery) SetField(v string) {
	o.Field = v
}

// GetNestedType returns the NestedType field value if set, zero value otherwise.
func (o *TypeAheadQuery) GetNestedType() string {
	if o == nil || isNil(o.NestedType) {
		var ret string
		return ret
	}
	return *o.NestedType
}

// GetNestedTypeOk returns a tuple with the NestedType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypeAheadQuery) GetNestedTypeOk() (*string, bool) {
	if o == nil || isNil(o.NestedType) {
		return nil, false
	}
	return o.NestedType, true
}

// HasNestedType returns a boolean if a field has been set.
func (o *TypeAheadQuery) HasNestedType() bool {
	if o != nil && !isNil(o.NestedType) {
		return true
	}

	return false
}

// SetNestedType gets a reference to the given string and assigns it to the NestedType field.
func (o *TypeAheadQuery) SetNestedType(v string) {
	o.NestedType = &v
}

// GetMaxExpansions returns the MaxExpansions field value if set, zero value otherwise.
func (o *TypeAheadQuery) GetMaxExpansions() int32 {
	if o == nil || isNil(o.MaxExpansions) {
		var ret int32
		return ret
	}
	return *o.MaxExpansions
}

// GetMaxExpansionsOk returns a tuple with the MaxExpansions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypeAheadQuery) GetMaxExpansionsOk() (*int32, bool) {
	if o == nil || isNil(o.MaxExpansions) {
		return nil, false
	}
	return o.MaxExpansions, true
}

// HasMaxExpansions returns a boolean if a field has been set.
func (o *TypeAheadQuery) HasMaxExpansions() bool {
	if o != nil && !isNil(o.MaxExpansions) {
		return true
	}

	return false
}

// SetMaxExpansions gets a reference to the given int32 and assigns it to the MaxExpansions field.
func (o *TypeAheadQuery) SetMaxExpansions(v int32) {
	o.MaxExpansions = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *TypeAheadQuery) GetSize() int32 {
	if o == nil || isNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypeAheadQuery) GetSizeOk() (*int32, bool) {
	if o == nil || isNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *TypeAheadQuery) HasSize() bool {
	if o != nil && !isNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *TypeAheadQuery) SetSize(v int32) {
	o.Size = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *TypeAheadQuery) GetSort() string {
	if o == nil || isNil(o.Sort) {
		var ret string
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypeAheadQuery) GetSortOk() (*string, bool) {
	if o == nil || isNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *TypeAheadQuery) HasSort() bool {
	if o != nil && !isNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given string and assigns it to the Sort field.
func (o *TypeAheadQuery) SetSort(v string) {
	o.Sort = &v
}

// GetSortByValue returns the SortByValue field value if set, zero value otherwise.
func (o *TypeAheadQuery) GetSortByValue() bool {
	if o == nil || isNil(o.SortByValue) {
		var ret bool
		return ret
	}
	return *o.SortByValue
}

// GetSortByValueOk returns a tuple with the SortByValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypeAheadQuery) GetSortByValueOk() (*bool, bool) {
	if o == nil || isNil(o.SortByValue) {
		return nil, false
	}
	return o.SortByValue, true
}

// HasSortByValue returns a boolean if a field has been set.
func (o *TypeAheadQuery) HasSortByValue() bool {
	if o != nil && !isNil(o.SortByValue) {
		return true
	}

	return false
}

// SetSortByValue gets a reference to the given bool and assigns it to the SortByValue field.
func (o *TypeAheadQuery) SetSortByValue(v bool) {
	o.SortByValue = &v
}

func (o TypeAheadQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TypeAheadQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["query"] = o.Query
	toSerialize["field"] = o.Field
	if !isNil(o.NestedType) {
		toSerialize["nestedType"] = o.NestedType
	}
	if !isNil(o.MaxExpansions) {
		toSerialize["maxExpansions"] = o.MaxExpansions
	}
	if !isNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !isNil(o.Sort) {
		toSerialize["sort"] = o.Sort
	}
	if !isNil(o.SortByValue) {
		toSerialize["sortByValue"] = o.SortByValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TypeAheadQuery) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"query",
		"field",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTypeAheadQuery := _TypeAheadQuery{}

	if err = json.Unmarshal(bytes, &varTypeAheadQuery); err == nil {
	*o = TypeAheadQuery(varTypeAheadQuery)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "query")
		delete(additionalProperties, "field")
		delete(additionalProperties, "nestedType")
		delete(additionalProperties, "maxExpansions")
		delete(additionalProperties, "size")
		delete(additionalProperties, "sort")
		delete(additionalProperties, "sortByValue")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTypeAheadQuery struct {
	value *TypeAheadQuery
	isSet bool
}

func (v NullableTypeAheadQuery) Get() *TypeAheadQuery {
	return v.value
}

func (v *NullableTypeAheadQuery) Set(val *TypeAheadQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableTypeAheadQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableTypeAheadQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypeAheadQuery(val *TypeAheadQuery) *NullableTypeAheadQuery {
	return &NullableTypeAheadQuery{value: val, isSet: true}
}

func (v NullableTypeAheadQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypeAheadQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

