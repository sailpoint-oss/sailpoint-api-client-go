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

// AccountRequestInfo If an account activity item is associated with an access request, captures details of that request.
type AccountRequestInfo struct {
	// Id of requested object
	RequestedObjectId *string `json:"requestedObjectId,omitempty"`
	// Human-readable name of requested object
	RequestedObjectName *string `json:"requestedObjectName,omitempty"`
	RequestedObjectType *RequestableObjectType `json:"requestedObjectType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountRequestInfo AccountRequestInfo

// NewAccountRequestInfo instantiates a new AccountRequestInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountRequestInfo() *AccountRequestInfo {
	this := AccountRequestInfo{}
	return &this
}

// NewAccountRequestInfoWithDefaults instantiates a new AccountRequestInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountRequestInfoWithDefaults() *AccountRequestInfo {
	this := AccountRequestInfo{}
	return &this
}

// GetRequestedObjectId returns the RequestedObjectId field value if set, zero value otherwise.
func (o *AccountRequestInfo) GetRequestedObjectId() string {
	if o == nil || isNil(o.RequestedObjectId) {
		var ret string
		return ret
	}
	return *o.RequestedObjectId
}

// GetRequestedObjectIdOk returns a tuple with the RequestedObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountRequestInfo) GetRequestedObjectIdOk() (*string, bool) {
	if o == nil || isNil(o.RequestedObjectId) {
		return nil, false
	}
	return o.RequestedObjectId, true
}

// HasRequestedObjectId returns a boolean if a field has been set.
func (o *AccountRequestInfo) HasRequestedObjectId() bool {
	if o != nil && !isNil(o.RequestedObjectId) {
		return true
	}

	return false
}

// SetRequestedObjectId gets a reference to the given string and assigns it to the RequestedObjectId field.
func (o *AccountRequestInfo) SetRequestedObjectId(v string) {
	o.RequestedObjectId = &v
}

// GetRequestedObjectName returns the RequestedObjectName field value if set, zero value otherwise.
func (o *AccountRequestInfo) GetRequestedObjectName() string {
	if o == nil || isNil(o.RequestedObjectName) {
		var ret string
		return ret
	}
	return *o.RequestedObjectName
}

// GetRequestedObjectNameOk returns a tuple with the RequestedObjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountRequestInfo) GetRequestedObjectNameOk() (*string, bool) {
	if o == nil || isNil(o.RequestedObjectName) {
		return nil, false
	}
	return o.RequestedObjectName, true
}

// HasRequestedObjectName returns a boolean if a field has been set.
func (o *AccountRequestInfo) HasRequestedObjectName() bool {
	if o != nil && !isNil(o.RequestedObjectName) {
		return true
	}

	return false
}

// SetRequestedObjectName gets a reference to the given string and assigns it to the RequestedObjectName field.
func (o *AccountRequestInfo) SetRequestedObjectName(v string) {
	o.RequestedObjectName = &v
}

// GetRequestedObjectType returns the RequestedObjectType field value if set, zero value otherwise.
func (o *AccountRequestInfo) GetRequestedObjectType() RequestableObjectType {
	if o == nil || isNil(o.RequestedObjectType) {
		var ret RequestableObjectType
		return ret
	}
	return *o.RequestedObjectType
}

// GetRequestedObjectTypeOk returns a tuple with the RequestedObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountRequestInfo) GetRequestedObjectTypeOk() (*RequestableObjectType, bool) {
	if o == nil || isNil(o.RequestedObjectType) {
		return nil, false
	}
	return o.RequestedObjectType, true
}

// HasRequestedObjectType returns a boolean if a field has been set.
func (o *AccountRequestInfo) HasRequestedObjectType() bool {
	if o != nil && !isNil(o.RequestedObjectType) {
		return true
	}

	return false
}

// SetRequestedObjectType gets a reference to the given RequestableObjectType and assigns it to the RequestedObjectType field.
func (o *AccountRequestInfo) SetRequestedObjectType(v RequestableObjectType) {
	o.RequestedObjectType = &v
}

func (o AccountRequestInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RequestedObjectId) {
		toSerialize["requestedObjectId"] = o.RequestedObjectId
	}
	if !isNil(o.RequestedObjectName) {
		toSerialize["requestedObjectName"] = o.RequestedObjectName
	}
	if !isNil(o.RequestedObjectType) {
		toSerialize["requestedObjectType"] = o.RequestedObjectType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccountRequestInfo) UnmarshalJSON(bytes []byte) (err error) {
	varAccountRequestInfo := _AccountRequestInfo{}

	if err = json.Unmarshal(bytes, &varAccountRequestInfo); err == nil {
		*o = AccountRequestInfo(varAccountRequestInfo)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "requestedObjectId")
		delete(additionalProperties, "requestedObjectName")
		delete(additionalProperties, "requestedObjectType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountRequestInfo struct {
	value *AccountRequestInfo
	isSet bool
}

func (v NullableAccountRequestInfo) Get() *AccountRequestInfo {
	return v.value
}

func (v *NullableAccountRequestInfo) Set(val *AccountRequestInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountRequestInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountRequestInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountRequestInfo(val *AccountRequestInfo) *NullableAccountRequestInfo {
	return &NullableAccountRequestInfo{value: val, isSet: true}
}

func (v NullableAccountRequestInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountRequestInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


