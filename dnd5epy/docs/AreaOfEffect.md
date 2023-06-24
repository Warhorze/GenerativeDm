# AreaOfEffect


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**size** | **float** |  | [optional] 
**type** | **str** |  | [optional] 

## Example

```python
from dnd5epy.models.area_of_effect import AreaOfEffect

# TODO update the JSON string below
json = "{}"
# create an instance of AreaOfEffect from a JSON string
area_of_effect_instance = AreaOfEffect.from_json(json)
# print the JSON string representation of the object
print AreaOfEffect.to_json()

# convert the object into a dict
area_of_effect_dict = area_of_effect_instance.to_dict()
# create an instance of AreaOfEffect from a dict
area_of_effect_form_dict = area_of_effect.from_dict(area_of_effect_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


