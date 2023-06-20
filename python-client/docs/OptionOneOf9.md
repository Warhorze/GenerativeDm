# OptionOneOf9


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**option_type** | **str** | Type of option; determines other attributes. | [optional] 
**name** | **str** | Name of the breath | [optional] 
**dc** | [**DC**](DC.md) |  | [optional] 
**damage** | [**List[Damage]**](Damage.md) | Damage dealt by the breath attack, if any. | [optional] 

## Example

```python
from openapi_client.models.option_one_of9 import OptionOneOf9

# TODO update the JSON string below
json = "{}"
# create an instance of OptionOneOf9 from a JSON string
option_one_of9_instance = OptionOneOf9.from_json(json)
# print the JSON string representation of the object
print OptionOneOf9.to_json()

# convert the object into a dict
option_one_of9_dict = option_one_of9_instance.to_dict()
# create an instance of OptionOneOf9 from a dict
option_one_of9_form_dict = option_one_of9.from_dict(option_one_of9_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


