# Trait

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **string** | Resource index for shorthand searching. | [optional] 
**Name** | Pointer to **string** | Name of the referenced resource. | [optional] 
**Url** | Pointer to **string** | URL of the referenced resource. | [optional] 
**Desc** | Pointer to **[]string** | Description of the resource. | [optional] 
**Races** | Pointer to [**[]APIReference**](APIReference.md) | List of &#x60;Races&#x60; that have access to the trait. | [optional] 
**Subraces** | Pointer to [**[]APIReference**](APIReference.md) | List of &#x60;Subraces&#x60; that have access to the trait. | [optional] 
**Proficiencies** | Pointer to [**[]APIReference**](APIReference.md) | List of &#x60;Proficiencies&#x60; this trait grants. | [optional] 
**ProficiencyChoices** | Pointer to [**Choice**](Choice.md) |  | [optional] 
**LanguageOptions** | Pointer to [**Choice**](Choice.md) |  | [optional] 
**TraitSpecific** | Pointer to [**TraitAllOfTraitSpecific**](TraitAllOfTraitSpecific.md) |  | [optional] 

## Methods

### NewTrait

`func NewTrait() *Trait`

NewTrait instantiates a new Trait object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraitWithDefaults

`func NewTraitWithDefaults() *Trait`

NewTraitWithDefaults instantiates a new Trait object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *Trait) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *Trait) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *Trait) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *Trait) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetName

`func (o *Trait) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Trait) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Trait) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Trait) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *Trait) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Trait) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Trait) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Trait) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDesc

`func (o *Trait) GetDesc() []string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *Trait) GetDescOk() (*[]string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *Trait) SetDesc(v []string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *Trait) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetRaces

`func (o *Trait) GetRaces() []APIReference`

GetRaces returns the Races field if non-nil, zero value otherwise.

### GetRacesOk

`func (o *Trait) GetRacesOk() (*[]APIReference, bool)`

GetRacesOk returns a tuple with the Races field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaces

`func (o *Trait) SetRaces(v []APIReference)`

SetRaces sets Races field to given value.

### HasRaces

`func (o *Trait) HasRaces() bool`

HasRaces returns a boolean if a field has been set.

### GetSubraces

`func (o *Trait) GetSubraces() []APIReference`

GetSubraces returns the Subraces field if non-nil, zero value otherwise.

### GetSubracesOk

`func (o *Trait) GetSubracesOk() (*[]APIReference, bool)`

GetSubracesOk returns a tuple with the Subraces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubraces

`func (o *Trait) SetSubraces(v []APIReference)`

SetSubraces sets Subraces field to given value.

### HasSubraces

`func (o *Trait) HasSubraces() bool`

HasSubraces returns a boolean if a field has been set.

### GetProficiencies

`func (o *Trait) GetProficiencies() []APIReference`

GetProficiencies returns the Proficiencies field if non-nil, zero value otherwise.

### GetProficienciesOk

`func (o *Trait) GetProficienciesOk() (*[]APIReference, bool)`

GetProficienciesOk returns a tuple with the Proficiencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProficiencies

`func (o *Trait) SetProficiencies(v []APIReference)`

SetProficiencies sets Proficiencies field to given value.

### HasProficiencies

`func (o *Trait) HasProficiencies() bool`

HasProficiencies returns a boolean if a field has been set.

### GetProficiencyChoices

`func (o *Trait) GetProficiencyChoices() Choice`

GetProficiencyChoices returns the ProficiencyChoices field if non-nil, zero value otherwise.

### GetProficiencyChoicesOk

`func (o *Trait) GetProficiencyChoicesOk() (*Choice, bool)`

GetProficiencyChoicesOk returns a tuple with the ProficiencyChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProficiencyChoices

`func (o *Trait) SetProficiencyChoices(v Choice)`

SetProficiencyChoices sets ProficiencyChoices field to given value.

### HasProficiencyChoices

`func (o *Trait) HasProficiencyChoices() bool`

HasProficiencyChoices returns a boolean if a field has been set.

### GetLanguageOptions

`func (o *Trait) GetLanguageOptions() Choice`

GetLanguageOptions returns the LanguageOptions field if non-nil, zero value otherwise.

### GetLanguageOptionsOk

`func (o *Trait) GetLanguageOptionsOk() (*Choice, bool)`

GetLanguageOptionsOk returns a tuple with the LanguageOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageOptions

`func (o *Trait) SetLanguageOptions(v Choice)`

SetLanguageOptions sets LanguageOptions field to given value.

### HasLanguageOptions

`func (o *Trait) HasLanguageOptions() bool`

HasLanguageOptions returns a boolean if a field has been set.

### GetTraitSpecific

`func (o *Trait) GetTraitSpecific() TraitAllOfTraitSpecific`

GetTraitSpecific returns the TraitSpecific field if non-nil, zero value otherwise.

### GetTraitSpecificOk

`func (o *Trait) GetTraitSpecificOk() (*TraitAllOfTraitSpecific, bool)`

GetTraitSpecificOk returns a tuple with the TraitSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraitSpecific

`func (o *Trait) SetTraitSpecific(v TraitAllOfTraitSpecific)`

SetTraitSpecific sets TraitSpecific field to given value.

### HasTraitSpecific

`func (o *Trait) HasTraitSpecific() bool`

HasTraitSpecific returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


