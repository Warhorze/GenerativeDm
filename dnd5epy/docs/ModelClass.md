# ModelClass

`Class` 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **str** | Resource index for shorthand searching. | [optional] 
**name** | **str** | Name of the referenced resource. | [optional] 
**url** | **str** | URL of the referenced resource. | [optional] 
**hit_die** | **float** | Hit die of the class. (ex: 12 &#x3D;&#x3D; 1d12). | [optional] 
**class_levels** | **str** | URL of the level resource for the class. | [optional] 
**multi_classing** | [**Multiclassing**](Multiclassing.md) |  | [optional] 
**spellcasting** | [**Spellcasting**](Spellcasting.md) |  | [optional] 
**spells** | **str** | URL of the spell resource list for the class. | [optional] 
**starting_equipment** | [**List[ClassAllOfStartingEquipment]**](ClassAllOfStartingEquipment.md) | List of equipment and their quantities all players of the class start with. | [optional] 
**starting_equipment_options** | [**List[Choice]**](Choice.md) | List of choices of starting equipment. | [optional] 
**proficiency_choices** | [**List[Choice]**](Choice.md) | List of choices of starting proficiencies. | [optional] 
**proficiencies** | [**List[APIReference]**](APIReference.md) | List of starting proficiencies for all new characters of this class. | [optional] 
**saving_throws** | [**List[APIReference]**](APIReference.md) | Saving throws the class is proficient in. | [optional] 
**subclasses** | [**List[APIReference]**](APIReference.md) | List of all possible subclasses this class can specialize in. | [optional] 

## Example

```python
from dnd5epy.models.model_class import ModelClass

# TODO update the JSON string below
json = "{}"
# create an instance of ModelClass from a JSON string
model_class_instance = ModelClass.from_json(json)
# print the JSON string representation of the object
print ModelClass.to_json()

# convert the object into a dict
model_class_dict = model_class_instance.to_dict()
# create an instance of ModelClass from a dict
model_class_form_dict = model_class.from_dict(model_class_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


