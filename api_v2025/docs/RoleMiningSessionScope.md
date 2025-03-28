# RoleMiningSessionScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityIds** | Pointer to **[]string** | The list of identities for this role mining session. | [optional] 
**Criteria** | Pointer to **NullableString** | The \&quot;search\&quot; criteria that produces the list of identities for this role mining session. | [optional] 
**AttributeFilterCriteria** | Pointer to **[]map[string]interface{}** | The filter criteria for this role mining session. | [optional] 

## Methods

### NewRoleMiningSessionScope

`func NewRoleMiningSessionScope() *RoleMiningSessionScope`

NewRoleMiningSessionScope instantiates a new RoleMiningSessionScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleMiningSessionScopeWithDefaults

`func NewRoleMiningSessionScopeWithDefaults() *RoleMiningSessionScope`

NewRoleMiningSessionScopeWithDefaults instantiates a new RoleMiningSessionScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityIds

`func (o *RoleMiningSessionScope) GetIdentityIds() []string`

GetIdentityIds returns the IdentityIds field if non-nil, zero value otherwise.

### GetIdentityIdsOk

`func (o *RoleMiningSessionScope) GetIdentityIdsOk() (*[]string, bool)`

GetIdentityIdsOk returns a tuple with the IdentityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityIds

`func (o *RoleMiningSessionScope) SetIdentityIds(v []string)`

SetIdentityIds sets IdentityIds field to given value.

### HasIdentityIds

`func (o *RoleMiningSessionScope) HasIdentityIds() bool`

HasIdentityIds returns a boolean if a field has been set.

### GetCriteria

`func (o *RoleMiningSessionScope) GetCriteria() string`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *RoleMiningSessionScope) GetCriteriaOk() (*string, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *RoleMiningSessionScope) SetCriteria(v string)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *RoleMiningSessionScope) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### SetCriteriaNil

`func (o *RoleMiningSessionScope) SetCriteriaNil(b bool)`

 SetCriteriaNil sets the value for Criteria to be an explicit nil

### UnsetCriteria
`func (o *RoleMiningSessionScope) UnsetCriteria()`

UnsetCriteria ensures that no value is present for Criteria, not even an explicit nil
### GetAttributeFilterCriteria

`func (o *RoleMiningSessionScope) GetAttributeFilterCriteria() []map[string]interface{}`

GetAttributeFilterCriteria returns the AttributeFilterCriteria field if non-nil, zero value otherwise.

### GetAttributeFilterCriteriaOk

`func (o *RoleMiningSessionScope) GetAttributeFilterCriteriaOk() (*[]map[string]interface{}, bool)`

GetAttributeFilterCriteriaOk returns a tuple with the AttributeFilterCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeFilterCriteria

`func (o *RoleMiningSessionScope) SetAttributeFilterCriteria(v []map[string]interface{})`

SetAttributeFilterCriteria sets AttributeFilterCriteria field to given value.

### HasAttributeFilterCriteria

`func (o *RoleMiningSessionScope) HasAttributeFilterCriteria() bool`

HasAttributeFilterCriteria returns a boolean if a field has been set.

### SetAttributeFilterCriteriaNil

`func (o *RoleMiningSessionScope) SetAttributeFilterCriteriaNil(b bool)`

 SetAttributeFilterCriteriaNil sets the value for AttributeFilterCriteria to be an explicit nil

### UnsetAttributeFilterCriteria
`func (o *RoleMiningSessionScope) UnsetAttributeFilterCriteria()`

UnsetAttributeFilterCriteria ensures that no value is present for AttributeFilterCriteria, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


