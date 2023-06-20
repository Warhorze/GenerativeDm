# OptionOneOf10


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**option_type** | **str** | Type of option; determines other attributes. | [optional] 
**damage_type** | [**APIReference**](APIReference.md) |  | [optional] 
**damage_dice** | **str** | Damage expressed in dice (e.g. \&quot;13d6\&quot;). | [optional] 
**notes** | **str** | Information regarding the damage. | [optional] 

## Example

```python
from openapi_client.models.option_one_of10 import OptionOneOf10

# TODO update the JSON string below
json = "{}"
# create an instance of OptionOneOf10 from a JSON string
option_one_of10_instance = OptionOneOf10.from_json(json)
# print the JSON string representation of the object
print OptionOneOf10.to_json()

# convert the object into a dict
option_one_of10_dict = option_one_of10_instance.to_dict()
# create an instance of OptionOneOf10 from a dict
option_one_of10_form_dict = option_one_of10.from_dict(option_one_of10_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


