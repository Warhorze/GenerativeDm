# MonsterSense


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**passive_perception** | **float** | The monster&#39;s passive perception (wisdom) score. | [optional] 
**blindsight** | **str** | A monster with blindsight can perceive its surroundings without relying on sight, within a specific radius. | [optional] 
**darkvision** | **str** | A monster with darkvision can see in the dark within a specific radius. | [optional] 
**tremorsense** | **str** | A monster with tremorsense can detect and pinpoint the origin of vibrations within a specific radius, provided that the monster and the source of the vibrations are in contact with the same ground or substance. | [optional] 
**truesight** | **str** | A monster with truesight can, out to a specific range, see in normal and magical darkness, see invisible creatures and objects, automatically detect visual illusions and succeed on saving throws against them, and perceive the original form of a shapechanger or a creature that is transformed by magic. Furthermore, the monster can see into the Ethereal Plane within the same range. | [optional] 

## Example

```python
from openapi_client.models.monster_sense import MonsterSense

# TODO update the JSON string below
json = "{}"
# create an instance of MonsterSense from a JSON string
monster_sense_instance = MonsterSense.from_json(json)
# print the JSON string representation of the object
print MonsterSense.to_json()

# convert the object into a dict
monster_sense_dict = monster_sense_instance.to_dict()
# create an instance of MonsterSense from a dict
monster_sense_form_dict = monster_sense.from_dict(monster_sense_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


