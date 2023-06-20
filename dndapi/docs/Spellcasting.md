# Spellcasting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | Pointer to **float32** | Level at which the class can start using its spellcasting abilities. | [optional] 
**Info** | Pointer to [**[]SpellcastingInfoInner**](SpellcastingInfoInner.md) | Descriptions of the class&#39; ability to cast spells. | [optional] 
**SpellcastingAbility** | Pointer to [**SpellcastingSpellcastingAbility**](SpellcastingSpellcastingAbility.md) |  | [optional] 

## Methods

### NewSpellcasting

`func NewSpellcasting() *Spellcasting`

NewSpellcasting instantiates a new Spellcasting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpellcastingWithDefaults

`func NewSpellcastingWithDefaults() *Spellcasting`

NewSpellcastingWithDefaults instantiates a new Spellcasting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *Spellcasting) GetLevel() float32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *Spellcasting) GetLevelOk() (*float32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *Spellcasting) SetLevel(v float32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *Spellcasting) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetInfo

`func (o *Spellcasting) GetInfo() []SpellcastingInfoInner`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Spellcasting) GetInfoOk() (*[]SpellcastingInfoInner, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Spellcasting) SetInfo(v []SpellcastingInfoInner)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Spellcasting) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetSpellcastingAbility

`func (o *Spellcasting) GetSpellcastingAbility() SpellcastingSpellcastingAbility`

GetSpellcastingAbility returns the SpellcastingAbility field if non-nil, zero value otherwise.

### GetSpellcastingAbilityOk

`func (o *Spellcasting) GetSpellcastingAbilityOk() (*SpellcastingSpellcastingAbility, bool)`

GetSpellcastingAbilityOk returns a tuple with the SpellcastingAbility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpellcastingAbility

`func (o *Spellcasting) SetSpellcastingAbility(v SpellcastingSpellcastingAbility)`

SetSpellcastingAbility sets SpellcastingAbility field to given value.

### HasSpellcastingAbility

`func (o *Spellcasting) HasSpellcastingAbility() bool`

HasSpellcastingAbility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


