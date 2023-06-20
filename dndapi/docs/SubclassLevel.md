# SubclassLevel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **string** | Resource index for shorthand searching. | [optional] 
**Url** | Pointer to **string** | URL of the referenced resource. | [optional] 
**Level** | Pointer to **float32** | Number value for the current level object. | [optional] 
**AbilityScoreBonuses** | Pointer to **float32** | Total number of ability score bonuses gained, added from previous levels. | [optional] 
**ProfBonus** | Pointer to **float32** | Proficiency bonus for this class at the specified level. | [optional] 
**Features** | Pointer to [**[]APIReference**](APIReference.md) | List of features gained at this level. | [optional] 
**Spellcasting** | Pointer to [**SubclassLevelSpellcasting**](SubclassLevelSpellcasting.md) |  | [optional] 
**Classspecific** | Pointer to **map[string]interface{}** | Class specific information such as dice values for bard songs and number of warlock invocations. | [optional] 

## Methods

### NewSubclassLevel

`func NewSubclassLevel() *SubclassLevel`

NewSubclassLevel instantiates a new SubclassLevel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubclassLevelWithDefaults

`func NewSubclassLevelWithDefaults() *SubclassLevel`

NewSubclassLevelWithDefaults instantiates a new SubclassLevel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *SubclassLevel) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *SubclassLevel) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *SubclassLevel) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *SubclassLevel) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetUrl

`func (o *SubclassLevel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SubclassLevel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SubclassLevel) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SubclassLevel) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetLevel

`func (o *SubclassLevel) GetLevel() float32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *SubclassLevel) GetLevelOk() (*float32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *SubclassLevel) SetLevel(v float32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *SubclassLevel) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetAbilityScoreBonuses

`func (o *SubclassLevel) GetAbilityScoreBonuses() float32`

GetAbilityScoreBonuses returns the AbilityScoreBonuses field if non-nil, zero value otherwise.

### GetAbilityScoreBonusesOk

`func (o *SubclassLevel) GetAbilityScoreBonusesOk() (*float32, bool)`

GetAbilityScoreBonusesOk returns a tuple with the AbilityScoreBonuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbilityScoreBonuses

`func (o *SubclassLevel) SetAbilityScoreBonuses(v float32)`

SetAbilityScoreBonuses sets AbilityScoreBonuses field to given value.

### HasAbilityScoreBonuses

`func (o *SubclassLevel) HasAbilityScoreBonuses() bool`

HasAbilityScoreBonuses returns a boolean if a field has been set.

### GetProfBonus

`func (o *SubclassLevel) GetProfBonus() float32`

GetProfBonus returns the ProfBonus field if non-nil, zero value otherwise.

### GetProfBonusOk

`func (o *SubclassLevel) GetProfBonusOk() (*float32, bool)`

GetProfBonusOk returns a tuple with the ProfBonus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfBonus

`func (o *SubclassLevel) SetProfBonus(v float32)`

SetProfBonus sets ProfBonus field to given value.

### HasProfBonus

`func (o *SubclassLevel) HasProfBonus() bool`

HasProfBonus returns a boolean if a field has been set.

### GetFeatures

`func (o *SubclassLevel) GetFeatures() []APIReference`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *SubclassLevel) GetFeaturesOk() (*[]APIReference, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *SubclassLevel) SetFeatures(v []APIReference)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *SubclassLevel) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetSpellcasting

`func (o *SubclassLevel) GetSpellcasting() SubclassLevelSpellcasting`

GetSpellcasting returns the Spellcasting field if non-nil, zero value otherwise.

### GetSpellcastingOk

`func (o *SubclassLevel) GetSpellcastingOk() (*SubclassLevelSpellcasting, bool)`

GetSpellcastingOk returns a tuple with the Spellcasting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpellcasting

`func (o *SubclassLevel) SetSpellcasting(v SubclassLevelSpellcasting)`

SetSpellcasting sets Spellcasting field to given value.

### HasSpellcasting

`func (o *SubclassLevel) HasSpellcasting() bool`

HasSpellcasting returns a boolean if a field has been set.

### GetClassspecific

`func (o *SubclassLevel) GetClassspecific() map[string]interface{}`

GetClassspecific returns the Classspecific field if non-nil, zero value otherwise.

### GetClassspecificOk

`func (o *SubclassLevel) GetClassspecificOk() (*map[string]interface{}, bool)`

GetClassspecificOk returns a tuple with the Classspecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassspecific

`func (o *SubclassLevel) SetClassspecific(v map[string]interface{})`

SetClassspecific sets Classspecific field to given value.

### HasClassspecific

`func (o *SubclassLevel) HasClassspecific() bool`

HasClassspecific returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


