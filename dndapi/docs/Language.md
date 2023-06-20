# Language

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **string** | Resource index for shorthand searching. | [optional] 
**Name** | Pointer to **string** | Name of the referenced resource. | [optional] 
**Url** | Pointer to **string** | URL of the referenced resource. | [optional] 
**Desc** | Pointer to **string** | Brief description of the language. | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Script** | Pointer to **string** | Script used for writing in the language. | [optional] 
**TypicalSpeakers** | Pointer to **[]string** | List of races that tend to speak the language. | [optional] 

## Methods

### NewLanguage

`func NewLanguage() *Language`

NewLanguage instantiates a new Language object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanguageWithDefaults

`func NewLanguageWithDefaults() *Language`

NewLanguageWithDefaults instantiates a new Language object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *Language) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *Language) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *Language) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *Language) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetName

`func (o *Language) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Language) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Language) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Language) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *Language) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Language) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Language) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Language) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDesc

`func (o *Language) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *Language) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *Language) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *Language) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetType

`func (o *Language) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Language) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Language) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Language) HasType() bool`

HasType returns a boolean if a field has been set.

### GetScript

`func (o *Language) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *Language) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *Language) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *Language) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetTypicalSpeakers

`func (o *Language) GetTypicalSpeakers() []string`

GetTypicalSpeakers returns the TypicalSpeakers field if non-nil, zero value otherwise.

### GetTypicalSpeakersOk

`func (o *Language) GetTypicalSpeakersOk() (*[]string, bool)`

GetTypicalSpeakersOk returns a tuple with the TypicalSpeakers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypicalSpeakers

`func (o *Language) SetTypicalSpeakers(v []string)`

SetTypicalSpeakers sets TypicalSpeakers field to given value.

### HasTypicalSpeakers

`func (o *Language) HasTypicalSpeakers() bool`

HasTypicalSpeakers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


