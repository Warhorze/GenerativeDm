# FeatureAllOfPrerequisitesInner


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**type** | **str** |  | [optional] 
**level** | **float** |  | [optional] 
**feature** | **str** |  | [optional] 
**spell** | **str** |  | [optional] 

## Example

```python
from dnd5epy.models.feature_all_of_prerequisites_inner import FeatureAllOfPrerequisitesInner

# TODO update the JSON string below
json = "{}"
# create an instance of FeatureAllOfPrerequisitesInner from a JSON string
feature_all_of_prerequisites_inner_instance = FeatureAllOfPrerequisitesInner.from_json(json)
# print the JSON string representation of the object
print FeatureAllOfPrerequisitesInner.to_json()

# convert the object into a dict
feature_all_of_prerequisites_inner_dict = feature_all_of_prerequisites_inner_instance.to_dict()
# create an instance of FeatureAllOfPrerequisitesInner from a dict
feature_all_of_prerequisites_inner_form_dict = feature_all_of_prerequisites_inner.from_dict(feature_all_of_prerequisites_inner_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


