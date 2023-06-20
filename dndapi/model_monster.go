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

// checks if the Monster type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Monster{}

// Monster `Monster` 
type Monster struct {
	// Resource index for shorthand searching.
	Index *string `json:"index,omitempty"`
	// Name of the referenced resource.
	Name *string `json:"name,omitempty"`
	// URL of the referenced resource.
	Url *string `json:"url,omitempty"`
	// Description of the resource.
	Desc []string `json:"desc,omitempty"`
	// A monster's ability to charm or intimidate a player.
	Charisma *float32 `json:"charisma,omitempty"`
	// How sturdy a monster is.\"
	Constitution *float32 `json:"constitution,omitempty"`
	// The monster's ability for swift movement or stealth
	Dexterity *float32 `json:"dexterity,omitempty"`
	// The monster's ability to outsmart a player.
	Intelligence *float32 `json:"intelligence,omitempty"`
	// How hard a monster can hit a player.
	Strength *float32 `json:"strength,omitempty"`
	// A monster's ability to ascertain the player's plan.
	Wisdom *float32 `json:"wisdom,omitempty"`
	// The image url of the monster.
	Image *string `json:"image,omitempty"`
	// The size of the monster ranging from Tiny to Gargantuan.\"
	Size *string `json:"size,omitempty"`
	// The type of monster.
	Type *string `json:"type,omitempty"`
	// The sub-category of a monster used for classification of monsters.\"
	Subtype *string `json:"subtype,omitempty"`
	// A creature's general moral and personal attitudes.
	Alignments *string `json:"alignments,omitempty"`
	// The difficulty for a player to successfully deal damage to a monster.
	ArmorClass []MonsterArmorClass `json:"armor_class,omitempty"`
	// The hit points of a monster determine how much damage it is able to take before it can be defeated.
	HitPoints *float32 `json:"hit_points,omitempty"`
	// The hit die of a monster can be used to make a version of the same monster whose hit points are determined by the roll of the die. For example: A monster with 2d6 would have its hit points determine by rolling a 6 sided die twice.
	HitDice *string `json:"hit_dice,omitempty"`
	// The roll for determining a monster's hit points, which consists of the hit dice (e.g. 18d10) and the modifier determined by its Constitution (e.g. +36). For example, 18d10+36
	HitPointsRoll *string `json:"hit_points_roll,omitempty"`
	// A list of actions that are available to the monster to take during combat.
	Actions []MonsterAction `json:"actions,omitempty"`
	// A list of legendary actions that are available to the monster to take during combat.
	LegendaryActions []MonsterAction `json:"legendary_actions,omitempty"`
	// A monster's challenge rating is a guideline number that says when a monster becomes an appropriate challenge against the party's average level. For example. A group of 4 players with an average level of 4 would have an appropriate combat challenge against a monster with a challenge rating of 4 but a monster with a challenge rating of 8 against the same group of players would pose a significant threat.
	ChallengeRating *float32 `json:"challenge_rating,omitempty"`
	// A list of conditions that a monster is immune to.
	ConditionImmunities []APIReference `json:"condition_immunities,omitempty"`
	// A list of damage types that a monster will take double damage from.
	DamageImmunities []string `json:"damage_immunities,omitempty"`
	// A list of damage types that a monster will take half damage from.
	DamageResistances []string `json:"damage_resistances,omitempty"`
	// A list of damage types that a monster will take double damage from.
	DamageVulnerabilities []string `json:"damage_vulnerabilities,omitempty"`
	// List of other related monster entries that are of the same form. Only applicable to Lycanthropes that have multiple forms.
	Forms []APIReference `json:"forms,omitempty"`
	// The languages a monster is able to speak.
	Languages *string `json:"languages,omitempty"`
	// A list of proficiencies of a monster.
	Proficiencies []MonsterProficiency `json:"proficiencies,omitempty"`
	// A list of reactions that is available to the monster to take during combat.
	Reactions []MonsterAction `json:"reactions,omitempty"`
	Senses *MonsterAllOfSenses `json:"senses,omitempty"`
	// A list of the monster's special abilities.
	SpecialAbilities []MonsterSpecialAbility `json:"special_abilities,omitempty"`
	Speed *MonsterAllOfSpeed `json:"speed,omitempty"`
	// The number of experience points (XP) a monster is worth is based on its challenge rating.
	Xp *float32 `json:"xp,omitempty"`
}

