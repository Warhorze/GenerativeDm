# OptionOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OptionType** | Pointer to **string** | Type of option; determines other attributes. | [optional] 
**Item** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 

## Methods

### NewOptionOneOf

`func NewOptionOneOf() *OptionOneOf`

NewOptionOneOf instantiates a new OptionOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionOneOfWithDefaults

`func NewOptionOneOfWithDefaults() *OptionOneOf`

NewOptionOneOfWithDefaults instantiates a new OptionOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptionType

`func (o *OptionOneOf) GetOptionType() string`

GetOptionType returns the OptionType field if non-nil, zero value otherwise.

### GetOptionTypeOk

`func (o *OptionOneOf) GetOptionTypeOk() (*string, bool)`

GetOptionTypeOk returns a tuple with the OptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionType

`func (o *OptionOneOf) SetOptionType(v string)`

SetOptionType sets OptionType field to given value.

### HasOptionType

`func (o *OptionOneOf) HasOptionType() bool`

HasOptionType returns a boolean if a field has been set.

### GetItem

`func (o *OptionOneOf) GetItem() APIReference`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *OptionOneOf) GetItemOk() (*APIReference, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *OptionOneOf) SetItem(v APIReference)`

SetItem sets Item field to given value.

### HasItem

`func (o *OptionOneOf) HasItem() bool`

HasItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


