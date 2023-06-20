# Feature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **string** | Resource index for shorthand searching. | [optional] 
**Name** | Pointer to **string** | Name of the referenced resource. | [optional] 
**Url** | Pointer to **string** | URL of the referenced resource. | [optional] 
**Desc** | Pointer to **[]string** | Description of the resource. | [optional] 
**Level** | Pointer to **float32** | The level this feature is gained. | [optional] 
**Class** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 
**Subclass** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 
**Parent** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 
**Prerequisites** | Pointer to [**[]FeatureAllOfPrerequisitesInner**](FeatureAllOfPrerequisitesInner.md) | The prerequisites for this feature. | [optional] 
**FeatureSpecific** | Pointer to **map[string]interface{}** | Information specific to this feature. | [optional] 

## Methods

### NewFeature

`func NewFeature() *Feature`

NewFeature instantiates a new Feature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureWithDefaults

`func NewFeatureWithDefaults() *Feature`

NewFeatureWithDefaults instantiates a new Feature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *Feature) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *Feature) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *Feature) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *Feature) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetName

`func (o *Feature) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Feature) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Feature) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Feature) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *Feature) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Feature) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Feature) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Feature) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDesc

`func (o *Feature) GetDesc() []string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *Feature) GetDescOk() (*[]string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *Feature) SetDesc(v []string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *Feature) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetLevel

`func (o *Feature) GetLevel() float32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *Feature) GetLevelOk() (*float32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *Feature) SetLevel(v float32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *Feature) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetClass

`func (o *Feature) GetClass() APIReference`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *Feature) GetClassOk() (*APIReference, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *Feature) SetClass(v APIReference)`

SetClass sets Class field to given value.

### HasClass

`func (o *Feature) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetSubclass

`func (o *Feature) GetSubclass() APIReference`

GetSubclass returns the Subclass field if non-nil, zero value otherwise.

### GetSubclassOk

`func (o *Feature) GetSubclassOk() (*APIReference, bool)`

GetSubclassOk returns a tuple with the Subclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubclass

`func (o *Feature) SetSubclass(v APIReference)`

SetSubclass sets Subclass field to given value.

### HasSubclass

`func (o *Feature) HasSubclass() bool`

HasSubclass returns a boolean if a field has been set.

### GetParent

`func (o *Feature) GetParent() APIReference`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *Feature) GetParentOk() (*APIReference, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *Feature) SetParent(v APIReference)`

SetParent sets Parent field to given value.

### HasParent

`func (o *Feature) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPrerequisites

`func (o *Feature) GetPrerequisites() []FeatureAllOfPrerequisitesInner`

GetPrerequisites returns the Prerequisites field if non-nil, zero value otherwise.

### GetPrerequisitesOk

`func (o *Feature) GetPrerequisitesOk() (*[]FeatureAllOfPrerequisitesInner, bool)`

GetPrerequisitesOk returns a tuple with the Prerequisites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrerequisites

`func (o *Feature) SetPrerequisites(v []FeatureAllOfPrerequisitesInner)`

SetPrerequisites sets Prerequisites field to given value.

### HasPrerequisites

`func (o *Feature) HasPrerequisites() bool`

HasPrerequisites returns a boolean if a field has been set.

### GetFeatureSpecific

`func (o *Feature) GetFeatureSpecific() map[string]interface{}`

GetFeatureSpecific returns the FeatureSpecific field if non-nil, zero value otherwise.

### GetFeatureSpecificOk

`func (o *Feature) GetFeatureSpecificOk() (*map[string]interface{}, bool)`

GetFeatureSpecificOk returns a tuple with the FeatureSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureSpecific

`func (o *Feature) SetFeatureSpecific(v map[string]interface{})`

SetFeatureSpecific sets FeatureSpecific field to given value.

### HasFeatureSpecific

`func (o *Feature) HasFeatureSpecific() bool`

HasFeatureSpecific returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


