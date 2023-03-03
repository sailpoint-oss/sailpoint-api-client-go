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

// ListAccessProfiles401Response struct for ListAccessProfiles401Response
type ListAccessProfiles401Response struct {
	// A message describing the error
	Error map[string]interface{} `json:"error,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListAccessProfiles401Response ListAccessProfiles401Response

// NewListAccessProfiles401Response instantiates a new ListAccessProfiles401Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAccessProfiles401Response() *ListAccessProfiles401Response {
	this := ListAccessProfiles401Response{}
	return &this
}

// NewListAccessProfiles401ResponseWithDefaults instantiates a new ListAccessProfiles401Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAccessProfiles401ResponseWithDefaults() *ListAccessProfiles401Response {
	this := ListAccessProfiles401Response{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ListAccessProfiles401Response) GetError() map[string]interface{} {
	if o == nil || isNil(o.Error) {
		var ret map[string]interface{}
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAccessProfiles401Response) GetErrorOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Error) {
		return map[string]interface{}{}, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ListAccessProfiles401Response) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given map[string]interface{} and assigns it to the Error field.
func (o *ListAccessProfiles401Response) SetError(v map[string]interface{}) {
	o.Error = v
}

func (o ListAccessProfiles401Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ListAccessProfiles401Response) UnmarshalJSON(bytes []byte) (err error) {
	varListAccessProfiles401Response := _ListAccessProfiles401Response{}

	if err = json.Unmarshal(bytes, &varListAccessProfiles401Response); err == nil {
		*o = ListAccessProfiles401Response(varListAccessProfiles401Response)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "error")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListAccessProfiles401Response struct {
	value *ListAccessProfiles401Response
	isSet bool
}

func (v NullableListAccessProfiles401Response) Get() *ListAccessProfiles401Response {
	return v.value
}

func (v *NullableListAccessProfiles401Response) Set(val *ListAccessProfiles401Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListAccessProfiles401Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListAccessProfiles401Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAccessProfiles401Response(val *ListAccessProfiles401Response) *NullableListAccessProfiles401Response {
	return &NullableListAccessProfiles401Response{value: val, isSet: true}
}

func (v NullableListAccessProfiles401Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAccessProfiles401Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


