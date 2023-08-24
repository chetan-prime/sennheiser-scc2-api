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

// checks if the ApiSscStateSubscriptionsSessionUUIDRemovePut400Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiSscStateSubscriptionsSessionUUIDRemovePut400Response{}

// ApiSscStateSubscriptionsSessionUUIDRemovePut400Response struct for ApiSscStateSubscriptionsSessionUUIDRemovePut400Response
type ApiSscStateSubscriptionsSessionUUIDRemovePut400Response struct {
	Path *string `json:"path,omitempty"`
	Error *int32 `json:"error,omitempty"`
}

// NewApiSscStateSubscriptionsSessionUUIDRemovePut400Response instantiates a new ApiSscStateSubscriptionsSessionUUIDRemovePut400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSscStateSubscriptionsSessionUUIDRemovePut400Response() *ApiSscStateSubscriptionsSessionUUIDRemovePut400Response {
	this := ApiSscStateSubscriptionsSessionUUIDRemovePut400Response{}
	return &this
}

// NewApiSscStateSubscriptionsSessionUUIDRemovePut400ResponseWithDefaults instantiates a new ApiSscStateSubscriptionsSessionUUIDRemovePut400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSscStateSubscriptionsSessionUUIDRemovePut400ResponseWithDefaults() *ApiSscStateSubscriptionsSessionUUIDRemovePut400Response {
	this := ApiSscStateSubscriptionsSessionUUIDRemovePut400Response{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ApiSscStateSubscriptionsSessionUUIDRemovePut400Response) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSscStateSubscriptionsSessionUUIDRemovePut400Response) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ApiSscStateSubscriptionsSessionUUIDRemovePut400Response) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ApiSscStateSubscriptionsSessionUUIDRemovePut400Response) SetPath(v string) {
	o.Path = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ApiSscStateSubscriptionsSessionUUIDRemovePut400Response) GetError() int32 {
	if o == nil || IsNil(o.Error) {
		var ret int32
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSscStateSubscriptionsSessionUUIDRemovePut400Response) GetErrorOk() (*int32, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ApiSscStateSubscriptionsSessionUUIDRemovePut400Response) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given int32 and assigns it to the Error field.
func (o *ApiSscStateSubscriptionsSessionUUIDRemovePut400Response) SetError(v int32) {
	o.Error = &v
}

func (o ApiSscStateSubscriptionsSessionUUIDRemovePut400Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiSscStateSubscriptionsSessionUUIDRemovePut400Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableApiSscStateSubscriptionsSessionUUIDRemovePut400Response struct {
	value *ApiSscStateSubscriptionsSessionUUIDRemovePut400Response
	isSet bool
}

func (v NullableApiSscStateSubscriptionsSessionUUIDRemovePut400Response) Get() *ApiSscStateSubscriptionsSessionUUIDRemovePut400Response {
	return v.value
}

func (v *NullableApiSscStateSubscriptionsSessionUUIDRemovePut400Response) Set(val *ApiSscStateSubscriptionsSessionUUIDRemovePut400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiSscStateSubscriptionsSessionUUIDRemovePut400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiSscStateSubscriptionsSessionUUIDRemovePut400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiSscStateSubscriptionsSessionUUIDRemovePut400Response(val *ApiSscStateSubscriptionsSessionUUIDRemovePut400Response) *NullableApiSscStateSubscriptionsSessionUUIDRemovePut400Response {
	return &NullableApiSscStateSubscriptionsSessionUUIDRemovePut400Response{value: val, isSet: true}
}

func (v NullableApiSscStateSubscriptionsSessionUUIDRemovePut400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiSscStateSubscriptionsSessionUUIDRemovePut400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

