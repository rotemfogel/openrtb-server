package model

// 5.1 Content Categories
const (
	IAB1 string = "IAB1"
	IAB1_1 = "IAB1-1"
	IAB1_2 = "IAB1-2"
	IAB1_3 = "IAB1-3"
	IAB1_4 = "IAB1-4"
	IAB1_5 = "IAB1-5"
	IAB1_6 = "IAB1-6"
	IAB1_7 = "IAB1-7"
	IAB2 = "IAB2"
	IAB2_1 = "IAB2-1"
	IAB2_2 = "IAB2-2"
	IAB2_3 = "IAB2-3"
	IAB2_4 = "IAB2-4"
	IAB2_5 = "IAB2-5"
	IAB2_6 = "IAB2-6"
	IAB2_7 = "IAB2-7"
	IAB2_8 = "IAB2-8"
	IAB2_9 = "IAB2-9"
	IAB2_10 = "IAB2-10"
	IAB2_11 = "IAB2-11"
	IAB2_12 = "IAB2-12"
	IAB2_13 = "IAB2-13"
	IAB2_14 = "IAB2-14"
	IAB2_15 = "IAB2-15"
	IAB2_16 = "IAB2-16"
	IAB2_17 = "IAB2-17"
	IAB2_18 = "IAB2-18"
	IAB2_19 = "IAB2-19"
	IAB2_20 = "IAB2-20"
	IAB2_21 = "IAB2-21"
	IAB2_22 = "IAB2-22"
	IAB2_23 = "IAB2-23"
	IAB3 = "IAB3"
	IAB3_1 = "IAB3-1"
	IAB3_2 = "IAB3-2"
	IAB3_3 = "IAB3-3"
	IAB3_4 = "IAB3-4"
	IAB3_5 = "IAB3-5"
	IAB3_6 = "IAB3-6"
	IAB3_7 = "IAB3-7"
	IAB3_8 = "IAB3-8"
	IAB3_9 = "IAB3-9"
	IAB3_10 = "IAB3-10"
	IAB3_11 = "IAB3-11"
	IAB3_12 = "IAB3-12"
	IAB4 = "IAB4"
	IAB4_1 = "IAB4-1"
	IAB4_2 = "IAB4-2"
	IAB4_3 = "IAB4-3"
	IAB4_4 = "IAB4-4"
	IAB4_5 = "IAB4-5"
	IAB4_6 = "IAB4-6"
	IAB4_7 = "IAB4-7"
	IAB4_8 = "IAB4-8"
	IAB4_9 = "IAB4-9"
	IAB4_10 = "IAB4-10"
	IAB4_11 = "IAB4-11"
	IAB5 = "IAB5"
	IAB5_1 = "IAB5-1"
	IAB5_2 = "IAB5-2"
	IAB5_3 = "IAB5-3"
	IAB5_4 = "IAB5-4"
	IAB5_5 = "IAB5-5"
	IAB5_6 = "IAB5-6"
	IAB5_7 = "IAB5-7"
	IAB5_8 = "IAB5-8"
	IAB5_9 = "IAB5-9"
	IAB5_10 = "IAB5-10"
	IAB5_11 = "IAB5-11"
	IAB5_12 = "IAB5-12"
	IAB5_13 = "IAB5-13"
	IAB5_14 = "IAB5-14"
	IAB5_15 = "IAB5-15"
	IAB6 = "IAB6"
	IAB6_1 = "IAB6-1"
	IAB6_2 = "IAB6-2"
	IAB6_3 = "IAB6-3"
	IAB6_4 = "IAB6-4"
	IAB6_5 = "IAB6-5"
	IAB6_6 = "IAB6-6"
	IAB6_7 = "IAB6-7"
	IAB6_8 = "IAB6-8"
	IAB6_9 = "IAB6-9"
	IAB7 = "IAB7"
	IAB7_1 = "IAB7-1"
	IAB7_2 = "IAB7-2"
	IAB7_3 = "IAB7-3"
	IAB7_4 = "IAB7-4"
	IAB7_5 = "IAB7-5"
	IAB7_6 = "IAB7-6"
	IAB7_7 = "IAB7-7"
	IAB7_8 = "IAB7-8"
	IAB7_9 = "IAB7-9"
	IAB7_10 = "IAB7-10"
	IAB7_11 = "IAB7-11"
	IAB7_12 = "IAB7-12"
	IAB7_13 = "IAB7-13"
	IAB7_14 = "IAB7-14"
	IAB7_15 = "IAB7-15"
	IAB7_16 = "IAB7-16"
	IAB7_17 = "IAB7-17"
	IAB7_18 = "IAB7-18"
	IAB7_19 = "IAB7-19"
	IAB7_20 = "IAB7-20"
	IAB7_21 = "IAB7-21"
	IAB7_22 = "IAB7-22"
	IAB7_23 = "IAB7-23"
	IAB7_24 = "IAB7-24"
	IAB7_25 = "IAB7-25"
	IAB7_26 = "IAB7-26"
	IAB7_27 = "IAB7-27"
	IAB7_28 = "IAB7-28"
	IAB7_29 = "IAB7-29"
	IAB7_30 = "IAB7-30"
	IAB7_31 = "IAB7-31"
	IAB7_32 = "IAB7-32"
	IAB7_33 = "IAB7-33"
	IAB7_34 = "IAB7-34"
	IAB7_35 = "IAB7-35"
	IAB7_36 = "IAB7-36"
	IAB7_37 = "IAB7-37"
	IAB7_38 = "IAB7-38"
	IAB7_39 = "IAB7-39"
	IAB7_40 = "IAB7-40"
	IAB7_41 = "IAB7-41"
	IAB7_42 = "IAB7-42"
	IAB7_43 = "IAB7-43"
	IAB7_44 = "IAB7-44"
	IAB7_45 = "IAB7-45"
	IAB8 = "IAB8"
	IAB8_1 = "IAB8-1"
	IAB8_2 = "IAB8-2"
	IAB8_3 = "IAB8-3"
	IAB8_4 = "IAB8-4"
	IAB8_5 = "IAB8-5"
	IAB8_6 = "IAB8-6"
	IAB8_7 = "IAB8-7"
	IAB8_8 = "IAB8-8"
	IAB8_9 = "IAB8-9"
	IAB8_10 = "IAB8-10"
	IAB8_11 = "IAB8-11"
	IAB8_12 = "IAB8-12"
	IAB8_13 = "IAB8-13"
	IAB8_14 = "IAB8-14"
	IAB8_15 = "IAB8-15"
	IAB8_16 = "IAB8-16"
	IAB8_17 = "IAB8-17"
	IAB8_18 = "IAB8-18"
	IAB9 = "IAB9"
	IAB9_1 = "IAB9-1"
	IAB9_2 = "IAB9-2"
	IAB9_3 = "IAB9-3"
	IAB9_4 = "IAB9-4"
	IAB9_5 = "IAB9-5"
	IAB9_6 = "IAB9-6"
	IAB9_7 = "IAB9-7"
	IAB9_8 = "IAB9-8"
	IAB9_9 = "IAB9-9"
	IAB9_10 = "IAB9-10"
	IAB9_11 = "IAB9-11"
	IAB9_12 = "IAB9-12"
	IAB9_13 = "IAB9-13"
	IAB9_14 = "IAB9-14"
	IAB9_15 = "IAB9-15"
	IAB9_16 = "IAB9-16"
	IAB9_17 = "IAB9-17"
	IAB9_18 = "IAB9-18"
	IAB9_19 = "IAB9-19"
	IAB9_20 = "IAB9-20"
	IAB9_21 = "IAB9-21"
	IAB9_22 = "IAB9-22"
	IAB9_23 = "IAB9-23"
	IAB9_24 = "IAB9-24"
	IAB9_25 = "IAB9-25"
	IAB9_26 = "IAB9-26"
	IAB9_27 = "IAB9-27"
	IAB9_28 = "IAB9-28"
	IAB9_29 = "IAB9-29"
	IAB9_30 = "IAB9-30"
	IAB9_31 = "IAB9-31"
	IAB10 = "IAB10"
	IAB10_1 = "IAB10-1"
	IAB10_2 = "IAB10-2"
	IAB10_3 = "IAB10-3"
	IAB10_4 = "IAB10-4"
	IAB10_5 = "IAB10-5"
	IAB10_6 = "IAB10-6"
	IAB10_7 = "IAB10-7"
	IAB10_8 = "IAB10-8"
	IAB10_9 = "IAB10-9"
	IAB11 = "IAB11"
	IAB11_1 = "IAB11-1"
	IAB11_2 = "IAB11-2"
	IAB11_3 = "IAB11-3"
	IAB11_4 = "IAB11-4"
	IAB11_5 = "IAB11-5"
	IAB12 = "IAB12"
	IAB12_1 = "IAB12-1"
	IAB12_2 = "IAB12-2"
	IAB12_3 = "IAB12-3"
	IAB13 = "IAB13"
	IAB13_1 = "IAB13-1"
	IAB13_2 = "IAB13-2"
	IAB13_3 = "IAB13-3"
	IAB13_4 = "IAB13-4"
	IAB13_5 = "IAB13-5"
	IAB13_6 = "IAB13-6"
	IAB13_7 = "IAB13-7"
	IAB13_8 = "IAB13-8"
	IAB13_9 = "IAB13-9"
	IAB13_10 = "IAB13-10"
	IAB13_11 = "IAB13-11"
	IAB13_12 = "IAB13-12"
	IAB14 = "IAB14"
	IAB14_1 = "IAB14-1"
	IAB14_2 = "IAB14-2"
	IAB14_3 = "IAB14-3"
	IAB14_4 = "IAB14-4"
	IAB14_5 = "IAB14-5"
	IAB14_6 = "IAB14-6"
	IAB14_7 = "IAB14-7"
	IAB14_8 = "IAB14-8"
	IAB15 = "IAB15"
	IAB15_1 = "IAB15-1"
	IAB15_2 = "IAB15-2"
	IAB15_3 = "IAB15-3"
	IAB15_4 = "IAB15-4"
	IAB15_5 = "IAB15-5"
	IAB15_6 = "IAB15-6"
	IAB15_7 = "IAB15-7"
	IAB15_8 = "IAB15-8"
	IAB15_9 = "IAB15-9"
	IAB15_10 = "IAB15-10"
	IAB16 = "IAB16"
	IAB16_1 = "IAB16-1"
	IAB16_2 = "IAB16-2"
	IAB16_3 = "IAB16-3"
	IAB16_4 = "IAB16-4"
	IAB16_5 = "IAB16-5"
	IAB16_6 = "IAB16-6"
	IAB16_7 = "IAB16-7"
	IAB17 = "IAB17"
	IAB17_1 = "IAB17-1"
	IAB17_2 = "IAB17-2"
	IAB17_3 = "IAB17-3"
	IAB17_4 = "IAB17-4"
	IAB17_5 = "IAB17-5"
	IAB17_6 = "IAB17-6"
	IAB17_7 = "IAB17-7"
	IAB17_8 = "IAB17-8"
	IAB17_9 = "IAB17-9"
	IAB17_10 = "IAB17-10"
	IAB17_11 = "IAB17-11"
	IAB17_12 = "IAB17-12"
	IAB17_13 = "IAB17-13"
	IAB17_14 = "IAB17-14"
	IAB17_15 = "IAB17-15"
	IAB17_16 = "IAB17-16"
	IAB17_17 = "IAB17-17"
	IAB17_18 = "IAB17-18"
	IAB17_19 = "IAB17-19"
	IAB17_20 = "IAB17-20"
	IAB17_21 = "IAB17-21"
	IAB17_22 = "IAB17-22"
	IAB17_23 = "IAB17-23"
	IAB17_24 = "IAB17-24"
	IAB17_25 = "IAB17-25"
	IAB17_26 = "IAB17-26"
	IAB17_27 = "IAB17-27"
	IAB17_28 = "IAB17-28"
	IAB17_29 = "IAB17-29"
	IAB17_30 = "IAB17-30"
	IAB17_31 = "IAB17-31"
	IAB17_32 = "IAB17-32"
	IAB17_33 = "IAB17-33"
	IAB17_34 = "IAB17-34"
	IAB17_35 = "IAB17-35"
	IAB17_36 = "IAB17-36"
	IAB17_37 = "IAB17-37"
	IAB17_38 = "IAB17-38"
	IAB17_39 = "IAB17-39"
	IAB17_40 = "IAB17-40"
	IAB17_41 = "IAB17-41"
	IAB17_42 = "IAB17-42"
	IAB17_43 = "IAB17-43"
	IAB17_44 = "IAB17-44"
	IAB18 = "IAB18"
	IAB18_1 = "IAB18-1"
	IAB18_2 = "IAB18-2"
	IAB18_3 = "IAB18-3"
	IAB18_4 = "IAB18-4"
	IAB18_5 = "IAB18-5"
	IAB18_6 = "IAB18-6"
	IAB19 = "IAB19"
	IAB19_1 = "IAB19-1"
	IAB19_2 = "IAB19-2"
	IAB19_3 = "IAB19-3"
	IAB19_4 = "IAB19-4"
	IAB19_5 = "IAB19-5"
	IAB19_6 = "IAB19-6"
	IAB19_7 = "IAB19-7"
	IAB19_8 = "IAB19-8"
	IAB19_9 = "IAB19-9"
	IAB19_10 = "IAB19-10"
	IAB19_11 = "IAB19-11"
	IAB19_12 = "IAB19-12"
	IAB19_13 = "IAB19-13"
	IAB19_14 = "IAB19-14"
	IAB19_15 = "IAB19-15"
	IAB19_16 = "IAB19-16"
	IAB19_17 = "IAB19-17"
	IAB19_18 = "IAB19-18"
	IAB19_19 = "IAB19-19"
	IAB19_20 = "IAB19-20"
	IAB19_21 = "IAB19-21"
	IAB19_22 = "IAB19-22"
	IAB19_23 = "IAB19-23"
	IAB19_24 = "IAB19-24"
	IAB19_25 = "IAB19-25"
	IAB19_26 = "IAB19-26"
	IAB19_27 = "IAB19-27"
	IAB19_28 = "IAB19-28"
	IAB19_29 = "IAB19-29"
	IAB19_30 = "IAB19-30"
	IAB19_31 = "IAB19-31"
	IAB19_32 = "IAB19-32"
	IAB19_33 = "IAB19-33"
	IAB19_34 = "IAB19-34"
	IAB19_35 = "IAB19-35"
	IAB19_36 = "IAB19-36"
	IAB20 = "IAB20"
	IAB20_1 = "IAB20-1"
	IAB20_2 = "IAB20-2"
	IAB20_3 = "IAB20-3"
	IAB20_4 = "IAB20-4"
	IAB20_5 = "IAB20-5"
	IAB20_6 = "IAB20-6"
	IAB20_7 = "IAB20-7"
	IAB20_8 = "IAB20-8"
	IAB20_9 = "IAB20-9"
	IAB20_10 = "IAB20-10"
	IAB20_11 = "IAB20-11"
	IAB20_12 = "IAB20-12"
	IAB20_13 = "IAB20-13"
	IAB20_14 = "IAB20-14"
	IAB20_15 = "IAB20-15"
	IAB20_16 = "IAB20-16"
	IAB20_17 = "IAB20-17"
	IAB20_18 = "IAB20-18"
	IAB20_19 = "IAB20-19"
	IAB20_20 = "IAB20-20"
	IAB20_21 = "IAB20-21"
	IAB20_22 = "IAB20-22"
	IAB20_23 = "IAB20-23"
	IAB20_24 = "IAB20-24"
	IAB20_25 = "IAB20-25"
	IAB20_26 = "IAB20-26"
	IAB20_27 = "IAB20-27"
	IAB21 = "IAB21"
	IAB21_1 = "IAB21-1"
	IAB21_2 = "IAB21-2"
	IAB21_3 = "IAB21-3"
	IAB22 = "IAB22"
	IAB22_1 = "IAB22-1"
	IAB22_2 = "IAB22-2"
	IAB22_3 = "IAB22-3"
	IAB22_4 = "IAB22-4"
	IAB23 = "IAB23"
	IAB23_2 = "IAB23-2"
	IAB23_3 = "IAB23-3"
	IAB23_4 = "IAB23-4"
	IAB23_5 = "IAB23-5"
	IAB23_6 = "IAB23-6"
	IAB23_7 = "IAB23-7"
	IAB23_8 = "IAB23-8"
	IAB23_9 = "IAB23-9"
	IAB23_10 = "IAB23-10"
	IAB24 = "IAB24"
	IAB25 = "IAB25"
	IAB25_1 = "IAB25-1"
	IAB25_2 = "IAB25-2"
	IAB25_3 = "IAB25-3"
	IAB25_4 = "IAB25-4"
	IAB25_5 = "IAB25-5"
	IAB25_6 = "IAB25-6"
	IAB25_7 = "IAB25-7"
	IAB26 = "IAB26"
	IAB26_1 = "IAB26-1"
	IAB26_2 = "IAB26-2"
	IAB26_3 = "IAB26-3"
	IAB26_4 = "IAB26-4"
)
// 5.2 Banner Ad Types
const (
	BannerTypeXHTMLText int = iota + 1
	BannerTypeXHTML
	BannerTypeJS
	BannerTypeFrame
)

