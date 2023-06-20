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

// checks if the ClassLevel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClassLevel{}

// ClassLevel `ClassLevel` 
type ClassLevel struct {
	// Resource index for shorthand searching.
	Index *string `json:"index,omitempty"`
	// URL of the referenced resource.
	Url *string `json:"url,omitempty"`
	// The number value for the current level object.
	Level *float32 `json:"level,omitempty"`
	// Total number of ability score bonuses gained, added from previous levels.
	AbilityScoreBonuses *float32 `json:"ability_score_bonuses,omitempty"`
	// Proficiency bonus for this class at the specified level.
	ProfBonus *float32 `json:"prof_bonus,omitempty"`
	// Features automatically gained at this level.
	Features []APIReference `json:"features,omitempty"`
	Spellcasting *SubclassLevelSpellcasting `json:"spellcasting,omitempty"`
	ClassSpecific *ClassLevelClassSpecific `json:"class_specific,omitempty"`
}

// NewClassLevel instantiates a new ClassLevel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClassLevel() *ClassLevel {
	this := ClassLevel{}
	return &this
}

// NewClassLevelWithDefaults instantiates a new ClassLevel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClassLevelWithDefaults() *ClassLevel {
	this := ClassLevel{}
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *ClassLevel) GetIndex() string {
	if o == nil || IsNil(o.Index) {
		var ret string
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassLevel) GetIndexOk() (*string, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *ClassLevel) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given string and assigns it to the Index field.
func (o *ClassLevel) SetIndex(v string) {
	o.Index = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ClassLevel) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassLevel) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ClassLevel) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ClassLevel) SetUrl(v string) {
	o.Url = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *ClassLevel) GetLevel() float32 {
	if o == nil || IsNil(o.Level) {
		var ret float32
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassLevel) GetLevelOk() (*float32, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *ClassLevel) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given float32 and assigns it to the Level field.
func (o *ClassLevel) SetLevel(v float32) {
	o.Level = &v
}

// GetAbilityScoreBonuses returns the AbilityScoreBonuses field value if set, zero value otherwise.
func (o *ClassLevel) GetAbilityScoreBonuses() float32 {
	if o == nil || IsNil(o.AbilityScoreBonuses) {
		var ret float32
		return ret
	}
	return *o.AbilityScoreBonuses
}

// GetAbilityScoreBonusesOk returns a tuple with the AbilityScoreBonuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassLevel) GetAbilityScoreBonusesOk() (*float32, bool) {
	if o == nil || IsNil(o.AbilityScoreBonuses) {
		return nil, false
	}
	return o.AbilityScoreBonuses, true
}

// HasAbilityScoreBonuses returns a boolean if a field has been set.
func (o *ClassLevel) HasAbilityScoreBonuses() bool {
	if o != nil && !IsNil(o.AbilityScoreBonuses) {
		return true
	}

	return false
}

// SetAbilityScoreBonuses gets a reference to the given float32 and assigns it to the AbilityScoreBonuses field.
func (o *ClassLevel) SetAbilityScoreBonuses(v float32) {
	o.AbilityScoreBonuses = &v
}

// GetProfBonus returns the ProfBonus field value if set, zero value otherwise.
func (o *ClassLevel) GetProfBonus() float32 {
	if o == nil || IsNil(o.ProfBonus) {
		var ret float32
		return ret
	}
	return *o.ProfBonus
}

// GetProfBonusOk returns a tuple with the ProfBonus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassLevel) GetProfBonusOk() (*float32, bool) {
	if o == nil || IsNil(o.ProfBonus) {
		return nil, false
	}
	return o.ProfBonus, true
}

// HasProfBonus returns a boolean if a field has been set.
func (o *ClassLevel) HasProfBonus() bool {
	if o != nil && !IsNil(o.ProfBonus) {
		return true
	}

	return false
}

// SetProfBonus gets a reference to the given float32 and assigns it to the ProfBonus field.
func (o *ClassLevel) SetProfBonus(v float32) {
	o.ProfBonus = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *ClassLevel) GetFeatures() []APIReference {
	if o == nil || IsNil(o.Features) {
		var ret []APIReference
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassLevel) GetFeaturesOk() ([]APIReference, bool) {
	if o == nil || IsNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *ClassLevel) HasFeatures() bool {
	if o != nil && !IsNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []APIReference and assigns it to the Features field.
func (o *ClassLevel) SetFeatures(v []APIReference) {
	o.Features = v
}

// GetSpellcasting returns the Spellcasting field value if set, zero value otherwise.
func (o *ClassLevel) GetSpellcasting() SubclassLevelSpellcasting {
	if o == nil || IsNil(o.Spellcasting) {
		var ret SubclassLevelSpellcasting
		return ret
	}
	return *o.Spellcasting
}

// GetSpellcastingOk returns a tuple with the Spellcasting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassLevel) GetSpellcastingOk() (*SubclassLevelSpellcasting, bool) {
	if o == nil || IsNil(o.Spellcasting) {
		return nil, false
	}
	return o.Spellcasting, true
}

// HasSpellcasting returns a boolean if a field has been set.
func (o *ClassLevel) HasSpellcasting() bool {
	if o != nil && !IsNil(o.Spellcasting) {
		return true
	}

	return false
}

// SetSpellcasting gets a reference to the given SubclassLevelSpellcasting and assigns it to the Spellcasting field.
func (o *ClassLevel) SetSpellcasting(v SubclassLevelSpellcasting) {
	o.Spellcasting = &v
}

// GetClassSpecific returns the ClassSpecific field value if set, zero value otherwise.
func (o *ClassLevel) GetClassSpecific() ClassLevelClassSpecific {
	if o == nil || IsNil(o.ClassSpecific) {
		var ret ClassLevelClassSpecific
		return ret
	}
	return *o.ClassSpecific
}

// GetClassSpecificOk returns a tuple with the ClassSpecific field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassLevel) GetClassSpecificOk() (*ClassLevelClassSpecific, bool) {
	if o == nil || IsNil(o.ClassSpecific) {
		return nil, false
	}
	return o.ClassSpecific, true
}

// HasClassSpecific returns a boolean if a field has been set.
func (o *ClassLevel) HasClassSpecific() bool {
	if o != nil && !IsNil(o.ClassSpecific) {
		return true
	}

	return false
}

// SetClassSpecific gets a reference to the given ClassLevelClassSpecific and assigns it to the ClassSpecific field.
func (o *ClassLevel) SetClassSpecific(v ClassLevelClassSpecific) {
	o.ClassSpecific = &v
}

func (o ClassLevel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClassLevel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	if !IsNil(o.AbilityScoreBonuses) {
		toSerialize["ability_score_bonuses"] = o.AbilityScoreBonuses
	}
	if !IsNil(o.ProfBonus) {
		toSerialize["prof_bonus"] = o.ProfBonus
	}
	if !IsNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	if !IsNil(o.Spellcasting) {
		toSerialize["spellcasting"] = o.Spellcasting
	}
	if !IsNil(o.ClassSpecific) {
		toSerialize["class_specific"] = o.ClassSpecific
	}
	return toSerialize, nil
}

type NullableClassLevel struct {
	value *ClassLevel
	isSet bool
}

func (v NullableClassLevel) Get() *ClassLevel {
	return v.value
}

func (v *NullableClassLevel) Set(val *ClassLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableClassLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableClassLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClassLevel(val *ClassLevel) *NullableClassLevel {
	return &NullableClassLevel{value: val, isSet: true}
}

func (v NullableClassLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClassLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


