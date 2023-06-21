# WeaponProperty

WeaponProperty

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **str** | Resource index for shorthand searching. | [optional] 
**name** | **str** | Name of the referenced resource. | [optional] 
**url** | **str** | URL of the referenced resource. | [optional] 
**desc** | **List[str]** | Description of the resource. | [optional] 

## Example

```python
from dnd5epy.models.weapon_property import WeaponProperty

# TODO update the JSON string below
json = "{}"
# create an instance of WeaponProperty from a JSON string
weapon_property_instance = WeaponProperty.from_json(json)
# print the JSON string representation of the object
print WeaponProperty.to_json()

# convert the object into a dict
weapon_property_dict = weapon_property_instance.to_dict()
# create an instance of WeaponProperty from a dict
weapon_property_form_dict = weapon_property.from_dict(weapon_property_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


