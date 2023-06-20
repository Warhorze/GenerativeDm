# DamageType

`DamageType` 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **str** | Resource index for shorthand searching. | [optional] 
**name** | **str** | Name of the referenced resource. | [optional] 
**url** | **str** | URL of the referenced resource. | [optional] 
**desc** | **List[str]** | Description of the resource. | [optional] 

## Example

```python
from openapi_client.models.damage_type import DamageType

# TODO update the JSON string below
json = "{}"
# create an instance of DamageType from a JSON string
damage_type_instance = DamageType.from_json(json)
# print the JSON string representation of the object
print DamageType.to_json()

# convert the object into a dict
damage_type_dict = damage_type_instance.to_dict()
# create an instance of DamageType from a dict
damage_type_form_dict = damage_type.from_dict(damage_type_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


