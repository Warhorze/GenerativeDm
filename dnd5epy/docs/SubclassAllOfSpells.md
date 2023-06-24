# SubclassAllOfSpells


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**prerequisites** | [**List[SpellPrerequisite]**](SpellPrerequisite.md) |  | [optional] 
**spell** | [**APIReference**](APIReference.md) |  | [optional] 

## Example

```python
from dnd5epy.models.subclass_all_of_spells import SubclassAllOfSpells

# TODO update the JSON string below
json = "{}"
# create an instance of SubclassAllOfSpells from a JSON string
subclass_all_of_spells_instance = SubclassAllOfSpells.from_json(json)
# print the JSON string representation of the object
print SubclassAllOfSpells.to_json()

# convert the object into a dict
subclass_all_of_spells_dict = subclass_all_of_spells_instance.to_dict()
# create an instance of SubclassAllOfSpells from a dict
subclass_all_of_spells_form_dict = subclass_all_of_spells.from_dict(subclass_all_of_spells_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


