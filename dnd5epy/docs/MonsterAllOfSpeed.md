# MonsterAllOfSpeed

Speed for a monster determines how fast it can move per turn.

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**walk** | **str** | All creatures have a walking speed, simply called the monster’s speed. Creatures that have no form of ground-based locomotion have a walking speed of 0 feet. | [optional] 
**burrow** | **str** | A monster that has a burrowing speed can use that speed to move through sand, earth, mud, or ice. A monster can’t burrow through solid rock unless it has a special trait that allows it to do so. | [optional] 
**climb** | **str** | A monster that has a climbing speed can use all or part of its movement to move on vertical surfaces. The monster doesn’t need to spend extra movement to climb. | [optional] 
**fly** | **str** | A monster that has a flying speed can use all or part of its movement to fly. | [optional] 
**swim** | **str** | A monster that has a swimming speed doesn’t need to spend extra movement to swim. | [optional] 

## Example

```python
from openapi_client.models.monster_all_of_speed import MonsterAllOfSpeed

# TODO update the JSON string below
json = "{}"
# create an instance of MonsterAllOfSpeed from a JSON string
monster_all_of_speed_instance = MonsterAllOfSpeed.from_json(json)
# print the JSON string representation of the object
print MonsterAllOfSpeed.to_json()

# convert the object into a dict
monster_all_of_speed_dict = monster_all_of_speed_instance.to_dict()
# create an instance of MonsterAllOfSpeed from a dict
monster_all_of_speed_form_dict = monster_all_of_speed.from_dict(monster_all_of_speed_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


