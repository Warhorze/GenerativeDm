# OptionSetOneOf


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**option_set_type** | **str** | Type of option set; determines other attributes. | [optional] 
**options_array** | [**List[Option]**](Option.md) | Array of options to choose from. | [optional] 

## Example

```python
from openapi_client.models.option_set_one_of import OptionSetOneOf

# TODO update the JSON string below
json = "{}"
# create an instance of OptionSetOneOf from a JSON string
option_set_one_of_instance = OptionSetOneOf.from_json(json)
# print the JSON string representation of the object
print OptionSetOneOf.to_json()

# convert the object into a dict
option_set_one_of_dict = option_set_one_of_instance.to_dict()
# create an instance of OptionSetOneOf from a dict
option_set_one_of_form_dict = option_set_one_of.from_dict(option_set_one_of_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


