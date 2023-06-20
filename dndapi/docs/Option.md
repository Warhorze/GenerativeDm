# Option

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OptionType** | Pointer to **string** | Type of option; determines other attributes. | [optional] 
**Item** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 
**ActionName** | Pointer to **string** | The name of the action. | [optional] 
**Count** | Pointer to **float32** | Count | [optional] 
**Type** | Pointer to **string** | For attack options that can be melee, ranged, abilities, or thrown. | [optional] 
**Items** | Pointer to [**[]Option**](Option.md) |  | [optional] 
**Choice** | Pointer to [**Choice**](Choice.md) |  | [optional] 
**String** | Pointer to **string** | The string. | [optional] 
**Desc** | Pointer to **string** | A description of the ideal. | [optional] 
**Alignments** | Pointer to [**[]APIReference**](APIReference.md) | A list of alignments of those who might follow the ideal. | [optional] 
**Of** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 
**AbilityScore** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 
**MinimumScore** | Pointer to **float32** | The minimum score required to satisfy the prerequisite. | [optional] 
**Bonus** | Pointer to **float32** | The bonus being applied to the ability score | [optional] 
**Name** | Pointer to **string** | Name of the breath | [optional] 
**Dc** | Pointer to [**DC**](DC.md) |  | [optional] 
**Damage** | Pointer to [**[]Damage**](Damage.md) | Damage dealt by the breath attack, if any. | [optional] 
**DamageType** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 
**DamageDice** | Pointer to **string** | Damage expressed in dice (e.g. \&quot;13d6\&quot;). | [optional] 
**Notes** | Pointer to **string** | Information regarding the damage. | [optional] 

## Methods

### NewOption

`func NewOption() *Option`

NewOption instantiates a new Option object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionWithDefaults

`func NewOptionWithDefaults() *Option`

NewOptionWithDefaults instantiates a new Option object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptionType

`func (o *Option) GetOptionType() string`

GetOptionType returns the OptionType field if non-nil, zero value otherwise.

### GetOptionTypeOk

`func (o *Option) GetOptionTypeOk() (*string, bool)`

GetOptionTypeOk returns a tuple with the OptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionType

`func (o *Option) SetOptionType(v string)`

SetOptionType sets OptionType field to given value.

### HasOptionType

`func (o *Option) HasOptionType() bool`

HasOptionType returns a boolean if a field has been set.

### GetItem

`func (o *Option) GetItem() APIReference`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *Option) GetItemOk() (*APIReference, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *Option) SetItem(v APIReference)`

SetItem sets Item field to given value.

### HasItem

`func (o *Option) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetActionName

`func (o *Option) GetActionName() string`

GetActionName returns the ActionName field if non-nil, zero value otherwise.

### GetActionNameOk

`func (o *Option) GetActionNameOk() (*string, bool)`

GetActionNameOk returns a tuple with the ActionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionName

`func (o *Option) SetActionName(v string)`

SetActionName sets ActionName field to given value.

### HasActionName

`func (o *Option) HasActionName() bool`

HasActionName returns a boolean if a field has been set.

### GetCount

`func (o *Option) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Option) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Option) SetCount(v float32)`

SetCount sets Count field to given value.

### HasCount

`func (o *Option) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetType

`func (o *Option) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Option) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Option) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Option) HasType() bool`

HasType returns a boolean if a field has been set.

### GetItems

`func (o *Option) GetItems() []Option`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Option) GetItemsOk() (*[]Option, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Option) SetItems(v []Option)`

SetItems sets Items field to given value.

### HasItems

`func (o *Option) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetChoice

`func (o *Option) GetChoice() Choice`

GetChoice returns the Choice field if non-nil, zero value otherwise.

### GetChoiceOk

`func (o *Option) GetChoiceOk() (*Choice, bool)`

GetChoiceOk returns a tuple with the Choice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoice

`func (o *Option) SetChoice(v Choice)`

SetChoice sets Choice field to given value.

### HasChoice

`func (o *Option) HasChoice() bool`

HasChoice returns a boolean if a field has been set.

### GetString

`func (o *Option) GetString() string`

GetString returns the String field if non-nil, zero value otherwise.

### GetStringOk

`func (o *Option) GetStringOk() (*string, bool)`

GetStringOk returns a tuple with the String field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetString

`func (o *Option) SetString(v string)`

SetString sets String field to given value.

### HasString

`func (o *Option) HasString() bool`

HasString returns a boolean if a field has been set.

### GetDesc

`func (o *Option) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *Option) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *Option) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *Option) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetAlignments

`func (o *Option) GetAlignments() []APIReference`

GetAlignments returns the Alignments field if non-nil, zero value otherwise.

