# MonsterAttack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Dc** | Pointer to [**DC**](DC.md) |  | [optional] 
**Damage** | Pointer to [**Damage**](Damage.md) |  | [optional] 

## Methods

### NewMonsterAttack

`func NewMonsterAttack() *MonsterAttack`

NewMonsterAttack instantiates a new MonsterAttack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonsterAttackWithDefaults

`func NewMonsterAttackWithDefaults() *MonsterAttack`

NewMonsterAttackWithDefaults instantiates a new MonsterAttack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MonsterAttack) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MonsterAttack) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MonsterAttack) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MonsterAttack) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDc

`func (o *MonsterAttack) GetDc() DC`

GetDc returns the Dc field if non-nil, zero value otherwise.

### GetDcOk

`func (o *MonsterAttack) GetDcOk() (*DC, bool)`

GetDcOk returns a tuple with the Dc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDc

`func (o *MonsterAttack) SetDc(v DC)`

SetDc sets Dc field to given value.

### HasDc

`func (o *MonsterAttack) HasDc() bool`

HasDc returns a boolean if a field has been set.

### GetDamage

`func (o *MonsterAttack) GetDamage() Damage`

GetDamage returns the Damage field if non-nil, zero value otherwise.

### GetDamageOk

`func (o *MonsterAttack) GetDamageOk() (*Damage, bool)`

GetDamageOk returns a tuple with the Damage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamage

`func (o *MonsterAttack) SetDamage(v Damage)`

SetDamage sets Damage field to given value.

### HasDamage

`func (o *MonsterAttack) HasDamage() bool`

HasDamage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


