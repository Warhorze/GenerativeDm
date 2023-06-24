# OptionSetOneOf2


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**option_set_type** | **str** | Type of option set; determines other attributes. | [optional] 
**resource_list** | **str** | A reference (by URL) to a collection in the database. | [optional] 

## Example

```python
from dnd5epy.models.option_set_one_of2 import OptionSetOneOf2

# TODO update the JSON string below
json = "{}"
# create an instance of OptionSetOneOf2 from a JSON string
option_set_one_of2_instance = OptionSetOneOf2.from_json(json)
# print the JSON string representation of the object
print OptionSetOneOf2.to_json()

# convert the object into a dict
option_set_one_of2_dict = option_set_one_of2_instance.to_dict()
# create an instance of OptionSetOneOf2 from a dict
option_set_one_of2_form_dict = option_set_one_of2.from_dict(option_set_one_of2_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


