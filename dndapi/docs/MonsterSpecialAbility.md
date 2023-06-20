# MonsterSpecialAbility

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Desc** | Pointer to **string** |  | [optional] 
**AttackBonus** | Pointer to **float32** |  | [optional] 
**Damage** | Pointer to [**Damage**](Damage.md) |  | [optional] 
**Dc** | Pointer to [**DC**](DC.md) |  | [optional] 
**Spellcasting** | Pointer to [**MonsterSpellcasting**](MonsterSpellcasting.md) |  | [optional] 
**Usage** | Pointer to [**MonsterUsage**](MonsterUsage.md) |  | [optional] 

## Methods

### NewMonsterSpecialAbility

`func NewMonsterSpecialAbility() *MonsterSpecialAbility`

NewMonsterSpecialAbility instantiates a new MonsterSpecialAbility object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonsterSpecialAbilityWithDefaults

`func NewMonsterSpecialAbilityWithDefaults() *MonsterSpecialAbility`

NewMonsterSpecialAbilityWithDefaults instantiates a new MonsterSpecialAbility object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MonsterSpecialAbility) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MonsterSpecialAbility) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MonsterSpecialAbility) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MonsterSpecialAbility) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDesc

`func (o *MonsterSpecialAbility) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *MonsterSpecialAbility) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *MonsterSpecialAbility) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *MonsterSpecialAbility) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetAttackBonus

`func (o *MonsterSpecialAbility) GetAttackBonus() float32`

GetAttackBonus returns the AttackBonus field if non-nil, zero value otherwise.

### GetAttackBonusOk

`func (o *MonsterSpecialAbility) GetAttackBonusOk() (*float32, bool)`

GetAttackBonusOk returns a tuple with the AttackBonus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackBonus

`func (o *MonsterSpecialAbility) SetAttackBonus(v float32)`

SetAttackBonus sets AttackBonus field to given value.

### HasAttackBonus

`func (o *MonsterSpecialAbility) HasAttackBonus() bool`

HasAttackBonus returns a boolean if a field has been set.

### GetDamage

`func (o *MonsterSpecialAbility) GetDamage() Damage`

GetDamage returns the Damage field if non-nil, zero value otherwise.

### GetDamageOk

`func (o *MonsterSpecialAbility) GetDamageOk() (*Damage, bool)`

GetDamageOk returns a tuple with the Damage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamage

`func (o *MonsterSpecialAbility) SetDamage(v Damage)`

SetDamage sets Damage field to given value.

### HasDamage

`func (o *MonsterSpecialAbility) HasDamage() bool`

HasDamage returns a boolean if a field has been set.

### GetDc

`func (o *MonsterSpecialAbility) GetDc() DC`

GetDc returns the Dc field if non-nil, zero value otherwise.

### GetDcOk

`func (o *MonsterSpecialAbility) GetDcOk() (*DC, bool)`

GetDcOk returns a tuple with the Dc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDc

`func (o *MonsterSpecialAbility) SetDc(v DC)`

SetDc sets Dc field to given value.

### HasDc

`func (o *MonsterSpecialAbility) HasDc() bool`

HasDc returns a boolean if a field has been set.

### GetSpellcasting

`func (o *MonsterSpecialAbility) GetSpellcasting() MonsterSpellcasting`

GetSpellcasting returns the Spellcasting field if non-nil, zero value otherwise.

### GetSpellcastingOk

`func (o *MonsterSpecialAbility) GetSpellcastingOk() (*MonsterSpellcasting, bool)`

GetSpellcastingOk returns a tuple with the Spellcasting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpellcasting

`func (o *MonsterSpecialAbility) SetSpellcasting(v MonsterSpellcasting)`

SetSpellcasting sets Spellcasting field to given value.

### HasSpellcasting

`func (o *MonsterSpecialAbility) HasSpellcasting() bool`

HasSpellcasting returns a boolean if a field has been set.

### GetUsage

`func (o *MonsterSpecialAbility) GetUsage() MonsterUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *MonsterSpecialAbility) GetUsageOk() (*MonsterUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *MonsterSpecialAbility) SetUsage(v MonsterUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *MonsterSpecialAbility) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


