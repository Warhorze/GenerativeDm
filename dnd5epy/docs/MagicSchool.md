# MagicSchool

`MagicSchool` 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **str** | Resource index for shorthand searching. | [optional] 
**name** | **str** | Name of the referenced resource. | [optional] 
**url** | **str** | URL of the referenced resource. | [optional] 
**desc** | **str** | Brief description of the resource. | [optional] 

## Example

```python
from dnd5epy.models.magic_school import MagicSchool

# TODO update the JSON string below
json = "{}"
# create an instance of MagicSchool from a JSON string
magic_school_instance = MagicSchool.from_json(json)
# print the JSON string representation of the object
print MagicSchool.to_json()

# convert the object into a dict
magic_school_dict = magic_school_instance.to_dict()
# create an instance of MagicSchool from a dict
magic_school_form_dict = magic_school.from_dict(magic_school_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


