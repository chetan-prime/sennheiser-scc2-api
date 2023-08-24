# ApiAudioOutputsDanteFarEndPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gain** | Pointer to **int32** | Gain in dB | [optional] 
**NoiseGateEnabled** | Pointer to **bool** | Applies the settings from /api/audio/noiseGate | [optional] 
**EqualizerEnabled** | Pointer to **bool** | Applies the settings from /api/audio/equalizer | [optional] 
**Delay** | Pointer to **int32** | Delay in ms | [optional] 

## Methods

### NewApiAudioOutputsDanteFarEndPutRequest

`func NewApiAudioOutputsDanteFarEndPutRequest() *ApiAudioOutputsDanteFarEndPutRequest`

NewApiAudioOutputsDanteFarEndPutRequest instantiates a new ApiAudioOutputsDanteFarEndPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAudioOutputsDanteFarEndPutRequestWithDefaults

`func NewApiAudioOutputsDanteFarEndPutRequestWithDefaults() *ApiAudioOutputsDanteFarEndPutRequest`

NewApiAudioOutputsDanteFarEndPutRequestWithDefaults instantiates a new ApiAudioOutputsDanteFarEndPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGain

`func (o *ApiAudioOutputsDanteFarEndPutRequest) GetGain() int32`

GetGain returns the Gain field if non-nil, zero value otherwise.

### GetGainOk

`func (o *ApiAudioOutputsDanteFarEndPutRequest) GetGainOk() (*int32, bool)`

GetGainOk returns a tuple with the Gain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGain

`func (o *ApiAudioOutputsDanteFarEndPutRequest) SetGain(v int32)`

SetGain sets Gain field to given value.

### HasGain

`func (o *ApiAudioOutputsDanteFarEndPutRequest) HasGain() bool`

HasGain returns a boolean if a field has been set.

### GetNoiseGateEnabled

`func (o *ApiAudioOutputsDanteFarEndPutRequest) GetNoiseGateEnabled() bool`

GetNoiseGateEnabled returns the NoiseGateEnabled field if non-nil, zero value otherwise.

### GetNoiseGateEnabledOk

`func (o *ApiAudioOutputsDanteFarEndPutRequest) GetNoiseGateEnabledOk() (*bool, bool)`

GetNoiseGateEnabledOk returns a tuple with the NoiseGateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoiseGateEnabled

`func (o *ApiAudioOutputsDanteFarEndPutRequest) SetNoiseGateEnabled(v bool)`

SetNoiseGateEnabled sets NoiseGateEnabled field to given value.

### HasNoiseGateEnabled

`func (o *ApiAudioOutputsDanteFarEndPutRequest) HasNoiseGateEnabled() bool`

HasNoiseGateEnabled returns a boolean if a field has been set.

### GetEqualizerEnabled

`func (o *ApiAudioOutputsDanteFarEndPutRequest) GetEqualizerEnabled() bool`

GetEqualizerEnabled returns the EqualizerEnabled field if non-nil, zero value otherwise.

### GetEqualizerEnabledOk

`func (o *ApiAudioOutputsDanteFarEndPutRequest) GetEqualizerEnabledOk() (*bool, bool)`

GetEqualizerEnabledOk returns a tuple with the EqualizerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEqualizerEnabled

`func (o *ApiAudioOutputsDanteFarEndPutRequest) SetEqualizerEnabled(v bool)`

SetEqualizerEnabled sets EqualizerEnabled field to given value.

### HasEqualizerEnabled

`func (o *ApiAudioOutputsDanteFarEndPutRequest) HasEqualizerEnabled() bool`

HasEqualizerEnabled returns a boolean if a field has been set.

### GetDelay

`func (o *ApiAudioOutputsDanteFarEndPutRequest) GetDelay() int32`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *ApiAudioOutputsDanteFarEndPutRequest) GetDelayOk() (*int32, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *ApiAudioOutputsDanteFarEndPutRequest) SetDelay(v int32)`

SetDelay sets Delay field to given value.

### HasDelay

`func (o *ApiAudioOutputsDanteFarEndPutRequest) HasDelay() bool`

HasDelay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


