/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// LogsAggregateRequestPaging Paging settings
type LogsAggregateRequestPaging struct {
	// The returned paging point to use to get the next results
	After *string `json:"after,omitempty"`
}

// NewLogsAggregateRequestPaging instantiates a new LogsAggregateRequestPaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsAggregateRequestPaging() *LogsAggregateRequestPaging {
	this := LogsAggregateRequestPaging{}
	return &this
}

// NewLogsAggregateRequestPagingWithDefaults instantiates a new LogsAggregateRequestPaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsAggregateRequestPagingWithDefaults() *LogsAggregateRequestPaging {
	this := LogsAggregateRequestPaging{}
	return &this
}

// GetAfter returns the After field value if set, zero value otherwise.
func (o *LogsAggregateRequestPaging) GetAfter() string {
	if o == nil || o.After == nil {
		var ret string
		return ret
	}
	return *o.After
}

// GetAfterOk returns a tuple with the After field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsAggregateRequestPaging) GetAfterOk() (*string, bool) {
	if o == nil || o.After == nil {
		return nil, false
	}
	return o.After, true
}

// HasAfter returns a boolean if a field has been set.
func (o *LogsAggregateRequestPaging) HasAfter() bool {
	if o != nil && o.After != nil {
		return true
	}

	return false
}

// SetAfter gets a reference to the given string and assigns it to the After field.
func (o *LogsAggregateRequestPaging) SetAfter(v string) {
	o.After = &v
}

func (o LogsAggregateRequestPaging) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.After != nil {
		toSerialize["after"] = o.After
	}
	return json.Marshal(toSerialize)
}

type NullableLogsAggregateRequestPaging struct {
	value *LogsAggregateRequestPaging
	isSet bool
}

func (v NullableLogsAggregateRequestPaging) Get() *LogsAggregateRequestPaging {
	return v.value
}

func (v *NullableLogsAggregateRequestPaging) Set(val *LogsAggregateRequestPaging) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsAggregateRequestPaging) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsAggregateRequestPaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsAggregateRequestPaging(val *LogsAggregateRequestPaging) *NullableLogsAggregateRequestPaging {
	return &NullableLogsAggregateRequestPaging{value: val, isSet: true}
}

func (v NullableLogsAggregateRequestPaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsAggregateRequestPaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}