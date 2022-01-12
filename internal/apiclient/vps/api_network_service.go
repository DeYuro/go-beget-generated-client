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

// NetworkServiceApiService NetworkServiceApi service
type NetworkServiceApiService service

type ApiNetworkServiceCreatePrivateNetworkRequest struct {
	ctx _context.Context
	ApiService *NetworkServiceApiService
}


func (r ApiNetworkServiceCreatePrivateNetworkRequest) Execute() (NetworkCreatePrivateNetworkResponse, *_nethttp.Response, error) {
	return r.ApiService.NetworkServiceCreatePrivateNetworkExecute(r)
}

/*
NetworkServiceCreatePrivateNetwork Method for NetworkServiceCreatePrivateNetwork

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiNetworkServiceCreatePrivateNetworkRequest
*/
func (a *NetworkServiceApiService) NetworkServiceCreatePrivateNetwork(ctx _context.Context) ApiNetworkServiceCreatePrivateNetworkRequest {
	return ApiNetworkServiceCreatePrivateNetworkRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return NetworkCreatePrivateNetworkResponse
func (a *NetworkServiceApiService) NetworkServiceCreatePrivateNetworkExecute(r ApiNetworkServiceCreatePrivateNetworkRequest) (NetworkCreatePrivateNetworkResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  NetworkCreatePrivateNetworkResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkServiceApiService.NetworkServiceCreatePrivateNetwork")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/vps/private-network"

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

type ApiNetworkServiceGetNetworkInfoRequest struct {
	ctx _context.Context
	ApiService *NetworkServiceApiService
}


func (r ApiNetworkServiceGetNetworkInfoRequest) Execute() (NetworkGetNetworkInfoResponse, *_nethttp.Response, error) {
	return r.ApiService.NetworkServiceGetNetworkInfoExecute(r)
}

/*
NetworkServiceGetNetworkInfo Method for NetworkServiceGetNetworkInfo

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiNetworkServiceGetNetworkInfoRequest
*/
func (a *NetworkServiceApiService) NetworkServiceGetNetworkInfo(ctx _context.Context) ApiNetworkServiceGetNetworkInfoRequest {
	return ApiNetworkServiceGetNetworkInfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return NetworkGetNetworkInfoResponse
func (a *NetworkServiceApiService) NetworkServiceGetNetworkInfoExecute(r ApiNetworkServiceGetNetworkInfoRequest) (NetworkGetNetworkInfoResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  NetworkGetNetworkInfoResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkServiceApiService.NetworkServiceGetNetworkInfo")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/vps/network"

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

type ApiNetworkServiceOrderIpAddressRequest struct {
	ctx _context.Context
	ApiService *NetworkServiceApiService
	networkOrderIpAddressRequest *NetworkOrderIpAddressRequest
}

func (r ApiNetworkServiceOrderIpAddressRequest) NetworkOrderIpAddressRequest(networkOrderIpAddressRequest NetworkOrderIpAddressRequest) ApiNetworkServiceOrderIpAddressRequest {
	r.networkOrderIpAddressRequest = &networkOrderIpAddressRequest
	return r
}

func (r ApiNetworkServiceOrderIpAddressRequest) Execute() (NetworkOrderIpAddressResponse, *_nethttp.Response, error) {
	return r.ApiService.NetworkServiceOrderIpAddressExecute(r)
}

/*
NetworkServiceOrderIpAddress Method for NetworkServiceOrderIpAddress

Заказ дополнительных IP-адресов для пользователя

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiNetworkServiceOrderIpAddressRequest
*/
func (a *NetworkServiceApiService) NetworkServiceOrderIpAddress(ctx _context.Context) ApiNetworkServiceOrderIpAddressRequest {
	return ApiNetworkServiceOrderIpAddressRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return NetworkOrderIpAddressResponse
func (a *NetworkServiceApiService) NetworkServiceOrderIpAddressExecute(r ApiNetworkServiceOrderIpAddressRequest) (NetworkOrderIpAddressResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  NetworkOrderIpAddressResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkServiceApiService.NetworkServiceOrderIpAddress")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/vps/network"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.networkOrderIpAddressRequest == nil {
		return localVarReturnValue, nil, reportError("networkOrderIpAddressRequest is required and must be specified")
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
	localVarPostBody = r.networkOrderIpAddressRequest
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

type ApiNetworkServiceRemoveIpAddressRequest struct {
	ctx _context.Context
	ApiService *NetworkServiceApiService
	ipAddress string
}


func (r ApiNetworkServiceRemoveIpAddressRequest) Execute() (NetworkRemoveIpAddressResponse, *_nethttp.Response, error) {
	return r.ApiService.NetworkServiceRemoveIpAddressExecute(r)
}

/*
NetworkServiceRemoveIpAddress Method for NetworkServiceRemoveIpAddress

Отмена заказа дополнительного IP адреса пользователя

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ipAddress
 @return ApiNetworkServiceRemoveIpAddressRequest
*/
func (a *NetworkServiceApiService) NetworkServiceRemoveIpAddress(ctx _context.Context, ipAddress string) ApiNetworkServiceRemoveIpAddressRequest {
	return ApiNetworkServiceRemoveIpAddressRequest{
		ApiService: a,
		ctx: ctx,
		ipAddress: ipAddress,
	}
}

// Execute executes the request
//  @return NetworkRemoveIpAddressResponse
func (a *NetworkServiceApiService) NetworkServiceRemoveIpAddressExecute(r ApiNetworkServiceRemoveIpAddressRequest) (NetworkRemoveIpAddressResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  NetworkRemoveIpAddressResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkServiceApiService.NetworkServiceRemoveIpAddress")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/vps/network/{ip_address}"
	localVarPath = strings.Replace(localVarPath, "{"+"ip_address"+"}", _neturl.PathEscape(parameterToString(r.ipAddress, "")), -1)

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

type ApiNetworkServiceSuggestPrivateAddressRequest struct {
	ctx _context.Context
	ApiService *NetworkServiceApiService
	networkId string
	networkSuggestPrivateAddressRequest *NetworkSuggestPrivateAddressRequest
}

func (r ApiNetworkServiceSuggestPrivateAddressRequest) NetworkSuggestPrivateAddressRequest(networkSuggestPrivateAddressRequest NetworkSuggestPrivateAddressRequest) ApiNetworkServiceSuggestPrivateAddressRequest {
	r.networkSuggestPrivateAddressRequest = &networkSuggestPrivateAddressRequest
	return r
}

func (r ApiNetworkServiceSuggestPrivateAddressRequest) Execute() (NetworkSuggestPrivateAddressResponse, *_nethttp.Response, error) {
	return r.ApiService.NetworkServiceSuggestPrivateAddressExecute(r)
}

/*
NetworkServiceSuggestPrivateAddress Method for NetworkServiceSuggestPrivateAddress

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return ApiNetworkServiceSuggestPrivateAddressRequest
*/
func (a *NetworkServiceApiService) NetworkServiceSuggestPrivateAddress(ctx _context.Context, networkId string) ApiNetworkServiceSuggestPrivateAddressRequest {
	return ApiNetworkServiceSuggestPrivateAddressRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return NetworkSuggestPrivateAddressResponse
func (a *NetworkServiceApiService) NetworkServiceSuggestPrivateAddressExecute(r ApiNetworkServiceSuggestPrivateAddressRequest) (NetworkSuggestPrivateAddressResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  NetworkSuggestPrivateAddressResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkServiceApiService.NetworkServiceSuggestPrivateAddress")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/vps/private-network/{network_id}/suggested-address"
	localVarPath = strings.Replace(localVarPath, "{"+"network_id"+"}", _neturl.PathEscape(parameterToString(r.networkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.networkSuggestPrivateAddressRequest == nil {
		return localVarReturnValue, nil, reportError("networkSuggestPrivateAddressRequest is required and must be specified")
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
	localVarPostBody = r.networkSuggestPrivateAddressRequest
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