// NewMonster instantiates a new Monster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonster() *Monster {
	this := Monster{}
	return &this
}

// NewMonsterWithDefaults instantiates a new Monster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonsterWithDefaults() *Monster {
	this := Monster{}
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *Monster) GetIndex() string {
	if o == nil || IsNil(o.Index) {
		var ret string
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monster) GetIndexOk() (*string, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *Monster) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given string and assigns it to the Index field.
func (o *Monster) SetIndex(v string) {
	o.Index = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Monster) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monster) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Monster) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Monster) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Monster) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monster) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Monster) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Monster) SetUrl(v string) {
	o.Url = &v
}

// GetDesc returns the Desc field value if set, zero value otherwise.
func (o *Monster) GetDesc() []string {
	if o == nil || IsNil(o.Desc) {
		var ret []string
		return ret
	}
	return o.Desc
}

// GetDescOk returns a tuple with the Desc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monster) GetDescOk() ([]string, bool) {
	if o == nil || IsNil(o.Desc) {
		return nil, false
	}
	return o.Desc, true
}

// HasDesc returns a boolean if a field has been set.
func (o *Monster) HasDesc() bool {
	if o != nil && !IsNil(o.Desc) {
		return true
	}

	return false
}

// SetDesc gets a reference to the given []string and assigns it to the Desc field.
func (o *Monster) SetDesc(v []string) {
	o.Desc = v
}

// GetCharisma returns the Charisma field value if set, zero value otherwise.
func (o *Monster) GetCharisma() float32 {
	if o == nil || IsNil(o.Charisma) {
		var ret float32
		return ret
	}
	return *o.Charisma
}

// GetCharismaOk returns a tuple with the Charisma field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monster) GetCharismaOk() (*float32, bool) {
	if o == nil || IsNil(o.Charisma) {
		return nil, false
	}
	return o.Charisma, true
}

// HasCharisma returns a boolean if a field has been set.
func (o *Monster) HasCharisma() bool {
	if o != nil && !IsNil(o.Charisma) {
		return true
	}

	return false
}

// SetCharisma gets a reference to the given float32 and assigns it to the Charisma field.
func (o *Monster) SetCharisma(v float32) {
	o.Charisma = &v
}

// GetConstitution returns the Constitution field value if set, zero value otherwise.
func (o *Monster) GetConstitution() float32 {
	if o == nil || IsNil(o.Constitution) {
		var ret float32
		return ret
	}
	return *o.Constitution
}

// GetConstitutionOk returns a tuple with the Constitution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monster) GetConstitutionOk() (*float32, bool) {
	if o == nil || IsNil(o.Constitution) {
		return nil, false
	}
	return o.Constitution, true
}

// HasConstitution returns a boolean if a field has been set.
func (o *Monster) HasConstitution() bool {
	if o != nil && !IsNil(o.Constitution) {
		return true
	}

	return false
}

// SetConstitution gets a reference to the given float32 and assigns it to the Constitution field.
func (o *Monster) SetConstitution(v float32) {
	o.Constitution = &v
}

// GetDexterity returns the Dexterity field value if set, zero value otherwise.
func (o *Monster) GetDexterity() float32 {
	if o == nil || IsNil(o.Dexterity) {
		var ret float32
		return ret
	}
	return *o.Dexterity
}

// GetDexterityOk returns a tuple with the Dexterity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monster) GetDexterityOk() (*float32, bool) {
	if o == nil || IsNil(o.Dexterity) {
		return nil, false
	}
	return o.Dexterity, true
}

