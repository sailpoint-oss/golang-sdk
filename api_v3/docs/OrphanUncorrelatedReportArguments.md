# OrphanUncorrelatedReportArguments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SelectedFormats** | Pointer to **[]string** | Output report file formats. This are formats for calling get endpoint as a query parameter &#39;fileFormat&#39;.  In case report won&#39;t have this argument there will be [&#39;CSV&#39;, &#39;PDF&#39;] as default. | [optional] 

## Methods

### NewOrphanUncorrelatedReportArguments

`func NewOrphanUncorrelatedReportArguments() *OrphanUncorrelatedReportArguments`

NewOrphanUncorrelatedReportArguments instantiates a new OrphanUncorrelatedReportArguments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrphanUncorrelatedReportArgumentsWithDefaults

`func NewOrphanUncorrelatedReportArgumentsWithDefaults() *OrphanUncorrelatedReportArguments`

NewOrphanUncorrelatedReportArgumentsWithDefaults instantiates a new OrphanUncorrelatedReportArguments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelectedFormats

`func (o *OrphanUncorrelatedReportArguments) GetSelectedFormats() []string`

GetSelectedFormats returns the SelectedFormats field if non-nil, zero value otherwise.

### GetSelectedFormatsOk

`func (o *OrphanUncorrelatedReportArguments) GetSelectedFormatsOk() (*[]string, bool)`

GetSelectedFormatsOk returns a tuple with the SelectedFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedFormats

`func (o *OrphanUncorrelatedReportArguments) SetSelectedFormats(v []string)`

SetSelectedFormats sets SelectedFormats field to given value.

### HasSelectedFormats

`func (o *OrphanUncorrelatedReportArguments) HasSelectedFormats() bool`

HasSelectedFormats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


