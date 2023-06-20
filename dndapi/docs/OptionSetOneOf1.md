# OptionSetOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OptionSetType** | Pointer to **string** | Type of option set; determines other attributes. | [optional] 
**EquipmentCategory** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 

## Methods

### NewOptionSetOneOf1

`func NewOptionSetOneOf1() *OptionSetOneOf1`

NewOptionSetOneOf1 instantiates a new OptionSetOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionSetOneOf1WithDefaults

`func NewOptionSetOneOf1WithDefaults() *OptionSetOneOf1`

NewOptionSetOneOf1WithDefaults instantiates a new OptionSetOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptionSetType

`func (o *OptionSetOneOf1) GetOptionSetType() string`

GetOptionSetType returns the OptionSetType field if non-nil, zero value otherwise.

### GetOptionSetTypeOk

`func (o *OptionSetOneOf1) GetOptionSetTypeOk() (*string, bool)`

GetOptionSetTypeOk returns a tuple with the OptionSetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSetType

`func (o *OptionSetOneOf1) SetOptionSetType(v string)`

SetOptionSetType sets OptionSetType field to given value.

### HasOptionSetType

`func (o *OptionSetOneOf1) HasOptionSetType() bool`

HasOptionSetType returns a boolean if a field has been set.

### GetEquipmentCategory

`func (o *OptionSetOneOf1) GetEquipmentCategory() APIReference`

GetEquipmentCategory returns the EquipmentCategory field if non-nil, zero value otherwise.

### GetEquipmentCategoryOk

`func (o *OptionSetOneOf1) GetEquipmentCategoryOk() (*APIReference, bool)`

GetEquipmentCategoryOk returns a tuple with the EquipmentCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentCategory

`func (o *OptionSetOneOf1) SetEquipmentCategory(v APIReference)`

SetEquipmentCategory sets EquipmentCategory field to given value.

### HasEquipmentCategory

`func (o *OptionSetOneOf1) HasEquipmentCategory() bool`

HasEquipmentCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


