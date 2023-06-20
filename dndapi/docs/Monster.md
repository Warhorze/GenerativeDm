# Monster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **string** | Resource index for shorthand searching. | [optional] 
**Name** | Pointer to **string** | Name of the referenced resource. | [optional] 
**Url** | Pointer to **string** | URL of the referenced resource. | [optional] 
**Desc** | Pointer to **[]string** | Description of the resource. | [optional] 
**Charisma** | Pointer to **float32** | A monster&#39;s ability to charm or intimidate a player. | [optional] 
**Constitution** | Pointer to **float32** | How sturdy a monster is.\&quot; | [optional] 
**Dexterity** | Pointer to **float32** | The monster&#39;s ability for swift movement or stealth | [optional] 
**Intelligence** | Pointer to **float32** | The monster&#39;s ability to outsmart a player. | [optional] 
**Strength** | Pointer to **float32** | How hard a monster can hit a player. | [optional] 
**Wisdom** | Pointer to **float32** | A monster&#39;s ability to ascertain the player&#39;s plan. | [optional] 
**Image** | Pointer to **string** | The image url of the monster. | [optional] 
**Size** | Pointer to **string** | The size of the monster ranging from Tiny to Gargantuan.\&quot; | [optional] 
**Type** | Pointer to **string** | The type of monster. | [optional] 
**Subtype** | Pointer to **string** | The sub-category of a monster used for classification of monsters.\&quot; | [optional] 
**Alignments** | Pointer to **string** | A creature&#39;s general moral and personal attitudes. | [optional] 
**ArmorClass** | Pointer to [**[]MonsterArmorClass**](MonsterArmorClass.md) | The difficulty for a player to successfully deal damage to a monster. | [optional] 
**HitPoints** | Pointer to **float32** | The hit points of a monster determine how much damage it is able to take before it can be defeated. | [optional] 
**HitDice** | Pointer to **string** | The hit die of a monster can be used to make a version of the same monster whose hit points are determined by the roll of the die. For example: A monster with 2d6 would have its hit points determine by rolling a 6 sided die twice. | [optional] 
**HitPointsRoll** | Pointer to **string** | The roll for determining a monster&#39;s hit points, which consists of the hit dice (e.g. 18d10) and the modifier determined by its Constitution (e.g. +36). For example, 18d10+36 | [optional] 
**Actions** | Pointer to [**[]MonsterAction**](MonsterAction.md) | A list of actions that are available to the monster to take during combat. | [optional] 
**LegendaryActions** | Pointer to [**[]MonsterAction**](MonsterAction.md) | A list of legendary actions that are available to the monster to take during combat. | [optional] 
**ChallengeRating** | Pointer to **float32** | A monster&#39;s challenge rating is a guideline number that says when a monster becomes an appropriate challenge against the party&#39;s average level. For example. A group of 4 players with an average level of 4 would have an appropriate combat challenge against a monster with a challenge rating of 4 but a monster with a challenge rating of 8 against the same group of players would pose a significant threat. | [optional] 
**ConditionImmunities** | Pointer to [**[]APIReference**](APIReference.md) | A list of conditions that a monster is immune to. | [optional] 
**DamageImmunities** | Pointer to **[]string** | A list of damage types that a monster will take double damage from. | [optional] 
**DamageResistances** | Pointer to **[]string** | A list of damage types that a monster will take half damage from. | [optional] 
**DamageVulnerabilities** | Pointer to **[]string** | A list of damage types that a monster will take double damage from. | [optional] 
**Forms** | Pointer to [**[]APIReference**](APIReference.md) | List of other related monster entries that are of the same form. Only applicable to Lycanthropes that have multiple forms. | [optional] 
**Languages** | Pointer to **string** | The languages a monster is able to speak. | [optional] 
**Proficiencies** | Pointer to [**[]MonsterProficiency**](MonsterProficiency.md) | A list of proficiencies of a monster. | [optional] 
**Reactions** | Pointer to [**[]MonsterAction**](MonsterAction.md) | A list of reactions that is available to the monster to take during combat. | [optional] 
**Senses** | Pointer to [**MonsterAllOfSenses**](MonsterAllOfSenses.md) |  | [optional] 
**SpecialAbilities** | Pointer to [**[]MonsterSpecialAbility**](MonsterSpecialAbility.md) | A list of the monster&#39;s special abilities. | [optional] 
**Speed** | Pointer to [**MonsterAllOfSpeed**](MonsterAllOfSpeed.md) |  | [optional] 
**Xp** | Pointer to **float32** | The number of experience points (XP) a monster is worth is based on its challenge rating. | [optional] 

## Methods

### NewMonster

`func NewMonster() *Monster`

NewMonster instantiates a new Monster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonsterWithDefaults

`func NewMonsterWithDefaults() *Monster`

