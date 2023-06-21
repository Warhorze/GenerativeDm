# SpellAllOfDamage


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**damage_at_character_level** | **Dict[str, object]** |  | [optional] 
**damage_type** | [**APIReference**](APIReference.md) |  | [optional] 
**damage_at_slot_level** | **Dict[str, object]** |  | [optional] 

## Example

```python
from dnd5epy.models.spell_all_of_damage import SpellAllOfDamage

# TODO update the JSON string below
json = "{}"
# create an instance of SpellAllOfDamage from a JSON string
spell_all_of_damage_instance = SpellAllOfDamage.from_json(json)
# print the JSON string representation of the object
print SpellAllOfDamage.to_json()

# convert the object into a dict
spell_all_of_damage_dict = spell_all_of_damage_instance.to_dict()
# create an instance of SpellAllOfDamage from a dict
spell_all_of_damage_form_dict = spell_all_of_damage.from_dict(spell_all_of_damage_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


