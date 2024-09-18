package models

// Funding Rate

type Fund struct {
	Code string     `json:"code"`
	Data []FundData `json:"data"`
	Msg  string     `json:"msg"`
}

type FundData struct {
	FundingRate     string `json:"fundingRate"`
	FundingTime     string `json:"fundingTime"`
	InstID          string `json:"instId"`
	InstType        string `json:"instType"`
	Method          string `json:"method"`
	MaxFundingRate  string `json:"maxFundingRate"`
	MinFundingRate  string `json:"minFundingRate"`
	NextFundingRate string `json:"nextFundingRate"`
	NextFundingTime string `json:"nextFundingTime"`
	Premium         string `json:"premium"`
	SettFundingRate string `json:"settFundingRate"`
	SettState       string `json:"settState"`
	Ts              string `json:"ts"`
}

// List all tickers(swap and spot)

type Symbol struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		InstType  string `json:"instType"`
		InstID    string `json:"instId"`
		Last      string `json:"last"`
		LastSz    string `json:"lastSz"`
		AskPx     string `json:"askPx"`
		AskSz     string `json:"askSz"`
		BidPx     string `json:"bidPx"`
		BidSz     string `json:"bidSz"`
		Open24H   string `json:"open24h"`
		High24H   string `json:"high24h"`
		Low24H    string `json:"low24h"`
		VolCcy24H string `json:"volCcy24h"`
		Vol24H    string `json:"vol24h"`
		SodUtc0   string `json:"sodUtc0"`
		SodUtc8   string `json:"sodUtc8"`
		Ts        string `json:"ts"`
	} `json:"data"`
}
