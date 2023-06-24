# openapi_client.RacesApi

All URIs are relative to *https://www.dnd5epypyapi.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_races_index_get**](RacesApi.md#api_races_index_get) | **GET** /api/races/{index} | Get a race by index.
[**api_races_index_proficiencies_get**](RacesApi.md#api_races_index_proficiencies_get) | **GET** /api/races/{index}/proficiencies | Get proficiencies available for a race.
[**api_races_index_subraces_get**](RacesApi.md#api_races_index_subraces_get) | **GET** /api/races/{index}/subraces | Get subraces available for a race.
[**api_races_index_traits_get**](RacesApi.md#api_races_index_traits_get) | **GET** /api/races/{index}/traits | Get traits available for a race.


# **api_races_index_get**
> Race api_races_index_get(index)

Get a race by index.

Each race grants your character ability and skill bonuses as well as racial traits.

### Example

```python
import time
import os
import openapi_client
from openapi_client.models.race import Race
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
    api_instance = openapi_client.RacesApi(api_client)
    index = 'elf' # str | The `index` of the race to get. 

    try:
        # Get a race by index.
        api_response = api_instance.api_races_index_get(index)
        print("The response of RacesApi->api_races_index_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling RacesApi->api_races_index_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the race to get.  | 

### Return type

[**Race**](Race.md)

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

# **api_races_index_proficiencies_get**
> APIReferenceList api_races_index_proficiencies_get(index)

Get proficiencies available for a race.

### Example

```python
import time
import os
import openapi_client
from openapi_client.models.api_reference_list import APIReferenceList
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
    api_instance = openapi_client.RacesApi(api_client)
    index = 'elf' # str | The `index` of the race to get. 

    try:
        # Get proficiencies available for a race.
        api_response = api_instance.api_races_index_proficiencies_get(index)
        print("The response of RacesApi->api_races_index_proficiencies_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling RacesApi->api_races_index_proficiencies_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the race to get.  | 

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
**200** | List of proficiencies for the race. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **api_races_index_subraces_get**
> APIReferenceList api_races_index_subraces_get(index)

Get subraces available for a race.

### Example

```python
import time
import os
import openapi_client
from openapi_client.models.api_reference_list import APIReferenceList
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
    api_instance = openapi_client.RacesApi(api_client)
    index = 'elf' # str | The `index` of the race to get. 

    try:
        # Get subraces available for a race.
        api_response = api_instance.api_races_index_subraces_get(index)
        print("The response of RacesApi->api_races_index_subraces_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling RacesApi->api_races_index_subraces_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the race to get.  | 

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
**200** | List of subraces for the race. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **api_races_index_traits_get**
> APIReferenceList api_races_index_traits_get(index)

Get traits available for a race.

### Example

```python
import time
import os
import openapi_client
from openapi_client.models.api_reference_list import APIReferenceList
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
    api_instance = openapi_client.RacesApi(api_client)
    index = 'elf' # str | The `index` of the race to get. 

    try:
        # Get traits available for a race.
        api_response = api_instance.api_races_index_traits_get(index)
        print("The response of RacesApi->api_races_index_traits_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling RacesApi->api_races_index_traits_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the race to get.  | 

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
**200** | List of traits for the race. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

