# \DeviceAPI

All URIs are relative to *https://:443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiDeviceIdentificationGet**](DeviceAPI.md#ApiDeviceIdentificationGet) | **Get** /api/device/identification | Get the state of device identification
[**ApiDeviceIdentificationPut**](DeviceAPI.md#ApiDeviceIdentificationPut) | **Put** /api/device/identification | Set the state of device identification
[**ApiDeviceIdentityGet**](DeviceAPI.md#ApiDeviceIdentityGet) | **Get** /api/device/identity | Get the device identity
[**ApiDeviceLedsRingGet**](DeviceAPI.md#ApiDeviceLedsRingGet) | **Get** /api/device/leds/ring | Get the current led ring settings
[**ApiDeviceLedsRingPut**](DeviceAPI.md#ApiDeviceLedsRingPut) | **Put** /api/device/leds/ring | Set the current led ring brightness and colors
[**ApiDevicePowerPoeDaisychainGet**](DeviceAPI.md#ApiDevicePowerPoeDaisychainGet) | **Get** /api/device/power/poe/daisychain | Get the daisy chain and PoE state
[**ApiDeviceSiteGet**](DeviceAPI.md#ApiDeviceSiteGet) | **Get** /api/device/site | Get the device site information
[**ApiDeviceStateGet**](DeviceAPI.md#ApiDeviceStateGet) | **Get** /api/device/state | Get the device state



## ApiDeviceIdentificationGet

> ApiDeviceIdentificationGet200Response ApiDeviceIdentificationGet(ctx).Execute()

Get the state of device identification

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
    resp, r, err := apiClient.DeviceAPI.ApiDeviceIdentificationGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.ApiDeviceIdentificationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDeviceIdentificationGet`: ApiDeviceIdentificationGet200Response
    fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.ApiDeviceIdentificationGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiDeviceIdentificationGetRequest struct via the builder pattern


### Return type

[**ApiDeviceIdentificationGet200Response**](ApiDeviceIdentificationGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDeviceIdentificationPut

> ApiDeviceIdentificationPutRequest ApiDeviceIdentificationPut(ctx).ApiDeviceIdentificationPutRequest(apiDeviceIdentificationPutRequest).Execute()

Set the state of device identification

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
    apiDeviceIdentificationPutRequest := *openapiclient.NewApiDeviceIdentificationPutRequest() // ApiDeviceIdentificationPutRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceAPI.ApiDeviceIdentificationPut(context.Background()).ApiDeviceIdentificationPutRequest(apiDeviceIdentificationPutRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.ApiDeviceIdentificationPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDeviceIdentificationPut`: ApiDeviceIdentificationPutRequest
    fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.ApiDeviceIdentificationPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiDeviceIdentificationPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiDeviceIdentificationPutRequest** | [**ApiDeviceIdentificationPutRequest**](ApiDeviceIdentificationPutRequest.md) |  | 

### Return type

[**ApiDeviceIdentificationPutRequest**](ApiDeviceIdentificationPutRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDeviceIdentityGet

> ApiDeviceIdentityGet200Response ApiDeviceIdentityGet(ctx).Execute()

Get the device identity

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
    resp, r, err := apiClient.DeviceAPI.ApiDeviceIdentityGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.ApiDeviceIdentityGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDeviceIdentityGet`: ApiDeviceIdentityGet200Response
    fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.ApiDeviceIdentityGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiDeviceIdentityGetRequest struct via the builder pattern


### Return type

[**ApiDeviceIdentityGet200Response**](ApiDeviceIdentityGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDeviceLedsRingGet

> ApiDeviceLedsRingGet200Response ApiDeviceLedsRingGet(ctx).Execute()

Get the current led ring settings

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
    resp, r, err := apiClient.DeviceAPI.ApiDeviceLedsRingGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.ApiDeviceLedsRingGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDeviceLedsRingGet`: ApiDeviceLedsRingGet200Response
    fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.ApiDeviceLedsRingGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiDeviceLedsRingGetRequest struct via the builder pattern


### Return type

[**ApiDeviceLedsRingGet200Response**](ApiDeviceLedsRingGet200Response.md)

### Authorization

[basicAuthApi](../README.md#basicAuthApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDeviceLedsRingPut

> ApiDeviceLedsRingPut(ctx).ApiDeviceLedsRingPutRequest(apiDeviceLedsRingPutRequest).Execute()

Set the current led ring brightness and colors

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
    apiDeviceLedsRingPutRequest := *openapiclient.NewApiDeviceLedsRingPutRequest() // ApiDeviceLedsRingPutRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeviceAPI.ApiDeviceLedsRingPut(context.Background()).ApiDeviceLedsRingPutRequest(apiDeviceLedsRingPutRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.ApiDeviceLedsRingPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiDeviceLedsRingPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiDeviceLedsRingPutRequest** | [**ApiDeviceLedsRingPutRequest**](ApiDeviceLedsRingPutRequest.md) |  | 

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


## ApiDevicePowerPoeDaisychainGet

> ApiDevicePowerPoeDaisychainGet200Response ApiDevicePowerPoeDaisychainGet(ctx).Execute()

Get the daisy chain and PoE state

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
    resp, r, err := apiClient.DeviceAPI.ApiDevicePowerPoeDaisychainGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.ApiDevicePowerPoeDaisychainGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDevicePowerPoeDaisychainGet`: ApiDevicePowerPoeDaisychainGet200Response
    fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.ApiDevicePowerPoeDaisychainGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiDevicePowerPoeDaisychainGetRequest struct via the builder pattern


### Return type

[**ApiDevicePowerPoeDaisychainGet200Response**](ApiDevicePowerPoeDaisychainGet200Response.md)

### Authorization

[basicAuthApi](../README.md#basicAuthApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDeviceSiteGet

> ApiDeviceSiteGet200Response ApiDeviceSiteGet(ctx).Execute()

Get the device site information

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
    resp, r, err := apiClient.DeviceAPI.ApiDeviceSiteGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.ApiDeviceSiteGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDeviceSiteGet`: ApiDeviceSiteGet200Response
    fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.ApiDeviceSiteGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiDeviceSiteGetRequest struct via the builder pattern


### Return type

[**ApiDeviceSiteGet200Response**](ApiDeviceSiteGet200Response.md)

### Authorization

[basicAuthApi](../README.md#basicAuthApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDeviceStateGet

> ApiDeviceStateGet200Response ApiDeviceStateGet(ctx).Execute()

Get the device state

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
    resp, r, err := apiClient.DeviceAPI.ApiDeviceStateGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.ApiDeviceStateGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDeviceStateGet`: ApiDeviceStateGet200Response
    fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.ApiDeviceStateGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiDeviceStateGetRequest struct via the builder pattern


### Return type

[**ApiDeviceStateGet200Response**](ApiDeviceStateGet200Response.md)

### Authorization

[basicAuthApi](../README.md#basicAuthApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

