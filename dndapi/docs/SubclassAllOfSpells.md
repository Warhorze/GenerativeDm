# SubclassAllOfSpells

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prerequisites** | Pointer to [**[]SpellPrerequisite**](SpellPrerequisite.md) |  | [optional] 
**Spell** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 

## Methods

### NewSubclassAllOfSpells

`func NewSubclassAllOfSpells() *SubclassAllOfSpells`

NewSubclassAllOfSpells instantiates a new SubclassAllOfSpells object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubclassAllOfSpellsWithDefaults

`func NewSubclassAllOfSpellsWithDefaults() *SubclassAllOfSpells`

NewSubclassAllOfSpellsWithDefaults instantiates a new SubclassAllOfSpells object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrerequisites

`func (o *SubclassAllOfSpells) GetPrerequisites() []SpellPrerequisite`

GetPrerequisites returns the Prerequisites field if non-nil, zero value otherwise.

### GetPrerequisitesOk

`func (o *SubclassAllOfSpells) GetPrerequisitesOk() (*[]SpellPrerequisite, bool)`

GetPrerequisitesOk returns a tuple with the Prerequisites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrerequisites

`func (o *SubclassAllOfSpells) SetPrerequisites(v []SpellPrerequisite)`

SetPrerequisites sets Prerequisites field to given value.

### HasPrerequisites

`func (o *SubclassAllOfSpells) HasPrerequisites() bool`

HasPrerequisites returns a boolean if a field has been set.

### GetSpell

`func (o *SubclassAllOfSpells) GetSpell() APIReference`

GetSpell returns the Spell field if non-nil, zero value otherwise.

### GetSpellOk

`func (o *SubclassAllOfSpells) GetSpellOk() (*APIReference, bool)`

GetSpellOk returns a tuple with the Spell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpell

`func (o *SubclassAllOfSpells) SetSpell(v APIReference)`

SetSpell sets Spell field to given value.

### HasSpell

`func (o *SubclassAllOfSpells) HasSpell() bool`

HasSpell returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


