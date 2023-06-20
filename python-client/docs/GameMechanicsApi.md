# openapi_client.GameMechanicsApi

All URIs are relative to *https://www.dnd5eapi.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_conditions_index_get**](GameMechanicsApi.md#api_conditions_index_get) | **GET** /api/conditions/{index} | Get a condition by index.
[**api_damage_types_index_get**](GameMechanicsApi.md#api_damage_types_index_get) | **GET** /api/damage-types/{index} | Get a damage type by index.
[**api_magic_schools_index_get**](GameMechanicsApi.md#api_magic_schools_index_get) | **GET** /api/magic-schools/{index} | Get a magic school by index.


# **api_conditions_index_get**
> Condition api_conditions_index_get(index)

Get a condition by index.

# Condition  A condition alters a creature’s capabilities in a variety of ways and can  arise as a result of a spell, a class feature, a monster’s attack, or other  effect. Most conditions, such as blinded, are impairments, but a few, such  as invisible, can be advantageous. 

### Example

```python
import time
import os
import openapi_client
from openapi_client.models.condition import Condition
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://www.dnd5eapi.co
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://www.dnd5eapi.co"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.GameMechanicsApi(api_client)
    index = 'blinded' # str | The `index` of the condition to get. 

    try:
        # Get a condition by index.
        api_response = api_instance.api_conditions_index_get(index)
        print("The response of GameMechanicsApi->api_conditions_index_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling GameMechanicsApi->api_conditions_index_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the condition to get.  | 

### Return type

[**Condition**](Condition.md)

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

# **api_damage_types_index_get**
> DamageType api_damage_types_index_get(index)

Get a damage type by index.

# Damage type  Different attacks, damaging spells, and other harmful effects deal different  types of damage. Damage types have no rules of their own, but other rules,  such as damage resistance, rely on the types. 

### Example

```python
import time
import os
import openapi_client
from openapi_client.models.damage_type import DamageType
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://www.dnd5eapi.co
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://www.dnd5eapi.co"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.GameMechanicsApi(api_client)
    index = 'acid' # str | The `index` of the damage type to get. 

    try:
        # Get a damage type by index.
        api_response = api_instance.api_damage_types_index_get(index)
        print("The response of GameMechanicsApi->api_damage_types_index_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling GameMechanicsApi->api_damage_types_index_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the damage type to get.  | 

### Return type

[**DamageType**](DamageType.md)

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

# **api_magic_schools_index_get**
> MagicSchool api_magic_schools_index_get(index)

Get a magic school by index.

# Magic School  Academies of magic group spells into eight categories called schools of  magic. Scholars, particularly wizards, apply these categories to all spells,  believing that all magic functions in essentially the same way, whether it  derives from rigorous study or is bestowed by a deity. 

### Example

```python
import time
import os
import openapi_client
from openapi_client.models.magic_school import MagicSchool
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://www.dnd5eapi.co
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://www.dnd5eapi.co"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.GameMechanicsApi(api_client)
    index = 'abjuration' # str | The `index` of the magic school to get. 

    try:
        # Get a magic school by index.
        api_response = api_instance.api_magic_schools_index_get(index)
        print("The response of GameMechanicsApi->api_magic_schools_index_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling GameMechanicsApi->api_magic_schools_index_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the magic school to get.  | 

### Return type

[**MagicSchool**](MagicSchool.md)

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

