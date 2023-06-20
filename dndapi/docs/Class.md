# Class

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **string** | Resource index for shorthand searching. | [optional] 
**Name** | Pointer to **string** | Name of the referenced resource. | [optional] 
**Url** | Pointer to **string** | URL of the referenced resource. | [optional] 
**HitDie** | Pointer to **float32** | Hit die of the class. (ex: 12 &#x3D;&#x3D; 1d12). | [optional] 
**ClassLevels** | Pointer to **string** | URL of the level resource for the class. | [optional] 
**MultiClassing** | Pointer to [**Multiclassing**](Multiclassing.md) |  | [optional] 
**Spellcasting** | Pointer to [**Spellcasting**](Spellcasting.md) |  | [optional] 
**Spells** | Pointer to **string** | URL of the spell resource list for the class. | [optional] 
**StartingEquipment** | Pointer to [**[]ClassAllOfStartingEquipment**](ClassAllOfStartingEquipment.md) | List of equipment and their quantities all players of the class start with. | [optional] 
**StartingEquipmentOptions** | Pointer to [**[]Choice**](Choice.md) | List of choices of starting equipment. | [optional] 
**ProficiencyChoices** | Pointer to [**[]Choice**](Choice.md) | List of choices of starting proficiencies. | [optional] 
**Proficiencies** | Pointer to [**[]APIReference**](APIReference.md) | List of starting proficiencies for all new characters of this class. | [optional] 
**SavingThrows** | Pointer to [**[]APIReference**](APIReference.md) | Saving throws the class is proficient in. | [optional] 
**Subclasses** | Pointer to [**[]APIReference**](APIReference.md) | List of all possible subclasses this class can specialize in. | [optional] 

## Methods

### NewClass

`func NewClass() *Class`

NewClass instantiates a new Class object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassWithDefaults

`func NewClassWithDefaults() *Class`

NewClassWithDefaults instantiates a new Class object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *Class) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *Class) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *Class) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *Class) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetName

