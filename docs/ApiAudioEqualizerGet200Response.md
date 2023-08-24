# ApiAudioEqualizerGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gains** | Pointer to **[]int32** | Equalizer bands with cut-off frequencies of 125, 250, 500, 1000, 2000, 4000, 8000 and Q value of 1.4142 | [optional] 

## Methods

### NewApiAudioEqualizerGet200Response

`func NewApiAudioEqualizerGet200Response() *ApiAudioEqualizerGet200Response`

NewApiAudioEqualizerGet200Response instantiates a new ApiAudioEqualizerGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAudioEqualizerGet200ResponseWithDefaults

`func NewApiAudioEqualizerGet200ResponseWithDefaults() *ApiAudioEqualizerGet200Response`

NewApiAudioEqualizerGet200ResponseWithDefaults instantiates a new ApiAudioEqualizerGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGains

`func (o *ApiAudioEqualizerGet200Response) GetGains() []int32`

GetGains returns the Gains field if non-nil, zero value otherwise.

### GetGainsOk

`func (o *ApiAudioEqualizerGet200Response) GetGainsOk() (*[]int32, bool)`

GetGainsOk returns a tuple with the Gains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGains

`func (o *ApiAudioEqualizerGet200Response) SetGains(v []int32)`

SetGains sets Gains field to given value.

### HasGains

`func (o *ApiAudioEqualizerGet200Response) HasGains() bool`

HasGains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


