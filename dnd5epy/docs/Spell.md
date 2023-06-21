# Spell

`Spell` 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **str** | Resource index for shorthand searching. | [optional] 
**name** | **str** | Name of the referenced resource. | [optional] 
**url** | **str** | URL of the referenced resource. | [optional] 
**desc** | **List[str]** | Description of the resource. | [optional] 
**higher_level** | **List[str]** | List of descriptions for casting the spell at higher levels. | [optional] 
**range** | **str** | Range of the spell, usually expressed in feet. | [optional] 
**components** | **List[str]** | List of shorthand for required components of the spell. V: verbal S: somatic M: material  | [optional] 
**material** | **str** | Material component for the spell to be cast. | [optional] 
**area_of_effect** | [**AreaOfEffect**](AreaOfEffect.md) |  | [optional] 
**ritual** | **bool** | Determines if a spell can be cast in a 10-min(in-game) ritual. | [optional] 
**duration** | **str** | How long the spell effect lasts. | [optional] 
**concentration** | **bool** | Determines if a spell needs concentration to persist. | [optional] 
**casting_time** | **str** | How long it takes for the spell to activate. | [optional] 
**level** | **float** | Level of the spell. | [optional] 
**attack_type** | **str** | Attack type of the spell. | [optional] 
**damage** | [**SpellAllOfDamage**](SpellAllOfDamage.md) |  | [optional] 
**school** | [**APIReference**](APIReference.md) |  | [optional] 
**classes** | [**List[APIReference]**](APIReference.md) | List of classes that are able to learn the spell. | [optional] 
**subclasses** | [**List[APIReference]**](APIReference.md) | List of subclasses that have access to the spell. | [optional] 

## Example

```python
from dnd5epy.models.spell import Spell

# TODO update the JSON string below
json = "{}"
# create an instance of Spell from a JSON string
spell_instance = Spell.from_json(json)
# print the JSON string representation of the object
print Spell.to_json()

# convert the object into a dict
spell_dict = spell_instance.to_dict()
# create an instance of Spell from a dict
spell_form_dict = spell.from_dict(spell_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


