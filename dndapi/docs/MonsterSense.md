# MonsterSense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PassivePerception** | Pointer to **float32** | The monster&#39;s passive perception (wisdom) score. | [optional] 
**Blindsight** | Pointer to **string** | A monster with blindsight can perceive its surroundings without relying on sight, within a specific radius. | [optional] 
**Darkvision** | Pointer to **string** | A monster with darkvision can see in the dark within a specific radius. | [optional] 
**Tremorsense** | Pointer to **string** | A monster with tremorsense can detect and pinpoint the origin of vibrations within a specific radius, provided that the monster and the source of the vibrations are in contact with the same ground or substance. | [optional] 
**Truesight** | Pointer to **string** | A monster with truesight can, out to a specific range, see in normal and magical darkness, see invisible creatures and objects, automatically detect visual illusions and succeed on saving throws against them, and perceive the original form of a shapechanger or a creature that is transformed by magic. Furthermore, the monster can see into the Ethereal Plane within the same range. | [optional] 

## Methods

### NewMonsterSense

`func NewMonsterSense() *MonsterSense`

NewMonsterSense instantiates a new MonsterSense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonsterSenseWithDefaults

`func NewMonsterSenseWithDefaults() *MonsterSense`

NewMonsterSenseWithDefaults instantiates a new MonsterSense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassivePerception

`func (o *MonsterSense) GetPassivePerception() float32`

GetPassivePerception returns the PassivePerception field if non-nil, zero value otherwise.

### GetPassivePerceptionOk

`func (o *MonsterSense) GetPassivePerceptionOk() (*float32, bool)`

GetPassivePerceptionOk returns a tuple with the PassivePerception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassivePerception

`func (o *MonsterSense) SetPassivePerception(v float32)`

SetPassivePerception sets PassivePerception field to given value.

### HasPassivePerception

`func (o *MonsterSense) HasPassivePerception() bool`

HasPassivePerception returns a boolean if a field has been set.

### GetBlindsight

`func (o *MonsterSense) GetBlindsight() string`

GetBlindsight returns the Blindsight field if non-nil, zero value otherwise.

### GetBlindsightOk

`func (o *MonsterSense) GetBlindsightOk() (*string, bool)`

GetBlindsightOk returns a tuple with the Blindsight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlindsight

`func (o *MonsterSense) SetBlindsight(v string)`

SetBlindsight sets Blindsight field to given value.

### HasBlindsight

`func (o *MonsterSense) HasBlindsight() bool`

HasBlindsight returns a boolean if a field has been set.

### GetDarkvision

`func (o *MonsterSense) GetDarkvision() string`

GetDarkvision returns the Darkvision field if non-nil, zero value otherwise.

### GetDarkvisionOk

`func (o *MonsterSense) GetDarkvisionOk() (*string, bool)`

GetDarkvisionOk returns a tuple with the Darkvision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkvision

`func (o *MonsterSense) SetDarkvision(v string)`

SetDarkvision sets Darkvision field to given value.

### HasDarkvision

`func (o *MonsterSense) HasDarkvision() bool`

HasDarkvision returns a boolean if a field has been set.

### GetTremorsense

`func (o *MonsterSense) GetTremorsense() string`

GetTremorsense returns the Tremorsense field if non-nil, zero value otherwise.

### GetTremorsenseOk

`func (o *MonsterSense) GetTremorsenseOk() (*string, bool)`

GetTremorsenseOk returns a tuple with the Tremorsense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTremorsense

`func (o *MonsterSense) SetTremorsense(v string)`

SetTremorsense sets Tremorsense field to given value.

### HasTremorsense

`func (o *MonsterSense) HasTremorsense() bool`

HasTremorsense returns a boolean if a field has been set.

### GetTruesight

`func (o *MonsterSense) GetTruesight() string`

GetTruesight returns the Truesight field if non-nil, zero value otherwise.

### GetTruesightOk

`func (o *MonsterSense) GetTruesightOk() (*string, bool)`

GetTruesightOk returns a tuple with the Truesight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruesight

`func (o *MonsterSense) SetTruesight(v string)`

SetTruesight sets Truesight field to given value.

### HasTruesight

`func (o *MonsterSense) HasTruesight() bool`

HasTruesight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


