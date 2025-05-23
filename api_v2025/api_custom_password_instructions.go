/*
Identity Security Cloud V2025 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2025
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2025

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// CustomPasswordInstructionsAPIService CustomPasswordInstructionsAPI service
type CustomPasswordInstructionsAPIService service

type ApiCreateCustomPasswordInstructionsRequest struct {
	ctx context.Context
	ApiService *CustomPasswordInstructionsAPIService
	xSailPointExperimental *string
	customPasswordInstruction *CustomPasswordInstruction
}

// Use this header to enable this experimental API.
func (r ApiCreateCustomPasswordInstructionsRequest) XSailPointExperimental(xSailPointExperimental string) ApiCreateCustomPasswordInstructionsRequest {
	r.xSailPointExperimental = &xSailPointExperimental
	return r
}

func (r ApiCreateCustomPasswordInstructionsRequest) CustomPasswordInstruction(customPasswordInstruction CustomPasswordInstruction) ApiCreateCustomPasswordInstructionsRequest {
	r.customPasswordInstruction = &customPasswordInstruction
	return r
}

func (r ApiCreateCustomPasswordInstructionsRequest) Execute() (*CustomPasswordInstruction, *http.Response, error) {
	return r.ApiService.CreateCustomPasswordInstructionsExecute(r)
}

/*
CreateCustomPasswordInstructions Create custom password instructions

This API creates the custom password instructions for the specified page ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateCustomPasswordInstructionsRequest
*/
func (a *CustomPasswordInstructionsAPIService) CreateCustomPasswordInstructions(ctx context.Context) ApiCreateCustomPasswordInstructionsRequest {
	return ApiCreateCustomPasswordInstructionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CustomPasswordInstruction
func (a *CustomPasswordInstructionsAPIService) CreateCustomPasswordInstructionsExecute(r ApiCreateCustomPasswordInstructionsRequest) (*CustomPasswordInstruction, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomPasswordInstruction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomPasswordInstructionsAPIService.CreateCustomPasswordInstructions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/custom-password-instructions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	
	if r.xSailPointExperimental == nil {
		headerxSailPointExperimental := "true"
		r.xSailPointExperimental = &headerxSailPointExperimental
	}
	
	if r.xSailPointExperimental == nil {
		return localVarReturnValue, nil, reportError("xSailPointExperimental is required and must be specified")
	}
	
	if r.xSailPointExperimental == nil {
		headerxSailPointExperimental := "true"
		r.xSailPointExperimental = &headerxSailPointExperimental
	}
	
	if r.customPasswordInstruction == nil {
		return localVarReturnValue, nil, reportError("customPasswordInstruction is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-SailPoint-Experimental", r.xSailPointExperimental, "", "")
	// body params
	localVarPostBody = r.customPasswordInstruction
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponseDto
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponseDto
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorResponseDto
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteCustomPasswordInstructionsRequest struct {
	ctx context.Context
	ApiService *CustomPasswordInstructionsAPIService
	pageId string
	xSailPointExperimental *string
	locale *string
}

// Use this header to enable this experimental API.
func (r ApiDeleteCustomPasswordInstructionsRequest) XSailPointExperimental(xSailPointExperimental string) ApiDeleteCustomPasswordInstructionsRequest {
	r.xSailPointExperimental = &xSailPointExperimental
	return r
}

// The locale for the custom instructions, a BCP47 language tag. The default value is \\\&quot;default\\\&quot;.
func (r ApiDeleteCustomPasswordInstructionsRequest) Locale(locale string) ApiDeleteCustomPasswordInstructionsRequest {
	r.locale = &locale
	return r
}

func (r ApiDeleteCustomPasswordInstructionsRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteCustomPasswordInstructionsExecute(r)
}

/*
DeleteCustomPasswordInstructions Delete custom password instructions by page id

This API delete the custom password instructions for the specified page ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pageId The page ID of custom password instructions to delete.
 @return ApiDeleteCustomPasswordInstructionsRequest
*/
func (a *CustomPasswordInstructionsAPIService) DeleteCustomPasswordInstructions(ctx context.Context, pageId string) ApiDeleteCustomPasswordInstructionsRequest {
	return ApiDeleteCustomPasswordInstructionsRequest{
		ApiService: a,
		ctx: ctx,
		pageId: pageId,
	}
}

// Execute executes the request
func (a *CustomPasswordInstructionsAPIService) DeleteCustomPasswordInstructionsExecute(r ApiDeleteCustomPasswordInstructionsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomPasswordInstructionsAPIService.DeleteCustomPasswordInstructions")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/custom-password-instructions/{pageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"pageId"+"}", url.PathEscape(parameterValueToString(r.pageId, "pageId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	
	if r.xSailPointExperimental == nil {
		headerxSailPointExperimental := "true"
		r.xSailPointExperimental = &headerxSailPointExperimental
	}
	
	if r.xSailPointExperimental == nil {
		return nil, reportError("xSailPointExperimental is required and must be specified")
	}

	if r.locale != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "locale", r.locale, "", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-SailPoint-Experimental", r.xSailPointExperimental, "", "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponseDto
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponseDto
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponseDto
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorResponseDto
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetCustomPasswordInstructionsRequest struct {
	ctx context.Context
	ApiService *CustomPasswordInstructionsAPIService
	pageId string
	xSailPointExperimental *string
	locale *string
}

// Use this header to enable this experimental API.
func (r ApiGetCustomPasswordInstructionsRequest) XSailPointExperimental(xSailPointExperimental string) ApiGetCustomPasswordInstructionsRequest {
	r.xSailPointExperimental = &xSailPointExperimental
	return r
}

// The locale for the custom instructions, a BCP47 language tag. The default value is \\\&quot;default\\\&quot;.
func (r ApiGetCustomPasswordInstructionsRequest) Locale(locale string) ApiGetCustomPasswordInstructionsRequest {
	r.locale = &locale
	return r
}

func (r ApiGetCustomPasswordInstructionsRequest) Execute() (*CustomPasswordInstruction, *http.Response, error) {
	return r.ApiService.GetCustomPasswordInstructionsExecute(r)
}

/*
GetCustomPasswordInstructions Get custom password instructions by page id

This API returns the custom password instructions for the specified page ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pageId The page ID of custom password instructions to query.
 @return ApiGetCustomPasswordInstructionsRequest
*/
func (a *CustomPasswordInstructionsAPIService) GetCustomPasswordInstructions(ctx context.Context, pageId string) ApiGetCustomPasswordInstructionsRequest {
	return ApiGetCustomPasswordInstructionsRequest{
		ApiService: a,
		ctx: ctx,
		pageId: pageId,
	}
}

// Execute executes the request
//  @return CustomPasswordInstruction
func (a *CustomPasswordInstructionsAPIService) GetCustomPasswordInstructionsExecute(r ApiGetCustomPasswordInstructionsRequest) (*CustomPasswordInstruction, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomPasswordInstruction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomPasswordInstructionsAPIService.GetCustomPasswordInstructions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/custom-password-instructions/{pageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"pageId"+"}", url.PathEscape(parameterValueToString(r.pageId, "pageId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	
	if r.xSailPointExperimental == nil {
		headerxSailPointExperimental := "true"
		r.xSailPointExperimental = &headerxSailPointExperimental
	}
	
	if r.xSailPointExperimental == nil {
		return localVarReturnValue, nil, reportError("xSailPointExperimental is required and must be specified")
	}

	if r.locale != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "locale", r.locale, "", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-SailPoint-Experimental", r.xSailPointExperimental, "", "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponseDto
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponseDto
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponseDto
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorResponseDto
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