// HasDexterity returns a boolean if a field has been set.
func (o *Monster) HasDexterity() bool {
	if o != nil && !IsNil(o.Dexterity) {
		return true
	}

	return false
}

// SetDexterity gets a reference to the given float32 and assigns it to the Dexterity field.
func (o *Monster) SetDexterity(v float32) {
	o.Dexterity = &v
}

// GetIntelligence returns the Intelligence field value if set, zero value otherwise.
func (o *Monster) GetIntelligence() float32 {
	if o == nil || IsNil(o.Intelligence) {
		var ret float32
		return ret
	}
	return *o.Intelligence
}

// GetIntelligenceOk returns a tuple with the Intelligence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monster) GetIntelligenceOk() (*float32, bool) {
	if o == nil || IsNil(o.Intelligence) {
		return nil, false
	}
	return o.Intelligence, true
}

// HasIntelligence returns a boolean if a field has been set.
func (o *Monster) HasIntelligence() bool {
	if o != nil && !IsNil(o.Intelligence) {
		return true
	}

	return false
}

// SetIntelligence gets a reference to the given float32 and assigns it to the Intelligence field.
func (o *Monster) SetIntelligence(v float32) {
	o.Intelligence = &v
}

// GetStrength returns the Strength field value if set, zero value otherwise.
func (o *Monster) GetStrength() float32 {
	if o == nil || IsNil(o.Strength) {
		var ret float32
		return ret
	}
	return *o.Strength
}

// GetStrengthOk returns a tuple with the Strength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monster) GetStrengthOk() (*float32, bool) {
	if o == nil || IsNil(o.Strength) {
		return nil, false
	}
	return o.Strength, true
}

// HasStrength returns a boolean if a field has been set.
func (o *Monster) HasStrength() bool {
	if o != nil && !IsNil(o.Strength) {
		return true
	}

	return false
}

// SetStrength gets a reference to the given float32 and assigns it to the Strength field.
func (o *Monster) SetStrength(v float32) {
	o.Strength = &v
}

// GetWisdom returns the Wisdom field value if set, zero value otherwise.
func (o *Monster) GetWisdom() float32 {
	if o == nil || IsNil(o.Wisdom) {
		var ret float32
		return ret
	}
	return *o.Wisdom
}

// GetWisdomOk returns a tuple with the Wisdom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monster) GetWisdomOk() (*float32, bool) {
	if o == nil || IsNil(o.Wisdom) {
		return nil, false
	}
	return o.Wisdom, true
}

// HasWisdom returns a boolean if a field has been set.
func (o *Monster) HasWisdom() bool {
	if o != nil && !IsNil(o.Wisdom) {
		return true
	}

	return false
}

