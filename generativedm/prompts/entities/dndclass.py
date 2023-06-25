import json
from dataclasses import dataclass
from dataclasses_json import dataclass_json


@dataclass_json
@dataclass()
class CharacterBuilder:
    task: str
    taskRules: list[str]
    character: dict
    proficiency_choices: list
    starting_equipment_choices: list
    game_settings: dict


@dataclass_json
@dataclass()
class Character:
    name: str
    character_class: str
    description: str


@dataclass_json
@dataclass()
class ProficiencyChoice:
    pass
    # Add relevant attributes here


@dataclass_json
@dataclass()
class StartingEquipmentChoice:
    pass
    # Add relevant attributes here


# Load and parse the JSON data
def load_character_prompt():
    filename = "/home/ron/Documents/projects/generativedm/generativedm/prompts/entities/class.json"
    with open(filename, 'r') as file:
        json_data = json.load(file)
    return CharacterBuilder.from_dict(json_data)


