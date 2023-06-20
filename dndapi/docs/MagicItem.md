# MagicItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **string** | Resource index for shorthand searching. | [optional] 
**Name** | Pointer to **string** | Name of the referenced resource. | [optional] 
**Url** | Pointer to **string** | URL of the referenced resource. | [optional] 
**Desc** | Pointer to **[]string** | Description of the resource. | [optional] 
**EquipmentCategory** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 
**Rarity** | Pointer to [**MagicItemAllOfRarity**](MagicItemAllOfRarity.md) |  | [optional] 
**Variants** | Pointer to [**[]APIReference**](APIReference.md) |  | [optional] 
**Variant** | Pointer to **bool** | Whether this is a variant or not | [optional] 

## Methods

### NewMagicItem

`func NewMagicItem() *MagicItem`

NewMagicItem instantiates a new MagicItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMagicItemWithDefaults

`func NewMagicItemWithDefaults() *MagicItem`

NewMagicItemWithDefaults instantiates a new MagicItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *MagicItem) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *MagicItem) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *MagicItem) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *MagicItem) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetName

`func (o *MagicItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MagicItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MagicItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MagicItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *MagicItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MagicItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MagicItem) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *MagicItem) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDesc

`func (o *MagicItem) GetDesc() []string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *MagicItem) GetDescOk() (*[]string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *MagicItem) SetDesc(v []string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *MagicItem) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetEquipmentCategory

`func (o *MagicItem) GetEquipmentCategory() APIReference`

GetEquipmentCategory returns the EquipmentCategory field if non-nil, zero value otherwise.

### GetEquipmentCategoryOk

`func (o *MagicItem) GetEquipmentCategoryOk() (*APIReference, bool)`

GetEquipmentCategoryOk returns a tuple with the EquipmentCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentCategory

`func (o *MagicItem) SetEquipmentCategory(v APIReference)`

SetEquipmentCategory sets EquipmentCategory field to given value.

### HasEquipmentCategory

`func (o *MagicItem) HasEquipmentCategory() bool`

HasEquipmentCategory returns a boolean if a field has been set.

### GetRarity

`func (o *MagicItem) GetRarity() MagicItemAllOfRarity`

GetRarity returns the Rarity field if non-nil, zero value otherwise.

### GetRarityOk

`func (o *MagicItem) GetRarityOk() (*MagicItemAllOfRarity, bool)`

GetRarityOk returns a tuple with the Rarity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRarity

`func (o *MagicItem) SetRarity(v MagicItemAllOfRarity)`

SetRarity sets Rarity field to given value.

### HasRarity

`func (o *MagicItem) HasRarity() bool`

HasRarity returns a boolean if a field has been set.

### GetVariants

`func (o *MagicItem) GetVariants() []APIReference`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *MagicItem) GetVariantsOk() (*[]APIReference, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *MagicItem) SetVariants(v []APIReference)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *MagicItem) HasVariants() bool`

HasVariants returns a boolean if a field has been set.

### GetVariant

`func (o *MagicItem) GetVariant() bool`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *MagicItem) GetVariantOk() (*bool, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *MagicItem) SetVariant(v bool)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *MagicItem) HasVariant() bool`

HasVariant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


