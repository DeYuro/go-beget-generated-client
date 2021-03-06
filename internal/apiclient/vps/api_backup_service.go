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

// BackupServiceApiService BackupServiceApi service
type BackupServiceApiService service

type ApiBackupServiceGetAvailableCopiesRequest struct {
	ctx _context.Context
	ApiService *BackupServiceApiService
}


func (r ApiBackupServiceGetAvailableCopiesRequest) Execute() (BackupGetAvailableCopiesResponse, *_nethttp.Response, error) {
	return r.ApiService.BackupServiceGetAvailableCopiesExecute(r)
}

/*
BackupServiceGetAvailableCopies Method for BackupServiceGetAvailableCopies

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiBackupServiceGetAvailableCopiesRequest
*/
func (a *BackupServiceApiService) BackupServiceGetAvailableCopies(ctx _context.Context) ApiBackupServiceGetAvailableCopiesRequest {
	return ApiBackupServiceGetAvailableCopiesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BackupGetAvailableCopiesResponse
func (a *BackupServiceApiService) BackupServiceGetAvailableCopiesExecute(r ApiBackupServiceGetAvailableCopiesRequest) (BackupGetAvailableCopiesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  BackupGetAvailableCopiesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BackupServiceApiService.BackupServiceGetAvailableCopies")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/vps/backup"

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

type ApiBackupServiceGetBackupFileListRequest struct {
	ctx _context.Context
	ApiService *BackupServiceApiService
	id string
	copyId string
	path *string
}

// ???????? ?????? ??????????????????????
func (r ApiBackupServiceGetBackupFileListRequest) Path(path string) ApiBackupServiceGetBackupFileListRequest {
	r.path = &path
	return r
}

func (r ApiBackupServiceGetBackupFileListRequest) Execute() (BackupGetBackupFileListResponse, *_nethttp.Response, error) {
	return r.ApiService.BackupServiceGetBackupFileListExecute(r)
}

/*
BackupServiceGetBackupFileList Method for BackupServiceGetBackupFileList

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @param copyId
 @return ApiBackupServiceGetBackupFileListRequest
*/
func (a *BackupServiceApiService) BackupServiceGetBackupFileList(ctx _context.Context, id string, copyId string) ApiBackupServiceGetBackupFileListRequest {
	return ApiBackupServiceGetBackupFileListRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		copyId: copyId,
	}
}

// Execute executes the request
//  @return BackupGetBackupFileListResponse
func (a *BackupServiceApiService) BackupServiceGetBackupFileListExecute(r ApiBackupServiceGetBackupFileListRequest) (BackupGetBackupFileListResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  BackupGetBackupFileListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BackupServiceApiService.BackupServiceGetBackupFileList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/vps/{id}/backup/{copy_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"copy_id"+"}", _neturl.PathEscape(parameterToString(r.copyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.path != nil {
		localVarQueryParams.Add("path", parameterToString(*r.path, ""))
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

type ApiBackupServiceGetOrdersRequest struct {
	ctx _context.Context
	ApiService *BackupServiceApiService
	limit *string
	offset *string
}

// ???????????????????? ?????????????? ???? ???????????????? (???? 1 ???? 100)
func (r ApiBackupServiceGetOrdersRequest) Limit(limit string) ApiBackupServiceGetOrdersRequest {
	r.limit = &limit
	return r
}
// ???????????????? ???????????????????????? ?????????????? (??????????????????) ????????????
func (r ApiBackupServiceGetOrdersRequest) Offset(offset string) ApiBackupServiceGetOrdersRequest {
	r.offset = &offset
	return r
}

func (r ApiBackupServiceGetOrdersRequest) Execute() (BackupGetOrdersResponse, *_nethttp.Response, error) {
	return r.ApiService.BackupServiceGetOrdersExecute(r)
}

/*
BackupServiceGetOrders Method for BackupServiceGetOrders

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiBackupServiceGetOrdersRequest
*/
func (a *BackupServiceApiService) BackupServiceGetOrders(ctx _context.Context) ApiBackupServiceGetOrdersRequest {
	return ApiBackupServiceGetOrdersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BackupGetOrdersResponse
func (a *BackupServiceApiService) BackupServiceGetOrdersExecute(r ApiBackupServiceGetOrdersRequest) (BackupGetOrdersResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  BackupGetOrdersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BackupServiceApiService.BackupServiceGetOrders")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/vps/backup/orders"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
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

type ApiBackupServiceRestoreFileRequest struct {
	ctx _context.Context
	ApiService *BackupServiceApiService
	id string
	copyId string
	backupRestoreFileRequest *BackupRestoreFileRequest
}

func (r ApiBackupServiceRestoreFileRequest) BackupRestoreFileRequest(backupRestoreFileRequest BackupRestoreFileRequest) ApiBackupServiceRestoreFileRequest {
	r.backupRestoreFileRequest = &backupRestoreFileRequest
	return r
}

func (r ApiBackupServiceRestoreFileRequest) Execute() (BackupRestoreFileResponse, *_nethttp.Response, error) {
	return r.ApiService.BackupServiceRestoreFileExecute(r)
}

/*
BackupServiceRestoreFile Method for BackupServiceRestoreFile

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @param copyId
 @return ApiBackupServiceRestoreFileRequest
*/
func (a *BackupServiceApiService) BackupServiceRestoreFile(ctx _context.Context, id string, copyId string) ApiBackupServiceRestoreFileRequest {
	return ApiBackupServiceRestoreFileRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		copyId: copyId,
	}
}

// Execute executes the request
//  @return BackupRestoreFileResponse
func (a *BackupServiceApiService) BackupServiceRestoreFileExecute(r ApiBackupServiceRestoreFileRequest) (BackupRestoreFileResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  BackupRestoreFileResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BackupServiceApiService.BackupServiceRestoreFile")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/vps/{id}/backup/{copy_id}/file"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"copy_id"+"}", _neturl.PathEscape(parameterToString(r.copyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.backupRestoreFileRequest == nil {
		return localVarReturnValue, nil, reportError("backupRestoreFileRequest is required and must be specified")
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
	localVarPostBody = r.backupRestoreFileRequest
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

type ApiBackupServiceRestoreServerRequest struct {
	ctx _context.Context
	ApiService *BackupServiceApiService
	id string
	copyId string
	backupRestoreServerRequest *BackupRestoreServerRequest
}

func (r ApiBackupServiceRestoreServerRequest) BackupRestoreServerRequest(backupRestoreServerRequest BackupRestoreServerRequest) ApiBackupServiceRestoreServerRequest {
	r.backupRestoreServerRequest = &backupRestoreServerRequest
	return r
}

func (r ApiBackupServiceRestoreServerRequest) Execute() (BackupRestoreServerResponse, *_nethttp.Response, error) {
	return r.ApiService.BackupServiceRestoreServerExecute(r)
}

/*
BackupServiceRestoreServer Method for BackupServiceRestoreServer

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @param copyId
 @return ApiBackupServiceRestoreServerRequest
*/
func (a *BackupServiceApiService) BackupServiceRestoreServer(ctx _context.Context, id string, copyId string) ApiBackupServiceRestoreServerRequest {
	return ApiBackupServiceRestoreServerRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		copyId: copyId,
	}
}

// Execute executes the request
//  @return BackupRestoreServerResponse
func (a *BackupServiceApiService) BackupServiceRestoreServerExecute(r ApiBackupServiceRestoreServerRequest) (BackupRestoreServerResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  BackupRestoreServerResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BackupServiceApiService.BackupServiceRestoreServer")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/vps/{id}/backup/{copy_id}/server"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"copy_id"+"}", _neturl.PathEscape(parameterToString(r.copyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.backupRestoreServerRequest == nil {
		return localVarReturnValue, nil, reportError("backupRestoreServerRequest is required and must be specified")
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
	localVarPostBody = r.backupRestoreServerRequest
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
