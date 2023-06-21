# ResourceDescription


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**desc** | **List[str]** | Description of the resource. | [optional] 

## Example

```python
from dnd5epy.models.resource_description import ResourceDescription

# TODO update the JSON string below
json = "{}"
# create an instance of ResourceDescription from a JSON string
resource_description_instance = ResourceDescription.from_json(json)
# print the JSON string representation of the object
print ResourceDescription.to_json()

# convert the object into a dict
resource_description_dict = resource_description_instance.to_dict()
# create an instance of ResourceDescription from a dict
resource_description_form_dict = resource_description.from_dict(resource_description_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


