# OptionSet

`Option Set` 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**option_set_type** | **str** | Type of option set; determines other attributes. | [optional] 
**options_array** | [**List[Option]**](Option.md) | Array of options to choose from. | [optional] 
**equipment_category** | [**APIReference**](APIReference.md) |  | [optional] 
**resource_list** | **str** | A reference (by URL) to a collection in the database. | [optional] 

## Example

```python
from dnd5epy.models.option_set import OptionSet

# TODO update the JSON string below
json = "{}"
# create an instance of OptionSet from a JSON string
option_set_instance = OptionSet.from_json(json)
# print the JSON string representation of the object
print OptionSet.to_json()

# convert the object into a dict
option_set_dict = option_set_instance.to_dict()
# create an instance of OptionSet from a dict
option_set_form_dict = option_set.from_dict(option_set_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


