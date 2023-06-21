# Proficiency

`Proficiency` 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **str** | Resource index for shorthand searching. | [optional] 
**name** | **str** | Name of the referenced resource. | [optional] 
**url** | **str** | URL of the referenced resource. | [optional] 
**type** | **str** | The general category of the proficiency. | [optional] 
**classes** | [**List[APIReference]**](APIReference.md) | Classes that start with this proficiency. | [optional] 
**races** | [**List[APIReference]**](APIReference.md) | Races that start with this proficiency. | [optional] 
**reference** | [**ProficiencyAllOfReference**](ProficiencyAllOfReference.md) |  | [optional] 

## Example

```python
from openapi_client.models.proficiency import Proficiency

# TODO update the JSON string below
json = "{}"
# create an instance of Proficiency from a JSON string
proficiency_instance = Proficiency.from_json(json)
# print the JSON string representation of the object
print Proficiency.to_json()

# convert the object into a dict
proficiency_dict = proficiency_instance.to_dict()
# create an instance of Proficiency from a dict
proficiency_form_dict = proficiency.from_dict(proficiency_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


