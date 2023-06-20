# Trait

`Trait` 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **str** | Resource index for shorthand searching. | [optional] 
**name** | **str** | Name of the referenced resource. | [optional] 
**url** | **str** | URL of the referenced resource. | [optional] 
**desc** | **List[str]** | Description of the resource. | [optional] 
**races** | [**List[APIReference]**](APIReference.md) | List of &#x60;Races&#x60; that have access to the trait. | [optional] 
**subraces** | [**List[APIReference]**](APIReference.md) | List of &#x60;Subraces&#x60; that have access to the trait. | [optional] 
**proficiencies** | [**List[APIReference]**](APIReference.md) | List of &#x60;Proficiencies&#x60; this trait grants. | [optional] 
**proficiency_choices** | [**Choice**](Choice.md) |  | [optional] 
**language_options** | [**Choice**](Choice.md) |  | [optional] 
**trait_specific** | [**TraitAllOfTraitSpecific**](TraitAllOfTraitSpecific.md) |  | [optional] 

## Example

```python
from openapi_client.models.trait import Trait

# TODO update the JSON string below
json = "{}"
# create an instance of Trait from a JSON string
trait_instance = Trait.from_json(json)
# print the JSON string representation of the object
print Trait.to_json()

# convert the object into a dict
trait_dict = trait_instance.to_dict()
# create an instance of Trait from a dict
trait_form_dict = trait.from_dict(trait_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


