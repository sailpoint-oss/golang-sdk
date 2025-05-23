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

// checks if the ResourceObjectsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceObjectsResponse{}

// ResourceObjectsResponse Response model for peek resource objects from source connectors.
type ResourceObjectsResponse struct {
	// ID of the source
	Id *string `json:"id,omitempty"`
	// Name of the source
	Name *string `json:"name,omitempty"`
	// The number of objects that were fetched by the connector.
	ObjectCount *int32 `json:"objectCount,omitempty"`
	// The number of milliseconds spent on the entire request.
	ElapsedMillis *int32 `json:"elapsedMillis,omitempty"`
	// Fetched objects from the source connector.
	ResourceObjects []ResourceObject `json:"resourceObjects,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceObjectsResponse ResourceObjectsResponse

// NewResourceObjectsResponse instantiates a new ResourceObjectsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceObjectsResponse() *ResourceObjectsResponse {
	this := ResourceObjectsResponse{}
	return &this
}

// NewResourceObjectsResponseWithDefaults instantiates a new ResourceObjectsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceObjectsResponseWithDefaults() *ResourceObjectsResponse {
	this := ResourceObjectsResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResourceObjectsResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceObjectsResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResourceObjectsResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ResourceObjectsResponse) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResourceObjectsResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceObjectsResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResourceObjectsResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ResourceObjectsResponse) SetName(v string) {
	o.Name = &v
}

// GetObjectCount returns the ObjectCount field value if set, zero value otherwise.
func (o *ResourceObjectsResponse) GetObjectCount() int32 {
	if o == nil || IsNil(o.ObjectCount) {
		var ret int32
		return ret
	}
	return *o.ObjectCount
}

// GetObjectCountOk returns a tuple with the ObjectCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceObjectsResponse) GetObjectCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ObjectCount) {
		return nil, false
	}
	return o.ObjectCount, true
}

// HasObjectCount returns a boolean if a field has been set.
func (o *ResourceObjectsResponse) HasObjectCount() bool {
	if o != nil && !IsNil(o.ObjectCount) {
		return true
	}

	return false
}

// SetObjectCount gets a reference to the given int32 and assigns it to the ObjectCount field.
func (o *ResourceObjectsResponse) SetObjectCount(v int32) {
	o.ObjectCount = &v
}

// GetElapsedMillis returns the ElapsedMillis field value if set, zero value otherwise.
func (o *ResourceObjectsResponse) GetElapsedMillis() int32 {
	if o == nil || IsNil(o.ElapsedMillis) {
		var ret int32
		return ret
	}
	return *o.ElapsedMillis
}

// GetElapsedMillisOk returns a tuple with the ElapsedMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceObjectsResponse) GetElapsedMillisOk() (*int32, bool) {
	if o == nil || IsNil(o.ElapsedMillis) {
		return nil, false
	}
	return o.ElapsedMillis, true
}

// HasElapsedMillis returns a boolean if a field has been set.
func (o *ResourceObjectsResponse) HasElapsedMillis() bool {
	if o != nil && !IsNil(o.ElapsedMillis) {
		return true
	}

	return false
}

// SetElapsedMillis gets a reference to the given int32 and assigns it to the ElapsedMillis field.
func (o *ResourceObjectsResponse) SetElapsedMillis(v int32) {
	o.ElapsedMillis = &v
}

// GetResourceObjects returns the ResourceObjects field value if set, zero value otherwise.
func (o *ResourceObjectsResponse) GetResourceObjects() []ResourceObject {
	if o == nil || IsNil(o.ResourceObjects) {
		var ret []ResourceObject
		return ret
	}
	return o.ResourceObjects
}

// GetResourceObjectsOk returns a tuple with the ResourceObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceObjectsResponse) GetResourceObjectsOk() ([]ResourceObject, bool) {
	if o == nil || IsNil(o.ResourceObjects) {
		return nil, false
	}
	return o.ResourceObjects, true
}

// HasResourceObjects returns a boolean if a field has been set.
func (o *ResourceObjectsResponse) HasResourceObjects() bool {
	if o != nil && !IsNil(o.ResourceObjects) {
		return true
	}

	return false
}

// SetResourceObjects gets a reference to the given []ResourceObject and assigns it to the ResourceObjects field.
func (o *ResourceObjectsResponse) SetResourceObjects(v []ResourceObject) {
	o.ResourceObjects = v
}

func (o ResourceObjectsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceObjectsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ObjectCount) {
		toSerialize["objectCount"] = o.ObjectCount
	}
	if !IsNil(o.ElapsedMillis) {
		toSerialize["elapsedMillis"] = o.ElapsedMillis
	}
	if !IsNil(o.ResourceObjects) {
		toSerialize["resourceObjects"] = o.ResourceObjects
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceObjectsResponse) UnmarshalJSON(data []byte) (err error) {
	varResourceObjectsResponse := _ResourceObjectsResponse{}

	err = json.Unmarshal(data, &varResourceObjectsResponse)

	if err != nil {
		return err
	}

	*o = ResourceObjectsResponse(varResourceObjectsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "objectCount")
		delete(additionalProperties, "elapsedMillis")
		delete(additionalProperties, "resourceObjects")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceObjectsResponse struct {
	value *ResourceObjectsResponse
	isSet bool
}

func (v NullableResourceObjectsResponse) Get() *ResourceObjectsResponse {
	return v.value
}

func (v *NullableResourceObjectsResponse) Set(val *ResourceObjectsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceObjectsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceObjectsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceObjectsResponse(val *ResourceObjectsResponse) *NullableResourceObjectsResponse {
	return &NullableResourceObjectsResponse{value: val, isSet: true}
}

func (v NullableResourceObjectsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceObjectsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


