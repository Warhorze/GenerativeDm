

import json 
from generativedm.prompts.entities.dndclass import load_character_prompt, CharacterBuilder

from generativedm.config.config_files import APIkeys
from langchain.llms import OpenAI
import pprint as pp


llm = OpenAI(openai_api_key= APIkeys.OpenAI)



character_prompt = load_character_prompt()
path = "/home/ron/Documents/projects/generativedm/data/"
name = "Eldric"

haracter_prompt = load_character_prompt()
path = "/home/ron/Documents/projects/generativedm/data/"
name = "Eldric"

with open(f'{path}characters/{name}.json', 'r') as character_file:
    character_data = json.load(character_file)

with open(f'{path}/story/story.json', 'r') as story_file:
    story_data = json.load(story_file)

char = [char for char in story_data['characters'] if char['name'] == name][0]

# Create a new instance of CharacterBuilder with updated values
character = CharacterBuilder(
    task=character_prompt.task,
    taskRules=character_prompt.taskRules,
    character=char,
    proficiency_choices=character_data['proficiency_choices'],
    starting_equipment_choices=character_data['starting_equipment_choices'],
    game_settings=character_prompt.game_settings
)
print(char['description'])
print(30 * '-')
print(character_data['proficiency_choices'][0]['description'])
print(30 * '-')
print(character_data['starting_equipment_choices'][0]['description'])


#Dolly-v2, 8b seems to work pretty well and fit into local memory.. which is essential during development.

