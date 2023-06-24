# MonsterSpellcasting


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ability** | [**APIReference**](APIReference.md) |  | [optional] 
**dc** | **float** |  | [optional] 
**modifier** | **float** |  | [optional] 
**components_required** | **List[str]** |  | [optional] 
**school** | **str** |  | [optional] 
**slots** | **Dict[str, float]** |  | [optional] 
**spells** | [**List[MonsterSpell]**](MonsterSpell.md) |  | [optional] 

## Example

```python
from dnd5epy.models.monster_spellcasting import MonsterSpellcasting

# TODO update the JSON string below
json = "{}"
# create an instance of MonsterSpellcasting from a JSON string
monster_spellcasting_instance = MonsterSpellcasting.from_json(json)
# print the JSON string representation of the object
print MonsterSpellcasting.to_json()

# convert the object into a dict
monster_spellcasting_dict = monster_spellcasting_instance.to_dict()
# create an instance of MonsterSpellcasting from a dict
monster_spellcasting_form_dict = monster_spellcasting.from_dict(monster_spellcasting_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


