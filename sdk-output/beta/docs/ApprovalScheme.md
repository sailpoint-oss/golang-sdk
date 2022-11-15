# ApprovalScheme

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApproverType** | Pointer to **string** | Describes the individual or group that is responsible for an approval step. Values are as follows.  **OWNER**: Owner of the associated Access Profile or Role  **SOURCE_OWNER**: Owner of the Source associated with an Access Profile  **MANAGER**: Manager of the Identity making the request  **GOVERNANCE_GROUP**: A Governance Group, the ID of which is specified by the **approverId** field | [optional] 
**ApproverId** | Pointer to **string** | Id of the specific approver, used only when approverType is GOVERNANCE_GROUP | [optional] 

## Methods

### NewApprovalScheme

`func NewApprovalScheme() *ApprovalScheme`

NewApprovalScheme instantiates a new ApprovalScheme object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalSchemeWithDefaults

`func NewApprovalSchemeWithDefaults() *ApprovalScheme`

NewApprovalSchemeWithDefaults instantiates a new ApprovalScheme object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApproverType

`func (o *ApprovalScheme) GetApproverType() string`

GetApproverType returns the ApproverType field if non-nil, zero value otherwise.

### GetApproverTypeOk

`func (o *ApprovalScheme) GetApproverTypeOk() (*string, bool)`

GetApproverTypeOk returns a tuple with the ApproverType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverType

`func (o *ApprovalScheme) SetApproverType(v string)`

SetApproverType sets ApproverType field to given value.

### HasApproverType

`func (o *ApprovalScheme) HasApproverType() bool`

HasApproverType returns a boolean if a field has been set.

### GetApproverId

`func (o *ApprovalScheme) GetApproverId() string`

GetApproverId returns the ApproverId field if non-nil, zero value otherwise.

### GetApproverIdOk

`func (o *ApprovalScheme) GetApproverIdOk() (*string, bool)`

GetApproverIdOk returns a tuple with the ApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverId

`func (o *ApprovalScheme) SetApproverId(v string)`

SetApproverId sets ApproverId field to given value.

### HasApproverId

`func (o *ApprovalScheme) HasApproverId() bool`

HasApproverId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


