from dnd5epy.api.common_api import CommonApi
import json

def main():
    # Instantiate the API client
    
    client = CommonApi()
    # getting all possible endpoints for the common API class
    # endpoints_names = client.api_get()
    output  = {}
    character_endpoints = ['ability-scores', 'alignments', 'backgrounds', 'classes', 'languages', 'proficiencies', "skills" , "races"] 
    for key in character_endpoints:
        options = client.api_endpoint_get(key)
        output[key] = [opts['index'] for opts in options.to_dict()["results"]]

    # Save the output to a JSON file
    with open("character_build_input.json", "w") as file:
        json.dump(output, file, indent=4)

if __name__ == "__main__":
    main()  