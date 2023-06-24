# SubclassLevelSpellcasting

Summary of spells known at this level.

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cantrips_known** | **float** |  | [optional] 
**spells_known** | **float** |  | [optional] 
**spell_slots_level_1** | **float** |  | [optional] 
**spell_slots_level_2** | **float** |  | [optional] 
**spell_slots_level_3** | **float** |  | [optional] 
**spell_slots_level_4** | **float** |  | [optional] 
**spell_slots_level_5** | **float** |  | [optional] 
**spell_slots_level_6** | **float** |  | [optional] 
**spell_slots_level_7** | **float** |  | [optional] 
**spell_slots_level_8** | **float** |  | [optional] 
**spell_slots_level_9** | **float** |  | [optional] 

## Example

```python
from dnd5epy.models.subclass_level_spellcasting import SubclassLevelSpellcasting

# TODO update the JSON string below
json = "{}"
# create an instance of SubclassLevelSpellcasting from a JSON string
subclass_level_spellcasting_instance = SubclassLevelSpellcasting.from_json(json)
# print the JSON string representation of the object
print SubclassLevelSpellcasting.to_json()

# convert the object into a dict
subclass_level_spellcasting_dict = subclass_level_spellcasting_instance.to_dict()
# create an instance of SubclassLevelSpellcasting from a dict
subclass_level_spellcasting_form_dict = subclass_level_spellcasting.from_dict(subclass_level_spellcasting_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


