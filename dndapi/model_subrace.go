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

// checks if the Subrace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Subrace{}

// Subrace `Subrace` 
type Subrace struct {
	// Resource index for shorthand searching.
	Index *string `json:"index,omitempty"`
	// Name of the referenced resource.
	Name *string `json:"name,omitempty"`
	// URL of the referenced resource.
	Url *string `json:"url,omitempty"`
	// Description of the subrace.
	Desc *string `json:"desc,omitempty"`
	Race *SubraceAllOfRace `json:"race,omitempty"`
	// Additional ability bonuses for the subrace.
	AbilityBonuses []AbilityBonus `json:"ability_bonuses,omitempty"`
	// Starting proficiencies for all new characters of the subrace.
	StartingProficiencies []APIReference `json:"starting_proficiencies,omitempty"`
	// Starting languages for all new characters of the subrace.
	Languages []APIReference `json:"languages,omitempty"`
	LanguageOptions *Choice `json:"language_options,omitempty"`
	// List of traits that for the subrace.
	RacialTraits []APIReference `json:"racial_traits,omitempty"`
}

// NewSubrace instantiates a new Subrace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubrace() *Subrace {
	this := Subrace{}
	return &this
}

// NewSubraceWithDefaults instantiates a new Subrace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubraceWithDefaults() *Subrace {
	this := Subrace{}
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *Subrace) GetIndex() string {
	if o == nil || IsNil(o.Index) {
		var ret string
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subrace) GetIndexOk() (*string, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *Subrace) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given string and assigns it to the Index field.
func (o *Subrace) SetIndex(v string) {
	o.Index = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Subrace) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subrace) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Subrace) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Subrace) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Subrace) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subrace) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Subrace) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Subrace) SetUrl(v string) {
	o.Url = &v
}

// GetDesc returns the Desc field value if set, zero value otherwise.
func (o *Subrace) GetDesc() string {
	if o == nil || IsNil(o.Desc) {
		var ret string
		return ret
	}
	return *o.Desc
}

// GetDescOk returns a tuple with the Desc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subrace) GetDescOk() (*string, bool) {
	if o == nil || IsNil(o.Desc) {
		return nil, false
	}
	return o.Desc, true
}

// HasDesc returns a boolean if a field has been set.
func (o *Subrace) HasDesc() bool {
	if o != nil && !IsNil(o.Desc) {
		return true
	}

	return false
}

// SetDesc gets a reference to the given string and assigns it to the Desc field.
func (o *Subrace) SetDesc(v string) {
	o.Desc = &v
}

// GetRace returns the Race field value if set, zero value otherwise.
func (o *Subrace) GetRace() SubraceAllOfRace {
	if o == nil || IsNil(o.Race) {
		var ret SubraceAllOfRace
		return ret
	}
	return *o.Race
}

// GetRaceOk returns a tuple with the Race field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subrace) GetRaceOk() (*SubraceAllOfRace, bool) {
	if o == nil || IsNil(o.Race) {
		return nil, false
	}
	return o.Race, true
}

// HasRace returns a boolean if a field has been set.
func (o *Subrace) HasRace() bool {
	if o != nil && !IsNil(o.Race) {
		return true
	}

	return false
}

// SetRace gets a reference to the given SubraceAllOfRace and assigns it to the Race field.
func (o *Subrace) SetRace(v SubraceAllOfRace) {
	o.Race = &v
}

// GetAbilityBonuses returns the AbilityBonuses field value if set, zero value otherwise.
func (o *Subrace) GetAbilityBonuses() []AbilityBonus {
	if o == nil || IsNil(o.AbilityBonuses) {
		var ret []AbilityBonus
		return ret
	}
	return o.AbilityBonuses
}

// GetAbilityBonusesOk returns a tuple with the AbilityBonuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subrace) GetAbilityBonusesOk() ([]AbilityBonus, bool) {
	if o == nil || IsNil(o.AbilityBonuses) {
		return nil, false
	}
	return o.AbilityBonuses, true
}

// HasAbilityBonuses returns a boolean if a field has been set.
func (o *Subrace) HasAbilityBonuses() bool {
	if o != nil && !IsNil(o.AbilityBonuses) {
		return true
	}

	return false
}

// SetAbilityBonuses gets a reference to the given []AbilityBonus and assigns it to the AbilityBonuses field.
func (o *Subrace) SetAbilityBonuses(v []AbilityBonus) {
	o.AbilityBonuses = v
}

// GetStartingProficiencies returns the StartingProficiencies field value if set, zero value otherwise.
func (o *Subrace) GetStartingProficiencies() []APIReference {
	if o == nil || IsNil(o.StartingProficiencies) {
		var ret []APIReference
		return ret
	}
	return o.StartingProficiencies
}

