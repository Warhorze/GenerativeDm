# Race

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **string** | Resource index for shorthand searching. | [optional] 
**Name** | Pointer to **string** | Name of the referenced resource. | [optional] 
**Url** | Pointer to **string** | URL of the referenced resource. | [optional] 
**Speed** | Pointer to **float32** | Base move speed for this race (in feet per round). | [optional] 
**AbilityBonuses** | Pointer to [**[]AbilityBonus**](AbilityBonus.md) | Racial bonuses to ability scores. | [optional] 
**Alignment** | Pointer to **string** | Flavor description of likely alignments this race takes. | [optional] 
**Age** | Pointer to **string** | Flavor description of possible ages for this race. | [optional] 
**Size** | Pointer to **string** | Size class of this race. | [optional] 
**SizeDescription** | Pointer to **string** | Flavor description of height and weight for this race. | [optional] 
**StartingProficiencies** | Pointer to [**[]APIReference**](APIReference.md) | Starting proficiencies for all new characters of this race. | [optional] 
**StartingProficiencyOptions** | Pointer to [**Choice**](Choice.md) |  | [optional] 
**Languages** | Pointer to [**[]APIReference**](APIReference.md) | Starting languages for all new characters of this race. | [optional] 
**LanguageDesc** | Pointer to **string** | Flavor description of the languages this race knows. | [optional] 
**Traits** | Pointer to [**[]APIReference**](APIReference.md) | Racial traits that provide benefits to its members. | [optional] 
**Subraces** | Pointer to [**[]APIReference**](APIReference.md) | All possible subraces that this race includes. | [optional] 

## Methods

### NewRace

`func NewRace() *Race`

NewRace instantiates a new Race object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRaceWithDefaults

`func NewRaceWithDefaults() *Race`

NewRaceWithDefaults instantiates a new Race object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *Race) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *Race) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *Race) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *Race) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetName

`func (o *Race) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Race) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Race) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Race) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *Race) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Race) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Race) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Race) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetSpeed

`func (o *Race) GetSpeed() float32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *Race) GetSpeedOk() (*float32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *Race) SetSpeed(v float32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *Race) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetAbilityBonuses

`func (o *Race) GetAbilityBonuses() []AbilityBonus`

GetAbilityBonuses returns the AbilityBonuses field if non-nil, zero value otherwise.

### GetAbilityBonusesOk

`func (o *Race) GetAbilityBonusesOk() (*[]AbilityBonus, bool)`

GetAbilityBonusesOk returns a tuple with the AbilityBonuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbilityBonuses

`func (o *Race) SetAbilityBonuses(v []AbilityBonus)`

SetAbilityBonuses sets AbilityBonuses field to given value.

### HasAbilityBonuses

`func (o *Race) HasAbilityBonuses() bool`

HasAbilityBonuses returns a boolean if a field has been set.

### GetAlignment

`func (o *Race) GetAlignment() string`

GetAlignment returns the Alignment field if non-nil, zero value otherwise.

### GetAlignmentOk

`func (o *Race) GetAlignmentOk() (*string, bool)`

GetAlignmentOk returns a tuple with the Alignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignment

`func (o *Race) SetAlignment(v string)`

SetAlignment sets Alignment field to given value.

### HasAlignment

`func (o *Race) HasAlignment() bool`

HasAlignment returns a boolean if a field has been set.

### GetAge

`func (o *Race) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *Race) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *Race) SetAge(v string)`

SetAge sets Age field to given value.

### HasAge

`func (o *Race) HasAge() bool`

HasAge returns a boolean if a field has been set.

### GetSize

`func (o *Race) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Race) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Race) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *Race) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSizeDescription

`func (o *Race) GetSizeDescription() string`

GetSizeDescription returns the SizeDescription field if non-nil, zero value otherwise.

### GetSizeDescriptionOk

`func (o *Race) GetSizeDescriptionOk() (*string, bool)`

GetSizeDescriptionOk returns a tuple with the SizeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeDescription

`func (o *Race) SetSizeDescription(v string)`

SetSizeDescription sets SizeDescription field to given value.

### HasSizeDescription

`func (o *Race) HasSizeDescription() bool`

HasSizeDescription returns a boolean if a field has been set.

### GetStartingProficiencies

`func (o *Race) GetStartingProficiencies() []APIReference`

GetStartingProficiencies returns the StartingProficiencies field if non-nil, zero value otherwise.

### GetStartingProficienciesOk

`func (o *Race) GetStartingProficienciesOk() (*[]APIReference, bool)`

GetStartingProficienciesOk returns a tuple with the StartingProficiencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingProficiencies

`func (o *Race) SetStartingProficiencies(v []APIReference)`

SetStartingProficiencies sets StartingProficiencies field to given value.

### HasStartingProficiencies

`func (o *Race) HasStartingProficiencies() bool`

HasStartingProficiencies returns a boolean if a field has been set.

### GetStartingProficiencyOptions

`func (o *Race) GetStartingProficiencyOptions() Choice`

GetStartingProficiencyOptions returns the StartingProficiencyOptions field if non-nil, zero value otherwise.

### GetStartingProficiencyOptionsOk

`func (o *Race) GetStartingProficiencyOptionsOk() (*Choice, bool)`

GetStartingProficiencyOptionsOk returns a tuple with the StartingProficiencyOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingProficiencyOptions

`func (o *Race) SetStartingProficiencyOptions(v Choice)`

SetStartingProficiencyOptions sets StartingProficiencyOptions field to given value.

### HasStartingProficiencyOptions

`func (o *Race) HasStartingProficiencyOptions() bool`

HasStartingProficiencyOptions returns a boolean if a field has been set.

### GetLanguages

`func (o *Race) GetLanguages() []APIReference`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *Race) GetLanguagesOk() (*[]APIReference, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *Race) SetLanguages(v []APIReference)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *Race) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetLanguageDesc

`func (o *Race) GetLanguageDesc() string`

GetLanguageDesc returns the LanguageDesc field if non-nil, zero value otherwise.

### GetLanguageDescOk

`func (o *Race) GetLanguageDescOk() (*string, bool)`

GetLanguageDescOk returns a tuple with the LanguageDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageDesc

`func (o *Race) SetLanguageDesc(v string)`

SetLanguageDesc sets LanguageDesc field to given value.

### HasLanguageDesc

`func (o *Race) HasLanguageDesc() bool`

HasLanguageDesc returns a boolean if a field has been set.

### GetTraits

`func (o *Race) GetTraits() []APIReference`

GetTraits returns the Traits field if non-nil, zero value otherwise.

### GetTraitsOk

`func (o *Race) GetTraitsOk() (*[]APIReference, bool)`

GetTraitsOk returns a tuple with the Traits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraits

`func (o *Race) SetTraits(v []APIReference)`

SetTraits sets Traits field to given value.

### HasTraits

`func (o *Race) HasTraits() bool`

HasTraits returns a boolean if a field has been set.

### GetSubraces

`func (o *Race) GetSubraces() []APIReference`

GetSubraces returns the Subraces field if non-nil, zero value otherwise.

### GetSubracesOk

`func (o *Race) GetSubracesOk() (*[]APIReference, bool)`

GetSubracesOk returns a tuple with the Subraces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubraces

`func (o *Race) SetSubraces(v []APIReference)`

SetSubraces sets Subraces field to given value.

### HasSubraces

`func (o *Race) HasSubraces() bool`

HasSubraces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


