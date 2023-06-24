# Race

`Race` 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **str** | Resource index for shorthand searching. | [optional] 
**name** | **str** | Name of the referenced resource. | [optional] 
**url** | **str** | URL of the referenced resource. | [optional] 
**speed** | **float** | Base move speed for this race (in feet per round). | [optional] 
**ability_bonuses** | [**List[AbilityBonus]**](AbilityBonus.md) | Racial bonuses to ability scores. | [optional] 
**alignment** | **str** | Flavor description of likely alignments this race takes. | [optional] 
**age** | **str** | Flavor description of possible ages for this race. | [optional] 
**size** | **str** | Size class of this race. | [optional] 
**size_description** | **str** | Flavor description of height and weight for this race. | [optional] 
**starting_proficiencies** | [**List[APIReference]**](APIReference.md) | Starting proficiencies for all new characters of this race. | [optional] 
**starting_proficiency_options** | [**Choice**](Choice.md) |  | [optional] 
**languages** | [**List[APIReference]**](APIReference.md) | Starting languages for all new characters of this race. | [optional] 
**language_desc** | **str** | Flavor description of the languages this race knows. | [optional] 
**traits** | [**List[APIReference]**](APIReference.md) | Racial traits that provide benefits to its members. | [optional] 
**subraces** | [**List[APIReference]**](APIReference.md) | All possible subraces that this race includes. | [optional] 

## Example

```python
from dnd5epy.models.race import Race

# TODO update the JSON string below
json = "{}"
# create an instance of Race from a JSON string
race_instance = Race.from_json(json)
# print the JSON string representation of the object
print Race.to_json()

# convert the object into a dict
race_dict = race_instance.to_dict()
# create an instance of Race from a dict
race_form_dict = race.from_dict(race_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


