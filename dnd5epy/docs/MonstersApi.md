# dnd5epy.MonstersApi

All URIs are relative to *https://www.dnd5epypyapi.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_monsters_get**](MonstersApi.md#api_monsters_get) | **GET** /api/monsters | Get list of monsters with optional filtering
[**api_monsters_index_get**](MonstersApi.md#api_monsters_index_get) | **GET** /api/monsters/{index} | Get monster by index.


# **api_monsters_get**
> APIReferenceList api_monsters_get(challenge_rating=challenge_rating)

Get list of monsters with optional filtering

### Example

```python
import time
import os
import dnd5epy
from dnd5epy.models.api_reference_list import APIReferenceList
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
    api_instance = dnd5epy.MonstersApi(api_client)
    challenge_rating = [[1]] # List[float] | The challenge rating or ratings to filter on. (optional)

    try:
        # Get list of monsters with optional filtering
        api_response = api_instance.api_monsters_get(challenge_rating=challenge_rating)
        print("The response of MonstersApi->api_monsters_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MonstersApi->api_monsters_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **challenge_rating** | [**List[float]**](float.md)| The challenge rating or ratings to filter on. | [optional] 

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

# **api_monsters_index_get**
> Monster api_monsters_index_get(index)

Get monster by index.

### Example

```python
import time
import os
import dnd5epy
from dnd5epy.models.monster import Monster
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
    api_instance = dnd5epy.MonstersApi(api_client)
    index = 'aboleth' # str | The `index` of the `Monster` to get. 

    try:
        # Get monster by index.
        api_response = api_instance.api_monsters_index_get(index)
        print("The response of MonstersApi->api_monsters_index_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MonstersApi->api_monsters_index_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the &#x60;Monster&#x60; to get.  | 

### Return type

[**Monster**](Monster.md)

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

