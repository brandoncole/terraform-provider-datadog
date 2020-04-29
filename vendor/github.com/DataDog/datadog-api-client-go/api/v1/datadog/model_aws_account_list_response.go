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

// AWSAccountListResponse TODO.
type AWSAccountListResponse struct {
	// TODO.
	Accounts *[]AWSAccount `json:"accounts,omitempty"`
}

// NewAWSAccountListResponse instantiates a new AWSAccountListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAWSAccountListResponse() *AWSAccountListResponse {
	this := AWSAccountListResponse{}
	return &this
}

// NewAWSAccountListResponseWithDefaults instantiates a new AWSAccountListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSAccountListResponseWithDefaults() *AWSAccountListResponse {
	this := AWSAccountListResponse{}
	return &this
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *AWSAccountListResponse) GetAccounts() []AWSAccount {
	if o == nil || o.Accounts == nil {
		var ret []AWSAccount
		return ret
	}
	return *o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountListResponse) GetAccountsOk() (*[]AWSAccount, bool) {
	if o == nil || o.Accounts == nil {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *AWSAccountListResponse) HasAccounts() bool {
	if o != nil && o.Accounts != nil {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []AWSAccount and assigns it to the Accounts field.
func (o *AWSAccountListResponse) SetAccounts(v []AWSAccount) {
	o.Accounts = &v
}

func (o AWSAccountListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Accounts != nil {
		toSerialize["accounts"] = o.Accounts
	}
	return json.Marshal(toSerialize)
}

type NullableAWSAccountListResponse struct {
	value *AWSAccountListResponse
	isSet bool
}

func (v NullableAWSAccountListResponse) Get() *AWSAccountListResponse {
	return v.value
}

func (v *NullableAWSAccountListResponse) Set(val *AWSAccountListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAWSAccountListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAWSAccountListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAWSAccountListResponse(val *AWSAccountListResponse) *NullableAWSAccountListResponse {
	return &NullableAWSAccountListResponse{value: val, isSet: true}
}

func (v NullableAWSAccountListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAWSAccountListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
