---
id: configuration-hub
title: ConfigurationHub
pagination_label: ConfigurationHub
sidebar_label: ConfigurationHub
sidebar_class_name: gosdk
keywords: ['go', 'golang', 'sdk', 'ConfigurationHub'] 
slug: /tools/sdk/go/v3/methods/configuration-hub
tags: ['SDK', 'Software Development Kit', 'ConfigurationHub']
---


# ConfigurationHub

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateObjectMapping**](#create-object-mapping) | **Post** /configuration-hub/object-mappings/{sourceOrg} | Creates an object mapping
[**CreateObjectMappings**](#create-object-mappings) | **Post** /configuration-hub/object-mappings/{sourceOrg}/bulk-create | Bulk creates object mappings
[**DeleteObjectMapping**](#delete-object-mapping) | **Delete** /configuration-hub/object-mappings/{sourceOrg}/{objectMappingId} | Deletes an object mapping
[**GetObjectMappings**](#get-object-mappings) | **Get** /configuration-hub/object-mappings/{sourceOrg} | Gets list of object mappings
[**UpdateObjectMappings**](#update-object-mappings) | **Post** /configuration-hub/object-mappings/{sourceOrg}/bulk-patch | Bulk updates object mappings



## create-object-mapping


This creates an object mapping between current org and source org.
The request will need the following security scope:
- sp:config-object-mapping:manage

### Parameters 
Param Type | Name | Data Type | Required  | Description
------------- | ------------- | ------------- | ------------- | ------------- 
Path   | sourceOrg | **string** | True  | The name of the source org.
 Body  | objectMappingRequest | [**ObjectMappingRequest**](../models/object-mapping-request) | True  | The object mapping request body.

	
### Return type

[**ObjectMappingResponse**](../models/object-mapping-response)

### Responses
Code | Description  | Data Type
------------- | ------------- | -------------
200 | The created object mapping between current org and source org. | ObjectMappingResponse
400 | Client Error - Returned if the request body is invalid. | ErrorResponseDto
401 | Unauthorized - Returned if there is no authorization header, or if the JWT token is expired. | ListAccessProfiles401Response
403 | Forbidden - Returned if the user you are running as, doesn&#39;t have access to this end-point. | ErrorResponseDto
404 | Not Found - returned if the request URL refers to a resource or object that does not exist | ErrorResponseDto
429 | Too Many Requests - Returned in response to too many requests in a given period of time - rate limited. The Retry-After header in the response includes how long to wait before trying again. | ListAccessProfiles429Response
500 | Internal Server Error - Returned if there is an unexpected error. | ErrorResponseDto


### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) 


## create-object-mappings


This creates a set of object mappings (Max 25) between current org and source org.
The request will need the following security scope:
- sp:config-object-mapping:manage

### Parameters 
Param Type | Name | Data Type | Required  | Description
------------- | ------------- | ------------- | ------------- | ------------- 
Path   | sourceOrg | **string** | True  | The name of the source org.
 Body  | objectMappingBulkCreateRequest | [**ObjectMappingBulkCreateRequest**](../models/object-mapping-bulk-create-request) | True  | The bulk create object mapping request body.

	
### Return type

[**ObjectMappingBulkCreateResponse**](../models/object-mapping-bulk-create-response)

### Responses
Code | Description  | Data Type
------------- | ------------- | -------------
200 | The created object mapping between current org and source org. | ObjectMappingBulkCreateResponse
400 | Client Error - Returned if the request body is invalid. | ErrorResponseDto
401 | Unauthorized - Returned if there is no authorization header, or if the JWT token is expired. | ListAccessProfiles401Response
403 | Forbidden - Returned if the user you are running as, doesn&#39;t have access to this end-point. | ErrorResponseDto
404 | Not Found - returned if the request URL refers to a resource or object that does not exist | ErrorResponseDto
429 | Too Many Requests - Returned in response to too many requests in a given period of time - rate limited. The Retry-After header in the response includes how long to wait before trying again. | ListAccessProfiles429Response
500 | Internal Server Error - Returned if there is an unexpected error. | ErrorResponseDto


### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) 


## delete-object-mapping


This deletes an existing object mapping.
The request will need the following security scope:
- sp:config-object-mapping:manage

### Parameters 
Param Type | Name | Data Type | Required  | Description
------------- | ------------- | ------------- | ------------- | ------------- 
Path   | sourceOrg | **string** | True  | The name of the source org.
Path   | objectMappingId | **string** | True  | The id of the object mapping to be deleted.

	
### Return type

 (empty response body)

### Responses
Code | Description  | Data Type
------------- | ------------- | -------------
204 | No content - indicates the request was successful but there is no content to be returned in the response. | 
400 | Client Error - Returned if the request body is invalid. | ErrorResponseDto
401 | Unauthorized - Returned if there is no authorization header, or if the JWT token is expired. | ListAccessProfiles401Response
403 | Forbidden - Returned if the user you are running as, doesn&#39;t have access to this end-point. | ErrorResponseDto
404 | Not Found - returned if the request URL refers to a resource or object that does not exist | ErrorResponseDto
429 | Too Many Requests - Returned in response to too many requests in a given period of time - rate limited. The Retry-After header in the response includes how long to wait before trying again. | ListAccessProfiles429Response
500 | Internal Server Error - Returned if there is an unexpected error. | ErrorResponseDto


### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) 


## get-object-mappings


This gets a list of existing object mappings between current org and source org.
The request will need the following security scope:
- sp:config-object-mapping:read

### Parameters 
Param Type | Name | Data Type | Required  | Description
------------- | ------------- | ------------- | ------------- | ------------- 
Path   | sourceOrg | **string** | True  | The name of the source org.

	
### Return type

[**[]ObjectMappingResponse**](../models/object-mapping-response)

### Responses
Code | Description  | Data Type
------------- | ------------- | -------------
200 | List of existing object mappings between current org and source org. | []ObjectMappingResponse
400 | Client Error - Returned if the request body is invalid. | ErrorResponseDto
401 | Unauthorized - Returned if there is no authorization header, or if the JWT token is expired. | ListAccessProfiles401Response
403 | Forbidden - Returned if the user you are running as, doesn&#39;t have access to this end-point. | ErrorResponseDto
404 | Not Found - returned if the request URL refers to a resource or object that does not exist | ErrorResponseDto
429 | Too Many Requests - Returned in response to too many requests in a given period of time - rate limited. The Retry-After header in the response includes how long to wait before trying again. | ListAccessProfiles429Response
500 | Internal Server Error - Returned if there is an unexpected error. | ErrorResponseDto


### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) 


## update-object-mappings


This updates a set of object mappings, only enabled and targetValue fields can be updated.
The request will need the following security scope:
- sp:config-object-mapping:manage

### Parameters 
Param Type | Name | Data Type | Required  | Description
------------- | ------------- | ------------- | ------------- | ------------- 
Path   | sourceOrg | **string** | True  | The name of the source org.
 Body  | objectMappingBulkPatchRequest | [**ObjectMappingBulkPatchRequest**](../models/object-mapping-bulk-patch-request) | True  | The object mapping request body.

	
### Return type

[**ObjectMappingBulkPatchResponse**](../models/object-mapping-bulk-patch-response)

### Responses
Code | Description  | Data Type
------------- | ------------- | -------------
200 | The updated object mappings. | ObjectMappingBulkPatchResponse
400 | Client Error - Returned if the request body is invalid. | ErrorResponseDto
401 | Unauthorized - Returned if there is no authorization header, or if the JWT token is expired. | ListAccessProfiles401Response
403 | Forbidden - Returned if the user you are running as, doesn&#39;t have access to this end-point. | ErrorResponseDto
404 | Not Found - returned if the request URL refers to a resource or object that does not exist | ErrorResponseDto
429 | Too Many Requests - Returned in response to too many requests in a given period of time - rate limited. The Retry-After header in the response includes how long to wait before trying again. | ListAccessProfiles429Response
500 | Internal Server Error - Returned if there is an unexpected error. | ErrorResponseDto


### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) 

