# ApiAudioInputsMicrophoneBeamPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstallationType** | Pointer to **string** | The audio signal will be optimized for the chosen installation type. | [optional] 
**SourceDetectionThreshold** | Pointer to **string** | Threshold for detecting the speaker depending on the room noise level. | [optional] 
**Offset** | Pointer to **int32** | Offset to the beam&#39;s azimuth angle. This will affect &#x60;/audio/metering/beamDirection&#x60; and zone settings. | [optional] 

## Methods

### NewApiAudioInputsMicrophoneBeamPutRequest

`func NewApiAudioInputsMicrophoneBeamPutRequest() *ApiAudioInputsMicrophoneBeamPutRequest`

NewApiAudioInputsMicrophoneBeamPutRequest instantiates a new ApiAudioInputsMicrophoneBeamPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAudioInputsMicrophoneBeamPutRequestWithDefaults

`func NewApiAudioInputsMicrophoneBeamPutRequestWithDefaults() *ApiAudioInputsMicrophoneBeamPutRequest`

NewApiAudioInputsMicrophoneBeamPutRequestWithDefaults instantiates a new ApiAudioInputsMicrophoneBeamPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstallationType

`func (o *ApiAudioInputsMicrophoneBeamPutRequest) GetInstallationType() string`

GetInstallationType returns the InstallationType field if non-nil, zero value otherwise.

### GetInstallationTypeOk

`func (o *ApiAudioInputsMicrophoneBeamPutRequest) GetInstallationTypeOk() (*string, bool)`

GetInstallationTypeOk returns a tuple with the InstallationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationType

`func (o *ApiAudioInputsMicrophoneBeamPutRequest) SetInstallationType(v string)`

SetInstallationType sets InstallationType field to given value.

### HasInstallationType

`func (o *ApiAudioInputsMicrophoneBeamPutRequest) HasInstallationType() bool`

HasInstallationType returns a boolean if a field has been set.

### GetSourceDetectionThreshold

`func (o *ApiAudioInputsMicrophoneBeamPutRequest) GetSourceDetectionThreshold() string`

GetSourceDetectionThreshold returns the SourceDetectionThreshold field if non-nil, zero value otherwise.

### GetSourceDetectionThresholdOk

`func (o *ApiAudioInputsMicrophoneBeamPutRequest) GetSourceDetectionThresholdOk() (*string, bool)`

GetSourceDetectionThresholdOk returns a tuple with the SourceDetectionThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDetectionThreshold

`func (o *ApiAudioInputsMicrophoneBeamPutRequest) SetSourceDetectionThreshold(v string)`

SetSourceDetectionThreshold sets SourceDetectionThreshold field to given value.

### HasSourceDetectionThreshold

`func (o *ApiAudioInputsMicrophoneBeamPutRequest) HasSourceDetectionThreshold() bool`

HasSourceDetectionThreshold returns a boolean if a field has been set.

### GetOffset

`func (o *ApiAudioInputsMicrophoneBeamPutRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ApiAudioInputsMicrophoneBeamPutRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ApiAudioInputsMicrophoneBeamPutRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ApiAudioInputsMicrophoneBeamPutRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


