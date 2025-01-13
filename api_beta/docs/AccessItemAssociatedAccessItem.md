# AccessItemAssociatedAccessItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessType** | Pointer to **string** | the access item type. role in this case | [optional] 
**Id** | Pointer to **string** | the access item id | [optional] 
**Name** | Pointer to **string** | the access profile name | [optional] 
**SourceName** | Pointer to **string** | the associated source name if it exists | [optional] 
**SourceId** | Pointer to **string** | the id of the source | [optional] 
**Description** | Pointer to **string** | the description for the role | [optional] 
**DisplayName** | Pointer to **string** | the role display name | [optional] 
**EntitlementCount** | Pointer to **string** | the number of entitlements the account will create | [optional] 
**AppDisplayName** | Pointer to **string** | the name of app | [optional] 
**NativeIdentity** | Pointer to **string** | the native identifier used to uniquely identify an acccount | [optional] 
**Attribute** | Pointer to **string** | the entitlement attribute | [optional] 
**Value** | Pointer to **string** | the associated value | [optional] 
**EntitlementType** | Pointer to **string** | the type of entitlement | [optional] 

## Methods

### NewAccessItemAssociatedAccessItem

`func NewAccessItemAssociatedAccessItem() *AccessItemAssociatedAccessItem`

NewAccessItemAssociatedAccessItem instantiates a new AccessItemAssociatedAccessItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessItemAssociatedAccessItemWithDefaults

`func NewAccessItemAssociatedAccessItemWithDefaults() *AccessItemAssociatedAccessItem`

NewAccessItemAssociatedAccessItemWithDefaults instantiates a new AccessItemAssociatedAccessItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessType

`func (o *AccessItemAssociatedAccessItem) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *AccessItemAssociatedAccessItem) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *AccessItemAssociatedAccessItem) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *AccessItemAssociatedAccessItem) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetId

`func (o *AccessItemAssociatedAccessItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessItemAssociatedAccessItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessItemAssociatedAccessItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccessItemAssociatedAccessItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AccessItemAssociatedAccessItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessItemAssociatedAccessItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessItemAssociatedAccessItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccessItemAssociatedAccessItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSourceName

`func (o *AccessItemAssociatedAccessItem) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *AccessItemAssociatedAccessItem) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *AccessItemAssociatedAccessItem) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *AccessItemAssociatedAccessItem) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### GetSourceId

`func (o *AccessItemAssociatedAccessItem) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *AccessItemAssociatedAccessItem) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *AccessItemAssociatedAccessItem) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *AccessItemAssociatedAccessItem) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetDescription

`func (o *AccessItemAssociatedAccessItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccessItemAssociatedAccessItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccessItemAssociatedAccessItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccessItemAssociatedAccessItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *AccessItemAssociatedAccessItem) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AccessItemAssociatedAccessItem) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AccessItemAssociatedAccessItem) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AccessItemAssociatedAccessItem) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEntitlementCount

`func (o *AccessItemAssociatedAccessItem) GetEntitlementCount() string`

GetEntitlementCount returns the EntitlementCount field if non-nil, zero value otherwise.

### GetEntitlementCountOk

`func (o *AccessItemAssociatedAccessItem) GetEntitlementCountOk() (*string, bool)`

GetEntitlementCountOk returns a tuple with the EntitlementCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementCount

`func (o *AccessItemAssociatedAccessItem) SetEntitlementCount(v string)`

SetEntitlementCount sets EntitlementCount field to given value.

### HasEntitlementCount

`func (o *AccessItemAssociatedAccessItem) HasEntitlementCount() bool`

HasEntitlementCount returns a boolean if a field has been set.

### GetAppDisplayName

`func (o *AccessItemAssociatedAccessItem) GetAppDisplayName() string`

GetAppDisplayName returns the AppDisplayName field if non-nil, zero value otherwise.

### GetAppDisplayNameOk

`func (o *AccessItemAssociatedAccessItem) GetAppDisplayNameOk() (*string, bool)`

GetAppDisplayNameOk returns a tuple with the AppDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDisplayName

`func (o *AccessItemAssociatedAccessItem) SetAppDisplayName(v string)`

SetAppDisplayName sets AppDisplayName field to given value.

### HasAppDisplayName

`func (o *AccessItemAssociatedAccessItem) HasAppDisplayName() bool`

HasAppDisplayName returns a boolean if a field has been set.

### GetNativeIdentity

`func (o *AccessItemAssociatedAccessItem) GetNativeIdentity() string`

GetNativeIdentity returns the NativeIdentity field if non-nil, zero value otherwise.

### GetNativeIdentityOk

`func (o *AccessItemAssociatedAccessItem) GetNativeIdentityOk() (*string, bool)`

GetNativeIdentityOk returns a tuple with the NativeIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeIdentity

`func (o *AccessItemAssociatedAccessItem) SetNativeIdentity(v string)`

SetNativeIdentity sets NativeIdentity field to given value.

### HasNativeIdentity

`func (o *AccessItemAssociatedAccessItem) HasNativeIdentity() bool`

HasNativeIdentity returns a boolean if a field has been set.

### GetAttribute

`func (o *AccessItemAssociatedAccessItem) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *AccessItemAssociatedAccessItem) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *AccessItemAssociatedAccessItem) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *AccessItemAssociatedAccessItem) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetValue

`func (o *AccessItemAssociatedAccessItem) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AccessItemAssociatedAccessItem) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AccessItemAssociatedAccessItem) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *AccessItemAssociatedAccessItem) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetEntitlementType

`func (o *AccessItemAssociatedAccessItem) GetEntitlementType() string`

GetEntitlementType returns the EntitlementType field if non-nil, zero value otherwise.

### GetEntitlementTypeOk

`func (o *AccessItemAssociatedAccessItem) GetEntitlementTypeOk() (*string, bool)`

GetEntitlementTypeOk returns a tuple with the EntitlementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementType

`func (o *AccessItemAssociatedAccessItem) SetEntitlementType(v string)`

SetEntitlementType sets EntitlementType field to given value.

### HasEntitlementType

`func (o *AccessItemAssociatedAccessItem) HasEntitlementType() bool`

HasEntitlementType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


