# ApiAudioRoomInUseConfigGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TriggerTime** | Pointer to **int32** | Time in s of activity (i.e. activityLevel &gt; threshold) required for room in use to indicate &#39;true&#39; | [optional] [default to 15]
**ReleaseTime** | Pointer to **int32** | Time in s of inactivity required (i.e. activity_level &lt; threshold) for room in use to indicate &#39;false&#39; | [optional] [default to 300]
**Threshold** | Pointer to **int32** | Activity detection threshold in dB for room in use to turn true | [optional] [default to 10]

## Methods

### NewApiAudioRoomInUseConfigGet200Response

`func NewApiAudioRoomInUseConfigGet200Response() *ApiAudioRoomInUseConfigGet200Response`

NewApiAudioRoomInUseConfigGet200Response instantiates a new ApiAudioRoomInUseConfigGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAudioRoomInUseConfigGet200ResponseWithDefaults

`func NewApiAudioRoomInUseConfigGet200ResponseWithDefaults() *ApiAudioRoomInUseConfigGet200Response`

NewApiAudioRoomInUseConfigGet200ResponseWithDefaults instantiates a new ApiAudioRoomInUseConfigGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTriggerTime

`func (o *ApiAudioRoomInUseConfigGet200Response) GetTriggerTime() int32`

GetTriggerTime returns the TriggerTime field if non-nil, zero value otherwise.

### GetTriggerTimeOk

`func (o *ApiAudioRoomInUseConfigGet200Response) GetTriggerTimeOk() (*int32, bool)`

GetTriggerTimeOk returns a tuple with the TriggerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerTime

`func (o *ApiAudioRoomInUseConfigGet200Response) SetTriggerTime(v int32)`

SetTriggerTime sets TriggerTime field to given value.

### HasTriggerTime

`func (o *ApiAudioRoomInUseConfigGet200Response) HasTriggerTime() bool`

HasTriggerTime returns a boolean if a field has been set.

### GetReleaseTime

`func (o *ApiAudioRoomInUseConfigGet200Response) GetReleaseTime() int32`

GetReleaseTime returns the ReleaseTime field if non-nil, zero value otherwise.

### GetReleaseTimeOk

`func (o *ApiAudioRoomInUseConfigGet200Response) GetReleaseTimeOk() (*int32, bool)`

GetReleaseTimeOk returns a tuple with the ReleaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseTime

`func (o *ApiAudioRoomInUseConfigGet200Response) SetReleaseTime(v int32)`

SetReleaseTime sets ReleaseTime field to given value.

### HasReleaseTime

`func (o *ApiAudioRoomInUseConfigGet200Response) HasReleaseTime() bool`

HasReleaseTime returns a boolean if a field has been set.

### GetThreshold

`func (o *ApiAudioRoomInUseConfigGet200Response) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *ApiAudioRoomInUseConfigGet200Response) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *ApiAudioRoomInUseConfigGet200Response) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *ApiAudioRoomInUseConfigGet200Response) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


