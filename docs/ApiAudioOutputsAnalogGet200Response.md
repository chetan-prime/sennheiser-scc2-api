# ApiAudioOutputsAnalogGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gain** | Pointer to **int32** | Gain in dB | [optional] [default to 0]
**Switch** | Pointer to **string** |  | [optional] [default to "FarendOutput"]

## Methods

### NewApiAudioOutputsAnalogGet200Response

`func NewApiAudioOutputsAnalogGet200Response() *ApiAudioOutputsAnalogGet200Response`

NewApiAudioOutputsAnalogGet200Response instantiates a new ApiAudioOutputsAnalogGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAudioOutputsAnalogGet200ResponseWithDefaults

`func NewApiAudioOutputsAnalogGet200ResponseWithDefaults() *ApiAudioOutputsAnalogGet200Response`

NewApiAudioOutputsAnalogGet200ResponseWithDefaults instantiates a new ApiAudioOutputsAnalogGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGain

`func (o *ApiAudioOutputsAnalogGet200Response) GetGain() int32`

GetGain returns the Gain field if non-nil, zero value otherwise.

### GetGainOk

`func (o *ApiAudioOutputsAnalogGet200Response) GetGainOk() (*int32, bool)`

GetGainOk returns a tuple with the Gain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGain

`func (o *ApiAudioOutputsAnalogGet200Response) SetGain(v int32)`

SetGain sets Gain field to given value.

### HasGain

`func (o *ApiAudioOutputsAnalogGet200Response) HasGain() bool`

HasGain returns a boolean if a field has been set.

### GetSwitch

`func (o *ApiAudioOutputsAnalogGet200Response) GetSwitch() string`

GetSwitch returns the Switch field if non-nil, zero value otherwise.

### GetSwitchOk

`func (o *ApiAudioOutputsAnalogGet200Response) GetSwitchOk() (*string, bool)`

GetSwitchOk returns a tuple with the Switch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitch

`func (o *ApiAudioOutputsAnalogGet200Response) SetSwitch(v string)`

SetSwitch sets Switch field to given value.

### HasSwitch

`func (o *ApiAudioOutputsAnalogGet200Response) HasSwitch() bool`

HasSwitch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


