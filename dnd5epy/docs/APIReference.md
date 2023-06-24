# APIReference

`APIReference` 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **str** | Resource index for shorthand searching. | [optional] 
**name** | **str** | Name of the referenced resource. | [optional] 
**url** | **str** | URL of the referenced resource. | [optional] 

## Example

```python
from openapi_client.models.api_reference import APIReference

# TODO update the JSON string below
json = "{}"
# create an instance of APIReference from a JSON string
api_reference_instance = APIReference.from_json(json)
# print the JSON string representation of the object
print APIReference.to_json()

# convert the object into a dict
api_reference_dict = api_reference_instance.to_dict()
# create an instance of APIReference from a dict
api_reference_form_dict = api_reference.from_dict(api_reference_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


