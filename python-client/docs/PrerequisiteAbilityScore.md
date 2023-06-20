# PrerequisiteAbilityScore


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **str** | Resource index for shorthand searching. | [optional] 
**name** | **str** | Name of the referenced resource. | [optional] 
**url** | **str** | URL of the referenced resource. | [optional] 

## Example

```python
from openapi_client.models.prerequisite_ability_score import PrerequisiteAbilityScore

# TODO update the JSON string below
json = "{}"
# create an instance of PrerequisiteAbilityScore from a JSON string
prerequisite_ability_score_instance = PrerequisiteAbilityScore.from_json(json)
# print the JSON string representation of the object
print PrerequisiteAbilityScore.to_json()

# convert the object into a dict
prerequisite_ability_score_dict = prerequisite_ability_score_instance.to_dict()
# create an instance of PrerequisiteAbilityScore from a dict
prerequisite_ability_score_form_dict = prerequisite_ability_score.from_dict(prerequisite_ability_score_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


