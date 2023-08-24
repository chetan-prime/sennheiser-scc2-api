# ApiAudioNoiseGatePutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Threshold** | Pointer to **int32** | Threshold in dB | [optional] 
**HoldTime** | Pointer to **int32** | Hold time in ms | [optional] 

## Methods

### NewApiAudioNoiseGatePutRequest

`func NewApiAudioNoiseGatePutRequest() *ApiAudioNoiseGatePutRequest`

NewApiAudioNoiseGatePutRequest instantiates a new ApiAudioNoiseGatePutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAudioNoiseGatePutRequestWithDefaults

`func NewApiAudioNoiseGatePutRequestWithDefaults() *ApiAudioNoiseGatePutRequest`

NewApiAudioNoiseGatePutRequestWithDefaults instantiates a new ApiAudioNoiseGatePutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreshold

`func (o *ApiAudioNoiseGatePutRequest) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *ApiAudioNoiseGatePutRequest) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *ApiAudioNoiseGatePutRequest) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *ApiAudioNoiseGatePutRequest) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetHoldTime

`func (o *ApiAudioNoiseGatePutRequest) GetHoldTime() int32`

GetHoldTime returns the HoldTime field if non-nil, zero value otherwise.

### GetHoldTimeOk

`func (o *ApiAudioNoiseGatePutRequest) GetHoldTimeOk() (*int32, bool)`

GetHoldTimeOk returns a tuple with the HoldTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldTime

`func (o *ApiAudioNoiseGatePutRequest) SetHoldTime(v int32)`

SetHoldTime sets HoldTime field to given value.

### HasHoldTime

`func (o *ApiAudioNoiseGatePutRequest) HasHoldTime() bool`

HasHoldTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


