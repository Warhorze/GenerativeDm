# Background

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **string** | Resource index for shorthand searching. | [optional] 
**Name** | Pointer to **string** | Name of the referenced resource. | [optional] 
**Url** | Pointer to **string** | URL of the referenced resource. | [optional] 
**StartingProficiencies** | Pointer to [**[]APIReference**](APIReference.md) | Starting proficiencies for all new characters of this background. | [optional] 
**StartingEquipment** | Pointer to [**[]APIReference**](APIReference.md) | Starting equipment for all new characters of this background. | [optional] 
**StartingEquipmentOptions** | Pointer to [**Choice**](Choice.md) |  | [optional] 
**LanguageOptions** | Pointer to [**Choice**](Choice.md) |  | [optional] 
**Feature** | Pointer to [**BackgroundAllOfFeature**](BackgroundAllOfFeature.md) |  | [optional] 
**PersonalityTraits** | Pointer to **map[string]interface{}** | Choice of personality traits for this background. | [optional] 
**Ideals** | Pointer to [**Choice**](Choice.md) |  | [optional] 
**Bonds** | Pointer to [**Choice**](Choice.md) |  | [optional] 
**Flaws** | Pointer to [**Choice**](Choice.md) |  | [optional] 

## Methods

### NewBackground

`func NewBackground() *Background`

NewBackground instantiates a new Background object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackgroundWithDefaults

`func NewBackgroundWithDefaults() *Background`

NewBackgroundWithDefaults instantiates a new Background object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *Background) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *Background) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *Background) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *Background) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetName

`func (o *Background) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Background) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Background) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Background) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *Background) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Background) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Background) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Background) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetStartingProficiencies

`func (o *Background) GetStartingProficiencies() []APIReference`

GetStartingProficiencies returns the StartingProficiencies field if non-nil, zero value otherwise.

### GetStartingProficienciesOk

`func (o *Background) GetStartingProficienciesOk() (*[]APIReference, bool)`

GetStartingProficienciesOk returns a tuple with the StartingProficiencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingProficiencies

`func (o *Background) SetStartingProficiencies(v []APIReference)`

SetStartingProficiencies sets StartingProficiencies field to given value.

### HasStartingProficiencies

`func (o *Background) HasStartingProficiencies() bool`

HasStartingProficiencies returns a boolean if a field has been set.

### GetStartingEquipment

`func (o *Background) GetStartingEquipment() []APIReference`

GetStartingEquipment returns the StartingEquipment field if non-nil, zero value otherwise.

### GetStartingEquipmentOk

`func (o *Background) GetStartingEquipmentOk() (*[]APIReference, bool)`

GetStartingEquipmentOk returns a tuple with the StartingEquipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingEquipment

`func (o *Background) SetStartingEquipment(v []APIReference)`

SetStartingEquipment sets StartingEquipment field to given value.

### HasStartingEquipment

`func (o *Background) HasStartingEquipment() bool`

HasStartingEquipment returns a boolean if a field has been set.

### GetStartingEquipmentOptions

`func (o *Background) GetStartingEquipmentOptions() Choice`

GetStartingEquipmentOptions returns the StartingEquipmentOptions field if non-nil, zero value otherwise.

### GetStartingEquipmentOptionsOk

`func (o *Background) GetStartingEquipmentOptionsOk() (*Choice, bool)`

GetStartingEquipmentOptionsOk returns a tuple with the StartingEquipmentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingEquipmentOptions

`func (o *Background) SetStartingEquipmentOptions(v Choice)`

SetStartingEquipmentOptions sets StartingEquipmentOptions field to given value.

### HasStartingEquipmentOptions

`func (o *Background) HasStartingEquipmentOptions() bool`

HasStartingEquipmentOptions returns a boolean if a field has been set.

### GetLanguageOptions

`func (o *Background) GetLanguageOptions() Choice`

GetLanguageOptions returns the LanguageOptions field if non-nil, zero value otherwise.

### GetLanguageOptionsOk

`func (o *Background) GetLanguageOptionsOk() (*Choice, bool)`

GetLanguageOptionsOk returns a tuple with the LanguageOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageOptions

`func (o *Background) SetLanguageOptions(v Choice)`

SetLanguageOptions sets LanguageOptions field to given value.

### HasLanguageOptions

`func (o *Background) HasLanguageOptions() bool`

HasLanguageOptions returns a boolean if a field has been set.

### GetFeature

`func (o *Background) GetFeature() BackgroundAllOfFeature`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *Background) GetFeatureOk() (*BackgroundAllOfFeature, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *Background) SetFeature(v BackgroundAllOfFeature)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *Background) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### GetPersonalityTraits

`func (o *Background) GetPersonalityTraits() map[string]interface{}`

GetPersonalityTraits returns the PersonalityTraits field if non-nil, zero value otherwise.

### GetPersonalityTraitsOk

`func (o *Background) GetPersonalityTraitsOk() (*map[string]interface{}, bool)`

GetPersonalityTraitsOk returns a tuple with the PersonalityTraits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalityTraits

`func (o *Background) SetPersonalityTraits(v map[string]interface{})`

SetPersonalityTraits sets PersonalityTraits field to given value.

### HasPersonalityTraits

`func (o *Background) HasPersonalityTraits() bool`

HasPersonalityTraits returns a boolean if a field has been set.

### GetIdeals

`func (o *Background) GetIdeals() Choice`

GetIdeals returns the Ideals field if non-nil, zero value otherwise.

### GetIdealsOk

`func (o *Background) GetIdealsOk() (*Choice, bool)`

GetIdealsOk returns a tuple with the Ideals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdeals

`func (o *Background) SetIdeals(v Choice)`

SetIdeals sets Ideals field to given value.

### HasIdeals

`func (o *Background) HasIdeals() bool`

HasIdeals returns a boolean if a field has been set.

### GetBonds

`func (o *Background) GetBonds() Choice`

GetBonds returns the Bonds field if non-nil, zero value otherwise.

### GetBondsOk

`func (o *Background) GetBondsOk() (*Choice, bool)`

GetBondsOk returns a tuple with the Bonds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonds

`func (o *Background) SetBonds(v Choice)`

SetBonds sets Bonds field to given value.

### HasBonds

`func (o *Background) HasBonds() bool`

HasBonds returns a boolean if a field has been set.

### GetFlaws

`func (o *Background) GetFlaws() Choice`

GetFlaws returns the Flaws field if non-nil, zero value otherwise.

### GetFlawsOk

`func (o *Background) GetFlawsOk() (*Choice, bool)`

GetFlawsOk returns a tuple with the Flaws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlaws

`func (o *Background) SetFlaws(v Choice)`

SetFlaws sets Flaws field to given value.

### HasFlaws

`func (o *Background) HasFlaws() bool`

HasFlaws returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


