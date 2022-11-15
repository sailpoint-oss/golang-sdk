# GetEvents200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessItem** | Pointer to [**AccessItemAssociatedAccessItem**](AccessItemAssociatedAccessItem.md) |  | [optional] 
**IdentityId** | Pointer to **string** | the identity id | [optional] 
**EventType** | Pointer to **string** | the event type | [optional] 
**Dt** | Pointer to **string** | the date of event | [optional] 
**GovernanceEvent** | Pointer to [**CorrelatedGovernanceEvent**](CorrelatedGovernanceEvent.md) |  | [optional] 
**Changes** | Pointer to [**[]AttributeChange**](AttributeChange.md) |  | [optional] 
**AccessRequest** | Pointer to [**AccessRequestResponse**](AccessRequestResponse.md) |  | [optional] 
**CertificationId** | Pointer to **string** | the id of the certification item | [optional] 
**CertificationName** | Pointer to **string** | the certification item name | [optional] 
**SignedDate** | Pointer to **string** | the date ceritification was signed | [optional] 
**Certifiers** | Pointer to [**[]CertifierResponse**](CertifierResponse.md) | this field is deprecated and may go away | [optional] 
**Reviewers** | Pointer to [**[]CertifierResponse**](CertifierResponse.md) | The list of identities who review this certification | [optional] 
**Signer** | Pointer to [**CertifierResponse**](CertifierResponse.md) |  | [optional] 
**Account** | Pointer to [**AccountStatusChangedAccount**](AccountStatusChangedAccount.md) |  | [optional] 
**StatusChange** | Pointer to [**AccountStatusChangedStatusChange**](AccountStatusChangedStatusChange.md) |  | [optional] 

## Methods

### NewGetEvents200ResponseInner

`func NewGetEvents200ResponseInner() *GetEvents200ResponseInner`

NewGetEvents200ResponseInner instantiates a new GetEvents200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseInnerWithDefaults

`func NewGetEvents200ResponseInnerWithDefaults() *GetEvents200ResponseInner`

NewGetEvents200ResponseInnerWithDefaults instantiates a new GetEvents200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessItem

`func (o *GetEvents200ResponseInner) GetAccessItem() AccessItemAssociatedAccessItem`

GetAccessItem returns the AccessItem field if non-nil, zero value otherwise.

### GetAccessItemOk

`func (o *GetEvents200ResponseInner) GetAccessItemOk() (*AccessItemAssociatedAccessItem, bool)`

GetAccessItemOk returns a tuple with the AccessItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessItem

`func (o *GetEvents200ResponseInner) SetAccessItem(v AccessItemAssociatedAccessItem)`

SetAccessItem sets AccessItem field to given value.

### HasAccessItem

`func (o *GetEvents200ResponseInner) HasAccessItem() bool`

HasAccessItem returns a boolean if a field has been set.

### GetIdentityId

`func (o *GetEvents200ResponseInner) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *GetEvents200ResponseInner) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *GetEvents200ResponseInner) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *GetEvents200ResponseInner) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetEventType

`func (o *GetEvents200ResponseInner) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *GetEvents200ResponseInner) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *GetEvents200ResponseInner) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *GetEvents200ResponseInner) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetDt

`func (o *GetEvents200ResponseInner) GetDt() string`

GetDt returns the Dt field if non-nil, zero value otherwise.

### GetDtOk

`func (o *GetEvents200ResponseInner) GetDtOk() (*string, bool)`

GetDtOk returns a tuple with the Dt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDt

`func (o *GetEvents200ResponseInner) SetDt(v string)`

SetDt sets Dt field to given value.

### HasDt

`func (o *GetEvents200ResponseInner) HasDt() bool`

HasDt returns a boolean if a field has been set.

### GetGovernanceEvent

`func (o *GetEvents200ResponseInner) GetGovernanceEvent() CorrelatedGovernanceEvent`

GetGovernanceEvent returns the GovernanceEvent field if non-nil, zero value otherwise.

### GetGovernanceEventOk

`func (o *GetEvents200ResponseInner) GetGovernanceEventOk() (*CorrelatedGovernanceEvent, bool)`

GetGovernanceEventOk returns a tuple with the GovernanceEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovernanceEvent

`func (o *GetEvents200ResponseInner) SetGovernanceEvent(v CorrelatedGovernanceEvent)`

SetGovernanceEvent sets GovernanceEvent field to given value.

### HasGovernanceEvent

`func (o *GetEvents200ResponseInner) HasGovernanceEvent() bool`

HasGovernanceEvent returns a boolean if a field has been set.

### GetChanges

`func (o *GetEvents200ResponseInner) GetChanges() []AttributeChange`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *GetEvents200ResponseInner) GetChangesOk() (*[]AttributeChange, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *GetEvents200ResponseInner) SetChanges(v []AttributeChange)`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *GetEvents200ResponseInner) HasChanges() bool`

HasChanges returns a boolean if a field has been set.

### GetAccessRequest

`func (o *GetEvents200ResponseInner) GetAccessRequest() AccessRequestResponse`

GetAccessRequest returns the AccessRequest field if non-nil, zero value otherwise.

### GetAccessRequestOk

`func (o *GetEvents200ResponseInner) GetAccessRequestOk() (*AccessRequestResponse, bool)`

