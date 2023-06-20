/*
D&D 5e API

# Introduction  Welcome to the dnd5eapi, the Dungeons & Dragons 5th Edition API! This documentation should help you familiarize yourself with the resources available and how to consume them with HTTP requests. Read through the getting started section before you dive in. Most of your problems should be solved just by reading through it.  ## Getting Started  Let's make our first API request to the D&D 5th Edition API!  Open up a terminal and use [curl](http://curl.haxx.se/) or [httpie](http://httpie.org/) to make an API request for a resource. You can also scroll through the definitions below and send requests directly from the endpoint documentation!  For example, if you paste and run this `curl` command: ```bash curl -X GET \"https://www.dnd5eapi.co/api/ability-scores/cha\" -H \"Accept: application/json\" ```  We should see a result containing details about the Charisma ability score: ```bash {   \"index\": \"cha\",   \"name\": \"CHA\",   \"full_name\": \"Charisma\",   \"desc\": [     \"Charisma measures your ability to interact effectively with others. It       includes such factors as confidence and eloquence, and it can represent       a charming or commanding personality.\",     \"A Charisma check might arise when you try to influence or entertain       others, when you try to make an impression or tell a convincing lie,       or when you are navigating a tricky social situation. The Deception,       Intimidation, Performance, and Persuasion skills reflect aptitude in       certain kinds of Charisma checks.\"   ],   \"skills\": [     {       \"name\": \"Deception\",       \"index\": \"deception\",       \"url\": \"/api/skills/deception\"     },     {       \"name\": \"Intimidation\",       \"index\": \"intimidation\",       \"url\": \"/api/skills/intimidation\"     },     {       \"name\": \"Performance\",       \"index\": \"performance\",       \"url\": \"/api/skills/performance\"     },     {       \"name\": \"Persuasion\",       \"index\": \"persuasion\",       \"url\": \"/api/skills/persuasion\"     }   ],   \"url\": \"/api/ability-scores/cha\" } ```  ## Authentication  The dnd5eapi is a completely open API. No authentication is required to query and get data. This also means that we've limited what you can do to just `GET`-ing the data. If you find a mistake in the data, feel free to [message us](https://discord.gg/TQuYTv7).  ## GraphQL  This API supports [GraphQL](https://graphql.org/). The GraphQL URL for this API is `https://www.dnd5eapi.co/graphql`. Most of your questions regarding the GraphQL schema can be answered by querying the endpoint with the Apollo sandbox explorer.  ## Schemas  Definitions of all schemas will be accessible in a future update. Two of the most common schemas are described here.  ### `APIReference` Represents a minimal representation of a resource. The detailed representation of the referenced resource can be retrieved by making a request to the referenced `URL`. ``` APIReference {   index     string   name      string   url       string } ``` <hr>  ### `DC` Represents a difficulty check. ``` DC {   dc_type       APIReference   dc_value      number   success_type  \"none\" | \"half\" | \"other\" } ``` <hr>  ### `Damage` Represents damage. ``` Damage {   damage_type     APIReference   damage_dice     string } ``` <hr>  ### `Choice` Represents a choice made by a player. Commonly seen related to decisions made during character creation or combat (e.g.: the description of the cleric class, under **Proficiencies**, states \"Skills: Choose two from History, Insight, Medicine, Persuasion, and Religion\" [[SRD p15]](https://media.wizards.com/2016/downloads/DND/SRD-OGL_V5.1.pdf#page=15)) ``` Choice {   desc      string   choose    number   type      string   from      OptionSet } ``` <hr>  ### `OptionSet` The OptionSet structure provides the options to be chosen from, or sufficient data to fetch and interpret the options. All OptionSets have an `option_set_type` attribute that indicates the structure of the object that contains the options. The possible values are `options_array`, `equipment_category`, and `reference_list`. Other attributes on the OptionSet depend on the value of this attribute. - `options_array`   - `options` (array): An array of Option objects. Each item in the array represents an option that can be chosen. - `equipment_category`   - `equipment_category` (APIReference): A reference to an EquipmentCategory. Each item in the EquipmentCategory's `equipment` array represents one option that can be chosen. - `resource_list`   - `resource_list_url` (string): A reference (by URL) to a collection in the database. The URL may include query parameters. Each item in the resulting ResourceList's `results` array represents one option that can be chosen. <hr>  ### `Option` When the options are given in an `options_array`, each item in the array inherits from the Option structure. All Options have an `option_type` attribute that indicates the structure of the option. The value of this attribute indicates how the option should be handled, and each type has different attributes. The possible values and their corresponding attributes are listed below. - `reference` - A terminal option. Contains a reference to a Document that can be added to the list of options chosen.   - `item` (APIReference): A reference to the chosen item. - `action` - A terminal option. Contains information describing an action, for use within Multiattack actions.   - `action_name` (string): The name of the action, according to its `name` attribute.   - `count` (number | string): The number of times this action can be repeated if this option is chosen.   - `type` (string = `\"melee\" | \"ranged\" | \"ability\" | \"magic\"`, optional): For attack actions that can be either melee, ranged, abilities, or magic. - `multiple` - When this option is chosen, all of its child options are chosen, and must be resolved the same way as a normal option.   - `items` (array): An array of Option objects. All of them must be taken if the option is chosen. - `choice` - A nested choice. If this option is chosen, the Choice structure contained within must be resolved like a normal Choice structure, and the results are the chosen options.   - `choice` (Choice): The Choice to resolve. - `string` - A terminal option. Contains a reference to a string.   - `string` (string): The string. - `ideal` - A terminal option. Contains information about an ideal.   - `desc` (string): A description of the ideal.   - `alignments` (ApiReference[]): A list of alignments of those who might follow the ideal. - `counted_reference` - A terminal option. Contains a reference to something else in the API along with a count.   - `count` (number): Count.   - `of` (ApiReference): Thing being referenced. - `score_prerequisite` - A terminal option. Contains a reference to an ability score and a minimum score.   - `ability_score` (ApiReference): Ability score being referenced.   - `minimum_score` (number): The minimum score required to satisfy the prerequisite. - `ability_bonus` - A terminal option. Contains a reference to an ability score and a bonus   - `ability_score` (ApiReference): Ability score being referenced   - `bonus` (number): The bonus being applied to the ability score - `breath` - A terminal option: Contains a reference to information about a breath attack.   - `name` (string): Name of the breath.   - `dc` (DC): Difficulty check of the breath attack.   - `damage` ([Damage]): Damage dealt by the breath attack, if any. - `damage` - A terminal option. Contains information about damage.   - `damage_type` (ApiReference): Reference to type of damage.   - `damage_dice` (string): Damage expressed in dice (e.g. \"13d6\").   - `notes` (string): Information regarding the damage.  ## FAQ  ### What is the SRD? The SRD, or Systems Reference Document, contains guidelines for publishing content under the OGL. This allows for some of the data for D&D 5e to be open source. The API only covers data that can be found in the SRD. [Here's a link to the full text of the SRD.](https://media.wizards.com/2016/downloads/DND/SRD-OGL_V5.1.pdf)  ### What is the OGL? The Open Game License (OGL) is a public copyright license by Wizards of the Coast that may be used by tabletop role-playing game developers to grant permission to modify, copy, and redistribute some of the content designed for their games, notably game mechanics. However, they must share-alike copies and derivative works. [More information about the OGL can be found here.](https://en.wikipedia.org/wiki/Open_Game_License)  ### A monster, spell, subclass, etc. is missing from the API / Database. Can I add it? Please check if the data is within the SRD. If it is, feel free to open an issue or PR to add it yourself. Otherwise, due to legal reasons, we cannot add it.  ### Can this API be self hosted? Yes it can! You can also host the data yourself if you don't want to use the API at all. You can also make changes and add extra data if you like. However, it is up to you to merge in new changes to the data and API.  #### Can I publish is on <insert platform>? Is this free use? Yes, you can. The API itself is under the [MIT license](https://opensource.org/licenses/MIT), and the underlying data accessible via the API is supported under the SRD and OGL.  # Status Page  The status page for the API can be found here: https://5e-bits.github.io/dnd-uptime/  # Chat  Come hang out with us [on Discord](https://discord.gg/TQuYTv7)!  # Contribute  This API is built from two repositories.   - The repo containing the data lives here: https://github.com/bagelbits/5e-database   - The repo with the API implementation lives here: https://github.com/bagelbits/5e-srd-api  This is a evolving API and having fresh ideas are always welcome! You can open an issue in either repo, open a PR for changes, or just discuss with other users in this discord. 

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// ClassLevelClassSpecific Class specific information such as dice values for bard songs and number of warlock invocations.
type ClassLevelClassSpecific struct {
	ClassLevelClassSpecificAnyOf *ClassLevelClassSpecificAnyOf
	ClassLevelClassSpecificAnyOf1 *ClassLevelClassSpecificAnyOf1
	ClassLevelClassSpecificAnyOf10 *ClassLevelClassSpecificAnyOf10
	ClassLevelClassSpecificAnyOf11 *ClassLevelClassSpecificAnyOf11
	ClassLevelClassSpecificAnyOf2 *ClassLevelClassSpecificAnyOf2
	ClassLevelClassSpecificAnyOf3 *ClassLevelClassSpecificAnyOf3
	ClassLevelClassSpecificAnyOf4 *ClassLevelClassSpecificAnyOf4
	ClassLevelClassSpecificAnyOf5 *ClassLevelClassSpecificAnyOf5
	ClassLevelClassSpecificAnyOf6 *ClassLevelClassSpecificAnyOf6
	ClassLevelClassSpecificAnyOf7 *ClassLevelClassSpecificAnyOf7
	ClassLevelClassSpecificAnyOf8 *ClassLevelClassSpecificAnyOf8
	ClassLevelClassSpecificAnyOf9 *ClassLevelClassSpecificAnyOf9
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ClassLevelClassSpecific) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ClassLevelClassSpecificAnyOf
	err = json.Unmarshal(data, &dst.ClassLevelClassSpecificAnyOf);
	if err == nil {
		jsonClassLevelClassSpecificAnyOf, _ := json.Marshal(dst.ClassLevelClassSpecificAnyOf)
		if string(jsonClassLevelClassSpecificAnyOf) == "{}" { // empty struct
			dst.ClassLevelClassSpecificAnyOf = nil
		} else {
			return nil // data stored in dst.ClassLevelClassSpecificAnyOf, return on the first match
		}
	} else {
		dst.ClassLevelClassSpecificAnyOf = nil
	}

	// try to unmarshal JSON data into ClassLevelClassSpecificAnyOf1
	err = json.Unmarshal(data, &dst.ClassLevelClassSpecificAnyOf1);
	if err == nil {
		jsonClassLevelClassSpecificAnyOf1, _ := json.Marshal(dst.ClassLevelClassSpecificAnyOf1)
		if string(jsonClassLevelClassSpecificAnyOf1) == "{}" { // empty struct
			dst.ClassLevelClassSpecificAnyOf1 = nil
		} else {
			return nil // data stored in dst.ClassLevelClassSpecificAnyOf1, return on the first match
		}
	} else {
		dst.ClassLevelClassSpecificAnyOf1 = nil
	}

	// try to unmarshal JSON data into ClassLevelClassSpecificAnyOf10
	err = json.Unmarshal(data, &dst.ClassLevelClassSpecificAnyOf10);
	if err == nil {
		jsonClassLevelClassSpecificAnyOf10, _ := json.Marshal(dst.ClassLevelClassSpecificAnyOf10)
		if string(jsonClassLevelClassSpecificAnyOf10) == "{}" { // empty struct
			dst.ClassLevelClassSpecificAnyOf10 = nil
		} else {
			return nil // data stored in dst.ClassLevelClassSpecificAnyOf10, return on the first match
		}
	} else {
		dst.ClassLevelClassSpecificAnyOf10 = nil
	}

	// try to unmarshal JSON data into ClassLevelClassSpecificAnyOf11
	err = json.Unmarshal(data, &dst.ClassLevelClassSpecificAnyOf11);
	if err == nil {
		jsonClassLevelClassSpecificAnyOf11, _ := json.Marshal(dst.ClassLevelClassSpecificAnyOf11)
		if string(jsonClassLevelClassSpecificAnyOf11) == "{}" { // empty struct
			dst.ClassLevelClassSpecificAnyOf11 = nil
		} else {
			return nil // data stored in dst.ClassLevelClassSpecificAnyOf11, return on the first match
		}
	} else {
		dst.ClassLevelClassSpecificAnyOf11 = nil
	}

	// try to unmarshal JSON data into ClassLevelClassSpecificAnyOf2
	err = json.Unmarshal(data, &dst.ClassLevelClassSpecificAnyOf2);
	if err == nil {
		jsonClassLevelClassSpecificAnyOf2, _ := json.Marshal(dst.ClassLevelClassSpecificAnyOf2)
		if string(jsonClassLevelClassSpecificAnyOf2) == "{}" { // empty struct
			dst.ClassLevelClassSpecificAnyOf2 = nil
		} else {
			return nil // data stored in dst.ClassLevelClassSpecificAnyOf2, return on the first match
		}
	} else {
		dst.ClassLevelClassSpecificAnyOf2 = nil
	}

	// try to unmarshal JSON data into ClassLevelClassSpecificAnyOf3
	err = json.Unmarshal(data, &dst.ClassLevelClassSpecificAnyOf3);
	if err == nil {
		jsonClassLevelClassSpecificAnyOf3, _ := json.Marshal(dst.ClassLevelClassSpecificAnyOf3)
		if string(jsonClassLevelClassSpecificAnyOf3) == "{}" { // empty struct
			dst.ClassLevelClassSpecificAnyOf3 = nil
		} else {
			return nil // data stored in dst.ClassLevelClassSpecificAnyOf3, return on the first match
		}
	} else {
		dst.ClassLevelClassSpecificAnyOf3 = nil
	}

	// try to unmarshal JSON data into ClassLevelClassSpecificAnyOf4
	err = json.Unmarshal(data, &dst.ClassLevelClassSpecificAnyOf4);
	if err == nil {
		jsonClassLevelClassSpecificAnyOf4, _ := json.Marshal(dst.ClassLevelClassSpecificAnyOf4)
		if string(jsonClassLevelClassSpecificAnyOf4) == "{}" { // empty struct
			dst.ClassLevelClassSpecificAnyOf4 = nil
		} else {
			return nil // data stored in dst.ClassLevelClassSpecificAnyOf4, return on the first match
		}
	} else {
		dst.ClassLevelClassSpecificAnyOf4 = nil
	}

	// try to unmarshal JSON data into ClassLevelClassSpecificAnyOf5
	err = json.Unmarshal(data, &dst.ClassLevelClassSpecificAnyOf5);
	if err == nil {
		jsonClassLevelClassSpecificAnyOf5, _ := json.Marshal(dst.ClassLevelClassSpecificAnyOf5)
		if string(jsonClassLevelClassSpecificAnyOf5) == "{}" { // empty struct
			dst.ClassLevelClassSpecificAnyOf5 = nil
		} else {
			return nil // data stored in dst.ClassLevelClassSpecificAnyOf5, return on the first match
		}
	} else {
		dst.ClassLevelClassSpecificAnyOf5 = nil
	}

	// try to unmarshal JSON data into ClassLevelClassSpecificAnyOf6
	err = json.Unmarshal(data, &dst.ClassLevelClassSpecificAnyOf6);
	if err == nil {
		jsonClassLevelClassSpecificAnyOf6, _ := json.Marshal(dst.ClassLevelClassSpecificAnyOf6)
		if string(jsonClassLevelClassSpecificAnyOf6) == "{}" { // empty struct
			dst.ClassLevelClassSpecificAnyOf6 = nil
		} else {
			return nil // data stored in dst.ClassLevelClassSpecificAnyOf6, return on the first match
		}
	} else {
		dst.ClassLevelClassSpecificAnyOf6 = nil
	}

	// try to unmarshal JSON data into ClassLevelClassSpecificAnyOf7
	err = json.Unmarshal(data, &dst.ClassLevelClassSpecificAnyOf7);
	if err == nil {
		jsonClassLevelClassSpecificAnyOf7, _ := json.Marshal(dst.ClassLevelClassSpecificAnyOf7)
		if string(jsonClassLevelClassSpecificAnyOf7) == "{}" { // empty struct
			dst.ClassLevelClassSpecificAnyOf7 = nil
		} else {
			return nil // data stored in dst.ClassLevelClassSpecificAnyOf7, return on the first match
		}
	} else {
		dst.ClassLevelClassSpecificAnyOf7 = nil
	}

	// try to unmarshal JSON data into ClassLevelClassSpecificAnyOf8
	err = json.Unmarshal(data, &dst.ClassLevelClassSpecificAnyOf8);
	if err == nil {
		jsonClassLevelClassSpecificAnyOf8, _ := json.Marshal(dst.ClassLevelClassSpecificAnyOf8)
		if string(jsonClassLevelClassSpecificAnyOf8) == "{}" { // empty struct
			dst.ClassLevelClassSpecificAnyOf8 = nil
		} else {
			return nil // data stored in dst.ClassLevelClassSpecificAnyOf8, return on the first match
		}
	} else {
		dst.ClassLevelClassSpecificAnyOf8 = nil
	}

	// try to unmarshal JSON data into ClassLevelClassSpecificAnyOf9
	err = json.Unmarshal(data, &dst.ClassLevelClassSpecificAnyOf9);
	if err == nil {
		jsonClassLevelClassSpecificAnyOf9, _ := json.Marshal(dst.ClassLevelClassSpecificAnyOf9)
		if string(jsonClassLevelClassSpecificAnyOf9) == "{}" { // empty struct
			dst.ClassLevelClassSpecificAnyOf9 = nil
		} else {
			return nil // data stored in dst.ClassLevelClassSpecificAnyOf9, return on the first match
		}
	} else {
		dst.ClassLevelClassSpecificAnyOf9 = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ClassLevelClassSpecific)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ClassLevelClassSpecific) MarshalJSON() ([]byte, error) {
	if src.ClassLevelClassSpecificAnyOf != nil {
		return json.Marshal(&src.ClassLevelClassSpecificAnyOf)
	}

	if src.ClassLevelClassSpecificAnyOf1 != nil {
		return json.Marshal(&src.ClassLevelClassSpecificAnyOf1)
	}

	if src.ClassLevelClassSpecificAnyOf10 != nil {
		return json.Marshal(&src.ClassLevelClassSpecificAnyOf10)
	}

	if src.ClassLevelClassSpecificAnyOf11 != nil {
		return json.Marshal(&src.ClassLevelClassSpecificAnyOf11)
	}

	if src.ClassLevelClassSpecificAnyOf2 != nil {
		return json.Marshal(&src.ClassLevelClassSpecificAnyOf2)
	}

	if src.ClassLevelClassSpecificAnyOf3 != nil {
		return json.Marshal(&src.ClassLevelClassSpecificAnyOf3)
	}

	if src.ClassLevelClassSpecificAnyOf4 != nil {
		return json.Marshal(&src.ClassLevelClassSpecificAnyOf4)
	}

	if src.ClassLevelClassSpecificAnyOf5 != nil {
		return json.Marshal(&src.ClassLevelClassSpecificAnyOf5)
	}

	if src.ClassLevelClassSpecificAnyOf6 != nil {
		return json.Marshal(&src.ClassLevelClassSpecificAnyOf6)
	}

	if src.ClassLevelClassSpecificAnyOf7 != nil {
		return json.Marshal(&src.ClassLevelClassSpecificAnyOf7)
	}

	if src.ClassLevelClassSpecificAnyOf8 != nil {
		return json.Marshal(&src.ClassLevelClassSpecificAnyOf8)
	}

	if src.ClassLevelClassSpecificAnyOf9 != nil {
		return json.Marshal(&src.ClassLevelClassSpecificAnyOf9)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableClassLevelClassSpecific struct {
	value *ClassLevelClassSpecific
	isSet bool
}

func (v NullableClassLevelClassSpecific) Get() *ClassLevelClassSpecific {
	return v.value
}

func (v *NullableClassLevelClassSpecific) Set(val *ClassLevelClassSpecific) {
	v.value = val
	v.isSet = true
}

func (v NullableClassLevelClassSpecific) IsSet() bool {
	return v.isSet
}

func (v *NullableClassLevelClassSpecific) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClassLevelClassSpecific(val *ClassLevelClassSpecific) *NullableClassLevelClassSpecific {
	return &NullableClassLevelClassSpecific{value: val, isSet: true}
}

func (v NullableClassLevelClassSpecific) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClassLevelClassSpecific) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


