# ReportResultReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**DtoType**](DtoType.md) |  | [optional] 
**Id** | Pointer to **string** | ID of the object to which this reference applies | [optional] 
**Name** | Pointer to **string** | Human-readable display name of the object to which this reference applies | [optional] 
**Status** | Pointer to **string** | Status of a violation report | [optional] 

## Methods

### NewReportResultReference

`func NewReportResultReference() *ReportResultReference`

NewReportResultReference instantiates a new ReportResultReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportResultReferenceWithDefaults

`func NewReportResultReferenceWithDefaults() *ReportResultReference`

NewReportResultReferenceWithDefaults instantiates a new ReportResultReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ReportResultReference) GetType() DtoType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReportResultReference) GetTypeOk() (*DtoType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReportResultReference) SetType(v DtoType)`

SetType sets Type field to given value.

### HasType

`func (o *ReportResultReference) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *ReportResultReference) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReportResultReference) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReportResultReference) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReportResultReference) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ReportResultReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReportResultReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReportResultReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReportResultReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *ReportResultReference) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReportResultReference) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReportResultReference) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReportResultReference) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


