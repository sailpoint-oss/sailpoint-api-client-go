/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
)

// ServiceDeskIntegrationDtoAllOfOwnerRef Reference to the identity that is the owner of this Service Desk integration
type ServiceDeskIntegrationDtoAllOfOwnerRef struct {
	// The type of object being referenced
	Type *string `json:"type,omitempty"`
	// ID of the identity
	Id *string `json:"id,omitempty"`
	// Human-readable display name of the identity
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceDeskIntegrationDtoAllOfOwnerRef ServiceDeskIntegrationDtoAllOfOwnerRef

// NewServiceDeskIntegrationDtoAllOfOwnerRef instantiates a new ServiceDeskIntegrationDtoAllOfOwnerRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceDeskIntegrationDtoAllOfOwnerRef() *ServiceDeskIntegrationDtoAllOfOwnerRef {
	this := ServiceDeskIntegrationDtoAllOfOwnerRef{}
	return &this
}

// NewServiceDeskIntegrationDtoAllOfOwnerRefWithDefaults instantiates a new ServiceDeskIntegrationDtoAllOfOwnerRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceDeskIntegrationDtoAllOfOwnerRefWithDefaults() *ServiceDeskIntegrationDtoAllOfOwnerRef {
	this := ServiceDeskIntegrationDtoAllOfOwnerRef{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ServiceDeskIntegrationDtoAllOfOwnerRef) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDeskIntegrationDtoAllOfOwnerRef) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ServiceDeskIntegrationDtoAllOfOwnerRef) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ServiceDeskIntegrationDtoAllOfOwnerRef) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServiceDeskIntegrationDtoAllOfOwnerRef) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDeskIntegrationDtoAllOfOwnerRef) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceDeskIntegrationDtoAllOfOwnerRef) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServiceDeskIntegrationDtoAllOfOwnerRef) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceDeskIntegrationDtoAllOfOwnerRef) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDeskIntegrationDtoAllOfOwnerRef) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceDeskIntegrationDtoAllOfOwnerRef) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceDeskIntegrationDtoAllOfOwnerRef) SetName(v string) {
	o.Name = &v
}

func (o ServiceDeskIntegrationDtoAllOfOwnerRef) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServiceDeskIntegrationDtoAllOfOwnerRef) UnmarshalJSON(bytes []byte) (err error) {
	varServiceDeskIntegrationDtoAllOfOwnerRef := _ServiceDeskIntegrationDtoAllOfOwnerRef{}

	if err = json.Unmarshal(bytes, &varServiceDeskIntegrationDtoAllOfOwnerRef); err == nil {
		*o = ServiceDeskIntegrationDtoAllOfOwnerRef(varServiceDeskIntegrationDtoAllOfOwnerRef)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceDeskIntegrationDtoAllOfOwnerRef struct {
	value *ServiceDeskIntegrationDtoAllOfOwnerRef
	isSet bool
}

func (v NullableServiceDeskIntegrationDtoAllOfOwnerRef) Get() *ServiceDeskIntegrationDtoAllOfOwnerRef {
	return v.value
}

func (v *NullableServiceDeskIntegrationDtoAllOfOwnerRef) Set(val *ServiceDeskIntegrationDtoAllOfOwnerRef) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceDeskIntegrationDtoAllOfOwnerRef) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceDeskIntegrationDtoAllOfOwnerRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceDeskIntegrationDtoAllOfOwnerRef(val *ServiceDeskIntegrationDtoAllOfOwnerRef) *NullableServiceDeskIntegrationDtoAllOfOwnerRef {
	return &NullableServiceDeskIntegrationDtoAllOfOwnerRef{value: val, isSet: true}
}

func (v NullableServiceDeskIntegrationDtoAllOfOwnerRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceDeskIntegrationDtoAllOfOwnerRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


