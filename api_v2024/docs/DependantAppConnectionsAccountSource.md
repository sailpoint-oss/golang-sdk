# DependantAppConnectionsAccountSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseForPasswordManagement** | Pointer to **bool** | Use this Account Source for password management | [optional] [default to false]
**PasswordPolicies** | Pointer to [**[]DependantAppConnectionsAccountSourcePasswordPoliciesInner**](DependantAppConnectionsAccountSourcePasswordPoliciesInner.md) | A list of Password Policies for this Account Source | [optional] 

## Methods

### NewDependantAppConnectionsAccountSource

`func NewDependantAppConnectionsAccountSource() *DependantAppConnectionsAccountSource`

NewDependantAppConnectionsAccountSource instantiates a new DependantAppConnectionsAccountSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDependantAppConnectionsAccountSourceWithDefaults

`func NewDependantAppConnectionsAccountSourceWithDefaults() *DependantAppConnectionsAccountSource`

NewDependantAppConnectionsAccountSourceWithDefaults instantiates a new DependantAppConnectionsAccountSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUseForPasswordManagement

`func (o *DependantAppConnectionsAccountSource) GetUseForPasswordManagement() bool`

GetUseForPasswordManagement returns the UseForPasswordManagement field if non-nil, zero value otherwise.

### GetUseForPasswordManagementOk

`func (o *DependantAppConnectionsAccountSource) GetUseForPasswordManagementOk() (*bool, bool)`

GetUseForPasswordManagementOk returns a tuple with the UseForPasswordManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseForPasswordManagement

`func (o *DependantAppConnectionsAccountSource) SetUseForPasswordManagement(v bool)`

SetUseForPasswordManagement sets UseForPasswordManagement field to given value.

### HasUseForPasswordManagement

`func (o *DependantAppConnectionsAccountSource) HasUseForPasswordManagement() bool`

HasUseForPasswordManagement returns a boolean if a field has been set.

### GetPasswordPolicies

`func (o *DependantAppConnectionsAccountSource) GetPasswordPolicies() []DependantAppConnectionsAccountSourcePasswordPoliciesInner`

GetPasswordPolicies returns the PasswordPolicies field if non-nil, zero value otherwise.

### GetPasswordPoliciesOk

`func (o *DependantAppConnectionsAccountSource) GetPasswordPoliciesOk() (*[]DependantAppConnectionsAccountSourcePasswordPoliciesInner, bool)`

GetPasswordPoliciesOk returns a tuple with the PasswordPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicies

`func (o *DependantAppConnectionsAccountSource) SetPasswordPolicies(v []DependantAppConnectionsAccountSourcePasswordPoliciesInner)`

SetPasswordPolicies sets PasswordPolicies field to given value.

### HasPasswordPolicies

`func (o *DependantAppConnectionsAccountSource) HasPasswordPolicies() bool`

HasPasswordPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


