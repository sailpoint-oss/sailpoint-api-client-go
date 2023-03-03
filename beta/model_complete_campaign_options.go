/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"encoding/json"
)

// CompleteCampaignOptions struct for CompleteCampaignOptions
type CompleteCampaignOptions struct {
	// Determines whether to auto-approve(APPROVE) or auto-revoke(REVOKE) upon campaign completion.
	AutoCompleteAction *string `json:"autoCompleteAction,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CompleteCampaignOptions CompleteCampaignOptions

// NewCompleteCampaignOptions instantiates a new CompleteCampaignOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompleteCampaignOptions() *CompleteCampaignOptions {
	this := CompleteCampaignOptions{}
	var autoCompleteAction string = "APPROVE"
	this.AutoCompleteAction = &autoCompleteAction
	return &this
}

// NewCompleteCampaignOptionsWithDefaults instantiates a new CompleteCampaignOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompleteCampaignOptionsWithDefaults() *CompleteCampaignOptions {
	this := CompleteCampaignOptions{}
	var autoCompleteAction string = "APPROVE"
	this.AutoCompleteAction = &autoCompleteAction
	return &this
}

// GetAutoCompleteAction returns the AutoCompleteAction field value if set, zero value otherwise.
func (o *CompleteCampaignOptions) GetAutoCompleteAction() string {
	if o == nil || isNil(o.AutoCompleteAction) {
		var ret string
		return ret
	}
	return *o.AutoCompleteAction
}

// GetAutoCompleteActionOk returns a tuple with the AutoCompleteAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompleteCampaignOptions) GetAutoCompleteActionOk() (*string, bool) {
	if o == nil || isNil(o.AutoCompleteAction) {
		return nil, false
	}
	return o.AutoCompleteAction, true
}

// HasAutoCompleteAction returns a boolean if a field has been set.
func (o *CompleteCampaignOptions) HasAutoCompleteAction() bool {
	if o != nil && !isNil(o.AutoCompleteAction) {
		return true
	}

	return false
}

// SetAutoCompleteAction gets a reference to the given string and assigns it to the AutoCompleteAction field.
func (o *CompleteCampaignOptions) SetAutoCompleteAction(v string) {
	o.AutoCompleteAction = &v
}

func (o CompleteCampaignOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AutoCompleteAction) {
		toSerialize["autoCompleteAction"] = o.AutoCompleteAction
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CompleteCampaignOptions) UnmarshalJSON(bytes []byte) (err error) {
	varCompleteCampaignOptions := _CompleteCampaignOptions{}

	if err = json.Unmarshal(bytes, &varCompleteCampaignOptions); err == nil {
		*o = CompleteCampaignOptions(varCompleteCampaignOptions)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "autoCompleteAction")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCompleteCampaignOptions struct {
	value *CompleteCampaignOptions
	isSet bool
}

func (v NullableCompleteCampaignOptions) Get() *CompleteCampaignOptions {
	return v.value
}

func (v *NullableCompleteCampaignOptions) Set(val *CompleteCampaignOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableCompleteCampaignOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableCompleteCampaignOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompleteCampaignOptions(val *CompleteCampaignOptions) *NullableCompleteCampaignOptions {
	return &NullableCompleteCampaignOptions{value: val, isSet: true}
}

func (v NullableCompleteCampaignOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompleteCampaignOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


