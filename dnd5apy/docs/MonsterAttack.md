# MonsterAttack


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** |  | [optional] 
**dc** | [**DC**](DC.md) |  | [optional] 
**damage** | [**Damage**](Damage.md) |  | [optional] 

## Example

```python
from openapi_client.models.monster_attack import MonsterAttack

# TODO update the JSON string below
json = "{}"
# create an instance of MonsterAttack from a JSON string
monster_attack_instance = MonsterAttack.from_json(json)
# print the JSON string representation of the object
print MonsterAttack.to_json()

# convert the object into a dict
monster_attack_dict = monster_attack_instance.to_dict()
# create an instance of MonsterAttack from a dict
monster_attack_form_dict = monster_attack.from_dict(monster_attack_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


