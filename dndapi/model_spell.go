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

// checks if the Spell type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Spell{}

// Spell `Spell` 
type Spell struct {
	// Resource index for shorthand searching.
	Index *string `json:"index,omitempty"`
	// Name of the referenced resource.
	Name *string `json:"name,omitempty"`
	// URL of the referenced resource.
	Url *string `json:"url,omitempty"`
	// Description of the resource.
	Desc []string `json:"desc,omitempty"`
	// List of descriptions for casting the spell at higher levels.
	HigherLevel []string `json:"higher_level,omitempty"`
	// Range of the spell, usually expressed in feet.
	Range *string `json:"range,omitempty"`
	// List of shorthand for required components of the spell. V: verbal S: somatic M: material 
	Components []string `json:"components,omitempty"`
	// Material component for the spell to be cast.
	Material *string `json:"material,omitempty"`
	AreaOfEffect *AreaOfEffect `json:"area_of_effect,omitempty"`
	// Determines if a spell can be cast in a 10-min(in-game) ritual.
	Ritual *bool `json:"ritual,omitempty"`
	// How long the spell effect lasts.
	Duration *string `json:"duration,omitempty"`
	// Determines if a spell needs concentration to persist.
	Concentration *bool `json:"concentration,omitempty"`
	// How long it takes for the spell to activate.
	CastingTime *string `json:"casting_time,omitempty"`
	// Level of the spell.
	Level *float32 `json:"level,omitempty"`
	// Attack type of the spell.
	AttackType *string `json:"attack_type,omitempty"`
	Damage *SpellAllOfDamage `json:"damage,omitempty"`
	School *APIReference `json:"school,omitempty"`
	// List of classes that are able to learn the spell.
	Classes []APIReference `json:"classes,omitempty"`
	// List of subclasses that have access to the spell.
	Subclasses []APIReference `json:"subclasses,omitempty"`
}

// NewSpell instantiates a new Spell object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpell() *Spell {
	this := Spell{}
	return &this
}

// NewSpellWithDefaults instantiates a new Spell object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpellWithDefaults() *Spell {
	this := Spell{}
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *Spell) GetIndex() string {
	if o == nil || IsNil(o.Index) {
		var ret string
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spell) GetIndexOk() (*string, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *Spell) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given string and assigns it to the Index field.
func (o *Spell) SetIndex(v string) {
	o.Index = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Spell) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spell) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Spell) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Spell) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Spell) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spell) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Spell) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Spell) SetUrl(v string) {
	o.Url = &v
}

// GetDesc returns the Desc field value if set, zero value otherwise.
func (o *Spell) GetDesc() []string {
	if o == nil || IsNil(o.Desc) {
		var ret []string
		return ret
	}
	return o.Desc
}

// GetDescOk returns a tuple with the Desc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spell) GetDescOk() ([]string, bool) {
	if o == nil || IsNil(o.Desc) {
		return nil, false
	}
	return o.Desc, true
}

// HasDesc returns a boolean if a field has been set.
func (o *Spell) HasDesc() bool {
	if o != nil && !IsNil(o.Desc) {
		return true
	}

	return false
}

// SetDesc gets a reference to the given []string and assigns it to the Desc field.
func (o *Spell) SetDesc(v []string) {
	o.Desc = v
}

// GetHigherLevel returns the HigherLevel field value if set, zero value otherwise.
func (o *Spell) GetHigherLevel() []string {
	if o == nil || IsNil(o.HigherLevel) {
		var ret []string
		return ret
	}
	return o.HigherLevel
}

// GetHigherLevelOk returns a tuple with the HigherLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spell) GetHigherLevelOk() ([]string, bool) {
	if o == nil || IsNil(o.HigherLevel) {
		return nil, false
	}
	return o.HigherLevel, true
}

// HasHigherLevel returns a boolean if a field has been set.
func (o *Spell) HasHigherLevel() bool {
	if o != nil && !IsNil(o.HigherLevel) {
		return true
	}

	return false
}