GetAccessRequestOk returns a tuple with the AccessRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRequest

`func (o *GetEvents200ResponseInner) SetAccessRequest(v AccessRequestResponse)`

SetAccessRequest sets AccessRequest field to given value.

### HasAccessRequest

`func (o *GetEvents200ResponseInner) HasAccessRequest() bool`

HasAccessRequest returns a boolean if a field has been set.

### GetCertificationId

`func (o *GetEvents200ResponseInner) GetCertificationId() string`

GetCertificationId returns the CertificationId field if non-nil, zero value otherwise.

### GetCertificationIdOk

`func (o *GetEvents200ResponseInner) GetCertificationIdOk() (*string, bool)`

GetCertificationIdOk returns a tuple with the CertificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificationId

`func (o *GetEvents200ResponseInner) SetCertificationId(v string)`

SetCertificationId sets CertificationId field to given value.

### HasCertificationId

`func (o *GetEvents200ResponseInner) HasCertificationId() bool`

HasCertificationId returns a boolean if a field has been set.

### GetCertificationName

`func (o *GetEvents200ResponseInner) GetCertificationName() string`

GetCertificationName returns the CertificationName field if non-nil, zero value otherwise.

### GetCertificationNameOk

`func (o *GetEvents200ResponseInner) GetCertificationNameOk() (*string, bool)`

GetCertificationNameOk returns a tuple with the CertificationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificationName

`func (o *GetEvents200ResponseInner) SetCertificationName(v string)`

SetCertificationName sets CertificationName field to given value.

### HasCertificationName

`func (o *GetEvents200ResponseInner) HasCertificationName() bool`

HasCertificationName returns a boolean if a field has been set.

### GetSignedDate

`func (o *GetEvents200ResponseInner) GetSignedDate() string`

GetSignedDate returns the SignedDate field if non-nil, zero value otherwise.

### GetSignedDateOk

`func (o *GetEvents200ResponseInner) GetSignedDateOk() (*string, bool)`

GetSignedDateOk returns a tuple with the SignedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedDate

`func (o *GetEvents200ResponseInner) SetSignedDate(v string)`

SetSignedDate sets SignedDate field to given value.

### HasSignedDate

`func (o *GetEvents200ResponseInner) HasSignedDate() bool`

HasSignedDate returns a boolean if a field has been set.

### GetCertifiers

`func (o *GetEvents200ResponseInner) GetCertifiers() []CertifierResponse`

GetCertifiers returns the Certifiers field if non-nil, zero value otherwise.

### GetCertifiersOk

`func (o *GetEvents200ResponseInner) GetCertifiersOk() (*[]CertifierResponse, bool)`

GetCertifiersOk returns a tuple with the Certifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertifiers

`func (o *GetEvents200ResponseInner) SetCertifiers(v []CertifierResponse)`

SetCertifiers sets Certifiers field to given value.

### HasCertifiers

`func (o *GetEvents200ResponseInner) HasCertifiers() bool`

HasCertifiers returns a boolean if a field has been set.

### GetReviewers

`func (o *GetEvents200ResponseInner) GetReviewers() []CertifierResponse`

GetReviewers returns the Reviewers field if non-nil, zero value otherwise.

### GetReviewersOk

`func (o *GetEvents200ResponseInner) GetReviewersOk() (*[]CertifierResponse, bool)`

GetReviewersOk returns a tuple with the Reviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewers

`func (o *GetEvents200ResponseInner) SetReviewers(v []CertifierResponse)`

SetReviewers sets Reviewers field to given value.

### HasReviewers

`func (o *GetEvents200ResponseInner) HasReviewers() bool`

HasReviewers returns a boolean if a field has been set.

### GetSigner

`func (o *GetEvents200ResponseInner) GetSigner() CertifierResponse`

GetSigner returns the Signer field if non-nil, zero value otherwise.

### GetSignerOk

`func (o *GetEvents200ResponseInner) GetSignerOk() (*CertifierResponse, bool)`

GetSignerOk returns a tuple with the Signer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigner

`func (o *GetEvents200ResponseInner) SetSigner(v CertifierResponse)`

SetSigner sets Signer field to given value.

### HasSigner

`func (o *GetEvents200ResponseInner) HasSigner() bool`

HasSigner returns a boolean if a field has been set.

### GetAccount

`func (o *GetEvents200ResponseInner) GetAccount() AccountStatusChangedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetEvents200ResponseInner) GetAccountOk() (*AccountStatusChangedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetEvents200ResponseInner) SetAccount(v AccountStatusChangedAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetEvents200ResponseInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetStatusChange

`func (o *GetEvents200ResponseInner) GetStatusChange() AccountStatusChangedStatusChange`

GetStatusChange returns the StatusChange field if non-nil, zero value otherwise.

### GetStatusChangeOk

`func (o *GetEvents200ResponseInner) GetStatusChangeOk() (*AccountStatusChangedStatusChange, bool)`

GetStatusChangeOk returns a tuple with the StatusChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusChange

`func (o *GetEvents200ResponseInner) SetStatusChange(v AccountStatusChangedStatusChange)`

SetStatusChange sets StatusChange field to given value.

### HasStatusChange

`func (o *GetEvents200ResponseInner) HasStatusChange() bool`

HasStatusChange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


