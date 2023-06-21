# AbilityScore

`AbilityScore` 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **str** | Resource index for shorthand searching. | [optional] 
**name** | **str** | Name of the referenced resource. | [optional] 
**url** | **str** | URL of the referenced resource. | [optional] 
**desc** | **List[str]** | Description of the resource. | [optional] 
**full_name** | **str** | Full name of the ability score. | [optional] 
**skills** | [**List[APIReference]**](APIReference.md) | List of skills that use this ability score. | [optional] 

## Example

```python
from dnd5epy.models.ability_score import AbilityScore

# TODO update the JSON string below
json = "{}"
# create an instance of AbilityScore from a JSON string
ability_score_instance = AbilityScore.from_json(json)
# print the JSON string representation of the object
print AbilityScore.to_json()

# convert the object into a dict
ability_score_dict = ability_score_instance.to_dict()
# create an instance of AbilityScore from a dict
ability_score_form_dict = ability_score.from_dict(ability_score_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


