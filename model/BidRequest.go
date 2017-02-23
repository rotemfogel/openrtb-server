package model

import (
	"errors"
)

// Validation errors
var (
	ErrorMissingId = errors.New("openrtb::model bidrequest ID missing")
	ErrorMissingImps = errors.New("openrtb::model bidrequest has no impressions")
	ErrorMutipleSources = errors.New("openrtb::model bidrequest has multiple sources") // has site and app
)

type BidRequest struct {
	Id      string       `json:"id"` // Unique ID of the bid request
	Imp     []Imp         `json:"imp,omitempty"`
	Site    *Site        `json:"site,omitempty"`
	App     *App         `json:"app,omitempty"`
	Device  *Device      `json:"device,omitempty"`
	User    *User        `json:"user,omitempty"`
	Test    int          `json:"test,omitempty"`
	At      int          `json:"at"`
	Tmax    int          `json:"tmax,omitempty"`
	Wseat   []string     `json:"wseat,omitempty"`
	Bseat   []string     `json:"bseat,omitempty"`
	Allimps int          `json:"allimps,omitempty"`
	Cur     []string     `json:"cur,omitempty"`
	Wlang   []string     `json:"cur,omitempty"`
	Bcat    []string     `json:"bcat,omitempty"`
	Badv    []string     `json:"badv,omitempty"`
	Bapp    []string     `json:"bapp,omitempty"`
	Source  int          `json:"source",omitempty`
	Regs    *Regs        `json:"regs,omitempty"`
	Ext     Ext          `json:"ext,omitempty"`
}

func NewBidRequest() *BidRequest {
	return &BidRequest{Test:0, At:2, Cur:"USD"}
}
func (req *BidRequest) addImp(imp *Imp) *BidRequest {
	req.Imp = append(req.Imp, &imp)
	return req
}

// Validates the request
func (req *BidRequest) Validate() error {
	if req.Id == "" {
		return ErrorMissingId
	} else if len(req.Imp) == 0 {
		return ErrorMissingImps
	} else if req.Site != nil && req.App != nil {
		return ErrorMutipleSources
	}

	for _, imp := range req.Imp {
		if err := (&imp).Validate(); err != nil {
			return err
		}
	}
	return nil
}