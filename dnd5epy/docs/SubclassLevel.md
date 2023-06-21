# SubclassLevel

`SubclassLevel` 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **str** | Resource index for shorthand searching. | [optional] 
**url** | **str** | URL of the referenced resource. | [optional] 
**level** | **float** | Number value for the current level object. | [optional] 
**ability_score_bonuses** | **float** | Total number of ability score bonuses gained, added from previous levels. | [optional] 
**prof_bonus** | **float** | Proficiency bonus for this class at the specified level. | [optional] 
**features** | [**List[APIReference]**](APIReference.md) | List of features gained at this level. | [optional] 
**spellcasting** | [**SubclassLevelSpellcasting**](SubclassLevelSpellcasting.md) |  | [optional] 
**classspecific** | **Dict[str, object]** | Class specific information such as dice values for bard songs and number of warlock invocations. | [optional] 

## Example

```python
from dnd5epy.models.subclass_level import SubclassLevel

# TODO update the JSON string below
json = "{}"
# create an instance of SubclassLevel from a JSON string
subclass_level_instance = SubclassLevel.from_json(json)
# print the JSON string representation of the object
print SubclassLevel.to_json()

# convert the object into a dict
subclass_level_dict = subclass_level_instance.to_dict()
# create an instance of SubclassLevel from a dict
subclass_level_form_dict = subclass_level.from_dict(subclass_level_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


