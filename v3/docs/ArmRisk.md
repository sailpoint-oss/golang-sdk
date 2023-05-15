# ArmRisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Risk Name | [optional] [readonly] 
**Description** | Pointer to **string** | Description | [optional] [readonly] 
**Owners** | Pointer to **[]string** | List of risk owners&#39; SAP usernames | [optional] [readonly] 
**ExternalReference** | Pointer to **string** | URL to \&quot;what if\&quot; details in ARM | [optional] [readonly] 
**Rating** | Pointer to **string** | Risk Rating | [optional] [readonly] 
**BusinessFunctions** | Pointer to **map[string]string** | A map from business function codes to \&quot;Left side\&quot; or \&quot;Right side\&quot; | [optional] [readonly] 
**Approvers** | Pointer to **[]string** | List of risk approvers&#39; SAP usernames | [optional] [readonly] 
**MitigatingControls** | Pointer to **string** | What was done to mitigate risks | [optional] [readonly] 
**CorrectionAdvice** | Pointer to **string** | Recommendation on how to resolve risk | [optional] [readonly] 

## Methods

### NewArmRisk

`func NewArmRisk() *ArmRisk`

NewArmRisk instantiates a new ArmRisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArmRiskWithDefaults

`func NewArmRiskWithDefaults() *ArmRisk`

NewArmRiskWithDefaults instantiates a new ArmRisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ArmRisk) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArmRisk) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArmRisk) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ArmRisk) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ArmRisk) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ArmRisk) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ArmRisk) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ArmRisk) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwners

`func (o *ArmRisk) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *ArmRisk) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *ArmRisk) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *ArmRisk) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetExternalReference

`func (o *ArmRisk) GetExternalReference() string`

GetExternalReference returns the ExternalReference field if non-nil, zero value otherwise.

### GetExternalReferenceOk

`func (o *ArmRisk) GetExternalReferenceOk() (*string, bool)`

GetExternalReferenceOk returns a tuple with the ExternalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReference

`func (o *ArmRisk) SetExternalReference(v string)`

SetExternalReference sets ExternalReference field to given value.

### HasExternalReference

`func (o *ArmRisk) HasExternalReference() bool`

HasExternalReference returns a boolean if a field has been set.

### GetRating

`func (o *ArmRisk) GetRating() string`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *ArmRisk) GetRatingOk() (*string, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *ArmRisk) SetRating(v string)`

SetRating sets Rating field to given value.

### HasRating

`func (o *ArmRisk) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetBusinessFunctions

`func (o *ArmRisk) GetBusinessFunctions() map[string]string`

GetBusinessFunctions returns the BusinessFunctions field if non-nil, zero value otherwise.

### GetBusinessFunctionsOk

`func (o *ArmRisk) GetBusinessFunctionsOk() (*map[string]string, bool)`

GetBusinessFunctionsOk returns a tuple with the BusinessFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessFunctions

`func (o *ArmRisk) SetBusinessFunctions(v map[string]string)`

SetBusinessFunctions sets BusinessFunctions field to given value.

### HasBusinessFunctions

`func (o *ArmRisk) HasBusinessFunctions() bool`

HasBusinessFunctions returns a boolean if a field has been set.

### GetApprovers

`func (o *ArmRisk) GetApprovers() []string`

GetApprovers returns the Approvers field if non-nil, zero value otherwise.

### GetApproversOk

`func (o *ArmRisk) GetApproversOk() (*[]string, bool)`

GetApproversOk returns a tuple with the Approvers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovers

`func (o *ArmRisk) SetApprovers(v []string)`

SetApprovers sets Approvers field to given value.

### HasApprovers

`func (o *ArmRisk) HasApprovers() bool`

HasApprovers returns a boolean if a field has been set.

### GetMitigatingControls

`func (o *ArmRisk) GetMitigatingControls() string`

GetMitigatingControls returns the MitigatingControls field if non-nil, zero value otherwise.

### GetMitigatingControlsOk

`func (o *ArmRisk) GetMitigatingControlsOk() (*string, bool)`

GetMitigatingControlsOk returns a tuple with the MitigatingControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigatingControls

`func (o *ArmRisk) SetMitigatingControls(v string)`

SetMitigatingControls sets MitigatingControls field to given value.

### HasMitigatingControls

`func (o *ArmRisk) HasMitigatingControls() bool`

HasMitigatingControls returns a boolean if a field has been set.

### GetCorrectionAdvice

`func (o *ArmRisk) GetCorrectionAdvice() string`

GetCorrectionAdvice returns the CorrectionAdvice field if non-nil, zero value otherwise.

### GetCorrectionAdviceOk

`func (o *ArmRisk) GetCorrectionAdviceOk() (*string, bool)`

GetCorrectionAdviceOk returns a tuple with the CorrectionAdvice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrectionAdvice

`func (o *ArmRisk) SetCorrectionAdvice(v string)`

SetCorrectionAdvice sets CorrectionAdvice field to given value.

### HasCorrectionAdvice

`func (o *ArmRisk) HasCorrectionAdvice() bool`

HasCorrectionAdvice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


