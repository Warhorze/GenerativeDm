# MonsterArmorClassOneOf3


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**type** | **str** |  | [optional] 
**value** | **float** |  | [optional] 
**spell** | [**APIReference**](APIReference.md) |  | [optional] 
**desc** | **str** |  | [optional] 

## Example

```python
from dnd5epy.models.monster_armor_class_one_of3 import MonsterArmorClassOneOf3

# TODO update the JSON string below
json = "{}"
# create an instance of MonsterArmorClassOneOf3 from a JSON string
monster_armor_class_one_of3_instance = MonsterArmorClassOneOf3.from_json(json)
# print the JSON string representation of the object
print MonsterArmorClassOneOf3.to_json()

# convert the object into a dict
monster_armor_class_one_of3_dict = monster_armor_class_one_of3_instance.to_dict()
# create an instance of MonsterArmorClassOneOf3 from a dict
monster_armor_class_one_of3_form_dict = monster_armor_class_one_of3.from_dict(monster_armor_class_one_of3_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


