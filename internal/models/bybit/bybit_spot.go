package bybit

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
}
