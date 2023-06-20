# \MonstersAPI

All URIs are relative to *https://www.dnd5eapi.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiMonstersGet**](MonstersAPI.md#ApiMonstersGet) | **Get** /api/monsters | Get list of monsters with optional filtering
[**ApiMonstersIndexGet**](MonstersAPI.md#ApiMonstersIndexGet) | **Get** /api/monsters/{index} | Get monster by index.



## ApiMonstersGet

> APIReferenceList ApiMonstersGet(ctx).ChallengeRating(challengeRating).Execute()

Get list of monsters with optional filtering

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
    challengeRating := []float32{float32(123)} // []float32 | The challenge rating or ratings to filter on. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonstersAPI.ApiMonstersGet(context.Background()).ChallengeRating(challengeRating).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonstersAPI.ApiMonstersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiMonstersGet`: APIReferenceList
    fmt.Fprintf(os.Stdout, "Response from `MonstersAPI.ApiMonstersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiMonstersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **challengeRating** | **[]float32** | The challenge rating or ratings to filter on. | 

### Return type

[**APIReferenceList**](APIReferenceList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiMonstersIndexGet

> Monster ApiMonstersIndexGet(ctx, index).Execute()

Get monster by index.

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
    index := "aboleth" // string | The `index` of the `Monster` to get. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonstersAPI.ApiMonstersIndexGet(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonstersAPI.ApiMonstersIndexGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiMonstersIndexGet`: Monster
    fmt.Fprintf(os.Stdout, "Response from `MonstersAPI.ApiMonstersIndexGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the &#x60;Monster&#x60; to get.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiMonstersIndexGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Monster**](Monster.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

