# OptionOneOf8


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**option_type** | **str** | Type of option; determines other attributes. | [optional] 
**ability_score** | [**APIReference**](APIReference.md) |  | [optional] 
**bonus** | **float** | The bonus being applied to the ability score | [optional] 

## Example

```python
from dnd5epy.models.option_one_of8 import OptionOneOf8

# TODO update the JSON string below
json = "{}"
# create an instance of OptionOneOf8 from a JSON string
option_one_of8_instance = OptionOneOf8.from_json(json)
# print the JSON string representation of the object
print OptionOneOf8.to_json()

# convert the object into a dict
option_one_of8_dict = option_one_of8_instance.to_dict()
# create an instance of OptionOneOf8 from a dict
option_one_of8_form_dict = option_one_of8.from_dict(option_one_of8_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


