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

// ViolationContextPolicyAllOf struct for ViolationContextPolicyAllOf
type ViolationContextPolicyAllOf struct {
	Type *string `json:"type,omitempty"`
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ViolationContextPolicyAllOf ViolationContextPolicyAllOf

// NewViolationContextPolicyAllOf instantiates a new ViolationContextPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViolationContextPolicyAllOf() *ViolationContextPolicyAllOf {
	this := ViolationContextPolicyAllOf{}
	return &this
}

// NewViolationContextPolicyAllOfWithDefaults instantiates a new ViolationContextPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViolationContextPolicyAllOfWithDefaults() *ViolationContextPolicyAllOf {
	this := ViolationContextPolicyAllOf{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ViolationContextPolicyAllOf) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolationContextPolicyAllOf) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ViolationContextPolicyAllOf) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ViolationContextPolicyAllOf) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ViolationContextPolicyAllOf) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolationContextPolicyAllOf) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ViolationContextPolicyAllOf) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ViolationContextPolicyAllOf) SetName(v string) {
	o.Name = &v
}

func (o ViolationContextPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ViolationContextPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varViolationContextPolicyAllOf := _ViolationContextPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varViolationContextPolicyAllOf); err == nil {
		*o = ViolationContextPolicyAllOf(varViolationContextPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableViolationContextPolicyAllOf struct {
	value *ViolationContextPolicyAllOf
	isSet bool
}

func (v NullableViolationContextPolicyAllOf) Get() *ViolationContextPolicyAllOf {
	return v.value
}

func (v *NullableViolationContextPolicyAllOf) Set(val *ViolationContextPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableViolationContextPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableViolationContextPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViolationContextPolicyAllOf(val *ViolationContextPolicyAllOf) *NullableViolationContextPolicyAllOf {
	return &NullableViolationContextPolicyAllOf{value: val, isSet: true}
}

func (v NullableViolationContextPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViolationContextPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


