# openapi_client.CharacterDataApi

All URIs are relative to *https://www.dnd5epypyapi.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_ability_scores_index_get**](CharacterDataApi.md#api_ability_scores_index_get) | **GET** /api/ability-scores/{index} | Get an ability score by index.
[**api_alignments_index_get**](CharacterDataApi.md#api_alignments_index_get) | **GET** /api/alignments/{index} | Get an alignment by index.
[**api_backgrounds_index_get**](CharacterDataApi.md#api_backgrounds_index_get) | **GET** /api/backgrounds/{index} | Get a background by index.
[**api_languages_index_get**](CharacterDataApi.md#api_languages_index_get) | **GET** /api/languages/{index} | Get a language by index.
[**api_proficiencies_index_get**](CharacterDataApi.md#api_proficiencies_index_get) | **GET** /api/proficiencies/{index} | Get a proficiency by index.
[**api_skills_index_get**](CharacterDataApi.md#api_skills_index_get) | **GET** /api/skills/{index} | Get a skill by index.


# **api_ability_scores_index_get**
> AbilityScore api_ability_scores_index_get(index)

Get an ability score by index.

# Ability Score  Represents one of the six abilities that describes a creature's physical and mental characteristics. The three main rolls of the game - the ability check, the saving throw, and the attack roll - rely on the ability scores. [[SRD p76](https://media.wizards.com/2016/downloads/DND/SRD-OGL_V5.1.pdf#page=76)] 

### Example

```python
import time
import os
import openapi_client
from openapi_client.models.ability_score import AbilityScore
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://www.dnd5epypyapi.co
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://www.dnd5epypyapi.co"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.CharacterDataApi(api_client)
    index = 'cha' # str | The `index` of the ability score to get. 

    try:
        # Get an ability score by index.
        api_response = api_instance.api_ability_scores_index_get(index)
        print("The response of CharacterDataApi->api_ability_scores_index_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CharacterDataApi->api_ability_scores_index_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the ability score to get.  | 

### Return type

[**AbilityScore**](AbilityScore.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **api_alignments_index_get**
> Alignment api_alignments_index_get(index)

Get an alignment by index.

# Alignment  A typical creature in the game world has an alignment, which broadly describes its moral and personal attitudes. Alignment is a combination of two factors: one identifies morality (good, evil, or neutral), and the other describes attitudes toward society and order (lawful, chaotic, or neutral). Thus, nine distinct alignments define the possible combinations.[[SRD p58](https://media.wizards.com/2016/downloads/DND/SRD-OGL_V5.1.pdf#page=58)] 

### Example

```python
import time
import os
import openapi_client
from openapi_client.models.alignment import Alignment
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://www.dnd5epypyapi.co
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://www.dnd5epypyapi.co"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.CharacterDataApi(api_client)
    index = 'chaotic-neutral' # str | The `index` of the alignment to get. 

    try:
        # Get an alignment by index.
        api_response = api_instance.api_alignments_index_get(index)
        print("The response of CharacterDataApi->api_alignments_index_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CharacterDataApi->api_alignments_index_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the alignment to get.  | 

### Return type

[**Alignment**](Alignment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **api_backgrounds_index_get**
> Background api_backgrounds_index_get(index)

Get a background by index.

# Background  Every story has a beginning. Your character's background reveals where you came from, how you became an adventurer, and your place in the world. Choosing a background provides you with important story cues about your character's identity. [[SRD p60](https://media.wizards.com/2016/downloads/DND/SRD-OGL_V5.1.pdf#page=60)]  _Note:_ acolyte is the only background included in the SRD. 

### Example

```python
import time
import os
import openapi_client
from openapi_client.models.background import Background
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://www.dnd5epypyapi.co
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://www.dnd5epypyapi.co"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.CharacterDataApi(api_client)
    index = 'acolyte' # str | The `index` of the background to get. 

    try:
        # Get a background by index.
        api_response = api_instance.api_backgrounds_index_get(index)
        print("The response of CharacterDataApi->api_backgrounds_index_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CharacterDataApi->api_backgrounds_index_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the background to get.  | 

### Return type

[**Background**](Background.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **api_languages_index_get**
> Language api_languages_index_get(index)

Get a language by index.

# Language  Your race indicates the languages your character can speak by default, and your background might give you access to one or more additional languages of your choice. [[SRD p59](https://media.wizards.com/2016/downloads/DND/SRD-OGL_V5.1.pdf#page=59)] 

### Example

```python
import time
import os
import openapi_client
from openapi_client.models.language import Language
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://www.dnd5epypyapi.co
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://www.dnd5epypyapi.co"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.CharacterDataApi(api_client)
    index = 'abyssal' # str | The `index` of the language to get. 

    try:
        # Get a language by index.
        api_response = api_instance.api_languages_index_get(index)
        print("The response of CharacterDataApi->api_languages_index_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CharacterDataApi->api_languages_index_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the language to get.  | 

### Return type

[**Language**](Language.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **api_proficiencies_index_get**
> Proficiency api_proficiencies_index_get(index)

Get a proficiency by index.

# Proficiency   By virtue of race, class, and background a character is proficient at using certain skills, weapons, and equipment. Characters can also gain additional proficiencies at higher levels or by multiclassing. A characters starting proficiencies are determined during character creation. 

### Example

```python
import time
import os
import openapi_client
from openapi_client.models.proficiency import Proficiency
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://www.dnd5epypyapi.co
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://www.dnd5epypyapi.co"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.CharacterDataApi(api_client)
    index = 'medium-armor' # str | The `index` of the proficiency to get.  Available values can be found in the [`ResourceList`](#get-/api/-endpoint-) for `proficiencies`. 

    try:
        # Get a proficiency by index.
        api_response = api_instance.api_proficiencies_index_get(index)
        print("The response of CharacterDataApi->api_proficiencies_index_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CharacterDataApi->api_proficiencies_index_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the proficiency to get.  Available values can be found in the [&#x60;ResourceList&#x60;](#get-/api/-endpoint-) for &#x60;proficiencies&#x60;.  | 

### Return type

[**Proficiency**](Proficiency.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **api_skills_index_get**
> Skill api_skills_index_get(index)

Get a skill by index.

# Skill  Each ability covers a broad range of capabilities, including skills that a character or a monster can be proficient in. A skill represents a specific aspect of an ability score, and an individual's proficiency in a skill demonstrates a focus on that aspect. [[SRD p77](https://media.wizards.com/2016/downloads/DND/SRD-OGL_V5.1.pdf#page=77)] 

### Example

```python
import time
import os
import openapi_client
from openapi_client.models.skill import Skill
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://www.dnd5epypyapi.co
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://www.dnd5epypyapi.co"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.CharacterDataApi(api_client)
    index = 'nature' # str | The `index` of the skill to get. 

    try:
        # Get a skill by index.
        api_response = api_instance.api_skills_index_get(index)
        print("The response of CharacterDataApi->api_skills_index_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CharacterDataApi->api_skills_index_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the skill to get.  | 

### Return type

[**Skill**](Skill.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

