# VAClusterStatusChangeEventApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The GUID of the application | 
**Name** | **string** | The name of the application | 
**Attributes** | **map[string]interface{}** | Custom map of attributes for a source.  This will only be populated if type is &#x60;SOURCE&#x60; and the source has a proxy. | 

## Methods

### NewVAClusterStatusChangeEventApplication

`func NewVAClusterStatusChangeEventApplication(id string, name string, attributes map[string]interface{}, ) *VAClusterStatusChangeEventApplication`

NewVAClusterStatusChangeEventApplication instantiates a new VAClusterStatusChangeEventApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVAClusterStatusChangeEventApplicationWithDefaults

`func NewVAClusterStatusChangeEventApplicationWithDefaults() *VAClusterStatusChangeEventApplication`

NewVAClusterStatusChangeEventApplicationWithDefaults instantiates a new VAClusterStatusChangeEventApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VAClusterStatusChangeEventApplication) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VAClusterStatusChangeEventApplication) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VAClusterStatusChangeEventApplication) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *VAClusterStatusChangeEventApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VAClusterStatusChangeEventApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VAClusterStatusChangeEventApplication) SetName(v string)`

SetName sets Name field to given value.


### GetAttributes

`func (o *VAClusterStatusChangeEventApplication) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *VAClusterStatusChangeEventApplication) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *VAClusterStatusChangeEventApplication) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.


### SetAttributesNil

`func (o *VAClusterStatusChangeEventApplication) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *VAClusterStatusChangeEventApplication) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


