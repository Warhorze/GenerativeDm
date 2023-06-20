# DC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DcType** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 
**DcValue** | Pointer to **float32** | Value to beat | [optional] 
**SuccessType** | Pointer to **string** | Result of a successful save. Can be \\\&quot;none\\\&quot;, \\\&quot;half\\\&quot;, or \\\&quot;other\\\&quot; | [optional] 

## Methods

### NewDC

`func NewDC() *DC`

NewDC instantiates a new DC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDCWithDefaults

`func NewDCWithDefaults() *DC`

NewDCWithDefaults instantiates a new DC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDcType

`func (o *DC) GetDcType() APIReference`

GetDcType returns the DcType field if non-nil, zero value otherwise.

### GetDcTypeOk

`func (o *DC) GetDcTypeOk() (*APIReference, bool)`

GetDcTypeOk returns a tuple with the DcType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcType

`func (o *DC) SetDcType(v APIReference)`

SetDcType sets DcType field to given value.

### HasDcType

`func (o *DC) HasDcType() bool`

HasDcType returns a boolean if a field has been set.

### GetDcValue

`func (o *DC) GetDcValue() float32`

GetDcValue returns the DcValue field if non-nil, zero value otherwise.

### GetDcValueOk

`func (o *DC) GetDcValueOk() (*float32, bool)`

GetDcValueOk returns a tuple with the DcValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcValue

`func (o *DC) SetDcValue(v float32)`

SetDcValue sets DcValue field to given value.

### HasDcValue

`func (o *DC) HasDcValue() bool`

HasDcValue returns a boolean if a field has been set.

### GetSuccessType

`func (o *DC) GetSuccessType() string`

GetSuccessType returns the SuccessType field if non-nil, zero value otherwise.

### GetSuccessTypeOk

`func (o *DC) GetSuccessTypeOk() (*string, bool)`

GetSuccessTypeOk returns a tuple with the SuccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessType

`func (o *DC) SetSuccessType(v string)`

SetSuccessType sets SuccessType field to given value.

### HasSuccessType

`func (o *DC) HasSuccessType() bool`

HasSuccessType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


