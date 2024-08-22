package connector

import (
	"encoding/json"
	"fmt"
	"github.com/Bochka27/ARBB/internal/models/bybit"
	"io"
	"net/http"
	"sort"
	"strconv"
	"time"
)

func NewBybitF() {
	url := "https://api.bybit.com/v5/market/tickers?category=linear"
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

	body, err := io.ReadAll(res.Body)
	var cResp *bybit.MarketTickersResponseF
	if err := json.Unmarshal(body, &cResp); err != nil {
		fmt.Println("Can not unmarshal JSON")

		sort.SliceStable(cResp.Result.List, func(i, j int) bool {
			return cResp.Result.List[i].FundingRate > cResp.Result.List[j].FundingRate
		})

		needSlice := cResp.Result.List
		fmt.Println(len(needSlice))
		for _, c := range needSlice {
			parseInt, err := strconv.ParseInt(c.NextFundingTime, 10, 64)
			if err != nil {
				return
			}
			fundingRate, err := strconv.ParseFloat(c.FundingRate, 64)
			if err != nil {
				continue
			}
			fmt.Println(c.Symbol, c.LastPrice, fundingRate*100, time.UnixMilli(parseInt).Format("2006-01-02 15:04:05"))
		}
	}
}

func NewBybitS() {
	url := "https://api.bybit.com/v5/market/tickers?category=spot"
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

	body, err := io.ReadAll(res.Body)
	var cResp *bybit.MarketTickersResponseS
	if err := json.Unmarshal(body, &cResp); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	needSlice := cResp.Result.List
	fmt.Println(len(needSlice))
	for _, c := range needSlice {
		fmt.Println(c.Symbol, c.LastPrice)
	}
}
