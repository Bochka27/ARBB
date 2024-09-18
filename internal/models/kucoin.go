package models

// Futures

type KucoinF struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	Retry   bool   `json:"retry"`
	Data    []struct {
		Sequence     int64  `json:"sequence"`
		Symbol       string `json:"symbol"`
		Side         string `json:"side"`
		Size         int    `json:"size"`
		TradeID      string `json:"tradeId"`
		Price        string `json:"price"`
		BestBidPrice string `json:"bestBidPrice"`
		BestBidSize  int    `json:"bestBidSize"`
		BestAskPrice string `json:"bestAskPrice"`
		BestAskSize  int    `json:"bestAskSize"`
		Ts           int64  `json:"ts"`
	} `json:"data"`
}

type KucoinFR struct {
	Code string  `json:"code"`
	Data []KData `json:"data"`
}

type KData struct {
	Symbol                  string   `json:"symbol"`
	RootSymbol              string   `json:"rootSymbol"`
	Type                    string   `json:"type"`
	FirstOpenDate           any      `json:"firstOpenDate"`
	ExpireDate              any      `json:"expireDate"`
	SettleDate              any      `json:"settleDate"`
	BaseCurrency            string   `json:"baseCurrency"`
	QuoteCurrency           string   `json:"quoteCurrency"`
	SettleCurrency          string   `json:"settleCurrency"`
	MaxOrderQty             float64  `json:"maxOrderQty"`
	MaxPrice                float64  `json:"maxPrice"`
	LotSize                 int64    `json:"lotSize"`
	TickSize                float64  `json:"tickSize"`
	IndexPriceTickSize      float64  `json:"indexPriceTickSize"`
	Multiplier              float64  `json:"multiplier"`
	InitialMargin           float64  `json:"initialMargin"`
	MaintainMargin          float64  `json:"maintainMargin"`
	MaxRiskLimit            float64  `json:"maxRiskLimit"`
	MinRiskLimit            float64  `json:"minRiskLimit"`
	RiskStep                float64  `json:"riskStep"`
	MakerFeeRate            float64  `json:"makerFeeRate"`
	TakerFeeRate            float64  `json:"takerFeeRate"`
	TakerFixFee             float64  `json:"takerFixFee"`
	MakerFixFee             float64  `json:"makerFixFee"`
	SettlementFee           any      `json:"settlementFee"`
	IsDeleverage            bool     `json:"isDeleverage"`
	IsQuanto                bool     `json:"isQuanto"`
	IsInverse               bool     `json:"isInverse"`
	MarkMethod              string   `json:"markMethod"`
	FairMethod              string   `json:"fairMethod"`
	FundingBaseSymbol       string   `json:"fundingBaseSymbol"`
	FundingQuoteSymbol      string   `json:"fundingQuoteSymbol"`
	FundingRateSymbol       string   `json:"fundingRateSymbol"`
	IndexSymbol             string   `json:"indexSymbol"`
	SettlementSymbol        string   `json:"settlementSymbol"`
	Status                  string   `json:"status"`
	FundingFeeRate          float64  `json:"fundingFeeRate"`
	PredictedFundingFeeRate float64  `json:"predictedFundingFeeRate"`
	FundingRateGranularity  float64  `json:"fundingRateGranularity"`
	OpenInterest            string   `json:"openInterest"`
	TurnoverOf24H           float64  `json:"turnoverOf24h"`
	VolumeOf24H             float64  `json:"volumeOf24h"`
	MarkPrice               float64  `json:"markPrice"`
	IndexPrice              float64  `json:"indexPrice"`
	LastTradePrice          float64  `json:"lastTradePrice"`
	NextFundingRateTime     int64    `json:"nextFundingRateTime"`
	MaxLeverage             int64    `json:"maxLeverage"`
	SourceExchanges         []string `json:"sourceExchanges"`
	PremiumsSymbol1M        string   `json:"premiumsSymbol1M"`
	PremiumsSymbol8H        string   `json:"premiumsSymbol8H"`
	FundingBaseSymbol1M     string   `json:"fundingBaseSymbol1M"`
	FundingQuoteSymbol1M    string   `json:"fundingQuoteSymbol1M"`
	LowPrice                float64  `json:"lowPrice"`
	HighPrice               float64  `json:"highPrice"`
	PriceChgPct             float64  `json:"priceChgPct"`
	PriceChg                float64  `json:"priceChg"`
	K                       float64  `json:"k"`
	M                       float64  `json:"m"`
	F                       float64  `json:"f"`
	MmrLimit                float64  `json:"mmrLimit"`
	MmrLevConstant          float64  `json:"mmrLevConstant"`
}

// Spot

type KucoinS struct {
	Code string `json:"code"`
	Data *Data  `json:"data"`
}

type Data struct {
	Time   int64     `json:"time"`
	Ticker []*Ticker `json:"ticker"`
}

type Ticker struct {
	Symbol           string `json:"symbol"`
	SymbolName       string `json:"symbolName"`
	Buy              string `json:"buy"`
	BestBidSize      string `json:"bestBidSize"`
	Sell             string `json:"sell"`
	BestAskSize      string `json:"bestAskSize"`
	ChangeRate       string `json:"changeRate"`
	ChangePrice      string `json:"changePrice"`
	High             string `json:"high"`
	Low              string `json:"low"`
	Vol              string `json:"vol"`
	VolValue         string `json:"volValue"`
	Last             string `json:"last"`
	AveragePrice     string `json:"averagePrice"`
	TakerFeeRate     string `json:"takerFeeRate"`
	MakerFeeRate     string `json:"makerFeeRate"`
	TakerCoefficient string `json:"takerCoefficient"`
	MakerCoefficient string `json:"makerCoefficient"`
}
