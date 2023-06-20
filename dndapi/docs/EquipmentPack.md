# EquipmentPack

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
**Contents** | Pointer to [**[]APIReference**](APIReference.md) | The list of adventuring gear in the pack. | [optional] 

## Methods

### NewEquipmentPack

`func NewEquipmentPack() *EquipmentPack`

NewEquipmentPack instantiates a new EquipmentPack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentPackWithDefaults

`func NewEquipmentPackWithDefaults() *EquipmentPack`

NewEquipmentPackWithDefaults instantiates a new EquipmentPack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *EquipmentPack) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *EquipmentPack) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *EquipmentPack) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *EquipmentPack) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetName

`func (o *EquipmentPack) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EquipmentPack) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EquipmentPack) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EquipmentPack) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *EquipmentPack) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EquipmentPack) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EquipmentPack) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *EquipmentPack) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDesc

`func (o *EquipmentPack) GetDesc() []string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *EquipmentPack) GetDescOk() (*[]string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *EquipmentPack) SetDesc(v []string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *EquipmentPack) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetEquipmentCategory

`func (o *EquipmentPack) GetEquipmentCategory() APIReference`

GetEquipmentCategory returns the EquipmentCategory field if non-nil, zero value otherwise.

### GetEquipmentCategoryOk

`func (o *EquipmentPack) GetEquipmentCategoryOk() (*APIReference, bool)`

GetEquipmentCategoryOk returns a tuple with the EquipmentCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentCategory

`func (o *EquipmentPack) SetEquipmentCategory(v APIReference)`

SetEquipmentCategory sets EquipmentCategory field to given value.

### HasEquipmentCategory

`func (o *EquipmentPack) HasEquipmentCategory() bool`

HasEquipmentCategory returns a boolean if a field has been set.

### GetGearCategory

`func (o *EquipmentPack) GetGearCategory() APIReference`

GetGearCategory returns the GearCategory field if non-nil, zero value otherwise.

### GetGearCategoryOk

`func (o *EquipmentPack) GetGearCategoryOk() (*APIReference, bool)`

GetGearCategoryOk returns a tuple with the GearCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGearCategory

`func (o *EquipmentPack) SetGearCategory(v APIReference)`

SetGearCategory sets GearCategory field to given value.

### HasGearCategory

`func (o *EquipmentPack) HasGearCategory() bool`

HasGearCategory returns a boolean if a field has been set.

### GetCost

`func (o *EquipmentPack) GetCost() Cost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *EquipmentPack) GetCostOk() (*Cost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *EquipmentPack) SetCost(v Cost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *EquipmentPack) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetContents

`func (o *EquipmentPack) GetContents() []APIReference`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *EquipmentPack) GetContentsOk() (*[]APIReference, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *EquipmentPack) SetContents(v []APIReference)`

SetContents sets Contents field to given value.

### HasContents

`func (o *EquipmentPack) HasContents() bool`

HasContents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


