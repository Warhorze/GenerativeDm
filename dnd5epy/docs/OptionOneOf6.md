# OptionOneOf6


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**option_type** | **str** | Type of option; determines other attributes. | [optional] 
**count** | **float** | Count | [optional] 
**of** | [**APIReference**](APIReference.md) |  | [optional] 

## Example

```python
from dnd5epy.models.option_one_of6 import OptionOneOf6

# TODO update the JSON string below
json = "{}"
# create an instance of OptionOneOf6 from a JSON string
option_one_of6_instance = OptionOneOf6.from_json(json)
# print the JSON string representation of the object
print OptionOneOf6.to_json()

# convert the object into a dict
option_one_of6_dict = option_one_of6_instance.to_dict()
# create an instance of OptionOneOf6 from a dict
option_one_of6_form_dict = option_one_of6.from_dict(option_one_of6_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