NewMonsterWithDefaults instantiates a new Monster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *Monster) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *Monster) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *Monster) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *Monster) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetName

`func (o *Monster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Monster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Monster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Monster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *Monster) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Monster) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Monster) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Monster) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDesc

`func (o *Monster) GetDesc() []string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *Monster) GetDescOk() (*[]string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *Monster) SetDesc(v []string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *Monster) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetCharisma

`func (o *Monster) GetCharisma() float32`

GetCharisma returns the Charisma field if non-nil, zero value otherwise.

### GetCharismaOk

`func (o *Monster) GetCharismaOk() (*float32, bool)`

GetCharismaOk returns a tuple with the Charisma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharisma

`func (o *Monster) SetCharisma(v float32)`

SetCharisma sets Charisma field to given value.

### HasCharisma

`func (o *Monster) HasCharisma() bool`

HasCharisma returns a boolean if a field has been set.

### GetConstitution

`func (o *Monster) GetConstitution() float32`

GetConstitution returns the Constitution field if non-nil, zero value otherwise.

### GetConstitutionOk

`func (o *Monster) GetConstitutionOk() (*float32, bool)`

GetConstitutionOk returns a tuple with the Constitution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstitution

`func (o *Monster) SetConstitution(v float32)`

SetConstitution sets Constitution field to given value.

### HasConstitution

`func (o *Monster) HasConstitution() bool`

HasConstitution returns a boolean if a field has been set.

### GetDexterity

`func (o *Monster) GetDexterity() float32`

GetDexterity returns the Dexterity field if non-nil, zero value otherwise.

### GetDexterityOk

`func (o *Monster) GetDexterityOk() (*float32, bool)`

GetDexterityOk returns a tuple with the Dexterity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDexterity

`func (o *Monster) SetDexterity(v float32)`

SetDexterity sets Dexterity field to given value.

### HasDexterity

`func (o *Monster) HasDexterity() bool`

HasDexterity returns a boolean if a field has been set.

### GetIntelligence

`func (o *Monster) GetIntelligence() float32`

GetIntelligence returns the Intelligence field if non-nil, zero value otherwise.

### GetIntelligenceOk

`func (o *Monster) GetIntelligenceOk() (*float32, bool)`

GetIntelligenceOk returns a tuple with the Intelligence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntelligence

`func (o *Monster) SetIntelligence(v float32)`

SetIntelligence sets Intelligence field to given value.

### HasIntelligence

`func (o *Monster) HasIntelligence() bool`

HasIntelligence returns a boolean if a field has been set.

### GetStrength

`func (o *Monster) GetStrength() float32`

GetStrength returns the Strength field if non-nil, zero value otherwise.

### GetStrengthOk

`func (o *Monster) GetStrengthOk() (*float32, bool)`

GetStrengthOk returns a tuple with the Strength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrength

`func (o *Monster) SetStrength(v float32)`

SetStrength sets Strength field to given value.

### HasStrength

`func (o *Monster) HasStrength() bool`

HasStrength returns a boolean if a field has been set.

### GetWisdom

`func (o *Monster) GetWisdom() float32`

GetWisdom returns the Wisdom field if non-nil, zero value otherwise.

### GetWisdomOk

`func (o *Monster) GetWisdomOk() (*float32, bool)`

GetWisdomOk returns a tuple with the Wisdom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWisdom

`func (o *Monster) SetWisdom(v float32)`

SetWisdom sets Wisdom field to given value.

### HasWisdom

`func (o *Monster) HasWisdom() bool`

HasWisdom returns a boolean if a field has been set.

### GetImage

`func (o *Monster) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Monster) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Monster) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *Monster) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetSize

`func (o *Monster) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Monster) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Monster) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *Monster) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *Monster) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Monster) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Monster) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Monster) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSubtype

