# MonsterMultiAttackAction


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**action_name** | **str** |  | [optional] 
**count** | **float** |  | [optional] 
**type** | **str** |  | [optional] 

## Example

```python
from dnd5epy.models.monster_multi_attack_action import MonsterMultiAttackAction

# TODO update the JSON string below
json = "{}"
# create an instance of MonsterMultiAttackAction from a JSON string
monster_multi_attack_action_instance = MonsterMultiAttackAction.from_json(json)
# print the JSON string representation of the object
print MonsterMultiAttackAction.to_json()

# convert the object into a dict
monster_multi_attack_action_dict = monster_multi_attack_action_instance.to_dict()
# create an instance of MonsterMultiAttackAction from a dict
monster_multi_attack_action_form_dict = monster_multi_attack_action.from_dict(monster_multi_attack_action_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


