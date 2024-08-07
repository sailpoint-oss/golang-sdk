/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	"fmt"
)

// checks if the SearchExportReportArguments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchExportReportArguments{}

// SearchExportReportArguments Arguments for Search Export report (SEARCH_EXPORT)
type SearchExportReportArguments struct {
	// The names of the Elasticsearch indices in which to search. If none are provided, then all indices will be searched.
	Indices []Index `json:"indices,omitempty"`
	// The filters to be applied for each filtered field name.
	Filters *map[string]Filter `json:"filters,omitempty"`
	Query Query `json:"query"`
	// Indicates whether nested objects from returned search results should be included.
	IncludeNested *bool `json:"includeNested,omitempty"`
	// The fields to be used to sort the search results. Use + or - to specify the sort direction.
	Sort []string `json:"sort,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SearchExportReportArguments SearchExportReportArguments

// NewSearchExportReportArguments instantiates a new SearchExportReportArguments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchExportReportArguments(query Query) *SearchExportReportArguments {
	this := SearchExportReportArguments{}
	this.Query = query
	var includeNested bool = true
	this.IncludeNested = &includeNested
	return &this
}

// NewSearchExportReportArgumentsWithDefaults instantiates a new SearchExportReportArguments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchExportReportArgumentsWithDefaults() *SearchExportReportArguments {
	this := SearchExportReportArguments{}
	var includeNested bool = true
	this.IncludeNested = &includeNested
	return &this
}

// GetIndices returns the Indices field value if set, zero value otherwise.
func (o *SearchExportReportArguments) GetIndices() []Index {
	if o == nil || isNil(o.Indices) {
		var ret []Index
		return ret
	}
	return o.Indices
}

// GetIndicesOk returns a tuple with the Indices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchExportReportArguments) GetIndicesOk() ([]Index, bool) {
	if o == nil || isNil(o.Indices) {
		return nil, false
	}
	return o.Indices, true
}

// HasIndices returns a boolean if a field has been set.
func (o *SearchExportReportArguments) HasIndices() bool {
	if o != nil && !isNil(o.Indices) {
		return true
	}

	return false
}

// SetIndices gets a reference to the given []Index and assigns it to the Indices field.
func (o *SearchExportReportArguments) SetIndices(v []Index) {
	o.Indices = v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *SearchExportReportArguments) GetFilters() map[string]Filter {
	if o == nil || isNil(o.Filters) {
		var ret map[string]Filter
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchExportReportArguments) GetFiltersOk() (*map[string]Filter, bool) {
	if o == nil || isNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *SearchExportReportArguments) HasFilters() bool {
	if o != nil && !isNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given map[string]Filter and assigns it to the Filters field.
func (o *SearchExportReportArguments) SetFilters(v map[string]Filter) {
	o.Filters = &v
}

// GetQuery returns the Query field value
func (o *SearchExportReportArguments) GetQuery() Query {
	if o == nil {
		var ret Query
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *SearchExportReportArguments) GetQueryOk() (*Query, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *SearchExportReportArguments) SetQuery(v Query) {
	o.Query = v
}

// GetIncludeNested returns the IncludeNested field value if set, zero value otherwise.
func (o *SearchExportReportArguments) GetIncludeNested() bool {
	if o == nil || isNil(o.IncludeNested) {
		var ret bool
		return ret
	}
	return *o.IncludeNested
}

// GetIncludeNestedOk returns a tuple with the IncludeNested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchExportReportArguments) GetIncludeNestedOk() (*bool, bool) {
	if o == nil || isNil(o.IncludeNested) {
		return nil, false
	}
	return o.IncludeNested, true
}

// HasIncludeNested returns a boolean if a field has been set.
func (o *SearchExportReportArguments) HasIncludeNested() bool {
	if o != nil && !isNil(o.IncludeNested) {
		return true
	}

	return false
}

// SetIncludeNested gets a reference to the given bool and assigns it to the IncludeNested field.
func (o *SearchExportReportArguments) SetIncludeNested(v bool) {
	o.IncludeNested = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *SearchExportReportArguments) GetSort() []string {
	if o == nil || isNil(o.Sort) {
		var ret []string
		return ret
	}
	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchExportReportArguments) GetSortOk() ([]string, bool) {
	if o == nil || isNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *SearchExportReportArguments) HasSort() bool {
	if o != nil && !isNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given []string and assigns it to the Sort field.
func (o *SearchExportReportArguments) SetSort(v []string) {
	o.Sort = v
}

func (o SearchExportReportArguments) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchExportReportArguments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Indices) {
		toSerialize["indices"] = o.Indices
	}
	if !isNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	toSerialize["query"] = o.Query
	if !isNil(o.IncludeNested) {
		toSerialize["includeNested"] = o.IncludeNested
	}
	if !isNil(o.Sort) {
		toSerialize["sort"] = o.Sort
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SearchExportReportArguments) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"query",
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

	varSearchExportReportArguments := _SearchExportReportArguments{}

	if err = json.Unmarshal(bytes, &varSearchExportReportArguments); err == nil {
	*o = SearchExportReportArguments(varSearchExportReportArguments)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "indices")
		delete(additionalProperties, "filters")
		delete(additionalProperties, "query")
		delete(additionalProperties, "includeNested")
		delete(additionalProperties, "sort")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSearchExportReportArguments struct {
	value *SearchExportReportArguments
	isSet bool
}

func (v NullableSearchExportReportArguments) Get() *SearchExportReportArguments {
	return v.value
}

func (v *NullableSearchExportReportArguments) Set(val *SearchExportReportArguments) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchExportReportArguments) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchExportReportArguments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchExportReportArguments(val *SearchExportReportArguments) *NullableSearchExportReportArguments {
	return &NullableSearchExportReportArguments{value: val, isSet: true}
}

func (v NullableSearchExportReportArguments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchExportReportArguments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


