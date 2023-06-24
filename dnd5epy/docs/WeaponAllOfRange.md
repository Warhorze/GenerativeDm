# WeaponAllOfRange


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**normal** | **float** | The weapon&#39;s normal range in feet. | [optional] 
**long** | **float** | The weapon&#39;s long range in feet. | [optional] 

## Example

```python
from dnd5epy.models.weapon_all_of_range import WeaponAllOfRange

# TODO update the JSON string below
json = "{}"
# create an instance of WeaponAllOfRange from a JSON string
weapon_all_of_range_instance = WeaponAllOfRange.from_json(json)
# print the JSON string representation of the object
print WeaponAllOfRange.to_json()

# convert the object into a dict
weapon_all_of_range_dict = weapon_all_of_range_instance.to_dict()
# create an instance of WeaponAllOfRange from a dict
weapon_all_of_range_form_dict = weapon_all_of_range.from_dict(weapon_all_of_range_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


