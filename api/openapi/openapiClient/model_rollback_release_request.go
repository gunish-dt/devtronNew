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

// RollbackReleaseRequest struct for RollbackReleaseRequest
type RollbackReleaseRequest struct {
	// Installed App Id if the app is installed from chart store
	InstalledAppId *int32 `json:"installedAppId,omitempty"`
	// Installed App Version Id if the app is installed from chart store
	InstalledAppVersionId *int32 `json:"installedAppVersionId,omitempty"`
	// helm App Id if the application is installed from using helm (for example \"clusterId|namespace|appName\" )
	HAppId *string `json:"hAppId,omitempty"`
	// rollback to this version
	Version *int32 `json:"version,omitempty"`
}

// NewRollbackReleaseRequest instantiates a new RollbackReleaseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRollbackReleaseRequest() *RollbackReleaseRequest {
	this := RollbackReleaseRequest{}
	return &this
}

// NewRollbackReleaseRequestWithDefaults instantiates a new RollbackReleaseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRollbackReleaseRequestWithDefaults() *RollbackReleaseRequest {
	this := RollbackReleaseRequest{}
	return &this
}

// GetInstalledAppId returns the InstalledAppId field value if set, zero value otherwise.
func (o *RollbackReleaseRequest) GetInstalledAppId() int32 {
	if o == nil || o.InstalledAppId == nil {
		var ret int32
		return ret
	}
	return *o.InstalledAppId
}

// GetInstalledAppIdOk returns a tuple with the InstalledAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RollbackReleaseRequest) GetInstalledAppIdOk() (*int32, bool) {
	if o == nil || o.InstalledAppId == nil {
		return nil, false
	}
	return o.InstalledAppId, true
}

// HasInstalledAppId returns a boolean if a field has been set.
func (o *RollbackReleaseRequest) HasInstalledAppId() bool {
	if o != nil && o.InstalledAppId != nil {
		return true
	}

	return false
}

// SetInstalledAppId gets a reference to the given int32 and assigns it to the InstalledAppId field.
func (o *RollbackReleaseRequest) SetInstalledAppId(v int32) {
	o.InstalledAppId = &v
}

// GetInstalledAppVersionId returns the InstalledAppVersionId field value if set, zero value otherwise.
func (o *RollbackReleaseRequest) GetInstalledAppVersionId() int32 {
	if o == nil || o.InstalledAppVersionId == nil {
		var ret int32
		return ret
	}
	return *o.InstalledAppVersionId
}

// GetInstalledAppVersionIdOk returns a tuple with the InstalledAppVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RollbackReleaseRequest) GetInstalledAppVersionIdOk() (*int32, bool) {
	if o == nil || o.InstalledAppVersionId == nil {
		return nil, false
	}
	return o.InstalledAppVersionId, true
}

// HasInstalledAppVersionId returns a boolean if a field has been set.
func (o *RollbackReleaseRequest) HasInstalledAppVersionId() bool {
	if o != nil && o.InstalledAppVersionId != nil {
		return true
	}

	return false
}

// SetInstalledAppVersionId gets a reference to the given int32 and assigns it to the InstalledAppVersionId field.
func (o *RollbackReleaseRequest) SetInstalledAppVersionId(v int32) {
	o.InstalledAppVersionId = &v
}

// GetHAppId returns the HAppId field value if set, zero value otherwise.
func (o *RollbackReleaseRequest) GetHAppId() string {
	if o == nil || o.HAppId == nil {
		var ret string
		return ret
	}
	return *o.HAppId
}

// GetHAppIdOk returns a tuple with the HAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RollbackReleaseRequest) GetHAppIdOk() (*string, bool) {
	if o == nil || o.HAppId == nil {
		return nil, false
	}
	return o.HAppId, true
}

// HasHAppId returns a boolean if a field has been set.
func (o *RollbackReleaseRequest) HasHAppId() bool {
	if o != nil && o.HAppId != nil {
		return true
	}

	return false
}

// SetHAppId gets a reference to the given string and assigns it to the HAppId field.
func (o *RollbackReleaseRequest) SetHAppId(v string) {
	o.HAppId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *RollbackReleaseRequest) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RollbackReleaseRequest) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *RollbackReleaseRequest) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *RollbackReleaseRequest) SetVersion(v int32) {
	o.Version = &v
}

func (o RollbackReleaseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InstalledAppId != nil {
		toSerialize["installedAppId"] = o.InstalledAppId
	}
	if o.InstalledAppVersionId != nil {
		toSerialize["installedAppVersionId"] = o.InstalledAppVersionId
	}
	if o.HAppId != nil {
		toSerialize["hAppId"] = o.HAppId
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableRollbackReleaseRequest struct {
	value *RollbackReleaseRequest
	isSet bool
}

func (v NullableRollbackReleaseRequest) Get() *RollbackReleaseRequest {
	return v.value
}

func (v *NullableRollbackReleaseRequest) Set(val *RollbackReleaseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRollbackReleaseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRollbackReleaseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRollbackReleaseRequest(val *RollbackReleaseRequest) *NullableRollbackReleaseRequest {
	return &NullableRollbackReleaseRequest{value: val, isSet: true}
}

func (v NullableRollbackReleaseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRollbackReleaseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