// SetWisdom gets a reference to the given float32 and assigns it to the Wisdom field.
func (o *Monster) SetWisdom(v float32) {
	o.Wisdom = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *Monster) GetImage() string {
	if o == nil || IsNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monster) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *Monster) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *Monster) SetImage(v string) {
	o.Image = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *Monster) GetSize() string {
	if o == nil || IsNil(o.Size) {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monster) GetSizeOk() (*string, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *Monster) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *Monster) SetSize(v string) {
	o.Size = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Monster) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monster) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Monster) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Monster) SetType(v string) {
	o.Type = &v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *Monster) GetSubtype() string {
	if o == nil || IsNil(o.Subtype) {
		var ret string
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monster) GetSubtypeOk() (*string, bool) {
	if o == nil || IsNil(o.Subtype) {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *Monster) HasSubtype() bool {
	if o != nil && !IsNil(o.Subtype) {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given string and assigns it to the Subtype field.
func (o *Monster) SetSubtype(v string) {
	o.Subtype = &v
}

// GetAlignments returns the Alignments field value if set, zero value otherwise.
func (o *Monster) GetAlignments() string {
	if o == nil || IsNil(o.Alignments) {
		var ret string
		return ret
	}
	return *o.Alignments
}

// GetAlignmentsOk returns a tuple with the Alignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monster) GetAlignmentsOk() (*string, bool) {
	if o == nil || IsNil(o.Alignments) {
		return nil, false
	}
	return o.Alignments, true
}

// HasAlignments returns a boolean if a field has been set.
func (o *Monster) HasAlignments() bool {
	if o != nil && !IsNil(o.Alignments) {
		return true
	}

	return false
}

// SetAlignments gets a reference to the given string and assigns it to the Alignments field.
func (o *Monster) SetAlignments(v string) {
	o.Alignments = &v
}

// GetArmorClass returns the ArmorClass field value if set, zero value otherwise.
func (o *Monster) GetArmorClass() []MonsterArmorClass {
	if o == nil || IsNil(o.ArmorClass) {
		var ret []MonsterArmorClass
		return ret
	}
	return o.ArmorClass
}

// GetArmorClassOk returns a tuple with the ArmorClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monster) GetArmorClassOk() ([]MonsterArmorClass, bool) {
	if o == nil || IsNil(o.ArmorClass) {
		return nil, false
	}
	return o.ArmorClass, true
}

// HasArmorClass returns a boolean if a field has been set.
func (o *Monster) HasArmorClass() bool {
	if o != nil && !IsNil(o.ArmorClass) {
		return true
	}

	return false
}

// SetArmorClass gets a reference to the given []MonsterArmorClass and assigns it to the ArmorClass field.
func (o *Monster) SetArmorClass(v []MonsterArmorClass) {
	o.ArmorClass = v
}

// GetHitPoints returns the HitPoints field value if set, zero value otherwise.
func (o *Monster) GetHitPoints() float32 {
	if o == nil || IsNil(o.HitPoints) {
		var ret float32
		return ret
	}
	return *o.HitPoints
}

// GetHitPointsOk returns a tuple with the HitPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monster) GetHitPointsOk() (*float32, bool) {
	if o == nil || IsNil(o.HitPoints) {
		return nil, false
	}
	return o.HitPoints, true
}

// HasHitPoints returns a boolean if a field has been set.
func (o *Monster) HasHitPoints() bool {
	if o != nil && !IsNil(o.HitPoints) {
		return true
	}

	return false
}

// SetHitPoints gets a reference to the given float32 and assigns it to the HitPoints field.
func (o *Monster) SetHitPoints(v float32) {
	o.HitPoints = &v
}

// GetHitDice returns the HitDice field value if set, zero value otherwise.
func (o *Monster) GetHitDice() string {
	if o == nil || IsNil(o.HitDice) {
		var ret string
		return ret
	}
	return *o.HitDice
}

// GetHitDiceOk returns a tuple with the HitDice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monster) GetHitDiceOk() (*string, bool) {
	if o == nil || IsNil(o.HitDice) {
		return nil, false
	}
	return o.HitDice, true
}

// HasHitDice returns a boolean if a field has been set.
func (o *Monster) HasHitDice() bool {
	if o != nil && !IsNil(o.HitDice) {
		return true
	}

	return false
}

// SetHitDice gets a reference to the given string and assigns it to the HitDice field.
func (o *Monster) SetHitDice(v string) {
	o.HitDice = &v
}

// GetHitPointsRoll returns the HitPointsRoll field value if set, zero value otherwise.
func (o *Monster) GetHitPointsRoll() string {
	if o == nil || IsNil(o.HitPointsRoll) {
		var ret string
		return ret
	}
	return *o.HitPointsRoll
}

// GetHitPointsRollOk returns a tuple with the HitPointsRoll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monster) GetHitPointsRollOk() (*string, bool) {
	if o == nil || IsNil(o.HitPointsRoll) {
		return nil, false
	}
	return o.HitPointsRoll, true
}

// HasHitPointsRoll returns a boolean if a field has been set.
func (o *Monster) HasHitPointsRoll() bool {
	if o != nil && !IsNil(o.HitPointsRoll) {
		return true
	}

	return false
}

