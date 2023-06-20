# OptionOneOf6

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OptionType** | Pointer to **string** | Type of option; determines other attributes. | [optional] 
**Count** | Pointer to **float32** | Count | [optional] 
**Of** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 

## Methods

### NewOptionOneOf6

`func NewOptionOneOf6() *OptionOneOf6`

NewOptionOneOf6 instantiates a new OptionOneOf6 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionOneOf6WithDefaults

`func NewOptionOneOf6WithDefaults() *OptionOneOf6`

NewOptionOneOf6WithDefaults instantiates a new OptionOneOf6 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptionType

`func (o *OptionOneOf6) GetOptionType() string`

GetOptionType returns the OptionType field if non-nil, zero value otherwise.

### GetOptionTypeOk

`func (o *OptionOneOf6) GetOptionTypeOk() (*string, bool)`

GetOptionTypeOk returns a tuple with the OptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionType

`func (o *OptionOneOf6) SetOptionType(v string)`

SetOptionType sets OptionType field to given value.

### HasOptionType

`func (o *OptionOneOf6) HasOptionType() bool`

HasOptionType returns a boolean if a field has been set.

### GetCount

`func (o *OptionOneOf6) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *OptionOneOf6) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *OptionOneOf6) SetCount(v float32)`

SetCount sets Count field to given value.

### HasCount

`func (o *OptionOneOf6) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetOf

`func (o *OptionOneOf6) GetOf() APIReference`

GetOf returns the Of field if non-nil, zero value otherwise.

### GetOfOk

`func (o *OptionOneOf6) GetOfOk() (*APIReference, bool)`

GetOfOk returns a tuple with the Of field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOf

`func (o *OptionOneOf6) SetOf(v APIReference)`

SetOf sets Of field to given value.

### HasOf

`func (o *OptionOneOf6) HasOf() bool`

HasOf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