// SetHigherLevel gets a reference to the given []string and assigns it to the HigherLevel field.
func (o *Spell) SetHigherLevel(v []string) {
	o.HigherLevel = v
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *Spell) GetRange() string {
	if o == nil || IsNil(o.Range) {
		var ret string
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spell) GetRangeOk() (*string, bool) {
	if o == nil || IsNil(o.Range) {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *Spell) HasRange() bool {
	if o != nil && !IsNil(o.Range) {
		return true
	}

	return false
}

// SetRange gets a reference to the given string and assigns it to the Range field.
func (o *Spell) SetRange(v string) {
	o.Range = &v
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *Spell) GetComponents() []string {
	if o == nil || IsNil(o.Components) {
		var ret []string
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spell) GetComponentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Components) {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *Spell) HasComponents() bool {
	if o != nil && !IsNil(o.Components) {
		return true
	}

	return false
}

// SetComponents gets a reference to the given []string and assigns it to the Components field.
func (o *Spell) SetComponents(v []string) {
	o.Components = v
}

// GetMaterial returns the Material field value if set, zero value otherwise.
func (o *Spell) GetMaterial() string {
	if o == nil || IsNil(o.Material) {
		var ret string
		return ret
	}
	return *o.Material
}

// GetMaterialOk returns a tuple with the Material field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spell) GetMaterialOk() (*string, bool) {
	if o == nil || IsNil(o.Material) {
		return nil, false
	}
	return o.Material, true
}

// HasMaterial returns a boolean if a field has been set.
func (o *Spell) HasMaterial() bool {
	if o != nil && !IsNil(o.Material) {
		return true
	}

	return false
}

// SetMaterial gets a reference to the given string and assigns it to the Material field.
func (o *Spell) SetMaterial(v string) {
	o.Material = &v
}

// GetAreaOfEffect returns the AreaOfEffect field value if set, zero value otherwise.
func (o *Spell) GetAreaOfEffect() AreaOfEffect {
	if o == nil || IsNil(o.AreaOfEffect) {
		var ret AreaOfEffect
		return ret
	}
	return *o.AreaOfEffect
}

// GetAreaOfEffectOk returns a tuple with the AreaOfEffect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spell) GetAreaOfEffectOk() (*AreaOfEffect, bool) {
	if o == nil || IsNil(o.AreaOfEffect) {
		return nil, false
	}
	return o.AreaOfEffect, true
}

// HasAreaOfEffect returns a boolean if a field has been set.
func (o *Spell) HasAreaOfEffect() bool {
	if o != nil && !IsNil(o.AreaOfEffect) {
		return true
	}

	return false
}

// SetAreaOfEffect gets a reference to the given AreaOfEffect and assigns it to the AreaOfEffect field.
func (o *Spell) SetAreaOfEffect(v AreaOfEffect) {
	o.AreaOfEffect = &v
}

// GetRitual returns the Ritual field value if set, zero value otherwise.
func (o *Spell) GetRitual() bool {
	if o == nil || IsNil(o.Ritual) {
		var ret bool
		return ret
	}
	return *o.Ritual
}

// GetRitualOk returns a tuple with the Ritual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spell) GetRitualOk() (*bool, bool) {
	if o == nil || IsNil(o.Ritual) {
		return nil, false
	}
	return o.Ritual, true
}

// HasRitual returns a boolean if a field has been set.
func (o *Spell) HasRitual() bool {
	if o != nil && !IsNil(o.Ritual) {
		return true
	}

	return false
}

// SetRitual gets a reference to the given bool and assigns it to the Ritual field.
func (o *Spell) SetRitual(v bool) {
	o.Ritual = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *Spell) GetDuration() string {
	if o == nil || IsNil(o.Duration) {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spell) GetDurationOk() (*string, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *Spell) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *Spell) SetDuration(v string) {
	o.Duration = &v
}

