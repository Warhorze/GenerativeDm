# OptionOneOf3


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**option_type** | **str** | Type of option; determines other attributes. | [optional] 
**choice** | [**Choice**](Choice.md) |  | [optional] 

## Example

```python
from openapi_client.models.option_one_of3 import OptionOneOf3

# TODO update the JSON string below
json = "{}"
# create an instance of OptionOneOf3 from a JSON string
option_one_of3_instance = OptionOneOf3.from_json(json)
# print the JSON string representation of the object
print OptionOneOf3.to_json()

# convert the object into a dict
option_one_of3_dict = option_one_of3_instance.to_dict()
# create an instance of OptionOneOf3 from a dict
option_one_of3_form_dict = option_one_of3.from_dict(option_one_of3_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


