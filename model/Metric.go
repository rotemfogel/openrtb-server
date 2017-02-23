package model

import "errors"

// Validation errors
var (
	ErrorMissingType = errors.New("openrtb::model Metric type missing")
	ErrorWrongValue = errors.New("openrtb::model Metric value is wrong")
)

type Metric struct {
	Type   string      `json:"type"`
	Value  float32     `json:"value"`
	Vendor string      `json:"vendor",omitempty`
	Ext    Ext         `json:"ext".omitempty`
}

func (m *Metric) Validate() error {
	if m.Type == "" {
		return ErrorMissingType
	}
	if m.Value < 0 && m.Value > 1.0 {
		return ErrorWrongValue
	}
	return nil
}