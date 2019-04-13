package govalidator

import (
	"bytes"
	"encoding/json"
)

// Int describes a custom type of built-in int data type
type Int struct {
	Value int  `json:"value"`
	IsSet bool `json:"isSet"`
}

var null = []byte("null")

// UnmarshalJSON ...
func (i *Int) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, null) {
		return nil
	}
	i.IsSet = true
	var temp int
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	return nil
}

// MarshalJSON ...
func (i *Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Value)
}

// Int64 describes a custom type of built-in int64 data type
type Int64 struct {
	Value int64 `json:"value"`
	IsSet bool  `json:"isSet"`
}

// UnmarshalJSON ...
func (i *Int64) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, null) {
		return nil
	}
	i.IsSet = true
	var temp int64
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	return nil
}

// MarshalJSON ...
func (i *Int64) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Value)
}

// Float32 describes a custom type of built-in float32 data type
type Float32 struct {
	Value float32 `json:"value"`
	IsSet bool    `json:"isSet"`
}

// UnmarshalJSON ...
func (i *Float32) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, null) {
		return nil
	}
	i.IsSet = true
	var temp float32
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	return nil
}

// MarshalJSON ...
func (i *Float32) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Value)
}

// Float64 describes a custom type of built-in float64 data type
type Float64 struct {
	Value float64 `json:"value"`
	IsSet bool    `json:"isSet"`
}

// UnmarshalJSON ...
func (i *Float64) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, null) {
		return nil
	}
	i.IsSet = true
	var temp float64
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	return nil
}

// MarshalJSON ...
func (i *Float64) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Value)
}

// Bool describes a custom type of built-in bool data type
type Bool struct {
	Value bool `json:"value"`
	IsSet bool `json:"isSet"`
}

// UnmarshalJSON ...
func (i *Bool) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, null) {
		return nil
	}
	i.IsSet = true
	var temp bool
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	return nil
}

// MarshalJSON ...
func (i *Bool) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Value)
}
