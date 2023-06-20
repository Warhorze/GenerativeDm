# EquipmentCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **string** | Resource index for shorthand searching. | [optional] 
**Name** | Pointer to **string** | Name of the referenced resource. | [optional] 
**Url** | Pointer to **string** | URL of the referenced resource. | [optional] 
**Equipment** | Pointer to [**[]APIReference**](APIReference.md) | A list of the equipment that falls into this category. | [optional] 

## Methods

### NewEquipmentCategory

`func NewEquipmentCategory() *EquipmentCategory`

NewEquipmentCategory instantiates a new EquipmentCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentCategoryWithDefaults

`func NewEquipmentCategoryWithDefaults() *EquipmentCategory`

NewEquipmentCategoryWithDefaults instantiates a new EquipmentCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *EquipmentCategory) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *EquipmentCategory) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *EquipmentCategory) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *EquipmentCategory) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetName

`func (o *EquipmentCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EquipmentCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EquipmentCategory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EquipmentCategory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *EquipmentCategory) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EquipmentCategory) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EquipmentCategory) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *EquipmentCategory) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetEquipment

`func (o *EquipmentCategory) GetEquipment() []APIReference`

GetEquipment returns the Equipment field if non-nil, zero value otherwise.

### GetEquipmentOk

`func (o *EquipmentCategory) GetEquipmentOk() (*[]APIReference, bool)`

GetEquipmentOk returns a tuple with the Equipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipment

`func (o *EquipmentCategory) SetEquipment(v []APIReference)`

SetEquipment sets Equipment field to given value.

### HasEquipment

`func (o *EquipmentCategory) HasEquipment() bool`

HasEquipment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


