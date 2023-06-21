# Monster

`Monster` 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **str** | Resource index for shorthand searching. | [optional] 
**name** | **str** | Name of the referenced resource. | [optional] 
**url** | **str** | URL of the referenced resource. | [optional] 
**desc** | **List[str]** | Description of the resource. | [optional] 
**charisma** | **float** | A monster&#39;s ability to charm or intimidate a player. | [optional] 
**constitution** | **float** | How sturdy a monster is.\&quot; | [optional] 
**dexterity** | **float** | The monster&#39;s ability for swift movement or stealth | [optional] 
**intelligence** | **float** | The monster&#39;s ability to outsmart a player. | [optional] 
**strength** | **float** | How hard a monster can hit a player. | [optional] 
**wisdom** | **float** | A monster&#39;s ability to ascertain the player&#39;s plan. | [optional] 
**image** | **str** | The image url of the monster. | [optional] 
**size** | **str** | The size of the monster ranging from Tiny to Gargantuan.\&quot; | [optional] 
**type** | **str** | The type of monster. | [optional] 
**subtype** | **str** | The sub-category of a monster used for classification of monsters.\&quot; | [optional] 
**alignments** | **str** | A creature&#39;s general moral and personal attitudes. | [optional] 
**armor_class** | [**List[MonsterArmorClass]**](MonsterArmorClass.md) | The difficulty for a player to successfully deal damage to a monster. | [optional] 
**hit_points** | **float** | The hit points of a monster determine how much damage it is able to take before it can be defeated. | [optional] 
**hit_dice** | **str** | The hit die of a monster can be used to make a version of the same monster whose hit points are determined by the roll of the die. For example: A monster with 2d6 would have its hit points determine by rolling a 6 sided die twice. | [optional] 
**hit_points_roll** | **str** | The roll for determining a monster&#39;s hit points, which consists of the hit dice (e.g. 18d10) and the modifier determined by its Constitution (e.g. +36). For example, 18d10+36 | [optional] 
**actions** | [**List[MonsterAction]**](MonsterAction.md) | A list of actions that are available to the monster to take during combat. | [optional] 
**legendary_actions** | [**List[MonsterAction]**](MonsterAction.md) | A list of legendary actions that are available to the monster to take during combat. | [optional] 
**challenge_rating** | **float** | A monster&#39;s challenge rating is a guideline number that says when a monster becomes an appropriate challenge against the party&#39;s average level. For example. A group of 4 players with an average level of 4 would have an appropriate combat challenge against a monster with a challenge rating of 4 but a monster with a challenge rating of 8 against the same group of players would pose a significant threat. | [optional] 
**condition_immunities** | [**List[APIReference]**](APIReference.md) | A list of conditions that a monster is immune to. | [optional] 
**damage_immunities** | **List[str]** | A list of damage types that a monster will take double damage from. | [optional] 
**damage_resistances** | **List[str]** | A list of damage types that a monster will take half damage from. | [optional] 
**damage_vulnerabilities** | **List[str]** | A list of damage types that a monster will take double damage from. | [optional] 
**forms** | [**List[APIReference]**](APIReference.md) | List of other related monster entries that are of the same form. Only applicable to Lycanthropes that have multiple forms. | [optional] 
**languages** | **str** | The languages a monster is able to speak. | [optional] 
**proficiencies** | [**List[MonsterProficiency]**](MonsterProficiency.md) | A list of proficiencies of a monster. | [optional] 
**reactions** | [**List[MonsterAction]**](MonsterAction.md) | A list of reactions that is available to the monster to take during combat. | [optional] 
**senses** | [**MonsterAllOfSenses**](MonsterAllOfSenses.md) |  | [optional] 
**special_abilities** | [**List[MonsterSpecialAbility]**](MonsterSpecialAbility.md) | A list of the monster&#39;s special abilities. | [optional] 
**speed** | [**MonsterAllOfSpeed**](MonsterAllOfSpeed.md) |  | [optional] 
**xp** | **float** | The number of experience points (XP) a monster is worth is based on its challenge rating. | [optional] 

## Example

```python
from dnd5epy.models.monster import Monster

# TODO update the JSON string below
json = "{}"
# create an instance of Monster from a JSON string
monster_instance = Monster.from_json(json)
# print the JSON string representation of the object
print Monster.to_json()

# convert the object into a dict
monster_dict = monster_instance.to_dict()
# create an instance of Monster from a dict
monster_form_dict = monster.from_dict(monster_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


