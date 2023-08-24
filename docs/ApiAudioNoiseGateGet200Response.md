# ApiAudioNoiseGateGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Threshold** | Pointer to **int32** | Threshold in dB | [optional] [default to -80]
**HoldTime** | Pointer to **int32** | Hold time in ms | [optional] [default to 350]

## Methods

### NewApiAudioNoiseGateGet200Response

`func NewApiAudioNoiseGateGet200Response() *ApiAudioNoiseGateGet200Response`

NewApiAudioNoiseGateGet200Response instantiates a new ApiAudioNoiseGateGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAudioNoiseGateGet200ResponseWithDefaults

`func NewApiAudioNoiseGateGet200ResponseWithDefaults() *ApiAudioNoiseGateGet200Response`

NewApiAudioNoiseGateGet200ResponseWithDefaults instantiates a new ApiAudioNoiseGateGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreshold

`func (o *ApiAudioNoiseGateGet200Response) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *ApiAudioNoiseGateGet200Response) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *ApiAudioNoiseGateGet200Response) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *ApiAudioNoiseGateGet200Response) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetHoldTime

`func (o *ApiAudioNoiseGateGet200Response) GetHoldTime() int32`

GetHoldTime returns the HoldTime field if non-nil, zero value otherwise.

### GetHoldTimeOk

`func (o *ApiAudioNoiseGateGet200Response) GetHoldTimeOk() (*int32, bool)`

GetHoldTimeOk returns a tuple with the HoldTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldTime

`func (o *ApiAudioNoiseGateGet200Response) SetHoldTime(v int32)`

SetHoldTime sets HoldTime field to given value.

### HasHoldTime

`func (o *ApiAudioNoiseGateGet200Response) HasHoldTime() bool`

HasHoldTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


