/*
TCC M 3rd party API Release

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ApiDeviceIdentificationGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiDeviceIdentificationGet200Response{}

// ApiDeviceIdentificationGet200Response struct for ApiDeviceIdentificationGet200Response
type ApiDeviceIdentificationGet200Response struct {
	Visual *bool `json:"visual,omitempty"`
}

// NewApiDeviceIdentificationGet200Response instantiates a new ApiDeviceIdentificationGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiDeviceIdentificationGet200Response() *ApiDeviceIdentificationGet200Response {
	this := ApiDeviceIdentificationGet200Response{}
	var visual bool = false
	this.Visual = &visual
	return &this
}

// NewApiDeviceIdentificationGet200ResponseWithDefaults instantiates a new ApiDeviceIdentificationGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiDeviceIdentificationGet200ResponseWithDefaults() *ApiDeviceIdentificationGet200Response {
	this := ApiDeviceIdentificationGet200Response{}
	var visual bool = false
	this.Visual = &visual
	return &this
}

// GetVisual returns the Visual field value if set, zero value otherwise.
func (o *ApiDeviceIdentificationGet200Response) GetVisual() bool {
	if o == nil || IsNil(o.Visual) {
		var ret bool
		return ret
	}
	return *o.Visual
}

// GetVisualOk returns a tuple with the Visual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDeviceIdentificationGet200Response) GetVisualOk() (*bool, bool) {
	if o == nil || IsNil(o.Visual) {
		return nil, false
	}
	return o.Visual, true
}

// HasVisual returns a boolean if a field has been set.
func (o *ApiDeviceIdentificationGet200Response) HasVisual() bool {
	if o != nil && !IsNil(o.Visual) {
		return true
	}

	return false
}

// SetVisual gets a reference to the given bool and assigns it to the Visual field.
func (o *ApiDeviceIdentificationGet200Response) SetVisual(v bool) {
	o.Visual = &v
}

func (o ApiDeviceIdentificationGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiDeviceIdentificationGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Visual) {
		toSerialize["visual"] = o.Visual
	}
	return toSerialize, nil
}

type NullableApiDeviceIdentificationGet200Response struct {
	value *ApiDeviceIdentificationGet200Response
	isSet bool
}

func (v NullableApiDeviceIdentificationGet200Response) Get() *ApiDeviceIdentificationGet200Response {
	return v.value
}

func (v *NullableApiDeviceIdentificationGet200Response) Set(val *ApiDeviceIdentificationGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiDeviceIdentificationGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiDeviceIdentificationGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiDeviceIdentificationGet200Response(val *ApiDeviceIdentificationGet200Response) *NullableApiDeviceIdentificationGet200Response {
	return &NullableApiDeviceIdentificationGet200Response{value: val, isSet: true}
}

func (v NullableApiDeviceIdentificationGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiDeviceIdentificationGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

