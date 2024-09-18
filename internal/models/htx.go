package models

// Futures data

type HtxFutures struct {
	Status string `json:"status"`
	Ticks  []struct {
		BusinessType string    `json:"business_type"`
		ContractCode string    `json:"contract_code"`
		Ask          []float64 `json:"ask"`
		Bid          []float64 `json:"bid"`
		Mrid         int       `json:"mrid"`
		Ts           int64     `json:"ts"`
	} `json:"ticks"`
	Ts int64 `json:"ts"`
}

type HtxFR struct {
	Status string `json:"status"`
	Data   []struct {
		FundingRate     string `json:"funding_rate"`
		ContractCode    string `json:"contract_code"`
		Symbol          string `json:"symbol"`
		FeeAsset        string `json:"fee_asset"`
		FundingTime     string `json:"funding_time"`
		EstimatedRate   string `json:"estimated_rate"`
		NextFundingTime string `json:"next_funding_time"`
	} `json:"data"`
	Ts int64 `json:"ts"`
}

//Spot data

type HtxSpotPrice struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   []struct {
		Symbol  string  `json:"symbol"`
		Open    float64 `json:"open"`
		High    float64 `json:"high"`
		Low     float64 `json:"low"`
		Close   float64 `json:"close"`
		Amount  float64 `json:"amount"`
		Vol     float64 `json:"vol"`
		Count   int     `json:"count"`
		Bid     float64 `json:"bid"`
		BidSize float64 `json:"bidSize"`
		Ask     float64 `json:"ask"`
		AskSize float64 `json:"askSize"`
	} `json:"data"`
}
