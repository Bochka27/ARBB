package connector

import (
	"encoding/json"
	"fmt"
	"github.com/Bochka27/ARBB/internal/models"
	"strconv"
)

func FundingData(ticker string) any {
	url := urlOFR + ticker
	response := DoGet(url)

	var dt *models.Fund
	if err := json.Unmarshal(response, &dt); err != nil {
		fmt.Println(err)
		return nil
	}

	funding, err := strconv.ParseFloat(dt.Data[0].FundingRate, 64)
	if err != nil {
		fmt.Println("Can not parse funding rate")
		return nil
	}

	dtfr := funding * 100
	return dtfr
}

func OkxF() {
	url := urlOF
	response := DoGet(url)

	var dt *models.Symbol
	if err := json.Unmarshal(response, &dt); err != nil {
		fmt.Println(err)
	}

	dtf := dt.Data
	for _, c := range dtf {
		f := FundingData(c.InstID)
		fmt.Println(c.InstID, c.AskPx, c.BidPx, f)
	}
}

func OkxS() {
	url := urlOSL
	response := DoGet(url)

	var dt *models.Symbol
	if err := json.Unmarshal(response, &dt); err != nil {
		fmt.Println(err)
		return
	}

	dts := dt.Data
	for _, c := range dts {
		fmt.Println(c.InstID, c.AskPx, c.BidPx)
	}
}
