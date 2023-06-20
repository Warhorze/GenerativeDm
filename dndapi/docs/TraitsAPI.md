# \TraitsAPI

All URIs are relative to *https://www.dnd5eapi.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiTraitsIndexGet**](TraitsAPI.md#ApiTraitsIndexGet) | **Get** /api/traits/{index} | Get a trait by index.



## ApiTraitsIndexGet

> Trait ApiTraitsIndexGet(ctx, index).Execute()

Get a trait by index.

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
    index := "trance" // string | The `index` of the `Trait` to get.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TraitsAPI.ApiTraitsIndexGet(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TraitsAPI.ApiTraitsIndexGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTraitsIndexGet`: Trait
    fmt.Fprintf(os.Stdout, "Response from `TraitsAPI.ApiTraitsIndexGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the &#x60;Trait&#x60; to get. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTraitsIndexGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Trait**](Trait.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

