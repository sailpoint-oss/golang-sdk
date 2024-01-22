# ProvisioningPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | the provisioning policy name | 
**Description** | Pointer to **string** | the description of the provisioning policy | [optional] 
**UsageType** | Pointer to [**UsageType**](UsageType.md) |  | [optional] 
**Fields** | Pointer to [**[]FieldDetailsDto**](FieldDetailsDto.md) |  | [optional] 

## Methods

### NewProvisioningPolicy

`func NewProvisioningPolicy(name string, ) *ProvisioningPolicy`

NewProvisioningPolicy instantiates a new ProvisioningPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningPolicyWithDefaults

`func NewProvisioningPolicyWithDefaults() *ProvisioningPolicy`

NewProvisioningPolicyWithDefaults instantiates a new ProvisioningPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProvisioningPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProvisioningPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProvisioningPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ProvisioningPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProvisioningPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProvisioningPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProvisioningPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUsageType

`func (o *ProvisioningPolicy) GetUsageType() UsageType`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *ProvisioningPolicy) GetUsageTypeOk() (*UsageType, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *ProvisioningPolicy) SetUsageType(v UsageType)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *ProvisioningPolicy) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.

### GetFields

`func (o *ProvisioningPolicy) GetFields() []FieldDetailsDto`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ProvisioningPolicy) GetFieldsOk() (*[]FieldDetailsDto, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ProvisioningPolicy) SetFields(v []FieldDetailsDto)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ProvisioningPolicy) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


