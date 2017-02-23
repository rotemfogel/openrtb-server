package model

import (
	"encoding/json"
	"errors"
)

// Validation errors
var (
	ErrInvalidVideoNoMimes = errors.New("openrtb::model video has no mimes")
	ErrInvalidVideoNoProtocols = errors.New("openrtb::model video protocols missing")
)

// The "video" object must be included directly in the impression object if the impression offered
// for auction is an in-stream video ad opportunity.
type Video struct {
	Mimes          []string  `json:"mimes,omitempty"`
	Minduration    int       `json:"minduration,omitempty"`
	Maxduration    int       `json:"maxduration,omitempty"`
	Protocols      []int     `json:"protocols,omitempty"`
	Protocol       int       `json:"protocol,omitempty"`
	W              int       `json:"w,omitempty"`
	H              int       `json:"h,omitempty"`
	Startdelay     int       `json:"startdelay,omitempty"`
	placement      int        `json:"placement,omitempty"`
	Linearity      int       `json:"linearity,omitempty"`
	Skip           int       `json:"skip,omitempty"`
	Skipmin        int       `json:"skipmin,omitempty"`
	Skipafter      int       `json:"skipafter,omitempty"`
	Sequence       int       `json:"sequence,omitempty"`
	BAttr          []int     `json:"battr,omitempty"`
	MaxExtended    int       `json:"maxextended,omitempty"`
	MinBitrate     int       `json:"minbitrate,omitempty"`
	MaxBitrate     int       `json:"maxbitrate,omitempty"`
	BoxingAllowed  *int      `json:"boxingallowed,omitempty"`
	PlaybackMethod []int     `json:"playbackmethod,omitempty"`
	Delivery       []int     `json:"delivery,omitempty"`
	Pos            int       `json:"pos,omitempty"`
	CompanionAd    []Banner  `json:"companionad,omitempty"`
	Api            []int     `json:"api,omitempty"`
	CompanionType  []int     `json:"companiontype,omitempty"`
	Ext            Ext       `json:"ext,omitempty"`
}

func NewVideo() *Video {
	return &Video{Skipmin:0, Skipafter:0, BoxingAllowed:1}
}

type jsonVideo Video

// Validates the object
func (v *Video) Validate() error {
	if len(v.Mimes) == 0 {
		return ErrInvalidVideoNoMimes
	}
	if v.Protocol == 0 {
		if len(v.Protocols) == 0 {
			return ErrInvalidVideoNoProtocols
		}
		v.Protocols[0] = append(v.Protocols, v.Protocol)
	}
	return nil
}

// MarshalJSON custom marshalling with normalization
func (v *Video) MarshalJSON() ([]byte, error) {
	v.normalize()
	return json.Marshal((*jsonVideo)(v))
}

// UnmarshalJSON custom unmarshalling with normalization
func (v *Video) UnmarshalJSON(data []byte) error {
	var h jsonVideo
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}

	*v = (Video)(h)
	v.normalize()
	return nil
}

func (v *Video) normalize() {
	if v.Sequence == 0 {
		v.Sequence = 1
	}
	if v.Linearity == 0 {
		v.Linearity = Linear
	}
}
