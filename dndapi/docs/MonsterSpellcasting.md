# MonsterSpellcasting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ability** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 
**Dc** | Pointer to **float32** |  | [optional] 
**Modifier** | Pointer to **float32** |  | [optional] 
**ComponentsRequired** | Pointer to **[]string** |  | [optional] 
**School** | Pointer to **string** |  | [optional] 
**Slots** | Pointer to **map[string]float32** |  | [optional] 
**Spells** | Pointer to [**[]MonsterSpell**](MonsterSpell.md) |  | [optional] 

## Methods

### NewMonsterSpellcasting

`func NewMonsterSpellcasting() *MonsterSpellcasting`

NewMonsterSpellcasting instantiates a new MonsterSpellcasting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonsterSpellcastingWithDefaults

`func NewMonsterSpellcastingWithDefaults() *MonsterSpellcasting`

NewMonsterSpellcastingWithDefaults instantiates a new MonsterSpellcasting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbility

`func (o *MonsterSpellcasting) GetAbility() APIReference`

GetAbility returns the Ability field if non-nil, zero value otherwise.

### GetAbilityOk

`func (o *MonsterSpellcasting) GetAbilityOk() (*APIReference, bool)`

GetAbilityOk returns a tuple with the Ability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbility

`func (o *MonsterSpellcasting) SetAbility(v APIReference)`

SetAbility sets Ability field to given value.

### HasAbility

`func (o *MonsterSpellcasting) HasAbility() bool`

HasAbility returns a boolean if a field has been set.

### GetDc

`func (o *MonsterSpellcasting) GetDc() float32`

GetDc returns the Dc field if non-nil, zero value otherwise.

### GetDcOk

`func (o *MonsterSpellcasting) GetDcOk() (*float32, bool)`

GetDcOk returns a tuple with the Dc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDc

`func (o *MonsterSpellcasting) SetDc(v float32)`

SetDc sets Dc field to given value.

### HasDc

`func (o *MonsterSpellcasting) HasDc() bool`

HasDc returns a boolean if a field has been set.

### GetModifier

`func (o *MonsterSpellcasting) GetModifier() float32`

GetModifier returns the Modifier field if non-nil, zero value otherwise.

### GetModifierOk

`func (o *MonsterSpellcasting) GetModifierOk() (*float32, bool)`

GetModifierOk returns a tuple with the Modifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifier

`func (o *MonsterSpellcasting) SetModifier(v float32)`

SetModifier sets Modifier field to given value.

### HasModifier

`func (o *MonsterSpellcasting) HasModifier() bool`

HasModifier returns a boolean if a field has been set.

### GetComponentsRequired

`func (o *MonsterSpellcasting) GetComponentsRequired() []string`

GetComponentsRequired returns the ComponentsRequired field if non-nil, zero value otherwise.

### GetComponentsRequiredOk

`func (o *MonsterSpellcasting) GetComponentsRequiredOk() (*[]string, bool)`

GetComponentsRequiredOk returns a tuple with the ComponentsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentsRequired

`func (o *MonsterSpellcasting) SetComponentsRequired(v []string)`

SetComponentsRequired sets ComponentsRequired field to given value.

### HasComponentsRequired

`func (o *MonsterSpellcasting) HasComponentsRequired() bool`

HasComponentsRequired returns a boolean if a field has been set.

### GetSchool

`func (o *MonsterSpellcasting) GetSchool() string`

GetSchool returns the School field if non-nil, zero value otherwise.

### GetSchoolOk

`func (o *MonsterSpellcasting) GetSchoolOk() (*string, bool)`

GetSchoolOk returns a tuple with the School field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchool

`func (o *MonsterSpellcasting) SetSchool(v string)`

SetSchool sets School field to given value.

### HasSchool

`func (o *MonsterSpellcasting) HasSchool() bool`

HasSchool returns a boolean if a field has been set.

### GetSlots

`func (o *MonsterSpellcasting) GetSlots() map[string]float32`

GetSlots returns the Slots field if non-nil, zero value otherwise.

### GetSlotsOk

`func (o *MonsterSpellcasting) GetSlotsOk() (*map[string]float32, bool)`

GetSlotsOk returns a tuple with the Slots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlots

`func (o *MonsterSpellcasting) SetSlots(v map[string]float32)`

SetSlots sets Slots field to given value.

### HasSlots

`func (o *MonsterSpellcasting) HasSlots() bool`

HasSlots returns a boolean if a field has been set.

### GetSpells

`func (o *MonsterSpellcasting) GetSpells() []MonsterSpell`

GetSpells returns the Spells field if non-nil, zero value otherwise.

### GetSpellsOk

`func (o *MonsterSpellcasting) GetSpellsOk() (*[]MonsterSpell, bool)`

GetSpellsOk returns a tuple with the Spells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpells

`func (o *MonsterSpellcasting) SetSpells(v []MonsterSpell)`

SetSpells sets Spells field to given value.

### HasSpells

`func (o *MonsterSpellcasting) HasSpells() bool`

HasSpells returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


