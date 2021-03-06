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

// InstalledAppInfo struct for InstalledAppInfo
type InstalledAppInfo struct {
	// appId
	AppId *int32 `json:"appId,omitempty"`
	// installedAppId
	InstalledAppId *int32 `json:"installedAppId,omitempty"`
	// environment Name
	EnvironmentName *string `json:"environmentName,omitempty"`
	// EA_ONLY/FULL
	AppOfferingMode *string `json:"appOfferingMode,omitempty"`
	// App store chart Id
	AppStoreChartId *float32 `json:"appStoreChartId,omitempty"`
	// App store installed app version Id
	InstalledAppVersionId *float32 `json:"installedAppVersionId,omitempty"`
	// Cluster Id
	ClusterId *float32 `json:"clusterId,omitempty"`
	// Environment Id
	EnvironmentId *float32 `json:"environmentId,omitempty"`
}

// NewInstalledAppInfo instantiates a new InstalledAppInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstalledAppInfo() *InstalledAppInfo {
	this := InstalledAppInfo{}
	return &this
}

// NewInstalledAppInfoWithDefaults instantiates a new InstalledAppInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstalledAppInfoWithDefaults() *InstalledAppInfo {
	this := InstalledAppInfo{}
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *InstalledAppInfo) GetAppId() int32 {
	if o == nil || o.AppId == nil {
		var ret int32
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstalledAppInfo) GetAppIdOk() (*int32, bool) {
	if o == nil || o.AppId == nil {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *InstalledAppInfo) HasAppId() bool {
	if o != nil && o.AppId != nil {
		return true
	}

	return false
}

// SetAppId gets a reference to the given int32 and assigns it to the AppId field.
func (o *InstalledAppInfo) SetAppId(v int32) {
	o.AppId = &v
}

// GetInstalledAppId returns the InstalledAppId field value if set, zero value otherwise.
func (o *InstalledAppInfo) GetInstalledAppId() int32 {
	if o == nil || o.InstalledAppId == nil {
		var ret int32
		return ret
	}
	return *o.InstalledAppId
}

// GetInstalledAppIdOk returns a tuple with the InstalledAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstalledAppInfo) GetInstalledAppIdOk() (*int32, bool) {
	if o == nil || o.InstalledAppId == nil {
		return nil, false
	}
	return o.InstalledAppId, true
}

// HasInstalledAppId returns a boolean if a field has been set.
func (o *InstalledAppInfo) HasInstalledAppId() bool {
	if o != nil && o.InstalledAppId != nil {
		return true
	}

	return false
}

// SetInstalledAppId gets a reference to the given int32 and assigns it to the InstalledAppId field.
func (o *InstalledAppInfo) SetInstalledAppId(v int32) {
	o.InstalledAppId = &v
}

// GetEnvironmentName returns the EnvironmentName field value if set, zero value otherwise.
func (o *InstalledAppInfo) GetEnvironmentName() string {
	if o == nil || o.EnvironmentName == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentName
}

// GetEnvironmentNameOk returns a tuple with the EnvironmentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstalledAppInfo) GetEnvironmentNameOk() (*string, bool) {
	if o == nil || o.EnvironmentName == nil {
		return nil, false
	}
	return o.EnvironmentName, true
}

// HasEnvironmentName returns a boolean if a field has been set.
func (o *InstalledAppInfo) HasEnvironmentName() bool {
	if o != nil && o.EnvironmentName != nil {
		return true
	}

	return false
}

// SetEnvironmentName gets a reference to the given string and assigns it to the EnvironmentName field.
func (o *InstalledAppInfo) SetEnvironmentName(v string) {
	o.EnvironmentName = &v
}

// GetAppOfferingMode returns the AppOfferingMode field value if set, zero value otherwise.
func (o *InstalledAppInfo) GetAppOfferingMode() string {
	if o == nil || o.AppOfferingMode == nil {
		var ret string
		return ret
	}
	return *o.AppOfferingMode
}

// GetAppOfferingModeOk returns a tuple with the AppOfferingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstalledAppInfo) GetAppOfferingModeOk() (*string, bool) {
	if o == nil || o.AppOfferingMode == nil {
		return nil, false
	}
	return o.AppOfferingMode, true
}

// HasAppOfferingMode returns a boolean if a field has been set.
func (o *InstalledAppInfo) HasAppOfferingMode() bool {
	if o != nil && o.AppOfferingMode != nil {
		return true
	}

	return false
}

// SetAppOfferingMode gets a reference to the given string and assigns it to the AppOfferingMode field.
func (o *InstalledAppInfo) SetAppOfferingMode(v string) {
	o.AppOfferingMode = &v
}