// 5.3 Creative Attributes
const (
	AudioAd = iota + 1
	AudioAdUserInitiated
	ExpandableAutomatic
	ExpandableUserInitiated
	ExpandableUserInitiatedRollover
	InBannerVideoAdAutoPlay
	InBannerVideoAdUserInitiated
	Pop
	ProvocativeOrSuggestiveImagery
	ShakyFlashingFlickeringExtremeAnimationSmileys
	Surveys
	TextOnly
	UserInteractive
	WindowsDialogOrAlertStyle
	HasAudioOnOffButton
	AdProvidesSkipButton
	AdobeFlash
)

// 5.4 Ad Position
const (
	AdPositionUnknown int = iota
	AdPositionAboveTheFold
	AdPositionDEPRECATED
	AdPositionBelowTheFold
	AdPositionHeader
	AdPositionFooter
	AdPositionSidebar
	AdPositionFullScreen
)

// 5.5 Expandable Direction
const (
	ExpandableLeft int = iota + 1
	ExpandableRight
	ExpandableUp
	ExpandableDown
	ExpandableFullScreen
)

// 5.6 API Frameworks
const (
	APIFrameworkVPAID1 int = iota + 1
	APIFrameworkVPAID2
	APIFrameworkMRAID1
	APIFrameworkORMMA
	APIFrameworkMRAID2
)

