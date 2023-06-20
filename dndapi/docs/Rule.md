# Rule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **string** | Resource index for shorthand searching. | [optional] 
**Name** | Pointer to **string** | Name of the referenced resource. | [optional] 
**Url** | Pointer to **string** | URL of the referenced resource. | [optional] 
**Desc** | Pointer to **string** | Description of the rule. | [optional] 
**Subsections** | Pointer to [**[]APIReference**](APIReference.md) | List of sections for each subheading underneath the rule in the SRD. | [optional] 

## Methods

### NewRule

`func NewRule() *Rule`

NewRule instantiates a new Rule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleWithDefaults

`func NewRuleWithDefaults() *Rule`

NewRuleWithDefaults instantiates a new Rule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *Rule) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *Rule) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *Rule) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *Rule) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetName

`func (o *Rule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Rule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Rule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Rule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *Rule) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Rule) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Rule) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Rule) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDesc

`func (o *Rule) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *Rule) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *Rule) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *Rule) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetSubsections

`func (o *Rule) GetSubsections() []APIReference`

GetSubsections returns the Subsections field if non-nil, zero value otherwise.

### GetSubsectionsOk

`func (o *Rule) GetSubsectionsOk() (*[]APIReference, bool)`

GetSubsectionsOk returns a tuple with the Subsections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsections

`func (o *Rule) SetSubsections(v []APIReference)`

SetSubsections sets Subsections field to given value.

### HasSubsections

`func (o *Rule) HasSubsections() bool`

HasSubsections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


