# dnd5epy.SpellsApi

All URIs are relative to *https://www.dnd5eapi.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_spells_get**](SpellsApi.md#api_spells_get) | **GET** /api/spells | Get list of spells with optional filtering.
[**api_spells_index_get**](SpellsApi.md#api_spells_index_get) | **GET** /api/spells/{index} | Get a spell by index.


# **api_spells_get**
> APIReferenceList api_spells_get(level=level, school=school)

Get list of spells with optional filtering.

### Example

```python
import time
import os
import dnd5epy
from dnd5epy.models.api_reference_list import APIReferenceList
from dnd5epy.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://www.dnd5eapi.co
# See configuration.py for a list of all supported configuration parameters.
configuration = dnd5epy.Configuration(
    host = "https://www.dnd5eapi.co"
)


# Enter a context with an instance of the API client
with dnd5epy.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = dnd5epy.SpellsApi(api_client)
    level = [[1]] # List[int] | The level or levels to filter on. (optional)
    school = ['[\"illusion\"]'] # List[str] | The magic school or schools to filter on. (optional)

    try:
        # Get list of spells with optional filtering.
        api_response = api_instance.api_spells_get(level=level, school=school)
        print("The response of SpellsApi->api_spells_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SpellsApi->api_spells_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **level** | [**List[int]**](int.md)| The level or levels to filter on. | [optional] 
 **school** | [**List[str]**](str.md)| The magic school or schools to filter on. | [optional] 

### Return type

[**APIReferenceList**](APIReferenceList.md)

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

# **api_spells_index_get**
> Spell api_spells_index_get(index)

Get a spell by index.

### Example

```python
import time
import os
import dnd5epy
from dnd5epy.models.spell import Spell
from dnd5epy.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://www.dnd5eapi.co
# See configuration.py for a list of all supported configuration parameters.
configuration = dnd5epy.Configuration(
    host = "https://www.dnd5eapi.co"
)


# Enter a context with an instance of the API client
with dnd5epy.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = dnd5epy.SpellsApi(api_client)
    index = 'sacred-flame' # str | The `index` of the `Spell` to get.  Available values can be found in the [`ResourceList`](#get-/api/-endpoint-) for `spells`. 

    try:
        # Get a spell by index.
        api_response = api_instance.api_spells_index_get(index)
        print("The response of SpellsApi->api_spells_index_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SpellsApi->api_spells_index_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the &#x60;Spell&#x60; to get.  Available values can be found in the [&#x60;ResourceList&#x60;](#get-/api/-endpoint-) for &#x60;spells&#x60;.  | 

### Return type

[**Spell**](Spell.md)

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

