# OptionOneOf9

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OptionType** | Pointer to **string** | Type of option; determines other attributes. | [optional] 
**Name** | Pointer to **string** | Name of the breath | [optional] 
**Dc** | Pointer to [**DC**](DC.md) |  | [optional] 
**Damage** | Pointer to [**[]Damage**](Damage.md) | Damage dealt by the breath attack, if any. | [optional] 

## Methods

### NewOptionOneOf9

`func NewOptionOneOf9() *OptionOneOf9`

NewOptionOneOf9 instantiates a new OptionOneOf9 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionOneOf9WithDefaults

`func NewOptionOneOf9WithDefaults() *OptionOneOf9`

NewOptionOneOf9WithDefaults instantiates a new OptionOneOf9 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptionType

`func (o *OptionOneOf9) GetOptionType() string`

GetOptionType returns the OptionType field if non-nil, zero value otherwise.

### GetOptionTypeOk

`func (o *OptionOneOf9) GetOptionTypeOk() (*string, bool)`

GetOptionTypeOk returns a tuple with the OptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionType

`func (o *OptionOneOf9) SetOptionType(v string)`

SetOptionType sets OptionType field to given value.

### HasOptionType

`func (o *OptionOneOf9) HasOptionType() bool`

HasOptionType returns a boolean if a field has been set.

### GetName

`func (o *OptionOneOf9) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OptionOneOf9) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OptionOneOf9) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OptionOneOf9) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDc

`func (o *OptionOneOf9) GetDc() DC`

GetDc returns the Dc field if non-nil, zero value otherwise.

### GetDcOk

`func (o *OptionOneOf9) GetDcOk() (*DC, bool)`

GetDcOk returns a tuple with the Dc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDc

`func (o *OptionOneOf9) SetDc(v DC)`

SetDc sets Dc field to given value.

### HasDc

`func (o *OptionOneOf9) HasDc() bool`

HasDc returns a boolean if a field has been set.

### GetDamage

`func (o *OptionOneOf9) GetDamage() []Damage`

GetDamage returns the Damage field if non-nil, zero value otherwise.

### GetDamageOk

`func (o *OptionOneOf9) GetDamageOk() (*[]Damage, bool)`

GetDamageOk returns a tuple with the Damage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamage

`func (o *OptionOneOf9) SetDamage(v []Damage)`

SetDamage sets Damage field to given value.

### HasDamage

`func (o *OptionOneOf9) HasDamage() bool`

HasDamage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


