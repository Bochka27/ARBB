package connector

import (
	"encoding/json"
	"fmt"
	"github.com/Bochka27/ARBB/internal/models"
)

func KuCoinF() {
	url := urlKCF
	response := DoGet(url)

	var dt *models.KucoinF
	if err := json.Unmarshal(response, &dt); err != nil {
		fmt.Println(err)
		return
	}

	for i, dtf := range dt.Data {
		fmt.Println(i, dtf.Symbol, dtf.BestAskPrice, dtf.BestBidPrice)
	}
}

func KuCoinFR() {
	url := urlKCFR
	response := DoGet(url)

	var dt *models.KucoinFR
	if err := json.Unmarshal(response, &dt); err != nil {
		fmt.Println(err)
	}

	for i, dtfr := range dt.Data {
		fmt.Println(i, dtfr.Symbol, dtfr.FundingFeeRate, dtfr.PredictedFundingFeeRate)
	}
}

func KuCoinS() {
	url := urlKCS
	response := DoGet(url)

	var dt *models.KucoinS
	if err := json.Unmarshal(response, &dt); err != nil {
		fmt.Println(err)
		return
	}

	for i, dts := range dt.Data.Ticker {
		fmt.Println(i, dts.Symbol, dts.Buy, dts.Sell)
	}
}
