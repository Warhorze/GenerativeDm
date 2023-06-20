# openapi_client.TraitsApi

All URIs are relative to *https://www.dnd5eapi.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_traits_index_get**](TraitsApi.md#api_traits_index_get) | **GET** /api/traits/{index} | Get a trait by index.


# **api_traits_index_get**
> Trait api_traits_index_get(index)

Get a trait by index.

### Example

```python
import time
import os
import openapi_client
from openapi_client.models.trait import Trait
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
    api_instance = openapi_client.TraitsApi(api_client)
    index = 'trance' # str | The `index` of the `Trait` to get.

    try:
        # Get a trait by index.
        api_response = api_instance.api_traits_index_get(index)
        print("The response of TraitsApi->api_traits_index_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling TraitsApi->api_traits_index_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the &#x60;Trait&#x60; to get. | 

### Return type

[**Trait**](Trait.md)

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

