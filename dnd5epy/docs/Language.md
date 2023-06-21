# Language

`Language` 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **str** | Resource index for shorthand searching. | [optional] 
**name** | **str** | Name of the referenced resource. | [optional] 
**url** | **str** | URL of the referenced resource. | [optional] 
**desc** | **str** | Brief description of the language. | [optional] 
**type** | **str** |  | [optional] 
**script** | **str** | Script used for writing in the language. | [optional] 
**typical_speakers** | **List[str]** | List of races that tend to speak the language. | [optional] 

## Example

```python
from dnd5epy.models.language import Language

# TODO update the JSON string below
json = "{}"
# create an instance of Language from a JSON string
language_instance = Language.from_json(json)
# print the JSON string representation of the object
print Language.to_json()

# convert the object into a dict
language_dict = language_instance.to_dict()
# create an instance of Language from a dict
language_form_dict = language.from_dict(language_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


