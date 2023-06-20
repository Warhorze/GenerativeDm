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

// Option - `Option` 
type Option struct {
	OptionOneOf *OptionOneOf
	OptionOneOf1 *OptionOneOf1
	OptionOneOf10 *OptionOneOf10
	OptionOneOf2 *OptionOneOf2
	OptionOneOf3 *OptionOneOf3
	OptionOneOf4 *OptionOneOf4
	OptionOneOf5 *OptionOneOf5
	OptionOneOf6 *OptionOneOf6
	OptionOneOf7 *OptionOneOf7
	OptionOneOf8 *OptionOneOf8
	OptionOneOf9 *OptionOneOf9
}

// OptionOneOfAsOption is a convenience function that returns OptionOneOf wrapped in Option
func OptionOneOfAsOption(v *OptionOneOf) Option {
	return Option{
		OptionOneOf: v,
	}
}

// OptionOneOf1AsOption is a convenience function that returns OptionOneOf1 wrapped in Option
func OptionOneOf1AsOption(v *OptionOneOf1) Option {
	return Option{
		OptionOneOf1: v,
	}
}

// OptionOneOf10AsOption is a convenience function that returns OptionOneOf10 wrapped in Option
func OptionOneOf10AsOption(v *OptionOneOf10) Option {
	return Option{
		OptionOneOf10: v,
	}
}

// OptionOneOf2AsOption is a convenience function that returns OptionOneOf2 wrapped in Option
func OptionOneOf2AsOption(v *OptionOneOf2) Option {
	return Option{
		OptionOneOf2: v,
	}
}

// OptionOneOf3AsOption is a convenience function that returns OptionOneOf3 wrapped in Option
func OptionOneOf3AsOption(v *OptionOneOf3) Option {
	return Option{
		OptionOneOf3: v,
	}
}

// OptionOneOf4AsOption is a convenience function that returns OptionOneOf4 wrapped in Option
func OptionOneOf4AsOption(v *OptionOneOf4) Option {
	return Option{
		OptionOneOf4: v,
	}
}

// OptionOneOf5AsOption is a convenience function that returns OptionOneOf5 wrapped in Option
func OptionOneOf5AsOption(v *OptionOneOf5) Option {
	return Option{
		OptionOneOf5: v,
	}
}

// OptionOneOf6AsOption is a convenience function that returns OptionOneOf6 wrapped in Option
func OptionOneOf6AsOption(v *OptionOneOf6) Option {
	return Option{
		OptionOneOf6: v,
	}
}

// OptionOneOf7AsOption is a convenience function that returns OptionOneOf7 wrapped in Option
func OptionOneOf7AsOption(v *OptionOneOf7) Option {
	return Option{
		OptionOneOf7: v,
	}
}

// OptionOneOf8AsOption is a convenience function that returns OptionOneOf8 wrapped in Option
func OptionOneOf8AsOption(v *OptionOneOf8) Option {
	return Option{
		OptionOneOf8: v,
	}
}

