# ApiAudioInputsMicrophoneBeamDirectionGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Azimuth** | Pointer to **int32** | Angle in degrees | [optional] 
**Elevation** | Pointer to **int32** | Angle in degrees | [optional] 
**BeamFreezeActive** | Pointer to **bool** | Indicates whether the beam is in beam-freeze, hence farend activity is recognized | [optional] 

## Methods

### NewApiAudioInputsMicrophoneBeamDirectionGet200Response

`func NewApiAudioInputsMicrophoneBeamDirectionGet200Response() *ApiAudioInputsMicrophoneBeamDirectionGet200Response`

NewApiAudioInputsMicrophoneBeamDirectionGet200Response instantiates a new ApiAudioInputsMicrophoneBeamDirectionGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAudioInputsMicrophoneBeamDirectionGet200ResponseWithDefaults

`func NewApiAudioInputsMicrophoneBeamDirectionGet200ResponseWithDefaults() *ApiAudioInputsMicrophoneBeamDirectionGet200Response`

NewApiAudioInputsMicrophoneBeamDirectionGet200ResponseWithDefaults instantiates a new ApiAudioInputsMicrophoneBeamDirectionGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzimuth

`func (o *ApiAudioInputsMicrophoneBeamDirectionGet200Response) GetAzimuth() int32`

GetAzimuth returns the Azimuth field if non-nil, zero value otherwise.

### GetAzimuthOk

`func (o *ApiAudioInputsMicrophoneBeamDirectionGet200Response) GetAzimuthOk() (*int32, bool)`

GetAzimuthOk returns a tuple with the Azimuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzimuth

`func (o *ApiAudioInputsMicrophoneBeamDirectionGet200Response) SetAzimuth(v int32)`

SetAzimuth sets Azimuth field to given value.

### HasAzimuth

`func (o *ApiAudioInputsMicrophoneBeamDirectionGet200Response) HasAzimuth() bool`

HasAzimuth returns a boolean if a field has been set.

### GetElevation

`func (o *ApiAudioInputsMicrophoneBeamDirectionGet200Response) GetElevation() int32`

GetElevation returns the Elevation field if non-nil, zero value otherwise.

### GetElevationOk

`func (o *ApiAudioInputsMicrophoneBeamDirectionGet200Response) GetElevationOk() (*int32, bool)`

GetElevationOk returns a tuple with the Elevation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElevation

`func (o *ApiAudioInputsMicrophoneBeamDirectionGet200Response) SetElevation(v int32)`

SetElevation sets Elevation field to given value.

### HasElevation

`func (o *ApiAudioInputsMicrophoneBeamDirectionGet200Response) HasElevation() bool`

HasElevation returns a boolean if a field has been set.

### GetBeamFreezeActive

`func (o *ApiAudioInputsMicrophoneBeamDirectionGet200Response) GetBeamFreezeActive() bool`

GetBeamFreezeActive returns the BeamFreezeActive field if non-nil, zero value otherwise.

### GetBeamFreezeActiveOk

`func (o *ApiAudioInputsMicrophoneBeamDirectionGet200Response) GetBeamFreezeActiveOk() (*bool, bool)`

GetBeamFreezeActiveOk returns a tuple with the BeamFreezeActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeamFreezeActive

`func (o *ApiAudioInputsMicrophoneBeamDirectionGet200Response) SetBeamFreezeActive(v bool)`

SetBeamFreezeActive sets BeamFreezeActive field to given value.

### HasBeamFreezeActive

`func (o *ApiAudioInputsMicrophoneBeamDirectionGet200Response) HasBeamFreezeActive() bool`

HasBeamFreezeActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


