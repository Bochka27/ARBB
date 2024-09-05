package connector

import (
	"encoding/json"
	"fmt"
	"github.com/Bochka27/ARBB/internal/models/mexc"
	"io"
	"net/http"
	"time"
)

func NewMexcF(symbol string) {
	url := "https://contract.mexc.com/api/v1/contract/funding_rate/" + symbol
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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	var cResp *mexc.FundingRate
	if err := json.Unmarshal(body, &cResp); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	data := cResp.Data

	fmt.Println(symbol, data.FundingRate, "CurrentTime = "+time.UnixMilli(data.Timestamp).Format("2006-01-02 15:04:05"), "NextSettleTime = "+time.UnixMilli(data.NextSettleTime).Format("2006-01-02 15:04:05"))
}

func NewMexcS() {
	url := "https://contract.mexc.com/api/v1/contract/detail"
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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	var cResp *mexc.DetailInformation
	if err := json.Unmarshal(body, &cResp); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	needSlice := cResp.Data
	fmt.Println(len(needSlice))
	for _, c := range needSlice {
		NewMexcF(c.Symbol)
	}
}
