# CreateSavedSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the saved search.  | [optional] 
**Description** | Pointer to **NullableString** | The description of the saved search.  | [optional] 
**Public** | Pointer to **bool** | Indicates if the saved search is public.  | [optional] [default to false]
**Created** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**Modified** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**Indices** | [**[]Index**](Index.md) | The names of the Elasticsearch indices in which to search.  | 
**Columns** | Pointer to [**map[string][]Column**](array.md) | The columns to be returned (specifies the order in which they will be presented) for each document type.  The currently supported document types are: _accessprofile_, _accountactivity_, _account_, _aggregation_, _entitlement_, _event_, _identity_, and _role_.  | [optional] 
**Query** | **string** | The search query using Elasticsearch [Query String Query](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/query-dsl-query-string-query.html#query-string) syntax from the Query DSL.  | 
**Fields** | Pointer to **[]string** | The fields to be searched against in a multi-field query.  | [optional] 
**Sort** | Pointer to **[]string** | The fields to be used to sort the search results.  | [optional] 
**Filters** | Pointer to [**NullableSavedSearchDetailFilters**](SavedSearchDetailFilters.md) |  | [optional] 

## Methods

### NewCreateSavedSearchRequest

`func NewCreateSavedSearchRequest(indices []Index, query string, ) *CreateSavedSearchRequest`

NewCreateSavedSearchRequest instantiates a new CreateSavedSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSavedSearchRequestWithDefaults

`func NewCreateSavedSearchRequestWithDefaults() *CreateSavedSearchRequest`

NewCreateSavedSearchRequestWithDefaults instantiates a new CreateSavedSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateSavedSearchRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSavedSearchRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSavedSearchRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateSavedSearchRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateSavedSearchRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSavedSearchRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateSavedSearchRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateSavedSearchRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateSavedSearchRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateSavedSearchRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPublic

`func (o *CreateSavedSearchRequest) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *CreateSavedSearchRequest) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *CreateSavedSearchRequest) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *CreateSavedSearchRequest) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetCreated

`func (o *CreateSavedSearchRequest) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CreateSavedSearchRequest) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CreateSavedSearchRequest) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CreateSavedSearchRequest) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *CreateSavedSearchRequest) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *CreateSavedSearchRequest) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetModified

`func (o *CreateSavedSearchRequest) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *CreateSavedSearchRequest) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *CreateSavedSearchRequest) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *CreateSavedSearchRequest) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModifiedNil

`func (o *CreateSavedSearchRequest) SetModifiedNil(b bool)`

 SetModifiedNil sets the value for Modified to be an explicit nil

### UnsetModified
`func (o *CreateSavedSearchRequest) UnsetModified()`

UnsetModified ensures that no value is present for Modified, not even an explicit nil
### GetIndices

`func (o *CreateSavedSearchRequest) GetIndices() []Index`

GetIndices returns the Indices field if non-nil, zero value otherwise.

### GetIndicesOk

`func (o *CreateSavedSearchRequest) GetIndicesOk() (*[]Index, bool)`

GetIndicesOk returns a tuple with the Indices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndices

`func (o *CreateSavedSearchRequest) SetIndices(v []Index)`

SetIndices sets Indices field to given value.


### GetColumns

`func (o *CreateSavedSearchRequest) GetColumns() map[string][]Column`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *CreateSavedSearchRequest) GetColumnsOk() (*map[string][]Column, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *CreateSavedSearchRequest) SetColumns(v map[string][]Column)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *CreateSavedSearchRequest) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetQuery

`func (o *CreateSavedSearchRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CreateSavedSearchRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CreateSavedSearchRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetFields

`func (o *CreateSavedSearchRequest) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *CreateSavedSearchRequest) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *CreateSavedSearchRequest) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *CreateSavedSearchRequest) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *CreateSavedSearchRequest) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *CreateSavedSearchRequest) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil
### GetSort

`func (o *CreateSavedSearchRequest) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *CreateSavedSearchRequest) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *CreateSavedSearchRequest) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *CreateSavedSearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetFilters

`func (o *CreateSavedSearchRequest) GetFilters() SavedSearchDetailFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *CreateSavedSearchRequest) GetFiltersOk() (*SavedSearchDetailFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *CreateSavedSearchRequest) SetFilters(v SavedSearchDetailFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *CreateSavedSearchRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *CreateSavedSearchRequest) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *CreateSavedSearchRequest) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


