# \SpellsAPI

All URIs are relative to *https://www.dnd5eapi.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiSpellsGet**](SpellsAPI.md#ApiSpellsGet) | **Get** /api/spells | Get list of spells with optional filtering.
[**ApiSpellsIndexGet**](SpellsAPI.md#ApiSpellsIndexGet) | **Get** /api/spells/{index} | Get a spell by index.



## ApiSpellsGet

> APIReferenceList ApiSpellsGet(ctx).Level(level).School(school).Execute()

Get list of spells with optional filtering.

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
    level := []int32{int32(123)} // []int32 | The level or levels to filter on. (optional)
    school := []string{"Inner_example"} // []string | The magic school or schools to filter on. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpellsAPI.ApiSpellsGet(context.Background()).Level(level).School(school).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpellsAPI.ApiSpellsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSpellsGet`: APIReferenceList
    fmt.Fprintf(os.Stdout, "Response from `SpellsAPI.ApiSpellsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiSpellsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **level** | **[]int32** | The level or levels to filter on. | 
 **school** | **[]string** | The magic school or schools to filter on. | 

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


## ApiSpellsIndexGet

> Spell ApiSpellsIndexGet(ctx, index).Execute()

Get a spell by index.

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
    index := "sacred-flame" // string | The `index` of the `Spell` to get.  Available values can be found in the [`ResourceList`](#get-/api/-endpoint-) for `spells`. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpellsAPI.ApiSpellsIndexGet(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpellsAPI.ApiSpellsIndexGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSpellsIndexGet`: Spell
    fmt.Fprintf(os.Stdout, "Response from `SpellsAPI.ApiSpellsIndexGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the &#x60;Spell&#x60; to get.  Available values can be found in the [&#x60;ResourceList&#x60;](#get-/api/-endpoint-) for &#x60;spells&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSpellsIndexGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Spell**](Spell.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

