package models

type FundingRate struct {
	Success bool `json:"success"`
	Code    int  `json:"code"`
	Data    struct {
		Symbol         string  `json:"symbol"`
		FundingRate    float64 `json:"fundingRate"`
		MaxFundingRate float64 `json:"maxFundingRate"`
		MinFundingRate float64 `json:"minFundingRate"`
		CollectCycle   int     `json:"collectCycle"`
		NextSettleTime int64   `json:"nextSettleTime"`
		Timestamp      int64   `json:"timestamp"`
	} `json:"data"`
}
