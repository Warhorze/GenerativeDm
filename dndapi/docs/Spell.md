# Spell

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **string** | Resource index for shorthand searching. | [optional] 
**Name** | Pointer to **string** | Name of the referenced resource. | [optional] 
**Url** | Pointer to **string** | URL of the referenced resource. | [optional] 
**Desc** | Pointer to **[]string** | Description of the resource. | [optional] 
**HigherLevel** | Pointer to **[]string** | List of descriptions for casting the spell at higher levels. | [optional] 
**Range** | Pointer to **string** | Range of the spell, usually expressed in feet. | [optional] 
**Components** | Pointer to **[]string** | List of shorthand for required components of the spell. V: verbal S: somatic M: material  | [optional] 
**Material** | Pointer to **string** | Material component for the spell to be cast. | [optional] 
**AreaOfEffect** | Pointer to [**AreaOfEffect**](AreaOfEffect.md) |  | [optional] 
**Ritual** | Pointer to **bool** | Determines if a spell can be cast in a 10-min(in-game) ritual. | [optional] 
**Duration** | Pointer to **string** | How long the spell effect lasts. | [optional] 
**Concentration** | Pointer to **bool** | Determines if a spell needs concentration to persist. | [optional] 
**CastingTime** | Pointer to **string** | How long it takes for the spell to activate. | [optional] 
**Level** | Pointer to **float32** | Level of the spell. | [optional] 
**AttackType** | Pointer to **string** | Attack type of the spell. | [optional] 
**Damage** | Pointer to [**SpellAllOfDamage**](SpellAllOfDamage.md) |  | [optional] 
**School** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 
**Classes** | Pointer to [**[]APIReference**](APIReference.md) | List of classes that are able to learn the spell. | [optional] 
**Subclasses** | Pointer to [**[]APIReference**](APIReference.md) | List of subclasses that have access to the spell. | [optional] 

## Methods

### NewSpell

`func NewSpell() *Spell`

NewSpell instantiates a new Spell object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpellWithDefaults

`func NewSpellWithDefaults() *Spell`

NewSpellWithDefaults instantiates a new Spell object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *Spell) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *Spell) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *Spell) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *Spell) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetName

`func (o *Spell) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Spell) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Spell) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Spell) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *Spell) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Spell) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Spell) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Spell) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDesc

`func (o *Spell) GetDesc() []string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *Spell) GetDescOk() (*[]string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *Spell) SetDesc(v []string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *Spell) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetHigherLevel

`func (o *Spell) GetHigherLevel() []string`

GetHigherLevel returns the HigherLevel field if non-nil, zero value otherwise.

### GetHigherLevelOk

`func (o *Spell) GetHigherLevelOk() (*[]string, bool)`

GetHigherLevelOk returns a tuple with the HigherLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigherLevel

`func (o *Spell) SetHigherLevel(v []string)`

SetHigherLevel sets HigherLevel field to given value.

### HasHigherLevel

`func (o *Spell) HasHigherLevel() bool`

HasHigherLevel returns a boolean if a field has been set.

### GetRange

`func (o *Spell) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *Spell) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *Spell) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *Spell) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetComponents

`func (o *Spell) GetComponents() []string`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *Spell) GetComponentsOk() (*[]string, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *Spell) SetComponents(v []string)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *Spell) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetMaterial

`func (o *Spell) GetMaterial() string`

GetMaterial returns the Material field if non-nil, zero value otherwise.

### GetMaterialOk

`func (o *Spell) GetMaterialOk() (*string, bool)`

GetMaterialOk returns a tuple with the Material field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterial

`func (o *Spell) SetMaterial(v string)`

SetMaterial sets Material field to given value.

### HasMaterial

`func (o *Spell) HasMaterial() bool`

HasMaterial returns a boolean if a field has been set.

### GetAreaOfEffect

`func (o *Spell) GetAreaOfEffect() AreaOfEffect`

GetAreaOfEffect returns the AreaOfEffect field if non-nil, zero value otherwise.

### GetAreaOfEffectOk

`func (o *Spell) GetAreaOfEffectOk() (*AreaOfEffect, bool)`

GetAreaOfEffectOk returns a tuple with the AreaOfEffect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaOfEffect

`func (o *Spell) SetAreaOfEffect(v AreaOfEffect)`

SetAreaOfEffect sets AreaOfEffect field to given value.

### HasAreaOfEffect

`func (o *Spell) HasAreaOfEffect() bool`

HasAreaOfEffect returns a boolean if a field has been set.

### GetRitual

`func (o *Spell) GetRitual() bool`

GetRitual returns the Ritual field if non-nil, zero value otherwise.

### GetRitualOk

`func (o *Spell) GetRitualOk() (*bool, bool)`

GetRitualOk returns a tuple with the Ritual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRitual

`func (o *Spell) SetRitual(v bool)`

SetRitual sets Ritual field to given value.

### HasRitual

`func (o *Spell) HasRitual() bool`

HasRitual returns a boolean if a field has been set.

### GetDuration

`func (o *Spell) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Spell) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Spell) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *Spell) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetConcentration

