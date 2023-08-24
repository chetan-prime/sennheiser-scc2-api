# \FastResourceAPI

All URIs are relative to *https://:443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAudioInputsMicrophoneBeamDirectionGet**](FastResourceAPI.md#ApiAudioInputsMicrophoneBeamDirectionGet) | **Get** /api/audio/inputs/microphone/beam/direction | Get the current direction of the beam
[**ApiAudioInputsMicrophoneLevelGet**](FastResourceAPI.md#ApiAudioInputsMicrophoneLevelGet) | **Get** /api/audio/inputs/microphone/level | Get the current microphone input level
[**ApiAudioInputsReferenceLevelGet**](FastResourceAPI.md#ApiAudioInputsReferenceLevelGet) | **Get** /api/audio/inputs/reference/level | Get the current level of the digital reference input
[**ApiAudioRoomInUseActivityLevelGet**](FastResourceAPI.md#ApiAudioRoomInUseActivityLevelGet) | **Get** /api/audio/roomInUse/activityLevel | Get the current room in use activity level



## ApiAudioInputsMicrophoneBeamDirectionGet

> ApiAudioInputsMicrophoneBeamDirectionGet200Response ApiAudioInputsMicrophoneBeamDirectionGet(ctx).Execute()

Get the current direction of the beam

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
    resp, r, err := apiClient.FastResourceAPI.ApiAudioInputsMicrophoneBeamDirectionGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FastResourceAPI.ApiAudioInputsMicrophoneBeamDirectionGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAudioInputsMicrophoneBeamDirectionGet`: ApiAudioInputsMicrophoneBeamDirectionGet200Response
    fmt.Fprintf(os.Stdout, "Response from `FastResourceAPI.ApiAudioInputsMicrophoneBeamDirectionGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiAudioInputsMicrophoneBeamDirectionGetRequest struct via the builder pattern


### Return type

[**ApiAudioInputsMicrophoneBeamDirectionGet200Response**](ApiAudioInputsMicrophoneBeamDirectionGet200Response.md)

### Authorization

[basicAuthApi](../README.md#basicAuthApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAudioInputsMicrophoneLevelGet

> ApiAudioInputsMicrophoneLevelGet200Response ApiAudioInputsMicrophoneLevelGet(ctx).Execute()

Get the current microphone input level

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
    resp, r, err := apiClient.FastResourceAPI.ApiAudioInputsMicrophoneLevelGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FastResourceAPI.ApiAudioInputsMicrophoneLevelGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAudioInputsMicrophoneLevelGet`: ApiAudioInputsMicrophoneLevelGet200Response
    fmt.Fprintf(os.Stdout, "Response from `FastResourceAPI.ApiAudioInputsMicrophoneLevelGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiAudioInputsMicrophoneLevelGetRequest struct via the builder pattern


### Return type

[**ApiAudioInputsMicrophoneLevelGet200Response**](ApiAudioInputsMicrophoneLevelGet200Response.md)

### Authorization

[basicAuthApi](../README.md#basicAuthApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAudioInputsReferenceLevelGet

> ApiAudioInputsReferenceLevelGet200Response ApiAudioInputsReferenceLevelGet(ctx).Execute()

Get the current level of the digital reference input

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
    resp, r, err := apiClient.FastResourceAPI.ApiAudioInputsReferenceLevelGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FastResourceAPI.ApiAudioInputsReferenceLevelGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAudioInputsReferenceLevelGet`: ApiAudioInputsReferenceLevelGet200Response
    fmt.Fprintf(os.Stdout, "Response from `FastResourceAPI.ApiAudioInputsReferenceLevelGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiAudioInputsReferenceLevelGetRequest struct via the builder pattern


### Return type

[**ApiAudioInputsReferenceLevelGet200Response**](ApiAudioInputsReferenceLevelGet200Response.md)

### Authorization

[basicAuthApi](../README.md#basicAuthApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAudioRoomInUseActivityLevelGet

> ApiAudioRoomInUseActivityLevelGet200Response ApiAudioRoomInUseActivityLevelGet(ctx).Execute()

Get the current room in use activity level

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
    resp, r, err := apiClient.FastResourceAPI.ApiAudioRoomInUseActivityLevelGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FastResourceAPI.ApiAudioRoomInUseActivityLevelGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAudioRoomInUseActivityLevelGet`: ApiAudioRoomInUseActivityLevelGet200Response
    fmt.Fprintf(os.Stdout, "Response from `FastResourceAPI.ApiAudioRoomInUseActivityLevelGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiAudioRoomInUseActivityLevelGetRequest struct via the builder pattern


### Return type

[**ApiAudioRoomInUseActivityLevelGet200Response**](ApiAudioRoomInUseActivityLevelGet200Response.md)

### Authorization

[basicAuthApi](../README.md#basicAuthApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

