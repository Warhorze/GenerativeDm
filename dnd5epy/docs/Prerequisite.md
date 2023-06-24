# Prerequisite

`Prerequisite` 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ability_score** | [**PrerequisiteAbilityScore**](PrerequisiteAbilityScore.md) |  | [optional] 
**minimum_score** | **float** | Minimum score to meet the prerequisite. | [optional] 

## Example

```python
from dnd5epy.models.prerequisite import Prerequisite

# TODO update the JSON string below
json = "{}"
# create an instance of Prerequisite from a JSON string
prerequisite_instance = Prerequisite.from_json(json)
# print the JSON string representation of the object
print Prerequisite.to_json()

# convert the object into a dict
prerequisite_dict = prerequisite_instance.to_dict()
# create an instance of Prerequisite from a dict
prerequisite_form_dict = prerequisite.from_dict(prerequisite_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