// 5.7 Video Linearity
const (
	VideoLinearityLinear  int = iota + 1
	VideoLinearityNonLinear
)

// 5.8 Video and Audio Bid Response Protocols
const (
	VAST1 int = iota + 1
	VAST2
	VAST3
	VAST1Wrapper
	VAST2Wrapper
	VAST3Wrapper
	VAST4
	VAST4Wrapper
	DAAST1
	DAAST1Wrapper
)

// 5.9 Video Playback Methods
const (
	VideoPlaybackMethodAutoSoundOn int = iota + 1
	VideoPlaybackMethodAutoSoundOff
	VideoPlaybackMethodClickToPlay
	VideoPlaybackMethodMouseOver
)

// 5.10 Video Start Delay
const (
	VideoStartDelayPreRoll         = 0
	VideoStartDelayGenericMidRoll  = -1
	VideoStartDelayGenericPostRoll = -2
)

// 5.11 Video Quality
const (
	VideoQualityUnknown int = iota
	VideoQualityProfessional
	VideoQualityProsumer
	VideoQualityUGC
)

// 5.12 VAST Companion Types
const (
	VASTCompanionStatic int = iota + 1
	VASTCompanionHTML
	VASTCompanionIFrame
)

// 5.13 Content Delivery Methods
const (
	ContentDeliveryStreaming int = iota + 1
	ContentDeliveryProgressive
	ContentDeliveryDownload
)

