# RuleSection

`RuleSection` 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **str** | Resource index for shorthand searching. | [optional] 
**name** | **str** | Name of the referenced resource. | [optional] 
**url** | **str** | URL of the referenced resource. | [optional] 
**desc** | **str** | Description of the rule. | [optional] 

## Example

```python
from dnd5epy.models.rule_section import RuleSection

# TODO update the JSON string below
json = "{}"
# create an instance of RuleSection from a JSON string
rule_section_instance = RuleSection.from_json(json)
# print the JSON string representation of the object
print RuleSection.to_json()

# convert the object into a dict
rule_section_dict = rule_section_instance.to_dict()
# create an instance of RuleSection from a dict
rule_section_form_dict = rule_section.from_dict(rule_section_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