// SetHitPointsRoll gets a reference to the given string and assigns it to the HitPointsRoll field.
func (o *Monster) SetHitPointsRoll(v string) {
	o.HitPointsRoll = &v
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *Monster) GetActions() []MonsterAction {
	if o == nil || IsNil(o.Actions) {
		var ret []MonsterAction
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monster) GetActionsOk() ([]MonsterAction, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *Monster) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []MonsterAction and assigns it to the Actions field.
func (o *Monster) SetActions(v []MonsterAction) {
	o.Actions = v
}

// GetLegendaryActions returns the LegendaryActions field value if set, zero value otherwise.
func (o *Monster) GetLegendaryActions() []MonsterAction {
	if o == nil || IsNil(o.LegendaryActions) {
		var ret []MonsterAction
		return ret
	}
	return o.LegendaryActions
}

// GetLegendaryActionsOk returns a tuple with the LegendaryActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monster) GetLegendaryActionsOk() ([]MonsterAction, bool) {
	if o == nil || IsNil(o.LegendaryActions) {
		return nil, false
	}
	return o.LegendaryActions, true
}

// HasLegendaryActions returns a boolean if a field has been set.
func (o *Monster) HasLegendaryActions() bool {
	if o != nil && !IsNil(o.LegendaryActions) {
		return true
	}

	return false
}

// SetLegendaryActions gets a reference to the given []MonsterAction and assigns it to the LegendaryActions field.
func (o *Monster) SetLegendaryActions(v []MonsterAction) {
	o.LegendaryActions = v
}

// GetChallengeRating returns the ChallengeRating field value if set, zero value otherwise.
func (o *Monster) GetChallengeRating() float32 {
	if o == nil || IsNil(o.ChallengeRating) {
		var ret float32
		return ret
	}
	return *o.ChallengeRating
}

// GetChallengeRatingOk returns a tuple with the ChallengeRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monster) GetChallengeRatingOk() (*float32, bool) {
	if o == nil || IsNil(o.ChallengeRating) {
		return nil, false
	}
	return o.ChallengeRating, true
}

// HasChallengeRating returns a boolean if a field has been set.
func (o *Monster) HasChallengeRating() bool {
	if o != nil && !IsNil(o.ChallengeRating) {
		return true
	}

	return false
}

// SetChallengeRating gets a reference to the given float32 and assigns it to the ChallengeRating field.
func (o *Monster) SetChallengeRating(v float32) {
	o.ChallengeRating = &v
}

// GetConditionImmunities returns the ConditionImmunities field value if set, zero value otherwise.
func (o *Monster) GetConditionImmunities() []APIReference {
	if o == nil || IsNil(o.ConditionImmunities) {
		var ret []APIReference
		return ret
	}
	return o.ConditionImmunities
}

// GetConditionImmunitiesOk returns a tuple with the ConditionImmunities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monster) GetConditionImmunitiesOk() ([]APIReference, bool) {
	if o == nil || IsNil(o.ConditionImmunities) {
		return nil, false
	}
	return o.ConditionImmunities, true
}

// HasConditionImmunities returns a boolean if a field has been set.
func (o *Monster) HasConditionImmunities() bool {
	if o != nil && !IsNil(o.ConditionImmunities) {
		return true
	}

	return false
}

// SetConditionImmunities gets a reference to the given []APIReference and assigns it to the ConditionImmunities field.
func (o *Monster) SetConditionImmunities(v []APIReference) {
	o.ConditionImmunities = v
}

// GetDamageImmunities returns the DamageImmunities field value if set, zero value otherwise.
func (o *Monster) GetDamageImmunities() []string {
	if o == nil || IsNil(o.DamageImmunities) {
		var ret []string
		return ret
	}
	return o.DamageImmunities
}

// GetDamageImmunitiesOk returns a tuple with the DamageImmunities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monster) GetDamageImmunitiesOk() ([]string, bool) {
	if o == nil || IsNil(o.DamageImmunities) {
		return nil, false
	}
	return o.DamageImmunities, true
}

// HasDamageImmunities returns a boolean if a field has been set.
func (o *Monster) HasDamageImmunities() bool {
	if o != nil && !IsNil(o.DamageImmunities) {
		return true
	}

	return false
}

