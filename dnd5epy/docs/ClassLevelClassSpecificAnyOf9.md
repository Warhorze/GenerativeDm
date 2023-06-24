# ClassLevelClassSpecificAnyOf9

Bard Sorcerer Specific Features

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**sorcery_points** | **float** |  | [optional] 
**metamagic_known** | **float** |  | [optional] 
**creating_spell_slots** | [**List[ClassLevelClassSpecificAnyOf9CreatingSpellSlotsInner]**](ClassLevelClassSpecificAnyOf9CreatingSpellSlotsInner.md) |  | [optional] 

## Example

```python
from dnd5epy.models.class_level_class_specific_any_of9 import ClassLevelClassSpecificAnyOf9

# TODO update the JSON string below
json = "{}"
# create an instance of ClassLevelClassSpecificAnyOf9 from a JSON string
class_level_class_specific_any_of9_instance = ClassLevelClassSpecificAnyOf9.from_json(json)
# print the JSON string representation of the object
print ClassLevelClassSpecificAnyOf9.to_json()

# convert the object into a dict
class_level_class_specific_any_of9_dict = class_level_class_specific_any_of9_instance.to_dict()
# create an instance of ClassLevelClassSpecificAnyOf9 from a dict
class_level_class_specific_any_of9_form_dict = class_level_class_specific_any_of9.from_dict(class_level_class_specific_any_of9_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


