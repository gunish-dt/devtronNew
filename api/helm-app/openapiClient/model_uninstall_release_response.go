/*
Devtron Labs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// UninstallReleaseResponse struct for UninstallReleaseResponse
type UninstallReleaseResponse struct {
	// success or failure
	Success *bool `json:"success,omitempty"`
}

// NewUninstallReleaseResponse instantiates a new UninstallReleaseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUninstallReleaseResponse() *UninstallReleaseResponse {
	this := UninstallReleaseResponse{}
	return &this
}

// NewUninstallReleaseResponseWithDefaults instantiates a new UninstallReleaseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUninstallReleaseResponseWithDefaults() *UninstallReleaseResponse {
	this := UninstallReleaseResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *UninstallReleaseResponse) GetSuccess() bool {
	if o == nil || o.Success == nil {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UninstallReleaseResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || o.Success == nil {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *UninstallReleaseResponse) HasSuccess() bool {
	if o != nil && o.Success != nil {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *UninstallReleaseResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o UninstallReleaseResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Success != nil {
		toSerialize["success"] = o.Success
	}
	return json.Marshal(toSerialize)
}

type NullableUninstallReleaseResponse struct {
	value *UninstallReleaseResponse
	isSet bool
}

func (v NullableUninstallReleaseResponse) Get() *UninstallReleaseResponse {
	return v.value
}

func (v *NullableUninstallReleaseResponse) Set(val *UninstallReleaseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUninstallReleaseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUninstallReleaseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUninstallReleaseResponse(val *UninstallReleaseResponse) *NullableUninstallReleaseResponse {
	return &NullableUninstallReleaseResponse{value: val, isSet: true}
}

func (v NullableUninstallReleaseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUninstallReleaseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


