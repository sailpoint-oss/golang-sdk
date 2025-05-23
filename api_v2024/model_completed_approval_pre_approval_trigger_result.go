/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	
)

// checks if the CompletedApprovalPreApprovalTriggerResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompletedApprovalPreApprovalTriggerResult{}

// CompletedApprovalPreApprovalTriggerResult If the access request submitted event trigger is configured and this access request was intercepted by it, then this is the result of the trigger's decision to either approve or deny the request.
type CompletedApprovalPreApprovalTriggerResult struct {
	// The comment from the trigger
	Comment *string `json:"comment,omitempty"`
	Decision *CompletedApprovalState `json:"decision,omitempty"`
	// The name of the approver
	Reviewer *string `json:"reviewer,omitempty"`
	// The date and time the trigger decided on the request
	Date *SailPointTime `json:"date,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CompletedApprovalPreApprovalTriggerResult CompletedApprovalPreApprovalTriggerResult

// NewCompletedApprovalPreApprovalTriggerResult instantiates a new CompletedApprovalPreApprovalTriggerResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompletedApprovalPreApprovalTriggerResult() *CompletedApprovalPreApprovalTriggerResult {
	this := CompletedApprovalPreApprovalTriggerResult{}
	return &this
}

// NewCompletedApprovalPreApprovalTriggerResultWithDefaults instantiates a new CompletedApprovalPreApprovalTriggerResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompletedApprovalPreApprovalTriggerResultWithDefaults() *CompletedApprovalPreApprovalTriggerResult {
	this := CompletedApprovalPreApprovalTriggerResult{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *CompletedApprovalPreApprovalTriggerResult) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletedApprovalPreApprovalTriggerResult) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *CompletedApprovalPreApprovalTriggerResult) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *CompletedApprovalPreApprovalTriggerResult) SetComment(v string) {
	o.Comment = &v
}

// GetDecision returns the Decision field value if set, zero value otherwise.
func (o *CompletedApprovalPreApprovalTriggerResult) GetDecision() CompletedApprovalState {
	if o == nil || IsNil(o.Decision) {
		var ret CompletedApprovalState
		return ret
	}
	return *o.Decision
}

// GetDecisionOk returns a tuple with the Decision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletedApprovalPreApprovalTriggerResult) GetDecisionOk() (*CompletedApprovalState, bool) {
	if o == nil || IsNil(o.Decision) {
		return nil, false
	}
	return o.Decision, true
}

// HasDecision returns a boolean if a field has been set.
func (o *CompletedApprovalPreApprovalTriggerResult) HasDecision() bool {
	if o != nil && !IsNil(o.Decision) {
		return true
	}

	return false
}

// SetDecision gets a reference to the given CompletedApprovalState and assigns it to the Decision field.
func (o *CompletedApprovalPreApprovalTriggerResult) SetDecision(v CompletedApprovalState) {
	o.Decision = &v
}

// GetReviewer returns the Reviewer field value if set, zero value otherwise.
func (o *CompletedApprovalPreApprovalTriggerResult) GetReviewer() string {
	if o == nil || IsNil(o.Reviewer) {
		var ret string
		return ret
	}
	return *o.Reviewer
}

// GetReviewerOk returns a tuple with the Reviewer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletedApprovalPreApprovalTriggerResult) GetReviewerOk() (*string, bool) {
	if o == nil || IsNil(o.Reviewer) {
		return nil, false
	}
	return o.Reviewer, true
}

// HasReviewer returns a boolean if a field has been set.
func (o *CompletedApprovalPreApprovalTriggerResult) HasReviewer() bool {
	if o != nil && !IsNil(o.Reviewer) {
		return true
	}

	return false
}

// SetReviewer gets a reference to the given string and assigns it to the Reviewer field.
func (o *CompletedApprovalPreApprovalTriggerResult) SetReviewer(v string) {
	o.Reviewer = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *CompletedApprovalPreApprovalTriggerResult) GetDate() SailPointTime {
	if o == nil || IsNil(o.Date) {
		var ret SailPointTime
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletedApprovalPreApprovalTriggerResult) GetDateOk() (*SailPointTime, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *CompletedApprovalPreApprovalTriggerResult) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given SailPointTime and assigns it to the Date field.
func (o *CompletedApprovalPreApprovalTriggerResult) SetDate(v SailPointTime) {
	o.Date = &v
}

func (o CompletedApprovalPreApprovalTriggerResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompletedApprovalPreApprovalTriggerResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Decision) {
		toSerialize["decision"] = o.Decision
	}
	if !IsNil(o.Reviewer) {
		toSerialize["reviewer"] = o.Reviewer
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CompletedApprovalPreApprovalTriggerResult) UnmarshalJSON(data []byte) (err error) {
	varCompletedApprovalPreApprovalTriggerResult := _CompletedApprovalPreApprovalTriggerResult{}

	err = json.Unmarshal(data, &varCompletedApprovalPreApprovalTriggerResult)

	if err != nil {
		return err
	}

	*o = CompletedApprovalPreApprovalTriggerResult(varCompletedApprovalPreApprovalTriggerResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "comment")
		delete(additionalProperties, "decision")
		delete(additionalProperties, "reviewer")
		delete(additionalProperties, "date")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCompletedApprovalPreApprovalTriggerResult struct {
	value *CompletedApprovalPreApprovalTriggerResult
	isSet bool
}

func (v NullableCompletedApprovalPreApprovalTriggerResult) Get() *CompletedApprovalPreApprovalTriggerResult {
	return v.value
}

func (v *NullableCompletedApprovalPreApprovalTriggerResult) Set(val *CompletedApprovalPreApprovalTriggerResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCompletedApprovalPreApprovalTriggerResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCompletedApprovalPreApprovalTriggerResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompletedApprovalPreApprovalTriggerResult(val *CompletedApprovalPreApprovalTriggerResult) *NullableCompletedApprovalPreApprovalTriggerResult {
	return &NullableCompletedApprovalPreApprovalTriggerResult{value: val, isSet: true}
}

func (v NullableCompletedApprovalPreApprovalTriggerResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompletedApprovalPreApprovalTriggerResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


