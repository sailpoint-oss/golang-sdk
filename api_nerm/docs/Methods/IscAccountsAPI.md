---
id: nerm-isc-accounts
title: IscAccounts
pagination_label: IscAccounts
sidebar_label: IscAccounts
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'IscAccounts', 'NERMIscAccounts'] 
slug: /tools/sdk/go/nerm/methods/isc-accounts
tags: ['SDK', 'Software Development Kit', 'IscAccounts', 'NERMIscAccounts']
---

# IscAccountsAPI
   
All URIs are relative to *https://acmeco.nonemployee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get-schema-mapped-profiles-collection**](#get-schema-mapped-profiles-collection) | **Get** `/isc/accounts` | Get Profiles
[**get-single-schema-mapped-profile**](#get-single-schema-mapped-profile) | **Get** `/isc/accounts/{id}` | Get Profile
[**update-profile**](#update-profile) | **Patch** `/isc/accounts/{id}` | Update Profile


## get-schema-mapped-profiles-collection
Get Profiles
Retrieve schema-mapped profiles collection. It returns a collection of stored profiles, optionally using schema-mapped field names if requested.

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-schema-mapped-profiles-collection)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSchemaMappedProfilesCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of items to return. | 
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | 
 **useSchema** | **bool** | The use_schema parameter returns schema-mapped field names in the Profiles API response when set to true. | [default to false]
 **overrideSyncToggle** | **bool** | The override_sync_toggle parameter returns all non-employee and assignment profiles regardless of the sync status when set to true. | [default to false]
 **category** | **string** | Filters profiles by profile type category | 
 **status** | **string** | The status of the profile. It can be multiple values with comma separated. | 

### Return type

[**GetSchemaMappedProfilesCollection200Response**](../models/get-schema-mapped-profiles-collection200-response)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    limit := 5 // int32 | The maximum number of items to return. (optional) # int32 | The maximum number of items to return. (optional)
    offset := 5 // int32 | The number of items to skip before starting to collect the result set. (optional) # int32 | The number of items to skip before starting to collect the result set. (optional)
    useSchema := true // bool | The use_schema parameter returns schema-mapped field names in the Profiles API response when set to true. (optional) (default to false) # bool | The use_schema parameter returns schema-mapped field names in the Profiles API response when set to true. (optional) (default to false)
    overrideSyncToggle := true // bool | The override_sync_toggle parameter returns all non-employee and assignment profiles regardless of the sync status when set to true. (optional) (default to false) # bool | The override_sync_toggle parameter returns all non-employee and assignment profiles regardless of the sync status when set to true. (optional) (default to false)
    category := `non-employee` // string | Filters profiles by profile type category (optional) # string | Filters profiles by profile type category (optional)
    status := `active,on leave` // string | The status of the profile. It can be multiple values with comma separated. (optional) # string | The status of the profile. It can be multiple values with comma separated. (optional)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.IscAccountsAPI.GetSchemaMappedProfilesCollection(context.Background()).Execute()
	  //resp, r, err := apiClient.NERM.IscAccountsAPI.GetSchemaMappedProfilesCollection(context.Background()).Limit(limit).Offset(offset).UseSchema(useSchema).OverrideSyncToggle(overrideSyncToggle).Category(category).Status(status).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `IscAccountsAPI.GetSchemaMappedProfilesCollection``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchemaMappedProfilesCollection`: GetSchemaMappedProfilesCollection200Response
    fmt.Fprintf(os.Stdout, "Response from `IscAccountsAPI.GetSchemaMappedProfilesCollection`: %v\n", resp)
}
```

[[Back to top]](#)

## get-single-schema-mapped-profile
Get Profile
It returns a single stored profile, optionally with schema-mapped field names.

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-single-schema-mapped-profile)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSingleSchemaMappedProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSingleSchemaMappedProfile200Response**](../models/get-single-schema-mapped-profile200-response)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    id := `1246d8b3-ac29-4015-8154-dea4434a73fa` // string | ID of the object to retrieve, update, or delete # string | ID of the object to retrieve, update, or delete

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.IscAccountsAPI.GetSingleSchemaMappedProfile(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.IscAccountsAPI.GetSingleSchemaMappedProfile(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `IscAccountsAPI.GetSingleSchemaMappedProfile``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSingleSchemaMappedProfile`: GetSingleSchemaMappedProfile200Response
    fmt.Fprintf(os.Stdout, "Response from `IscAccountsAPI.GetSingleSchemaMappedProfile`: %v\n", resp)
}
```

[[Back to top]](#)

## update-profile
Update Profile
Updates a profile only through ISC schema-mapped attributes, performs a reverse mapping to match the NERM attributes to update.

[API Spec](https://developer.sailpoint.com/docs/api/NERM/update-profile)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateProfileRequest** | [**UpdateProfileRequest**](../models/update-profile-request) |  | 

### Return type

[**GetSingleSchemaMappedProfile200Response**](../models/get-single-schema-mapped-profile200-response)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    id := `1246d8b3-ac29-4015-8154-dea4434a73fa` // string | ID of the object to retrieve, update, or delete # string | ID of the object to retrieve, update, or delete
    updateprofilerequest := []byte(``) // UpdateProfileRequest |  (optional)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.IscAccountsAPI.UpdateProfile(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.IscAccountsAPI.UpdateProfile(context.Background(), id).UpdateProfileRequest(updateProfileRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `IscAccountsAPI.UpdateProfile``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProfile`: GetSingleSchemaMappedProfile200Response
    fmt.Fprintf(os.Stdout, "Response from `IscAccountsAPI.UpdateProfile`: %v\n", resp)
}
```

[[Back to top]](#)

