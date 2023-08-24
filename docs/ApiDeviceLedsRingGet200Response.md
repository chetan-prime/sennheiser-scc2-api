# ApiDeviceLedsRingGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brightness** | Pointer to **int32** | The brightness of the led ring can be controlled in 6 steps, where 0 is off and 5 is full brightness. | [optional] [default to 5]
**ShowFarendActivity** | Pointer to **bool** |  | [optional] [default to false]
**MicOn** | Pointer to [**ApiDeviceLedsRingGet200ResponseMicOn**](ApiDeviceLedsRingGet200ResponseMicOn.md) |  | [optional] 
**MicMute** | Pointer to [**ApiDeviceLedsRingGet200ResponseMicOn**](ApiDeviceLedsRingGet200ResponseMicOn.md) |  | [optional] 
**MicCustom** | Pointer to [**ApiDeviceLedsRingGet200ResponseMicCustom**](ApiDeviceLedsRingGet200ResponseMicCustom.md) |  | [optional] 

## Methods

### NewApiDeviceLedsRingGet200Response

`func NewApiDeviceLedsRingGet200Response() *ApiDeviceLedsRingGet200Response`

NewApiDeviceLedsRingGet200Response instantiates a new ApiDeviceLedsRingGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiDeviceLedsRingGet200ResponseWithDefaults

`func NewApiDeviceLedsRingGet200ResponseWithDefaults() *ApiDeviceLedsRingGet200Response`

NewApiDeviceLedsRingGet200ResponseWithDefaults instantiates a new ApiDeviceLedsRingGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrightness

`func (o *ApiDeviceLedsRingGet200Response) GetBrightness() int32`

GetBrightness returns the Brightness field if non-nil, zero value otherwise.

### GetBrightnessOk

`func (o *ApiDeviceLedsRingGet200Response) GetBrightnessOk() (*int32, bool)`

GetBrightnessOk returns a tuple with the Brightness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrightness

`func (o *ApiDeviceLedsRingGet200Response) SetBrightness(v int32)`

SetBrightness sets Brightness field to given value.

### HasBrightness

`func (o *ApiDeviceLedsRingGet200Response) HasBrightness() bool`

HasBrightness returns a boolean if a field has been set.

### GetShowFarendActivity

`func (o *ApiDeviceLedsRingGet200Response) GetShowFarendActivity() bool`

GetShowFarendActivity returns the ShowFarendActivity field if non-nil, zero value otherwise.

### GetShowFarendActivityOk

`func (o *ApiDeviceLedsRingGet200Response) GetShowFarendActivityOk() (*bool, bool)`

GetShowFarendActivityOk returns a tuple with the ShowFarendActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowFarendActivity

`func (o *ApiDeviceLedsRingGet200Response) SetShowFarendActivity(v bool)`

SetShowFarendActivity sets ShowFarendActivity field to given value.

### HasShowFarendActivity

`func (o *ApiDeviceLedsRingGet200Response) HasShowFarendActivity() bool`

HasShowFarendActivity returns a boolean if a field has been set.

### GetMicOn

`func (o *ApiDeviceLedsRingGet200Response) GetMicOn() ApiDeviceLedsRingGet200ResponseMicOn`

GetMicOn returns the MicOn field if non-nil, zero value otherwise.

### GetMicOnOk

`func (o *ApiDeviceLedsRingGet200Response) GetMicOnOk() (*ApiDeviceLedsRingGet200ResponseMicOn, bool)`

GetMicOnOk returns a tuple with the MicOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicOn

`func (o *ApiDeviceLedsRingGet200Response) SetMicOn(v ApiDeviceLedsRingGet200ResponseMicOn)`

SetMicOn sets MicOn field to given value.

### HasMicOn

`func (o *ApiDeviceLedsRingGet200Response) HasMicOn() bool`

HasMicOn returns a boolean if a field has been set.

### GetMicMute

`func (o *ApiDeviceLedsRingGet200Response) GetMicMute() ApiDeviceLedsRingGet200ResponseMicOn`

GetMicMute returns the MicMute field if non-nil, zero value otherwise.

### GetMicMuteOk

`func (o *ApiDeviceLedsRingGet200Response) GetMicMuteOk() (*ApiDeviceLedsRingGet200ResponseMicOn, bool)`

GetMicMuteOk returns a tuple with the MicMute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicMute

`func (o *ApiDeviceLedsRingGet200Response) SetMicMute(v ApiDeviceLedsRingGet200ResponseMicOn)`

SetMicMute sets MicMute field to given value.

### HasMicMute

`func (o *ApiDeviceLedsRingGet200Response) HasMicMute() bool`

HasMicMute returns a boolean if a field has been set.

### GetMicCustom

`func (o *ApiDeviceLedsRingGet200Response) GetMicCustom() ApiDeviceLedsRingGet200ResponseMicCustom`

GetMicCustom returns the MicCustom field if non-nil, zero value otherwise.

### GetMicCustomOk

`func (o *ApiDeviceLedsRingGet200Response) GetMicCustomOk() (*ApiDeviceLedsRingGet200ResponseMicCustom, bool)`

GetMicCustomOk returns a tuple with the MicCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCustom

`func (o *ApiDeviceLedsRingGet200Response) SetMicCustom(v ApiDeviceLedsRingGet200ResponseMicCustom)`

SetMicCustom sets MicCustom field to given value.

### HasMicCustom

`func (o *ApiDeviceLedsRingGet200Response) HasMicCustom() bool`

HasMicCustom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