`func (o *Class) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Class) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Class) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Class) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *Class) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Class) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Class) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Class) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetHitDie

`func (o *Class) GetHitDie() float32`

GetHitDie returns the HitDie field if non-nil, zero value otherwise.

### GetHitDieOk

`func (o *Class) GetHitDieOk() (*float32, bool)`

GetHitDieOk returns a tuple with the HitDie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitDie

`func (o *Class) SetHitDie(v float32)`

SetHitDie sets HitDie field to given value.

### HasHitDie

`func (o *Class) HasHitDie() bool`

HasHitDie returns a boolean if a field has been set.

### GetClassLevels

`func (o *Class) GetClassLevels() string`

GetClassLevels returns the ClassLevels field if non-nil, zero value otherwise.

### GetClassLevelsOk

`func (o *Class) GetClassLevelsOk() (*string, bool)`

GetClassLevelsOk returns a tuple with the ClassLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassLevels

`func (o *Class) SetClassLevels(v string)`

SetClassLevels sets ClassLevels field to given value.

### HasClassLevels

`func (o *Class) HasClassLevels() bool`

HasClassLevels returns a boolean if a field has been set.

### GetMultiClassing

`func (o *Class) GetMultiClassing() Multiclassing`

GetMultiClassing returns the MultiClassing field if non-nil, zero value otherwise.

### GetMultiClassingOk

`func (o *Class) GetMultiClassingOk() (*Multiclassing, bool)`

GetMultiClassingOk returns a tuple with the MultiClassing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiClassing

`func (o *Class) SetMultiClassing(v Multiclassing)`

SetMultiClassing sets MultiClassing field to given value.

### HasMultiClassing

`func (o *Class) HasMultiClassing() bool`

HasMultiClassing returns a boolean if a field has been set.

### GetSpellcasting

`func (o *Class) GetSpellcasting() Spellcasting`

GetSpellcasting returns the Spellcasting field if non-nil, zero value otherwise.

### GetSpellcastingOk

`func (o *Class) GetSpellcastingOk() (*Spellcasting, bool)`

GetSpellcastingOk returns a tuple with the Spellcasting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpellcasting

`func (o *Class) SetSpellcasting(v Spellcasting)`

SetSpellcasting sets Spellcasting field to given value.

### HasSpellcasting

`func (o *Class) HasSpellcasting() bool`

HasSpellcasting returns a boolean if a field has been set.

### GetSpells

`func (o *Class) GetSpells() string`

GetSpells returns the Spells field if non-nil, zero value otherwise.

### GetSpellsOk

`func (o *Class) GetSpellsOk() (*string, bool)`

GetSpellsOk returns a tuple with the Spells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpells

`func (o *Class) SetSpells(v string)`

SetSpells sets Spells field to given value.

### HasSpells

`func (o *Class) HasSpells() bool`

HasSpells returns a boolean if a field has been set.

### GetStartingEquipment

`func (o *Class) GetStartingEquipment() []ClassAllOfStartingEquipment`

GetStartingEquipment returns the StartingEquipment field if non-nil, zero value otherwise.

### GetStartingEquipmentOk

`func (o *Class) GetStartingEquipmentOk() (*[]ClassAllOfStartingEquipment, bool)`

GetStartingEquipmentOk returns a tuple with the StartingEquipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingEquipment

`func (o *Class) SetStartingEquipment(v []ClassAllOfStartingEquipment)`

SetStartingEquipment sets StartingEquipment field to given value.

### HasStartingEquipment

`func (o *Class) HasStartingEquipment() bool`

HasStartingEquipment returns a boolean if a field has been set.

### GetStartingEquipmentOptions

`func (o *Class) GetStartingEquipmentOptions() []Choice`

GetStartingEquipmentOptions returns the StartingEquipmentOptions field if non-nil, zero value otherwise.

### GetStartingEquipmentOptionsOk

`func (o *Class) GetStartingEquipmentOptionsOk() (*[]Choice, bool)`

GetStartingEquipmentOptionsOk returns a tuple with the StartingEquipmentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingEquipmentOptions

`func (o *Class) SetStartingEquipmentOptions(v []Choice)`

SetStartingEquipmentOptions sets StartingEquipmentOptions field to given value.

### HasStartingEquipmentOptions

`func (o *Class) HasStartingEquipmentOptions() bool`

HasStartingEquipmentOptions returns a boolean if a field has been set.

### GetProficiencyChoices

`func (o *Class) GetProficiencyChoices() []Choice`

GetProficiencyChoices returns the ProficiencyChoices field if non-nil, zero value otherwise.

### GetProficiencyChoicesOk

`func (o *Class) GetProficiencyChoicesOk() (*[]Choice, bool)`

GetProficiencyChoicesOk returns a tuple with the ProficiencyChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProficiencyChoices

`func (o *Class) SetProficiencyChoices(v []Choice)`

SetProficiencyChoices sets ProficiencyChoices field to given value.

### HasProficiencyChoices

`func (o *Class) HasProficiencyChoices() bool`

HasProficiencyChoices returns a boolean if a field has been set.

### GetProficiencies

`func (o *Class) GetProficiencies() []APIReference`

GetProficiencies returns the Proficiencies field if non-nil, zero value otherwise.

### GetProficienciesOk

`func (o *Class) GetProficienciesOk() (*[]APIReference, bool)`

GetProficienciesOk returns a tuple with the Proficiencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProficiencies

`func (o *Class) SetProficiencies(v []APIReference)`

SetProficiencies sets Proficiencies field to given value.

### HasProficiencies

`func (o *Class) HasProficiencies() bool`

HasProficiencies returns a boolean if a field has been set.

### GetSavingThrows

`func (o *Class) GetSavingThrows() []APIReference`

GetSavingThrows returns the SavingThrows field if non-nil, zero value otherwise.

### GetSavingThrowsOk

`func (o *Class) GetSavingThrowsOk() (*[]APIReference, bool)`

GetSavingThrowsOk returns a tuple with the SavingThrows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavingThrows

`func (o *Class) SetSavingThrows(v []APIReference)`

SetSavingThrows sets SavingThrows field to given value.

### HasSavingThrows

`func (o *Class) HasSavingThrows() bool`

HasSavingThrows returns a boolean if a field has been set.

### GetSubclasses

`func (o *Class) GetSubclasses() []APIReference`

GetSubclasses returns the Subclasses field if non-nil, zero value otherwise.

### GetSubclassesOk

`func (o *Class) GetSubclassesOk() (*[]APIReference, bool)`

GetSubclassesOk returns a tuple with the Subclasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubclasses

`func (o *Class) SetSubclasses(v []APIReference)`

SetSubclasses sets Subclasses field to given value.

### HasSubclasses

`func (o *Class) HasSubclasses() bool`

HasSubclasses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


