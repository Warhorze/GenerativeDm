# Skill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **string** | Resource index for shorthand searching. | [optional] 
**Name** | Pointer to **string** | Name of the referenced resource. | [optional] 
**Url** | Pointer to **string** | URL of the referenced resource. | [optional] 
**Desc** | Pointer to **[]string** | Description of the resource. | [optional] 
**AbilityScore** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 

## Methods

### NewSkill

`func NewSkill() *Skill`

NewSkill instantiates a new Skill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillWithDefaults

`func NewSkillWithDefaults() *Skill`

NewSkillWithDefaults instantiates a new Skill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *Skill) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *Skill) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *Skill) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *Skill) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetName

`func (o *Skill) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Skill) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Skill) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Skill) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *Skill) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Skill) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Skill) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Skill) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDesc

`func (o *Skill) GetDesc() []string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *Skill) GetDescOk() (*[]string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *Skill) SetDesc(v []string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *Skill) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetAbilityScore

`func (o *Skill) GetAbilityScore() APIReference`

GetAbilityScore returns the AbilityScore field if non-nil, zero value otherwise.

### GetAbilityScoreOk

`func (o *Skill) GetAbilityScoreOk() (*APIReference, bool)`

GetAbilityScoreOk returns a tuple with the AbilityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbilityScore

`func (o *Skill) SetAbilityScore(v APIReference)`

SetAbilityScore sets AbilityScore field to given value.

### HasAbilityScore

`func (o *Skill) HasAbilityScore() bool`

HasAbilityScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


