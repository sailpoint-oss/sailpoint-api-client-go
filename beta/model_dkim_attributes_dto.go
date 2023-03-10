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

// DkimAttributesDto DKIM attributes for a domain / identity
type DkimAttributesDto struct {
	// The identity or domain address
	Address *string `json:"address,omitempty"`
	// Whether or not DKIM has been enabled for this domain / identity
	DkimEnabled *bool `json:"dkimEnabled,omitempty"`
	// The tokens to be added to a DNS for verification
	DkimTokens []string `json:"dkimTokens,omitempty"`
	// The current status if the domain /identity has been verified. Ie Success, Failed, Pending
	DkimVerificationStatus *string `json:"dkimVerificationStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DkimAttributesDto DkimAttributesDto

// NewDkimAttributesDto instantiates a new DkimAttributesDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDkimAttributesDto() *DkimAttributesDto {
	this := DkimAttributesDto{}
	return &this
}

// NewDkimAttributesDtoWithDefaults instantiates a new DkimAttributesDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDkimAttributesDtoWithDefaults() *DkimAttributesDto {
	this := DkimAttributesDto{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *DkimAttributesDto) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DkimAttributesDto) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *DkimAttributesDto) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *DkimAttributesDto) SetAddress(v string) {
	o.Address = &v
}

// GetDkimEnabled returns the DkimEnabled field value if set, zero value otherwise.
func (o *DkimAttributesDto) GetDkimEnabled() bool {
	if o == nil || isNil(o.DkimEnabled) {
		var ret bool
		return ret
	}
	return *o.DkimEnabled
}

// GetDkimEnabledOk returns a tuple with the DkimEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DkimAttributesDto) GetDkimEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.DkimEnabled) {
		return nil, false
	}
	return o.DkimEnabled, true
}

// HasDkimEnabled returns a boolean if a field has been set.
func (o *DkimAttributesDto) HasDkimEnabled() bool {
	if o != nil && !isNil(o.DkimEnabled) {
		return true
	}

	return false
}

// SetDkimEnabled gets a reference to the given bool and assigns it to the DkimEnabled field.
func (o *DkimAttributesDto) SetDkimEnabled(v bool) {
	o.DkimEnabled = &v
}

// GetDkimTokens returns the DkimTokens field value if set, zero value otherwise.
func (o *DkimAttributesDto) GetDkimTokens() []string {
	if o == nil || isNil(o.DkimTokens) {
		var ret []string
		return ret
	}
	return o.DkimTokens
}

// GetDkimTokensOk returns a tuple with the DkimTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DkimAttributesDto) GetDkimTokensOk() ([]string, bool) {
	if o == nil || isNil(o.DkimTokens) {
		return nil, false
	}
	return o.DkimTokens, true
}

// HasDkimTokens returns a boolean if a field has been set.
func (o *DkimAttributesDto) HasDkimTokens() bool {
	if o != nil && !isNil(o.DkimTokens) {
		return true
	}

	return false
}

// SetDkimTokens gets a reference to the given []string and assigns it to the DkimTokens field.
func (o *DkimAttributesDto) SetDkimTokens(v []string) {
	o.DkimTokens = v
}

// GetDkimVerificationStatus returns the DkimVerificationStatus field value if set, zero value otherwise.
func (o *DkimAttributesDto) GetDkimVerificationStatus() string {
	if o == nil || isNil(o.DkimVerificationStatus) {
		var ret string
		return ret
	}
	return *o.DkimVerificationStatus
}

// GetDkimVerificationStatusOk returns a tuple with the DkimVerificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DkimAttributesDto) GetDkimVerificationStatusOk() (*string, bool) {
	if o == nil || isNil(o.DkimVerificationStatus) {
		return nil, false
	}
	return o.DkimVerificationStatus, true
}

// HasDkimVerificationStatus returns a boolean if a field has been set.
func (o *DkimAttributesDto) HasDkimVerificationStatus() bool {
	if o != nil && !isNil(o.DkimVerificationStatus) {
		return true
	}

	return false
}

// SetDkimVerificationStatus gets a reference to the given string and assigns it to the DkimVerificationStatus field.
func (o *DkimAttributesDto) SetDkimVerificationStatus(v string) {
	o.DkimVerificationStatus = &v
}

func (o DkimAttributesDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !isNil(o.DkimEnabled) {
		toSerialize["dkimEnabled"] = o.DkimEnabled
	}
	if !isNil(o.DkimTokens) {
		toSerialize["dkimTokens"] = o.DkimTokens
	}
	if !isNil(o.DkimVerificationStatus) {
		toSerialize["dkimVerificationStatus"] = o.DkimVerificationStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DkimAttributesDto) UnmarshalJSON(bytes []byte) (err error) {
	varDkimAttributesDto := _DkimAttributesDto{}

	if err = json.Unmarshal(bytes, &varDkimAttributesDto); err == nil {
		*o = DkimAttributesDto(varDkimAttributesDto)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "dkimEnabled")
		delete(additionalProperties, "dkimTokens")
		delete(additionalProperties, "dkimVerificationStatus")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDkimAttributesDto struct {
	value *DkimAttributesDto
	isSet bool
}

func (v NullableDkimAttributesDto) Get() *DkimAttributesDto {
	return v.value
}

func (v *NullableDkimAttributesDto) Set(val *DkimAttributesDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDkimAttributesDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDkimAttributesDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDkimAttributesDto(val *DkimAttributesDto) *NullableDkimAttributesDto {
	return &NullableDkimAttributesDto{value: val, isSet: true}
}

func (v NullableDkimAttributesDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDkimAttributesDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


