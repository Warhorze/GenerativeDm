# MonsterUsage


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**type** | **str** |  | [optional] 
**rest_types** | **List[str]** |  | [optional] 
**times** | **float** |  | [optional] 

## Example

```python
from dnd5epy.models.monster_usage import MonsterUsage

# TODO update the JSON string below
json = "{}"
# create an instance of MonsterUsage from a JSON string
monster_usage_instance = MonsterUsage.from_json(json)
# print the JSON string representation of the object
print MonsterUsage.to_json()

# convert the object into a dict
monster_usage_dict = monster_usage_instance.to_dict()
# create an instance of MonsterUsage from a dict
monster_usage_form_dict = monster_usage.from_dict(monster_usage_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


