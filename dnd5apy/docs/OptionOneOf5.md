# OptionOneOf5


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**option_type** | **str** | Type of option; determines other attributes. | [optional] 
**desc** | **str** | A description of the ideal. | [optional] 
**alignments** | [**List[APIReference]**](APIReference.md) | A list of alignments of those who might follow the ideal. | [optional] 

## Example

```python
from openapi_client.models.option_one_of5 import OptionOneOf5

# TODO update the JSON string below
json = "{}"
# create an instance of OptionOneOf5 from a JSON string
option_one_of5_instance = OptionOneOf5.from_json(json)
# print the JSON string representation of the object
print OptionOneOf5.to_json()

# convert the object into a dict
option_one_of5_dict = option_one_of5_instance.to_dict()
# create an instance of OptionOneOf5 from a dict
option_one_of5_form_dict = option_one_of5.from_dict(option_one_of5_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


