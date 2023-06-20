# Subrace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **string** | Resource index for shorthand searching. | [optional] 
**Name** | Pointer to **string** | Name of the referenced resource. | [optional] 
**Url** | Pointer to **string** | URL of the referenced resource. | [optional] 
**Desc** | Pointer to **string** | Description of the subrace. | [optional] 
**Race** | Pointer to [**SubraceAllOfRace**](SubraceAllOfRace.md) |  | [optional] 
**AbilityBonuses** | Pointer to [**[]AbilityBonus**](AbilityBonus.md) | Additional ability bonuses for the subrace. | [optional] 
**StartingProficiencies** | Pointer to [**[]APIReference**](APIReference.md) | Starting proficiencies for all new characters of the subrace. | [optional] 
**Languages** | Pointer to [**[]APIReference**](APIReference.md) | Starting languages for all new characters of the subrace. | [optional] 
**LanguageOptions** | Pointer to [**Choice**](Choice.md) |  | [optional] 
**RacialTraits** | Pointer to [**[]APIReference**](APIReference.md) | List of traits that for the subrace. | [optional] 

## Methods

### NewSubrace

`func NewSubrace() *Subrace`

NewSubrace instantiates a new Subrace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubraceWithDefaults

`func NewSubraceWithDefaults() *Subrace`

NewSubraceWithDefaults instantiates a new Subrace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *Subrace) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *Subrace) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *Subrace) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *Subrace) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetName

`func (o *Subrace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Subrace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Subrace) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Subrace) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *Subrace) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Subrace) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Subrace) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Subrace) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDesc

`func (o *Subrace) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *Subrace) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *Subrace) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *Subrace) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetRace

`func (o *Subrace) GetRace() SubraceAllOfRace`

GetRace returns the Race field if non-nil, zero value otherwise.

### GetRaceOk

`func (o *Subrace) GetRaceOk() (*SubraceAllOfRace, bool)`

GetRaceOk returns a tuple with the Race field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRace

`func (o *Subrace) SetRace(v SubraceAllOfRace)`

SetRace sets Race field to given value.

### HasRace

`func (o *Subrace) HasRace() bool`

HasRace returns a boolean if a field has been set.

### GetAbilityBonuses

`func (o *Subrace) GetAbilityBonuses() []AbilityBonus`

GetAbilityBonuses returns the AbilityBonuses field if non-nil, zero value otherwise.

### GetAbilityBonusesOk

`func (o *Subrace) GetAbilityBonusesOk() (*[]AbilityBonus, bool)`

GetAbilityBonusesOk returns a tuple with the AbilityBonuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbilityBonuses

`func (o *Subrace) SetAbilityBonuses(v []AbilityBonus)`

SetAbilityBonuses sets AbilityBonuses field to given value.

### HasAbilityBonuses

`func (o *Subrace) HasAbilityBonuses() bool`

HasAbilityBonuses returns a boolean if a field has been set.

### GetStartingProficiencies

`func (o *Subrace) GetStartingProficiencies() []APIReference`

GetStartingProficiencies returns the StartingProficiencies field if non-nil, zero value otherwise.

### GetStartingProficienciesOk

`func (o *Subrace) GetStartingProficienciesOk() (*[]APIReference, bool)`

GetStartingProficienciesOk returns a tuple with the StartingProficiencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingProficiencies

`func (o *Subrace) SetStartingProficiencies(v []APIReference)`

SetStartingProficiencies sets StartingProficiencies field to given value.

### HasStartingProficiencies

`func (o *Subrace) HasStartingProficiencies() bool`

HasStartingProficiencies returns a boolean if a field has been set.

### GetLanguages

`func (o *Subrace) GetLanguages() []APIReference`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *Subrace) GetLanguagesOk() (*[]APIReference, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *Subrace) SetLanguages(v []APIReference)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *Subrace) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetLanguageOptions

`func (o *Subrace) GetLanguageOptions() Choice`

GetLanguageOptions returns the LanguageOptions field if non-nil, zero value otherwise.

### GetLanguageOptionsOk

`func (o *Subrace) GetLanguageOptionsOk() (*Choice, bool)`

GetLanguageOptionsOk returns a tuple with the LanguageOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageOptions

`func (o *Subrace) SetLanguageOptions(v Choice)`

SetLanguageOptions sets LanguageOptions field to given value.

### HasLanguageOptions

`func (o *Subrace) HasLanguageOptions() bool`

HasLanguageOptions returns a boolean if a field has been set.

### GetRacialTraits

`func (o *Subrace) GetRacialTraits() []APIReference`

GetRacialTraits returns the RacialTraits field if non-nil, zero value otherwise.

### GetRacialTraitsOk

`func (o *Subrace) GetRacialTraitsOk() (*[]APIReference, bool)`

GetRacialTraitsOk returns a tuple with the RacialTraits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRacialTraits

`func (o *Subrace) SetRacialTraits(v []APIReference)`

SetRacialTraits sets RacialTraits field to given value.

### HasRacialTraits

`func (o *Subrace) HasRacialTraits() bool`

HasRacialTraits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


