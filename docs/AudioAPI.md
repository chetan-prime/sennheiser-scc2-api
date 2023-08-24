# \AudioAPI

All URIs are relative to *https://:443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAudioEqualizerGet**](AudioAPI.md#ApiAudioEqualizerGet) | **Get** /api/audio/equalizer | Get the current equalizer settings
[**ApiAudioEqualizerPut**](AudioAPI.md#ApiAudioEqualizerPut) | **Put** /api/audio/equalizer | Set the current equalizer settings
[**ApiAudioInputsDanteReferenceGet**](AudioAPI.md#ApiAudioInputsDanteReferenceGet) | **Get** /api/audio/inputs/dante/reference | Get settings of the reference input
[**ApiAudioInputsDanteReferencePut**](AudioAPI.md#ApiAudioInputsDanteReferencePut) | **Put** /api/audio/inputs/dante/reference | Set reference input settings
[**ApiAudioInputsMicrophoneBeamDirectionGet**](AudioAPI.md#ApiAudioInputsMicrophoneBeamDirectionGet) | **Get** /api/audio/inputs/microphone/beam/direction | Get the current direction of the beam
[**ApiAudioInputsMicrophoneBeamGet**](AudioAPI.md#ApiAudioInputsMicrophoneBeamGet) | **Get** /api/audio/inputs/microphone/beam | Get the current beam settings
[**ApiAudioInputsMicrophoneBeamPut**](AudioAPI.md#ApiAudioInputsMicrophoneBeamPut) | **Put** /api/audio/inputs/microphone/beam | Set the current beam settings
[**ApiAudioInputsMicrophoneExclusionZonesGet**](AudioAPI.md#ApiAudioInputsMicrophoneExclusionZonesGet) | **Get** /api/audio/inputs/microphone/exclusionZones | Get the supported exclusion zone ids
[**ApiAudioInputsMicrophoneExclusionZonesIdGet**](AudioAPI.md#ApiAudioInputsMicrophoneExclusionZonesIdGet) | **Get** /api/audio/inputs/microphone/exclusionZones/{id} | Get the current exclusion zone settings of zone number &#x60;id&#x60;
[**ApiAudioInputsMicrophoneExclusionZonesIdPut**](AudioAPI.md#ApiAudioInputsMicrophoneExclusionZonesIdPut) | **Put** /api/audio/inputs/microphone/exclusionZones/{id} | Set the current exclusion zone settings of zone number &#x60;id&#x60;
[**ApiAudioInputsMicrophoneLevelGet**](AudioAPI.md#ApiAudioInputsMicrophoneLevelGet) | **Get** /api/audio/inputs/microphone/level | Get the current microphone input level
[**ApiAudioInputsMicrophonePriorityZonesGet**](AudioAPI.md#ApiAudioInputsMicrophonePriorityZonesGet) | **Get** /api/audio/inputs/microphone/priorityZones | Get the supported priority zone ids
[**ApiAudioInputsMicrophonePriorityZonesIdGet**](AudioAPI.md#ApiAudioInputsMicrophonePriorityZonesIdGet) | **Get** /api/audio/inputs/microphone/priorityZones/{id} | Get the current priority zone settings of zone number &#x60;id&#x60;
[**ApiAudioInputsMicrophonePriorityZonesIdPut**](AudioAPI.md#ApiAudioInputsMicrophonePriorityZonesIdPut) | **Put** /api/audio/inputs/microphone/priorityZones/{id} | Set the current priority zone settings of zone number &#x60;id&#x60;
[**ApiAudioInputsReferenceLevelGet**](AudioAPI.md#ApiAudioInputsReferenceLevelGet) | **Get** /api/audio/inputs/reference/level | Get the current level of the digital reference input
[**ApiAudioNoiseGateGet**](AudioAPI.md#ApiAudioNoiseGateGet) | **Get** /api/audio/noiseGate | Get the current noise gate settings
[**ApiAudioNoiseGatePut**](AudioAPI.md#ApiAudioNoiseGatePut) | **Put** /api/audio/noiseGate | Set the current noise gate settings
[**ApiAudioOutputsAnalogGet**](AudioAPI.md#ApiAudioOutputsAnalogGet) | **Get** /api/audio/outputs/analog | Get the current settings of the analog output
[**ApiAudioOutputsAnalogPut**](AudioAPI.md#ApiAudioOutputsAnalogPut) | **Put** /api/audio/outputs/analog | Set the analog output settings
[**ApiAudioOutputsDanteFarEndGet**](AudioAPI.md#ApiAudioOutputsDanteFarEndGet) | **Get** /api/audio/outputs/dante/farEnd | Get the current settings of the far end output
[**ApiAudioOutputsDanteFarEndPut**](AudioAPI.md#ApiAudioOutputsDanteFarEndPut) | **Put** /api/audio/outputs/dante/farEnd | Set the far end output settings
[**ApiAudioOutputsDanteLocalGet**](AudioAPI.md#ApiAudioOutputsDanteLocalGet) | **Get** /api/audio/outputs/dante/local | Get the current settings of the local output
[**ApiAudioOutputsDanteLocalPut**](AudioAPI.md#ApiAudioOutputsDanteLocalPut) | **Put** /api/audio/outputs/dante/local | Set the local output settings
[**ApiAudioOutputsGlobalMuteGet**](AudioAPI.md#ApiAudioOutputsGlobalMuteGet) | **Get** /api/audio/outputs/global/mute | Get the mute status of the device
[**ApiAudioOutputsGlobalMutePut**](AudioAPI.md#ApiAudioOutputsGlobalMutePut) | **Put** /api/audio/outputs/global/mute | Mute all audio outputs
[**ApiAudioRoomInUseActivityLevelGet**](AudioAPI.md#ApiAudioRoomInUseActivityLevelGet) | **Get** /api/audio/roomInUse/activityLevel | Get the current room in use activity level
[**ApiAudioRoomInUseConfigGet**](AudioAPI.md#ApiAudioRoomInUseConfigGet) | **Get** /api/audio/roomInUse/config | Get the configuration of the room in use feature
[**ApiAudioRoomInUseGet**](AudioAPI.md#ApiAudioRoomInUseGet) | **Get** /api/audio/roomInUse | Get the current room in use state
[**ApiAudioVoiceLiftGet**](AudioAPI.md#ApiAudioVoiceLiftGet) | **Get** /api/audio/voiceLift | Get the current voice lift settings
[**ApiAudioVoiceLiftPut**](AudioAPI.md#ApiAudioVoiceLiftPut) | **Put** /api/audio/voiceLift | Set the current voice lift settings



## ApiAudioEqualizerGet

> ApiAudioEqualizerGet200Response ApiAudioEqualizerGet(ctx).Execute()

Get the current equalizer settings

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
    resp, r, err := apiClient.AudioAPI.ApiAudioEqualizerGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.ApiAudioEqualizerGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAudioEqualizerGet`: ApiAudioEqualizerGet200Response
    fmt.Fprintf(os.Stdout, "Response from `AudioAPI.ApiAudioEqualizerGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiAudioEqualizerGetRequest struct via the builder pattern


### Return type

[**ApiAudioEqualizerGet200Response**](ApiAudioEqualizerGet200Response.md)

### Authorization

[basicAuthApi](../README.md#basicAuthApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAudioEqualizerPut

> ApiAudioEqualizerPut(ctx).ApiAudioEqualizerPutRequest(apiAudioEqualizerPutRequest).Execute()

Set the current equalizer settings



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
    apiAudioEqualizerPutRequest := *openapiclient.NewApiAudioEqualizerPutRequest() // ApiAudioEqualizerPutRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AudioAPI.ApiAudioEqualizerPut(context.Background()).ApiAudioEqualizerPutRequest(apiAudioEqualizerPutRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.ApiAudioEqualizerPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAudioEqualizerPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiAudioEqualizerPutRequest** | [**ApiAudioEqualizerPutRequest**](ApiAudioEqualizerPutRequest.md) |  | 

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


## ApiAudioInputsDanteReferenceGet

> ApiAudioInputsDanteReferenceGet200Response ApiAudioInputsDanteReferenceGet(ctx).Execute()

Get settings of the reference input

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
    resp, r, err := apiClient.AudioAPI.ApiAudioInputsDanteReferenceGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.ApiAudioInputsDanteReferenceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAudioInputsDanteReferenceGet`: ApiAudioInputsDanteReferenceGet200Response
    fmt.Fprintf(os.Stdout, "Response from `AudioAPI.ApiAudioInputsDanteReferenceGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiAudioInputsDanteReferenceGetRequest struct via the builder pattern


### Return type

[**ApiAudioInputsDanteReferenceGet200Response**](ApiAudioInputsDanteReferenceGet200Response.md)

### Authorization

[basicAuthApi](../README.md#basicAuthApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAudioInputsDanteReferencePut

> ApiAudioInputsDanteReferencePut(ctx).ApiAudioInputsDanteReferencePutRequest(apiAudioInputsDanteReferencePutRequest).Execute()

Set reference input settings

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
    apiAudioInputsDanteReferencePutRequest := *openapiclient.NewApiAudioInputsDanteReferencePutRequest() // ApiAudioInputsDanteReferencePutRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AudioAPI.ApiAudioInputsDanteReferencePut(context.Background()).ApiAudioInputsDanteReferencePutRequest(apiAudioInputsDanteReferencePutRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.ApiAudioInputsDanteReferencePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAudioInputsDanteReferencePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiAudioInputsDanteReferencePutRequest** | [**ApiAudioInputsDanteReferencePutRequest**](ApiAudioInputsDanteReferencePutRequest.md) |  | 

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
    resp, r, err := apiClient.AudioAPI.ApiAudioInputsMicrophoneBeamDirectionGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.ApiAudioInputsMicrophoneBeamDirectionGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAudioInputsMicrophoneBeamDirectionGet`: ApiAudioInputsMicrophoneBeamDirectionGet200Response
    fmt.Fprintf(os.Stdout, "Response from `AudioAPI.ApiAudioInputsMicrophoneBeamDirectionGet`: %v\n", resp)
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


## ApiAudioInputsMicrophoneBeamGet

> ApiAudioInputsMicrophoneBeamGet200Response ApiAudioInputsMicrophoneBeamGet(ctx).Execute()

Get the current beam settings

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
    resp, r, err := apiClient.AudioAPI.ApiAudioInputsMicrophoneBeamGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.ApiAudioInputsMicrophoneBeamGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAudioInputsMicrophoneBeamGet`: ApiAudioInputsMicrophoneBeamGet200Response
    fmt.Fprintf(os.Stdout, "Response from `AudioAPI.ApiAudioInputsMicrophoneBeamGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiAudioInputsMicrophoneBeamGetRequest struct via the builder pattern


### Return type

[**ApiAudioInputsMicrophoneBeamGet200Response**](ApiAudioInputsMicrophoneBeamGet200Response.md)

### Authorization

[basicAuthApi](../README.md#basicAuthApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAudioInputsMicrophoneBeamPut

> ApiAudioInputsMicrophoneBeamPut(ctx).ApiAudioInputsMicrophoneBeamPutRequest(apiAudioInputsMicrophoneBeamPutRequest).Execute()

Set the current beam settings

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
    apiAudioInputsMicrophoneBeamPutRequest := *openapiclient.NewApiAudioInputsMicrophoneBeamPutRequest() // ApiAudioInputsMicrophoneBeamPutRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AudioAPI.ApiAudioInputsMicrophoneBeamPut(context.Background()).ApiAudioInputsMicrophoneBeamPutRequest(apiAudioInputsMicrophoneBeamPutRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.ApiAudioInputsMicrophoneBeamPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAudioInputsMicrophoneBeamPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiAudioInputsMicrophoneBeamPutRequest** | [**ApiAudioInputsMicrophoneBeamPutRequest**](ApiAudioInputsMicrophoneBeamPutRequest.md) |  | 

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


## ApiAudioInputsMicrophoneExclusionZonesGet

> []ApiAudioInputsMicrophoneExclusionZonesGet200ResponseInner ApiAudioInputsMicrophoneExclusionZonesGet(ctx).Execute()

Get the supported exclusion zone ids

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
    resp, r, err := apiClient.AudioAPI.ApiAudioInputsMicrophoneExclusionZonesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.ApiAudioInputsMicrophoneExclusionZonesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAudioInputsMicrophoneExclusionZonesGet`: []ApiAudioInputsMicrophoneExclusionZonesGet200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `AudioAPI.ApiAudioInputsMicrophoneExclusionZonesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiAudioInputsMicrophoneExclusionZonesGetRequest struct via the builder pattern


### Return type

[**[]ApiAudioInputsMicrophoneExclusionZonesGet200ResponseInner**](ApiAudioInputsMicrophoneExclusionZonesGet200ResponseInner.md)

### Authorization

[basicAuthApi](../README.md#basicAuthApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAudioInputsMicrophoneExclusionZonesIdGet

> ApiAudioInputsMicrophoneExclusionZonesIdGet200Response ApiAudioInputsMicrophoneExclusionZonesIdGet(ctx, id).Execute()

Get the current exclusion zone settings of zone number `id`

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
    id := int32(0) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AudioAPI.ApiAudioInputsMicrophoneExclusionZonesIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.ApiAudioInputsMicrophoneExclusionZonesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAudioInputsMicrophoneExclusionZonesIdGet`: ApiAudioInputsMicrophoneExclusionZonesIdGet200Response
    fmt.Fprintf(os.Stdout, "Response from `AudioAPI.ApiAudioInputsMicrophoneExclusionZonesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAudioInputsMicrophoneExclusionZonesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiAudioInputsMicrophoneExclusionZonesIdGet200Response**](ApiAudioInputsMicrophoneExclusionZonesIdGet200Response.md)

### Authorization

[basicAuthApi](../README.md#basicAuthApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAudioInputsMicrophoneExclusionZonesIdPut

> ApiAudioInputsMicrophoneExclusionZonesIdPut(ctx, id).ApiAudioInputsMicrophoneExclusionZonesIdPutRequest(apiAudioInputsMicrophoneExclusionZonesIdPutRequest).Execute()

Set the current exclusion zone settings of zone number `id`

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
    id := int32(0) // int32 | 
    apiAudioInputsMicrophoneExclusionZonesIdPutRequest := *openapiclient.NewApiAudioInputsMicrophoneExclusionZonesIdPutRequest() // ApiAudioInputsMicrophoneExclusionZonesIdPutRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AudioAPI.ApiAudioInputsMicrophoneExclusionZonesIdPut(context.Background(), id).ApiAudioInputsMicrophoneExclusionZonesIdPutRequest(apiAudioInputsMicrophoneExclusionZonesIdPutRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.ApiAudioInputsMicrophoneExclusionZonesIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAudioInputsMicrophoneExclusionZonesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiAudioInputsMicrophoneExclusionZonesIdPutRequest** | [**ApiAudioInputsMicrophoneExclusionZonesIdPutRequest**](ApiAudioInputsMicrophoneExclusionZonesIdPutRequest.md) |  | 

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
    resp, r, err := apiClient.AudioAPI.ApiAudioInputsMicrophoneLevelGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.ApiAudioInputsMicrophoneLevelGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAudioInputsMicrophoneLevelGet`: ApiAudioInputsMicrophoneLevelGet200Response
    fmt.Fprintf(os.Stdout, "Response from `AudioAPI.ApiAudioInputsMicrophoneLevelGet`: %v\n", resp)
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


## ApiAudioInputsMicrophonePriorityZonesGet

> []ApiAudioInputsMicrophonePriorityZonesGet200ResponseInner ApiAudioInputsMicrophonePriorityZonesGet(ctx).Execute()

Get the supported priority zone ids

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
    resp, r, err := apiClient.AudioAPI.ApiAudioInputsMicrophonePriorityZonesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.ApiAudioInputsMicrophonePriorityZonesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAudioInputsMicrophonePriorityZonesGet`: []ApiAudioInputsMicrophonePriorityZonesGet200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `AudioAPI.ApiAudioInputsMicrophonePriorityZonesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiAudioInputsMicrophonePriorityZonesGetRequest struct via the builder pattern


### Return type

[**[]ApiAudioInputsMicrophonePriorityZonesGet200ResponseInner**](ApiAudioInputsMicrophonePriorityZonesGet200ResponseInner.md)

### Authorization

[basicAuthApi](../README.md#basicAuthApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAudioInputsMicrophonePriorityZonesIdGet

> ApiAudioInputsMicrophonePriorityZonesIdGet200Response ApiAudioInputsMicrophonePriorityZonesIdGet(ctx, id).Execute()

Get the current priority zone settings of zone number `id`

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
    id := int32(0) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AudioAPI.ApiAudioInputsMicrophonePriorityZonesIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.ApiAudioInputsMicrophonePriorityZonesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAudioInputsMicrophonePriorityZonesIdGet`: ApiAudioInputsMicrophonePriorityZonesIdGet200Response
    fmt.Fprintf(os.Stdout, "Response from `AudioAPI.ApiAudioInputsMicrophonePriorityZonesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAudioInputsMicrophonePriorityZonesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiAudioInputsMicrophonePriorityZonesIdGet200Response**](ApiAudioInputsMicrophonePriorityZonesIdGet200Response.md)

### Authorization

[basicAuthApi](../README.md#basicAuthApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAudioInputsMicrophonePriorityZonesIdPut

> ApiAudioInputsMicrophonePriorityZonesIdPut(ctx, id).ApiAudioInputsMicrophonePriorityZonesIdPutRequest(apiAudioInputsMicrophonePriorityZonesIdPutRequest).Execute()

Set the current priority zone settings of zone number `id`

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
    id := int32(0) // int32 | 
    apiAudioInputsMicrophonePriorityZonesIdPutRequest := *openapiclient.NewApiAudioInputsMicrophonePriorityZonesIdPutRequest() // ApiAudioInputsMicrophonePriorityZonesIdPutRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AudioAPI.ApiAudioInputsMicrophonePriorityZonesIdPut(context.Background(), id).ApiAudioInputsMicrophonePriorityZonesIdPutRequest(apiAudioInputsMicrophonePriorityZonesIdPutRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.ApiAudioInputsMicrophonePriorityZonesIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAudioInputsMicrophonePriorityZonesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiAudioInputsMicrophonePriorityZonesIdPutRequest** | [**ApiAudioInputsMicrophonePriorityZonesIdPutRequest**](ApiAudioInputsMicrophonePriorityZonesIdPutRequest.md) |  | 

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
    resp, r, err := apiClient.AudioAPI.ApiAudioInputsReferenceLevelGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.ApiAudioInputsReferenceLevelGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAudioInputsReferenceLevelGet`: ApiAudioInputsReferenceLevelGet200Response
    fmt.Fprintf(os.Stdout, "Response from `AudioAPI.ApiAudioInputsReferenceLevelGet`: %v\n", resp)
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


## ApiAudioNoiseGateGet

> ApiAudioNoiseGateGet200Response ApiAudioNoiseGateGet(ctx).Execute()

Get the current noise gate settings

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
    resp, r, err := apiClient.AudioAPI.ApiAudioNoiseGateGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.ApiAudioNoiseGateGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAudioNoiseGateGet`: ApiAudioNoiseGateGet200Response
    fmt.Fprintf(os.Stdout, "Response from `AudioAPI.ApiAudioNoiseGateGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiAudioNoiseGateGetRequest struct via the builder pattern


### Return type

[**ApiAudioNoiseGateGet200Response**](ApiAudioNoiseGateGet200Response.md)

### Authorization

[basicAuthApi](../README.md#basicAuthApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAudioNoiseGatePut

> ApiAudioNoiseGatePut(ctx).ApiAudioNoiseGatePutRequest(apiAudioNoiseGatePutRequest).Execute()

Set the current noise gate settings



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
    apiAudioNoiseGatePutRequest := *openapiclient.NewApiAudioNoiseGatePutRequest() // ApiAudioNoiseGatePutRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AudioAPI.ApiAudioNoiseGatePut(context.Background()).ApiAudioNoiseGatePutRequest(apiAudioNoiseGatePutRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.ApiAudioNoiseGatePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAudioNoiseGatePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiAudioNoiseGatePutRequest** | [**ApiAudioNoiseGatePutRequest**](ApiAudioNoiseGatePutRequest.md) |  | 

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


## ApiAudioOutputsAnalogGet

> ApiAudioOutputsAnalogGet200Response ApiAudioOutputsAnalogGet(ctx).Execute()

Get the current settings of the analog output

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
    resp, r, err := apiClient.AudioAPI.ApiAudioOutputsAnalogGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.ApiAudioOutputsAnalogGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAudioOutputsAnalogGet`: ApiAudioOutputsAnalogGet200Response
    fmt.Fprintf(os.Stdout, "Response from `AudioAPI.ApiAudioOutputsAnalogGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiAudioOutputsAnalogGetRequest struct via the builder pattern


### Return type

[**ApiAudioOutputsAnalogGet200Response**](ApiAudioOutputsAnalogGet200Response.md)

### Authorization

[basicAuthApi](../README.md#basicAuthApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAudioOutputsAnalogPut

> ApiAudioOutputsAnalogPut(ctx).ApiAudioOutputsAnalogPutRequest(apiAudioOutputsAnalogPutRequest).Execute()

Set the analog output settings

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
    apiAudioOutputsAnalogPutRequest := *openapiclient.NewApiAudioOutputsAnalogPutRequest() // ApiAudioOutputsAnalogPutRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AudioAPI.ApiAudioOutputsAnalogPut(context.Background()).ApiAudioOutputsAnalogPutRequest(apiAudioOutputsAnalogPutRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.ApiAudioOutputsAnalogPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAudioOutputsAnalogPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiAudioOutputsAnalogPutRequest** | [**ApiAudioOutputsAnalogPutRequest**](ApiAudioOutputsAnalogPutRequest.md) |  | 

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


## ApiAudioOutputsDanteFarEndGet

> ApiAudioOutputsDanteFarEndGet200Response ApiAudioOutputsDanteFarEndGet(ctx).Execute()

Get the current settings of the far end output

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
    resp, r, err := apiClient.AudioAPI.ApiAudioOutputsDanteFarEndGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.ApiAudioOutputsDanteFarEndGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAudioOutputsDanteFarEndGet`: ApiAudioOutputsDanteFarEndGet200Response
    fmt.Fprintf(os.Stdout, "Response from `AudioAPI.ApiAudioOutputsDanteFarEndGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiAudioOutputsDanteFarEndGetRequest struct via the builder pattern


### Return type

[**ApiAudioOutputsDanteFarEndGet200Response**](ApiAudioOutputsDanteFarEndGet200Response.md)

### Authorization

[basicAuthApi](../README.md#basicAuthApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAudioOutputsDanteFarEndPut

> ApiAudioOutputsDanteFarEndPut(ctx).ApiAudioOutputsDanteFarEndPutRequest(apiAudioOutputsDanteFarEndPutRequest).Execute()

Set the far end output settings

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
    apiAudioOutputsDanteFarEndPutRequest := *openapiclient.NewApiAudioOutputsDanteFarEndPutRequest() // ApiAudioOutputsDanteFarEndPutRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AudioAPI.ApiAudioOutputsDanteFarEndPut(context.Background()).ApiAudioOutputsDanteFarEndPutRequest(apiAudioOutputsDanteFarEndPutRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.ApiAudioOutputsDanteFarEndPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAudioOutputsDanteFarEndPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiAudioOutputsDanteFarEndPutRequest** | [**ApiAudioOutputsDanteFarEndPutRequest**](ApiAudioOutputsDanteFarEndPutRequest.md) |  | 

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


## ApiAudioOutputsDanteLocalGet

> ApiAudioOutputsDanteLocalGet200Response ApiAudioOutputsDanteLocalGet(ctx).Execute()

Get the current settings of the local output

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
    resp, r, err := apiClient.AudioAPI.ApiAudioOutputsDanteLocalGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.ApiAudioOutputsDanteLocalGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAudioOutputsDanteLocalGet`: ApiAudioOutputsDanteLocalGet200Response
    fmt.Fprintf(os.Stdout, "Response from `AudioAPI.ApiAudioOutputsDanteLocalGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiAudioOutputsDanteLocalGetRequest struct via the builder pattern


### Return type

[**ApiAudioOutputsDanteLocalGet200Response**](ApiAudioOutputsDanteLocalGet200Response.md)

### Authorization

[basicAuthApi](../README.md#basicAuthApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAudioOutputsDanteLocalPut

> ApiAudioOutputsDanteLocalPut(ctx).ApiAudioOutputsDanteLocalPutRequest(apiAudioOutputsDanteLocalPutRequest).Execute()

Set the local output settings

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
    apiAudioOutputsDanteLocalPutRequest := *openapiclient.NewApiAudioOutputsDanteLocalPutRequest() // ApiAudioOutputsDanteLocalPutRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AudioAPI.ApiAudioOutputsDanteLocalPut(context.Background()).ApiAudioOutputsDanteLocalPutRequest(apiAudioOutputsDanteLocalPutRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.ApiAudioOutputsDanteLocalPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAudioOutputsDanteLocalPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiAudioOutputsDanteLocalPutRequest** | [**ApiAudioOutputsDanteLocalPutRequest**](ApiAudioOutputsDanteLocalPutRequest.md) |  | 

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


## ApiAudioOutputsGlobalMuteGet

> ApiAudioOutputsGlobalMuteGet200Response ApiAudioOutputsGlobalMuteGet(ctx).Execute()

Get the mute status of the device

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
    resp, r, err := apiClient.AudioAPI.ApiAudioOutputsGlobalMuteGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.ApiAudioOutputsGlobalMuteGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAudioOutputsGlobalMuteGet`: ApiAudioOutputsGlobalMuteGet200Response
    fmt.Fprintf(os.Stdout, "Response from `AudioAPI.ApiAudioOutputsGlobalMuteGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiAudioOutputsGlobalMuteGetRequest struct via the builder pattern


### Return type

[**ApiAudioOutputsGlobalMuteGet200Response**](ApiAudioOutputsGlobalMuteGet200Response.md)

### Authorization

[basicAuthApi](../README.md#basicAuthApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAudioOutputsGlobalMutePut

> ApiAudioOutputsGlobalMutePut(ctx).ApiAudioOutputsGlobalMutePutRequest(apiAudioOutputsGlobalMutePutRequest).Execute()

Mute all audio outputs

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
    apiAudioOutputsGlobalMutePutRequest := *openapiclient.NewApiAudioOutputsGlobalMutePutRequest() // ApiAudioOutputsGlobalMutePutRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AudioAPI.ApiAudioOutputsGlobalMutePut(context.Background()).ApiAudioOutputsGlobalMutePutRequest(apiAudioOutputsGlobalMutePutRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.ApiAudioOutputsGlobalMutePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAudioOutputsGlobalMutePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiAudioOutputsGlobalMutePutRequest** | [**ApiAudioOutputsGlobalMutePutRequest**](ApiAudioOutputsGlobalMutePutRequest.md) |  | 

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
    resp, r, err := apiClient.AudioAPI.ApiAudioRoomInUseActivityLevelGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.ApiAudioRoomInUseActivityLevelGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAudioRoomInUseActivityLevelGet`: ApiAudioRoomInUseActivityLevelGet200Response
    fmt.Fprintf(os.Stdout, "Response from `AudioAPI.ApiAudioRoomInUseActivityLevelGet`: %v\n", resp)
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


## ApiAudioRoomInUseConfigGet

> ApiAudioRoomInUseConfigGet200Response ApiAudioRoomInUseConfigGet(ctx).Execute()

Get the configuration of the room in use feature

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
    resp, r, err := apiClient.AudioAPI.ApiAudioRoomInUseConfigGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.ApiAudioRoomInUseConfigGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAudioRoomInUseConfigGet`: ApiAudioRoomInUseConfigGet200Response
    fmt.Fprintf(os.Stdout, "Response from `AudioAPI.ApiAudioRoomInUseConfigGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiAudioRoomInUseConfigGetRequest struct via the builder pattern


### Return type

[**ApiAudioRoomInUseConfigGet200Response**](ApiAudioRoomInUseConfigGet200Response.md)

### Authorization

[basicAuthApi](../README.md#basicAuthApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAudioRoomInUseGet

> ApiAudioRoomInUseGet200Response ApiAudioRoomInUseGet(ctx).Execute()

Get the current room in use state

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
    resp, r, err := apiClient.AudioAPI.ApiAudioRoomInUseGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.ApiAudioRoomInUseGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAudioRoomInUseGet`: ApiAudioRoomInUseGet200Response
    fmt.Fprintf(os.Stdout, "Response from `AudioAPI.ApiAudioRoomInUseGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiAudioRoomInUseGetRequest struct via the builder pattern


### Return type

[**ApiAudioRoomInUseGet200Response**](ApiAudioRoomInUseGet200Response.md)

### Authorization

[basicAuthApi](../README.md#basicAuthApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAudioVoiceLiftGet

> ApiAudioVoiceLiftGet200Response ApiAudioVoiceLiftGet(ctx).Execute()

Get the current voice lift settings

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
    resp, r, err := apiClient.AudioAPI.ApiAudioVoiceLiftGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.ApiAudioVoiceLiftGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAudioVoiceLiftGet`: ApiAudioVoiceLiftGet200Response
    fmt.Fprintf(os.Stdout, "Response from `AudioAPI.ApiAudioVoiceLiftGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiAudioVoiceLiftGetRequest struct via the builder pattern


### Return type

[**ApiAudioVoiceLiftGet200Response**](ApiAudioVoiceLiftGet200Response.md)

### Authorization

[basicAuthApi](../README.md#basicAuthApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAudioVoiceLiftPut

> ApiAudioVoiceLiftPut(ctx).ApiAudioVoiceLiftPutRequest(apiAudioVoiceLiftPutRequest).Execute()

Set the current voice lift settings



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
    apiAudioVoiceLiftPutRequest := *openapiclient.NewApiAudioVoiceLiftPutRequest() // ApiAudioVoiceLiftPutRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AudioAPI.ApiAudioVoiceLiftPut(context.Background()).ApiAudioVoiceLiftPutRequest(apiAudioVoiceLiftPutRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.ApiAudioVoiceLiftPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAudioVoiceLiftPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiAudioVoiceLiftPutRequest** | [**ApiAudioVoiceLiftPutRequest**](ApiAudioVoiceLiftPutRequest.md) |  | 

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

