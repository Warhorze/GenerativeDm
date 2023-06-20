# Gear

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **string** | Resource index for shorthand searching. | [optional] 
**Name** | Pointer to **string** | Name of the referenced resource. | [optional] 
**Url** | Pointer to **string** | URL of the referenced resource. | [optional] 
**Desc** | Pointer to **[]string** | Description of the resource. | [optional] 
**EquipmentCategory** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 
**GearCategory** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 
**Cost** | Pointer to [**Cost**](Cost.md) |  | [optional] 
**Weight** | Pointer to **float32** | How much the equipment weighs. | [optional] 

## Methods

### NewGear

`func NewGear() *Gear`

NewGear instantiates a new Gear object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGearWithDefaults

`func NewGearWithDefaults() *Gear`

NewGearWithDefaults instantiates a new Gear object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *Gear) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *Gear) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *Gear) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *Gear) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetName

`func (o *Gear) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Gear) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Gear) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Gear) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *Gear) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Gear) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Gear) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Gear) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDesc

`func (o *Gear) GetDesc() []string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *Gear) GetDescOk() (*[]string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *Gear) SetDesc(v []string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *Gear) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetEquipmentCategory

`func (o *Gear) GetEquipmentCategory() APIReference`

GetEquipmentCategory returns the EquipmentCategory field if non-nil, zero value otherwise.

### GetEquipmentCategoryOk

`func (o *Gear) GetEquipmentCategoryOk() (*APIReference, bool)`

GetEquipmentCategoryOk returns a tuple with the EquipmentCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentCategory

`func (o *Gear) SetEquipmentCategory(v APIReference)`

SetEquipmentCategory sets EquipmentCategory field to given value.

### HasEquipmentCategory

`func (o *Gear) HasEquipmentCategory() bool`

HasEquipmentCategory returns a boolean if a field has been set.

### GetGearCategory

`func (o *Gear) GetGearCategory() APIReference`

GetGearCategory returns the GearCategory field if non-nil, zero value otherwise.

### GetGearCategoryOk

`func (o *Gear) GetGearCategoryOk() (*APIReference, bool)`

GetGearCategoryOk returns a tuple with the GearCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGearCategory

`func (o *Gear) SetGearCategory(v APIReference)`

SetGearCategory sets GearCategory field to given value.

### HasGearCategory

`func (o *Gear) HasGearCategory() bool`

HasGearCategory returns a boolean if a field has been set.

### GetCost

`func (o *Gear) GetCost() Cost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *Gear) GetCostOk() (*Cost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *Gear) SetCost(v Cost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *Gear) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetWeight

`func (o *Gear) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *Gear) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *Gear) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *Gear) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


