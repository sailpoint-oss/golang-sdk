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
[**DeleteUploadedBackup**](#delete-uploaded-backup) | **Delete** /configuration-hub/backups/uploads/{id} | Deletes an uploaded backup file
[**GetObjectMappings**](#get-object-mappings) | **Get** /configuration-hub/object-mappings/{sourceOrg} | Gets list of object mappings
[**GetUploadedBackup**](#get-uploaded-backup) | **Get** /configuration-hub/backups/uploads/{id} | Get an uploaded backup&#39;s information
[**GetUploadedBackups**](#get-uploaded-backups) | **Get** /configuration-hub/backups/uploads | Gets list of Uploaded backups
[**ImportUploadedBackup**](#import-uploaded-backup) | **Post** /configuration-hub/backups/uploads | Uploads a backup file
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


## delete-uploaded-backup


This deletes an Uploaded backup based on job ID.
On success, this endpoint will return an empty response.
The job id can be obtained from the response after a successful upload, or the list uploads endpoint.
The following scopes are required to access this endpoint: sp:config:manage

### Parameters 
Param Type | Name | Data Type | Required  | Description
------------- | ------------- | ------------- | ------------- | ------------- 
Path   | id | **string** | True  | The id of the uploaded backup.

	
### Return type

 (empty response body)

### Responses
Code | Description  | Data Type
------------- | ------------- | -------------
204 | No content - indicates the request was successful but there is no content to be returned in the response. | 
400 | Client Error - Returned if the request body is invalid. | ErrorResponseDto
401 | Unauthorized - Returned if there is no authorization header, or if the JWT token is expired. | ListAccessProfiles401Response
403 | Forbidden - Returned if the user you are running as, doesn&#39;t have access to this end-point. | ErrorResponseDto
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


## get-uploaded-backup


Returns all the information and status of an upload job.
- sp:config-backups:read

### Parameters 
Param Type | Name | Data Type | Required  | Description
------------- | ------------- | ------------- | ------------- | ------------- 
Path   | id | **string** | True  | The id of the uploaded backup.

	
### Return type

**map[string]interface{}**

### Responses
Code | Description  | Data Type
------------- | ------------- | -------------
200 | List of existing uploaded backups. | map[string]interface{}
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


## get-uploaded-backups


Returns a list of the current uploaded backups associated with the current tenant.
A filter "status" can be added to only return the Completed, Failed, or Successful uploads

### Parameters 
Param Type | Name | Data Type | Required  | Description
------------- | ------------- | ------------- | ------------- | ------------- 
  Query | status | **string** |   (optional) | Filter listed uploaded backups by status of operation

	
### Return type

[**[]UploadsResponse**](../models/uploads-response)

### Responses
Code | Description  | Data Type
------------- | ------------- | -------------
200 | List of existing uploaded backups. | []UploadsResponse
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


## import-uploaded-backup


This post will upload a JSON backup file into a tenant. Configuration files can be managed and deployed via Configuration Hub by uploading a json file which contains configuration data. The JSON file should be the same as the one used by our import endpoints. The object types that currently support by upload file functionality are the same as the ones supported by our regular backup functionality. here: [SaaS Configuration](https://developer.sailpoint.com/idn/docs/saas-configuration/#supported-objects).

The request will need the following security scope:
- sp:config:manage

### Parameters 
Param Type | Name | Data Type | Required  | Description
------------- | ------------- | ------------- | ------------- | ------------- 
   | data | ***os.File** | True  | JSON file containing the objects to be imported.
   | name | **string** | True  | Name that will be assigned to the uploaded file.

	
### Return type

[**UploadsRequest**](../models/uploads-request)

### Responses
Code | Description  | Data Type
------------- | ------------- | -------------
202 | Upload job accepted and queued for processing. | UploadsRequest
400 | Client Error - Returned if the request body is invalid.  | ErrorResponseDto
401 | Unauthorized - Returned if there is no authorization header, or if the JWT token is expired. | ListAccessProfiles401Response
403 | Forbidden - Returned if the user you are running as, doesn&#39;t have access to this end-point. | ErrorResponseDto
429 | Too Many Requests - Returned in response to too many requests in a given period of time - rate limited. The Retry-After header in the response includes how long to wait before trying again. | ListAccessProfiles429Response
500 | Internal Server Error - Returned if there is an unexpected error. | ErrorResponseDto


### HTTP request headers

- **Content-Type**: multipart/form-data
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

