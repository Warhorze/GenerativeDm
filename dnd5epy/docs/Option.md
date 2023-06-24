# Option

`Option` 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**option_type** | **str** | Type of option; determines other attributes. | [optional] 
**item** | [**APIReference**](APIReference.md) |  | [optional] 
**action_name** | **str** | The name of the action. | [optional] 
**count** | **float** | Count | [optional] 
**type** | **str** | For attack options that can be melee, ranged, abilities, or thrown. | [optional] 
**items** | [**List[Option]**](Option.md) |  | [optional] 
**choice** | [**Choice**](Choice.md) |  | [optional] 
**string** | **str** | The string. | [optional] 
**desc** | **str** | A description of the ideal. | [optional] 
**alignments** | [**List[APIReference]**](APIReference.md) | A list of alignments of those who might follow the ideal. | [optional] 
**of** | [**APIReference**](APIReference.md) |  | [optional] 
**ability_score** | [**APIReference**](APIReference.md) |  | [optional] 
**minimum_score** | **float** | The minimum score required to satisfy the prerequisite. | [optional] 
**bonus** | **float** | The bonus being applied to the ability score | [optional] 
**name** | **str** | Name of the breath | [optional] 
**dc** | [**DC**](DC.md) |  | [optional] 
**damage** | [**List[Damage]**](Damage.md) | Damage dealt by the breath attack, if any. | [optional] 
**damage_type** | [**APIReference**](APIReference.md) |  | [optional] 
**damage_dice** | **str** | Damage expressed in dice (e.g. \&quot;13d6\&quot;). | [optional] 
**notes** | **str** | Information regarding the damage. | [optional] 

## Example

```python
from openapi_client.models.option import Option

# TODO update the JSON string below
json = "{}"
# create an instance of Option from a JSON string
option_instance = Option.from_json(json)
# print the JSON string representation of the object
print Option.to_json()

# convert the object into a dict
option_dict = option_instance.to_dict()
# create an instance of Option from a dict
option_form_dict = option.from_dict(option_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


