# SubraceAllOfRace

Parent race for the subrace.

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **str** | Resource index for shorthand searching. | [optional] 
**name** | **str** | Name of the referenced resource. | [optional] 
**url** | **str** | URL of the referenced resource. | [optional] 

## Example

```python
from openapi_client.models.subrace_all_of_race import SubraceAllOfRace

# TODO update the JSON string below
json = "{}"
# create an instance of SubraceAllOfRace from a JSON string
subrace_all_of_race_instance = SubraceAllOfRace.from_json(json)
# print the JSON string representation of the object
print SubraceAllOfRace.to_json()

# convert the object into a dict
subrace_all_of_race_dict = subrace_all_of_race_instance.to_dict()
# create an instance of SubraceAllOfRace from a dict
subrace_all_of_race_form_dict = subrace_all_of_race.from_dict(subrace_all_of_race_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


