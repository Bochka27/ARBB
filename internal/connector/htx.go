package connector

import (
	"encoding/json"
	"fmt"
	"github.com/Bochka27/ARBB/internal/models"
	"strconv"
)

func HtxF() {
	url := urlHF
	response := DoGet(url)

	var dt *models.HtxFutures
	if err := json.Unmarshal(response, &dt); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(len(dt.Ticks))

	for i, dtf := range dt.Ticks {
		if len(dtf.Ask) < 1 && len(dtf.Bid) < 1 {
			continue
		}
		fmt.Println(i, dtf.ContractCode, dtf.Ask[0], dtf.Ask[1], dtf.Bid[0], dtf.Bid[1], dt.Ts)
	}
}

func HFR() {
	url := urlHFR
	response := DoGet(url)

	var dt *models.HtxFR
	if err := json.Unmarshal(response, &dt); err != nil {
		fmt.Println(err)
		return
	}

	for i, dtfr := range dt.Data {
		fr, err := strconv.ParseFloat(dtfr.FundingRate, 64)
		if err != nil {
			continue
		}
		fmt.Sprint(len(dt.Data))
		fmt.Println(i, dtfr.Symbol, fr*100)
	}
}

func HtxS() {
	url := urlHS
	response := DoGet(url)

	var dt *models.HtxSpotPrice
	if err := json.Unmarshal(response, &dt); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(dt.Data))
	for _, dts := range dt.Data {
		fmt.Println(dts.Symbol, dts.Ask, dts.Bid)
	}
}
