# ClassLevelClassSpecificAnyOf

Barbarian Class Specific Features

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**rage_count** | **float** |  | [optional] 
**rage_damage_bonus** | **float** |  | [optional] 
**brutal_critical_dice** | **float** |  | [optional] 

## Example

```python
from dnd5epy.models.class_level_class_specific_any_of import ClassLevelClassSpecificAnyOf

# TODO update the JSON string below
json = "{}"
# create an instance of ClassLevelClassSpecificAnyOf from a JSON string
class_level_class_specific_any_of_instance = ClassLevelClassSpecificAnyOf.from_json(json)
# print the JSON string representation of the object
print ClassLevelClassSpecificAnyOf.to_json()

# convert the object into a dict
class_level_class_specific_any_of_dict = class_level_class_specific_any_of_instance.to_dict()
# create an instance of ClassLevelClassSpecificAnyOf from a dict
class_level_class_specific_any_of_form_dict = class_level_class_specific_any_of.from_dict(class_level_class_specific_any_of_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


