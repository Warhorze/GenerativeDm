# Equipment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **string** | Resource index for shorthand searching. | [optional] 
**Name** | Pointer to **string** | Name of the referenced resource. | [optional] 
**Url** | Pointer to **string** | URL of the referenced resource. | [optional] 
**Desc** | Pointer to **[]string** | Description of the resource. | [optional] 
**EquipmentCategory** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 
**WeaponCategory** | Pointer to **string** | The category of weapon this falls into. | [optional] 
**WeaponRange** | Pointer to **string** | Whether this is a Melee or Ranged weapon. | [optional] 
**CategoryRange** | Pointer to **string** | A combination of weapon_category and weapon_range. | [optional] 
**Range** | Pointer to [**WeaponAllOfRange**](WeaponAllOfRange.md) |  | [optional] 
**Damage** | Pointer to [**Damage**](Damage.md) |  | [optional] 
**TwoHandedDamage** | Pointer to [**Damage**](Damage.md) |  | [optional] 
**Properties** | Pointer to [**[]APIReference**](APIReference.md) | A list of the properties this weapon has. | [optional] 
**Cost** | Pointer to [**Cost**](Cost.md) |  | [optional] 
**Weight** | Pointer to **float32** | How much the equipment weighs. | [optional] 
**ArmorCategory** | Pointer to **string** | The category of armor this falls into. | [optional] 
**ArmorClass** | Pointer to **map[string]string** | Details on how to calculate armor class. | [optional] 
**StrMinimum** | Pointer to **float32** | Minimum STR required to use this armor. | [optional] 
**StealthDisadvantage** | Pointer to **bool** | Whether the armor gives disadvantage for Stealth. | [optional] 
**GearCategory** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 
**Contents** | Pointer to [**[]APIReference**](APIReference.md) | The list of adventuring gear in the pack. | [optional] 

## Methods

### NewEquipment

`func NewEquipment() *Equipment`

NewEquipment instantiates a new Equipment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentWithDefaults

`func NewEquipmentWithDefaults() *Equipment`

NewEquipmentWithDefaults instantiates a new Equipment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *Equipment) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *Equipment) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *Equipment) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *Equipment) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetName

`func (o *Equipment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Equipment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Equipment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Equipment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *Equipment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Equipment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Equipment) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Equipment) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDesc

`func (o *Equipment) GetDesc() []string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *Equipment) GetDescOk() (*[]string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *Equipment) SetDesc(v []string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *Equipment) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetEquipmentCategory

`func (o *Equipment) GetEquipmentCategory() APIReference`

GetEquipmentCategory returns the EquipmentCategory field if non-nil, zero value otherwise.

### GetEquipmentCategoryOk

`func (o *Equipment) GetEquipmentCategoryOk() (*APIReference, bool)`

GetEquipmentCategoryOk returns a tuple with the EquipmentCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentCategory

`func (o *Equipment) SetEquipmentCategory(v APIReference)`

SetEquipmentCategory sets EquipmentCategory field to given value.

### HasEquipmentCategory

`func (o *Equipment) HasEquipmentCategory() bool`

HasEquipmentCategory returns a boolean if a field has been set.

### GetWeaponCategory

`func (o *Equipment) GetWeaponCategory() string`

GetWeaponCategory returns the WeaponCategory field if non-nil, zero value otherwise.

### GetWeaponCategoryOk

`func (o *Equipment) GetWeaponCategoryOk() (*string, bool)`

GetWeaponCategoryOk returns a tuple with the WeaponCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeaponCategory

`func (o *Equipment) SetWeaponCategory(v string)`

SetWeaponCategory sets WeaponCategory field to given value.

### HasWeaponCategory

`func (o *Equipment) HasWeaponCategory() bool`

HasWeaponCategory returns a boolean if a field has been set.

### GetWeaponRange

`func (o *Equipment) GetWeaponRange() string`

GetWeaponRange returns the WeaponRange field if non-nil, zero value otherwise.

### GetWeaponRangeOk

`func (o *Equipment) GetWeaponRangeOk() (*string, bool)`

GetWeaponRangeOk returns a tuple with the WeaponRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeaponRange

`func (o *Equipment) SetWeaponRange(v string)`

SetWeaponRange sets WeaponRange field to given value.

### HasWeaponRange

`func (o *Equipment) HasWeaponRange() bool`

HasWeaponRange returns a boolean if a field has been set.

### GetCategoryRange

`func (o *Equipment) GetCategoryRange() string`

GetCategoryRange returns the CategoryRange field if non-nil, zero value otherwise.

### GetCategoryRangeOk

`func (o *Equipment) GetCategoryRangeOk() (*string, bool)`

GetCategoryRangeOk returns a tuple with the CategoryRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryRange

`func (o *Equipment) SetCategoryRange(v string)`

SetCategoryRange sets CategoryRange field to given value.

### HasCategoryRange

`func (o *Equipment) HasCategoryRange() bool`

HasCategoryRange returns a boolean if a field has been set.

### GetRange

`func (o *Equipment) GetRange() WeaponAllOfRange`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *Equipment) GetRangeOk() (*WeaponAllOfRange, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *Equipment) SetRange(v WeaponAllOfRange)`

SetRange sets Range field to given value.

### HasRange

`func (o *Equipment) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetDamage

