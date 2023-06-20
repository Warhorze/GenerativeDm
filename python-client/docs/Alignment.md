# Alignment

`Alignment` 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **str** | Resource index for shorthand searching. | [optional] 
**name** | **str** | Name of the referenced resource. | [optional] 
**url** | **str** | URL of the referenced resource. | [optional] 
**desc** | **str** | Brief description of the resource. | [optional] 
**abbreviation** | **str** | Abbreviation/initials/acronym for the alignment. | [optional] 

## Example

```python
from openapi_client.models.alignment import Alignment

# TODO update the JSON string below
json = "{}"
# create an instance of Alignment from a JSON string
alignment_instance = Alignment.from_json(json)
# print the JSON string representation of the object
print Alignment.to_json()

# convert the object into a dict
alignment_dict = alignment_instance.to_dict()
# create an instance of Alignment from a dict
alignment_form_dict = alignment.from_dict(alignment_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


