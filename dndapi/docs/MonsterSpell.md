# MonsterSpell

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Level** | Pointer to **float32** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Usage** | Pointer to [**MonsterUsage**](MonsterUsage.md) |  | [optional] 

## Methods

### NewMonsterSpell

`func NewMonsterSpell() *MonsterSpell`

NewMonsterSpell instantiates a new MonsterSpell object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonsterSpellWithDefaults

`func NewMonsterSpellWithDefaults() *MonsterSpell`

NewMonsterSpellWithDefaults instantiates a new MonsterSpell object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MonsterSpell) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MonsterSpell) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MonsterSpell) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MonsterSpell) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLevel

`func (o *MonsterSpell) GetLevel() float32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *MonsterSpell) GetLevelOk() (*float32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *MonsterSpell) SetLevel(v float32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *MonsterSpell) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetUrl

`func (o *MonsterSpell) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MonsterSpell) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MonsterSpell) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *MonsterSpell) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsage

`func (o *MonsterSpell) GetUsage() MonsterUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *MonsterSpell) GetUsageOk() (*MonsterUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *MonsterSpell) SetUsage(v MonsterUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *MonsterSpell) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


