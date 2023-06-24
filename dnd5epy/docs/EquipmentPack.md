# EquipmentPack

`EquipmentPack` 

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
**contents** | [**List[APIReference]**](APIReference.md) | The list of adventuring gear in the pack. | [optional] 

## Example

```python
from openapi_client.models.equipment_pack import EquipmentPack

# TODO update the JSON string below
json = "{}"
# create an instance of EquipmentPack from a JSON string
equipment_pack_instance = EquipmentPack.from_json(json)
# print the JSON string representation of the object
print EquipmentPack.to_json()

# convert the object into a dict
equipment_pack_dict = equipment_pack_instance.to_dict()
# create an instance of EquipmentPack from a dict
equipment_pack_form_dict = equipment_pack.from_dict(equipment_pack_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


