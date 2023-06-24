from dataclasses import dataclass, field
import json
import pprint as pp

@dataclass
class DnDCharacterStats:
    name :str
    describtion : str
    level : int
    hitpoints : int
    ability_scores: dict
    alignment: str
    background: str
    char_class: str
    languages: list
    proficiencies: list
    race: str
    skills: list

    def __post_init__(self):
        pp.pprint(self.__dict__)
        self.name = ""
        self.describtion = ""
        self.hitpoints = 
        #TODO add langchain model to reutnrn name + description

    @classmethod
    def from_json(cls, json_data):
        return cls(**json_data)
    
    def to_json(self):
        return json.dumps(self.__dict__)

# Example usage:
json_data = '''
{
  "abilityScores": {
    "cha": 14,
    "con": 12,
    "dex": 16,
    "int": 10,
    "str": 13,
    "wis": 8
  },
  "alignment": "chaotic-good",
  "background": "acolyte",
  "class": "rogue",
  "languages": [
    "common",
    "elvish"
  ],
  "proficiencies": [
    "darts",
    "shortswords",
    "thieves-tools",
    "skill-acrobatics",
    "skill-deception",
    "skill-perception",
    "skill-stealth"
  ],
  "race": "elf",
  "skills": [
    "acrobatics",
    "perception",
    "stealth"
  ]
}
'''

# Note the renaming in the loaded json data to match the class variable names
json_data_dict = json.loads(json_data)
json_data_dict['char_class'] = json_data_dict.pop('class')
json_data_dict['ability_scores'] = json_data_dict.pop('abilityScores')

character = DnDCharacterStats.from_json(json_data_dict)

# Accessing the character's attributes
#print(character.ability_scores)
#print(character.alignment)
#print(character.background)
#print(character.char_class)
#print(character.languages)
#print(character.proficiencies)

#print(character.race)
#print(character.skills)
