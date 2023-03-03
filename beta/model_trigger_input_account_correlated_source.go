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

// TriggerInputAccountCorrelatedSource The source from which the account came from.
type TriggerInputAccountCorrelatedSource struct {
	// ID of the object to which this reference applies
	Id string `json:"id"`
	// The type of object that is referenced
	Type string `json:"type"`
	// Human-readable display name of the object to which this reference applies
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _TriggerInputAccountCorrelatedSource TriggerInputAccountCorrelatedSource

// NewTriggerInputAccountCorrelatedSource instantiates a new TriggerInputAccountCorrelatedSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerInputAccountCorrelatedSource(id string, type_ string, name string) *TriggerInputAccountCorrelatedSource {
	this := TriggerInputAccountCorrelatedSource{}
	this.Id = id
	this.Type = type_
	this.Name = name
	return &this
}

// NewTriggerInputAccountCorrelatedSourceWithDefaults instantiates a new TriggerInputAccountCorrelatedSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerInputAccountCorrelatedSourceWithDefaults() *TriggerInputAccountCorrelatedSource {
	this := TriggerInputAccountCorrelatedSource{}
	return &this
}

// GetId returns the Id field value
func (o *TriggerInputAccountCorrelatedSource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccountCorrelatedSource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TriggerInputAccountCorrelatedSource) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *TriggerInputAccountCorrelatedSource) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccountCorrelatedSource) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TriggerInputAccountCorrelatedSource) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *TriggerInputAccountCorrelatedSource) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccountCorrelatedSource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TriggerInputAccountCorrelatedSource) SetName(v string) {
	o.Name = v
}

func (o TriggerInputAccountCorrelatedSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TriggerInputAccountCorrelatedSource) UnmarshalJSON(bytes []byte) (err error) {
	varTriggerInputAccountCorrelatedSource := _TriggerInputAccountCorrelatedSource{}

	if err = json.Unmarshal(bytes, &varTriggerInputAccountCorrelatedSource); err == nil {
		*o = TriggerInputAccountCorrelatedSource(varTriggerInputAccountCorrelatedSource)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTriggerInputAccountCorrelatedSource struct {
	value *TriggerInputAccountCorrelatedSource
	isSet bool
}

func (v NullableTriggerInputAccountCorrelatedSource) Get() *TriggerInputAccountCorrelatedSource {
	return v.value
}

func (v *NullableTriggerInputAccountCorrelatedSource) Set(val *TriggerInputAccountCorrelatedSource) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerInputAccountCorrelatedSource) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerInputAccountCorrelatedSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerInputAccountCorrelatedSource(val *TriggerInputAccountCorrelatedSource) *NullableTriggerInputAccountCorrelatedSource {
	return &NullableTriggerInputAccountCorrelatedSource{value: val, isSet: true}
}

func (v NullableTriggerInputAccountCorrelatedSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerInputAccountCorrelatedSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


