# Condition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **string** | Resource index for shorthand searching. | [optional] 
**Name** | Pointer to **string** | Name of the referenced resource. | [optional] 
**Url** | Pointer to **string** | URL of the referenced resource. | [optional] 
**Desc** | Pointer to **[]string** | Description of the resource. | [optional] 

## Methods

### NewCondition

`func NewCondition() *Condition`

NewCondition instantiates a new Condition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionWithDefaults

`func NewConditionWithDefaults() *Condition`

NewConditionWithDefaults instantiates a new Condition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *Condition) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *Condition) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *Condition) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *Condition) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetName

`func (o *Condition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Condition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Condition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Condition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *Condition) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Condition) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Condition) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Condition) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDesc

`func (o *Condition) GetDesc() []string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *Condition) GetDescOk() (*[]string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *Condition) SetDesc(v []string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *Condition) HasDesc() bool`

HasDesc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


