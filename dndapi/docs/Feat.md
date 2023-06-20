# Feat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **string** | Resource index for shorthand searching. | [optional] 
**Name** | Pointer to **string** | Name of the referenced resource. | [optional] 
**Url** | Pointer to **string** | URL of the referenced resource. | [optional] 
**Desc** | Pointer to **[]string** | Description of the resource. | [optional] 
**Prerequisites** | Pointer to [**[]Prerequisite**](Prerequisite.md) | An object of APIReferences to ability scores and minimum scores. | [optional] 

## Methods

### NewFeat

`func NewFeat() *Feat`

NewFeat instantiates a new Feat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatWithDefaults

`func NewFeatWithDefaults() *Feat`

NewFeatWithDefaults instantiates a new Feat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *Feat) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *Feat) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *Feat) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *Feat) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetName

`func (o *Feat) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Feat) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Feat) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Feat) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *Feat) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Feat) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Feat) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Feat) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDesc

`func (o *Feat) GetDesc() []string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *Feat) GetDescOk() (*[]string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *Feat) SetDesc(v []string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *Feat) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetPrerequisites

`func (o *Feat) GetPrerequisites() []Prerequisite`

GetPrerequisites returns the Prerequisites field if non-nil, zero value otherwise.

### GetPrerequisitesOk

`func (o *Feat) GetPrerequisitesOk() (*[]Prerequisite, bool)`

GetPrerequisitesOk returns a tuple with the Prerequisites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrerequisites

`func (o *Feat) SetPrerequisites(v []Prerequisite)`

SetPrerequisites sets Prerequisites field to given value.

### HasPrerequisites

`func (o *Feat) HasPrerequisites() bool`

HasPrerequisites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


