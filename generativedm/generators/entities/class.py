from dnd5apy.api import ClassApi
import pprint as pp

from typing import List, Dict, Any, Union
from dataclasses import dataclass


from typing import List, Tuple

@dataclass
class ClassStats: 
    items : Dict[str, List[Tuple[str, float]]]  # will hold the list of (index, quantity) for each item category

def get_stats(data: Dict[str, Any], keyword: str) -> ClassStats:
    item_data = []
    for item in data.get(keyword, []):
        match item:
            case {"index": index}:  # proficiency items have no quantity, so we'll use 1 as a default
                item_data.append({'index': index, "count": 1.0})
            case {"equipment": {"index": index}, "quantity": quantity}:  # starting equipment items
                item_data.append({'index': index, "count": quantity})
            case _:
                print(f"Unmatched item: {item}")
    return ClassStats(items={keyword: item_data})
    
@dataclass
class Choice:
    type: str
    description: str
    quantity: float
    options: List[Dict[str, Union[str, int]]]

def get_choices(data: Dict[str, Any], keyword: str) -> List[Choice]:
    choices = []
    for _, item in enumerate(data.get(keyword, [])):
        options = []
        for option in item.get('_from', {}).get('options', []):
            match option:
                case {"item": item, "count": count}:
                    options.append({'index': item.get('index', ''), 'count': count})
                case {"option_type": "multiple", "items": sub_items}:
                    for sub_item in sub_items:
                        options.append({'index': sub_item['of'].get('index', ''), 'count': sub_item.get('count', 1)})
                case {"option_type": "counted_reference", "count": count, "of": {"index": index} }:
                    options.append({'index': index, 'count': count})
                case {"option_type": "reference", "item": {"index": index} }:
                    options.append({'index': index, 'count': 1})
                case _:
                    print("Unmatched option: ", option)
        choice = Choice(
            type=item.get('type'),
            description=item.get('desc', ''),
            quantity=item.get('choose', 0),
            options=options
        )
        choices.append(choice)
    return choices
