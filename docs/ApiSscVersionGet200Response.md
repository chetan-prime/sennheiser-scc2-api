# ApiSscVersionGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to **string** | The version of SSC protocol. | [optional] 
**Schema** | Pointer to **string** | This is the schema version of the API of the TCC M device. Semantic versioning must be used. So additional paths or properties will increase the minor number while breaking changes increase the integer part before the decimal point. | [optional] 

## Methods

### NewApiSscVersionGet200Response

`func NewApiSscVersionGet200Response() *ApiSscVersionGet200Response`

NewApiSscVersionGet200Response instantiates a new ApiSscVersionGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSscVersionGet200ResponseWithDefaults

`func NewApiSscVersionGet200ResponseWithDefaults() *ApiSscVersionGet200Response`

NewApiSscVersionGet200ResponseWithDefaults instantiates a new ApiSscVersionGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *ApiSscVersionGet200Response) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ApiSscVersionGet200Response) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ApiSscVersionGet200Response) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ApiSscVersionGet200Response) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSchema

`func (o *ApiSscVersionGet200Response) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *ApiSscVersionGet200Response) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *ApiSscVersionGet200Response) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *ApiSscVersionGet200Response) HasSchema() bool`

HasSchema returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


