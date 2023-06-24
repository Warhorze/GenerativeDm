# MonsterArmorClassOneOf2


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**type** | **str** |  | [optional] 
**value** | **float** |  | [optional] 
**armor** | [**List[APIReference]**](APIReference.md) |  | [optional] 
**desc** | **str** |  | [optional] 

## Example

```python
from dnd5epy.models.monster_armor_class_one_of2 import MonsterArmorClassOneOf2

# TODO update the JSON string below
json = "{}"
# create an instance of MonsterArmorClassOneOf2 from a JSON string
monster_armor_class_one_of2_instance = MonsterArmorClassOneOf2.from_json(json)
# print the JSON string representation of the object
print MonsterArmorClassOneOf2.to_json()

# convert the object into a dict
monster_armor_class_one_of2_dict = monster_armor_class_one_of2_instance.to_dict()
# create an instance of MonsterArmorClassOneOf2 from a dict
monster_armor_class_one_of2_form_dict = monster_armor_class_one_of2.from_dict(monster_armor_class_one_of2_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


