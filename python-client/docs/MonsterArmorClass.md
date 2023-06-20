# MonsterArmorClass

The armor class of a monster.

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**type** | **str** |  | [optional] 
**value** | **float** |  | [optional] 
**desc** | **str** |  | [optional] 
**armor** | [**List[APIReference]**](APIReference.md) |  | [optional] 
**spell** | [**APIReference**](APIReference.md) |  | [optional] 
**condition** | [**APIReference**](APIReference.md) |  | [optional] 

## Example

```python
from openapi_client.models.monster_armor_class import MonsterArmorClass

# TODO update the JSON string below
json = "{}"
# create an instance of MonsterArmorClass from a JSON string
monster_armor_class_instance = MonsterArmorClass.from_json(json)
# print the JSON string representation of the object
print MonsterArmorClass.to_json()

# convert the object into a dict
monster_armor_class_dict = monster_armor_class_instance.to_dict()
# create an instance of MonsterArmorClass from a dict
monster_armor_class_form_dict = monster_armor_class.from_dict(monster_armor_class_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


