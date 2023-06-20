# \FeaturesAPI

All URIs are relative to *https://www.dnd5eapi.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiFeaturesIndexGet**](FeaturesAPI.md#ApiFeaturesIndexGet) | **Get** /api/features/{index} | Get a feature by index.



## ApiFeaturesIndexGet

> Feature ApiFeaturesIndexGet(ctx, index).Execute()

Get a feature by index.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    index := "action-surge-1-use" // string | The `index` of the feature to get.  Available values can be found in the [`ResourceList`](#get-/api/-endpoint-) for `features`. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesAPI.ApiFeaturesIndexGet(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.ApiFeaturesIndexGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiFeaturesIndexGet`: Feature
    fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.ApiFeaturesIndexGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the feature to get.  Available values can be found in the [&#x60;ResourceList&#x60;](#get-/api/-endpoint-) for &#x60;features&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiFeaturesIndexGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Feature**](Feature.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

