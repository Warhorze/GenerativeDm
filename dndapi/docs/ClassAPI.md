# \ClassAPI

All URIs are relative to *https://www.dnd5eapi.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiClassesIndexGet**](ClassAPI.md#ApiClassesIndexGet) | **Get** /api/classes/{index} | Get a class by index.
[**ApiClassesIndexMultiClassingGet**](ClassAPI.md#ApiClassesIndexMultiClassingGet) | **Get** /api/classes/{index}/multi-classing | Get multiclassing resource for a class.
[**ApiClassesIndexSpellcastingGet**](ClassAPI.md#ApiClassesIndexSpellcastingGet) | **Get** /api/classes/{index}/spellcasting | Get spellcasting info for a class.



## ApiClassesIndexGet

> Class ApiClassesIndexGet(ctx, index).Execute()

Get a class by index.



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
    index := "paladin" // string | The `index` of the class to get. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClassAPI.ApiClassesIndexGet(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClassAPI.ApiClassesIndexGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiClassesIndexGet`: Class
    fmt.Fprintf(os.Stdout, "Response from `ClassAPI.ApiClassesIndexGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the class to get.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiClassesIndexGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Class**](Class.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiClassesIndexMultiClassingGet

> Multiclassing ApiClassesIndexMultiClassingGet(ctx, index).Execute()

Get multiclassing resource for a class.

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
    index := "paladin" // string | The `index` of the class to get. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClassAPI.ApiClassesIndexMultiClassingGet(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClassAPI.ApiClassesIndexMultiClassingGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiClassesIndexMultiClassingGet`: Multiclassing
    fmt.Fprintf(os.Stdout, "Response from `ClassAPI.ApiClassesIndexMultiClassingGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the class to get.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiClassesIndexMultiClassingGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Multiclassing**](Multiclassing.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiClassesIndexSpellcastingGet

> Spellcasting ApiClassesIndexSpellcastingGet(ctx, index).Execute()

Get spellcasting info for a class.

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
    index := "paladin" // string | The `index` of the class to get. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClassAPI.ApiClassesIndexSpellcastingGet(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClassAPI.ApiClassesIndexSpellcastingGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiClassesIndexSpellcastingGet`: Spellcasting
    fmt.Fprintf(os.Stdout, "Response from `ClassAPI.ApiClassesIndexSpellcastingGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the class to get.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiClassesIndexSpellcastingGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Spellcasting**](Spellcasting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

