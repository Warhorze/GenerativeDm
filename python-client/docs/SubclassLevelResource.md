# SubclassLevelResource


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **str** |  | [optional] 
**url** | **str** |  | [optional] 
**level** | **float** |  | [optional] 
**features** | [**List[APIReference]**](APIReference.md) |  | [optional] 
**var_class** | [**APIReference**](APIReference.md) |  | [optional] 
**subclass** | [**APIReference**](APIReference.md) |  | [optional] 

## Example

```python
from openapi_client.models.subclass_level_resource import SubclassLevelResource

# TODO update the JSON string below
json = "{}"
# create an instance of SubclassLevelResource from a JSON string
subclass_level_resource_instance = SubclassLevelResource.from_json(json)
# print the JSON string representation of the object
print SubclassLevelResource.to_json()

# convert the object into a dict
subclass_level_resource_dict = subclass_level_resource_instance.to_dict()
# create an instance of SubclassLevelResource from a dict
subclass_level_resource_form_dict = subclass_level_resource.from_dict(subclass_level_resource_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


