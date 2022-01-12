/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// SnapshotServiceApiService SnapshotServiceApi service
type SnapshotServiceApiService service

type ApiSnapshotServiceCreateRequest struct {
	ctx _context.Context
	ApiService *SnapshotServiceApiService
	snapshotCreateRequest *SnapshotCreateRequest
}

func (r ApiSnapshotServiceCreateRequest) SnapshotCreateRequest(snapshotCreateRequest SnapshotCreateRequest) ApiSnapshotServiceCreateRequest {
	r.snapshotCreateRequest = &snapshotCreateRequest
	return r
}

func (r ApiSnapshotServiceCreateRequest) Execute() (SnapshotCreateResponse, *_nethttp.Response, error) {
	return r.ApiService.SnapshotServiceCreateExecute(r)
}

/*
SnapshotServiceCreate Method for SnapshotServiceCreate

create a snapshot

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSnapshotServiceCreateRequest
*/
func (a *SnapshotServiceApiService) SnapshotServiceCreate(ctx _context.Context) ApiSnapshotServiceCreateRequest {
	return ApiSnapshotServiceCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SnapshotCreateResponse
func (a *SnapshotServiceApiService) SnapshotServiceCreateExecute(r ApiSnapshotServiceCreateRequest) (SnapshotCreateResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  SnapshotCreateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotServiceApiService.SnapshotServiceCreate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/vps/snapshot"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.snapshotCreateRequest == nil {
		return localVarReturnValue, nil, reportError("snapshotCreateRequest is required and must be specified")
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
	// body params
	localVarPostBody = r.snapshotCreateRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSnapshotServiceCreateCalculatorRequest struct {
	ctx _context.Context
	ApiService *SnapshotServiceApiService
	snapshotCreateCalculatorRequest *SnapshotCreateCalculatorRequest
}

func (r ApiSnapshotServiceCreateCalculatorRequest) SnapshotCreateCalculatorRequest(snapshotCreateCalculatorRequest SnapshotCreateCalculatorRequest) ApiSnapshotServiceCreateCalculatorRequest {
	r.snapshotCreateCalculatorRequest = &snapshotCreateCalculatorRequest
	return r
}

func (r ApiSnapshotServiceCreateCalculatorRequest) Execute() (SnapshotCreateCalculatorResponse, *_nethttp.Response, error) {
	return r.ApiService.SnapshotServiceCreateCalculatorExecute(r)
}

/*
SnapshotServiceCreateCalculator Method for SnapshotServiceCreateCalculator

create cost calculator for snapshot creating

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSnapshotServiceCreateCalculatorRequest
*/
func (a *SnapshotServiceApiService) SnapshotServiceCreateCalculator(ctx _context.Context) ApiSnapshotServiceCreateCalculatorRequest {
	return ApiSnapshotServiceCreateCalculatorRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SnapshotCreateCalculatorResponse
func (a *SnapshotServiceApiService) SnapshotServiceCreateCalculatorExecute(r ApiSnapshotServiceCreateCalculatorRequest) (SnapshotCreateCalculatorResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  SnapshotCreateCalculatorResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotServiceApiService.SnapshotServiceCreateCalculator")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/vps/snapshot/calculator"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.snapshotCreateCalculatorRequest == nil {
		return localVarReturnValue, nil, reportError("snapshotCreateCalculatorRequest is required and must be specified")
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
	// body params
	localVarPostBody = r.snapshotCreateCalculatorRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSnapshotServiceEditRequest struct {
	ctx _context.Context
	ApiService *SnapshotServiceApiService
	id string
	snapshotEditRequest *SnapshotEditRequest
}

func (r ApiSnapshotServiceEditRequest) SnapshotEditRequest(snapshotEditRequest SnapshotEditRequest) ApiSnapshotServiceEditRequest {
	r.snapshotEditRequest = &snapshotEditRequest
	return r
}

func (r ApiSnapshotServiceEditRequest) Execute() (SnapshotEditResponse, *_nethttp.Response, error) {
	return r.ApiService.SnapshotServiceEditExecute(r)
}

/*
SnapshotServiceEdit Method for SnapshotServiceEdit

edit snapshot user fields

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiSnapshotServiceEditRequest
*/
func (a *SnapshotServiceApiService) SnapshotServiceEdit(ctx _context.Context, id string) ApiSnapshotServiceEditRequest {
	return ApiSnapshotServiceEditRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return SnapshotEditResponse
func (a *SnapshotServiceApiService) SnapshotServiceEditExecute(r ApiSnapshotServiceEditRequest) (SnapshotEditResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  SnapshotEditResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotServiceApiService.SnapshotServiceEdit")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/vps/snapshot/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.snapshotEditRequest == nil {
		return localVarReturnValue, nil, reportError("snapshotEditRequest is required and must be specified")
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
	// body params
	localVarPostBody = r.snapshotEditRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSnapshotServiceGetAllRequest struct {
	ctx _context.Context
	ApiService *SnapshotServiceApiService
}


func (r ApiSnapshotServiceGetAllRequest) Execute() (SnapshotGetAllResponse, *_nethttp.Response, error) {
	return r.ApiService.SnapshotServiceGetAllExecute(r)
}

/*
SnapshotServiceGetAll Method for SnapshotServiceGetAll

get all snapshots

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSnapshotServiceGetAllRequest
*/
func (a *SnapshotServiceApiService) SnapshotServiceGetAll(ctx _context.Context) ApiSnapshotServiceGetAllRequest {
	return ApiSnapshotServiceGetAllRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SnapshotGetAllResponse
func (a *SnapshotServiceApiService) SnapshotServiceGetAllExecute(r ApiSnapshotServiceGetAllRequest) (SnapshotGetAllResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  SnapshotGetAllResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotServiceApiService.SnapshotServiceGetAll")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/vps/snapshot"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSnapshotServiceGetAllRestoresRequest struct {
	ctx _context.Context
	ApiService *SnapshotServiceApiService
	id *string
}

// snapshot id. Optional, full list at 0
func (r ApiSnapshotServiceGetAllRestoresRequest) Id(id string) ApiSnapshotServiceGetAllRestoresRequest {
	r.id = &id
	return r
}

func (r ApiSnapshotServiceGetAllRestoresRequest) Execute() (SnapshotGetAllRestoresResponse, *_nethttp.Response, error) {
	return r.ApiService.SnapshotServiceGetAllRestoresExecute(r)
}

/*
SnapshotServiceGetAllRestores Method for SnapshotServiceGetAllRestores

get restore list

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSnapshotServiceGetAllRestoresRequest
*/
func (a *SnapshotServiceApiService) SnapshotServiceGetAllRestores(ctx _context.Context) ApiSnapshotServiceGetAllRestoresRequest {
	return ApiSnapshotServiceGetAllRestoresRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SnapshotGetAllRestoresResponse
func (a *SnapshotServiceApiService) SnapshotServiceGetAllRestoresExecute(r ApiSnapshotServiceGetAllRestoresRequest) (SnapshotGetAllRestoresResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  SnapshotGetAllRestoresResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotServiceApiService.SnapshotServiceGetAllRestores")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/vps/snapshot/restore"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.id != nil {
		localVarQueryParams.Add("id", parameterToString(*r.id, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSnapshotServiceRemoveRequest struct {
	ctx _context.Context
	ApiService *SnapshotServiceApiService
	id string
}


func (r ApiSnapshotServiceRemoveRequest) Execute() (SnapshotRemoveResponse, *_nethttp.Response, error) {
	return r.ApiService.SnapshotServiceRemoveExecute(r)
}

/*
SnapshotServiceRemove Method for SnapshotServiceRemove

remove the snapshot

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiSnapshotServiceRemoveRequest
*/
func (a *SnapshotServiceApiService) SnapshotServiceRemove(ctx _context.Context, id string) ApiSnapshotServiceRemoveRequest {
	return ApiSnapshotServiceRemoveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return SnapshotRemoveResponse
func (a *SnapshotServiceApiService) SnapshotServiceRemoveExecute(r ApiSnapshotServiceRemoveRequest) (SnapshotRemoveResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  SnapshotRemoveResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotServiceApiService.SnapshotServiceRemove")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/vps/snapshot/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSnapshotServiceRestoreRequest struct {
	ctx _context.Context
	ApiService *SnapshotServiceApiService
	id string
	snapshotRestoreRequest *SnapshotRestoreRequest
}

func (r ApiSnapshotServiceRestoreRequest) SnapshotRestoreRequest(snapshotRestoreRequest SnapshotRestoreRequest) ApiSnapshotServiceRestoreRequest {
	r.snapshotRestoreRequest = &snapshotRestoreRequest
	return r
}

func (r ApiSnapshotServiceRestoreRequest) Execute() (SnapshotRestoreResponse, *_nethttp.Response, error) {
	return r.ApiService.SnapshotServiceRestoreExecute(r)
}

/*
SnapshotServiceRestore Method for SnapshotServiceRestore

create a restore for the snapshot

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiSnapshotServiceRestoreRequest
*/
func (a *SnapshotServiceApiService) SnapshotServiceRestore(ctx _context.Context, id string) ApiSnapshotServiceRestoreRequest {
	return ApiSnapshotServiceRestoreRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return SnapshotRestoreResponse
func (a *SnapshotServiceApiService) SnapshotServiceRestoreExecute(r ApiSnapshotServiceRestoreRequest) (SnapshotRestoreResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  SnapshotRestoreResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotServiceApiService.SnapshotServiceRestore")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/vps/snapshot/{id}/restore"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.snapshotRestoreRequest == nil {
		return localVarReturnValue, nil, reportError("snapshotRestoreRequest is required and must be specified")
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
	// body params
	localVarPostBody = r.snapshotRestoreRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
