# OptionOneOf7

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OptionType** | Pointer to **string** | Type of option; determines other attributes. | [optional] 
**AbilityScore** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 
**MinimumScore** | Pointer to **float32** | The minimum score required to satisfy the prerequisite. | [optional] 

## Methods

### NewOptionOneOf7

`func NewOptionOneOf7() *OptionOneOf7`

NewOptionOneOf7 instantiates a new OptionOneOf7 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionOneOf7WithDefaults

`func NewOptionOneOf7WithDefaults() *OptionOneOf7`

NewOptionOneOf7WithDefaults instantiates a new OptionOneOf7 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptionType

`func (o *OptionOneOf7) GetOptionType() string`

GetOptionType returns the OptionType field if non-nil, zero value otherwise.

### GetOptionTypeOk

`func (o *OptionOneOf7) GetOptionTypeOk() (*string, bool)`

GetOptionTypeOk returns a tuple with the OptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionType

`func (o *OptionOneOf7) SetOptionType(v string)`

SetOptionType sets OptionType field to given value.

### HasOptionType

`func (o *OptionOneOf7) HasOptionType() bool`

HasOptionType returns a boolean if a field has been set.

### GetAbilityScore

`func (o *OptionOneOf7) GetAbilityScore() APIReference`

GetAbilityScore returns the AbilityScore field if non-nil, zero value otherwise.

### GetAbilityScoreOk

`func (o *OptionOneOf7) GetAbilityScoreOk() (*APIReference, bool)`

GetAbilityScoreOk returns a tuple with the AbilityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbilityScore

`func (o *OptionOneOf7) SetAbilityScore(v APIReference)`

SetAbilityScore sets AbilityScore field to given value.

### HasAbilityScore

`func (o *OptionOneOf7) HasAbilityScore() bool`

HasAbilityScore returns a boolean if a field has been set.

### GetMinimumScore

`func (o *OptionOneOf7) GetMinimumScore() float32`

GetMinimumScore returns the MinimumScore field if non-nil, zero value otherwise.

### GetMinimumScoreOk

`func (o *OptionOneOf7) GetMinimumScoreOk() (*float32, bool)`

GetMinimumScoreOk returns a tuple with the MinimumScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumScore

`func (o *OptionOneOf7) SetMinimumScore(v float32)`

SetMinimumScore sets MinimumScore field to given value.

### HasMinimumScore

`func (o *OptionOneOf7) HasMinimumScore() bool`

HasMinimumScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