// 5.14 Content Context
const (
	ContextVideo int = iota + 1
	ContextGame
	ContextMusic
	ContextApplication
	ContextText
	ContextOther
	ContextUnknown
)

// 5.15 QAG Media Ratings
const (
	QAGAll int = iota + 1
	QAGOver12
	QAGMature
)

// 5.16 Location Type
const (
	LocationTypeGPS int = iota + 1
	LocationTypeIP
	LocationTypeUser
)

// 5.17 Device Type
const (
	DeviceTypeUnknown int = iota
	DeviceTypeMobile
	DeviceTypePC
	DeviceTypeTV
	DeviceTypePhone
	DeviceTypeTablet
	DeviceTypeConnected
	DeviceTypeSetTopBox
)

// 5.18 Connection Type
const (
	ConnTypeUnknown int = iota
	ConnTypeEthernet
	ConnTypeWIFI
	ConnTypeCell
	ConnTypeCell2G
	ConnTypeCell3G
	ConnTypeCell4G
)

// 5.19 No-Bid Reason Codes
const (
	NBRUnknownError int = iota
	NBRTechnicalError
	NBRInvalidRequest
	NBRKnownSpider
	NBRSuspectedNonHuman
	NBRProxyIP
	NBRUnsupportedDevice
	NBRBlockedSite
	NBRUnmatchedUser
)

