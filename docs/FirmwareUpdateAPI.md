# \FirmwareUpdateAPI

All URIs are relative to *https://:443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiFirmwareUpdateStateGet**](FirmwareUpdateAPI.md#ApiFirmwareUpdateStateGet) | **Get** /api/firmware/update/state | Get the state of a firmware update



## ApiFirmwareUpdateStateGet

> ApiFirmwareUpdateStateGet200Response ApiFirmwareUpdateStateGet(ctx).Execute()

Get the state of a firmware update

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
    resp, r, err := apiClient.FirmwareUpdateAPI.ApiFirmwareUpdateStateGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirmwareUpdateAPI.ApiFirmwareUpdateStateGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiFirmwareUpdateStateGet`: ApiFirmwareUpdateStateGet200Response
    fmt.Fprintf(os.Stdout, "Response from `FirmwareUpdateAPI.ApiFirmwareUpdateStateGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiFirmwareUpdateStateGetRequest struct via the builder pattern


### Return type

[**ApiFirmwareUpdateStateGet200Response**](ApiFirmwareUpdateStateGet200Response.md)

### Authorization

[basicAuthApi](../README.md#basicAuthApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

