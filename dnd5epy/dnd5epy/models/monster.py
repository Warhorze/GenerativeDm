# coding: utf-8

"""
    D&D 5e API

    # Introduction  Welcome to the dnd5eapi, the Dungeons & Dragons 5th Edition API! This documentation should help you familiarize yourself with the resources available and how to consume them with HTTP requests. Read through the getting started section before you dive in. Most of your problems should be solved just by reading through it.  ## Getting Started  Let's make our first API request to the D&D 5th Edition API!  Open up a terminal and use [curl](http://curl.haxx.se/) or [httpie](http://httpie.org/) to make an API request for a resource. You can also scroll through the definitions below and send requests directly from the endpoint documentation!  For example, if you paste and run this `curl` command: ```bash curl -X GET \"https://www.dnd5eapi.co/api/ability-scores/cha\" -H \"Accept: application/json\" ```  We should see a result containing details about the Charisma ability score: ```bash {   \"index\": \"cha\",   \"name\": \"CHA\",   \"full_name\": \"Charisma\",   \"desc\": [     \"Charisma measures your ability to interact effectively with others. It       includes such factors as confidence and eloquence, and it can represent       a charming or commanding personality.\",     \"A Charisma check might arise when you try to influence or entertain       others, when you try to make an impression or tell a convincing lie,       or when you are navigating a tricky social situation. The Deception,       Intimidation, Performance, and Persuasion skills reflect aptitude in       certain kinds of Charisma checks.\"   ],   \"skills\": [     {       \"name\": \"Deception\",       \"index\": \"deception\",       \"url\": \"/api/skills/deception\"     },     {       \"name\": \"Intimidation\",       \"index\": \"intimidation\",       \"url\": \"/api/skills/intimidation\"     },     {       \"name\": \"Performance\",       \"index\": \"performance\",       \"url\": \"/api/skills/performance\"     },     {       \"name\": \"Persuasion\",       \"index\": \"persuasion\",       \"url\": \"/api/skills/persuasion\"     }   ],   \"url\": \"/api/ability-scores/cha\" } ```  ## Authentication  The dnd5eapi is a completely open API. No authentication is required to query and get data. This also means that we've limited what you can do to just `GET`-ing the data. If you find a mistake in the data, feel free to [message us](https://discord.gg/TQuYTv7).  ## GraphQL  This API supports [GraphQL](https://graphql.org/). The GraphQL URL for this API is `https://www.dnd5eapi.co/graphql`. Most of your questions regarding the GraphQL schema can be answered by querying the endpoint with the Apollo sandbox explorer.  ## Schemas  Definitions of all schemas will be accessible in a future update. Two of the most common schemas are described here.  ### `APIReference` Represents a minimal representation of a resource. The detailed representation of the referenced resource can be retrieved by making a request to the referenced `URL`. ``` APIReference {   index     string   name      string   url       string } ``` <hr>  ### `DC` Represents a difficulty check. ``` DC {   dc_type       APIReference   dc_value      number   success_type  \"none\" | \"half\" | \"other\" } ``` <hr>  ### `Damage` Represents damage. ``` Damage {   damage_type     APIReference   damage_dice     string } ``` <hr>  ### `Choice` Represents a choice made by a player. Commonly seen related to decisions made during character creation or combat (e.g.: the description of the cleric class, under **Proficiencies**, states \"Skills: Choose two from History, Insight, Medicine, Persuasion, and Religion\" [[SRD p15]](https://media.wizards.com/2016/downloads/DND/SRD-OGL_V5.1.pdf#page=15)) ``` Choice {   desc      string   choose    number   type      string   from      OptionSet } ``` <hr>  ### `OptionSet` The OptionSet structure provides the options to be chosen from, or sufficient data to fetch and interpret the options. All OptionSets have an `option_set_type` attribute that indicates the structure of the object that contains the options. The possible values are `options_array`, `equipment_category`, and `reference_list`. Other attributes on the OptionSet depend on the value of this attribute. - `options_array`   - `options` (array): An array of Option objects. Each item in the array represents an option that can be chosen. - `equipment_category`   - `equipment_category` (APIReference): A reference to an EquipmentCategory. Each item in the EquipmentCategory's `equipment` array represents one option that can be chosen. - `resource_list`   - `resource_list_url` (string): A reference (by URL) to a collection in the database. The URL may include query parameters. Each item in the resulting ResourceList's `results` array represents one option that can be chosen. <hr>  ### `Option` When the options are given in an `options_array`, each item in the array inherits from the Option structure. All Options have an `option_type` attribute that indicates the structure of the option. The value of this attribute indicates how the option should be handled, and each type has different attributes. The possible values and their corresponding attributes are listed below. - `reference` - A terminal option. Contains a reference to a Document that can be added to the list of options chosen.   - `item` (APIReference): A reference to the chosen item. - `action` - A terminal option. Contains information describing an action, for use within Multiattack actions.   - `action_name` (string): The name of the action, according to its `name` attribute.   - `count` (number | string): The number of times this action can be repeated if this option is chosen.   - `type` (string = `\"melee\" | \"ranged\" | \"ability\" | \"magic\"`, optional): For attack actions that can be either melee, ranged, abilities, or magic. - `multiple` - When this option is chosen, all of its child options are chosen, and must be resolved the same way as a normal option.   - `items` (array): An array of Option objects. All of them must be taken if the option is chosen. - `choice` - A nested choice. If this option is chosen, the Choice structure contained within must be resolved like a normal Choice structure, and the results are the chosen options.   - `choice` (Choice): The Choice to resolve. - `string` - A terminal option. Contains a reference to a string.   - `string` (string): The string. - `ideal` - A terminal option. Contains information about an ideal.   - `desc` (string): A description of the ideal.   - `alignments` (ApiReference[]): A list of alignments of those who might follow the ideal. - `counted_reference` - A terminal option. Contains a reference to something else in the API along with a count.   - `count` (number): Count.   - `of` (ApiReference): Thing being referenced. - `score_prerequisite` - A terminal option. Contains a reference to an ability score and a minimum score.   - `ability_score` (ApiReference): Ability score being referenced.   - `minimum_score` (number): The minimum score required to satisfy the prerequisite. - `ability_bonus` - A terminal option. Contains a reference to an ability score and a bonus   - `ability_score` (ApiReference): Ability score being referenced   - `bonus` (number): The bonus being applied to the ability score - `breath` - A terminal option: Contains a reference to information about a breath attack.   - `name` (string): Name of the breath.   - `dc` (DC): Difficulty check of the breath attack.   - `damage` ([Damage]): Damage dealt by the breath attack, if any. - `damage` - A terminal option. Contains information about damage.   - `damage_type` (ApiReference): Reference to type of damage.   - `damage_dice` (string): Damage expressed in dice (e.g. \"13d6\").   - `notes` (string): Information regarding the damage.  ## FAQ  ### What is the SRD? The SRD, or Systems Reference Document, contains guidelines for publishing content under the OGL. This allows for some of the data for D&D 5e to be open source. The API only covers data that can be found in the SRD. [Here's a link to the full text of the SRD.](https://media.wizards.com/2016/downloads/DND/SRD-OGL_V5.1.pdf)  ### What is the OGL? The Open Game License (OGL) is a public copyright license by Wizards of the Coast that may be used by tabletop role-playing game developers to grant permission to modify, copy, and redistribute some of the content designed for their games, notably game mechanics. However, they must share-alike copies and derivative works. [More information about the OGL can be found here.](https://en.wikipedia.org/wiki/Open_Game_License)  ### A monster, spell, subclass, etc. is missing from the API / Database. Can I add it? Please check if the data is within the SRD. If it is, feel free to open an issue or PR to add it yourself. Otherwise, due to legal reasons, we cannot add it.  ### Can this API be self hosted? Yes it can! You can also host the data yourself if you don't want to use the API at all. You can also make changes and add extra data if you like. However, it is up to you to merge in new changes to the data and API.  #### Can I publish is on <insert platform>? Is this free use? Yes, you can. The API itself is under the [MIT license](https://opensource.org/licenses/MIT), and the underlying data accessible via the API is supported under the SRD and OGL.  # Status Page  The status page for the API can be found here: https://5e-bits.github.io/dnd-uptime/  # Chat  Come hang out with us [on Discord](https://discord.gg/TQuYTv7)!  # Contribute  This API is built from two repositories.   - The repo containing the data lives here: https://github.com/bagelbits/5e-database   - The repo with the API implementation lives here: https://github.com/bagelbits/5e-srd-api  This is a evolving API and having fresh ideas are always welcome! You can open an issue in either repo, open a PR for changes, or just discuss with other users in this discord.   # noqa: E501

    The version of the OpenAPI document: 0.1
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""


from __future__ import annotations
import pprint
import re  # noqa: F401
import json


from typing import List, Optional, Union
from pydantic import BaseModel, Field, StrictFloat, StrictInt, StrictStr, confloat, conint, conlist, validator
from dnd5epy.models.api_reference import APIReference
from dnd5epy.models.monster_action import MonsterAction
from dnd5epy.models.monster_all_of_senses import MonsterAllOfSenses
from dnd5epy.models.monster_all_of_speed import MonsterAllOfSpeed
from dnd5epy.models.monster_armor_class import MonsterArmorClass
from dnd5epy.models.monster_proficiency import MonsterProficiency
from dnd5epy.models.monster_special_ability import MonsterSpecialAbility

class Monster(BaseModel):
    """
    `Monster` 
    """
    index: Optional[StrictStr] = Field(None, description="Resource index for shorthand searching.")
    name: Optional[StrictStr] = Field(None, description="Name of the referenced resource.")
    url: Optional[StrictStr] = Field(None, description="URL of the referenced resource.")
    desc: Optional[conlist(StrictStr)] = Field(None, description="Description of the resource.")
    charisma: Optional[Union[StrictFloat, StrictInt]] = Field(None, description="A monster's ability to charm or intimidate a player.")
    constitution: Optional[Union[StrictFloat, StrictInt]] = Field(None, description="How sturdy a monster is.\"")
    dexterity: Optional[Union[StrictFloat, StrictInt]] = Field(None, description="The monster's ability for swift movement or stealth")
    intelligence: Optional[Union[StrictFloat, StrictInt]] = Field(None, description="The monster's ability to outsmart a player.")
    strength: Optional[Union[StrictFloat, StrictInt]] = Field(None, description="How hard a monster can hit a player.")
    wisdom: Optional[Union[StrictFloat, StrictInt]] = Field(None, description="A monster's ability to ascertain the player's plan.")
    image: Optional[StrictStr] = Field(None, description="The image url of the monster.")
    size: Optional[StrictStr] = Field(None, description="The size of the monster ranging from Tiny to Gargantuan.\"")
    type: Optional[StrictStr] = Field(None, description="The type of monster.")
    subtype: Optional[StrictStr] = Field(None, description="The sub-category of a monster used for classification of monsters.\"")
    alignments: Optional[StrictStr] = Field(None, description="A creature's general moral and personal attitudes.")
    armor_class: Optional[conlist(MonsterArmorClass)] = Field(None, description="The difficulty for a player to successfully deal damage to a monster.")
    hit_points: Optional[Union[StrictFloat, StrictInt]] = Field(None, description="The hit points of a monster determine how much damage it is able to take before it can be defeated.")
    hit_dice: Optional[StrictStr] = Field(None, description="The hit die of a monster can be used to make a version of the same monster whose hit points are determined by the roll of the die. For example: A monster with 2d6 would have its hit points determine by rolling a 6 sided die twice.")
    hit_points_roll: Optional[StrictStr] = Field(None, description="The roll for determining a monster's hit points, which consists of the hit dice (e.g. 18d10) and the modifier determined by its Constitution (e.g. +36). For example, 18d10+36")
    actions: Optional[conlist(MonsterAction)] = Field(None, description="A list of actions that are available to the monster to take during combat.")
    legendary_actions: Optional[conlist(MonsterAction)] = Field(None, description="A list of legendary actions that are available to the monster to take during combat.")
    challenge_rating: Optional[Union[confloat(le=21, ge=0, strict=True), conint(le=21, ge=0, strict=True)]] = Field(None, description="A monster's challenge rating is a guideline number that says when a monster becomes an appropriate challenge against the party's average level. For example. A group of 4 players with an average level of 4 would have an appropriate combat challenge against a monster with a challenge rating of 4 but a monster with a challenge rating of 8 against the same group of players would pose a significant threat.")
    condition_immunities: Optional[conlist(APIReference)] = Field(None, description="A list of conditions that a monster is immune to.")
    damage_immunities: Optional[conlist(StrictStr)] = Field(None, description="A list of damage types that a monster will take double damage from.")
    damage_resistances: Optional[conlist(StrictStr)] = Field(None, description="A list of damage types that a monster will take half damage from.")
    damage_vulnerabilities: Optional[conlist(StrictStr)] = Field(None, description="A list of damage types that a monster will take double damage from.")
    forms: Optional[conlist(APIReference)] = Field(None, description="List of other related monster entries that are of the same form. Only applicable to Lycanthropes that have multiple forms.")
    languages: Optional[StrictStr] = Field(None, description="The languages a monster is able to speak.")
    proficiencies: Optional[conlist(MonsterProficiency)] = Field(None, description="A list of proficiencies of a monster.")
    reactions: Optional[conlist(MonsterAction)] = Field(None, description="A list of reactions that is available to the monster to take during combat.")
    senses: Optional[MonsterAllOfSenses] = None
    special_abilities: Optional[conlist(MonsterSpecialAbility)] = Field(None, description="A list of the monster's special abilities.")
    speed: Optional[MonsterAllOfSpeed] = None
    xp: Optional[Union[StrictFloat, StrictInt]] = Field(None, description="The number of experience points (XP) a monster is worth is based on its challenge rating.")
    __properties = ["index", "name", "url", "desc", "charisma", "constitution", "dexterity", "intelligence", "strength", "wisdom", "image", "size", "type", "subtype", "alignments", "armor_class", "hit_points", "hit_dice", "hit_points_roll", "actions", "legendary_actions", "challenge_rating", "condition_immunities", "damage_immunities", "damage_resistances", "damage_vulnerabilities", "forms", "languages", "proficiencies", "reactions", "senses", "special_abilities", "speed", "xp"]

    @validator('size')
    def size_validate_enum(cls, value):
        """Validates the enum"""
        if value is None:
            return value

        if value not in ('Tiny', 'Small', 'Medium', 'Large', 'Huge', 'Gargantuan'):
            raise ValueError("must be one of enum values ('Tiny', 'Small', 'Medium', 'Large', 'Huge', 'Gargantuan')")
        return value

    @validator('alignments')
    def alignments_validate_enum(cls, value):
        """Validates the enum"""
        if value is None:
            return value

        if value not in ('chaotic neutral', 'chaotic evil', 'chaotic good', 'lawful neutral', 'lawful evil', 'lawful good', 'neutral', 'neutral evil', 'neutral good', 'any alignment', 'unaligned'):
            raise ValueError("must be one of enum values ('chaotic neutral', 'chaotic evil', 'chaotic good', 'lawful neutral', 'lawful evil', 'lawful good', 'neutral', 'neutral evil', 'neutral good', 'any alignment', 'unaligned')")
        return value

    class Config:
        """Pydantic configuration"""
        allow_population_by_field_name = True
        validate_assignment = True

    def to_str(self) -> str:
        """Returns the string representation of the model using alias"""
        return pprint.pformat(self.dict(by_alias=True))

    def to_json(self) -> str:
        """Returns the JSON representation of the model using alias"""
        return json.dumps(self.to_dict())

    @classmethod
    def from_json(cls, json_str: str) -> Monster:
        """Create an instance of Monster from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        # override the default output from pydantic by calling `to_dict()` of each item in armor_class (list)
        _items = []
        if self.armor_class:
            for _item in self.armor_class:
                if _item:
                    _items.append(_item.to_dict())
            _dict['armor_class'] = _items
        # override the default output from pydantic by calling `to_dict()` of each item in actions (list)
        _items = []
        if self.actions:
            for _item in self.actions:
                if _item:
                    _items.append(_item.to_dict())
            _dict['actions'] = _items
        # override the default output from pydantic by calling `to_dict()` of each item in legendary_actions (list)
        _items = []
        if self.legendary_actions:
            for _item in self.legendary_actions:
                if _item:
                    _items.append(_item.to_dict())
            _dict['legendary_actions'] = _items
        # override the default output from pydantic by calling `to_dict()` of each item in condition_immunities (list)
        _items = []
        if self.condition_immunities:
            for _item in self.condition_immunities:
                if _item:
                    _items.append(_item.to_dict())
            _dict['condition_immunities'] = _items
        # override the default output from pydantic by calling `to_dict()` of each item in forms (list)
        _items = []
        if self.forms:
            for _item in self.forms:
                if _item:
                    _items.append(_item.to_dict())
            _dict['forms'] = _items
        # override the default output from pydantic by calling `to_dict()` of each item in proficiencies (list)
        _items = []
        if self.proficiencies:
            for _item in self.proficiencies:
                if _item:
                    _items.append(_item.to_dict())
            _dict['proficiencies'] = _items
        # override the default output from pydantic by calling `to_dict()` of each item in reactions (list)
        _items = []
        if self.reactions:
            for _item in self.reactions:
                if _item:
                    _items.append(_item.to_dict())
            _dict['reactions'] = _items
        # override the default output from pydantic by calling `to_dict()` of senses
        if self.senses:
            _dict['senses'] = self.senses.to_dict()
        # override the default output from pydantic by calling `to_dict()` of each item in special_abilities (list)
        _items = []
        if self.special_abilities:
            for _item in self.special_abilities:
                if _item:
                    _items.append(_item.to_dict())
            _dict['special_abilities'] = _items
        # override the default output from pydantic by calling `to_dict()` of speed
        if self.speed:
            _dict['speed'] = self.speed.to_dict()
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> Monster:
        """Create an instance of Monster from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return Monster.parse_obj(obj)

        _obj = Monster.parse_obj({
            "index": obj.get("index"),
            "name": obj.get("name"),
            "url": obj.get("url"),
            "desc": obj.get("desc"),
            "charisma": obj.get("charisma"),
            "constitution": obj.get("constitution"),
            "dexterity": obj.get("dexterity"),
            "intelligence": obj.get("intelligence"),
            "strength": obj.get("strength"),
            "wisdom": obj.get("wisdom"),
            "image": obj.get("image"),
            "size": obj.get("size"),
            "type": obj.get("type"),
            "subtype": obj.get("subtype"),
            "alignments": obj.get("alignments"),
            "armor_class": [MonsterArmorClass.from_dict(_item) for _item in obj.get("armor_class")] if obj.get("armor_class") is not None else None,
            "hit_points": obj.get("hit_points"),
            "hit_dice": obj.get("hit_dice"),
            "hit_points_roll": obj.get("hit_points_roll"),
            "actions": [MonsterAction.from_dict(_item) for _item in obj.get("actions")] if obj.get("actions") is not None else None,
            "legendary_actions": [MonsterAction.from_dict(_item) for _item in obj.get("legendary_actions")] if obj.get("legendary_actions") is not None else None,
            "challenge_rating": obj.get("challenge_rating"),
            "condition_immunities": [APIReference.from_dict(_item) for _item in obj.get("condition_immunities")] if obj.get("condition_immunities") is not None else None,
            "damage_immunities": obj.get("damage_immunities"),
            "damage_resistances": obj.get("damage_resistances"),
            "damage_vulnerabilities": obj.get("damage_vulnerabilities"),
            "forms": [APIReference.from_dict(_item) for _item in obj.get("forms")] if obj.get("forms") is not None else None,
            "languages": obj.get("languages"),
            "proficiencies": [MonsterProficiency.from_dict(_item) for _item in obj.get("proficiencies")] if obj.get("proficiencies") is not None else None,
            "reactions": [MonsterAction.from_dict(_item) for _item in obj.get("reactions")] if obj.get("reactions") is not None else None,
            "senses": MonsterAllOfSenses.from_dict(obj.get("senses")) if obj.get("senses") is not None else None,
            "special_abilities": [MonsterSpecialAbility.from_dict(_item) for _item in obj.get("special_abilities")] if obj.get("special_abilities") is not None else None,
            "speed": MonsterAllOfSpeed.from_dict(obj.get("speed")) if obj.get("speed") is not None else None,
            "xp": obj.get("xp")
        })
        return _obj

