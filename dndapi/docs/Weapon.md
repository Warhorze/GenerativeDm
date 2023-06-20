# Weapon

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

## Methods

### NewWeapon

`func NewWeapon() *Weapon`

NewWeapon instantiates a new Weapon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWeaponWithDefaults

`func NewWeaponWithDefaults() *Weapon`

NewWeaponWithDefaults instantiates a new Weapon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *Weapon) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *Weapon) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *Weapon) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *Weapon) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetName

`func (o *Weapon) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Weapon) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Weapon) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Weapon) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *Weapon) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Weapon) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Weapon) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Weapon) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDesc

`func (o *Weapon) GetDesc() []string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *Weapon) GetDescOk() (*[]string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *Weapon) SetDesc(v []string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *Weapon) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetEquipmentCategory

`func (o *Weapon) GetEquipmentCategory() APIReference`

GetEquipmentCategory returns the EquipmentCategory field if non-nil, zero value otherwise.

### GetEquipmentCategoryOk

`func (o *Weapon) GetEquipmentCategoryOk() (*APIReference, bool)`

GetEquipmentCategoryOk returns a tuple with the EquipmentCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentCategory

`func (o *Weapon) SetEquipmentCategory(v APIReference)`

SetEquipmentCategory sets EquipmentCategory field to given value.

### HasEquipmentCategory

`func (o *Weapon) HasEquipmentCategory() bool`

HasEquipmentCategory returns a boolean if a field has been set.

### GetWeaponCategory

`func (o *Weapon) GetWeaponCategory() string`

GetWeaponCategory returns the WeaponCategory field if non-nil, zero value otherwise.

### GetWeaponCategoryOk

`func (o *Weapon) GetWeaponCategoryOk() (*string, bool)`

GetWeaponCategoryOk returns a tuple with the WeaponCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeaponCategory

`func (o *Weapon) SetWeaponCategory(v string)`

SetWeaponCategory sets WeaponCategory field to given value.

### HasWeaponCategory

`func (o *Weapon) HasWeaponCategory() bool`

HasWeaponCategory returns a boolean if a field has been set.

### GetWeaponRange

`func (o *Weapon) GetWeaponRange() string`

GetWeaponRange returns the WeaponRange field if non-nil, zero value otherwise.

### GetWeaponRangeOk

`func (o *Weapon) GetWeaponRangeOk() (*string, bool)`

GetWeaponRangeOk returns a tuple with the WeaponRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeaponRange

`func (o *Weapon) SetWeaponRange(v string)`

SetWeaponRange sets WeaponRange field to given value.

### HasWeaponRange

`func (o *Weapon) HasWeaponRange() bool`

HasWeaponRange returns a boolean if a field has been set.

### GetCategoryRange

`func (o *Weapon) GetCategoryRange() string`

GetCategoryRange returns the CategoryRange field if non-nil, zero value otherwise.

### GetCategoryRangeOk

`func (o *Weapon) GetCategoryRangeOk() (*string, bool)`

GetCategoryRangeOk returns a tuple with the CategoryRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryRange

`func (o *Weapon) SetCategoryRange(v string)`

SetCategoryRange sets CategoryRange field to given value.

### HasCategoryRange

`func (o *Weapon) HasCategoryRange() bool`

HasCategoryRange returns a boolean if a field has been set.

### GetRange

`func (o *Weapon) GetRange() WeaponAllOfRange`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *Weapon) GetRangeOk() (*WeaponAllOfRange, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *Weapon) SetRange(v WeaponAllOfRange)`

SetRange sets Range field to given value.

### HasRange

`func (o *Weapon) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetDamage

`func (o *Weapon) GetDamage() Damage`

GetDamage returns the Damage field if non-nil, zero value otherwise.

### GetDamageOk

`func (o *Weapon) GetDamageOk() (*Damage, bool)`

GetDamageOk returns a tuple with the Damage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamage

`func (o *Weapon) SetDamage(v Damage)`

SetDamage sets Damage field to given value.

### HasDamage

`func (o *Weapon) HasDamage() bool`

HasDamage returns a boolean if a field has been set.

### GetTwoHandedDamage

`func (o *Weapon) GetTwoHandedDamage() Damage`

GetTwoHandedDamage returns the TwoHandedDamage field if non-nil, zero value otherwise.

### GetTwoHandedDamageOk

`func (o *Weapon) GetTwoHandedDamageOk() (*Damage, bool)`

GetTwoHandedDamageOk returns a tuple with the TwoHandedDamage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoHandedDamage

`func (o *Weapon) SetTwoHandedDamage(v Damage)`

SetTwoHandedDamage sets TwoHandedDamage field to given value.

### HasTwoHandedDamage

`func (o *Weapon) HasTwoHandedDamage() bool`

HasTwoHandedDamage returns a boolean if a field has been set.

### GetProperties

`func (o *Weapon) GetProperties() []APIReference`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Weapon) GetPropertiesOk() (*[]APIReference, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Weapon) SetProperties(v []APIReference)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Weapon) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetCost

`func (o *Weapon) GetCost() Cost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *Weapon) GetCostOk() (*Cost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *Weapon) SetCost(v Cost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *Weapon) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetWeight

`func (o *Weapon) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *Weapon) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *Weapon) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *Weapon) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


