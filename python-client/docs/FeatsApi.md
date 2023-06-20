# openapi_client.FeatsApi

All URIs are relative to *https://www.dnd5eapi.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_feats_index_get**](FeatsApi.md#api_feats_index_get) | **GET** /api/feats/{index} | Get a feat by index.


# **api_feats_index_get**
> Feat api_feats_index_get(index)

Get a feat by index.

# Feat   A feat is a boon a character can receive at level up instead of an ability score increase. 

### Example

```python
import time
import os
import openapi_client
from openapi_client.models.feat import Feat
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
    api_instance = openapi_client.FeatsApi(api_client)
    index = 'index_example' # str | The `index` of the feat to get. 

    try:
        # Get a feat by index.
        api_response = api_instance.api_feats_index_get(index)
        print("The response of FeatsApi->api_feats_index_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FeatsApi->api_feats_index_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the feat to get.  | 

### Return type

[**Feat**](Feat.md)

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

