# AbilityScore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **string** | Resource index for shorthand searching. | [optional] 
**Name** | Pointer to **string** | Name of the referenced resource. | [optional] 
**Url** | Pointer to **string** | URL of the referenced resource. | [optional] 
**Desc** | Pointer to **[]string** | Description of the resource. | [optional] 
**FullName** | Pointer to **string** | Full name of the ability score. | [optional] 
**Skills** | Pointer to [**[]APIReference**](APIReference.md) | List of skills that use this ability score. | [optional] 

## Methods

### NewAbilityScore

`func NewAbilityScore() *AbilityScore`

NewAbilityScore instantiates a new AbilityScore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbilityScoreWithDefaults

`func NewAbilityScoreWithDefaults() *AbilityScore`

NewAbilityScoreWithDefaults instantiates a new AbilityScore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *AbilityScore) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *AbilityScore) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *AbilityScore) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *AbilityScore) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetName

`func (o *AbilityScore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AbilityScore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AbilityScore) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AbilityScore) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *AbilityScore) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AbilityScore) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AbilityScore) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AbilityScore) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDesc

`func (o *AbilityScore) GetDesc() []string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *AbilityScore) GetDescOk() (*[]string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *AbilityScore) SetDesc(v []string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *AbilityScore) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetFullName

`func (o *AbilityScore) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *AbilityScore) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *AbilityScore) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *AbilityScore) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetSkills

`func (o *AbilityScore) GetSkills() []APIReference`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *AbilityScore) GetSkillsOk() (*[]APIReference, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *AbilityScore) SetSkills(v []APIReference)`

SetSkills sets Skills field to given value.

### HasSkills

`func (o *AbilityScore) HasSkills() bool`

HasSkills returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


