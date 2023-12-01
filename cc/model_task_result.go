package cc

import (
	"context"
	"log"
	"time"
)

// checks if the TaskResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskResult{}

// TaskResult
type TaskResult struct {
	Description      *string `json:"description,omitempty"`
	Id               *string `json:"id,omitempty"`
	Launcher         *string `json:"launcher,omitempty"`
	Name             *string `json:"name,omitempty"`
	Type             *string `json:"type,omitempty"`
	Created          *int64  `json:"created,omitempty"`
	CompletionStatus *string `json:"completionStatus,omitempty"`
	Completed        *int64  `json:"completed,omitempty"`
	Progress         *string `json:"progress,omitempty"`

	// attributes       System.Management.Automation.PSCustomObject attributes=@{qpocJobId=88ae88c8-730a-4eb2-8d7f-bfdaf8e91a31}
	// completed        object completed=null
	// launched         object launched=null
	// messages         Object[] messages=System.Object[]
	// parentName       object parentName=null
	// returns          Object[] returns=System.Object[]
}

type _TaskResult TaskResult

type NullableTaskResult struct {
	value *TaskResult
	isSet bool
}

// NewTaskResult instantiates a new TaskResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskResult() *TaskResult {
	this := TaskResult{}
	return &this
}

func (o TaskResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Launcher) {
		toSerialize["launcher"] = o.Launcher
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !isNil(o.CompletionStatus) {
		toSerialize["completionstatus"] = o.CompletionStatus
	}
	if !isNil(o.Completed) {
		toSerialize["completed"] = o.Completed
	}
	if !isNil(o.Progress) {
		toSerialize["progress"] = o.Progress
	}

	return toSerialize, nil
}

// GetPolicyName returns the PolicyName field value if set, zero value otherwise.
func (o *TaskResult) GetPolicyName() string {
	if o == nil || isNil(o.CompletionStatus) {
		var ret string
		return ret
	}
	return *o.CompletionStatus
}

// WaitForTaskCompletion logic from https://github.com/pcassaretto-cmc/terraform-provider-identitynow/blob/c52b91fb218e4eebadf129985dfc4be814b3d05f/resource_source_entitlement_aggregation.go
func (o TaskResult) WaitForTaskCompletion(a APIClient) (bool, *string, error) {
	for {
		log.Printf("[WaitForTaskCompletion] polling taskId:%s", *o.Id)
		time.Sleep(200 * time.Millisecond)

		task, httpResp, err := a.TaskResultsApi.TaskResults(context.TODO(), *o.Id).Execute()
		if err != nil {
			log.Printf("httpResp:%s", httpResp)
			return false, nil, err
		}

		// log.Printf("[WaitForTaskCompletion] task completion status:%v progress:%v", *task.CompletionStatus, *task.Progress)

		// As there is no documentation around entitlement aggregation in the
		if task.CompletionStatus == nil {
			time.Sleep(20 * time.Second)
			continue
		}

		// This sleep is required for the API objects to become available post aggregation assuming the aggregation completed successfully
		time.Sleep(2 * time.Second)

		return true, task.CompletionStatus, nil
	}
}
