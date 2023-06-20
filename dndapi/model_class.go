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

// checks if the Class type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Class{}

// Class `Class` 
type Class struct {
	// Resource index for shorthand searching.
	Index *string `json:"index,omitempty"`
	// Name of the referenced resource.
	Name *string `json:"name,omitempty"`
	// URL of the referenced resource.
	Url *string `json:"url,omitempty"`
	// Hit die of the class. (ex: 12 == 1d12).
	HitDie *float32 `json:"hit_die,omitempty"`
	// URL of the level resource for the class.
	ClassLevels *string `json:"class_levels,omitempty"`
	MultiClassing *Multiclassing `json:"multi_classing,omitempty"`
	Spellcasting *Spellcasting `json:"spellcasting,omitempty"`
	// URL of the spell resource list for the class.
	Spells *string `json:"spells,omitempty"`
	// List of equipment and their quantities all players of the class start with.
	StartingEquipment []ClassAllOfStartingEquipment `json:"starting_equipment,omitempty"`
	// List of choices of starting equipment.
	StartingEquipmentOptions []Choice `json:"starting_equipment_options,omitempty"`
	// List of choices of starting proficiencies.
	ProficiencyChoices []Choice `json:"proficiency_choices,omitempty"`
	// List of starting proficiencies for all new characters of this class.
	Proficiencies []APIReference `json:"proficiencies,omitempty"`
	// Saving throws the class is proficient in.
	SavingThrows []APIReference `json:"saving_throws,omitempty"`
	// List of all possible subclasses this class can specialize in.
	Subclasses []APIReference `json:"subclasses,omitempty"`
}

// NewClass instantiates a new Class object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClass() *Class {
	this := Class{}
	return &this
}

// NewClassWithDefaults instantiates a new Class object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClassWithDefaults() *Class {
	this := Class{}
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *Class) GetIndex() string {
	if o == nil || IsNil(o.Index) {
		var ret string
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Class) GetIndexOk() (*string, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *Class) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given string and assigns it to the Index field.
func (o *Class) SetIndex(v string) {
	o.Index = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Class) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Class) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Class) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Class) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Class) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Class) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Class) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Class) SetUrl(v string) {
	o.Url = &v
}

// GetHitDie returns the HitDie field value if set, zero value otherwise.
func (o *Class) GetHitDie() float32 {
	if o == nil || IsNil(o.HitDie) {
		var ret float32
		return ret
	}
	return *o.HitDie
}

// GetHitDieOk returns a tuple with the HitDie field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Class) GetHitDieOk() (*float32, bool) {
	if o == nil || IsNil(o.HitDie) {
		return nil, false
	}
	return o.HitDie, true
}

// HasHitDie returns a boolean if a field has been set.
func (o *Class) HasHitDie() bool {
	if o != nil && !IsNil(o.HitDie) {
		return true
	}

	return false
}

// SetHitDie gets a reference to the given float32 and assigns it to the HitDie field.
func (o *Class) SetHitDie(v float32) {
	o.HitDie = &v
}

// GetClassLevels returns the ClassLevels field value if set, zero value otherwise.
func (o *Class) GetClassLevels() string {
	if o == nil || IsNil(o.ClassLevels) {
		var ret string
		return ret
	}
	return *o.ClassLevels
}

// GetClassLevelsOk returns a tuple with the ClassLevels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Class) GetClassLevelsOk() (*string, bool) {
	if o == nil || IsNil(o.ClassLevels) {
		return nil, false
	}
	return o.ClassLevels, true
}

// HasClassLevels returns a boolean if a field has been set.
func (o *Class) HasClassLevels() bool {
	if o != nil && !IsNil(o.ClassLevels) {
		return true
	}

	return false
}

// SetClassLevels gets a reference to the given string and assigns it to the ClassLevels field.
func (o *Class) SetClassLevels(v string) {
	o.ClassLevels = &v
}

// GetMultiClassing returns the MultiClassing field value if set, zero value otherwise.
func (o *Class) GetMultiClassing() Multiclassing {
	if o == nil || IsNil(o.MultiClassing) {
		var ret Multiclassing
		return ret
	}
	return *o.MultiClassing
}

// GetMultiClassingOk returns a tuple with the MultiClassing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Class) GetMultiClassingOk() (*Multiclassing, bool) {
	if o == nil || IsNil(o.MultiClassing) {
		return nil, false
	}
	return o.MultiClassing, true
}

// HasMultiClassing returns a boolean if a field has been set.
func (o *Class) HasMultiClassing() bool {
	if o != nil && !IsNil(o.MultiClassing) {
		return true
	}

	return false
}

// SetMultiClassing gets a reference to the given Multiclassing and assigns it to the MultiClassing field.
func (o *Class) SetMultiClassing(v Multiclassing) {
	o.MultiClassing = &v
}

