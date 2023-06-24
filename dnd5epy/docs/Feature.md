# Feature

`Feature` 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **str** | Resource index for shorthand searching. | [optional] 
**name** | **str** | Name of the referenced resource. | [optional] 
**url** | **str** | URL of the referenced resource. | [optional] 
**desc** | **List[str]** | Description of the resource. | [optional] 
**level** | **float** | The level this feature is gained. | [optional] 
**var_class** | [**APIReference**](APIReference.md) |  | [optional] 
**subclass** | [**APIReference**](APIReference.md) |  | [optional] 
**parent** | [**APIReference**](APIReference.md) |  | [optional] 
**prerequisites** | [**List[FeatureAllOfPrerequisitesInner]**](FeatureAllOfPrerequisitesInner.md) | The prerequisites for this feature. | [optional] 
**feature_specific** | **Dict[str, object]** | Information specific to this feature. | [optional] 

## Example

```python
from dnd5epy.models.feature import Feature

# TODO update the JSON string below
json = "{}"
# create an instance of Feature from a JSON string
feature_instance = Feature.from_json(json)
# print the JSON string representation of the object
print Feature.to_json()

# convert the object into a dict
feature_dict = feature_instance.to_dict()
# create an instance of Feature from a dict
feature_form_dict = feature.from_dict(feature_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


