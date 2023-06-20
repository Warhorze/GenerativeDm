# \SubclassesAPI

All URIs are relative to *https://www.dnd5eapi.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiSubclassesIndexFeaturesGet**](SubclassesAPI.md#ApiSubclassesIndexFeaturesGet) | **Get** /api/subclasses/{index}/features | Get features available for a subclass.
[**ApiSubclassesIndexGet**](SubclassesAPI.md#ApiSubclassesIndexGet) | **Get** /api/subclasses/{index} | Get a subclass by index.
[**ApiSubclassesIndexLevelsGet**](SubclassesAPI.md#ApiSubclassesIndexLevelsGet) | **Get** /api/subclasses/{index}/levels | Get all level resources for a subclass.
[**ApiSubclassesIndexLevelsSubclassLevelFeaturesGet**](SubclassesAPI.md#ApiSubclassesIndexLevelsSubclassLevelFeaturesGet) | **Get** /api/subclasses/{index}/levels/{subclass_level}/features | Get features of the requested spell level available to the class.
[**ApiSubclassesIndexLevelsSubclassLevelGet**](SubclassesAPI.md#ApiSubclassesIndexLevelsSubclassLevelGet) | **Get** /api/subclasses/{index}/levels/{subclass_level} | Get level resources for a subclass and level.



## ApiSubclassesIndexFeaturesGet

> APIReferenceList ApiSubclassesIndexFeaturesGet(ctx, index).Execute()

Get features available for a subclass.

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
    index := "fiend" // string | The `index` of the subclass to get. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubclassesAPI.ApiSubclassesIndexFeaturesGet(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubclassesAPI.ApiSubclassesIndexFeaturesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSubclassesIndexFeaturesGet`: APIReferenceList
    fmt.Fprintf(os.Stdout, "Response from `SubclassesAPI.ApiSubclassesIndexFeaturesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the subclass to get.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSubclassesIndexFeaturesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ApiSubclassesIndexGet

> Subclass ApiSubclassesIndexGet(ctx, index).Execute()

Get a subclass by index.



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
    index := "fiend" // string | The `index` of the subclass to get. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubclassesAPI.ApiSubclassesIndexGet(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubclassesAPI.ApiSubclassesIndexGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSubclassesIndexGet`: Subclass
    fmt.Fprintf(os.Stdout, "Response from `SubclassesAPI.ApiSubclassesIndexGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the subclass to get.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSubclassesIndexGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Subclass**](Subclass.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiSubclassesIndexLevelsGet

> []SubclassLevelResource ApiSubclassesIndexLevelsGet(ctx, index).Execute()

Get all level resources for a subclass.

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
    index := "fiend" // string | The `index` of the subclass to get. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubclassesAPI.ApiSubclassesIndexLevelsGet(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubclassesAPI.ApiSubclassesIndexLevelsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSubclassesIndexLevelsGet`: []SubclassLevelResource
    fmt.Fprintf(os.Stdout, "Response from `SubclassesAPI.ApiSubclassesIndexLevelsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the subclass to get.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSubclassesIndexLevelsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SubclassLevelResource**](SubclassLevelResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiSubclassesIndexLevelsSubclassLevelFeaturesGet

> APIReferenceList ApiSubclassesIndexLevelsSubclassLevelFeaturesGet(ctx, index, subclassLevel).Execute()

Get features of the requested spell level available to the class.

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
    index := "fiend" // string | The `index` of the subclass to get. 
    subclassLevel := int32(6) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubclassesAPI.ApiSubclassesIndexLevelsSubclassLevelFeaturesGet(context.Background(), index, subclassLevel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubclassesAPI.ApiSubclassesIndexLevelsSubclassLevelFeaturesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSubclassesIndexLevelsSubclassLevelFeaturesGet`: APIReferenceList
    fmt.Fprintf(os.Stdout, "Response from `SubclassesAPI.ApiSubclassesIndexLevelsSubclassLevelFeaturesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the subclass to get.  | 
**subclassLevel** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSubclassesIndexLevelsSubclassLevelFeaturesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## ApiSubclassesIndexLevelsSubclassLevelGet

> SubclassLevel ApiSubclassesIndexLevelsSubclassLevelGet(ctx, index, subclassLevel).Execute()

Get level resources for a subclass and level.

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
    index := "fiend" // string | The `index` of the subclass to get. 
    subclassLevel := int32(6) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubclassesAPI.ApiSubclassesIndexLevelsSubclassLevelGet(context.Background(), index, subclassLevel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubclassesAPI.ApiSubclassesIndexLevelsSubclassLevelGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSubclassesIndexLevelsSubclassLevelGet`: SubclassLevel
    fmt.Fprintf(os.Stdout, "Response from `SubclassesAPI.ApiSubclassesIndexLevelsSubclassLevelGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the subclass to get.  | 
**subclassLevel** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSubclassesIndexLevelsSubclassLevelGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SubclassLevel**](SubclassLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

