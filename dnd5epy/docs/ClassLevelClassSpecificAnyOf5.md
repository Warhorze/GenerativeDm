# ClassLevelClassSpecificAnyOf5

Monk Class Specific Features

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ki_points** | **float** |  | [optional] 
**unarmored_movement** | **float** |  | [optional] 
**martial_arts** | [**ClassLevelClassSpecificAnyOf5MartialArts**](ClassLevelClassSpecificAnyOf5MartialArts.md) |  | [optional] 

## Example

```python
from dnd5epy.models.class_level_class_specific_any_of5 import ClassLevelClassSpecificAnyOf5

# TODO update the JSON string below
json = "{}"
# create an instance of ClassLevelClassSpecificAnyOf5 from a JSON string
class_level_class_specific_any_of5_instance = ClassLevelClassSpecificAnyOf5.from_json(json)
# print the JSON string representation of the object
print ClassLevelClassSpecificAnyOf5.to_json()

# convert the object into a dict
class_level_class_specific_any_of5_dict = class_level_class_specific_any_of5_instance.to_dict()
# create an instance of ClassLevelClassSpecificAnyOf5 from a dict
class_level_class_specific_any_of5_form_dict = class_level_class_specific_any_of5.from_dict(class_level_class_specific_any_of5_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


