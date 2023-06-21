# ProficiencyAllOfReference

`APIReference` to the full description of the related resource. 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **str** | Resource index for shorthand searching. | [optional] 
**name** | **str** | Name of the referenced resource. | [optional] 
**url** | **str** | URL of the referenced resource. | [optional] 

## Example

```python
from dnd5epy.models.proficiency_all_of_reference import ProficiencyAllOfReference

# TODO update the JSON string below
json = "{}"
# create an instance of ProficiencyAllOfReference from a JSON string
proficiency_all_of_reference_instance = ProficiencyAllOfReference.from_json(json)
# print the JSON string representation of the object
print ProficiencyAllOfReference.to_json()

# convert the object into a dict
proficiency_all_of_reference_dict = proficiency_all_of_reference_instance.to_dict()
# create an instance of ProficiencyAllOfReference from a dict
proficiency_all_of_reference_form_dict = proficiency_all_of_reference.from_dict(proficiency_all_of_reference_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


