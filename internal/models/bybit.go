package models

// Futures data

type MarketTickersResponseF struct {
	RetCode    int            `json:"retCode"`
	RetMsg     string         `json:"retMsg"`
	Result     MarketTickersF `json:"result"`
	RetExtInfo struct{}       `json:"retExtInfo"`
	Time       int            `json:"time"`
}
type MarketTickersF struct {
	Category string         `json:"category"`
	List     []*TickerInfoF `json:"list"`
}
type TickerInfoF struct {
	Symbol          string `json:"symbol"`
	LastPrice       string `json:"lastPrice"`
	Volume24h       string `json:"volume24h"`
	FundingRate     string `json:"fundingRate"`
	NextFundingTime string `json:"nextFundingTime"`
	Ask1Size        string `json:"ask1Size"`
	Bid1Price       string `json:"bid1Price"`
	Ask1Price       string `json:"ask1Price"`
	Bid1Size        string `json:"bid1Size"`
	Basis           string `json:"basis"`
}

// Spot data

type MarketTickersResponseS struct {
	RetCode    int            `json:"retCode"`
	RetMsg     string         `json:"retMsg"`
	Result     MarketTickersS `json:"result"`
	RetExtInfo struct{}       `json:"retExtInfo"`
	Time       int            `json:"time"`
}
type MarketTickersS struct {
	Category string         `json:"category"`
	List     []*TickerInfoS `json:"list"`
}
type TickerInfoS struct {
	Symbol    string `json:"symbol"`
	LastPrice string `json:"lastPrice"`
	Volume24h string `json:"volume24h"`
	Ask1Size  string `json:"ask1Size"`
	Bid1Price string `json:"bid1Price"`
	Ask1Price string `json:"ask1Price"`
	Bid1Size  string `json:"bid1Size"`
	Basis     string `json:"basis"`
}
