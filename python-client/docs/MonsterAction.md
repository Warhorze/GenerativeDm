# MonsterAction

Action available to a `Monster` in addition to the standard creature actions.

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** |  | [optional] 
**desc** | **str** |  | [optional] 
**action_options** | [**Choice**](Choice.md) |  | [optional] 
**actions** | [**List[MonsterMultiAttackAction]**](MonsterMultiAttackAction.md) |  | [optional] 
**options** | [**Choice**](Choice.md) |  | [optional] 
**multiattack_type** | **str** |  | [optional] 
**attack_bonus** | **float** |  | [optional] 
**dc** | [**DC**](DC.md) |  | [optional] 
**attacks** | [**List[MonsterAttack]**](MonsterAttack.md) |  | [optional] 
**damage** | [**List[Damage]**](Damage.md) |  | [optional] 

## Example

```python
from openapi_client.models.monster_action import MonsterAction

# TODO update the JSON string below
json = "{}"
# create an instance of MonsterAction from a JSON string
monster_action_instance = MonsterAction.from_json(json)
# print the JSON string representation of the object
print MonsterAction.to_json()

# convert the object into a dict
monster_action_dict = monster_action_instance.to_dict()
# create an instance of MonsterAction from a dict
monster_action_form_dict = monster_action.from_dict(monster_action_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


