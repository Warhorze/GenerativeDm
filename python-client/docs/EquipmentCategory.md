# EquipmentCategory

`EquipmentCategory` 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **str** | Resource index for shorthand searching. | [optional] 
**name** | **str** | Name of the referenced resource. | [optional] 
**url** | **str** | URL of the referenced resource. | [optional] 
**equipment** | [**List[APIReference]**](APIReference.md) | A list of the equipment that falls into this category. | [optional] 

## Example

```python
from openapi_client.models.equipment_category import EquipmentCategory

# TODO update the JSON string below
json = "{}"
# create an instance of EquipmentCategory from a JSON string
equipment_category_instance = EquipmentCategory.from_json(json)
# print the JSON string representation of the object
print EquipmentCategory.to_json()

# convert the object into a dict
equipment_category_dict = equipment_category_instance.to_dict()
# create an instance of EquipmentCategory from a dict
equipment_category_form_dict = equipment_category.from_dict(equipment_category_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


