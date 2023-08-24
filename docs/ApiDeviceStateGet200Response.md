# ApiDeviceStateGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] [default to "Normal"]
**Warnings** | Pointer to **[]string** |  | [optional] [default to []]

## Methods

### NewApiDeviceStateGet200Response

`func NewApiDeviceStateGet200Response() *ApiDeviceStateGet200Response`

NewApiDeviceStateGet200Response instantiates a new ApiDeviceStateGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiDeviceStateGet200ResponseWithDefaults

`func NewApiDeviceStateGet200ResponseWithDefaults() *ApiDeviceStateGet200Response`

NewApiDeviceStateGet200ResponseWithDefaults instantiates a new ApiDeviceStateGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ApiDeviceStateGet200Response) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApiDeviceStateGet200Response) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ApiDeviceStateGet200Response) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ApiDeviceStateGet200Response) HasState() bool`

HasState returns a boolean if a field has been set.

### GetWarnings

`func (o *ApiDeviceStateGet200Response) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *ApiDeviceStateGet200Response) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *ApiDeviceStateGet200Response) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *ApiDeviceStateGet200Response) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


