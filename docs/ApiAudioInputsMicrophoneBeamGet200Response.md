# ApiAudioInputsMicrophoneBeamGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstallationType** | Pointer to **string** | The audio signal will be optimized for the chosen installation type. | [optional] [default to "FlushMounted"]
**SourceDetectionThreshold** | Pointer to **string** | Threshold for detecting the speaker depending on the room noise level. | [optional] [default to "NormalRoom"]
**Offset** | Pointer to **int32** | Offset to the beam&#39;s azimuth angle. This will affect the beam metering and zone settings. | [optional] [default to 0]

## Methods

### NewApiAudioInputsMicrophoneBeamGet200Response

`func NewApiAudioInputsMicrophoneBeamGet200Response() *ApiAudioInputsMicrophoneBeamGet200Response`

NewApiAudioInputsMicrophoneBeamGet200Response instantiates a new ApiAudioInputsMicrophoneBeamGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAudioInputsMicrophoneBeamGet200ResponseWithDefaults

`func NewApiAudioInputsMicrophoneBeamGet200ResponseWithDefaults() *ApiAudioInputsMicrophoneBeamGet200Response`

NewApiAudioInputsMicrophoneBeamGet200ResponseWithDefaults instantiates a new ApiAudioInputsMicrophoneBeamGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstallationType

`func (o *ApiAudioInputsMicrophoneBeamGet200Response) GetInstallationType() string`

GetInstallationType returns the InstallationType field if non-nil, zero value otherwise.

### GetInstallationTypeOk

`func (o *ApiAudioInputsMicrophoneBeamGet200Response) GetInstallationTypeOk() (*string, bool)`

GetInstallationTypeOk returns a tuple with the InstallationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationType

`func (o *ApiAudioInputsMicrophoneBeamGet200Response) SetInstallationType(v string)`

SetInstallationType sets InstallationType field to given value.

### HasInstallationType

`func (o *ApiAudioInputsMicrophoneBeamGet200Response) HasInstallationType() bool`

HasInstallationType returns a boolean if a field has been set.

### GetSourceDetectionThreshold

`func (o *ApiAudioInputsMicrophoneBeamGet200Response) GetSourceDetectionThreshold() string`

GetSourceDetectionThreshold returns the SourceDetectionThreshold field if non-nil, zero value otherwise.

### GetSourceDetectionThresholdOk

`func (o *ApiAudioInputsMicrophoneBeamGet200Response) GetSourceDetectionThresholdOk() (*string, bool)`

GetSourceDetectionThresholdOk returns a tuple with the SourceDetectionThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDetectionThreshold

`func (o *ApiAudioInputsMicrophoneBeamGet200Response) SetSourceDetectionThreshold(v string)`

SetSourceDetectionThreshold sets SourceDetectionThreshold field to given value.

### HasSourceDetectionThreshold

`func (o *ApiAudioInputsMicrophoneBeamGet200Response) HasSourceDetectionThreshold() bool`

HasSourceDetectionThreshold returns a boolean if a field has been set.

### GetOffset

`func (o *ApiAudioInputsMicrophoneBeamGet200Response) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ApiAudioInputsMicrophoneBeamGet200Response) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ApiAudioInputsMicrophoneBeamGet200Response) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ApiAudioInputsMicrophoneBeamGet200Response) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


