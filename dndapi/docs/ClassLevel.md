# ClassLevel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **string** | Resource index for shorthand searching. | [optional] 
**Url** | Pointer to **string** | URL of the referenced resource. | [optional] 
**Level** | Pointer to **float32** | The number value for the current level object. | [optional] 
**AbilityScoreBonuses** | Pointer to **float32** | Total number of ability score bonuses gained, added from previous levels. | [optional] 
**ProfBonus** | Pointer to **float32** | Proficiency bonus for this class at the specified level. | [optional] 
**Features** | Pointer to [**[]APIReference**](APIReference.md) | Features automatically gained at this level. | [optional] 
**Spellcasting** | Pointer to [**SubclassLevelSpellcasting**](SubclassLevelSpellcasting.md) |  | [optional] 
**ClassSpecific** | Pointer to [**ClassLevelClassSpecific**](ClassLevelClassSpecific.md) |  | [optional] 

## Methods

### NewClassLevel

`func NewClassLevel() *ClassLevel`

NewClassLevel instantiates a new ClassLevel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassLevelWithDefaults

`func NewClassLevelWithDefaults() *ClassLevel`

NewClassLevelWithDefaults instantiates a new ClassLevel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *ClassLevel) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ClassLevel) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ClassLevel) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *ClassLevel) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetUrl

`func (o *ClassLevel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ClassLevel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ClassLevel) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ClassLevel) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetLevel

`func (o *ClassLevel) GetLevel() float32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ClassLevel) GetLevelOk() (*float32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ClassLevel) SetLevel(v float32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *ClassLevel) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetAbilityScoreBonuses

`func (o *ClassLevel) GetAbilityScoreBonuses() float32`

GetAbilityScoreBonuses returns the AbilityScoreBonuses field if non-nil, zero value otherwise.

### GetAbilityScoreBonusesOk

`func (o *ClassLevel) GetAbilityScoreBonusesOk() (*float32, bool)`

GetAbilityScoreBonusesOk returns a tuple with the AbilityScoreBonuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbilityScoreBonuses

`func (o *ClassLevel) SetAbilityScoreBonuses(v float32)`

SetAbilityScoreBonuses sets AbilityScoreBonuses field to given value.

### HasAbilityScoreBonuses

`func (o *ClassLevel) HasAbilityScoreBonuses() bool`

HasAbilityScoreBonuses returns a boolean if a field has been set.

### GetProfBonus

`func (o *ClassLevel) GetProfBonus() float32`

GetProfBonus returns the ProfBonus field if non-nil, zero value otherwise.

### GetProfBonusOk

`func (o *ClassLevel) GetProfBonusOk() (*float32, bool)`

GetProfBonusOk returns a tuple with the ProfBonus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfBonus

`func (o *ClassLevel) SetProfBonus(v float32)`

SetProfBonus sets ProfBonus field to given value.

### HasProfBonus

`func (o *ClassLevel) HasProfBonus() bool`

HasProfBonus returns a boolean if a field has been set.

### GetFeatures

`func (o *ClassLevel) GetFeatures() []APIReference`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ClassLevel) GetFeaturesOk() (*[]APIReference, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ClassLevel) SetFeatures(v []APIReference)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *ClassLevel) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetSpellcasting

`func (o *ClassLevel) GetSpellcasting() SubclassLevelSpellcasting`

GetSpellcasting returns the Spellcasting field if non-nil, zero value otherwise.

### GetSpellcastingOk

`func (o *ClassLevel) GetSpellcastingOk() (*SubclassLevelSpellcasting, bool)`

GetSpellcastingOk returns a tuple with the Spellcasting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpellcasting

`func (o *ClassLevel) SetSpellcasting(v SubclassLevelSpellcasting)`

SetSpellcasting sets Spellcasting field to given value.

### HasSpellcasting

`func (o *ClassLevel) HasSpellcasting() bool`

HasSpellcasting returns a boolean if a field has been set.

### GetClassSpecific

`func (o *ClassLevel) GetClassSpecific() ClassLevelClassSpecific`

GetClassSpecific returns the ClassSpecific field if non-nil, zero value otherwise.

### GetClassSpecificOk

`func (o *ClassLevel) GetClassSpecificOk() (*ClassLevelClassSpecific, bool)`

GetClassSpecificOk returns a tuple with the ClassSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassSpecific

`func (o *ClassLevel) SetClassSpecific(v ClassLevelClassSpecific)`

SetClassSpecific sets ClassSpecific field to given value.

### HasClassSpecific

`func (o *ClassLevel) HasClassSpecific() bool`

HasClassSpecific returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