// GetSpellcasting returns the Spellcasting field value if set, zero value otherwise.
func (o *Class) GetSpellcasting() Spellcasting {
	if o == nil || IsNil(o.Spellcasting) {
		var ret Spellcasting
		return ret
	}
	return *o.Spellcasting
}

// GetSpellcastingOk returns a tuple with the Spellcasting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Class) GetSpellcastingOk() (*Spellcasting, bool) {
	if o == nil || IsNil(o.Spellcasting) {
		return nil, false
	}
	return o.Spellcasting, true
}

// HasSpellcasting returns a boolean if a field has been set.
func (o *Class) HasSpellcasting() bool {
	if o != nil && !IsNil(o.Spellcasting) {
		return true
	}

	return false
}

// SetSpellcasting gets a reference to the given Spellcasting and assigns it to the Spellcasting field.
func (o *Class) SetSpellcasting(v Spellcasting) {
	o.Spellcasting = &v
}

// GetSpells returns the Spells field value if set, zero value otherwise.
func (o *Class) GetSpells() string {
	if o == nil || IsNil(o.Spells) {
		var ret string
		return ret
	}
	return *o.Spells
}

// GetSpellsOk returns a tuple with the Spells field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Class) GetSpellsOk() (*string, bool) {
	if o == nil || IsNil(o.Spells) {
		return nil, false
	}
	return o.Spells, true
}

// HasSpells returns a boolean if a field has been set.
func (o *Class) HasSpells() bool {
	if o != nil && !IsNil(o.Spells) {
		return true
	}

	return false
}

// SetSpells gets a reference to the given string and assigns it to the Spells field.
func (o *Class) SetSpells(v string) {
	o.Spells = &v
}

// GetStartingEquipment returns the StartingEquipment field value if set, zero value otherwise.
func (o *Class) GetStartingEquipment() []ClassAllOfStartingEquipment {
	if o == nil || IsNil(o.StartingEquipment) {
		var ret []ClassAllOfStartingEquipment
		return ret
	}
	return o.StartingEquipment
}

// GetStartingEquipmentOk returns a tuple with the StartingEquipment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Class) GetStartingEquipmentOk() ([]ClassAllOfStartingEquipment, bool) {
	if o == nil || IsNil(o.StartingEquipment) {
		return nil, false
	}
	return o.StartingEquipment, true
}

// HasStartingEquipment returns a boolean if a field has been set.
func (o *Class) HasStartingEquipment() bool {
	if o != nil && !IsNil(o.StartingEquipment) {
		return true
	}

	return false
}

// SetStartingEquipment gets a reference to the given []ClassAllOfStartingEquipment and assigns it to the StartingEquipment field.
func (o *Class) SetStartingEquipment(v []ClassAllOfStartingEquipment) {
	o.StartingEquipment = v
}

// GetStartingEquipmentOptions returns the StartingEquipmentOptions field value if set, zero value otherwise.
func (o *Class) GetStartingEquipmentOptions() []Choice {
	if o == nil || IsNil(o.StartingEquipmentOptions) {
		var ret []Choice
		return ret
	}
	return o.StartingEquipmentOptions
}

// GetStartingEquipmentOptionsOk returns a tuple with the StartingEquipmentOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Class) GetStartingEquipmentOptionsOk() ([]Choice, bool) {
	if o == nil || IsNil(o.StartingEquipmentOptions) {
		return nil, false
	}
	return o.StartingEquipmentOptions, true
}

// HasStartingEquipmentOptions returns a boolean if a field has been set.
func (o *Class) HasStartingEquipmentOptions() bool {
	if o != nil && !IsNil(o.StartingEquipmentOptions) {
		return true
	}

	return false
}

// SetStartingEquipmentOptions gets a reference to the given []Choice and assigns it to the StartingEquipmentOptions field.
func (o *Class) SetStartingEquipmentOptions(v []Choice) {
	o.StartingEquipmentOptions = v
}

// GetProficiencyChoices returns the ProficiencyChoices field value if set, zero value otherwise.
func (o *Class) GetProficiencyChoices() []Choice {
	if o == nil || IsNil(o.ProficiencyChoices) {
		var ret []Choice
		return ret
	}
	return o.ProficiencyChoices
}

// GetProficiencyChoicesOk returns a tuple with the ProficiencyChoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Class) GetProficiencyChoicesOk() ([]Choice, bool) {
	if o == nil || IsNil(o.ProficiencyChoices) {
		return nil, false
	}
	return o.ProficiencyChoices, true
}

// HasProficiencyChoices returns a boolean if a field has been set.
func (o *Class) HasProficiencyChoices() bool {
	if o != nil && !IsNil(o.ProficiencyChoices) {
		return true
	}

	return false
}

