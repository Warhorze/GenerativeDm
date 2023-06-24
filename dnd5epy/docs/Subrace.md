# Subrace

`Subrace` 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **str** | Resource index for shorthand searching. | [optional] 
**name** | **str** | Name of the referenced resource. | [optional] 
**url** | **str** | URL of the referenced resource. | [optional] 
**desc** | **str** | Description of the subrace. | [optional] 
**race** | [**SubraceAllOfRace**](SubraceAllOfRace.md) |  | [optional] 
**ability_bonuses** | [**List[AbilityBonus]**](AbilityBonus.md) | Additional ability bonuses for the subrace. | [optional] 
**starting_proficiencies** | [**List[APIReference]**](APIReference.md) | Starting proficiencies for all new characters of the subrace. | [optional] 
**languages** | [**List[APIReference]**](APIReference.md) | Starting languages for all new characters of the subrace. | [optional] 
**language_options** | [**Choice**](Choice.md) |  | [optional] 
**racial_traits** | [**List[APIReference]**](APIReference.md) | List of traits that for the subrace. | [optional] 

## Example

```python
from openapi_client.models.subrace import Subrace

# TODO update the JSON string below
json = "{}"
# create an instance of Subrace from a JSON string
subrace_instance = Subrace.from_json(json)
# print the JSON string representation of the object
print Subrace.to_json()

# convert the object into a dict
subrace_dict = subrace_instance.to_dict()
# create an instance of Subrace from a dict
subrace_form_dict = subrace.from_dict(subrace_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


