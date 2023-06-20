# \FeatsAPI

All URIs are relative to *https://www.dnd5eapi.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiFeatsIndexGet**](FeatsAPI.md#ApiFeatsIndexGet) | **Get** /api/feats/{index} | Get a feat by index.



## ApiFeatsIndexGet

> Feat ApiFeatsIndexGet(ctx, index).Execute()

Get a feat by index.



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
    index := "index_example" // string | The `index` of the feat to get. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatsAPI.ApiFeatsIndexGet(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatsAPI.ApiFeatsIndexGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiFeatsIndexGet`: Feat
    fmt.Fprintf(os.Stdout, "Response from `FeatsAPI.ApiFeatsIndexGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the feat to get.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiFeatsIndexGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Feat**](Feat.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

