package model

import "encoding/json"

type Ext struct {
	m map[string]interface{}
}

func (e *Ext) Get(key string) (interface{}) {
	if e.m == nil {
		return nil
	}
	return e.m[key]
}

func (e *Ext) Put(key string, val interface{}) {
	if e.m == nil {
		e.m = make(map[string]interface{})
	}
	e.m[key] = val
}

func (e *Ext) UnmarshalJSON(data []byte) error {
	var mm map[string]interface{}
	err := json.Unmarshal(data, mm)
	if (err == nil) {
		*e = Ext{m:mm}
	}
	return err
}

func (e *Ext) Marshal(v interface{}) ([]byte, error) {
	if e.m == nil {
		return nil, nil
	}
	return json.Marshal(e.m)
}