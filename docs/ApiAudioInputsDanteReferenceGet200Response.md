# ApiAudioInputsDanteReferenceGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gain** | Pointer to **int32** | The manual gain for the far end detection in dB | [optional] [default to 0]
**FarEndAutoAdjustEnabled** | Pointer to **bool** | Indication whether automatic threshold adjustment based on noise floor estimation is enabled | [optional] [default to true]

## Methods

### NewApiAudioInputsDanteReferenceGet200Response

`func NewApiAudioInputsDanteReferenceGet200Response() *ApiAudioInputsDanteReferenceGet200Response`

NewApiAudioInputsDanteReferenceGet200Response instantiates a new ApiAudioInputsDanteReferenceGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAudioInputsDanteReferenceGet200ResponseWithDefaults

`func NewApiAudioInputsDanteReferenceGet200ResponseWithDefaults() *ApiAudioInputsDanteReferenceGet200Response`

NewApiAudioInputsDanteReferenceGet200ResponseWithDefaults instantiates a new ApiAudioInputsDanteReferenceGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGain

`func (o *ApiAudioInputsDanteReferenceGet200Response) GetGain() int32`

GetGain returns the Gain field if non-nil, zero value otherwise.

### GetGainOk

`func (o *ApiAudioInputsDanteReferenceGet200Response) GetGainOk() (*int32, bool)`

GetGainOk returns a tuple with the Gain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGain

`func (o *ApiAudioInputsDanteReferenceGet200Response) SetGain(v int32)`

SetGain sets Gain field to given value.

### HasGain

`func (o *ApiAudioInputsDanteReferenceGet200Response) HasGain() bool`

HasGain returns a boolean if a field has been set.

### GetFarEndAutoAdjustEnabled

`func (o *ApiAudioInputsDanteReferenceGet200Response) GetFarEndAutoAdjustEnabled() bool`

GetFarEndAutoAdjustEnabled returns the FarEndAutoAdjustEnabled field if non-nil, zero value otherwise.

### GetFarEndAutoAdjustEnabledOk

`func (o *ApiAudioInputsDanteReferenceGet200Response) GetFarEndAutoAdjustEnabledOk() (*bool, bool)`

GetFarEndAutoAdjustEnabledOk returns a tuple with the FarEndAutoAdjustEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFarEndAutoAdjustEnabled

`func (o *ApiAudioInputsDanteReferenceGet200Response) SetFarEndAutoAdjustEnabled(v bool)`

SetFarEndAutoAdjustEnabled sets FarEndAutoAdjustEnabled field to given value.

### HasFarEndAutoAdjustEnabled

`func (o *ApiAudioInputsDanteReferenceGet200Response) HasFarEndAutoAdjustEnabled() bool`

HasFarEndAutoAdjustEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


