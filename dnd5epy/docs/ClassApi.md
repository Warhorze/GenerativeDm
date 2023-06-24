# dnd5epy.ClassApi

All URIs are relative to *https://www.dnd5epypyapi.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_classes_index_get**](ClassApi.md#api_classes_index_get) | **GET** /api/classes/{index} | Get a class by index.
[**api_classes_index_multi_classing_get**](ClassApi.md#api_classes_index_multi_classing_get) | **GET** /api/classes/{index}/multi-classing | Get multiclassing resource for a class.
[**api_classes_index_spellcasting_get**](ClassApi.md#api_classes_index_spellcasting_get) | **GET** /api/classes/{index}/spellcasting | Get spellcasting info for a class.


# **api_classes_index_get**
> ModelClass api_classes_index_get(index)

Get a class by index.

# Class  A character class is a fundamental part of the identity and nature of characters in the Dungeons & Dragons role-playing game. A character's capabilities, strengths, and weaknesses are largely defined by its class. A character's class affects a character's available skills and abilities. [[SRD p8-55](https://media.wizards.com/2016/downloads/DND/SRD-OGL_V5.1.pdf#page=8)] 

### Example

```python
import time
import os
import dnd5epy
from dnd5epy.models.model_class import ModelClass
from dnd5epy.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://www.dnd5epypyapi.co
# See configuration.py for a list of all supported configuration parameters.
configuration = dnd5epy.Configuration(
    host = "https://www.dnd5epypyapi.co"
)


# Enter a context with an instance of the API client
with dnd5epy.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = dnd5epy.ClassApi(api_client)
    index = 'paladin' # str | The `index` of the class to get. 

    try:
        # Get a class by index.
        api_response = api_instance.api_classes_index_get(index)
        print("The response of ClassApi->api_classes_index_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ClassApi->api_classes_index_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the class to get.  | 

### Return type

[**ModelClass**](ModelClass.md)

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

# **api_classes_index_multi_classing_get**
> Multiclassing api_classes_index_multi_classing_get(index)

Get multiclassing resource for a class.

### Example

```python
import time
import os
import dnd5epy
from dnd5epy.models.multiclassing import Multiclassing
from dnd5epy.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://www.dnd5epypyapi.co
# See configuration.py for a list of all supported configuration parameters.
configuration = dnd5epy.Configuration(
    host = "https://www.dnd5epypyapi.co"
)


# Enter a context with an instance of the API client
with dnd5epy.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = dnd5epy.ClassApi(api_client)
    index = 'paladin' # str | The `index` of the class to get. 

    try:
        # Get multiclassing resource for a class.
        api_response = api_instance.api_classes_index_multi_classing_get(index)
        print("The response of ClassApi->api_classes_index_multi_classing_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ClassApi->api_classes_index_multi_classing_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the class to get.  | 

### Return type

[**Multiclassing**](Multiclassing.md)

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

# **api_classes_index_spellcasting_get**
> Spellcasting api_classes_index_spellcasting_get(index)

Get spellcasting info for a class.

### Example

```python
import time
import os
import dnd5epy
from dnd5epy.models.spellcasting import Spellcasting
from dnd5epy.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://www.dnd5epypyapi.co
# See configuration.py for a list of all supported configuration parameters.
configuration = dnd5epy.Configuration(
    host = "https://www.dnd5epypyapi.co"
)


# Enter a context with an instance of the API client
with dnd5epy.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = dnd5epy.ClassApi(api_client)
    index = 'paladin' # str | The `index` of the class to get. 

    try:
        # Get spellcasting info for a class.
        api_response = api_instance.api_classes_index_spellcasting_get(index)
        print("The response of ClassApi->api_classes_index_spellcasting_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ClassApi->api_classes_index_spellcasting_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the class to get.  | 

### Return type

[**Spellcasting**](Spellcasting.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**404** | Not found. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

