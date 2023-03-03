# ViolationOwnerAssignmentConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignmentRule** | Pointer to **string** | Details about the violations owner. MANAGER - identity&#39;s manager STATIC - Governance Group or Identity | [optional] 
**OwnerRef** | Pointer to [**BaseReferenceDto**](BaseReferenceDto.md) |  | [optional] 

## Methods

### NewViolationOwnerAssignmentConfig

`func NewViolationOwnerAssignmentConfig() *ViolationOwnerAssignmentConfig`

NewViolationOwnerAssignmentConfig instantiates a new ViolationOwnerAssignmentConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViolationOwnerAssignmentConfigWithDefaults

`func NewViolationOwnerAssignmentConfigWithDefaults() *ViolationOwnerAssignmentConfig`

NewViolationOwnerAssignmentConfigWithDefaults instantiates a new ViolationOwnerAssignmentConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignmentRule

`func (o *ViolationOwnerAssignmentConfig) GetAssignmentRule() string`

GetAssignmentRule returns the AssignmentRule field if non-nil, zero value otherwise.

### GetAssignmentRuleOk

`func (o *ViolationOwnerAssignmentConfig) GetAssignmentRuleOk() (*string, bool)`

GetAssignmentRuleOk returns a tuple with the AssignmentRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentRule

`func (o *ViolationOwnerAssignmentConfig) SetAssignmentRule(v string)`

SetAssignmentRule sets AssignmentRule field to given value.

### HasAssignmentRule

`func (o *ViolationOwnerAssignmentConfig) HasAssignmentRule() bool`

HasAssignmentRule returns a boolean if a field has been set.

### GetOwnerRef

`func (o *ViolationOwnerAssignmentConfig) GetOwnerRef() BaseReferenceDto`

GetOwnerRef returns the OwnerRef field if non-nil, zero value otherwise.

### GetOwnerRefOk

`func (o *ViolationOwnerAssignmentConfig) GetOwnerRefOk() (*BaseReferenceDto, bool)`

GetOwnerRefOk returns a tuple with the OwnerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerRef

`func (o *ViolationOwnerAssignmentConfig) SetOwnerRef(v BaseReferenceDto)`

SetOwnerRef sets OwnerRef field to given value.

### HasOwnerRef

`func (o *ViolationOwnerAssignmentConfig) HasOwnerRef() bool`

HasOwnerRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


