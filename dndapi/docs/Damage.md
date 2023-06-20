# Damage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DamageDice** | Pointer to **string** |  | [optional] 
**DamageType** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 

## Methods

### NewDamage

`func NewDamage() *Damage`

NewDamage instantiates a new Damage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDamageWithDefaults

`func NewDamageWithDefaults() *Damage`

NewDamageWithDefaults instantiates a new Damage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDamageDice

`func (o *Damage) GetDamageDice() string`

GetDamageDice returns the DamageDice field if non-nil, zero value otherwise.

### GetDamageDiceOk

`func (o *Damage) GetDamageDiceOk() (*string, bool)`

GetDamageDiceOk returns a tuple with the DamageDice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamageDice

`func (o *Damage) SetDamageDice(v string)`

SetDamageDice sets DamageDice field to given value.

### HasDamageDice

`func (o *Damage) HasDamageDice() bool`

HasDamageDice returns a boolean if a field has been set.

### GetDamageType

`func (o *Damage) GetDamageType() APIReference`

GetDamageType returns the DamageType field if non-nil, zero value otherwise.

### GetDamageTypeOk

`func (o *Damage) GetDamageTypeOk() (*APIReference, bool)`

GetDamageTypeOk returns a tuple with the DamageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamageType

`func (o *Damage) SetDamageType(v APIReference)`

SetDamageType sets DamageType field to given value.

### HasDamageType

`func (o *Damage) HasDamageType() bool`

HasDamageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


