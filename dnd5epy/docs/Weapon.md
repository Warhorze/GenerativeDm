# Weapon

`Weapon` 

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

## Example

```python
from dnd5epy.models.weapon import Weapon

# TODO update the JSON string below
json = "{}"
# create an instance of Weapon from a JSON string
weapon_instance = Weapon.from_json(json)
# print the JSON string representation of the object
print Weapon.to_json()

# convert the object into a dict
weapon_dict = weapon_instance.to_dict()
# create an instance of Weapon from a dict
weapon_form_dict = weapon.from_dict(weapon_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


