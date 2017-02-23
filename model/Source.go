package model

type Source struct {
	Fd     int    `json:"fd",omitempty`
	Tid    string `json:"tid",omitempty`
	Pchain string `json:"pchain",omitrmpty`
	Ext    Ext    `json:"ext,omitempty"`
}