// GetConcentration returns the Concentration field value if set, zero value otherwise.
func (o *Spell) GetConcentration() bool {
	if o == nil || IsNil(o.Concentration) {
		var ret bool
		return ret
	}
	return *o.Concentration
}

// GetConcentrationOk returns a tuple with the Concentration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spell) GetConcentrationOk() (*bool, bool) {
	if o == nil || IsNil(o.Concentration) {
		return nil, false
	}
	return o.Concentration, true
}

// HasConcentration returns a boolean if a field has been set.
func (o *Spell) HasConcentration() bool {
	if o != nil && !IsNil(o.Concentration) {
		return true
	}

	return false
}

// SetConcentration gets a reference to the given bool and assigns it to the Concentration field.
func (o *Spell) SetConcentration(v bool) {
	o.Concentration = &v
}

// GetCastingTime returns the CastingTime field value if set, zero value otherwise.
func (o *Spell) GetCastingTime() string {
	if o == nil || IsNil(o.CastingTime) {
		var ret string
		return ret
	}
	return *o.CastingTime
}

// GetCastingTimeOk returns a tuple with the CastingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spell) GetCastingTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CastingTime) {
		return nil, false
	}
	return o.CastingTime, true
}

// HasCastingTime returns a boolean if a field has been set.
func (o *Spell) HasCastingTime() bool {
	if o != nil && !IsNil(o.CastingTime) {
		return true
	}

	return false
}

// SetCastingTime gets a reference to the given string and assigns it to the CastingTime field.
func (o *Spell) SetCastingTime(v string) {
	o.CastingTime = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *Spell) GetLevel() float32 {
	if o == nil || IsNil(o.Level) {
		var ret float32
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spell) GetLevelOk() (*float32, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *Spell) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given float32 and assigns it to the Level field.
func (o *Spell) SetLevel(v float32) {
	o.Level = &v
}

// GetAttackType returns the AttackType field value if set, zero value otherwise.
func (o *Spell) GetAttackType() string {
	if o == nil || IsNil(o.AttackType) {
		var ret string
		return ret
	}
	return *o.AttackType
}

// GetAttackTypeOk returns a tuple with the AttackType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spell) GetAttackTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AttackType) {
		return nil, false
	}
	return o.AttackType, true
}

// HasAttackType returns a boolean if a field has been set.
func (o *Spell) HasAttackType() bool {
	if o != nil && !IsNil(o.AttackType) {
		return true
	}

	return false
}

// SetAttackType gets a reference to the given string and assigns it to the AttackType field.
func (o *Spell) SetAttackType(v string) {
	o.AttackType = &v
}

// GetDamage returns the Damage field value if set, zero value otherwise.
func (o *Spell) GetDamage() SpellAllOfDamage {
	if o == nil || IsNil(o.Damage) {
		var ret SpellAllOfDamage
		return ret
	}
	return *o.Damage
}

// GetDamageOk returns a tuple with the Damage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spell) GetDamageOk() (*SpellAllOfDamage, bool) {
	if o == nil || IsNil(o.Damage) {
		return nil, false
	}
	return o.Damage, true
}

// HasDamage returns a boolean if a field has been set.
func (o *Spell) HasDamage() bool {
	if o != nil && !IsNil(o.Damage) {
		return true
	}

	return false
}

// SetDamage gets a reference to the given SpellAllOfDamage and assigns it to the Damage field.
func (o *Spell) SetDamage(v SpellAllOfDamage) {
	o.Damage = &v
}

// GetSchool returns the School field value if set, zero value otherwise.
func (o *Spell) GetSchool() APIReference {
	if o == nil || IsNil(o.School) {
		var ret APIReference
		return ret
	}
	return *o.School
}

// GetSchoolOk returns a tuple with the School field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spell) GetSchoolOk() (*APIReference, bool) {
	if o == nil || IsNil(o.School) {
		return nil, false
	}
	return o.School, true
}

