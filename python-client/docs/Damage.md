# Damage

`Damage` 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**damage_dice** | **str** |  | [optional] 
**damage_type** | [**APIReference**](APIReference.md) |  | [optional] 

## Example

```python
from openapi_client.models.damage import Damage

# TODO update the JSON string below
json = "{}"
# create an instance of Damage from a JSON string
damage_instance = Damage.from_json(json)
# print the JSON string representation of the object
print Damage.to_json()

# convert the object into a dict
damage_dict = damage_instance.to_dict()
# create an instance of Damage from a dict
damage_form_dict = damage.from_dict(damage_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


