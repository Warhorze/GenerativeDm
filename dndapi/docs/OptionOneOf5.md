# OptionOneOf5

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OptionType** | Pointer to **string** | Type of option; determines other attributes. | [optional] 
**Desc** | Pointer to **string** | A description of the ideal. | [optional] 
**Alignments** | Pointer to [**[]APIReference**](APIReference.md) | A list of alignments of those who might follow the ideal. | [optional] 

## Methods

### NewOptionOneOf5

`func NewOptionOneOf5() *OptionOneOf5`

NewOptionOneOf5 instantiates a new OptionOneOf5 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionOneOf5WithDefaults

`func NewOptionOneOf5WithDefaults() *OptionOneOf5`

NewOptionOneOf5WithDefaults instantiates a new OptionOneOf5 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptionType

`func (o *OptionOneOf5) GetOptionType() string`

GetOptionType returns the OptionType field if non-nil, zero value otherwise.

### GetOptionTypeOk

`func (o *OptionOneOf5) GetOptionTypeOk() (*string, bool)`

GetOptionTypeOk returns a tuple with the OptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionType

`func (o *OptionOneOf5) SetOptionType(v string)`

SetOptionType sets OptionType field to given value.

### HasOptionType

`func (o *OptionOneOf5) HasOptionType() bool`

HasOptionType returns a boolean if a field has been set.

### GetDesc

`func (o *OptionOneOf5) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *OptionOneOf5) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *OptionOneOf5) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *OptionOneOf5) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetAlignments

`func (o *OptionOneOf5) GetAlignments() []APIReference`

GetAlignments returns the Alignments field if non-nil, zero value otherwise.

### GetAlignmentsOk

`func (o *OptionOneOf5) GetAlignmentsOk() (*[]APIReference, bool)`

GetAlignmentsOk returns a tuple with the Alignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignments

`func (o *OptionOneOf5) SetAlignments(v []APIReference)`

SetAlignments sets Alignments field to given value.

### HasAlignments

`func (o *OptionOneOf5) HasAlignments() bool`

HasAlignments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


