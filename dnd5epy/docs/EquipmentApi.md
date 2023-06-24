# openapi_client.EquipmentApi

All URIs are relative to *https://www.dnd5epypyapi.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_equipment_categories_index_get**](EquipmentApi.md#api_equipment_categories_index_get) | **GET** /api/equipment-categories/{index} | Get an equipment category by index.
[**api_equipment_index_get**](EquipmentApi.md#api_equipment_index_get) | **GET** /api/equipment/{index} | Get an equipment item by index.
[**api_magic_items_index_get**](EquipmentApi.md#api_magic_items_index_get) | **GET** /api/magic-items/{index} | Get a magic item by index.
[**api_weapon_properties_index_get**](EquipmentApi.md#api_weapon_properties_index_get) | **GET** /api/weapon-properties/{index} | Get a weapon property by index.


# **api_equipment_categories_index_get**
> EquipmentCategory api_equipment_categories_index_get(index)

Get an equipment category by index.

These are the categories that various equipment fall under.

### Example

```python
import time
import os
import openapi_client
from openapi_client.models.equipment_category import EquipmentCategory
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
    api_instance = openapi_client.EquipmentApi(api_client)
    index = 'waterborne-vehicles' # str | The `index` of the equipment category score to get.  Available values can be found in the resource list for this endpoint. 

    try:
        # Get an equipment category by index.
        api_response = api_instance.api_equipment_categories_index_get(index)
        print("The response of EquipmentApi->api_equipment_categories_index_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling EquipmentApi->api_equipment_categories_index_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the equipment category score to get.  Available values can be found in the resource list for this endpoint.  | 

### Return type

[**EquipmentCategory**](EquipmentCategory.md)

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

# **api_equipment_index_get**
> Equipment api_equipment_index_get(index)

Get an equipment item by index.

# Equipment  Opportunities abound to find treasure, equipment, weapons, armor, and more  in the dungeons you explore. Normally, you can sell your treasures and  trinkets when you return to a town or other settlement, provided that you  can find buyers and merchants interested in your loot. 

### Example

```python
import time
import os
import openapi_client
from openapi_client.models.equipment import Equipment
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
    api_instance = openapi_client.EquipmentApi(api_client)
    index = 'club' # str | The `index` of the equipment to get.  Available values can be found in the [`ResourceList`](#get-/api/-endpoint-) for `equipment`. 

    try:
        # Get an equipment item by index.
        api_response = api_instance.api_equipment_index_get(index)
        print("The response of EquipmentApi->api_equipment_index_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling EquipmentApi->api_equipment_index_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the equipment to get.  Available values can be found in the [&#x60;ResourceList&#x60;](#get-/api/-endpoint-) for &#x60;equipment&#x60;.  | 

### Return type

[**Equipment**](Equipment.md)

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

# **api_magic_items_index_get**
> MagicItem api_magic_items_index_get(index)

Get a magic item by index.

These are the various magic items you can find in the game.

### Example

```python
import time
import os
import openapi_client
from openapi_client.models.magic_item import MagicItem
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
    api_instance = openapi_client.EquipmentApi(api_client)
    index = 'adamantine-armor' # str | The `index` of the magic item to get.  Available values can be found in the resource list for this endpoint. 

    try:
        # Get a magic item by index.
        api_response = api_instance.api_magic_items_index_get(index)
        print("The response of EquipmentApi->api_magic_items_index_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling EquipmentApi->api_magic_items_index_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the magic item to get.  Available values can be found in the resource list for this endpoint.  | 

### Return type

[**MagicItem**](MagicItem.md)

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

# **api_weapon_properties_index_get**
> WeaponProperty api_weapon_properties_index_get(index)

Get a weapon property by index.

### Example

```python
import time
import os
import openapi_client
from openapi_client.models.weapon_property import WeaponProperty
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
    api_instance = openapi_client.EquipmentApi(api_client)
    index = 'ammunition' # str | The `index` of the weapon property to get. 

    try:
        # Get a weapon property by index.
        api_response = api_instance.api_weapon_properties_index_get(index)
        print("The response of EquipmentApi->api_weapon_properties_index_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling EquipmentApi->api_weapon_properties_index_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the weapon property to get.  | 

### Return type

[**WeaponProperty**](WeaponProperty.md)

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

