package models

// Futures

type ListF []struct {
	Contract              string `json:"contract"`
	Last                  string `json:"last"`
	Low24H                string `json:"low_24h"`
	High24H               string `json:"high_24h"`
	ChangePercentage      string `json:"change_percentage"`
	TotalSize             string `json:"total_size"`
	Volume24H             string `json:"volume_24h"`
	Volume24HBtc          string `json:"volume_24h_btc"`
	Volume24HUsd          string `json:"volume_24h_usd"`
	Volume24HBase         string `json:"volume_24h_base"`
	Volume24HQuote        string `json:"volume_24h_quote"`
	Volume24HSettle       string `json:"volume_24h_settle"`
	MarkPrice             string `json:"mark_price"`
	FundingRate           string `json:"funding_rate"`
	FundingRateIndicative string `json:"funding_rate_indicative"`
	IndexPrice            string `json:"index_price"`
	HighestBid            string `json:"highest_bid"`
	LowestAsk             string `json:"lowest_ask"`
}

// Spot

type ListS []struct {
	CurrencyPair     string `json:"currency_pair"`
	Last             string `json:"last"`
	LowestAsk        string `json:"lowest_ask"`
	HighestBid       string `json:"highest_bid"`
	ChangePercentage string `json:"change_percentage"`
	ChangeUtc0       string `json:"change_utc0"`
	ChangeUtc8       string `json:"change_utc8"`
	BaseVolume       string `json:"base_volume"`
	QuoteVolume      string `json:"quote_volume"`
	High24H          string `json:"high_24h"`
	Low24H           string `json:"low_24h"`
	EtfNetValue      string `json:"etf_net_value"`
	EtfPreNetValue   string `json:"etf_pre_net_value"`
	EtfPreTimestamp  int    `json:"etf_pre_timestamp"`
	EtfLeverage      string `json:"etf_leverage"`
}
