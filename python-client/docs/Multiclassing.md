# Multiclassing

`Multiclassing` 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**prerequisites** | [**List[Prerequisite]**](Prerequisite.md) | List of prerequisites that must be met. | [optional] 
**prerequisite_options** | [**List[Choice]**](Choice.md) | List of choices of prerequisites to meet for. | [optional] 
**proficiencies** | [**List[APIReference]**](APIReference.md) | List of proficiencies available when multiclassing. | [optional] 
**proficiency_choices** | [**List[Choice]**](Choice.md) | List of choices of proficiencies that are given when multiclassing. | [optional] 

## Example

```python
from openapi_client.models.multiclassing import Multiclassing

# TODO update the JSON string below
json = "{}"
# create an instance of Multiclassing from a JSON string
multiclassing_instance = Multiclassing.from_json(json)
# print the JSON string representation of the object
print Multiclassing.to_json()

# convert the object into a dict
multiclassing_dict = multiclassing_instance.to_dict()
# create an instance of Multiclassing from a dict
multiclassing_form_dict = multiclassing.from_dict(multiclassing_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


