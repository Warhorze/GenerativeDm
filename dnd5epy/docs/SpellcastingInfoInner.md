# SpellcastingInfoInner


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Feature name. | [optional] 
**desc** | **List[str]** | Feature description. | [optional] 

## Example

```python
from dnd5epy.models.spellcasting_info_inner import SpellcastingInfoInner

# TODO update the JSON string below
json = "{}"
# create an instance of SpellcastingInfoInner from a JSON string
spellcasting_info_inner_instance = SpellcastingInfoInner.from_json(json)
# print the JSON string representation of the object
print SpellcastingInfoInner.to_json()

# convert the object into a dict
spellcasting_info_inner_dict = spellcasting_info_inner_instance.to_dict()
# create an instance of SpellcastingInfoInner from a dict
spellcasting_info_inner_form_dict = spellcasting_info_inner.from_dict(spellcasting_info_inner_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


