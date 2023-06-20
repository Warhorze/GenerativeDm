# MonsterSpell


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** |  | [optional] 
**level** | **float** |  | [optional] 
**url** | **str** |  | [optional] 
**usage** | [**MonsterUsage**](MonsterUsage.md) |  | [optional] 

## Example

```python
from openapi_client.models.monster_spell import MonsterSpell

# TODO update the JSON string below
json = "{}"
# create an instance of MonsterSpell from a JSON string
monster_spell_instance = MonsterSpell.from_json(json)
# print the JSON string representation of the object
print MonsterSpell.to_json()

# convert the object into a dict
monster_spell_dict = monster_spell_instance.to_dict()
# create an instance of MonsterSpell from a dict
monster_spell_form_dict = monster_spell.from_dict(monster_spell_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


