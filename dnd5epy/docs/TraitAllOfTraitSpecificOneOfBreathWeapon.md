# TraitAllOfTraitSpecificOneOfBreathWeapon

The breath weapon action associated with a draconic ancestry.

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** |  | [optional] 
**desc** | **str** |  | [optional] 
**area_of_effect** | [**AreaOfEffect**](AreaOfEffect.md) |  | [optional] 
**damage** | [**TraitAllOfTraitSpecificOneOfBreathWeaponDamage**](TraitAllOfTraitSpecificOneOfBreathWeaponDamage.md) |  | [optional] 
**dc** | [**DC**](DC.md) |  | [optional] 
**usage** | [**TraitAllOfTraitSpecificOneOfBreathWeaponUsage**](TraitAllOfTraitSpecificOneOfBreathWeaponUsage.md) |  | [optional] 

## Example

```python
from dnd5epy.models.trait_all_of_trait_specific_one_of_breath_weapon import TraitAllOfTraitSpecificOneOfBreathWeapon

# TODO update the JSON string below
json = "{}"
# create an instance of TraitAllOfTraitSpecificOneOfBreathWeapon from a JSON string
trait_all_of_trait_specific_one_of_breath_weapon_instance = TraitAllOfTraitSpecificOneOfBreathWeapon.from_json(json)
# print the JSON string representation of the object
print TraitAllOfTraitSpecificOneOfBreathWeapon.to_json()

# convert the object into a dict
trait_all_of_trait_specific_one_of_breath_weapon_dict = trait_all_of_trait_specific_one_of_breath_weapon_instance.to_dict()
# create an instance of TraitAllOfTraitSpecificOneOfBreathWeapon from a dict
trait_all_of_trait_specific_one_of_breath_weapon_form_dict = trait_all_of_trait_specific_one_of_breath_weapon.from_dict(trait_all_of_trait_specific_one_of_breath_weapon_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


