package connector

import (
	"encoding/json"
	"fmt"
	"github.com/Bochka27/ARBB/internal/models"
	"strconv"
)

func GateF() {
	url := urlGF
	response := DoGet(url)

	var dt models.ListF
	if err := json.Unmarshal(response, &dt); err != nil {
		fmt.Println(err)
		return
	}

	for i, dtf := range dt {
		fr, err := strconv.ParseFloat(dtf.FundingRate, 64)
		if err != nil {
			continue
		}
		fmt.Println(i, dtf.Contract, dtf.LowestAsk, dtf.HighestBid, fr*100)
	}
}

func GateS() {
	url := urlGS
	response := DoGet(url)

	var dt models.ListS
	if err := json.Unmarshal(response, &dt); err != nil {
		fmt.Println(err)
		return
	}

	for i, dts := range dt {
		fmt.Println(i, dts.CurrencyPair, dts.LowestAsk, dts.HighestBid)
	}
}
