# openapi_client.SubclassesApi

All URIs are relative to *https://www.dnd5epypyapi.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_subclasses_index_features_get**](SubclassesApi.md#api_subclasses_index_features_get) | **GET** /api/subclasses/{index}/features | Get features available for a subclass.
[**api_subclasses_index_get**](SubclassesApi.md#api_subclasses_index_get) | **GET** /api/subclasses/{index} | Get a subclass by index.
[**api_subclasses_index_levels_get**](SubclassesApi.md#api_subclasses_index_levels_get) | **GET** /api/subclasses/{index}/levels | Get all level resources for a subclass.
[**api_subclasses_index_levels_subclass_level_features_get**](SubclassesApi.md#api_subclasses_index_levels_subclass_level_features_get) | **GET** /api/subclasses/{index}/levels/{subclass_level}/features | Get features of the requested spell level available to the class.
[**api_subclasses_index_levels_subclass_level_get**](SubclassesApi.md#api_subclasses_index_levels_subclass_level_get) | **GET** /api/subclasses/{index}/levels/{subclass_level} | Get level resources for a subclass and level.


# **api_subclasses_index_features_get**
> APIReferenceList api_subclasses_index_features_get(index)

Get features available for a subclass.

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
    api_instance = openapi_client.SubclassesApi(api_client)
    index = 'fiend' # str | The `index` of the subclass to get. 

    try:
        # Get features available for a subclass.
        api_response = api_instance.api_subclasses_index_features_get(index)
        print("The response of SubclassesApi->api_subclasses_index_features_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SubclassesApi->api_subclasses_index_features_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the subclass to get.  | 

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
**200** | List of features for the subclass. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **api_subclasses_index_get**
> Subclass api_subclasses_index_get(index)

Get a subclass by index.

Subclasses reflect the different paths a class may take as levels are gained.

### Example

```python
import time
import os
import openapi_client
from openapi_client.models.subclass import Subclass
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
    api_instance = openapi_client.SubclassesApi(api_client)
    index = 'fiend' # str | The `index` of the subclass to get. 

    try:
        # Get a subclass by index.
        api_response = api_instance.api_subclasses_index_get(index)
        print("The response of SubclassesApi->api_subclasses_index_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SubclassesApi->api_subclasses_index_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the subclass to get.  | 

### Return type

[**Subclass**](Subclass.md)

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

# **api_subclasses_index_levels_get**
> List[SubclassLevelResource] api_subclasses_index_levels_get(index)

Get all level resources for a subclass.

### Example

```python
import time
import os
import openapi_client
from openapi_client.models.subclass_level_resource import SubclassLevelResource
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
    api_instance = openapi_client.SubclassesApi(api_client)
    index = 'fiend' # str | The `index` of the subclass to get. 

    try:
        # Get all level resources for a subclass.
        api_response = api_instance.api_subclasses_index_levels_get(index)
        print("The response of SubclassesApi->api_subclasses_index_levels_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SubclassesApi->api_subclasses_index_levels_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the subclass to get.  | 

### Return type

[**List[SubclassLevelResource]**](SubclassLevelResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of level resource for the subclass. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **api_subclasses_index_levels_subclass_level_features_get**
> APIReferenceList api_subclasses_index_levels_subclass_level_features_get(index, subclass_level)

Get features of the requested spell level available to the class.

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
    api_instance = openapi_client.SubclassesApi(api_client)
    index = 'fiend' # str | The `index` of the subclass to get. 
    subclass_level = 6 # int | 

    try:
        # Get features of the requested spell level available to the class.
        api_response = api_instance.api_subclasses_index_levels_subclass_level_features_get(index, subclass_level)
        print("The response of SubclassesApi->api_subclasses_index_levels_subclass_level_features_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SubclassesApi->api_subclasses_index_levels_subclass_level_features_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the subclass to get.  | 
 **subclass_level** | **int**|  | 

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
**200** | List of features for the subclass and level. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **api_subclasses_index_levels_subclass_level_get**
> SubclassLevel api_subclasses_index_levels_subclass_level_get(index, subclass_level)

Get level resources for a subclass and level.

### Example

```python
import time
import os
import openapi_client
from openapi_client.models.subclass_level import SubclassLevel
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
    api_instance = openapi_client.SubclassesApi(api_client)
    index = 'fiend' # str | The `index` of the subclass to get. 
    subclass_level = 6 # int | 

    try:
        # Get level resources for a subclass and level.
        api_response = api_instance.api_subclasses_index_levels_subclass_level_get(index, subclass_level)
        print("The response of SubclassesApi->api_subclasses_index_levels_subclass_level_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SubclassesApi->api_subclasses_index_levels_subclass_level_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the subclass to get.  | 
 **subclass_level** | **int**|  | 

### Return type

[**SubclassLevel**](SubclassLevel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Level resource for the subclass and level. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

