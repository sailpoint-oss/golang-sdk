# CampaignReportAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportType** | Pointer to [**ReportType**](ReportType.md) |  | [optional] 
**LastRunAt** | Pointer to **time.Time** | The most recent date and time this report was run | [optional] [readonly] 

## Methods

### NewCampaignReportAllOf

`func NewCampaignReportAllOf() *CampaignReportAllOf`

NewCampaignReportAllOf instantiates a new CampaignReportAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignReportAllOfWithDefaults

`func NewCampaignReportAllOfWithDefaults() *CampaignReportAllOf`

NewCampaignReportAllOfWithDefaults instantiates a new CampaignReportAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportType

`func (o *CampaignReportAllOf) GetReportType() ReportType`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *CampaignReportAllOf) GetReportTypeOk() (*ReportType, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *CampaignReportAllOf) SetReportType(v ReportType)`

SetReportType sets ReportType field to given value.

### HasReportType

`func (o *CampaignReportAllOf) HasReportType() bool`

HasReportType returns a boolean if a field has been set.

### GetLastRunAt

`func (o *CampaignReportAllOf) GetLastRunAt() time.Time`

GetLastRunAt returns the LastRunAt field if non-nil, zero value otherwise.

### GetLastRunAtOk

`func (o *CampaignReportAllOf) GetLastRunAtOk() (*time.Time, bool)`

GetLastRunAtOk returns a tuple with the LastRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunAt

`func (o *CampaignReportAllOf) SetLastRunAt(v time.Time)`

SetLastRunAt sets LastRunAt field to given value.

### HasLastRunAt

`func (o *CampaignReportAllOf) HasLastRunAt() bool`

HasLastRunAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