// SetProficiencyChoices gets a reference to the given []Choice and assigns it to the ProficiencyChoices field.
func (o *Class) SetProficiencyChoices(v []Choice) {
	o.ProficiencyChoices = v
}

// GetProficiencies returns the Proficiencies field value if set, zero value otherwise.
func (o *Class) GetProficiencies() []APIReference {
	if o == nil || IsNil(o.Proficiencies) {
		var ret []APIReference
		return ret
	}
	return o.Proficiencies
}

// GetProficienciesOk returns a tuple with the Proficiencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Class) GetProficienciesOk() ([]APIReference, bool) {
	if o == nil || IsNil(o.Proficiencies) {
		return nil, false
	}
	return o.Proficiencies, true
}

// HasProficiencies returns a boolean if a field has been set.
func (o *Class) HasProficiencies() bool {
	if o != nil && !IsNil(o.Proficiencies) {
		return true
	}

	return false
}

// SetProficiencies gets a reference to the given []APIReference and assigns it to the Proficiencies field.
func (o *Class) SetProficiencies(v []APIReference) {
	o.Proficiencies = v
}

// GetSavingThrows returns the SavingThrows field value if set, zero value otherwise.
func (o *Class) GetSavingThrows() []APIReference {
	if o == nil || IsNil(o.SavingThrows) {
		var ret []APIReference
		return ret
	}
	return o.SavingThrows
}

// GetSavingThrowsOk returns a tuple with the SavingThrows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Class) GetSavingThrowsOk() ([]APIReference, bool) {
	if o == nil || IsNil(o.SavingThrows) {
		return nil, false
	}
	return o.SavingThrows, true
}

// HasSavingThrows returns a boolean if a field has been set.
func (o *Class) HasSavingThrows() bool {
	if o != nil && !IsNil(o.SavingThrows) {
		return true
	}

	return false
}

// SetSavingThrows gets a reference to the given []APIReference and assigns it to the SavingThrows field.
func (o *Class) SetSavingThrows(v []APIReference) {
	o.SavingThrows = v
}

// GetSubclasses returns the Subclasses field value if set, zero value otherwise.
func (o *Class) GetSubclasses() []APIReference {
	if o == nil || IsNil(o.Subclasses) {
		var ret []APIReference
		return ret
	}
	return o.Subclasses
}

// GetSubclassesOk returns a tuple with the Subclasses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Class) GetSubclassesOk() ([]APIReference, bool) {
	if o == nil || IsNil(o.Subclasses) {
		return nil, false
	}
	return o.Subclasses, true
}

// HasSubclasses returns a boolean if a field has been set.
func (o *Class) HasSubclasses() bool {
	if o != nil && !IsNil(o.Subclasses) {
		return true
	}

	return false
}

// SetSubclasses gets a reference to the given []APIReference and assigns it to the Subclasses field.
func (o *Class) SetSubclasses(v []APIReference) {
	o.Subclasses = v
}

func (o Class) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Class) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.HitDie) {
		toSerialize["hit_die"] = o.HitDie
	}
	if !IsNil(o.ClassLevels) {
		toSerialize["class_levels"] = o.ClassLevels
	}
	if !IsNil(o.MultiClassing) {
		toSerialize["multi_classing"] = o.MultiClassing
	}
	if !IsNil(o.Spellcasting) {
		toSerialize["spellcasting"] = o.Spellcasting
	}
	if !IsNil(o.Spells) {
		toSerialize["spells"] = o.Spells
	}
	if !IsNil(o.StartingEquipment) {
		toSerialize["starting_equipment"] = o.StartingEquipment
	}
	if !IsNil(o.StartingEquipmentOptions) {
		toSerialize["starting_equipment_options"] = o.StartingEquipmentOptions
	}
	if !IsNil(o.ProficiencyChoices) {
		toSerialize["proficiency_choices"] = o.ProficiencyChoices
	}
	if !IsNil(o.Proficiencies) {
		toSerialize["proficiencies"] = o.Proficiencies
	}
	if !IsNil(o.SavingThrows) {
		toSerialize["saving_throws"] = o.SavingThrows
	}
	if !IsNil(o.Subclasses) {
		toSerialize["subclasses"] = o.Subclasses
	}
	return toSerialize, nil
}

type NullableClass struct {
	value *Class
	isSet bool
}

func (v NullableClass) Get() *Class {
	return v.value
}

func (v *NullableClass) Set(val *Class) {
	v.value = val
	v.isSet = true
}

func (v NullableClass) IsSet() bool {
	return v.isSet
}

func (v *NullableClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClass(val *Class) *NullableClass {
	return &NullableClass{value: val, isSet: true}
}

func (v NullableClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


