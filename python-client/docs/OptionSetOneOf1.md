# OptionSetOneOf1


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**option_set_type** | **str** | Type of option set; determines other attributes. | [optional] 
**equipment_category** | [**APIReference**](APIReference.md) |  | [optional] 

## Example

```python
from openapi_client.models.option_set_one_of1 import OptionSetOneOf1

# TODO update the JSON string below
json = "{}"
# create an instance of OptionSetOneOf1 from a JSON string
option_set_one_of1_instance = OptionSetOneOf1.from_json(json)
# print the JSON string representation of the object
print OptionSetOneOf1.to_json()

# convert the object into a dict
option_set_one_of1_dict = option_set_one_of1_instance.to_dict()
# create an instance of OptionSetOneOf1 from a dict
option_set_one_of1_form_dict = option_set_one_of1.from_dict(option_set_one_of1_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