// OptionOneOf9AsOption is a convenience function that returns OptionOneOf9 wrapped in Option
func OptionOneOf9AsOption(v *OptionOneOf9) Option {
	return Option{
		OptionOneOf9: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *Option) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into OptionOneOf
	err = newStrictDecoder(data).Decode(&dst.OptionOneOf)
	if err == nil {
		jsonOptionOneOf, _ := json.Marshal(dst.OptionOneOf)
		if string(jsonOptionOneOf) == "{}" { // empty struct
			dst.OptionOneOf = nil
		} else {
			match++
		}
	} else {
		dst.OptionOneOf = nil
	}

	// try to unmarshal data into OptionOneOf1
	err = newStrictDecoder(data).Decode(&dst.OptionOneOf1)
	if err == nil {
		jsonOptionOneOf1, _ := json.Marshal(dst.OptionOneOf1)
		if string(jsonOptionOneOf1) == "{}" { // empty struct
			dst.OptionOneOf1 = nil
		} else {
			match++
		}
	} else {
		dst.OptionOneOf1 = nil
	}

	// try to unmarshal data into OptionOneOf10
	err = newStrictDecoder(data).Decode(&dst.OptionOneOf10)
	if err == nil {
		jsonOptionOneOf10, _ := json.Marshal(dst.OptionOneOf10)
		if string(jsonOptionOneOf10) == "{}" { // empty struct
			dst.OptionOneOf10 = nil
		} else {
			match++
		}
	} else {
		dst.OptionOneOf10 = nil
	}

	// try to unmarshal data into OptionOneOf2
	err = newStrictDecoder(data).Decode(&dst.OptionOneOf2)
	if err == nil {
		jsonOptionOneOf2, _ := json.Marshal(dst.OptionOneOf2)
		if string(jsonOptionOneOf2) == "{}" { // empty struct
			dst.OptionOneOf2 = nil
		} else {
			match++
		}
	} else {
		dst.OptionOneOf2 = nil
	}

	// try to unmarshal data into OptionOneOf3
	err = newStrictDecoder(data).Decode(&dst.OptionOneOf3)
	if err == nil {
		jsonOptionOneOf3, _ := json.Marshal(dst.OptionOneOf3)
		if string(jsonOptionOneOf3) == "{}" { // empty struct
			dst.OptionOneOf3 = nil
		} else {
			match++
		}
	} else {
		dst.OptionOneOf3 = nil
	}

	// try to unmarshal data into OptionOneOf4
	err = newStrictDecoder(data).Decode(&dst.OptionOneOf4)
	if err == nil {
		jsonOptionOneOf4, _ := json.Marshal(dst.OptionOneOf4)
		if string(jsonOptionOneOf4) == "{}" { // empty struct
			dst.OptionOneOf4 = nil
		} else {
			match++
		}
	} else {
		dst.OptionOneOf4 = nil
	}

	// try to unmarshal data into OptionOneOf5
	err = newStrictDecoder(data).Decode(&dst.OptionOneOf5)
	if err == nil {
		jsonOptionOneOf5, _ := json.Marshal(dst.OptionOneOf5)
		if string(jsonOptionOneOf5) == "{}" { // empty struct
			dst.OptionOneOf5 = nil
		} else {
			match++
		}
	} else {
		dst.OptionOneOf5 = nil
	}

	// try to unmarshal data into OptionOneOf6
	err = newStrictDecoder(data).Decode(&dst.OptionOneOf6)
	if err == nil {
		jsonOptionOneOf6, _ := json.Marshal(dst.OptionOneOf6)
		if string(jsonOptionOneOf6) == "{}" { // empty struct
			dst.OptionOneOf6 = nil
		} else {
			match++
		}
	} else {
		dst.OptionOneOf6 = nil
	}

	// try to unmarshal data into OptionOneOf7
	err = newStrictDecoder(data).Decode(&dst.OptionOneOf7)
	if err == nil {
		jsonOptionOneOf7, _ := json.Marshal(dst.OptionOneOf7)
		if string(jsonOptionOneOf7) == "{}" { // empty struct
			dst.OptionOneOf7 = nil
		} else {
			match++
		}
	} else {
		dst.OptionOneOf7 = nil
	}

	// try to unmarshal data into OptionOneOf8
	err = newStrictDecoder(data).Decode(&dst.OptionOneOf8)
	if err == nil {
		jsonOptionOneOf8, _ := json.Marshal(dst.OptionOneOf8)
		if string(jsonOptionOneOf8) == "{}" { // empty struct
			dst.OptionOneOf8 = nil
		} else {
			match++
		}
	} else {
		dst.OptionOneOf8 = nil
	}

	// try to unmarshal data into OptionOneOf9
	err = newStrictDecoder(data).Decode(&dst.OptionOneOf9)
	if err == nil {
		jsonOptionOneOf9, _ := json.Marshal(dst.OptionOneOf9)
		if string(jsonOptionOneOf9) == "{}" { // empty struct
			dst.OptionOneOf9 = nil
		} else {
			match++
		}
	} else {
		dst.OptionOneOf9 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.OptionOneOf = nil
		dst.OptionOneOf1 = nil
		dst.OptionOneOf10 = nil
		dst.OptionOneOf2 = nil
		dst.OptionOneOf3 = nil
		dst.OptionOneOf4 = nil
		dst.OptionOneOf5 = nil
		dst.OptionOneOf6 = nil
		dst.OptionOneOf7 = nil
		dst.OptionOneOf8 = nil
		dst.OptionOneOf9 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Option)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Option)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Option) MarshalJSON() ([]byte, error) {
	if src.OptionOneOf != nil {
		return json.Marshal(&src.OptionOneOf)
	}

	if src.OptionOneOf1 != nil {
		return json.Marshal(&src.OptionOneOf1)
	}

	if src.OptionOneOf10 != nil {
		return json.Marshal(&src.OptionOneOf10)
	}

	if src.OptionOneOf2 != nil {
		return json.Marshal(&src.OptionOneOf2)
	}

	if src.OptionOneOf3 != nil {
		return json.Marshal(&src.OptionOneOf3)
	}

	if src.OptionOneOf4 != nil {
		return json.Marshal(&src.OptionOneOf4)
	}

	if src.OptionOneOf5 != nil {
		return json.Marshal(&src.OptionOneOf5)
	}

	if src.OptionOneOf6 != nil {
		return json.Marshal(&src.OptionOneOf6)
	}

	if src.OptionOneOf7 != nil {
		return json.Marshal(&src.OptionOneOf7)
	}

	if src.OptionOneOf8 != nil {
		return json.Marshal(&src.OptionOneOf8)
	}

	if src.OptionOneOf9 != nil {
		return json.Marshal(&src.OptionOneOf9)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Option) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.OptionOneOf != nil {
		return obj.OptionOneOf
	}

	if obj.OptionOneOf1 != nil {
		return obj.OptionOneOf1
	}

	if obj.OptionOneOf10 != nil {
		return obj.OptionOneOf10
	}

	if obj.OptionOneOf2 != nil {
		return obj.OptionOneOf2
	}

	if obj.OptionOneOf3 != nil {
		return obj.OptionOneOf3
	}

	if obj.OptionOneOf4 != nil {
		return obj.OptionOneOf4
	}

	if obj.OptionOneOf5 != nil {
		return obj.OptionOneOf5
	}

	if obj.OptionOneOf6 != nil {
		return obj.OptionOneOf6
	}

	if obj.OptionOneOf7 != nil {
		return obj.OptionOneOf7
	}

	if obj.OptionOneOf8 != nil {
		return obj.OptionOneOf8
	}

	if obj.OptionOneOf9 != nil {
		return obj.OptionOneOf9
	}

	// all schemas are nil
	return nil
}

type NullableOption struct {
	value *Option
	isSet bool
}

func (v NullableOption) Get() *Option {
	return v.value
}

func (v *NullableOption) Set(val *Option) {
	v.value = val
	v.isSet = true
}

func (v NullableOption) IsSet() bool {
	return v.isSet
}

func (v *NullableOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOption(val *Option) *NullableOption {
	return &NullableOption{value: val, isSet: true}
}

func (v NullableOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


