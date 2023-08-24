# ApiDeviceLedsRingPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brightness** | Pointer to **int32** | The brightness of the led ring can be controlled in 6 steps, where 0 is off and 5 is full brightness. | [optional] 
**ShowFarendActivity** | Pointer to **bool** |  | [optional] 
**MicOn** | Pointer to [**ApiDeviceLedsRingPutRequestMicOn**](ApiDeviceLedsRingPutRequestMicOn.md) |  | [optional] 
**MicMute** | Pointer to [**ApiDeviceLedsRingPutRequestMicOn**](ApiDeviceLedsRingPutRequestMicOn.md) |  | [optional] 
**MicCustom** | Pointer to [**ApiDeviceLedsRingPutRequestMicCustom**](ApiDeviceLedsRingPutRequestMicCustom.md) |  | [optional] 

## Methods

### NewApiDeviceLedsRingPutRequest

`func NewApiDeviceLedsRingPutRequest() *ApiDeviceLedsRingPutRequest`

NewApiDeviceLedsRingPutRequest instantiates a new ApiDeviceLedsRingPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiDeviceLedsRingPutRequestWithDefaults

`func NewApiDeviceLedsRingPutRequestWithDefaults() *ApiDeviceLedsRingPutRequest`

NewApiDeviceLedsRingPutRequestWithDefaults instantiates a new ApiDeviceLedsRingPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrightness

`func (o *ApiDeviceLedsRingPutRequest) GetBrightness() int32`

GetBrightness returns the Brightness field if non-nil, zero value otherwise.

### GetBrightnessOk

`func (o *ApiDeviceLedsRingPutRequest) GetBrightnessOk() (*int32, bool)`

GetBrightnessOk returns a tuple with the Brightness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrightness

`func (o *ApiDeviceLedsRingPutRequest) SetBrightness(v int32)`

SetBrightness sets Brightness field to given value.

### HasBrightness

`func (o *ApiDeviceLedsRingPutRequest) HasBrightness() bool`

HasBrightness returns a boolean if a field has been set.

### GetShowFarendActivity

`func (o *ApiDeviceLedsRingPutRequest) GetShowFarendActivity() bool`

GetShowFarendActivity returns the ShowFarendActivity field if non-nil, zero value otherwise.

### GetShowFarendActivityOk

`func (o *ApiDeviceLedsRingPutRequest) GetShowFarendActivityOk() (*bool, bool)`

GetShowFarendActivityOk returns a tuple with the ShowFarendActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowFarendActivity

`func (o *ApiDeviceLedsRingPutRequest) SetShowFarendActivity(v bool)`

SetShowFarendActivity sets ShowFarendActivity field to given value.

### HasShowFarendActivity

`func (o *ApiDeviceLedsRingPutRequest) HasShowFarendActivity() bool`

HasShowFarendActivity returns a boolean if a field has been set.

### GetMicOn

`func (o *ApiDeviceLedsRingPutRequest) GetMicOn() ApiDeviceLedsRingPutRequestMicOn`

GetMicOn returns the MicOn field if non-nil, zero value otherwise.

### GetMicOnOk

`func (o *ApiDeviceLedsRingPutRequest) GetMicOnOk() (*ApiDeviceLedsRingPutRequestMicOn, bool)`

GetMicOnOk returns a tuple with the MicOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicOn

`func (o *ApiDeviceLedsRingPutRequest) SetMicOn(v ApiDeviceLedsRingPutRequestMicOn)`

SetMicOn sets MicOn field to given value.

### HasMicOn

`func (o *ApiDeviceLedsRingPutRequest) HasMicOn() bool`

HasMicOn returns a boolean if a field has been set.

### GetMicMute

`func (o *ApiDeviceLedsRingPutRequest) GetMicMute() ApiDeviceLedsRingPutRequestMicOn`

GetMicMute returns the MicMute field if non-nil, zero value otherwise.

### GetMicMuteOk

`func (o *ApiDeviceLedsRingPutRequest) GetMicMuteOk() (*ApiDeviceLedsRingPutRequestMicOn, bool)`

GetMicMuteOk returns a tuple with the MicMute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicMute

`func (o *ApiDeviceLedsRingPutRequest) SetMicMute(v ApiDeviceLedsRingPutRequestMicOn)`

SetMicMute sets MicMute field to given value.

### HasMicMute

`func (o *ApiDeviceLedsRingPutRequest) HasMicMute() bool`

HasMicMute returns a boolean if a field has been set.

### GetMicCustom

`func (o *ApiDeviceLedsRingPutRequest) GetMicCustom() ApiDeviceLedsRingPutRequestMicCustom`

GetMicCustom returns the MicCustom field if non-nil, zero value otherwise.

### GetMicCustomOk

`func (o *ApiDeviceLedsRingPutRequest) GetMicCustomOk() (*ApiDeviceLedsRingPutRequestMicCustom, bool)`

GetMicCustomOk returns a tuple with the MicCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCustom

`func (o *ApiDeviceLedsRingPutRequest) SetMicCustom(v ApiDeviceLedsRingPutRequestMicCustom)`

SetMicCustom sets MicCustom field to given value.

### HasMicCustom

`func (o *ApiDeviceLedsRingPutRequest) HasMicCustom() bool`

HasMicCustom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


