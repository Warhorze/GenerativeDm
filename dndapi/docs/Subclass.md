# Subclass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **string** | Resource index for shorthand searching. | [optional] 
**Name** | Pointer to **string** | Name of the referenced resource. | [optional] 
**Url** | Pointer to **string** | URL of the referenced resource. | [optional] 
**Desc** | Pointer to **[]string** | Description of the resource. | [optional] 
**Class** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 
**SubclassFlavor** | Pointer to **string** | Lore-friendly flavor text for a classes respective subclass. | [optional] 
**SubclassLevels** | Pointer to **string** | Resource url that shows the subclass level progression. | [optional] 
**Spells** | Pointer to [**[]SubclassAllOfSpells**](SubclassAllOfSpells.md) |  | [optional] 

## Methods

### NewSubclass

`func NewSubclass() *Subclass`

NewSubclass instantiates a new Subclass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubclassWithDefaults

`func NewSubclassWithDefaults() *Subclass`

NewSubclassWithDefaults instantiates a new Subclass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *Subclass) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *Subclass) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *Subclass) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *Subclass) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetName

`func (o *Subclass) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Subclass) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Subclass) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Subclass) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *Subclass) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Subclass) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Subclass) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Subclass) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDesc

`func (o *Subclass) GetDesc() []string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *Subclass) GetDescOk() (*[]string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *Subclass) SetDesc(v []string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *Subclass) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetClass

`func (o *Subclass) GetClass() APIReference`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *Subclass) GetClassOk() (*APIReference, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *Subclass) SetClass(v APIReference)`

SetClass sets Class field to given value.

### HasClass

`func (o *Subclass) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetSubclassFlavor

`func (o *Subclass) GetSubclassFlavor() string`

GetSubclassFlavor returns the SubclassFlavor field if non-nil, zero value otherwise.

### GetSubclassFlavorOk

`func (o *Subclass) GetSubclassFlavorOk() (*string, bool)`

GetSubclassFlavorOk returns a tuple with the SubclassFlavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubclassFlavor

`func (o *Subclass) SetSubclassFlavor(v string)`

SetSubclassFlavor sets SubclassFlavor field to given value.

### HasSubclassFlavor

`func (o *Subclass) HasSubclassFlavor() bool`

HasSubclassFlavor returns a boolean if a field has been set.

### GetSubclassLevels

`func (o *Subclass) GetSubclassLevels() string`

GetSubclassLevels returns the SubclassLevels field if non-nil, zero value otherwise.

### GetSubclassLevelsOk

`func (o *Subclass) GetSubclassLevelsOk() (*string, bool)`

GetSubclassLevelsOk returns a tuple with the SubclassLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubclassLevels

`func (o *Subclass) SetSubclassLevels(v string)`

SetSubclassLevels sets SubclassLevels field to given value.

### HasSubclassLevels

`func (o *Subclass) HasSubclassLevels() bool`

HasSubclassLevels returns a boolean if a field has been set.

### GetSpells

`func (o *Subclass) GetSpells() []SubclassAllOfSpells`

GetSpells returns the Spells field if non-nil, zero value otherwise.

### GetSpellsOk

`func (o *Subclass) GetSpellsOk() (*[]SubclassAllOfSpells, bool)`

GetSpellsOk returns a tuple with the Spells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpells

`func (o *Subclass) SetSpells(v []SubclassAllOfSpells)`

SetSpells sets Spells field to given value.

### HasSpells

`func (o *Subclass) HasSpells() bool`

HasSpells returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


