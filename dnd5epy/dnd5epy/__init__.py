# coding: utf-8

# flake8: noqa

"""
    D&D 5e API

    # Introduction  Welcome to the dnd5eapi, the Dungeons & Dragons 5th Edition API! This documentation should help you familiarize yourself with the resources available and how to consume them with HTTP requests. Read through the getting started section before you dive in. Most of your problems should be solved just by reading through it.  ## Getting Started  Let's make our first API request to the D&D 5th Edition API!  Open up a terminal and use [curl](http://curl.haxx.se/) or [httpie](http://httpie.org/) to make an API request for a resource. You can also scroll through the definitions below and send requests directly from the endpoint documentation!  For example, if you paste and run this `curl` command: ```bash curl -X GET \"https://www.dnd5eapi.co/api/ability-scores/cha\" -H \"Accept: application/json\" ```  We should see a result containing details about the Charisma ability score: ```bash {   \"index\": \"cha\",   \"name\": \"CHA\",   \"full_name\": \"Charisma\",   \"desc\": [     \"Charisma measures your ability to interact effectively with others. It       includes such factors as confidence and eloquence, and it can represent       a charming or commanding personality.\",     \"A Charisma check might arise when you try to influence or entertain       others, when you try to make an impression or tell a convincing lie,       or when you are navigating a tricky social situation. The Deception,       Intimidation, Performance, and Persuasion skills reflect aptitude in       certain kinds of Charisma checks.\"   ],   \"skills\": [     {       \"name\": \"Deception\",       \"index\": \"deception\",       \"url\": \"/api/skills/deception\"     },     {       \"name\": \"Intimidation\",       \"index\": \"intimidation\",       \"url\": \"/api/skills/intimidation\"     },     {       \"name\": \"Performance\",       \"index\": \"performance\",       \"url\": \"/api/skills/performance\"     },     {       \"name\": \"Persuasion\",       \"index\": \"persuasion\",       \"url\": \"/api/skills/persuasion\"     }   ],   \"url\": \"/api/ability-scores/cha\" } ```  ## Authentication  The dnd5eapi is a completely open API. No authentication is required to query and get data. This also means that we've limited what you can do to just `GET`-ing the data. If you find a mistake in the data, feel free to [message us](https://discord.gg/TQuYTv7).  ## GraphQL  This API supports [GraphQL](https://graphql.org/). The GraphQL URL for this API is `https://www.dnd5eapi.co/graphql`. Most of your questions regarding the GraphQL schema can be answered by querying the endpoint with the Apollo sandbox explorer.  ## Schemas  Definitions of all schemas will be accessible in a future update. Two of the most common schemas are described here.  ### `APIReference` Represents a minimal representation of a resource. The detailed representation of the referenced resource can be retrieved by making a request to the referenced `URL`. ``` APIReference {   index     string   name      string   url       string } ``` <hr>  ### `DC` Represents a difficulty check. ``` DC {   dc_type       APIReference   dc_value      number   success_type  \"none\" | \"half\" | \"other\" } ``` <hr>  ### `Damage` Represents damage. ``` Damage {   damage_type     APIReference   damage_dice     string } ``` <hr>  ### `Choice` Represents a choice made by a player. Commonly seen related to decisions made during character creation or combat (e.g.: the description of the cleric class, under **Proficiencies**, states \"Skills: Choose two from History, Insight, Medicine, Persuasion, and Religion\" [[SRD p15]](https://media.wizards.com/2016/downloads/DND/SRD-OGL_V5.1.pdf#page=15)) ``` Choice {   desc      string   choose    number   type      string   from      OptionSet } ``` <hr>  ### `OptionSet` The OptionSet structure provides the options to be chosen from, or sufficient data to fetch and interpret the options. All OptionSets have an `option_set_type` attribute that indicates the structure of the object that contains the options. The possible values are `options_array`, `equipment_category`, and `reference_list`. Other attributes on the OptionSet depend on the value of this attribute. - `options_array`   - `options` (array): An array of Option objects. Each item in the array represents an option that can be chosen. - `equipment_category`   - `equipment_category` (APIReference): A reference to an EquipmentCategory. Each item in the EquipmentCategory's `equipment` array represents one option that can be chosen. - `resource_list`   - `resource_list_url` (string): A reference (by URL) to a collection in the database. The URL may include query parameters. Each item in the resulting ResourceList's `results` array represents one option that can be chosen. <hr>  ### `Option` When the options are given in an `options_array`, each item in the array inherits from the Option structure. All Options have an `option_type` attribute that indicates the structure of the option. The value of this attribute indicates how the option should be handled, and each type has different attributes. The possible values and their corresponding attributes are listed below. - `reference` - A terminal option. Contains a reference to a Document that can be added to the list of options chosen.   - `item` (APIReference): A reference to the chosen item. - `action` - A terminal option. Contains information describing an action, for use within Multiattack actions.   - `action_name` (string): The name of the action, according to its `name` attribute.   - `count` (number | string): The number of times this action can be repeated if this option is chosen.   - `type` (string = `\"melee\" | \"ranged\" | \"ability\" | \"magic\"`, optional): For attack actions that can be either melee, ranged, abilities, or magic. - `multiple` - When this option is chosen, all of its child options are chosen, and must be resolved the same way as a normal option.   - `items` (array): An array of Option objects. All of them must be taken if the option is chosen. - `choice` - A nested choice. If this option is chosen, the Choice structure contained within must be resolved like a normal Choice structure, and the results are the chosen options.   - `choice` (Choice): The Choice to resolve. - `string` - A terminal option. Contains a reference to a string.   - `string` (string): The string. - `ideal` - A terminal option. Contains information about an ideal.   - `desc` (string): A description of the ideal.   - `alignments` (ApiReference[]): A list of alignments of those who might follow the ideal. - `counted_reference` - A terminal option. Contains a reference to something else in the API along with a count.   - `count` (number): Count.   - `of` (ApiReference): Thing being referenced. - `score_prerequisite` - A terminal option. Contains a reference to an ability score and a minimum score.   - `ability_score` (ApiReference): Ability score being referenced.   - `minimum_score` (number): The minimum score required to satisfy the prerequisite. - `ability_bonus` - A terminal option. Contains a reference to an ability score and a bonus   - `ability_score` (ApiReference): Ability score being referenced   - `bonus` (number): The bonus being applied to the ability score - `breath` - A terminal option: Contains a reference to information about a breath attack.   - `name` (string): Name of the breath.   - `dc` (DC): Difficulty check of the breath attack.   - `damage` ([Damage]): Damage dealt by the breath attack, if any. - `damage` - A terminal option. Contains information about damage.   - `damage_type` (ApiReference): Reference to type of damage.   - `damage_dice` (string): Damage expressed in dice (e.g. \"13d6\").   - `notes` (string): Information regarding the damage.  ## FAQ  ### What is the SRD? The SRD, or Systems Reference Document, contains guidelines for publishing content under the OGL. This allows for some of the data for D&D 5e to be open source. The API only covers data that can be found in the SRD. [Here's a link to the full text of the SRD.](https://media.wizards.com/2016/downloads/DND/SRD-OGL_V5.1.pdf)  ### What is the OGL? The Open Game License (OGL) is a public copyright license by Wizards of the Coast that may be used by tabletop role-playing game developers to grant permission to modify, copy, and redistribute some of the content designed for their games, notably game mechanics. However, they must share-alike copies and derivative works. [More information about the OGL can be found here.](https://en.wikipedia.org/wiki/Open_Game_License)  ### A monster, spell, subclass, etc. is missing from the API / Database. Can I add it? Please check if the data is within the SRD. If it is, feel free to open an issue or PR to add it yourself. Otherwise, due to legal reasons, we cannot add it.  ### Can this API be self hosted? Yes it can! You can also host the data yourself if you don't want to use the API at all. You can also make changes and add extra data if you like. However, it is up to you to merge in new changes to the data and API.  #### Can I publish is on <insert platform>? Is this free use? Yes, you can. The API itself is under the [MIT license](https://opensource.org/licenses/MIT), and the underlying data accessible via the API is supported under the SRD and OGL.  # Status Page  The status page for the API can be found here: https://5e-bits.github.io/dnd-uptime/  # Chat  Come hang out with us [on Discord](https://discord.gg/TQuYTv7)!  # Contribute  This API is built from two repositories.   - The repo containing the data lives here: https://github.com/bagelbits/5e-database   - The repo with the API implementation lives here: https://github.com/bagelbits/5e-srd-api  This is a evolving API and having fresh ideas are always welcome! You can open an issue in either repo, open a PR for changes, or just discuss with other users in this discord.   # noqa: E501

    The version of the OpenAPI document: 0.1
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""


__version__ = "1.0.0"

# import apis into sdk package
from dnd5epy.api.character_data_api import CharacterDataApi
from dnd5epy.api.class_api import ClassApi
from dnd5epy.api.class_levels_api import ClassLevelsApi
from dnd5epy.api.class_resource_lists_api import ClassResourceListsApi
from dnd5epy.api.common_api import CommonApi
from dnd5epy.api.equipment_api import EquipmentApi
from dnd5epy.api.feats_api import FeatsApi
from dnd5epy.api.features_api import FeaturesApi
from dnd5epy.api.game_mechanics_api import GameMechanicsApi
from dnd5epy.api.monsters_api import MonstersApi
from dnd5epy.api.races_api import RacesApi
from dnd5epy.api.rules_api import RulesApi
from dnd5epy.api.spells_api import SpellsApi
from dnd5epy.api.subclasses_api import SubclassesApi
from dnd5epy.api.subraces_api import SubracesApi
from dnd5epy.api.traits_api import TraitsApi

# import ApiClient
from dnd5epy.api_response import ApiResponse
from dnd5epy.api_client import ApiClient
from dnd5epy.configuration import Configuration
from dnd5epy.exceptions import OpenApiException
from dnd5epy.exceptions import ApiTypeError
from dnd5epy.exceptions import ApiValueError
from dnd5epy.exceptions import ApiKeyError
from dnd5epy.exceptions import ApiAttributeError
from dnd5epy.exceptions import ApiException

# import models into sdk package
from dnd5epy.models.api_reference import APIReference
from dnd5epy.models.api_reference_list import APIReferenceList
from dnd5epy.models.ability_bonus import AbilityBonus
from dnd5epy.models.ability_score import AbilityScore
from dnd5epy.models.alignment import Alignment
from dnd5epy.models.area_of_effect import AreaOfEffect
from dnd5epy.models.armor import Armor
from dnd5epy.models.background import Background
from dnd5epy.models.background_all_of_feature import BackgroundAllOfFeature
from dnd5epy.models.choice import Choice
from dnd5epy.models.class_all_of_starting_equipment import ClassAllOfStartingEquipment
from dnd5epy.models.class_level import ClassLevel
from dnd5epy.models.class_level_class_specific import ClassLevelClassSpecific
from dnd5epy.models.class_level_class_specific_any_of import ClassLevelClassSpecificAnyOf
from dnd5epy.models.class_level_class_specific_any_of1 import ClassLevelClassSpecificAnyOf1
from dnd5epy.models.class_level_class_specific_any_of10 import ClassLevelClassSpecificAnyOf10
from dnd5epy.models.class_level_class_specific_any_of11 import ClassLevelClassSpecificAnyOf11
from dnd5epy.models.class_level_class_specific_any_of2 import ClassLevelClassSpecificAnyOf2
from dnd5epy.models.class_level_class_specific_any_of3 import ClassLevelClassSpecificAnyOf3
from dnd5epy.models.class_level_class_specific_any_of4 import ClassLevelClassSpecificAnyOf4
from dnd5epy.models.class_level_class_specific_any_of5 import ClassLevelClassSpecificAnyOf5
from dnd5epy.models.class_level_class_specific_any_of5_martial_arts import ClassLevelClassSpecificAnyOf5MartialArts
from dnd5epy.models.class_level_class_specific_any_of6 import ClassLevelClassSpecificAnyOf6
from dnd5epy.models.class_level_class_specific_any_of7 import ClassLevelClassSpecificAnyOf7
from dnd5epy.models.class_level_class_specific_any_of8 import ClassLevelClassSpecificAnyOf8
from dnd5epy.models.class_level_class_specific_any_of9 import ClassLevelClassSpecificAnyOf9
from dnd5epy.models.class_level_class_specific_any_of9_creating_spell_slots_inner import ClassLevelClassSpecificAnyOf9CreatingSpellSlotsInner
from dnd5epy.models.condition import Condition
from dnd5epy.models.cost import Cost
from dnd5epy.models.dc import DC
from dnd5epy.models.damage import Damage
from dnd5epy.models.damage_at_character_level import DamageAtCharacterLevel
from dnd5epy.models.damage_at_slot_level import DamageAtSlotLevel
from dnd5epy.models.damage_type import DamageType
from dnd5epy.models.equipment import Equipment
from dnd5epy.models.equipment_category import EquipmentCategory
from dnd5epy.models.equipment_pack import EquipmentPack
from dnd5epy.models.error_response import ErrorResponse
from dnd5epy.models.feat import Feat
from dnd5epy.models.feature import Feature
from dnd5epy.models.feature_all_of_prerequisites_inner import FeatureAllOfPrerequisitesInner
from dnd5epy.models.feature_all_of_prerequisites_inner_any_of import FeatureAllOfPrerequisitesInnerAnyOf
from dnd5epy.models.feature_all_of_prerequisites_inner_any_of1 import FeatureAllOfPrerequisitesInnerAnyOf1
from dnd5epy.models.feature_all_of_prerequisites_inner_any_of2 import FeatureAllOfPrerequisitesInnerAnyOf2
from dnd5epy.models.gear import Gear
from dnd5epy.models.language import Language
from dnd5epy.models.magic_item import MagicItem
from dnd5epy.models.magic_item_all_of_rarity import MagicItemAllOfRarity
from dnd5epy.models.magic_school import MagicSchool
from dnd5epy.models.model_class import ModelClass
from dnd5epy.models.monster import Monster
from dnd5epy.models.monster_ability import MonsterAbility
from dnd5epy.models.monster_action import MonsterAction
from dnd5epy.models.monster_all_of_senses import MonsterAllOfSenses
from dnd5epy.models.monster_all_of_speed import MonsterAllOfSpeed
from dnd5epy.models.monster_armor_class import MonsterArmorClass
from dnd5epy.models.monster_armor_class_one_of import MonsterArmorClassOneOf
from dnd5epy.models.monster_armor_class_one_of1 import MonsterArmorClassOneOf1
from dnd5epy.models.monster_armor_class_one_of2 import MonsterArmorClassOneOf2
from dnd5epy.models.monster_armor_class_one_of3 import MonsterArmorClassOneOf3
from dnd5epy.models.monster_armor_class_one_of4 import MonsterArmorClassOneOf4
from dnd5epy.models.monster_attack import MonsterAttack
from dnd5epy.models.monster_multi_attack_action import MonsterMultiAttackAction
from dnd5epy.models.monster_proficiency import MonsterProficiency
from dnd5epy.models.monster_sense import MonsterSense
from dnd5epy.models.monster_special_ability import MonsterSpecialAbility
from dnd5epy.models.monster_spell import MonsterSpell
from dnd5epy.models.monster_spellcasting import MonsterSpellcasting
from dnd5epy.models.monster_usage import MonsterUsage
from dnd5epy.models.multiclassing import Multiclassing
from dnd5epy.models.option import Option
from dnd5epy.models.option_one_of import OptionOneOf
from dnd5epy.models.option_one_of1 import OptionOneOf1
from dnd5epy.models.option_one_of10 import OptionOneOf10
from dnd5epy.models.option_one_of2 import OptionOneOf2
from dnd5epy.models.option_one_of3 import OptionOneOf3
from dnd5epy.models.option_one_of4 import OptionOneOf4
from dnd5epy.models.option_one_of5 import OptionOneOf5
from dnd5epy.models.option_one_of6 import OptionOneOf6
from dnd5epy.models.option_one_of7 import OptionOneOf7
from dnd5epy.models.option_one_of8 import OptionOneOf8
from dnd5epy.models.option_one_of9 import OptionOneOf9
from dnd5epy.models.option_set import OptionSet
from dnd5epy.models.option_set_one_of import OptionSetOneOf
from dnd5epy.models.option_set_one_of1 import OptionSetOneOf1
from dnd5epy.models.option_set_one_of2 import OptionSetOneOf2
from dnd5epy.models.prerequisite import Prerequisite
from dnd5epy.models.prerequisite_ability_score import PrerequisiteAbilityScore
from dnd5epy.models.proficiency import Proficiency
from dnd5epy.models.proficiency_all_of_reference import ProficiencyAllOfReference
from dnd5epy.models.race import Race
from dnd5epy.models.resource_description import ResourceDescription
from dnd5epy.models.rule import Rule
from dnd5epy.models.rule_section import RuleSection
from dnd5epy.models.skill import Skill
from dnd5epy.models.spell import Spell
from dnd5epy.models.spell_all_of_damage import SpellAllOfDamage
from dnd5epy.models.spell_prerequisite import SpellPrerequisite
from dnd5epy.models.spellcasting import Spellcasting
from dnd5epy.models.spellcasting_info_inner import SpellcastingInfoInner
from dnd5epy.models.spellcasting_spellcasting_ability import SpellcastingSpellcastingAbility
from dnd5epy.models.subclass import Subclass
from dnd5epy.models.subclass_all_of_spells import SubclassAllOfSpells
from dnd5epy.models.subclass_level import SubclassLevel
from dnd5epy.models.subclass_level_resource import SubclassLevelResource
from dnd5epy.models.subclass_level_spellcasting import SubclassLevelSpellcasting
from dnd5epy.models.subrace import Subrace
from dnd5epy.models.subrace_all_of_race import SubraceAllOfRace
from dnd5epy.models.trait import Trait
from dnd5epy.models.trait_all_of_trait_specific import TraitAllOfTraitSpecific
from dnd5epy.models.trait_all_of_trait_specific_one_of import TraitAllOfTraitSpecificOneOf
from dnd5epy.models.trait_all_of_trait_specific_one_of_breath_weapon import TraitAllOfTraitSpecificOneOfBreathWeapon
from dnd5epy.models.trait_all_of_trait_specific_one_of_breath_weapon_damage import TraitAllOfTraitSpecificOneOfBreathWeaponDamage
from dnd5epy.models.trait_all_of_trait_specific_one_of_breath_weapon_usage import TraitAllOfTraitSpecificOneOfBreathWeaponUsage
from dnd5epy.models.trait_all_of_trait_specific_one_of_damage_type import TraitAllOfTraitSpecificOneOfDamageType
from dnd5epy.models.weapon import Weapon
from dnd5epy.models.weapon_all_of_range import WeaponAllOfRange
from dnd5epy.models.weapon_property import WeaponProperty
