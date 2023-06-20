# Multiclassing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prerequisites** | Pointer to [**[]Prerequisite**](Prerequisite.md) | List of prerequisites that must be met. | [optional] 
**PrerequisiteOptions** | Pointer to [**[]Choice**](Choice.md) | List of choices of prerequisites to meet for. | [optional] 
**Proficiencies** | Pointer to [**[]APIReference**](APIReference.md) | List of proficiencies available when multiclassing. | [optional] 
**ProficiencyChoices** | Pointer to [**[]Choice**](Choice.md) | List of choices of proficiencies that are given when multiclassing. | [optional] 

## Methods

### NewMulticlassing

`func NewMulticlassing() *Multiclassing`

NewMulticlassing instantiates a new Multiclassing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMulticlassingWithDefaults

`func NewMulticlassingWithDefaults() *Multiclassing`

NewMulticlassingWithDefaults instantiates a new Multiclassing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrerequisites

`func (o *Multiclassing) GetPrerequisites() []Prerequisite`

GetPrerequisites returns the Prerequisites field if non-nil, zero value otherwise.

### GetPrerequisitesOk

`func (o *Multiclassing) GetPrerequisitesOk() (*[]Prerequisite, bool)`

GetPrerequisitesOk returns a tuple with the Prerequisites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrerequisites

`func (o *Multiclassing) SetPrerequisites(v []Prerequisite)`

SetPrerequisites sets Prerequisites field to given value.

### HasPrerequisites

`func (o *Multiclassing) HasPrerequisites() bool`

HasPrerequisites returns a boolean if a field has been set.

### GetPrerequisiteOptions

`func (o *Multiclassing) GetPrerequisiteOptions() []Choice`

GetPrerequisiteOptions returns the PrerequisiteOptions field if non-nil, zero value otherwise.

### GetPrerequisiteOptionsOk

`func (o *Multiclassing) GetPrerequisiteOptionsOk() (*[]Choice, bool)`

GetPrerequisiteOptionsOk returns a tuple with the PrerequisiteOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrerequisiteOptions

`func (o *Multiclassing) SetPrerequisiteOptions(v []Choice)`

SetPrerequisiteOptions sets PrerequisiteOptions field to given value.

### HasPrerequisiteOptions

`func (o *Multiclassing) HasPrerequisiteOptions() bool`

HasPrerequisiteOptions returns a boolean if a field has been set.

### GetProficiencies

`func (o *Multiclassing) GetProficiencies() []APIReference`

GetProficiencies returns the Proficiencies field if non-nil, zero value otherwise.

### GetProficienciesOk

`func (o *Multiclassing) GetProficienciesOk() (*[]APIReference, bool)`

GetProficienciesOk returns a tuple with the Proficiencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProficiencies

`func (o *Multiclassing) SetProficiencies(v []APIReference)`

SetProficiencies sets Proficiencies field to given value.

### HasProficiencies

`func (o *Multiclassing) HasProficiencies() bool`

HasProficiencies returns a boolean if a field has been set.

### GetProficiencyChoices

`func (o *Multiclassing) GetProficiencyChoices() []Choice`

GetProficiencyChoices returns the ProficiencyChoices field if non-nil, zero value otherwise.

### GetProficiencyChoicesOk

`func (o *Multiclassing) GetProficiencyChoicesOk() (*[]Choice, bool)`

GetProficiencyChoicesOk returns a tuple with the ProficiencyChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProficiencyChoices

`func (o *Multiclassing) SetProficiencyChoices(v []Choice)`

SetProficiencyChoices sets ProficiencyChoices field to given value.

### HasProficiencyChoices

`func (o *Multiclassing) HasProficiencyChoices() bool`

HasProficiencyChoices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


