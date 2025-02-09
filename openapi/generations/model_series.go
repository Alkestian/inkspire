/*
Inkspire

This app allows for making lists of various types across different media and will eventually have fuzzy searching and rating functionality added to it. 

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Series type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Series{}

// Series struct for Series
type Series struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

type _Series Series

// NewSeries instantiates a new Series object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeries(id string, name string) *Series {
	this := Series{}
	this.Id = id
	this.Name = name
	return &this
}

// NewSeriesWithDefaults instantiates a new Series object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeriesWithDefaults() *Series {
	this := Series{}
	return &this
}

// GetId returns the Id field value
func (o *Series) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Series) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Series) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Series) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Series) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Series) SetName(v string) {
	o.Name = v
}

func (o Series) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Series) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *Series) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSeries := _Series{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSeries)

	if err != nil {
		return err
	}

	*o = Series(varSeries)

	return err
}

type NullableSeries struct {
	value *Series
	isSet bool
}

func (v NullableSeries) Get() *Series {
	return v.value
}

func (v *NullableSeries) Set(val *Series) {
	v.value = val
	v.isSet = true
}

func (v NullableSeries) IsSet() bool {
	return v.isSet
}

func (v *NullableSeries) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeries(val *Series) *NullableSeries {
	return &NullableSeries{value: val, isSet: true}
}

func (v NullableSeries) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeries) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