`func (o *Equipment) GetDamage() Damage`

GetDamage returns the Damage field if non-nil, zero value otherwise.

### GetDamageOk

`func (o *Equipment) GetDamageOk() (*Damage, bool)`

GetDamageOk returns a tuple with the Damage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamage

`func (o *Equipment) SetDamage(v Damage)`

SetDamage sets Damage field to given value.

### HasDamage

`func (o *Equipment) HasDamage() bool`

HasDamage returns a boolean if a field has been set.

### GetTwoHandedDamage

`func (o *Equipment) GetTwoHandedDamage() Damage`

GetTwoHandedDamage returns the TwoHandedDamage field if non-nil, zero value otherwise.

### GetTwoHandedDamageOk

`func (o *Equipment) GetTwoHandedDamageOk() (*Damage, bool)`

GetTwoHandedDamageOk returns a tuple with the TwoHandedDamage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoHandedDamage

`func (o *Equipment) SetTwoHandedDamage(v Damage)`

SetTwoHandedDamage sets TwoHandedDamage field to given value.

### HasTwoHandedDamage

`func (o *Equipment) HasTwoHandedDamage() bool`

HasTwoHandedDamage returns a boolean if a field has been set.

### GetProperties

`func (o *Equipment) GetProperties() []APIReference`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Equipment) GetPropertiesOk() (*[]APIReference, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Equipment) SetProperties(v []APIReference)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Equipment) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetCost

`func (o *Equipment) GetCost() Cost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *Equipment) GetCostOk() (*Cost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *Equipment) SetCost(v Cost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *Equipment) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetWeight

`func (o *Equipment) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *Equipment) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *Equipment) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *Equipment) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetArmorCategory

`func (o *Equipment) GetArmorCategory() string`

GetArmorCategory returns the ArmorCategory field if non-nil, zero value otherwise.

### GetArmorCategoryOk

`func (o *Equipment) GetArmorCategoryOk() (*string, bool)`

GetArmorCategoryOk returns a tuple with the ArmorCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArmorCategory

`func (o *Equipment) SetArmorCategory(v string)`

SetArmorCategory sets ArmorCategory field to given value.

### HasArmorCategory

`func (o *Equipment) HasArmorCategory() bool`

HasArmorCategory returns a boolean if a field has been set.

### GetArmorClass

`func (o *Equipment) GetArmorClass() map[string]string`

GetArmorClass returns the ArmorClass field if non-nil, zero value otherwise.

### GetArmorClassOk

`func (o *Equipment) GetArmorClassOk() (*map[string]string, bool)`

GetArmorClassOk returns a tuple with the ArmorClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArmorClass

`func (o *Equipment) SetArmorClass(v map[string]string)`

SetArmorClass sets ArmorClass field to given value.

### HasArmorClass

`func (o *Equipment) HasArmorClass() bool`

HasArmorClass returns a boolean if a field has been set.

### GetStrMinimum

`func (o *Equipment) GetStrMinimum() float32`

GetStrMinimum returns the StrMinimum field if non-nil, zero value otherwise.

### GetStrMinimumOk

`func (o *Equipment) GetStrMinimumOk() (*float32, bool)`

GetStrMinimumOk returns a tuple with the StrMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrMinimum

`func (o *Equipment) SetStrMinimum(v float32)`

SetStrMinimum sets StrMinimum field to given value.

### HasStrMinimum

`func (o *Equipment) HasStrMinimum() bool`

HasStrMinimum returns a boolean if a field has been set.

### GetStealthDisadvantage

`func (o *Equipment) GetStealthDisadvantage() bool`

GetStealthDisadvantage returns the StealthDisadvantage field if non-nil, zero value otherwise.

### GetStealthDisadvantageOk

`func (o *Equipment) GetStealthDisadvantageOk() (*bool, bool)`

GetStealthDisadvantageOk returns a tuple with the StealthDisadvantage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStealthDisadvantage

`func (o *Equipment) SetStealthDisadvantage(v bool)`

SetStealthDisadvantage sets StealthDisadvantage field to given value.

### HasStealthDisadvantage

`func (o *Equipment) HasStealthDisadvantage() bool`

HasStealthDisadvantage returns a boolean if a field has been set.

### GetGearCategory

`func (o *Equipment) GetGearCategory() APIReference`

GetGearCategory returns the GearCategory field if non-nil, zero value otherwise.

### GetGearCategoryOk

`func (o *Equipment) GetGearCategoryOk() (*APIReference, bool)`

GetGearCategoryOk returns a tuple with the GearCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGearCategory

`func (o *Equipment) SetGearCategory(v APIReference)`

SetGearCategory sets GearCategory field to given value.

### HasGearCategory

`func (o *Equipment) HasGearCategory() bool`

HasGearCategory returns a boolean if a field has been set.

### GetContents

`func (o *Equipment) GetContents() []APIReference`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *Equipment) GetContentsOk() (*[]APIReference, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *Equipment) SetContents(v []APIReference)`

SetContents sets Contents field to given value.

### HasContents

`func (o *Equipment) HasContents() bool`

HasContents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


