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

// PasswordChangeRequest struct for PasswordChangeRequest
type PasswordChangeRequest struct {
	// The identity ID that requested the password change
	IdentityId *string `json:"identityId,omitempty"`
	// The RSA encrypted password
	EncryptedPassword *string `json:"encryptedPassword,omitempty"`
	// The encryption key ID
	PublicKeyId *string `json:"publicKeyId,omitempty"`
	// Account ID of the account This is specified per account schema in the source configuration. It is used to distinguish accounts. More info can be found here https://community.sailpoint.com/t5/IdentityNow-Connectors/How-do-I-designate-an-account-attribute-as-the-Account-ID-for-a/ta-p/80350
	AccountId *string `json:"accountId,omitempty"`
	// The ID of the source for which identity is requesting the password change
	SourceId *string `json:"sourceId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordChangeRequest PasswordChangeRequest

// NewPasswordChangeRequest instantiates a new PasswordChangeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordChangeRequest() *PasswordChangeRequest {
	this := PasswordChangeRequest{}
	return &this
}

// NewPasswordChangeRequestWithDefaults instantiates a new PasswordChangeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordChangeRequestWithDefaults() *PasswordChangeRequest {
	this := PasswordChangeRequest{}
	return &this
}

// GetIdentityId returns the IdentityId field value if set, zero value otherwise.
func (o *PasswordChangeRequest) GetIdentityId() string {
	if o == nil || isNil(o.IdentityId) {
		var ret string
		return ret
	}
	return *o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordChangeRequest) GetIdentityIdOk() (*string, bool) {
	if o == nil || isNil(o.IdentityId) {
		return nil, false
	}
	return o.IdentityId, true
}

// HasIdentityId returns a boolean if a field has been set.
func (o *PasswordChangeRequest) HasIdentityId() bool {
	if o != nil && !isNil(o.IdentityId) {
		return true
	}

	return false
}

// SetIdentityId gets a reference to the given string and assigns it to the IdentityId field.
func (o *PasswordChangeRequest) SetIdentityId(v string) {
	o.IdentityId = &v
}

// GetEncryptedPassword returns the EncryptedPassword field value if set, zero value otherwise.
func (o *PasswordChangeRequest) GetEncryptedPassword() string {
	if o == nil || isNil(o.EncryptedPassword) {
		var ret string
		return ret
	}
	return *o.EncryptedPassword
}

// GetEncryptedPasswordOk returns a tuple with the EncryptedPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordChangeRequest) GetEncryptedPasswordOk() (*string, bool) {
	if o == nil || isNil(o.EncryptedPassword) {
		return nil, false
	}
	return o.EncryptedPassword, true
}

// HasEncryptedPassword returns a boolean if a field has been set.
func (o *PasswordChangeRequest) HasEncryptedPassword() bool {
	if o != nil && !isNil(o.EncryptedPassword) {
		return true
	}

	return false
}

// SetEncryptedPassword gets a reference to the given string and assigns it to the EncryptedPassword field.
func (o *PasswordChangeRequest) SetEncryptedPassword(v string) {
	o.EncryptedPassword = &v
}

// GetPublicKeyId returns the PublicKeyId field value if set, zero value otherwise.
func (o *PasswordChangeRequest) GetPublicKeyId() string {
	if o == nil || isNil(o.PublicKeyId) {
		var ret string
		return ret
	}
	return *o.PublicKeyId
}

// GetPublicKeyIdOk returns a tuple with the PublicKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordChangeRequest) GetPublicKeyIdOk() (*string, bool) {
	if o == nil || isNil(o.PublicKeyId) {
		return nil, false
	}
	return o.PublicKeyId, true
}

// HasPublicKeyId returns a boolean if a field has been set.
func (o *PasswordChangeRequest) HasPublicKeyId() bool {
	if o != nil && !isNil(o.PublicKeyId) {
		return true
	}

	return false
}

// SetPublicKeyId gets a reference to the given string and assigns it to the PublicKeyId field.
func (o *PasswordChangeRequest) SetPublicKeyId(v string) {
	o.PublicKeyId = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *PasswordChangeRequest) GetAccountId() string {
	if o == nil || isNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordChangeRequest) GetAccountIdOk() (*string, bool) {
	if o == nil || isNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *PasswordChangeRequest) HasAccountId() bool {
	if o != nil && !isNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *PasswordChangeRequest) SetAccountId(v string) {
	o.AccountId = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *PasswordChangeRequest) GetSourceId() string {
	if o == nil || isNil(o.SourceId) {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordChangeRequest) GetSourceIdOk() (*string, bool) {
	if o == nil || isNil(o.SourceId) {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *PasswordChangeRequest) HasSourceId() bool {
	if o != nil && !isNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *PasswordChangeRequest) SetSourceId(v string) {
	o.SourceId = &v
}

func (o PasswordChangeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IdentityId) {
		toSerialize["identityId"] = o.IdentityId
	}
	if !isNil(o.EncryptedPassword) {
		toSerialize["encryptedPassword"] = o.EncryptedPassword
	}
	if !isNil(o.PublicKeyId) {
		toSerialize["publicKeyId"] = o.PublicKeyId
	}
	if !isNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !isNil(o.SourceId) {
		toSerialize["sourceId"] = o.SourceId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PasswordChangeRequest) UnmarshalJSON(bytes []byte) (err error) {
	varPasswordChangeRequest := _PasswordChangeRequest{}

	if err = json.Unmarshal(bytes, &varPasswordChangeRequest); err == nil {
		*o = PasswordChangeRequest(varPasswordChangeRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "identityId")
		delete(additionalProperties, "encryptedPassword")
		delete(additionalProperties, "publicKeyId")
		delete(additionalProperties, "accountId")
		delete(additionalProperties, "sourceId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePasswordChangeRequest struct {
	value *PasswordChangeRequest
	isSet bool
}

func (v NullablePasswordChangeRequest) Get() *PasswordChangeRequest {
	return v.value
}

func (v *NullablePasswordChangeRequest) Set(val *PasswordChangeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordChangeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordChangeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordChangeRequest(val *PasswordChangeRequest) *NullablePasswordChangeRequest {
	return &NullablePasswordChangeRequest{value: val, isSet: true}
}

func (v NullablePasswordChangeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordChangeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