// HasSchool returns a boolean if a field has been set.
func (o *Spell) HasSchool() bool {
	if o != nil && !IsNil(o.School) {
		return true
	}

	return false
}

// SetSchool gets a reference to the given APIReference and assigns it to the School field.
func (o *Spell) SetSchool(v APIReference) {
	o.School = &v
}

// GetClasses returns the Classes field value if set, zero value otherwise.
func (o *Spell) GetClasses() []APIReference {
	if o == nil || IsNil(o.Classes) {
		var ret []APIReference
		return ret
	}
	return o.Classes
}

// GetClassesOk returns a tuple with the Classes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spell) GetClassesOk() ([]APIReference, bool) {
	if o == nil || IsNil(o.Classes) {
		return nil, false
	}
	return o.Classes, true
}

// HasClasses returns a boolean if a field has been set.
func (o *Spell) HasClasses() bool {
	if o != nil && !IsNil(o.Classes) {
		return true
	}

	return false
}

// SetClasses gets a reference to the given []APIReference and assigns it to the Classes field.
func (o *Spell) SetClasses(v []APIReference) {
	o.Classes = v
}

// GetSubclasses returns the Subclasses field value if set, zero value otherwise.
func (o *Spell) GetSubclasses() []APIReference {
	if o == nil || IsNil(o.Subclasses) {
		var ret []APIReference
		return ret
	}
	return o.Subclasses
}

// GetSubclassesOk returns a tuple with the Subclasses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spell) GetSubclassesOk() ([]APIReference, bool) {
	if o == nil || IsNil(o.Subclasses) {
		return nil, false
	}
	return o.Subclasses, true
}

// HasSubclasses returns a boolean if a field has been set.
func (o *Spell) HasSubclasses() bool {
	if o != nil && !IsNil(o.Subclasses) {
		return true
	}

	return false
}

// SetSubclasses gets a reference to the given []APIReference and assigns it to the Subclasses field.
func (o *Spell) SetSubclasses(v []APIReference) {
	o.Subclasses = v
}

func (o Spell) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Spell) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.HigherLevel) {
		toSerialize["higher_level"] = o.HigherLevel
	}
	if !IsNil(o.Range) {
		toSerialize["range"] = o.Range
	}
	if !IsNil(o.Components) {
		toSerialize["components"] = o.Components
	}
	if !IsNil(o.Material) {
		toSerialize["material"] = o.Material
	}
	if !IsNil(o.AreaOfEffect) {
		toSerialize["area_of_effect"] = o.AreaOfEffect
	}
	if !IsNil(o.Ritual) {
		toSerialize["ritual"] = o.Ritual
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.Concentration) {
		toSerialize["concentration"] = o.Concentration
	}
	if !IsNil(o.CastingTime) {
		toSerialize["casting_time"] = o.CastingTime
	}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	if !IsNil(o.AttackType) {
		toSerialize["attack_type"] = o.AttackType
	}
	if !IsNil(o.Damage) {
		toSerialize["damage"] = o.Damage
	}
	if !IsNil(o.School) {
		toSerialize["school"] = o.School
	}
	if !IsNil(o.Classes) {
		toSerialize["classes"] = o.Classes
	}
	if !IsNil(o.Subclasses) {
		toSerialize["subclasses"] = o.Subclasses
	}
	return toSerialize, nil
}

type NullableSpell struct {
	value *Spell
	isSet bool
}

func (v NullableSpell) Get() *Spell {
	return v.value
}

func (v *NullableSpell) Set(val *Spell) {
	v.value = val
	v.isSet = true
}

func (v NullableSpell) IsSet() bool {
	return v.isSet
}

func (v *NullableSpell) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpell(val *Spell) *NullableSpell {
	return &NullableSpell{value: val, isSet: true}
}

func (v NullableSpell) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpell) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


