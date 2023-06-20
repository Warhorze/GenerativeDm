# DamageAtCharacterLevel

'Spell Damage' 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**damage_at_character_level** | **Dict[str, object]** |  | [optional] 
**damage_type** | [**APIReference**](APIReference.md) |  | [optional] 

## Example

```python
from openapi_client.models.damage_at_character_level import DamageAtCharacterLevel

# TODO update the JSON string below
json = "{}"
# create an instance of DamageAtCharacterLevel from a JSON string
damage_at_character_level_instance = DamageAtCharacterLevel.from_json(json)
# print the JSON string representation of the object
print DamageAtCharacterLevel.to_json()

# convert the object into a dict
damage_at_character_level_dict = damage_at_character_level_instance.to_dict()
# create an instance of DamageAtCharacterLevel from a dict
damage_at_character_level_form_dict = damage_at_character_level.from_dict(damage_at_character_level_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