// GetStartingProficienciesOk returns a tuple with the StartingProficiencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subrace) GetStartingProficienciesOk() ([]APIReference, bool) {
	if o == nil || IsNil(o.StartingProficiencies) {
		return nil, false
	}
	return o.StartingProficiencies, true
}

// HasStartingProficiencies returns a boolean if a field has been set.
func (o *Subrace) HasStartingProficiencies() bool {
	if o != nil && !IsNil(o.StartingProficiencies) {
		return true
	}

	return false
}

// SetStartingProficiencies gets a reference to the given []APIReference and assigns it to the StartingProficiencies field.
func (o *Subrace) SetStartingProficiencies(v []APIReference) {
	o.StartingProficiencies = v
}

// GetLanguages returns the Languages field value if set, zero value otherwise.
func (o *Subrace) GetLanguages() []APIReference {
	if o == nil || IsNil(o.Languages) {
		var ret []APIReference
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subrace) GetLanguagesOk() ([]APIReference, bool) {
	if o == nil || IsNil(o.Languages) {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *Subrace) HasLanguages() bool {
	if o != nil && !IsNil(o.Languages) {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []APIReference and assigns it to the Languages field.
func (o *Subrace) SetLanguages(v []APIReference) {
	o.Languages = v
}

// GetLanguageOptions returns the LanguageOptions field value if set, zero value otherwise.
func (o *Subrace) GetLanguageOptions() Choice {
	if o == nil || IsNil(o.LanguageOptions) {
		var ret Choice
		return ret
	}
	return *o.LanguageOptions
}

// GetLanguageOptionsOk returns a tuple with the LanguageOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subrace) GetLanguageOptionsOk() (*Choice, bool) {
	if o == nil || IsNil(o.LanguageOptions) {
		return nil, false
	}
	return o.LanguageOptions, true
}

// HasLanguageOptions returns a boolean if a field has been set.
func (o *Subrace) HasLanguageOptions() bool {
	if o != nil && !IsNil(o.LanguageOptions) {
		return true
	}

	return false
}

// SetLanguageOptions gets a reference to the given Choice and assigns it to the LanguageOptions field.
func (o *Subrace) SetLanguageOptions(v Choice) {
	o.LanguageOptions = &v
}

// GetRacialTraits returns the RacialTraits field value if set, zero value otherwise.
func (o *Subrace) GetRacialTraits() []APIReference {
	if o == nil || IsNil(o.RacialTraits) {
		var ret []APIReference
		return ret
	}
	return o.RacialTraits
}

// GetRacialTraitsOk returns a tuple with the RacialTraits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subrace) GetRacialTraitsOk() ([]APIReference, bool) {
	if o == nil || IsNil(o.RacialTraits) {
		return nil, false
	}
	return o.RacialTraits, true
}

// HasRacialTraits returns a boolean if a field has been set.
func (o *Subrace) HasRacialTraits() bool {
	if o != nil && !IsNil(o.RacialTraits) {
		return true
	}

	return false
}

// SetRacialTraits gets a reference to the given []APIReference and assigns it to the RacialTraits field.
func (o *Subrace) SetRacialTraits(v []APIReference) {
	o.RacialTraits = v
}

func (o Subrace) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Subrace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Desc) {
		toSerialize["desc"] = o.Desc
	}
	if !IsNil(o.Race) {
		toSerialize["race"] = o.Race
	}
	if !IsNil(o.AbilityBonuses) {
		toSerialize["ability_bonuses"] = o.AbilityBonuses
	}
	if !IsNil(o.StartingProficiencies) {
		toSerialize["starting_proficiencies"] = o.StartingProficiencies
	}
	if !IsNil(o.Languages) {
		toSerialize["languages"] = o.Languages
	}
	if !IsNil(o.LanguageOptions) {
		toSerialize["language_options"] = o.LanguageOptions
	}
	if !IsNil(o.RacialTraits) {
		toSerialize["racial_traits"] = o.RacialTraits
	}
	return toSerialize, nil
}

type NullableSubrace struct {
	value *Subrace
	isSet bool
}

func (v NullableSubrace) Get() *Subrace {
	return v.value
}

func (v *NullableSubrace) Set(val *Subrace) {
	v.value = val
	v.isSet = true
}

func (v NullableSubrace) IsSet() bool {
	return v.isSet
}

func (v *NullableSubrace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubrace(val *Subrace) *NullableSubrace {
	return &NullableSubrace{value: val, isSet: true}
}

func (v NullableSubrace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubrace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