type Common struct {
	Id     string    `json:"id,omitempty"`
	Name   string    `json:"name,omitempty"`
	Cat    []string  `json:"cat,omitempty"`
	Domain string    `json:"domain,omitempty"`
	Ext    Ext       `json:"ext,omitempty"`
}

type Publisher Common

type Producer Common

// Note that the Geo Object may appear in one or both the Device Object and the User Object.
// This is intentional, since the information may be derived from either a device-oriented source
// (such as IP geo lookup), or by user registration information (for example provided to a publisher
// through a user registration).
type Geo struct {
	Lat           float64   `json:"lat,omitempty"`
	Lon           float64   `json:"lon,omitempty"`
	Type          int       `json:"type,omitempty"`
	Accuracy      int       `json:"accuracy,omitempty"`
	Lastfix       int       `json:"lastfix,omitempty"`
	Ipservice     int       `json:"ipservice,omitempty"`
	Country       string    `json:"country,omitempty"`
	Region        string    `json:"region,omitempty"`
	Regionfips104 string    `json:"regionFIPS104,omitempty"`
	Metro         string    `json:"metro,omitempty"`
	City          string    `json:"city,omitempty"`
	Zip           string    `json:"zip,omitempty"`
	Utcoffset     int       `json:"utcoffset,omitempty"`
	Ext           Ext       `json:"ext,omitempty"`
}