// GetAppStoreChartId returns the AppStoreChartId field value if set, zero value otherwise.
func (o *InstalledAppInfo) GetAppStoreChartId() float32 {
	if o == nil || o.AppStoreChartId == nil {
		var ret float32
		return ret
	}
	return *o.AppStoreChartId
}

// GetAppStoreChartIdOk returns a tuple with the AppStoreChartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstalledAppInfo) GetAppStoreChartIdOk() (*float32, bool) {
	if o == nil || o.AppStoreChartId == nil {
		return nil, false
	}
	return o.AppStoreChartId, true
}

// HasAppStoreChartId returns a boolean if a field has been set.
func (o *InstalledAppInfo) HasAppStoreChartId() bool {
	if o != nil && o.AppStoreChartId != nil {
		return true
	}

	return false
}

// SetAppStoreChartId gets a reference to the given float32 and assigns it to the AppStoreChartId field.
func (o *InstalledAppInfo) SetAppStoreChartId(v float32) {
	o.AppStoreChartId = &v
}

// GetInstalledAppVersionId returns the InstalledAppVersionId field value if set, zero value otherwise.
func (o *InstalledAppInfo) GetInstalledAppVersionId() float32 {
	if o == nil || o.InstalledAppVersionId == nil {
		var ret float32
		return ret
	}
	return *o.InstalledAppVersionId
}

// GetInstalledAppVersionIdOk returns a tuple with the InstalledAppVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstalledAppInfo) GetInstalledAppVersionIdOk() (*float32, bool) {
	if o == nil || o.InstalledAppVersionId == nil {
		return nil, false
	}
	return o.InstalledAppVersionId, true
}

// HasInstalledAppVersionId returns a boolean if a field has been set.
func (o *InstalledAppInfo) HasInstalledAppVersionId() bool {
	if o != nil && o.InstalledAppVersionId != nil {
		return true
	}

	return false
}

// SetInstalledAppVersionId gets a reference to the given float32 and assigns it to the InstalledAppVersionId field.
func (o *InstalledAppInfo) SetInstalledAppVersionId(v float32) {
	o.InstalledAppVersionId = &v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *InstalledAppInfo) GetClusterId() float32 {
	if o == nil || o.ClusterId == nil {
		var ret float32
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstalledAppInfo) GetClusterIdOk() (*float32, bool) {
	if o == nil || o.ClusterId == nil {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *InstalledAppInfo) HasClusterId() bool {
	if o != nil && o.ClusterId != nil {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given float32 and assigns it to the ClusterId field.
func (o *InstalledAppInfo) SetClusterId(v float32) {
	o.ClusterId = &v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *InstalledAppInfo) GetEnvironmentId() float32 {
	if o == nil || o.EnvironmentId == nil {
		var ret float32
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstalledAppInfo) GetEnvironmentIdOk() (*float32, bool) {
	if o == nil || o.EnvironmentId == nil {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *InstalledAppInfo) HasEnvironmentId() bool {
	if o != nil && o.EnvironmentId != nil {
		return true
	}

	return false
}

// SetEnvironmentId gets a reference to the given float32 and assigns it to the EnvironmentId field.
func (o *InstalledAppInfo) SetEnvironmentId(v float32) {
	o.EnvironmentId = &v
}

func (o InstalledAppInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppId != nil {
		toSerialize["appId"] = o.AppId
	}
	if o.InstalledAppId != nil {
		toSerialize["installedAppId"] = o.InstalledAppId
	}
	if o.EnvironmentName != nil {
		toSerialize["environmentName"] = o.EnvironmentName
	}
	if o.AppOfferingMode != nil {
		toSerialize["appOfferingMode"] = o.AppOfferingMode
	}
	if o.AppStoreChartId != nil {
		toSerialize["appStoreChartId"] = o.AppStoreChartId
	}
	if o.InstalledAppVersionId != nil {
		toSerialize["installedAppVersionId"] = o.InstalledAppVersionId
	}
	if o.ClusterId != nil {
		toSerialize["clusterId"] = o.ClusterId
	}
	if o.EnvironmentId != nil {
		toSerialize["environmentId"] = o.EnvironmentId
	}
	return json.Marshal(toSerialize)
}

type NullableInstalledAppInfo struct {
	value *InstalledAppInfo
	isSet bool
}

func (v NullableInstalledAppInfo) Get() *InstalledAppInfo {
	return v.value
}

func (v *NullableInstalledAppInfo) Set(val *InstalledAppInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableInstalledAppInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableInstalledAppInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstalledAppInfo(val *InstalledAppInfo) *NullableInstalledAppInfo {
	return &NullableInstalledAppInfo{value: val, isSet: true}
}

func (v NullableInstalledAppInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstalledAppInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


