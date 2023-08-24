# \SSCAPI

All URIs are relative to *https://:443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiSscSchemaGet**](SSCAPI.md#ApiSscSchemaGet) | **Get** /api/ssc/schema | Get the address tree
[**ApiSscStateSubscriptionsGet**](SSCAPI.md#ApiSscStateSubscriptionsGet) | **Get** /api/ssc/state/subscriptions | Start a subscription
[**ApiSscStateSubscriptionsSessionUUIDAddPut**](SSCAPI.md#ApiSscStateSubscriptionsSessionUUIDAddPut) | **Put** /api/ssc/state/subscriptions/{sessionUUID}/add | Add resource(s) to the subscription list
[**ApiSscStateSubscriptionsSessionUUIDDelete**](SSCAPI.md#ApiSscStateSubscriptionsSessionUUIDDelete) | **Delete** /api/ssc/state/subscriptions/{sessionUUID} | End an existing subscription
[**ApiSscStateSubscriptionsSessionUUIDGet**](SSCAPI.md#ApiSscStateSubscriptionsSessionUUIDGet) | **Get** /api/ssc/state/subscriptions/{sessionUUID} | Get the subscription list
[**ApiSscStateSubscriptionsSessionUUIDPut**](SSCAPI.md#ApiSscStateSubscriptionsSessionUUIDPut) | **Put** /api/ssc/state/subscriptions/{sessionUUID} | Set or change the list of subscriptions associated with the sessionUUID
[**ApiSscStateSubscriptionsSessionUUIDRemovePut**](SSCAPI.md#ApiSscStateSubscriptionsSessionUUIDRemovePut) | **Put** /api/ssc/state/subscriptions/{sessionUUID}/remove | Remove resource(s) from the subscription list
[**ApiSscVersionGet**](SSCAPI.md#ApiSscVersionGet) | **Get** /api/ssc/version | Get the schema version



## ApiSscSchemaGet

> string ApiSscSchemaGet(ctx).Execute()

Get the address tree

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "//"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SSCAPI.ApiSscSchemaGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSCAPI.ApiSscSchemaGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSscSchemaGet`: string
    fmt.Fprintf(os.Stdout, "Response from `SSCAPI.ApiSscSchemaGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiSscSchemaGetRequest struct via the builder pattern


### Return type

**string**

### Authorization

[basicAuthApi](../README.md#basicAuthApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/openapi+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiSscStateSubscriptionsGet

> string ApiSscStateSubscriptionsGet(ctx).Execute()

Start a subscription



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "//"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SSCAPI.ApiSscStateSubscriptionsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSCAPI.ApiSscStateSubscriptionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSscStateSubscriptionsGet`: string
    fmt.Fprintf(os.Stdout, "Response from `SSCAPI.ApiSscStateSubscriptionsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiSscStateSubscriptionsGetRequest struct via the builder pattern


### Return type

**string**

### Authorization

[basicAuthApi](../README.md#basicAuthApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiSscStateSubscriptionsSessionUUIDAddPut

> ApiSscStateSubscriptionsSessionUUIDAddPut(ctx, sessionUUID).RequestBody(requestBody).Execute()

Add resource(s) to the subscription list



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "//"
)

func main() {
    sessionUUID := "sessionUUID_example" // string | 
    requestBody := []string{"/device/site"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SSCAPI.ApiSscStateSubscriptionsSessionUUIDAddPut(context.Background(), sessionUUID).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSCAPI.ApiSscStateSubscriptionsSessionUUIDAddPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSscStateSubscriptionsSessionUUIDAddPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** |  | 

### Return type

 (empty response body)

### Authorization

[basicAuthApi](../README.md#basicAuthApi)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiSscStateSubscriptionsSessionUUIDDelete

> ApiSscStateSubscriptionsSessionUUIDDelete(ctx, sessionUUID).Execute()

End an existing subscription



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "//"
)

func main() {
    sessionUUID := "31875a94-29e6-4fb0-ab4d-7f0bbd6e1bc8" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SSCAPI.ApiSscStateSubscriptionsSessionUUIDDelete(context.Background(), sessionUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSCAPI.ApiSscStateSubscriptionsSessionUUIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSscStateSubscriptionsSessionUUIDDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basicAuthApi](../README.md#basicAuthApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiSscStateSubscriptionsSessionUUIDGet

> []string ApiSscStateSubscriptionsSessionUUIDGet(ctx, sessionUUID).Execute()

Get the subscription list



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "//"
)

func main() {
    sessionUUID := "31875a94-29e6-4fb0-ab4d-7f0bbd6e1bc8" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SSCAPI.ApiSscStateSubscriptionsSessionUUIDGet(context.Background(), sessionUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSCAPI.ApiSscStateSubscriptionsSessionUUIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSscStateSubscriptionsSessionUUIDGet`: []string
    fmt.Fprintf(os.Stdout, "Response from `SSCAPI.ApiSscStateSubscriptionsSessionUUIDGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSscStateSubscriptionsSessionUUIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[basicAuthApi](../README.md#basicAuthApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiSscStateSubscriptionsSessionUUIDPut

> ApiSscStateSubscriptionsSessionUUIDPut(ctx, sessionUUID).RequestBody(requestBody).Execute()

Set or change the list of subscriptions associated with the sessionUUID

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "//"
)

func main() {
    sessionUUID := "31875a94-29e6-4fb0-ab4d-7f0bbd6e1bc8" // string | 
    requestBody := []string{"/device/site"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SSCAPI.ApiSscStateSubscriptionsSessionUUIDPut(context.Background(), sessionUUID).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSCAPI.ApiSscStateSubscriptionsSessionUUIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSscStateSubscriptionsSessionUUIDPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** |  | 

### Return type

 (empty response body)

### Authorization

[basicAuthApi](../README.md#basicAuthApi)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiSscStateSubscriptionsSessionUUIDRemovePut

> ApiSscStateSubscriptionsSessionUUIDRemovePut(ctx, sessionUUID).RequestBody(requestBody).Execute()

Remove resource(s) from the subscription list



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "//"
)

func main() {
    sessionUUID := "sessionUUID_example" // string | 
    requestBody := []string{"/api/device/site"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SSCAPI.ApiSscStateSubscriptionsSessionUUIDRemovePut(context.Background(), sessionUUID).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSCAPI.ApiSscStateSubscriptionsSessionUUIDRemovePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSscStateSubscriptionsSessionUUIDRemovePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** |  | 

### Return type

 (empty response body)

### Authorization

[basicAuthApi](../README.md#basicAuthApi)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiSscVersionGet

> ApiSscVersionGet200Response ApiSscVersionGet(ctx).Execute()

Get the schema version

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "//"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SSCAPI.ApiSscVersionGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSCAPI.ApiSscVersionGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSscVersionGet`: ApiSscVersionGet200Response
    fmt.Fprintf(os.Stdout, "Response from `SSCAPI.ApiSscVersionGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiSscVersionGetRequest struct via the builder pattern


### Return type

[**ApiSscVersionGet200Response**](ApiSscVersionGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

