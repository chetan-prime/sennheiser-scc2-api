# ApiAudioOutputsAnalogPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gain** | Pointer to **int32** | Gain in dB | [optional] 
**Switch** | Pointer to **string** |  | [optional] 

## Methods

### NewApiAudioOutputsAnalogPutRequest

`func NewApiAudioOutputsAnalogPutRequest() *ApiAudioOutputsAnalogPutRequest`

NewApiAudioOutputsAnalogPutRequest instantiates a new ApiAudioOutputsAnalogPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAudioOutputsAnalogPutRequestWithDefaults

`func NewApiAudioOutputsAnalogPutRequestWithDefaults() *ApiAudioOutputsAnalogPutRequest`

NewApiAudioOutputsAnalogPutRequestWithDefaults instantiates a new ApiAudioOutputsAnalogPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGain

`func (o *ApiAudioOutputsAnalogPutRequest) GetGain() int32`

GetGain returns the Gain field if non-nil, zero value otherwise.

### GetGainOk

`func (o *ApiAudioOutputsAnalogPutRequest) GetGainOk() (*int32, bool)`

GetGainOk returns a tuple with the Gain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGain

`func (o *ApiAudioOutputsAnalogPutRequest) SetGain(v int32)`

SetGain sets Gain field to given value.

### HasGain

`func (o *ApiAudioOutputsAnalogPutRequest) HasGain() bool`

HasGain returns a boolean if a field has been set.

### GetSwitch

`func (o *ApiAudioOutputsAnalogPutRequest) GetSwitch() string`

GetSwitch returns the Switch field if non-nil, zero value otherwise.

### GetSwitchOk

`func (o *ApiAudioOutputsAnalogPutRequest) GetSwitchOk() (*string, bool)`

GetSwitchOk returns a tuple with the Switch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitch

`func (o *ApiAudioOutputsAnalogPutRequest) SetSwitch(v string)`

SetSwitch sets Switch field to given value.

### HasSwitch

`func (o *ApiAudioOutputsAnalogPutRequest) HasSwitch() bool`

HasSwitch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


