# AbilityBonus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bonus** | Pointer to **float32** | Bonus amount for this ability score. | [optional] 
**AbilityScore** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 

## Methods

### NewAbilityBonus

`func NewAbilityBonus() *AbilityBonus`

NewAbilityBonus instantiates a new AbilityBonus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbilityBonusWithDefaults

`func NewAbilityBonusWithDefaults() *AbilityBonus`

NewAbilityBonusWithDefaults instantiates a new AbilityBonus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBonus

`func (o *AbilityBonus) GetBonus() float32`

GetBonus returns the Bonus field if non-nil, zero value otherwise.

### GetBonusOk

`func (o *AbilityBonus) GetBonusOk() (*float32, bool)`

GetBonusOk returns a tuple with the Bonus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonus

`func (o *AbilityBonus) SetBonus(v float32)`

SetBonus sets Bonus field to given value.

### HasBonus

`func (o *AbilityBonus) HasBonus() bool`

HasBonus returns a boolean if a field has been set.

### GetAbilityScore

`func (o *AbilityBonus) GetAbilityScore() APIReference`

GetAbilityScore returns the AbilityScore field if non-nil, zero value otherwise.

### GetAbilityScoreOk

`func (o *AbilityBonus) GetAbilityScoreOk() (*APIReference, bool)`

GetAbilityScoreOk returns a tuple with the AbilityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbilityScore

`func (o *AbilityBonus) SetAbilityScore(v APIReference)`

SetAbilityScore sets AbilityScore field to given value.

### HasAbilityScore

`func (o *AbilityBonus) HasAbilityScore() bool`

HasAbilityScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


