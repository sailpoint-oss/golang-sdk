# ClusterManualUpgrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | Pointer to [**[]ClusterManualUpgradeJobsInner**](ClusterManualUpgradeJobsInner.md) | List of job objects for the upgrade request. | [optional] 

## Methods

### NewClusterManualUpgrade

`func NewClusterManualUpgrade() *ClusterManualUpgrade`

NewClusterManualUpgrade instantiates a new ClusterManualUpgrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterManualUpgradeWithDefaults

`func NewClusterManualUpgradeWithDefaults() *ClusterManualUpgrade`

NewClusterManualUpgradeWithDefaults instantiates a new ClusterManualUpgrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobs

`func (o *ClusterManualUpgrade) GetJobs() []ClusterManualUpgradeJobsInner`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *ClusterManualUpgrade) GetJobsOk() (*[]ClusterManualUpgradeJobsInner, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *ClusterManualUpgrade) SetJobs(v []ClusterManualUpgradeJobsInner)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *ClusterManualUpgrade) HasJobs() bool`

HasJobs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


