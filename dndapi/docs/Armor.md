# Armor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **string** | Resource index for shorthand searching. | [optional] 
**Name** | Pointer to **string** | Name of the referenced resource. | [optional] 
**Url** | Pointer to **string** | URL of the referenced resource. | [optional] 
**Desc** | Pointer to **[]string** | Description of the resource. | [optional] 
**EquipmentCategory** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 
**ArmorCategory** | Pointer to **string** | The category of armor this falls into. | [optional] 
**ArmorClass** | Pointer to **map[string]string** | Details on how to calculate armor class. | [optional] 
**StrMinimum** | Pointer to **float32** | Minimum STR required to use this armor. | [optional] 
**StealthDisadvantage** | Pointer to **bool** | Whether the armor gives disadvantage for Stealth. | [optional] 
**Cost** | Pointer to [**Cost**](Cost.md) |  | [optional] 
**Weight** | Pointer to **float32** | How much the equipment weighs. | [optional] 

## Methods

### NewArmor

`func NewArmor() *Armor`

NewArmor instantiates a new Armor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArmorWithDefaults

`func NewArmorWithDefaults() *Armor`

NewArmorWithDefaults instantiates a new Armor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *Armor) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *Armor) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *Armor) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *Armor) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetName

`func (o *Armor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Armor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Armor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Armor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *Armor) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Armor) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Armor) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Armor) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDesc

`func (o *Armor) GetDesc() []string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *Armor) GetDescOk() (*[]string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *Armor) SetDesc(v []string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *Armor) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetEquipmentCategory

`func (o *Armor) GetEquipmentCategory() APIReference`

GetEquipmentCategory returns the EquipmentCategory field if non-nil, zero value otherwise.

### GetEquipmentCategoryOk

`func (o *Armor) GetEquipmentCategoryOk() (*APIReference, bool)`

GetEquipmentCategoryOk returns a tuple with the EquipmentCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentCategory

`func (o *Armor) SetEquipmentCategory(v APIReference)`

SetEquipmentCategory sets EquipmentCategory field to given value.

### HasEquipmentCategory

`func (o *Armor) HasEquipmentCategory() bool`

HasEquipmentCategory returns a boolean if a field has been set.

### GetArmorCategory

`func (o *Armor) GetArmorCategory() string`

GetArmorCategory returns the ArmorCategory field if non-nil, zero value otherwise.

### GetArmorCategoryOk

`func (o *Armor) GetArmorCategoryOk() (*string, bool)`

GetArmorCategoryOk returns a tuple with the ArmorCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArmorCategory

`func (o *Armor) SetArmorCategory(v string)`

SetArmorCategory sets ArmorCategory field to given value.

### HasArmorCategory

`func (o *Armor) HasArmorCategory() bool`

HasArmorCategory returns a boolean if a field has been set.

### GetArmorClass

`func (o *Armor) GetArmorClass() map[string]string`

GetArmorClass returns the ArmorClass field if non-nil, zero value otherwise.

### GetArmorClassOk

`func (o *Armor) GetArmorClassOk() (*map[string]string, bool)`

GetArmorClassOk returns a tuple with the ArmorClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArmorClass

`func (o *Armor) SetArmorClass(v map[string]string)`

SetArmorClass sets ArmorClass field to given value.

### HasArmorClass

`func (o *Armor) HasArmorClass() bool`

HasArmorClass returns a boolean if a field has been set.

### GetStrMinimum

`func (o *Armor) GetStrMinimum() float32`

GetStrMinimum returns the StrMinimum field if non-nil, zero value otherwise.

### GetStrMinimumOk

`func (o *Armor) GetStrMinimumOk() (*float32, bool)`

GetStrMinimumOk returns a tuple with the StrMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrMinimum

`func (o *Armor) SetStrMinimum(v float32)`

SetStrMinimum sets StrMinimum field to given value.

### HasStrMinimum

`func (o *Armor) HasStrMinimum() bool`

HasStrMinimum returns a boolean if a field has been set.

### GetStealthDisadvantage

`func (o *Armor) GetStealthDisadvantage() bool`

GetStealthDisadvantage returns the StealthDisadvantage field if non-nil, zero value otherwise.

### GetStealthDisadvantageOk

`func (o *Armor) GetStealthDisadvantageOk() (*bool, bool)`

GetStealthDisadvantageOk returns a tuple with the StealthDisadvantage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStealthDisadvantage

`func (o *Armor) SetStealthDisadvantage(v bool)`

SetStealthDisadvantage sets StealthDisadvantage field to given value.

### HasStealthDisadvantage

`func (o *Armor) HasStealthDisadvantage() bool`

HasStealthDisadvantage returns a boolean if a field has been set.

### GetCost

`func (o *Armor) GetCost() Cost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *Armor) GetCostOk() (*Cost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *Armor) SetCost(v Cost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *Armor) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetWeight

`func (o *Armor) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *Armor) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *Armor) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *Armor) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


