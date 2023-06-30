# RefreshIdentitiesRequestRefreshArgs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorrelateEntitlements** | Pointer to **bool** | This will refresh entitlement, role, and access profile calculations. | [optional] 
**PromoteAttributes** | Pointer to **bool** | This will calculate identity attributes. | [optional] 
**RefreshManagerStatus** | Pointer to **bool** | This recalculates manager correlation and manager status. Note: This is computationally expensive to run.  | [optional] 
**SynchronizeAttributes** | Pointer to **bool** | Enables attribute synchronization. | [optional] 
**PruneIdentities** | Pointer to **bool** | Option that will enable deletion of an identity objects if there are no account objects. Note: This is not typically used in IdentityNow, except by guidance from SailPoint.  | [optional] 
**Provision** | Pointer to **bool** | Enables provisioning of role assignments with entitlements that are not currently fulfilled. | [optional] 

## Methods

### NewRefreshIdentitiesRequestRefreshArgs

`func NewRefreshIdentitiesRequestRefreshArgs() *RefreshIdentitiesRequestRefreshArgs`

NewRefreshIdentitiesRequestRefreshArgs instantiates a new RefreshIdentitiesRequestRefreshArgs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefreshIdentitiesRequestRefreshArgsWithDefaults

`func NewRefreshIdentitiesRequestRefreshArgsWithDefaults() *RefreshIdentitiesRequestRefreshArgs`

NewRefreshIdentitiesRequestRefreshArgsWithDefaults instantiates a new RefreshIdentitiesRequestRefreshArgs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorrelateEntitlements

`func (o *RefreshIdentitiesRequestRefreshArgs) GetCorrelateEntitlements() bool`

GetCorrelateEntitlements returns the CorrelateEntitlements field if non-nil, zero value otherwise.

### GetCorrelateEntitlementsOk

`func (o *RefreshIdentitiesRequestRefreshArgs) GetCorrelateEntitlementsOk() (*bool, bool)`

GetCorrelateEntitlementsOk returns a tuple with the CorrelateEntitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelateEntitlements

`func (o *RefreshIdentitiesRequestRefreshArgs) SetCorrelateEntitlements(v bool)`

SetCorrelateEntitlements sets CorrelateEntitlements field to given value.

### HasCorrelateEntitlements

`func (o *RefreshIdentitiesRequestRefreshArgs) HasCorrelateEntitlements() bool`

HasCorrelateEntitlements returns a boolean if a field has been set.

### GetPromoteAttributes

`func (o *RefreshIdentitiesRequestRefreshArgs) GetPromoteAttributes() bool`

GetPromoteAttributes returns the PromoteAttributes field if non-nil, zero value otherwise.

### GetPromoteAttributesOk

`func (o *RefreshIdentitiesRequestRefreshArgs) GetPromoteAttributesOk() (*bool, bool)`

GetPromoteAttributesOk returns a tuple with the PromoteAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoteAttributes

`func (o *RefreshIdentitiesRequestRefreshArgs) SetPromoteAttributes(v bool)`

SetPromoteAttributes sets PromoteAttributes field to given value.

### HasPromoteAttributes

`func (o *RefreshIdentitiesRequestRefreshArgs) HasPromoteAttributes() bool`

HasPromoteAttributes returns a boolean if a field has been set.

### GetRefreshManagerStatus

`func (o *RefreshIdentitiesRequestRefreshArgs) GetRefreshManagerStatus() bool`

GetRefreshManagerStatus returns the RefreshManagerStatus field if non-nil, zero value otherwise.

### GetRefreshManagerStatusOk

`func (o *RefreshIdentitiesRequestRefreshArgs) GetRefreshManagerStatusOk() (*bool, bool)`

GetRefreshManagerStatusOk returns a tuple with the RefreshManagerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshManagerStatus

`func (o *RefreshIdentitiesRequestRefreshArgs) SetRefreshManagerStatus(v bool)`

SetRefreshManagerStatus sets RefreshManagerStatus field to given value.

### HasRefreshManagerStatus

`func (o *RefreshIdentitiesRequestRefreshArgs) HasRefreshManagerStatus() bool`

HasRefreshManagerStatus returns a boolean if a field has been set.

### GetSynchronizeAttributes

`func (o *RefreshIdentitiesRequestRefreshArgs) GetSynchronizeAttributes() bool`

GetSynchronizeAttributes returns the SynchronizeAttributes field if non-nil, zero value otherwise.

### GetSynchronizeAttributesOk

`func (o *RefreshIdentitiesRequestRefreshArgs) GetSynchronizeAttributesOk() (*bool, bool)`

GetSynchronizeAttributesOk returns a tuple with the SynchronizeAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronizeAttributes

`func (o *RefreshIdentitiesRequestRefreshArgs) SetSynchronizeAttributes(v bool)`

SetSynchronizeAttributes sets SynchronizeAttributes field to given value.

### HasSynchronizeAttributes

`func (o *RefreshIdentitiesRequestRefreshArgs) HasSynchronizeAttributes() bool`

HasSynchronizeAttributes returns a boolean if a field has been set.

### GetPruneIdentities

`func (o *RefreshIdentitiesRequestRefreshArgs) GetPruneIdentities() bool`

GetPruneIdentities returns the PruneIdentities field if non-nil, zero value otherwise.

### GetPruneIdentitiesOk

`func (o *RefreshIdentitiesRequestRefreshArgs) GetPruneIdentitiesOk() (*bool, bool)`

GetPruneIdentitiesOk returns a tuple with the PruneIdentities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPruneIdentities

`func (o *RefreshIdentitiesRequestRefreshArgs) SetPruneIdentities(v bool)`

SetPruneIdentities sets PruneIdentities field to given value.

### HasPruneIdentities

`func (o *RefreshIdentitiesRequestRefreshArgs) HasPruneIdentities() bool`

HasPruneIdentities returns a boolean if a field has been set.

### GetProvision

`func (o *RefreshIdentitiesRequestRefreshArgs) GetProvision() bool`

GetProvision returns the Provision field if non-nil, zero value otherwise.

### GetProvisionOk

`func (o *RefreshIdentitiesRequestRefreshArgs) GetProvisionOk() (*bool, bool)`

GetProvisionOk returns a tuple with the Provision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvision

`func (o *RefreshIdentitiesRequestRefreshArgs) SetProvision(v bool)`

SetProvision sets Provision field to given value.

### HasProvision

`func (o *RefreshIdentitiesRequestRefreshArgs) HasProvision() bool`

HasProvision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


