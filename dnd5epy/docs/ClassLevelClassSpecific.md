# ClassLevelClassSpecific

Class specific information such as dice values for bard songs and number of warlock invocations.

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**rage_count** | **float** |  | [optional] 
**rage_damage_bonus** | **float** |  | [optional] 
**brutal_critical_dice** | **float** |  | [optional] 
**bardic_inspiration_dice** | **float** |  | [optional] 
**song_of_rest_die** | **float** |  | [optional] 
**magical_secrets_max_5** | **float** |  | [optional] 
**magical_secrets_max_7** | **float** |  | [optional] 
**magical_secrets_max_9** | **float** |  | [optional] 
**channel_divinity_charges** | **float** |  | [optional] 
**destroy_undead_cr** | **float** |  | [optional] 
**wild_shape_max_cr** | **float** |  | [optional] 
**wild_shape_swim** | **bool** |  | [optional] 
**wild_shape_fly** | **bool** |  | [optional] 
**action_surges** | **float** |  | [optional] 
**indomitable_uses** | **float** |  | [optional] 
**extra_attacks** | **float** |  | [optional] 
**ki_points** | **float** |  | [optional] 
**unarmored_movement** | **float** |  | [optional] 
**martial_arts** | [**ClassLevelClassSpecificAnyOf5MartialArts**](ClassLevelClassSpecificAnyOf5MartialArts.md) |  | [optional] 
**aura_range** | **float** |  | [optional] 
**favored_enemies** | **float** |  | [optional] 
**favored_terrain** | **float** |  | [optional] 
**sneak_attack** | [**ClassLevelClassSpecificAnyOf5MartialArts**](ClassLevelClassSpecificAnyOf5MartialArts.md) |  | [optional] 
**sorcery_points** | **float** |  | [optional] 
**metamagic_known** | **float** |  | [optional] 
**creating_spell_slots** | [**List[ClassLevelClassSpecificAnyOf9CreatingSpellSlotsInner]**](ClassLevelClassSpecificAnyOf9CreatingSpellSlotsInner.md) |  | [optional] 
**invocations_known** | **float** |  | [optional] 
**mystic_arcanum_level_6** | **float** |  | [optional] 
**mystic_arcanum_level_7** | **float** |  | [optional] 
**mystic_arcanum_level_8** | **float** |  | [optional] 
**mystic_arcanum_level_9** | **float** |  | [optional] 
**arcane_recover_levels** | **float** |  | [optional] 

## Example

```python
from dnd5epy.models.class_level_class_specific import ClassLevelClassSpecific

# TODO update the JSON string below
json = "{}"
# create an instance of ClassLevelClassSpecific from a JSON string
class_level_class_specific_instance = ClassLevelClassSpecific.from_json(json)
# print the JSON string representation of the object
print ClassLevelClassSpecific.to_json()

# convert the object into a dict
class_level_class_specific_dict = class_level_class_specific_instance.to_dict()
# create an instance of ClassLevelClassSpecific from a dict
class_level_class_specific_form_dict = class_level_class_specific.from_dict(class_level_class_specific_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


