# OptionOneOf2


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**option_type** | **str** | Type of option; determines other attributes. | [optional] 
**items** | [**List[Option]**](Option.md) |  | [optional] 

## Example

```python
from dnd5epy.models.option_one_of2 import OptionOneOf2

# TODO update the JSON string below
json = "{}"
# create an instance of OptionOneOf2 from a JSON string
option_one_of2_instance = OptionOneOf2.from_json(json)
# print the JSON string representation of the object
print OptionOneOf2.to_json()

# convert the object into a dict
option_one_of2_dict = option_one_of2_instance.to_dict()
# create an instance of OptionOneOf2 from a dict
option_one_of2_form_dict = option_one_of2.from_dict(option_one_of2_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


