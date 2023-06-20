# OptionOneOf8

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OptionType** | Pointer to **string** | Type of option; determines other attributes. | [optional] 
**AbilityScore** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 
**Bonus** | Pointer to **float32** | The bonus being applied to the ability score | [optional] 

## Methods

### NewOptionOneOf8

`func NewOptionOneOf8() *OptionOneOf8`

NewOptionOneOf8 instantiates a new OptionOneOf8 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionOneOf8WithDefaults

`func NewOptionOneOf8WithDefaults() *OptionOneOf8`

NewOptionOneOf8WithDefaults instantiates a new OptionOneOf8 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptionType

`func (o *OptionOneOf8) GetOptionType() string`

GetOptionType returns the OptionType field if non-nil, zero value otherwise.

### GetOptionTypeOk

`func (o *OptionOneOf8) GetOptionTypeOk() (*string, bool)`

GetOptionTypeOk returns a tuple with the OptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionType

`func (o *OptionOneOf8) SetOptionType(v string)`

SetOptionType sets OptionType field to given value.

### HasOptionType

`func (o *OptionOneOf8) HasOptionType() bool`

HasOptionType returns a boolean if a field has been set.

### GetAbilityScore

`func (o *OptionOneOf8) GetAbilityScore() APIReference`

GetAbilityScore returns the AbilityScore field if non-nil, zero value otherwise.

### GetAbilityScoreOk

`func (o *OptionOneOf8) GetAbilityScoreOk() (*APIReference, bool)`

GetAbilityScoreOk returns a tuple with the AbilityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbilityScore

`func (o *OptionOneOf8) SetAbilityScore(v APIReference)`

SetAbilityScore sets AbilityScore field to given value.

### HasAbilityScore

`func (o *OptionOneOf8) HasAbilityScore() bool`

HasAbilityScore returns a boolean if a field has been set.

### GetBonus

`func (o *OptionOneOf8) GetBonus() float32`

GetBonus returns the Bonus field if non-nil, zero value otherwise.

### GetBonusOk

`func (o *OptionOneOf8) GetBonusOk() (*float32, bool)`

GetBonusOk returns a tuple with the Bonus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonus

`func (o *OptionOneOf8) SetBonus(v float32)`

SetBonus sets Bonus field to given value.

### HasBonus

`func (o *OptionOneOf8) HasBonus() bool`

HasBonus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


