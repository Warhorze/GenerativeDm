# TraitAllOfTraitSpecific

Information specific to this trait

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**desc** | **str** | Description of the choice to be made. | [optional] 
**choose** | **float** | Number of items to pick from the list. | [optional] 
**type** | **str** | Type of the resources to choose from. | [optional] 
**var_from** | [**OptionSet**](OptionSet.md) |  | [optional] 
**damage_type** | [**TraitAllOfTraitSpecificOneOfDamageType**](TraitAllOfTraitSpecificOneOfDamageType.md) |  | [optional] 
**breath_weapon** | [**TraitAllOfTraitSpecificOneOfBreathWeapon**](TraitAllOfTraitSpecificOneOfBreathWeapon.md) |  | [optional] 

## Example

```python
from openapi_client.models.trait_all_of_trait_specific import TraitAllOfTraitSpecific

# TODO update the JSON string below
json = "{}"
# create an instance of TraitAllOfTraitSpecific from a JSON string
trait_all_of_trait_specific_instance = TraitAllOfTraitSpecific.from_json(json)
# print the JSON string representation of the object
print TraitAllOfTraitSpecific.to_json()

# convert the object into a dict
trait_all_of_trait_specific_dict = trait_all_of_trait_specific_instance.to_dict()
# create an instance of TraitAllOfTraitSpecific from a dict
trait_all_of_trait_specific_form_dict = trait_all_of_trait_specific.from_dict(trait_all_of_trait_specific_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