type User struct {
	Id         string    `json:"id,omitempty"`
	Buyerid    string    `json:"buyerid,omitempty"`
	Buyeruis   string    `json:"buyeruid,omitempty"`
	Yob        int       `json:"yob,omitempty"`
	Gender     string    `json:"gender,omitempty"`
	Keywords   string    `json:"keywords,omitempty"`
	Customdata string    `json:"customdata,omitempty"`
	Geo        *Geo      `json:"geo,omitempty"`
	Data       []Data    `json:"data,omitempty"`
	Ext        Ext       `json:"ext,omitempty"`
}

type Data struct {
	Id      string    `json:"id,omitempty"`
	Name    string    `json:"name,omitempty"`
	Segment []Segment `json:"segment,omitempty"`
	Ext     Ext       `json:"ext,omitempty"`
}

type Segment struct {
	Id    string    `json:"id,omitempty"`
	Name  string    `json:"name,omitempty"`
	Value string    `json:"value,omitempty"`
	Ext   Ext       `json:"ext,omitempty"`
}

type Format struct {
	W      int       `json:"w",omitempty`
	H      int       `json:"h",omitempty`
	Wratio int       `json:"wratio",omitempty`
	Hratio int       `json:"hratio",omitempty`
	Wmin   int       `json:"wmin",omitempty`
	Ext    Ext       `json:"ext",omitempty`
}