### GetAlignmentsOk

`func (o *Option) GetAlignmentsOk() (*[]APIReference, bool)`

GetAlignmentsOk returns a tuple with the Alignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignments

`func (o *Option) SetAlignments(v []APIReference)`

SetAlignments sets Alignments field to given value.

### HasAlignments

`func (o *Option) HasAlignments() bool`

HasAlignments returns a boolean if a field has been set.

### GetOf

`func (o *Option) GetOf() APIReference`

GetOf returns the Of field if non-nil, zero value otherwise.

### GetOfOk

`func (o *Option) GetOfOk() (*APIReference, bool)`

GetOfOk returns a tuple with the Of field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOf

`func (o *Option) SetOf(v APIReference)`

SetOf sets Of field to given value.

### HasOf

`func (o *Option) HasOf() bool`

HasOf returns a boolean if a field has been set.

### GetAbilityScore

`func (o *Option) GetAbilityScore() APIReference`

GetAbilityScore returns the AbilityScore field if non-nil, zero value otherwise.

### GetAbilityScoreOk

`func (o *Option) GetAbilityScoreOk() (*APIReference, bool)`

GetAbilityScoreOk returns a tuple with the AbilityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbilityScore

`func (o *Option) SetAbilityScore(v APIReference)`

SetAbilityScore sets AbilityScore field to given value.

### HasAbilityScore

`func (o *Option) HasAbilityScore() bool`

HasAbilityScore returns a boolean if a field has been set.

### GetMinimumScore

`func (o *Option) GetMinimumScore() float32`

GetMinimumScore returns the MinimumScore field if non-nil, zero value otherwise.

### GetMinimumScoreOk

`func (o *Option) GetMinimumScoreOk() (*float32, bool)`

GetMinimumScoreOk returns a tuple with the MinimumScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumScore

`func (o *Option) SetMinimumScore(v float32)`

SetMinimumScore sets MinimumScore field to given value.

### HasMinimumScore

`func (o *Option) HasMinimumScore() bool`

HasMinimumScore returns a boolean if a field has been set.

### GetBonus

`func (o *Option) GetBonus() float32`

GetBonus returns the Bonus field if non-nil, zero value otherwise.

### GetBonusOk

`func (o *Option) GetBonusOk() (*float32, bool)`

GetBonusOk returns a tuple with the Bonus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonus

`func (o *Option) SetBonus(v float32)`

SetBonus sets Bonus field to given value.

### HasBonus

`func (o *Option) HasBonus() bool`

HasBonus returns a boolean if a field has been set.

### GetName

`func (o *Option) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Option) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Option) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Option) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDc

`func (o *Option) GetDc() DC`

GetDc returns the Dc field if non-nil, zero value otherwise.

### GetDcOk

`func (o *Option) GetDcOk() (*DC, bool)`

GetDcOk returns a tuple with the Dc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDc

`func (o *Option) SetDc(v DC)`

SetDc sets Dc field to given value.

### HasDc

`func (o *Option) HasDc() bool`

HasDc returns a boolean if a field has been set.

### GetDamage

`func (o *Option) GetDamage() []Damage`

GetDamage returns the Damage field if non-nil, zero value otherwise.

### GetDamageOk

`func (o *Option) GetDamageOk() (*[]Damage, bool)`

GetDamageOk returns a tuple with the Damage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamage

`func (o *Option) SetDamage(v []Damage)`

SetDamage sets Damage field to given value.

### HasDamage

`func (o *Option) HasDamage() bool`

HasDamage returns a boolean if a field has been set.

### GetDamageType

`func (o *Option) GetDamageType() APIReference`

GetDamageType returns the DamageType field if non-nil, zero value otherwise.

### GetDamageTypeOk

`func (o *Option) GetDamageTypeOk() (*APIReference, bool)`

GetDamageTypeOk returns a tuple with the DamageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamageType

`func (o *Option) SetDamageType(v APIReference)`

SetDamageType sets DamageType field to given value.

### HasDamageType

`func (o *Option) HasDamageType() bool`

HasDamageType returns a boolean if a field has been set.

### GetDamageDice

`func (o *Option) GetDamageDice() string`

GetDamageDice returns the DamageDice field if non-nil, zero value otherwise.

### GetDamageDiceOk

`func (o *Option) GetDamageDiceOk() (*string, bool)`

GetDamageDiceOk returns a tuple with the DamageDice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamageDice

`func (o *Option) SetDamageDice(v string)`

SetDamageDice sets DamageDice field to given value.

### HasDamageDice

`func (o *Option) HasDamageDice() bool`

HasDamageDice returns a boolean if a field has been set.

### GetNotes

`func (o *Option) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Option) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Option) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Option) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


