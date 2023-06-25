from dnd5apy.api import ClassApi
from typing import List, Dict, Any, Union, Tuple
from dataclasses import dataclass
import json
import argparse

@dataclass
class Proficiencies:
    items: Dict[str, List[str]]  # will hold the list of index for each proficiency category

@dataclass
@dataclass
class EquipmentData:
    items: Dict[str, List[Tuple[str, int]]]  # will hold the list of index for each equipment category

@dataclass
class Choice:
    type: str
    description: str
    quantity: int
    options: List[Dict[str, str]]


def get_stats(data: Dict[str, Any], keyword: str) -> Union[Proficiencies, EquipmentData]:
    item_data = []
    for item in data.get(keyword, []):
        match item:
            case {"index": index}:  # proficiency items have no quantity, so we'll just capture the index
                item_data.append(index)
            case {"equipment": {"index": index}, "quantity": quantity}:  # starting equipment items
                item_data.append(index)
            case _:
                print(f"Unmatched item: {item}")

    # Return an instance of the appropriate class based on the type
    if keyword == 'proficiencies':
        return Proficiencies(items={keyword: item_data})
    elif keyword in ['equipment', 'starting_equipment']:
        return EquipmentData(items={keyword: item_data})
    else:
        raise ValueError(f'Unknown type: {keyword}')


def get_choices(data: Dict[str, Any], keyword: str) -> List[Choice]:
    choices = []
    for _, item in enumerate(data.get(keyword, [])):
        options = []
        for option in item.get('_from', {}).get('options', []):
            match option:
                case {"item": item, "count": count}:
                    options.append(item.get('index', ''))
                case {"option_type": "multiple", "items": sub_items}:
                    for sub_item in sub_items:
                        if 'of' in sub_item:
                            options.append(sub_item['of'].get('index', ''))
                        else:
                            options.append(sub_item.get('index', ''))
                case {"option_type": "counted_reference", "count": count, "of": {"index": index} }:
                    options.append(index)
                case {"option_type": "reference", "item": {"index": index} }:
                    options.append(index)
                case _:
                    print("Unmatched option: ", option)
        choice = Choice(
            type=item.get('type'),
            description=item.get('desc', ''),
            quantity=item.get('choose', 1),  # quantity of choices that need to be made
            options=options
        )
        choices.append(choice)
    return choices

def main(index, name, dryrun):
    path = "/home/ron/Documents/projects/generativedm/data/characters/"
    if not dryrun:
        class_client = ClassApi()
        data = class_client.api_classes_index_get(index).to_dict()
    elif dryrun:
        file_path = f"{path}Jody.json"
        with open(file_path, "r") as file:
            data = json.load(file)
    
    print(data)

    hit_die = data['hit_die']
    starter_items = get_stats(data, 'starting_equipment')
    proficiency_skills = get_stats(data, 'proficiencies')
    proficiency_choices = get_choices(data, 'proficiency_choices')
    starting_equipment_choices = get_choices(data, 'starting_equipment_options')
    
    output_data = {
        'hit_die': hit_die,
        'proficiency_skills': proficiency_skills.items,
        'proficiency_choices': [choice.__dict__ for choice in proficiency_choices],
        'starting_equipment_choices': [choice.__dict__ for choice in starting_equipment_choices],
        'starter_items': starter_items.items
    }
    
    with open(f"{path}{name}.json", 'w') as f:
        json.dump(output_data, f, indent=4)

if __name__ == "__main__": 
    parser = argparse.ArgumentParser(description='Process D&D class data and save as JSON.')
    parser.add_argument('--index', type=str, help='Class index', required=True)
    parser.add_argument('--name', type=str, help='Character name', required=True)
    parser.add_argument('--dryrun', action='store_true', help='Perform a dry run')


    args = parser.parse_args()
    main(args.index, args.name, args.dryrun)