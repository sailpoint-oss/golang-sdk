# CreatePersonalAccessTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the personal access token (PAT) to be created. Cannot be the same as another PAT owned by the user for whom this PAT is being created. | 
**Scope** | Pointer to **[]string** | Scopes of the personal  access token. If no scope is specified, the token will be created with the default scope \&quot;sp:scopes:all\&quot;. This means the personal access token will have all the rights of the owner who created it. | [optional] 

## Methods

### NewCreatePersonalAccessTokenRequest

`func NewCreatePersonalAccessTokenRequest(name string, ) *CreatePersonalAccessTokenRequest`

NewCreatePersonalAccessTokenRequest instantiates a new CreatePersonalAccessTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePersonalAccessTokenRequestWithDefaults

`func NewCreatePersonalAccessTokenRequestWithDefaults() *CreatePersonalAccessTokenRequest`

NewCreatePersonalAccessTokenRequestWithDefaults instantiates a new CreatePersonalAccessTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreatePersonalAccessTokenRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePersonalAccessTokenRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePersonalAccessTokenRequest) SetName(v string)`

SetName sets Name field to given value.


### GetScope

`func (o *CreatePersonalAccessTokenRequest) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreatePersonalAccessTokenRequest) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreatePersonalAccessTokenRequest) SetScope(v []string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CreatePersonalAccessTokenRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *CreatePersonalAccessTokenRequest) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *CreatePersonalAccessTokenRequest) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


