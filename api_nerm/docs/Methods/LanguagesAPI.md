---
id: nerm-languages
title: Languages
pagination_label: Languages
sidebar_label: Languages
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Languages', 'NERMLanguages'] 
slug: /tools/sdk/go/nerm/methods/languages
tags: ['SDK', 'Software Development Kit', 'Languages', 'NERMLanguages']
---

# LanguagesAPI
   
All URIs are relative to *https://acmeco.nonemployee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**patch-language**](#patch-language) | **Patch** `/languages/{locale}` | Update a language by locale


## patch-language
Update a language by locale
Update a language by locale

[API Spec](https://developer.sailpoint.com/docs/api/NERM/patch-language)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**languageLocale** | **string** | Language locale of the object | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchLanguageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchLanguageRequest** | [**PatchLanguageRequest**](../models/patch-language-request) |  | 

### Return type

[**PatchLanguageRequest**](../models/patch-language-request)

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
  "encoding/json"
    NERM "github.com/sailpoint-oss/golang-sdk/v2/api_nerm"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    languageLocale := `es` // string | Language locale of the object # string | Language locale of the object
    patchlanguagerequest := []byte(``) // PatchLanguageRequest | 

    var patchLanguageRequest NERM.PatchLanguageRequest
    if err := json.Unmarshal(patchlanguagerequest, &patchLanguageRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.LanguagesAPI.PatchLanguage(context.Background(), languageLocale).PatchLanguageRequest(patchLanguageRequest).Execute()
	  //resp, r, err := apiClient.NERM.LanguagesAPI.PatchLanguage(context.Background(), languageLocale).PatchLanguageRequest(patchLanguageRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `LanguagesAPI.PatchLanguage``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchLanguage`: PatchLanguageRequest
    fmt.Fprintf(os.Stdout, "Response from `LanguagesAPI.PatchLanguage`: %v\n", resp)
}
```

[[Back to top]](#)

