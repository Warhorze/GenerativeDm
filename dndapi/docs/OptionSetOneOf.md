# OptionSetOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OptionSetType** | Pointer to **string** | Type of option set; determines other attributes. | [optional] 
**OptionsArray** | Pointer to [**[]Option**](Option.md) | Array of options to choose from. | [optional] 

## Methods

### NewOptionSetOneOf

`func NewOptionSetOneOf() *OptionSetOneOf`

NewOptionSetOneOf instantiates a new OptionSetOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionSetOneOfWithDefaults

`func NewOptionSetOneOfWithDefaults() *OptionSetOneOf`

NewOptionSetOneOfWithDefaults instantiates a new OptionSetOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptionSetType

`func (o *OptionSetOneOf) GetOptionSetType() string`

GetOptionSetType returns the OptionSetType field if non-nil, zero value otherwise.

### GetOptionSetTypeOk

`func (o *OptionSetOneOf) GetOptionSetTypeOk() (*string, bool)`

GetOptionSetTypeOk returns a tuple with the OptionSetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSetType

`func (o *OptionSetOneOf) SetOptionSetType(v string)`

SetOptionSetType sets OptionSetType field to given value.

### HasOptionSetType

`func (o *OptionSetOneOf) HasOptionSetType() bool`

HasOptionSetType returns a boolean if a field has been set.

### GetOptionsArray

`func (o *OptionSetOneOf) GetOptionsArray() []Option`

GetOptionsArray returns the OptionsArray field if non-nil, zero value otherwise.

### GetOptionsArrayOk

`func (o *OptionSetOneOf) GetOptionsArrayOk() (*[]Option, bool)`

GetOptionsArrayOk returns a tuple with the OptionsArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionsArray

`func (o *OptionSetOneOf) SetOptionsArray(v []Option)`

SetOptionsArray sets OptionsArray field to given value.

### HasOptionsArray

`func (o *OptionSetOneOf) HasOptionsArray() bool`

HasOptionsArray returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


