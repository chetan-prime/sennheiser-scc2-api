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

// checks if the ApiAudioInputsMicrophoneExclusionZonesIdPutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAudioInputsMicrophoneExclusionZonesIdPutRequest{}

// ApiAudioInputsMicrophoneExclusionZonesIdPutRequest struct for ApiAudioInputsMicrophoneExclusionZonesIdPutRequest
type ApiAudioInputsMicrophoneExclusionZonesIdPutRequest struct {
	Enabled *bool `json:"enabled,omitempty"`
	Elevation *ApiAudioInputsMicrophoneExclusionZonesIdGet200ResponseElevation `json:"elevation,omitempty"`
	Azimuth *ApiAudioInputsMicrophoneExclusionZonesIdGet200ResponseAzimuth `json:"azimuth,omitempty"`
}

// NewApiAudioInputsMicrophoneExclusionZonesIdPutRequest instantiates a new ApiAudioInputsMicrophoneExclusionZonesIdPutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAudioInputsMicrophoneExclusionZonesIdPutRequest() *ApiAudioInputsMicrophoneExclusionZonesIdPutRequest {
	this := ApiAudioInputsMicrophoneExclusionZonesIdPutRequest{}
	return &this
}

// NewApiAudioInputsMicrophoneExclusionZonesIdPutRequestWithDefaults instantiates a new ApiAudioInputsMicrophoneExclusionZonesIdPutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAudioInputsMicrophoneExclusionZonesIdPutRequestWithDefaults() *ApiAudioInputsMicrophoneExclusionZonesIdPutRequest {
	this := ApiAudioInputsMicrophoneExclusionZonesIdPutRequest{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApiAudioInputsMicrophoneExclusionZonesIdPutRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAudioInputsMicrophoneExclusionZonesIdPutRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApiAudioInputsMicrophoneExclusionZonesIdPutRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApiAudioInputsMicrophoneExclusionZonesIdPutRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetElevation returns the Elevation field value if set, zero value otherwise.
func (o *ApiAudioInputsMicrophoneExclusionZonesIdPutRequest) GetElevation() ApiAudioInputsMicrophoneExclusionZonesIdGet200ResponseElevation {
	if o == nil || IsNil(o.Elevation) {
		var ret ApiAudioInputsMicrophoneExclusionZonesIdGet200ResponseElevation
		return ret
	}
	return *o.Elevation
}

// GetElevationOk returns a tuple with the Elevation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAudioInputsMicrophoneExclusionZonesIdPutRequest) GetElevationOk() (*ApiAudioInputsMicrophoneExclusionZonesIdGet200ResponseElevation, bool) {
	if o == nil || IsNil(o.Elevation) {
		return nil, false
	}
	return o.Elevation, true
}

// HasElevation returns a boolean if a field has been set.
func (o *ApiAudioInputsMicrophoneExclusionZonesIdPutRequest) HasElevation() bool {
	if o != nil && !IsNil(o.Elevation) {
		return true
	}

	return false
}

// SetElevation gets a reference to the given ApiAudioInputsMicrophoneExclusionZonesIdGet200ResponseElevation and assigns it to the Elevation field.
func (o *ApiAudioInputsMicrophoneExclusionZonesIdPutRequest) SetElevation(v ApiAudioInputsMicrophoneExclusionZonesIdGet200ResponseElevation) {
	o.Elevation = &v
}

// GetAzimuth returns the Azimuth field value if set, zero value otherwise.
func (o *ApiAudioInputsMicrophoneExclusionZonesIdPutRequest) GetAzimuth() ApiAudioInputsMicrophoneExclusionZonesIdGet200ResponseAzimuth {
	if o == nil || IsNil(o.Azimuth) {
		var ret ApiAudioInputsMicrophoneExclusionZonesIdGet200ResponseAzimuth
		return ret
	}
	return *o.Azimuth
}

// GetAzimuthOk returns a tuple with the Azimuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAudioInputsMicrophoneExclusionZonesIdPutRequest) GetAzimuthOk() (*ApiAudioInputsMicrophoneExclusionZonesIdGet200ResponseAzimuth, bool) {
	if o == nil || IsNil(o.Azimuth) {
		return nil, false
	}
	return o.Azimuth, true
}

// HasAzimuth returns a boolean if a field has been set.
func (o *ApiAudioInputsMicrophoneExclusionZonesIdPutRequest) HasAzimuth() bool {
	if o != nil && !IsNil(o.Azimuth) {
		return true
	}

	return false
}

// SetAzimuth gets a reference to the given ApiAudioInputsMicrophoneExclusionZonesIdGet200ResponseAzimuth and assigns it to the Azimuth field.
func (o *ApiAudioInputsMicrophoneExclusionZonesIdPutRequest) SetAzimuth(v ApiAudioInputsMicrophoneExclusionZonesIdGet200ResponseAzimuth) {
	o.Azimuth = &v
}

func (o ApiAudioInputsMicrophoneExclusionZonesIdPutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAudioInputsMicrophoneExclusionZonesIdPutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Elevation) {
		toSerialize["elevation"] = o.Elevation
	}
	if !IsNil(o.Azimuth) {
		toSerialize["azimuth"] = o.Azimuth
	}
	return toSerialize, nil
}

type NullableApiAudioInputsMicrophoneExclusionZonesIdPutRequest struct {
	value *ApiAudioInputsMicrophoneExclusionZonesIdPutRequest
	isSet bool
}

func (v NullableApiAudioInputsMicrophoneExclusionZonesIdPutRequest) Get() *ApiAudioInputsMicrophoneExclusionZonesIdPutRequest {
	return v.value
}

func (v *NullableApiAudioInputsMicrophoneExclusionZonesIdPutRequest) Set(val *ApiAudioInputsMicrophoneExclusionZonesIdPutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAudioInputsMicrophoneExclusionZonesIdPutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAudioInputsMicrophoneExclusionZonesIdPutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAudioInputsMicrophoneExclusionZonesIdPutRequest(val *ApiAudioInputsMicrophoneExclusionZonesIdPutRequest) *NullableApiAudioInputsMicrophoneExclusionZonesIdPutRequest {
	return &NullableApiAudioInputsMicrophoneExclusionZonesIdPutRequest{value: val, isSet: true}
}

func (v NullableApiAudioInputsMicrophoneExclusionZonesIdPutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAudioInputsMicrophoneExclusionZonesIdPutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

