# WeaponProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **string** | Resource index for shorthand searching. | [optional] 
**Name** | Pointer to **string** | Name of the referenced resource. | [optional] 
**Url** | Pointer to **string** | URL of the referenced resource. | [optional] 
**Desc** | Pointer to **[]string** | Description of the resource. | [optional] 

## Methods

### NewWeaponProperty

`func NewWeaponProperty() *WeaponProperty`

NewWeaponProperty instantiates a new WeaponProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWeaponPropertyWithDefaults

`func NewWeaponPropertyWithDefaults() *WeaponProperty`

NewWeaponPropertyWithDefaults instantiates a new WeaponProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *WeaponProperty) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *WeaponProperty) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *WeaponProperty) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *WeaponProperty) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetName

`func (o *WeaponProperty) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WeaponProperty) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WeaponProperty) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WeaponProperty) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *WeaponProperty) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WeaponProperty) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WeaponProperty) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WeaponProperty) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDesc

`func (o *WeaponProperty) GetDesc() []string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *WeaponProperty) GetDescOk() (*[]string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *WeaponProperty) SetDesc(v []string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *WeaponProperty) HasDesc() bool`

HasDesc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


