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

// checks if the ApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth{}

// ApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth struct for ApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth
type ApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth struct {
	// Angle in degrees
	Min *int32 `json:"min,omitempty"`
	// Angle in degrees
	Max *int32 `json:"max,omitempty"`
}

// NewApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth instantiates a new ApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth() *ApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth {
	this := ApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth{}
	return &this
}

// NewApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuthWithDefaults instantiates a new ApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuthWithDefaults() *ApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth {
	this := ApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *ApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth) GetMin() int32 {
	if o == nil || IsNil(o.Min) {
		var ret int32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth) GetMinOk() (*int32, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *ApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given int32 and assigns it to the Min field.
func (o *ApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth) SetMin(v int32) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *ApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth) GetMax() int32 {
	if o == nil || IsNil(o.Max) {
		var ret int32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth) GetMaxOk() (*int32, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *ApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given int32 and assigns it to the Max field.
func (o *ApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth) SetMax(v int32) {
	o.Max = &v
}

func (o ApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	return toSerialize, nil
}

type NullableApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth struct {
	value *ApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth
	isSet bool
}

func (v NullableApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth) Get() *ApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth {
	return v.value
}

func (v *NullableApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth) Set(val *ApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth(val *ApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth) *NullableApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth {
	return &NullableApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth{value: val, isSet: true}
}

func (v NullableApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

