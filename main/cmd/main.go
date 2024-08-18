package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type MarketTickersResponse struct {
	RetCode    int           `json:"retCode"`
	RetMsg     string        `json:"retMsg"`
	Result     MarketTickers `json:"result"`
	RetExtInfo struct{}      `json:"retExtInfo"`
	Time       int           `json:"time"`
}
type MarketTickers struct {
	Category string        `json:"category"`
	List     []*TickerInfo `json:"list"`
}
type TickerInfo struct {
	Symbol          string `json:"symbol"`
	LastPrice       string `json:"lastPrice"`
	Volume24h       string `json:"volume24h"`
	FundingRate     string `json:"fundingRate"`
	NextFundingTime string `json:"nextFundingTime"`
}

/*func (c MarketTickersResponse) TextOutput() string {
	p := fmt.Sprintf(
		"retCode: %s\nretMsg: %s\nresult: %s\nretExtInfo: %s\ntime: %s\n",
		c.RetCode, c.RetMsg, c.Result, c.RetExtInfo, c.Time)
	return p
}*/

func main() {
	//a := app.NewApp()
	//
	//if err := a.Run(); err != nil {
	//	panic(err)
	//}

	url := "https://api-testnet.bybit.com/v5/market/tickers?category=linear"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	/*var cResp MarketTickersResponse

	if err := json.NewDecoder(res.Body).Decode(&cResp); err != nil {
		fmt.Println("Decode error", err)
	} else {
		fmt.Println(cResp.TextOutput())
	}*/
	body, err := io.ReadAll(res.Body)
	var cResp MarketTickersResponse
	if err := json.Unmarshal(body, &cResp); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	for _, c := range cResp.Result.List {

		fmt.Println(c.Symbol, c.LastPrice, c.FundingRate, c.NextFundingTime)
	}

}
