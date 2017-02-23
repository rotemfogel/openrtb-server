package model

type Banner struct {
	Format   []Format  `json:"format,omitempty"`
	W        int       `json:"w,omitempty"`
	H        int       `json:"h,omitempty"`
	WMax     int       `json:"wmax,omitempty"`
	HMax     int       `json:"hmax,omitempty"`
	WMin     int       `json:"wmin,omitempty"`
	HMin     int       `json:"hmin,omitempty"`
	Btype    []int     `json:"btype,omitempty"`
	Battr    []int     `json:"battr,omitempty"`
	Pos      int       `json:"pos,omitempty"`
	Mimes    []string  `json:"mimes,omitempty"`
	Topframe int       `json:"topframe,omitempty"`
	Expdir   []int     `json:"expdir,omitempty"`
	Api      []int     `json:"api,omitempty"`
	Id       string    `json:"id,omitempty"`
	Vcm      int       `json:"vcm,omitempty"`
	Ext      Ext       `json:"ext,omitempty"`
}

func NewBanner() *Banner {
	return &Banner{Topframe:0}
}
