# OptionOneOf


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**option_type** | **str** | Type of option; determines other attributes. | [optional] 
**item** | [**APIReference**](APIReference.md) |  | [optional] 

## Example

```python
from dnd5epy.models.option_one_of import OptionOneOf

# TODO update the JSON string below
json = "{}"
# create an instance of OptionOneOf from a JSON string
option_one_of_instance = OptionOneOf.from_json(json)
# print the JSON string representation of the object
print OptionOneOf.to_json()

# convert the object into a dict
option_one_of_dict = option_one_of_instance.to_dict()
# create an instance of OptionOneOf from a dict
option_one_of_form_dict = option_one_of.from_dict(option_one_of_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


