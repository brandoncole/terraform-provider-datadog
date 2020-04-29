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

// LogsIndexesOrder Object containing the ordered list of log index names.
type LogsIndexesOrder struct {
	// Array of strings identifying by their name(s) the index(es) of your organization. Logs are tested against the query filter of each index one by one, following the order of the array. Logs are eventually stored in the first matching index.
	IndexNames []string `json:"index_names"`
}

// NewLogsIndexesOrder instantiates a new LogsIndexesOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsIndexesOrder(indexNames []string) *LogsIndexesOrder {
	this := LogsIndexesOrder{}
	this.IndexNames = indexNames
	return &this
}

// NewLogsIndexesOrderWithDefaults instantiates a new LogsIndexesOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsIndexesOrderWithDefaults() *LogsIndexesOrder {
	this := LogsIndexesOrder{}
	return &this
}

// GetIndexNames returns the IndexNames field value
func (o *LogsIndexesOrder) GetIndexNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.IndexNames
}

// GetIndexNamesOk returns a tuple with the IndexNames field value
// and a boolean to check if the value has been set.
func (o *LogsIndexesOrder) GetIndexNamesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndexNames, true
}

// SetIndexNames sets field value
func (o *LogsIndexesOrder) SetIndexNames(v []string) {
	o.IndexNames = v
}

func (o LogsIndexesOrder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["index_names"] = o.IndexNames
	}
	return json.Marshal(toSerialize)
}

type NullableLogsIndexesOrder struct {
	value *LogsIndexesOrder
	isSet bool
}

func (v NullableLogsIndexesOrder) Get() *LogsIndexesOrder {
	return v.value
}

func (v *NullableLogsIndexesOrder) Set(val *LogsIndexesOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsIndexesOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsIndexesOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsIndexesOrder(val *LogsIndexesOrder) *NullableLogsIndexesOrder {
	return &NullableLogsIndexesOrder{value: val, isSet: true}
}

func (v NullableLogsIndexesOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsIndexesOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
