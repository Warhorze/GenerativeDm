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

// checks if the Weapon type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Weapon{}

// Weapon `Weapon` 
type Weapon struct {
	// Resource index for shorthand searching.
	Index *string `json:"index,omitempty"`
	// Name of the referenced resource.
	Name *string `json:"name,omitempty"`
	// URL of the referenced resource.
	Url *string `json:"url,omitempty"`
	// Description of the resource.
	Desc []string `json:"desc,omitempty"`
	EquipmentCategory *APIReference `json:"equipment_category,omitempty"`
	// The category of weapon this falls into.
	WeaponCategory *string `json:"weapon_category,omitempty"`
	// Whether this is a Melee or Ranged weapon.
	WeaponRange *string `json:"weapon_range,omitempty"`
	// A combination of weapon_category and weapon_range.
	CategoryRange *string `json:"category_range,omitempty"`
	Range *WeaponAllOfRange `json:"range,omitempty"`
	Damage *Damage `json:"damage,omitempty"`
	TwoHandedDamage *Damage `json:"two_handed_damage,omitempty"`
	// A list of the properties this weapon has.
	Properties []APIReference `json:"properties,omitempty"`
	Cost *Cost `json:"cost,omitempty"`
	// How much the equipment weighs.
	Weight *float32 `json:"weight,omitempty"`
}

// NewWeapon instantiates a new Weapon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeapon() *Weapon {
	this := Weapon{}
	return &this
}

// NewWeaponWithDefaults instantiates a new Weapon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeaponWithDefaults() *Weapon {
	this := Weapon{}
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *Weapon) GetIndex() string {
	if o == nil || IsNil(o.Index) {
		var ret string
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Weapon) GetIndexOk() (*string, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *Weapon) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given string and assigns it to the Index field.
func (o *Weapon) SetIndex(v string) {
	o.Index = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Weapon) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Weapon) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Weapon) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Weapon) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Weapon) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Weapon) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Weapon) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Weapon) SetUrl(v string) {
	o.Url = &v
}

// GetDesc returns the Desc field value if set, zero value otherwise.
func (o *Weapon) GetDesc() []string {
	if o == nil || IsNil(o.Desc) {
		var ret []string
		return ret
	}
	return o.Desc
}

// GetDescOk returns a tuple with the Desc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Weapon) GetDescOk() ([]string, bool) {
	if o == nil || IsNil(o.Desc) {
		return nil, false
	}
	return o.Desc, true
}

// HasDesc returns a boolean if a field has been set.
func (o *Weapon) HasDesc() bool {
	if o != nil && !IsNil(o.Desc) {
		return true
	}

	return false
}

// SetDesc gets a reference to the given []string and assigns it to the Desc field.
func (o *Weapon) SetDesc(v []string) {
	o.Desc = v
}

// GetEquipmentCategory returns the EquipmentCategory field value if set, zero value otherwise.
func (o *Weapon) GetEquipmentCategory() APIReference {
	if o == nil || IsNil(o.EquipmentCategory) {
		var ret APIReference
		return ret
	}
	return *o.EquipmentCategory
}

// GetEquipmentCategoryOk returns a tuple with the EquipmentCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Weapon) GetEquipmentCategoryOk() (*APIReference, bool) {
	if o == nil || IsNil(o.EquipmentCategory) {
		return nil, false
	}
	return o.EquipmentCategory, true
}

// HasEquipmentCategory returns a boolean if a field has been set.
func (o *Weapon) HasEquipmentCategory() bool {
	if o != nil && !IsNil(o.EquipmentCategory) {
		return true
	}

	return false
}

// SetEquipmentCategory gets a reference to the given APIReference and assigns it to the EquipmentCategory field.
func (o *Weapon) SetEquipmentCategory(v APIReference) {
	o.EquipmentCategory = &v
}

// GetWeaponCategory returns the WeaponCategory field value if set, zero value otherwise.
func (o *Weapon) GetWeaponCategory() string {
	if o == nil || IsNil(o.WeaponCategory) {
		var ret string
		return ret
	}
	return *o.WeaponCategory
}

// GetWeaponCategoryOk returns a tuple with the WeaponCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Weapon) GetWeaponCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.WeaponCategory) {
		return nil, false
	}
	return o.WeaponCategory, true
}

// HasWeaponCategory returns a boolean if a field has been set.
func (o *Weapon) HasWeaponCategory() bool {
	if o != nil && !IsNil(o.WeaponCategory) {
		return true
	}

	return false
}

// SetWeaponCategory gets a reference to the given string and assigns it to the WeaponCategory field.
func (o *Weapon) SetWeaponCategory(v string) {
	o.WeaponCategory = &v
}

// GetWeaponRange returns the WeaponRange field value if set, zero value otherwise.
func (o *Weapon) GetWeaponRange() string {
	if o == nil || IsNil(o.WeaponRange) {
		var ret string
		return ret
	}
	return *o.WeaponRange
}

// GetWeaponRangeOk returns a tuple with the WeaponRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Weapon) GetWeaponRangeOk() (*string, bool) {
	if o == nil || IsNil(o.WeaponRange) {
		return nil, false
	}
	return o.WeaponRange, true
}

// HasWeaponRange returns a boolean if a field has been set.
func (o *Weapon) HasWeaponRange() bool {
	if o != nil && !IsNil(o.WeaponRange) {
		return true
	}

	return false
}

// SetWeaponRange gets a reference to the given string and assigns it to the WeaponRange field.
func (o *Weapon) SetWeaponRange(v string) {
	o.WeaponRange = &v
}

