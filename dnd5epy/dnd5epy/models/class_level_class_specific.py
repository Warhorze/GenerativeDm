# coding: utf-8

"""
    D&D 5e API

    # Introduction  Welcome to the dnd5eapi, the Dungeons & Dragons 5th Edition API! This documentation should help you familiarize yourself with the resources available and how to consume them with HTTP requests. Read through the getting started section before you dive in. Most of your problems should be solved just by reading through it.  ## Getting Started  Let's make our first API request to the D&D 5th Edition API!  Open up a terminal and use [curl](http://curl.haxx.se/) or [httpie](http://httpie.org/) to make an API request for a resource. You can also scroll through the definitions below and send requests directly from the endpoint documentation!  For example, if you paste and run this `curl` command: ```bash curl -X GET \"https://www.dnd5eapi.co/api/ability-scores/cha\" -H \"Accept: application/json\" ```  We should see a result containing details about the Charisma ability score: ```bash {   \"index\": \"cha\",   \"name\": \"CHA\",   \"full_name\": \"Charisma\",   \"desc\": [     \"Charisma measures your ability to interact effectively with others. It       includes such factors as confidence and eloquence, and it can represent       a charming or commanding personality.\",     \"A Charisma check might arise when you try to influence or entertain       others, when you try to make an impression or tell a convincing lie,       or when you are navigating a tricky social situation. The Deception,       Intimidation, Performance, and Persuasion skills reflect aptitude in       certain kinds of Charisma checks.\"   ],   \"skills\": [     {       \"name\": \"Deception\",       \"index\": \"deception\",       \"url\": \"/api/skills/deception\"     },     {       \"name\": \"Intimidation\",       \"index\": \"intimidation\",       \"url\": \"/api/skills/intimidation\"     },     {       \"name\": \"Performance\",       \"index\": \"performance\",       \"url\": \"/api/skills/performance\"     },     {       \"name\": \"Persuasion\",       \"index\": \"persuasion\",       \"url\": \"/api/skills/persuasion\"     }   ],   \"url\": \"/api/ability-scores/cha\" } ```  ## Authentication  The dnd5eapi is a completely open API. No authentication is required to query and get data. This also means that we've limited what you can do to just `GET`-ing the data. If you find a mistake in the data, feel free to [message us](https://discord.gg/TQuYTv7).  ## GraphQL  This API supports [GraphQL](https://graphql.org/). The GraphQL URL for this API is `https://www.dnd5eapi.co/graphql`. Most of your questions regarding the GraphQL schema can be answered by querying the endpoint with the Apollo sandbox explorer.  ## Schemas  Definitions of all schemas will be accessible in a future update. Two of the most common schemas are described here.  ### `APIReference` Represents a minimal representation of a resource. The detailed representation of the referenced resource can be retrieved by making a request to the referenced `URL`. ``` APIReference {   index     string   name      string   url       string } ``` <hr>  ### `DC` Represents a difficulty check. ``` DC {   dc_type       APIReference   dc_value      number   success_type  \"none\" | \"half\" | \"other\" } ``` <hr>  ### `Damage` Represents damage. ``` Damage {   damage_type     APIReference   damage_dice     string } ``` <hr>  ### `Choice` Represents a choice made by a player. Commonly seen related to decisions made during character creation or combat (e.g.: the description of the cleric class, under **Proficiencies**, states \"Skills: Choose two from History, Insight, Medicine, Persuasion, and Religion\" [[SRD p15]](https://media.wizards.com/2016/downloads/DND/SRD-OGL_V5.1.pdf#page=15)) ``` Choice {   desc      string   choose    number   type      string   from      OptionSet } ``` <hr>  ### `OptionSet` The OptionSet structure provides the options to be chosen from, or sufficient data to fetch and interpret the options. All OptionSets have an `option_set_type` attribute that indicates the structure of the object that contains the options. The possible values are `options_array`, `equipment_category`, and `reference_list`. Other attributes on the OptionSet depend on the value of this attribute. - `options_array`   - `options` (array): An array of Option objects. Each item in the array represents an option that can be chosen. - `equipment_category`   - `equipment_category` (APIReference): A reference to an EquipmentCategory. Each item in the EquipmentCategory's `equipment` array represents one option that can be chosen. - `resource_list`   - `resource_list_url` (string): A reference (by URL) to a collection in the database. The URL may include query parameters. Each item in the resulting ResourceList's `results` array represents one option that can be chosen. <hr>  ### `Option` When the options are given in an `options_array`, each item in the array inherits from the Option structure. All Options have an `option_type` attribute that indicates the structure of the option. The value of this attribute indicates how the option should be handled, and each type has different attributes. The possible values and their corresponding attributes are listed below. - `reference` - A terminal option. Contains a reference to a Document that can be added to the list of options chosen.   - `item` (APIReference): A reference to the chosen item. - `action` - A terminal option. Contains information describing an action, for use within Multiattack actions.   - `action_name` (string): The name of the action, according to its `name` attribute.   - `count` (number | string): The number of times this action can be repeated if this option is chosen.   - `type` (string = `\"melee\" | \"ranged\" | \"ability\" | \"magic\"`, optional): For attack actions that can be either melee, ranged, abilities, or magic. - `multiple` - When this option is chosen, all of its child options are chosen, and must be resolved the same way as a normal option.   - `items` (array): An array of Option objects. All of them must be taken if the option is chosen. - `choice` - A nested choice. If this option is chosen, the Choice structure contained within must be resolved like a normal Choice structure, and the results are the chosen options.   - `choice` (Choice): The Choice to resolve. - `string` - A terminal option. Contains a reference to a string.   - `string` (string): The string. - `ideal` - A terminal option. Contains information about an ideal.   - `desc` (string): A description of the ideal.   - `alignments` (ApiReference[]): A list of alignments of those who might follow the ideal. - `counted_reference` - A terminal option. Contains a reference to something else in the API along with a count.   - `count` (number): Count.   - `of` (ApiReference): Thing being referenced. - `score_prerequisite` - A terminal option. Contains a reference to an ability score and a minimum score.   - `ability_score` (ApiReference): Ability score being referenced.   - `minimum_score` (number): The minimum score required to satisfy the prerequisite. - `ability_bonus` - A terminal option. Contains a reference to an ability score and a bonus   - `ability_score` (ApiReference): Ability score being referenced   - `bonus` (number): The bonus being applied to the ability score - `breath` - A terminal option: Contains a reference to information about a breath attack.   - `name` (string): Name of the breath.   - `dc` (DC): Difficulty check of the breath attack.   - `damage` ([Damage]): Damage dealt by the breath attack, if any. - `damage` - A terminal option. Contains information about damage.   - `damage_type` (ApiReference): Reference to type of damage.   - `damage_dice` (string): Damage expressed in dice (e.g. \"13d6\").   - `notes` (string): Information regarding the damage.  ## FAQ  ### What is the SRD? The SRD, or Systems Reference Document, contains guidelines for publishing content under the OGL. This allows for some of the data for D&D 5e to be open source. The API only covers data that can be found in the SRD. [Here's a link to the full text of the SRD.](https://media.wizards.com/2016/downloads/DND/SRD-OGL_V5.1.pdf)  ### What is the OGL? The Open Game License (OGL) is a public copyright license by Wizards of the Coast that may be used by tabletop role-playing game developers to grant permission to modify, copy, and redistribute some of the content designed for their games, notably game mechanics. However, they must share-alike copies and derivative works. [More information about the OGL can be found here.](https://en.wikipedia.org/wiki/Open_Game_License)  ### A monster, spell, subclass, etc. is missing from the API / Database. Can I add it? Please check if the data is within the SRD. If it is, feel free to open an issue or PR to add it yourself. Otherwise, due to legal reasons, we cannot add it.  ### Can this API be self hosted? Yes it can! You can also host the data yourself if you don't want to use the API at all. You can also make changes and add extra data if you like. However, it is up to you to merge in new changes to the data and API.  #### Can I publish is on <insert platform>? Is this free use? Yes, you can. The API itself is under the [MIT license](https://opensource.org/licenses/MIT), and the underlying data accessible via the API is supported under the SRD and OGL.  # Status Page  The status page for the API can be found here: https://5e-bits.github.io/dnd-uptime/  # Chat  Come hang out with us [on Discord](https://discord.gg/TQuYTv7)!  # Contribute  This API is built from two repositories.   - The repo containing the data lives here: https://github.com/bagelbits/5e-database   - The repo with the API implementation lives here: https://github.com/bagelbits/5e-srd-api  This is a evolving API and having fresh ideas are always welcome! You can open an issue in either repo, open a PR for changes, or just discuss with other users in this discord.   # noqa: E501

    The version of the OpenAPI document: 0.1
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""


from __future__ import annotations
from inspect import getfullargspec
import json
import pprint
import re  # noqa: F401

from typing import Optional
from pydantic import BaseModel, Field, StrictStr, ValidationError, validator
from dnd5epy.models.class_level_class_specific_any_of import ClassLevelClassSpecificAnyOf
from dnd5epy.models.class_level_class_specific_any_of1 import ClassLevelClassSpecificAnyOf1
from dnd5epy.models.class_level_class_specific_any_of10 import ClassLevelClassSpecificAnyOf10
from dnd5epy.models.class_level_class_specific_any_of11 import ClassLevelClassSpecificAnyOf11
from dnd5epy.models.class_level_class_specific_any_of2 import ClassLevelClassSpecificAnyOf2
from dnd5epy.models.class_level_class_specific_any_of3 import ClassLevelClassSpecificAnyOf3
from dnd5epy.models.class_level_class_specific_any_of4 import ClassLevelClassSpecificAnyOf4
from dnd5epy.models.class_level_class_specific_any_of5 import ClassLevelClassSpecificAnyOf5
from dnd5epy.models.class_level_class_specific_any_of6 import ClassLevelClassSpecificAnyOf6
from dnd5epy.models.class_level_class_specific_any_of7 import ClassLevelClassSpecificAnyOf7
from dnd5epy.models.class_level_class_specific_any_of8 import ClassLevelClassSpecificAnyOf8
from dnd5epy.models.class_level_class_specific_any_of9 import ClassLevelClassSpecificAnyOf9
from typing import Any, List
from pydantic import StrictStr, Field

CLASSLEVELCLASSSPECIFIC_ANY_OF_SCHEMAS = ["ClassLevelClassSpecificAnyOf", "ClassLevelClassSpecificAnyOf1", "ClassLevelClassSpecificAnyOf10", "ClassLevelClassSpecificAnyOf11", "ClassLevelClassSpecificAnyOf2", "ClassLevelClassSpecificAnyOf3", "ClassLevelClassSpecificAnyOf4", "ClassLevelClassSpecificAnyOf5", "ClassLevelClassSpecificAnyOf6", "ClassLevelClassSpecificAnyOf7", "ClassLevelClassSpecificAnyOf8", "ClassLevelClassSpecificAnyOf9"]

class ClassLevelClassSpecific(BaseModel):
    """
    Class specific information such as dice values for bard songs and number of warlock invocations.
    """

    # data type: ClassLevelClassSpecificAnyOf
    anyof_schema_1_validator: Optional[ClassLevelClassSpecificAnyOf] = None
    # data type: ClassLevelClassSpecificAnyOf1
    anyof_schema_2_validator: Optional[ClassLevelClassSpecificAnyOf1] = None
    # data type: ClassLevelClassSpecificAnyOf2
    anyof_schema_3_validator: Optional[ClassLevelClassSpecificAnyOf2] = None
    # data type: ClassLevelClassSpecificAnyOf3
    anyof_schema_4_validator: Optional[ClassLevelClassSpecificAnyOf3] = None
    # data type: ClassLevelClassSpecificAnyOf4
    anyof_schema_5_validator: Optional[ClassLevelClassSpecificAnyOf4] = None
    # data type: ClassLevelClassSpecificAnyOf5
    anyof_schema_6_validator: Optional[ClassLevelClassSpecificAnyOf5] = None
    # data type: ClassLevelClassSpecificAnyOf6
    anyof_schema_7_validator: Optional[ClassLevelClassSpecificAnyOf6] = None
    # data type: ClassLevelClassSpecificAnyOf7
    anyof_schema_8_validator: Optional[ClassLevelClassSpecificAnyOf7] = None
    # data type: ClassLevelClassSpecificAnyOf8
    anyof_schema_9_validator: Optional[ClassLevelClassSpecificAnyOf8] = None
    # data type: ClassLevelClassSpecificAnyOf9
    anyof_schema_10_validator: Optional[ClassLevelClassSpecificAnyOf9] = None
    # data type: ClassLevelClassSpecificAnyOf10
    anyof_schema_11_validator: Optional[ClassLevelClassSpecificAnyOf10] = None
    # data type: ClassLevelClassSpecificAnyOf11
    anyof_schema_12_validator: Optional[ClassLevelClassSpecificAnyOf11] = None
    actual_instance: Any
    any_of_schemas: List[str] = Field(CLASSLEVELCLASSSPECIFIC_ANY_OF_SCHEMAS, const=True)

    class Config:
        validate_assignment = True

    def __init__(self, *args, **kwargs):
        if args:
            if len(args) > 1:
                raise ValueError("If a position argument is used, only 1 is allowed to set `actual_instance`")
            if kwargs:
                raise ValueError("If a position argument is used, keyword arguments cannot be used.")
            super().__init__(actual_instance=args[0])
        else:
            super().__init__(**kwargs)

    @validator('actual_instance')
    def actual_instance_must_validate_anyof(cls, v):
        instance = ClassLevelClassSpecific.construct()
        error_messages = []
        # validate data type: ClassLevelClassSpecificAnyOf
        if not isinstance(v, ClassLevelClassSpecificAnyOf):
            error_messages.append(f"Error! Input type `{type(v)}` is not `ClassLevelClassSpecificAnyOf`")
        else:
            return v

        # validate data type: ClassLevelClassSpecificAnyOf1
        if not isinstance(v, ClassLevelClassSpecificAnyOf1):
            error_messages.append(f"Error! Input type `{type(v)}` is not `ClassLevelClassSpecificAnyOf1`")
        else:
            return v

        # validate data type: ClassLevelClassSpecificAnyOf2
        if not isinstance(v, ClassLevelClassSpecificAnyOf2):
            error_messages.append(f"Error! Input type `{type(v)}` is not `ClassLevelClassSpecificAnyOf2`")
        else:
            return v

        # validate data type: ClassLevelClassSpecificAnyOf3
        if not isinstance(v, ClassLevelClassSpecificAnyOf3):
            error_messages.append(f"Error! Input type `{type(v)}` is not `ClassLevelClassSpecificAnyOf3`")
        else:
            return v

        # validate data type: ClassLevelClassSpecificAnyOf4
        if not isinstance(v, ClassLevelClassSpecificAnyOf4):
            error_messages.append(f"Error! Input type `{type(v)}` is not `ClassLevelClassSpecificAnyOf4`")
        else:
            return v

        # validate data type: ClassLevelClassSpecificAnyOf5
        if not isinstance(v, ClassLevelClassSpecificAnyOf5):
            error_messages.append(f"Error! Input type `{type(v)}` is not `ClassLevelClassSpecificAnyOf5`")
        else:
            return v

        # validate data type: ClassLevelClassSpecificAnyOf6
        if not isinstance(v, ClassLevelClassSpecificAnyOf6):
            error_messages.append(f"Error! Input type `{type(v)}` is not `ClassLevelClassSpecificAnyOf6`")
        else:
            return v

        # validate data type: ClassLevelClassSpecificAnyOf7
        if not isinstance(v, ClassLevelClassSpecificAnyOf7):
            error_messages.append(f"Error! Input type `{type(v)}` is not `ClassLevelClassSpecificAnyOf7`")
        else:
            return v

        # validate data type: ClassLevelClassSpecificAnyOf8
        if not isinstance(v, ClassLevelClassSpecificAnyOf8):
            error_messages.append(f"Error! Input type `{type(v)}` is not `ClassLevelClassSpecificAnyOf8`")
        else:
            return v

        # validate data type: ClassLevelClassSpecificAnyOf9
        if not isinstance(v, ClassLevelClassSpecificAnyOf9):
            error_messages.append(f"Error! Input type `{type(v)}` is not `ClassLevelClassSpecificAnyOf9`")
        else:
            return v

        # validate data type: ClassLevelClassSpecificAnyOf10
        if not isinstance(v, ClassLevelClassSpecificAnyOf10):
            error_messages.append(f"Error! Input type `{type(v)}` is not `ClassLevelClassSpecificAnyOf10`")
        else:
            return v

        # validate data type: ClassLevelClassSpecificAnyOf11
        if not isinstance(v, ClassLevelClassSpecificAnyOf11):
            error_messages.append(f"Error! Input type `{type(v)}` is not `ClassLevelClassSpecificAnyOf11`")
        else:
            return v

        if error_messages:
            # no match
            raise ValueError("No match found when setting the actual_instance in ClassLevelClassSpecific with anyOf schemas: ClassLevelClassSpecificAnyOf, ClassLevelClassSpecificAnyOf1, ClassLevelClassSpecificAnyOf10, ClassLevelClassSpecificAnyOf11, ClassLevelClassSpecificAnyOf2, ClassLevelClassSpecificAnyOf3, ClassLevelClassSpecificAnyOf4, ClassLevelClassSpecificAnyOf5, ClassLevelClassSpecificAnyOf6, ClassLevelClassSpecificAnyOf7, ClassLevelClassSpecificAnyOf8, ClassLevelClassSpecificAnyOf9. Details: " + ", ".join(error_messages))
        else:
            return v

    @classmethod
    def from_dict(cls, obj: dict) -> ClassLevelClassSpecific:
        return cls.from_json(json.dumps(obj))

    @classmethod
    def from_json(cls, json_str: str) -> ClassLevelClassSpecific:
        """Returns the object represented by the json string"""
        instance = ClassLevelClassSpecific.construct()
        error_messages = []
        # anyof_schema_1_validator: Optional[ClassLevelClassSpecificAnyOf] = None
        try:
            instance.actual_instance = ClassLevelClassSpecificAnyOf.from_json(json_str)
            return instance
        except (ValidationError, ValueError) as e:
             error_messages.append(str(e))
        # anyof_schema_2_validator: Optional[ClassLevelClassSpecificAnyOf1] = None
        try:
            instance.actual_instance = ClassLevelClassSpecificAnyOf1.from_json(json_str)
            return instance
        except (ValidationError, ValueError) as e:
             error_messages.append(str(e))
        # anyof_schema_3_validator: Optional[ClassLevelClassSpecificAnyOf2] = None
        try:
            instance.actual_instance = ClassLevelClassSpecificAnyOf2.from_json(json_str)
            return instance
        except (ValidationError, ValueError) as e:
             error_messages.append(str(e))
        # anyof_schema_4_validator: Optional[ClassLevelClassSpecificAnyOf3] = None
        try:
            instance.actual_instance = ClassLevelClassSpecificAnyOf3.from_json(json_str)
            return instance
        except (ValidationError, ValueError) as e:
             error_messages.append(str(e))
        # anyof_schema_5_validator: Optional[ClassLevelClassSpecificAnyOf4] = None
        try:
            instance.actual_instance = ClassLevelClassSpecificAnyOf4.from_json(json_str)
            return instance
        except (ValidationError, ValueError) as e:
             error_messages.append(str(e))
        # anyof_schema_6_validator: Optional[ClassLevelClassSpecificAnyOf5] = None
        try:
            instance.actual_instance = ClassLevelClassSpecificAnyOf5.from_json(json_str)
            return instance
        except (ValidationError, ValueError) as e:
             error_messages.append(str(e))
        # anyof_schema_7_validator: Optional[ClassLevelClassSpecificAnyOf6] = None
        try:
            instance.actual_instance = ClassLevelClassSpecificAnyOf6.from_json(json_str)
            return instance
        except (ValidationError, ValueError) as e:
             error_messages.append(str(e))
        # anyof_schema_8_validator: Optional[ClassLevelClassSpecificAnyOf7] = None
        try:
            instance.actual_instance = ClassLevelClassSpecificAnyOf7.from_json(json_str)
            return instance
        except (ValidationError, ValueError) as e:
             error_messages.append(str(e))
        # anyof_schema_9_validator: Optional[ClassLevelClassSpecificAnyOf8] = None
        try:
            instance.actual_instance = ClassLevelClassSpecificAnyOf8.from_json(json_str)
            return instance
        except (ValidationError, ValueError) as e:
             error_messages.append(str(e))
        # anyof_schema_10_validator: Optional[ClassLevelClassSpecificAnyOf9] = None
        try:
            instance.actual_instance = ClassLevelClassSpecificAnyOf9.from_json(json_str)
            return instance
        except (ValidationError, ValueError) as e:
             error_messages.append(str(e))
        # anyof_schema_11_validator: Optional[ClassLevelClassSpecificAnyOf10] = None
        try:
            instance.actual_instance = ClassLevelClassSpecificAnyOf10.from_json(json_str)
            return instance
        except (ValidationError, ValueError) as e:
             error_messages.append(str(e))
        # anyof_schema_12_validator: Optional[ClassLevelClassSpecificAnyOf11] = None
        try:
            instance.actual_instance = ClassLevelClassSpecificAnyOf11.from_json(json_str)
            return instance
        except (ValidationError, ValueError) as e:
             error_messages.append(str(e))

        if error_messages:
            # no match
            raise ValueError("No match found when deserializing the JSON string into ClassLevelClassSpecific with anyOf schemas: ClassLevelClassSpecificAnyOf, ClassLevelClassSpecificAnyOf1, ClassLevelClassSpecificAnyOf10, ClassLevelClassSpecificAnyOf11, ClassLevelClassSpecificAnyOf2, ClassLevelClassSpecificAnyOf3, ClassLevelClassSpecificAnyOf4, ClassLevelClassSpecificAnyOf5, ClassLevelClassSpecificAnyOf6, ClassLevelClassSpecificAnyOf7, ClassLevelClassSpecificAnyOf8, ClassLevelClassSpecificAnyOf9. Details: " + ", ".join(error_messages))
        else:
            return instance

    def to_json(self) -> str:
        """Returns the JSON representation of the actual instance"""
        if self.actual_instance is None:
            return "null"

        to_json = getattr(self.actual_instance, "to_json", None)
        if callable(to_json):
            return self.actual_instance.to_json()
        else:
            return json.dumps(self.actual_instance)

    def to_dict(self) -> dict:
        """Returns the dict representation of the actual instance"""
        if self.actual_instance is None:
            return "null"

        to_json = getattr(self.actual_instance, "to_json", None)
        if callable(to_json):
            return self.actual_instance.to_dict()
        else:
            return json.dumps(self.actual_instance)

    def to_str(self) -> str:
        """Returns the string representation of the actual instance"""
        return pprint.pformat(self.dict())

