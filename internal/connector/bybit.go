package connector

import (
	"encoding/json"
	"fmt"
	"github.com/Bochka27/ARBB/internal/models"
	"sort"
	"strconv"
	"time"
)

func BybitF() {
	url := urlBF
	response := DoGet(url)

	var dt *models.MarketTickersResponseF
	if err := json.Unmarshal(response, &dt); err != nil {
		fmt.Println(err)

		sort.SliceStable(dt.Result.List, func(i, j int) bool {
			return dt.Result.List[i].FundingRate > dt.Result.List[j].FundingRate
		})

		dtf := dt.Result.List
		fmt.Println(len(dtf))
		for _, c := range dtf {
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

func BybitS() {
	url := urlBS
	response := DoGet(url)

	var dt *models.MarketTickersResponseS
	if err := json.Unmarshal(response, &dt); err != nil {
		fmt.Println(err)
	}

	dts := dt.Result.List
	fmt.Println(len(dts))
	for _, c := range dts {
		fmt.Println(c.Symbol, c.Ask1Price, c.Bid1Price)
	}
}
