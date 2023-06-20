# MagicItem

`MagicItem` 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **str** | Resource index for shorthand searching. | [optional] 
**name** | **str** | Name of the referenced resource. | [optional] 
**url** | **str** | URL of the referenced resource. | [optional] 
**desc** | **List[str]** | Description of the resource. | [optional] 
**equipment_category** | [**APIReference**](APIReference.md) |  | [optional] 
**rarity** | [**MagicItemAllOfRarity**](MagicItemAllOfRarity.md) |  | [optional] 
**variants** | [**List[APIReference]**](APIReference.md) |  | [optional] 
**variant** | **bool** | Whether this is a variant or not | [optional] 

## Example

```python
from openapi_client.models.magic_item import MagicItem

# TODO update the JSON string below
json = "{}"
# create an instance of MagicItem from a JSON string
magic_item_instance = MagicItem.from_json(json)
# print the JSON string representation of the object
print MagicItem.to_json()

# convert the object into a dict
magic_item_dict = magic_item_instance.to_dict()
# create an instance of MagicItem from a dict
magic_item_form_dict = magic_item.from_dict(magic_item_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


