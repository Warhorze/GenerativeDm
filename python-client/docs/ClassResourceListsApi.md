# openapi_client.ClassResourceListsApi

All URIs are relative to *https://www.dnd5eapi.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_classes_index_features_get**](ClassResourceListsApi.md#api_classes_index_features_get) | **GET** /api/classes/{index}/features | Get features available for a class.
[**api_classes_index_proficiencies_get**](ClassResourceListsApi.md#api_classes_index_proficiencies_get) | **GET** /api/classes/{index}/proficiencies | Get proficiencies available for a class.
[**api_classes_index_spells_get**](ClassResourceListsApi.md#api_classes_index_spells_get) | **GET** /api/classes/{index}/spells | Get spells available for a class.
[**api_classes_index_subclasses_get**](ClassResourceListsApi.md#api_classes_index_subclasses_get) | **GET** /api/classes/{index}/subclasses | Get subclasses available for a class.


# **api_classes_index_features_get**
> APIReferenceList api_classes_index_features_get(index)

Get features available for a class.

### Example

```python
import time
import os
import openapi_client
from openapi_client.models.api_reference_list import APIReferenceList
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
    api_instance = openapi_client.ClassResourceListsApi(api_client)
    index = 'paladin' # str | The `index` of the class to get. 

    try:
        # Get features available for a class.
        api_response = api_instance.api_classes_index_features_get(index)
        print("The response of ClassResourceListsApi->api_classes_index_features_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ClassResourceListsApi->api_classes_index_features_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the class to get.  | 

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
**200** | List of features for the class. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **api_classes_index_proficiencies_get**
> APIReferenceList api_classes_index_proficiencies_get(index)

Get proficiencies available for a class.

### Example

```python
import time
import os
import openapi_client
from openapi_client.models.api_reference_list import APIReferenceList
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
    api_instance = openapi_client.ClassResourceListsApi(api_client)
    index = 'paladin' # str | The `index` of the class to get. 

    try:
        # Get proficiencies available for a class.
        api_response = api_instance.api_classes_index_proficiencies_get(index)
        print("The response of ClassResourceListsApi->api_classes_index_proficiencies_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ClassResourceListsApi->api_classes_index_proficiencies_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the class to get.  | 

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
**200** | List of proficiencies for the class. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **api_classes_index_spells_get**
> APIReferenceList api_classes_index_spells_get(index)

Get spells available for a class.

### Example

```python
import time
import os
import openapi_client
from openapi_client.models.api_reference_list import APIReferenceList
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
    api_instance = openapi_client.ClassResourceListsApi(api_client)
    index = 'paladin' # str | The `index` of the class to get. 

    try:
        # Get spells available for a class.
        api_response = api_instance.api_classes_index_spells_get(index)
        print("The response of ClassResourceListsApi->api_classes_index_spells_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ClassResourceListsApi->api_classes_index_spells_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the class to get.  | 

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

# **api_classes_index_subclasses_get**
> APIReferenceList api_classes_index_subclasses_get(index)

Get subclasses available for a class.

### Example

```python
import time
import os
import openapi_client
from openapi_client.models.api_reference_list import APIReferenceList
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
    api_instance = openapi_client.ClassResourceListsApi(api_client)
    index = 'paladin' # str | The `index` of the class to get. 

    try:
        # Get subclasses available for a class.
        api_response = api_instance.api_classes_index_subclasses_get(index)
        print("The response of ClassResourceListsApi->api_classes_index_subclasses_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ClassResourceListsApi->api_classes_index_subclasses_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the class to get.  | 

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

