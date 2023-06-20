# APIReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **string** | Resource index for shorthand searching. | [optional] 
**Name** | Pointer to **string** | Name of the referenced resource. | [optional] 
**Url** | Pointer to **string** | URL of the referenced resource. | [optional] 

## Methods

### NewAPIReference

`func NewAPIReference() *APIReference`

NewAPIReference instantiates a new APIReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIReferenceWithDefaults

`func NewAPIReferenceWithDefaults() *APIReference`

NewAPIReferenceWithDefaults instantiates a new APIReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *APIReference) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *APIReference) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *APIReference) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *APIReference) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetName

`func (o *APIReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *APIReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *APIReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *APIReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *APIReference) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *APIReference) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *APIReference) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *APIReference) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


