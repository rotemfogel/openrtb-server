package model

import "errors"

// Validation errors
var (
	ErrorMissingId = errors.New("openrtb::model Imp id missing")
	ErrorMutipleObjects = errors.New("openrtb::model Imp must have only one of banner/video/audio/native")
)

type Imp struct {
	Id                string    `json:"id"`
	Metric            []Metric  `json:"metric",omitempty`
	Banner            *Banner   `json:"banner",omitempty`
	Video             *Video    `json:"video",omitempty`
	Audio             *Audio    `json:"audio",omitempty`
	Native            *Native   `json:"native",omitempty`
	Pmp               *Pmp      `json:"pmp",omitempty`
	Displaymanager    string    `json:"displaymanager",omitempty`
	Displaymanagerver string    `json:"displaymanagerver",omitempty`
	Instl             int       `json:"instl"`
	Tagid             string    `json:"tagid",omitempty`
	Bidfloor          float32   `json:"bidfloor"`
	Bidfloorcur       string    `json:"bidfloorcur"`
	Clickbrowser      int       `json:"clickbrowser",omitempty`
	Secure            int       `json:"secure",omitempty`
	Iframebuster      []string  `json:"iframebuster",omitempty`
	Exp               int       `json:"exp",omitempty`
	Ext               Ext       `json:"ext",omitempty`
}

func NewImp() *Imp {
	return &Imp{Instl:0, Bidfloor:0.0, Bidfloorcur:"USD"}
}

func (imp *Imp) Validate() error {
	if imp.Id == "" {
		return ErrorMissingId
	}
	if imp.Banner != nil && imp.Video != nil && imp.Audio != nil && imp.Native != nil {
		return ErrorMutipleObjects
	}
	// TODO: validate banner/video/audio/native
	if imp.Video != nil {
		if err := imp.Video.Validate(); err != nil {
			return err
		}
	}
	for _, metric := range imp.Metric {
		if err := (&metric).Validate(); err != nil {
			return err
		}
	}
	return nil
}