// SetDamageImmunities gets a reference to the given []string and assigns it to the DamageImmunities field.
func (o *Monster) SetDamageImmunities(v []string) {
	o.DamageImmunities = v
}

// GetDamageResistances returns the DamageResistances field value if set, zero value otherwise.
func (o *Monster) GetDamageResistances() []string {
	if o == nil || IsNil(o.DamageResistances) {
		var ret []string
		return ret
	}
	return o.DamageResistances
}

// GetDamageResistancesOk returns a tuple with the DamageResistances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monster) GetDamageResistancesOk() ([]string, bool) {
	if o == nil || IsNil(o.DamageResistances) {
		return nil, false
	}
	return o.DamageResistances, true
}

// HasDamageResistances returns a boolean if a field has been set.
func (o *Monster) HasDamageResistances() bool {
	if o != nil && !IsNil(o.DamageResistances) {
		return true
	}

	return false
}

// SetDamageResistances gets a reference to the given []string and assigns it to the DamageResistances field.
func (o *Monster) SetDamageResistances(v []string) {
	o.DamageResistances = v
}

// GetDamageVulnerabilities returns the DamageVulnerabilities field value if set, zero value otherwise.
func (o *Monster) GetDamageVulnerabilities() []string {
	if o == nil || IsNil(o.DamageVulnerabilities) {
		var ret []string
		return ret
	}
	return o.DamageVulnerabilities
}

// GetDamageVulnerabilitiesOk returns a tuple with the DamageVulnerabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monster) GetDamageVulnerabilitiesOk() ([]string, bool) {
	if o == nil || IsNil(o.DamageVulnerabilities) {
		return nil, false
	}
	return o.DamageVulnerabilities, true
}

// HasDamageVulnerabilities returns a boolean if a field has been set.
func (o *Monster) HasDamageVulnerabilities() bool {
	if o != nil && !IsNil(o.DamageVulnerabilities) {
		return true
	}

	return false
}

// SetDamageVulnerabilities gets a reference to the given []string and assigns it to the DamageVulnerabilities field.
func (o *Monster) SetDamageVulnerabilities(v []string) {
	o.DamageVulnerabilities = v
}

// GetForms returns the Forms field value if set, zero value otherwise.
func (o *Monster) GetForms() []APIReference {
	if o == nil || IsNil(o.Forms) {
		var ret []APIReference
		return ret
	}
	return o.Forms
}

// GetFormsOk returns a tuple with the Forms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monster) GetFormsOk() ([]APIReference, bool) {
	if o == nil || IsNil(o.Forms) {
		return nil, false
	}
	return o.Forms, true
}

// HasForms returns a boolean if a field has been set.
func (o *Monster) HasForms() bool {
	if o != nil && !IsNil(o.Forms) {
		return true
	}

	return false
}

// SetForms gets a reference to the given []APIReference and assigns it to the Forms field.
func (o *Monster) SetForms(v []APIReference) {
	o.Forms = v
}

