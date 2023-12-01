package cc

// checks if the LoadEntitlementResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadEntitlementResponse{}

// LoadEntitlementResponse
type LoadEntitlementResponse struct {
	Success *bool                `json:"success"`
	Task    *TaskResult `json:"task"`
}

type _LoadEntitlementResponse LoadEntitlementResponse

func (o LoadEntitlementResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !isNil(o.Task) {
		toSerialize["task"] = o.Task
	}

	return toSerialize, nil
}
