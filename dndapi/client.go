/*
D&D 5e API

# Introduction  Welcome to the dnd5eapi, the Dungeons & Dragons 5th Edition API! This documentation should help you familiarize yourself with the resources available and how to consume them with HTTP requests. Read through the getting started section before you dive in. Most of your problems should be solved just by reading through it.  ## Getting Started  Let's make our first API request to the D&D 5th Edition API!  Open up a terminal and use [curl](http://curl.haxx.se/) or [httpie](http://httpie.org/) to make an API request for a resource. You can also scroll through the definitions below and send requests directly from the endpoint documentation!  For example, if you paste and run this `curl` command: ```bash curl -X GET \"https://www.dnd5eapi.co/api/ability-scores/cha\" -H \"Accept: application/json\" ```  We should see a result containing details about the Charisma ability score: ```bash {   \"index\": \"cha\",   \"name\": \"CHA\",   \"full_name\": \"Charisma\",   \"desc\": [     \"Charisma measures your ability to interact effectively with others. It       includes such factors as confidence and eloquence, and it can represent       a charming or commanding personality.\",     \"A Charisma check might arise when you try to influence or entertain       others, when you try to make an impression or tell a convincing lie,       or when you are navigating a tricky social situation. The Deception,       Intimidation, Performance, and Persuasion skills reflect aptitude in       certain kinds of Charisma checks.\"   ],   \"skills\": [     {       \"name\": \"Deception\",       \"index\": \"deception\",       \"url\": \"/api/skills/deception\"     },     {       \"name\": \"Intimidation\",       \"index\": \"intimidation\",       \"url\": \"/api/skills/intimidation\"     },     {       \"name\": \"Performance\",       \"index\": \"performance\",       \"url\": \"/api/skills/performance\"     },     {       \"name\": \"Persuasion\",       \"index\": \"persuasion\",       \"url\": \"/api/skills/persuasion\"     }   ],   \"url\": \"/api/ability-scores/cha\" } ```  ## Authentication  The dnd5eapi is a completely open API. No authentication is required to query and get data. This also means that we've limited what you can do to just `GET`-ing the data. If you find a mistake in the data, feel free to [message us](https://discord.gg/TQuYTv7).  ## GraphQL  This API supports [GraphQL](https://graphql.org/). The GraphQL URL for this API is `https://www.dnd5eapi.co/graphql`. Most of your questions regarding the GraphQL schema can be answered by querying the endpoint with the Apollo sandbox explorer.  ## Schemas  Definitions of all schemas will be accessible in a future update. Two of the most common schemas are described here.  ### `APIReference` Represents a minimal representation of a resource. The detailed representation of the referenced resource can be retrieved by making a request to the referenced `URL`. ``` APIReference {   index     string   name      string   url       string } ``` <hr>  ### `DC` Represents a difficulty check. ``` DC {   dc_type       APIReference   dc_value      number   success_type  \"none\" | \"half\" | \"other\" } ``` <hr>  ### `Damage` Represents damage. ``` Damage {   damage_type     APIReference   damage_dice     string } ``` <hr>  ### `Choice` Represents a choice made by a player. Commonly seen related to decisions made during character creation or combat (e.g.: the description of the cleric class, under **Proficiencies**, states \"Skills: Choose two from History, Insight, Medicine, Persuasion, and Religion\" [[SRD p15]](https://media.wizards.com/2016/downloads/DND/SRD-OGL_V5.1.pdf#page=15)) ``` Choice {   desc      string   choose    number   type      string   from      OptionSet } ``` <hr>  ### `OptionSet` The OptionSet structure provides the options to be chosen from, or sufficient data to fetch and interpret the options. All OptionSets have an `option_set_type` attribute that indicates the structure of the object that contains the options. The possible values are `options_array`, `equipment_category`, and `reference_list`. Other attributes on the OptionSet depend on the value of this attribute. - `options_array`   - `options` (array): An array of Option objects. Each item in the array represents an option that can be chosen. - `equipment_category`   - `equipment_category` (APIReference): A reference to an EquipmentCategory. Each item in the EquipmentCategory's `equipment` array represents one option that can be chosen. - `resource_list`   - `resource_list_url` (string): A reference (by URL) to a collection in the database. The URL may include query parameters. Each item in the resulting ResourceList's `results` array represents one option that can be chosen. <hr>  ### `Option` When the options are given in an `options_array`, each item in the array inherits from the Option structure. All Options have an `option_type` attribute that indicates the structure of the option. The value of this attribute indicates how the option should be handled, and each type has different attributes. The possible values and their corresponding attributes are listed below. - `reference` - A terminal option. Contains a reference to a Document that can be added to the list of options chosen.   - `item` (APIReference): A reference to the chosen item. - `action` - A terminal option. Contains information describing an action, for use within Multiattack actions.   - `action_name` (string): The name of the action, according to its `name` attribute.   - `count` (number | string): The number of times this action can be repeated if this option is chosen.   - `type` (string = `\"melee\" | \"ranged\" | \"ability\" | \"magic\"`, optional): For attack actions that can be either melee, ranged, abilities, or magic. - `multiple` - When this option is chosen, all of its child options are chosen, and must be resolved the same way as a normal option.   - `items` (array): An array of Option objects. All of them must be taken if the option is chosen. - `choice` - A nested choice. If this option is chosen, the Choice structure contained within must be resolved like a normal Choice structure, and the results are the chosen options.   - `choice` (Choice): The Choice to resolve. - `string` - A terminal option. Contains a reference to a string.   - `string` (string): The string. - `ideal` - A terminal option. Contains information about an ideal.   - `desc` (string): A description of the ideal.   - `alignments` (ApiReference[]): A list of alignments of those who might follow the ideal. - `counted_reference` - A terminal option. Contains a reference to something else in the API along with a count.   - `count` (number): Count.   - `of` (ApiReference): Thing being referenced. - `score_prerequisite` - A terminal option. Contains a reference to an ability score and a minimum score.   - `ability_score` (ApiReference): Ability score being referenced.   - `minimum_score` (number): The minimum score required to satisfy the prerequisite. - `ability_bonus` - A terminal option. Contains a reference to an ability score and a bonus   - `ability_score` (ApiReference): Ability score being referenced   - `bonus` (number): The bonus being applied to the ability score - `breath` - A terminal option: Contains a reference to information about a breath attack.   - `name` (string): Name of the breath.   - `dc` (DC): Difficulty check of the breath attack.   - `damage` ([Damage]): Damage dealt by the breath attack, if any. - `damage` - A terminal option. Contains information about damage.   - `damage_type` (ApiReference): Reference to type of damage.   - `damage_dice` (string): Damage expressed in dice (e.g. \"13d6\").   - `notes` (string): Information regarding the damage.  ## FAQ  ### What is the SRD? The SRD, or Systems Reference Document, contains guidelines for publishing content under the OGL. This allows for some of the data for D&D 5e to be open source. The API only covers data that can be found in the SRD. [Here's a link to the full text of the SRD.](https://media.wizards.com/2016/downloads/DND/SRD-OGL_V5.1.pdf)  ### What is the OGL? The Open Game License (OGL) is a public copyright license by Wizards of the Coast that may be used by tabletop role-playing game developers to grant permission to modify, copy, and redistribute some of the content designed for their games, notably game mechanics. However, they must share-alike copies and derivative works. [More information about the OGL can be found here.](https://en.wikipedia.org/wiki/Open_Game_License)  ### A monster, spell, subclass, etc. is missing from the API / Database. Can I add it? Please check if the data is within the SRD. If it is, feel free to open an issue or PR to add it yourself. Otherwise, due to legal reasons, we cannot add it.  ### Can this API be self hosted? Yes it can! You can also host the data yourself if you don't want to use the API at all. You can also make changes and add extra data if you like. However, it is up to you to merge in new changes to the data and API.  #### Can I publish is on <insert platform>? Is this free use? Yes, you can. The API itself is under the [MIT license](https://opensource.org/licenses/MIT), and the underlying data accessible via the API is supported under the SRD and OGL.  # Status Page  The status page for the API can be found here: https://5e-bits.github.io/dnd-uptime/  # Chat  Come hang out with us [on Discord](https://discord.gg/TQuYTv7)!  # Contribute  This API is built from two repositories.   - The repo containing the data lives here: https://github.com/bagelbits/5e-database   - The repo with the API implementation lives here: https://github.com/bagelbits/5e-srd-api  This is a evolving API and having fresh ideas are always welcome! You can open an issue in either repo, open a PR for changes, or just discuss with other users in this discord. 

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

)

var (
	jsonCheck = regexp.MustCompile(`(?i:(?:application|text)/(?:vnd\.[^;]+\+)?json)`)
	xmlCheck  = regexp.MustCompile(`(?i:(?:application|text)/xml)`)
	queryParamSplit = regexp.MustCompile(`(^|&)([^&]+)`)
	queryDescape    = strings.NewReplacer( "%5B", "[", "%5D", "]" )
)

// APIClient manages communication with the D&amp;D 5e API API v0.1
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	CharacterDataAPI *CharacterDataAPIService

	ClassAPI *ClassAPIService

	ClassLevelsAPI *ClassLevelsAPIService

	ClassResourceListsAPI *ClassResourceListsAPIService

	CommonAPI *CommonAPIService

	EquipmentAPI *EquipmentAPIService

	FeatsAPI *FeatsAPIService

	FeaturesAPI *FeaturesAPIService

	GameMechanicsAPI *GameMechanicsAPIService

	MonstersAPI *MonstersAPIService

	RacesAPI *RacesAPIService

	RulesAPI *RulesAPIService

	SpellsAPI *SpellsAPIService

	SubclassesAPI *SubclassesAPIService

	SubracesAPI *SubracesAPIService

	TraitsAPI *TraitsAPIService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.CharacterDataAPI = (*CharacterDataAPIService)(&c.common)
	c.ClassAPI = (*ClassAPIService)(&c.common)
	c.ClassLevelsAPI = (*ClassLevelsAPIService)(&c.common)
	c.ClassResourceListsAPI = (*ClassResourceListsAPIService)(&c.common)
	c.CommonAPI = (*CommonAPIService)(&c.common)
	c.EquipmentAPI = (*EquipmentAPIService)(&c.common)
	c.FeatsAPI = (*FeatsAPIService)(&c.common)
	c.FeaturesAPI = (*FeaturesAPIService)(&c.common)
	c.GameMechanicsAPI = (*GameMechanicsAPIService)(&c.common)
	c.MonstersAPI = (*MonstersAPIService)(&c.common)
	c.RacesAPI = (*RacesAPIService)(&c.common)
	c.RulesAPI = (*RulesAPIService)(&c.common)
	c.SpellsAPI = (*SpellsAPIService)(&c.common)
	c.SubclassesAPI = (*SubclassesAPIService)(&c.common)
	c.SubracesAPI = (*SubracesAPIService)(&c.common)
	c.TraitsAPI = (*TraitsAPIService)(&c.common)

	return c
}

func atoi(in string) (int, error) {
	return strconv.Atoi(in)
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insensitive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.EqualFold(a, needle) {
			return true
		}
	}
	return false
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj interface{}, expected string, name string) error {
	// Make sure there is an object.
	if obj == nil {
		return nil
	}

	// Check the type is as expected.
	if reflect.TypeOf(obj).String() != expected {
		return fmt.Errorf("expected %s to be of type %s but received %s", name, expected, reflect.TypeOf(obj).String())
	}
	return nil
}

func parameterValueToString( obj interface{}, key string ) string {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		return fmt.Sprintf("%v", obj)
	}
	var param,ok = obj.(MappedNullable)
	if !ok {
		return ""
	}
	dataMap,err := param.ToMap()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%v", dataMap[key])
}

// parameterAddToHeaderOrQuery adds the provided object to the request header or url query
// supporting deep object syntax
func parameterAddToHeaderOrQuery(headerOrQueryParams interface{}, keyPrefix string, obj interface{}, collectionType string) {
	var v = reflect.ValueOf(obj)
	var value = ""
	if v == reflect.ValueOf(nil) {
		value = "null"
	} else {
		switch v.Kind() {
			case reflect.Invalid:
				value = "invalid"

			case reflect.Struct:
				if t,ok := obj.(MappedNullable); ok {
					dataMap,err := t.ToMap()
					if err != nil {
						return
					}
					parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, dataMap, collectionType)
					return
				}
				if t, ok := obj.(time.Time); ok {
					parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, t.Format(time.RFC3339), collectionType)
					return
				}
				value = v.Type().String() + " value"
			case reflect.Slice:
				var indValue = reflect.ValueOf(obj)
				if indValue == reflect.ValueOf(nil) {
					return
				}
				var lenIndValue = indValue.Len()
				for i:=0;i<lenIndValue;i++ {
					var arrayValue = indValue.Index(i)
					parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, arrayValue.Interface(), collectionType)
				}
				return

			case reflect.Map:
				var indValue = reflect.ValueOf(obj)
				if indValue == reflect.ValueOf(nil) {
					return
				}
				iter := indValue.MapRange()
				for iter.Next() {
					k,v := iter.Key(), iter.Value()
					parameterAddToHeaderOrQuery(headerOrQueryParams, fmt.Sprintf("%s[%s]", keyPrefix, k.String()), v.Interface(), collectionType)
				}
				return

			case reflect.Interface:
				fallthrough
			case reflect.Ptr:
				parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, v.Elem().Interface(), collectionType)
				return

			case reflect.Int, reflect.Int8, reflect.Int16,
				reflect.Int32, reflect.Int64:
				value = strconv.FormatInt(v.Int(), 10)
			case reflect.Uint, reflect.Uint8, reflect.Uint16,
				reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				value = strconv.FormatUint(v.Uint(), 10)
			case reflect.Float32, reflect.Float64:
				value = strconv.FormatFloat(v.Float(), 'g', -1, 32)
			case reflect.Bool:
				value = strconv.FormatBool(v.Bool())
			case reflect.String:
				value = v.String()
			default:
				value = v.Type().String() + " value"
		}
	}

	switch valuesMap := headerOrQueryParams.(type) {
		case url.Values:
			if collectionType == "csv" && valuesMap.Get(keyPrefix) != "" {
				valuesMap.Set(keyPrefix, valuesMap.Get(keyPrefix) + "," + value)
			} else {
				valuesMap.Add(keyPrefix, value)
			}
			break
		case map[string]string:
			valuesMap[keyPrefix] = value
			break
	}
}

// helper for converting interface{} parameters to json strings
func parameterToJson(obj interface{}) (string, error) {
	jsonBuf, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(jsonBuf), err
}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	if c.cfg.Debug {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		log.Printf("\n%s\n", string(dump))
	}

	resp, err := c.cfg.HTTPClient.Do(request)
	if err != nil {
		return resp, err
	}

	if c.cfg.Debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return resp, err
		}
		log.Printf("\n%s\n", string(dump))
	}
	return resp, err
}

// Allow modification of underlying config for alternate implementations and testing
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior
func (c *APIClient) GetConfig() *Configuration {
	return c.cfg
}

type formFile struct {
		fileBytes []byte
		fileName string
		formFileName string
}

// prepareRequest build the request
func (c *APIClient) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	formFiles []formFile) (localVarRequest *http.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// add form parameters and file if available.
	if strings.HasPrefix(headerParams["Content-Type"], "multipart/form-data") && len(formParams) > 0 || (len(formFiles) > 0) {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and multipart form at the same time.")
		}
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range formParams {
			for _, iv := range v {
				if strings.HasPrefix(k, "@") { // file
					err = addFile(w, k[1:], iv)
					if err != nil {
						return nil, err
					}
				} else { // form value
					w.WriteField(k, iv)
				}
			}
		}
		for _, formFile := range formFiles {
			if len(formFile.fileBytes) > 0 && formFile.fileName != "" {
				w.Boundary()
				part, err := w.CreateFormFile(formFile.formFileName, filepath.Base(formFile.fileName))
				if err != nil {
						return nil, err
				}
				_, err = part.Write(formFile.fileBytes)
				if err != nil {
						return nil, err
				}
			}
		}

		// Set the Boundary in the Content-Type
		headerParams["Content-Type"] = w.FormDataContentType()

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
		w.Close()
	}

	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(formParams) > 0 {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and x-www-form-urlencoded form at the same time.")
		}
		body = &bytes.Buffer{}
		body.WriteString(formParams.Encode())
		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Override request host, if applicable
	if c.cfg.Host != "" {
		url.Host = c.cfg.Host
	}

	// Override request scheme, if applicable
	if c.cfg.Scheme != "" {
		url.Scheme = c.cfg.Scheme
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = queryParamSplit.ReplaceAllStringFunc(query.Encode(), func(s string) string {
		pieces := strings.Split(s, "=")
		pieces[0] = queryDescape.Replace(pieces[0])
		return strings.Join(pieces, "=")
	})

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers[h] = []string{v}
		}
		localVarRequest.Header = headers
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("User-Agent", c.cfg.UserAgent)

	if ctx != nil {
		// add context to the request
		localVarRequest = localVarRequest.WithContext(ctx)

		// Walk through any authentication.

	}

	for header, value := range c.cfg.DefaultHeader {
		localVarRequest.Header.Add(header, value)
	}
	return localVarRequest, nil
}

func (c *APIClient) decode(v interface{}, b []byte, contentType string) (err error) {
	if len(b) == 0 {
		return nil
	}
	if s, ok := v.(*string); ok {
		*s = string(b)
		return nil
	}
	if f, ok := v.(*os.File); ok {
		f, err = os.CreateTemp("", "HttpClientFile")
		if err != nil {
			return
		}
		_, err = f.Write(b)
		if err != nil {
			return
		}
		_, err = f.Seek(0, io.SeekStart)
		return
	}
	if f, ok := v.(**os.File); ok {
		*f, err = os.CreateTemp("", "HttpClientFile")
		if err != nil {
			return
		}
		_, err = (*f).Write(b)
		if err != nil {
			return
		}
		_, err = (*f).Seek(0, io.SeekStart)
		return
	}
	if xmlCheck.MatchString(contentType) {
		if err = xml.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}
	if jsonCheck.MatchString(contentType) {
		if actualObj, ok := v.(interface{ GetActualInstance() interface{} }); ok { // oneOf, anyOf schemas
			if unmarshalObj, ok := actualObj.(interface{ UnmarshalJSON([]byte) error }); ok { // make sure it has UnmarshalJSON defined
				if err = unmarshalObj.UnmarshalJSON(b); err != nil {
					return err
				}
			} else {
				return errors.New("Unknown type with GetActualInstance but no unmarshalObj.UnmarshalJSON defined")
			}
		} else if err = json.Unmarshal(b, v); err != nil { // simple model
			return err
		}
		return nil
	}
	return errors.New("undefined response type")
}

// Add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

// Prevent trying to import "fmt"
func reportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

// A wrapper for strict JSON decoding
func newStrictDecoder(data []byte) *json.Decoder {
	dec := json.NewDecoder(bytes.NewBuffer(data))
	dec.DisallowUnknownFields()
	return dec
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if fp, ok := body.(*os.File); ok {
		_, err = bodyBuf.ReadFrom(fp)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if jsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if xmlCheck.MatchString(contentType) {
		err = xml.NewEncoder(bodyBuf).Encode(body)
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("invalid body type %s\n", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		} else {
			expires = now.Add(lifetime)
		}
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func strlen(s string) int {
	return utf8.RuneCountInString(s)
}

// GenericOpenAPIError Provides access to the body, error and model on returned errors.
type GenericOpenAPIError struct {
	body  []byte
	error string
	model interface{}
}

// Error returns non-empty string if there was an error.
func (e GenericOpenAPIError) Error() string {
	return e.error
}

// Body returns the raw bytes of the response
func (e GenericOpenAPIError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericOpenAPIError) Model() interface{} {
	return e.model
}

// format error message using title and detail when model implements rfc7807
func formatErrorMessage(status string, v interface{}) string {
	str := ""
	metaValue := reflect.ValueOf(v).Elem()

	if metaValue.Kind() == reflect.Struct {
		field := metaValue.FieldByName("Title")
		if field != (reflect.Value{}) {
			str = fmt.Sprintf("%s", field.Interface())
		}

		field = metaValue.FieldByName("Detail")
		if field != (reflect.Value{}) {
			str = fmt.Sprintf("%s (%s)", str, field.Interface())
		}
	}

	return strings.TrimSpace(fmt.Sprintf("%s %s", status, str))
}
