# SpellPrerequisite

`SpellPrerequisite` 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **str** | Resource index for shorthand searching. | [optional] 
**name** | **str** | Name of the referenced resource. | [optional] 
**url** | **str** | URL of the referenced resource. | [optional] 
**type** | **str** | The type of prerequisite. | [optional] 

## Example

```python
from dnd5epy.models.spell_prerequisite import SpellPrerequisite

# TODO update the JSON string below
json = "{}"
# create an instance of SpellPrerequisite from a JSON string
spell_prerequisite_instance = SpellPrerequisite.from_json(json)
# print the JSON string representation of the object
print SpellPrerequisite.to_json()

# convert the object into a dict
spell_prerequisite_dict = spell_prerequisite_instance.to_dict()
# create an instance of SpellPrerequisite from a dict
spell_prerequisite_form_dict = spell_prerequisite.from_dict(spell_prerequisite_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


