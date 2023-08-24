# ApiDevicePowerPoeDaisychainGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SufficientPower** | Pointer to **bool** | Indicates whether there is sufficient power on the PoE input to power a subsequent device on the PoE output. | [optional] 
**InUse** | Pointer to **bool** | Indicates whether there is a subsequent device being powered on the PoE output. | [optional] 

## Methods

### NewApiDevicePowerPoeDaisychainGet200Response

`func NewApiDevicePowerPoeDaisychainGet200Response() *ApiDevicePowerPoeDaisychainGet200Response`

NewApiDevicePowerPoeDaisychainGet200Response instantiates a new ApiDevicePowerPoeDaisychainGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiDevicePowerPoeDaisychainGet200ResponseWithDefaults

`func NewApiDevicePowerPoeDaisychainGet200ResponseWithDefaults() *ApiDevicePowerPoeDaisychainGet200Response`

NewApiDevicePowerPoeDaisychainGet200ResponseWithDefaults instantiates a new ApiDevicePowerPoeDaisychainGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSufficientPower

`func (o *ApiDevicePowerPoeDaisychainGet200Response) GetSufficientPower() bool`

GetSufficientPower returns the SufficientPower field if non-nil, zero value otherwise.

### GetSufficientPowerOk

`func (o *ApiDevicePowerPoeDaisychainGet200Response) GetSufficientPowerOk() (*bool, bool)`

GetSufficientPowerOk returns a tuple with the SufficientPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSufficientPower

`func (o *ApiDevicePowerPoeDaisychainGet200Response) SetSufficientPower(v bool)`

SetSufficientPower sets SufficientPower field to given value.

### HasSufficientPower

`func (o *ApiDevicePowerPoeDaisychainGet200Response) HasSufficientPower() bool`

HasSufficientPower returns a boolean if a field has been set.

### GetInUse

`func (o *ApiDevicePowerPoeDaisychainGet200Response) GetInUse() bool`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *ApiDevicePowerPoeDaisychainGet200Response) GetInUseOk() (*bool, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *ApiDevicePowerPoeDaisychainGet200Response) SetInUse(v bool)`

SetInUse sets InUse field to given value.

### HasInUse

`func (o *ApiDevicePowerPoeDaisychainGet200Response) HasInUse() bool`

HasInUse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


