# ApiAudioOutputsDanteFarEndGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gain** | Pointer to **int32** | Gain in dB | [optional] [default to 12]
**NoiseGateEnabled** | Pointer to **bool** | Indication whether the settings from /api/audio/noiseGate are applied for this output | [optional] [default to false]
**EqualizerEnabled** | Pointer to **bool** | Indication whether the settings from /api/audio/equalizer are applied for this output | [optional] [default to false]
**Delay** | Pointer to **int32** | Delay in ms | [optional] [default to 0]

## Methods

### NewApiAudioOutputsDanteFarEndGet200Response

`func NewApiAudioOutputsDanteFarEndGet200Response() *ApiAudioOutputsDanteFarEndGet200Response`

NewApiAudioOutputsDanteFarEndGet200Response instantiates a new ApiAudioOutputsDanteFarEndGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAudioOutputsDanteFarEndGet200ResponseWithDefaults

`func NewApiAudioOutputsDanteFarEndGet200ResponseWithDefaults() *ApiAudioOutputsDanteFarEndGet200Response`

NewApiAudioOutputsDanteFarEndGet200ResponseWithDefaults instantiates a new ApiAudioOutputsDanteFarEndGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGain

`func (o *ApiAudioOutputsDanteFarEndGet200Response) GetGain() int32`

GetGain returns the Gain field if non-nil, zero value otherwise.

### GetGainOk

`func (o *ApiAudioOutputsDanteFarEndGet200Response) GetGainOk() (*int32, bool)`

GetGainOk returns a tuple with the Gain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGain

`func (o *ApiAudioOutputsDanteFarEndGet200Response) SetGain(v int32)`

SetGain sets Gain field to given value.

### HasGain

`func (o *ApiAudioOutputsDanteFarEndGet200Response) HasGain() bool`

HasGain returns a boolean if a field has been set.

### GetNoiseGateEnabled

`func (o *ApiAudioOutputsDanteFarEndGet200Response) GetNoiseGateEnabled() bool`

GetNoiseGateEnabled returns the NoiseGateEnabled field if non-nil, zero value otherwise.

### GetNoiseGateEnabledOk

`func (o *ApiAudioOutputsDanteFarEndGet200Response) GetNoiseGateEnabledOk() (*bool, bool)`

GetNoiseGateEnabledOk returns a tuple with the NoiseGateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoiseGateEnabled

`func (o *ApiAudioOutputsDanteFarEndGet200Response) SetNoiseGateEnabled(v bool)`

SetNoiseGateEnabled sets NoiseGateEnabled field to given value.

### HasNoiseGateEnabled

`func (o *ApiAudioOutputsDanteFarEndGet200Response) HasNoiseGateEnabled() bool`

HasNoiseGateEnabled returns a boolean if a field has been set.

### GetEqualizerEnabled

`func (o *ApiAudioOutputsDanteFarEndGet200Response) GetEqualizerEnabled() bool`

GetEqualizerEnabled returns the EqualizerEnabled field if non-nil, zero value otherwise.

### GetEqualizerEnabledOk

`func (o *ApiAudioOutputsDanteFarEndGet200Response) GetEqualizerEnabledOk() (*bool, bool)`

GetEqualizerEnabledOk returns a tuple with the EqualizerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEqualizerEnabled

`func (o *ApiAudioOutputsDanteFarEndGet200Response) SetEqualizerEnabled(v bool)`

SetEqualizerEnabled sets EqualizerEnabled field to given value.

### HasEqualizerEnabled

`func (o *ApiAudioOutputsDanteFarEndGet200Response) HasEqualizerEnabled() bool`

HasEqualizerEnabled returns a boolean if a field has been set.

### GetDelay

`func (o *ApiAudioOutputsDanteFarEndGet200Response) GetDelay() int32`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *ApiAudioOutputsDanteFarEndGet200Response) GetDelayOk() (*int32, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *ApiAudioOutputsDanteFarEndGet200Response) SetDelay(v int32)`

SetDelay sets Delay field to given value.

### HasDelay

`func (o *ApiAudioOutputsDanteFarEndGet200Response) HasDelay() bool`

HasDelay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


