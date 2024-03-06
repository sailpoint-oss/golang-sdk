# ProvisioningCriteriaLevel3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | Pointer to [**ProvisioningCriteriaOperation**](ProvisioningCriteriaOperation.md) |  | [optional] 
**Attribute** | Pointer to **NullableString** | Name of the Account attribute to be tested. If **operation** is one of EQUALS, NOT_EQUALS, CONTAINS, or HAS, this field is required. Otherwise, specifying it is an error. | [optional] 
**Value** | Pointer to **string** | String value to test the Account attribute w/r/t the specified operation. If the operation is one of EQUALS, NOT_EQUALS, or CONTAINS, this field is required. Otherwise, specifying it is an error. If the Attribute is not String-typed, it will be converted to the appropriate type. | [optional] 

## Methods

### NewProvisioningCriteriaLevel3

`func NewProvisioningCriteriaLevel3() *ProvisioningCriteriaLevel3`

NewProvisioningCriteriaLevel3 instantiates a new ProvisioningCriteriaLevel3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningCriteriaLevel3WithDefaults

`func NewProvisioningCriteriaLevel3WithDefaults() *ProvisioningCriteriaLevel3`

NewProvisioningCriteriaLevel3WithDefaults instantiates a new ProvisioningCriteriaLevel3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *ProvisioningCriteriaLevel3) GetOperation() ProvisioningCriteriaOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ProvisioningCriteriaLevel3) GetOperationOk() (*ProvisioningCriteriaOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ProvisioningCriteriaLevel3) SetOperation(v ProvisioningCriteriaOperation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *ProvisioningCriteriaLevel3) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetAttribute

`func (o *ProvisioningCriteriaLevel3) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *ProvisioningCriteriaLevel3) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *ProvisioningCriteriaLevel3) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *ProvisioningCriteriaLevel3) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### SetAttributeNil

`func (o *ProvisioningCriteriaLevel3) SetAttributeNil(b bool)`

 SetAttributeNil sets the value for Attribute to be an explicit nil

### UnsetAttribute
`func (o *ProvisioningCriteriaLevel3) UnsetAttribute()`

UnsetAttribute ensures that no value is present for Attribute, not even an explicit nil
### GetValue

`func (o *ProvisioningCriteriaLevel3) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProvisioningCriteriaLevel3) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProvisioningCriteriaLevel3) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ProvisioningCriteriaLevel3) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