// GetLanguages returns the Languages field value if set, zero value otherwise.
func (o *Monster) GetLanguages() string {
	if o == nil || IsNil(o.Languages) {
		var ret string
		return ret
	}
	return *o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monster) GetLanguagesOk() (*string, bool) {
	if o == nil || IsNil(o.Languages) {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *Monster) HasLanguages() bool {
	if o != nil && !IsNil(o.Languages) {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given string and assigns it to the Languages field.
func (o *Monster) SetLanguages(v string) {
	o.Languages = &v
}

// GetProficiencies returns the Proficiencies field value if set, zero value otherwise.
func (o *Monster) GetProficiencies() []MonsterProficiency {
	if o == nil || IsNil(o.Proficiencies) {
		var ret []MonsterProficiency
		return ret
	}
	return o.Proficiencies
}

// GetProficienciesOk returns a tuple with the Proficiencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monster) GetProficienciesOk() ([]MonsterProficiency, bool) {
	if o == nil || IsNil(o.Proficiencies) {
		return nil, false
	}
	return o.Proficiencies, true
}

// HasProficiencies returns a boolean if a field has been set.
func (o *Monster) HasProficiencies() bool {
	if o != nil && !IsNil(o.Proficiencies) {
		return true
	}

	return false
}

// SetProficiencies gets a reference to the given []MonsterProficiency and assigns it to the Proficiencies field.
func (o *Monster) SetProficiencies(v []MonsterProficiency) {
	o.Proficiencies = v
}

// GetReactions returns the Reactions field value if set, zero value otherwise.
func (o *Monster) GetReactions() []MonsterAction {
	if o == nil || IsNil(o.Reactions) {
		var ret []MonsterAction
		return ret
	}
	return o.Reactions
}

// GetReactionsOk returns a tuple with the Reactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monster) GetReactionsOk() ([]MonsterAction, bool) {
	if o == nil || IsNil(o.Reactions) {
		return nil, false
	}
	return o.Reactions, true
}

// HasReactions returns a boolean if a field has been set.
func (o *Monster) HasReactions() bool {
	if o != nil && !IsNil(o.Reactions) {
		return true
	}

	return false
}

// SetReactions gets a reference to the given []MonsterAction and assigns it to the Reactions field.
func (o *Monster) SetReactions(v []MonsterAction) {
	o.Reactions = v
}

// GetSenses returns the Senses field value if set, zero value otherwise.
func (o *Monster) GetSenses() MonsterAllOfSenses {
	if o == nil || IsNil(o.Senses) {
		var ret MonsterAllOfSenses
		return ret
	}
	return *o.Senses
}

// GetSensesOk returns a tuple with the Senses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monster) GetSensesOk() (*MonsterAllOfSenses, bool) {
	if o == nil || IsNil(o.Senses) {
		return nil, false
	}
	return o.Senses, true
}

// HasSenses returns a boolean if a field has been set.
func (o *Monster) HasSenses() bool {
	if o != nil && !IsNil(o.Senses) {
		return true
	}

	return false
}

// SetSenses gets a reference to the given MonsterAllOfSenses and assigns it to the Senses field.
func (o *Monster) SetSenses(v MonsterAllOfSenses) {
	o.Senses = &v
}

// GetSpecialAbilities returns the SpecialAbilities field value if set, zero value otherwise.
func (o *Monster) GetSpecialAbilities() []MonsterSpecialAbility {
	if o == nil || IsNil(o.SpecialAbilities) {
		var ret []MonsterSpecialAbility
		return ret
	}
	return o.SpecialAbilities
}

// GetSpecialAbilitiesOk returns a tuple with the SpecialAbilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monster) GetSpecialAbilitiesOk() ([]MonsterSpecialAbility, bool) {
	if o == nil || IsNil(o.SpecialAbilities) {
		return nil, false
	}
	return o.SpecialAbilities, true
}

// HasSpecialAbilities returns a boolean if a field has been set.
func (o *Monster) HasSpecialAbilities() bool {
	if o != nil && !IsNil(o.SpecialAbilities) {
		return true
	}

	return false
}

