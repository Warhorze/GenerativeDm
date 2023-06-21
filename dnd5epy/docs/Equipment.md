# Equipment

`Equipment` 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **str** | Resource index for shorthand searching. | [optional] 
**name** | **str** | Name of the referenced resource. | [optional] 
**url** | **str** | URL of the referenced resource. | [optional] 
**desc** | **List[str]** | Description of the resource. | [optional] 
**equipment_category** | [**APIReference**](APIReference.md) |  | [optional] 
**weapon_category** | **str** | The category of weapon this falls into. | [optional] 
**weapon_range** | **str** | Whether this is a Melee or Ranged weapon. | [optional] 
**category_range** | **str** | A combination of weapon_category and weapon_range. | [optional] 
**range** | [**WeaponAllOfRange**](WeaponAllOfRange.md) |  | [optional] 
**damage** | [**Damage**](Damage.md) |  | [optional] 
**two_handed_damage** | [**Damage**](Damage.md) |  | [optional] 
**properties** | [**List[APIReference]**](APIReference.md) | A list of the properties this weapon has. | [optional] 
**cost** | [**Cost**](Cost.md) |  | [optional] 
**weight** | **float** | How much the equipment weighs. | [optional] 
**armor_category** | **str** | The category of armor this falls into. | [optional] 
**armor_class** | **Dict[str, str]** | Details on how to calculate armor class. | [optional] 
**str_minimum** | **float** | Minimum STR required to use this armor. | [optional] 
**stealth_disadvantage** | **bool** | Whether the armor gives disadvantage for Stealth. | [optional] 
**gear_category** | [**APIReference**](APIReference.md) |  | [optional] 
**contents** | [**List[APIReference]**](APIReference.md) | The list of adventuring gear in the pack. | [optional] 

## Example

```python
from dnd5epy.models.equipment import Equipment

# TODO update the JSON string below
json = "{}"
# create an instance of Equipment from a JSON string
equipment_instance = Equipment.from_json(json)
# print the JSON string representation of the object
print Equipment.to_json()

# convert the object into a dict
equipment_dict = equipment_instance.to_dict()
# create an instance of Equipment from a dict
equipment_form_dict = equipment.from_dict(equipment_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


