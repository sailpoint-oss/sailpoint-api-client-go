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

// TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner struct for TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner
type TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner struct {
	// The unique ID of the access item being requested.
	Id string `json:"id"`
	// The human friendly name of the access item.
	Name string `json:"name"`
	// Detailed description of the access item.
	Description NullableString `json:"description,omitempty"`
	// The type of access item.
	Type map[string]interface{} `json:"type"`
	// The action to perform on the access item.
	Operation map[string]interface{} `json:"operation"`
	// A comment from the identity requesting the access.
	Comment NullableString `json:"comment,omitempty"`
	// Additional customer defined metadata about the access item.
	ClientMetadata map[string]interface{} `json:"clientMetadata,omitempty"`
	// A list of one or more approvers for the access request.
	ApprovalInfo []TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner `json:"approvalInfo"`
	AdditionalProperties map[string]interface{}
}

type _TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner

// NewTriggerInputAccessRequestPostApprovalRequestedItemsStatusInner instantiates a new TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerInputAccessRequestPostApprovalRequestedItemsStatusInner(id string, name string, type_ map[string]interface{}, operation map[string]interface{}, approvalInfo []TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner {
	this := TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner{}
	this.Id = id
	this.Name = name
	this.Type = type_
	this.Operation = operation
	this.ApprovalInfo = approvalInfo
	return &this
}

// NewTriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerWithDefaults instantiates a new TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerWithDefaults() *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner {
	this := TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner{}
	return &this
}

// GetId returns the Id field value
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) GetDescription() string {
	if o == nil || isNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) UnsetDescription() {
	o.Description.Unset()
}

// GetType returns the Type field value
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) GetType() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) GetTypeOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Type, true
}

// SetType sets field value
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) SetType(v map[string]interface{}) {
	o.Type = v
}

// GetOperation returns the Operation field value
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) GetOperation() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) GetOperationOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Operation, true
}

// SetOperation sets field value
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) SetOperation(v map[string]interface{}) {
	o.Operation = v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) GetComment() string {
	if o == nil || isNil(o.Comment.Get()) {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) UnsetComment() {
	o.Comment.Unset()
}

// GetClientMetadata returns the ClientMetadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) GetClientMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ClientMetadata
}

// GetClientMetadataOk returns a tuple with the ClientMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) GetClientMetadataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.ClientMetadata) {
		return map[string]interface{}{}, false
	}
	return o.ClientMetadata, true
}

// HasClientMetadata returns a boolean if a field has been set.
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) HasClientMetadata() bool {
	if o != nil && isNil(o.ClientMetadata) {
		return true
	}

	return false
}

// SetClientMetadata gets a reference to the given map[string]interface{} and assigns it to the ClientMetadata field.
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) SetClientMetadata(v map[string]interface{}) {
	o.ClientMetadata = v
}

// GetApprovalInfo returns the ApprovalInfo field value
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) GetApprovalInfo() []TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner {
	if o == nil {
		var ret []TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner
		return ret
	}

	return o.ApprovalInfo
}

// GetApprovalInfoOk returns a tuple with the ApprovalInfo field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) GetApprovalInfoOk() ([]TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApprovalInfo, true
}

// SetApprovalInfo sets field value
func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) SetApprovalInfo(v []TriggerInputAccessRequestPostApprovalRequestedItemsStatusInnerApprovalInfoInner) {
	o.ApprovalInfo = v
}

func (o TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["operation"] = o.Operation
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.ClientMetadata != nil {
		toSerialize["clientMetadata"] = o.ClientMetadata
	}
	if true {
		toSerialize["approvalInfo"] = o.ApprovalInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) UnmarshalJSON(bytes []byte) (err error) {
	varTriggerInputAccessRequestPostApprovalRequestedItemsStatusInner := _TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner{}

	if err = json.Unmarshal(bytes, &varTriggerInputAccessRequestPostApprovalRequestedItemsStatusInner); err == nil {
		*o = TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner(varTriggerInputAccessRequestPostApprovalRequestedItemsStatusInner)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "type")
		delete(additionalProperties, "operation")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "clientMetadata")
		delete(additionalProperties, "approvalInfo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTriggerInputAccessRequestPostApprovalRequestedItemsStatusInner struct {
	value *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner
	isSet bool
}

func (v NullableTriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) Get() *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner {
	return v.value
}

func (v *NullableTriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) Set(val *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerInputAccessRequestPostApprovalRequestedItemsStatusInner(val *TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) *NullableTriggerInputAccessRequestPostApprovalRequestedItemsStatusInner {
	return &NullableTriggerInputAccessRequestPostApprovalRequestedItemsStatusInner{value: val, isSet: true}
}

func (v NullableTriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerInputAccessRequestPostApprovalRequestedItemsStatusInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


