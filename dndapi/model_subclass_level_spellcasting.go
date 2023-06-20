/*
D&D 5e API

# Introduction  Welcome to the dnd5eapi, the Dungeons & Dragons 5th Edition API! This documentation should help you familiarize yourself with the resources available and how to consume them with HTTP requests. Read through the getting started section before you dive in. Most of your problems should be solved just by reading through it.  ## Getting Started  Let's make our first API request to the D&D 5th Edition API!  Open up a terminal and use [curl](http://curl.haxx.se/) or [httpie](http://httpie.org/) to make an API request for a resource. You can also scroll through the definitions below and send requests directly from the endpoint documentation!  For example, if you paste and run this `curl` command: ```bash curl -X GET \"https://www.dnd5eapi.co/api/ability-scores/cha\" -H \"Accept: application/json\" ```  We should see a result containing details about the Charisma ability score: ```bash {   \"index\": \"cha\",   \"name\": \"CHA\",   \"full_name\": \"Charisma\",   \"desc\": [     \"Charisma measures your ability to interact effectively with others. It       includes such factors as confidence and eloquence, and it can represent       a charming or commanding personality.\",     \"A Charisma check might arise when you try to influence or entertain       others, when you try to make an impression or tell a convincing lie,       or when you are navigating a tricky social situation. The Deception,       Intimidation, Performance, and Persuasion skills reflect aptitude in       certain kinds of Charisma checks.\"   ],   \"skills\": [     {       \"name\": \"Deception\",       \"index\": \"deception\",       \"url\": \"/api/skills/deception\"     },     {       \"name\": \"Intimidation\",       \"index\": \"intimidation\",       \"url\": \"/api/skills/intimidation\"     },     {       \"name\": \"Performance\",       \"index\": \"performance\",       \"url\": \"/api/skills/performance\"     },     {       \"name\": \"Persuasion\",       \"index\": \"persuasion\",       \"url\": \"/api/skills/persuasion\"     }   ],   \"url\": \"/api/ability-scores/cha\" } ```  ## Authentication  The dnd5eapi is a completely open API. No authentication is required to query and get data. This also means that we've limited what you can do to just `GET`-ing the data. If you find a mistake in the data, feel free to [message us](https://discord.gg/TQuYTv7).  ## GraphQL  This API supports [GraphQL](https://graphql.org/). The GraphQL URL for this API is `https://www.dnd5eapi.co/graphql`. Most of your questions regarding the GraphQL schema can be answered by querying the endpoint with the Apollo sandbox explorer.  ## Schemas  Definitions of all schemas will be accessible in a future update. Two of the most common schemas are described here.  ### `APIReference` Represents a minimal representation of a resource. The detailed representation of the referenced resource can be retrieved by making a request to the referenced `URL`. ``` APIReference {   index     string   name      string   url       string } ``` <hr>  ### `DC` Represents a difficulty check. ``` DC {   dc_type       APIReference   dc_value      number   success_type  \"none\" | \"half\" | \"other\" } ``` <hr>  ### `Damage` Represents damage. ``` Damage {   damage_type     APIReference   damage_dice     string } ``` <hr>  ### `Choice` Represents a choice made by a player. Commonly seen related to decisions made during character creation or combat (e.g.: the description of the cleric class, under **Proficiencies**, states \"Skills: Choose two from History, Insight, Medicine, Persuasion, and Religion\" [[SRD p15]](https://media.wizards.com/2016/downloads/DND/SRD-OGL_V5.1.pdf#page=15)) ``` Choice {   desc      string   choose    number   type      string   from      OptionSet } ``` <hr>  ### `OptionSet` The OptionSet structure provides the options to be chosen from, or sufficient data to fetch and interpret the options. All OptionSets have an `option_set_type` attribute that indicates the structure of the object that contains the options. The possible values are `options_array`, `equipment_category`, and `reference_list`. Other attributes on the OptionSet depend on the value of this attribute. - `options_array`   - `options` (array): An array of Option objects. Each item in the array represents an option that can be chosen. - `equipment_category`   - `equipment_category` (APIReference): A reference to an EquipmentCategory. Each item in the EquipmentCategory's `equipment` array represents one option that can be chosen. - `resource_list`   - `resource_list_url` (string): A reference (by URL) to a collection in the database. The URL may include query parameters. Each item in the resulting ResourceList's `results` array represents one option that can be chosen. <hr>  ### `Option` When the options are given in an `options_array`, each item in the array inherits from the Option structure. All Options have an `option_type` attribute that indicates the structure of the option. The value of this attribute indicates how the option should be handled, and each type has different attributes. The possible values and their corresponding attributes are listed below. - `reference` - A terminal option. Contains a reference to a Document that can be added to the list of options chosen.   - `item` (APIReference): A reference to the chosen item. - `action` - A terminal option. Contains information describing an action, for use within Multiattack actions.   - `action_name` (string): The name of the action, according to its `name` attribute.   - `count` (number | string): The number of times this action can be repeated if this option is chosen.   - `type` (string = `\"melee\" | \"ranged\" | \"ability\" | \"magic\"`, optional): For attack actions that can be either melee, ranged, abilities, or magic. - `multiple` - When this option is chosen, all of its child options are chosen, and must be resolved the same way as a normal option.   - `items` (array): An array of Option objects. All of them must be taken if the option is chosen. - `choice` - A nested choice. If this option is chosen, the Choice structure contained within must be resolved like a normal Choice structure, and the results are the chosen options.   - `choice` (Choice): The Choice to resolve. - `string` - A terminal option. Contains a reference to a string.   - `string` (string): The string. - `ideal` - A terminal option. Contains information about an ideal.   - `desc` (string): A description of the ideal.   - `alignments` (ApiReference[]): A list of alignments of those who might follow the ideal. - `counted_reference` - A terminal option. Contains a reference to something else in the API along with a count.   - `count` (number): Count.   - `of` (ApiReference): Thing being referenced. - `score_prerequisite` - A terminal option. Contains a reference to an ability score and a minimum score.   - `ability_score` (ApiReference): Ability score being referenced.   - `minimum_score` (number): The minimum score required to satisfy the prerequisite. - `ability_bonus` - A terminal option. Contains a reference to an ability score and a bonus   - `ability_score` (ApiReference): Ability score being referenced   - `bonus` (number): The bonus being applied to the ability score - `breath` - A terminal option: Contains a reference to information about a breath attack.   - `name` (string): Name of the breath.   - `dc` (DC): Difficulty check of the breath attack.   - `damage` ([Damage]): Damage dealt by the breath attack, if any. - `damage` - A terminal option. Contains information about damage.   - `damage_type` (ApiReference): Reference to type of damage.   - `damage_dice` (string): Damage expressed in dice (e.g. \"13d6\").   - `notes` (string): Information regarding the damage.  ## FAQ  ### What is the SRD? The SRD, or Systems Reference Document, contains guidelines for publishing content under the OGL. This allows for some of the data for D&D 5e to be open source. The API only covers data that can be found in the SRD. [Here's a link to the full text of the SRD.](https://media.wizards.com/2016/downloads/DND/SRD-OGL_V5.1.pdf)  ### What is the OGL? The Open Game License (OGL) is a public copyright license by Wizards of the Coast that may be used by tabletop role-playing game developers to grant permission to modify, copy, and redistribute some of the content designed for their games, notably game mechanics. However, they must share-alike copies and derivative works. [More information about the OGL can be found here.](https://en.wikipedia.org/wiki/Open_Game_License)  ### A monster, spell, subclass, etc. is missing from the API / Database. Can I add it? Please check if the data is within the SRD. If it is, feel free to open an issue or PR to add it yourself. Otherwise, due to legal reasons, we cannot add it.  ### Can this API be self hosted? Yes it can! You can also host the data yourself if you don't want to use the API at all. You can also make changes and add extra data if you like. However, it is up to you to merge in new changes to the data and API.  #### Can I publish is on <insert platform>? Is this free use? Yes, you can. The API itself is under the [MIT license](https://opensource.org/licenses/MIT), and the underlying data accessible via the API is supported under the SRD and OGL.  # Status Page  The status page for the API can be found here: https://5e-bits.github.io/dnd-uptime/  # Chat  Come hang out with us [on Discord](https://discord.gg/TQuYTv7)!  # Contribute  This API is built from two repositories.   - The repo containing the data lives here: https://github.com/bagelbits/5e-database   - The repo with the API implementation lives here: https://github.com/bagelbits/5e-srd-api  This is a evolving API and having fresh ideas are always welcome! You can open an issue in either repo, open a PR for changes, or just discuss with other users in this discord. 

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the SubclassLevelSpellcasting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubclassLevelSpellcasting{}

// SubclassLevelSpellcasting Summary of spells known at this level.
type SubclassLevelSpellcasting struct {
	CantripsKnown *float32 `json:"cantrips_known,omitempty"`
	SpellsKnown *float32 `json:"spells_known,omitempty"`
	SpellSlotsLevel1 *float32 `json:"spell_slots_level_1,omitempty"`
	SpellSlotsLevel2 *float32 `json:"spell_slots_level_2,omitempty"`
	SpellSlotsLevel3 *float32 `json:"spell_slots_level_3,omitempty"`
	SpellSlotsLevel4 *float32 `json:"spell_slots_level_4,omitempty"`
	SpellSlotsLevel5 *float32 `json:"spell_slots_level_5,omitempty"`
	SpellSlotsLevel6 *float32 `json:"spell_slots_level_6,omitempty"`
	SpellSlotsLevel7 *float32 `json:"spell_slots_level_7,omitempty"`
	SpellSlotsLevel8 *float32 `json:"spell_slots_level_8,omitempty"`
	SpellSlotsLevel9 *float32 `json:"spell_slots_level_9,omitempty"`
}

// NewSubclassLevelSpellcasting instantiates a new SubclassLevelSpellcasting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubclassLevelSpellcasting() *SubclassLevelSpellcasting {
	this := SubclassLevelSpellcasting{}
	return &this
}

// NewSubclassLevelSpellcastingWithDefaults instantiates a new SubclassLevelSpellcasting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubclassLevelSpellcastingWithDefaults() *SubclassLevelSpellcasting {
	this := SubclassLevelSpellcasting{}
	return &this
}

// GetCantripsKnown returns the CantripsKnown field value if set, zero value otherwise.
func (o *SubclassLevelSpellcasting) GetCantripsKnown() float32 {
	if o == nil || IsNil(o.CantripsKnown) {
		var ret float32
		return ret
	}
	return *o.CantripsKnown
}

// GetCantripsKnownOk returns a tuple with the CantripsKnown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubclassLevelSpellcasting) GetCantripsKnownOk() (*float32, bool) {
	if o == nil || IsNil(o.CantripsKnown) {
		return nil, false
	}
	return o.CantripsKnown, true
}

// HasCantripsKnown returns a boolean if a field has been set.
func (o *SubclassLevelSpellcasting) HasCantripsKnown() bool {
	if o != nil && !IsNil(o.CantripsKnown) {
		return true
	}

	return false
}

// SetCantripsKnown gets a reference to the given float32 and assigns it to the CantripsKnown field.
func (o *SubclassLevelSpellcasting) SetCantripsKnown(v float32) {
	o.CantripsKnown = &v
}

// GetSpellsKnown returns the SpellsKnown field value if set, zero value otherwise.
func (o *SubclassLevelSpellcasting) GetSpellsKnown() float32 {
	if o == nil || IsNil(o.SpellsKnown) {
		var ret float32
		return ret
	}
	return *o.SpellsKnown
}

// GetSpellsKnownOk returns a tuple with the SpellsKnown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubclassLevelSpellcasting) GetSpellsKnownOk() (*float32, bool) {
	if o == nil || IsNil(o.SpellsKnown) {
		return nil, false
	}
	return o.SpellsKnown, true
}

// HasSpellsKnown returns a boolean if a field has been set.
func (o *SubclassLevelSpellcasting) HasSpellsKnown() bool {
	if o != nil && !IsNil(o.SpellsKnown) {
		return true
	}

	return false
}

// SetSpellsKnown gets a reference to the given float32 and assigns it to the SpellsKnown field.
func (o *SubclassLevelSpellcasting) SetSpellsKnown(v float32) {
	o.SpellsKnown = &v
}

// GetSpellSlotsLevel1 returns the SpellSlotsLevel1 field value if set, zero value otherwise.
func (o *SubclassLevelSpellcasting) GetSpellSlotsLevel1() float32 {
	if o == nil || IsNil(o.SpellSlotsLevel1) {
		var ret float32
		return ret
	}
	return *o.SpellSlotsLevel1
}

// GetSpellSlotsLevel1Ok returns a tuple with the SpellSlotsLevel1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubclassLevelSpellcasting) GetSpellSlotsLevel1Ok() (*float32, bool) {
	if o == nil || IsNil(o.SpellSlotsLevel1) {
		return nil, false
	}
	return o.SpellSlotsLevel1, true
}

// HasSpellSlotsLevel1 returns a boolean if a field has been set.
func (o *SubclassLevelSpellcasting) HasSpellSlotsLevel1() bool {
	if o != nil && !IsNil(o.SpellSlotsLevel1) {
		return true
	}

	return false
}

// SetSpellSlotsLevel1 gets a reference to the given float32 and assigns it to the SpellSlotsLevel1 field.
func (o *SubclassLevelSpellcasting) SetSpellSlotsLevel1(v float32) {
	o.SpellSlotsLevel1 = &v
}

// GetSpellSlotsLevel2 returns the SpellSlotsLevel2 field value if set, zero value otherwise.
func (o *SubclassLevelSpellcasting) GetSpellSlotsLevel2() float32 {
	if o == nil || IsNil(o.SpellSlotsLevel2) {
		var ret float32
		return ret
	}
	return *o.SpellSlotsLevel2
}

// GetSpellSlotsLevel2Ok returns a tuple with the SpellSlotsLevel2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubclassLevelSpellcasting) GetSpellSlotsLevel2Ok() (*float32, bool) {
	if o == nil || IsNil(o.SpellSlotsLevel2) {
		return nil, false
	}
	return o.SpellSlotsLevel2, true
}

// HasSpellSlotsLevel2 returns a boolean if a field has been set.
func (o *SubclassLevelSpellcasting) HasSpellSlotsLevel2() bool {
	if o != nil && !IsNil(o.SpellSlotsLevel2) {
		return true
	}

	return false
}

// SetSpellSlotsLevel2 gets a reference to the given float32 and assigns it to the SpellSlotsLevel2 field.
func (o *SubclassLevelSpellcasting) SetSpellSlotsLevel2(v float32) {
	o.SpellSlotsLevel2 = &v
}

// GetSpellSlotsLevel3 returns the SpellSlotsLevel3 field value if set, zero value otherwise.
func (o *SubclassLevelSpellcasting) GetSpellSlotsLevel3() float32 {
	if o == nil || IsNil(o.SpellSlotsLevel3) {
		var ret float32
		return ret
	}
	return *o.SpellSlotsLevel3
}

// GetSpellSlotsLevel3Ok returns a tuple with the SpellSlotsLevel3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubclassLevelSpellcasting) GetSpellSlotsLevel3Ok() (*float32, bool) {
	if o == nil || IsNil(o.SpellSlotsLevel3) {
		return nil, false
	}
	return o.SpellSlotsLevel3, true
}

// HasSpellSlotsLevel3 returns a boolean if a field has been set.
func (o *SubclassLevelSpellcasting) HasSpellSlotsLevel3() bool {
	if o != nil && !IsNil(o.SpellSlotsLevel3) {
		return true
	}

	return false
}

// SetSpellSlotsLevel3 gets a reference to the given float32 and assigns it to the SpellSlotsLevel3 field.
func (o *SubclassLevelSpellcasting) SetSpellSlotsLevel3(v float32) {
	o.SpellSlotsLevel3 = &v
}

// GetSpellSlotsLevel4 returns the SpellSlotsLevel4 field value if set, zero value otherwise.
func (o *SubclassLevelSpellcasting) GetSpellSlotsLevel4() float32 {
	if o == nil || IsNil(o.SpellSlotsLevel4) {
		var ret float32
		return ret
	}
	return *o.SpellSlotsLevel4
}

// GetSpellSlotsLevel4Ok returns a tuple with the SpellSlotsLevel4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubclassLevelSpellcasting) GetSpellSlotsLevel4Ok() (*float32, bool) {
	if o == nil || IsNil(o.SpellSlotsLevel4) {
		return nil, false
	}
	return o.SpellSlotsLevel4, true
}

// HasSpellSlotsLevel4 returns a boolean if a field has been set.
func (o *SubclassLevelSpellcasting) HasSpellSlotsLevel4() bool {
	if o != nil && !IsNil(o.SpellSlotsLevel4) {
		return true
	}

	return false
}

// SetSpellSlotsLevel4 gets a reference to the given float32 and assigns it to the SpellSlotsLevel4 field.
func (o *SubclassLevelSpellcasting) SetSpellSlotsLevel4(v float32) {
	o.SpellSlotsLevel4 = &v
}

// GetSpellSlotsLevel5 returns the SpellSlotsLevel5 field value if set, zero value otherwise.
func (o *SubclassLevelSpellcasting) GetSpellSlotsLevel5() float32 {
	if o == nil || IsNil(o.SpellSlotsLevel5) {
		var ret float32
		return ret
	}
	return *o.SpellSlotsLevel5
}

// GetSpellSlotsLevel5Ok returns a tuple with the SpellSlotsLevel5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubclassLevelSpellcasting) GetSpellSlotsLevel5Ok() (*float32, bool) {
	if o == nil || IsNil(o.SpellSlotsLevel5) {
		return nil, false
	}
	return o.SpellSlotsLevel5, true
}

// HasSpellSlotsLevel5 returns a boolean if a field has been set.
func (o *SubclassLevelSpellcasting) HasSpellSlotsLevel5() bool {
	if o != nil && !IsNil(o.SpellSlotsLevel5) {
		return true
	}

	return false
}

// SetSpellSlotsLevel5 gets a reference to the given float32 and assigns it to the SpellSlotsLevel5 field.
func (o *SubclassLevelSpellcasting) SetSpellSlotsLevel5(v float32) {
	o.SpellSlotsLevel5 = &v
}

// GetSpellSlotsLevel6 returns the SpellSlotsLevel6 field value if set, zero value otherwise.
func (o *SubclassLevelSpellcasting) GetSpellSlotsLevel6() float32 {
	if o == nil || IsNil(o.SpellSlotsLevel6) {
		var ret float32
		return ret
	}
	return *o.SpellSlotsLevel6
}

// GetSpellSlotsLevel6Ok returns a tuple with the SpellSlotsLevel6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubclassLevelSpellcasting) GetSpellSlotsLevel6Ok() (*float32, bool) {
	if o == nil || IsNil(o.SpellSlotsLevel6) {
		return nil, false
	}
	return o.SpellSlotsLevel6, true
}

// HasSpellSlotsLevel6 returns a boolean if a field has been set.
func (o *SubclassLevelSpellcasting) HasSpellSlotsLevel6() bool {
	if o != nil && !IsNil(o.SpellSlotsLevel6) {
		return true
	}

	return false
}

// SetSpellSlotsLevel6 gets a reference to the given float32 and assigns it to the SpellSlotsLevel6 field.
func (o *SubclassLevelSpellcasting) SetSpellSlotsLevel6(v float32) {
	o.SpellSlotsLevel6 = &v
}

// GetSpellSlotsLevel7 returns the SpellSlotsLevel7 field value if set, zero value otherwise.
func (o *SubclassLevelSpellcasting) GetSpellSlotsLevel7() float32 {
	if o == nil || IsNil(o.SpellSlotsLevel7) {
		var ret float32
		return ret
	}
	return *o.SpellSlotsLevel7
}

// GetSpellSlotsLevel7Ok returns a tuple with the SpellSlotsLevel7 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubclassLevelSpellcasting) GetSpellSlotsLevel7Ok() (*float32, bool) {
	if o == nil || IsNil(o.SpellSlotsLevel7) {
		return nil, false
	}
	return o.SpellSlotsLevel7, true
}

// HasSpellSlotsLevel7 returns a boolean if a field has been set.
func (o *SubclassLevelSpellcasting) HasSpellSlotsLevel7() bool {
	if o != nil && !IsNil(o.SpellSlotsLevel7) {
		return true
	}

	return false
}

// SetSpellSlotsLevel7 gets a reference to the given float32 and assigns it to the SpellSlotsLevel7 field.
func (o *SubclassLevelSpellcasting) SetSpellSlotsLevel7(v float32) {
	o.SpellSlotsLevel7 = &v
}

// GetSpellSlotsLevel8 returns the SpellSlotsLevel8 field value if set, zero value otherwise.
func (o *SubclassLevelSpellcasting) GetSpellSlotsLevel8() float32 {
	if o == nil || IsNil(o.SpellSlotsLevel8) {
		var ret float32
		return ret
	}
	return *o.SpellSlotsLevel8
}

// GetSpellSlotsLevel8Ok returns a tuple with the SpellSlotsLevel8 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubclassLevelSpellcasting) GetSpellSlotsLevel8Ok() (*float32, bool) {
	if o == nil || IsNil(o.SpellSlotsLevel8) {
		return nil, false
	}
	return o.SpellSlotsLevel8, true
}

// HasSpellSlotsLevel8 returns a boolean if a field has been set.
func (o *SubclassLevelSpellcasting) HasSpellSlotsLevel8() bool {
	if o != nil && !IsNil(o.SpellSlotsLevel8) {
		return true
	}

	return false
}

// SetSpellSlotsLevel8 gets a reference to the given float32 and assigns it to the SpellSlotsLevel8 field.
func (o *SubclassLevelSpellcasting) SetSpellSlotsLevel8(v float32) {
	o.SpellSlotsLevel8 = &v
}

// GetSpellSlotsLevel9 returns the SpellSlotsLevel9 field value if set, zero value otherwise.
func (o *SubclassLevelSpellcasting) GetSpellSlotsLevel9() float32 {
	if o == nil || IsNil(o.SpellSlotsLevel9) {
		var ret float32
		return ret
	}
	return *o.SpellSlotsLevel9
}

// GetSpellSlotsLevel9Ok returns a tuple with the SpellSlotsLevel9 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubclassLevelSpellcasting) GetSpellSlotsLevel9Ok() (*float32, bool) {
	if o == nil || IsNil(o.SpellSlotsLevel9) {
		return nil, false
	}
	return o.SpellSlotsLevel9, true
}

// HasSpellSlotsLevel9 returns a boolean if a field has been set.
func (o *SubclassLevelSpellcasting) HasSpellSlotsLevel9() bool {
	if o != nil && !IsNil(o.SpellSlotsLevel9) {
		return true
	}

	return false
}

// SetSpellSlotsLevel9 gets a reference to the given float32 and assigns it to the SpellSlotsLevel9 field.
func (o *SubclassLevelSpellcasting) SetSpellSlotsLevel9(v float32) {
	o.SpellSlotsLevel9 = &v
}

func (o SubclassLevelSpellcasting) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubclassLevelSpellcasting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CantripsKnown) {
		toSerialize["cantrips_known"] = o.CantripsKnown
	}
	if !IsNil(o.SpellsKnown) {
		toSerialize["spells_known"] = o.SpellsKnown
	}
	if !IsNil(o.SpellSlotsLevel1) {
		toSerialize["spell_slots_level_1"] = o.SpellSlotsLevel1
	}
	if !IsNil(o.SpellSlotsLevel2) {
		toSerialize["spell_slots_level_2"] = o.SpellSlotsLevel2
	}
	if !IsNil(o.SpellSlotsLevel3) {
		toSerialize["spell_slots_level_3"] = o.SpellSlotsLevel3
	}
	if !IsNil(o.SpellSlotsLevel4) {
		toSerialize["spell_slots_level_4"] = o.SpellSlotsLevel4
	}
	if !IsNil(o.SpellSlotsLevel5) {
		toSerialize["spell_slots_level_5"] = o.SpellSlotsLevel5
	}
	if !IsNil(o.SpellSlotsLevel6) {
		toSerialize["spell_slots_level_6"] = o.SpellSlotsLevel6
	}
	if !IsNil(o.SpellSlotsLevel7) {
		toSerialize["spell_slots_level_7"] = o.SpellSlotsLevel7
	}
	if !IsNil(o.SpellSlotsLevel8) {
		toSerialize["spell_slots_level_8"] = o.SpellSlotsLevel8
	}
	if !IsNil(o.SpellSlotsLevel9) {
		toSerialize["spell_slots_level_9"] = o.SpellSlotsLevel9
	}
	return toSerialize, nil
}

type NullableSubclassLevelSpellcasting struct {
	value *SubclassLevelSpellcasting
	isSet bool
}

func (v NullableSubclassLevelSpellcasting) Get() *SubclassLevelSpellcasting {
	return v.value
}

func (v *NullableSubclassLevelSpellcasting) Set(val *SubclassLevelSpellcasting) {
	v.value = val
	v.isSet = true
}

func (v NullableSubclassLevelSpellcasting) IsSet() bool {
	return v.isSet
}

func (v *NullableSubclassLevelSpellcasting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubclassLevelSpellcasting(val *SubclassLevelSpellcasting) *NullableSubclassLevelSpellcasting {
	return &NullableSubclassLevelSpellcasting{value: val, isSet: true}
}

func (v NullableSubclassLevelSpellcasting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubclassLevelSpellcasting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


