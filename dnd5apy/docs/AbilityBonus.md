# AbilityBonus


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**bonus** | **float** | Bonus amount for this ability score. | [optional] 
**ability_score** | [**APIReference**](APIReference.md) |  | [optional] 

## Example

```python
from openapi_client.models.ability_bonus import AbilityBonus

# TODO update the JSON string below
json = "{}"
# create an instance of AbilityBonus from a JSON string
ability_bonus_instance = AbilityBonus.from_json(json)
# print the JSON string representation of the object
print AbilityBonus.to_json()

# convert the object into a dict
ability_bonus_dict = ability_bonus_instance.to_dict()
# create an instance of AbilityBonus from a dict
ability_bonus_form_dict = ability_bonus.from_dict(ability_bonus_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


