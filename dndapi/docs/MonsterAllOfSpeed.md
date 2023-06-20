# MonsterAllOfSpeed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Walk** | Pointer to **string** | All creatures have a walking speed, simply called the monster’s speed. Creatures that have no form of ground-based locomotion have a walking speed of 0 feet. | [optional] 
**Burrow** | Pointer to **string** | A monster that has a burrowing speed can use that speed to move through sand, earth, mud, or ice. A monster can’t burrow through solid rock unless it has a special trait that allows it to do so. | [optional] 
**Climb** | Pointer to **string** | A monster that has a climbing speed can use all or part of its movement to move on vertical surfaces. The monster doesn’t need to spend extra movement to climb. | [optional] 
**Fly** | Pointer to **string** | A monster that has a flying speed can use all or part of its movement to fly. | [optional] 
**Swim** | Pointer to **string** | A monster that has a swimming speed doesn’t need to spend extra movement to swim. | [optional] 

## Methods

### NewMonsterAllOfSpeed

`func NewMonsterAllOfSpeed() *MonsterAllOfSpeed`

NewMonsterAllOfSpeed instantiates a new MonsterAllOfSpeed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonsterAllOfSpeedWithDefaults

`func NewMonsterAllOfSpeedWithDefaults() *MonsterAllOfSpeed`

NewMonsterAllOfSpeedWithDefaults instantiates a new MonsterAllOfSpeed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalk

`func (o *MonsterAllOfSpeed) GetWalk() string`

GetWalk returns the Walk field if non-nil, zero value otherwise.

### GetWalkOk

`func (o *MonsterAllOfSpeed) GetWalkOk() (*string, bool)`

GetWalkOk returns a tuple with the Walk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalk

`func (o *MonsterAllOfSpeed) SetWalk(v string)`

SetWalk sets Walk field to given value.

### HasWalk

`func (o *MonsterAllOfSpeed) HasWalk() bool`

HasWalk returns a boolean if a field has been set.

### GetBurrow

`func (o *MonsterAllOfSpeed) GetBurrow() string`

GetBurrow returns the Burrow field if non-nil, zero value otherwise.

### GetBurrowOk

`func (o *MonsterAllOfSpeed) GetBurrowOk() (*string, bool)`

GetBurrowOk returns a tuple with the Burrow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurrow

`func (o *MonsterAllOfSpeed) SetBurrow(v string)`

SetBurrow sets Burrow field to given value.

### HasBurrow

`func (o *MonsterAllOfSpeed) HasBurrow() bool`

HasBurrow returns a boolean if a field has been set.

### GetClimb

`func (o *MonsterAllOfSpeed) GetClimb() string`

GetClimb returns the Climb field if non-nil, zero value otherwise.

### GetClimbOk

`func (o *MonsterAllOfSpeed) GetClimbOk() (*string, bool)`

GetClimbOk returns a tuple with the Climb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClimb

`func (o *MonsterAllOfSpeed) SetClimb(v string)`

SetClimb sets Climb field to given value.

### HasClimb

`func (o *MonsterAllOfSpeed) HasClimb() bool`

HasClimb returns a boolean if a field has been set.

### GetFly

`func (o *MonsterAllOfSpeed) GetFly() string`

GetFly returns the Fly field if non-nil, zero value otherwise.

### GetFlyOk

`func (o *MonsterAllOfSpeed) GetFlyOk() (*string, bool)`

GetFlyOk returns a tuple with the Fly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFly

`func (o *MonsterAllOfSpeed) SetFly(v string)`

SetFly sets Fly field to given value.

### HasFly

`func (o *MonsterAllOfSpeed) HasFly() bool`

HasFly returns a boolean if a field has been set.

### GetSwim

`func (o *MonsterAllOfSpeed) GetSwim() string`

GetSwim returns the Swim field if non-nil, zero value otherwise.

### GetSwimOk

`func (o *MonsterAllOfSpeed) GetSwimOk() (*string, bool)`

GetSwimOk returns a tuple with the Swim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwim

`func (o *MonsterAllOfSpeed) SetSwim(v string)`

SetSwim sets Swim field to given value.

### HasSwim

`func (o *MonsterAllOfSpeed) HasSwim() bool`

HasSwim returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


