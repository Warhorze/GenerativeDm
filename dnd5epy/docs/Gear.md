# Gear

`Gear` 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **str** | Resource index for shorthand searching. | [optional] 
**name** | **str** | Name of the referenced resource. | [optional] 
**url** | **str** | URL of the referenced resource. | [optional] 
**desc** | **List[str]** | Description of the resource. | [optional] 
**equipment_category** | [**APIReference**](APIReference.md) |  | [optional] 
**gear_category** | [**APIReference**](APIReference.md) |  | [optional] 
**cost** | [**Cost**](Cost.md) |  | [optional] 
**weight** | **float** | How much the equipment weighs. | [optional] 

## Example

```python
from dnd5epy.models.gear import Gear

# TODO update the JSON string below
json = "{}"
# create an instance of Gear from a JSON string
gear_instance = Gear.from_json(json)
# print the JSON string representation of the object
print Gear.to_json()

# convert the object into a dict
gear_dict = gear_instance.to_dict()
# create an instance of Gear from a dict
gear_form_dict = gear.from_dict(gear_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


