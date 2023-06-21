# APIReferenceList

`APIReferenceList` 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**count** | **float** | Total number of resources available. | [optional] 
**results** | [**List[APIReference]**](APIReference.md) |  | [optional] 

## Example

```python
from dnd5epy.models.api_reference_list import APIReferenceList

# TODO update the JSON string below
json = "{}"
# create an instance of APIReferenceList from a JSON string
api_reference_list_instance = APIReferenceList.from_json(json)
# print the JSON string representation of the object
print APIReferenceList.to_json()

# convert the object into a dict
api_reference_list_dict = api_reference_list_instance.to_dict()
# create an instance of APIReferenceList from a dict
api_reference_list_form_dict = api_reference_list.from_dict(api_reference_list_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


