# Subclass

`Subclass` 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **str** | Resource index for shorthand searching. | [optional] 
**name** | **str** | Name of the referenced resource. | [optional] 
**url** | **str** | URL of the referenced resource. | [optional] 
**desc** | **List[str]** | Description of the resource. | [optional] 
**var_class** | [**APIReference**](APIReference.md) |  | [optional] 
**subclass_flavor** | **str** | Lore-friendly flavor text for a classes respective subclass. | [optional] 
**subclass_levels** | **str** | Resource url that shows the subclass level progression. | [optional] 
**spells** | [**List[SubclassAllOfSpells]**](SubclassAllOfSpells.md) |  | [optional] 

## Example

```python
from openapi_client.models.subclass import Subclass

# TODO update the JSON string below
json = "{}"
# create an instance of Subclass from a JSON string
subclass_instance = Subclass.from_json(json)
# print the JSON string representation of the object
print Subclass.to_json()

# convert the object into a dict
subclass_dict = subclass_instance.to_dict()
# create an instance of Subclass from a dict
subclass_form_dict = subclass.from_dict(subclass_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


