# ApiFirmwareUpdateStateGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceVersion** | Pointer to **string** | The firmware version of the system running at the moment. | [optional] 
**DanteVersion** | Pointer to **string** | The Dante firmware version of the system running at the moment. | [optional] 
**State** | Pointer to **string** |  | [optional] [default to "Idle"]
**Progress** | Pointer to **int32** | Progress in percent | [optional] [default to 0]
**LastStatus** | Pointer to **string** | Value is reset to None after reboot and on every start of FW update | [optional] [default to "None"]

## Methods

### NewApiFirmwareUpdateStateGet200Response

`func NewApiFirmwareUpdateStateGet200Response() *ApiFirmwareUpdateStateGet200Response`

NewApiFirmwareUpdateStateGet200Response instantiates a new ApiFirmwareUpdateStateGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiFirmwareUpdateStateGet200ResponseWithDefaults

`func NewApiFirmwareUpdateStateGet200ResponseWithDefaults() *ApiFirmwareUpdateStateGet200Response`

NewApiFirmwareUpdateStateGet200ResponseWithDefaults instantiates a new ApiFirmwareUpdateStateGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceVersion

`func (o *ApiFirmwareUpdateStateGet200Response) GetDeviceVersion() string`

GetDeviceVersion returns the DeviceVersion field if non-nil, zero value otherwise.

### GetDeviceVersionOk

`func (o *ApiFirmwareUpdateStateGet200Response) GetDeviceVersionOk() (*string, bool)`

GetDeviceVersionOk returns a tuple with the DeviceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceVersion

`func (o *ApiFirmwareUpdateStateGet200Response) SetDeviceVersion(v string)`

SetDeviceVersion sets DeviceVersion field to given value.

### HasDeviceVersion

`func (o *ApiFirmwareUpdateStateGet200Response) HasDeviceVersion() bool`

HasDeviceVersion returns a boolean if a field has been set.

### GetDanteVersion

`func (o *ApiFirmwareUpdateStateGet200Response) GetDanteVersion() string`

GetDanteVersion returns the DanteVersion field if non-nil, zero value otherwise.

### GetDanteVersionOk

`func (o *ApiFirmwareUpdateStateGet200Response) GetDanteVersionOk() (*string, bool)`

GetDanteVersionOk returns a tuple with the DanteVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDanteVersion

`func (o *ApiFirmwareUpdateStateGet200Response) SetDanteVersion(v string)`

SetDanteVersion sets DanteVersion field to given value.

### HasDanteVersion

`func (o *ApiFirmwareUpdateStateGet200Response) HasDanteVersion() bool`

HasDanteVersion returns a boolean if a field has been set.

### GetState

`func (o *ApiFirmwareUpdateStateGet200Response) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApiFirmwareUpdateStateGet200Response) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ApiFirmwareUpdateStateGet200Response) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ApiFirmwareUpdateStateGet200Response) HasState() bool`

HasState returns a boolean if a field has been set.

### GetProgress

`func (o *ApiFirmwareUpdateStateGet200Response) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ApiFirmwareUpdateStateGet200Response) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ApiFirmwareUpdateStateGet200Response) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *ApiFirmwareUpdateStateGet200Response) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetLastStatus

`func (o *ApiFirmwareUpdateStateGet200Response) GetLastStatus() string`

GetLastStatus returns the LastStatus field if non-nil, zero value otherwise.

### GetLastStatusOk

`func (o *ApiFirmwareUpdateStateGet200Response) GetLastStatusOk() (*string, bool)`

GetLastStatusOk returns a tuple with the LastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatus

`func (o *ApiFirmwareUpdateStateGet200Response) SetLastStatus(v string)`

SetLastStatus sets LastStatus field to given value.

### HasLastStatus

`func (o *ApiFirmwareUpdateStateGet200Response) HasLastStatus() bool`

HasLastStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


