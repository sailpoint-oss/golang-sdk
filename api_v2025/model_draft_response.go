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

// checks if the DraftResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DraftResponse{}

// DraftResponse struct for DraftResponse
type DraftResponse struct {
	// Unique id assigned to this job.
	JobId *string `json:"jobId,omitempty"`
	// Status of the job.
	Status *string `json:"status,omitempty"`
	// Type of the job, will always be CREATE_DRAFT for this type of job.
	Type *string `json:"type,omitempty"`
	// Message providing information about the outcome of the draft process.
	Message *string `json:"message,omitempty"`
	// The name of user that that initiated the draft process.
	RequesterName *string `json:"requesterName,omitempty"`
	// Whether or not a file was generated for this draft.
	FileExists *bool `json:"fileExists,omitempty"`
	// The time the job was started.
	Created *SailPointTime `json:"created,omitempty"`
	// The time of the last update to the job.
	Modified *SailPointTime `json:"modified,omitempty"`
	// The time the job was completed.
	Completed *SailPointTime `json:"completed,omitempty"`
	// Name of the draft.
	Name *string `json:"name,omitempty"`
	// Tenant owner of the backup from which the draft was generated.
	SourceTenant *string `json:"sourceTenant,omitempty"`
	// Id of the backup from which the draft was generated.
	SourceBackupId *string `json:"sourceBackupId,omitempty"`
	// Name of the backup from which the draft was generated.
	SourceBackupName *string `json:"sourceBackupName,omitempty"`
	// Denotes the origin of the source backup from which the draft was generated. - RESTORE - Same tenant. - PROMOTE - Different tenant. - UPLOAD - Uploaded configuration.
	Mode *string `json:"mode,omitempty"`
	// Approval status of the draft used to determine whether or not the draft can be deployed.
	ApprovalStatus *string `json:"approvalStatus,omitempty"`
	// List of comments that have been exchanged between an approval requester and an approver.
	ApprovalComment []ApprovalComment `json:"approvalComment,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DraftResponse DraftResponse

// NewDraftResponse instantiates a new DraftResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDraftResponse() *DraftResponse {
	this := DraftResponse{}
	var fileExists bool = true
	this.FileExists = &fileExists
	return &this
}

// NewDraftResponseWithDefaults instantiates a new DraftResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDraftResponseWithDefaults() *DraftResponse {
	this := DraftResponse{}
	var fileExists bool = true
	this.FileExists = &fileExists
	return &this
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *DraftResponse) GetJobId() string {
	if o == nil || IsNil(o.JobId) {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DraftResponse) GetJobIdOk() (*string, bool) {
	if o == nil || IsNil(o.JobId) {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *DraftResponse) HasJobId() bool {
	if o != nil && !IsNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
func (o *DraftResponse) SetJobId(v string) {
	o.JobId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DraftResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DraftResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DraftResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DraftResponse) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DraftResponse) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DraftResponse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DraftResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DraftResponse) SetType(v string) {
	o.Type = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DraftResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DraftResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DraftResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *DraftResponse) SetMessage(v string) {
	o.Message = &v
}

// GetRequesterName returns the RequesterName field value if set, zero value otherwise.
func (o *DraftResponse) GetRequesterName() string {
	if o == nil || IsNil(o.RequesterName) {
		var ret string
		return ret
	}
	return *o.RequesterName
}

// GetRequesterNameOk returns a tuple with the RequesterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DraftResponse) GetRequesterNameOk() (*string, bool) {
	if o == nil || IsNil(o.RequesterName) {
		return nil, false
	}
	return o.RequesterName, true
}

// HasRequesterName returns a boolean if a field has been set.
func (o *DraftResponse) HasRequesterName() bool {
	if o != nil && !IsNil(o.RequesterName) {
		return true
	}

	return false
}

// SetRequesterName gets a reference to the given string and assigns it to the RequesterName field.
func (o *DraftResponse) SetRequesterName(v string) {
	o.RequesterName = &v
}

// GetFileExists returns the FileExists field value if set, zero value otherwise.
func (o *DraftResponse) GetFileExists() bool {
	if o == nil || IsNil(o.FileExists) {
		var ret bool
		return ret
	}
	return *o.FileExists
}

// GetFileExistsOk returns a tuple with the FileExists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DraftResponse) GetFileExistsOk() (*bool, bool) {
	if o == nil || IsNil(o.FileExists) {
		return nil, false
	}
	return o.FileExists, true
}

// HasFileExists returns a boolean if a field has been set.
func (o *DraftResponse) HasFileExists() bool {
	if o != nil && !IsNil(o.FileExists) {
		return true
	}

	return false
}

// SetFileExists gets a reference to the given bool and assigns it to the FileExists field.
func (o *DraftResponse) SetFileExists(v bool) {
	o.FileExists = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *DraftResponse) GetCreated() SailPointTime {
	if o == nil || IsNil(o.Created) {
		var ret SailPointTime
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DraftResponse) GetCreatedOk() (*SailPointTime, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *DraftResponse) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given SailPointTime and assigns it to the Created field.
func (o *DraftResponse) SetCreated(v SailPointTime) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *DraftResponse) GetModified() SailPointTime {
	if o == nil || IsNil(o.Modified) {
		var ret SailPointTime
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DraftResponse) GetModifiedOk() (*SailPointTime, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *DraftResponse) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given SailPointTime and assigns it to the Modified field.
func (o *DraftResponse) SetModified(v SailPointTime) {
	o.Modified = &v
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *DraftResponse) GetCompleted() SailPointTime {
	if o == nil || IsNil(o.Completed) {
		var ret SailPointTime
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DraftResponse) GetCompletedOk() (*SailPointTime, bool) {
	if o == nil || IsNil(o.Completed) {
		return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *DraftResponse) HasCompleted() bool {
	if o != nil && !IsNil(o.Completed) {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given SailPointTime and assigns it to the Completed field.
func (o *DraftResponse) SetCompleted(v SailPointTime) {
	o.Completed = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DraftResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DraftResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DraftResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DraftResponse) SetName(v string) {
	o.Name = &v
}

// GetSourceTenant returns the SourceTenant field value if set, zero value otherwise.
func (o *DraftResponse) GetSourceTenant() string {
	if o == nil || IsNil(o.SourceTenant) {
		var ret string
		return ret
	}
	return *o.SourceTenant
}

// GetSourceTenantOk returns a tuple with the SourceTenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DraftResponse) GetSourceTenantOk() (*string, bool) {
	if o == nil || IsNil(o.SourceTenant) {
		return nil, false
	}
	return o.SourceTenant, true
}

// HasSourceTenant returns a boolean if a field has been set.
func (o *DraftResponse) HasSourceTenant() bool {
	if o != nil && !IsNil(o.SourceTenant) {
		return true
	}

	return false
}

// SetSourceTenant gets a reference to the given string and assigns it to the SourceTenant field.
func (o *DraftResponse) SetSourceTenant(v string) {
	o.SourceTenant = &v
}

// GetSourceBackupId returns the SourceBackupId field value if set, zero value otherwise.
func (o *DraftResponse) GetSourceBackupId() string {
	if o == nil || IsNil(o.SourceBackupId) {
		var ret string
		return ret
	}
	return *o.SourceBackupId
}

// GetSourceBackupIdOk returns a tuple with the SourceBackupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DraftResponse) GetSourceBackupIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceBackupId) {
		return nil, false
	}
	return o.SourceBackupId, true
}

// HasSourceBackupId returns a boolean if a field has been set.
func (o *DraftResponse) HasSourceBackupId() bool {
	if o != nil && !IsNil(o.SourceBackupId) {
		return true
	}

	return false
}

// SetSourceBackupId gets a reference to the given string and assigns it to the SourceBackupId field.
func (o *DraftResponse) SetSourceBackupId(v string) {
	o.SourceBackupId = &v
}

// GetSourceBackupName returns the SourceBackupName field value if set, zero value otherwise.
func (o *DraftResponse) GetSourceBackupName() string {
	if o == nil || IsNil(o.SourceBackupName) {
		var ret string
		return ret
	}
	return *o.SourceBackupName
}

// GetSourceBackupNameOk returns a tuple with the SourceBackupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DraftResponse) GetSourceBackupNameOk() (*string, bool) {
	if o == nil || IsNil(o.SourceBackupName) {
		return nil, false
	}
	return o.SourceBackupName, true
}

// HasSourceBackupName returns a boolean if a field has been set.
func (o *DraftResponse) HasSourceBackupName() bool {
	if o != nil && !IsNil(o.SourceBackupName) {
		return true
	}

	return false
}

// SetSourceBackupName gets a reference to the given string and assigns it to the SourceBackupName field.
func (o *DraftResponse) SetSourceBackupName(v string) {
	o.SourceBackupName = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *DraftResponse) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DraftResponse) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *DraftResponse) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *DraftResponse) SetMode(v string) {
	o.Mode = &v
}

// GetApprovalStatus returns the ApprovalStatus field value if set, zero value otherwise.
func (o *DraftResponse) GetApprovalStatus() string {
	if o == nil || IsNil(o.ApprovalStatus) {
		var ret string
		return ret
	}
	return *o.ApprovalStatus
}

// GetApprovalStatusOk returns a tuple with the ApprovalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DraftResponse) GetApprovalStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ApprovalStatus) {
		return nil, false
	}
	return o.ApprovalStatus, true
}

// HasApprovalStatus returns a boolean if a field has been set.
func (o *DraftResponse) HasApprovalStatus() bool {
	if o != nil && !IsNil(o.ApprovalStatus) {
		return true
	}

	return false
}

// SetApprovalStatus gets a reference to the given string and assigns it to the ApprovalStatus field.
func (o *DraftResponse) SetApprovalStatus(v string) {
	o.ApprovalStatus = &v
}

// GetApprovalComment returns the ApprovalComment field value if set, zero value otherwise.
func (o *DraftResponse) GetApprovalComment() []ApprovalComment {
	if o == nil || IsNil(o.ApprovalComment) {
		var ret []ApprovalComment
		return ret
	}
	return o.ApprovalComment
}

// GetApprovalCommentOk returns a tuple with the ApprovalComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DraftResponse) GetApprovalCommentOk() ([]ApprovalComment, bool) {
	if o == nil || IsNil(o.ApprovalComment) {
		return nil, false
	}
	return o.ApprovalComment, true
}

// HasApprovalComment returns a boolean if a field has been set.
func (o *DraftResponse) HasApprovalComment() bool {
	if o != nil && !IsNil(o.ApprovalComment) {
		return true
	}

	return false
}

// SetApprovalComment gets a reference to the given []ApprovalComment and assigns it to the ApprovalComment field.
func (o *DraftResponse) SetApprovalComment(v []ApprovalComment) {
	o.ApprovalComment = v
}

func (o DraftResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DraftResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JobId) {
		toSerialize["jobId"] = o.JobId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.RequesterName) {
		toSerialize["requesterName"] = o.RequesterName
	}
	if !IsNil(o.FileExists) {
		toSerialize["fileExists"] = o.FileExists
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !IsNil(o.Completed) {
		toSerialize["completed"] = o.Completed
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.SourceTenant) {
		toSerialize["sourceTenant"] = o.SourceTenant
	}
	if !IsNil(o.SourceBackupId) {
		toSerialize["sourceBackupId"] = o.SourceBackupId
	}
	if !IsNil(o.SourceBackupName) {
		toSerialize["sourceBackupName"] = o.SourceBackupName
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.ApprovalStatus) {
		toSerialize["approvalStatus"] = o.ApprovalStatus
	}
	if !IsNil(o.ApprovalComment) {
		toSerialize["approvalComment"] = o.ApprovalComment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DraftResponse) UnmarshalJSON(data []byte) (err error) {
	varDraftResponse := _DraftResponse{}

	err = json.Unmarshal(data, &varDraftResponse)

	if err != nil {
		return err
	}

	*o = DraftResponse(varDraftResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "jobId")
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		delete(additionalProperties, "message")
		delete(additionalProperties, "requesterName")
		delete(additionalProperties, "fileExists")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "completed")
		delete(additionalProperties, "name")
		delete(additionalProperties, "sourceTenant")
		delete(additionalProperties, "sourceBackupId")
		delete(additionalProperties, "sourceBackupName")
		delete(additionalProperties, "mode")
		delete(additionalProperties, "approvalStatus")
		delete(additionalProperties, "approvalComment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDraftResponse struct {
	value *DraftResponse
	isSet bool
}

func (v NullableDraftResponse) Get() *DraftResponse {
	return v.value
}

func (v *NullableDraftResponse) Set(val *DraftResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDraftResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDraftResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDraftResponse(val *DraftResponse) *NullableDraftResponse {
	return &NullableDraftResponse{value: val, isSet: true}
}

func (v NullableDraftResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDraftResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


