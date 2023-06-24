# MonsterAbility

`Monster Ability` 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**charisma** | **float** | A monster&#39;s ability to charm or intimidate a player. | [optional] 
**constitution** | **float** | How sturdy a monster is.\&quot; | [optional] 
**dexterity** | **float** | The monster&#39;s ability for swift movement or stealth | [optional] 
**intelligence** | **float** | The monster&#39;s ability to outsmart a player. | [optional] 
**strength** | **float** | How hard a monster can hit a player. | [optional] 
**wisdom** | **float** | A monster&#39;s ability to ascertain the player&#39;s plan. | [optional] 

## Example

```python
from openapi_client.models.monster_ability import MonsterAbility

# TODO update the JSON string below
json = "{}"
# create an instance of MonsterAbility from a JSON string
monster_ability_instance = MonsterAbility.from_json(json)
# print the JSON string representation of the object
print MonsterAbility.to_json()

# convert the object into a dict
monster_ability_dict = monster_ability_instance.to_dict()
# create an instance of MonsterAbility from a dict
monster_ability_form_dict = monster_ability.from_dict(monster_ability_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


