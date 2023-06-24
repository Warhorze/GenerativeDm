# DC

`DC` 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**dc_type** | [**APIReference**](APIReference.md) |  | [optional] 
**dc_value** | **float** | Value to beat | [optional] 
**success_type** | **str** | Result of a successful save. Can be \\\&quot;none\\\&quot;, \\\&quot;half\\\&quot;, or \\\&quot;other\\\&quot; | [optional] 

## Example

```python
from openapi_client.models.dc import DC

# TODO update the JSON string below
json = "{}"
# create an instance of DC from a JSON string
dc_instance = DC.from_json(json)
# print the JSON string representation of the object
print DC.to_json()

# convert the object into a dict
dc_dict = dc_instance.to_dict()
# create an instance of DC from a dict
dc_form_dict = dc.from_dict(dc_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


