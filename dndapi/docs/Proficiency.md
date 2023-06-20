# Proficiency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **string** | Resource index for shorthand searching. | [optional] 
**Name** | Pointer to **string** | Name of the referenced resource. | [optional] 
**Url** | Pointer to **string** | URL of the referenced resource. | [optional] 
**Type** | Pointer to **string** | The general category of the proficiency. | [optional] 
**Classes** | Pointer to [**[]APIReference**](APIReference.md) | Classes that start with this proficiency. | [optional] 
**Races** | Pointer to [**[]APIReference**](APIReference.md) | Races that start with this proficiency. | [optional] 
**Reference** | Pointer to [**ProficiencyAllOfReference**](ProficiencyAllOfReference.md) |  | [optional] 

## Methods

### NewProficiency

`func NewProficiency() *Proficiency`

NewProficiency instantiates a new Proficiency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProficiencyWithDefaults

`func NewProficiencyWithDefaults() *Proficiency`

NewProficiencyWithDefaults instantiates a new Proficiency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *Proficiency) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *Proficiency) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *Proficiency) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *Proficiency) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetName

`func (o *Proficiency) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Proficiency) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Proficiency) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Proficiency) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *Proficiency) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Proficiency) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Proficiency) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Proficiency) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetType

`func (o *Proficiency) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Proficiency) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Proficiency) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Proficiency) HasType() bool`

HasType returns a boolean if a field has been set.

### GetClasses

`func (o *Proficiency) GetClasses() []APIReference`

GetClasses returns the Classes field if non-nil, zero value otherwise.

### GetClassesOk

`func (o *Proficiency) GetClassesOk() (*[]APIReference, bool)`

GetClassesOk returns a tuple with the Classes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClasses

`func (o *Proficiency) SetClasses(v []APIReference)`

SetClasses sets Classes field to given value.

### HasClasses

`func (o *Proficiency) HasClasses() bool`

HasClasses returns a boolean if a field has been set.

### GetRaces

`func (o *Proficiency) GetRaces() []APIReference`

GetRaces returns the Races field if non-nil, zero value otherwise.

### GetRacesOk

`func (o *Proficiency) GetRacesOk() (*[]APIReference, bool)`

GetRacesOk returns a tuple with the Races field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaces

`func (o *Proficiency) SetRaces(v []APIReference)`

SetRaces sets Races field to given value.

### HasRaces

`func (o *Proficiency) HasRaces() bool`

HasRaces returns a boolean if a field has been set.

### GetReference

`func (o *Proficiency) GetReference() ProficiencyAllOfReference`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *Proficiency) GetReferenceOk() (*ProficiencyAllOfReference, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *Proficiency) SetReference(v ProficiencyAllOfReference)`

SetReference sets Reference field to given value.

### HasReference

`func (o *Proficiency) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