data = {
    "hit_die": 8.0,
    "class_levels": "/api/classes/rogue/levels",
    "multi_classing": {
        "prerequisites": [
            {
                "ability_score": {
                    "index": "dex",
                    "name": "DEX",
                    "url": "/api/ability-scores/dex",
                },
                "minimum_score": 13.0,
            }
        ],
        "prerequisite_options": None,
        "proficiencies": [
            {
                "index": "light-armor",
                "name": "Light Armor",
                "url": "/api/proficiencies/light-armor",
            },
            {
                "index": "thieves-tools",
                "name": "Thieves' Tools",
                "url": "/api/proficiencies/thieves-tools",
            },
        ],
        "proficiency_choices": [
            {
                "desc": None,
                "choose": 1.0,
                "type": "proficiencies",
                "_from": {
                    "option_set_type": "options_array",
                    "options": [
                        {
                            "option_type": "reference",
                            "item": {
                                "index": "skill-acrobatics",
                                "name": "Skill: Acrobatics",
                                "url": "/api/proficiencies/skill-acrobatics",
                            },
                        },
                        {
                            "option_type": "reference",
                            "item": {
                                "index": "skill-athletics",
                                "name": "Skill: Athletics",
                                "url": "/api/proficiencies/skill-athletics",
                            },
                        },
                        {
                            "option_type": "reference",
                            "item": {
                                "index": "skill-deception",
                                "name": "Skill: Deception",
                                "url": "/api/proficiencies/skill-deception",
                            },
                        },
                        {
                            "option_type": "reference",
                            "item": {
                                "index": "skill-insight",
                                "name": "Skill: Insight",
                                "url": "/api/proficiencies/skill-insight",
                            },
                        },
                        {
                            "option_type": "reference",
                            "item": {
                                "index": "skill-intimidation",
                                "name": "Skill: Intimidation",
                                "url": "/api/proficiencies/skill-intimidation",
                            },
                        },
                        {
                            "option_type": "reference",
                            "item": {
                                "index": "skill-investigation",
                                "name": "Skill: Investigation",
                                "url": "/api/proficiencies/skill-investigation",
                            },
                        },
                        {
                            "option_type": "reference",
                            "item": {
                                "index": "skill-perception",
                                "name": "Skill: Perception",
                                "url": "/api/proficiencies/skill-perception",
                            },
                        },
                        {
                            "option_type": "reference",
                            "item": {
                                "index": "skill-performance",
                                "name": "Skill: Performance",
                                "url": "/api/proficiencies/skill-performance",
                            },
                        },
                        {
                            "option_type": "reference",
                            "item": {
                                "index": "skill-persuasion",
                                "name": "Skill: Persuasion",
                                "url": "/api/proficiencies/skill-persuasion",
                            },
                        },
                        {
                            "option_type": "reference",
                            "item": {
                                "index": "skill-sleight-of-hand",
                                "name": "Skill: Sleight of Hand",
                                "url": "/api/proficiencies/skill-sleight-of-hand",
                            },
                        },
                        {
                            "option_type": "reference",
                            "item": {
                                "index": "skill-stealth",
                                "name": "Skill: Stealth",
                                "url": "/api/proficiencies/skill-stealth",
                            },
                        },
                    ],
                },
            }
        ],
    },
    "spellcasting": None,
    "spells": None,
    "starting_equipment": [
        {
            "quantity": 1.0,
            "equipment": {
                "index": "leather-armor",
                "name": "Leather Armor",
                "url": "/api/equipment/leather-armor",
            },
        },
        {
            "quantity": 2.0,
            "equipment": {
                "index": "dagger",
                "name": "Dagger",
                "url": "/api/equipment/dagger",
            },
        },
        {
            "quantity": 1.0,
            "equipment": {
                "index": "thieves-tools",
                "name": "Thieves' Tools",
                "url": "/api/equipment/thieves-tools",
            },
        },
    ],
    "starting_equipment_options": [
        {
            "desc": "(a) a rapier or (b) a shortsword",
            "choose": 1.0,
            "type": "equipment",
            "_from": {
                "option_set_type": "options_array",
                "options": [
                    {
                        "option_type": "counted_reference",
                        "count": 1,
                        "of": {
                            "index": "rapier",
                            "name": "Rapier",
                            "url": "/api/equipment/rapier",
                        },
                    },
                    {
                        "option_type": "counted_reference",
                        "count": 1,
                        "of": {
                            "index": "shortsword",
                            "name": "Shortsword",
                            "url": "/api/equipment/shortsword",
                        },
                    },
                ],
            },
        },
        {
            "desc": "(a) a shortbow and quiver of 20 arrows or (b) a shortsword",
            "choose": 1.0,
            "type": "equipment",
            "_from": {
                "option_set_type": "options_array",
                "options": [
                    {
                        "option_type": "multiple",
                        "items": [
                            {
                                "option_type": "counted_reference",
                                "count": 1,
                                "of": {
                                    "index": "shortbow",
                                    "name": "Shortbow",
                                    "url": "/api/equipment/shortbow",
                                },
                            },
                            {
                                "option_type": "counted_reference",
                                "count": 20,
                                "of": {
                                    "index": "arrow",
                                    "name": "Arrow",
                                    "url": "/api/equipment/arrow",
                                },
                            },
                        ],
                    },
                    {
                        "option_type": "counted_reference",
                        "count": 1,
                        "of": {
                            "index": "shortsword",
                            "name": "Shortsword",
                            "url": "/api/equipment/shortsword",
                        },
                    },
                ],
            },
        },
        {
            "desc": "(a) a burglar’s pack, (b) a dungeoneer’s pack, or (c) an explorer’s pack",
            "choose": 1.0,
            "type": "equipment",
            "_from": {
                "option_set_type": "options_array",
                "options": [
                    {
                        "option_type": "counted_reference",
                        "count": 1,
                        "of": {
                            "index": "burglars-pack",
                            "name": "Burglar's Pack",
                            "url": "/api/equipment/burglars-pack",
                        },
                    },
                    {
                        "option_type": "counted_reference",
                        "count": 1,
                        "of": {
                            "index": "dungeoneers-pack",
                            "name": "Dungeoneer's Pack",
                            "url": "/api/equipment/dungeoneers-pack",
                        },
                    },
                    {
                        "option_type": "counted_reference",
                        "count": 1,
                        "of": {
                            "index": "explorers-pack",
                            "name": "Explorer's Pack",
                            "url": "/api/equipment/explorers-pack",
                        },
                    },
                ],
            },
        },
    ],
    "proficiency_choices": [
        {
            "desc": "Choose four from Acrobatics, Athletics, Deception, Insight, Intimidation, Investigation, Perception, Performance, Persuasion, Sleight of Hand, and Stealth",
            "choose": 4.0,
            "type": "proficiencies",
            "_from": {
                "option_set_type": "options_array",
                "options": [
                    {
                        "option_type": "reference",
                        "item": {
                            "index": "skill-acrobatics",
                            "name": "Skill: Acrobatics",
                            "url": "/api/proficiencies/skill-acrobatics",
                        },
                    },
                    {
                        "option_type": "reference",
                        "item": {
                            "index": "skill-athletics",
                            "name": "Skill: Athletics",
                            "url": "/api/proficiencies/skill-athletics",
                        },
                    },
                    {
                        "option_type": "reference",
                        "item": {
                            "index": "skill-deception",
                            "name": "Skill: Deception",
                            "url": "/api/proficiencies/skill-deception",
                        },
                    },
                    {
                        "option_type": "reference",
                        "item": {
                            "index": "skill-insight",
                            "name": "Skill: Insight",
                            "url": "/api/proficiencies/skill-insight",
                        },
                    },
                    {
                        "option_type": "reference",
                        "item": {
                            "index": "skill-intimidation",
                            "name": "Skill: Intimidation",
                            "url": "/api/proficiencies/skill-intimidation",
                        },
                    },
                    {
                        "option_type": "reference",
                        "item": {
                            "index": "skill-investigation",
                            "name": "Skill: Investigation",
                            "url": "/api/proficiencies/skill-investigation",
                        },
                    },
                    {
                        "option_type": "reference",
                        "item": {
                            "index": "skill-perception",
                            "name": "Skill: Perception",
                            "url": "/api/proficiencies/skill-perception",
                        },
                    },
                    {
                        "option_type": "reference",
                        "item": {
                            "index": "skill-performance",
                            "name": "Skill: Performance",
                            "url": "/api/proficiencies/skill-performance",
                        },
                    },
                    {
                        "option_type": "reference",
                        "item": {
                            "index": "skill-persuasion",
                            "name": "Skill: Persuasion",
                            "url": "/api/proficiencies/skill-persuasion",
                        },
                    },
                    {
                        "option_type": "reference",
                        "item": {
                            "index": "skill-sleight-of-hand",
                            "name": "Skill: Sleight of Hand",
                            "url": "/api/proficiencies/skill-sleight-of-hand",
                        },
                    },
                    {
                        "option_type": "reference",
                        "item": {
                            "index": "skill-stealth",
                            "name": "Skill: Stealth",
                            "url": "/api/proficiencies/skill-stealth",
                        },
                    },
                ],
            },
        }
    ],
    "proficiencies": [
        {
            "index": "light-armor",
            "name": "Light Armor",
            "url": "/api/proficiencies/light-armor",
        },
        {
            "index": "simple-weapons",
            "name": "Simple Weapons",
            "url": "/api/proficiencies/simple-weapons",
        },
        {
            "index": "longswords",
            "name": "Longswords",
            "url": "/api/proficiencies/longswords",
        },
        {"index": "rapiers", "name": "Rapiers", "url": "/api/proficiencies/rapiers"},
        {
            "index": "shortswords",
            "name": "Shortswords",
            "url": "/api/proficiencies/shortswords",
        },
        {
            "index": "hand-crossbows",
            "name": "Hand crossbows",
            "url": "/api/proficiencies/hand-crossbows",
        },
        {
            "index": "thieves-tools",
            "name": "Thieves' Tools",
            "url": "/api/proficiencies/thieves-tools",
        },
        {
            "index": "saving-throw-dex",
            "name": "Saving Throw: DEX",
            "url": "/api/proficiencies/saving-throw-dex",
        },
        {
            "index": "saving-throw-int",
            "name": "Saving Throw: INT",
            "url": "/api/proficiencies/saving-throw-int",
        },
    ],
    "saving_throws": [
        {"index": "dex", "name": "DEX", "url": "/api/ability-scores/dex"},
        {"index": "int", "name": "INT", "url": "/api/ability-scores/int"},
    ],
    "subclasses": [{"index": "thief", "name": "Thief", "url": "/api/subclasses/thief"}],
    "index": "rogue",
    "name": "Rogue",
    "url": "/api/classes/rogue",
}


index = "rogue"
class_client = ClassApi()
# data= class_client.api_classes_index_get(index).to_dict()

starter_items = get_stats(data, 'starting_equipment')
proficiency_skils = get_stats(data, 'proficiencies')
proficiency_choices = get_choices(data, 'proficiency_choices')
starting_equipment_choices = get_choices(data, 'starting_equipment_options')

print(proficiency_choices)
print( 40 * '-')
print(starting_equipment_choices)
print( 40 * '-')
print(proficiency_skils)
print( 40 * '-')
print(starter_items)

#TODO create an item class, and a skill class?
#TODO convert the choices to be added to the statsclass
#TODO write a promt that makes choices.
