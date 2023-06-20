# openapi_client.RulesApi

All URIs are relative to *https://www.dnd5eapi.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_rule_sections_index_get**](RulesApi.md#api_rule_sections_index_get) | **GET** /api/rule-sections/{index} | Get a rule section by index.
[**api_rules_index_get**](RulesApi.md#api_rules_index_get) | **GET** /api/rules/{index} | Get a rule by index.


# **api_rule_sections_index_get**
> RuleSection api_rule_sections_index_get(index)

Get a rule section by index.

Rule sections represent a sub-heading and text that can be found underneath a rule heading in the SRD.

### Example

```python
import time
import os
import openapi_client
from openapi_client.models.rule_section import RuleSection
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
    api_instance = openapi_client.RulesApi(api_client)
    index = 'traps' # str | The `index` of the rule section to get. 

    try:
        # Get a rule section by index.
        api_response = api_instance.api_rule_sections_index_get(index)
        print("The response of RulesApi->api_rule_sections_index_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling RulesApi->api_rule_sections_index_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the rule section to get.  | 

### Return type

[**RuleSection**](RuleSection.md)

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

# **api_rules_index_get**
> Rule api_rules_index_get(index)

Get a rule by index.

# Rule   Rules are pages in the SRD that document the mechanics of Dungeons and Dragons.  Rules have descriptions which is the text directly underneath the rule heading  in the SRD. Rules also have subsections for each heading underneath the rule in the SRD. 

### Example

```python
import time
import os
import openapi_client
from openapi_client.models.rule import Rule
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
    api_instance = openapi_client.RulesApi(api_client)
    index = 'adventuring' # str | The `index` of the rule to get. 

    try:
        # Get a rule by index.
        api_response = api_instance.api_rules_index_get(index)
        print("The response of RulesApi->api_rules_index_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling RulesApi->api_rules_index_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **str**| The &#x60;index&#x60; of the rule to get.  | 

### Return type

[**Rule**](Rule.md)

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

