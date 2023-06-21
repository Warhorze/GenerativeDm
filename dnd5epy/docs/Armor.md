# Armor

`Armor` 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **str** | Resource index for shorthand searching. | [optional] 
**name** | **str** | Name of the referenced resource. | [optional] 
**url** | **str** | URL of the referenced resource. | [optional] 
**desc** | **List[str]** | Description of the resource. | [optional] 
**equipment_category** | [**APIReference**](APIReference.md) |  | [optional] 
**armor_category** | **str** | The category of armor this falls into. | [optional] 
**armor_class** | **Dict[str, str]** | Details on how to calculate armor class. | [optional] 
**str_minimum** | **float** | Minimum STR required to use this armor. | [optional] 
**stealth_disadvantage** | **bool** | Whether the armor gives disadvantage for Stealth. | [optional] 
**cost** | [**Cost**](Cost.md) |  | [optional] 
**weight** | **float** | How much the equipment weighs. | [optional] 

## Example

```python
from dnd5epy.models.armor import Armor

# TODO update the JSON string below
json = "{}"
# create an instance of Armor from a JSON string
armor_instance = Armor.from_json(json)
# print the JSON string representation of the object
print Armor.to_json()

# convert the object into a dict
armor_dict = armor_instance.to_dict()
# create an instance of Armor from a dict
armor_form_dict = armor.from_dict(armor_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