// GetCategoryRange returns the CategoryRange field value if set, zero value otherwise.
func (o *Weapon) GetCategoryRange() string {
	if o == nil || IsNil(o.CategoryRange) {
		var ret string
		return ret
	}
	return *o.CategoryRange
}

// GetCategoryRangeOk returns a tuple with the CategoryRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Weapon) GetCategoryRangeOk() (*string, bool) {
	if o == nil || IsNil(o.CategoryRange) {
		return nil, false
	}
	return o.CategoryRange, true
}

// HasCategoryRange returns a boolean if a field has been set.
func (o *Weapon) HasCategoryRange() bool {
	if o != nil && !IsNil(o.CategoryRange) {
		return true
	}

	return false
}

// SetCategoryRange gets a reference to the given string and assigns it to the CategoryRange field.
func (o *Weapon) SetCategoryRange(v string) {
	o.CategoryRange = &v
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *Weapon) GetRange() WeaponAllOfRange {
	if o == nil || IsNil(o.Range) {
		var ret WeaponAllOfRange
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Weapon) GetRangeOk() (*WeaponAllOfRange, bool) {
	if o == nil || IsNil(o.Range) {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *Weapon) HasRange() bool {
	if o != nil && !IsNil(o.Range) {
		return true
	}

	return false
}

// SetRange gets a reference to the given WeaponAllOfRange and assigns it to the Range field.
func (o *Weapon) SetRange(v WeaponAllOfRange) {
	o.Range = &v
}

// GetDamage returns the Damage field value if set, zero value otherwise.
func (o *Weapon) GetDamage() Damage {
	if o == nil || IsNil(o.Damage) {
		var ret Damage
		return ret
	}
	return *o.Damage
}

// GetDamageOk returns a tuple with the Damage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Weapon) GetDamageOk() (*Damage, bool) {
	if o == nil || IsNil(o.Damage) {
		return nil, false
	}
	return o.Damage, true
}

// HasDamage returns a boolean if a field has been set.
func (o *Weapon) HasDamage() bool {
	if o != nil && !IsNil(o.Damage) {
		return true
	}

	return false
}

// SetDamage gets a reference to the given Damage and assigns it to the Damage field.
func (o *Weapon) SetDamage(v Damage) {
	o.Damage = &v
}

// GetTwoHandedDamage returns the TwoHandedDamage field value if set, zero value otherwise.
func (o *Weapon) GetTwoHandedDamage() Damage {
	if o == nil || IsNil(o.TwoHandedDamage) {
		var ret Damage
		return ret
	}
	return *o.TwoHandedDamage
}

// GetTwoHandedDamageOk returns a tuple with the TwoHandedDamage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Weapon) GetTwoHandedDamageOk() (*Damage, bool) {
	if o == nil || IsNil(o.TwoHandedDamage) {
		return nil, false
	}
	return o.TwoHandedDamage, true
}

// HasTwoHandedDamage returns a boolean if a field has been set.
func (o *Weapon) HasTwoHandedDamage() bool {
	if o != nil && !IsNil(o.TwoHandedDamage) {
		return true
	}

	return false
}

// SetTwoHandedDamage gets a reference to the given Damage and assigns it to the TwoHandedDamage field.
func (o *Weapon) SetTwoHandedDamage(v Damage) {
	o.TwoHandedDamage = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *Weapon) GetProperties() []APIReference {
	if o == nil || IsNil(o.Properties) {
		var ret []APIReference
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Weapon) GetPropertiesOk() ([]APIReference, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *Weapon) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []APIReference and assigns it to the Properties field.
func (o *Weapon) SetProperties(v []APIReference) {
	o.Properties = v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *Weapon) GetCost() Cost {
	if o == nil || IsNil(o.Cost) {
		var ret Cost
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Weapon) GetCostOk() (*Cost, bool) {
	if o == nil || IsNil(o.Cost) {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *Weapon) HasCost() bool {
	if o != nil && !IsNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given Cost and assigns it to the Cost field.
func (o *Weapon) SetCost(v Cost) {
	o.Cost = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *Weapon) GetWeight() float32 {
	if o == nil || IsNil(o.Weight) {
		var ret float32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Weapon) GetWeightOk() (*float32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *Weapon) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given float32 and assigns it to the Weight field.
func (o *Weapon) SetWeight(v float32) {
	o.Weight = &v
}

func (o Weapon) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Weapon) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.EquipmentCategory) {
		toSerialize["equipment_category"] = o.EquipmentCategory
	}
	if !IsNil(o.WeaponCategory) {
		toSerialize["weapon_category"] = o.WeaponCategory
	}
	if !IsNil(o.WeaponRange) {
		toSerialize["weapon_range"] = o.WeaponRange
	}
	if !IsNil(o.CategoryRange) {
		toSerialize["category_range"] = o.CategoryRange
	}
	if !IsNil(o.Range) {
		toSerialize["range"] = o.Range
	}
	if !IsNil(o.Damage) {
		toSerialize["damage"] = o.Damage
	}
	if !IsNil(o.TwoHandedDamage) {
		toSerialize["two_handed_damage"] = o.TwoHandedDamage
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.Cost) {
		toSerialize["cost"] = o.Cost
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	return toSerialize, nil
}

type NullableWeapon struct {
	value *Weapon
	isSet bool
}

func (v NullableWeapon) Get() *Weapon {
	return v.value
}

func (v *NullableWeapon) Set(val *Weapon) {
	v.value = val
	v.isSet = true
}

func (v NullableWeapon) IsSet() bool {
	return v.isSet
}

func (v *NullableWeapon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeapon(val *Weapon) *NullableWeapon {
	return &NullableWeapon{value: val, isSet: true}
}

func (v NullableWeapon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeapon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