// SetSpecialAbilities gets a reference to the given []MonsterSpecialAbility and assigns it to the SpecialAbilities field.
func (o *Monster) SetSpecialAbilities(v []MonsterSpecialAbility) {
	o.SpecialAbilities = v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *Monster) GetSpeed() MonsterAllOfSpeed {
	if o == nil || IsNil(o.Speed) {
		var ret MonsterAllOfSpeed
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monster) GetSpeedOk() (*MonsterAllOfSpeed, bool) {
	if o == nil || IsNil(o.Speed) {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *Monster) HasSpeed() bool {
	if o != nil && !IsNil(o.Speed) {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given MonsterAllOfSpeed and assigns it to the Speed field.
func (o *Monster) SetSpeed(v MonsterAllOfSpeed) {
	o.Speed = &v
}

// GetXp returns the Xp field value if set, zero value otherwise.
func (o *Monster) GetXp() float32 {
	if o == nil || IsNil(o.Xp) {
		var ret float32
		return ret
	}
	return *o.Xp
}

// GetXpOk returns a tuple with the Xp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monster) GetXpOk() (*float32, bool) {
	if o == nil || IsNil(o.Xp) {
		return nil, false
	}
	return o.Xp, true
}

// HasXp returns a boolean if a field has been set.
func (o *Monster) HasXp() bool {
	if o != nil && !IsNil(o.Xp) {
		return true
	}

	return false
}

// SetXp gets a reference to the given float32 and assigns it to the Xp field.
func (o *Monster) SetXp(v float32) {
	o.Xp = &v
}

func (o Monster) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Monster) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Charisma) {
		toSerialize["charisma"] = o.Charisma
	}
	if !IsNil(o.Constitution) {
		toSerialize["constitution"] = o.Constitution
	}
	if !IsNil(o.Dexterity) {
		toSerialize["dexterity"] = o.Dexterity
	}
	if !IsNil(o.Intelligence) {
		toSerialize["intelligence"] = o.Intelligence
	}
	if !IsNil(o.Strength) {
		toSerialize["strength"] = o.Strength
	}
	if !IsNil(o.Wisdom) {
		toSerialize["wisdom"] = o.Wisdom
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Subtype) {
		toSerialize["subtype"] = o.Subtype
	}
	if !IsNil(o.Alignments) {
		toSerialize["alignments"] = o.Alignments
	}
	if !IsNil(o.ArmorClass) {
		toSerialize["armor_class"] = o.ArmorClass
	}
	if !IsNil(o.HitPoints) {
		toSerialize["hit_points"] = o.HitPoints
	}
	if !IsNil(o.HitDice) {
		toSerialize["hit_dice"] = o.HitDice
	}
	if !IsNil(o.HitPointsRoll) {
		toSerialize["hit_points_roll"] = o.HitPointsRoll
	}
	if !IsNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}
	if !IsNil(o.LegendaryActions) {
		toSerialize["legendary_actions"] = o.LegendaryActions
	}
	if !IsNil(o.ChallengeRating) {
		toSerialize["challenge_rating"] = o.ChallengeRating
	}
	if !IsNil(o.ConditionImmunities) {
		toSerialize["condition_immunities"] = o.ConditionImmunities
	}
	if !IsNil(o.DamageImmunities) {
		toSerialize["damage_immunities"] = o.DamageImmunities
	}
	if !IsNil(o.DamageResistances) {
		toSerialize["damage_resistances"] = o.DamageResistances
	}
	if !IsNil(o.DamageVulnerabilities) {
		toSerialize["damage_vulnerabilities"] = o.DamageVulnerabilities
	}
	if !IsNil(o.Forms) {
		toSerialize["forms"] = o.Forms
	}
	if !IsNil(o.Languages) {
		toSerialize["languages"] = o.Languages
	}
	if !IsNil(o.Proficiencies) {
		toSerialize["proficiencies"] = o.Proficiencies
	}
	if !IsNil(o.Reactions) {
		toSerialize["reactions"] = o.Reactions
	}
	if !IsNil(o.Senses) {
		toSerialize["senses"] = o.Senses
	}
	if !IsNil(o.SpecialAbilities) {
		toSerialize["special_abilities"] = o.SpecialAbilities
	}
	if !IsNil(o.Speed) {
		toSerialize["speed"] = o.Speed
	}
	if !IsNil(o.Xp) {
		toSerialize["xp"] = o.Xp
	}
	return toSerialize, nil
}

type NullableMonster struct {
	value *Monster
	isSet bool
}

func (v NullableMonster) Get() *Monster {
	return v.value
}

func (v *NullableMonster) Set(val *Monster) {
	v.value = val
	v.isSet = true
}

func (v NullableMonster) IsSet() bool {
	return v.isSet
}

func (v *NullableMonster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonster(val *Monster) *NullableMonster {
	return &NullableMonster{value: val, isSet: true}
}

func (v NullableMonster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


