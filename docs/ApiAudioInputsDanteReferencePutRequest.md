# ApiAudioInputsDanteReferencePutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gain** | Pointer to **int32** | The manual gain for the far end detection in dB. Must not be sent if &#x60;farEndAutoAdjustEnabled&#x60; is &#x60;true&#x60;. | [optional] 
**FarEndAutoAdjustEnabled** | Pointer to **bool** | Enable automatic threshold adjustment based on noise floor estimation. If &#x60;true&#x60; must not be sent together with &#x60;gain&#x60;. | [optional] 

## Methods

### NewApiAudioInputsDanteReferencePutRequest

`func NewApiAudioInputsDanteReferencePutRequest() *ApiAudioInputsDanteReferencePutRequest`

NewApiAudioInputsDanteReferencePutRequest instantiates a new ApiAudioInputsDanteReferencePutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAudioInputsDanteReferencePutRequestWithDefaults

`func NewApiAudioInputsDanteReferencePutRequestWithDefaults() *ApiAudioInputsDanteReferencePutRequest`

NewApiAudioInputsDanteReferencePutRequestWithDefaults instantiates a new ApiAudioInputsDanteReferencePutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGain

`func (o *ApiAudioInputsDanteReferencePutRequest) GetGain() int32`

GetGain returns the Gain field if non-nil, zero value otherwise.

### GetGainOk

`func (o *ApiAudioInputsDanteReferencePutRequest) GetGainOk() (*int32, bool)`

GetGainOk returns a tuple with the Gain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGain

`func (o *ApiAudioInputsDanteReferencePutRequest) SetGain(v int32)`

SetGain sets Gain field to given value.

### HasGain

`func (o *ApiAudioInputsDanteReferencePutRequest) HasGain() bool`

HasGain returns a boolean if a field has been set.

### GetFarEndAutoAdjustEnabled

`func (o *ApiAudioInputsDanteReferencePutRequest) GetFarEndAutoAdjustEnabled() bool`

GetFarEndAutoAdjustEnabled returns the FarEndAutoAdjustEnabled field if non-nil, zero value otherwise.

### GetFarEndAutoAdjustEnabledOk

`func (o *ApiAudioInputsDanteReferencePutRequest) GetFarEndAutoAdjustEnabledOk() (*bool, bool)`

GetFarEndAutoAdjustEnabledOk returns a tuple with the FarEndAutoAdjustEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFarEndAutoAdjustEnabled

`func (o *ApiAudioInputsDanteReferencePutRequest) SetFarEndAutoAdjustEnabled(v bool)`

SetFarEndAutoAdjustEnabled sets FarEndAutoAdjustEnabled field to given value.

### HasFarEndAutoAdjustEnabled

`func (o *ApiAudioInputsDanteReferencePutRequest) HasFarEndAutoAdjustEnabled() bool`

HasFarEndAutoAdjustEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


