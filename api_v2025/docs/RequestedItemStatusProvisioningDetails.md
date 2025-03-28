# RequestedItemStatusProvisioningDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderedSubPhaseReferences** | Pointer to **string** | Ordered CSV of sub phase references to objects that contain more information about provisioning. For example, this can contain \&quot;manualWorkItemDetails\&quot; which indicate that there is further information in that object for this phase. | [optional] 

## Methods

### NewRequestedItemStatusProvisioningDetails

`func NewRequestedItemStatusProvisioningDetails() *RequestedItemStatusProvisioningDetails`

NewRequestedItemStatusProvisioningDetails instantiates a new RequestedItemStatusProvisioningDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestedItemStatusProvisioningDetailsWithDefaults

`func NewRequestedItemStatusProvisioningDetailsWithDefaults() *RequestedItemStatusProvisioningDetails`

NewRequestedItemStatusProvisioningDetailsWithDefaults instantiates a new RequestedItemStatusProvisioningDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderedSubPhaseReferences

`func (o *RequestedItemStatusProvisioningDetails) GetOrderedSubPhaseReferences() string`

GetOrderedSubPhaseReferences returns the OrderedSubPhaseReferences field if non-nil, zero value otherwise.

### GetOrderedSubPhaseReferencesOk

`func (o *RequestedItemStatusProvisioningDetails) GetOrderedSubPhaseReferencesOk() (*string, bool)`

GetOrderedSubPhaseReferencesOk returns a tuple with the OrderedSubPhaseReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderedSubPhaseReferences

`func (o *RequestedItemStatusProvisioningDetails) SetOrderedSubPhaseReferences(v string)`

SetOrderedSubPhaseReferences sets OrderedSubPhaseReferences field to given value.

### HasOrderedSubPhaseReferences

`func (o *RequestedItemStatusProvisioningDetails) HasOrderedSubPhaseReferences() bool`

HasOrderedSubPhaseReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


