# Feat

`Feat` 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **str** | Resource index for shorthand searching. | [optional] 
**name** | **str** | Name of the referenced resource. | [optional] 
**url** | **str** | URL of the referenced resource. | [optional] 
**desc** | **List[str]** | Description of the resource. | [optional] 
**prerequisites** | [**List[Prerequisite]**](Prerequisite.md) | An object of APIReferences to ability scores and minimum scores. | [optional] 

## Example

```python
from dnd5epy.models.feat import Feat

# TODO update the JSON string below
json = "{}"
# create an instance of Feat from a JSON string
feat_instance = Feat.from_json(json)
# print the JSON string representation of the object
print Feat.to_json()

# convert the object into a dict
feat_dict = feat_instance.to_dict()
# create an instance of Feat from a dict
feat_form_dict = feat.from_dict(feat_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


