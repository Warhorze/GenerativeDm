# OptionOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OptionType** | Pointer to **string** | Type of option; determines other attributes. | [optional] 
**ActionName** | Pointer to **string** | The name of the action. | [optional] 
**Count** | Pointer to **float32** | The number of times this action can be repeated if chosen. | [optional] 
**Type** | Pointer to **string** | For attack options that can be melee, ranged, abilities, or thrown. | [optional] 

## Methods

### NewOptionOneOf1

`func NewOptionOneOf1() *OptionOneOf1`

NewOptionOneOf1 instantiates a new OptionOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionOneOf1WithDefaults

`func NewOptionOneOf1WithDefaults() *OptionOneOf1`

NewOptionOneOf1WithDefaults instantiates a new OptionOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptionType

`func (o *OptionOneOf1) GetOptionType() string`

GetOptionType returns the OptionType field if non-nil, zero value otherwise.

### GetOptionTypeOk

`func (o *OptionOneOf1) GetOptionTypeOk() (*string, bool)`

GetOptionTypeOk returns a tuple with the OptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionType

`func (o *OptionOneOf1) SetOptionType(v string)`

SetOptionType sets OptionType field to given value.

### HasOptionType

`func (o *OptionOneOf1) HasOptionType() bool`

HasOptionType returns a boolean if a field has been set.

### GetActionName

`func (o *OptionOneOf1) GetActionName() string`

GetActionName returns the ActionName field if non-nil, zero value otherwise.

### GetActionNameOk

`func (o *OptionOneOf1) GetActionNameOk() (*string, bool)`

GetActionNameOk returns a tuple with the ActionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionName

`func (o *OptionOneOf1) SetActionName(v string)`

SetActionName sets ActionName field to given value.

### HasActionName

`func (o *OptionOneOf1) HasActionName() bool`

HasActionName returns a boolean if a field has been set.

### GetCount

`func (o *OptionOneOf1) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *OptionOneOf1) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *OptionOneOf1) SetCount(v float32)`

SetCount sets Count field to given value.

### HasCount

`func (o *OptionOneOf1) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetType

`func (o *OptionOneOf1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OptionOneOf1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OptionOneOf1) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OptionOneOf1) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


