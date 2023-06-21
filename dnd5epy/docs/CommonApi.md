# dnd5epy.CommonApi

All URIs are relative to *https://www.dnd5eapi.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_endpoint_get**](CommonApi.md#api_endpoint_get) | **GET** /api/{endpoint} | Get list of all available resources for an endpoint.
[**api_get**](CommonApi.md#api_get) | **GET** /api | Get all resource URLs.


# **api_endpoint_get**
> APIReferenceList api_endpoint_get(endpoint)

Get list of all available resources for an endpoint.

Currently only the [`/spells`](#get-/api/spells) and [`/monsters`](#get-/api/monsters) endpoints support filtering with query parameters. Use of these query parameters is documented under the respective [Spells](#tag--Spells) and [Monsters](#tag--Monsters) sections. 

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
    api_instance = dnd5epy.CommonApi(api_client)
    endpoint = 'ability-scores' # str | 

    try:
        # Get list of all available resources for an endpoint.
        api_response = api_instance.api_endpoint_get(endpoint)
        print("The response of CommonApi->api_endpoint_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CommonApi->api_endpoint_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endpoint** | **str**|  | 

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

# **api_get**
> Dict[str, str] api_get()

Get all resource URLs.

Making a request to the API's base URL returns an object containing available endpoints.

### Example

```python
import time
import os
import dnd5epy
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
    api_instance = dnd5epy.CommonApi(api_client)

    try:
        # Get all resource URLs.
        api_response = api_instance.api_get()
        print("The response of CommonApi->api_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CommonApi->api_get: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

### Return type

**Dict[str, str]**

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

