# OptionOneOf10

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OptionType** | Pointer to **string** | Type of option; determines other attributes. | [optional] 
**DamageType** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 
**DamageDice** | Pointer to **string** | Damage expressed in dice (e.g. \&quot;13d6\&quot;). | [optional] 
**Notes** | Pointer to **string** | Information regarding the damage. | [optional] 

## Methods

### NewOptionOneOf10

`func NewOptionOneOf10() *OptionOneOf10`

NewOptionOneOf10 instantiates a new OptionOneOf10 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionOneOf10WithDefaults

`func NewOptionOneOf10WithDefaults() *OptionOneOf10`

NewOptionOneOf10WithDefaults instantiates a new OptionOneOf10 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptionType

`func (o *OptionOneOf10) GetOptionType() string`

GetOptionType returns the OptionType field if non-nil, zero value otherwise.

### GetOptionTypeOk

`func (o *OptionOneOf10) GetOptionTypeOk() (*string, bool)`

GetOptionTypeOk returns a tuple with the OptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionType

`func (o *OptionOneOf10) SetOptionType(v string)`

SetOptionType sets OptionType field to given value.

### HasOptionType

`func (o *OptionOneOf10) HasOptionType() bool`

HasOptionType returns a boolean if a field has been set.

### GetDamageType

`func (o *OptionOneOf10) GetDamageType() APIReference`

GetDamageType returns the DamageType field if non-nil, zero value otherwise.

### GetDamageTypeOk

`func (o *OptionOneOf10) GetDamageTypeOk() (*APIReference, bool)`

GetDamageTypeOk returns a tuple with the DamageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamageType

`func (o *OptionOneOf10) SetDamageType(v APIReference)`

SetDamageType sets DamageType field to given value.

### HasDamageType

`func (o *OptionOneOf10) HasDamageType() bool`

HasDamageType returns a boolean if a field has been set.

### GetDamageDice

`func (o *OptionOneOf10) GetDamageDice() string`

GetDamageDice returns the DamageDice field if non-nil, zero value otherwise.

### GetDamageDiceOk

`func (o *OptionOneOf10) GetDamageDiceOk() (*string, bool)`

GetDamageDiceOk returns a tuple with the DamageDice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamageDice

`func (o *OptionOneOf10) SetDamageDice(v string)`

SetDamageDice sets DamageDice field to given value.

### HasDamageDice

`func (o *OptionOneOf10) HasDamageDice() bool`

HasDamageDice returns a boolean if a field has been set.

### GetNotes

`func (o *OptionOneOf10) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *OptionOneOf10) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *OptionOneOf10) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *OptionOneOf10) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


