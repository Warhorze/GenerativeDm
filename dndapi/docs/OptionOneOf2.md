# OptionOneOf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OptionType** | Pointer to **string** | Type of option; determines other attributes. | [optional] 
**Items** | Pointer to [**[]Option**](Option.md) |  | [optional] 

## Methods

### NewOptionOneOf2

`func NewOptionOneOf2() *OptionOneOf2`

NewOptionOneOf2 instantiates a new OptionOneOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionOneOf2WithDefaults

`func NewOptionOneOf2WithDefaults() *OptionOneOf2`

NewOptionOneOf2WithDefaults instantiates a new OptionOneOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptionType

`func (o *OptionOneOf2) GetOptionType() string`

GetOptionType returns the OptionType field if non-nil, zero value otherwise.

### GetOptionTypeOk

`func (o *OptionOneOf2) GetOptionTypeOk() (*string, bool)`

GetOptionTypeOk returns a tuple with the OptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionType

`func (o *OptionOneOf2) SetOptionType(v string)`

SetOptionType sets OptionType field to given value.

### HasOptionType

`func (o *OptionOneOf2) HasOptionType() bool`

HasOptionType returns a boolean if a field has been set.

### GetItems

`func (o *OptionOneOf2) GetItems() []Option`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *OptionOneOf2) GetItemsOk() (*[]Option, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *OptionOneOf2) SetItems(v []Option)`

SetItems sets Items field to given value.

### HasItems

`func (o *OptionOneOf2) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


