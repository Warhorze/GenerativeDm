# Condition

`Condition` 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **str** | Resource index for shorthand searching. | [optional] 
**name** | **str** | Name of the referenced resource. | [optional] 
**url** | **str** | URL of the referenced resource. | [optional] 
**desc** | **List[str]** | Description of the resource. | [optional] 

## Example

```python
from dnd5epy.models.condition import Condition

# TODO update the JSON string below
json = "{}"
# create an instance of Condition from a JSON string
condition_instance = Condition.from_json(json)
# print the JSON string representation of the object
print Condition.to_json()

# convert the object into a dict
condition_dict = condition_instance.to_dict()
# create an instance of Condition from a dict
condition_form_dict = condition.from_dict(condition_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

