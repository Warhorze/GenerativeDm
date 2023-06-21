# OptionOneOf1


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**option_type** | **str** | Type of option; determines other attributes. | [optional] 
**action_name** | **str** | The name of the action. | [optional] 
**count** | **float** | The number of times this action can be repeated if chosen. | [optional] 
**type** | **str** | For attack options that can be melee, ranged, abilities, or thrown. | [optional] 

## Example

```python
from dnd5epy.models.option_one_of1 import OptionOneOf1

# TODO update the JSON string below
json = "{}"
# create an instance of OptionOneOf1 from a JSON string
option_one_of1_instance = OptionOneOf1.from_json(json)
# print the JSON string representation of the object
print OptionOneOf1.to_json()

# convert the object into a dict
option_one_of1_dict = option_one_of1_instance.to_dict()
# create an instance of OptionOneOf1 from a dict
option_one_of1_form_dict = option_one_of1.from_dict(option_one_of1_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


