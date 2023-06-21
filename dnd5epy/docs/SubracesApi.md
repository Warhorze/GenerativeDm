# dnd5epy.SubracesApi

All URIs are relative to *https://www.dnd5eapi.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_subraces_index_get**](SubracesApi.md#api_subraces_index_get) | **GET** /api/subraces/{index} | Get a subrace by index.
[**api_subraces_index_proficiencies_get**](SubracesApi.md#api_subraces_index_proficiencies_get) | **GET** /api/subraces/{index}/proficiencies | Get proficiences available for a subrace.
[**api_subraces_index_traits_get**](SubracesApi.md#api_subraces_index_traits_get) | **GET** /api/subraces/{index}/traits | Get traits available for a subrace.


# **api_subraces_index_get**
> Subrace api_subraces_index_get(index)

Get a subrace by index.

Subraces reflect the different varieties of a certain parent race.

### Example

```python
import time
import os
import dnd5epy
from dnd5epy.models.subrace import Subrace
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
    api_instance = dnd5epy.SubracesApi(api_client)
    index = 'hill-dwarf' # str | The `index` of the subrace to get. 

    try:
        # Get a subrace by index.
        api_response = api_instance.api_subraces_index_get(index)
        print("The response of SubracesApi->api_subraces_index_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SubracesApi->api_subraces_index_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the subrace to get.  | 

### Return type

[**Subrace**](Subrace.md)

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

# **api_subraces_index_proficiencies_get**
> APIReferenceList api_subraces_index_proficiencies_get(index)

Get proficiences available for a subrace.

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
    api_instance = dnd5epy.SubracesApi(api_client)
    index = 'hill-dwarf' # str | The `index` of the subrace to get. 

    try:
        # Get proficiences available for a subrace.
        api_response = api_instance.api_subraces_index_proficiencies_get(index)
        print("The response of SubracesApi->api_subraces_index_proficiencies_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SubracesApi->api_subraces_index_proficiencies_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the subrace to get.  | 

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
**200** | List of proficiences for the subrace. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **api_subraces_index_traits_get**
> APIReferenceList api_subraces_index_traits_get(index)

Get traits available for a subrace.

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
    api_instance = dnd5epy.SubracesApi(api_client)
    index = 'hill-dwarf' # str | The `index` of the subrace to get. 

    try:
        # Get traits available for a subrace.
        api_response = api_instance.api_subraces_index_traits_get(index)
        print("The response of SubracesApi->api_subraces_index_traits_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SubracesApi->api_subraces_index_traits_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the subrace to get.  | 

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
**200** | List of traits for the subrace. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

