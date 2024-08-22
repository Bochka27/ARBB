package bybit

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
}
