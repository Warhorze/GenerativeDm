# openapi_client.FeaturesApi

All URIs are relative to *https://www.dnd5epypyapi.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_features_index_get**](FeaturesApi.md#api_features_index_get) | **GET** /api/features/{index} | Get a feature by index.


# **api_features_index_get**
> Feature api_features_index_get(index)

Get a feature by index.

# Feature   When you gain a new level in a class, you get its features for that level.  You don’t, however, receive the class’s starting Equipment, and a few  features have additional rules when you’re multiclassing: Channel Divinity,  Extra Attack, Unarmored Defense, and Spellcasting. 

### Example

```python
import time
import os
import openapi_client
from openapi_client.models.feature import Feature
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
    api_instance = openapi_client.FeaturesApi(api_client)
    index = 'action-surge-1-use' # str | The `index` of the feature to get.  Available values can be found in the [`ResourceList`](#get-/api/-endpoint-) for `features`. 

    try:
        # Get a feature by index.
        api_response = api_instance.api_features_index_get(index)
        print("The response of FeaturesApi->api_features_index_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FeaturesApi->api_features_index_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the feature to get.  Available values can be found in the [&#x60;ResourceList&#x60;](#get-/api/-endpoint-) for &#x60;features&#x60;.  | 

### Return type

[**Feature**](Feature.md)

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

