# MonsterArmorClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **float32** |  | [optional] 
**Desc** | Pointer to **string** |  | [optional] 
**Armor** | Pointer to [**[]APIReference**](APIReference.md) |  | [optional] 
**Spell** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 
**Condition** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 

## Methods

### NewMonsterArmorClass

`func NewMonsterArmorClass() *MonsterArmorClass`

NewMonsterArmorClass instantiates a new MonsterArmorClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonsterArmorClassWithDefaults

`func NewMonsterArmorClassWithDefaults() *MonsterArmorClass`

NewMonsterArmorClassWithDefaults instantiates a new MonsterArmorClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MonsterArmorClass) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MonsterArmorClass) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MonsterArmorClass) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MonsterArmorClass) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *MonsterArmorClass) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MonsterArmorClass) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MonsterArmorClass) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *MonsterArmorClass) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDesc

`func (o *MonsterArmorClass) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *MonsterArmorClass) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *MonsterArmorClass) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *MonsterArmorClass) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetArmor

`func (o *MonsterArmorClass) GetArmor() []APIReference`

GetArmor returns the Armor field if non-nil, zero value otherwise.

### GetArmorOk

`func (o *MonsterArmorClass) GetArmorOk() (*[]APIReference, bool)`

GetArmorOk returns a tuple with the Armor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArmor

`func (o *MonsterArmorClass) SetArmor(v []APIReference)`

SetArmor sets Armor field to given value.

### HasArmor

`func (o *MonsterArmorClass) HasArmor() bool`

HasArmor returns a boolean if a field has been set.

### GetSpell

`func (o *MonsterArmorClass) GetSpell() APIReference`

GetSpell returns the Spell field if non-nil, zero value otherwise.

### GetSpellOk

`func (o *MonsterArmorClass) GetSpellOk() (*APIReference, bool)`

GetSpellOk returns a tuple with the Spell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpell

`func (o *MonsterArmorClass) SetSpell(v APIReference)`

SetSpell sets Spell field to given value.

### HasSpell

`func (o *MonsterArmorClass) HasSpell() bool`

HasSpell returns a boolean if a field has been set.

### GetCondition

`func (o *MonsterArmorClass) GetCondition() APIReference`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *MonsterArmorClass) GetConditionOk() (*APIReference, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *MonsterArmorClass) SetCondition(v APIReference)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *MonsterArmorClass) HasCondition() bool`

HasCondition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


