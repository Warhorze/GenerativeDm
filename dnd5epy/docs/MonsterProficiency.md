# MonsterProficiency


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**value** | **float** |  | [optional] 
**proficiency** | [**APIReference**](APIReference.md) |  | [optional] 

## Example

```python
from dnd5epy.models.monster_proficiency import MonsterProficiency

# TODO update the JSON string below
json = "{}"
# create an instance of MonsterProficiency from a JSON string
monster_proficiency_instance = MonsterProficiency.from_json(json)
# print the JSON string representation of the object
print MonsterProficiency.to_json()

# convert the object into a dict
monster_proficiency_dict = monster_proficiency_instance.to_dict()
# create an instance of MonsterProficiency from a dict
monster_proficiency_form_dict = monster_proficiency.from_dict(monster_proficiency_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


