# Choice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Desc** | Pointer to **string** | Description of the choice to be made. | [optional] 
**Choose** | Pointer to **float32** | Number of items to pick from the list. | [optional] 
**Type** | Pointer to **string** | Type of the resources to choose from. | [optional] 
**From** | Pointer to [**OptionSet**](OptionSet.md) |  | [optional] 

## Methods

### NewChoice

`func NewChoice() *Choice`

NewChoice instantiates a new Choice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChoiceWithDefaults

`func NewChoiceWithDefaults() *Choice`

NewChoiceWithDefaults instantiates a new Choice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesc

`func (o *Choice) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *Choice) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *Choice) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *Choice) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetChoose

`func (o *Choice) GetChoose() float32`

GetChoose returns the Choose field if non-nil, zero value otherwise.

### GetChooseOk

`func (o *Choice) GetChooseOk() (*float32, bool)`

GetChooseOk returns a tuple with the Choose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoose

`func (o *Choice) SetChoose(v float32)`

SetChoose sets Choose field to given value.

### HasChoose

`func (o *Choice) HasChoose() bool`

HasChoose returns a boolean if a field has been set.

### GetType

`func (o *Choice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Choice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Choice) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Choice) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFrom

`func (o *Choice) GetFrom() OptionSet`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Choice) GetFromOk() (*OptionSet, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *Choice) SetFrom(v OptionSet)`

SetFrom sets From field to given value.

### HasFrom

`func (o *Choice) HasFrom() bool`

HasFrom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


