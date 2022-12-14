/*
SailPoint SaaS API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointv2sdk

import (
	"encoding/json"
	"time"
)

// ProvisioningActivity struct for ProvisioningActivity
type ProvisioningActivity struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	DateCreated *time.Time `json:"dateCreated,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	Operation *string `json:"operation,omitempty"`
	Status *string `json:"status,omitempty"`
	SourceId *string `json:"sourceId,omitempty"`
	SourceName *string `json:"sourceName,omitempty"`
	AccountName *string `json:"accountName,omitempty"`
	OwnerName *string `json:"ownerName,omitempty"`
	ApproverName *string `json:"approverName,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	Errors []string `json:"errors,omitempty"`
	ProvisioningPlan *string `json:"provisioningPlan,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProvisioningActivity ProvisioningActivity

// NewProvisioningActivity instantiates a new ProvisioningActivity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningActivity() *ProvisioningActivity {
	this := ProvisioningActivity{}
	return &this
}

// NewProvisioningActivityWithDefaults instantiates a new ProvisioningActivity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningActivityWithDefaults() *ProvisioningActivity {
	this := ProvisioningActivity{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProvisioningActivity) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningActivity) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProvisioningActivity) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProvisioningActivity) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProvisioningActivity) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningActivity) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProvisioningActivity) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProvisioningActivity) SetName(v string) {
	o.Name = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *ProvisioningActivity) GetDateCreated() time.Time {
	if o == nil || isNil(o.DateCreated) {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningActivity) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.DateCreated) {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *ProvisioningActivity) HasDateCreated() bool {
	if o != nil && !isNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *ProvisioningActivity) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *ProvisioningActivity) GetLastUpdated() time.Time {
	if o == nil || isNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningActivity) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *ProvisioningActivity) HasLastUpdated() bool {
	if o != nil && !isNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *ProvisioningActivity) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *ProvisioningActivity) GetOperation() string {
	if o == nil || isNil(o.Operation) {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningActivity) GetOperationOk() (*string, bool) {
	if o == nil || isNil(o.Operation) {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *ProvisioningActivity) HasOperation() bool {
	if o != nil && !isNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *ProvisioningActivity) SetOperation(v string) {
	o.Operation = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ProvisioningActivity) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningActivity) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ProvisioningActivity) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ProvisioningActivity) SetStatus(v string) {
	o.Status = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *ProvisioningActivity) GetSourceId() string {
	if o == nil || isNil(o.SourceId) {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningActivity) GetSourceIdOk() (*string, bool) {
	if o == nil || isNil(o.SourceId) {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *ProvisioningActivity) HasSourceId() bool {
	if o != nil && !isNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *ProvisioningActivity) SetSourceId(v string) {
	o.SourceId = &v
}

// GetSourceName returns the SourceName field value if set, zero value otherwise.
func (o *ProvisioningActivity) GetSourceName() string {
	if o == nil || isNil(o.SourceName) {
		var ret string
		return ret
	}
	return *o.SourceName
}

// GetSourceNameOk returns a tuple with the SourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningActivity) GetSourceNameOk() (*string, bool) {
	if o == nil || isNil(o.SourceName) {
		return nil, false
	}
	return o.SourceName, true
}

// HasSourceName returns a boolean if a field has been set.
func (o *ProvisioningActivity) HasSourceName() bool {
	if o != nil && !isNil(o.SourceName) {
		return true
	}

	return false
}

// SetSourceName gets a reference to the given string and assigns it to the SourceName field.
func (o *ProvisioningActivity) SetSourceName(v string) {
	o.SourceName = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *ProvisioningActivity) GetAccountName() string {
	if o == nil || isNil(o.AccountName) {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningActivity) GetAccountNameOk() (*string, bool) {
	if o == nil || isNil(o.AccountName) {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *ProvisioningActivity) HasAccountName() bool {
	if o != nil && !isNil(o.AccountName) {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *ProvisioningActivity) SetAccountName(v string) {
	o.AccountName = &v
}

// GetOwnerName returns the OwnerName field value if set, zero value otherwise.
func (o *ProvisioningActivity) GetOwnerName() string {
	if o == nil || isNil(o.OwnerName) {
		var ret string
		return ret
	}
	return *o.OwnerName
}

// GetOwnerNameOk returns a tuple with the OwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningActivity) GetOwnerNameOk() (*string, bool) {
	if o == nil || isNil(o.OwnerName) {
		return nil, false
	}
	return o.OwnerName, true
}

// HasOwnerName returns a boolean if a field has been set.
func (o *ProvisioningActivity) HasOwnerName() bool {
	if o != nil && !isNil(o.OwnerName) {
		return true
	}

	return false
}

// SetOwnerName gets a reference to the given string and assigns it to the OwnerName field.
func (o *ProvisioningActivity) SetOwnerName(v string) {
	o.OwnerName = &v
}

// GetApproverName returns the ApproverName field value if set, zero value otherwise.
func (o *ProvisioningActivity) GetApproverName() string {
	if o == nil || isNil(o.ApproverName) {
		var ret string
		return ret
	}
	return *o.ApproverName
}

// GetApproverNameOk returns a tuple with the ApproverName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningActivity) GetApproverNameOk() (*string, bool) {
	if o == nil || isNil(o.ApproverName) {
		return nil, false
	}
	return o.ApproverName, true
}

// HasApproverName returns a boolean if a field has been set.
func (o *ProvisioningActivity) HasApproverName() bool {
	if o != nil && !isNil(o.ApproverName) {
		return true
	}

	return false
}

// SetApproverName gets a reference to the given string and assigns it to the ApproverName field.
func (o *ProvisioningActivity) SetApproverName(v string) {
	o.ApproverName = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ProvisioningActivity) GetWarnings() []string {
	if o == nil || isNil(o.Warnings) {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningActivity) GetWarningsOk() ([]string, bool) {
	if o == nil || isNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ProvisioningActivity) HasWarnings() bool {
	if o != nil && !isNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *ProvisioningActivity) SetWarnings(v []string) {
	o.Warnings = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ProvisioningActivity) GetErrors() []string {
	if o == nil || isNil(o.Errors) {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningActivity) GetErrorsOk() ([]string, bool) {
	if o == nil || isNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ProvisioningActivity) HasErrors() bool {
	if o != nil && !isNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *ProvisioningActivity) SetErrors(v []string) {
	o.Errors = v
}

// GetProvisioningPlan returns the ProvisioningPlan field value if set, zero value otherwise.
func (o *ProvisioningActivity) GetProvisioningPlan() string {
	if o == nil || isNil(o.ProvisioningPlan) {
		var ret string
		return ret
	}
	return *o.ProvisioningPlan
}

// GetProvisioningPlanOk returns a tuple with the ProvisioningPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningActivity) GetProvisioningPlanOk() (*string, bool) {
	if o == nil || isNil(o.ProvisioningPlan) {
		return nil, false
	}
	return o.ProvisioningPlan, true
}

// HasProvisioningPlan returns a boolean if a field has been set.
func (o *ProvisioningActivity) HasProvisioningPlan() bool {
	if o != nil && !isNil(o.ProvisioningPlan) {
		return true
	}

	return false
}

// SetProvisioningPlan gets a reference to the given string and assigns it to the ProvisioningPlan field.
func (o *ProvisioningActivity) SetProvisioningPlan(v string) {
	o.ProvisioningPlan = &v
}

func (o ProvisioningActivity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.DateCreated) {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if !isNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !isNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.SourceId) {
		toSerialize["sourceId"] = o.SourceId
	}
	if !isNil(o.SourceName) {
		toSerialize["sourceName"] = o.SourceName
	}
	if !isNil(o.AccountName) {
		toSerialize["accountName"] = o.AccountName
	}
	if !isNil(o.OwnerName) {
		toSerialize["ownerName"] = o.OwnerName
	}
	if !isNil(o.ApproverName) {
		toSerialize["approverName"] = o.ApproverName
	}
	if !isNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	if !isNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !isNil(o.ProvisioningPlan) {
		toSerialize["provisioningPlan"] = o.ProvisioningPlan
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProvisioningActivity) UnmarshalJSON(bytes []byte) (err error) {
	varProvisioningActivity := _ProvisioningActivity{}

	if err = json.Unmarshal(bytes, &varProvisioningActivity); err == nil {
		*o = ProvisioningActivity(varProvisioningActivity)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "dateCreated")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "operation")
		delete(additionalProperties, "status")
		delete(additionalProperties, "sourceId")
		delete(additionalProperties, "sourceName")
		delete(additionalProperties, "accountName")
		delete(additionalProperties, "ownerName")
		delete(additionalProperties, "approverName")
		delete(additionalProperties, "warnings")
		delete(additionalProperties, "errors")
		delete(additionalProperties, "provisioningPlan")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProvisioningActivity struct {
	value *ProvisioningActivity
	isSet bool
}

func (v NullableProvisioningActivity) Get() *ProvisioningActivity {
	return v.value
}

func (v *NullableProvisioningActivity) Set(val *ProvisioningActivity) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningActivity(val *ProvisioningActivity) *NullableProvisioningActivity {
	return &NullableProvisioningActivity{value: val, isSet: true}
}

func (v NullableProvisioningActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