`func (o *Spell) GetConcentration() bool`

GetConcentration returns the Concentration field if non-nil, zero value otherwise.

### GetConcentrationOk

`func (o *Spell) GetConcentrationOk() (*bool, bool)`

GetConcentrationOk returns a tuple with the Concentration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcentration

`func (o *Spell) SetConcentration(v bool)`

SetConcentration sets Concentration field to given value.

### HasConcentration

`func (o *Spell) HasConcentration() bool`

HasConcentration returns a boolean if a field has been set.

### GetCastingTime

`func (o *Spell) GetCastingTime() string`

GetCastingTime returns the CastingTime field if non-nil, zero value otherwise.

### GetCastingTimeOk

`func (o *Spell) GetCastingTimeOk() (*string, bool)`

GetCastingTimeOk returns a tuple with the CastingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastingTime

`func (o *Spell) SetCastingTime(v string)`

SetCastingTime sets CastingTime field to given value.

### HasCastingTime

`func (o *Spell) HasCastingTime() bool`

HasCastingTime returns a boolean if a field has been set.

### GetLevel

`func (o *Spell) GetLevel() float32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *Spell) GetLevelOk() (*float32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *Spell) SetLevel(v float32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *Spell) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetAttackType

`func (o *Spell) GetAttackType() string`

GetAttackType returns the AttackType field if non-nil, zero value otherwise.

### GetAttackTypeOk

`func (o *Spell) GetAttackTypeOk() (*string, bool)`

GetAttackTypeOk returns a tuple with the AttackType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackType

`func (o *Spell) SetAttackType(v string)`

SetAttackType sets AttackType field to given value.

### HasAttackType

`func (o *Spell) HasAttackType() bool`

HasAttackType returns a boolean if a field has been set.

### GetDamage

`func (o *Spell) GetDamage() SpellAllOfDamage`

GetDamage returns the Damage field if non-nil, zero value otherwise.

### GetDamageOk

`func (o *Spell) GetDamageOk() (*SpellAllOfDamage, bool)`

GetDamageOk returns a tuple with the Damage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamage

`func (o *Spell) SetDamage(v SpellAllOfDamage)`

SetDamage sets Damage field to given value.

### HasDamage

`func (o *Spell) HasDamage() bool`

HasDamage returns a boolean if a field has been set.

### GetSchool

`func (o *Spell) GetSchool() APIReference`

GetSchool returns the School field if non-nil, zero value otherwise.

### GetSchoolOk

`func (o *Spell) GetSchoolOk() (*APIReference, bool)`

GetSchoolOk returns a tuple with the School field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchool

`func (o *Spell) SetSchool(v APIReference)`

SetSchool sets School field to given value.

### HasSchool

`func (o *Spell) HasSchool() bool`

HasSchool returns a boolean if a field has been set.

### GetClasses

`func (o *Spell) GetClasses() []APIReference`

GetClasses returns the Classes field if non-nil, zero value otherwise.

### GetClassesOk

`func (o *Spell) GetClassesOk() (*[]APIReference, bool)`

GetClassesOk returns a tuple with the Classes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClasses

`func (o *Spell) SetClasses(v []APIReference)`

SetClasses sets Classes field to given value.

### HasClasses

`func (o *Spell) HasClasses() bool`

HasClasses returns a boolean if a field has been set.

### GetSubclasses

`func (o *Spell) GetSubclasses() []APIReference`

GetSubclasses returns the Subclasses field if non-nil, zero value otherwise.

### GetSubclassesOk

`func (o *Spell) GetSubclassesOk() (*[]APIReference, bool)`

GetSubclassesOk returns a tuple with the Subclasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubclasses

`func (o *Spell) SetSubclasses(v []APIReference)`

SetSubclasses sets Subclasses field to given value.

### HasSubclasses

`func (o *Spell) HasSubclasses() bool`

HasSubclasses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


