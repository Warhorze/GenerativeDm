# SpellcastingSpellcastingAbility

Reference to the `AbilityScore` used for spellcasting by the class.

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **str** | Resource index for shorthand searching. | [optional] 
**name** | **str** | Name of the referenced resource. | [optional] 
**url** | **str** | URL of the referenced resource. | [optional] 

## Example

```python
from openapi_client.models.spellcasting_spellcasting_ability import SpellcastingSpellcastingAbility

# TODO update the JSON string below
json = "{}"
# create an instance of SpellcastingSpellcastingAbility from a JSON string
spellcasting_spellcasting_ability_instance = SpellcastingSpellcastingAbility.from_json(json)
# print the JSON string representation of the object
print SpellcastingSpellcastingAbility.to_json()

# convert the object into a dict
spellcasting_spellcasting_ability_dict = spellcasting_spellcasting_ability_instance.to_dict()
# create an instance of SpellcastingSpellcastingAbility from a dict
spellcasting_spellcasting_ability_form_dict = spellcasting_spellcasting_ability.from_dict(spellcasting_spellcasting_ability_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