`func (o *Monster) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *Monster) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *Monster) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *Monster) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetAlignments

`func (o *Monster) GetAlignments() string`

GetAlignments returns the Alignments field if non-nil, zero value otherwise.

### GetAlignmentsOk

`func (o *Monster) GetAlignmentsOk() (*string, bool)`

GetAlignmentsOk returns a tuple with the Alignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignments

`func (o *Monster) SetAlignments(v string)`

SetAlignments sets Alignments field to given value.

### HasAlignments

`func (o *Monster) HasAlignments() bool`

HasAlignments returns a boolean if a field has been set.

### GetArmorClass

`func (o *Monster) GetArmorClass() []MonsterArmorClass`

GetArmorClass returns the ArmorClass field if non-nil, zero value otherwise.

### GetArmorClassOk

`func (o *Monster) GetArmorClassOk() (*[]MonsterArmorClass, bool)`

GetArmorClassOk returns a tuple with the ArmorClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArmorClass

`func (o *Monster) SetArmorClass(v []MonsterArmorClass)`

SetArmorClass sets ArmorClass field to given value.

### HasArmorClass

`func (o *Monster) HasArmorClass() bool`

HasArmorClass returns a boolean if a field has been set.

### GetHitPoints

`func (o *Monster) GetHitPoints() float32`

GetHitPoints returns the HitPoints field if non-nil, zero value otherwise.

### GetHitPointsOk

`func (o *Monster) GetHitPointsOk() (*float32, bool)`

GetHitPointsOk returns a tuple with the HitPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitPoints

`func (o *Monster) SetHitPoints(v float32)`

SetHitPoints sets HitPoints field to given value.

### HasHitPoints

`func (o *Monster) HasHitPoints() bool`

HasHitPoints returns a boolean if a field has been set.

### GetHitDice

`func (o *Monster) GetHitDice() string`

GetHitDice returns the HitDice field if non-nil, zero value otherwise.

### GetHitDiceOk

`func (o *Monster) GetHitDiceOk() (*string, bool)`

GetHitDiceOk returns a tuple with the HitDice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitDice

`func (o *Monster) SetHitDice(v string)`

SetHitDice sets HitDice field to given value.

### HasHitDice

`func (o *Monster) HasHitDice() bool`

HasHitDice returns a boolean if a field has been set.

### GetHitPointsRoll

`func (o *Monster) GetHitPointsRoll() string`

GetHitPointsRoll returns the HitPointsRoll field if non-nil, zero value otherwise.

### GetHitPointsRollOk

`func (o *Monster) GetHitPointsRollOk() (*string, bool)`

GetHitPointsRollOk returns a tuple with the HitPointsRoll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitPointsRoll

`func (o *Monster) SetHitPointsRoll(v string)`

SetHitPointsRoll sets HitPointsRoll field to given value.

### HasHitPointsRoll

`func (o *Monster) HasHitPointsRoll() bool`

HasHitPointsRoll returns a boolean if a field has been set.

### GetActions

`func (o *Monster) GetActions() []MonsterAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *Monster) GetActionsOk() (*[]MonsterAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *Monster) SetActions(v []MonsterAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *Monster) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetLegendaryActions

`func (o *Monster) GetLegendaryActions() []MonsterAction`

GetLegendaryActions returns the LegendaryActions field if non-nil, zero value otherwise.

### GetLegendaryActionsOk

`func (o *Monster) GetLegendaryActionsOk() (*[]MonsterAction, bool)`

GetLegendaryActionsOk returns a tuple with the LegendaryActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegendaryActions

`func (o *Monster) SetLegendaryActions(v []MonsterAction)`

SetLegendaryActions sets LegendaryActions field to given value.

### HasLegendaryActions

`func (o *Monster) HasLegendaryActions() bool`

HasLegendaryActions returns a boolean if a field has been set.

### GetChallengeRating

`func (o *Monster) GetChallengeRating() float32`

GetChallengeRating returns the ChallengeRating field if non-nil, zero value otherwise.

### GetChallengeRatingOk

`func (o *Monster) GetChallengeRatingOk() (*float32, bool)`

GetChallengeRatingOk returns a tuple with the ChallengeRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeRating

`func (o *Monster) SetChallengeRating(v float32)`

SetChallengeRating sets ChallengeRating field to given value.

### HasChallengeRating

`func (o *Monster) HasChallengeRating() bool`

HasChallengeRating returns a boolean if a field has been set.

### GetConditionImmunities

`func (o *Monster) GetConditionImmunities() []APIReference`

GetConditionImmunities returns the ConditionImmunities field if non-nil, zero value otherwise.

### GetConditionImmunitiesOk

`func (o *Monster) GetConditionImmunitiesOk() (*[]APIReference, bool)`

GetConditionImmunitiesOk returns a tuple with the ConditionImmunities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionImmunities

`func (o *Monster) SetConditionImmunities(v []APIReference)`

SetConditionImmunities sets ConditionImmunities field to given value.

### HasConditionImmunities

`func (o *Monster) HasConditionImmunities() bool`

HasConditionImmunities returns a boolean if a field has been set.

### GetDamageImmunities

`func (o *Monster) GetDamageImmunities() []string`

GetDamageImmunities returns the DamageImmunities field if non-nil, zero value otherwise.

### GetDamageImmunitiesOk

`func (o *Monster) GetDamageImmunitiesOk() (*[]string, bool)`

GetDamageImmunitiesOk returns a tuple with the DamageImmunities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamageImmunities

`func (o *Monster) SetDamageImmunities(v []string)`

SetDamageImmunities sets DamageImmunities field to given value.

### HasDamageImmunities

`func (o *Monster) HasDamageImmunities() bool`

HasDamageImmunities returns a boolean if a field has been set.

### GetDamageResistances

`func (o *Monster) GetDamageResistances() []string`

