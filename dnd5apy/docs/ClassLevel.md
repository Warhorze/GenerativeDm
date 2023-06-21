# ClassLevel

`ClassLevel` 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **str** | Resource index for shorthand searching. | [optional] 
**url** | **str** | URL of the referenced resource. | [optional] 
**level** | **float** | The number value for the current level object. | [optional] 
**ability_score_bonuses** | **float** | Total number of ability score bonuses gained, added from previous levels. | [optional] 
**prof_bonus** | **float** | Proficiency bonus for this class at the specified level. | [optional] 
**features** | [**List[APIReference]**](APIReference.md) | Features automatically gained at this level. | [optional] 
**spellcasting** | [**SubclassLevelSpellcasting**](SubclassLevelSpellcasting.md) |  | [optional] 
**class_specific** | [**ClassLevelClassSpecific**](ClassLevelClassSpecific.md) |  | [optional] 

## Example

```python
from openapi_client.models.class_level import ClassLevel

# TODO update the JSON string below
json = "{}"
# create an instance of ClassLevel from a JSON string
class_level_instance = ClassLevel.from_json(json)
# print the JSON string representation of the object
print ClassLevel.to_json()

# convert the object into a dict
class_level_dict = class_level_instance.to_dict()
# create an instance of ClassLevel from a dict
class_level_form_dict = class_level.from_dict(class_level_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


