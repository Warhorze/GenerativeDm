# OptionSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OptionSetType** | Pointer to **string** | Type of option set; determines other attributes. | [optional] 
**OptionsArray** | Pointer to [**[]Option**](Option.md) | Array of options to choose from. | [optional] 
**EquipmentCategory** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 
**ResourceList** | Pointer to **string** | A reference (by URL) to a collection in the database. | [optional] 

## Methods

### NewOptionSet

`func NewOptionSet() *OptionSet`

NewOptionSet instantiates a new OptionSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionSetWithDefaults

`func NewOptionSetWithDefaults() *OptionSet`

NewOptionSetWithDefaults instantiates a new OptionSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptionSetType

`func (o *OptionSet) GetOptionSetType() string`

GetOptionSetType returns the OptionSetType field if non-nil, zero value otherwise.

### GetOptionSetTypeOk

`func (o *OptionSet) GetOptionSetTypeOk() (*string, bool)`

GetOptionSetTypeOk returns a tuple with the OptionSetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSetType

`func (o *OptionSet) SetOptionSetType(v string)`

SetOptionSetType sets OptionSetType field to given value.

### HasOptionSetType

`func (o *OptionSet) HasOptionSetType() bool`

HasOptionSetType returns a boolean if a field has been set.

### GetOptionsArray

`func (o *OptionSet) GetOptionsArray() []Option`

GetOptionsArray returns the OptionsArray field if non-nil, zero value otherwise.

### GetOptionsArrayOk

`func (o *OptionSet) GetOptionsArrayOk() (*[]Option, bool)`

GetOptionsArrayOk returns a tuple with the OptionsArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionsArray

`func (o *OptionSet) SetOptionsArray(v []Option)`

SetOptionsArray sets OptionsArray field to given value.

### HasOptionsArray

`func (o *OptionSet) HasOptionsArray() bool`

HasOptionsArray returns a boolean if a field has been set.

### GetEquipmentCategory

`func (o *OptionSet) GetEquipmentCategory() APIReference`

GetEquipmentCategory returns the EquipmentCategory field if non-nil, zero value otherwise.

### GetEquipmentCategoryOk

`func (o *OptionSet) GetEquipmentCategoryOk() (*APIReference, bool)`

GetEquipmentCategoryOk returns a tuple with the EquipmentCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentCategory

`func (o *OptionSet) SetEquipmentCategory(v APIReference)`

SetEquipmentCategory sets EquipmentCategory field to given value.

### HasEquipmentCategory

`func (o *OptionSet) HasEquipmentCategory() bool`

HasEquipmentCategory returns a boolean if a field has been set.

### GetResourceList

`func (o *OptionSet) GetResourceList() string`

GetResourceList returns the ResourceList field if non-nil, zero value otherwise.

### GetResourceListOk

`func (o *OptionSet) GetResourceListOk() (*string, bool)`

GetResourceListOk returns a tuple with the ResourceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceList

`func (o *OptionSet) SetResourceList(v string)`

SetResourceList sets ResourceList field to given value.

### HasResourceList

`func (o *OptionSet) HasResourceList() bool`

HasResourceList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


