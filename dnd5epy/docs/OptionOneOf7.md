# OptionOneOf7


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**option_type** | **str** | Type of option; determines other attributes. | [optional] 
**ability_score** | [**APIReference**](APIReference.md) |  | [optional] 
**minimum_score** | **float** | The minimum score required to satisfy the prerequisite. | [optional] 

## Example

```python
from dnd5epy.models.option_one_of7 import OptionOneOf7

# TODO update the JSON string below
json = "{}"
# create an instance of OptionOneOf7 from a JSON string
option_one_of7_instance = OptionOneOf7.from_json(json)
# print the JSON string representation of the object
print OptionOneOf7.to_json()

# convert the object into a dict
option_one_of7_dict = option_one_of7_instance.to_dict()
# create an instance of OptionOneOf7 from a dict
option_one_of7_form_dict = option_one_of7.from_dict(option_one_of7_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


