# MonsterAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Desc** | Pointer to **string** |  | [optional] 
**ActionOptions** | Pointer to [**Choice**](Choice.md) |  | [optional] 
**Actions** | Pointer to [**[]MonsterMultiAttackAction**](MonsterMultiAttackAction.md) |  | [optional] 
**Options** | Pointer to [**Choice**](Choice.md) |  | [optional] 
**MultiattackType** | Pointer to **string** |  | [optional] 
**AttackBonus** | Pointer to **float32** |  | [optional] 
**Dc** | Pointer to [**DC**](DC.md) |  | [optional] 
**Attacks** | Pointer to [**[]MonsterAttack**](MonsterAttack.md) |  | [optional] 
**Damage** | Pointer to [**[]Damage**](Damage.md) |  | [optional] 

## Methods

### NewMonsterAction

`func NewMonsterAction() *MonsterAction`

NewMonsterAction instantiates a new MonsterAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonsterActionWithDefaults

`func NewMonsterActionWithDefaults() *MonsterAction`

NewMonsterActionWithDefaults instantiates a new MonsterAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MonsterAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MonsterAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MonsterAction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MonsterAction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDesc

`func (o *MonsterAction) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *MonsterAction) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *MonsterAction) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *MonsterAction) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetActionOptions

`func (o *MonsterAction) GetActionOptions() Choice`

GetActionOptions returns the ActionOptions field if non-nil, zero value otherwise.

### GetActionOptionsOk

`func (o *MonsterAction) GetActionOptionsOk() (*Choice, bool)`

GetActionOptionsOk returns a tuple with the ActionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionOptions

`func (o *MonsterAction) SetActionOptions(v Choice)`

SetActionOptions sets ActionOptions field to given value.

### HasActionOptions

`func (o *MonsterAction) HasActionOptions() bool`

HasActionOptions returns a boolean if a field has been set.

### GetActions

`func (o *MonsterAction) GetActions() []MonsterMultiAttackAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *MonsterAction) GetActionsOk() (*[]MonsterMultiAttackAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *MonsterAction) SetActions(v []MonsterMultiAttackAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *MonsterAction) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetOptions

`func (o *MonsterAction) GetOptions() Choice`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *MonsterAction) GetOptionsOk() (*Choice, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *MonsterAction) SetOptions(v Choice)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *MonsterAction) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetMultiattackType

`func (o *MonsterAction) GetMultiattackType() string`

GetMultiattackType returns the MultiattackType field if non-nil, zero value otherwise.

### GetMultiattackTypeOk

`func (o *MonsterAction) GetMultiattackTypeOk() (*string, bool)`

GetMultiattackTypeOk returns a tuple with the MultiattackType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiattackType

`func (o *MonsterAction) SetMultiattackType(v string)`

SetMultiattackType sets MultiattackType field to given value.

### HasMultiattackType

`func (o *MonsterAction) HasMultiattackType() bool`

HasMultiattackType returns a boolean if a field has been set.

### GetAttackBonus

`func (o *MonsterAction) GetAttackBonus() float32`

GetAttackBonus returns the AttackBonus field if non-nil, zero value otherwise.

### GetAttackBonusOk

`func (o *MonsterAction) GetAttackBonusOk() (*float32, bool)`

GetAttackBonusOk returns a tuple with the AttackBonus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackBonus

`func (o *MonsterAction) SetAttackBonus(v float32)`

SetAttackBonus sets AttackBonus field to given value.

### HasAttackBonus

`func (o *MonsterAction) HasAttackBonus() bool`

HasAttackBonus returns a boolean if a field has been set.

### GetDc

`func (o *MonsterAction) GetDc() DC`

GetDc returns the Dc field if non-nil, zero value otherwise.

### GetDcOk

`func (o *MonsterAction) GetDcOk() (*DC, bool)`

GetDcOk returns a tuple with the Dc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDc

`func (o *MonsterAction) SetDc(v DC)`

SetDc sets Dc field to given value.

### HasDc

`func (o *MonsterAction) HasDc() bool`

HasDc returns a boolean if a field has been set.

### GetAttacks

`func (o *MonsterAction) GetAttacks() []MonsterAttack`

GetAttacks returns the Attacks field if non-nil, zero value otherwise.

### GetAttacksOk

`func (o *MonsterAction) GetAttacksOk() (*[]MonsterAttack, bool)`

GetAttacksOk returns a tuple with the Attacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttacks

`func (o *MonsterAction) SetAttacks(v []MonsterAttack)`

SetAttacks sets Attacks field to given value.

### HasAttacks

`func (o *MonsterAction) HasAttacks() bool`

HasAttacks returns a boolean if a field has been set.

### GetDamage

`func (o *MonsterAction) GetDamage() []Damage`

GetDamage returns the Damage field if non-nil, zero value otherwise.

### GetDamageOk

`func (o *MonsterAction) GetDamageOk() (*[]Damage, bool)`

GetDamageOk returns a tuple with the Damage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamage

`func (o *MonsterAction) SetDamage(v []Damage)`

SetDamage sets Damage field to given value.

### HasDamage

`func (o *MonsterAction) HasDamage() bool`

HasDamage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


