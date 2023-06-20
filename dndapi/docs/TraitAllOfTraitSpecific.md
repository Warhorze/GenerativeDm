# TraitAllOfTraitSpecific

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Desc** | Pointer to **string** | Description of the choice to be made. | [optional] 
**Choose** | Pointer to **float32** | Number of items to pick from the list. | [optional] 
**Type** | Pointer to **string** | Type of the resources to choose from. | [optional] 
**From** | Pointer to [**OptionSet**](OptionSet.md) |  | [optional] 
**DamageType** | Pointer to [**TraitAllOfTraitSpecificOneOfDamageType**](TraitAllOfTraitSpecificOneOfDamageType.md) |  | [optional] 
**BreathWeapon** | Pointer to [**TraitAllOfTraitSpecificOneOfBreathWeapon**](TraitAllOfTraitSpecificOneOfBreathWeapon.md) |  | [optional] 

## Methods

### NewTraitAllOfTraitSpecific

`func NewTraitAllOfTraitSpecific() *TraitAllOfTraitSpecific`

NewTraitAllOfTraitSpecific instantiates a new TraitAllOfTraitSpecific object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraitAllOfTraitSpecificWithDefaults

`func NewTraitAllOfTraitSpecificWithDefaults() *TraitAllOfTraitSpecific`

NewTraitAllOfTraitSpecificWithDefaults instantiates a new TraitAllOfTraitSpecific object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesc

`func (o *TraitAllOfTraitSpecific) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *TraitAllOfTraitSpecific) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *TraitAllOfTraitSpecific) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *TraitAllOfTraitSpecific) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetChoose

`func (o *TraitAllOfTraitSpecific) GetChoose() float32`

GetChoose returns the Choose field if non-nil, zero value otherwise.

### GetChooseOk

`func (o *TraitAllOfTraitSpecific) GetChooseOk() (*float32, bool)`

GetChooseOk returns a tuple with the Choose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoose

`func (o *TraitAllOfTraitSpecific) SetChoose(v float32)`

SetChoose sets Choose field to given value.

### HasChoose

`func (o *TraitAllOfTraitSpecific) HasChoose() bool`

HasChoose returns a boolean if a field has been set.

### GetType

`func (o *TraitAllOfTraitSpecific) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TraitAllOfTraitSpecific) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TraitAllOfTraitSpecific) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TraitAllOfTraitSpecific) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFrom

`func (o *TraitAllOfTraitSpecific) GetFrom() OptionSet`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *TraitAllOfTraitSpecific) GetFromOk() (*OptionSet, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *TraitAllOfTraitSpecific) SetFrom(v OptionSet)`

SetFrom sets From field to given value.

### HasFrom

`func (o *TraitAllOfTraitSpecific) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetDamageType

`func (o *TraitAllOfTraitSpecific) GetDamageType() TraitAllOfTraitSpecificOneOfDamageType`

GetDamageType returns the DamageType field if non-nil, zero value otherwise.

### GetDamageTypeOk

`func (o *TraitAllOfTraitSpecific) GetDamageTypeOk() (*TraitAllOfTraitSpecificOneOfDamageType, bool)`

GetDamageTypeOk returns a tuple with the DamageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamageType

`func (o *TraitAllOfTraitSpecific) SetDamageType(v TraitAllOfTraitSpecificOneOfDamageType)`

SetDamageType sets DamageType field to given value.

### HasDamageType

`func (o *TraitAllOfTraitSpecific) HasDamageType() bool`

HasDamageType returns a boolean if a field has been set.

### GetBreathWeapon

`func (o *TraitAllOfTraitSpecific) GetBreathWeapon() TraitAllOfTraitSpecificOneOfBreathWeapon`

GetBreathWeapon returns the BreathWeapon field if non-nil, zero value otherwise.

### GetBreathWeaponOk

`func (o *TraitAllOfTraitSpecific) GetBreathWeaponOk() (*TraitAllOfTraitSpecificOneOfBreathWeapon, bool)`

GetBreathWeaponOk returns a tuple with the BreathWeapon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreathWeapon

`func (o *TraitAllOfTraitSpecific) SetBreathWeapon(v TraitAllOfTraitSpecificOneOfBreathWeapon)`

SetBreathWeapon sets BreathWeapon field to given value.

### HasBreathWeapon

`func (o *TraitAllOfTraitSpecific) HasBreathWeapon() bool`

HasBreathWeapon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


