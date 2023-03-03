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

// RoleMiningSessionStatus struct for RoleMiningSessionStatus
type RoleMiningSessionStatus struct {
	// The role mining session status. Can be one of these states - CREATED|UPDATED|IDENTITIES_OBTAINED|PRUNE_THRESHOLD_OBTAINED|POTENTIAL_ROLES_PROCESSING|POTENTIAL_ROLES_CREATED
	State *string `json:"state,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleMiningSessionStatus RoleMiningSessionStatus

// NewRoleMiningSessionStatus instantiates a new RoleMiningSessionStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleMiningSessionStatus() *RoleMiningSessionStatus {
	this := RoleMiningSessionStatus{}
	return &this
}

// NewRoleMiningSessionStatusWithDefaults instantiates a new RoleMiningSessionStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleMiningSessionStatusWithDefaults() *RoleMiningSessionStatus {
	this := RoleMiningSessionStatus{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *RoleMiningSessionStatus) GetState() string {
	if o == nil || isNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningSessionStatus) GetStateOk() (*string, bool) {
	if o == nil || isNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *RoleMiningSessionStatus) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *RoleMiningSessionStatus) SetState(v string) {
	o.State = &v
}

func (o RoleMiningSessionStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RoleMiningSessionStatus) UnmarshalJSON(bytes []byte) (err error) {
	varRoleMiningSessionStatus := _RoleMiningSessionStatus{}

	if err = json.Unmarshal(bytes, &varRoleMiningSessionStatus); err == nil {
		*o = RoleMiningSessionStatus(varRoleMiningSessionStatus)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "state")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleMiningSessionStatus struct {
	value *RoleMiningSessionStatus
	isSet bool
}

func (v NullableRoleMiningSessionStatus) Get() *RoleMiningSessionStatus {
	return v.value
}

func (v *NullableRoleMiningSessionStatus) Set(val *RoleMiningSessionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleMiningSessionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleMiningSessionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleMiningSessionStatus(val *RoleMiningSessionStatus) *NullableRoleMiningSessionStatus {
	return &NullableRoleMiningSessionStatus{value: val, isSet: true}
}

func (v NullableRoleMiningSessionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleMiningSessionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


