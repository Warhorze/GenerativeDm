# Prerequisite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbilityScore** | Pointer to [**PrerequisiteAbilityScore**](PrerequisiteAbilityScore.md) |  | [optional] 
**MinimumScore** | Pointer to **float32** | Minimum score to meet the prerequisite. | [optional] 

## Methods

### NewPrerequisite

`func NewPrerequisite() *Prerequisite`

NewPrerequisite instantiates a new Prerequisite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrerequisiteWithDefaults

`func NewPrerequisiteWithDefaults() *Prerequisite`

NewPrerequisiteWithDefaults instantiates a new Prerequisite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbilityScore

`func (o *Prerequisite) GetAbilityScore() PrerequisiteAbilityScore`

GetAbilityScore returns the AbilityScore field if non-nil, zero value otherwise.

### GetAbilityScoreOk

`func (o *Prerequisite) GetAbilityScoreOk() (*PrerequisiteAbilityScore, bool)`

GetAbilityScoreOk returns a tuple with the AbilityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbilityScore

`func (o *Prerequisite) SetAbilityScore(v PrerequisiteAbilityScore)`

SetAbilityScore sets AbilityScore field to given value.

### HasAbilityScore

`func (o *Prerequisite) HasAbilityScore() bool`

HasAbilityScore returns a boolean if a field has been set.

### GetMinimumScore

`func (o *Prerequisite) GetMinimumScore() float32`

GetMinimumScore returns the MinimumScore field if non-nil, zero value otherwise.

### GetMinimumScoreOk

`func (o *Prerequisite) GetMinimumScoreOk() (*float32, bool)`

GetMinimumScoreOk returns a tuple with the MinimumScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumScore

`func (o *Prerequisite) SetMinimumScore(v float32)`

SetMinimumScore sets MinimumScore field to given value.

### HasMinimumScore

`func (o *Prerequisite) HasMinimumScore() bool`

HasMinimumScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