GetDamageResistances returns the DamageResistances field if non-nil, zero value otherwise.

### GetDamageResistancesOk

`func (o *Monster) GetDamageResistancesOk() (*[]string, bool)`

GetDamageResistancesOk returns a tuple with the DamageResistances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamageResistances

`func (o *Monster) SetDamageResistances(v []string)`

SetDamageResistances sets DamageResistances field to given value.

### HasDamageResistances

`func (o *Monster) HasDamageResistances() bool`

HasDamageResistances returns a boolean if a field has been set.

### GetDamageVulnerabilities

`func (o *Monster) GetDamageVulnerabilities() []string`

GetDamageVulnerabilities returns the DamageVulnerabilities field if non-nil, zero value otherwise.

### GetDamageVulnerabilitiesOk

`func (o *Monster) GetDamageVulnerabilitiesOk() (*[]string, bool)`

GetDamageVulnerabilitiesOk returns a tuple with the DamageVulnerabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamageVulnerabilities

`func (o *Monster) SetDamageVulnerabilities(v []string)`

SetDamageVulnerabilities sets DamageVulnerabilities field to given value.

### HasDamageVulnerabilities

`func (o *Monster) HasDamageVulnerabilities() bool`

HasDamageVulnerabilities returns a boolean if a field has been set.

### GetForms

`func (o *Monster) GetForms() []APIReference`

GetForms returns the Forms field if non-nil, zero value otherwise.

### GetFormsOk

`func (o *Monster) GetFormsOk() (*[]APIReference, bool)`

GetFormsOk returns a tuple with the Forms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForms

`func (o *Monster) SetForms(v []APIReference)`

SetForms sets Forms field to given value.

### HasForms

`func (o *Monster) HasForms() bool`

HasForms returns a boolean if a field has been set.

### GetLanguages

`func (o *Monster) GetLanguages() string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *Monster) GetLanguagesOk() (*string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *Monster) SetLanguages(v string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *Monster) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetProficiencies

`func (o *Monster) GetProficiencies() []MonsterProficiency`

GetProficiencies returns the Proficiencies field if non-nil, zero value otherwise.

### GetProficienciesOk

`func (o *Monster) GetProficienciesOk() (*[]MonsterProficiency, bool)`

GetProficienciesOk returns a tuple with the Proficiencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProficiencies

`func (o *Monster) SetProficiencies(v []MonsterProficiency)`

SetProficiencies sets Proficiencies field to given value.

### HasProficiencies

`func (o *Monster) HasProficiencies() bool`

HasProficiencies returns a boolean if a field has been set.

### GetReactions

`func (o *Monster) GetReactions() []MonsterAction`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *Monster) GetReactionsOk() (*[]MonsterAction, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *Monster) SetReactions(v []MonsterAction)`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *Monster) HasReactions() bool`

HasReactions returns a boolean if a field has been set.

### GetSenses

`func (o *Monster) GetSenses() MonsterAllOfSenses`

GetSenses returns the Senses field if non-nil, zero value otherwise.

### GetSensesOk

`func (o *Monster) GetSensesOk() (*MonsterAllOfSenses, bool)`

GetSensesOk returns a tuple with the Senses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenses

`func (o *Monster) SetSenses(v MonsterAllOfSenses)`

SetSenses sets Senses field to given value.

### HasSenses

`func (o *Monster) HasSenses() bool`

HasSenses returns a boolean if a field has been set.

### GetSpecialAbilities

`func (o *Monster) GetSpecialAbilities() []MonsterSpecialAbility`

GetSpecialAbilities returns the SpecialAbilities field if non-nil, zero value otherwise.

### GetSpecialAbilitiesOk

`func (o *Monster) GetSpecialAbilitiesOk() (*[]MonsterSpecialAbility, bool)`

GetSpecialAbilitiesOk returns a tuple with the SpecialAbilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialAbilities

`func (o *Monster) SetSpecialAbilities(v []MonsterSpecialAbility)`

SetSpecialAbilities sets SpecialAbilities field to given value.

### HasSpecialAbilities

`func (o *Monster) HasSpecialAbilities() bool`

HasSpecialAbilities returns a boolean if a field has been set.

### GetSpeed

`func (o *Monster) GetSpeed() MonsterAllOfSpeed`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *Monster) GetSpeedOk() (*MonsterAllOfSpeed, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *Monster) SetSpeed(v MonsterAllOfSpeed)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *Monster) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetXp

`func (o *Monster) GetXp() float32`

GetXp returns the Xp field if non-nil, zero value otherwise.

### GetXpOk

`func (o *Monster) GetXpOk() (*float32, bool)`

GetXpOk returns a tuple with the Xp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXp

`func (o *Monster) SetXp(v float32)`

SetXp sets Xp field to given value.

### HasXp

`func (o *Monster) HasXp() bool`

HasXp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


