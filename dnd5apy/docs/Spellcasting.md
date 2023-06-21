# Spellcasting

`Spellcasting` 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**level** | **float** | Level at which the class can start using its spellcasting abilities. | [optional] 
**info** | [**List[SpellcastingInfoInner]**](SpellcastingInfoInner.md) | Descriptions of the class&#39; ability to cast spells. | [optional] 
**spellcasting_ability** | [**SpellcastingSpellcastingAbility**](SpellcastingSpellcastingAbility.md) |  | [optional] 

## Example

```python
from openapi_client.models.spellcasting import Spellcasting

# TODO update the JSON string below
json = "{}"
# create an instance of Spellcasting from a JSON string
spellcasting_instance = Spellcasting.from_json(json)
# print the JSON string representation of the object
print Spellcasting.to_json()

# convert the object into a dict
spellcasting_dict = spellcasting_instance.to_dict()
# create an instance of Spellcasting from a dict
spellcasting_form_dict = spellcasting.from_dict(spellcasting_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


