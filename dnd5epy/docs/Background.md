# Background

`Background` 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **str** | Resource index for shorthand searching. | [optional] 
**name** | **str** | Name of the referenced resource. | [optional] 
**url** | **str** | URL of the referenced resource. | [optional] 
**starting_proficiencies** | [**List[APIReference]**](APIReference.md) | Starting proficiencies for all new characters of this background. | [optional] 
**starting_equipment** | [**List[APIReference]**](APIReference.md) | Starting equipment for all new characters of this background. | [optional] 
**starting_equipment_options** | [**Choice**](Choice.md) |  | [optional] 
**language_options** | [**Choice**](Choice.md) |  | [optional] 
**feature** | [**BackgroundAllOfFeature**](BackgroundAllOfFeature.md) |  | [optional] 
**personality_traits** | **object** | Choice of personality traits for this background. | [optional] 
**ideals** | [**Choice**](Choice.md) |  | [optional] 
**bonds** | [**Choice**](Choice.md) |  | [optional] 
**flaws** | [**Choice**](Choice.md) |  | [optional] 

## Example

```python
from dnd5epy.models.background import Background

# TODO update the JSON string below
json = "{}"
# create an instance of Background from a JSON string
background_instance = Background.from_json(json)
# print the JSON string representation of the object
print Background.to_json()

# convert the object into a dict
background_dict = background_instance.to_dict()
# create an instance of Background from a dict
background_form_dict = background.from_dict(background_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


