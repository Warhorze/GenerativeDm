# MonsterSpecialAbility


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** |  | [optional] 
**desc** | **str** |  | [optional] 
**attack_bonus** | **float** |  | [optional] 
**damage** | [**Damage**](Damage.md) |  | [optional] 
**dc** | [**DC**](DC.md) |  | [optional] 
**spellcasting** | [**MonsterSpellcasting**](MonsterSpellcasting.md) |  | [optional] 
**usage** | [**MonsterUsage**](MonsterUsage.md) |  | [optional] 

## Example

```python
from dnd5epy.models.monster_special_ability import MonsterSpecialAbility

# TODO update the JSON string below
json = "{}"
# create an instance of MonsterSpecialAbility from a JSON string
monster_special_ability_instance = MonsterSpecialAbility.from_json(json)
# print the JSON string representation of the object
print MonsterSpecialAbility.to_json()

# convert the object into a dict
monster_special_ability_dict = monster_special_ability_instance.to_dict()
# create an instance of MonsterSpecialAbility from a dict
monster_special_ability_form_dict = monster_special_ability.from_dict(monster_special_ability_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


